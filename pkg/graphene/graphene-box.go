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
// #include <stdlib.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_box_get_type()), F: marshalBox},
	})
}

// Box: 3D box, described as the volume between a minimum and a maximum
// vertices.
//
// An instance of this type is always passed by reference.
type Box struct {
	*box
}

// box is the struct that's finalized.
type box struct {
	native *C.graphene_box_t
}

func marshalBox(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Box{&box{(*C.graphene_box_t)(b)}}, nil
}

// NewBoxAlloc constructs a struct Box.
func NewBoxAlloc() *Box {
	var _cret *C.graphene_box_t // in

	_cret = C.graphene_box_alloc()

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_box)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.graphene_box_free((*C.graphene_box_t)(intern.C))
		},
	)

	return _box
}

// ContainsBox checks whether the #graphene_box_t a contains the given
// #graphene_box_t b.
func (a *Box) ContainsBox(b *Box) bool {
	var _arg0 *C.graphene_box_t // out
	var _arg1 *C.graphene_box_t // out
	var _cret C._Bool           // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_box_contains_box(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ContainsPoint checks whether box contains the given point.
func (box *Box) ContainsPoint(point *Point3D) bool {
	var _arg0 *C.graphene_box_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C._Bool               // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	_cret = C.graphene_box_contains_point(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(point)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Equal checks whether the two given boxes are equal.
func (a *Box) Equal(b *Box) bool {
	var _arg0 *C.graphene_box_t // out
	var _arg1 *C.graphene_box_t // out
	var _cret C._Bool           // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_box_equal(_arg0, _arg1)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Expand expands the dimensions of box to include the coordinates at point.
func (box *Box) Expand(point *Point3D) *Box {
	var _arg0 *C.graphene_box_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.graphene_box_t      // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	C.graphene_box_expand(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(point)

	var _res *Box // out

	_res = (*Box)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// ExpandScalar expands the dimensions of box by the given scalar value.
//
// If scalar is positive, the #graphene_box_t will grow; if scalar is negative,
// the #graphene_box_t will shrink.
func (box *Box) ExpandScalar(scalar float32) *Box {
	var _arg0 *C.graphene_box_t // out
	var _arg1 C.float           // out
	var _arg2 C.graphene_box_t  // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	_arg1 = C.float(scalar)

	C.graphene_box_expand_scalar(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(scalar)

	var _res *Box // out

	_res = (*Box)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// ExpandVec3 expands the dimensions of box to include the coordinates of the
// given vector.
func (box *Box) ExpandVec3(vec *Vec3) *Box {
	var _arg0 *C.graphene_box_t  // out
	var _arg1 *C.graphene_vec3_t // out
	var _arg2 C.graphene_box_t   // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(vec)))

	C.graphene_box_expand_vec3(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(vec)

	var _res *Box // out

	_res = (*Box)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// BoundingSphere computes the bounding #graphene_sphere_t capable of containing
// the given #graphene_box_t.
func (box *Box) BoundingSphere() *Sphere {
	var _arg0 *C.graphene_box_t   // out
	var _arg1 C.graphene_sphere_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	C.graphene_box_get_bounding_sphere(_arg0, &_arg1)
	runtime.KeepAlive(box)

	var _sphere *Sphere // out

	_sphere = (*Sphere)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _sphere
}

// Center retrieves the coordinates of the center of a #graphene_box_t.
func (box *Box) Center() *Point3D {
	var _arg0 *C.graphene_box_t    // out
	var _arg1 C.graphene_point3d_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	C.graphene_box_get_center(_arg0, &_arg1)
	runtime.KeepAlive(box)

	var _center *Point3D // out

	_center = (*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _center
}

// Depth retrieves the size of the box on the Z axis.
func (box *Box) Depth() float32 {
	var _arg0 *C.graphene_box_t // out
	var _cret C.float           // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	_cret = C.graphene_box_get_depth(_arg0)
	runtime.KeepAlive(box)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Height retrieves the size of the box on the Y axis.
func (box *Box) Height() float32 {
	var _arg0 *C.graphene_box_t // out
	var _cret C.float           // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	_cret = C.graphene_box_get_height(_arg0)
	runtime.KeepAlive(box)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Max retrieves the coordinates of the maximum point of the given
// #graphene_box_t.
func (box *Box) Max() *Point3D {
	var _arg0 *C.graphene_box_t    // out
	var _arg1 C.graphene_point3d_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	C.graphene_box_get_max(_arg0, &_arg1)
	runtime.KeepAlive(box)

	var _max *Point3D // out

	_max = (*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _max
}

// Min retrieves the coordinates of the minimum point of the given
// #graphene_box_t.
func (box *Box) Min() *Point3D {
	var _arg0 *C.graphene_box_t    // out
	var _arg1 C.graphene_point3d_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	C.graphene_box_get_min(_arg0, &_arg1)
	runtime.KeepAlive(box)

	var _min *Point3D // out

	_min = (*Point3D)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _min
}

// Size retrieves the size of the box on all three axes, and stores it into the
// given size vector.
func (box *Box) Size() *Vec3 {
	var _arg0 *C.graphene_box_t // out
	var _arg1 C.graphene_vec3_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	C.graphene_box_get_size(_arg0, &_arg1)
	runtime.KeepAlive(box)

	var _size *Vec3 // out

	_size = (*Vec3)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _size
}

// Vertices computes the vertices of the given #graphene_box_t.
func (box *Box) Vertices() [8]Vec3 {
	var _arg0 *C.graphene_box_t    // out
	var _arg1 [8]C.graphene_vec3_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	C.graphene_box_get_vertices(_arg0, &_arg1[0])
	runtime.KeepAlive(box)

	var _vertices [8]Vec3 // out

	{
		src := &_arg1
		for i := 0; i < 8; i++ {
			_vertices[i] = *(*Vec3)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}

	return _vertices
}

// Width retrieves the size of the box on the X axis.
func (box *Box) Width() float32 {
	var _arg0 *C.graphene_box_t // out
	var _cret C.float           // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))

	_cret = C.graphene_box_get_width(_arg0)
	runtime.KeepAlive(box)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Init initializes the given #graphene_box_t with two vertices.
func (box *Box) Init(min, max *Point3D) *Box {
	var _arg0 *C.graphene_box_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 *C.graphene_point3d_t // out
	var _cret *C.graphene_box_t     // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	if min != nil {
		_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(min)))
	}
	if max != nil {
		_arg2 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(max)))
	}

	_cret = C.graphene_box_init(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)

	var _ret *Box // out

	_ret = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ret
}

// InitFromBox initializes the given #graphene_box_t with the vertices of
// another #graphene_box_t.
func (box *Box) InitFromBox(src *Box) *Box {
	var _arg0 *C.graphene_box_t // out
	var _arg1 *C.graphene_box_t // out
	var _cret *C.graphene_box_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	_arg1 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.graphene_box_init_from_box(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(src)

	var _ret *Box // out

	_ret = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ret
}

// InitFromPoints initializes the given #graphene_box_t with the given array of
// vertices.
//
// If n_points is 0, the returned box is initialized with graphene_box_empty().
func (box *Box) InitFromPoints(points []Point3D) *Box {
	var _arg0 *C.graphene_box_t     // out
	var _arg2 *C.graphene_point3d_t // out
	var _arg1 C.uint
	var _cret *C.graphene_box_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	_arg1 = (C.uint)(len(points))
	_arg2 = (*C.graphene_point3d_t)(C.calloc(C.size_t(len(points)), C.size_t(C.sizeof_graphene_point3d_t)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.graphene_point3d_t)(_arg2), len(points))
		for i := range points {
			out[i] = *(*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer((&points[i]))))
		}
	}

	_cret = C.graphene_box_init_from_points(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(points)

	var _ret *Box // out

	_ret = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ret
}

// InitFromVec3 initializes the given #graphene_box_t with two vertices stored
// inside #graphene_vec3_t.
func (box *Box) InitFromVec3(min, max *Vec3) *Box {
	var _arg0 *C.graphene_box_t  // out
	var _arg1 *C.graphene_vec3_t // out
	var _arg2 *C.graphene_vec3_t // out
	var _cret *C.graphene_box_t  // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	if min != nil {
		_arg1 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(min)))
	}
	if max != nil {
		_arg2 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(max)))
	}

	_cret = C.graphene_box_init_from_vec3(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)

	var _ret *Box // out

	_ret = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ret
}

// InitFromVectors initializes the given #graphene_box_t with the given array of
// vertices.
//
// If n_vectors is 0, the returned box is initialized with graphene_box_empty().
func (box *Box) InitFromVectors(vectors []Vec3) *Box {
	var _arg0 *C.graphene_box_t  // out
	var _arg2 *C.graphene_vec3_t // out
	var _arg1 C.uint
	var _cret *C.graphene_box_t // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(box)))
	_arg1 = (C.uint)(len(vectors))
	_arg2 = (*C.graphene_vec3_t)(C.calloc(C.size_t(len(vectors)), C.size_t(C.sizeof_graphene_vec3_t)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.graphene_vec3_t)(_arg2), len(vectors))
		for i := range vectors {
			out[i] = *(*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer((&vectors[i]))))
		}
	}

	_cret = C.graphene_box_init_from_vectors(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(vectors)

	var _ret *Box // out

	_ret = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ret
}

// Intersection intersects the two given #graphene_box_t.
//
// If the two boxes do not intersect, res will contain a degenerate box
// initialized with graphene_box_empty().
func (a *Box) Intersection(b *Box) (*Box, bool) {
	var _arg0 *C.graphene_box_t // out
	var _arg1 *C.graphene_box_t // out
	var _arg2 C.graphene_box_t  // in
	var _cret C._Bool           // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(b)))

	_cret = C.graphene_box_intersection(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Box // out
	var _ok bool  // out

	_res = (*Box)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret {
		_ok = true
	}

	return _res, _ok
}

// Union unions the two given #graphene_box_t.
func (a *Box) Union(b *Box) *Box {
	var _arg0 *C.graphene_box_t // out
	var _arg1 *C.graphene_box_t // out
	var _arg2 C.graphene_box_t  // in

	_arg0 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(a)))
	_arg1 = (*C.graphene_box_t)(gextras.StructNative(unsafe.Pointer(b)))

	C.graphene_box_union(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)

	var _res *Box // out

	_res = (*Box)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _res
}

// BoxEmpty: degenerate #graphene_box_t that can only be expanded.
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxEmpty() *Box {
	var _cret *C.graphene_box_t // in

	_cret = C.graphene_box_empty()

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _box
}

// BoxInfinite: degenerate #graphene_box_t that cannot be expanded.
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxInfinite() *Box {
	var _cret *C.graphene_box_t // in

	_cret = C.graphene_box_infinite()

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _box
}

// BoxMinusOne with the minimum vertex set at (-1, -1, -1) and the maximum
// vertex set at (0, 0, 0).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxMinusOne() *Box {
	var _cret *C.graphene_box_t // in

	_cret = C.graphene_box_minus_one()

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _box
}

// BoxOne with the minimum vertex set at (0, 0, 0) and the maximum vertex set at
// (1, 1, 1).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxOne() *Box {
	var _cret *C.graphene_box_t // in

	_cret = C.graphene_box_one()

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _box
}

// BoxOneMinusOne with the minimum vertex set at (-1, -1, -1) and the maximum
// vertex set at (1, 1, 1).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxOneMinusOne() *Box {
	var _cret *C.graphene_box_t // in

	_cret = C.graphene_box_one_minus_one()

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _box
}

// BoxZero with both the minimum and maximum vertices set at (0, 0, 0).
//
// The returned value is owned by Graphene and should not be modified or freed.
func BoxZero() *Box {
	var _cret *C.graphene_box_t // in

	_cret = C.graphene_box_zero()

	var _box *Box // out

	_box = (*Box)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _box
}
