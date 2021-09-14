// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_sphere_get_type()), F: marshalSphere},
	})
}

// Sphere: sphere, represented by its center and radius.
//
// An instance of this type is always passed by reference.
type Sphere struct {
	*sphere
}

// sphere is the struct that's finalized.
type sphere struct {
	native *C.graphene_sphere_t
}

func marshalSphere(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Sphere{&sphere{(*C.graphene_sphere_t)(unsafe.Pointer(b))}}, nil
}

// NewSphereAlloc constructs a struct Sphere.
func NewSphereAlloc() *Sphere {
	var _cret *C.graphene_sphere_t // in

	_cret = C.graphene_sphere_alloc()

	var _sphere *Sphere // out

	_sphere = (*Sphere)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_sphere)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_sphere_free((*C.graphene_sphere_t)(intern.C))
		},
	)

	return _sphere
}

// ContainsPoint checks whether the given point is contained in the volume of a
// #graphene_sphere_t.
func (s *Sphere) ContainsPoint(point *Point3D) bool {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	_cret = C.graphene_sphere_contains_point(_arg0, _arg1)
	runtime.KeepAlive(s)
	runtime.KeepAlive(point)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Distance computes the distance of the given point from the surface of a
// #graphene_sphere_t.
func (s *Sphere) Distance(point *Point3D) float32 {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C.float               // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	_cret = C.graphene_sphere_distance(_arg0, _arg1)
	runtime.KeepAlive(s)
	runtime.KeepAlive(point)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Equal checks whether two #graphene_sphere_t are equal.
func (a *Sphere) Equal(b *Sphere) bool {
	var _arg0 *C.graphene_sphere_t // out
	var _arg1 *C.graphene_sphere_t // out
	var _cret C._Bool              // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_sphere_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// BoundingBox computes the bounding box capable of containing the given
// #graphene_sphere_t.
func (s *Sphere) BoundingBox() Box {
	var _arg0 *C.graphene_sphere_t // out
	var _arg1 C.graphene_box_t     // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))

	C.graphene_sphere_get_bounding_box(_arg0, &_arg1)
	runtime.KeepAlive(s)

	var _box Box // out

	_box = *(*Box)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _box
}

// Center retrieves the coordinates of the center of a #graphene_sphere_t.
func (s *Sphere) Center() Point3D {
	var _arg0 *C.graphene_sphere_t // out
	var _arg1 C.graphene_point3d_t // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))

	C.graphene_sphere_get_center(_arg0, &_arg1)
	runtime.KeepAlive(s)

	var _center Point3D // out

	_center = *(*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _center
}

// Radius retrieves the radius of a #graphene_sphere_t.
func (s *Sphere) Radius() float32 {
	var _arg0 *C.graphene_sphere_t // out
	var _cret C.float              // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))

	_cret = C.graphene_sphere_get_radius(_arg0)
	runtime.KeepAlive(s)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes the given #graphene_sphere_t with the given center and
// radius.
func (s *Sphere) Init(center *Point3D, radius float32) *Sphere {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.float               // out
	var _cret *C.graphene_sphere_t  // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))
	if center != nil {
		_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(center)))
	}
	_arg2 = C.float(radius)

	_cret = C.graphene_sphere_init(_arg0, _arg1, _arg2)
	runtime.KeepAlive(s)
	runtime.KeepAlive(center)
	runtime.KeepAlive(radius)

	var _sphere *Sphere // out

	_sphere = (*Sphere)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _sphere
}

// InitFromPoints initializes the given #graphene_sphere_t using the given array
// of 3D coordinates so that the sphere includes them.
//
// The center of the sphere can either be specified, or will be center of the 3D
// volume that encompasses all points.
func (s *Sphere) InitFromPoints(points []Point3D, center *Point3D) *Sphere {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg2 *C.graphene_point3d_t // out
	var _arg1 C.uint
	var _arg3 *C.graphene_point3d_t // out
	var _cret *C.graphene_sphere_t  // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))
	_arg1 = (C.uint)(len(points))
	if len(points) > 0 {
		_arg2 = (*C.graphene_point3d_t)(unsafe.Pointer(&points[0]))
	}
	if center != nil {
		_arg3 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(center)))
	}

	_cret = C.graphene_sphere_init_from_points(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(s)
	runtime.KeepAlive(points)
	runtime.KeepAlive(center)

	var _sphere *Sphere // out

	_sphere = (*Sphere)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _sphere
}

// InitFromVectors initializes the given #graphene_sphere_t using the given
// array of 3D coordinates so that the sphere includes them.
//
// The center of the sphere can either be specified, or will be center of the 3D
// volume that encompasses all vectors.
func (s *Sphere) InitFromVectors(vectors []Vec3, center *Point3D) *Sphere {
	var _arg0 *C.graphene_sphere_t // out
	var _arg2 *C.graphene_vec3_t   // out
	var _arg1 C.uint
	var _arg3 *C.graphene_point3d_t // out
	var _cret *C.graphene_sphere_t  // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))
	_arg1 = (C.uint)(len(vectors))
	_arg2 = (*C.graphene_vec3_t)(C.malloc(C.ulong(len(vectors)) * C.ulong(C.sizeof_graphene_vec3_t)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.graphene_vec3_t)(_arg2), len(vectors))
		for i := range vectors {
			out[i] = *(*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer((&vectors[i]))))
		}
	}
	if center != nil {
		_arg3 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(center)))
	}

	_cret = C.graphene_sphere_init_from_vectors(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(s)
	runtime.KeepAlive(vectors)
	runtime.KeepAlive(center)

	var _sphere *Sphere // out

	_sphere = (*Sphere)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _sphere
}

// IsEmpty checks whether the sphere has a zero radius.
func (s *Sphere) IsEmpty() bool {
	var _arg0 *C.graphene_sphere_t // out
	var _cret C._Bool              // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))

	_cret = C.graphene_sphere_is_empty(_arg0)
	runtime.KeepAlive(s)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Translate translates the center of the given #graphene_sphere_t using the
// point coordinates as the delta of the translation.
func (s *Sphere) Translate(point *Point3D) Sphere {
	var _arg0 *C.graphene_sphere_t  // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.graphene_sphere_t   // in

	_arg0 = (*C.graphene_sphere_t)(gextras.StructNative(unsafe.Pointer(s)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	C.graphene_sphere_translate(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(s)
	runtime.KeepAlive(point)

	var _res Sphere // out

	_res = *(*Sphere)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}
