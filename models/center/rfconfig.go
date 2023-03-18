// Code generated for Django model center.RFConfig. DO NOT EDIT.

/*
  Command used to generate:

  DJANGO_SETTINGS_MODULE=application.settings ../djan-go-rm/djan-go-rm.py --gomodule github.com/rkojedzinszky/thermo-center center heatcontrol

  https://github.com/rkojedzinszky/djan-go-rm
*/

package center

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/rkojedzinszky/thermo-center/models"
)

// Rfconfig mirrors model center.RFConfig
type Rfconfig struct {
	existsInDB bool

	id        int32
	RfChannel int32
	rfProfile int32
	NetworkId int32
	AesKey    string
}

// RfconfigList is a list of Rfconfig
type RfconfigList []*Rfconfig

// RfconfigQS represents a queryset for center.RFConfig
type RfconfigQS struct {
	condFragments models.AndFragment
	order         []string
	forClause     string
}

func (qs RfconfigQS) filter(c string, p interface{}) RfconfigQS {
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
func (qs RfconfigQS) Or(exprs ...RfconfigQS) RfconfigQS {
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

// GetID returns Rfconfig.ID
func (r *Rfconfig) GetID() int32 {
	return r.id
}

// IDEq filters for id being equal to argument
func (qs RfconfigQS) IDEq(v int32) RfconfigQS {
	return qs.filter(`"id" =`, v)
}

// IDNe filters for id being not equal to argument
func (qs RfconfigQS) IDNe(v int32) RfconfigQS {
	return qs.filter(`"id" <>`, v)
}

// IDLt filters for id being less than argument
func (qs RfconfigQS) IDLt(v int32) RfconfigQS {
	return qs.filter(`"id" <`, v)
}

// IDLe filters for id being less than or equal to argument
func (qs RfconfigQS) IDLe(v int32) RfconfigQS {
	return qs.filter(`"id" <=`, v)
}

// IDGt filters for id being greater than argument
func (qs RfconfigQS) IDGt(v int32) RfconfigQS {
	return qs.filter(`"id" >`, v)
}

// IDGe filters for id being greater than or equal to argument
func (qs RfconfigQS) IDGe(v int32) RfconfigQS {
	return qs.filter(`"id" >=`, v)
}

type inRfconfigid []interface{}

func (in inRfconfigid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfconfigQS) IDIn(values []int32) RfconfigQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inRfconfigid(vals),
	)

	return qs
}

type notinRfconfigid []interface{}

func (in notinRfconfigid) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfconfigQS) IDNotIn(values []int32) RfconfigQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinRfconfigid(vals),
	)

	return qs
}

// OrderByID sorts result by ID in ascending order
func (qs RfconfigQS) OrderByID() RfconfigQS {
	qs.order = append(qs.order, `"id"`)

	return qs
}

// OrderByIDDesc sorts result by ID in descending order
func (qs RfconfigQS) OrderByIDDesc() RfconfigQS {
	qs.order = append(qs.order, `"id" DESC`)

	return qs
}

// RfChannelEq filters for RfChannel being equal to argument
func (qs RfconfigQS) RfChannelEq(v int32) RfconfigQS {
	return qs.filter(`"rf_channel" =`, v)
}

// RfChannelNe filters for RfChannel being not equal to argument
func (qs RfconfigQS) RfChannelNe(v int32) RfconfigQS {
	return qs.filter(`"rf_channel" <>`, v)
}

// RfChannelLt filters for RfChannel being less than argument
func (qs RfconfigQS) RfChannelLt(v int32) RfconfigQS {
	return qs.filter(`"rf_channel" <`, v)
}

// RfChannelLe filters for RfChannel being less than or equal to argument
func (qs RfconfigQS) RfChannelLe(v int32) RfconfigQS {
	return qs.filter(`"rf_channel" <=`, v)
}

// RfChannelGt filters for RfChannel being greater than argument
func (qs RfconfigQS) RfChannelGt(v int32) RfconfigQS {
	return qs.filter(`"rf_channel" >`, v)
}

// RfChannelGe filters for RfChannel being greater than or equal to argument
func (qs RfconfigQS) RfChannelGe(v int32) RfconfigQS {
	return qs.filter(`"rf_channel" >=`, v)
}

