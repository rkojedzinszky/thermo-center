/*
  AUTO-GENERATED file for Django model center.SensorResync

  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

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

// GetID returns Sensorresync.ID
func (s *Sensorresync) GetID() int32 {
	return s.id
}

// IDEq filters for id being equal to argument
func (qs SensorresyncQS) IDEq(v int32) SensorresyncQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs SensorresyncQS) IDNe(v int32) SensorresyncQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs SensorresyncQS) IDLt(v int32) SensorresyncQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs SensorresyncQS) IDLe(v int32) SensorresyncQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs SensorresyncQS) IDGt(v int32) SensorresyncQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs SensorresyncQS) IDGe(v int32) SensorresyncQS {
	return qs.filter(`"id" >=`, v)
}

type inSensorresyncid struct {
	values []interface{}
}

func (in *inSensorresyncid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorresyncQS) IDIn(values []int32) SensorresyncQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inSensorresyncid{
			values: vals,
		},
	)

	return qs
}

type notinSensorresyncid struct {
	values []interface{}
}

func (in *notinSensorresyncid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorresyncQS) IDNotIn(values []int32) SensorresyncQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinSensorresyncid{
			values: vals,
		},
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs SensorresyncQS) OrderByID() SensorresyncQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs SensorresyncQS) OrderByIDDesc() SensorresyncQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetSensor returns Sensor
func (s *Sensorresync) GetSensor(db models.DBInterface) (*Sensor, error) {
	return SensorQS{}.IDEq(s.sensor).First(db)
}

// SetSensor sets foreign key pointer to Sensor
func (s *Sensorresync) SetSensor(ptr *Sensor) error {
	if ptr != nil {
		s.sensor = ptr.ID
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
	return qs.filter(`"sensor_id" =`, v.ID)
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

type inSensorresyncTs struct {
	values []interface{}
}

func (in *inSensorresyncTs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"ts" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorresyncQS) TsIn(values []time.Time) SensorresyncQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inSensorresyncTs{
			values: vals,
		},
	)

	return qs
}

type notinSensorresyncTs struct {
	values []interface{}
}

func (in *notinSensorresyncTs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"ts" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorresyncQS) TsNotIn(values []time.Time) SensorresyncQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinSensorresyncTs{
			values: vals,
		},
	)

	return qs
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

	s += " LIMIT 1"

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

// Delete deletes rows matching queryset filters
func (qs SensorresyncQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "center_sensorresync"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs SensorresyncQS) Update() SensorresyncUpdateQS {
	return SensorresyncUpdateQS{condFragments: qs.condFragments}
}

// SensorresyncUpdateQS represents an updated queryset for center.SensorResync
type SensorresyncUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs SensorresyncUpdateQS) update(c string, v interface{}) SensorresyncUpdateQS {
	var frag models.ConditionFragment

	if v == nil {
		frag = &models.ConstantFragment{
			Constant: c + " = NULL",
		}
	} else {
		frag = &models.UnaryFragment{
			Frag:  c + " =",
			Param: v,
		}
	}

	uqs.updates = append(uqs.updates, frag)

	return uqs
}

// SetID sets ID to the given value
func (uqs SensorresyncUpdateQS) SetID(v int32) SensorresyncUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetSensor sets foreign key pointer to Sensor
func (uqs SensorresyncUpdateQS) SetSensor(ptr *Sensor) SensorresyncUpdateQS {
	if ptr != nil {
		return uqs.update(`"sensor_id"`, ptr.ID)
	}

	return uqs.update(`"sensor_id"`, nil)
} // SetTs sets Ts to the given value
func (uqs SensorresyncUpdateQS) SetTs(v time.Time) SensorresyncUpdateQS {
	return uqs.update(`"ts"`, v)
}

// Exec executes the update operation
func (uqs SensorresyncUpdateQS) Exec(db models.DBInterface) (int64, error) {
	if len(uqs.updates) == 0 {
		return 0, nil
	}

	c := &models.PositionalCounter{}

	var params []interface{}

	var sets []string
	for _, set := range uqs.updates {
		s, p := set.GetConditionFragment(c)

		sets = append(sets, s)
		params = append(params, p...)
	}

	ws, wp := SensorresyncQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "center_sensorresync" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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
