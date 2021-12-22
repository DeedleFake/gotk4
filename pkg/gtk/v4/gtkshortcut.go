// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_get_type()), F: marshalShortcutter},
	})
}

// Shortcut: GtkShortcut describes a keyboard shortcut.
//
// It contains a description of how to trigger the shortcut via a
// gtk.ShortcutTrigger and a way to activate the shortcut on a widget via a
// gtk.ShortcutAction.
//
// The actual work is usually done via gtk.ShortcutController, which decides if
// and when to activate a shortcut. Using that controller directly however is
// rarely necessary as various higher level convenience APIs exist on Widgets
// that make it easier to use shortcuts in GTK.
//
// GtkShortcut does provide functionality to make it easy for users to work with
// shortcuts, either by providing informational strings for display purposes or
// by allowing shortcuts to be configured.
type Shortcut struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Shortcut)(nil)
)

func wrapShortcut(obj *externglib.Object) *Shortcut {
	return &Shortcut{
		Object: obj,
	}
}

func marshalShortcutter(p uintptr) (interface{}, error) {
	return wrapShortcut(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewShortcut creates a new GtkShortcut that is triggered by trigger and then
// activates action.
//
// The function takes the following parameters:
//
//    - trigger (optional) that will trigger the shortcut.
//    - action (optional) that will be activated upon triggering.
//
// The function returns the following values:
//
//    - shortcut: new GtkShortcut.
//
func NewShortcut(trigger ShortcutTriggerer, action ShortcutActioner) *Shortcut {
	var _arg1 *C.GtkShortcutTrigger // out
	var _arg2 *C.GtkShortcutAction  // out
	var _cret *C.GtkShortcut        // in

	if trigger != nil {
		_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer(trigger.Native()))
		C.g_object_ref(C.gpointer(trigger.Native()))
	}
	if action != nil {
		_arg2 = (*C.GtkShortcutAction)(unsafe.Pointer(action.Native()))
		C.g_object_ref(C.gpointer(action.Native()))
	}

	_cret = C.gtk_shortcut_new(_arg1, _arg2)
	runtime.KeepAlive(trigger)
	runtime.KeepAlive(action)

	var _shortcut *Shortcut // out

	_shortcut = wrapShortcut(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _shortcut
}

// Action gets the action that is activated by this shortcut.
//
// The function returns the following values:
//
//    - shortcutAction (optional): action.
//
func (self *Shortcut) Action() ShortcutActioner {
	var _arg0 *C.GtkShortcut       // out
	var _cret *C.GtkShortcutAction // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_get_action(_arg0)
	runtime.KeepAlive(self)

	var _shortcutAction ShortcutActioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(ShortcutActioner)
				return ok
			})
			rv, ok := casted.(ShortcutActioner)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.ShortcutActioner")
			}
			_shortcutAction = rv
		}
	}

	return _shortcutAction
}

// Arguments gets the arguments that are passed when activating the shortcut.
//
// The function returns the following values:
//
//    - variant (optional): arguments.
//
func (self *Shortcut) Arguments() *glib.Variant {
	var _arg0 *C.GtkShortcut // out
	var _cret *C.GVariant    // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_get_arguments(_arg0)
	runtime.KeepAlive(self)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_variant_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// Trigger gets the trigger used to trigger self.
//
// The function returns the following values:
//
//    - shortcutTrigger (optional): trigger used.
//
func (self *Shortcut) Trigger() ShortcutTriggerer {
	var _arg0 *C.GtkShortcut        // out
	var _cret *C.GtkShortcutTrigger // in

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_shortcut_get_trigger(_arg0)
	runtime.KeepAlive(self)

	var _shortcutTrigger ShortcutTriggerer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(ShortcutTriggerer)
				return ok
			})
			rv, ok := casted.(ShortcutTriggerer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.ShortcutTriggerer")
			}
			_shortcutTrigger = rv
		}
	}

	return _shortcutTrigger
}

// SetAction sets the new action for self to be action.
//
// The function takes the following parameters:
//
//    - action (optional): new action. If the action is NULL, the nothing action
//      will be used.
//
func (self *Shortcut) SetAction(action ShortcutActioner) {
	var _arg0 *C.GtkShortcut       // out
	var _arg1 *C.GtkShortcutAction // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))
	if action != nil {
		_arg1 = (*C.GtkShortcutAction)(unsafe.Pointer(action.Native()))
		C.g_object_ref(C.gpointer(action.Native()))
	}

	C.gtk_shortcut_set_action(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(action)
}

// SetArguments sets the arguments to pass when activating the shortcut.
//
// The function takes the following parameters:
//
//    - args (optional) arguments to pass when activating self.
//
func (self *Shortcut) SetArguments(args *glib.Variant) {
	var _arg0 *C.GtkShortcut // out
	var _arg1 *C.GVariant    // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))
	if args != nil {
		_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(args)))
	}

	C.gtk_shortcut_set_arguments(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(args)
}

// SetTrigger sets the new trigger for self to be trigger.
//
// The function takes the following parameters:
//
//    - trigger (optional): new trigger. If the trigger is NULL, the never
//      trigger will be used.
//
func (self *Shortcut) SetTrigger(trigger ShortcutTriggerer) {
	var _arg0 *C.GtkShortcut        // out
	var _arg1 *C.GtkShortcutTrigger // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(self.Native()))
	if trigger != nil {
		_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer(trigger.Native()))
		C.g_object_ref(C.gpointer(trigger.Native()))
	}

	C.gtk_shortcut_set_trigger(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(trigger)
}
