// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_mapped_file_get_type()), F: marshalMappedFile},
	})
}

// MappedFile: the File represents a file mapping created with
// g_mapped_file_new(). It has only private members and should not be accessed
// directly.
type MappedFile struct {
	native C.GMappedFile
}

// WrapMappedFile wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMappedFile(ptr unsafe.Pointer) *MappedFile {
	return (*MappedFile)(ptr)
}

func marshalMappedFile(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*MappedFile)(unsafe.Pointer(b)), nil
}

// NewMappedFile constructs a struct MappedFile.
func NewMappedFile(filename string, writable bool) (*MappedFile, error) {
	var _arg1 *C.gchar       // out
	var _arg2 C.gboolean     // out
	var _cret *C.GMappedFile // in
	var _cerr *C.GError      // in

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	if writable {
		_arg2 = C.TRUE
	}

	_cret = C.g_mapped_file_new(_arg1, _arg2, &_cerr)

	var _mappedFile *MappedFile // out
	var _goerr error            // out

	_mappedFile = (*MappedFile)(unsafe.Pointer(*C.GMappedFile))
	C.g_mapped_file_ref(_cret)
	runtime.SetFinalizer(_mappedFile, func(v *MappedFile) {
		C.g_mapped_file_unref((*C.GMappedFile)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _mappedFile, _goerr
}

// NewMappedFileFromFd constructs a struct MappedFile.
func NewMappedFileFromFd(fd int, writable bool) (*MappedFile, error) {
	var _arg1 C.gint         // out
	var _arg2 C.gboolean     // out
	var _cret *C.GMappedFile // in
	var _cerr *C.GError      // in

	_arg1 = C.gint(fd)
	if writable {
		_arg2 = C.TRUE
	}

	_cret = C.g_mapped_file_new_from_fd(_arg1, _arg2, &_cerr)

	var _mappedFile *MappedFile // out
	var _goerr error            // out

	_mappedFile = (*MappedFile)(unsafe.Pointer(*C.GMappedFile))
	C.g_mapped_file_ref(_cret)
	runtime.SetFinalizer(_mappedFile, func(v *MappedFile) {
		C.g_mapped_file_unref((*C.GMappedFile)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _mappedFile, _goerr
}

// Native returns the underlying C source pointer.
func (m *MappedFile) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}

// Free: this call existed before File had refcounting and is currently exactly
// the same as g_mapped_file_unref().
//
// Deprecated: since version 2.22.
func (f *MappedFile) free() {
	var _arg0 *C.GMappedFile // out

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(*MappedFile))

	C.g_mapped_file_free(_arg0)
}

// Contents returns the contents of a File.
//
// Note that the contents may not be zero-terminated, even if the File is backed
// by a text file.
//
// If the file is empty then nil is returned.
func (f *MappedFile) Contents() string {
	var _arg0 *C.GMappedFile // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(*MappedFile))

	_cret = C.g_mapped_file_get_contents(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Length returns the length of the contents of a File.
func (f *MappedFile) Length() uint {
	var _arg0 *C.GMappedFile // out
	var _cret C.gsize        // in

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(*MappedFile))

	_cret = C.g_mapped_file_get_length(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Ref increments the reference count of @file by one. It is safe to call this
// function from any thread.
func (f *MappedFile) ref() *MappedFile {
	var _arg0 *C.GMappedFile // out
	var _cret *C.GMappedFile // in

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(*MappedFile))

	_cret = C.g_mapped_file_ref(_arg0)

	var _mappedFile *MappedFile // out

	_mappedFile = (*MappedFile)(unsafe.Pointer(*C.GMappedFile))
	C.g_mapped_file_ref(_cret)
	runtime.SetFinalizer(_mappedFile, func(v *MappedFile) {
		C.g_mapped_file_unref((*C.GMappedFile)(unsafe.Pointer(v)))
	})

	return _mappedFile
}

// Unref decrements the reference count of @file by one. If the reference count
// drops to 0, unmaps the buffer of @file and frees it.
//
// It is safe to call this function from any thread.
//
// Since 2.22
func (f *MappedFile) unref() {
	var _arg0 *C.GMappedFile // out

	_arg0 = (*C.GMappedFile)(unsafe.Pointer(*MappedFile))

	C.g_mapped_file_unref(_arg0)
}
