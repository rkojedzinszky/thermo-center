package aggregator

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rkojedzinszky/thermo-center/v5/aggregator/sensorvalue"
	"github.com/rkojedzinszky/thermo-center/v5/models/center"
	"github.com/rkojedzinszky/thermo-center/v5/models/heatcontrol"
)

const defaultCarbonMetricPathTemplate = `sensor.{{ printf "%02x" .SensorID }}.{{ .Metric }}`

// aggregator serves the aggregator API
type aggregator struct {
	db                *pgxpool.Pool
	location          *time.Location
	mc                *memcache.Client
	mqtt              *MqttClient
	graphite          *GraphiteSender
	updateProbability float64
	processID         []byte

	localLock    []time.Time
	localLockMtx sync.Mutex

	UnimplementedAggregatorServer
}

// NewAggregator instantiates a new aggregator
func NewAggregator(db *pgxpool.Pool, location *time.Location) (AggregatorServer, error) {
	a := &aggregator{
		db:       db,
		location: location,
	}

	// Init components
	mcServerList := new(memcache.ServerList)
	err := mcServerList.SetServers(fmt.Sprintf("%s:%s", getenv("MEMCACHED_HOST", "memcached"), getenv("MEMCACHED_PORT", "11211")))
	if err != nil {
		return nil, err
	}

	a.mc = memcache.NewFromSelector(mcServerList)

	a.mqtt = NewMqttClient(getenv("MQTT_HOST", "mqtt"), getenvInt("MQTT_PORT", 1883))
	go a.mqtt.Run()

	if graphiteHost := getenv("CARBON_LINE_RECEIVER_HOST", ""); graphiteHost != "" {
		a.graphite = NewGraphiteSender(
			graphiteHost,
			getenvInt("CARBON_LINE_RECEIVER_PORT", 2003),
			getenv("CARBON_LINE_RECEIVER_METRIC_PATH_TEMPLATE", defaultCarbonMetricPathTemplate),
		)
		if a.graphite != nil {
			go a.graphite.Run()
		}
	}

	a.updateProbability = getenvFloat64("SENSOR_DB_UPDATE_PROBABILITY", 0.01)

	hostname, _ := os.Hostname()
	a.processID = []byte(fmt.Sprintf("%s-%d", hostname, os.Getpid()))

	a.localLock = make([]time.Time, 128)

	return a, nil
}

type sensorStat struct {
	Sensor *center.Sensor
	Stat   map[string]interface{}
}

const (
	sensorLockKey            = "tc/sensor/%02x/lock"
	sensorLockTimeout        = 2
	sensorCacheKeyTemplate   = "sensor.%02x"
	sensorControlKeyTemplate = "tc/sensor/%02x/control"

	sensorCacheFlag    = 16
	sensorCacheTimeout = 120
)

func (a *aggregator) now() time.Time {
	return time.Now().In(a.location)
}

func (a *aggregator) lockLocal(id uint8) bool {
	now := time.Now()
	a.localLockMtx.Lock()
	defer a.localLockMtx.Unlock()

	if now.Before(a.localLock[id]) {
		return false
	}

	a.localLock[id] = now.Add(sensorLockTimeout * time.Second)

	return true
}

func (a *aggregator) lockSensor(id uint8) (bool, error) {
	id &= 0x7f

	if a.lockLocal(id) == false {
		return false, nil
	}

	key := fmt.Sprintf(sensorLockKey, id)

	err := a.mc.Add(&memcache.Item{
		Key:        key,
		Expiration: sensorLockTimeout,
		Value:      a.processID,
	})

	if err == nil {
		return true, nil
	}

	if err == memcache.ErrNotStored {
		if _, err := a.mc.Get(key); err == nil {
			return false, nil
		}

		return false, fmt.Errorf("Locking sensor %02x failed, however could not determine owner", id)
	}

	return false, err
}

func sensorControlKey(s *center.Sensor) string {
	return fmt.Sprintf(sensorControlKeyTemplate, s.ID)
}

// update sequence values from cache
func (a *aggregator) validateSensorPacket(ctx context.Context, s *center.Sensor, p *SensorPacket) (float64, bool) {
	now := float64(time.Now().UnixNano()) / 1e9

	var avg float64
	var valid bool
	var save bool

	if s.LastTsf.Valid {
		interval := now - s.LastTsf.Float64

		if s.LastSeq.Valid {
			diff := (p.Seq - uint32(s.LastSeq.Int32)) & 0x7fffffff

			if diff != 0 {
				avg = interval / float64(diff)
				if 25 <= avg && avg <= 35 {
					valid = true
				}
			}
		} else {
			if interval <= 35 {
				valid = true
				save = true
			}
		}
	}

	if valid {
		s.LastSeq.Int32 = int32(p.Seq)
		s.LastSeq.Valid = true
		s.LastTsf.Float64 = now
		s.LastTsf.Valid = true

		if save || rand.Float64() < a.updateProbability {
			center.SensorQS{}.IDEq(s.ID).Update().SetLastSeq(s.LastSeq).SetLastTsf(s.LastTsf).Exec(ctx, a.db)
		}
	}

	return avg, valid
}

