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

// Stock is an object representing the database table.
type Stock struct {
	ID            int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CompanyID     int       `boil:"company_id" json:"company_id" toml:"company_id" yaml:"company_id"`
	FKWarehouseID int       `boil:"fk_warehouse_id" json:"fk_warehouse_id" toml:"fk_warehouse_id" yaml:"fk_warehouse_id"`
	FKVariantID   int       `boil:"fk_variant_id" json:"fk_variant_id" toml:"fk_variant_id" yaml:"fk_variant_id"`
	Quantity      int       `boil:"quantity" json:"quantity" toml:"quantity" yaml:"quantity"`
	CreateTime    time.Time `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime    time.Time `boil:"update_time" json:"update_time" toml:"update_time" yaml:"update_time"`

	R *stockR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StockColumns = struct {
	ID            string
	CompanyID     string
	FKWarehouseID string
	FKVariantID   string
	Quantity      string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	CompanyID:     "company_id",
	FKWarehouseID: "fk_warehouse_id",
	FKVariantID:   "fk_variant_id",
	Quantity:      "quantity",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

var StockTableColumns = struct {
	ID            string
	CompanyID     string
	FKWarehouseID string
	FKVariantID   string
	Quantity      string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "stock.id",
	CompanyID:     "stock.company_id",
	FKWarehouseID: "stock.fk_warehouse_id",
	FKVariantID:   "stock.fk_variant_id",
	Quantity:      "stock.quantity",
	CreateTime:    "stock.create_time",
	UpdateTime:    "stock.update_time",
}

// Generated where

var StockWhere = struct {
	ID            whereHelperint
	CompanyID     whereHelperint
	FKWarehouseID whereHelperint
	FKVariantID   whereHelperint
	Quantity      whereHelperint
	CreateTime    whereHelpertime_Time
	UpdateTime    whereHelpertime_Time
}{
	ID:            whereHelperint{field: "[products].[stock].[id]"},
	CompanyID:     whereHelperint{field: "[products].[stock].[company_id]"},
	FKWarehouseID: whereHelperint{field: "[products].[stock].[fk_warehouse_id]"},
	FKVariantID:   whereHelperint{field: "[products].[stock].[fk_variant_id]"},
	Quantity:      whereHelperint{field: "[products].[stock].[quantity]"},
	CreateTime:    whereHelpertime_Time{field: "[products].[stock].[create_time]"},
	UpdateTime:    whereHelpertime_Time{field: "[products].[stock].[update_time]"},
}

// StockRels is where relationship names are stored.
var StockRels = struct {
	FKVariant   string
	FKWarehouse string
}{
	FKVariant:   "FKVariant",
	FKWarehouse: "FKWarehouse",
}

// stockR is where relationships are stored.
type stockR struct {
	FKVariant   *Variant   `boil:"FKVariant" json:"FKVariant" toml:"FKVariant" yaml:"FKVariant"`
	FKWarehouse *Warehouse `boil:"FKWarehouse" json:"FKWarehouse" toml:"FKWarehouse" yaml:"FKWarehouse"`
}

// NewStruct creates a new relationship struct
func (*stockR) NewStruct() *stockR {
	return &stockR{}
}

func (r *stockR) GetFKVariant() *Variant {
	if r == nil {
		return nil
	}
	return r.FKVariant
}

func (r *stockR) GetFKWarehouse() *Warehouse {
	if r == nil {
		return nil
	}
	return r.FKWarehouse
}

// stockL is where Load methods for each relationship are stored.
type stockL struct{}

var (
	stockAllColumns            = []string{"id", "company_id", "fk_warehouse_id", "fk_variant_id", "quantity", "create_time", "update_time"}
	stockColumnsWithoutDefault = []string{"company_id", "fk_warehouse_id", "fk_variant_id", "quantity"}
	stockColumnsWithDefault    = []string{"id", "create_time", "update_time"}
	stockPrimaryKeyColumns     = []string{"id"}
	stockGeneratedColumns      = []string{"id"}
)

type (
	// StockSlice is an alias for a slice of pointers to Stock.
	// This should almost always be used instead of []Stock.
	StockSlice []*Stock
	// StockHook is the signature for custom Stock hook methods
	StockHook func(context.Context, boil.ContextExecutor, *Stock) error

	stockQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockType                 = reflect.TypeOf(&Stock{})
	stockMapping              = queries.MakeStructMapping(stockType)
	stockPrimaryKeyMapping, _ = queries.BindMapping(stockType, stockMapping, stockPrimaryKeyColumns)
	stockInsertCacheMut       sync.RWMutex
	stockInsertCache          = make(map[string]insertCache)
	stockUpdateCacheMut       sync.RWMutex
	stockUpdateCache          = make(map[string]updateCache)
	stockUpsertCacheMut       sync.RWMutex
	stockUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var stockAfterSelectHooks []StockHook

var stockBeforeInsertHooks []StockHook
var stockAfterInsertHooks []StockHook

var stockBeforeUpdateHooks []StockHook
var stockAfterUpdateHooks []StockHook

var stockBeforeDeleteHooks []StockHook
var stockAfterDeleteHooks []StockHook

var stockBeforeUpsertHooks []StockHook
var stockAfterUpsertHooks []StockHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Stock) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Stock) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Stock) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Stock) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Stock) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Stock) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Stock) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Stock) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Stock) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockHook registers your hook function for all future operations.
func AddStockHook(hookPoint boil.HookPoint, stockHook StockHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		stockAfterSelectHooks = append(stockAfterSelectHooks, stockHook)
	case boil.BeforeInsertHook:
		stockBeforeInsertHooks = append(stockBeforeInsertHooks, stockHook)
	case boil.AfterInsertHook:
		stockAfterInsertHooks = append(stockAfterInsertHooks, stockHook)
	case boil.BeforeUpdateHook:
		stockBeforeUpdateHooks = append(stockBeforeUpdateHooks, stockHook)
	case boil.AfterUpdateHook:
		stockAfterUpdateHooks = append(stockAfterUpdateHooks, stockHook)
	case boil.BeforeDeleteHook:
		stockBeforeDeleteHooks = append(stockBeforeDeleteHooks, stockHook)
	case boil.AfterDeleteHook:
		stockAfterDeleteHooks = append(stockAfterDeleteHooks, stockHook)
	case boil.BeforeUpsertHook:
		stockBeforeUpsertHooks = append(stockBeforeUpsertHooks, stockHook)
	case boil.AfterUpsertHook:
		stockAfterUpsertHooks = append(stockAfterUpsertHooks, stockHook)
	}
}

// One returns a single stock record from the query.
func (q stockQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Stock, error) {
	o := &Stock{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Stock records from the query.
func (q stockQuery) All(ctx context.Context, exec boil.ContextExecutor) (StockSlice, error) {
	var o []*Stock

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stock slice")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Stock records in the query.
func (q stockQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stockQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock exists")
	}

	return count > 0, nil
}

// FKVariant pointed to by the foreign key.
func (o *Stock) FKVariant(mods ...qm.QueryMod) variantQuery {
	queryMods := []qm.QueryMod{
		qm.Where("[id] = ?", o.FKVariantID),
	}

	queryMods = append(queryMods, mods...)

	return Variants(queryMods...)
}

// FKWarehouse pointed to by the foreign key.
func (o *Stock) FKWarehouse(mods ...qm.QueryMod) warehouseQuery {
	queryMods := []qm.QueryMod{
		qm.Where("[id] = ?", o.FKWarehouseID),
	}

	queryMods = append(queryMods, mods...)

	return Warehouses(queryMods...)
}

// LoadFKVariant allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stockL) LoadFKVariant(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStock interface{}, mods queries.Applicator) error {
	var slice []*Stock
	var object *Stock

	if singular {
		var ok bool
		object, ok = maybeStock.(*Stock)
		if !ok {
			object = new(Stock)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeStock)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeStock))
			}
		}
	} else {
		s, ok := maybeStock.(*[]*Stock)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeStock)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeStock))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stockR{}
		}
		args = append(args, object.FKVariantID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stockR{}
			}

			for _, a := range args {
				if a == obj.FKVariantID {
					continue Outer
				}
			}

			args = append(args, obj.FKVariantID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`products.variant`),
		qm.WhereIn(`products.variant.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Variant")
	}

	var resultSlice []*Variant
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Variant")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for variant")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for variant")
	}

	if len(variantAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.FKVariant = foreign
		if foreign.R == nil {
			foreign.R = &variantR{}
		}
		foreign.R.FKVariantStocks = append(foreign.R.FKVariantStocks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FKVariantID == foreign.ID {
				local.R.FKVariant = foreign
				if foreign.R == nil {
					foreign.R = &variantR{}
				}
				foreign.R.FKVariantStocks = append(foreign.R.FKVariantStocks, local)
				break
			}
		}
	}

	return nil
}

