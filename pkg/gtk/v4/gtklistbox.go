// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_ListBoxForeachFunc(GtkListBox*, GtkListBoxRow*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_box_get_type()), F: marshalListBox},
		{T: externglib.Type(C.gtk_list_box_row_get_type()), F: marshalListBoxRow},
	})
}

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a
// `GListModel` with gtk_list_box_bind_model() for each item that gets added to
// the model.
type ListBoxCreateWidgetFunc func(item *externglib.Object, userData interface{}) (widget *WidgetClass)

//export gotk4_ListBoxCreateWidgetFunc
func gotk4_ListBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) (cret *C.GtkWidget) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out
	var userData interface{}    // out

	item = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*externglib.Object)
	userData = box.Get(uintptr(arg1))

	fn := v.(ListBoxCreateWidgetFunc)
	widget := fn(item, userData)

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return cret
}

// ListBoxFilterFunc: will be called whenever the row changes or is added and
// lets you control if the row should be visible or not.
type ListBoxFilterFunc func(row *ListBoxRowClass, userData interface{}) (ok bool)

//export gotk4_ListBoxFilterFunc
func gotk4_ListBoxFilterFunc(arg0 *C.GtkListBoxRow, arg1 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var row *ListBoxRowClass // out
	var userData interface{} // out

	row = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*ListBoxRowClass)
	userData = box.Get(uintptr(arg1))

	fn := v.(ListBoxFilterFunc)
	ok := fn(row, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ListBoxForeachFunc: function used by gtk_list_box_selected_foreach().
//
// It will be called on every selected child of the @box.
type ListBoxForeachFunc func(box *ListBoxClass, row *ListBoxRowClass, userData interface{})

//export gotk4_ListBoxForeachFunc
func gotk4_ListBoxForeachFunc(arg0 *C.GtkListBox, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var box *ListBoxClass    // out
	var row *ListBoxRowClass // out
	var userData interface{} // out

	box = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*ListBoxClass)
	row = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*ListBoxRowClass)
	userData = box.Get(uintptr(arg2))

	fn := v.(ListBoxForeachFunc)
	fn(box, row, userData)
}

// ListBoxSortFunc: compare two rows to determine which should be first.
type ListBoxSortFunc func(row1 *ListBoxRowClass, row2 *ListBoxRowClass, userData interface{}) (gint int)

//export gotk4_ListBoxSortFunc
func gotk4_ListBoxSortFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) (cret C.int) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row1 *ListBoxRowClass // out
	var row2 *ListBoxRowClass // out
	var userData interface{}  // out

	row1 = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*ListBoxRowClass)
	row2 = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*ListBoxRowClass)
	userData = box.Get(uintptr(arg2))

	fn := v.(ListBoxSortFunc)
	gint := fn(row1, row2, userData)

	cret = C.int(gint)

	return cret
}

// ListBoxUpdateHeaderFunc: whenever @row changes or which row is before @row
// changes this is called, which lets you update the header on @row.
//
// You may remove or set a new one via [method@Gtk.ListBoxRow.set_header] or
// just change the state of the current header widget.
type ListBoxUpdateHeaderFunc func(row *ListBoxRowClass, before *ListBoxRowClass, userData interface{})

//export gotk4_ListBoxUpdateHeaderFunc
func gotk4_ListBoxUpdateHeaderFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row *ListBoxRowClass    // out
	var before *ListBoxRowClass // out
	var userData interface{}    // out

	row = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*ListBoxRowClass)
	before = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*ListBoxRowClass)
	userData = box.Get(uintptr(arg2))

	fn := v.(ListBoxUpdateHeaderFunc)
	fn(row, before, userData)
}

