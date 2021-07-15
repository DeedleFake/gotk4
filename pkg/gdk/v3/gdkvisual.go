// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_visual_type_get_type()), F: marshalVisualType},
		{T: externglib.Type(C.gdk_visual_get_type()), F: marshalVisualer},
	})
}

// VisualType: set of values that describe the manner in which the pixel values
// for a visual are converted into RGB values for display.
type VisualType int

const (
	// StaticGray: each pixel value indexes a grayscale value directly.
	VisualTypeStaticGray VisualType = iota
	// Grayscale: each pixel is an index into a color map that maps pixel values
	// into grayscale values. The color map can be changed by an application.
	VisualTypeGrayscale
	// StaticColor: each pixel value is an index into a predefined, unmodifiable
	// color map that maps pixel values into RGB values.
	VisualTypeStaticColor
	// PseudoColor: each pixel is an index into a color map that maps pixel
	// values into rgb values. The color map can be changed by an application.
	VisualTypePseudoColor
	// TrueColor: each pixel value directly contains red, green, and blue
	// components. Use gdk_visual_get_red_pixel_details(), etc, to obtain
	// information about how the components are assembled into a pixel value.
	VisualTypeTrueColor
	// DirectColor: each pixel value contains red, green, and blue components as
	// for GDK_VISUAL_TRUE_COLOR, but the components are mapped via a color
	// table into the final output table instead of being converted directly.
	VisualTypeDirectColor
)

func marshalVisualType(p uintptr) (interface{}, error) {
	return VisualType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// QueryDepths: this function returns the available bit depths for the default
// screen. It’s equivalent to listing the visuals (gdk_list_visuals()) and then
// looking at the depth field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
func QueryDepths() []int {
	var _arg1 *C.gint
	var _arg2 C.gint // in

	C.gdk_query_depths(&_arg1, &_arg2)

	var _depths []int

	_depths = unsafe.Slice((*int)(unsafe.Pointer(_arg1)), _arg2)
	runtime.SetFinalizer(&_depths, func(v *[]int) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})

	return _depths
}

// QueryVisualTypes: this function returns the available visual types for the
// default screen. It’s equivalent to listing the visuals (gdk_list_visuals())
// and then looking at the type field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
func QueryVisualTypes() []VisualType {
	var _arg1 *C.GdkVisualType
	var _arg2 C.gint // in

	C.gdk_query_visual_types(&_arg1, &_arg2)

	var _visualTypes []VisualType

	{
		src := unsafe.Slice(_arg1, _arg2)
		_visualTypes = make([]VisualType, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_visualTypes[i] = VisualType(src[i])
		}
	}

	return _visualTypes
}

// Visual contains information about a particular visual.
type Visual struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Visual)(nil)

func wrapVisual(obj *externglib.Object) *Visual {
	return &Visual{
		Object: obj,
	}
}

func marshalVisualer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVisual(obj), nil
}

