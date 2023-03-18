// Code generated for Django model center.Sensor. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package center

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/rkojedzinszky/thermo-center/models"
)

// Sensor mirrors model center.Sensor
type Sensor struct {
	existsInDB bool

	ID      int32
	Name    string
	LastSeq sql.NullInt32
	LastTsf sql.NullFloat64
}

// SensorList is a list of Sensor
type SensorList []*Sensor

// SensorQS represents a queryset for center.Sensor
type SensorQS struct {
	condFragments models.AndFragment
	order         []string
	forClause     string
}

func (qs SensorQS) filter(c string, p interface{}) SensorQS {
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
func (qs SensorQS) Or(exprs ...SensorQS) SensorQS {
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

// IDEq filters for ID being equal to argument
func (qs SensorQS) IDEq(v int32) SensorQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for ID being not equal to argument
func (qs SensorQS) IDNe(v int32) SensorQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for ID being less than argument
func (qs SensorQS) IDLt(v int32) SensorQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for ID being less than or equal to argument
func (qs SensorQS) IDLe(v int32) SensorQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for ID being greater than argument
func (qs SensorQS) IDGt(v int32) SensorQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for ID being greater than or equal to argument
func (qs SensorQS) IDGe(v int32) SensorQS {
	return qs.filter(`"id" >=`, v)
}

type inSensorID []interface{}

func (in inSensorID) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorQS) IDIn(values []int32) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inSensorID(vals),
	)

	return qs
}

type notinSensorID []interface{}

func (in notinSensorID) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorQS) IDNotIn(values []int32) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinSensorID(vals),
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs SensorQS) OrderByID() SensorQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs SensorQS) OrderByIDDesc() SensorQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// NameEq filters for Name being equal to argument
func (qs SensorQS) NameEq(v string) SensorQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs SensorQS) NameNe(v string) SensorQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs SensorQS) NameLt(v string) SensorQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs SensorQS) NameLe(v string) SensorQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs SensorQS) NameGt(v string) SensorQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs SensorQS) NameGe(v string) SensorQS {
	return qs.filter(`"name" >=`, v)
}

type inSensorName []interface{}

func (in inSensorName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorQS) NameIn(values []string) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inSensorName(vals),
	)

	return qs
}

type notinSensorName []interface{}

func (in notinSensorName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorQS) NameNotIn(values []string) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinSensorName(vals),
	)

	return qs
}

// OrderByName sorts result by Name in ascending order
func (qs SensorQS) OrderByName() SensorQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs SensorQS) OrderByNameDesc() SensorQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

// LastSeqIsNull filters for LastSeq being null
func (qs SensorQS) LastSeqIsNull() SensorQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"last_seq" IS NULL`,
		},
	)
	return qs
}

// LastSeqIsNotNull filters for LastSeq being not null
func (qs SensorQS) LastSeqIsNotNull() SensorQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"last_seq" IS NOT NULL`,
		},
	)
	return qs
}

// LastSeqEq filters for LastSeq being equal to argument
func (qs SensorQS) LastSeqEq(v int32) SensorQS {
	return qs.filter(`"last_seq" =`, v)
}

// LastSeqNe filters for LastSeq being not equal to argument
func (qs SensorQS) LastSeqNe(v int32) SensorQS {
	return qs.filter(`"last_seq" <>`, v)
}

// LastSeqLt filters for LastSeq being less than argument
func (qs SensorQS) LastSeqLt(v int32) SensorQS {
	return qs.filter(`"last_seq" <`, v)
}

// LastSeqLe filters for LastSeq being less than or equal to argument
func (qs SensorQS) LastSeqLe(v int32) SensorQS {
	return qs.filter(`"last_seq" <=`, v)
}

// LastSeqGt filters for LastSeq being greater than argument
func (qs SensorQS) LastSeqGt(v int32) SensorQS {
	return qs.filter(`"last_seq" >`, v)
}

// LastSeqGe filters for LastSeq being greater than or equal to argument
func (qs SensorQS) LastSeqGe(v int32) SensorQS {
	return qs.filter(`"last_seq" >=`, v)
}

type inSensorLastSeq []interface{}

func (in inSensorLastSeq) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"last_seq" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorQS) LastSeqIn(values []int32) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inSensorLastSeq(vals),
	)

	return qs
}

type notinSensorLastSeq []interface{}

func (in notinSensorLastSeq) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"last_seq" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorQS) LastSeqNotIn(values []int32) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinSensorLastSeq(vals),
	)

	return qs
}

// OrderByLastSeq sorts result by LastSeq in ascending order
func (qs SensorQS) OrderByLastSeq() SensorQS {
	qs.order = append(qs.order, `"last_seq"`)

	return qs
}

// OrderByLastSeqDesc sorts result by LastSeq in descending order
func (qs SensorQS) OrderByLastSeqDesc() SensorQS {
	qs.order = append(qs.order, `"last_seq" DESC`)

	return qs
}

// LastTsfIsNull filters for LastTsf being null
func (qs SensorQS) LastTsfIsNull() SensorQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"last_tsf" IS NULL`,
		},
	)
	return qs
}

// LastTsfIsNotNull filters for LastTsf being not null
func (qs SensorQS) LastTsfIsNotNull() SensorQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"last_tsf" IS NOT NULL`,
		},
	)
	return qs
}

// LastTsfEq filters for LastTsf being equal to argument
func (qs SensorQS) LastTsfEq(v float64) SensorQS {
	return qs.filter(`"last_tsf" =`, v)
}

// LastTsfNe filters for LastTsf being not equal to argument
func (qs SensorQS) LastTsfNe(v float64) SensorQS {
	return qs.filter(`"last_tsf" <>`, v)
}

