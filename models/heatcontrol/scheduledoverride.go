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

// GetId returns Scheduledoverride.Id
func (s *Scheduledoverride) GetId() int32 {
	return s.id
}

// IdEq filters for id being equal to argument
func (qs ScheduledoverrideQS) IdEq(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" =`, v)
}

// IdNe filters for id being not equal to argument
func (qs ScheduledoverrideQS) IdNe(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" <>`, v)
}

// IdLt filters for id being less than argument
func (qs ScheduledoverrideQS) IdLt(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" <`, v)
}

// IdLe filters for id being less than or equal to argument
func (qs ScheduledoverrideQS) IdLe(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" <=`, v)
}

// IdGt filters for id being greater than argument
func (qs ScheduledoverrideQS) IdGt(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" >`, v)
}

// IdGe filters for id being greater than or equal to argument
func (qs ScheduledoverrideQS) IdGe(v int32) ScheduledoverrideQS {
	return qs.filter(`"id" >=`, v)
}

// OrderById sorts result by Id in ascending order
func (qs ScheduledoverrideQS) OrderById() ScheduledoverrideQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIdDesc sorts result by Id in descending order
func (qs ScheduledoverrideQS) OrderByIdDesc() ScheduledoverrideQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// GetControl returns Control
func (s *Scheduledoverride) GetControl(db models.DBInterface) (*Control, error) {
	return ControlQS{}.IdEq(s.control).First(db)
}

// SetControl sets foreign key pointer to Control
func (s *Scheduledoverride) SetControl(ptr *Control) error {
	if ptr != nil {
		s.control = ptr.GetId()
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
	return qs.filter(`"control_id" =`, v.GetId())
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
