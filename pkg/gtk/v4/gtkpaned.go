// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_Paned_ConnectToggleHandleFocus(gpointer, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectMoveHandle(gpointer, GtkScrollType, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectCycleHandleFocus(gpointer, gboolean, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectCycleChildFocus(gpointer, gboolean, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectCancelPosition(gpointer, guintptr);
// extern gboolean _gotk4_gtk4_Paned_ConnectAcceptPosition(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypePaned = coreglib.Type(C.gtk_paned_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePaned, F: marshalPaned},
	})
}

// Paned: GtkPaned has two panes, arranged either horizontally or vertically.
//
// !An example GtkPaned (panes.png)
//
// The division between the two panes is adjustable by the user by dragging a
// handle.
//
// Child widgets are added to the panes of the widget with
// gtk.Paned.SetStartChild() and gtk.Paned.SetEndChild(). The division between
// the two children is set by default from the size requests of the children,
// but it can be adjusted by the user.
//
// A paned widget draws a separator between the two child widgets and a small
// handle that the user can drag to adjust the division. It does not draw any
// relief around the children or around the separator. (The space in which the
// separator is called the gutter.) Often, it is useful to put each child inside
// a gtk.Frame so that the gutter appears as a ridge. No separator is drawn if
// one of the children is missing.
//
// Each child has two options that can be set, resize and shrink. If resize is
// true, then when the GtkPaned is resized, that child will expand or shrink
// along with the paned widget. If shrink is true, then that child can be made
// smaller than its requisition by the user. Setting shrink to FALSE allows the
// application to set a minimum size. If resize is false for both children, then
// this is treated as if resize is true for both children.
//
// The application can set the position of the slider as if it were set by the
// user, by calling gtk.Paned.SetPosition().
//
// CSS nodes
//
//    paned
//    ├── <child>
//    ├── separator[.wide]
//    ╰── <child>
//
//
// GtkPaned has a main CSS node with name paned, and a subnode for the separator
// with name separator. The subnode gets a .wide style class when the paned is
// supposed to be wide.
//
// In horizontal orientation, the nodes are arranged based on the text
// direction, so in left-to-right mode, :first-child will select the leftmost
// child, while it will select the rightmost child in RTL layouts.
//
// Creating a paned widget with minimum sizes.
//
//    GtkWidget *hpaned = gtk_paned_new (GTK_ORIENTATION_HORIZONTAL);
//    GtkWidget *frame1 = gtk_frame_new (NULL);
//    GtkWidget *frame2 = gtk_frame_new (NULL);
//
//    gtk_widget_set_size_request (hpaned, 200, -1);
//
//    gtk_paned_set_start_child (GTK_PANED (hpaned), frame1);
//    gtk_paned_set_start_child_resize (GTK_PANED (hpaned), TRUE);
//    gtk_paned_set_start_child_shrink (GTK_PANED (hpaned), FALSE);
//    gtk_widget_set_size_request (frame1, 50, -1);
//
//    gtk_paned_set_end_child (GTK_PANED (hpaned), frame2);
//    gtk_paned_set_end_child_resize (GTK_PANED (hpaned), FALSE);
//    gtk_paned_set_end_child_shrink (GTK_PANED (hpaned), FALSE);
//    gtk_widget_set_size_request (frame2, 50, -1);.
type Paned struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*Paned)(nil)
	_ coreglib.Objector = (*Paned)(nil)
)

func wrapPaned(obj *coreglib.Object) *Paned {
	return &Paned{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalPaned(p uintptr) (interface{}, error) {
	return wrapPaned(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAcceptPosition is emitted to accept the current position of the handle
// when moving it using key bindings.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Return or Space.
func (paned *Paned) ConnectAcceptPosition(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "accept-position", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectAcceptPosition), f)
}

// ConnectCancelPosition is emitted to cancel moving the position of the handle
// using key bindings.
//
// The position of the handle will be reset to the value prior to moving it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Escape.
func (paned *Paned) ConnectCancelPosition(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "cancel-position", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectCancelPosition), f)
}

// ConnectCycleChildFocus is emitted to cycle the focus between the children of
// the paned.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding is F6.
func (paned *Paned) ConnectCycleChildFocus(f func(reversed bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "cycle-child-focus", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectCycleChildFocus), f)
}

// ConnectCycleHandleFocus is emitted to cycle whether the paned should grab
// focus to allow the user to change position of the handle by using key
// bindings.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is F8.
func (paned *Paned) ConnectCycleHandleFocus(f func(reversed bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "cycle-handle-focus", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectCycleHandleFocus), f)
}

// ConnectMoveHandle is emitted to move the handle with key bindings.
//
// This is a keybinding signal (class.SignalAction.html).
func (paned *Paned) ConnectMoveHandle(f func(scrollType ScrollType) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "move-handle", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectMoveHandle), f)
}

