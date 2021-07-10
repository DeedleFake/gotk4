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
		{T: externglib.Type(C.gtk_grid_get_type()), F: marshalGridder},
	})
}

// Gridder describes Grid's methods.
type Gridder interface {
	gextras.Objector

	Attach(child Widgetter, column int, row int, width int, height int)
	BaselineRow() int
	ChildAt(column int, row int) *Widget
	ColumnHomogeneous() bool
	ColumnSpacing() uint
	RowBaselinePosition(row int) BaselinePosition
	RowHomogeneous() bool
	RowSpacing() uint
	InsertColumn(position int)
	InsertRow(position int)
	QueryChild(child Widgetter) (column int, row int, width int, height int)
	Remove(child Widgetter)
	RemoveColumn(position int)
	RemoveRow(position int)
	SetBaselineRow(row int)
	SetColumnHomogeneous(homogeneous bool)
	SetColumnSpacing(spacing uint)
	SetRowHomogeneous(homogeneous bool)
	SetRowSpacing(spacing uint)
}

// Grid: `GtkGrid` is a container which arranges its child widgets in rows and
// columns.
//
// !An example GtkGrid (grid.png)
//
// It supports arbitrary positions and horizontal/vertical spans.
//
// Children are added using [method@Gtk.Grid.attach]. They can span multiple
// rows or columns. It is also possible to add a child next to an existing
// child, using [method@Gtk.Grid.attach_next_to]. To remove a child from the
// grid, use [method@Gtk.Grid.remove].
//
// The behaviour of `GtkGrid` when several children occupy the same grid cell is
// undefined.
//
//
// GtkGrid as GtkBuildable
//
// Every child in a `GtkGrid` has access to a custom [iface@Gtk.Buildable]
// element, called ´<layout>´. It can by used to specify a position in the grid
// and optionally spans. All properties that can be used in the ´<layout>´
// element are implemented by [class@Gtk.GridLayoutChild].
//
// It is implemented by `GtkWidget` using [class@Gtk.LayoutManager].
//
// To showcase it, here is a simple example:
//
// “`xml <object class="GtkGrid" id="my_grid"> <child> <object class="GtkButton"
// id="button1"> <property name="label">Button 1</property> <layout> <property
// name="column">0</property> <property name="row">0</property> </layout>
// </object> </child> <child> <object class="GtkButton" id="button2"> <property
// name="label">Button 2</property> <layout> <property
// name="column">1</property> <property name="row">0</property> </layout>
// </object> </child> <child> <object class="GtkButton" id="button3"> <property
// name="label">Button 3</property> <layout> <property
// name="column">2</property> <property name="row">0</property> <property
// name="row-span">2</property> </layout> </object> </child> <child> <object
// class="GtkButton" id="button4"> <property name="label">Button 4</property>
// <layout> <property name="column">0</property> <property
// name="row">1</property> <property name="column-span">2</property> </layout>
// </object> </child> </object> “`
//
// It organizes the first two buttons side-by-side in one cell each. The third
// button is in the last column but spans across two rows. This is defined by
// the ´row-span´ property. The last button is located in the second row and
// spans across two columns, which is defined by the ´column-span´ property.
//
//
// CSS nodes
//
// `GtkGrid` uses a single CSS node with name `grid`.
//
//
// Accessibility
//
// `GtkGrid` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Grid struct {
	*externglib.Object

	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ Gridder = (*Grid)(nil)

func wrapGridder(obj *externglib.Object) Gridder {
	return &Grid{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalGridder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGridder(obj), nil
}

// NewGrid creates a new grid widget.
func NewGrid() *Grid {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_grid_new()

	var _grid *Grid // out

	_grid = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Grid)

	return _grid
}