func (a *aggregator) loadPidControl(s *center.Sensor, pid *pidController) error {
	cache, err := a.mc.Get(sensorControlKey(s))
	switch err {
	case memcache.ErrCacheMiss:
	case nil:
		json.Unmarshal(cache.Value, pid)
	default:
		return err
	}

	return nil
}

func (a *aggregator) savePidControl(s *center.Sensor, pid *pidController) error {
	data, err := json.Marshal(pid)
	if err != nil {
		return err
	}

	if err := a.mc.Set(&memcache.Item{
		Key:        sensorControlKey(s),
		Value:      data,
		Expiration: sensorCacheTimeout,
	}); err != nil {
		return err
	}

	return nil
}

func sensorCacheKey(s *center.Sensor) string {
	return fmt.Sprintf(sensorCacheKeyTemplate, s.ID)
}

// save metric values in cache
func (a *aggregator) saveCache(s *center.Sensor, cache map[string]interface{}) error {
	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	if err := a.mc.Set(&memcache.Item{
		Key:        sensorCacheKey(s),
		Value:      data,
		Flags:      16,
		Expiration: sensorCacheTimeout,
	}); err != nil {
		return err
	}

	return nil
}

func (a *aggregator) getTargetTemp(ctx context.Context, c *heatcontrol.Control) (*float64, error) {
	now := a.now()

	var so *heatcontrol.Scheduledoverride
	var ipe *heatcontrol.Instantprofileentry
	var err error

	so, err = c.Scheduledoverride().StartLe(now).EndGt(now).OrderByIDDesc().First(ctx, a.db)

	if err != nil {
		return nil, err
	}

	if so != nil {
		return &so.TargetTemp, nil
	}

	ipe, err = c.Instantprofileentry().ActiveEq(true).First(ctx, a.db)

	if err != nil {
		return nil, err
	}

	if ipe != nil {
		if ipe.TargetTemp.Valid {
			return &ipe.TargetTemp.Float64, nil
		}

		return nil, nil
	}

	// Reading profile is hardcoded here
	var targetTemp *float64
	row := a.db.QueryRow(ctx, `
		select "heatcontrol_profile"."target_temp" from "heatcontrol_profile"
		where "control_id" = $1 and daytype_id =
			(select "daytype_id" from heatcontrol_calendar where day = $2::date)
		and start <= $3::time order by "heatcontrol_profile"."start" desc`,
		c.GetID(), now, now,
	)

	err = row.Scan(&targetTemp)
	if err == nil {
		if targetTemp != nil {
			return targetTemp, nil
		}
	}

	if err != pgx.ErrNoRows {
		return nil, err
	}

	return nil, nil
}

// FeedSensorPacket processes a packet received by a receiver
func (a *aggregator) FeedSensorPacket(ctx context.Context, p *SensorPacket) (*FeedResponse, error) {

	ok, err := a.lockSensor(uint8(p.Id))

	if !ok {
		if err != nil {
			fmt.Println(err)
		}

		return &FeedResponse{
			Processed: false,
		}, nil
	}

	sensor, err := center.SensorQS{}.IDEq(int32(p.Id)).First(ctx, a.db)
	if err != nil || sensor == nil {
		return &FeedResponse{
			Processed: false,
		}, err
	}

	cache := make(map[string]interface{})

	intvl, valid := a.validateSensorPacket(ctx, sensor, p)

	cache["valid"] = valid
	if valid {
		cache["intvl"] = intvl

		metrics := sensorvalue.ParseBytes(p.Raw)
		metrics = append(metrics, sensorvalue.NewLQI(p.Lqi), sensorvalue.NewRSSI(p.Rssi))

		for _, m := range metrics {
			cache[m.Metric()] = m.Value()
		}

		control, err := heatcontrol.ControlQS{}.SensorEq(sensor).First(ctx, a.db)
		if err != nil {
			return &FeedResponse{
				Processed: false,
			}, err
		}

		if control != nil {
			var pid pidController

			if err = a.loadPidControl(sensor, &pid); err != nil {
				return &FeedResponse{
					Processed: false,
				}, err
			}

			targetTemp, err := a.getTargetTemp(ctx, control)
			if err != nil {
				return &FeedResponse{
					Processed: false,
				}, err
			}

			if targetTemp != nil {
				e := *targetTemp - cache["Temperature"].(float64)

				var intabsmax *float64
				if control.Intabsmax.Valid {
					intabsmax = &control.Intabsmax.Float64
				}
				pid.feed(e, intabsmax)

				pcv := pid.value(control.Kp, control.Ki, control.Kd)

				cache["pidcontrol"] = pcv

				a.savePidControl(sensor, &pid)
			}
		}

	}

	// last_tsf is available through the API, so set it in cache
	if valid {
		cache["last_tsf"] = sensor.LastTsf.Float64
	} else {
		cache["last_tsf"] = nil
	}

	a.saveCache(sensor, cache)

	ss := sensorStat{
		Sensor: sensor,
		Stat:   cache,
	}

	a.mqtt.Push(ss)

	if valid {
		if a.graphite != nil {
			a.graphite.Push(ss)
		}
	}

	return &FeedResponse{
		Processed: true,
	}, nil
}