// ListBox: `GtkListBox` is a vertical list.
//
// A `GtkListBox` only contains `GtkListBoxRow` children. These rows can by
// dynamically sorted and filtered, and headers can be added dynamically
// depending on the row content. It also allows keyboard and mouse navigation
// and selection like a typical list.
//
// Using `GtkListBox` is often an alternative to `GtkTreeView`, especially when
// the list contents has a more complicated layout than what is allowed by a
// `GtkCellRenderer`, or when the contents is interactive (i.e. has a button in
// it).
//
// Although a `GtkListBox` must have only `GtkListBoxRow` children, you can add
// any kind of widget to it via [method@Gtk.ListBox.prepend],
// [method@Gtk.ListBox.append] and [method@Gtk.ListBox.insert] and a
// `GtkListBoxRow` widget will automatically be inserted between the list and
// the widget.
//
// `GtkListBoxRows` can be marked as activatable or selectable. If a row is
// activatable, [signal@Gtk.ListBox::row-activated] will be emitted for it when
// the user tries to activate it. If it is selectable, the row will be marked as
// selected when the user tries to select it.
//
//
// GtkListBox as GtkBuildable
//
// The `GtkListBox` implementation of the `GtkBuildable` interface supports
// setting a child as the placeholder by specifying “placeholder” as the “type”
// attribute of a <child> element. See [method@Gtk.ListBox.set_placeholder] for
// info.
//
// CSS nodes
//
//    list[.separators][.rich-list][.navigation-sidebar]
//    ╰── row[.activatable]
//
// `GtkListBox` uses a single CSS node named list. It may carry the .separators
// style class, when the [property@Gtk.ListBox:show-separators] property is set.
// Each `GtkListBoxRow` uses a single CSS node named row. The row nodes get the
// .activatable style class added when appropriate.
//
// The main list node may also carry style classes to select the style of list
// presentation (section-list-widget.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// `GtkListBox` uses the GTK_ACCESSIBLE_ROLE_LIST role and `GtkListBoxRow` uses
// the GTK_ACCESSIBLE_ROLE_LIST_ITEM role.
type ListBox interface {
	gextras.Objector

	// Append a widget to the list.
	//
	// If a sort function is set, the widget will actually be inserted at the
	// calculated position.
	Append(child Widget)
	// DragHighlightRow: add a drag highlight to a row.
	//
	// This is a helper function for implementing DnD onto a `GtkListBox`. The
	// passed in @row will be highlighted by setting the
	// GTK_STATE_FLAG_DROP_ACTIVE state and any previously highlighted row will
	// be unhighlighted.
	//
	// The row will also be unhighlighted when the widget gets a drag leave
	// event.
	DragHighlightRow(row ListBoxRow)
	// DragUnhighlightRow: if a row has previously been highlighted via
	// gtk_list_box_drag_highlight_row(), it will have the highlight removed.
	DragUnhighlightRow()
	// ActivateOnSingleClick returns whether rows activate on single clicks.
	ActivateOnSingleClick() bool
	// Adjustment gets the adjustment (if any) that the widget uses to for
	// vertical scrolling.
	Adjustment() *AdjustmentClass
	// RowAtIndex gets the n-th child in the list (not counting headers).
	//
	// If @index_ is negative or larger than the number of items in the list,
	// nil is returned.
	RowAtIndex(index_ int) *ListBoxRowClass
	// RowAtY gets the row at the @y position.
	RowAtY(y int) *ListBoxRowClass
	// SelectedRow gets the selected row, or nil if no rows are selected.
	//
	// Note that the box may allow multiple selection, in which case you should
	// use [method@Gtk.ListBox.selected_foreach] to find all selected rows.
	SelectedRow() *ListBoxRowClass
	// SelectionMode gets the selection mode of the listbox.
	SelectionMode() SelectionMode
	// ShowSeparators returns whether the list box should show separators
	// between rows.
	ShowSeparators() bool
	// Insert the @child into the @box at @position.
	//
	// If a sort function is set, the widget will actually be inserted at the
	// calculated position.
	//
	// If @position is -1, or larger than the total number of items in the @box,
	// then the @child will be appended to the end.
	Insert(child Widget, position int)
	// InvalidateFilter: update the filtering for all rows.
	//
	// Call this when result of the filter function on the @box is changed due
	// to an external factor. For instance, this would be used if the filter
	// function just looked for a specific search string and the entry with the
	// search string has changed.
	InvalidateFilter()
	// InvalidateHeaders: update the separators for all rows.
	//
	// Call this when result of the header function on the @box is changed due
	// to an external factor.
	InvalidateHeaders()
	// InvalidateSort: update the sorting for all rows.
	//
	// Call this when result of the sort function on the @box is changed due to
	// an external factor.
	InvalidateSort()
	// Prepend a widget to the list.
	//
	// If a sort function is set, the widget will actually be inserted at the
	// calculated position.
	Prepend(child Widget)
	// Remove removes a child from @box.
	Remove(child Widget)
	// SelectAll: select all children of @box, if the selection mode allows it.
	SelectAll()
	// SelectRow: make @row the currently selected row.
	SelectRow(row ListBoxRow)
	// SelectedForeach calls a function for each selected child.
	//
	// Note that the selection cannot be modified from within this function.
	SelectedForeach(fn ListBoxForeachFunc)
	// SetActivateOnSingleClick: if @single is true, rows will be activated when
	// you click on them, otherwise you need to double-click.
	SetActivateOnSingleClick(single bool)
	// SetAdjustment sets the adjustment (if any) that the widget uses to for
	// vertical scrolling.
	//
	// For instance, this is used to get the page size for PageUp/Down key
	// handling.
	//
	// In the normal case when the @box is packed inside a `GtkScrolledWindow`
	// the adjustment from that will be picked up automatically, so there is no
	// need to manually do that.
	SetAdjustment(adjustment Adjustment)
	// SetPlaceholder sets the placeholder widget that is shown in the list when
	// it doesn't display any visible children.
	SetPlaceholder(placeholder Widget)
	// SetShowSeparators sets whether the list box should show separators
	// between rows.
	SetShowSeparators(showSeparators bool)
	// UnselectAll: unselect all children of @box, if the selection mode allows
	// it.
	UnselectAll()
	// UnselectRow unselects a single row of @box, if the selection mode allows
	// it.
	UnselectRow(row ListBoxRow)
}

