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
		{T: externglib.Type(C.gtk_column_view_get_type()), F: marshalColumnView},
	})
}

// ColumnView: `GtkColumnView` presents a large dynamic list of items using
// multiple columns with headers.
//
// `GtkColumnView` uses the factories of its columns to generate a cell widget
// for each column, for each visible item and displays them together as the row
// for this item.
//
// The [property@Gtk.ColumnView:show-row-separators] and
// [propertyGtk.ColumnView:show-column-separators] properties offer a simple way
// to display separators between the rows or columns.
//
// `GtkColumnView` allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on *rubberband selection*, using
// [property@Gtk.ColumnView:enable-rubberband].
//
// The column view supports sorting that can be customized by the user by
// clicking on column headers. To set this up, the `GtkSorter` returned by
// [method@Gtk.ColumnView.get_sorter] must be attached to a sort model for the
// data that the view is showing, and the columns must have sorters attached to
// them by calling [method@Gtk.ColumnViewColumn.set_sorter]. The initial sort
// order can be set with [method@Gtk.ColumnView.sort_by_column].
//
// The column view also supports interactive resizing and reordering of columns,
// via Drag-and-Drop of the column headers. This can be enabled or disabled with
// the [property@Gtk.ColumnView:reorderable] and
// [property@Gtk.ColumnViewColumn:resizable] properties.
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
//
// CSS nodes
//
// “`
// columnview[.column-separators][.rich-list][.navigation-sidebar][.data-table]
// ├── header │ ├── <column header> ┊ ┊ │ ╰── <column header> │ ├── listview │ ┊
// ╰── [rubberband] “`
//
// `GtkColumnView` uses a single CSS node named columnview. It may carry the
// .column-separators style class, when
// [property@Gtk.ColumnView:show-column-separators] property is set. Header
// widgets appear below a node with name header. The rows are contained in a
// `GtkListView` widget, so there is a listview node with the same structure as
// for a standalone `GtkListView` widget. If
// [property@Gtk.ColumnView:show-row-separators] is set, it will be passed on to
// the list view, causing its CSS node to carry the .separators style class. For
// rubberband selection, a node with name rubberband is used.
//
// The main columnview node may also carry style classes to select the style of
// list presentation (section-list-widget.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// `GtkColumnView` uses the GTK_ACCESSIBLE_ROLE_TREE_GRID role, header title
// widgets are using the GTK_ACCESSIBLE_ROLE_COLUMN_HEADER role. The row widgets
// are using the GTK_ACCESSIBLE_ROLE_ROW role, and individual cells are using
// the GTK_ACCESSIBLE_ROLE_GRID_CELL role
type ColumnView interface {
	gextras.Objector

	// AppendColumn appends the @column to the end of the columns in @self.
	AppendColumn(column ColumnViewColumn)
	// EnableRubberband returns whether rows can be selected by dragging with
	// the mouse.
	EnableRubberband() bool
	// Reorderable returns whether columns are reorderable.
	Reorderable() bool
	// ShowColumnSeparators returns whether the list should show separators
	// between columns.
	ShowColumnSeparators() bool
	// ShowRowSeparators returns whether the list should show separators between
	// rows.
	ShowRowSeparators() bool
	// SingleClickActivate returns whether rows will be activated on single
	// click and selected on hover.
	SingleClickActivate() bool
	// Sorter returns a special sorter that reflects the users sorting choices
	// in the column view.
	//
	// To allow users to customizable sorting by clicking on column headers,
	// this sorter needs to be set on the sort model underneath the model that
	// is displayed by the view.
	//
	// See [method@Gtk.ColumnViewColumn.set_sorter] for setting up per-column
	// sorting.
	//
	// Here is an example: “`c gtk_column_view_column_set_sorter (column,
	// sorter); gtk_column_view_append_column (view, column); sorter =
	// g_object_ref (gtk_column_view_get_sorter (view))); model =
	// gtk_sort_list_model_new (store, sorter); selection = gtk_no_selection_new
	// (model); gtk_column_view_set_model (view, selection); “`
	Sorter() *SorterClass
	// InsertColumn inserts a column at the given position in the columns of
	// @self.
	//
	// If @column is already a column of @self, it will be repositioned.
	InsertColumn(position uint, column ColumnViewColumn)
	// RemoveColumn removes the @column from the list of columns of @self.
	RemoveColumn(column ColumnViewColumn)
	// SetEnableRubberband sets whether selections can be changed by dragging
	// with the mouse.
	SetEnableRubberband(enableRubberband bool)
	// SetReorderable sets whether columns should be reorderable by dragging.
	SetReorderable(reorderable bool)
	// SetShowColumnSeparators sets whether the list should show separators
	// between columns.
	SetShowColumnSeparators(showColumnSeparators bool)
	// SetShowRowSeparators sets whether the list should show separators between
	// rows.
	SetShowRowSeparators(showRowSeparators bool)
	// SetSingleClickActivate sets whether rows should be activated on single
	// click and selected on hover.
	SetSingleClickActivate(singleClickActivate bool)
}

