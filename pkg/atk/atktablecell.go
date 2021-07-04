// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_table_cell_get_type()), F: marshalTableCell},
	})
}

// TableCell: being Table a component which present elements ordered via rows
// and columns, an TableCell is the interface which each of those elements, so
// "cells" should implement.
//
// See also Table.
type TableCell interface {
	Object

	// ColumnSpan returns a reference to the accessible of the containing table.
	ColumnSpan() int
	// Position returns a reference to the accessible of the containing table.
	Position() (row int, column int, ok bool)
	// RowColumnSpan returns a reference to the accessible of the containing
	// table.
	RowColumnSpan() (row int, column int, rowSpan int, columnSpan int, ok bool)
	// RowSpan returns a reference to the accessible of the containing table.
	RowSpan() int
	// Table returns a reference to the accessible of the containing table.
	Table() Object
}

// tableCell implements the TableCell interface.
type tableCell struct {
	Object
}

var _ TableCell = (*tableCell)(nil)

// WrapTableCell wraps a GObject to a type that implements
// interface TableCell. It is primarily used internally.
func WrapTableCell(obj *externglib.Object) TableCell {
	return tableCell{
		Object: WrapObject(obj),
	}
}

func marshalTableCell(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTableCell(obj), nil
}

func (c tableCell) ColumnSpan() int {
	var _arg0 *C.AtkTableCell // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(c.Native()))

	_cret = C.atk_table_cell_get_column_span(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c tableCell) Position() (row int, column int, ok bool) {
	var _arg0 *C.AtkTableCell // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(c.Native()))

	_cret = C.atk_table_cell_get_position(_arg0, &_arg1, &_arg2)

	var _row int    // out
	var _column int // out
	var _ok bool    // out

	_row = int(_arg1)
	_column = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _row, _column, _ok
}

func (c tableCell) RowColumnSpan() (row int, column int, rowSpan int, columnSpan int, ok bool) {
	var _arg0 *C.AtkTableCell // out
	var _arg1 C.gint          // in
	var _arg2 C.gint          // in
	var _arg3 C.gint          // in
	var _arg4 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(c.Native()))

	_cret = C.atk_table_cell_get_row_column_span(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)

	var _row int        // out
	var _column int     // out
	var _rowSpan int    // out
	var _columnSpan int // out
	var _ok bool        // out

	_row = int(_arg1)
	_column = int(_arg2)
	_rowSpan = int(_arg3)
	_columnSpan = int(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _row, _column, _rowSpan, _columnSpan, _ok
}

func (c tableCell) RowSpan() int {
	var _arg0 *C.AtkTableCell // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(c.Native()))

	_cret = C.atk_table_cell_get_row_span(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c tableCell) Table() Object {
	var _arg0 *C.AtkTableCell // out
	var _cret *C.AtkObject    // in

	_arg0 = (*C.AtkTableCell)(unsafe.Pointer(c.Native()))

	_cret = C.atk_table_cell_get_table(_arg0)

	var _object Object // out

	_object = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Object)

	return _object
}
