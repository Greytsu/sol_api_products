// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Bundle is an object representing the database table.
type Bundle struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CompanyID  int       `boil:"company_id" json:"company_id" toml:"company_id" yaml:"company_id"`
	Reference  string    `boil:"reference" json:"reference" toml:"reference" yaml:"reference"`
	Name       string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Price      float64   `boil:"price" json:"price" toml:"price" yaml:"price"`
	CreateTime time.Time `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime time.Time `boil:"update_time" json:"update_time" toml:"update_time" yaml:"update_time"`

	R *bundleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bundleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BundleColumns = struct {
	ID         string
	CompanyID  string
	Reference  string
	Name       string
	Price      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	CompanyID:  "company_id",
	Reference:  "reference",
	Name:       "name",
	Price:      "price",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

var BundleTableColumns = struct {
	ID         string
	CompanyID  string
	Reference  string
	Name       string
	Price      string
	CreateTime string
	UpdateTime string
}{
	ID:         "bundle.id",
	CompanyID:  "bundle.company_id",
	Reference:  "bundle.reference",
	Name:       "bundle.name",
	Price:      "bundle.price",
	CreateTime: "bundle.create_time",
	UpdateTime: "bundle.update_time",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BundleWhere = struct {
	ID         whereHelperint
	CompanyID  whereHelperint
	Reference  whereHelperstring
	Name       whereHelperstring
	Price      whereHelperfloat64
	CreateTime whereHelpertime_Time
	UpdateTime whereHelpertime_Time
}{
	ID:         whereHelperint{field: "[products].[bundle].[id]"},
	CompanyID:  whereHelperint{field: "[products].[bundle].[company_id]"},
	Reference:  whereHelperstring{field: "[products].[bundle].[reference]"},
	Name:       whereHelperstring{field: "[products].[bundle].[name]"},
	Price:      whereHelperfloat64{field: "[products].[bundle].[price]"},
	CreateTime: whereHelpertime_Time{field: "[products].[bundle].[create_time]"},
	UpdateTime: whereHelpertime_Time{field: "[products].[bundle].[update_time]"},
}

// BundleRels is where relationship names are stored.
var BundleRels = struct {
	FKBundleBundleElements string
}{
	FKBundleBundleElements: "FKBundleBundleElements",
}

// bundleR is where relationships are stored.
type bundleR struct {
	FKBundleBundleElements BundleElementSlice `boil:"FKBundleBundleElements" json:"FKBundleBundleElements" toml:"FKBundleBundleElements" yaml:"FKBundleBundleElements"`
}

// NewStruct creates a new relationship struct
func (*bundleR) NewStruct() *bundleR {
	return &bundleR{}
}

func (r *bundleR) GetFKBundleBundleElements() BundleElementSlice {
	if r == nil {
		return nil
	}
	return r.FKBundleBundleElements
}

// bundleL is where Load methods for each relationship are stored.
type bundleL struct{}

var (
	bundleAllColumns            = []string{"id", "company_id", "reference", "name", "price", "create_time", "update_time"}
	bundleColumnsWithoutDefault = []string{"company_id", "reference", "name", "price"}
	bundleColumnsWithDefault    = []string{"id", "create_time", "update_time"}
	bundlePrimaryKeyColumns     = []string{"id"}
	bundleGeneratedColumns      = []string{"id"}
)

type (
	// BundleSlice is an alias for a slice of pointers to Bundle.
	// This should almost always be used instead of []Bundle.
	BundleSlice []*Bundle
	// BundleHook is the signature for custom Bundle hook methods
	BundleHook func(context.Context, boil.ContextExecutor, *Bundle) error

	bundleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bundleType                 = reflect.TypeOf(&Bundle{})
	bundleMapping              = queries.MakeStructMapping(bundleType)
	bundlePrimaryKeyMapping, _ = queries.BindMapping(bundleType, bundleMapping, bundlePrimaryKeyColumns)
	bundleInsertCacheMut       sync.RWMutex
	bundleInsertCache          = make(map[string]insertCache)
	bundleUpdateCacheMut       sync.RWMutex
	bundleUpdateCache          = make(map[string]updateCache)
	bundleUpsertCacheMut       sync.RWMutex
	bundleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bundleAfterSelectHooks []BundleHook

var bundleBeforeInsertHooks []BundleHook
var bundleAfterInsertHooks []BundleHook

var bundleBeforeUpdateHooks []BundleHook
var bundleAfterUpdateHooks []BundleHook

var bundleBeforeDeleteHooks []BundleHook
var bundleAfterDeleteHooks []BundleHook

var bundleBeforeUpsertHooks []BundleHook
var bundleAfterUpsertHooks []BundleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Bundle) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Bundle) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Bundle) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Bundle) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Bundle) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Bundle) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Bundle) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Bundle) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Bundle) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bundleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBundleHook registers your hook function for all future operations.
func AddBundleHook(hookPoint boil.HookPoint, bundleHook BundleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		bundleAfterSelectHooks = append(bundleAfterSelectHooks, bundleHook)
	case boil.BeforeInsertHook:
		bundleBeforeInsertHooks = append(bundleBeforeInsertHooks, bundleHook)
	case boil.AfterInsertHook:
		bundleAfterInsertHooks = append(bundleAfterInsertHooks, bundleHook)
	case boil.BeforeUpdateHook:
		bundleBeforeUpdateHooks = append(bundleBeforeUpdateHooks, bundleHook)
	case boil.AfterUpdateHook:
		bundleAfterUpdateHooks = append(bundleAfterUpdateHooks, bundleHook)
	case boil.BeforeDeleteHook:
		bundleBeforeDeleteHooks = append(bundleBeforeDeleteHooks, bundleHook)
	case boil.AfterDeleteHook:
		bundleAfterDeleteHooks = append(bundleAfterDeleteHooks, bundleHook)
	case boil.BeforeUpsertHook:
		bundleBeforeUpsertHooks = append(bundleBeforeUpsertHooks, bundleHook)
	case boil.AfterUpsertHook:
		bundleAfterUpsertHooks = append(bundleAfterUpsertHooks, bundleHook)
	}
}

// One returns a single bundle record from the query.
func (q bundleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Bundle, error) {
	o := &Bundle{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bundle")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Bundle records from the query.
func (q bundleQuery) All(ctx context.Context, exec boil.ContextExecutor) (BundleSlice, error) {
	var o []*Bundle

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Bundle slice")
	}

	if len(bundleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Bundle records in the query.
func (q bundleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bundle rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bundleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bundle exists")
	}

	return count > 0, nil
}

// FKBundleBundleElements retrieves all the bundle_element's BundleElements with an executor via fk_bundle_id column.
func (o *Bundle) FKBundleBundleElements(mods ...qm.QueryMod) bundleElementQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("[products].[bundle_element].[fk_bundle_id]=?", o.ID),
	)

	return BundleElements(queryMods...)
}

// LoadFKBundleBundleElements allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bundleL) LoadFKBundleBundleElements(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBundle interface{}, mods queries.Applicator) error {
	var slice []*Bundle
	var object *Bundle

	if singular {
		var ok bool
		object, ok = maybeBundle.(*Bundle)
		if !ok {
			object = new(Bundle)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBundle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBundle))
			}
		}
	} else {
		s, ok := maybeBundle.(*[]*Bundle)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBundle)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBundle))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bundleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bundleR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`products.bundle_element`),
		qm.WhereIn(`products.bundle_element.fk_bundle_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load bundle_element")
	}

	var resultSlice []*BundleElement
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice bundle_element")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on bundle_element")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bundle_element")
	}

	if len(bundleElementAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.FKBundleBundleElements = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &bundleElementR{}
			}
			foreign.R.FKBundle = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.FKBundleID {
				local.R.FKBundleBundleElements = append(local.R.FKBundleBundleElements, foreign)
				if foreign.R == nil {
					foreign.R = &bundleElementR{}
				}
				foreign.R.FKBundle = local
				break
			}
		}
	}

	return nil
}

