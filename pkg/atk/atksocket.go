// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_socket_get_type()), F: marshalSocketter},
	})
}

// SocketOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketOverrider interface {
	// Embed embeds the children of an Plug as the children of the Socket. The
	// plug may be in the same process or in a different process.
	//
	// The class item used by this function should be filled in by the IPC layer
	// (usually at-spi2-atk). The implementor of the AtkSocket should call this
	// function and pass the id for the plug as returned by atk_plug_get_id().
	// It is the responsibility of the application to pass the plug id on to the
	// process implementing the Socket as needed.
	Embed(plugId string)
}

// Socket: together with Plug, Socket provides the ability to embed accessibles
// from one process into another in a fashion that is transparent to assistive
// technologies. Socket works as the container of Plug, embedding it using the
// method atk_socket_embed(). Any accessible contained in the Plug will appear
// to the assistive technologies as being inside the application that created
// the Socket.
//
// The communication between a Socket and a Plug is done by the IPC layer of the
// accessibility framework, normally implemented by the D-Bus based
// implementation of AT-SPI (at-spi2). If that is the case, at-spi-atk2 is the
// responsible to implement the abstract methods atk_plug_get_id() and
// atk_socket_embed(), so an ATK implementor shouldn't reimplement them. The
// process that contains the Plug is responsible to send the ID returned by
// atk_plug_id() to the process that contains the Socket, so it could call the
// method atk_socket_embed() in order to embed it.
//
// For the same reasons, an implementor doesn't need to implement
// atk_object_get_n_accessible_children() and atk_object_ref_accessible_child().
// All the logic related to those functions will be implemented by the IPC
// layer.
type Socket struct {
	ObjectClass

	Component
	*externglib.Object
}

func wrapSocket(obj *externglib.Object) *Socket {
	return &Socket{
		ObjectClass: ObjectClass{
			Object: obj,
		},
		Component: Component{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalSocketter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocket(obj), nil
}

// NewSocket creates a new Socket.
func NewSocket() *Socket {
	var _cret *C.AtkObject // in

	_cret = C.atk_socket_new()

	var _socket *Socket // out

	_socket = wrapSocket(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socket
}

// Embed embeds the children of an Plug as the children of the Socket. The plug
// may be in the same process or in a different process.
//
// The class item used by this function should be filled in by the IPC layer
// (usually at-spi2-atk). The implementor of the AtkSocket should call this
// function and pass the id for the plug as returned by atk_plug_get_id(). It is
// the responsibility of the application to pass the plug id on to the process
// implementing the Socket as needed.
func (obj *Socket) Embed(plugId string) {
	var _arg0 *C.AtkSocket // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.AtkSocket)(unsafe.Pointer(obj.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(plugId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.atk_socket_embed(_arg0, _arg1)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(plugId)
}

// IsOccupied determines whether or not the socket has an embedded plug.
func (obj *Socket) IsOccupied() bool {
	var _arg0 *C.AtkSocket // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AtkSocket)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_socket_is_occupied(_arg0)
	runtime.KeepAlive(obj)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