// ListBoxClass implements the ListBox interface.
type ListBoxClass struct {
	*externglib.Object
	WidgetClass
	AccessibleIface
	BuildableIface
	ConstraintTargetIface
}

var _ ListBox = (*ListBoxClass)(nil)

func wrapListBox(obj *externglib.Object) ListBox {
	return &ListBoxClass{
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
	}
}

func marshalListBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBox(obj), nil
}

// NewListBox creates a new `GtkListBox` container.
func NewListBox() *ListBoxClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_new()

	var _listBox *ListBoxClass // out

	_listBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxClass)

	return _listBox
}

// Append a widget to the list.
//
// If a sort function is set, the widget will actually be inserted at the
// calculated position.
func (box *ListBoxClass) Append(child Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_append(_arg0, _arg1)
}

// DragHighlightRow: add a drag highlight to a row.
//
// This is a helper function for implementing DnD onto a `GtkListBox`. The
// passed in @row will be highlighted by setting the GTK_STATE_FLAG_DROP_ACTIVE
// state and any previously highlighted row will be unhighlighted.
//
// The row will also be unhighlighted when the widget gets a drag leave event.
func (box *ListBoxClass) DragHighlightRow(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_drag_highlight_row(_arg0, _arg1)
}

// DragUnhighlightRow: if a row has previously been highlighted via
// gtk_list_box_drag_highlight_row(), it will have the highlight removed.
func (box *ListBoxClass) DragUnhighlightRow() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_drag_unhighlight_row(_arg0)
}

// ActivateOnSingleClick returns whether rows activate on single clicks.
func (box *ListBoxClass) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkListBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_activate_on_single_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Adjustment gets the adjustment (if any) that the widget uses to for vertical
// scrolling.
func (box *ListBoxClass) Adjustment() *AdjustmentClass {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_adjustment(_arg0)

	var _adjustment *AdjustmentClass // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*AdjustmentClass)

	return _adjustment
}

// RowAtIndex gets the n-th child in the list (not counting headers).
//
// If @index_ is negative or larger than the number of items in the list, nil is
// returned.
func (box *ListBoxClass) RowAtIndex(index_ int) *ListBoxRowClass {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.int            // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.int(index_)

	_cret = C.gtk_list_box_get_row_at_index(_arg0, _arg1)

	var _listBoxRow *ListBoxRowClass // out

	_listBoxRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxRowClass)

	return _listBoxRow
}

// RowAtY gets the row at the @y position.
func (box *ListBoxClass) RowAtY(y int) *ListBoxRowClass {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.int            // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.int(y)

	_cret = C.gtk_list_box_get_row_at_y(_arg0, _arg1)

	var _listBoxRow *ListBoxRowClass // out

	_listBoxRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxRowClass)

	return _listBoxRow
}

// SelectedRow gets the selected row, or nil if no rows are selected.
//
// Note that the box may allow multiple selection, in which case you should use
// [method@Gtk.ListBox.selected_foreach] to find all selected rows.
func (box *ListBoxClass) SelectedRow() *ListBoxRowClass {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selected_row(_arg0)

	var _listBoxRow *ListBoxRowClass // out

	_listBoxRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxRowClass)

	return _listBoxRow
}

// SelectionMode gets the selection mode of the listbox.
func (box *ListBoxClass) SelectionMode() SelectionMode {
	var _arg0 *C.GtkListBox      // out
	var _cret C.GtkSelectionMode // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_selection_mode(_arg0)

	var _selectionMode SelectionMode // out

	_selectionMode = (SelectionMode)(_cret)

	return _selectionMode
}