type inRfconfigRfChannel []interface{}

func (in inRfconfigRfChannel) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"rf_channel" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfconfigQS) RfChannelIn(values []int32) RfconfigQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inRfconfigRfChannel(vals),
	)

	return qs
}

type notinRfconfigRfChannel []interface{}

func (in notinRfconfigRfChannel) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"rf_channel" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfconfigQS) RfChannelNotIn(values []int32) RfconfigQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinRfconfigRfChannel(vals),
	)

	return qs
}

// OrderByRfChannel sorts result by RfChannel in ascending order
func (qs RfconfigQS) OrderByRfChannel() RfconfigQS {
	qs.order = append(qs.order, `"rf_channel"`)

	return qs
}

// OrderByRfChannelDesc sorts result by RfChannel in descending order
func (qs RfconfigQS) OrderByRfChannelDesc() RfconfigQS {
	qs.order = append(qs.order, `"rf_channel" DESC`)

	return qs
}

// GetRfProfile returns Rfprofile
func (r *Rfconfig) GetRfProfile(ctx context.Context, db models.DBInterface) (*Rfprofile, error) {
	return RfprofileQS{}.IDEq(r.rfProfile).First(ctx, db)
}

// SetRfProfile sets foreign key pointer to Rfprofile
func (r *Rfconfig) SetRfProfile(ptr *Rfprofile) error {
	if ptr != nil {
		r.rfProfile = ptr.GetID()
	} else {
		return fmt.Errorf("Rfconfig.SetRfProfile: non-null field received null value")
	}

	return nil
}

// GetRfProfileRaw returns Rfconfig.RfProfile
func (r *Rfconfig) GetRfProfileRaw() int32 {
	return r.rfProfile
}

// RfProfileEq filters for rfProfile being equal to argument
func (qs RfconfigQS) RfProfileEq(v *Rfprofile) RfconfigQS {
	return qs.filter(`"rf_profile_id" =`, v.GetID())
}

// RfProfileRawEq filters for rfProfile being equal to raw argument
func (qs RfconfigQS) RfProfileRawEq(v int32) RfconfigQS {
	return qs.filter(`"rf_profile_id" =`, v)
}

type inRfconfigrfProfileRfprofile struct {
	qs RfprofileQS
}

func (in *inRfconfigrfProfileRfprofile) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	s, p := in.qs.QueryId(c)

	return `"rf_profile_id" IN (` + s + `)`, p
}

func (qs RfconfigQS) RfProfileIn(oqs RfprofileQS) RfconfigQS {
	qs.condFragments = append(
		qs.condFragments,
		&inRfconfigrfProfileRfprofile{
			qs: oqs,
		},
	)

	return qs
}

// OrderByRfProfile sorts result by RfProfile in ascending order
func (qs RfconfigQS) OrderByRfProfile() RfconfigQS {
	qs.order = append(qs.order, `"rf_profile_id"`)

	return qs
}

// OrderByRfProfileDesc sorts result by RfProfile in descending order
func (qs RfconfigQS) OrderByRfProfileDesc() RfconfigQS {
	qs.order = append(qs.order, `"rf_profile_id" DESC`)

	return qs
}

// NetworkIdEq filters for NetworkId being equal to argument
func (qs RfconfigQS) NetworkIdEq(v int32) RfconfigQS {
	return qs.filter(`"network_id" =`, v)
}

// NetworkIdNe filters for NetworkId being not equal to argument
func (qs RfconfigQS) NetworkIdNe(v int32) RfconfigQS {
	return qs.filter(`"network_id" <>`, v)
}

// NetworkIdLt filters for NetworkId being less than argument
func (qs RfconfigQS) NetworkIdLt(v int32) RfconfigQS {
	return qs.filter(`"network_id" <`, v)
}

// NetworkIdLe filters for NetworkId being less than or equal to argument
func (qs RfconfigQS) NetworkIdLe(v int32) RfconfigQS {
	return qs.filter(`"network_id" <=`, v)
}