// LoadFKWarehouse allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stockL) LoadFKWarehouse(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStock interface{}, mods queries.Applicator) error {
	var slice []*Stock
	var object *Stock

	if singular {
		var ok bool
		object, ok = maybeStock.(*Stock)
		if !ok {
			object = new(Stock)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeStock)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeStock))
			}
		}
	} else {
		s, ok := maybeStock.(*[]*Stock)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeStock)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeStock))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stockR{}
		}
		args = append(args, object.FKWarehouseID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stockR{}
			}

			for _, a := range args {
				if a == obj.FKWarehouseID {
					continue Outer
				}
			}

			args = append(args, obj.FKWarehouseID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`products.warehouse`),
		qm.WhereIn(`products.warehouse.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Warehouse")
	}

	var resultSlice []*Warehouse
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Warehouse")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for warehouse")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for warehouse")
	}

	if len(warehouseAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.FKWarehouse = foreign
		if foreign.R == nil {
			foreign.R = &warehouseR{}
		}
		foreign.R.FKWarehouseStocks = append(foreign.R.FKWarehouseStocks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FKWarehouseID == foreign.ID {
				local.R.FKWarehouse = foreign
				if foreign.R == nil {
					foreign.R = &warehouseR{}
				}
				foreign.R.FKWarehouseStocks = append(foreign.R.FKWarehouseStocks, local)
				break
			}
		}
	}

	return nil
}

// SetFKVariant of the stock to the related item.
// Sets o.R.FKVariant to related.
// Adds o to related.R.FKVariantStocks.
func (o *Stock) SetFKVariant(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Variant) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE [products].[stock] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, []string{"fk_variant_id"}),
		strmangle.WhereClause("[", "]", 2, stockPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FKVariantID = related.ID
	if o.R == nil {
		o.R = &stockR{
			FKVariant: related,
		}
	} else {
		o.R.FKVariant = related
	}

	if related.R == nil {
		related.R = &variantR{
			FKVariantStocks: StockSlice{o},
		}
	} else {
		related.R.FKVariantStocks = append(related.R.FKVariantStocks, o)
	}

	return nil
}

// SetFKWarehouse of the stock to the related item.
// Sets o.R.FKWarehouse to related.
// Adds o to related.R.FKWarehouseStocks.
func (o *Stock) SetFKWarehouse(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Warehouse) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE [products].[stock] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, []string{"fk_warehouse_id"}),
		strmangle.WhereClause("[", "]", 2, stockPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FKWarehouseID = related.ID
	if o.R == nil {
		o.R = &stockR{
			FKWarehouse: related,
		}
	} else {
		o.R.FKWarehouse = related
	}

	if related.R == nil {
		related.R = &warehouseR{
			FKWarehouseStocks: StockSlice{o},
		}
	} else {
		related.R.FKWarehouseStocks = append(related.R.FKWarehouseStocks, o)
	}

	return nil
}

// Stocks retrieves all the records using an executor.
func Stocks(mods ...qm.QueryMod) stockQuery {
	mods = append(mods, qm.From("[products].[stock]"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"[products].[stock].*"})
	}

	return stockQuery{q}
}

// FindStock retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStock(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Stock, error) {
	stockObj := &Stock{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from [products].[stock] where [id]=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stockObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock")
	}

	if err = stockObj.doAfterSelectHooks(ctx, exec); err != nil {
		return stockObj, err
	}

	return stockObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Stock) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stock provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stockInsertCacheMut.RLock()
	cache, cached := stockInsertCache[key]
	stockInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stockAllColumns,
			stockColumnsWithDefault,
			stockColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, stockGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockType, stockMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO [products].[stock] ([%s]) %%sVALUES (%s)%%s", strings.Join(wl, "],["), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO [products].[stock] %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into stock")
	}

	if !cached {
		stockInsertCacheMut.Lock()
		stockInsertCache[key] = cache
		stockInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Stock.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Stock) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	stockUpdateCacheMut.RLock()
	cache, cached := stockUpdateCache[key]
	stockUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stockAllColumns,
			stockPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, stockGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stock, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE [products].[stock] SET %s WHERE %s",
			strmangle.SetParamNames("[", "]", 1, wl),
			strmangle.WhereClause("[", "]", len(wl)+1, stockPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, append(wl, stockPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update stock row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stock")
	}

	if !cached {
		stockUpdateCacheMut.Lock()
		stockUpdateCache[key] = cache
		stockUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q stockQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stock")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE [products].[stock] SET %s WHERE %s",
		strmangle.SetParamNames("[", "]", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, stockPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stock")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Stock) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stock provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockColumnsWithDefault, o)

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

	stockUpsertCacheMut.RLock()
	cache, cached := stockUpsertCache[key]
	stockUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stockAllColumns,
			stockColumnsWithDefault,
			stockColumnsWithoutDefault,
			nzDefaults,
		)

		insert = strmangle.SetComplement(insert, stockGeneratedColumns)

		for i, v := range insert {
			if strmangle.ContainsAny(stockPrimaryKeyColumns, v) && strmangle.ContainsAny(stockColumnsWithDefault, v) {
				insert = append(insert[:i], insert[i+1:]...)
			}
		}
		if len(insert) == 0 {
			return errors.New("models: unable to upsert stock, could not build insert column list")
		}

		update := updateColumns.UpdateColumnSet(
			stockAllColumns,
			stockPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, stockGeneratedColumns)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert stock, could not build update column list")
		}

		cache.query = buildUpsertQueryMSSQL(dialect, "[products].[stock]", stockPrimaryKeyColumns, update, insert, ret)

		whitelist := make([]string, len(stockPrimaryKeyColumns))
		copy(whitelist, stockPrimaryKeyColumns)
		whitelist = append(whitelist, update...)
		whitelist = append(whitelist, insert...)

		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockType, stockMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert stock")
	}

	if !cached {
		stockUpsertCacheMut.Lock()
		stockUpsertCache[key] = cache
		stockUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Stock record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stock) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Stock provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockPrimaryKeyMapping)
	sql := "DELETE FROM [products].[stock] WHERE [id]=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stock")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stockQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stockQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stock")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(stockBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM [products].[stock] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stockPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stock")
	}

	if len(stockAfterDeleteHooks) != 0 {
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
func (o *Stock) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStock(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StockSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT [products].[stock].* FROM [products].[stock] WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stockPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockSlice")
	}

	*o = slice

	return nil
}

// StockExists checks if the Stock row exists.
func StockExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select case when exists(select top(1) 1 from [products].[stock] where [id]=$1) then 1 else 0 end"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock exists")
	}

	return exists, nil
}

// Exists checks if the Stock row exists.
func (o *Stock) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return StockExists(ctx, exec, o.ID)
}
