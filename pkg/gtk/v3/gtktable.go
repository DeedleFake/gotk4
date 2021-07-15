// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_attach_options_get_type()), F: marshalAttachOptions},
		{T: externglib.Type(C.gtk_table_get_type()), F: marshalTabler},
	})
}

// AttachOptions denotes the expansion properties that a widget will have when
// it (or its parent) is resized.
type AttachOptions int

const (
	// AttachOptionsExpand: widget should expand to take up any extra space in
	// its container that has been allocated.
	AttachOptionsExpand AttachOptions = 0b1
	// AttachOptionsShrink: widget should shrink as and when possible.
	AttachOptionsShrink AttachOptions = 0b10
	// AttachOptionsFill: widget should fill the space allocated to it.
	AttachOptionsFill AttachOptions = 0b100
)

func marshalAttachOptions(p uintptr) (interface{}, error) {
	return AttachOptions(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Table functions allow the programmer to arrange widgets in rows and columns,
// making it easy to align many widgets next to each other, horizontally and
// vertically.
//
// Tables are created with a call to gtk_table_new(), the size of which can
// later be changed with gtk_table_resize().
//
// Widgets can be added to a table using gtk_table_attach() or the more
// convenient (but slightly less flexible) gtk_table_attach_defaults().
//
// To alter the space next to a specific row, use gtk_table_set_row_spacing(),
// and for a column, gtk_table_set_col_spacing(). The gaps between all rows or
// columns can be changed by calling gtk_table_set_row_spacings() or
// gtk_table_set_col_spacings() respectively. Note that spacing is added between
// the children, while padding added by gtk_table_attach() is added on either
// side of the widget it belongs to.
//
// gtk_table_set_homogeneous(), can be used to set whether all cells in the
// table will resize themselves to the size of the largest widget in the table.
//
// > Table has been deprecated. Use Grid instead. It provides the same >
// capabilities as GtkTable for arranging widgets in a rectangular grid, but >
// does support height-for-width geometry management.
type Table struct {
	Container
}

var _ gextras.Nativer = (*Table)(nil)

func wrapTable(obj *externglib.Object) *Table {
	return &Table{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalTabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTable(obj), nil
}

// NewTable: used to create a new table widget. An initial size must be given by
// specifying how many rows and columns the table should have, although this can
// be changed later with gtk_table_resize(). rows and columns must both be in
// the range 1 .. 65535. For historical reasons, 0 is accepted as well and is
// silently interpreted as 1.
//
// Deprecated: Use gtk_grid_new().
func NewTable(rows uint, columns uint, homogeneous bool) *Table {
	var _arg1 C.guint      // out
	var _arg2 C.guint      // out
	var _arg3 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = C.guint(rows)
	_arg2 = C.guint(columns)
	if homogeneous {
		_arg3 = C.TRUE
	}

	_cret = C.gtk_table_new(_arg1, _arg2, _arg3)

	var _table *Table // out

	_table = wrapTable(externglib.Take(unsafe.Pointer(_cret)))

	return _table
}

// Attach adds a widget to a table. The number of “cells” that a widget will
// occupy is specified by left_attach, right_attach, top_attach and
// bottom_attach. These each represent the leftmost, rightmost, uppermost and
// lowest column and row numbers of the table. (Columns and rows are indexed
// from zero).
//
// To make a button occupy the lower right cell of a 2x2 table, use
//
//    gtk_table_attach (table, button,
//                      1, 2, // left, right attach
//                      1, 2, // top, bottom attach
//                      xoptions, yoptions,
//                      xpadding, ypadding);
//
// If you want to make the button span the entire bottom row, use left_attach ==
// 0 and right_attach = 2 instead.
//
// Deprecated: Use gtk_grid_attach() with Grid. Note that the attach arguments
// differ between those two functions.
func (table *Table) Attach(child Widgetter, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint) {
	var _arg0 *C.GtkTable        // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.guint            // out
	var _arg3 C.guint            // out
	var _arg4 C.guint            // out
	var _arg5 C.guint            // out
	var _arg6 C.GtkAttachOptions // out
	var _arg7 C.GtkAttachOptions // out
	var _arg8 C.guint            // out
	var _arg9 C.guint            // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)
	_arg6 = C.GtkAttachOptions(xoptions)
	_arg7 = C.GtkAttachOptions(yoptions)
	_arg8 = C.guint(xpadding)
	_arg9 = C.guint(ypadding)

	C.gtk_table_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// AttachDefaults as there are many options associated with gtk_table_attach(),
// this convenience function provides the programmer with a means to add
// children to a table with identical padding and expansion options. The values
// used for the AttachOptions are GTK_EXPAND | GTK_FILL, and the padding is set
// to 0.
//
// Deprecated: Use gtk_grid_attach() with Grid. Note that the attach arguments
// differ between those two functions.
func (table *Table) AttachDefaults(widget Widgetter, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint) {
	var _arg0 *C.GtkTable  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.guint      // out
	var _arg3 C.guint      // out
	var _arg4 C.guint      // out
	var _arg5 C.guint      // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)

	C.gtk_table_attach_defaults(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// ColSpacing gets the amount of space between column col, and column col + 1.
// See gtk_table_set_col_spacing().
//
// Deprecated: Grid does not offer a replacement for this functionality.
func (table *Table) ColSpacing(column uint) uint {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(column)

	_cret = C.gtk_table_get_col_spacing(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// DefaultColSpacing gets the default column spacing for the table. This is the
// spacing that will be used for newly added columns. (See
// gtk_table_set_col_spacings())
//
// Deprecated: Use gtk_grid_get_column_spacing() with Grid.
func (table *Table) DefaultColSpacing() uint {
	var _arg0 *C.GtkTable // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))

	_cret = C.gtk_table_get_default_col_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// DefaultRowSpacing gets the default row spacing for the table. This is the
// spacing that will be used for newly added rows. (See
// gtk_table_set_row_spacings())
//
// Deprecated: Use gtk_grid_get_row_spacing() with Grid.
func (table *Table) DefaultRowSpacing() uint {
	var _arg0 *C.GtkTable // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))

	_cret = C.gtk_table_get_default_row_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Homogeneous returns whether the table cells are all constrained to the same
// width and height. (See gtk_table_set_homogeneous ())
//
// Deprecated: Use gtk_grid_get_row_homogeneous() and
// gtk_grid_get_column_homogeneous() with Grid.
func (table *Table) Homogeneous() bool {
	var _arg0 *C.GtkTable // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))

	_cret = C.gtk_table_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpacing gets the amount of space between row row, and row row + 1. See
// gtk_table_set_row_spacing().
//
// Deprecated: Grid does not offer a replacement for this functionality.
func (table *Table) RowSpacing(row uint) uint {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(row)

	_cret = C.gtk_table_get_row_spacing(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Size gets the number of rows and columns in the table.
//
// Deprecated: Grid does not expose the number of columns and rows.
func (table *Table) Size() (rows uint, columns uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // in
	var _arg2 C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))

	C.gtk_table_get_size(_arg0, &_arg1, &_arg2)

	var _rows uint    // out
	var _columns uint // out

	_rows = uint(_arg1)
	_columns = uint(_arg2)

	return _rows, _columns
}

// Resize: if you need to change a table’s size after it has been created, this
// function allows you to do so.
//
// Deprecated: Grid resizes automatically.
func (table *Table) Resize(rows uint, columns uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(rows)
	_arg2 = C.guint(columns)

	C.gtk_table_resize(_arg0, _arg1, _arg2)
}

// SetColSpacing alters the amount of space between a given table column and the
// following column.
//
// Deprecated: Use gtk_widget_set_margin_start() and gtk_widget_set_margin_end()
// on the widgets contained in the row if you need this functionality. Grid does
// not support per-row spacing.
func (table *Table) SetColSpacing(column uint, spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(column)
	_arg2 = C.guint(spacing)

	C.gtk_table_set_col_spacing(_arg0, _arg1, _arg2)
}

// SetColSpacings sets the space between every column in table equal to spacing.
//
// Deprecated: Use gtk_grid_set_column_spacing() with Grid.
func (table *Table) SetColSpacings(spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_table_set_col_spacings(_arg0, _arg1)
}

// SetHomogeneous changes the homogenous property of table cells, ie. whether
// all cells are an equal size or not.
//
// Deprecated: Use gtk_grid_set_row_homogeneous() and
// gtk_grid_set_column_homogeneous() with Grid.
func (table *Table) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_table_set_homogeneous(_arg0, _arg1)
}

// SetRowSpacing changes the space between a given table row and the subsequent
// row.
//
// Deprecated: Use gtk_widget_set_margin_top() and
// gtk_widget_set_margin_bottom() on the widgets contained in the row if you
// need this functionality. Grid does not support per-row spacing.
func (table *Table) SetRowSpacing(row uint, spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(row)
	_arg2 = C.guint(spacing)

	C.gtk_table_set_row_spacing(_arg0, _arg1, _arg2)
}

// SetRowSpacings sets the space between every row in table equal to spacing.
//
// Deprecated: Use gtk_grid_set_row_spacing() with Grid.
func (table *Table) SetRowSpacings(spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_table_set_row_spacings(_arg0, _arg1)
}

type TableChild struct {
	nocopy gextras.NoCopy
	native *C.GtkTableChild
}

func (t *TableChild) Widget() *Widget {
	var v *Widget // out
	v = wrapWidget(externglib.Take(unsafe.Pointer(t.native.widget)))
	return v
}

func (t *TableChild) LeftAttach() uint16 {
	var v uint16 // out
	v = uint16(t.native.left_attach)
	return v
}

func (t *TableChild) RightAttach() uint16 {
	var v uint16 // out
	v = uint16(t.native.right_attach)
	return v
}

func (t *TableChild) TopAttach() uint16 {
	var v uint16 // out
	v = uint16(t.native.top_attach)
	return v
}

func (t *TableChild) BottomAttach() uint16 {
	var v uint16 // out
	v = uint16(t.native.bottom_attach)
	return v
}

func (t *TableChild) Xpadding() uint16 {
	var v uint16 // out
	v = uint16(t.native.xpadding)
	return v
}

func (t *TableChild) Ypadding() uint16 {
	var v uint16 // out
	v = uint16(t.native.ypadding)
	return v
}

type TableRowCol struct {
	nocopy gextras.NoCopy
	native *C.GtkTableRowCol
}

func (t *TableRowCol) Requisition() uint16 {
	var v uint16 // out
	v = uint16(t.native.requisition)
	return v
}

func (t *TableRowCol) Allocation() uint16 {
	var v uint16 // out
	v = uint16(t.native.allocation)
	return v
}

func (t *TableRowCol) Spacing() uint16 {
	var v uint16 // out
	v = uint16(t.native.spacing)
	return v
}
