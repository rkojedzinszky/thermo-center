// Code generated for Django model center.ConfigureSensorTask. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center/v5 center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package center

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/rkojedzinszky/thermo-center/v5/models"
)

// Configuresensortask mirrors model center.ConfigureSensorTask
type Configuresensortask struct {
	existsInDB bool

	id             int32
	sensor         int32
	Created        time.Time
	Started        sql.NullTime
	FirstDiscovery sql.NullTime
	LastDiscovery  sql.NullTime
	Finished       sql.NullTime
	Error          sql.NullString
}

// ConfiguresensortaskList is a list of Configuresensortask
type ConfiguresensortaskList []*Configuresensortask

// ConfiguresensortaskQS represents a queryset for center.ConfigureSensorTask
type ConfiguresensortaskQS struct {
	distinctOnFields []string
	condFragments    models.AndFragment
	order            []string
	forClause        string
}

func (qs ConfiguresensortaskQS) filter(c string, p interface{}) ConfiguresensortaskQS {
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
func (qs ConfiguresensortaskQS) Or(exprs ...ConfiguresensortaskQS) ConfiguresensortaskQS {
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

// BEGIN - center.ConfigureSensorTask.id

// GetID returns Configuresensortask.ID
func (c *Configuresensortask) GetID() int32 {
	return c.id
}

// IDEq filters for id being equal to argument
func (qs ConfiguresensortaskQS) IDEq(v int32) ConfiguresensortaskQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs ConfiguresensortaskQS) IDNe(v int32) ConfiguresensortaskQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs ConfiguresensortaskQS) IDLt(v int32) ConfiguresensortaskQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs ConfiguresensortaskQS) IDLe(v int32) ConfiguresensortaskQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs ConfiguresensortaskQS) IDGt(v int32) ConfiguresensortaskQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs ConfiguresensortaskQS) IDGe(v int32) ConfiguresensortaskQS {
	return qs.filter(`"id" >=`, v)
}

type inConfiguresensortaskid []interface{}

func (in inConfiguresensortaskid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) IDIn(values []int32) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inConfiguresensortaskid(vals),
	)

	return qs
}

type notinConfiguresensortaskid []interface{}

func (in notinConfiguresensortaskid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) IDNotIn(values []int32) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinConfiguresensortaskid(vals),
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs ConfiguresensortaskQS) OrderByID() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs ConfiguresensortaskQS) OrderByIDDesc() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// DistinctOnID marks field in queries to add to DISTINCT ON clause
func (qs ConfiguresensortaskQS) DistinctOnID() ConfiguresensortaskQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"id"`)

	return qs
}

// END - center.ConfigureSensorTask.id

// BEGIN - center.ConfigureSensorTask.sensor

// GetSensor returns Sensor
func (c *Configuresensortask) GetSensor(ctx context.Context, db models.DBInterface) (*Sensor, error) {
	return SensorQS{}.IDEq(c.sensor).First(ctx, db)
}

// SetSensor sets foreign key pointer to Sensor
func (c *Configuresensortask) SetSensor(ptr *Sensor) error {
	if ptr != nil {
		c.sensor = ptr.ID
	} else {
		return fmt.Errorf("Configuresensortask.SetSensor: non-null field received null value")
	}

	return nil
}

// GetSensorRaw returns Configuresensortask.Sensor
func (c *Configuresensortask) GetSensorRaw() int32 {
	return c.sensor
}

// SensorEq filters for sensor being equal to argument
func (qs ConfiguresensortaskQS) SensorEq(v *Sensor) ConfiguresensortaskQS {
	return qs.filter(`"sensor_id" =`, v.ID)
}

// SensorRawEq filters for sensor being equal to raw argument
func (qs ConfiguresensortaskQS) SensorRawEq(v int32) ConfiguresensortaskQS {
	return qs.filter(`"sensor_id" =`, v)
}

type inConfiguresensortasksensorSensor struct {
	qs SensorQS
}

func (in *inConfiguresensortasksensorSensor) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"sensor_id" IN (` + s + `)`, p
}

func (qs ConfiguresensortaskQS) SensorIn(oqs SensorQS) ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&inConfiguresensortasksensorSensor{
			qs: oqs,
		},
	)

	return qs
}

// OrderBySensor sorts result by Sensor in ascending order
func (qs ConfiguresensortaskQS) OrderBySensor() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"sensor_id"`)

	return qs
}

// OrderBySensorDesc sorts result by Sensor in descending order
func (qs ConfiguresensortaskQS) OrderBySensorDesc() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"sensor_id" DESC`)

	return qs
}

