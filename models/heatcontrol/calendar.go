// Code generated for Django model heatcontrol.Calendar. DO NOT EDIT.

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

// Calendar mirrors model heatcontrol.Calendar
type Calendar struct {
	existsInDB bool

	id      int32
	Day     time.Time
	daytype int32
}

// CalendarList is a list of Calendar
type CalendarList []*Calendar

// CalendarQS represents a queryset for heatcontrol.Calendar
type CalendarQS struct {
	condFragments models.AndFragment
	order         []string
	forClause     string
}

func (qs CalendarQS) filter(c string, p interface{}) CalendarQS {
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
func (qs CalendarQS) Or(exprs ...CalendarQS) CalendarQS {
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

// GetID returns Calendar.ID
func (c *Calendar) GetID() int32 {
	return c.id
}

// IDEq filters for id being equal to argument
func (qs CalendarQS) IDEq(v int32) CalendarQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs CalendarQS) IDNe(v int32) CalendarQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs CalendarQS) IDLt(v int32) CalendarQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs CalendarQS) IDLe(v int32) CalendarQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs CalendarQS) IDGt(v int32) CalendarQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs CalendarQS) IDGe(v int32) CalendarQS {
	return qs.filter(`"id" >=`, v)
}

type inCalendarid []interface{}

func (in inCalendarid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs CalendarQS) IDIn(values []int32) CalendarQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inCalendarid(vals),
	)

	return qs
}

type notinCalendarid []interface{}

func (in notinCalendarid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs CalendarQS) IDNotIn(values []int32) CalendarQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinCalendarid(vals),
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs CalendarQS) OrderByID() CalendarQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs CalendarQS) OrderByIDDesc() CalendarQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// DayEq filters for Day being equal to argument
func (qs CalendarQS) DayEq(v time.Time) CalendarQS {
	return qs.filter(`"day" =`, v)
}

// DayNe filters for Day being not equal to argument
func (qs CalendarQS) DayNe(v time.Time) CalendarQS {
	return qs.filter(`"day" <>`, v)
}

// DayLt filters for Day being less than argument
func (qs CalendarQS) DayLt(v time.Time) CalendarQS {
	return qs.filter(`"day" <`, v)
}

// DayLe filters for Day being less than or equal to argument
func (qs CalendarQS) DayLe(v time.Time) CalendarQS {
	return qs.filter(`"day" <=`, v)
}

// DayGt filters for Day being greater than argument
func (qs CalendarQS) DayGt(v time.Time) CalendarQS {
	return qs.filter(`"day" >`, v)
}

// DayGe filters for Day being greater than or equal to argument
func (qs CalendarQS) DayGe(v time.Time) CalendarQS {
	return qs.filter(`"day" >=`, v)
}

type inCalendarDay []interface{}

func (in inCalendarDay) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"day" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs CalendarQS) DayIn(values []time.Time) CalendarQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inCalendarDay(vals),
	)

	return qs
}

type notinCalendarDay []interface{}

func (in notinCalendarDay) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"day" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs CalendarQS) DayNotIn(values []time.Time) CalendarQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinCalendarDay(vals),
	)

	return qs
}

// OrderByDay sorts result by Day in ascending order
func (qs CalendarQS) OrderByDay() CalendarQS {
	qs.order = append(qs.order, `"day"`)

	return qs
}

// OrderByDayDesc sorts result by Day in descending order
func (qs CalendarQS) OrderByDayDesc() CalendarQS {
	qs.order = append(qs.order, `"day" DESC`)

	return qs
}

// GetDaytype returns Daytype
func (c *Calendar) GetDaytype(ctx context.Context, db models.DBInterface) (*Daytype, error) {
	return DaytypeQS{}.IDEq(c.daytype).First(ctx, db)
}

// SetDaytype sets foreign key pointer to Daytype
func (c *Calendar) SetDaytype(ptr *Daytype) error {
	if ptr != nil {
		c.daytype = ptr.GetID()
	} else {
		return fmt.Errorf("Calendar.SetDaytype: non-null field received null value")
	}

	return nil
}

// GetDaytypeRaw returns Calendar.Daytype
func (c *Calendar) GetDaytypeRaw() int32 {
	return c.daytype
}

// DaytypeEq filters for daytype being equal to argument
func (qs CalendarQS) DaytypeEq(v *Daytype) CalendarQS {
	return qs.filter(`"daytype_id" =`, v.GetID())
}

