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
		{T: externglib.Type(C.g_memory_input_stream_get_type()), F: marshalMemoryInputStream},
	})
}

// MemoryInputStream is a class for using arbitrary memory chunks as input for
// GIO streaming input operations.
//
// As of GLib 2.34, InputStream implements InputStream.
type MemoryInputStream interface {
	gextras.Objector

	privateMemoryInputStreamClass()
}

// MemoryInputStreamClass implements the MemoryInputStream interface.
type MemoryInputStreamClass struct {
	*externglib.Object
	InputStreamClass
	PollableInputStreamIface
	SeekableIface
}

var _ MemoryInputStream = (*MemoryInputStreamClass)(nil)

func wrapMemoryInputStream(obj *externglib.Object) MemoryInputStream {
	return &MemoryInputStreamClass{
		Object: obj,
		InputStreamClass: InputStreamClass{
			Object: obj,
		},
		PollableInputStreamIface: PollableInputStreamIface{
			InputStreamClass: InputStreamClass{
				Object: obj,
			},
		},
		SeekableIface: SeekableIface{
			Object: obj,
		},
	}
}

func marshalMemoryInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMemoryInputStream(obj), nil
}

// NewMemoryInputStream creates a new empty InputStream.
func NewMemoryInputStream() *MemoryInputStreamClass {
	var _cret *C.GInputStream // in

	_cret = C.g_memory_input_stream_new()

	var _memoryInputStream *MemoryInputStreamClass // out

	_memoryInputStream = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*MemoryInputStreamClass)

	return _memoryInputStream
}

func (*MemoryInputStreamClass) privateMemoryInputStreamClass() {}
