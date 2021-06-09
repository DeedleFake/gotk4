// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/graphene"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

// RoundedRect: a rectangular region with rounded corners.
//
// Application code should normalize rectangles using
// [method@Gsk.RoundedRect.normalize]; this function will ensure that the bounds
// of the rectangle are normalized and ensure that the corner values are
// positive and the corners do not overlap.
//
// All functions taking a `GskRoundedRect` as an argument will internally
// operate on a normalized copy; all functions returning a `GskRoundedRect` will
// always return a normalized one.
//
// The algorithm used for normalizing corner sizes is described in the CSS
// specification (https://drafts.csswg.org/css-backgrounds-3/#border-radius).
type RoundedRect struct {
	native C.GskRoundedRect
}

// WrapRoundedRect wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRoundedRect(ptr unsafe.Pointer) *RoundedRect {
	if ptr == nil {
		return nil
	}

	return (*RoundedRect)(ptr)
}

func marshalRoundedRect(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRoundedRect(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RoundedRect) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Bounds gets the field inside the struct.
func (r *RoundedRect) Bounds() graphene.Rect {
	var v graphene.Rect
	v = *graphene.WrapRect(unsafe.Pointer(&r.native.bounds))
	return v
}

// Corner gets the field inside the struct.
func (r *RoundedRect) Corner() [4]graphene.Size {
	var v [4]graphene.Size
	v = *(*[4]graphene.Size)(unsafe.Pointer(r.native.corner))
	return v
}

// ContainsPoint checks if the given @point is inside the rounded rectangle.
func (s *RoundedRect) ContainsPoint(point *graphene.Point) bool {
	var _arg0 *C.GskRoundedRect
	var _arg1 *C.graphene_point_t

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(point.Native()))

	var _cret C.gboolean

	cret = C.gsk_rounded_rect_contains_point(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// ContainsRect checks if the given @rect is contained inside the rounded
// rectangle.
func (s *RoundedRect) ContainsRect(rect *graphene.Rect) bool {
	var _arg0 *C.GskRoundedRect
	var _arg1 *C.graphene_rect_t

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(rect.Native()))

	var _cret C.gboolean

	cret = C.gsk_rounded_rect_contains_rect(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Init initializes the given `GskRoundedRect` with the given values.
//
// This function will implicitly normalize the `GskRoundedRect` before
// returning.
func (s *RoundedRect) Init(bounds *graphene.Rect, topLeft *graphene.Size, topRight *graphene.Size, bottomRight *graphene.Size, bottomLeft *graphene.Size) *RoundedRect {
	var _arg0 *C.GskRoundedRect
	var _arg1 *C.graphene_rect_t
	var _arg2 *C.graphene_size_t
	var _arg3 *C.graphene_size_t
	var _arg4 *C.graphene_size_t
	var _arg5 *C.graphene_size_t

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	_arg2 = (*C.graphene_size_t)(unsafe.Pointer(topLeft.Native()))
	_arg3 = (*C.graphene_size_t)(unsafe.Pointer(topRight.Native()))
	_arg4 = (*C.graphene_size_t)(unsafe.Pointer(bottomRight.Native()))
	_arg5 = (*C.graphene_size_t)(unsafe.Pointer(bottomLeft.Native()))

	var _cret *C.GskRoundedRect

	cret = C.gsk_rounded_rect_init(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _roundedRect *RoundedRect

	_roundedRect = WrapRoundedRect(unsafe.Pointer(_cret))

	return _roundedRect
}

// InitCopy initializes @self using the given @src rectangle.
//
// This function will not normalize the `GskRoundedRect`, so make sure the
// source is normalized.
func (s *RoundedRect) InitCopy(src *RoundedRect) *RoundedRect {
	var _arg0 *C.GskRoundedRect
	var _arg1 *C.GskRoundedRect

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GskRoundedRect)(unsafe.Pointer(src.Native()))

	var _cret *C.GskRoundedRect

	cret = C.gsk_rounded_rect_init_copy(_arg0, _arg1)

	var _roundedRect *RoundedRect

	_roundedRect = WrapRoundedRect(unsafe.Pointer(_cret))

	return _roundedRect
}

// InitFromRect initializes @self to the given @bounds and sets the radius of
// all four corners to @radius.
func (s *RoundedRect) InitFromRect(bounds *graphene.Rect, radius float32) *RoundedRect {
	var _arg0 *C.GskRoundedRect
	var _arg1 *C.graphene_rect_t
	var _arg2 C.float

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	_arg2 = C.float(radius)

	var _cret *C.GskRoundedRect

	cret = C.gsk_rounded_rect_init_from_rect(_arg0, _arg1, _arg2)

	var _roundedRect *RoundedRect

	_roundedRect = WrapRoundedRect(unsafe.Pointer(_cret))

	return _roundedRect
}

// IntersectsRect checks if part of the given @rect is contained inside the
// rounded rectangle.
func (s *RoundedRect) IntersectsRect(rect *graphene.Rect) bool {
	var _arg0 *C.GskRoundedRect
	var _arg1 *C.graphene_rect_t

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(rect.Native()))

	var _cret C.gboolean

	cret = C.gsk_rounded_rect_intersects_rect(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsRectilinear checks if all corners of @self are right angles and the
// rectangle covers all of its bounds.
//
// This information can be used to decide if [ctor@Gsk.ClipNode.new] or
// [ctor@Gsk.RoundedClipNode.new] should be called.
func (s *RoundedRect) IsRectilinear() bool {
	var _arg0 *C.GskRoundedRect

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.gsk_rounded_rect_is_rectilinear(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Normalize normalizes the passed rectangle.
//
// This function will ensure that the bounds of the rectangle are normalized and
// ensure that the corner values are positive and the corners do not overlap.
func (s *RoundedRect) Normalize() *RoundedRect {
	var _arg0 *C.GskRoundedRect

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))

	var _cret *C.GskRoundedRect

	cret = C.gsk_rounded_rect_normalize(_arg0)

	var _roundedRect *RoundedRect

	_roundedRect = WrapRoundedRect(unsafe.Pointer(_cret))

	return _roundedRect
}

// Offset offsets the bound's origin by @dx and @dy.
//
// The size and corners of the rectangle are unchanged.
func (s *RoundedRect) Offset(dx float32, dy float32) *RoundedRect {
	var _arg0 *C.GskRoundedRect
	var _arg1 C.float
	var _arg2 C.float

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(dx)
	_arg2 = C.float(dy)

	var _cret *C.GskRoundedRect

	cret = C.gsk_rounded_rect_offset(_arg0, _arg1, _arg2)

	var _roundedRect *RoundedRect

	_roundedRect = WrapRoundedRect(unsafe.Pointer(_cret))

	return _roundedRect
}

// Shrink shrinks (or grows) the given rectangle by moving the 4 sides according
// to the offsets given.
//
// The corner radii will be changed in a way that tries to keep the center of
// the corner circle intact. This emulates CSS behavior.
//
// This function also works for growing rectangles if you pass negative values
// for the @top, @right, @bottom or @left.
func (s *RoundedRect) Shrink(top float32, right float32, bottom float32, left float32) *RoundedRect {
	var _arg0 *C.GskRoundedRect
	var _arg1 C.float
	var _arg2 C.float
	var _arg3 C.float
	var _arg4 C.float

	_arg0 = (*C.GskRoundedRect)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(top)
	_arg2 = C.float(right)
	_arg3 = C.float(bottom)
	_arg4 = C.float(left)

	var _cret *C.GskRoundedRect

	cret = C.gsk_rounded_rect_shrink(_arg0, _arg1, _arg2, _arg3, _arg4)

	var _roundedRect *RoundedRect

	_roundedRect = WrapRoundedRect(unsafe.Pointer(_cret))

	return _roundedRect
}
