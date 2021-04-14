/*
  AUTO-GENERATED file for Django model heatcontrol.Control

  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

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
	condFragments models.AndFragment
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

// Or combines given expressions with OR operator
func (qs ControlQS) Or(exprs ...ControlQS) ControlQS {
	var o models.OrFragment

	for _, expr := range exprs {
		o = append(o, expr.condFragments)
	}

	qs.condFragments = append(
		qs.condFragments,
		o,
	)

	return qs
}

// GetID returns Control.ID
func (c *Control) GetID() int32 {
	return c.id
}

// IDEq filters for id being equal to argument
func (qs ControlQS) IDEq(v int32) ControlQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs ControlQS) IDNe(v int32) ControlQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs ControlQS) IDLt(v int32) ControlQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs ControlQS) IDLe(v int32) ControlQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs ControlQS) IDGt(v int32) ControlQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs ControlQS) IDGe(v int32) ControlQS {
	return qs.filter(`"id" >=`, v)
}

type inControlid struct {
	values []interface{}
}

func (in *inControlid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) IDIn(values []int32) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inControlid{
			values: vals,
		},
	)

	return qs
}

type notinControlid struct {
	values []interface{}
}

func (in *notinControlid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) IDNotIn(values []int32) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinControlid{
			values: vals,
		},
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs ControlQS) OrderByID() ControlQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs ControlQS) OrderByIDDesc() ControlQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetSensor returns center.Sensor
func (c *Control) GetSensor(db models.DBInterface) (*center.Sensor, error) {
	return center.SensorQS{}.IDEq(c.sensor).First(db)
}

// SetSensor sets foreign key pointer to center.Sensor
func (c *Control) SetSensor(ptr *center.Sensor) error {
	if ptr != nil {
		c.sensor = ptr.ID
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
	return qs.filter(`"sensor_id" =`, v.ID)
}

// SensorRawEq filters for sensor being equal to raw argument
func (qs ControlQS) SensorRawEq(v int32) ControlQS {
	return qs.filter(`"sensor_id" =`, v)
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

type inControlKp struct {
	values []interface{}
}

func (in *inControlKp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"kp" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) KpIn(values []float64) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inControlKp{
			values: vals,
		},
	)

	return qs
}

type notinControlKp struct {
	values []interface{}
}

func (in *notinControlKp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"kp" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) KpNotIn(values []float64) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinControlKp{
			values: vals,
		},
	)

	return qs
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

type inControlKi struct {
	values []interface{}
}

func (in *inControlKi) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"ki" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) KiIn(values []float64) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inControlKi{
			values: vals,
		},
	)

	return qs
}

type notinControlKi struct {
	values []interface{}
}

func (in *notinControlKi) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"ki" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) KiNotIn(values []float64) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinControlKi{
			values: vals,
		},
	)

	return qs
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

type inControlKd struct {
	values []interface{}
}

func (in *inControlKd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"kd" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) KdIn(values []float64) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inControlKd{
			values: vals,
		},
	)

	return qs
}

type notinControlKd struct {
	values []interface{}
}

func (in *notinControlKd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"kd" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) KdNotIn(values []float64) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinControlKd{
			values: vals,
		},
	)

	return qs
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

type inControlIntabsmax struct {
	values []interface{}
}

func (in *inControlIntabsmax) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"intabsmax" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) IntabsmaxIn(values []float64) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inControlIntabsmax{
			values: vals,
		},
	)

	return qs
}

type notinControlIntabsmax struct {
	values []interface{}
}

func (in *notinControlIntabsmax) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"intabsmax" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ControlQS) IntabsmaxNotIn(values []float64) ControlQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinControlIntabsmax{
			values: vals,
		},
	)

	return qs
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

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs ControlQS) ForUpdate() ControlQS {
	qs.forUpdate = true

	return qs
}

func (qs ControlQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

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

	s += " LIMIT 1"

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

// Delete deletes rows matching queryset filters
func (qs ControlQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_control"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs ControlQS) Update() ControlUpdateQS {
	return ControlUpdateQS{condFragments: qs.condFragments}
}

// ControlUpdateQS represents an updated queryset for heatcontrol.Control
type ControlUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs ControlUpdateQS) update(c string, v interface{}) ControlUpdateQS {
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
func (uqs ControlUpdateQS) SetID(v int32) ControlUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetSensor sets foreign key pointer to center.Sensor
func (uqs ControlUpdateQS) SetSensor(ptr *center.Sensor) ControlUpdateQS {
	if ptr != nil {
		return uqs.update(`"sensor_id"`, ptr.ID)
	}

	return uqs.update(`"sensor_id"`, nil)
} // SetKp sets Kp to the given value
func (uqs ControlUpdateQS) SetKp(v float64) ControlUpdateQS {
	return uqs.update(`"kp"`, v)
}

// SetKi sets Ki to the given value
func (uqs ControlUpdateQS) SetKi(v float64) ControlUpdateQS {
	return uqs.update(`"ki"`, v)
}

// SetKd sets Kd to the given value
func (uqs ControlUpdateQS) SetKd(v float64) ControlUpdateQS {
	return uqs.update(`"kd"`, v)
}

// SetIntabsmax sets Intabsmax to the given value
func (uqs ControlUpdateQS) SetIntabsmax(v sql.NullFloat64) ControlUpdateQS {
	return uqs.update(`"intabsmax"`, v)
}

// Exec executes the update operation
func (uqs ControlUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := ControlQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "heatcontrol_control" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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

// Profile returns the set of Profile referencing this Control instance
func (c *Control) Profile() ProfileQS {
	return ProfileQS{}.ControlEq(c)
}

// Scheduledoverride returns the set of Scheduledoverride referencing this Control instance
func (c *Control) Scheduledoverride() ScheduledoverrideQS {
	return ScheduledoverrideQS{}.ControlEq(c)
}

// Instantprofileentry returns the set of Instantprofileentry referencing this Control instance
func (c *Control) Instantprofileentry() InstantprofileentryQS {
	return InstantprofileentryQS{}.ControlEq(c)
}
