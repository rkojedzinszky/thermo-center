// Code generated for Django model heatcontrol.ScheduledOverride. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package heatcontrol

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
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

// ScheduledoverrideList is a list of Scheduledoverride
type ScheduledoverrideList []*Scheduledoverride

// ScheduledoverrideQS represents a queryset for heatcontrol.ScheduledOverride
type ScheduledoverrideQS struct {
	condFragments models.AndFragment
	order         []string
	forClause     string
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

// Or combines given expressions with OR operator
func (qs ScheduledoverrideQS) Or(exprs ...ScheduledoverrideQS) ScheduledoverrideQS {
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

type inScheduledoverrideid []interface{}

func (in inScheduledoverrideid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ScheduledoverrideQS) IDIn(values []int32) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inScheduledoverrideid(vals),
	)

	return qs
}

type notinScheduledoverrideid []interface{}

func (in notinScheduledoverrideid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ScheduledoverrideQS) IDNotIn(values []int32) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinScheduledoverrideid(vals),
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
func (s *Scheduledoverride) GetControl(ctx context.Context, db models.DBInterface) (*Control, error) {
	return ControlQS{}.IDEq(s.control).First(ctx, db)
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

// ControlRawEq filters for control being equal to raw argument
func (qs ScheduledoverrideQS) ControlRawEq(v int32) ScheduledoverrideQS {
	return qs.filter(`"control_id" =`, v)
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

type inScheduledoverrideStart []interface{}

func (in inScheduledoverrideStart) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"start" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ScheduledoverrideQS) StartIn(values []time.Time) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inScheduledoverrideStart(vals),
	)

	return qs
}

type notinScheduledoverrideStart []interface{}

func (in notinScheduledoverrideStart) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"start" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ScheduledoverrideQS) StartNotIn(values []time.Time) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinScheduledoverrideStart(vals),
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

type inScheduledoverrideEnd []interface{}

func (in inScheduledoverrideEnd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"end" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ScheduledoverrideQS) EndIn(values []time.Time) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inScheduledoverrideEnd(vals),
	)

	return qs
}

type notinScheduledoverrideEnd []interface{}

func (in notinScheduledoverrideEnd) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"end" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ScheduledoverrideQS) EndNotIn(values []time.Time) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinScheduledoverrideEnd(vals),
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

type inScheduledoverrideTargetTemp []interface{}

func (in inScheduledoverrideTargetTemp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"target_temp" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ScheduledoverrideQS) TargetTempIn(values []float64) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inScheduledoverrideTargetTemp(vals),
	)

	return qs
}

type notinScheduledoverrideTargetTemp []interface{}

func (in notinScheduledoverrideTargetTemp) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"target_temp" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ScheduledoverrideQS) TargetTempNotIn(values []float64) ScheduledoverrideQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinScheduledoverrideTargetTemp(vals),
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

// OrderByRandom randomizes result
func (qs ScheduledoverrideQS) OrderByRandom() ScheduledoverrideQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs ScheduledoverrideQS) ForUpdate() ScheduledoverrideQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs ScheduledoverrideQS) ForUpdateNowait() ScheduledoverrideQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs ScheduledoverrideQS) ForUpdateSkipLocked() ScheduledoverrideQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs ScheduledoverrideQS) ClearForUpdate() ScheduledoverrideQS {
	qs.forClause = ""

	return qs
}

func (qs ScheduledoverrideQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

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
	s += qs.forClause

	return `SELECT "id", "control_id", "start", "end", "target_temp" FROM "heatcontrol_scheduledoverride"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs ScheduledoverrideQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_scheduledoverride"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs ScheduledoverrideQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	row := db.QueryRow(ctx, `SELECT COUNT("id") FROM "heatcontrol_scheduledoverride"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs ScheduledoverrideQS) All(ctx context.Context, db models.DBInterface) (ScheduledoverrideList, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret ScheduledoverrideList
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
func (qs ScheduledoverrideQS) First(ctx context.Context, db models.DBInterface) (*Scheduledoverride, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Scheduledoverride{existsInDB: true}
	err := row.Scan(&obj.id, &obj.control, &obj.Start, &obj.End, &obj.TargetTemp)
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
func (qs ScheduledoverrideQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_scheduledoverride"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
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
func (uqs ScheduledoverrideUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (s *Scheduledoverride) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "heatcontrol_scheduledoverride" ("control_id", "start", "end", "target_temp") VALUES ($1, $2, $3, $4) RETURNING "id"`, s.control, s.Start, s.End, s.TargetTemp)

	if err := row.Scan(&s.id); err != nil {
		return err
	}

	s.existsInDB = true

	return nil
}

// update operation
func (s *Scheduledoverride) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "heatcontrol_scheduledoverride" SET "control_id" = $1, "start" = $2, "end" = $3, "target_temp" = $4 WHERE "id" = $5`, s.control, s.Start, s.End, s.TargetTemp, s.id)

	return err
}

// Save inserts or updates record
func (s *Scheduledoverride) Save(ctx context.Context, db models.DBInterface) error {
	if s.existsInDB {
		return s.update(ctx, db)
	}

	return s.insert(ctx, db)
}

// Delete removes row from database
func (s *Scheduledoverride) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "heatcontrol_scheduledoverride" WHERE "id" = $1`, s.id)

	s.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (sl ScheduledoverrideList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts ScheduledoverrideList

	for _, s := range sl {
		if s.existsInDB {
			if err := s.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, s)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 4*len(inserts))
	offs := 1
	for _, s := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3))
		vaa = append(vaa, s.control, s.Start, s.End, s.TargetTemp)
		offs += 4
	}

	qs := `INSERT INTO "heatcontrol_scheduledoverride" ("control_id", "start", "end", "target_temp") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, s := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&s.id); err != nil {
			return err
		}

		s.existsInDB = true
	}

	return nil
}
