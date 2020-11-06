/*
  AUTO-GENERATED file for Django model center.RFProfile

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

// Rfprofile mirrors model center.RFProfile
type Rfprofile struct {
	existsInDB bool

	id       int32
	Name     string
	Confregs string
}

// RfprofileQS represents a queryset for center.RFProfile
type RfprofileQS struct {
	condFragments []models.ConditionFragment
	order         []string
	forUpdate     bool
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

type inRfprofileid struct {
	values []interface{}
}

func (in *inRfprofileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs RfprofileQS) IDIn(values []int32) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inRfprofileid{
			values: vals,
		},
	)

	return qs
}

type notinRfprofileid struct {
	values []interface{}
}

func (in *notinRfprofileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs RfprofileQS) IDNotIn(values []int32) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinRfprofileid{
			values: vals,
		},
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

type inRfprofileName struct {
	values []interface{}
}

func (in *inRfprofileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs RfprofileQS) NameIn(values []string) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inRfprofileName{
			values: vals,
		},
	)

	return qs
}

type notinRfprofileName struct {
	values []interface{}
}

func (in *notinRfprofileName) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"name" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs RfprofileQS) NameNotIn(values []string) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinRfprofileName{
			values: vals,
		},
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

type inRfprofileConfregs struct {
	values []interface{}
}

func (in *inRfprofileConfregs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"confregs" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs RfprofileQS) ConfregsIn(values []string) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inRfprofileConfregs{
			values: vals,
		},
	)

	return qs
}

type notinRfprofileConfregs struct {
	values []interface{}
}

func (in *notinRfprofileConfregs) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"confregs" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs RfprofileQS) ConfregsNotIn(values []string) RfprofileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinRfprofileConfregs{
			values: vals,
		},
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

func (qs RfprofileQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
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
func (qs RfprofileQS) ForUpdate() RfprofileQS {
	qs.forUpdate = true

	return qs
}

func (qs RfprofileQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs RfprofileQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs RfprofileQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "name", "confregs" FROM "center_rfprofile"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs RfprofileQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "center_rfprofile"` + s, p
}

// All returns all rows matching queryset filters
func (qs RfprofileQS) All(db models.DBInterface) ([]*Rfprofile, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Rfprofile
	for rows.Next() {
		obj := Rfprofile{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.Name, &obj.Confregs); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs RfprofileQS) First(db models.DBInterface) (*Rfprofile, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Rfprofile{existsInDB: true}
	err := row.Scan(&obj.id, &obj.Name, &obj.Confregs)
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
func (qs RfprofileQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "center_rfprofile"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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
func (uqs RfprofileUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (r *Rfprofile) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "center_rfprofile" ("name", "confregs") VALUES ($1, $2) RETURNING "id"`, r.Name, r.Confregs)

	if err := row.Scan(&r.id); err != nil {
		return err
	}

	r.existsInDB = true

	return nil
}

// update operation
func (r *Rfprofile) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "center_rfprofile" SET "name" = $1, "confregs" = $2 WHERE "id" = $3`, r.Name, r.Confregs, r.id)

	return err
}

// Save inserts or updates record
func (r *Rfprofile) Save(db models.DBInterface) error {
	if r.existsInDB {
		return r.update(db)
	}

	return r.insert(db)
}

// Delete removes row from database
func (r *Rfprofile) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "center_rfprofile" WHERE "id" = $1`, r.id)

	r.existsInDB = false

	return err
}

// Rfconfig returns the set of Rfconfig referencing this Rfprofile instance
func (r *Rfprofile) Rfconfig() RfconfigQS {
	return RfconfigQS{}.RfProfileEq(r)
}
