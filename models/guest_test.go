// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testGuests(t *testing.T) {
	t.Parallel()

	query := Guests(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testGuestsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = guest.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuestsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Guests(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuestsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GuestSlice{guest}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testGuestsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := GuestExists(tx, guest.ID)
	if err != nil {
		t.Errorf("Unable to check if Guest exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GuestExistsG to return true, but got false.")
	}
}
func testGuestsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	guestFound, err := FindGuest(tx, guest.ID)
	if err != nil {
		t.Error(err)
	}

	if guestFound == nil {
		t.Error("want a record, got nil")
	}
}
func testGuestsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Guests(tx).Bind(guest); err != nil {
		t.Error(err)
	}
}

func testGuestsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Guests(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGuestsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestOne := &Guest{}
	guestTwo := &Guest{}
	if err = randomize.Struct(seed, guestOne, guestDBTypes, false, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}
	if err = randomize.Struct(seed, guestTwo, guestDBTypes, false, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = guestTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Guests(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGuestsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	guestOne := &Guest{}
	guestTwo := &Guest{}
	if err = randomize.Struct(seed, guestOne, guestDBTypes, false, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}
	if err = randomize.Struct(seed, guestTwo, guestDBTypes, false, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = guestTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func guestBeforeInsertHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func guestAfterInsertHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func guestAfterSelectHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func guestBeforeUpdateHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func guestAfterUpdateHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func guestBeforeDeleteHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func guestAfterDeleteHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func guestBeforeUpsertHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func guestAfterUpsertHook(e boil.Executor, o *Guest) error {
	*o = Guest{}
	return nil
}

func testGuestsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Guest{}
	o := &Guest{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, guestDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Guest object: %s", err)
	}

	AddGuestHook(boil.BeforeInsertHook, guestBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	guestBeforeInsertHooks = []GuestHook{}

	AddGuestHook(boil.AfterInsertHook, guestAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	guestAfterInsertHooks = []GuestHook{}

	AddGuestHook(boil.AfterSelectHook, guestAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	guestAfterSelectHooks = []GuestHook{}

	AddGuestHook(boil.BeforeUpdateHook, guestBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	guestBeforeUpdateHooks = []GuestHook{}

	AddGuestHook(boil.AfterUpdateHook, guestAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	guestAfterUpdateHooks = []GuestHook{}

	AddGuestHook(boil.BeforeDeleteHook, guestBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	guestBeforeDeleteHooks = []GuestHook{}

	AddGuestHook(boil.AfterDeleteHook, guestAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	guestAfterDeleteHooks = []GuestHook{}

	AddGuestHook(boil.BeforeUpsertHook, guestBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	guestBeforeUpsertHooks = []GuestHook{}

	AddGuestHook(boil.AfterUpsertHook, guestAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	guestAfterUpsertHooks = []GuestHook{}
}
func testGuestsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuestsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx, guestColumns...); err != nil {
		t.Error(err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuestsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = guest.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testGuestsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GuestSlice{guest}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testGuestsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Guests(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	guestDBTypes = map[string]string{`Data`: `json`, `ID`: `integer`}
	_            = bytes.MinRead
)

func testGuestsUpdate(t *testing.T) {
	t.Parallel()

	if len(guestColumns) == len(guestPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	if err = guest.Update(tx); err != nil {
		t.Error(err)
	}
}

func testGuestsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(guestColumns) == len(guestPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	guest := &Guest{}
	if err = randomize.Struct(seed, guest, guestDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, guest, guestDBTypes, true, guestPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(guestColumns, guestPrimaryKeyColumns) {
		fields = guestColumns
	} else {
		fields = strmangle.SetComplement(
			guestColumns,
			guestPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(guest))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := GuestSlice{guest}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testGuestsUpsert(t *testing.T) {
	t.Parallel()

	if len(guestColumns) == len(guestPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	guest := Guest{}
	if err = randomize.Struct(seed, &guest, guestDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guest.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Guest: %s", err)
	}

	count, err := Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &guest, guestDBTypes, false, guestPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Guest struct: %s", err)
	}

	if err = guest.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Guest: %s", err)
	}

	count, err = Guests(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