// NetworkIdGt filters for NetworkId being greater than argument
func (qs RfconfigQS) NetworkIdGt(v int32) RfconfigQS {
	return qs.filter(`"network_id" >`, v)
}

// NetworkIdGe filters for NetworkId being greater than or equal to argument
func (qs RfconfigQS) NetworkIdGe(v int32) RfconfigQS {
	return qs.filter(`"network_id" >=`, v)
}

type inRfconfigNetworkId []interface{}

func (in inRfconfigNetworkId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"network_id" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfconfigQS) NetworkIdIn(values []int32) RfconfigQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inRfconfigNetworkId(vals),
	)

	return qs
}

type notinRfconfigNetworkId []interface{}

func (in notinRfconfigNetworkId) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"network_id" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfconfigQS) NetworkIdNotIn(values []int32) RfconfigQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinRfconfigNetworkId(vals),
	)

	return qs
}

// OrderByNetworkId sorts result by NetworkId in ascending order
func (qs RfconfigQS) OrderByNetworkId() RfconfigQS {
	qs.order = append(qs.order, `"network_id"`)

	return qs
}

// OrderByNetworkIdDesc sorts result by NetworkId in descending order
func (qs RfconfigQS) OrderByNetworkIdDesc() RfconfigQS {
	qs.order = append(qs.order, `"network_id" DESC`)

	return qs
}

// AesKeyEq filters for AesKey being equal to argument
func (qs RfconfigQS) AesKeyEq(v string) RfconfigQS {
	return qs.filter(`"aes_key" =`, v)
}

// AesKeyNe filters for AesKey being not equal to argument
func (qs RfconfigQS) AesKeyNe(v string) RfconfigQS {
	return qs.filter(`"aes_key" <>`, v)
}

// AesKeyLt filters for AesKey being less than argument
func (qs RfconfigQS) AesKeyLt(v string) RfconfigQS {
	return qs.filter(`"aes_key" <`, v)
}

// AesKeyLe filters for AesKey being less than or equal to argument
func (qs RfconfigQS) AesKeyLe(v string) RfconfigQS {
	return qs.filter(`"aes_key" <=`, v)
}

// AesKeyGt filters for AesKey being greater than argument
func (qs RfconfigQS) AesKeyGt(v string) RfconfigQS {
	return qs.filter(`"aes_key" >`, v)
}

// AesKeyGe filters for AesKey being greater than or equal to argument
func (qs RfconfigQS) AesKeyGe(v string) RfconfigQS {
	return qs.filter(`"aes_key" >=`, v)
}

type inRfconfigAesKey []interface{}

func (in inRfconfigAesKey) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"aes_key" IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfconfigQS) AesKeyIn(values []string) RfconfigQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		inRfconfigAesKey(vals),
	)

	return qs
}

type notinRfconfigAesKey []interface{}

func (in notinRfconfigAesKey) GetConditionFragment(c *models.PositionalCounter) (string, []interface{}) {
	if len(in) == 0 {
		return `false`, nil
	}

	var params []string
	for range in {
		params = append(params, c.Get())
	}

	return `"aes_key" NOT IN (` + strings.Join(params, ", ") + `)`, in
}

func (qs RfconfigQS) AesKeyNotIn(values []string) RfconfigQS {
	var vals []interface{}
	for _, v := range values {
		vals = append(vals, v)
	}

	qs.condFragments = append(
		qs.condFragments,
		notinRfconfigAesKey(vals),
	)

	return qs
}

// OrderByAesKey sorts result by AesKey in ascending order
func (qs RfconfigQS) OrderByAesKey() RfconfigQS {
	qs.order = append(qs.order, `"aes_key"`)

	return qs
}

// OrderByAesKeyDesc sorts result by AesKey in descending order
func (qs RfconfigQS) OrderByAesKeyDesc() RfconfigQS {
	qs.order = append(qs.order, `"aes_key" DESC`)

	return qs
}

// OrderByRandom randomizes result
func (qs RfconfigQS) OrderByRandom() RfconfigQS {
	qs.order = append(qs.order, `random()`)

	return qs
}

// ForUpdate marks the queryset to use FOR UPDATE clause
func (qs RfconfigQS) ForUpdate() RfconfigQS {
	qs.forClause = " FOR UPDATE"

	return qs
}