// DistinctOnSensor marks field in queries to add to DISTINCT ON clause
func (qs ConfiguresensortaskQS) DistinctOnSensor() ConfiguresensortaskQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"sensor_id"`)

	return qs
}

// END - center.ConfigureSensorTask.sensor

// BEGIN - center.ConfigureSensorTask.created

// CreatedEq filters for Created being equal to argument
func (qs ConfiguresensortaskQS) CreatedEq(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"created" =`, v)
}

// CreatedNe filters for Created being not equal to argument
func (qs ConfiguresensortaskQS) CreatedNe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"created" <>`, v)
}

// CreatedLt filters for Created being less than argument
func (qs ConfiguresensortaskQS) CreatedLt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"created" <`, v)
}

// CreatedLe filters for Created being less than or equal to argument
func (qs ConfiguresensortaskQS) CreatedLe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"created" <=`, v)
}

// CreatedGt filters for Created being greater than argument
func (qs ConfiguresensortaskQS) CreatedGt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"created" >`, v)
}

// CreatedGe filters for Created being greater than or equal to argument
func (qs ConfiguresensortaskQS) CreatedGe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"created" >=`, v)
}

type inConfiguresensortaskCreated []interface{}

func (in inConfiguresensortaskCreated) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"created" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) CreatedIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inConfiguresensortaskCreated(vals),
	)

	return qs
}

type notinConfiguresensortaskCreated []interface{}

func (in notinConfiguresensortaskCreated) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"created" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) CreatedNotIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinConfiguresensortaskCreated(vals),
	)

	return qs
}

// OrderByCreated sorts result by Created in ascending order
func (qs ConfiguresensortaskQS) OrderByCreated() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"created"`)

	return qs
}

// OrderByCreatedDesc sorts result by Created in descending order
func (qs ConfiguresensortaskQS) OrderByCreatedDesc() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"created" DESC`)

	return qs
}

// DistinctOnCreated marks field in queries to add to DISTINCT ON clause
func (qs ConfiguresensortaskQS) DistinctOnCreated() ConfiguresensortaskQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"created"`)

	return qs
}

// END - center.ConfigureSensorTask.created

// BEGIN - center.ConfigureSensorTask.started

// StartedIsNull filters for Started being null
func (qs ConfiguresensortaskQS) StartedIsNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"started" IS NULL`,
		},
	)
	return qs
}

// StartedIsNotNull filters for Started being not null
func (qs ConfiguresensortaskQS) StartedIsNotNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"started" IS NOT NULL`,
		},
	)
	return qs
}

// StartedEq filters for Started being equal to argument
func (qs ConfiguresensortaskQS) StartedEq(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"started" =`, v)
}

// StartedNe filters for Started being not equal to argument
func (qs ConfiguresensortaskQS) StartedNe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"started" <>`, v)
}

// StartedLt filters for Started being less than argument
func (qs ConfiguresensortaskQS) StartedLt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"started" <`, v)
}

// StartedLe filters for Started being less than or equal to argument
func (qs ConfiguresensortaskQS) StartedLe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"started" <=`, v)
}

// StartedGt filters for Started being greater than argument
func (qs ConfiguresensortaskQS) StartedGt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"started" >`, v)
}

// StartedGe filters for Started being greater than or equal to argument
func (qs ConfiguresensortaskQS) StartedGe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"started" >=`, v)
}

type inConfiguresensortaskStarted []interface{}

func (in inConfiguresensortaskStarted) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"started" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) StartedIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inConfiguresensortaskStarted(vals),
	)

	return qs
}

type notinConfiguresensortaskStarted []interface{}

func (in notinConfiguresensortaskStarted) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"started" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) StartedNotIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinConfiguresensortaskStarted(vals),
	)

	return qs
}

// OrderByStarted sorts result by Started in ascending order
func (qs ConfiguresensortaskQS) OrderByStarted() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"started"`)

	return qs
}

// OrderByStartedDesc sorts result by Started in descending order
func (qs ConfiguresensortaskQS) OrderByStartedDesc() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"started" DESC`)

	return qs
}

