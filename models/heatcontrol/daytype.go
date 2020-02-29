// AUTO-GENERATED file for Django model heatcontrol.DayType

package heatcontrol

import (
	"database/sql"
	"github.com/rkojedzinszky/thermo-center/models"
	"strings"
)

// Daytype mirrors model heatcontrol.DayType
type Daytype struct {
	existsInDB bool

	id   int32
	Name string
}

// DaytypeQS represents a queryset for heatcontrol.DayType
type DaytypeQS struct {
	condFragments []models.ConditionFragment
	order         []string
	forUpdate     bool
}

func (qs DaytypeQS) filter(c string, p interface{}) DaytypeQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.UnaryFragment{
			Frag:  c,
			Param: p,
		},
	)
	return qs
}

// GetId returns Daytype.Id
func (d *Daytype) GetId() int32 {
	return d.id
}

// IdEq filters for id being equal to argument
func (qs DaytypeQS) IdEq(v int32) DaytypeQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for id being not equal to argument
func (qs DaytypeQS) IdNe(v int32) DaytypeQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for id being less than argument
func (qs DaytypeQS) IdLt(v int32) DaytypeQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for id being less than or equal to argument
func (qs DaytypeQS) IdLe(v int32) DaytypeQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for id being greater than argument
func (qs DaytypeQS) IdGt(v int32) DaytypeQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for id being greater than or equal to argument
func (qs DaytypeQS) IdGe(v int32) DaytypeQS {
	return qs.filter(`"id" >=`, v)
}

// OrderById sorts result by Id in ascending order
func (qs DaytypeQS) OrderById() DaytypeQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs DaytypeQS) OrderByIdDesc() DaytypeQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// NameEq filters for Name being equal to argument
func (qs DaytypeQS) NameEq(v string) DaytypeQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs DaytypeQS) NameNe(v string) DaytypeQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs DaytypeQS) NameLt(v string) DaytypeQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs DaytypeQS) NameLe(v string) DaytypeQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs DaytypeQS) NameGt(v string) DaytypeQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs DaytypeQS) NameGe(v string) DaytypeQS {
	return qs.filter(`"name" >=`, v)
}

// OrderByName sorts result by Name in ascending order
func (qs DaytypeQS) OrderByName() DaytypeQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs DaytypeQS) OrderByNameDesc() DaytypeQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

func (qs DaytypeQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
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
func (qs DaytypeQS) ForUpdate() DaytypeQS {
	qs.forUpdate = true

	return qs
}

func (qs DaytypeQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs DaytypeQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs DaytypeQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "name" FROM "heatcontrol_daytype"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs DaytypeQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_daytype"` + s, p
}

// All returns all rows matching queryset filters
func (qs DaytypeQS) All(db models.DBInterface) ([]*Daytype, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Daytype
	for rows.Next() {
		obj := Daytype{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.Name); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs DaytypeQS) First(db models.DBInterface) (*Daytype, error) {
	s, p := qs.queryFull()

	row := db.QueryRow(s, p...)

	obj := Daytype{existsInDB: true}
	err := row.Scan(&obj.id, &obj.Name)
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
func (d *Daytype) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "heatcontrol_daytype" ("name") VALUES ($1) RETURNING "id"`, d.Name)

	if err := row.Scan(&d.id); err != nil {
		return err
	}

	d.existsInDB = true

	return nil
}

// update operation
func (d *Daytype) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "heatcontrol_daytype" SET "name" = $1 WHERE "id" = $2`, d.Name, d.id)

	return err
}

// Save inserts or updates record
func (d *Daytype) Save(db models.DBInterface) error {
	if d.existsInDB {
		return d.update(db)
	}

	return d.insert(db)
}

// Delete removes row from database
func (d *Daytype) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "heatcontrol_daytype" WHERE "id" = $1`, d.id)

	d.existsInDB = false

	return err
}
