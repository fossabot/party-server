// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"github.com/vattle/sqlboiler/types"
)

// Guest is an object representing the database table.
type Guest struct {
	ID   int        `boil:"id" json:"id" toml:"id" yaml:"id"`
	Data types.JSON `boil:"data" json:"data" toml:"data" yaml:"data"`

	R *guestR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L guestL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// guestR is where relationships are stored.
type guestR struct {
}

// guestL is where Load methods for each relationship are stored.
type guestL struct{}

var (
	guestColumns               = []string{"id", "data"}
	guestColumnsWithoutDefault = []string{"data"}
	guestColumnsWithDefault    = []string{"id"}
	guestPrimaryKeyColumns     = []string{"id"}
)

type (
	// GuestSlice is an alias for a slice of pointers to Guest.
	// This should generally be used opposed to []Guest.
	GuestSlice []*Guest
	// GuestHook is the signature for custom Guest hook methods
	GuestHook func(boil.Executor, *Guest) error

	guestQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	guestType                 = reflect.TypeOf(&Guest{})
	guestMapping              = queries.MakeStructMapping(guestType)
	guestPrimaryKeyMapping, _ = queries.BindMapping(guestType, guestMapping, guestPrimaryKeyColumns)
	guestInsertCacheMut       sync.RWMutex
	guestInsertCache          = make(map[string]insertCache)
	guestUpdateCacheMut       sync.RWMutex
	guestUpdateCache          = make(map[string]updateCache)
	guestUpsertCacheMut       sync.RWMutex
	guestUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var guestBeforeInsertHooks []GuestHook
var guestBeforeUpdateHooks []GuestHook
var guestBeforeDeleteHooks []GuestHook
var guestBeforeUpsertHooks []GuestHook

var guestAfterInsertHooks []GuestHook
var guestAfterSelectHooks []GuestHook
var guestAfterUpdateHooks []GuestHook
var guestAfterDeleteHooks []GuestHook
var guestAfterUpsertHooks []GuestHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Guest) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range guestBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Guest) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range guestBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Guest) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range guestBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Guest) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range guestBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Guest) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range guestAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Guest) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range guestAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Guest) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range guestAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Guest) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range guestAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Guest) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range guestAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGuestHook registers your hook function for all future operations.
func AddGuestHook(hookPoint boil.HookPoint, guestHook GuestHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		guestBeforeInsertHooks = append(guestBeforeInsertHooks, guestHook)
	case boil.BeforeUpdateHook:
		guestBeforeUpdateHooks = append(guestBeforeUpdateHooks, guestHook)
	case boil.BeforeDeleteHook:
		guestBeforeDeleteHooks = append(guestBeforeDeleteHooks, guestHook)
	case boil.BeforeUpsertHook:
		guestBeforeUpsertHooks = append(guestBeforeUpsertHooks, guestHook)
	case boil.AfterInsertHook:
		guestAfterInsertHooks = append(guestAfterInsertHooks, guestHook)
	case boil.AfterSelectHook:
		guestAfterSelectHooks = append(guestAfterSelectHooks, guestHook)
	case boil.AfterUpdateHook:
		guestAfterUpdateHooks = append(guestAfterUpdateHooks, guestHook)
	case boil.AfterDeleteHook:
		guestAfterDeleteHooks = append(guestAfterDeleteHooks, guestHook)
	case boil.AfterUpsertHook:
		guestAfterUpsertHooks = append(guestAfterUpsertHooks, guestHook)
	}
}