// DistinctOnStarted marks field in queries to add to DISTINCT ON clause
func (qs ConfiguresensortaskQS) DistinctOnStarted() ConfiguresensortaskQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"started"`)

	return qs
}

// END - center.ConfigureSensorTask.started

// BEGIN - center.ConfigureSensorTask.first_discovery

// FirstDiscoveryIsNull filters for FirstDiscovery being null
func (qs ConfiguresensortaskQS) FirstDiscoveryIsNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"first_discovery" IS NULL`,
		},
	)
	return qs
}

// FirstDiscoveryIsNotNull filters for FirstDiscovery being not null
func (qs ConfiguresensortaskQS) FirstDiscoveryIsNotNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"first_discovery" IS NOT NULL`,
		},
	)
	return qs
}

// FirstDiscoveryEq filters for FirstDiscovery being equal to argument
func (qs ConfiguresensortaskQS) FirstDiscoveryEq(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"first_discovery" =`, v)
}

// FirstDiscoveryNe filters for FirstDiscovery being not equal to argument
func (qs ConfiguresensortaskQS) FirstDiscoveryNe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"first_discovery" <>`, v)
}

// FirstDiscoveryLt filters for FirstDiscovery being less than argument
func (qs ConfiguresensortaskQS) FirstDiscoveryLt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"first_discovery" <`, v)
}

// FirstDiscoveryLe filters for FirstDiscovery being less than or equal to argument
func (qs ConfiguresensortaskQS) FirstDiscoveryLe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"first_discovery" <=`, v)
}

// FirstDiscoveryGt filters for FirstDiscovery being greater than argument
func (qs ConfiguresensortaskQS) FirstDiscoveryGt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"first_discovery" >`, v)
}

// FirstDiscoveryGe filters for FirstDiscovery being greater than or equal to argument
func (qs ConfiguresensortaskQS) FirstDiscoveryGe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"first_discovery" >=`, v)
}

type inConfiguresensortaskFirstDiscovery []interface{}

func (in inConfiguresensortaskFirstDiscovery) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"first_discovery" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) FirstDiscoveryIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inConfiguresensortaskFirstDiscovery(vals),
	)

	return qs
}

type notinConfiguresensortaskFirstDiscovery []interface{}

func (in notinConfiguresensortaskFirstDiscovery) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"first_discovery" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) FirstDiscoveryNotIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinConfiguresensortaskFirstDiscovery(vals),
	)

	return qs
}

// OrderByFirstDiscovery sorts result by FirstDiscovery in ascending order
func (qs ConfiguresensortaskQS) OrderByFirstDiscovery() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"first_discovery"`)

	return qs
}

// OrderByFirstDiscoveryDesc sorts result by FirstDiscovery in descending order
func (qs ConfiguresensortaskQS) OrderByFirstDiscoveryDesc() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"first_discovery" DESC`)

	return qs
}

// DistinctOnFirstDiscovery marks field in queries to add to DISTINCT ON clause
func (qs ConfiguresensortaskQS) DistinctOnFirstDiscovery() ConfiguresensortaskQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"first_discovery"`)

	return qs
}

// END - center.ConfigureSensorTask.first_discovery

// BEGIN - center.ConfigureSensorTask.last_discovery

// LastDiscoveryIsNull filters for LastDiscovery being null
func (qs ConfiguresensortaskQS) LastDiscoveryIsNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"last_discovery" IS NULL`,
		},
	)
	return qs
}

// LastDiscoveryIsNotNull filters for LastDiscovery being not null
func (qs ConfiguresensortaskQS) LastDiscoveryIsNotNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"last_discovery" IS NOT NULL`,
		},
	)
	return qs
}

// LastDiscoveryEq filters for LastDiscovery being equal to argument
func (qs ConfiguresensortaskQS) LastDiscoveryEq(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"last_discovery" =`, v)
}

// LastDiscoveryNe filters for LastDiscovery being not equal to argument
func (qs ConfiguresensortaskQS) LastDiscoveryNe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"last_discovery" <>`, v)
}

// LastDiscoveryLt filters for LastDiscovery being less than argument
func (qs ConfiguresensortaskQS) LastDiscoveryLt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"last_discovery" <`, v)
}

// LastDiscoveryLe filters for LastDiscovery being less than or equal to argument
func (qs ConfiguresensortaskQS) LastDiscoveryLe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"last_discovery" <=`, v)
}

// LastDiscoveryGt filters for LastDiscovery being greater than argument
func (qs ConfiguresensortaskQS) LastDiscoveryGt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"last_discovery" >`, v)
}

// LastDiscoveryGe filters for LastDiscovery being greater than or equal to argument
func (qs ConfiguresensortaskQS) LastDiscoveryGe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"last_discovery" >=`, v)
}

