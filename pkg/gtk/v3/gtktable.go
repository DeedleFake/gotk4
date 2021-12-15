// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
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
type AttachOptions C.guint

const (
	// Expand: widget should expand to take up any extra space in its container
	// that has been allocated.
	Expand AttachOptions = 0b1
	// Shrink: widget should shrink as and when possible.
	Shrink AttachOptions = 0b10
	// Fill: widget should fill the space allocated to it.
	Fill AttachOptions = 0b100
)

func marshalAttachOptions(p uintptr) (any, error) {
	return AttachOptions(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AttachOptions.
func (a AttachOptions) String() string {
	if a == 0 {
		return "AttachOptions(0)"
	}

	var builder strings.Builder
	builder.Grow(18)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case Expand:
			builder.WriteString("Expand|")
		case Shrink:
			builder.WriteString("Shrink|")
		case Fill:
			builder.WriteString("Fill|")
		default:
			builder.WriteString(fmt.Sprintf("AttachOptions(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AttachOptions) Has(other AttachOptions) bool {
	return (a & other) == other
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

var (
	_ Containerer = (*Table)(nil)
)

func wrapTable(obj *externglib.Object) *Table {
	return &Table{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
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
	return wrapTable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTable: used to create a new table widget. An initial size must be given by
// specifying how many rows and columns the table should have, although this can
// be changed later with gtk_table_resize(). rows and columns must both be in
// the range 1 .. 65535. For historical reasons, 0 is accepted as well and is
// silently interpreted as 1.
//
// Deprecated: Use gtk_grid_new().
//
// The function takes the following parameters:
//
//    - rows: number of rows the new table should have.
//    - columns: number of columns the new table should have.
//    - homogeneous: if set to TRUE, all table cells are resized to the size of
//    the cell containing the largest widget.
//
func NewTable(rows, columns uint, homogeneous bool) *Table {
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
	runtime.KeepAlive(rows)
	runtime.KeepAlive(columns)
	runtime.KeepAlive(homogeneous)

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
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - leftAttach: column number to attach the left side of a child widget to.
//    - rightAttach: column number to attach the right side of a child widget
//    to.
//    - topAttach: row number to attach the top of a child widget to.
//    - bottomAttach: row number to attach the bottom of a child widget to.
//    - xoptions: used to specify the properties of the child widget when the
//    table is resized.
//    - yoptions: same as xoptions, except this field determines behaviour of
//    vertical resizing.
//    - xpadding: integer value specifying the padding on the left and right of
//    the widget being added to the table.
//    - ypadding: amount of padding above and below the child widget.
//
func (table *Table) Attach(child Widgetter, leftAttach, rightAttach, topAttach, bottomAttach uint, xoptions, yoptions AttachOptions, xpadding, ypadding uint) {
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
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)
	_arg6 = C.GtkAttachOptions(xoptions)
	_arg7 = C.GtkAttachOptions(yoptions)
	_arg8 = C.guint(xpadding)
	_arg9 = C.guint(ypadding)

	C.gtk_table_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
	runtime.KeepAlive(table)
	runtime.KeepAlive(child)
	runtime.KeepAlive(leftAttach)
	runtime.KeepAlive(rightAttach)
	runtime.KeepAlive(topAttach)
	runtime.KeepAlive(bottomAttach)
	runtime.KeepAlive(xoptions)
	runtime.KeepAlive(yoptions)
	runtime.KeepAlive(xpadding)
	runtime.KeepAlive(ypadding)
}

// AttachDefaults as there are many options associated with gtk_table_attach(),
// this convenience function provides the programmer with a means to add
// children to a table with identical padding and expansion options. The values
// used for the AttachOptions are GTK_EXPAND | GTK_FILL, and the padding is set
// to 0.
//
// Deprecated: Use gtk_grid_attach() with Grid. Note that the attach arguments
// differ between those two functions.
//
// The function takes the following parameters:
//
//    - widget: child widget to add.
//    - leftAttach: column number to attach the left side of the child widget
//    to.
//    - rightAttach: column number to attach the right side of the child widget
//    to.
//    - topAttach: row number to attach the top of the child widget to.
//    - bottomAttach: row number to attach the bottom of the child widget to.
//
func (table *Table) AttachDefaults(widget Widgetter, leftAttach, rightAttach, topAttach, bottomAttach uint) {
	var _arg0 *C.GtkTable  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.guint      // out
	var _arg3 C.guint      // out
	var _arg4 C.guint      // out
	var _arg5 C.guint      // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)

	C.gtk_table_attach_defaults(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(table)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(leftAttach)
	runtime.KeepAlive(rightAttach)
	runtime.KeepAlive(topAttach)
	runtime.KeepAlive(bottomAttach)
}

// ColSpacing gets the amount of space between column col, and column col + 1.
// See gtk_table_set_col_spacing().
//
// Deprecated: Grid does not offer a replacement for this functionality.
//
// The function takes the following parameters:
//
//    - column in the table, 0 indicates the first column.
//
func (table *Table) ColSpacing(column uint) uint {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(column)

	_cret = C.gtk_table_get_col_spacing(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(column)

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
	runtime.KeepAlive(table)

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
	runtime.KeepAlive(table)

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
	runtime.KeepAlive(table)

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
//
// The function takes the following parameters:
//
//    - row in the table, 0 indicates the first row.
//
func (table *Table) RowSpacing(row uint) uint {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(row)

	_cret = C.gtk_table_get_row_spacing(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(row)

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
	runtime.KeepAlive(table)

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
//
// The function takes the following parameters:
//
//    - rows: new number of rows.
//    - columns: new number of columns.
//
func (table *Table) Resize(rows, columns uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(rows)
	_arg2 = C.guint(columns)

	C.gtk_table_resize(_arg0, _arg1, _arg2)
	runtime.KeepAlive(table)
	runtime.KeepAlive(rows)
	runtime.KeepAlive(columns)
}

// SetColSpacing alters the amount of space between a given table column and the
// following column.
//
// Deprecated: Use gtk_widget_set_margin_start() and gtk_widget_set_margin_end()
// on the widgets contained in the row if you need this functionality. Grid does
// not support per-row spacing.
//
// The function takes the following parameters:
//
//    - column whose spacing should be changed.
//    - spacing: number of pixels that the spacing should take up.
//
func (table *Table) SetColSpacing(column, spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(column)
	_arg2 = C.guint(spacing)

	C.gtk_table_set_col_spacing(_arg0, _arg1, _arg2)
	runtime.KeepAlive(table)
	runtime.KeepAlive(column)
	runtime.KeepAlive(spacing)
}

// SetColSpacings sets the space between every column in table equal to spacing.
//
// Deprecated: Use gtk_grid_set_column_spacing() with Grid.
//
// The function takes the following parameters:
//
//    - spacing: number of pixels of space to place between every column in the
//    table.
//
func (table *Table) SetColSpacings(spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_table_set_col_spacings(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(spacing)
}

// SetHomogeneous changes the homogenous property of table cells, ie. whether
// all cells are an equal size or not.
//
// Deprecated: Use gtk_grid_set_row_homogeneous() and
// gtk_grid_set_column_homogeneous() with Grid.
//
// The function takes the following parameters:
//
//    - homogeneous: set to TRUE to ensure all table cells are the same size.
//    Set to FALSE if this is not your desired behaviour.
//
func (table *Table) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_table_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(homogeneous)
}

// SetRowSpacing changes the space between a given table row and the subsequent
// row.
//
// Deprecated: Use gtk_widget_set_margin_top() and
// gtk_widget_set_margin_bottom() on the widgets contained in the row if you
// need this functionality. Grid does not support per-row spacing.
//
// The function takes the following parameters:
//
//    - row number whose spacing will be changed.
//    - spacing: number of pixels that the spacing should take up.
//
func (table *Table) SetRowSpacing(row, spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(row)
	_arg2 = C.guint(spacing)

	C.gtk_table_set_row_spacing(_arg0, _arg1, _arg2)
	runtime.KeepAlive(table)
	runtime.KeepAlive(row)
	runtime.KeepAlive(spacing)
}

// SetRowSpacings sets the space between every row in table equal to spacing.
//
// Deprecated: Use gtk_grid_set_row_spacing() with Grid.
//
// The function takes the following parameters:
//
//    - spacing: number of pixels of space to place between every row in the
//    table.
//
func (table *Table) SetRowSpacings(spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(table.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_table_set_row_spacings(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(spacing)
}

// TableChild: instance of this type is always passed by reference.
type TableChild struct {
	*tableChild
}

// tableChild is the struct that's finalized.
type tableChild struct {
	native *C.GtkTableChild
}

func (t *TableChild) Widget() Widgetter {
	var v Widgetter // out
	{
		objptr := unsafe.Pointer(t.native.widget)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		v = rv
	}
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

// TableRowCol: instance of this type is always passed by reference.
type TableRowCol struct {
	*tableRowCol
}

// tableRowCol is the struct that's finalized.
type tableRowCol struct {
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
