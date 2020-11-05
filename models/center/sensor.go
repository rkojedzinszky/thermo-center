/*
  AUTO-GENERATED file for Django model center.Sensor

  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package center

import (
	"database/sql"
	"github.com/rkojedzinszky/thermo-center/models"
	"strings"
)

// Sensor mirrors model center.Sensor
type Sensor struct {
	existsInDB bool

	ID      int32
	Name    string
	LastSeq sql.NullInt32
	LastTsf sql.NullFloat64
}

// SensorQS represents a queryset for center.Sensor
type SensorQS struct {
	condFragments []models.ConditionFragment
	order         []string
	forUpdate     bool
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

type inSensorID struct {
	values []interface{}
}

func (in *inSensorID) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorQS) IDIn(values []int32) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inSensorID{
			values: vals,
		},
	)

	return qs
}

type notinSensorID struct {
	values []interface{}
}

func (in *notinSensorID) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorQS) IDNotIn(values []int32) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinSensorID{
			values: vals,
		},
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

type inSensorName struct {
	values []interface{}
}

func (in *inSensorName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorQS) NameIn(values []string) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inSensorName{
			values: vals,
		},
	)

	return qs
}

type notinSensorName struct {
	values []interface{}
}

func (in *notinSensorName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorQS) NameNotIn(values []string) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinSensorName{
			values: vals,
		},
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

type inSensorLastSeq struct {
	values []interface{}
}

func (in *inSensorLastSeq) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"last_seq" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorQS) LastSeqIn(values []int32) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inSensorLastSeq{
			values: vals,
		},
	)

	return qs
}

type notinSensorLastSeq struct {
	values []interface{}
}

func (in *notinSensorLastSeq) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"last_seq" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorQS) LastSeqNotIn(values []int32) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinSensorLastSeq{
			values: vals,
		},
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

type inSensorLastTsf struct {
	values []interface{}
}

func (in *inSensorLastTsf) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"last_tsf" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorQS) LastTsfIn(values []float64) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inSensorLastTsf{
			values: vals,
		},
	)

	return qs
}

type notinSensorLastTsf struct {
	values []interface{}
}

func (in *notinSensorLastTsf) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"last_tsf" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs SensorQS) LastTsfNotIn(values []float64) SensorQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinSensorLastTsf{
			values: vals,
		},
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

func (qs SensorQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
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
func (qs SensorQS) ForUpdate() SensorQS {
	qs.forUpdate = true

	return qs
}

func (qs SensorQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

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
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "name", "last_seq", "last_tsf" FROM "center_sensor"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs SensorQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "center_sensor"` + s, p
}

// All returns all rows matching queryset filters
func (qs SensorQS) All(db models.DBInterface) ([]*Sensor, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Sensor
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
func (qs SensorQS) First(db models.DBInterface) (*Sensor, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Sensor{existsInDB: true}
	err := row.Scan(&obj.ID, &obj.Name, &obj.LastSeq, &obj.LastTsf)
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
func (qs SensorQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "center_sensor"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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
func (uqs SensorUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (s *Sensor) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "center_sensor" ("name", "last_seq", "last_tsf") VALUES ($1, $2, $3)`, s.Name, s.LastSeq, s.LastTsf)

	if err != nil {
		return err
	}

	s.existsInDB = true

	return nil
}

// update operation
func (s *Sensor) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "center_sensor" SET "name" = $1, "last_seq" = $2, "last_tsf" = $3 WHERE "id" = $4`, s.Name, s.LastSeq, s.LastTsf, s.ID)

	return err
}

// Save inserts or updates record
func (s *Sensor) Save(db models.DBInterface) error {
	if s.existsInDB {
		return s.update(db)
	}

	return s.insert(db)
}

// Delete removes row from database
func (s *Sensor) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "center_sensor" WHERE "id" = $1`, s.ID)

	s.existsInDB = false

	return err
}

// Sensorresync returns the set of Sensorresync referencing this Sensor instance
func (s *Sensor) Sensorresync() SensorresyncQS {
	return SensorresyncQS{}.SensorEq(s)
}

// Configuresensortask returns the set of Configuresensortask referencing this Sensor instance
func (s *Sensor) Configuresensortask() ConfiguresensortaskQS {
	return ConfiguresensortaskQS{}.SensorEq(s)
}
