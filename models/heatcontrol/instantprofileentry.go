// AUTO-GENERATED file for Django model heatcontrol.InstantProfileEntry

package heatcontrol

import (
	"database/sql"
	"fmt"
	"github.com/rkojedzinszky/thermo-center/models"
	"strings"
)

// Instantprofileentry mirrors model heatcontrol.InstantProfileEntry
type Instantprofileentry struct {
	existsInDB bool

	id         int32
	profile    int32
	control    int32
	TargetTemp sql.NullFloat64
	Active     bool
}

// InstantprofileentryQS represents a queryset for heatcontrol.InstantProfileEntry
type InstantprofileentryQS struct {
	condFragments []models.ConditionFragment
	order         []string
	forUpdate     bool
}

func (qs InstantprofileentryQS) filter(c string, p interface{}) InstantprofileentryQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.UnaryFragment{
			Frag:  c,
			Param: p,
		},
	)
	return qs
}

// GetId returns Instantprofileentry.Id
func (i *Instantprofileentry) GetId() int32 {
	return i.id
}

// IdEq filters for id being equal to argument
func (qs InstantprofileentryQS) IdEq(v int32) InstantprofileentryQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for id being not equal to argument
func (qs InstantprofileentryQS) IdNe(v int32) InstantprofileentryQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for id being less than argument
func (qs InstantprofileentryQS) IdLt(v int32) InstantprofileentryQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for id being less than or equal to argument
func (qs InstantprofileentryQS) IdLe(v int32) InstantprofileentryQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for id being greater than argument
func (qs InstantprofileentryQS) IdGt(v int32) InstantprofileentryQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for id being greater than or equal to argument
func (qs InstantprofileentryQS) IdGe(v int32) InstantprofileentryQS {
	return qs.filter(`"id" >=`, v)
}

type inInstantprofileentryid struct {
	values []interface{}
}

func (in *inInstantprofileentryid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileentryQS) IdIn(values []int32) InstantprofileentryQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inInstantprofileentryid{
			values: vals,
		},
	)

	return qs
}

type notinInstantprofileentryid struct {
	values []interface{}
}

func (in *notinInstantprofileentryid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileentryQS) IdNotIn(values []int32) InstantprofileentryQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinInstantprofileentryid{
			values: vals,
		},
	)

	return qs
}

// OrderById sorts result by Id in ascending order
func (qs InstantprofileentryQS) OrderById() InstantprofileentryQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs InstantprofileentryQS) OrderByIdDesc() InstantprofileentryQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetProfile returns Instantprofile
func (i *Instantprofileentry) GetProfile(db models.DBInterface) (*Instantprofile, error) {
	return InstantprofileQS{}.IdEq(i.profile).First(db)
}

// SetProfile sets foreign key pointer to Instantprofile
func (i *Instantprofileentry) SetProfile(ptr *Instantprofile) error {
	if ptr != nil {
		i.profile = ptr.GetId()
	} else {
		return fmt.Errorf("Instantprofileentry.SetProfile: non-null field received null value")
	}

	return nil
}

// GetProfileRaw returns Instantprofileentry.Profile
func (i *Instantprofileentry) GetProfileRaw() int32 {
	return i.profile
}

// ProfileEq filters for profile being equal to argument
func (qs InstantprofileentryQS) ProfileEq(v *Instantprofile) InstantprofileentryQS {
	return qs.filter(`"profile_id" =`, v.GetId())
}

type inInstantprofileentryprofileInstantprofile struct {
	qs InstantprofileQS
}

