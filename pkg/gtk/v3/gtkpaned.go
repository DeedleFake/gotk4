// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_paned_get_type()), F: marshalPanedder},
	})
}

// PanedOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PanedOverrider interface {
	AcceptPosition() bool
	CancelPosition() bool
	CycleChildFocus(reverse bool) bool
	CycleHandleFocus(reverse bool) bool
	MoveHandle(scroll ScrollType) bool
	ToggleHandleFocus() bool
}

// Paned has two panes, arranged either horizontally or vertically. The division
// between the two panes is adjustable by the user by dragging a handle.
//
// Child widgets are added to the panes of the widget with gtk_paned_pack1() and
// gtk_paned_pack2(). The division between the two children is set by default
// from the size requests of the children, but it can be adjusted by the user.
//
// A paned widget draws a separator between the two child widgets and a small
// handle that the user can drag to adjust the division. It does not draw any
// relief around the children or around the separator. (The space in which the
// separator is called the gutter.) Often, it is useful to put each child inside
// a Frame with the shadow type set to GTK_SHADOW_IN so that the gutter appears
// as a ridge. No separator is drawn if one of the children is missing.
//
// Each child has two options that can be set, resize and shrink. If resize is
// true, then when the Paned is resized, that child will expand or shrink along
// with the paned widget. If shrink is true, then that child can be made smaller
// than its requisition by the user. Setting shrink to FALSE allows the
// application to set a minimum size. If resize is false for both children, then
// this is treated as if resize is true for both children.
//
// The application can set the position of the slider as if it were set by the
// user, by calling gtk_paned_set_position().
//
// CSS nodes
//
//    GtkWidget *hpaned = gtk_paned_new (GTK_ORIENTATION_HORIZONTAL);
//    GtkWidget *frame1 = gtk_frame_new (NULL);
//    GtkWidget *frame2 = gtk_frame_new (NULL);
//    gtk_frame_set_shadow_type (GTK_FRAME (frame1), GTK_SHADOW_IN);
//    gtk_frame_set_shadow_type (GTK_FRAME (frame2), GTK_SHADOW_IN);
//
//    gtk_widget_set_size_request (hpaned, 200, -1);
//
//    gtk_paned_pack1 (GTK_PANED (hpaned), frame1, TRUE, FALSE);
//    gtk_widget_set_size_request (frame1, 50, -1);
//
//    gtk_paned_pack2 (GTK_PANED (hpaned), frame2, FALSE, FALSE);
//    gtk_widget_set_size_request (frame2, 50, -1);
type Paned struct {
	Container

	Orientable
	*externglib.Object
}

func wrapPaned(obj *externglib.Object) *Paned {
	return &Paned{
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

func marshalPanedder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPaned(obj), nil
}

// NewPaned creates a new Paned widget.
func NewPaned(orientation Orientation) *Paned {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_paned_new(_arg1)
	runtime.KeepAlive(orientation)

	var _paned *Paned // out

	_paned = wrapPaned(externglib.Take(unsafe.Pointer(_cret)))

	return _paned
}

// Add1 adds a child to the top or left pane with default parameters. This is
// equivalent to gtk_paned_pack1 (paned, child, FALSE, TRUE).
func (paned *Paned) Add1(child Widgetter) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_add1(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
}

// Add2 adds a child to the bottom or right pane with default parameters. This
// is equivalent to gtk_paned_pack2 (paned, child, TRUE, TRUE).
func (paned *Paned) Add2(child Widgetter) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_add2(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
}

// Child1 obtains the first child of the paned widget.
func (paned *Paned) Child1() Widgetter {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_child1(_arg0)
	runtime.KeepAlive(paned)

	var _widget Widgetter // out

	if _cret != nil {
		{
			object := externglib.Take(unsafe.Pointer(_cret))
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Child2 obtains the second child of the paned widget.
func (paned *Paned) Child2() Widgetter {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_child2(_arg0)
	runtime.KeepAlive(paned)

	var _widget Widgetter // out

	if _cret != nil {
		{
			object := externglib.Take(unsafe.Pointer(_cret))
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// HandleWindow returns the Window of the handle. This function is useful when
// handling button or motion events because it enables the callback to
// distinguish between the window of the paned, a child and the handle.
func (paned *Paned) HandleWindow() gdk.Windower {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_handle_window(_arg0)
	runtime.KeepAlive(paned)

	var _window gdk.Windower // out

	{
		object := externglib.Take(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(gdk.Windower)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// Position obtains the position of the divider between the two panes.
func (paned *Paned) Position() int {
	var _arg0 *C.GtkPaned // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_position(_arg0)
	runtime.KeepAlive(paned)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WideHandle gets the Paned:wide-handle property.
func (paned *Paned) WideHandle() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_wide_handle(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Pack1 adds a child to the top or left pane.
func (paned *Paned) Pack1(child Widgetter, resize bool, shrink bool) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gboolean   // out
	var _arg3 C.gboolean   // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if resize {
		_arg2 = C.TRUE
	}
	if shrink {
		_arg3 = C.TRUE
	}

	C.gtk_paned_pack1(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
	runtime.KeepAlive(resize)
	runtime.KeepAlive(shrink)
}

// Pack2 adds a child to the bottom or right pane.
func (paned *Paned) Pack2(child Widgetter, resize bool, shrink bool) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gboolean   // out
	var _arg3 C.gboolean   // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if resize {
		_arg2 = C.TRUE
	}
	if shrink {
		_arg3 = C.TRUE
	}

	C.gtk_paned_pack2(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
	runtime.KeepAlive(resize)
	runtime.KeepAlive(shrink)
}

// SetPosition sets the position of the divider between the two panes.
func (paned *Paned) SetPosition(position int) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	_arg1 = C.gint(position)

	C.gtk_paned_set_position(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(position)
}

// SetWideHandle sets the Paned:wide-handle property.
func (paned *Paned) SetWideHandle(wide bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	if wide {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_wide_handle(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(wide)
}
