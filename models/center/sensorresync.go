// AUTO-GENERATED file for Django model center.SensorResync

package center

import (
	"database/sql"
	"fmt"
	"github.com/rkojedzinszky/thermo-center/models"
	"strings"
	"time"
)

// Sensorresync mirrors model center.SensorResync
type Sensorresync struct {
	existsInDB bool

	id     int32
	sensor int32
	Ts     time.Time
}

// SensorresyncQS represents a queryset for center.SensorResync
type SensorresyncQS struct {
	condFragments []models.ConditionFragment
	order         []string
	forUpdate     bool
}

func (qs SensorresyncQS) filter(c string, p interface{}) SensorresyncQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.UnaryFragment{
			Frag:  c,
			Param: p,
		},
	)
	return qs
}

// GetId returns Sensorresync.Id
func (s *Sensorresync) GetId() int32 {
	return s.id
}

// IdEq filters for id being equal to argument
func (qs SensorresyncQS) IdEq(v int32) SensorresyncQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for id being not equal to argument
func (qs SensorresyncQS) IdNe(v int32) SensorresyncQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for id being less than argument
func (qs SensorresyncQS) IdLt(v int32) SensorresyncQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for id being less than or equal to argument
func (qs SensorresyncQS) IdLe(v int32) SensorresyncQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for id being greater than argument
func (qs SensorresyncQS) IdGt(v int32) SensorresyncQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for id being greater than or equal to argument
func (qs SensorresyncQS) IdGe(v int32) SensorresyncQS {
	return qs.filter(`"id" >=`, v)
}

// OrderById sorts result by Id in ascending order
func (qs SensorresyncQS) OrderById() SensorresyncQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs SensorresyncQS) OrderByIdDesc() SensorresyncQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetSensor returns Sensor
func (s *Sensorresync) GetSensor(db models.DBInterface) (*Sensor, error) {
	return SensorQS{}.IdEq(s.sensor).First(db)
}

// SetSensor sets foreign key pointer to Sensor
func (s *Sensorresync) SetSensor(ptr *Sensor) error {
	if ptr != nil {
		s.sensor = ptr.Id
	} else {
		return fmt.Errorf("Sensorresync.SetSensor: non-null field received null value")
	}

	return nil
}

// GetSensorRaw returns Sensorresync.Sensor
func (s *Sensorresync) GetSensorRaw() int32 {
	return s.sensor
}

// SensorEq filters for sensor being equal to argument
func (qs SensorresyncQS) SensorEq(v *Sensor) SensorresyncQS {
	return qs.filter(`"sensor_id" =`, v.Id)
}

type inSensorresyncsensorSensor struct {
	qs SensorQS
}

func (in *inSensorresyncsensorSensor) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"sensor_id" IN (` + s + `)`, p
}

func (qs SensorresyncQS) SensorIn(oqs SensorQS) SensorresyncQS {
	qs.condFragments = append(
		qs.condFragments,
		&inSensorresyncsensorSensor{
			qs: oqs,
		},
	)

	return qs
}

// OrderBySensor sorts result by Sensor in ascending order
func (qs SensorresyncQS) OrderBySensor() SensorresyncQS {
	qs.order = append(qs.order, `"sensor_id"`)

	return qs
}

// OrderBySensorDesc sorts result by Sensor in descending order
func (qs SensorresyncQS) OrderBySensorDesc() SensorresyncQS {
	qs.order = append(qs.order, `"sensor_id" DESC`)

	return qs
}

// TsEq filters for Ts being equal to argument
func (qs SensorresyncQS) TsEq(v time.Time) SensorresyncQS {
	return qs.filter(`"ts" =`, v)
}

// TsNe filters for Ts being not equal to argument
func (qs SensorresyncQS) TsNe(v time.Time) SensorresyncQS {
	return qs.filter(`"ts" <>`, v)
}

// TsLt filters for Ts being less than argument
func (qs SensorresyncQS) TsLt(v time.Time) SensorresyncQS {
	return qs.filter(`"ts" <`, v)
}

// TsLe filters for Ts being less than or equal to argument
func (qs SensorresyncQS) TsLe(v time.Time) SensorresyncQS {
	return qs.filter(`"ts" <=`, v)
}

// TsGt filters for Ts being greater than argument
func (qs SensorresyncQS) TsGt(v time.Time) SensorresyncQS {
	return qs.filter(`"ts" >`, v)
}

// TsGe filters for Ts being greater than or equal to argument
func (qs SensorresyncQS) TsGe(v time.Time) SensorresyncQS {
	return qs.filter(`"ts" >=`, v)
}

// OrderByTs sorts result by Ts in ascending order
func (qs SensorresyncQS) OrderByTs() SensorresyncQS {
	qs.order = append(qs.order, `"ts"`)

	return qs
}

// OrderByTsDesc sorts result by Ts in descending order
func (qs SensorresyncQS) OrderByTsDesc() SensorresyncQS {
	qs.order = append(qs.order, `"ts" DESC`)

	return qs
}

func (qs SensorresyncQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	var conds []string
	var condp []interface{}

	for _, cond := range qs.condFragments {
		s, p := cond.GetConditionFragment(c)

		conds = append(conds, s)
		condp = append(condp, p...)
	}

	return strings.Join(conds, " AND "), condp
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs SensorresyncQS) ForUpdate() SensorresyncQS {
	qs.forUpdate = true

	return qs
}

func (qs SensorresyncQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs SensorresyncQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs SensorresyncQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "sensor_id", "ts" FROM "center_sensorresync"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs SensorresyncQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "center_sensorresync"` + s, p
}

// All returns all rows matching queryset filters
func (qs SensorresyncQS) All(db models.DBInterface) ([]*Sensorresync, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Sensorresync
	for rows.Next() {
		obj := Sensorresync{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.sensor, &obj.Ts); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs SensorresyncQS) First(db models.DBInterface) (*Sensorresync, error) {
	s, p := qs.queryFull()

	row := db.QueryRow(s, p...)

	obj := Sensorresync{existsInDB: true}
	err := row.Scan(&obj.id, &obj.sensor, &obj.Ts)
	switch err {
	case nil:
		return &obj, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

}

// insert operation
func (s *Sensorresync) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "center_sensorresync" ("sensor_id", "ts") VALUES ($1, $2) RETURNING "id"`, s.sensor, s.Ts)

	if err := row.Scan(&s.id); err != nil {
		return err
	}

	s.existsInDB = true

	return nil
}

// update operation
func (s *Sensorresync) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "center_sensorresync" SET "sensor_id" = $1, "ts" = $2 WHERE "id" = $3`, s.sensor, s.Ts, s.id)

	return err
}

// Save inserts or updates record
func (s *Sensorresync) Save(db models.DBInterface) error {
	if s.existsInDB {
		return s.update(db)
	}

	return s.insert(db)
}

// Delete removes row from database
func (s *Sensorresync) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "center_sensorresync" WHERE "id" = $1`, s.id)

	s.existsInDB = false

	return err
}
