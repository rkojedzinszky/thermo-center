package configurator

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/rkojedzinszky/thermo-center/models/center"
)

type configurator struct {
	db       *sql.DB
	location *time.Location
}

// NewConfigurator returns a new configurator
func NewConfigurator(db *sql.DB, location *time.Location) ConfiguratorServer {
	return &configurator{
		db:       db,
		location: location,
	}
}

func (c *configurator) GetRadioCfg(ctx context.Context, r *RadioCfgRequest) (*RadioCfgResponse, error) {
	if r.Cluster != 1 {
		return nil, fmt.Errorf("Invalid cluster ID received")
	}

	config, err := center.RfconfigQS{}.IDEq(int32(r.Cluster)).First(c.db)
	if err != nil {
		return nil, err
	}

	profile, err := config.GetRfProfile(c.db)
	if err != nil {
		return nil, err
	}

	radioConfig, _ := hex.DecodeString(profile.Confregs)
	radioConfig = append(radioConfig, byte(0x0a), byte(config.RfChannel))
	aesKey, _ := hex.DecodeString(config.AesKey)

	return &RadioCfgResponse{
		Network:     uint32(config.NetworkId),
		RadioConfig: radioConfig,
		AesKey:      aesKey,
	}, nil
}

func (c *configurator) TaskAcquire(ctx context.Context, t *Task) (*TaskDetails, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	task, err := center.ConfiguresensortaskQS{}.ForUpdate().IDEq(int32(t.TaskId)).StartedIsNull().First(tx)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, fmt.Errorf("TaskAcquire: task with id %d not found", t.TaskId)
	}

	task.Started.Time = time.Now().In(c.location)
	task.Started.Valid = true

	err = task.Save(tx)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	cfg, err := c.GetRadioCfg(ctx, &RadioCfgRequest{Cluster: 1})
	if err != nil {
		return nil, err
	}

	return &TaskDetails{
		TaskId:   uint32(task.GetID()),
		SensorId: uint32(task.GetSensorRaw()),
		Config:   cfg,
	}, nil
}

func (c *configurator) TaskDiscoveryReceived(ctx context.Context, t *Task) (*TaskUpdateResponse, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	task, err := center.ConfiguresensortaskQS{}.ForUpdate().IDEq(int32(t.TaskId)).StartedIsNotNull().FinishedIsNull().First(tx)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, fmt.Errorf("TaskDiscoveryReceived: task with id %d not found", t.TaskId)
	}

	task.LastDiscovery.Time = time.Now().In(c.location)
	task.LastDiscovery.Valid = true
	if task.FirstDiscovery.Valid == false {
		task.FirstDiscovery = task.LastDiscovery
	}

	err = task.Save(tx)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &TaskUpdateResponse{
		Success: true,
	}, nil
}

func (c *configurator) TaskFinished(ctx context.Context, t *TaskFinishedRequest) (*TaskUpdateResponse, error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	task, err := center.ConfiguresensortaskQS{}.ForUpdate().IDEq(int32(t.TaskId)).StartedIsNotNull().FinishedIsNull().First(tx)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, fmt.Errorf("TaskFinished: task with id %d not found", t.TaskId)
	}

	now := time.Now().In(c.location)

	task.Finished.Time = now
	task.Finished.Valid = true

	task.Error.String = t.Error
	task.Error.Valid = true

	err = task.Save(tx)
	if err != nil {
		return nil, err
	}

	if t.Error == "" {
		sensor, err := task.GetSensor(tx)
		if err != nil {
			return nil, err
		}

		sensor.LastSeq.Valid = false

		sensor.LastTsf.Float64 = float64(now.Unix())
		sensor.LastTsf.Valid = true

		if err = sensor.Save(tx); err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &TaskUpdateResponse{
		Success: true,
	}, nil
}
