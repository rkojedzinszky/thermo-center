// AUTO-GENERATED file for Django model heatcontrol.ScheduledOverride

package heatcontrol

import (
	"database/sql"
	"fmt"
	"github.com/rkojedzinszky/thermo-center/models"
	"strings"
	"time"
)

// Scheduledoverride mirrors model heatcontrol.ScheduledOverride
type Scheduledoverride struct {
	existsInDB bool

	id         int32
	control    int32
	Start      time.Time
	End        time.Time
	TargetTemp float64
}

// ScheduledoverrideQS represents a queryset for heatcontrol.ScheduledOverride
type ScheduledoverrideQS struct {
	condFragments []models.ConditionFragment
	order         []string
	forUpdate     bool
}

func (qs ScheduledoverrideQS) filter(c string, p interface{}) ScheduledoverrideQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.UnaryFragment{
			Frag:  c,
			Param: p,
		},
	)
	return qs
}

// GetID returns Scheduledoverride.ID
func (s *Scheduledoverride) GetID() int32 {
	return s.id
}

// IDEq filters for id being equal to argument
func (qs ScheduledoverrideQS) IDEq(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs ScheduledoverrideQS) IDNe(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs ScheduledoverrideQS) IDLt(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs ScheduledoverrideQS) IDLe(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs ScheduledoverrideQS) IDGt(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs ScheduledoverrideQS) IDGe(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" >=`, v)
}

type inScheduledoverrideid struct {
	values []interface{}
}

func (in *inScheduledoverrideid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ScheduledoverrideQS) IDIn(values []int32) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inScheduledoverrideid{
			values: vals,
		},
	)

	return qs
}

type notinScheduledoverrideid struct {
	values []interface{}
}

func (in *notinScheduledoverrideid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ScheduledoverrideQS) IDNotIn(values []int32) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinScheduledoverrideid{
			values: vals,
		},
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs ScheduledoverrideQS) OrderByID() ScheduledoverrideQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs ScheduledoverrideQS) OrderByIDDesc() ScheduledoverrideQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetControl returns Control
func (s *Scheduledoverride) GetControl(db models.DBInterface) (*Control, error) {
	return ControlQS{}.IDEq(s.control).First(db)
}

// SetControl sets foreign key pointer to Control
func (s *Scheduledoverride) SetControl(ptr *Control) error {
	if ptr != nil {
		s.control = ptr.GetID()
	} else {
		return fmt.Errorf("Scheduledoverride.SetControl: non-null field received null value")
	}

	return nil
}

// GetControlRaw returns Scheduledoverride.Control
func (s *Scheduledoverride) GetControlRaw() int32 {
	return s.control
}

// ControlEq filters for control being equal to argument
func (qs ScheduledoverrideQS) ControlEq(v *Control) ScheduledoverrideQS {
	return qs.filter(`"control_id" =`, v.GetID())
}

type inScheduledoverridecontrolControl struct {
	qs ControlQS
}

func (in *inScheduledoverridecontrolControl) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"control_id" IN (` + s + `)`, p
}

func (qs ScheduledoverrideQS) ControlIn(oqs ControlQS) ScheduledoverrideQS {
	qs.condFragments = append(
		qs.condFragments,
		&inScheduledoverridecontrolControl{
			qs: oqs,
		},
	)

	return qs
}

// OrderByControl sorts result by Control in ascending order
func (qs ScheduledoverrideQS) OrderByControl() ScheduledoverrideQS {
	qs.order = append(qs.order, `"control_id"`)

	return qs
}

// OrderByControlDesc sorts result by Control in descending order
func (qs ScheduledoverrideQS) OrderByControlDesc() ScheduledoverrideQS {
	qs.order = append(qs.order, `"control_id" DESC`)

	return qs
}

// StartEq filters for Start being equal to argument
func (qs ScheduledoverrideQS) StartEq(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"start" =`, v)
}

// StartNe filters for Start being not equal to argument
func (qs ScheduledoverrideQS) StartNe(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"start" <>`, v)
}

// StartLt filters for Start being less than argument
func (qs ScheduledoverrideQS) StartLt(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"start" <`, v)
}

// StartLe filters for Start being less than or equal to argument
func (qs ScheduledoverrideQS) StartLe(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"start" <=`, v)
}

// StartGt filters for Start being greater than argument
func (qs ScheduledoverrideQS) StartGt(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"start" >`, v)
}

