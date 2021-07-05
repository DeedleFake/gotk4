// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_simple_action_group_get_type()), F: marshalSimpleActionGroup},
	})
}

// SimpleActionGroup is a hash table filled with #GAction objects, implementing
// the Group and Map interfaces.
type SimpleActionGroup interface {
	gextras.Objector

	// AsActionGroup casts the class to the ActionGroup interface.
	AsActionGroup() ActionGroup
	// AsActionMap casts the class to the ActionMap interface.
	AsActionMap() ActionMap

	// AddEntriesSimpleActionGroup: convenience function for creating multiple
	// Action instances and adding them to the action group.
	//
	// Deprecated: since version 2.38.
	AddEntriesSimpleActionGroup(entries []ActionEntry, userData interface{})
	// InsertSimpleActionGroup adds an action to the action group.
	//
	// If the action group already contains an action with the same name as
	// @action then the old action is dropped from the group.
	//
	// The action group takes its own reference on @action.
	//
	// Deprecated: since version 2.38.
	InsertSimpleActionGroup(action Action)
	// LookupSimpleActionGroup looks up the action with the name @action_name in
	// the group.
	//
	// If no such action exists, returns nil.
	//
	// Deprecated: since version 2.38.
	LookupSimpleActionGroup(actionName string) Action
	// RemoveSimpleActionGroup removes the named action from the action group.
	//
	// If no action of this name is in the group then nothing happens.
	//
	// Deprecated: since version 2.38.
	RemoveSimpleActionGroup(actionName string)
}

// simpleActionGroup implements the SimpleActionGroup class.
type simpleActionGroup struct {
	gextras.Objector
}

// WrapSimpleActionGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapSimpleActionGroup(obj *externglib.Object) SimpleActionGroup {
	return simpleActionGroup{
		Objector: obj,
	}
}

func marshalSimpleActionGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSimpleActionGroup(obj), nil
}

// NewSimpleActionGroup creates a new, empty, ActionGroup.
func NewSimpleActionGroup() SimpleActionGroup {
	var _cret *C.GSimpleActionGroup // in

	_cret = C.g_simple_action_group_new()

	var _simpleActionGroup SimpleActionGroup // out

	_simpleActionGroup = WrapSimpleActionGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleActionGroup
}

func (s simpleActionGroup) AddEntriesSimpleActionGroup(entries []ActionEntry, userData interface{}) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.GActionEntry
	var _arg2 C.gint
	var _arg3 C.gpointer // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg2 = C.gint(len(entries))
	_arg1 = (*C.GActionEntry)(unsafe.Pointer(&entries[0]))
	_arg3 = (C.gpointer)(box.Assign(userData))

	C.g_simple_action_group_add_entries(_arg0, _arg1, _arg2, _arg3)
}

func (s simpleActionGroup) InsertSimpleActionGroup(action Action) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.GAction            // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(action.Native()))

	C.g_simple_action_group_insert(_arg0, _arg1)
}

func (s simpleActionGroup) LookupSimpleActionGroup(actionName string) Action {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.gchar              // out
	var _cret *C.GAction            // in

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_simple_action_group_lookup(_arg0, _arg1)

	var _action Action // out

	_action = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Action)

	return _action
}

func (s simpleActionGroup) RemoveSimpleActionGroup(actionName string) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_simple_action_group_remove(_arg0, _arg1)
}

func (s simpleActionGroup) AsActionGroup() ActionGroup {
	return WrapActionGroup(gextras.InternObject(s))
}

func (s simpleActionGroup) AsActionMap() ActionMap {
	return WrapActionMap(gextras.InternObject(s))
}
