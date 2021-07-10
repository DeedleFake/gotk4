// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_grid_get_type()), F: marshalGrid},
	})
}

// Grid is a container which arranges its child widgets in rows and columns,
// with arbitrary positions and horizontal/vertical spans.
//
// Children are added using gtk_grid_attach(). They can span multiple rows or
// columns. It is also possible to add a child next to an existing child, using
// gtk_grid_attach_next_to(). The behaviour of GtkGrid when several children
// occupy the same grid cell is undefined.
//
// GtkGrid can be used like a Box by just using gtk_container_add(), which will
// place children next to each other in the direction determined by the
// Orientable:orientation property. However, if all you want is a single row or
// column, then Box is the preferred widget.
//
//
// CSS nodes
//
// GtkGrid uses a single CSS node with name grid.
type Grid interface {
	gextras.Objector

	// Attach adds a widget to the grid.
	//
	// The position of @child is determined by @left and @top. The number of
	// “cells” that @child will occupy is determined by @width and @height.
	Attach(child Widget, left int, top int, width int, height int)
	// BaselineRow returns which row defines the global baseline of @grid.
	BaselineRow() int
	// ChildAt gets the child of @grid whose area covers the grid cell whose
	// upper left corner is at @left, @top.
	ChildAt(left int, top int) *WidgetClass
	// ColumnHomogeneous returns whether all columns of @grid have the same
	// width.
	ColumnHomogeneous() bool
	// ColumnSpacing returns the amount of space between the columns of @grid.
	ColumnSpacing() uint
	// RowBaselinePosition returns the baseline position of @row as set by
	// gtk_grid_set_row_baseline_position() or the default value
	// GTK_BASELINE_POSITION_CENTER.
	RowBaselinePosition(row int) BaselinePosition
	// RowHomogeneous returns whether all rows of @grid have the same height.
	RowHomogeneous() bool
	// RowSpacing returns the amount of space between the rows of @grid.
	RowSpacing() uint
	// InsertColumn inserts a column at the specified position.
	//
	// Children which are attached at or to the right of this position are moved
	// one column to the right. Children which span across this position are
	// grown to span the new column.
	InsertColumn(position int)
	// InsertRow inserts a row at the specified position.
	//
	// Children which are attached at or below this position are moved one row
	// down. Children which span across this position are grown to span the new
	// row.
	InsertRow(position int)
	// RemoveColumn removes a column from the grid.
	//
	// Children that are placed in this column are removed, spanning children
	// that overlap this column have their width reduced by one, and children
	// after the column are moved to the left.
	RemoveColumn(position int)
	// RemoveRow removes a row from the grid.
	//
	// Children that are placed in this row are removed, spanning children that
	// overlap this row have their height reduced by one, and children below the
	// row are moved up.
	RemoveRow(position int)
	// SetBaselineRow sets which row defines the global baseline for the entire
	// grid. Each row in the grid can have its own local baseline, but only one
	// of those is global, meaning it will be the baseline in the parent of the
	// @grid.
	SetBaselineRow(row int)
	// SetColumnHomogeneous sets whether all columns of @grid will have the same
	// width.
	SetColumnHomogeneous(homogeneous bool)
	// SetColumnSpacing sets the amount of space between columns of @grid.
	SetColumnSpacing(spacing uint)
	// SetRowHomogeneous sets whether all rows of @grid will have the same
	// height.
	SetRowHomogeneous(homogeneous bool)
	// SetRowSpacing sets the amount of space between rows of @grid.
	SetRowSpacing(spacing uint)
}

// GridClass implements the Grid interface.
type GridClass struct {
	*externglib.Object
	ContainerClass
	BuildableIface
	OrientableIface
}

var _ Grid = (*GridClass)(nil)

func wrapGrid(obj *externglib.Object) Grid {
	return &GridClass{
		Object: obj,
		ContainerClass: ContainerClass{
			Object: obj,
			WidgetClass: WidgetClass{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
		OrientableIface: OrientableIface{
			Object: obj,
		},
	}
}

func marshalGrid(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGrid(obj), nil
}

// NewGrid creates a new grid widget.
func NewGrid() *GridClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_grid_new()

	var _grid *GridClass // out

	_grid = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*GridClass)

	return _grid
}

