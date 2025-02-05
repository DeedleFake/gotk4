// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// CairoGetClipRectangle: this is a convenience function around
// cairo_clip_extents(). It rounds the clip extents to integer coordinates and
// returns a boolean indicating if a clip area exists.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//
// The function returns the following values:
//
//    - rect (optional): return location for the clip, or NULL.
//    - ok: TRUE if a clip rectangle exists, FALSE if all of cr is clipped and
//      all drawing can be skipped.
//
func CairoGetClipRectangle(cr *cairo.Context) (*Rectangle, bool) {
	var _arg1 *C.cairo_t     // out
	var _arg2 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	_cret = C.gdk_cairo_get_clip_rectangle(_arg1, &_arg2)
	runtime.KeepAlive(cr)

	var _rect *Rectangle // out
	var _ok bool         // out

	_rect = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// CairoRegionCreateFromSurface creates region that describes covers the area
// where the given surface is more than 50% opaque.
//
// This function takes into account device offsets that might be set with
// cairo_surface_set_device_offset().
//
// The function takes the following parameters:
//
//    - surface: cairo surface.
//
// The function returns the following values:
//
//    - region must be freed with cairo_region_destroy().
//
func CairoRegionCreateFromSurface(surface *cairo.Surface) *cairo.Region {
	var _arg1 *C.cairo_surface_t // out
	var _cret *C.cairo_region_t  // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_cairo_region_create_from_surface(_arg1)
	runtime.KeepAlive(surface)

	var _region *cairo.Region // out

	{
		_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
		_region = (*cairo.Region)(unsafe.Pointer(_pp))
	}
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.cairo_region_destroy((*C.cairo_region_t)(unsafe.Pointer(v.Native())))
	})

	return _region
}