type inConfiguresensortaskLastDiscovery []interface{}

func (in inConfiguresensortaskLastDiscovery) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"last_discovery" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) LastDiscoveryIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inConfiguresensortaskLastDiscovery(vals),
	)

	return qs
}

type notinConfiguresensortaskLastDiscovery []interface{}

func (in notinConfiguresensortaskLastDiscovery) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"last_discovery" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) LastDiscoveryNotIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinConfiguresensortaskLastDiscovery(vals),
	)

	return qs
}

// OrderByLastDiscovery sorts result by LastDiscovery in ascending order
func (qs ConfiguresensortaskQS) OrderByLastDiscovery() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"last_discovery"`)

	return qs
}

// OrderByLastDiscoveryDesc sorts result by LastDiscovery in descending order
func (qs ConfiguresensortaskQS) OrderByLastDiscoveryDesc() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"last_discovery" DESC`)

	return qs
}

// DistinctOnLastDiscovery marks field in queries to add to DISTINCT ON clause
func (qs ConfiguresensortaskQS) DistinctOnLastDiscovery() ConfiguresensortaskQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"last_discovery"`)

	return qs
}

// END - center.ConfigureSensorTask.last_discovery

// BEGIN - center.ConfigureSensorTask.finished

// FinishedIsNull filters for Finished being null
func (qs ConfiguresensortaskQS) FinishedIsNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"finished" IS NULL`,
		},
	)
	return qs
}

// FinishedIsNotNull filters for Finished being not null
func (qs ConfiguresensortaskQS) FinishedIsNotNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"finished" IS NOT NULL`,
		},
	)
	return qs
}

// FinishedEq filters for Finished being equal to argument
func (qs ConfiguresensortaskQS) FinishedEq(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"finished" =`, v)
}

// FinishedNe filters for Finished being not equal to argument
func (qs ConfiguresensortaskQS) FinishedNe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"finished" <>`, v)
}

// FinishedLt filters for Finished being less than argument
func (qs ConfiguresensortaskQS) FinishedLt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"finished" <`, v)
}

// FinishedLe filters for Finished being less than or equal to argument
func (qs ConfiguresensortaskQS) FinishedLe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"finished" <=`, v)
}

// FinishedGt filters for Finished being greater than argument
func (qs ConfiguresensortaskQS) FinishedGt(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"finished" >`, v)
}

// FinishedGe filters for Finished being greater than or equal to argument
func (qs ConfiguresensortaskQS) FinishedGe(v time.Time) ConfiguresensortaskQS {
	return qs.filter(`"finished" >=`, v)
}

type inConfiguresensortaskFinished []interface{}

func (in inConfiguresensortaskFinished) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"finished" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) FinishedIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inConfiguresensortaskFinished(vals),
	)

	return qs
}

type notinConfiguresensortaskFinished []interface{}

func (in notinConfiguresensortaskFinished) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"finished" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) FinishedNotIn(values []time.Time) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinConfiguresensortaskFinished(vals),
	)

	return qs
}

// OrderByFinished sorts result by Finished in ascending order
func (qs ConfiguresensortaskQS) OrderByFinished() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"finished"`)

	return qs
}

// OrderByFinishedDesc sorts result by Finished in descending order
func (qs ConfiguresensortaskQS) OrderByFinishedDesc() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"finished" DESC`)

	return qs
}

// DistinctOnFinished marks field in queries to add to DISTINCT ON clause
func (qs ConfiguresensortaskQS) DistinctOnFinished() ConfiguresensortaskQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"finished"`)

	return qs
}

// END - center.ConfigureSensorTask.finished

// BEGIN - center.ConfigureSensorTask.error

// ErrorIsNull filters for Error being null
func (qs ConfiguresensortaskQS) ErrorIsNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"error" IS NULL`,
		},
	)
	return qs
}

// ErrorIsNotNull filters for Error being not null
func (qs ConfiguresensortaskQS) ErrorIsNotNull() ConfiguresensortaskQS {
	qs.condFragments = append(
		qs.condFragments,
		&models.ConstantFragment{
			Constant: `"error" IS NOT NULL`,
		},
	)
	return qs
}

// ErrorEq filters for Error being equal to argument
func (qs ConfiguresensortaskQS) ErrorEq(v string) ConfiguresensortaskQS {
	return qs.filter(`"error" =`, v)
}

