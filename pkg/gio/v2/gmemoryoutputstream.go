// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeMemoryOutputStream = coreglib.Type(C.g_memory_output_stream_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMemoryOutputStream, F: marshalMemoryOutputStream},
	})
}

// MemoryOutputStreamOverrider contains methods that are overridable.
type MemoryOutputStreamOverrider interface {
}

// MemoryOutputStream is a class for using arbitrary memory chunks as output for
// GIO streaming output operations.
//
// As of GLib 2.34, OutputStream trivially implements OutputStream: it always
// polls as ready.
type MemoryOutputStream struct {
	_ [0]func() // equal guard
	OutputStream

	*coreglib.Object
	PollableOutputStream
	Seekable
}

var (
	_ OutputStreamer    = (*MemoryOutputStream)(nil)
	_ coreglib.Objector = (*MemoryOutputStream)(nil)
)

func initClassMemoryOutputStream(gclass unsafe.Pointer, goval any) {
}

func wrapMemoryOutputStream(obj *coreglib.Object) *MemoryOutputStream {
	return &MemoryOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
		Object: obj,
		PollableOutputStream: PollableOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalMemoryOutputStream(p uintptr) (interface{}, error) {
	return wrapMemoryOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMemoryOutputStreamResizable creates a new OutputStream, using g_realloc()
// and g_free() for memory allocation.
//
// The function returns the following values:
//
func NewMemoryOutputStreamResizable() *MemoryOutputStream {
	var _cret *C.GOutputStream // in

	_cret = C.g_memory_output_stream_new_resizable()

	var _memoryOutputStream *MemoryOutputStream // out

	_memoryOutputStream = wrapMemoryOutputStream(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryOutputStream
}

// Data gets any loaded data from the ostream.
//
// Note that the returned pointer may become invalid on the next write or
// truncate operation on the stream.
//
// The function returns the following values:
//
//    - gpointer (optional): pointer to the stream's data, or NULL if the data
//      has been stolen.
//
func (ostream *MemoryOutputStream) Data() unsafe.Pointer {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gpointer             // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(coreglib.InternObject(ostream).Native()))

	_cret = C.g_memory_output_stream_get_data(_arg0)
	runtime.KeepAlive(ostream)

	var _gpointer unsafe.Pointer // out

	_gpointer = (unsafe.Pointer)(unsafe.Pointer(_cret))

	return _gpointer
}

// DataSize returns the number of bytes from the start up to including the last
// byte written in the stream that has not been truncated away.
//
// The function returns the following values:
//
//    - gsize: number of bytes written to the stream.
//
func (ostream *MemoryOutputStream) DataSize() uint {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gsize                // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(coreglib.InternObject(ostream).Native()))

	_cret = C.g_memory_output_stream_get_data_size(_arg0)
	runtime.KeepAlive(ostream)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Size gets the size of the currently allocated data area (available from
// g_memory_output_stream_get_data()).
//
// You probably don't want to use this function on resizable streams. See
// g_memory_output_stream_get_data_size() instead. For resizable streams the
// size returned by this function is an implementation detail and may be change
// at any time in response to operations on the stream.
//
// If the stream is fixed-sized (ie: no realloc was passed to
// g_memory_output_stream_new()) then this is the maximum size of the stream and
// further writes will return G_IO_ERROR_NO_SPACE.
//
// In any case, if you want the number of bytes currently written to the stream,
// use g_memory_output_stream_get_data_size().
//
// The function returns the following values:
//
//    - gsize: number of bytes allocated for the data buffer.
//
func (ostream *MemoryOutputStream) Size() uint {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gsize                // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(coreglib.InternObject(ostream).Native()))

	_cret = C.g_memory_output_stream_get_size(_arg0)
	runtime.KeepAlive(ostream)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// StealAsBytes returns data from the ostream as a #GBytes. ostream must be
// closed before calling this function.
//
// The function returns the following values:
//
//    - bytes stream's data.
//
func (ostream *MemoryOutputStream) StealAsBytes() *glib.Bytes {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret *C.GBytes              // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(coreglib.InternObject(ostream).Native()))

	_cret = C.g_memory_output_stream_steal_as_bytes(_arg0)
	runtime.KeepAlive(ostream)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _bytes
}

// StealData gets any loaded data from the ostream. Ownership of the data is
// transferred to the caller; when no longer needed it must be freed using the
// free function set in ostream's OutputStream:destroy-function property.
//
// ostream must be closed before calling this function.
//
// The function returns the following values:
//
//    - gpointer (optional) stream's data, or NULL if it has previously been
//      stolen.
//
func (ostream *MemoryOutputStream) StealData() unsafe.Pointer {
	var _arg0 *C.GMemoryOutputStream // out
	var _cret C.gpointer             // in

	_arg0 = (*C.GMemoryOutputStream)(unsafe.Pointer(coreglib.InternObject(ostream).Native()))

	_cret = C.g_memory_output_stream_steal_data(_arg0)
	runtime.KeepAlive(ostream)

	var _gpointer unsafe.Pointer // out

	_gpointer = (unsafe.Pointer)(unsafe.Pointer(_cret))

	return _gpointer
}
