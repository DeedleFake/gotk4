// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_list_model_get_type()), F: marshalListModel},
	})
}

// ListModel is an interface that represents a mutable list of #GObjects. Its
// main intention is as a model for various widgets in user interfaces, such as
// list views, but it can also be used as a convenient method of returning lists
// of data, with support for updates.
//
// Each object in the list may also report changes in itself via some mechanism
// (normally the #GObject::notify signal). Taken together with the
// Model::items-changed signal, this provides for a list that can change its
// membership, and in which the members can change their individual properties.
//
// A good example would be the list of visible wireless network access points,
// where each access point can report dynamic properties such as signal
// strength.
//
// It is important to note that the Model itself does not report changes to the
// individual items. It only reports changes to the list membership. If you want
// to observe changes to the objects themselves then you need to connect signals
// to the objects that you are interested in.
//
// All items in a Model are of (or derived from) the same type.
// g_list_model_get_item_type() returns that type. The type may be an interface,
// in which case all objects in the list must implement it.
//
// The semantics are close to that of an array: g_list_model_get_n_items()
// returns the number of items in the list and g_list_model_get_item() returns
// an item at a (0-based) position. In order to allow implementations to
// calculate the list length lazily, you can also iterate over items: starting
// from 0, repeatedly call g_list_model_get_item() until it returns nil.
//
// An implementation may create objects lazily, but must take care to return the
// same object for a given position until all references to it are gone.
//
// On the other side, a consumer is expected only to hold references on objects
// that are currently "user visible", in order to facilitate the maximum level
// of laziness in the implementation of the list and to reduce the required
// number of signal connections at a given time.
//
// This interface is intended only to be used from a single thread. The thread
// in which it is appropriate to use it depends on the particular
// implementation, but typically it will be from the thread that owns the
// [thread-default main context][g-main-context-push-thread-default] in effect
// at the time that the model was created.
type ListModel interface {
	gextras.Objector

	// ItemType emits the Model::items-changed signal on @list.
	//
	// This function should only be called by classes implementing Model. It has
	// to be called after the internal representation of @list has been updated,
	// because handlers connected to this signal might query the new state of
	// the list.
	//
	// Implementations must only make changes to the model (as visible to its
	// consumer) in places that will not cause problems for that consumer. For
	// models that are driven directly by a write API (such as Store), changes
	// can be reported in response to uses of that API. For models that
	// represent remote data, changes should only be made from a fresh mainloop
	// dispatch. It is particularly not permitted to make changes in response to
	// a call to the Model consumer API.
	//
	// Stated another way: in general, it is assumed that code making a series
	// of accesses to the model via the API, without returning to the mainloop,
	// and without calling other code, will continue to view the same contents
	// of the model.
	ItemType() externglib.Type
	// NItems emits the Model::items-changed signal on @list.
	//
	// This function should only be called by classes implementing Model. It has
	// to be called after the internal representation of @list has been updated,
	// because handlers connected to this signal might query the new state of
	// the list.
	//
	// Implementations must only make changes to the model (as visible to its
	// consumer) in places that will not cause problems for that consumer. For
	// models that are driven directly by a write API (such as Store), changes
	// can be reported in response to uses of that API. For models that
	// represent remote data, changes should only be made from a fresh mainloop
	// dispatch. It is particularly not permitted to make changes in response to
	// a call to the Model consumer API.
	//
	// Stated another way: in general, it is assumed that code making a series
	// of accesses to the model via the API, without returning to the mainloop,
	// and without calling other code, will continue to view the same contents
	// of the model.
	NItems() uint
	// Object emits the Model::items-changed signal on @list.
	//
	// This function should only be called by classes implementing Model. It has
	// to be called after the internal representation of @list has been updated,
	// because handlers connected to this signal might query the new state of
	// the list.
	//
	// Implementations must only make changes to the model (as visible to its
	// consumer) in places that will not cause problems for that consumer. For
	// models that are driven directly by a write API (such as Store), changes
	// can be reported in response to uses of that API. For models that
	// represent remote data, changes should only be made from a fresh mainloop
	// dispatch. It is particularly not permitted to make changes in response to
	// a call to the Model consumer API.
	//
	// Stated another way: in general, it is assumed that code making a series
	// of accesses to the model via the API, without returning to the mainloop,
	// and without calling other code, will continue to view the same contents
	// of the model.
	Object(position uint) gextras.Objector
	// ItemsChanged emits the Model::items-changed signal on @list.
	//
	// This function should only be called by classes implementing Model. It has
	// to be called after the internal representation of @list has been updated,
	// because handlers connected to this signal might query the new state of
	// the list.
	//
	// Implementations must only make changes to the model (as visible to its
	// consumer) in places that will not cause problems for that consumer. For
	// models that are driven directly by a write API (such as Store), changes
	// can be reported in response to uses of that API. For models that
	// represent remote data, changes should only be made from a fresh mainloop
	// dispatch. It is particularly not permitted to make changes in response to
	// a call to the Model consumer API.
	//
	// Stated another way: in general, it is assumed that code making a series
	// of accesses to the model via the API, without returning to the mainloop,
	// and without calling other code, will continue to view the same contents
	// of the model.
	ItemsChanged(position uint, removed uint, added uint)
}

// listModel implements the ListModel interface.
type listModel struct {
	gextras.Objector
}

var _ ListModel = (*listModel)(nil)

// WrapListModel wraps a GObject to a type that implements
// interface ListModel. It is primarily used internally.
func WrapListModel(obj *externglib.Object) ListModel {
	return listModel{
		Objector: obj,
	}
}

func marshalListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListModel(obj), nil
}

func (l listModel) ItemType() externglib.Type {
	var _arg0 *C.GListModel // out
	var _cret C.GType       // in

	_arg0 = (*C.GListModel)(unsafe.Pointer(l.Native()))

	_cret = C.g_list_model_get_item_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

func (l listModel) NItems() uint {
	var _arg0 *C.GListModel // out
	var _cret C.guint       // in

	_arg0 = (*C.GListModel)(unsafe.Pointer(l.Native()))

	_cret = C.g_list_model_get_n_items(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (l listModel) Object(position uint) gextras.Objector {
	var _arg0 *C.GListModel // out
	var _arg1 C.guint       // out
	var _cret *C.GObject    // in

	_arg0 = (*C.GListModel)(unsafe.Pointer(l.Native()))
	_arg1 = C.guint(position)

	_cret = C.g_list_model_get_object(_arg0, _arg1)

	var _object gextras.Objector // out

	_object = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gextras.Objector)

	return _object
}

func (l listModel) ItemsChanged(position uint, removed uint, added uint) {
	var _arg0 *C.GListModel // out
	var _arg1 C.guint       // out
	var _arg2 C.guint       // out
	var _arg3 C.guint       // out

	_arg0 = (*C.GListModel)(unsafe.Pointer(l.Native()))
	_arg1 = C.guint(position)
	_arg2 = C.guint(removed)
	_arg3 = C.guint(added)

	C.g_list_model_items_changed(_arg0, _arg1, _arg2, _arg3)
}
