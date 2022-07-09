// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeRGBA returns the GType for the type RGBA.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRGBA() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "RGBA").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRGBA)
	return gtype
}

// RGBA: GdkRGBA is used to represent a color, in a way that is compatible with
// cairo’s notion of color.
//
// GdkRGBA is a convenient way to pass colors around. It’s based on cairo’s way
// to deal with colors and mirrors its behavior. All values are in the range
// from 0.0 to 1.0 inclusive. So the color (0.0, 0.0, 0.0, 0.0) represents
// transparent black and (1.0, 1.0, 1.0, 1.0) is opaque white. Other values will
// be clamped to this range when drawing.
//
// An instance of this type is always passed by reference.
type RGBA struct {
	*rgbA
}

// rgbA is the struct that's finalized.
type rgbA struct {
	native unsafe.Pointer
}

func marshalRGBA(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &RGBA{&rgbA{(unsafe.Pointer)(b)}}, nil
}

// Red: intensity of the red channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Red() float32 {
	offset := girepository.MustFind("Gdk", "RGBA").StructFieldOffset("red")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	var v float32 // out
	v = float32(*(*C.float)(unsafe.Pointer(&*valptr)))
	return v
}

// Green: intensity of the green channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Green() float32 {
	offset := girepository.MustFind("Gdk", "RGBA").StructFieldOffset("green")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	var v float32 // out
	v = float32(*(*C.float)(unsafe.Pointer(&*valptr)))
	return v
}

// Blue: intensity of the blue channel from 0.0 to 1.0 inclusive.
func (r *RGBA) Blue() float32 {
	offset := girepository.MustFind("Gdk", "RGBA").StructFieldOffset("blue")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	var v float32 // out
	v = float32(*(*C.float)(unsafe.Pointer(&*valptr)))
	return v
}

// Alpha: opacity of the color from 0.0 for completely translucent to 1.0 for
// opaque.
func (r *RGBA) Alpha() float32 {
	offset := girepository.MustFind("Gdk", "RGBA").StructFieldOffset("alpha")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	var v float32 // out
	v = float32(*(*C.float)(unsafe.Pointer(&*valptr)))
	return v
}

// Red: intensity of the red channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetRed(red float32) {
	offset := girepository.MustFind("Gdk", "RGBA").StructFieldOffset("red")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	*(*C.float)(unsafe.Pointer(&*valptr)) = C.float(red)
}

// Green: intensity of the green channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetGreen(green float32) {
	offset := girepository.MustFind("Gdk", "RGBA").StructFieldOffset("green")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	*(*C.float)(unsafe.Pointer(&*valptr)) = C.float(green)
}

// Blue: intensity of the blue channel from 0.0 to 1.0 inclusive.
func (r *RGBA) SetBlue(blue float32) {
	offset := girepository.MustFind("Gdk", "RGBA").StructFieldOffset("blue")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	*(*C.float)(unsafe.Pointer(&*valptr)) = C.float(blue)
}

// Alpha: opacity of the color from 0.0 for completely translucent to 1.0 for
// opaque.
func (r *RGBA) SetAlpha(alpha float32) {
	offset := girepository.MustFind("Gdk", "RGBA").StructFieldOffset("alpha")
	valptr := (*uintptr)(unsafe.Add(r.native, offset))
	*(*C.float)(unsafe.Pointer(&*valptr)) = C.float(alpha)
}

// Copy makes a copy of a GdkRGBA.
//
// The result must be freed through gdk.RGBA.Free().
//
// The function returns the following values:
//
//    - rgbA: newly allocated GdkRGBA, with the same contents as rgba.
//
func (rgba *RGBA) Copy() *RGBA {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))

	_info := girepository.MustFind("Gdk", "RGBA")
	_gret := _info.InvokeRecordMethod("copy", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(rgba)

	var _rgbA *RGBA // out

	_rgbA = (*RGBA)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_rgbA)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				args := [1]girepository.Argument{(*C.void)(intern.C)}
				girepository.MustFind("Gdk", "RGBA").InvokeRecordMethod("free", args[:], nil)
			}
		},
	)

	return _rgbA
}

