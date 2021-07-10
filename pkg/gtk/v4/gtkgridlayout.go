// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_grid_layout_get_type()), F: marshalGridLayout},
		{T: externglib.Type(C.gtk_grid_layout_child_get_type()), F: marshalGridLayoutChild},
	})
}

// GridLayout: `GtkGridLayout` is a layout manager which arranges child widgets
// in rows and columns.
//
// Children have an "attach point" defined by the horizontal and vertical index
// of the cell they occupy; children can span multiple rows or columns. The
// layout properties for setting the attach points and spans are set using the
// [class@Gtk.GridLayoutChild] associated to each child widget.
//
// The behaviour of `GtkGridLayout` when several children occupy the same grid
// cell is undefined.
//
// `GtkGridLayout` can be used like a `GtkBoxLayout` if all children are
// attached to the same row or column; however, if you only ever need a single
// row or column, you should consider using `GtkBoxLayout`.
type GridLayout interface {
	gextras.Objector

	// BaselineRow retrieves the row set with
	// gtk_grid_layout_set_baseline_row().
	BaselineRow() int
	// ColumnHomogeneous checks whether all columns of @grid should have the
	// same width.
	ColumnHomogeneous() bool
	// ColumnSpacing retrieves the spacing set with
	// gtk_grid_layout_set_column_spacing().
	ColumnSpacing() uint
	// RowBaselinePosition returns the baseline position of @row.
	//
	// If no value has been set with
	// [method@Gtk.GridLayout.set_row_baseline_position], the default value of
	// GTK_BASELINE_POSITION_CENTER is returned.
	RowBaselinePosition(row int) BaselinePosition
	// RowHomogeneous checks whether all rows of @grid should have the same
	// height.
	RowHomogeneous() bool
	// RowSpacing retrieves the spacing set with
	// gtk_grid_layout_set_row_spacing().
	RowSpacing() uint
	// SetBaselineRow sets which row defines the global baseline for the entire
	// grid.
	//
	// Each row in the grid can have its own local baseline, but only one of
	// those is global, meaning it will be the baseline in the parent of the
	// @grid.
	SetBaselineRow(row int)
	// SetColumnHomogeneous sets whether all columns of @grid should have the
	// same width.
	SetColumnHomogeneous(homogeneous bool)
	// SetColumnSpacing sets the amount of space to insert between consecutive
	// columns.
	SetColumnSpacing(spacing uint)
	// SetRowHomogeneous sets whether all rows of @grid should have the same
	// height.
	SetRowHomogeneous(homogeneous bool)
	// SetRowSpacing sets the amount of space to insert between consecutive
	// rows.
	SetRowSpacing(spacing uint)
}

// GridLayoutClass implements the GridLayout interface.
type GridLayoutClass struct {
	LayoutManagerClass
}

var _ GridLayout = (*GridLayoutClass)(nil)

func wrapGridLayout(obj *externglib.Object) GridLayout {
	return &GridLayoutClass{
		LayoutManagerClass: LayoutManagerClass{
			Object: obj,
		},
	}
}

func marshalGridLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGridLayout(obj), nil
}

// NewGridLayout creates a new `GtkGridLayout`.
func NewGridLayout() *GridLayoutClass {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_grid_layout_new()

	var _gridLayout *GridLayoutClass // out

	_gridLayout = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*GridLayoutClass)

	return _gridLayout
}