// ErrorNe filters for Error being not equal to argument
func (qs ConfiguresensortaskQS) ErrorNe(v string) ConfiguresensortaskQS {
	return qs.filter(`"error" <>`, v)
}

// ErrorLt filters for Error being less than argument
func (qs ConfiguresensortaskQS) ErrorLt(v string) ConfiguresensortaskQS {
	return qs.filter(`"error" <`, v)
}

// ErrorLe filters for Error being less than or equal to argument
func (qs ConfiguresensortaskQS) ErrorLe(v string) ConfiguresensortaskQS {
	return qs.filter(`"error" <=`, v)
}

// ErrorGt filters for Error being greater than argument
func (qs ConfiguresensortaskQS) ErrorGt(v string) ConfiguresensortaskQS {
	return qs.filter(`"error" >`, v)
}

// ErrorGe filters for Error being greater than or equal to argument
func (qs ConfiguresensortaskQS) ErrorGe(v string) ConfiguresensortaskQS {
	return qs.filter(`"error" >=`, v)
}

type inConfiguresensortaskError []interface{}

func (in inConfiguresensortaskError) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"error" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) ErrorIn(values []string) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inConfiguresensortaskError(vals),
	)

	return qs
}

type notinConfiguresensortaskError []interface{}

func (in notinConfiguresensortaskError) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"error" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs ConfiguresensortaskQS) ErrorNotIn(values []string) ConfiguresensortaskQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinConfiguresensortaskError(vals),
	)

	return qs
}

