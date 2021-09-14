// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// GtkWidget* _gotk4_gtk3_FlowBoxCreateWidgetFunc(gpointer, gpointer);
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk3_FlowBoxFilterFunc(GtkFlowBoxChild*, gpointer);
// gint _gotk4_gtk3_FlowBoxSortFunc(GtkFlowBoxChild*, GtkFlowBoxChild*, gpointer);
// void _gotk4_gtk3_FlowBoxForeachFunc(GtkFlowBox*, GtkFlowBoxChild*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_flow_box_get_type()), F: marshalFlowBoxer},
		{T: externglib.Type(C.gtk_flow_box_child_get_type()), F: marshalFlowBoxChilder},
	})
}

// FlowBoxCreateWidgetFunc: called for flow boxes that are bound to a Model with
// gtk_flow_box_bind_model() for each item that gets added to the model.
type FlowBoxCreateWidgetFunc func(item *externglib.Object) (widget Widgetter)

//export _gotk4_gtk3_FlowBoxCreateWidgetFunc
func _gotk4_gtk3_FlowBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) (cret *C.GtkWidget) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out

	item = externglib.Take(unsafe.Pointer(arg0))

	fn := v.(FlowBoxCreateWidgetFunc)
	widget := fn(item)

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return cret
}

// FlowBoxFilterFunc: function that will be called whenrever a child changes or
// is added. It lets you control if the child should be visible or not.
type FlowBoxFilterFunc func(child *FlowBoxChild) (ok bool)

//export _gotk4_gtk3_FlowBoxFilterFunc
func _gotk4_gtk3_FlowBoxFilterFunc(arg0 *C.GtkFlowBoxChild, arg1 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var child *FlowBoxChild // out

	child = wrapFlowBoxChild(externglib.Take(unsafe.Pointer(arg0)))

	fn := v.(FlowBoxFilterFunc)
	ok := fn(child)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FlowBoxForeachFunc: function used by gtk_flow_box_selected_foreach(). It will
// be called on every selected child of the box.
type FlowBoxForeachFunc func(box *FlowBox, child *FlowBoxChild)

//export _gotk4_gtk3_FlowBoxForeachFunc
func _gotk4_gtk3_FlowBoxForeachFunc(arg0 *C.GtkFlowBox, arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var box *FlowBox        // out
	var child *FlowBoxChild // out

	box = wrapFlowBox(externglib.Take(unsafe.Pointer(arg0)))
	child = wrapFlowBoxChild(externglib.Take(unsafe.Pointer(arg1)))

	fn := v.(FlowBoxForeachFunc)
	fn(box, child)
}

// FlowBoxSortFunc: function to compare two children to determine which should
// come first.
type FlowBoxSortFunc func(child1 *FlowBoxChild, child2 *FlowBoxChild) (gint int)

//export _gotk4_gtk3_FlowBoxSortFunc
func _gotk4_gtk3_FlowBoxSortFunc(arg0 *C.GtkFlowBoxChild, arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var child1 *FlowBoxChild // out
	var child2 *FlowBoxChild // out

	child1 = wrapFlowBoxChild(externglib.Take(unsafe.Pointer(arg0)))
	child2 = wrapFlowBoxChild(externglib.Take(unsafe.Pointer(arg1)))

	fn := v.(FlowBoxSortFunc)
	gint := fn(child1, child2)

	cret = C.gint(gint)

	return cret
}

// FlowBoxOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FlowBoxOverrider interface {
	ActivateCursorChild()
	ChildActivated(child *FlowBoxChild)
	MoveCursor(step MovementStep, count int) bool
	// SelectAll: select all children of box, if the selection mode allows it.
	SelectAll()
	SelectedChildrenChanged()
	ToggleCursorChild()
	// UnselectAll: unselect all children of box, if the selection mode allows
	// it.
	UnselectAll()
}

// FlowBox positions child widgets in sequence according to its orientation.
//
// For instance, with the horizontal orientation, the widgets will be arranged
// from left to right, starting a new row under the previous row when necessary.
// Reducing the width in this case will require more rows, so a larger height
// will be requested.
//
// Likewise, with the vertical orientation, the widgets will be arranged from
// top to bottom, starting a new column to the right when necessary. Reducing
// the height will require more columns, so a larger width will be requested.
//
// The size request of a GtkFlowBox alone may not be what you expect; if you
// need to be able to shrink it along both axes and dynamically reflow its
// children, you may have to wrap it in a ScrolledWindow to enable that.
//
// The children of a GtkFlowBox can be dynamically sorted and filtered.
//
// Although a GtkFlowBox must have only FlowBoxChild children, you can add any
// kind of widget to it via gtk_container_add(), and a GtkFlowBoxChild widget
// will automatically be inserted between the box and the widget.
//
// Also see ListBox.
//
// GtkFlowBox was added in GTK+ 3.12.
//
// CSS nodes
//
//    flowbox
//    ├── flowboxchild
//    │   ╰── <child>
//    ├── flowboxchild
//    │   ╰── <child>
//    ┊
//    ╰── [rubberband]
//
// GtkFlowBox uses a single CSS node with name flowbox. GtkFlowBoxChild uses a
// single CSS node with name flowboxchild. For rubberband selection, a subnode
// with name rubberband is used.
type FlowBox struct {
	Container

	Orientable
	*externglib.Object
}

