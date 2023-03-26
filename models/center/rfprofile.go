// Code generated for Django model center.RFProfile. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center/v5 center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package center

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/rkojedzinszky/thermo-center/v5/models"
)

// Rfprofile mirrors model center.RFProfile
type Rfprofile struct {
	existsInDB bool

	id       int32
	Name     string
	Confregs string
}

// RfprofileList is a list of Rfprofile
type RfprofileList []*Rfprofile

// RfprofileQS represents a queryset for center.RFProfile
type RfprofileQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs RfprofileQS) filter(c string, p interface{}) RfprofileQS {
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
func (qs RfprofileQS) Or(exprs ...RfprofileQS) RfprofileQS {
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

// BEGIN - center.RFProfile.id

// GetID returns Rfprofile.ID
func (r *Rfprofile) GetID() int32 {
	return r.id
}

// IDEq filters for id being equal to argument
func (qs RfprofileQS) IDEq(v int32) RfprofileQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs RfprofileQS) IDNe(v int32) RfprofileQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs RfprofileQS) IDLt(v int32) RfprofileQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs RfprofileQS) IDLe(v int32) RfprofileQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs RfprofileQS) IDGt(v int32) RfprofileQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs RfprofileQS) IDGe(v int32) RfprofileQS {
	return qs.filter(`"id" >=`, v)
}

type inRfprofileid []interface{}

func (in inRfprofileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfprofileQS) IDIn(values []int32) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inRfprofileid(vals),
	)

	return qs
}

type notinRfprofileid []interface{}

func (in notinRfprofileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfprofileQS) IDNotIn(values []int32) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinRfprofileid(vals),
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs RfprofileQS) OrderByID() RfprofileQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs RfprofileQS) OrderByIDDesc() RfprofileQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// DistinctOnID marks field in queries to add to DISTINCT ON clause
func (qs RfprofileQS) DistinctOnID() RfprofileQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"id"`)

	return qs
}

// END - center.RFProfile.id

// BEGIN - center.RFProfile.name

// NameEq filters for Name being equal to argument
func (qs RfprofileQS) NameEq(v string) RfprofileQS {
	return qs.filter(`"name" =`, v)
}

// NameNe filters for Name being not equal to argument
func (qs RfprofileQS) NameNe(v string) RfprofileQS {
	return qs.filter(`"name" <>`, v)
}

// NameLt filters for Name being less than argument
func (qs RfprofileQS) NameLt(v string) RfprofileQS {
	return qs.filter(`"name" <`, v)
}

// NameLe filters for Name being less than or equal to argument
func (qs RfprofileQS) NameLe(v string) RfprofileQS {
	return qs.filter(`"name" <=`, v)
}

// NameGt filters for Name being greater than argument
func (qs RfprofileQS) NameGt(v string) RfprofileQS {
	return qs.filter(`"name" >`, v)
}

// NameGe filters for Name being greater than or equal to argument
func (qs RfprofileQS) NameGe(v string) RfprofileQS {
	return qs.filter(`"name" >=`, v)
}

type inRfprofileName []interface{}

func (in inRfprofileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfprofileQS) NameIn(values []string) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inRfprofileName(vals),
	)

	return qs
}

type notinRfprofileName []interface{}

func (in notinRfprofileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfprofileQS) NameNotIn(values []string) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinRfprofileName(vals),
	)

	return qs
}

// OrderByName sorts result by Name in ascending order
func (qs RfprofileQS) OrderByName() RfprofileQS {
	qs.order = append(qs.order, `"name"`)

	return qs
}

// OrderByNameDesc sorts result by Name in descending order
func (qs RfprofileQS) OrderByNameDesc() RfprofileQS {
	qs.order = append(qs.order, `"name" DESC`)

	return qs
}

// DistinctOnName marks field in queries to add to DISTINCT ON clause
func (qs RfprofileQS) DistinctOnName() RfprofileQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"name"`)

	return qs
}

// END - center.RFProfile.name

// BEGIN - center.RFProfile.confregs

// ConfregsEq filters for Confregs being equal to argument
func (qs RfprofileQS) ConfregsEq(v string) RfprofileQS {
	return qs.filter(`"confregs" =`, v)
}

// ConfregsNe filters for Confregs being not equal to argument
func (qs RfprofileQS) ConfregsNe(v string) RfprofileQS {
	return qs.filter(`"confregs" <>`, v)
}

// ConfregsLt filters for Confregs being less than argument
func (qs RfprofileQS) ConfregsLt(v string) RfprofileQS {
	return qs.filter(`"confregs" <`, v)
}

// ConfregsLe filters for Confregs being less than or equal to argument
func (qs RfprofileQS) ConfregsLe(v string) RfprofileQS {
	return qs.filter(`"confregs" <=`, v)
}

// ConfregsGt filters for Confregs being greater than argument
func (qs RfprofileQS) ConfregsGt(v string) RfprofileQS {
	return qs.filter(`"confregs" >`, v)
}

