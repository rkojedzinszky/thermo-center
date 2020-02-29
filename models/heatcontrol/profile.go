// AUTO-GENERATED file for Django model heatcontrol.Profile

package heatcontrol

import (
	"database/sql"
	"fmt"
	"github.com/rkojedzinszky/thermo-center/models"
	"strings"
	"time"
)

// Profile mirrors model heatcontrol.Profile
type Profile struct {
	existsInDB bool

	id         int32
	control    int32
	daytype    int32
	Start      time.Time
	TargetTemp sql.NullFloat64
}

// ProfileQS represents a queryset for heatcontrol.Profile
type ProfileQS struct {
	condFragments []models.ConditionFragment
	order         []string
	forUpdate     bool
}

func (qs ProfileQS) filter(c string, p interface{}) ProfileQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.UnaryFragment{
			Frag:  c,
			Param: p,
		},
	)
	return qs
}

// GetId returns Profile.Id
func (p *Profile) GetId() int32 {
	return p.id
}

// IdEq filters for id being equal to argument
func (qs ProfileQS) IdEq(v int32) ProfileQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for id being not equal to argument
func (qs ProfileQS) IdNe(v int32) ProfileQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for id being less than argument
func (qs ProfileQS) IdLt(v int32) ProfileQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for id being less than or equal to argument
func (qs ProfileQS) IdLe(v int32) ProfileQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for id being greater than argument
func (qs ProfileQS) IdGt(v int32) ProfileQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for id being greater than or equal to argument
func (qs ProfileQS) IdGe(v int32) ProfileQS {
	return qs.filter(`"id" >=`, v)
}

// OrderById sorts result by Id in ascending order
func (qs ProfileQS) OrderById() ProfileQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs ProfileQS) OrderByIdDesc() ProfileQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetControl returns Control
func (p *Profile) GetControl(db models.DBInterface) (*Control, error) {
	return ControlQS{}.IdEq(p.control).First(db)
}

// SetControl sets foreign key pointer to Control
func (p *Profile) SetControl(ptr *Control) error {
	if ptr != nil {
		p.control = ptr.GetId()
	} else {
		return fmt.Errorf("Profile.SetControl: non-null field received null value")
	}

	return nil
}

// GetControlRaw returns Profile.Control
func (p *Profile) GetControlRaw() int32 {
	return p.control
}

// ControlEq filters for control being equal to argument
func (qs ProfileQS) ControlEq(v *Control) ProfileQS {
	return qs.filter(`"control_id" =`, v.GetId())
}

type inProfilecontrolControl struct {
	qs ControlQS
}

func (in *inProfilecontrolControl) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"control_id" IN (` + s + `)`, p
}

func (qs ProfileQS) ControlIn(oqs ControlQS) ProfileQS {
	qs.condFragments = append(
		qs.condFragments,
		&inProfilecontrolControl{
			qs: oqs,
		},
	)

	return qs
}

// OrderByControl sorts result by Control in ascending order
func (qs ProfileQS) OrderByControl() ProfileQS {
	qs.order = append(qs.order, `"control_id"`)

	return qs
}

// OrderByControlDesc sorts result by Control in descending order
func (qs ProfileQS) OrderByControlDesc() ProfileQS {
	qs.order = append(qs.order, `"control_id" DESC`)

	return qs
}

// GetDaytype returns Daytype
func (p *Profile) GetDaytype(db models.DBInterface) (*Daytype, error) {
	return DaytypeQS{}.IdEq(p.daytype).First(db)
}

// SetDaytype sets foreign key pointer to Daytype
func (p *Profile) SetDaytype(ptr *Daytype) error {
	if ptr != nil {
		p.daytype = ptr.GetId()
	} else {
		return fmt.Errorf("Profile.SetDaytype: non-null field received null value")
	}

	return nil
}

// GetDaytypeRaw returns Profile.Daytype
func (p *Profile) GetDaytypeRaw() int32 {
	return p.daytype
}

// DaytypeEq filters for daytype being equal to argument
func (qs ProfileQS) DaytypeEq(v *Daytype) ProfileQS {
	return qs.filter(`"daytype_id" =`, v.GetId())
}

type inProfiledaytypeDaytype struct {
	qs DaytypeQS
}

func (in *inProfiledaytypeDaytype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"daytype_id" IN (` + s + `)`, p
}

func (qs ProfileQS) DaytypeIn(oqs DaytypeQS) ProfileQS {
	qs.condFragments = append(
		qs.condFragments,
		&inProfiledaytypeDaytype{
			qs: oqs,
		},
	)

	return qs
}

// OrderByDaytype sorts result by Daytype in ascending order
func (qs ProfileQS) OrderByDaytype() ProfileQS {
	qs.order = append(qs.order, `"daytype_id"`)

	return qs
}

// OrderByDaytypeDesc sorts result by Daytype in descending order
func (qs ProfileQS) OrderByDaytypeDesc() ProfileQS {
	qs.order = append(qs.order, `"daytype_id" DESC`)

	return qs
}

// StartEq filters for Start being equal to argument
func (qs ProfileQS) StartEq(v time.Time) ProfileQS {
	return qs.filter(`"start" =`, v)
}