// StartGe filters for Start being greater than or equal to argument
func (qs ScheduledoverrideQS) StartGe(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"start" >=`, v)
}

type inScheduledoverrideStart struct {
	values []interface{}
}

func (in *inScheduledoverrideStart) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"start" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ScheduledoverrideQS) StartIn(values []time.Time) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inScheduledoverrideStart{
			values: vals,
		},
	)

	return qs
}

type notinScheduledoverrideStart struct {
	values []interface{}
}

func (in *notinScheduledoverrideStart) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"start" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ScheduledoverrideQS) StartNotIn(values []time.Time) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinScheduledoverrideStart{
			values: vals,
		},
	)

	return qs
}

// OrderByStart sorts result by Start in ascending order
func (qs ScheduledoverrideQS) OrderByStart() ScheduledoverrideQS {
	qs.order = append(qs.order, `"start"`)

	return qs
}

// OrderByStartDesc sorts result by Start in descending order
func (qs ScheduledoverrideQS) OrderByStartDesc() ScheduledoverrideQS {
	qs.order = append(qs.order, `"start" DESC`)

	return qs
}

// EndEq filters for End being equal to argument
func (qs ScheduledoverrideQS) EndEq(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"end" =`, v)
}

// EndNe filters for End being not equal to argument
func (qs ScheduledoverrideQS) EndNe(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"end" <>`, v)
}

// EndLt filters for End being less than argument
func (qs ScheduledoverrideQS) EndLt(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"end" <`, v)
}

// EndLe filters for End being less than or equal to argument
func (qs ScheduledoverrideQS) EndLe(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"end" <=`, v)
}

// EndGt filters for End being greater than argument
func (qs ScheduledoverrideQS) EndGt(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"end" >`, v)
}

// EndGe filters for End being greater than or equal to argument
func (qs ScheduledoverrideQS) EndGe(v time.Time) ScheduledoverrideQS {
	return qs.filter(`"end" >=`, v)
}

type inScheduledoverrideEnd struct {
	values []interface{}
}

func (in *inScheduledoverrideEnd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"end" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ScheduledoverrideQS) EndIn(values []time.Time) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inScheduledoverrideEnd{
			values: vals,
		},
	)

	return qs
}

type notinScheduledoverrideEnd struct {
	values []interface{}
}

func (in *notinScheduledoverrideEnd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"end" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ScheduledoverrideQS) EndNotIn(values []time.Time) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinScheduledoverrideEnd{
			values: vals,
		},
	)

	return qs
}

// OrderByEnd sorts result by End in ascending order
func (qs ScheduledoverrideQS) OrderByEnd() ScheduledoverrideQS {
	qs.order = append(qs.order, `"end"`)

	return qs
}

// OrderByEndDesc sorts result by End in descending order
func (qs ScheduledoverrideQS) OrderByEndDesc() ScheduledoverrideQS {
	qs.order = append(qs.order, `"end" DESC`)

	return qs
}

// TargetTempEq filters for TargetTemp being equal to argument
func (qs ScheduledoverrideQS) TargetTempEq(v float64) ScheduledoverrideQS {
	return qs.filter(`"target_temp" =`, v)
}

// TargetTempNe filters for TargetTemp being not equal to argument
func (qs ScheduledoverrideQS) TargetTempNe(v float64) ScheduledoverrideQS {
	return qs.filter(`"target_temp" <>`, v)
}

// TargetTempLt filters for TargetTemp being less than argument
func (qs ScheduledoverrideQS) TargetTempLt(v float64) ScheduledoverrideQS {
	return qs.filter(`"target_temp" <`, v)
}

// TargetTempLe filters for TargetTemp being less than or equal to argument
func (qs ScheduledoverrideQS) TargetTempLe(v float64) ScheduledoverrideQS {
	return qs.filter(`"target_temp" <=`, v)
}

// TargetTempGt filters for TargetTemp being greater than argument
func (qs ScheduledoverrideQS) TargetTempGt(v float64) ScheduledoverrideQS {
	return qs.filter(`"target_temp" >`, v)
}

// TargetTempGe filters for TargetTemp being greater than or equal to argument
func (qs ScheduledoverrideQS) TargetTempGe(v float64) ScheduledoverrideQS {
	return qs.filter(`"target_temp" >=`, v)
}

type inScheduledoverrideTargetTemp struct {
	values []interface{}
}

func (in *inScheduledoverrideTargetTemp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"target_temp" IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ScheduledoverrideQS) TargetTempIn(values []float64) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&inScheduledoverrideTargetTemp{
			values: vals,
		},
	)

	return qs
}

type notinScheduledoverrideTargetTemp struct {
	values []interface{}
}

func (in *notinScheduledoverrideTargetTemp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in.values) == 0 {
		return `false`, nil
	}

	var params []string
	for range in.values {
		params = append(params, c.Get())
	}

	return `"target_temp" NOT IN (` + strings.Join(params, ", ") + `)`, in.values
}

func (qs ScheduledoverrideQS) TargetTempNotIn(values []float64) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		&notinScheduledoverrideTargetTemp{
			values: vals,
		},
	)

	return qs
}

