// AUTO-GENERATED file for Django model heatcontrol.Control

package heatcontrol

import (
	"database/sql"
	"fmt"
	"github.com/rkojedzinszky/thermo-center/models"
	"github.com/rkojedzinszky/thermo-center/models/center"
	"strings"
)

// Control mirrors model heatcontrol.Control
type Control struct {
	existsInDB bool

	id        int32
	sensor    int32
	Kp        float64
	Ki        float64
	Kd        float64
	Intabsmax sql.NullFloat64
}

// ControlQS represents a queryset for heatcontrol.Control
type ControlQS struct {
	condFragments []models.ConditionFragment
	order         []string
	forUpdate     bool
}

func (qs ControlQS) filter(c string, p interface{}) ControlQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.UnaryFragment{
			Frag:  c,
			Param: p,
		},
	)
	return qs
}

// GetId returns Control.Id
func (c *Control) GetId() int32 {
	return c.id
}

// IdEq filters for id being equal to argument
func (qs ControlQS) IdEq(v int32) ControlQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for id being not equal to argument
func (qs ControlQS) IdNe(v int32) ControlQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for id being less than argument
func (qs ControlQS) IdLt(v int32) ControlQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for id being less than or equal to argument
func (qs ControlQS) IdLe(v int32) ControlQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for id being greater than argument
func (qs ControlQS) IdGt(v int32) ControlQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for id being greater than or equal to argument
func (qs ControlQS) IdGe(v int32) ControlQS {
	return qs.filter(`"id" >=`, v)
}

// OrderById sorts result by Id in ascending order
func (qs ControlQS) OrderById() ControlQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs ControlQS) OrderByIdDesc() ControlQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetSensor returns center.Sensor
func (c *Control) GetSensor(db models.DBInterface) (*center.Sensor, error) {
	return center.SensorQS{}.IdEq(c.sensor).First(db)
}

// SetSensor sets foreign key pointer to center.Sensor
func (c *Control) SetSensor(ptr *center.Sensor) error {
	if ptr != nil {
		c.sensor = ptr.Id
	} else {
		return fmt.Errorf("Control.SetSensor: non-null field received null value")
	}

	return nil
}

// GetSensorRaw returns Control.Sensor
func (c *Control) GetSensorRaw() int32 {
	return c.sensor
}

// SensorEq filters for sensor being equal to argument
func (qs ControlQS) SensorEq(v *center.Sensor) ControlQS {
	return qs.filter(`"sensor_id" =`, v.Id)
}

type inControlsensorSensor struct {
	qs center.SensorQS
}

func (in *inControlsensorSensor) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"sensor_id" IN (` + s + `)`, p
}

func (qs ControlQS) SensorIn(oqs center.SensorQS) ControlQS {
	qs.condFragments = append(
		qs.condFragments,
		&inControlsensorSensor{
			qs: oqs,
		},
	)

	return qs
}

// OrderBySensor sorts result by Sensor in ascending order
func (qs ControlQS) OrderBySensor() ControlQS {
	qs.order = append(qs.order, `"sensor_id"`)

	return qs
}

// OrderBySensorDesc sorts result by Sensor in descending order
func (qs ControlQS) OrderBySensorDesc() ControlQS {
	qs.order = append(qs.order, `"sensor_id" DESC`)

	return qs
}

// KpEq filters for Kp being equal to argument
func (qs ControlQS) KpEq(v float64) ControlQS {
	return qs.filter(`"kp" =`, v)
}

// KpNe filters for Kp being not equal to argument
func (qs ControlQS) KpNe(v float64) ControlQS {
	return qs.filter(`"kp" <>`, v)
}

// KpLt filters for Kp being less than argument
func (qs ControlQS) KpLt(v float64) ControlQS {
	return qs.filter(`"kp" <`, v)
}

// KpLe filters for Kp being less than or equal to argument
func (qs ControlQS) KpLe(v float64) ControlQS {
	return qs.filter(`"kp" <=`, v)
}

// KpGt filters for Kp being greater than argument
func (qs ControlQS) KpGt(v float64) ControlQS {
	return qs.filter(`"kp" >`, v)
}

// KpGe filters for Kp being greater than or equal to argument
func (qs ControlQS) KpGe(v float64) ControlQS {
	return qs.filter(`"kp" >=`, v)
}

// OrderByKp sorts result by Kp in ascending order
func (qs ControlQS) OrderByKp() ControlQS {
	qs.order = append(qs.order, `"kp"`)

	return qs
}

// OrderByKpDesc sorts result by Kp in descending order
func (qs ControlQS) OrderByKpDesc() ControlQS {
	qs.order = append(qs.order, `"kp" DESC`)

	return qs
}

// KiEq filters for Ki being equal to argument
func (qs ControlQS) KiEq(v float64) ControlQS {
	return qs.filter(`"ki" =`, v)
}

// KiNe filters for Ki being not equal to argument
func (qs ControlQS) KiNe(v float64) ControlQS {
	return qs.filter(`"ki" <>`, v)
}

// KiLt filters for Ki being less than argument
func (qs ControlQS) KiLt(v float64) ControlQS {
	return qs.filter(`"ki" <`, v)
}

// KiLe filters for Ki being less than or equal to argument
func (qs ControlQS) KiLe(v float64) ControlQS {
	return qs.filter(`"ki" <=`, v)
}

// KiGt filters for Ki being greater than argument
func (qs ControlQS) KiGt(v float64) ControlQS {
	return qs.filter(`"ki" >`, v)
}

// KiGe filters for Ki being greater than or equal to argument
func (qs ControlQS) KiGe(v float64) ControlQS {
	return qs.filter(`"ki" >=`, v)
}

// OrderByKi sorts result by Ki in ascending order
func (qs ControlQS) OrderByKi() ControlQS {
	qs.order = append(qs.order, `"ki"`)

	return qs
}

// OrderByKiDesc sorts result by Ki in descending order
func (qs ControlQS) OrderByKiDesc() ControlQS {
	qs.order = append(qs.order, `"ki" DESC`)

	return qs
}

// KdEq filters for Kd being equal to argument
func (qs ControlQS) KdEq(v float64) ControlQS {
	return qs.filter(`"kd" =`, v)
}

// KdNe filters for Kd being not equal to argument
func (qs ControlQS) KdNe(v float64) ControlQS {
	return qs.filter(`"kd" <>`, v)
}

// KdLt filters for Kd being less than argument
func (qs ControlQS) KdLt(v float64) ControlQS {
	return qs.filter(`"kd" <`, v)
}

// KdLe filters for Kd being less than or equal to argument
func (qs ControlQS) KdLe(v float64) ControlQS {
	return qs.filter(`"kd" <=`, v)
}

// KdGt filters for Kd being greater than argument
func (qs ControlQS) KdGt(v float64) ControlQS {
	return qs.filter(`"kd" >`, v)
}

// KdGe filters for Kd being greater than or equal to argument
func (qs ControlQS) KdGe(v float64) ControlQS {
	return qs.filter(`"kd" >=`, v)
}

// OrderByKd sorts result by Kd in ascending order
func (qs ControlQS) OrderByKd() ControlQS {
	qs.order = append(qs.order, `"kd"`)

	return qs
}

// OrderByKdDesc sorts result by Kd in descending order
func (qs ControlQS) OrderByKdDesc() ControlQS {
	qs.order = append(qs.order, `"kd" DESC`)

	return qs
}

// IntabsmaxIsNull filters for Intabsmax being null
func (qs ControlQS) IntabsmaxIsNull() ControlQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"intabsmax" IS NULL`,
		},
	)
	return qs
}

