// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_point3d_get_type()), F: marshalPoint3D},
	})
}

// Point3D: point with three components: X, Y, and Z.
type Point3D struct {
	nocopy gextras.NoCopy
	native *C.graphene_point3d_t
}

func marshalPoint3D(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Point3D{native: (*C.graphene_point3d_t)(unsafe.Pointer(b))}, nil
}

// NewPoint3DAlloc constructs a struct Point3D.
func NewPoint3DAlloc() *Point3D {
	var _cret *C.graphene_point3d_t // in

	_cret = C.graphene_point3d_alloc()

	var _point3D *Point3D // out

	_point3D = (*Point3D)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_point3D, func(v *Point3D) {
		C.graphene_point3d_free((*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _point3D
}

// X coordinate
func (p *Point3D) X() float32 {
	var v float32 // out
	v = float32(p.native.x)
	return v
}

// Y coordinate
func (p *Point3D) Y() float32 {
	var v float32 // out
	v = float32(p.native.y)
	return v
}

// Z coordinate
func (p *Point3D) Z() float32 {
	var v float32 // out
	v = float32(p.native.z)
	return v
}

// Cross computes the cross product of the two given #graphene_point3d_t.
func (a *Point3D) Cross(b *Point3D) Point3D {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.graphene_point3d_t  // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_point3d_cross(_arg0, _arg1, &_arg2)

	var _res Point3D // out

	_res = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// Distance computes the distance between the two given #graphene_point3d_t.
func (a *Point3D) Distance(b *Point3D) (Vec3, float32) {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.graphene_vec3_t     // in
	var _cret C.float               // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_point3d_distance(_arg0, _arg1, &_arg2)

	var _delta Vec3     // out
	var _gfloat float32 // out

	_delta = *(*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_gfloat = float32(_cret)

	return _delta, _gfloat
}

// Dot computes the dot product of the two given #graphene_point3d_t.
func (a *Point3D) Dot(b *Point3D) float32 {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C.float               // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_point3d_dot(_arg0, _arg1)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether two given points are equal.
func (a *Point3D) Equal(b *Point3D) bool {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_point3d_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Init initializes a #graphene_point3d_t with the given coordinates.
func (p *Point3D) Init(x float32, y float32, z float32) *Point3D {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 C.float               // out
	var _arg2 C.float               // out
	var _arg3 C.float               // out
	var _cret *C.graphene_point3d_t // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = C.float(x)
	_arg2 = C.float(y)
	_arg3 = C.float(z)

	_cret = C.graphene_point3d_init(_arg0, _arg1, _arg2, _arg3)

	var _point3D *Point3D // out

	_point3D = (*Point3D)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point3D
}

// InitFromPoint initializes a #graphene_point3d_t using the coordinates of
// another #graphene_point3d_t.
func (p *Point3D) InitFromPoint(src *Point3D) *Point3D {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret *C.graphene_point3d_t // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_point3d_init_from_point(_arg0, _arg1)

	var _point3D *Point3D // out

	_point3D = (*Point3D)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point3D
}

// InitFromVec3 initializes a #graphene_point3d_t using the components of a
// #graphene_vec3_t.
func (p *Point3D) InitFromVec3(v *Vec3) *Point3D {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_vec3_t    // out
	var _cret *C.graphene_point3d_t // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(v)))

	_cret = C.graphene_point3d_init_from_vec3(_arg0, _arg1)

	var _point3D *Point3D // out

	_point3D = (*Point3D)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point3D
}

// Interpolate: linearly interpolates each component of a and b using the
// provided factor, and places the result in res.
func (a *Point3D) Interpolate(b *Point3D, factor float64) Point3D {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.double              // out
	var _arg3 C.graphene_point3d_t  // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(b)))
	_arg2 = C.double(factor)

	C.graphene_point3d_interpolate(_arg0, _arg1, _arg2, &_arg3)

	var _res Point3D // out

	_res = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _res
}

// Length computes the length of the vector represented by the coordinates of
// the given #graphene_point3d_t.
func (p *Point3D) Length() float32 {
	var _arg0 *C.graphene_point3d_t // out
	var _cret C.float               // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))

	_cret = C.graphene_point3d_length(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Near checks whether the two points are near each other, within an epsilon
// factor.
func (a *Point3D) Near(b *Point3D, epsilon float32) bool {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.float               // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(b)))
	_arg2 = C.float(epsilon)

	_cret = C.graphene_point3d_near(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Normalize computes the normalization of the vector represented by the
// coordinates of the given #graphene_point3d_t.
func (p *Point3D) Normalize() Point3D {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 C.graphene_point3d_t  // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))

	C.graphene_point3d_normalize(_arg0, &_arg1)

	var _res Point3D // out

	_res = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _res
}

// NormalizeViewport normalizes the coordinates of a #graphene_point3d_t using
// the given viewport and clipping planes.
//
// The coordinates of the resulting #graphene_point3d_t will be in the [ -1, 1 ]
// range.
func (p *Point3D) NormalizeViewport(viewport *Rect, zNear float32, zFar float32) Point3D {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 *C.graphene_rect_t    // out
	var _arg2 C.float               // out
	var _arg3 C.float               // out
	var _arg4 C.graphene_point3d_t  // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(viewport)))
	_arg2 = C.float(zNear)
	_arg3 = C.float(zFar)

	C.graphene_point3d_normalize_viewport(_arg0, _arg1, _arg2, _arg3, &_arg4)

	var _res Point3D // out

	_res = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))

	return _res
}

// Scale scales the coordinates of the given #graphene_point3d_t by the given
// factor.
func (p *Point3D) Scale(factor float32) Point3D {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 C.float               // out
	var _arg2 C.graphene_point3d_t  // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))
	_arg1 = C.float(factor)

	C.graphene_point3d_scale(_arg0, _arg1, &_arg2)

	var _res Point3D // out

	_res = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// ToVec3 stores the coordinates of a #graphene_point3d_t into a
// #graphene_vec3_t.
func (p *Point3D) ToVec3() Vec3 {
	var _arg0 *C.graphene_point3d_t // out
	var _arg1 C.graphene_vec3_t     // in

	_arg0 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(p)))

	C.graphene_point3d_to_vec3(_arg0, &_arg1)

	var _v Vec3 // out

	_v = *(*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _v
}

// Point3DZero retrieves a constant point with all three coordinates set to 0.
func Point3DZero() *Point3D {
	var _cret *C.graphene_point3d_t // in

	_cret = C.graphene_point3d_zero()

	var _point3D *Point3D // out

	_point3D = (*Point3D)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _point3D
}