// OrderByError sorts result by Error in ascending order
func (qs ConfiguresensortaskQS) OrderByError() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"error"`)

	return qs
}

// OrderByErrorDesc sorts result by Error in descending order
func (qs ConfiguresensortaskQS) OrderByErrorDesc() ConfiguresensortaskQS {
	qs.order = append(qs.order, `"error" DESC`)

	return qs
}

// DistinctOnError marks field in queries to add to DISTINCT ON clause
func (qs ConfiguresensortaskQS) DistinctOnError() ConfiguresensortaskQS {
	qs.distinctOnFields = append(qs.distinctOnFields, `"error"`)

	return qs
}

// END - center.ConfigureSensorTask.error

// OrderByRandom randomizes result
func (qs ConfiguresensortaskQS) OrderByRandom() ConfiguresensortaskQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs ConfiguresensortaskQS) ForUpdate() ConfiguresensortaskQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs ConfiguresensortaskQS) ForUpdateNowait() ConfiguresensortaskQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs ConfiguresensortaskQS) ForUpdateSkipLocked() ConfiguresensortaskQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs ConfiguresensortaskQS) ClearForUpdate() ConfiguresensortaskQS {
	qs.forClause = ""

	return qs
}

func (qs ConfiguresensortaskQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs ConfiguresensortaskQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs ConfiguresensortaskQS) queryFull(distinctOnFields []string) (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	var distinctClause string
	if len(distinctOnFields) > 0 {
		distinctClause = fmt.Sprintf("DISTINCT ON (%s) ", strings.Join(distinctOnFields, ", "))
	}

	return `SELECT ` + distinctClause + `"id", "sensor_id", "created", "started", "first_discovery", "last_discovery", "finished", "error" FROM "center_configuresensortask"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs ConfiguresensortaskQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "center_configuresensortask"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs ConfiguresensortaskQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	var countClause string
	if len(qs.distinctOnFields) > 0 {
		countClause = fmt.Sprintf("DISTINCT (%s)", strings.Join(qs.distinctOnFields, ", "))
	} else {
		countClause = `"id"`
	}

	row := db.QueryRow(ctx, `SELECT COUNT(`+countClause+`) FROM "center_configuresensortask"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs ConfiguresensortaskQS) All(ctx context.Context, db models.DBInterface) (ConfiguresensortaskList, error) {
	s, p := qs.queryFull(qs.distinctOnFields)

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret ConfiguresensortaskList
	for rows.Next() {
		obj := Configuresensortask{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.sensor, &obj.Created, &obj.Started, &obj.FirstDiscovery, &obj.LastDiscovery, &obj.Finished, &obj.Error); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs ConfiguresensortaskQS) First(ctx context.Context, db models.DBInterface) (*Configuresensortask, error) {
	s, p := qs.queryFull(nil)

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Configuresensortask{existsInDB: true}
	err := row.Scan(&obj.id, &obj.sensor, &obj.Created, &obj.Started, &obj.FirstDiscovery, &obj.LastDiscovery, &obj.Finished, &obj.Error)
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
func (qs ConfiguresensortaskQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "center_configuresensortask"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs ConfiguresensortaskQS) Update() ConfiguresensortaskUpdateQS {
	return ConfiguresensortaskUpdateQS{condFragments: qs.condFragments}
}

// ConfiguresensortaskUpdateQS represents an updated queryset for center.ConfigureSensorTask
type ConfiguresensortaskUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs ConfiguresensortaskUpdateQS) update(c string, v interface{}) ConfiguresensortaskUpdateQS {
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
func (uqs ConfiguresensortaskUpdateQS) SetID(v int32) ConfiguresensortaskUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetSensor sets foreign key pointer to Sensor
func (uqs ConfiguresensortaskUpdateQS) SetSensor(ptr *Sensor) ConfiguresensortaskUpdateQS {
	if ptr != nil {
		return uqs.update(`"sensor_id"`, ptr.ID)
	}

	return uqs.update(`"sensor_id"`, nil)
} // SetCreated sets Created to the given value
func (uqs ConfiguresensortaskUpdateQS) SetCreated(v time.Time) ConfiguresensortaskUpdateQS {
	return uqs.update(`"created"`, v)
}

// SetStarted sets Started to the given value
func (uqs ConfiguresensortaskUpdateQS) SetStarted(v sql.NullTime) ConfiguresensortaskUpdateQS {
	return uqs.update(`"started"`, v)
}

// SetFirstDiscovery sets FirstDiscovery to the given value
func (uqs ConfiguresensortaskUpdateQS) SetFirstDiscovery(v sql.NullTime) ConfiguresensortaskUpdateQS {
	return uqs.update(`"first_discovery"`, v)
}

// SetLastDiscovery sets LastDiscovery to the given value
func (uqs ConfiguresensortaskUpdateQS) SetLastDiscovery(v sql.NullTime) ConfiguresensortaskUpdateQS {
	return uqs.update(`"last_discovery"`, v)
}

// SetFinished sets Finished to the given value
func (uqs ConfiguresensortaskUpdateQS) SetFinished(v sql.NullTime) ConfiguresensortaskUpdateQS {
	return uqs.update(`"finished"`, v)
}

// SetError sets Error to the given value
func (uqs ConfiguresensortaskUpdateQS) SetError(v sql.NullString) ConfiguresensortaskUpdateQS {
	return uqs.update(`"error"`, v)
}

// Exec executes the update operation
func (uqs ConfiguresensortaskUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := ConfiguresensortaskQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "center_configuresensortask" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (c *Configuresensortask) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "center_configuresensortask" ("sensor_id", "created", "started", "first_discovery", "last_discovery", "finished", "error") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING "id"`, c.sensor, c.Created, c.Started, c.FirstDiscovery, c.LastDiscovery, c.Finished, c.Error)

	if err := row.Scan(&c.id); err != nil {
		return err
	}

	c.existsInDB = true

	return nil
}

// update operation
func (c *Configuresensortask) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "center_configuresensortask" SET "sensor_id" = $1, "created" = $2, "started" = $3, "first_discovery" = $4, "last_discovery" = $5, "finished" = $6, "error" = $7 WHERE "id" = $8`, c.sensor, c.Created, c.Started, c.FirstDiscovery, c.LastDiscovery, c.Finished, c.Error, c.id)

	return err
}

// Save inserts or updates record
func (c *Configuresensortask) Save(ctx context.Context, db models.DBInterface) error {
	if c.existsInDB {
		return c.update(ctx, db)
	}

	return c.insert(ctx, db)
}

// Delete removes row from database
func (c *Configuresensortask) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "center_configuresensortask" WHERE "id" = $1`, c.id)

	c.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (cl ConfiguresensortaskList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts ConfiguresensortaskList

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
	vaa := make([]any, 0, 7*len(inserts))
	offs := 1
	for _, c := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3, offs+4, offs+5, offs+6))
		vaa = append(vaa, c.sensor, c.Created, c.Started, c.FirstDiscovery, c.LastDiscovery, c.Finished, c.Error)
		offs += 7
	}

	qs := `INSERT INTO "center_configuresensortask" ("sensor_id", "created", "started", "first_discovery", "last_discovery", "finished", "error") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "id"`
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