// OrderByTargetTemp sorts result by TargetTemp in ascending order
func (qs ScheduledoverrideQS) OrderByTargetTemp() ScheduledoverrideQS {
	qs.order = append(qs.order, `"target_temp"`)

	return qs
}

// OrderByTargetTempDesc sorts result by TargetTemp in descending order
func (qs ScheduledoverrideQS) OrderByTargetTempDesc() ScheduledoverrideQS {
	qs.order = append(qs.order, `"target_temp" DESC`)

	return qs
}

func (qs ScheduledoverrideQS) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
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
func (qs ScheduledoverrideQS) ForUpdate() ScheduledoverrideQS {
	qs.forUpdate = true

	return qs
}

func (qs ScheduledoverrideQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs ScheduledoverrideQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs ScheduledoverrideQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	if qs.forUpdate {
		s += " FOR UPDATE"
	}

	return `SELECT "id", "control_id", "start", "end", "target_temp" FROM "heatcontrol_scheduledoverride"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs ScheduledoverrideQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_scheduledoverride"` + s, p
}

// All returns all rows matching queryset filters
func (qs ScheduledoverrideQS) All(db models.DBInterface) ([]*Scheduledoverride, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []*Scheduledoverride
	for rows.Next() {
		obj := Scheduledoverride{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.control, &obj.Start, &obj.End, &obj.TargetTemp); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs ScheduledoverrideQS) First(db models.DBInterface) (*Scheduledoverride, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(s, p...)

	obj := Scheduledoverride{existsInDB: true}
	err := row.Scan(&obj.id, &obj.control, &obj.Start, &obj.End, &obj.TargetTemp)
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
func (qs ScheduledoverrideQS) Delete(db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_scheduledoverride"` + s

	result, err := db.Exec(s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs ScheduledoverrideQS) Update() ScheduledoverrideUpdateQS {
	return ScheduledoverrideUpdateQS{condFragments: qs.condFragments}
}

// ScheduledoverrideUpdateQS represents an updated queryset for heatcontrol.ScheduledOverride
type ScheduledoverrideUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs ScheduledoverrideUpdateQS) update(c string, v interface{}) ScheduledoverrideUpdateQS {
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
func (uqs ScheduledoverrideUpdateQS) SetID(v int32) ScheduledoverrideUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetControl sets foreign key pointer to Control
func (uqs ScheduledoverrideUpdateQS) SetControl(ptr *Control) ScheduledoverrideUpdateQS {
	if ptr != nil {
		return uqs.update(`"control_id"`, ptr.GetID())
	}

	return uqs.update(`"control_id"`, nil)
} // SetStart sets Start to the given value
func (uqs ScheduledoverrideUpdateQS) SetStart(v time.Time) ScheduledoverrideUpdateQS {
	return uqs.update(`"start"`, v)
}

// SetEnd sets End to the given value
func (uqs ScheduledoverrideUpdateQS) SetEnd(v time.Time) ScheduledoverrideUpdateQS {
	return uqs.update(`"end"`, v)
}

// SetTargetTemp sets TargetTemp to the given value
func (uqs ScheduledoverrideUpdateQS) SetTargetTemp(v float64) ScheduledoverrideUpdateQS {
	return uqs.update(`"target_temp"`, v)
}

// Exec executes the update operation
func (uqs ScheduledoverrideUpdateQS) Exec(db models.DBInterface) (int64, error) {
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

	ws, wp := ScheduledoverrideQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "heatcontrol_scheduledoverride" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// insert operation
func (s *Scheduledoverride) insert(db models.DBInterface) error {
	row := db.QueryRow(`INSERT INTO "heatcontrol_scheduledoverride" ("control_id", "start", "end", "target_temp") VALUES ($1, $2, $3, $4) RETURNING "id"`, s.control, s.Start, s.End, s.TargetTemp)

	if err := row.Scan(&s.id); err != nil {
		return err
	}

	s.existsInDB = true

	return nil
}

// update operation
func (s *Scheduledoverride) update(db models.DBInterface) error {
	_, err := db.Exec(`UPDATE "heatcontrol_scheduledoverride" SET "control_id" = $1, "start" = $2, "end" = $3, "target_temp" = $4 WHERE "id" = $5`, s.control, s.Start, s.End, s.TargetTemp, s.id)

	return err
}

// Save inserts or updates record
func (s *Scheduledoverride) Save(db models.DBInterface) error {
	if s.existsInDB {
		return s.update(db)
	}

	return s.insert(db)
}

// Delete removes row from database
func (s *Scheduledoverride) Delete(db models.DBInterface) error {
	_, err := db.Exec(`DELETE FROM "heatcontrol_scheduledoverride" WHERE "id" = $1`, s.id)

	s.existsInDB = false

	return err
}
