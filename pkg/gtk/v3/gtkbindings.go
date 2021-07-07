// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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

// BindingsActivate: find a key binding matching @keyval and @modifiers and
// activate the binding on @object.
func BindingsActivate(object gextras.Objector, keyval uint, modifiers gdk.ModifierType) bool {
	var _arg1 *C.GObject        // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)

	_cret = C.gtk_bindings_activate(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BindingsActivateEvent looks up key bindings for @object to find one matching
// @event, and if one was found, activate it.
func BindingsActivateEvent(object gextras.Objector, event *gdk.EventKey) bool {
	var _arg1 *C.GObject     // out
	var _arg2 *C.GdkEventKey // out
	var _cret C.gboolean     // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = (*C.GdkEventKey)(unsafe.Pointer(event))

	_cret = C.gtk_bindings_activate_event(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BindingArg holds the data associated with an argument for a key binding
// signal emission as stored in BindingSignal.
type BindingArg struct {
	native C.GtkBindingArg
}

// WrapBindingArg wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBindingArg(ptr unsafe.Pointer) *BindingArg {
	return (*BindingArg)(ptr)
}

// Native returns the underlying C source pointer.
func (b *BindingArg) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// ArgType: implementation detail
func (b *BindingArg) ArgType() externglib.Type {
	var v externglib.Type // out
	v = externglib.Type(b.native.arg_type)
	return v
}

// BindingEntry: each key binding element of a binding sets binding list is
// represented by a GtkBindingEntry.
type BindingEntry struct {
	native C.GtkBindingEntry
}

// WrapBindingEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBindingEntry(ptr unsafe.Pointer) *BindingEntry {
	return (*BindingEntry)(ptr)
}

// Native returns the underlying C source pointer.
func (b *BindingEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Keyval: key value to match
func (b *BindingEntry) Keyval() uint {
	var v uint // out
	v = uint(b.native.keyval)
	return v
}

// Modifiers: key modifiers to match
func (b *BindingEntry) Modifiers() gdk.ModifierType {
	var v gdk.ModifierType // out
	v = gdk.ModifierType(b.native.modifiers)
	return v
}

// SetNext: linked list of entries maintained by binding set
func (b *BindingEntry) SetNext() *BindingEntry {
	var v *BindingEntry // out
	v = (*BindingEntry)(unsafe.Pointer(b.native.set_next))
	return v
}

// HashNext: implementation detail
func (b *BindingEntry) HashNext() *BindingEntry {
	var v *BindingEntry // out
	v = (*BindingEntry)(unsafe.Pointer(b.native.hash_next))
	return v
}

// Signals: action signals of this entry
func (b *BindingEntry) Signals() *BindingSignal {
	var v *BindingSignal // out
	v = (*BindingSignal)(unsafe.Pointer(b.native.signals))
	return v
}

// BindingSignal stores the necessary information to activate a widget in
// response to a key press via a signal emission.
type BindingSignal struct {
	native C.GtkBindingSignal
}

// WrapBindingSignal wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBindingSignal(ptr unsafe.Pointer) *BindingSignal {
	return (*BindingSignal)(ptr)
}

// Native returns the underlying C source pointer.
func (b *BindingSignal) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Next: implementation detail
func (b *BindingSignal) Next() *BindingSignal {
	var v *BindingSignal // out
	v = (*BindingSignal)(unsafe.Pointer(b.native.next))
	return v
}

// SignalName: the action signal to be emitted
func (b *BindingSignal) SignalName() string {
	var v string // out
	v = C.GoString(b.native.signal_name)
	return v
}

// NArgs: number of arguments specified for the signal
func (b *BindingSignal) NArgs() uint {
	var v uint // out
	v = uint(b.native.n_args)
	return v
}