// ColumnViewClass implements the ColumnView interface.
type ColumnViewClass struct {
	*externglib.Object
	WidgetClass
	AccessibleIface
	BuildableIface
	ConstraintTargetIface
	ScrollableIface
}

var _ ColumnView = (*ColumnViewClass)(nil)

func wrapColumnView(obj *externglib.Object) ColumnView {
	return &ColumnViewClass{
		Object: obj,
		WidgetClass: WidgetClass{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			AccessibleIface: AccessibleIface{
				Object: obj,
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
			ConstraintTargetIface: ConstraintTargetIface{
				Object: obj,
			},
		},
		AccessibleIface: AccessibleIface{
			Object: obj,
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
		ConstraintTargetIface: ConstraintTargetIface{
			Object: obj,
		},
		ScrollableIface: ScrollableIface{
			Object: obj,
		},
	}
}

func marshalColumnView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColumnView(obj), nil
}

// AppendColumn appends the @column to the end of the columns in @self.
func (self *ColumnViewClass) AppendColumn(column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_append_column(_arg0, _arg1)
}

// EnableRubberband returns whether rows can be selected by dragging with the
// mouse.
func (self *ColumnViewClass) EnableRubberband() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_enable_rubberband(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reorderable returns whether columns are reorderable.
func (self *ColumnViewClass) Reorderable() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_reorderable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowColumnSeparators returns whether the list should show separators between
// columns.
func (self *ColumnViewClass) ShowColumnSeparators() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_show_column_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRowSeparators returns whether the list should show separators between
// rows.
func (self *ColumnViewClass) ShowRowSeparators() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_show_row_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SingleClickActivate returns whether rows will be activated on single click
// and selected on hover.
func (self *ColumnViewClass) SingleClickActivate() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_single_click_activate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sorter returns a special sorter that reflects the users sorting choices in
// the column view.
//
// To allow users to customizable sorting by clicking on column headers, this
// sorter needs to be set on the sort model underneath the model that is
// displayed by the view.
//
// See [method@Gtk.ColumnViewColumn.set_sorter] for setting up per-column
// sorting.
//
// Here is an example: “`c gtk_column_view_column_set_sorter (column, sorter);
// gtk_column_view_append_column (view, column); sorter = g_object_ref
// (gtk_column_view_get_sorter (view))); model = gtk_sort_list_model_new (store,
// sorter); selection = gtk_no_selection_new (model); gtk_column_view_set_model
// (view, selection); “`
func (self *ColumnViewClass) Sorter() *SorterClass {
	var _arg0 *C.GtkColumnView // out
	var _cret *C.GtkSorter     // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_column_view_get_sorter(_arg0)

	var _sorter *SorterClass // out

	_sorter = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*SorterClass)

	return _sorter
}

// InsertColumn inserts a column at the given position in the columns of @self.
//
// If @column is already a column of @self, it will be repositioned.
func (self *ColumnViewClass) InsertColumn(position uint, column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 C.guint                // out
	var _arg2 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_insert_column(_arg0, _arg1, _arg2)
}

// RemoveColumn removes the @column from the list of columns of @self.
func (self *ColumnViewClass) RemoveColumn(column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_remove_column(_arg0, _arg1)
}

// SetEnableRubberband sets whether selections can be changed by dragging with
// the mouse.
func (self *ColumnViewClass) SetEnableRubberband(enableRubberband bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if enableRubberband {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_enable_rubberband(_arg0, _arg1)
}

// SetReorderable sets whether columns should be reorderable by dragging.
func (self *ColumnViewClass) SetReorderable(reorderable bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_reorderable(_arg0, _arg1)
}

// SetShowColumnSeparators sets whether the list should show separators between
// columns.
func (self *ColumnViewClass) SetShowColumnSeparators(showColumnSeparators bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if showColumnSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_show_column_separators(_arg0, _arg1)
}

// SetShowRowSeparators sets whether the list should show separators between
// rows.
func (self *ColumnViewClass) SetShowRowSeparators(showRowSeparators bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if showRowSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_show_row_separators(_arg0, _arg1)
}

// SetSingleClickActivate sets whether rows should be activated on single click
// and selected on hover.
func (self *ColumnViewClass) SetSingleClickActivate(singleClickActivate bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(self.Native()))
	if singleClickActivate {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_single_click_activate(_arg0, _arg1)
}