// ShowSeparators returns whether the list box should show separators between
// rows.
func (box *ListBoxClass) ShowSeparators() bool {
	var _arg0 *C.GtkListBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_list_box_get_show_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Insert the @child into the @box at @position.
//
// If a sort function is set, the widget will actually be inserted at the
// calculated position.
//
// If @position is -1, or larger than the total number of items in the @box,
// then the @child will be appended to the end.
func (box *ListBoxClass) Insert(child Widget, position int) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(position)

	C.gtk_list_box_insert(_arg0, _arg1, _arg2)
}

// InvalidateFilter: update the filtering for all rows.
//
// Call this when result of the filter function on the @box is changed due to an
// external factor. For instance, this would be used if the filter function just
// looked for a specific search string and the entry with the search string has
// changed.
func (box *ListBoxClass) InvalidateFilter() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_filter(_arg0)
}

// InvalidateHeaders: update the separators for all rows.
//
// Call this when result of the header function on the @box is changed due to an
// external factor.
func (box *ListBoxClass) InvalidateHeaders() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_headers(_arg0)
}

// InvalidateSort: update the sorting for all rows.
//
// Call this when result of the sort function on the @box is changed due to an
// external factor.
func (box *ListBoxClass) InvalidateSort() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_invalidate_sort(_arg0)
}

// Prepend a widget to the list.
//
// If a sort function is set, the widget will actually be inserted at the
// calculated position.
func (box *ListBoxClass) Prepend(child Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_prepend(_arg0, _arg1)
}

// Remove removes a child from @box.
func (box *ListBoxClass) Remove(child Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_remove(_arg0, _arg1)
}

// SelectAll: select all children of @box, if the selection mode allows it.
func (box *ListBoxClass) SelectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_select_all(_arg0)
}

// SelectRow: make @row the currently selected row.
func (box *ListBoxClass) SelectRow(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_select_row(_arg0, _arg1)
}

// SelectedForeach calls a function for each selected child.
//
// Note that the selection cannot be modified from within this function.
func (box *ListBoxClass) SelectedForeach(fn ListBoxForeachFunc) {
	var _arg0 *C.GtkListBox           // out
	var _arg1 C.GtkListBoxForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*[0]byte)(C.gotk4_ListBoxForeachFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_list_box_selected_foreach(_arg0, _arg1, _arg2)
}

// SetActivateOnSingleClick: if @single is true, rows will be activated when you
// click on them, otherwise you need to double-click.
func (box *ListBoxClass) SetActivateOnSingleClick(single bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_set_activate_on_single_click(_arg0, _arg1)
}

// SetAdjustment sets the adjustment (if any) that the widget uses to for
// vertical scrolling.
//
// For instance, this is used to get the page size for PageUp/Down key handling.
//
// In the normal case when the @box is packed inside a `GtkScrolledWindow` the
// adjustment from that will be picked up automatically, so there is no need to
// manually do that.
func (box *ListBoxClass) SetAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_list_box_set_adjustment(_arg0, _arg1)
}

// SetPlaceholder sets the placeholder widget that is shown in the list when it
// doesn't display any visible children.
func (box *ListBoxClass) SetPlaceholder(placeholder Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(placeholder.Native()))

	C.gtk_list_box_set_placeholder(_arg0, _arg1)
}

// SetShowSeparators sets whether the list box should show separators between
// rows.
func (box *ListBoxClass) SetShowSeparators(showSeparators bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	if showSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_set_show_separators(_arg0, _arg1)
}

// UnselectAll: unselect all children of @box, if the selection mode allows it.
func (box *ListBoxClass) UnselectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))

	C.gtk_list_box_unselect_all(_arg0)
}

// UnselectRow unselects a single row of @box, if the selection mode allows it.
func (box *ListBoxClass) UnselectRow(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_unselect_row(_arg0, _arg1)
}

// ListBoxRowOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ListBoxRowOverrider interface {
	Activate()
}