// IntabsmaxIsNotNull filters for Intabsmax being not null
func (qs ControlQS) IntabsmaxIsNotNull() ControlQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"intabsmax" IS NOT NULL`,
		},
	)
	return qs
}

// IntabsmaxEq filters for Intabsmax being equal to argument
func (qs ControlQS) IntabsmaxEq(v float64) ControlQS {
	return qs.filter(`"intabsmax" =`, v)
}

// IntabsmaxNe filters for Intabsmax being not equal to argument
func (qs ControlQS) IntabsmaxNe(v float64) ControlQS {
	return qs.filter(`"intabsmax" <>`, v)
}

// IntabsmaxLt filters for Intabsmax being less than argument
func (qs ControlQS) IntabsmaxLt(v float64) ControlQS {
	return qs.filter(`"intabsmax" <`, v)
}

// IntabsmaxLe filters for Intabsmax being less than or equal to argument
func (qs ControlQS) IntabsmaxLe(v float64) ControlQS {
	return qs.filter(`"intabsmax" <=`, v)
}

// IntabsmaxGt filters for Intabsmax being greater than argument
func (qs ControlQS) IntabsmaxGt(v float64) ControlQS {
	return qs.filter(`"intabsmax" >`, v)
}

// IntabsmaxGe filters for Intabsmax being greater than or equal to argument
func (qs ControlQS) IntabsmaxGe(v float64) ControlQS {
	return qs.filter(`"intabsmax" >=`, v)
}

// OrderByIntabsmax sorts result by Intabsmax in ascending order
func (qs ControlQS) OrderByIntabsmax() ControlQS {
	qs.order = append(qs.order, `"intabsmax"`)

	return qs
}

// OrderByIntabsmaxDesc sorts result by Intabsmax in descending order
func (qs ControlQS) OrderByIntabsmaxDesc() ControlQS {
	qs.order = append(qs.order, `"intabsmax" DESC`)

	return qs
}

func (qs ControlQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
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
func (qs ControlQS) ForUpdate() ControlQS {
	qs.forUpdate = true

	return qs
}

func (qs ControlQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs ControlQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs ControlQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "sensor_id", "kp", "ki", "kd", "intabsmax" FROM "heatcontrol_control"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs ControlQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_control"` + s, p
}

// All returns all rows matching queryset filters
func (qs ControlQS) All(db models.DBInterface) ([]*Control, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Control
	for rows.Next() {
		obj := Control{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.sensor, &obj.Kp, &obj.Ki, &obj.Kd, &obj.Intabsmax); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs ControlQS) First(db models.DBInterface) (*Control, error) {
	s, p := qs.queryFull()

	row := db.QueryRow(s, p...)

	obj := Control{existsInDB: true}
	err := row.Scan(&obj.id, &obj.sensor, &obj.Kp, &obj.Ki, &obj.Kd, &obj.Intabsmax)
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
func (c *Control) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "heatcontrol_control" ("sensor_id", "kp", "ki", "kd", "intabsmax") VALUES ($1, $2, $3, $4, $5) RETURNING "id"`, c.sensor, c.Kp, c.Ki, c.Kd, c.Intabsmax)

	if err := row.Scan(&c.id); err != nil {
		return err
	}

	c.existsInDB = true

	return nil
}

// update operation
func (c *Control) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "heatcontrol_control" SET "sensor_id" = $1, "kp" = $2, "ki" = $3, "kd" = $4, "intabsmax" = $5 WHERE "id" = $6`, c.sensor, c.Kp, c.Ki, c.Kd, c.Intabsmax, c.id)

	return err
}

// Save inserts or updates record
func (c *Control) Save(db models.DBInterface) error {
	if c.existsInDB {
		return c.update(db)
	}

	return c.insert(db)
}

// Delete removes row from database
func (c *Control) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "heatcontrol_control" WHERE "id" = $1`, c.id)

	c.existsInDB = false

	return err
}