// Attach adds a widget to the grid.
//
// The position of @child is determined by @column and @row. The number of
// “cells” that @child will occupy is determined by @width and @height.
func (grid *Grid) Attach(child Widgetter, column int, row int, width int, height int) {
	var _arg0 *C.GtkGrid   // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.int        // out
	var _arg3 C.int        // out
	var _arg4 C.int        // out
	var _arg5 C.int        // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(column)
	_arg3 = C.int(row)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	C.gtk_grid_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// BaselineRow returns which row defines the global baseline of @grid.
func (grid *Grid) BaselineRow() int {
	var _arg0 *C.GtkGrid // out
	var _cret C.int      // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_baseline_row(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ChildAt gets the child of @grid whose area covers the grid cell at @column,
// @row.
func (grid *Grid) ChildAt(column int, row int) *Widget {
	var _arg0 *C.GtkGrid   // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(column)
	_arg2 = C.int(row)

	_cret = C.gtk_grid_get_child_at(_arg0, _arg1, _arg2)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// ColumnHomogeneous returns whether all columns of @grid have the same width.
func (grid *Grid) ColumnHomogeneous() bool {
	var _arg0 *C.GtkGrid // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_column_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ColumnSpacing returns the amount of space between the columns of @grid.
func (grid *Grid) ColumnSpacing() uint {
	var _arg0 *C.GtkGrid // out
	var _cret C.guint    // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_column_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RowBaselinePosition returns the baseline position of @row.
//
// See [method@Gtk.Grid.set_row_baseline_position].
func (grid *Grid) RowBaselinePosition(row int) BaselinePosition {
	var _arg0 *C.GtkGrid            // out
	var _arg1 C.int                 // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(row)

	_cret = C.gtk_grid_get_row_baseline_position(_arg0, _arg1)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = (BaselinePosition)(_cret)

	return _baselinePosition
}

// RowHomogeneous returns whether all rows of @grid have the same height.
func (grid *Grid) RowHomogeneous() bool {
	var _arg0 *C.GtkGrid // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_row_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpacing returns the amount of space between the rows of @grid.
func (grid *Grid) RowSpacing() uint {
	var _arg0 *C.GtkGrid // out
	var _cret C.guint    // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_row_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// InsertColumn inserts a column at the specified position.
//
// Children which are attached at or to the right of this position are moved one
// column to the right. Children which span across this position are grown to
// span the new column.
func (grid *Grid) InsertColumn(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(position)

	C.gtk_grid_insert_column(_arg0, _arg1)
}

// InsertRow inserts a row at the specified position.
//
// Children which are attached at or below this position are moved one row down.
// Children which span across this position are grown to span the new row.
func (grid *Grid) InsertRow(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(position)

	C.gtk_grid_insert_row(_arg0, _arg1)
}

// QueryChild queries the attach points and spans of @child inside the given
// `GtkGrid`.
func (grid *Grid) QueryChild(child Widgetter) (column int, row int, width int, height int) {
	var _arg0 *C.GtkGrid   // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.int        // in
	var _arg3 C.int        // in
	var _arg4 C.int        // in
	var _arg5 C.int        // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_grid_query_child(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _column int // out
	var _row int    // out
	var _width int  // out
	var _height int // out

	_column = int(_arg2)
	_row = int(_arg3)
	_width = int(_arg4)
	_height = int(_arg5)

	return _column, _row, _width, _height
}

// Remove removes a child from @grid.
//
// The child must have been added with [method@Gtk.Grid.attach] or
// [method@Gtk.Grid.attach_next_to].
func (grid *Grid) Remove(child Widgetter) {
	var _arg0 *C.GtkGrid   // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_grid_remove(_arg0, _arg1)
}

// RemoveColumn removes a column from the grid.
//
// Children that are placed in this column are removed, spanning children that
// overlap this column have their width reduced by one, and children after the
// column are moved to the left.
func (grid *Grid) RemoveColumn(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(position)

	C.gtk_grid_remove_column(_arg0, _arg1)
}

// RemoveRow removes a row from the grid.
//
// Children that are placed in this row are removed, spanning children that
// overlap this row have their height reduced by one, and children below the row
// are moved up.
func (grid *Grid) RemoveRow(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(position)

	C.gtk_grid_remove_row(_arg0, _arg1)
}

// SetBaselineRow sets which row defines the global baseline for the entire
// grid.
//
// Each row in the grid can have its own local baseline, but only one of those
// is global, meaning it will be the baseline in the parent of the @grid.
func (grid *Grid) SetBaselineRow(row int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.int(row)

	C.gtk_grid_set_baseline_row(_arg0, _arg1)
}

// SetColumnHomogeneous sets whether all columns of @grid will have the same
// width.
func (grid *Grid) SetColumnHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_set_column_homogeneous(_arg0, _arg1)
}

// SetColumnSpacing sets the amount of space between columns of @grid.
func (grid *Grid) SetColumnSpacing(spacing uint) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_set_column_spacing(_arg0, _arg1)
}

// SetRowHomogeneous sets whether all rows of @grid will have the same height.
func (grid *Grid) SetRowHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_set_row_homogeneous(_arg0, _arg1)
}

// SetRowSpacing sets the amount of space between rows of @grid.
func (grid *Grid) SetRowSpacing(spacing uint) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_set_row_spacing(_arg0, _arg1)
}