// BaselineRow retrieves the row set with gtk_grid_layout_set_baseline_row().
func (grid *GridLayoutClass) BaselineRow() int {
	var _arg0 *C.GtkGridLayout // out
	var _cret C.int            // in

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_layout_get_baseline_row(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ColumnHomogeneous checks whether all columns of @grid should have the same
// width.
func (grid *GridLayoutClass) ColumnHomogeneous() bool {
	var _arg0 *C.GtkGridLayout // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_layout_get_column_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ColumnSpacing retrieves the spacing set with
// gtk_grid_layout_set_column_spacing().
func (grid *GridLayoutClass) ColumnSpacing() uint {
	var _arg0 *C.GtkGridLayout // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_layout_get_column_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RowBaselinePosition returns the baseline position of @row.
//
// If no value has been set with
// [method@Gtk.GridLayout.set_row_baseline_position], the default value of
// GTK_BASELINE_POSITION_CENTER is returned.
func (grid *GridLayoutClass) RowBaselinePosition(row int) BaselinePosition {
	var _arg0 *C.GtkGridLayout      // out
	var _arg1 C.int                 // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(row)

	_cret = C.gtk_grid_layout_get_row_baseline_position(_arg0, _arg1)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = (BaselinePosition)(_cret)

	return _baselinePosition
}

// RowHomogeneous checks whether all rows of @grid should have the same height.
func (grid *GridLayoutClass) RowHomogeneous() bool {
	var _arg0 *C.GtkGridLayout // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_layout_get_row_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpacing retrieves the spacing set with gtk_grid_layout_set_row_spacing().
func (grid *GridLayoutClass) RowSpacing() uint {
	var _arg0 *C.GtkGridLayout // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_layout_get_row_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetBaselineRow sets which row defines the global baseline for the entire
// grid.
//
// Each row in the grid can have its own local baseline, but only one of those
// is global, meaning it will be the baseline in the parent of the @grid.
func (grid *GridLayoutClass) SetBaselineRow(row int) {
	var _arg0 *C.GtkGridLayout // out
	var _arg1 C.int            // out

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(row)

	C.gtk_grid_layout_set_baseline_row(_arg0, _arg1)
}

// SetColumnHomogeneous sets whether all columns of @grid should have the same
// width.
func (grid *GridLayoutClass) SetColumnHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGridLayout // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_layout_set_column_homogeneous(_arg0, _arg1)
}

// SetColumnSpacing sets the amount of space to insert between consecutive
// columns.
func (grid *GridLayoutClass) SetColumnSpacing(spacing uint) {
	var _arg0 *C.GtkGridLayout // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_layout_set_column_spacing(_arg0, _arg1)
}

// SetRowHomogeneous sets whether all rows of @grid should have the same height.
func (grid *GridLayoutClass) SetRowHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGridLayout // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_layout_set_row_homogeneous(_arg0, _arg1)
}

// SetRowSpacing sets the amount of space to insert between consecutive rows.
func (grid *GridLayoutClass) SetRowSpacing(spacing uint) {
	var _arg0 *C.GtkGridLayout // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkGridLayout)(unsafe.Pointer(grid.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_layout_set_row_spacing(_arg0, _arg1)
}

// GridLayoutChild: `GtkLayoutChild` subclass for children in a `GtkGridLayout`.
type GridLayoutChild interface {
	gextras.Objector

	// Column retrieves the column number to which @child attaches its left
	// side.
	Column() int
	// ColumnSpan retrieves the number of columns that @child spans to.
	ColumnSpan() int
	// Row retrieves the row number to which @child attaches its top side.
	Row() int
	// RowSpan retrieves the number of rows that @child spans to.
	RowSpan() int
	// SetColumn sets the column number to attach the left side of @child.
	SetColumn(column int)
	// SetColumnSpan sets the number of columns @child spans to.
	SetColumnSpan(span int)
	// SetRow sets the row to place @child in.
	SetRow(row int)
	// SetRowSpan sets the number of rows @child spans to.
	SetRowSpan(span int)
}

// GridLayoutChildClass implements the GridLayoutChild interface.
type GridLayoutChildClass struct {
	LayoutChildClass
}

var _ GridLayoutChild = (*GridLayoutChildClass)(nil)

func wrapGridLayoutChild(obj *externglib.Object) GridLayoutChild {
	return &GridLayoutChildClass{
		LayoutChildClass: LayoutChildClass{
			Object: obj,
		},
	}
}

func marshalGridLayoutChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGridLayoutChild(obj), nil
}

// Column retrieves the column number to which @child attaches its left side.
func (child *GridLayoutChildClass) Column() int {
	var _arg0 *C.GtkGridLayoutChild // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_grid_layout_child_get_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ColumnSpan retrieves the number of columns that @child spans to.
func (child *GridLayoutChildClass) ColumnSpan() int {
	var _arg0 *C.GtkGridLayoutChild // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_grid_layout_child_get_column_span(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Row retrieves the row number to which @child attaches its top side.
func (child *GridLayoutChildClass) Row() int {
	var _arg0 *C.GtkGridLayoutChild // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_grid_layout_child_get_row(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RowSpan retrieves the number of rows that @child spans to.
func (child *GridLayoutChildClass) RowSpan() int {
	var _arg0 *C.GtkGridLayoutChild // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_grid_layout_child_get_row_span(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetColumn sets the column number to attach the left side of @child.
func (child *GridLayoutChildClass) SetColumn(column int) {
	var _arg0 *C.GtkGridLayoutChild // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(child.Native()))
	_arg1 = C.int(column)

	C.gtk_grid_layout_child_set_column(_arg0, _arg1)
}

// SetColumnSpan sets the number of columns @child spans to.
func (child *GridLayoutChildClass) SetColumnSpan(span int) {
	var _arg0 *C.GtkGridLayoutChild // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(child.Native()))
	_arg1 = C.int(span)

	C.gtk_grid_layout_child_set_column_span(_arg0, _arg1)
}

// SetRow sets the row to place @child in.
func (child *GridLayoutChildClass) SetRow(row int) {
	var _arg0 *C.GtkGridLayoutChild // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(child.Native()))
	_arg1 = C.int(row)

	C.gtk_grid_layout_child_set_row(_arg0, _arg1)
}

// SetRowSpan sets the number of rows @child spans to.
func (child *GridLayoutChildClass) SetRowSpan(span int) {
	var _arg0 *C.GtkGridLayoutChild // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkGridLayoutChild)(unsafe.Pointer(child.Native()))
	_arg1 = C.int(span)

	C.gtk_grid_layout_child_set_row_span(_arg0, _arg1)
}