// ListBoxRow: `GtkListBoxRow` is the kind of widget that can be added to a
// `GtkListBox`.
type ListBoxRow interface {
	gextras.Objector

	// Changed marks @row as changed, causing any state that depends on this to
	// be updated.
	//
	// This affects sorting, filtering and headers.
	//
	// Note that calls to this method must be in sync with the data used for the
	// row functions. For instance, if the list is mirroring some external data
	// set, and *two* rows changed in the external data set then when you call
	// gtk_list_box_row_changed() on the first row the sort function must only
	// read the new data for the first of the two changed rows, otherwise the
	// resorting of the rows will be wrong.
	//
	// This generally means that if you don’t fully control the data model you
	// have to duplicate the data that affects the listbox row functions into
	// the row widgets themselves. Another alternative is to call
	// [method@Gtk.ListBox.invalidate_sort] on any model change, but that is
	// more expensive.
	Changed()
	// Activatable gets whether the row is activatable.
	Activatable() bool
	// Child gets the child widget of @row.
	Child() *WidgetClass
	// Header returns the current header of the @row.
	//
	// This can be used in a [callback@Gtk.ListBoxUpdateHeaderFunc] to see if
	// there is a header set already, and if so to update the state of it.
	Header() *WidgetClass
	// Index gets the current index of the @row in its `GtkListBox` container.
	Index() int
	// Selectable gets whether the row can be selected.
	Selectable() bool
	// IsSelected returns whether the child is currently selected in its
	// `GtkListBox` container.
	IsSelected() bool
	// SetActivatable: set whether the row is activatable.
	SetActivatable(activatable bool)
	// SetChild sets the child widget of @self.
	SetChild(child Widget)
	// SetHeader sets the current header of the @row.
	//
	// This is only allowed to be called from a
	// [callback@Gtk.ListBoxUpdateHeaderFunc]. It will replace any existing
	// header in the row, and be shown in front of the row in the listbox.
	SetHeader(header Widget)
	// SetSelectable: set whether the row can be selected.
	SetSelectable(selectable bool)
}

// ListBoxRowClass implements the ListBoxRow interface.
type ListBoxRowClass struct {
	*externglib.Object
	WidgetClass
	AccessibleIface
	ActionableIface
	BuildableIface
	ConstraintTargetIface
}

var _ ListBoxRow = (*ListBoxRowClass)(nil)

func wrapListBoxRow(obj *externglib.Object) ListBoxRow {
	return &ListBoxRowClass{
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
		ActionableIface: ActionableIface{
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
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
		ConstraintTargetIface: ConstraintTargetIface{
			Object: obj,
		},
	}
}

func marshalListBoxRow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListBoxRow(obj), nil
}

// NewListBoxRow creates a new `GtkListBoxRow`.
func NewListBoxRow() *ListBoxRowClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_row_new()

	var _listBoxRow *ListBoxRowClass // out

	_listBoxRow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ListBoxRowClass)

	return _listBoxRow
}

// Changed marks @row as changed, causing any state that depends on this to be
// updated.
//
// This affects sorting, filtering and headers.
//
// Note that calls to this method must be in sync with the data used for the row
// functions. For instance, if the list is mirroring some external data set, and
// *two* rows changed in the external data set then when you call
// gtk_list_box_row_changed() on the first row the sort function must only read
// the new data for the first of the two changed rows, otherwise the resorting
// of the rows will be wrong.
//
// This generally means that if you don’t fully control the data model you have
// to duplicate the data that affects the listbox row functions into the row
// widgets themselves. Another alternative is to call
// [method@Gtk.ListBox.invalidate_sort] on any model change, but that is more
// expensive.
func (row *ListBoxRowClass) Changed() {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_row_changed(_arg0)
}

// Activatable gets whether the row is activatable.
func (row *ListBoxRowClass) Activatable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child widget of @row.
func (row *ListBoxRowClass) Child() *WidgetClass {
	var _arg0 *C.GtkListBoxRow // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_child(_arg0)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// Header returns the current header of the @row.
//
// This can be used in a [callback@Gtk.ListBoxUpdateHeaderFunc] to see if there
// is a header set already, and if so to update the state of it.
func (row *ListBoxRowClass) Header() *WidgetClass {
	var _arg0 *C.GtkListBoxRow // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_header(_arg0)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// Index gets the current index of the @row in its `GtkListBox` container.
func (row *ListBoxRowClass) Index() int {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.int            // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_index(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Selectable gets whether the row can be selected.
func (row *ListBoxRowClass) Selectable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_get_selectable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSelected returns whether the child is currently selected in its
// `GtkListBox` container.
func (row *ListBoxRowClass) IsSelected() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	_cret = C.gtk_list_box_row_is_selected(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatable: set whether the row is activatable.
func (row *ListBoxRowClass) SetActivatable(activatable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if activatable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_activatable(_arg0, _arg1)
}

// SetChild sets the child widget of @self.
func (row *ListBoxRowClass) SetChild(child Widget) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_row_set_child(_arg0, _arg1)
}

// SetHeader sets the current header of the @row.
//
// This is only allowed to be called from a
// [callback@Gtk.ListBoxUpdateHeaderFunc]. It will replace any existing header
// in the row, and be shown in front of the row in the listbox.
func (row *ListBoxRowClass) SetHeader(header Widget) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(header.Native()))

	C.gtk_list_box_row_set_header(_arg0, _arg1)
}

// SetSelectable: set whether the row can be selected.
func (row *ListBoxRowClass) SetSelectable(selectable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))
	if selectable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_selectable(_arg0, _arg1)
}
