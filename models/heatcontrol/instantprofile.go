// Code generated for Django model heatcontrol.InstantProfile. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center/v5 center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package heatcontrol

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/rkojedzinszky/thermo-center/v5/models"
)

// Instantprofile mirrors model heatcontrol.InstantProfile
type Instantprofile struct {
	existsInDB bool

	id     int32
	Name   string
	Active bool
}

// InstantprofileList is a list of Instantprofile
type InstantprofileList []*Instantprofile

// InstantprofileQS represents a queryset for heatcontrol.InstantProfile
type InstantprofileQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
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

// BEGIN - heatcontrol.InstantProfile.id

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

type inInstantprofileid []interface{}

func (in inInstantprofileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs InstantprofileQS) IDIn(values []int32) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inInstantprofileid(vals),
	)

	return qs
}

type notinInstantprofileid []interface{}

func (in notinInstantprofileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs InstantprofileQS) IDNotIn(values []int32) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinInstantprofileid(vals),
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

// DistinctOnID marks field in queries to add to DISTINCT ON clause
func (qs InstantprofileQS) DistinctOnID() InstantprofileQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"id"`)

	return qs
}

// END - heatcontrol.InstantProfile.id

// BEGIN - heatcontrol.InstantProfile.name

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

type inInstantprofileName []interface{}

func (in inInstantprofileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs InstantprofileQS) NameIn(values []string) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inInstantprofileName(vals),
	)

	return qs
}

type notinInstantprofileName []interface{}

func (in notinInstantprofileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs InstantprofileQS) NameNotIn(values []string) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinInstantprofileName(vals),
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

// DistinctOnName marks field in queries to add to DISTINCT ON clause
func (qs InstantprofileQS) DistinctOnName() InstantprofileQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"name"`)

	return qs
}

// END - heatcontrol.InstantProfile.name

// BEGIN - heatcontrol.InstantProfile.active

// ActiveEq filters for Active being equal to argument
func (qs InstantprofileQS) ActiveEq(v bool) InstantprofileQS {
	return qs.filter(`"active" =`, v)
}

// ActiveNe filters for Active being not equal to argument
func (qs InstantprofileQS) ActiveNe(v bool) InstantprofileQS {
	return qs.filter(`"active" <>`, v)
}

type inInstantprofileActive []interface{}

func (in inInstantprofileActive) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"active" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs InstantprofileQS) ActiveIn(values []bool) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inInstantprofileActive(vals),
	)

	return qs
}

type notinInstantprofileActive []interface{}

func (in notinInstantprofileActive) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"active" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs InstantprofileQS) ActiveNotIn(values []bool) InstantprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinInstantprofileActive(vals),
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

// DistinctOnActive marks field in queries to add to DISTINCT ON clause
func (qs InstantprofileQS) DistinctOnActive() InstantprofileQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"active"`)

	return qs
}

// END - heatcontrol.InstantProfile.active

// OrderByRandom randomizes result
func (qs InstantprofileQS) OrderByRandom() InstantprofileQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs InstantprofileQS) ForUpdate() InstantprofileQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs InstantprofileQS) ForUpdateNowait() InstantprofileQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs InstantprofileQS) ForUpdateSkipLocked() InstantprofileQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs InstantprofileQS) ClearForUpdate() InstantprofileQS {
	qs.forClause = ""

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

func (qs InstantprofileQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"id", "name", "active" FROM "heatcontrol_instantprofile"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs InstantprofileQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_instantprofile"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs InstantprofileQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"id"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "heatcontrol_instantprofile"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs InstantprofileQS) All(ctx context.Context, db models.DBInterface) (InstantprofileList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret InstantprofileList
	for rows.Next() {
		obj := Instantprofile{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.Name, &obj.Active); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs InstantprofileQS) First(ctx context.Context, db models.DBInterface) (*Instantprofile, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Instantprofile{existsInDB: true}
	err := row.Scan(&obj.id, &obj.Name, &obj.Active)
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
func (qs InstantprofileQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_instantprofile"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
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
func (uqs InstantprofileUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (i *Instantprofile) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "heatcontrol_instantprofile" ("name", "active") VALUES ($1, $2) RETURNING "id"`, i.Name, i.Active)

	if err := row.Scan(&i.id); err != nil {
		return err
	}

	i.existsInDB = true

	return nil
}

// update operation
func (i *Instantprofile) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "heatcontrol_instantprofile" SET "name" = $1, "active" = $2 WHERE "id" = $3`, i.Name, i.Active, i.id)

	return err
}

// Save inserts or updates record
func (i *Instantprofile) Save(ctx context.Context, db models.DBInterface) error {
	if i.existsInDB {
		return i.update(ctx, db)
	}

	return i.insert(ctx, db)
}

// Delete removes row from database
func (i *Instantprofile) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "heatcontrol_instantprofile" WHERE "id" = $1`, i.id)

	i.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (il InstantprofileList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts InstantprofileList

	for _, i := range il {
		if i.existsInDB {
			if err := i.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, i)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 2*len(inserts))
	offs := 1
	for _, i := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d)", offs+0, offs+1))
		vaa = append(vaa, i.Name, i.Active)
		offs += 2
	}

	qs := `INSERT INTO "heatcontrol_instantprofile" ("name", "active") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, i := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&i.id); err != nil {
			return err
		}

		i.existsInDB = true
	}

	return nil
}

// Instantprofileentry returns the set of Instantprofileentry referencing this Instantprofile instance
func (i *Instantprofile) Instantprofileentry() InstantprofileentryQS {
	return InstantprofileentryQS{}.ProfileEq(i)
}