// OneP returns a single guest record from the query, and panics on error.
func (q guestQuery) OneP() *Guest {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single guest record from the query.
func (q guestQuery) One() (*Guest, error) {
	o := &Guest{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for guest")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Guest records from the query, and panics on error.
func (q guestQuery) AllP() GuestSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Guest records from the query.
func (q guestQuery) All() (GuestSlice, error) {
	var o GuestSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Guest slice")
	}

	if len(guestAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Guest records in the query, and panics on error.
func (q guestQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Guest records in the query.
func (q guestQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count guest rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q guestQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q guestQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if guest exists")
	}

	return count > 0, nil
}

// GuestsG retrieves all records.
func GuestsG(mods ...qm.QueryMod) guestQuery {
	return Guests(boil.GetDB(), mods...)
}

// Guests retrieves all the records using an executor.
func Guests(exec boil.Executor, mods ...qm.QueryMod) guestQuery {
	mods = append(mods, qm.From("\"guest\""))
	return guestQuery{NewQuery(exec, mods...)}
}

// FindGuestG retrieves a single record by ID.
func FindGuestG(id int, selectCols ...string) (*Guest, error) {
	return FindGuest(boil.GetDB(), id, selectCols...)
}

// FindGuestGP retrieves a single record by ID, and panics on error.
func FindGuestGP(id int, selectCols ...string) *Guest {
	retobj, err := FindGuest(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindGuest retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGuest(exec boil.Executor, id int, selectCols ...string) (*Guest, error) {
	guestObj := &Guest{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"guest\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(guestObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from guest")
	}

	return guestObj, nil
}

// FindGuestP retrieves a single record by ID with an executor, and panics on error.
func FindGuestP(exec boil.Executor, id int, selectCols ...string) *Guest {
	retobj, err := FindGuest(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Guest) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Guest) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Guest) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Guest) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no guest provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(guestColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	guestInsertCacheMut.RLock()
	cache, cached := guestInsertCache[key]
	guestInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			guestColumns,
			guestColumnsWithDefault,
			guestColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(guestType, guestMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(guestType, guestMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"guest\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"guest\" DEFAULT VALUES"
		}

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into guest")
	}

	if !cached {
		guestInsertCacheMut.Lock()
		guestInsertCache[key] = cache
		guestInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Guest record. See Update for
// whitelist behavior description.
func (o *Guest) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Guest record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Guest) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Guest, and panics on error.
// See Update for whitelist behavior description.
func (o *Guest) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Guest.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Guest) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	guestUpdateCacheMut.RLock()
	cache, cached := guestUpdateCache[key]
	guestUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(guestColumns, guestPrimaryKeyColumns, whitelist)
		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update guest, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"guest\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, guestPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(guestType, guestMapping, append(wl, guestPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update guest row")
	}

	if !cached {
		guestUpdateCacheMut.Lock()
		guestUpdateCache[key] = cache
		guestUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q guestQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q guestQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for guest")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GuestSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o GuestSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o GuestSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GuestSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"guest\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(guestPrimaryKeyColumns), len(colNames)+1, len(guestPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in guest slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Guest) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Guest) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Guest) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Guest) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no guest provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(guestColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	guestUpsertCacheMut.RLock()
	cache, cached := guestUpsertCache[key]
	guestUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			guestColumns,
			guestColumnsWithDefault,
			guestColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			guestColumns,
			guestPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert guest, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(guestPrimaryKeyColumns))
			copy(conflict, guestPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"guest\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(guestType, guestMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(guestType, guestMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert guest")
	}

	if !cached {
		guestUpsertCacheMut.Lock()
		guestUpsertCache[key] = cache
		guestUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Guest record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Guest) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Guest record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Guest) DeleteG() error {
	if o == nil {
		return errors.New("models: no Guest provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Guest record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Guest) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Guest record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Guest) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Guest provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), guestPrimaryKeyMapping)
	sql := "DELETE FROM \"guest\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from guest")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q guestQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q guestQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no guestQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from guest")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o GuestSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o GuestSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Guest slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o GuestSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GuestSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Guest slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(guestBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"guest\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, guestPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(guestPrimaryKeyColumns), 1, len(guestPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from guest slice")
	}

	if len(guestAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Guest) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Guest) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Guest) ReloadG() error {
	if o == nil {
		return errors.New("models: no Guest provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Guest) Reload(exec boil.Executor) error {
	ret, err := FindGuest(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *GuestSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *GuestSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GuestSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty GuestSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GuestSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	guests := GuestSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"guest\".* FROM \"guest\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, guestPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(guestPrimaryKeyColumns), 1, len(guestPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&guests)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GuestSlice")
	}

	*o = guests

	return nil
}

// GuestExists checks if the Guest row exists.
func GuestExists(exec boil.Executor, id int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"guest\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if guest exists")
	}

	return exists, nil
}

// GuestExistsG checks if the Guest row exists.
func GuestExistsG(id int) (bool, error) {
	return GuestExists(boil.GetDB(), id)
}

// GuestExistsGP checks if the Guest row exists. Panics on error.
func GuestExistsGP(id int) bool {
	e, err := GuestExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// GuestExistsP checks if the Guest row exists. Panics on error.
func GuestExistsP(exec boil.Executor, id int) bool {
	e, err := GuestExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}