// DaytypeRawEq filters for daytype being equal to raw argument
func (qs CalendarQS) DaytypeRawEq(v int32) CalendarQS {
	return qs.filter(`"daytype_id" =`, v)
}

type inCalendardaytypeDaytype struct {
	qs DaytypeQS
}

func (in *inCalendardaytypeDaytype) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"daytype_id" IN (` + s + `)`, p
}

func (qs CalendarQS) DaytypeIn(oqs DaytypeQS) CalendarQS {
	qs.condFragments = append(
		qs.condFragments,
		&inCalendardaytypeDaytype{
			qs: oqs,
		},
	)

	return qs
}

// OrderByDaytype sorts result by Daytype in ascending order
func (qs CalendarQS) OrderByDaytype() CalendarQS {
	qs.order = append(qs.order, `"daytype_id"`)

	return qs
}

// OrderByDaytypeDesc sorts result by Daytype in descending order
func (qs CalendarQS) OrderByDaytypeDesc() CalendarQS {
	qs.order = append(qs.order, `"daytype_id" DESC`)

	return qs
}

// OrderByRandom randomizes result
func (qs CalendarQS) OrderByRandom() CalendarQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs CalendarQS) ForUpdate() CalendarQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs CalendarQS) ForUpdateNowait() CalendarQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs CalendarQS) ForUpdateSkipLocked() CalendarQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs CalendarQS) ClearForUpdate() CalendarQS {
	qs.forClause = ""

	return qs
}

func (qs CalendarQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs CalendarQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs CalendarQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	return `SELECT "id", "day", "daytype_id" FROM "heatcontrol_calendar"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs CalendarQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "heatcontrol_calendar"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs CalendarQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	row := db.QueryRow(ctx, `SELECT COUNT("id") FROM "heatcontrol_calendar"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs CalendarQS) All(ctx context.Context, db models.DBInterface) (CalendarList, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret CalendarList
	for rows.Next() {
		obj := Calendar{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.Day, &obj.daytype); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs CalendarQS) First(ctx context.Context, db models.DBInterface) (*Calendar, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Calendar{existsInDB: true}
	err := row.Scan(&obj.id, &obj.Day, &obj.daytype)
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
func (qs CalendarQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "heatcontrol_calendar"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs CalendarQS) Update() CalendarUpdateQS {
	return CalendarUpdateQS{condFragments: qs.condFragments}
}

// CalendarUpdateQS represents an updated queryset for heatcontrol.Calendar
type CalendarUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs CalendarUpdateQS) update(c string, v interface{}) CalendarUpdateQS {
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
func (uqs CalendarUpdateQS) SetID(v int32) CalendarUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetDay sets Day to the given value
func (uqs CalendarUpdateQS) SetDay(v time.Time) CalendarUpdateQS {
	return uqs.update(`"day"`, v)
}

// SetDaytype sets foreign key pointer to Daytype
func (uqs CalendarUpdateQS) SetDaytype(ptr *Daytype) CalendarUpdateQS {
	if ptr != nil {
		return uqs.update(`"daytype_id"`, ptr.GetID())
	}

	return uqs.update(`"daytype_id"`, nil)
} // Exec executes the update operation
func (uqs CalendarUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := CalendarQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "heatcontrol_calendar" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (c *Calendar) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "heatcontrol_calendar" ("day", "daytype_id") VALUES ($1, $2) RETURNING "id"`, c.Day, c.daytype)

	if err := row.Scan(&c.id); err != nil {
		return err
	}

	c.existsInDB = true

	return nil
}

// update operation
func (c *Calendar) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "heatcontrol_calendar" SET "day" = $1, "daytype_id" = $2 WHERE "id" = $3`, c.Day, c.daytype, c.id)

	return err
}

// Save inserts or updates record
func (c *Calendar) Save(ctx context.Context, db models.DBInterface) error {
	if c.existsInDB {
		return c.update(ctx, db)
	}

	return c.insert(ctx, db)
}

// Delete removes row from database
func (c *Calendar) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "heatcontrol_calendar" WHERE "id" = $1`, c.id)

	c.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (cl CalendarList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts CalendarList

	for _, c := range cl {
		if c.existsInDB {
			if err := c.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, c)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 2*len(inserts))
	offs := 1
	for _, c := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d)", offs+0, offs+1))
		vaa = append(vaa, c.Day, c.daytype)
		offs += 2
	}

	qs := `INSERT INTO "heatcontrol_calendar" ("day", "daytype_id") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, c := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&c.id); err != nil {
			return err
		}

		c.existsInDB = true
	}

	return nil
}
