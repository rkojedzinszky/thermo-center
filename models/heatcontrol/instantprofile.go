// AUTO-GENERATED file for Django model heatcontrol.InstantProfile

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
	condFragments []models.ConditionFragment
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

// GetId returns Instantprofile.Id
func (i *Instantprofile) GetId() int32 {
	return i.id
}

// IdEq filters for id being equal to argument
func (qs InstantprofileQS) IdEq(v int32) InstantprofileQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for id being not equal to argument
func (qs InstantprofileQS) IdNe(v int32) InstantprofileQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for id being less than argument
func (qs InstantprofileQS) IdLt(v int32) InstantprofileQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for id being less than or equal to argument
func (qs InstantprofileQS) IdLe(v int32) InstantprofileQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for id being greater than argument
func (qs InstantprofileQS) IdGt(v int32) InstantprofileQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for id being greater than or equal to argument
func (qs InstantprofileQS) IdGe(v int32) InstantprofileQS {
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

func (qs InstantprofileQS) IdIn(values []int32) InstantprofileQS {
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

func (qs InstantprofileQS) IdNotIn(values []int32) InstantprofileQS {
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

// OrderById sorts result by Id in ascending order
func (qs InstantprofileQS) OrderById() InstantprofileQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs InstantprofileQS) OrderByIdDesc() InstantprofileQS {
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

func (qs InstantprofileQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
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
func (qs InstantprofileQS) ForUpdate() InstantprofileQS {
	qs.forUpdate = true

	return qs
}

func (qs InstantprofileQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

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
