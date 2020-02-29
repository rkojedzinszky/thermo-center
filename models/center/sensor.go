// AUTO-GENERATED file for Django model center.Sensor

package center

import (
	"database/sql"
	"github.com/rkojedzinszky/thermo-center/models"
	"strings"
)

// Sensor mirrors model center.Sensor
type Sensor struct {
	existsInDB bool

	Id      int32
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

// IdEq filters for Id being equal to argument
func (qs SensorQS) IdEq(v int32) SensorQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for Id being not equal to argument
func (qs SensorQS) IdNe(v int32) SensorQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for Id being less than argument
func (qs SensorQS) IdLt(v int32) SensorQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for Id being less than or equal to argument
func (qs SensorQS) IdLe(v int32) SensorQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for Id being greater than argument
func (qs SensorQS) IdGt(v int32) SensorQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for Id being greater than or equal to argument
func (qs SensorQS) IdGe(v int32) SensorQS {
	return qs.filter(`"id" >=`, v)
}

// OrderById sorts result by Id in ascending order
func (qs SensorQS) OrderById() SensorQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs SensorQS) OrderByIdDesc() SensorQS {
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
		if err = rows.Scan(&obj.Id, &obj.Name, &obj.LastSeq, &obj.LastTsf); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs SensorQS) First(db models.DBInterface) (*Sensor, error) {
	s, p := qs.queryFull()

	row := db.QueryRow(s, p...)

	obj := Sensor{existsInDB: true}
	err := row.Scan(&obj.Id, &obj.Name, &obj.LastSeq, &obj.LastTsf)
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
func (s *Sensor) insert(db models.DBInterface) error {
	_, err := db.Exec(`INSERT INTO "center_sensor" ("id", "name", "last_seq", "last_tsf") VALUES ($1, $2, $3, $4)`, s.Id, s.Name, s.LastSeq, s.LastTsf)

	if err != nil {
		return err
	}

	s.existsInDB = true

	return nil
}

// update operation
func (s *Sensor) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "center_sensor" SET "id" = $1, "name" = $2, "last_seq" = $3, "last_tsf" = $4 WHERE "id" = $5`, s.Id, s.Name, s.LastSeq, s.LastTsf, s.Id)

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
	_, err := db.Exec(`DELETE FROM "center_sensor" WHERE "id" = $1`, s.Id)

	s.existsInDB = false

	return err
}