// BitsPerRGB returns the number of significant bits per red, green and blue
// value.
//
// Not all GDK backend provide a meaningful value for this function.
//
// Deprecated: Use gdk_visual_get_red_pixel_details() and its variants to learn
// about the pixel layout of TrueColor and DirectColor visuals.
func (visual *Visual) BitsPerRGB() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	_cret = C.gdk_visual_get_bits_per_rgb(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// BluePixelDetails obtains values that are needed to calculate blue pixel
// values in TrueColor and DirectColor. The “mask” is the significant bits
// within the pixel. The “shift” is the number of bits left we must shift a
// primary for it to be in position (according to the "mask"). Finally,
// "precision" refers to how much precision the pixel value contains for a
// particular primary.
func (visual *Visual) BluePixelDetails() (mask uint32, shift int, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	C.gdk_visual_get_blue_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = uint32(_arg1)
	_shift = int(_arg2)
	_precision = int(_arg3)

	return _mask, _shift, _precision
}

// ByteOrder returns the byte order of this visual.
//
// The information returned by this function is only relevant when working with
// XImages, and not all backends return meaningful information for this.
//
// Deprecated: This information is not useful.
func (visual *Visual) ByteOrder() ByteOrder {
	var _arg0 *C.GdkVisual   // out
	var _cret C.GdkByteOrder // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	_cret = C.gdk_visual_get_byte_order(_arg0)

	var _byteOrder ByteOrder // out

	_byteOrder = ByteOrder(_cret)

	return _byteOrder
}

// ColormapSize returns the size of a colormap for this visual.
//
// You have to use platform-specific APIs to manipulate colormaps.
//
// Deprecated: This information is not useful, since GDK does not provide APIs
// to operate on colormaps.
func (visual *Visual) ColormapSize() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	_cret = C.gdk_visual_get_colormap_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Depth returns the bit depth of this visual.
func (visual *Visual) Depth() int {
	var _arg0 *C.GdkVisual // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	_cret = C.gdk_visual_get_depth(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GreenPixelDetails obtains values that are needed to calculate green pixel
// values in TrueColor and DirectColor. The “mask” is the significant bits
// within the pixel. The “shift” is the number of bits left we must shift a
// primary for it to be in position (according to the "mask"). Finally,
// "precision" refers to how much precision the pixel value contains for a
// particular primary.
func (visual *Visual) GreenPixelDetails() (mask uint32, shift int, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	C.gdk_visual_get_green_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = uint32(_arg1)
	_shift = int(_arg2)
	_precision = int(_arg3)

	return _mask, _shift, _precision
}

// RedPixelDetails obtains values that are needed to calculate red pixel values
// in TrueColor and DirectColor. The “mask” is the significant bits within the
// pixel. The “shift” is the number of bits left we must shift a primary for it
// to be in position (according to the "mask"). Finally, "precision" refers to
// how much precision the pixel value contains for a particular primary.
func (visual *Visual) RedPixelDetails() (mask uint32, shift int, precision int) {
	var _arg0 *C.GdkVisual // out
	var _arg1 C.guint32    // in
	var _arg2 C.gint       // in
	var _arg3 C.gint       // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	C.gdk_visual_get_red_pixel_details(_arg0, &_arg1, &_arg2, &_arg3)

	var _mask uint32   // out
	var _shift int     // out
	var _precision int // out

	_mask = uint32(_arg1)
	_shift = int(_arg2)
	_precision = int(_arg3)

	return _mask, _shift, _precision
}

// Screen gets the screen to which this visual belongs
func (visual *Visual) Screen() *Screen {
	var _arg0 *C.GdkVisual // out
	var _cret *C.GdkScreen // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	_cret = C.gdk_visual_get_screen(_arg0)

	var _screen *Screen // out

	_screen = wrapScreen(externglib.Take(unsafe.Pointer(_cret)))

	return _screen
}

// VisualType returns the type of visual this is (PseudoColor, TrueColor, etc).
func (visual *Visual) VisualType() VisualType {
	var _arg0 *C.GdkVisual    // out
	var _cret C.GdkVisualType // in

	_arg0 = (*C.GdkVisual)(unsafe.Pointer(visual.Native()))

	_cret = C.gdk_visual_get_visual_type(_arg0)

	var _visualType VisualType // out

	_visualType = VisualType(_cret)

	return _visualType
}

// VisualGetBest: get the visual with the most available colors for the default
// GDK screen. The return value should not be freed.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
func VisualGetBest() *Visual {
	var _cret *C.GdkVisual // in

	_cret = C.gdk_visual_get_best()

	var _visual *Visual // out

	_visual = wrapVisual(externglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// VisualGetBestDepth: get the best available depth for the default GDK screen.
// “Best” means “largest,” i.e. 32 preferred over 24 preferred over 8 bits per
// pixel.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
func VisualGetBestDepth() int {
	var _cret C.gint // in

	_cret = C.gdk_visual_get_best_depth()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// VisualGetBestType: return the best available visual type for the default GDK
// screen.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
func VisualGetBestType() VisualType {
	var _cret C.GdkVisualType // in

	_cret = C.gdk_visual_get_best_type()

	var _visualType VisualType // out

	_visualType = VisualType(_cret)

	return _visualType
}

// VisualGetBestWithBoth combines gdk_visual_get_best_with_depth() and
// gdk_visual_get_best_with_type().
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
func VisualGetBestWithBoth(depth int, visualType VisualType) *Visual {
	var _arg1 C.gint          // out
	var _arg2 C.GdkVisualType // out
	var _cret *C.GdkVisual    // in

	_arg1 = C.gint(depth)
	_arg2 = C.GdkVisualType(visualType)

	_cret = C.gdk_visual_get_best_with_both(_arg1, _arg2)

	var _visual *Visual // out

	_visual = wrapVisual(externglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// VisualGetBestWithDepth: get the best visual with depth depth for the default
// GDK screen. Color visuals and visuals with mutable colormaps are preferred
// over grayscale or fixed-colormap visuals. The return value should not be
// freed. NULL may be returned if no visual supports depth.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
func VisualGetBestWithDepth(depth int) *Visual {
	var _arg1 C.gint       // out
	var _cret *C.GdkVisual // in

	_arg1 = C.gint(depth)

	_cret = C.gdk_visual_get_best_with_depth(_arg1)

	var _visual *Visual // out

	_visual = wrapVisual(externglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// VisualGetBestWithType: get the best visual of the given visual_type for the
// default GDK screen. Visuals with higher color depths are considered better.
// The return value should not be freed. NULL may be returned if no visual has
// type visual_type.
//
// Deprecated: Visual selection should be done using
// gdk_screen_get_system_visual() and gdk_screen_get_rgba_visual().
func VisualGetBestWithType(visualType VisualType) *Visual {
	var _arg1 C.GdkVisualType // out
	var _cret *C.GdkVisual    // in

	_arg1 = C.GdkVisualType(visualType)

	_cret = C.gdk_visual_get_best_with_type(_arg1)

	var _visual *Visual // out

	_visual = wrapVisual(externglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// VisualGetSystem: get the system’s default visual for the default GDK screen.
// This is the visual for the root window of the display. The return value
// should not be freed.
//
// Deprecated: Use gdk_screen_get_system_visual (gdk_screen_get_default ()).
func VisualGetSystem() *Visual {
	var _cret *C.GdkVisual // in

	_cret = C.gdk_visual_get_system()

	var _visual *Visual // out

	_visual = wrapVisual(externglib.Take(unsafe.Pointer(_cret)))

	return _visual
}