// ForUpdateNowait marks the queryset to use FOR UPDATE NOWAIT clause
func (qs RfconfigQS) ForUpdateNowait() RfconfigQS {
	qs.forClause = " FOR UPDATE NOWAIT"

	return qs
}

// ForUpdateSkipLocked marks the queryset to use FOR UPDATE SKIP LOCKED clause
func (qs RfconfigQS) ForUpdateSkipLocked() RfconfigQS {
	qs.forClause = " FOR UPDATE SKIP LOCKED"

	return qs
}

// ClearForUpdate clears FOR UPDATE clause set on queryset
func (qs RfconfigQS) ClearForUpdate() RfconfigQS {
	qs.forClause = ""

	return qs
}

func (qs RfconfigQS) whereClause(c *models.PositionalCounter) (string, []interface{}) {
	if len(qs.condFragments) == 0 {
		return "", nil
	}

	cond, params := qs.condFragments.GetConditionFragment(c)

	return " WHERE " + cond, params
}

func (qs RfconfigQS) orderByClause() string {
	if len(qs.order) == 0 {
		return ""
	}

	return " ORDER BY " + strings.Join(qs.order, ", ")
}

func (qs RfconfigQS) queryFull() (string, []interface{}) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s += qs.orderByClause()
	s += qs.forClause

	return `SELECT "id", "rf_channel", "rf_profile_id", "network_id", "aes_key" FROM "center_rfconfig"` + s, p
}

// QueryId returns statement and parameters suitable for embedding in IN clause
func (qs RfconfigQS) QueryId(c *models.PositionalCounter) (string, []interface{}) {
	s, p := qs.whereClause(c)

	return `SELECT "id" FROM "center_rfconfig"` + s, p
}

// Count returns the number of rows matching queryset filters
func (qs RfconfigQS) Count(ctx context.Context, db models.DBInterface) (count int, err error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)

	row := db.QueryRow(ctx, `SELECT COUNT("id") FROM "center_rfconfig"`+s, p...)

	err = row.Scan(&count)

	return
}