// Attach adds a widget to the grid.
//
// The position of @child is determined by @left and @top. The number of “cells”
// that @child will occupy is determined by @width and @height.
func (grid *GridClass) Attach(child Widget, left int, top int, width int, height int) {
	var _arg0 *C.GtkGrid   // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out
	var _arg4 C.gint       // out
	var _arg5 C.gint       // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(left)
	_arg3 = C.gint(top)
	_arg4 = C.gint(width)
	_arg5 = C.gint(height)

	C.gtk_grid_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// BaselineRow returns which row defines the global baseline of @grid.
func (grid *GridClass) BaselineRow() int {
	var _arg0 *C.GtkGrid // out
	var _cret C.gint     // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_baseline_row(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ChildAt gets the child of @grid whose area covers the grid cell whose upper
// left corner is at @left, @top.
func (grid *GridClass) ChildAt(left int, top int) *WidgetClass {
	var _arg0 *C.GtkGrid   // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(left)
	_arg2 = C.gint(top)

	_cret = C.gtk_grid_get_child_at(_arg0, _arg1, _arg2)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// ColumnHomogeneous returns whether all columns of @grid have the same width.
func (grid *GridClass) ColumnHomogeneous() bool {
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
func (grid *GridClass) ColumnSpacing() uint {
	var _arg0 *C.GtkGrid // out
	var _cret C.guint    // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_column_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RowBaselinePosition returns the baseline position of @row as set by
// gtk_grid_set_row_baseline_position() or the default value
// GTK_BASELINE_POSITION_CENTER.
func (grid *GridClass) RowBaselinePosition(row int) BaselinePosition {
	var _arg0 *C.GtkGrid            // out
	var _arg1 C.gint                // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(row)

	_cret = C.gtk_grid_get_row_baseline_position(_arg0, _arg1)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = (BaselinePosition)(_cret)

	return _baselinePosition
}

// RowHomogeneous returns whether all rows of @grid have the same height.
func (grid *GridClass) RowHomogeneous() bool {
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
func (grid *GridClass) RowSpacing() uint {
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
func (grid *GridClass) InsertColumn(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(position)

	C.gtk_grid_insert_column(_arg0, _arg1)
}

// InsertRow inserts a row at the specified position.
//
// Children which are attached at or below this position are moved one row down.
// Children which span across this position are grown to span the new row.
func (grid *GridClass) InsertRow(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(position)

	C.gtk_grid_insert_row(_arg0, _arg1)
}

// RemoveColumn removes a column from the grid.
//
// Children that are placed in this column are removed, spanning children that
// overlap this column have their width reduced by one, and children after the
// column are moved to the left.
func (grid *GridClass) RemoveColumn(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(position)

	C.gtk_grid_remove_column(_arg0, _arg1)
}

// RemoveRow removes a row from the grid.
//
// Children that are placed in this row are removed, spanning children that
// overlap this row have their height reduced by one, and children below the row
// are moved up.
func (grid *GridClass) RemoveRow(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(position)

	C.gtk_grid_remove_row(_arg0, _arg1)
}

// SetBaselineRow sets which row defines the global baseline for the entire
// grid. Each row in the grid can have its own local baseline, but only one of
// those is global, meaning it will be the baseline in the parent of the @grid.
func (grid *GridClass) SetBaselineRow(row int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(row)

	C.gtk_grid_set_baseline_row(_arg0, _arg1)
}

// SetColumnHomogeneous sets whether all columns of @grid will have the same
// width.
func (grid *GridClass) SetColumnHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_set_column_homogeneous(_arg0, _arg1)
}

// SetColumnSpacing sets the amount of space between columns of @grid.
func (grid *GridClass) SetColumnSpacing(spacing uint) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_set_column_spacing(_arg0, _arg1)
}

// SetRowHomogeneous sets whether all rows of @grid will have the same height.
func (grid *GridClass) SetRowHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_set_row_homogeneous(_arg0, _arg1)
}

// SetRowSpacing sets the amount of space between rows of @grid.
func (grid *GridClass) SetRowSpacing(spacing uint) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_set_row_spacing(_arg0, _arg1)
}
