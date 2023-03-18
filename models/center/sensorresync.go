// Code generated for Django model center.SensorResync. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package center

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/rkojedzinszky/thermo-center/models"
)

// Sensorresync mirrors model center.SensorResync
type Sensorresync struct {
	existsInDB bool

	id     int32
	sensor int32
	Ts     time.Time
}

// SensorresyncList is a list of Sensorresync
type SensorresyncList []*Sensorresync

// SensorresyncQS represents a queryset for center.SensorResync
type SensorresyncQS struct {
	condFragments models.AndFragment
	order         []string
	forClause     string
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

// Or combines given expressions with OR operator
func (qs SensorresyncQS) Or(exprs ...SensorresyncQS) SensorresyncQS {
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

type inSensorresyncid []interface{}

func (in inSensorresyncid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorresyncQS) IDIn(values []int32) SensorresyncQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inSensorresyncid(vals),
	)

	return qs
}

type notinSensorresyncid []interface{}

func (in notinSensorresyncid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorresyncQS) IDNotIn(values []int32) SensorresyncQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinSensorresyncid(vals),
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
func (s *Sensorresync) GetSensor(ctx context.Context, db models.DBInterface) (*Sensor, error) {
	return SensorQS{}.IDEq(s.sensor).First(ctx, db)
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

// SensorRawEq filters for sensor being equal to raw argument
func (qs SensorresyncQS) SensorRawEq(v int32) SensorresyncQS {
	return qs.filter(`"sensor_id" =`, v)
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

type inSensorresyncTs []interface{}

func (in inSensorresyncTs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"ts" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorresyncQS) TsIn(values []time.Time) SensorresyncQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inSensorresyncTs(vals),
	)

	return qs
}

type notinSensorresyncTs []interface{}

func (in notinSensorresyncTs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"ts" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorresyncQS) TsNotIn(values []time.Time) SensorresyncQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinSensorresyncTs(vals),
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

// OrderByRandom randomizes result
func (qs SensorresyncQS) OrderByRandom() SensorresyncQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs SensorresyncQS) ForUpdate() SensorresyncQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs SensorresyncQS) ForUpdateNowait() SensorresyncQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs SensorresyncQS) ForUpdateSkipLocked() SensorresyncQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs SensorresyncQS) ClearForUpdate() SensorresyncQS {
	qs.forClause = ""

	return qs
}

func (qs SensorresyncQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

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
	s += qs.forClause

	return `SELECT "id", "sensor_id", "ts" FROM "center_sensorresync"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs SensorresyncQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "center_sensorresync"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs SensorresyncQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	row := db.QueryRow(ctx, `SELECT COUNT("id") FROM "center_sensorresync"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs SensorresyncQS) All(ctx context.Context, db models.DBInterface) (SensorresyncList, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret SensorresyncList
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
func (qs SensorresyncQS) First(ctx context.Context, db models.DBInterface) (*Sensorresync, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Sensorresync{existsInDB: true}
	err := row.Scan(&obj.id, &obj.sensor, &obj.Ts)
	switch err {
	case nil:
		return &obj, nil
	case pgx.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

// Delete deletes rows matching queryset filters
func (qs SensorresyncQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "center_sensorresync"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
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
func (uqs SensorresyncUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (s *Sensorresync) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "center_sensorresync" ("sensor_id", "ts") VALUES ($1, $2) RETURNING "id"`, s.sensor, s.Ts)

	if err := row.Scan(&s.id); err != nil {
		return err
	}

	s.existsInDB = true

	return nil
}

// update operation
func (s *Sensorresync) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "center_sensorresync" SET "sensor_id" = $1, "ts" = $2 WHERE "id" = $3`, s.sensor, s.Ts, s.id)

	return err
}

// Save inserts or updates record
func (s *Sensorresync) Save(ctx context.Context, db models.DBInterface) error {
	if s.existsInDB {
		return s.update(ctx, db)
	}

	return s.insert(ctx, db)
}

// Delete removes row from database
func (s *Sensorresync) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "center_sensorresync" WHERE "id" = $1`, s.id)

	s.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (sl SensorresyncList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts SensorresyncList

	for _, s := range sl {
		if s.existsInDB {
			if err := s.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, s)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 2*len(inserts))
	offs := 1
	for _, s := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d)", offs+0, offs+1))
		vaa = append(vaa, s.sensor, s.Ts)
		offs += 2
	}

	qs := `INSERT INTO "center_sensorresync" ("sensor_id", "ts") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, s := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&s.id); err != nil {
			return err
		}

		s.existsInDB = true
	}

	return nil
}
