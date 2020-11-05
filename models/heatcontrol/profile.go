/*
  AUTO-GENERATED file for Django model heatcontrol.Profile

  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

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

// GetID returns Profile.ID
func (p *Profile) GetID() int32 {
	return p.id
}

// IDEq filters for id being equal to argument
func (qs ProfileQS) IDEq(v int32) ProfileQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs ProfileQS) IDNe(v int32) ProfileQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs ProfileQS) IDLt(v int32) ProfileQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs ProfileQS) IDLe(v int32) ProfileQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs ProfileQS) IDGt(v int32) ProfileQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs ProfileQS) IDGe(v int32) ProfileQS {
	return qs.filter(`"id" >=`, v)
}

type inProfileid struct {
	values []interface{}
}

func (in *inProfileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ProfileQS) IDIn(values []int32) ProfileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inProfileid{
			values: vals,
		},
	)

	return qs
}

type notinProfileid struct {
	values []interface{}
}

func (in *notinProfileid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ProfileQS) IDNotIn(values []int32) ProfileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinProfileid{
			values: vals,
		},
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs ProfileQS) OrderByID() ProfileQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs ProfileQS) OrderByIDDesc() ProfileQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetControl returns Control
func (p *Profile) GetControl(db models.DBInterface) (*Control, error) {
	return ControlQS{}.IDEq(p.control).First(db)
}

// SetControl sets foreign key pointer to Control
func (p *Profile) SetControl(ptr *Control) error {
	if ptr != nil {
		p.control = ptr.GetID()
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
	return qs.filter(`"control_id" =`, v.GetID())
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
	return DaytypeQS{}.IDEq(p.daytype).First(db)
}

// SetDaytype sets foreign key pointer to Daytype
func (p *Profile) SetDaytype(ptr *Daytype) error {
	if ptr != nil {
		p.daytype = ptr.GetID()
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
	return qs.filter(`"daytype_id" =`, v.GetID())
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

type inProfileStart struct {
	values []interface{}
}

func (in *inProfileStart) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"start" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ProfileQS) StartIn(values []time.Time) ProfileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inProfileStart{
			values: vals,
		},
	)

	return qs
}

type notinProfileStart struct {
	values []interface{}
}

func (in *notinProfileStart) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"start" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ProfileQS) StartNotIn(values []time.Time) ProfileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinProfileStart{
			values: vals,
		},
	)

	return qs
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

type inProfileTargetTemp struct {
	values []interface{}
}

func (in *inProfileTargetTemp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"target_temp" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ProfileQS) TargetTempIn(values []float64) ProfileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inProfileTargetTemp{
			values: vals,
		},
	)

	return qs
}

type notinProfileTargetTemp struct {
	values []interface{}
}

func (in *notinProfileTargetTemp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"target_temp" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ProfileQS) TargetTempNotIn(values []float64) ProfileQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinProfileTargetTemp{
			values: vals,
		},
	)

	return qs
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

	s += " LIMIT 1"

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

// Delete deletes rows matching queryset filters
func (qs ProfileQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_profile"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs ProfileQS) Update() ProfileUpdateQS {
	return ProfileUpdateQS{condFragments: qs.condFragments}
}

// ProfileUpdateQS represents an updated queryset for heatcontrol.Profile
type ProfileUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs ProfileUpdateQS) update(c string, v interface{}) ProfileUpdateQS {
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
func (uqs ProfileUpdateQS) SetID(v int32) ProfileUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetControl sets foreign key pointer to Control
func (uqs ProfileUpdateQS) SetControl(ptr *Control) ProfileUpdateQS {
	if ptr != nil {
		return uqs.update(`"control_id"`, ptr.GetID())
	}

	return uqs.update(`"control_id"`, nil)
} // SetDaytype sets foreign key pointer to Daytype
func (uqs ProfileUpdateQS) SetDaytype(ptr *Daytype) ProfileUpdateQS {
	if ptr != nil {
		return uqs.update(`"daytype_id"`, ptr.GetID())
	}

	return uqs.update(`"daytype_id"`, nil)
} // SetStart sets Start to the given value
func (uqs ProfileUpdateQS) SetStart(v time.Time) ProfileUpdateQS {
	return uqs.update(`"start"`, v)
}

// SetTargetTemp sets TargetTemp to the given value
func (uqs ProfileUpdateQS) SetTargetTemp(v sql.NullFloat64) ProfileUpdateQS {
	return uqs.update(`"target_temp"`, v)
}

// Exec executes the update operation
func (uqs ProfileUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := ProfileQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "heatcontrol_profile" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
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