// All returns all rows matching queryset filters
func (qs RfconfigQS) All(ctx context.Context, db models.DBInterface) (RfconfigList, error) {
	s, p := qs.queryFull()

	rows, err := db.Query(ctx, s, p...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret RfconfigList
	for rows.Next() {
		obj := Rfconfig{existsInDB: true}
		if err = rows.Scan(&obj.id, &obj.RfChannel, &obj.rfProfile, &obj.NetworkId, &obj.AesKey); err != nil {
			return nil, err
		}
		ret = append(ret, &obj)
	}

	return ret, nil
}

// First returns the first row matching queryset filters, others are discarded
func (qs RfconfigQS) First(ctx context.Context, db models.DBInterface) (*Rfconfig, error) {
	s, p := qs.queryFull()

	s += " LIMIT 1"

	row := db.QueryRow(ctx, s, p...)

	obj := Rfconfig{existsInDB: true}
	err := row.Scan(&obj.id, &obj.RfChannel, &obj.rfProfile, &obj.NetworkId, &obj.AesKey)
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
func (qs RfconfigQS) Delete(ctx context.Context, db models.DBInterface) (int64, error) {
	c := &models.PositionalCounter{}

	s, p := qs.whereClause(c)
	s = `DELETE FROM "center_rfconfig"` + s

	result, err := db.Exec(ctx, s, p...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// Update returns an Update queryset inheriting all the filter conditions, which then can be
// used to specify columns to be updated. At the end, .Exec() must be called to do the real operation.
func (qs RfconfigQS) Update() RfconfigUpdateQS {
	return RfconfigUpdateQS{condFragments: qs.condFragments}
}

// RfconfigUpdateQS represents an updated queryset for center.RFConfig
type RfconfigUpdateQS struct {
	updates       []models.ConditionFragment
	condFragments []models.ConditionFragment
}

func (uqs RfconfigUpdateQS) update(c string, v interface{}) RfconfigUpdateQS {
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
func (uqs RfconfigUpdateQS) SetID(v int32) RfconfigUpdateQS {
	return uqs.update(`"id"`, v)
}

// SetRfChannel sets RfChannel to the given value
func (uqs RfconfigUpdateQS) SetRfChannel(v int32) RfconfigUpdateQS {
	return uqs.update(`"rf_channel"`, v)
}

// SetRfProfile sets foreign key pointer to Rfprofile
func (uqs RfconfigUpdateQS) SetRfProfile(ptr *Rfprofile) RfconfigUpdateQS {
	if ptr != nil {
		return uqs.update(`"rf_profile_id"`, ptr.GetID())
	}

	return uqs.update(`"rf_profile_id"`, nil)
} // SetNetworkId sets NetworkId to the given value
func (uqs RfconfigUpdateQS) SetNetworkId(v int32) RfconfigUpdateQS {
	return uqs.update(`"network_id"`, v)
}

// SetAesKey sets AesKey to the given value
func (uqs RfconfigUpdateQS) SetAesKey(v string) RfconfigUpdateQS {
	return uqs.update(`"aes_key"`, v)
}

// Exec executes the update operation
func (uqs RfconfigUpdateQS) Exec(ctx context.Context, db models.DBInterface) (int64, error) {
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

	ws, wp := RfconfigQS{condFragments: uqs.condFragments}.whereClause(c)

	st := `UPDATE "center_rfconfig" SET ` + strings.Join(sets, ", ") + ws

	params = append(params, wp...)

	result, err := db.Exec(ctx, st, params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

// insert operation
func (r *Rfconfig) insert(ctx context.Context, db models.DBInterface) error {
	row := db.QueryRow(ctx, `INSERT INTO "center_rfconfig" ("rf_channel", "rf_profile_id", "network_id", "aes_key") VALUES ($1, $2, $3, $4) RETURNING "id"`, r.RfChannel, r.rfProfile, r.NetworkId, r.AesKey)

	if err := row.Scan(&r.id); err != nil {
		return err
	}

	r.existsInDB = true

	return nil
}

// update operation
func (r *Rfconfig) update(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `UPDATE "center_rfconfig" SET "rf_channel" = $1, "rf_profile_id" = $2, "network_id" = $3, "aes_key" = $4 WHERE "id" = $5`, r.RfChannel, r.rfProfile, r.NetworkId, r.AesKey, r.id)

	return err
}

// Save inserts or updates record
func (r *Rfconfig) Save(ctx context.Context, db models.DBInterface) error {
	if r.existsInDB {
		return r.update(ctx, db)
	}

	return r.insert(ctx, db)
}

// Delete removes row from database
func (r *Rfconfig) Delete(ctx context.Context, db models.DBInterface) error {
	_, err := db.Exec(ctx, `DELETE FROM "center_rfconfig" WHERE "id" = $1`, r.id)

	r.existsInDB = false

	return err
}

// Save saves all elements, optimizing inserts in a batch
func (rl RfconfigList) Save(ctx context.Context, db models.DBInterface) error {
	var inserts RfconfigList

	for _, r := range rl {
		if r.existsInDB {
			if err := r.update(ctx, db); err != nil {
				return err
			}
		} else {
			inserts = append(inserts, r)
		}
	}

	if len(inserts) == 0 {
		return nil
	}

	vva := make([]string, 0, len(inserts))
	vaa := make([]any, 0, 4*len(inserts))
	offs := 1
	for _, r := range inserts {
		vva = append(vva, fmt.Sprintf("($%d, $%d, $%d, $%d)", offs+0, offs+1, offs+2, offs+3))
		vaa = append(vaa, r.RfChannel, r.rfProfile, r.NetworkId, r.AesKey)
		offs += 4
	}

	qs := `INSERT INTO "center_rfconfig" ("rf_channel", "rf_profile_id", "network_id", "aes_key") VALUES ` + strings.Join(vva, ", ") + ` RETURNING "id"`
	rows, err := db.Query(ctx, qs, vaa...)

	if err != nil {
		return err
	}
	defer rows.Close()

	for _, r := range inserts {
		if !rows.Next() {
			return rows.Err()
		}

		if err := rows.Scan(&r.id); err != nil {
			return err
		}

		r.existsInDB = true
	}

	return nil
}
