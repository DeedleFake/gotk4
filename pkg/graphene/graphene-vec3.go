// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_vec3_get_type()), F: marshalVec3},
	})
}

// Vec3: structure capable of holding a vector with three dimensions: x, y, and
// z.
//
// The contents of the #graphene_vec3_t structure are private and should never
// be accessed directly.
type Vec3 struct {
	native C.graphene_vec3_t
}

// WrapVec3 wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapVec3(ptr unsafe.Pointer) *Vec3 {
	return (*Vec3)(ptr)
}

func marshalVec3(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Vec3)(unsafe.Pointer(b)), nil
}

// NewVec3Alloc constructs a struct Vec3.
func NewVec3Alloc() *Vec3 {
	var _cret *C.graphene_vec3_t // in

	_cret = C.graphene_vec3_alloc()

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(*C.graphene_vec3_t))
	runtime.SetFinalizer(_vec3, func(v *Vec3) {
		C.free(unsafe.Pointer(v))
	})

	return _vec3
}

// Native returns the underlying C source pointer.
func (v *Vec3) Native() unsafe.Pointer {
	return unsafe.Pointer(&v.native)
}

// Dot computes the dot product of the two given vectors.
func (a *Vec3) Dot(b *Vec3) float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	_cret = C.graphene_vec3_dot(_arg0, _arg1)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether the two given #graphene_vec3_t are equal.
func (v *Vec3) Equal(v2 *Vec3) bool {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	_cret = C.graphene_vec3_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free frees the resources allocated by @v
func (v *Vec3) free() {
	var _arg0 *C.graphene_vec3_t // out

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	C.graphene_vec3_free(_arg0)
}

// X retrieves the first component of the given vector @v.
func (v *Vec3) X() float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	_cret = C.graphene_vec3_get_x(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Y retrieves the second component of the given vector @v.
func (v *Vec3) Y() float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	_cret = C.graphene_vec3_get_y(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Z retrieves the third component of the given vector @v.
func (v *Vec3) Z() float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	_cret = C.graphene_vec3_get_z(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes a #graphene_vec3_t using the given values.
//
// This function can be called multiple times.
func (v *Vec3) Init(x float32, y float32, z float32) *Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _arg3 C.float            // out
	var _cret *C.graphene_vec3_t // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(z)

	_cret = C.graphene_vec3_init(_arg0, _arg1, _arg2, _arg3)

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(*C.graphene_vec3_t))

	return _vec3
}

// InitFromFloat initializes a #graphene_vec3_t with the values from an array.
func (v *Vec3) InitFromFloat(src [3]float32) *Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.float
	var _cret *C.graphene_vec3_t // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))
	_arg1 = (*C.float)(unsafe.Pointer(&src))

	_cret = C.graphene_vec3_init_from_float(_arg0, _arg1)

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(*C.graphene_vec3_t))

	return _vec3
}

// InitFromVec3 initializes a #graphene_vec3_t with the values of another
// #graphene_vec3_t.
func (v *Vec3) InitFromVec3(src *Vec3) *Vec3 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _cret *C.graphene_vec3_t // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	_cret = C.graphene_vec3_init_from_vec3(_arg0, _arg1)

	var _vec3 *Vec3 // out

	_vec3 = (*Vec3)(unsafe.Pointer(*C.graphene_vec3_t))

	return _vec3
}

// Length retrieves the length of the given vector @v.
func (v *Vec3) Length() float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _cret C.float            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	_cret = C.graphene_vec3_length(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Near compares the two given #graphene_vec3_t vectors and checks whether their
// values are within the given @epsilon.
func (v *Vec3) Near(v2 *Vec3, epsilon float32) bool {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 *C.graphene_vec3_t // out
	var _arg2 C.float            // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))
	_arg2 = C.float(epsilon)

	_cret = C.graphene_vec3_near(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ToFloat copies the components of a #graphene_vec3_t into the given array.
func (v *Vec3) ToFloat() [3]float32 {
	var _arg0 *C.graphene_vec3_t // out
	var _arg1 [3]C.float

	_arg0 = (*C.graphene_vec3_t)(unsafe.Pointer(*Vec3))

	C.graphene_vec3_to_float(_arg0, &_arg1[0])

	var _dest [3]float32

	_dest = *(*[3]float32)(unsafe.Pointer(&_arg1))

	return _dest
}