func (in *inInstantprofileentryprofileInstantprofile) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"profile_id" IN (` + s + `)`, p
}

func (qs InstantprofileentryQS) ProfileIn(oqs InstantprofileQS) InstantprofileentryQS {
	qs.condFragments = append(
		qs.condFragments,
		&inInstantprofileentryprofileInstantprofile{
			qs: oqs,
		},
	)

	return qs
}

// OrderByProfile sorts result by Profile in ascending order
func (qs InstantprofileentryQS) OrderByProfile() InstantprofileentryQS {
	qs.order = append(qs.order, `"profile_id"`)

	return qs
}

// OrderByProfileDesc sorts result by Profile in descending order
func (qs InstantprofileentryQS) OrderByProfileDesc() InstantprofileentryQS {
	qs.order = append(qs.order, `"profile_id" DESC`)

	return qs
}

// GetControl returns Control
func (i *Instantprofileentry) GetControl(db models.DBInterface) (*Control, error) {
	return ControlQS{}.IdEq(i.control).First(db)
}

// SetControl sets foreign key pointer to Control
func (i *Instantprofileentry) SetControl(ptr *Control) error {
	if ptr != nil {
		i.control = ptr.GetId()
	} else {
		return fmt.Errorf("Instantprofileentry.SetControl: non-null field received null value")
	}

	return nil
}

// GetControlRaw returns Instantprofileentry.Control
func (i *Instantprofileentry) GetControlRaw() int32 {
	return i.control
}

// ControlEq filters for control being equal to argument
func (qs InstantprofileentryQS) ControlEq(v *Control) InstantprofileentryQS {
	return qs.filter(`"control_id" =`, v.GetId())
}

type inInstantprofileentrycontrolControl struct {
	qs ControlQS
}

func (in *inInstantprofileentrycontrolControl) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"control_id" IN (` + s + `)`, p
}

func (qs InstantprofileentryQS) ControlIn(oqs ControlQS) InstantprofileentryQS {
	qs.condFragments = append(
		qs.condFragments,
		&inInstantprofileentrycontrolControl{
			qs: oqs,
		},
	)

	return qs
}

// OrderByControl sorts result by Control in ascending order
func (qs InstantprofileentryQS) OrderByControl() InstantprofileentryQS {
	qs.order = append(qs.order, `"control_id"`)

	return qs
}

// OrderByControlDesc sorts result by Control in descending order
func (qs InstantprofileentryQS) OrderByControlDesc() InstantprofileentryQS {
	qs.order = append(qs.order, `"control_id" DESC`)

	return qs
}

// TargetTempIsNull filters for TargetTemp being null
func (qs InstantprofileentryQS) TargetTempIsNull() InstantprofileentryQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"target_temp" IS NULL`,
		},
	)
	return qs
}

// TargetTempIsNotNull filters for TargetTemp being not null
func (qs InstantprofileentryQS) TargetTempIsNotNull() InstantprofileentryQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"target_temp" IS NOT NULL`,
		},
	)
	return qs
}

// TargetTempEq filters for TargetTemp being equal to argument
func (qs InstantprofileentryQS) TargetTempEq(v float64) InstantprofileentryQS {
	return qs.filter(`"target_temp" =`, v)
}

// TargetTempNe filters for TargetTemp being not equal to argument
func (qs InstantprofileentryQS) TargetTempNe(v float64) InstantprofileentryQS {
	return qs.filter(`"target_temp" <>`, v)
}

// TargetTempLt filters for TargetTemp being less than argument
func (qs InstantprofileentryQS) TargetTempLt(v float64) InstantprofileentryQS {
	return qs.filter(`"target_temp" <`, v)
}

// TargetTempLe filters for TargetTemp being less than or equal to argument
func (qs InstantprofileentryQS) TargetTempLe(v float64) InstantprofileentryQS {
	return qs.filter(`"target_temp" <=`, v)
}

// TargetTempGt filters for TargetTemp being greater than argument
func (qs InstantprofileentryQS) TargetTempGt(v float64) InstantprofileentryQS {
	return qs.filter(`"target_temp" >`, v)
}

// TargetTempGe filters for TargetTemp being greater than or equal to argument
func (qs InstantprofileentryQS) TargetTempGe(v float64) InstantprofileentryQS {
	return qs.filter(`"target_temp" >=`, v)
}

type inInstantprofileentryTargetTemp struct {
	values []interface{}
}

func (in *inInstantprofileentryTargetTemp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"target_temp" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileentryQS) TargetTempIn(values []float64) InstantprofileentryQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inInstantprofileentryTargetTemp{
			values: vals,
		},
	)

	return qs
}

type notinInstantprofileentryTargetTemp struct {
	values []interface{}
}

func (in *notinInstantprofileentryTargetTemp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"target_temp" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileentryQS) TargetTempNotIn(values []float64) InstantprofileentryQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinInstantprofileentryTargetTemp{
			values: vals,
		},
	)

	return qs
}

// OrderByTargetTemp sorts result by TargetTemp in ascending order
func (qs InstantprofileentryQS) OrderByTargetTemp() InstantprofileentryQS {
	qs.order = append(qs.order, `"target_temp"`)

	return qs
}

