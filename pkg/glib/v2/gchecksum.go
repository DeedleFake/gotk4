// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

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
		{T: externglib.Type(C.g_checksum_get_type()), F: marshalChecksum},
	})
}

// ChecksumType: the hashing algorithm to be used by #GChecksum when performing
// the digest of some data.
//
// Note that the Type enumeration may be extended at a later date to include new
// hashing algorithm types.
type ChecksumType int

const (
	// MD5: use the MD5 hashing algorithm
	ChecksumTypeMD5 ChecksumType = iota
	// SHA1: use the SHA-1 hashing algorithm
	ChecksumTypeSHA1
	// SHA256: use the SHA-256 hashing algorithm
	ChecksumTypeSHA256
	// SHA512: use the SHA-512 hashing algorithm (Since: 2.36)
	ChecksumTypeSHA512
	// SHA384: use the SHA-384 hashing algorithm (Since: 2.51)
	ChecksumTypeSHA384
)

// Checksum: opaque structure representing a checksumming operation. To create a
// new GChecksum, use g_checksum_new(). To free a GChecksum, use
// g_checksum_free().
type Checksum struct {
	native C.GChecksum
}

// WrapChecksum wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapChecksum(ptr unsafe.Pointer) *Checksum {
	return (*Checksum)(ptr)
}

func marshalChecksum(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Checksum)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *Checksum) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Copy copies a #GChecksum. If @checksum has been closed, by calling
// g_checksum_get_string() or g_checksum_get_digest(), the copied checksum will
// be closed as well.
func (checksum *Checksum) Copy() *Checksum {
	var _arg0 *C.GChecksum // out
	var _cret *C.GChecksum // in

	_arg0 = (*C.GChecksum)(unsafe.Pointer(checksum))

	_cret = C.g_checksum_copy(_arg0)

	var _ret *Checksum // out

	_ret = (*Checksum)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *Checksum) {
		C.free(unsafe.Pointer(v))
	})

	return _ret
}

// Free frees the memory allocated for @checksum.
func (checksum *Checksum) free() {
	var _arg0 *C.GChecksum // out

	_arg0 = (*C.GChecksum)(unsafe.Pointer(checksum))

	C.g_checksum_free(_arg0)
}

// String gets the digest as a hexadecimal string.
//
// Once this function has been called the #GChecksum can no longer be updated
// with g_checksum_update().
//
// The hexadecimal characters will be lower case.
func (checksum *Checksum) String() string {
	var _arg0 *C.GChecksum // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GChecksum)(unsafe.Pointer(checksum))

	_cret = C.g_checksum_get_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Reset resets the state of the @checksum back to its initial state.
func (checksum *Checksum) Reset() {
	var _arg0 *C.GChecksum // out

	_arg0 = (*C.GChecksum)(unsafe.Pointer(checksum))

	C.g_checksum_reset(_arg0)
}

// Update feeds @data into an existing #GChecksum. The checksum must still be
// open, that is g_checksum_get_string() or g_checksum_get_digest() must not
// have been called on @checksum.
func (checksum *Checksum) Update(data []byte) {
	var _arg0 *C.GChecksum // out
	var _arg1 *C.guchar
	var _arg2 C.gssize

	_arg0 = (*C.GChecksum)(unsafe.Pointer(checksum))
	_arg2 = C.gssize(len(data))
	_arg1 = (*C.guchar)(unsafe.Pointer(&data[0]))

	C.g_checksum_update(_arg0, _arg1, _arg2)
}
