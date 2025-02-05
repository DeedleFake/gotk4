// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_Paned_ConnectToggleHandleFocus(gpointer, guintptr);
// extern gboolean _gotk4_gtk3_Paned_ConnectMoveHandle(gpointer, GtkScrollType, guintptr);
// extern gboolean _gotk4_gtk3_Paned_ConnectCycleHandleFocus(gpointer, gboolean, guintptr);
// extern gboolean _gotk4_gtk3_Paned_ConnectCycleChildFocus(gpointer, gboolean, guintptr);
// extern gboolean _gotk4_gtk3_Paned_ConnectCancelPosition(gpointer, guintptr);
// extern gboolean _gotk4_gtk3_Paned_ConnectAcceptPosition(gpointer, guintptr);
// extern gboolean _gotk4_gtk3_PanedClass_toggle_handle_focus(GtkPaned*);
// extern gboolean _gotk4_gtk3_PanedClass_move_handle(GtkPaned*, GtkScrollType);
// extern gboolean _gotk4_gtk3_PanedClass_cycle_handle_focus(GtkPaned*, gboolean);
// extern gboolean _gotk4_gtk3_PanedClass_cycle_child_focus(GtkPaned*, gboolean);
// extern gboolean _gotk4_gtk3_PanedClass_cancel_position(GtkPaned*);
// extern gboolean _gotk4_gtk3_PanedClass_accept_position(GtkPaned*);
// gboolean _gotk4_gtk3_Paned_virtual_accept_position(void* fnptr, GtkPaned* arg0) {
//   return ((gboolean (*)(GtkPaned*))(fnptr))(arg0);
// };
// gboolean _gotk4_gtk3_Paned_virtual_cancel_position(void* fnptr, GtkPaned* arg0) {
//   return ((gboolean (*)(GtkPaned*))(fnptr))(arg0);
// };
// gboolean _gotk4_gtk3_Paned_virtual_cycle_child_focus(void* fnptr, GtkPaned* arg0, gboolean arg1) {
//   return ((gboolean (*)(GtkPaned*, gboolean))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gtk3_Paned_virtual_cycle_handle_focus(void* fnptr, GtkPaned* arg0, gboolean arg1) {
//   return ((gboolean (*)(GtkPaned*, gboolean))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gtk3_Paned_virtual_move_handle(void* fnptr, GtkPaned* arg0, GtkScrollType arg1) {
//   return ((gboolean (*)(GtkPaned*, GtkScrollType))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gtk3_Paned_virtual_toggle_handle_focus(void* fnptr, GtkPaned* arg0) {
//   return ((gboolean (*)(GtkPaned*))(fnptr))(arg0);
// };
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

// PanedOverrides contains methods that are overridable.
type PanedOverrides struct {
	// The function returns the following values:
	//
	AcceptPosition func() bool
	// The function returns the following values:
	//
	CancelPosition func() bool
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	CycleChildFocus func(reverse bool) bool
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	CycleHandleFocus func(reverse bool) bool
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	MoveHandle func(scroll ScrollType) bool
	// The function returns the following values:
	//
	ToggleHandleFocus func() bool
}

func defaultPanedOverrides(v *Paned) PanedOverrides {
	return PanedOverrides{
		AcceptPosition:    v.acceptPosition,
		CancelPosition:    v.cancelPosition,
		CycleChildFocus:   v.cycleChildFocus,
		CycleHandleFocus:  v.cycleHandleFocus,
		MoveHandle:        v.moveHandle,
		ToggleHandleFocus: v.toggleHandleFocus,
	}
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
//    gtk_widget_set_size_request (frame2, 50, -1);.
type Paned struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	Orientable
}

var (
	_ Containerer       = (*Paned)(nil)
	_ coreglib.Objector = (*Paned)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Paned, *PanedClass, PanedOverrides](
		GTypePaned,
		initPanedClass,
		wrapPaned,
		defaultPanedOverrides,
	)
}