// LastTsfLt filters for LastTsf being less than argument
func (qs SensorQS) LastTsfLt(v float64) SensorQS {
	return qs.filter(`"last_tsf" <`, v)
}

// LastTsfLe filters for LastTsf being less than or equal to argument
func (qs SensorQS) LastTsfLe(v float64) SensorQS {
	return qs.filter(`"last_tsf" <=`, v)
}

// LastTsfGt filters for LastTsf being greater than argument
func (qs SensorQS) LastTsfGt(v float64) SensorQS {
	return qs.filter(`"last_tsf" >`, v)
}

// LastTsfGe filters for LastTsf being greater than or equal to argument
func (qs SensorQS) LastTsfGe(v float64) SensorQS {
	return qs.filter(`"last_tsf" >=`, v)
}

type inSensorLastTsf []interface{}

func (in inSensorLastTsf) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"last_tsf" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorQS) LastTsfIn(values []float64) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inSensorLastTsf(vals),
	)

	return qs
}

type notinSensorLastTsf []interface{}

func (in notinSensorLastTsf) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"last_tsf" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs SensorQS) LastTsfNotIn(values []float64) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinSensorLastTsf(vals),
	)

	return qs
}

// OrderByLastTsf sorts result by LastTsf in ascending order
func (qs SensorQS) OrderByLastTsf() SensorQS {
	qs.order = append(qs.order, `"last_tsf"`)

	return qs
}

// OrderByLastTsfDesc sorts result by LastTsf in descending order
func (qs SensorQS) OrderByLastTsfDesc() SensorQS {
	qs.order = append(qs.order, `"last_tsf" DESC`)

	return qs
}

// OrderByRandom randomizes result
func (qs SensorQS) OrderByRandom() SensorQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs SensorQS) ForUpdate() SensorQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs SensorQS) ForUpdateNowait() SensorQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs SensorQS) ForUpdateSkipLocked() SensorQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs SensorQS) ClearForUpdate() SensorQS {
	qs.forClause = ""

	return qs
}

func (qs SensorQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs SensorQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs SensorQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	return `SELECT "id", "name", "last_seq", "last_tsf" FROM "center_sensor"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs SensorQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "center_sensor"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs SensorQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	row := db.QueryRow(ctx, `SELECT COUNT("id") FROM "center_sensor"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs SensorQS) All(ctx context.Context, db models.DBInterface) (SensorList, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret SensorList
	for rows.Next() {
		obj := Sensor{existsInDB: true}
		if err = rows.Scan(&obj.ID, &obj.Name, &obj.LastSeq, &obj.LastTsf); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs SensorQS) First(ctx context.Context, db models.DBInterface) (*Sensor, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Sensor{existsInDB: true}
	err := row.Scan(&obj.ID, &obj.Name, &obj.LastSeq, &obj.LastTsf)
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
func (qs SensorQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "center_sensor"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs SensorQS) Update() SensorUpdateQS {
	return SensorUpdateQS{condFragments: qs.condFragments}
}

// SensorUpdateQS represents an updated queryset for center.Sensor
type SensorUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs SensorUpdateQS) update(c string, v interface{}) SensorUpdateQS {
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
func (uqs SensorUpdateQS) SetID(v int32) SensorUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetName sets Name to the given value
func (uqs SensorUpdateQS) SetName(v string) SensorUpdateQS {
	return uqs.update(`"name"`, v)
}

// SetLastSeq sets LastSeq to the given value
func (uqs SensorUpdateQS) SetLastSeq(v sql.NullInt32) SensorUpdateQS {
	return uqs.update(`"last_seq"`, v)
}

// SetLastTsf sets LastTsf to the given value
func (uqs SensorUpdateQS) SetLastTsf(v sql.NullFloat64) SensorUpdateQS {
	return uqs.update(`"last_tsf"`, v)
}

// Exec executes the update operation
func (uqs SensorUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := SensorQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "center_sensor" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (s *Sensor) insert(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `INSERT INTO "center_sensor" ("name", "last_seq", "last_tsf", "id") VALUES ($1, $2, $3, $4)`, s.Name, s.LastSeq, s.LastTsf, s.ID)

	if err != nil {
		return err
	}

	s.existsInDB = true

	return nil
}

// update operation
func (s *Sensor) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "center_sensor" SET "name" = $1, "last_seq" = $2, "last_tsf" = $3 WHERE "id" = $4`, s.Name, s.LastSeq, s.LastTsf, s.ID)

	return err
}

// Save inserts or updates record
func (s *Sensor) Save(ctx context.Context, db models.DBInterface) error {
	if s.existsInDB {
		return s.update(ctx, db)
	}

	return s.insert(ctx, db)
}

// Delete removes row from database
func (s *Sensor) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "center_sensor" WHERE "id" = $1`, s.ID)

	s.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (sl SensorList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts SensorList

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
	vaa := make([]any, 0, 4*len(inserts))
	offs := 1
	for _, s := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3))
		vaa = append(vaa, s.Name, s.LastSeq, s.LastTsf, s.ID)
		offs += 4
	}

	qs := `INSERT INTO "center_sensor" ("name", "last_seq", "last_tsf", "id") VALUES ` + strings.Join(vva, ", ")
	_, err := db.Exec(ctx, qs, vaa...)

	if err != nil {
		return err
	}

	for _, s := range inserts {
		s.existsInDB = true
	}

	return nil
}

// Sensorresync returns the set of Sensorresync referencing this Sensor instance
func (s *Sensor) Sensorresync() SensorresyncQS {
	return SensorresyncQS{}.SensorEq(s)
}

// Configuresensortask returns the set of Configuresensortask referencing this Sensor instance
func (s *Sensor) Configuresensortask() ConfiguresensortaskQS {
	return ConfiguresensortaskQS{}.SensorEq(s)
}
