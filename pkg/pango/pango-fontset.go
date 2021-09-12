// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "pango.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fontset_get_type()), F: marshalFontsetter},
		{T: externglib.Type(C.pango_fontset_simple_get_type()), F: marshalFontsetSimpler},
	})
}

// FontsetForeachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForeachFunc func(fontset Fontsetter, font Fonter) (ok bool)

//export _gotk4_pango1_FontsetForeachFunc
func _gotk4_pango1_FontsetForeachFunc(arg0 *C.PangoFontset, arg1 *C.PangoFont, arg2 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var fontset Fontsetter // out
	var font Fonter        // out

	fontset = (externglib.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(Fontsetter)
	font = (externglib.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(Fonter)

	fn := v.(FontsetForeachFunc)
	ok := fn(fontset, font)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontsetOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FontsetOverrider interface {
	// Foreach iterates through all the fonts in a fontset, calling func for
	// each one.
	//
	// If func returns TRUE, that stops the iteration.
	Foreach(fn FontsetForeachFunc)
	// Font returns the font in the fontset that contains the best glyph for a
	// Unicode character.
	Font(wc uint) Fonter
	Language() *Language
	// Metrics: get overall metric information for the fonts in the fontset.
	Metrics() *FontMetrics
}

// Fontset: PangoFontset represents a set of PangoFont to use when rendering
// text.
//
// A PAngoFontset is the result of resolving a PangoFontDescription against a
// particular PangoContext. It has operations for finding the component font for
// a particular Unicode character, and for finding a composite set of metrics
// for the entire fontset.
type Fontset struct {
	*externglib.Object
}

// Fontsetter describes Fontset's abstract methods.
type Fontsetter interface {
	externglib.Objector

	// Foreach iterates through all the fonts in a fontset, calling func for
	// each one.
	Foreach(fn FontsetForeachFunc)
	// Font returns the font in the fontset that contains the best glyph for a
	// Unicode character.
	Font(wc uint) Fonter
	// Metrics: get overall metric information for the fonts in the fontset.
	Metrics() *FontMetrics
}

var _ Fontsetter = (*Fontset)(nil)

func wrapFontset(obj *externglib.Object) *Fontset {
	return &Fontset{
		Object: obj,
	}
}

func marshalFontsetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontset(obj), nil
}

// Foreach iterates through all the fonts in a fontset, calling func for each
// one.
//
// If func returns TRUE, that stops the iteration.
func (fontset *Fontset) Foreach(fn FontsetForeachFunc) {
	var _arg0 *C.PangoFontset           // out
	var _arg1 C.PangoFontsetForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))
	_arg1 = (*[0]byte)(C._gotk4_pango1_FontsetForeachFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.pango_fontset_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(fn)
}

// Font returns the font in the fontset that contains the best glyph for a
// Unicode character.
func (fontset *Fontset) Font(wc uint) Fonter {
	var _arg0 *C.PangoFontset // out
	var _arg1 C.guint         // out
	var _cret *C.PangoFont    // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))
	_arg1 = C.guint(wc)

	_cret = C.pango_fontset_get_font(_arg0, _arg1)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(wc)

	var _font Fonter // out

	_font = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Fonter)

	return _font
}

// Metrics: get overall metric information for the fonts in the fontset.
func (fontset *Fontset) Metrics() *FontMetrics {
	var _arg0 *C.PangoFontset     // out
	var _cret *C.PangoFontMetrics // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))

	_cret = C.pango_fontset_get_metrics(_arg0)
	runtime.KeepAlive(fontset)

	var _fontMetrics *FontMetrics // out

	_fontMetrics = (*FontMetrics)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_fontMetrics)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_font_metrics_unref((*C.PangoFontMetrics)(intern.C))
		},
	)

	return _fontMetrics
}

// FontsetSimple: PangoFontsetSimple is a implementation of the abstract
// PangoFontset base class as an array of fonts.
//
// When creating a PangoFontsetSimple, you have to provide the array of fonts
// that make up the fontset.
type FontsetSimple struct {
	Fontset
}

func wrapFontsetSimple(obj *externglib.Object) *FontsetSimple {
	return &FontsetSimple{
		Fontset: Fontset{
			Object: obj,
		},
	}
}

func marshalFontsetSimpler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontsetSimple(obj), nil
}

// NewFontsetSimple creates a new PangoFontsetSimple for the given language.
func NewFontsetSimple(language *Language) *FontsetSimple {
	var _arg1 *C.PangoLanguage      // out
	var _cret *C.PangoFontsetSimple // in

	_arg1 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_fontset_simple_new(_arg1)
	runtime.KeepAlive(language)

	var _fontsetSimple *FontsetSimple // out

	_fontsetSimple = wrapFontsetSimple(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fontsetSimple
}

// Append adds a font to the fontset.
func (fontset *FontsetSimple) Append(font Fonter) {
	var _arg0 *C.PangoFontsetSimple // out
	var _arg1 *C.PangoFont          // out

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(fontset.Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))

	C.pango_fontset_simple_append(_arg0, _arg1)
	runtime.KeepAlive(fontset)
	runtime.KeepAlive(font)
}

// Size returns the number of fonts in the fontset.
func (fontset *FontsetSimple) Size() int {
	var _arg0 *C.PangoFontsetSimple // out
	var _cret C.int                 // in

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(fontset.Native()))

	_cret = C.pango_fontset_simple_size(_arg0)
	runtime.KeepAlive(fontset)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