// AddFKBundleBundleElements adds the given related objects to the existing relationships
// of the bundle, optionally inserting them as new records.
// Appends related to o.R.FKBundleBundleElements.
// Sets related.R.FKBundle appropriately.
func (o *Bundle) AddFKBundleBundleElements(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BundleElement) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.FKBundleID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE [products].[bundle_element] SET %s WHERE %s",
				strmangle.SetParamNames("[", "]", 1, []string{"fk_bundle_id"}),
				strmangle.WhereClause("[", "]", 2, bundleElementPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.FKBundleID = o.ID
		}
	}

	if o.R == nil {
		o.R = &bundleR{
			FKBundleBundleElements: related,
		}
	} else {
		o.R.FKBundleBundleElements = append(o.R.FKBundleBundleElements, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &bundleElementR{
				FKBundle: o,
			}
		} else {
			rel.R.FKBundle = o
		}
	}
	return nil
}

// Bundles retrieves all the records using an executor.
func Bundles(mods ...qm.QueryMod) bundleQuery {
	mods = append(mods, qm.From("[products].[bundle]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[products].[bundle].*"})
	}

	return bundleQuery{q}
}

// FindBundle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBundle(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Bundle, error) {
	bundleObj := &Bundle{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [products].[bundle] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bundleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bundle")
	}

	if err = bundleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return bundleObj, err
	}

	return bundleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Bundle) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bundle provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bundleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bundleInsertCacheMut.RLock()
	cache, cached := bundleInsertCache[key]
	bundleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bundleAllColumns,
			bundleColumnsWithDefault,
			bundleColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, bundleGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(bundleType, bundleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bundleType, bundleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [products].[bundle] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [products].[bundle] %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryOutput = fmt.Sprintf("OUTPUT INSERTED.[%s] ", strings.Join(returnColumns, "],INSERTED.["))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into bundle")
	}

	if !cached {
		bundleInsertCacheMut.Lock()
		bundleInsertCache[key] = cache
		bundleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Bundle.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Bundle) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bundleUpdateCacheMut.RLock()
	cache, cached := bundleUpdateCache[key]
	bundleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bundleAllColumns,
			bundlePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, bundleGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bundle, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [products].[bundle] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, bundlePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bundleType, bundleMapping, append(wl, bundlePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update bundle row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bundle")
	}

	if !cached {
		bundleUpdateCacheMut.Lock()
		bundleUpdateCache[key] = cache
		bundleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bundleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bundle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bundle")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BundleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bundlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [products].[bundle] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bundlePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bundle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bundle")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Bundle) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bundle provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bundleColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	bundleUpsertCacheMut.RLock()
	cache, cached := bundleUpsertCache[key]
	bundleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bundleAllColumns,
			bundleColumnsWithDefault,
			bundleColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, bundleGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(bundlePrimaryKeyColumns, v) && strmangle.ContainsAny(bundleColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert bundle, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			bundleAllColumns,
			bundlePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, bundleGeneratedColumns)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert bundle, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[products].[bundle]", bundlePrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(bundlePrimaryKeyColumns))
		copy(whitelist, bundlePrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(bundleType, bundleMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bundleType, bundleMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // MSSQL doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert bundle")
	}

	if !cached {
		bundleUpsertCacheMut.Lock()
		bundleUpsertCache[key] = cache
		bundleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Bundle record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Bundle) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Bundle provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bundlePrimaryKeyMapping)
	sql := "DELETE FROM [products].[bundle] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bundle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bundle")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bundleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bundleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bundle")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bundle")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BundleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bundleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bundlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [products].[bundle] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bundlePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bundle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bundle")
	}

	if len(bundleAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Bundle) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBundle(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BundleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BundleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bundlePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [products].[bundle].* FROM [products].[bundle] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bundlePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BundleSlice")
	}

	*o = slice

	return nil
}

// BundleExists checks if the Bundle row exists.
func BundleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [products].[bundle] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bundle exists")
	}

	return exists, nil
}

// Exists checks if the Bundle row exists.
func (o *Bundle) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BundleExists(ctx, exec, o.ID)
}
