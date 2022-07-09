// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// HSVToRGB converts a color from HSV space to RGB.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
//
// The function takes the following parameters:
//
//    - h: hue.
//    - s: saturation.
//    - v: value.
//
// The function returns the following values:
//
//    - r: return value for the red component.
//    - g: return value for the green component.
//    - b: return value for the blue component.
//
func HSVToRGB(h, s, v float32) (r, g, b float32) {
	var _args [3]girepository.Argument
	var _outs [3]girepository.Argument

	*(*C.float)(unsafe.Pointer(&_args[0])) = C.float(h)
	*(*C.float)(unsafe.Pointer(&_args[1])) = C.float(s)
	*(*C.float)(unsafe.Pointer(&_args[2])) = C.float(v)

	_info := girepository.MustFind("Gtk", "hsv_to_rgb")
	_info.Invoke(_args[:], _outs[:])

	runtime.KeepAlive(h)
	runtime.KeepAlive(s)
	runtime.KeepAlive(v)

	var _r float32 // out
	var _g float32 // out
	var _b float32 // out

	_r = *(*float32)(unsafe.Pointer(_outs[0]))
	_g = *(*float32)(unsafe.Pointer(_outs[1]))
	_b = *(*float32)(unsafe.Pointer(_outs[2]))

	return _r, _g, _b
}

// RGBToHSV converts a color from RGB space to HSV.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
//
// The function takes the following parameters:
//
//    - r: red.
//    - g: green.
//    - b: blue.
//
// The function returns the following values:
//
//    - h: return value for the hue component.
//    - s: return value for the saturation component.
//    - v: return value for the value component.
//
func RGBToHSV(r, g, b float32) (h, s, v float32) {
	var _args [3]girepository.Argument
	var _outs [3]girepository.Argument

	*(*C.float)(unsafe.Pointer(&_args[0])) = C.float(r)
	*(*C.float)(unsafe.Pointer(&_args[1])) = C.float(g)
	*(*C.float)(unsafe.Pointer(&_args[2])) = C.float(b)

	_info := girepository.MustFind("Gtk", "rgb_to_hsv")
	_info.Invoke(_args[:], _outs[:])

	runtime.KeepAlive(r)
	runtime.KeepAlive(g)
	runtime.KeepAlive(b)

	var _h float32 // out
	var _s float32 // out
	var _v float32 // out

	_h = *(*float32)(unsafe.Pointer(_outs[0]))
	_s = *(*float32)(unsafe.Pointer(_outs[1]))
	_v = *(*float32)(unsafe.Pointer(_outs[2]))

	return _h, _s, _v
}