func wrapFlowBox(obj *externglib.Object) *FlowBox {
	return &FlowBox{
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
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalFlowBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFlowBox(obj), nil
}

// NewFlowBox creates a GtkFlowBox.
func NewFlowBox() *FlowBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_flow_box_new()

	var _flowBox *FlowBox // out

	_flowBox = wrapFlowBox(externglib.Take(unsafe.Pointer(_cret)))

	return _flowBox
}

// BindModel binds model to box.
//
// If box was already bound to a model, that previous binding is destroyed.
//
// The contents of box are cleared and then filled with widgets that represent
// items from model. box is updated whenever model changes. If model is NULL,
// box is left empty.
//
// It is undefined to add or remove widgets directly (for example, with
// gtk_flow_box_insert() or gtk_container_add()) while box is bound to a model.
//
// Note that using a model is incompatible with the filtering and sorting
// functionality in GtkFlowBox. When using a model, filtering and sorting should
// be implemented by the model.
func (box *FlowBox) BindModel(model gio.ListModeller, createWidgetFunc FlowBoxCreateWidgetFunc) {
	var _arg0 *C.GtkFlowBox                // out
	var _arg1 *C.GListModel                // out
	var _arg2 C.GtkFlowBoxCreateWidgetFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}
	_arg2 = (*[0]byte)(C._gotk4_gtk3_FlowBoxCreateWidgetFunc)
	_arg3 = C.gpointer(gbox.Assign(createWidgetFunc))
	_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_flow_box_bind_model(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(box)
	runtime.KeepAlive(model)
	runtime.KeepAlive(createWidgetFunc)
}

// ActivateOnSingleClick returns whether children activate on single clicks.
func (box *FlowBox) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_flow_box_get_activate_on_single_click(_arg0)
	runtime.KeepAlive(box)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ChildAtIndex gets the nth child in the box.