// OrderByTargetTempDesc sorts result by TargetTemp in descending order
func (qs InstantprofileentryQS) OrderByTargetTempDesc() InstantprofileentryQS {
	qs.order = append(qs.order, `"target_temp" DESC`)

	return qs
}

// ActiveEq filters for Active being equal to argument
func (qs InstantprofileentryQS) ActiveEq(v bool) InstantprofileentryQS {
	return qs.filter(`"active" =`, v)
}

// ActiveNe filters for Active being not equal to argument
func (qs InstantprofileentryQS) ActiveNe(v bool) InstantprofileentryQS {
	return qs.filter(`"active" <>`, v)
}

// ActiveLt filters for Active being less than argument
func (qs InstantprofileentryQS) ActiveLt(v bool) InstantprofileentryQS {
	return qs.filter(`"active" <`, v)
}

// ActiveLe filters for Active being less than or equal to argument
func (qs InstantprofileentryQS) ActiveLe(v bool) InstantprofileentryQS {
	return qs.filter(`"active" <=`, v)
}

// ActiveGt filters for Active being greater than argument
func (qs InstantprofileentryQS) ActiveGt(v bool) InstantprofileentryQS {
	return qs.filter(`"active" >`, v)
}

// ActiveGe filters for Active being greater than or equal to argument
func (qs InstantprofileentryQS) ActiveGe(v bool) InstantprofileentryQS {
	return qs.filter(`"active" >=`, v)
}

type inInstantprofileentryActive struct {
	values []interface{}
}

func (in *inInstantprofileentryActive) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"active" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileentryQS) ActiveIn(values []bool) InstantprofileentryQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inInstantprofileentryActive{
			values: vals,
		},
	)

	return qs
}

type notinInstantprofileentryActive struct {
	values []interface{}
}

func (in *notinInstantprofileentryActive) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"active" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs InstantprofileentryQS) ActiveNotIn(values []bool) InstantprofileentryQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinInstantprofileentryActive{
			values: vals,
		},
	)

	return qs
}

// OrderByActive sorts result by Active in ascending order
func (qs InstantprofileentryQS) OrderByActive() InstantprofileentryQS {
	qs.order = append(qs.order, `"active"`)

	return qs
}

// OrderByActiveDesc sorts result by Active in descending order
func (qs InstantprofileentryQS) OrderByActiveDesc() InstantprofileentryQS {
	qs.order = append(qs.order, `"active" DESC`)

	return qs
}

func (qs InstantprofileentryQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
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
func (qs InstantprofileentryQS) ForUpdate() InstantprofileentryQS {
	qs.forUpdate = true

	return qs
}

func (qs InstantprofileentryQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs InstantprofileentryQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs InstantprofileentryQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "profile_id", "control_id", "target_temp", "active" FROM "heatcontrol_instantprofileentry"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs InstantprofileentryQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_instantprofileentry"` + s, p
}

// All returns all rows matching queryset filters
func (qs InstantprofileentryQS) All(db models.DBInterface) ([]*Instantprofileentry, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Instantprofileentry
	for rows.Next() {
		obj := Instantprofileentry{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.profile, &obj.control, &obj.TargetTemp, &obj.Active); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs InstantprofileentryQS) First(db models.DBInterface) (*Instantprofileentry, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Instantprofileentry{existsInDB: true}
	err := row.Scan(&obj.id, &obj.profile, &obj.control, &obj.TargetTemp, &obj.Active)
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
func (i *Instantprofileentry) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "heatcontrol_instantprofileentry" ("profile_id", "control_id", "target_temp", "active") VALUES ($1, $2, $3, $4) RETURNING "id"`, i.profile, i.control, i.TargetTemp, i.Active)

	if err := row.Scan(&i.id); err != nil {
		return err
	}

	i.existsInDB = true

	return nil
}

// update operation
func (i *Instantprofileentry) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "heatcontrol_instantprofileentry" SET "profile_id" = $1, "control_id" = $2, "target_temp" = $3, "active" = $4 WHERE "id" = $5`, i.profile, i.control, i.TargetTemp, i.Active, i.id)

	return err
}

// Save inserts or updates record
func (i *Instantprofileentry) Save(db models.DBInterface) error {
	if i.existsInDB {
		return i.update(db)
	}

	return i.insert(db)
}

// Delete removes row from database
func (i *Instantprofileentry) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "heatcontrol_instantprofileentry" WHERE "id" = $1`, i.id)

	i.existsInDB = false

	return err
}
