/*
  AUTO-GENERATED file for Django model heatcontrol.InstantProfile

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

// Instantprofile mirrors model heatcontrol.InstantProfile
type Instantprofile struct {
	existsInDB bool

	id     int32
	Name   string
	Active bool
}

// InstantprofileQS represents a queryset for heatcontrol.InstantProfile
type InstantprofileQS struct {
	condFragments models.AndFragment
	order         []string
	forUpdate     bool
}

func (qs InstantprofileQS) filter(c string, p interface{}) InstantprofileQS {
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
func (qs InstantprofileQS) Or(exprs ...InstantprofileQS) InstantprofileQS {
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

// GetID returns Instantprofile.ID
func (i *Instantprofile) GetID() int32 {
	return i.id
}

// IDEq filters for id being equal to argument
func (qs InstantprofileQS) IDEq(v int32) InstantprofileQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs InstantprofileQS) IDNe(v int32) InstantprofileQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs InstantprofileQS) IDLt(v int32) InstantprofileQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs InstantprofileQS) IDLe(v int32) InstantprofileQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs InstantprofileQS) IDGt(v int32) InstantprofileQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs InstantprofileQS) IDGe(v int32) InstantprofileQS {
	return qs.filter(`"id" >=`, v)
}

type inInstantprofileid struct {
	values []interface{}
}

func (in *inInstantprofileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileQS) IDIn(values []int32) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inInstantprofileid{
			values: vals,
		},
	)

	return qs
}

type notinInstantprofileid struct {
	values []interface{}
}

func (in *notinInstantprofileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileQS) IDNotIn(values []int32) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinInstantprofileid{
			values: vals,
		},
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs InstantprofileQS) OrderByID() InstantprofileQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs InstantprofileQS) OrderByIDDesc() InstantprofileQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// NameEq filters for Name being equal to argument
func (qs InstantprofileQS) NameEq(v string) InstantprofileQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs InstantprofileQS) NameNe(v string) InstantprofileQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs InstantprofileQS) NameLt(v string) InstantprofileQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs InstantprofileQS) NameLe(v string) InstantprofileQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs InstantprofileQS) NameGt(v string) InstantprofileQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs InstantprofileQS) NameGe(v string) InstantprofileQS {
	return qs.filter(`"name" >=`, v)
}

type inInstantprofileName struct {
	values []interface{}
}

func (in *inInstantprofileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileQS) NameIn(values []string) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inInstantprofileName{
			values: vals,
		},
	)

	return qs
}

type notinInstantprofileName struct {
	values []interface{}
}

func (in *notinInstantprofileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileQS) NameNotIn(values []string) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinInstantprofileName{
			values: vals,
		},
	)

	return qs
}

// OrderByName sorts result by Name in ascending order
func (qs InstantprofileQS) OrderByName() InstantprofileQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs InstantprofileQS) OrderByNameDesc() InstantprofileQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

// ActiveEq filters for Active being equal to argument
func (qs InstantprofileQS) ActiveEq(v bool) InstantprofileQS {
	return qs.filter(`"active" =`, v)
}

// ActiveNe filters for Active being not equal to argument
func (qs InstantprofileQS) ActiveNe(v bool) InstantprofileQS {
	return qs.filter(`"active" <>`, v)
}

// ActiveLt filters for Active being less than argument
func (qs InstantprofileQS) ActiveLt(v bool) InstantprofileQS {
	return qs.filter(`"active" <`, v)
}

// ActiveLe filters for Active being less than or equal to argument
func (qs InstantprofileQS) ActiveLe(v bool) InstantprofileQS {
	return qs.filter(`"active" <=`, v)
}

// ActiveGt filters for Active being greater than argument
func (qs InstantprofileQS) ActiveGt(v bool) InstantprofileQS {
	return qs.filter(`"active" >`, v)
}

// ActiveGe filters for Active being greater than or equal to argument
func (qs InstantprofileQS) ActiveGe(v bool) InstantprofileQS {
	return qs.filter(`"active" >=`, v)
}

type inInstantprofileActive struct {
	values []interface{}
}

func (in *inInstantprofileActive) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"active" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileQS) ActiveIn(values []bool) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inInstantprofileActive{
			values: vals,
		},
	)

	return qs
}

type notinInstantprofileActive struct {
	values []interface{}
}

func (in *notinInstantprofileActive) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"active" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileQS) ActiveNotIn(values []bool) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinInstantprofileActive{
			values: vals,
		},
	)

	return qs
}

// OrderByActive sorts result by Active in ascending order
func (qs InstantprofileQS) OrderByActive() InstantprofileQS {
	qs.order = append(qs.order, `"active"`)

	return qs
}

// OrderByActiveDesc sorts result by Active in descending order
func (qs InstantprofileQS) OrderByActiveDesc() InstantprofileQS {
	qs.order = append(qs.order, `"active" DESC`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs InstantprofileQS) ForUpdate() InstantprofileQS {
	qs.forUpdate = true

	return qs
}

func (qs InstantprofileQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs InstantprofileQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs InstantprofileQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "name", "active" FROM "heatcontrol_instantprofile"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs InstantprofileQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_instantprofile"` + s, p
}

// All returns all rows matching queryset filters
func (qs InstantprofileQS) All(db models.DBInterface) ([]*Instantprofile, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Instantprofile
	for rows.Next() {
		obj := Instantprofile{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.Name, &obj.Active); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs InstantprofileQS) First(db models.DBInterface) (*Instantprofile, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Instantprofile{existsInDB: true}
	err := row.Scan(&obj.id, &obj.Name, &obj.Active)
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
func (qs InstantprofileQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_instantprofile"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs InstantprofileQS) Update() InstantprofileUpdateQS {
	return InstantprofileUpdateQS{condFragments: qs.condFragments}
}

// InstantprofileUpdateQS represents an updated queryset for heatcontrol.InstantProfile
type InstantprofileUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs InstantprofileUpdateQS) update(c string, v interface{}) InstantprofileUpdateQS {
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
func (uqs InstantprofileUpdateQS) SetID(v int32) InstantprofileUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetName sets Name to the given value
func (uqs InstantprofileUpdateQS) SetName(v string) InstantprofileUpdateQS {
	return uqs.update(`"name"`, v)
}

// SetActive sets Active to the given value
func (uqs InstantprofileUpdateQS) SetActive(v bool) InstantprofileUpdateQS {
	return uqs.update(`"active"`, v)
}

// Exec executes the update operation
func (uqs InstantprofileUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := InstantprofileQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "heatcontrol_instantprofile" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (i *Instantprofile) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "heatcontrol_instantprofile" ("name", "active") VALUES ($1, $2) RETURNING "id"`, i.Name, i.Active)

	if err := row.Scan(&i.id); err != nil {
		return err
	}

	i.existsInDB = true

	return nil
}

// update operation
func (i *Instantprofile) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "heatcontrol_instantprofile" SET "name" = $1, "active" = $2 WHERE "id" = $3`, i.Name, i.Active, i.id)

	return err
}

// Save inserts or updates record
func (i *Instantprofile) Save(db models.DBInterface) error {
	if i.existsInDB {
		return i.update(db)
	}

	return i.insert(db)
}

// Delete removes row from database
func (i *Instantprofile) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "heatcontrol_instantprofile" WHERE "id" = $1`, i.id)

	i.existsInDB = false

	return err
}

// Instantprofileentry returns the set of Instantprofileentry referencing this Instantprofile instance
func (i *Instantprofile) Instantprofileentry() InstantprofileentryQS {
	return InstantprofileentryQS{}.ProfileEq(i)
}
