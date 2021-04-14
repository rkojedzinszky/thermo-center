/*
  AUTO-GENERATED file for Django model heatcontrol.DayType

  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

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
	condFragments models.AndFragment
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

// Or combines given expressions with OR operator
func (qs DaytypeQS) Or(exprs ...DaytypeQS) DaytypeQS {
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

// GetID returns Daytype.ID
func (d *Daytype) GetID() int32 {
	return d.id
}

// IDEq filters for id being equal to argument
func (qs DaytypeQS) IDEq(v int32) DaytypeQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs DaytypeQS) IDNe(v int32) DaytypeQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs DaytypeQS) IDLt(v int32) DaytypeQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs DaytypeQS) IDLe(v int32) DaytypeQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs DaytypeQS) IDGt(v int32) DaytypeQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs DaytypeQS) IDGe(v int32) DaytypeQS {
	return qs.filter(`"id" >=`, v)
}

type inDaytypeid struct {
	values []interface{}
}

func (in *inDaytypeid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs DaytypeQS) IDIn(values []int32) DaytypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDaytypeid{
			values: vals,
		},
	)

	return qs
}

type notinDaytypeid struct {
	values []interface{}
}

func (in *notinDaytypeid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs DaytypeQS) IDNotIn(values []int32) DaytypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDaytypeid{
			values: vals,
		},
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs DaytypeQS) OrderByID() DaytypeQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs DaytypeQS) OrderByIDDesc() DaytypeQS {
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

type inDaytypeName struct {
	values []interface{}
}

func (in *inDaytypeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs DaytypeQS) NameIn(values []string) DaytypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inDaytypeName{
			values: vals,
		},
	)

	return qs
}

type notinDaytypeName struct {
	values []interface{}
}

func (in *notinDaytypeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs DaytypeQS) NameNotIn(values []string) DaytypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinDaytypeName{
			values: vals,
		},
	)

	return qs
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

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs DaytypeQS) ForUpdate() DaytypeQS {
	qs.forUpdate = true

	return qs
}

func (qs DaytypeQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

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

	s += " LIMIT 1"

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

// Delete deletes rows matching queryset filters
func (qs DaytypeQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_daytype"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs DaytypeQS) Update() DaytypeUpdateQS {
	return DaytypeUpdateQS{condFragments: qs.condFragments}
}

// DaytypeUpdateQS represents an updated queryset for heatcontrol.DayType
type DaytypeUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs DaytypeUpdateQS) update(c string, v interface{}) DaytypeUpdateQS {
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
func (uqs DaytypeUpdateQS) SetID(v int32) DaytypeUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetName sets Name to the given value
func (uqs DaytypeUpdateQS) SetName(v string) DaytypeUpdateQS {
	return uqs.update(`"name"`, v)
}

// Exec executes the update operation
func (uqs DaytypeUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := DaytypeQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "heatcontrol_daytype" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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

// Calendar returns the set of Calendar referencing this Daytype instance
func (d *Daytype) Calendar() CalendarQS {
	return CalendarQS{}.DaytypeEq(d)
}

// Profile returns the set of Profile referencing this Daytype instance
func (d *Daytype) Profile() ProfileQS {
	return ProfileQS{}.DaytypeEq(d)
}
