// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeFilterInputStream returns the GType for the type FilterInputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFilterInputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "FilterInputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFilterInputStream)
	return gtype
}

// FilterInputStreamOverrider contains methods that are overridable.
type FilterInputStreamOverrider interface {
}

// FilterInputStream: base class for input stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterInputStream struct {
	_ [0]func() // equal guard
	InputStream
}

var (
	_ InputStreamer = (*FilterInputStream)(nil)
)

// FilterInputStreamer describes types inherited from class FilterInputStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FilterInputStreamer interface {
	coreglib.Objector
	baseFilterInputStream() *FilterInputStream
}

var _ FilterInputStreamer = (*FilterInputStream)(nil)

func classInitFilterInputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFilterInputStream(obj *coreglib.Object) *FilterInputStream {
	return &FilterInputStream{
		InputStream: InputStream{
			Object: obj,
		},
	}
}

func marshalFilterInputStream(p uintptr) (interface{}, error) {
	return wrapFilterInputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (stream *FilterInputStream) baseFilterInputStream() *FilterInputStream {
	return stream
}

// BaseFilterInputStream returns the underlying base object.
func BaseFilterInputStream(obj FilterInputStreamer) *FilterInputStream {
	return obj.baseFilterInputStream()
}

// BaseStream gets the base stream for the filter stream.
//
// The function returns the following values:
//
//    - inputStream: Stream.
//
func (stream *FilterInputStream) BaseStream() InputStreamer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "FilterInputStream")
	_gret := _info.InvokeClassMethod("get_base_stream", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _inputStream InputStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.InputStreamer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(InputStreamer)
			return ok
		})
		rv, ok := casted.(InputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
		}
		_inputStream = rv
	}

	return _inputStream
}

// CloseBaseStream returns whether the base stream will be closed when stream is
// closed.
//
// The function returns the following values:
//
//    - ok: TRUE if the base stream will be closed.
//
func (stream *FilterInputStream) CloseBaseStream() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "FilterInputStream")
	_gret := _info.InvokeClassMethod("get_close_base_stream", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetCloseBaseStream sets whether the base stream will be closed when stream is
// closed.
//
// The function takes the following parameters:
//
//    - closeBase: TRUE to close the base stream.
//
func (stream *FilterInputStream) SetCloseBaseStream(closeBase bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	if closeBase {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gio", "FilterInputStream")
	_info.InvokeClassMethod("set_close_base_stream", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(closeBase)
}
