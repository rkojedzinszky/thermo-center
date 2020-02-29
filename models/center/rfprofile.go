// AUTO-GENERATED file for Django model center.RFProfile

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

// GetId returns Rfprofile.Id
func (r *Rfprofile) GetId() int32 {
	return r.id
}

// IdEq filters for id being equal to argument
func (qs RfprofileQS) IdEq(v int32) RfprofileQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for id being not equal to argument
func (qs RfprofileQS) IdNe(v int32) RfprofileQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for id being less than argument
func (qs RfprofileQS) IdLt(v int32) RfprofileQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for id being less than or equal to argument
func (qs RfprofileQS) IdLe(v int32) RfprofileQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for id being greater than argument
func (qs RfprofileQS) IdGt(v int32) RfprofileQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for id being greater than or equal to argument
func (qs RfprofileQS) IdGe(v int32) RfprofileQS {
	return qs.filter(`"id" >=`, v)
}

// OrderById sorts result by Id in ascending order
func (qs RfprofileQS) OrderById() RfprofileQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs RfprofileQS) OrderByIdDesc() RfprofileQS {
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