// ConfregsGe filters for Confregs being greater than or equal to argument
func (qs RfprofileQS) ConfregsGe(v string) RfprofileQS {
	return qs.filter(`"confregs" >=`, v)
}

type inRfprofileConfregs []interface{}

func (in inRfprofileConfregs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"confregs" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfprofileQS) ConfregsIn(values []string) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inRfprofileConfregs(vals),
	)

	return qs
}

type notinRfprofileConfregs []interface{}

func (in notinRfprofileConfregs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"confregs" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfprofileQS) ConfregsNotIn(values []string) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinRfprofileConfregs(vals),
	)

	return qs
}

// OrderByConfregs sorts result by Confregs in ascending order
func (qs RfprofileQS) OrderByConfregs() RfprofileQS {
	qs.order = append(qs.order, `"confregs"`)

	return qs
}

// OrderByConfregsDesc sorts result by Confregs in descending order
func (qs RfprofileQS) OrderByConfregsDesc() RfprofileQS {
	qs.order = append(qs.order, `"confregs" DESC`)

	return qs
}

// DistinctOnConfregs marks field in queries to add to DISTINCT ON clause
func (qs RfprofileQS) DistinctOnConfregs() RfprofileQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"confregs"`)

	return qs
}

// END - center.RFProfile.confregs

// OrderByRandom randomizes result
func (qs RfprofileQS) OrderByRandom() RfprofileQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs RfprofileQS) ForUpdate() RfprofileQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs RfprofileQS) ForUpdateNowait() RfprofileQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs RfprofileQS) ForUpdateSkipLocked() RfprofileQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs RfprofileQS) ClearForUpdate() RfprofileQS {
	qs.forClause = ""

	return qs
}

func (qs RfprofileQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs RfprofileQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs RfprofileQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"id", "name", "confregs" FROM "center_rfprofile"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs RfprofileQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "center_rfprofile"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs RfprofileQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"id"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "center_rfprofile"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs RfprofileQS) All(ctx context.Context, db models.DBInterface) (RfprofileList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret RfprofileList
	for rows.Next() {
		obj := Rfprofile{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.Name, &obj.Confregs); err != nil {
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
func (qs RfprofileQS) First(ctx context.Context, db models.DBInterface) (*Rfprofile, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Rfprofile{existsInDB: true}
	err := row.Scan(&obj.id, &obj.Name, &obj.Confregs)
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
func (qs RfprofileQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "center_rfprofile"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs RfprofileQS) Update() RfprofileUpdateQS {
	return RfprofileUpdateQS{condFragments: qs.condFragments}
}

// RfprofileUpdateQS represents an updated queryset for center.RFProfile
type RfprofileUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs RfprofileUpdateQS) update(c string, v interface{}) RfprofileUpdateQS {
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
func (uqs RfprofileUpdateQS) SetID(v int32) RfprofileUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetName sets Name to the given value
func (uqs RfprofileUpdateQS) SetName(v string) RfprofileUpdateQS {
	return uqs.update(`"name"`, v)
}

// SetConfregs sets Confregs to the given value
func (uqs RfprofileUpdateQS) SetConfregs(v string) RfprofileUpdateQS {
	return uqs.update(`"confregs"`, v)
}

// Exec executes the update operation
func (uqs RfprofileUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := RfprofileQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "center_rfprofile" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (r *Rfprofile) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "center_rfprofile" ("name", "confregs") VALUES ($1, $2) RETURNING "id"`, r.Name, r.Confregs)

	if err := row.Scan(&r.id); err != nil {
		return err
	}

	r.existsInDB = true

	return nil
}

// update operation
func (r *Rfprofile) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "center_rfprofile" SET "name" = $1, "confregs" = $2 WHERE "id" = $3`, r.Name, r.Confregs, r.id)

	return err
}

// Save inserts or updates record
func (r *Rfprofile) Save(ctx context.Context, db models.DBInterface) error {
	if r.existsInDB {
		return r.update(ctx, db)
	}

	return r.insert(ctx, db)
}

// Delete removes row from database
func (r *Rfprofile) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "center_rfprofile" WHERE "id" = $1`, r.id)

	r.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (rl RfprofileList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts RfprofileList

	for _, r := range rl {
		if r.existsInDB {
			if err := r.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, r)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 2*len(inserts))
	offs := 1
	for _, r := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d)", offs+0, offs+1))
		vaa = append(vaa, r.Name, r.Confregs)
		offs += 2
	}

	qs := `INSERT INTO "center_rfprofile" ("name", "confregs") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, r := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&r.id); err != nil {
			return err
		}

		r.existsInDB = true
	}

	return nil
}

// Rfconfig returns the set of Rfconfig referencing this Rfprofile instance
func (r *Rfprofile) Rfconfig() RfconfigQS {
	return RfconfigQS{}.RfProfileEq(r)
}