func (box *FlowBox) ChildAtIndex(idx int) *FlowBoxChild {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 C.gint             // out
	var _cret *C.GtkFlowBoxChild // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.gint(idx)

	_cret = C.gtk_flow_box_get_child_at_index(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(idx)

	var _flowBoxChild *FlowBoxChild // out

	if _cret != nil {
		_flowBoxChild = wrapFlowBoxChild(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _flowBoxChild
}

// ChildAtPos gets the child in the (x, y) position.
func (box *FlowBox) ChildAtPos(x int, y int) *FlowBoxChild {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 C.gint             // out
	var _arg2 C.gint             // out
	var _cret *C.GtkFlowBoxChild // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_flow_box_get_child_at_pos(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _flowBoxChild *FlowBoxChild // out

	if _cret != nil {
		_flowBoxChild = wrapFlowBoxChild(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _flowBoxChild
}

// ColumnSpacing gets the horizontal spacing.
func (box *FlowBox) ColumnSpacing() uint {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.guint       // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_flow_box_get_column_spacing(_arg0)
	runtime.KeepAlive(box)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Homogeneous returns whether the box is homogeneous (all children are the same
// size). See gtk_box_set_homogeneous().
func (box *FlowBox) Homogeneous() bool {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_flow_box_get_homogeneous(_arg0)
	runtime.KeepAlive(box)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxChildrenPerLine gets the maximum number of children per line.
func (box *FlowBox) MaxChildrenPerLine() uint {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.guint       // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_flow_box_get_max_children_per_line(_arg0)
	runtime.KeepAlive(box)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MinChildrenPerLine gets the minimum number of children per line.
func (box *FlowBox) MinChildrenPerLine() uint {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.guint       // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_flow_box_get_min_children_per_line(_arg0)
	runtime.KeepAlive(box)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RowSpacing gets the vertical spacing.
func (box *FlowBox) RowSpacing() uint {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.guint       // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_flow_box_get_row_spacing(_arg0)
	runtime.KeepAlive(box)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SelectedChildren creates a list of all selected children.
func (box *FlowBox) SelectedChildren() []FlowBoxChild {
	var _arg0 *C.GtkFlowBox // out
	var _cret *C.GList      // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_flow_box_get_selected_children(_arg0)
	runtime.KeepAlive(box)

	var _list []FlowBoxChild // out

	_list = make([]FlowBoxChild, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkFlowBoxChild)(v)
		var dst FlowBoxChild // out
		dst = *wrapFlowBoxChild(externglib.Take(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// SelectionMode gets the selection mode of box.
func (box *FlowBox) SelectionMode() SelectionMode {
	var _arg0 *C.GtkFlowBox      // out
	var _cret C.GtkSelectionMode // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_flow_box_get_selection_mode(_arg0)
	runtime.KeepAlive(box)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

// Insert inserts the widget into box at position.
//
// If a sort function is set, the widget will actually be inserted at the
// calculated position and this function has the same effect as
// gtk_container_add().
//
// If position is -1, or larger than the total number of children in the box,
// then the widget will be appended to the end.
func (box *FlowBox) Insert(widget Widgetter, position int) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.gint(position)

	C.gtk_flow_box_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(position)
}

// InvalidateFilter updates the filtering for all children.
//
// Call this function when the result of the filter function on the box is
// changed due ot an external factor. For instance, this would be used if the
// filter function just looked for a specific search term, and the entry with
// the string has changed.
func (box *FlowBox) InvalidateFilter() {
	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	C.gtk_flow_box_invalidate_filter(_arg0)
	runtime.KeepAlive(box)
}

// InvalidateSort updates the sorting for all children.
//
// Call this when the result of the sort function on box is changed due to an
// external factor.
func (box *FlowBox) InvalidateSort() {
	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	C.gtk_flow_box_invalidate_sort(_arg0)
	runtime.KeepAlive(box)
}

// SelectAll: select all children of box, if the selection mode allows it.
func (box *FlowBox) SelectAll() {
	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	C.gtk_flow_box_select_all(_arg0)
	runtime.KeepAlive(box)
}

// SelectChild selects a single child of box, if the selection mode allows it.
func (box *FlowBox) SelectChild(child *FlowBoxChild) {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 *C.GtkFlowBoxChild // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_select_child(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// SelectedForeach calls a function for each selected child.
//
// Note that the selection cannot be modified from within this function.
func (box *FlowBox) SelectedForeach(fn FlowBoxForeachFunc) {
	var _arg0 *C.GtkFlowBox           // out
	var _arg1 C.GtkFlowBoxForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_FlowBoxForeachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.gtk_flow_box_selected_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(fn)
}

// SetActivateOnSingleClick: if single is TRUE, children will be activated when
// you click on them, otherwise you need to double-click.
func (box *FlowBox) SetActivateOnSingleClick(single bool) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_flow_box_set_activate_on_single_click(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(single)
}

// SetColumnSpacing sets the horizontal space to add between children. See the
// FlowBox:column-spacing property.
func (box *FlowBox) SetColumnSpacing(spacing uint) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_flow_box_set_column_spacing(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(spacing)
}

// SetFilterFunc: by setting a filter function on the box one can decide
// dynamically which of the children to show. For instance, to implement a
// search function that only shows the children matching the search terms.
//
// The filter_func will be called for each child after the call, and it will
// continue to be called each time a child changes (via
// gtk_flow_box_child_changed()) or when gtk_flow_box_invalidate_filter() is
// called.
//
// Note that using a filter function is incompatible with using a model (see
// gtk_flow_box_bind_model()).
func (box *FlowBox) SetFilterFunc(filterFunc FlowBoxFilterFunc) {
	var _arg0 *C.GtkFlowBox          // out
	var _arg1 C.GtkFlowBoxFilterFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	if filterFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_FlowBoxFilterFunc)
		_arg2 = C.gpointer(gbox.Assign(filterFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_flow_box_set_filter_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(box)
	runtime.KeepAlive(filterFunc)
}

// SetHAdjustment hooks up an adjustment to focus handling in box. The
// adjustment is also used for autoscrolling during rubberband selection. See
// gtk_scrolled_window_get_hadjustment() for a typical way of obtaining the
// adjustment, and gtk_flow_box_set_vadjustment()for setting the vertical
// adjustment.
//
// The adjustments have to be in pixel units and in the same coordinate system
// as the allocation for immediate children of the box.
func (box *FlowBox) SetHAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkFlowBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_flow_box_set_hadjustment(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(adjustment)
}

// SetHomogeneous sets the FlowBox:homogeneous property of box, controlling
// whether or not all children of box are given equal space in the box.
func (box *FlowBox) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_flow_box_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(homogeneous)
}

// SetMaxChildrenPerLine sets the maximum number of children to request and
// allocate space for in box’s orientation.
//
// Setting the maximum number of children per line limits the overall natural
// size request to be no more than n_children children long in the given
// orientation.
func (box *FlowBox) SetMaxChildrenPerLine(nChildren uint) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.guint(nChildren)

	C.gtk_flow_box_set_max_children_per_line(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(nChildren)
}

// SetMinChildrenPerLine sets the minimum number of children to line up in box’s
// orientation before flowing.
func (box *FlowBox) SetMinChildrenPerLine(nChildren uint) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.guint(nChildren)

	C.gtk_flow_box_set_min_children_per_line(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(nChildren)
}

// SetRowSpacing sets the vertical space to add between children. See the
// FlowBox:row-spacing property.
func (box *FlowBox) SetRowSpacing(spacing uint) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_flow_box_set_row_spacing(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(spacing)
}

// SetSelectionMode sets how selection works in box. See SelectionMode for
// details.
func (box *FlowBox) SetSelectionMode(mode SelectionMode) {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 C.GtkSelectionMode // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.GtkSelectionMode(mode)

	C.gtk_flow_box_set_selection_mode(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(mode)
}

// SetSortFunc: by setting a sort function on the box, one can dynamically
// reorder the children of the box, based on the contents of the children.
//
// The sort_func will be called for each child after the call, and will continue
// to be called each time a child changes (via gtk_flow_box_child_changed()) and
// when gtk_flow_box_invalidate_sort() is called.
//
// Note that using a sort function is incompatible with using a model (see
// gtk_flow_box_bind_model()).
func (box *FlowBox) SetSortFunc(sortFunc FlowBoxSortFunc) {
	var _arg0 *C.GtkFlowBox        // out
	var _arg1 C.GtkFlowBoxSortFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	if sortFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_FlowBoxSortFunc)
		_arg2 = C.gpointer(gbox.Assign(sortFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_flow_box_set_sort_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(box)
	runtime.KeepAlive(sortFunc)
}

// SetVAdjustment hooks up an adjustment to focus handling in box. The
// adjustment is also used for autoscrolling during rubberband selection. See
// gtk_scrolled_window_get_vadjustment() for a typical way of obtaining the
// adjustment, and gtk_flow_box_set_hadjustment()for setting the horizontal
// adjustment.
//
// The adjustments have to be in pixel units and in the same coordinate system
// as the allocation for immediate children of the box.
func (box *FlowBox) SetVAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkFlowBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_flow_box_set_vadjustment(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(adjustment)
}

// UnselectAll: unselect all children of box, if the selection mode allows it.
func (box *FlowBox) UnselectAll() {
	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))

	C.gtk_flow_box_unselect_all(_arg0)
	runtime.KeepAlive(box)
}

// UnselectChild unselects a single child of box, if the selection mode allows
// it.
func (box *FlowBox) UnselectChild(child *FlowBoxChild) {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 *C.GtkFlowBoxChild // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_unselect_child(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// FlowBoxChildOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FlowBoxChildOverrider interface {
	Activate()
}

type FlowBoxChild struct {
	Bin
}

func wrapFlowBoxChild(obj *externglib.Object) *FlowBoxChild {
	return &FlowBoxChild{
		Bin: Bin{
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
					Object: obj,
				},
			},
		},
	}
}

func marshalFlowBoxChilder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFlowBoxChild(obj), nil
}

// NewFlowBoxChild creates a new FlowBoxChild, to be used as a child of a
// FlowBox.
func NewFlowBoxChild() *FlowBoxChild {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_flow_box_child_new()

	var _flowBoxChild *FlowBoxChild // out

	_flowBoxChild = wrapFlowBoxChild(externglib.Take(unsafe.Pointer(_cret)))

	return _flowBoxChild
}

// Changed marks child as changed, causing any state that depends on this to be
// updated. This affects sorting and filtering.
//
// Note that calls to this method must be in sync with the data used for the
// sorting and filtering functions. For instance, if the list is mirroring some
// external data set, and *two* children changed in the external data set when
// you call gtk_flow_box_child_changed() on the first child, the sort function
// must only read the new data for the first of the two changed children,
// otherwise the resorting of the children will be wrong.
//
// This generally means that if you don’t fully control the data model, you have
// to duplicate the data that affects the sorting and filtering functions into
// the widgets themselves. Another alternative is to call
// gtk_flow_box_invalidate_sort() on any model change, but that is more
// expensive.
func (child *FlowBoxChild) Changed() {
	var _arg0 *C.GtkFlowBoxChild // out

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_child_changed(_arg0)
	runtime.KeepAlive(child)
}

// Index gets the current index of the child in its FlowBox container.
func (child *FlowBoxChild) Index() int {
	var _arg0 *C.GtkFlowBoxChild // out
	var _cret C.gint             // in

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_flow_box_child_get_index(_arg0)
	runtime.KeepAlive(child)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsSelected returns whether the child is currently selected in its FlowBox
// container.
func (child *FlowBoxChild) IsSelected() bool {
	var _arg0 *C.GtkFlowBoxChild // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_flow_box_child_is_selected(_arg0)
	runtime.KeepAlive(child)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
