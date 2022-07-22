// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_SwitchClass_state_set(GtkSwitch*, gboolean);
// extern gboolean _gotk4_gtk3_Switch_ConnectStateSet(gpointer, gboolean, guintptr);
// extern void _gotk4_gtk3_SwitchClass_activate(GtkSwitch*);
// extern void _gotk4_gtk3_Switch_ConnectActivate(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeSwitch = coreglib.Type(C.gtk_switch_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSwitch, F: marshalSwitch},
	})
}

// SwitchOverrider contains methods that are overridable.
type SwitchOverrider interface {
	Activate()
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	StateSet(state bool) bool
}

// Switch is a widget that has two states: on or off. The user can control which
// state should be active by clicking the empty area, or by dragging the handle.
//
// GtkSwitch can also handle situations where the underlying state changes with
// a delay. See Switch::state-set for details.
//
// CSS nodes
//
//    switch
//    ╰── slider
//
// GtkSwitch has two css nodes, the main node with the name switch and a subnode
// named slider. Neither of them is using any style classes.
type Switch struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Actionable
	Activatable
}

var (
	_ Widgetter         = (*Switch)(nil)
	_ coreglib.Objector = (*Switch)(nil)
)

func classInitSwitcher(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkSwitchClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ Activate() }); ok {
		pclass.activate = (*[0]byte)(C._gotk4_gtk3_SwitchClass_activate)
	}

	if _, ok := goval.(interface{ StateSet(state bool) bool }); ok {
		pclass.state_set = (*[0]byte)(C._gotk4_gtk3_SwitchClass_state_set)
	}
}

//export _gotk4_gtk3_SwitchClass_activate
func _gotk4_gtk3_SwitchClass_activate(arg0 *C.GtkSwitch) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Activate() })

	iface.Activate()
}

//export _gotk4_gtk3_SwitchClass_state_set
func _gotk4_gtk3_SwitchClass_state_set(arg0 *C.GtkSwitch, arg1 C.gboolean) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ StateSet(state bool) bool })

	var _state bool // out

	if arg1 != 0 {
		_state = true
	}

	ok := iface.StateSet(_state)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapSwitch(obj *coreglib.Object) *Switch {
	return &Switch{
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
		Object: obj,
		Actionable: Actionable{
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
		Activatable: Activatable{
			Object: obj,
		},
	}
}

func marshalSwitch(p uintptr) (interface{}, error) {
	return wrapSwitch(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Switch_ConnectActivate
func _gotk4_gtk3_Switch_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectActivate signal on GtkSwitch is an action signal and emitting it
// causes the switch to animate. Applications should never connect to this
// signal, but use the notify::active signal.
func (sw *Switch) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sw, "activate", false, unsafe.Pointer(C._gotk4_gtk3_Switch_ConnectActivate), f)
}

//export _gotk4_gtk3_Switch_ConnectStateSet
func _gotk4_gtk3_Switch_ConnectStateSet(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) (cret C.gboolean) {
	var f func(state bool) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(state bool) (ok bool))
	}

	var _state bool // out

	if arg1 != 0 {
		_state = true
	}

	ok := f(_state)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectStateSet signal on GtkSwitch is emitted to change the underlying
// state. It is emitted when the user changes the switch position. The default
// handler keeps the state in sync with the Switch:active property.
//
// To implement delayed state change, applications can connect to this signal,
// initiate the change of the underlying state, and call gtk_switch_set_state()
// when the underlying state change is complete. The signal handler should
// return TRUE to prevent the default handler from running.
//
// Visually, the underlying state is represented by the trough color of the
// switch, while the Switch:active property is represented by the position of
// the switch.
func (sw *Switch) ConnectStateSet(f func(state bool) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sw, "state-set", false, unsafe.Pointer(C._gotk4_gtk3_Switch_ConnectStateSet), f)
}

// NewSwitch creates a new Switch widget.
//
// The function returns the following values:
//
//    - _switch: newly created Switch instance.
//
func NewSwitch() *Switch {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_switch_new()

	var __switch *Switch // out

	__switch = wrapSwitch(coreglib.Take(unsafe.Pointer(_cret)))

	return __switch
}

// Active gets whether the Switch is in its “on” or “off” state.
//
// The function returns the following values:
//
//    - ok: TRUE if the Switch is active, and FALSE otherwise.
//
func (sw *Switch) Active() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(coreglib.InternObject(sw).Native()))

	_cret = C.gtk_switch_get_active(_arg0)
	runtime.KeepAlive(sw)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// State gets the underlying state of the Switch.
//
// The function returns the following values:
//
//    - ok: underlying state.
//
func (sw *Switch) State() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(coreglib.InternObject(sw).Native()))

	_cret = C.gtk_switch_get_state(_arg0)
	runtime.KeepAlive(sw)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive changes the state of sw to the desired one.
//
// The function takes the following parameters:
//
//    - isActive: TRUE if sw should be active, and FALSE otherwise.
//
func (sw *Switch) SetActive(isActive bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(coreglib.InternObject(sw).Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_active(_arg0, _arg1)
	runtime.KeepAlive(sw)
	runtime.KeepAlive(isActive)
}

// SetState sets the underlying state of the Switch.
//
// Normally, this is the same as Switch:active, unless the switch is set up for
// delayed state changes. This function is typically called from a
// Switch::state-set signal handler.
//
// See Switch::state-set for details.
//
// The function takes the following parameters:
//
//    - state: new state.
//
func (sw *Switch) SetState(state bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(coreglib.InternObject(sw).Native()))
	if state {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_state(_arg0, _arg1)
	runtime.KeepAlive(sw)
	runtime.KeepAlive(state)
}