// StartNe filters for Start being not equal to argument
func (qs ProfileQS) StartNe(v time.Time) ProfileQS {
	return qs.filter(`"start" <>`, v)
}

// StartLt filters for Start being less than argument
func (qs ProfileQS) StartLt(v time.Time) ProfileQS {
	return qs.filter(`"start" <`, v)
}

// StartLe filters for Start being less than or equal to argument
func (qs ProfileQS) StartLe(v time.Time) ProfileQS {
	return qs.filter(`"start" <=`, v)
}

// StartGt filters for Start being greater than argument
func (qs ProfileQS) StartGt(v time.Time) ProfileQS {
	return qs.filter(`"start" >`, v)
}

// StartGe filters for Start being greater than or equal to argument
func (qs ProfileQS) StartGe(v time.Time) ProfileQS {
	return qs.filter(`"start" >=`, v)
}

// OrderByStart sorts result by Start in ascending order
func (qs ProfileQS) OrderByStart() ProfileQS {
	qs.order = append(qs.order, `"start"`)

	return qs
}

// OrderByStartDesc sorts result by Start in descending order
func (qs ProfileQS) OrderByStartDesc() ProfileQS {
	qs.order = append(qs.order, `"start" DESC`)

	return qs
}

// TargetTempIsNull filters for TargetTemp being null
func (qs ProfileQS) TargetTempIsNull() ProfileQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"target_temp" IS NULL`,
		},
	)
	return qs
}

// TargetTempIsNotNull filters for TargetTemp being not null
func (qs ProfileQS) TargetTempIsNotNull() ProfileQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"target_temp" IS NOT NULL`,
		},
	)
	return qs
}

// TargetTempEq filters for TargetTemp being equal to argument
func (qs ProfileQS) TargetTempEq(v float64) ProfileQS {
	return qs.filter(`"target_temp" =`, v)
}

// TargetTempNe filters for TargetTemp being not equal to argument
func (qs ProfileQS) TargetTempNe(v float64) ProfileQS {
	return qs.filter(`"target_temp" <>`, v)
}

// TargetTempLt filters for TargetTemp being less than argument
func (qs ProfileQS) TargetTempLt(v float64) ProfileQS {
	return qs.filter(`"target_temp" <`, v)
}

// TargetTempLe filters for TargetTemp being less than or equal to argument
func (qs ProfileQS) TargetTempLe(v float64) ProfileQS {
	return qs.filter(`"target_temp" <=`, v)
}

// TargetTempGt filters for TargetTemp being greater than argument
func (qs ProfileQS) TargetTempGt(v float64) ProfileQS {
	return qs.filter(`"target_temp" >`, v)
}

// TargetTempGe filters for TargetTemp being greater than or equal to argument
func (qs ProfileQS) TargetTempGe(v float64) ProfileQS {
	return qs.filter(`"target_temp" >=`, v)
}

// OrderByTargetTemp sorts result by TargetTemp in ascending order
func (qs ProfileQS) OrderByTargetTemp() ProfileQS {
	qs.order = append(qs.order, `"target_temp"`)

	return qs
}

// OrderByTargetTempDesc sorts result by TargetTemp in descending order
func (qs ProfileQS) OrderByTargetTempDesc() ProfileQS {
	qs.order = append(qs.order, `"target_temp" DESC`)

	return qs
}

func (qs ProfileQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
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
func (qs ProfileQS) ForUpdate() ProfileQS {
	qs.forUpdate = true

	return qs
}

func (qs ProfileQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs ProfileQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs ProfileQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "control_id", "daytype_id", "start", "target_temp" FROM "heatcontrol_profile"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs ProfileQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_profile"` + s, p
}

// All returns all rows matching queryset filters
func (qs ProfileQS) All(db models.DBInterface) ([]*Profile, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Profile
	for rows.Next() {
		obj := Profile{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.control, &obj.daytype, &obj.Start, &obj.TargetTemp); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs ProfileQS) First(db models.DBInterface) (*Profile, error) {
	s, p := qs.queryFull()

	row := db.QueryRow(s, p...)

	obj := Profile{existsInDB: true}
	err := row.Scan(&obj.id, &obj.control, &obj.daytype, &obj.Start, &obj.TargetTemp)
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
func (p *Profile) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "heatcontrol_profile" ("control_id", "daytype_id", "start", "target_temp") VALUES ($1, $2, $3, $4) RETURNING "id"`, p.control, p.daytype, p.Start, p.TargetTemp)

	if err := row.Scan(&p.id); err != nil {
		return err
	}

	p.existsInDB = true

	return nil
}

// update operation
func (p *Profile) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "heatcontrol_profile" SET "control_id" = $1, "daytype_id" = $2, "start" = $3, "target_temp" = $4 WHERE "id" = $5`, p.control, p.daytype, p.Start, p.TargetTemp, p.id)

	return err
}

// Save inserts or updates record
func (p *Profile) Save(db models.DBInterface) error {
	if p.existsInDB {
		return p.update(db)
	}

	return p.insert(db)
}

// Delete removes row from database
func (p *Profile) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "heatcontrol_profile" WHERE "id" = $1`, p.id)

	p.existsInDB = false

	return err
}
