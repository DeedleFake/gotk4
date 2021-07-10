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
		{T: externglib.Type(C.g_memory_input_stream_get_type()), F: marshalMemoryInputStreamer},
	})
}

// MemoryInputStreamer describes MemoryInputStream's methods.
type MemoryInputStreamer interface {
	gextras.Objector

	privateMemoryInputStream()
}

// MemoryInputStream is a class for using arbitrary memory chunks as input for
// GIO streaming input operations.
//
// As of GLib 2.34, InputStream implements InputStream.
type MemoryInputStream struct {
	*externglib.Object

	InputStream
	PollableInputStream
	Seekable
}

var _ MemoryInputStreamer = (*MemoryInputStream)(nil)

func wrapMemoryInputStreamer(obj *externglib.Object) MemoryInputStreamer {
	return &MemoryInputStream{
		Object: obj,
		InputStream: InputStream{
			Object: obj,
		},
		PollableInputStream: PollableInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalMemoryInputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMemoryInputStreamer(obj), nil
}

// NewMemoryInputStream creates a new empty InputStream.
func NewMemoryInputStream() *MemoryInputStream {
	var _cret *C.GInputStream // in

	_cret = C.g_memory_input_stream_new()

	var _memoryInputStream *MemoryInputStream // out

	_memoryInputStream = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*MemoryInputStream)

	return _memoryInputStream
}

func (*MemoryInputStream) privateMemoryInputStream() {}