func initPanedClass(gclass unsafe.Pointer, overrides PanedOverrides, classInitFunc func(*PanedClass)) {
	pclass := (*C.GtkPanedClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypePaned))))

	if overrides.AcceptPosition != nil {
		pclass.accept_position = (*[0]byte)(C._gotk4_gtk3_PanedClass_accept_position)
	}

	if overrides.CancelPosition != nil {
		pclass.cancel_position = (*[0]byte)(C._gotk4_gtk3_PanedClass_cancel_position)
	}

	if overrides.CycleChildFocus != nil {
		pclass.cycle_child_focus = (*[0]byte)(C._gotk4_gtk3_PanedClass_cycle_child_focus)
	}

	if overrides.CycleHandleFocus != nil {
		pclass.cycle_handle_focus = (*[0]byte)(C._gotk4_gtk3_PanedClass_cycle_handle_focus)
	}

	if overrides.MoveHandle != nil {
		pclass.move_handle = (*[0]byte)(C._gotk4_gtk3_PanedClass_move_handle)
	}

	if overrides.ToggleHandleFocus != nil {
		pclass.toggle_handle_focus = (*[0]byte)(C._gotk4_gtk3_PanedClass_toggle_handle_focus)
	}

	if classInitFunc != nil {
		class := (*PanedClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPaned(obj *coreglib.Object) *Paned {
	return &Paned{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
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
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalPaned(p uintptr) (interface{}, error) {
	return wrapPaned(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAcceptPosition signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted to accept the current position of the handle when moving it
// using key bindings.
//
// The default binding for this signal is Return or Space.
func (paned *Paned) ConnectAcceptPosition(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "accept-position", false, unsafe.Pointer(C._gotk4_gtk3_Paned_ConnectAcceptPosition), f)
}

// ConnectCancelPosition signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted to cancel moving the position of the handle using key bindings.
// The position of the handle will be reset to the value prior to moving it.
//
// The default binding for this signal is Escape.
func (paned *Paned) ConnectCancelPosition(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "cancel-position", false, unsafe.Pointer(C._gotk4_gtk3_Paned_ConnectCancelPosition), f)
}

// ConnectCycleChildFocus signal is a [keybinding signal][GtkBindingSignal]
// which gets emitted to cycle the focus between the children of the paned.
//
// The default binding is f6.
func (paned *Paned) ConnectCycleChildFocus(f func(reversed bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "cycle-child-focus", false, unsafe.Pointer(C._gotk4_gtk3_Paned_ConnectCycleChildFocus), f)
}

// ConnectCycleHandleFocus signal is a [keybinding signal][GtkBindingSignal]
// which gets emitted to cycle whether the paned should grab focus to allow the
// user to change position of the handle by using key bindings.
//
// The default binding for this signal is f8.
func (paned *Paned) ConnectCycleHandleFocus(f func(reversed bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "cycle-handle-focus", false, unsafe.Pointer(C._gotk4_gtk3_Paned_ConnectCycleHandleFocus), f)
}

// ConnectMoveHandle signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted to move the handle when the user is using key bindings to move
// it.
func (paned *Paned) ConnectMoveHandle(f func(scrollType ScrollType) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "move-handle", false, unsafe.Pointer(C._gotk4_gtk3_Paned_ConnectMoveHandle), f)
}

// ConnectToggleHandleFocus is a [keybinding signal][GtkBindingSignal] which
// gets emitted to accept the current position of the handle and then move focus
// to the next widget in the focus chain.
//
// The default binding is Tab.
func (paned *Paned) ConnectToggleHandleFocus(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(paned, "toggle-handle-focus", false, unsafe.Pointer(C._gotk4_gtk3_Paned_ConnectToggleHandleFocus), f)
}

// NewPaned creates a new Paned widget.
//
// The function takes the following parameters:
//
//    - orientation paned’s orientation.
//
// The function returns the following values:
//
//    - paned: new Paned.
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

// Add1 adds a child to the top or left pane with default parameters. This is
// equivalent to gtk_paned_pack1 (paned, child, FALSE, TRUE).
//
// The function takes the following parameters:
//
//    - child to add.
//
func (paned *Paned) Add1(child Widgetter) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_paned_add1(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
}

// Add2 adds a child to the bottom or right pane with default parameters. This
// is equivalent to gtk_paned_pack2 (paned, child, TRUE, TRUE).
//
// The function takes the following parameters:
//
//    - child to add.
//
func (paned *Paned) Add2(child Widgetter) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_paned_add2(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
}

// Child1 obtains the first child of the paned widget.
//
// The function returns the following values:
//
//    - widget (optional): first child, or NULL if it is not set.
//
func (paned *Paned) Child1() Widgetter {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_child1(_arg0)
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

// Child2 obtains the second child of the paned widget.
//
// The function returns the following values:
//
//    - widget (optional): second child, or NULL if it is not set.
//
func (paned *Paned) Child2() Widgetter {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_child2(_arg0)
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

// HandleWindow returns the Window of the handle. This function is useful when
// handling button or motion events because it enables the callback to
// distinguish between the window of the paned, a child and the handle.
//
// The function returns the following values:
//
//    - window paned’s handle window.
//
func (paned *Paned) HandleWindow() gdk.Windower {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_handle_window(_arg0)
	runtime.KeepAlive(paned)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Windower)
			return ok
		})
		rv, ok := casted.(gdk.Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// Position obtains the position of the divider between the two panes.
//
// The function returns the following values:
//
//    - gint: position of the divider.
//
func (paned *Paned) Position() int {
	var _arg0 *C.GtkPaned // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C.gtk_paned_get_position(_arg0)
	runtime.KeepAlive(paned)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WideHandle gets the Paned:wide-handle property.
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

// Pack1 adds a child to the top or left pane.
//
// The function takes the following parameters:
//
//    - child to add.
//    - resize: should this child expand when the paned widget is resized.
//    - shrink: can this child be made smaller than its requisition.
//
func (paned *Paned) Pack1(child Widgetter, resize, shrink bool) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gboolean   // out
	var _arg3 C.gboolean   // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
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
//
// The function takes the following parameters:
//
//    - child to add.
//    - resize: should this child expand when the paned widget is resized.
//    - shrink: can this child be made smaller than its requisition.
//
func (paned *Paned) Pack2(child Widgetter, resize, shrink bool) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gboolean   // out
	var _arg3 C.gboolean   // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
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
//
// The function takes the following parameters:
//
//    - position: pixel position of divider, a negative value means that the
//      position is unset.
//
func (paned *Paned) SetPosition(position int) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = C.gint(position)

	C.gtk_paned_set_position(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(position)
}

// SetWideHandle sets the Paned:wide-handle property.
//
// The function takes the following parameters:
//
//    - wide: new value for the Paned:wide-handle property.
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

// The function returns the following values:
//
func (paned *Paned) acceptPosition() bool {
	gclass := (*C.GtkPanedClass)(coreglib.PeekParentClass(paned))
	fnarg := gclass.accept_position

	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C._gotk4_gtk3_Paned_virtual_accept_position(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
func (paned *Paned) cancelPosition() bool {
	gclass := (*C.GtkPanedClass)(coreglib.PeekParentClass(paned))
	fnarg := gclass.cancel_position

	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C._gotk4_gtk3_Paned_virtual_cancel_position(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (paned *Paned) cycleChildFocus(reverse bool) bool {
	gclass := (*C.GtkPanedClass)(coreglib.PeekParentClass(paned))
	fnarg := gclass.cycle_child_focus

	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	if reverse {
		_arg1 = C.TRUE
	}

	_cret = C._gotk4_gtk3_Paned_virtual_cycle_child_focus(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(reverse)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (paned *Paned) cycleHandleFocus(reverse bool) bool {
	gclass := (*C.GtkPanedClass)(coreglib.PeekParentClass(paned))
	fnarg := gclass.cycle_handle_focus

	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	if reverse {
		_arg1 = C.TRUE
	}

	_cret = C._gotk4_gtk3_Paned_virtual_cycle_handle_focus(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(reverse)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (paned *Paned) moveHandle(scroll ScrollType) bool {
	gclass := (*C.GtkPanedClass)(coreglib.PeekParentClass(paned))
	fnarg := gclass.move_handle

	var _arg0 *C.GtkPaned     // out
	var _arg1 C.GtkScrollType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))
	_arg1 = C.GtkScrollType(scroll)

	_cret = C._gotk4_gtk3_Paned_virtual_move_handle(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(scroll)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
func (paned *Paned) toggleHandleFocus() bool {
	gclass := (*C.GtkPanedClass)(coreglib.PeekParentClass(paned))
	fnarg := gclass.toggle_handle_focus

	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(coreglib.InternObject(paned).Native()))

	_cret = C._gotk4_gtk3_Paned_virtual_toggle_handle_focus(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PanedClass: instance of this type is always passed by reference.
type PanedClass struct {
	*panedClass
}

// panedClass is the struct that's finalized.
type panedClass struct {
	native *C.GtkPanedClass
}

func (p *PanedClass) ParentClass() *ContainerClass {
	valptr := &p.native.parent_class
	var _v *ContainerClass // out
	_v = (*ContainerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
