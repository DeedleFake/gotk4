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
		{T: externglib.Type(C.gtk_layout_manager_get_type()), F: marshalLayoutManager},
	})
}

// LayoutManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type LayoutManagerOverrider interface {
	// Allocate assigns the given @width, @height, and @baseline to a @widget,
	// and computes the position and sizes of the children of the @widget using
	// the layout management policy of @manager.
	Allocate(widget Widget, width int, height int, baseline int)
	// CreateLayoutChild: create a LayoutChild instance for the given @for_child
	// widget.
	CreateLayoutChild(widget Widget, forChild Widget) *LayoutChildClass
	RequestMode(widget Widget) SizeRequestMode
	Root()
	Unroot()
}

// LayoutManager: layout managers are delegate classes that handle the preferred
// size and the allocation of a widget.
//
// You typically subclass `GtkLayoutManager` if you want to implement a layout
// policy for the children of a widget, or if you want to determine the size of
// a widget depending on its contents.
//
// Each `GtkWidget` can only have a `GtkLayoutManager` instance associated to it
// at any given time; it is possible, though, to replace the layout manager
// instance using [method@Gtk.Widget.set_layout_manager].
//
//
// Layout properties
//
// A layout manager can expose properties for controlling the layout of each
// child, by creating an object type derived from [class@Gtk.LayoutChild] and
// installing the properties on it as normal `GObject` properties.
//
// Each `GtkLayoutChild` instance storing the layout properties for a specific
// child is created through the [method@Gtk.LayoutManager.get_layout_child]
// method; a `GtkLayoutManager` controls the creation of its `GtkLayoutChild`
// instances by overriding the GtkLayoutManagerClass.create_layout_child()
// virtual function. The typical implementation should look like:
//
// “`c static GtkLayoutChild * create_layout_child (GtkLayoutManager *manager,
// GtkWidget *container, GtkWidget *child) { return g_object_new
// (your_layout_child_get_type (), "layout-manager", manager, "child-widget",
// child, NULL); } “`
//
// The [property@Gtk.LayoutChild:layout-manager] and
// [property@Gtk.LayoutChild:child-widget] properties on the newly created
// `GtkLayoutChild` instance are mandatory. The `GtkLayoutManager` will cache
// the newly created `GtkLayoutChild` instance until the widget is removed from
// its parent, or the parent removes the layout manager.
//
// Each `GtkLayoutManager` instance creating a `GtkLayoutChild` should use
// [method@Gtk.LayoutManager.get_layout_child] every time it needs to query the
// layout properties; each `GtkLayoutChild` instance should call
// [method@Gtk.LayoutManager.layout_changed] every time a property is updated,
// in order to queue a new size measuring and allocation.
type LayoutManager interface {
	gextras.Objector

	// Allocate assigns the given @width, @height, and @baseline to a @widget,
	// and computes the position and sizes of the children of the @widget using
	// the layout management policy of @manager.
	Allocate(widget Widget, width int, height int, baseline int)
	// LayoutChild retrieves a `GtkLayoutChild` instance for the
	// `GtkLayoutManager`, creating one if necessary.
	//
	// The @child widget must be a child of the widget using @manager.
	//
	// The `GtkLayoutChild` instance is owned by the `GtkLayoutManager`, and is
	// guaranteed to exist as long as @child is a child of the `GtkWidget` using
	// the given `GtkLayoutManager`.
	LayoutChild(child Widget) *LayoutChildClass
	// RequestMode retrieves the request mode of @manager.
	RequestMode() SizeRequestMode
	// Widget retrieves the `GtkWidget` using the given `GtkLayoutManager`.
	Widget() *WidgetClass
	// LayoutChanged queues a resize on the `GtkWidget` using @manager, if any.
	//
	// This function should be called by subclasses of `GtkLayoutManager` in
	// response to changes to their layout management policies.
	LayoutChanged()
}

// LayoutManagerClass implements the LayoutManager interface.
type LayoutManagerClass struct {
	*externglib.Object
}

var _ LayoutManager = (*LayoutManagerClass)(nil)

func wrapLayoutManager(obj *externglib.Object) LayoutManager {
	return &LayoutManagerClass{
		Object: obj,
	}
}

func marshalLayoutManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLayoutManager(obj), nil
}

// Allocate assigns the given @width, @height, and @baseline to a @widget, and
// computes the position and sizes of the children of the @widget using the
// layout management policy of @manager.
func (manager *LayoutManagerClass) Allocate(widget Widget, width int, height int, baseline int) {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _arg4 C.int               // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.int(width)
	_arg3 = C.int(height)
	_arg4 = C.int(baseline)

	C.gtk_layout_manager_allocate(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// LayoutChild retrieves a `GtkLayoutChild` instance for the `GtkLayoutManager`,
// creating one if necessary.
//
// The @child widget must be a child of the widget using @manager.
//
// The `GtkLayoutChild` instance is owned by the `GtkLayoutManager`, and is
// guaranteed to exist as long as @child is a child of the `GtkWidget` using the
// given `GtkLayoutManager`.
func (manager *LayoutManagerClass) LayoutChild(child Widget) *LayoutChildClass {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.GtkLayoutChild   // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_layout_manager_get_layout_child(_arg0, _arg1)

	var _layoutChild *LayoutChildClass // out

	_layoutChild = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LayoutChildClass)

	return _layoutChild
}

// RequestMode retrieves the request mode of @manager.
func (manager *LayoutManagerClass) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkLayoutManager  // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_layout_manager_get_request_mode(_arg0)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = (SizeRequestMode)(_cret)

	return _sizeRequestMode
}

// Widget retrieves the `GtkWidget` using the given `GtkLayoutManager`.
func (manager *LayoutManagerClass) Widget() *WidgetClass {
	var _arg0 *C.GtkLayoutManager // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_layout_manager_get_widget(_arg0)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// LayoutChanged queues a resize on the `GtkWidget` using @manager, if any.
//
// This function should be called by subclasses of `GtkLayoutManager` in
// response to changes to their layout management policies.
func (manager *LayoutManagerClass) LayoutChanged() {
	var _arg0 *C.GtkLayoutManager // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(manager.Native()))

	C.gtk_layout_manager_layout_changed(_arg0)
}