// ConnectToggleHandleFocus is emitted to accept the current position of the
// handle and then move focus to the next widget in the focus chain.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding is Tab.
func (paned *Paned) ConnectToggleHandleFocus(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "toggle-handle-focus", false, unsafe.Pointer(C._gotk4_gtk4_Paned_ConnectToggleHandleFocus), f)
}

// NewPaned creates a new GtkPaned widget.
//
// The function takes the following parameters:
//
//    - orientation paned’s orientation.
//
// The function returns the following values:
//
//    - paned: new GtkPaned.
//
func NewPaned(orientation Orientation) *Paned {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_paned_new(_arg1)
	runtime.KeepAlive(orientation)

	var _paned *Paned // out

	_paned = wrapPaned(coreglib.Take(unsafe.Pointer(_cret)))

	return _paned
}

// EndChild retrieves the end child of the given GtkPaned.
//
// See also: GtkPaned:end-child.
//
// The function returns the following values:
//
//    - widget (optional): end child widget.
//
func (paned *Paned) EndChild() Widgetter {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_end_child(_arg0)
	runtime.KeepAlive(paned)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Position obtains the position of the divider between the two panes.
//
// The function returns the following values:
//
//    - gint: position of the divider.
//
func (paned *Paned) Position() int {
	var _arg0 *C.GtkPaned // out
	var _cret C.int       // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_position(_arg0)
	runtime.KeepAlive(paned)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ResizeEndChild returns whether the end child can be resized.
//
// The function returns the following values:
//
//    - ok: TRUE if the end child is resizable.
//
func (paned *Paned) ResizeEndChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_resize_end_child(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResizeStartChild returns whether the start child can be resized.
//
// The function returns the following values:
//
//    - ok: TRUE if the start child is resizable.
//
func (paned *Paned) ResizeStartChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_resize_start_child(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShrinkEndChild returns whether the end child can be shrunk.
//
// The function returns the following values:
//
//    - ok: TRUE if the end child is shrinkable.
//
func (paned *Paned) ShrinkEndChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_shrink_end_child(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShrinkStartChild returns whether the start child can be shrunk.
//
// The function returns the following values:
//
//    - ok: TRUE if the start child is shrinkable.
//
func (paned *Paned) ShrinkStartChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_shrink_start_child(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartChild retrieves the start child of the given GtkPaned.
//
// See also: GtkPaned:start-child.
//
// The function returns the following values:
//
//    - widget (optional): start child widget.
//
func (paned *Paned) StartChild() Widgetter {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_start_child(_arg0)
	runtime.KeepAlive(paned)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// WideHandle gets whether the separator should be wide.
//
// The function returns the following values:
//
//    - ok: TRUE if the paned should have a wide handle.
//
func (paned *Paned) WideHandle() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_wide_handle(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetEndChild sets the end child of paned to child.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//
func (paned *Paned) SetEndChild(child Widgetter) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_paned_set_end_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
}

// SetPosition sets the position of the divider between the two panes.
//
// The function takes the following parameters:
//
//    - position: pixel position of divider, a negative value means that the
//      position is unset.
//
func (paned *Paned) SetPosition(position int) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = C.int(position)

	C.gtk_paned_set_position(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(position)
}

// SetResizeEndChild sets the GtkPaned:resize-end-child property.
//
// The function takes the following parameters:
//
//    - resize: TRUE to let the end child be resized.
//
func (paned *Paned) SetResizeEndChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_resize_end_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(resize)
}

// SetResizeStartChild sets the GtkPaned:resize-start-child property.
//
// The function takes the following parameters:
//
//    - resize: TRUE to let the start child be resized.
//
func (paned *Paned) SetResizeStartChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_resize_start_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(resize)
}

// SetShrinkEndChild sets the GtkPaned:shrink-end-child property.
//
// The function takes the following parameters:
//
//    - resize: TRUE to let the end child be shrunk.
//
func (paned *Paned) SetShrinkEndChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_shrink_end_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(resize)
}

// SetShrinkStartChild sets the GtkPaned:shrink-start-child property.
//
// The function takes the following parameters:
//
//    - resize: TRUE to let the start child be shrunk.
//
func (paned *Paned) SetShrinkStartChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_shrink_start_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(resize)
}

// SetStartChild sets the start child of paned to child.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//
func (paned *Paned) SetStartChild(child Widgetter) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_paned_set_start_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
}

// SetWideHandle sets whether the separator should be wide.
//
// The function takes the following parameters:
//
//    - wide: new value for the gtk.Paned:wide-handle property.
//
func (paned *Paned) SetWideHandle(wide bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	if wide {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_wide_handle(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(wide)
}