// Equal compares two GdkRGBA colors.
//
// The function takes the following parameters:
//
//    - p2: another GdkRGBA.
//
// The function returns the following values:
//
//    - ok: TRUE if the two colors compare equal.
//
func (p1 *RGBA) Equal(p2 *RGBA) bool {
	var _args [2]girepository.Argument

	*(*C.gpointer)(unsafe.Pointer(&_args[0])) = C.gpointer(gextras.StructNative(unsafe.Pointer(p1)))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = C.gpointer(gextras.StructNative(unsafe.Pointer(p2)))

	_info := girepository.MustFind("Gdk", "RGBA")
	_gret := _info.InvokeRecordMethod("equal", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Hash: hash function suitable for using for a hash table that stores GdkRGBAs.
//
// The function returns the following values:
//
//    - guint: hash value for p.
//
func (p *RGBA) Hash() uint32 {
	var _args [1]girepository.Argument

	*(*C.gpointer)(unsafe.Pointer(&_args[0])) = C.gpointer(gextras.StructNative(unsafe.Pointer(p)))

	_info := girepository.MustFind("Gdk", "RGBA")
	_gret := _info.InvokeRecordMethod("hash", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(p)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// IsClear checks if an rgba value is transparent.
//
// That is, drawing with the value would not produce any change.
//
// The function returns the following values:
//
//    - ok: TRUE if the rgba is clear.
//
func (rgba *RGBA) IsClear() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))

	_info := girepository.MustFind("Gdk", "RGBA")
	_gret := _info.InvokeRecordMethod("is_clear", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(rgba)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsOpaque checks if an rgba value is opaque.
//
// That is, drawing with the value will not retain any results from previous
// contents.
//
// The function returns the following values:
//
//    - ok: TRUE if the rgba is opaque.
//
func (rgba *RGBA) IsOpaque() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))

	_info := girepository.MustFind("Gdk", "RGBA")
	_gret := _info.InvokeRecordMethod("is_opaque", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(rgba)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Parse parses a textual representation of a color.
//
// The string can be either one of:
//
// - A standard name (Taken from the X11 rgb.txt file).
//
// - A hexadecimal value in the form “\#rgb”, “\#rrggbb”, “\#rrrgggbbb” or
// ”\#rrrrggggbbbb”
//
// - A hexadecimal value in the form “\#rgba”, “\#rrggbbaa”, or
// ”\#rrrrggggbbbbaaaa”
//
// - A RGB color in the form “rgb(r,g,b)” (In this case the color will have full
// opacity)
//
// - A RGBA color in the form “rgba(r,g,b,a)”
//
// Where “r”, “g”, “b” and “a” are respectively the red, green, blue and alpha
// color values. In the last two cases, “r”, “g”, and “b” are either integers in
// the range 0 to 255 or percentage values in the range 0% to 100%, and a is a
// floating point value in the range 0 to 1.
//
// The function takes the following parameters:
//
//    - spec: string specifying the color.
//
// The function returns the following values:
//
//    - ok: TRUE if the parsing succeeded.
//
func (rgba *RGBA) Parse(spec string) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(spec)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gdk", "RGBA")
	_gret := _info.InvokeRecordMethod("parse", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(rgba)
	runtime.KeepAlive(spec)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// String returns a textual specification of rgba in the form rgb(r,g,b) or
// rgba(r,g,b,a), where “r”, “g”, “b” and “a” represent the red, green, blue and
// alpha values respectively. “r”, “g”, and “b” are represented as integers in
// the range 0 to 255, and “a” is represented as a floating point value in the
// range 0 to 1.
//
// These string forms are string forms that are supported by the CSS3 colors
// module, and can be parsed by gdk.RGBA.Parse().
//
// Note that this string representation may lose some precision, since “r”, “g”
// and “b” are represented as 8-bit integers. If this is a concern, you should
// use a different representation.
//
// The function returns the following values:
//
//    - utf8: newly allocated text string.
//
func (rgba *RGBA) String() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rgba)))

	_info := girepository.MustFind("Gdk", "RGBA")
	_gret := _info.InvokeRecordMethod("to_string", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(rgba)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
