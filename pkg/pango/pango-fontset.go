// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
//
// gboolean gotk4_FontsetForeachFunc(PangoFontset*, PangoFont*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fontset_get_type()), F: marshalFontset},
		{T: externglib.Type(C.pango_fontset_simple_get_type()), F: marshalFontsetSimple},
	})
}

// FontsetForeachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForeachFunc func(fontset Fontset, font Font, ok bool)

//export gotk4_FontsetForeachFunc
func gotk4_FontsetForeachFunc(arg0 *C.PangoFontset, arg1 *C.PangoFont, arg2 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var fontset Fontset // out
	var font Font       // out

	fontset = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(Fontset)
	font = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(Font)

	fn := v.(FontsetForeachFunc)
	ok := fn(fontset, font)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// Fontset: a `PangoFontset` represents a set of `PangoFont` to use when
// rendering text.
//
// A `PAngoFontset` is the result of resolving a `PangoFontDescription` against
// a particular `PangoContext`. It has operations for finding the component font
// for a particular Unicode character, and for finding a composite set of
// metrics for the entire fontset.
type Fontset interface {
	gextras.Objector

	ForeachFontset(fn FontsetForeachFunc)

	Font(wc uint) Font

	Metrics() *FontMetrics
}

// fontset implements the Fontset class.
type fontset struct {
	gextras.Objector
}

// WrapFontset wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontset(obj *externglib.Object) Fontset {
	return fontset{
		Objector: obj,
	}
}

func marshalFontset(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontset(obj), nil
}

func (f fontset) ForeachFontset(fn FontsetForeachFunc) {
	var _arg0 *C.PangoFontset           // out
	var _arg1 C.PangoFontsetForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(f.Native()))
	_arg1 = (*[0]byte)(C.gotk4_FontsetForeachFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.pango_fontset_foreach(_arg0, _arg1, _arg2)
}

func (f fontset) Font(wc uint) Font {
	var _arg0 *C.PangoFontset // out
	var _arg1 C.guint         // out
	var _cret *C.PangoFont    // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(f.Native()))
	_arg1 = C.guint(wc)

	_cret = C.pango_fontset_get_font(_arg0, _arg1)

	var _font Font // out

	_font = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Font)

	return _font
}

func (f fontset) Metrics() *FontMetrics {
	var _arg0 *C.PangoFontset     // out
	var _cret *C.PangoFontMetrics // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(f.Native()))

	_cret = C.pango_fontset_get_metrics(_arg0)

	var _fontMetrics *FontMetrics // out

	_fontMetrics = (*FontMetrics)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_fontMetrics, func(v **FontMetrics) {
		C.free(unsafe.Pointer(v))
	})

	return _fontMetrics
}

// FontsetSimple: `PangoFontsetSimple` is a implementation of the abstract
// `PangoFontset` base class as an array of fonts.
//
// When creating a `PangoFontsetSimple`, you have to provide the array of fonts
// that make up the fontset.
type FontsetSimple interface {
	Fontset

	AppendFontsetSimple(font Font)

	SizeFontsetSimple() int
}

// fontsetSimple implements the FontsetSimple class.
type fontsetSimple struct {
	Fontset
}

// WrapFontsetSimple wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontsetSimple(obj *externglib.Object) FontsetSimple {
	return fontsetSimple{
		Fontset: WrapFontset(obj),
	}
}

func marshalFontsetSimple(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontsetSimple(obj), nil
}

func NewFontsetSimple(language *Language) FontsetSimple {
	var _arg1 *C.PangoLanguage      // out
	var _cret *C.PangoFontsetSimple // in

	_arg1 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	_cret = C.pango_fontset_simple_new(_arg1)

	var _fontsetSimple FontsetSimple // out

	_fontsetSimple = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FontsetSimple)

	return _fontsetSimple
}

func (f fontsetSimple) AppendFontsetSimple(font Font) {
	var _arg0 *C.PangoFontsetSimple // out
	var _arg1 *C.PangoFont          // out

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))

	C.pango_fontset_simple_append(_arg0, _arg1)
}

func (f fontsetSimple) SizeFontsetSimple() int {
	var _arg0 *C.PangoFontsetSimple // out
	var _cret C.int                 // in

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(f.Native()))

	_cret = C.pango_fontset_simple_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
