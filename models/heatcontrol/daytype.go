// Code generated for Django model heatcontrol.DayType. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package heatcontrol

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/rkojedzinszky/thermo-center/models"
)

// Daytype mirrors model heatcontrol.DayType
type Daytype struct {
	existsInDB bool

	id   int32
	Name string
}

// DaytypeList is a list of Daytype
type DaytypeList []*Daytype

// DaytypeQS represents a queryset for heatcontrol.DayType
type DaytypeQS struct {
	condFragments models.AndFragment
	order         []string
	forClause     string
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

type inDaytypeid []interface{}

func (in inDaytypeid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs DaytypeQS) IDIn(values []int32) DaytypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDaytypeid(vals),
	)

	return qs
}

type notinDaytypeid []interface{}

func (in notinDaytypeid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs DaytypeQS) IDNotIn(values []int32) DaytypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDaytypeid(vals),
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

type inDaytypeName []interface{}

func (in inDaytypeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs DaytypeQS) NameIn(values []string) DaytypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inDaytypeName(vals),
	)

	return qs
}

type notinDaytypeName []interface{}

func (in notinDaytypeName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs DaytypeQS) NameNotIn(values []string) DaytypeQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinDaytypeName(vals),
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

// OrderByRandom randomizes result
func (qs DaytypeQS) OrderByRandom() DaytypeQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs DaytypeQS) ForUpdate() DaytypeQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs DaytypeQS) ForUpdateNowait() DaytypeQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs DaytypeQS) ForUpdateSkipLocked() DaytypeQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs DaytypeQS) ClearForUpdate() DaytypeQS {
	qs.forClause = ""

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
	s += qs.forClause

	return `SELECT "id", "name" FROM "heatcontrol_daytype"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs DaytypeQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_daytype"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs DaytypeQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	row := db.QueryRow(ctx, `SELECT COUNT("id") FROM "heatcontrol_daytype"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs DaytypeQS) All(ctx context.Context, db models.DBInterface) (DaytypeList, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret DaytypeList
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
func (qs DaytypeQS) First(ctx context.Context, db models.DBInterface) (*Daytype, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Daytype{existsInDB: true}
	err := row.Scan(&obj.id, &obj.Name)
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
func (qs DaytypeQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_daytype"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
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
func (uqs DaytypeUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (d *Daytype) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "heatcontrol_daytype" ("name") VALUES ($1) RETURNING "id"`, d.Name)

	if err := row.Scan(&d.id); err != nil {
		return err
	}

	d.existsInDB = true

	return nil
}

// update operation
func (d *Daytype) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "heatcontrol_daytype" SET "name" = $1 WHERE "id" = $2`, d.Name, d.id)

	return err
}

// Save inserts or updates record
func (d *Daytype) Save(ctx context.Context, db models.DBInterface) error {
	if d.existsInDB {
		return d.update(ctx, db)
	}

	return d.insert(ctx, db)
}

// Delete removes row from database
func (d *Daytype) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "heatcontrol_daytype" WHERE "id" = $1`, d.id)

	d.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (dl DaytypeList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts DaytypeList

	for _, d := range dl {
		if d.existsInDB {
			if err := d.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, d)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 1*len(inserts))
	offs := 1
	for _, d := range inserts {
		vva = append(vva, fmt.Sprintf("($%d)", offs+0))
		vaa = append(vaa, d.Name)
		offs += 1
	}

	qs := `INSERT INTO "heatcontrol_daytype" ("name") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, d := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&d.id); err != nil {
			return err
		}

		d.existsInDB = true
	}

	return nil
}

// Calendar returns the set of Calendar referencing this Daytype instance
func (d *Daytype) Calendar() CalendarQS {
	return CalendarQS{}.DaytypeEq(d)
}

// Profile returns the set of Profile referencing this Daytype instance
func (d *Daytype) Profile() ProfileQS {
	return ProfileQS{}.DaytypeEq(d)
}
