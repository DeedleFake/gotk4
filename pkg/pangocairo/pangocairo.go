// Code generated by girgen. DO NOT EDIT.

package pangocairo

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangocairo pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pangocairo.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_cairo_font_get_type()), F: marshalFont},
		{T: externglib.Type(C.pango_cairo_font_map_get_type()), F: marshalFontMap},
	})
}

// ShapeRendererFunc: function type for rendering attributes of type
// PANGO_ATTR_SHAPE with Pango's Cairo renderer.
type ShapeRendererFunc func(cr *cairo.Context, attr *pango.AttrShape, doPath bool, data interface{})

//export gotk4_ShapeRendererFunc
func gotk4_ShapeRendererFunc(arg0 *C.cairo_t, arg1 *C.PangoAttrShape, arg2 C.gboolean, arg3 C.gpointer) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var cr *cairo.Context     // out
	var attr *pango.AttrShape // out
	var doPath bool           // out
	var data interface{}      // out

	cr = (*cairo.Context)(unsafe.Pointer(*C.cairo_t))
	attr = (*pango.AttrShape)(unsafe.Pointer(*C.PangoAttrShape))
	if arg2 != 0 {
		doPath = true
	}
	data = box.Get(uintptr(arg3))

	fn := v.(ShapeRendererFunc)
	fn(cr, attr, doPath, data)
}

// ContextGetFontOptions retrieves any font rendering options previously set
// with [func@PangoCairo.context_set_font_options].
//
// This function does not report options that are derived from the target
// surface by [func@update_context].
func ContextGetFontOptions(context pango.Context) *cairo.FontOptions {
	var _arg1 *C.PangoContext         // out
	var _cret *C.cairo_font_options_t // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer((&pango.Context).Native()))

	_cret = C.pango_cairo_context_get_font_options(_arg1)

	var _fontOptions *cairo.FontOptions // out

	_fontOptions = (*cairo.FontOptions)(unsafe.Pointer(*C.cairo_font_options_t))

	return _fontOptions
}

// ContextGetResolution gets the resolution for the context. See
// [func@PangoCairo.context_set_resolution]
func ContextGetResolution(context pango.Context) float64 {
	var _arg1 *C.PangoContext // out
	var _cret C.double        // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer((&pango.Context).Native()))

	_cret = C.pango_cairo_context_get_resolution(_arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ContextSetFontOptions sets the font options used when rendering text with
// this context.
//
// These options override any options that [func@update_context] derives from
// the target surface.
func ContextSetFontOptions(context pango.Context, options *cairo.FontOptions) {
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.cairo_font_options_t // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer((&pango.Context).Native()))
	_arg2 = (*C.cairo_font_options_t)(unsafe.Pointer(*cairo.FontOptions))

	C.pango_cairo_context_set_font_options(_arg1, _arg2)
}

// ContextSetResolution sets the resolution for the context.
//
// This is a scale factor between points specified in a `PangoFontDescription`
// and Cairo units. The default value is 96, meaning that a 10 point font will
// be 13 units high. (10 * 96. / 72. = 13.3).
func ContextSetResolution(context pango.Context, dpi float64) {
	var _arg1 *C.PangoContext // out
	var _arg2 C.double        // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer((&pango.Context).Native()))
	_arg2 = C.double(dpi)

	C.pango_cairo_context_set_resolution(_arg1, _arg2)
}

// CreateContext creates a context object set up to match the current
// transformation and target surface of the Cairo context.
//
// This context can then be used to create a layout using
// [ctor@Pango.Layout.new].
//
// This function is a convenience function that creates a context using the
// default font map, then updates it to @cr. If you just need to create a layout
// for use with @cr and do not need to access `PangoContext` directly, you can
// use [func@create_layout] instead.
func CreateContext(cr *cairo.Context) *pango.ContextClass {
	var _arg1 *C.cairo_t      // out
	var _cret *C.PangoContext // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))

	_cret = C.pango_cairo_create_context(_arg1)

	var _context *pango.ContextClass // out

	_context = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*pango.ContextClass)

	return _context
}

// CreateLayout creates a layout object set up to match the current
// transformation and target surface of the Cairo context.
//
// This layout can then be used for text measurement with functions like
// [method@Pango.Layout.get_size] or drawing with functions like
// [func@show_layout]. If you change the transformation or target surface for
// @cr, you need to call [func@update_layout].
//
// This function is the most convenient way to use Cairo with Pango, however it
// is slightly inefficient since it creates a separate `PangoContext` object for
// each layout. This might matter in an application that was laying out large
// amounts of text.
func CreateLayout(cr *cairo.Context) *pango.LayoutClass {
	var _arg1 *C.cairo_t     // out
	var _cret *C.PangoLayout // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))

	_cret = C.pango_cairo_create_layout(_arg1)

	var _layout *pango.LayoutClass // out

	_layout = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*pango.LayoutClass)

	return _layout
}

// ErrorUnderlinePath: add a squiggly line to the current path in the specified
// cairo context that approximately covers the given rectangle in the style of
// an underline used to indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ErrorUnderlinePath(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.cairo_t // out
	var _arg2 C.double   // out
	var _arg3 C.double   // out
	var _arg4 C.double   // out
	var _arg5 C.double   // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.pango_cairo_error_underline_path(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// GlyphStringPath adds the glyphs in @glyphs to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be at the
// current point of the cairo context.
func GlyphStringPath(cr *cairo.Context, font pango.Font, glyphs *pango.GlyphString) {
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.PangoFont        // out
	var _arg3 *C.PangoGlyphString // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = (*C.PangoFont)(unsafe.Pointer((&pango.Font).Native()))
	_arg3 = (*C.PangoGlyphString)(unsafe.Pointer(*pango.GlyphString))

	C.pango_cairo_glyph_string_path(_arg1, _arg2, _arg3)
}

// LayoutPath adds the text in a `PangoLayout` to the current path in the
// specified cairo context.
//
// The top-left corner of the `PangoLayout` will be at the current point of the
// cairo context.
func LayoutPath(cr *cairo.Context, layout pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer((&pango.Layout).Native()))

	C.pango_cairo_layout_path(_arg1, _arg2)
}

// ShowErrorUnderline: draw a squiggly line in the specified cairo context that
// approximately covers the given rectangle in the style of an underline used to
// indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ShowErrorUnderline(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.cairo_t // out
	var _arg2 C.double   // out
	var _arg3 C.double   // out
	var _arg4 C.double   // out
	var _arg5 C.double   // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.pango_cairo_show_error_underline(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// ShowGlyphItem draws the glyphs in @glyph_item in the specified cairo context,
//
// embedding the text associated with the glyphs in the output if the output
// format supports it (PDF for example), otherwise it acts similar to
// [func@show_glyph_string].
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
//
// Note that @text is the start of the text for layout, which is then indexed by
// `glyph_item->item->offset`.
func ShowGlyphItem(cr *cairo.Context, text string, glyphItem *pango.GlyphItem) {
	var _arg1 *C.cairo_t        // out
	var _arg2 *C.char           // out
	var _arg3 *C.PangoGlyphItem // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.PangoGlyphItem)(unsafe.Pointer(*pango.GlyphItem))

	C.pango_cairo_show_glyph_item(_arg1, _arg2, _arg3)
}

// ShowGlyphString draws the glyphs in @glyphs in the specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
func ShowGlyphString(cr *cairo.Context, font pango.Font, glyphs *pango.GlyphString) {
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.PangoFont        // out
	var _arg3 *C.PangoGlyphString // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = (*C.PangoFont)(unsafe.Pointer((&pango.Font).Native()))
	_arg3 = (*C.PangoGlyphString)(unsafe.Pointer(*pango.GlyphString))

	C.pango_cairo_show_glyph_string(_arg1, _arg2, _arg3)
}

// ShowLayout draws a `PangoLayout` in the specified cairo context.
//
// The top-left corner of the `PangoLayout` will be drawn at the current point
// of the cairo context.
func ShowLayout(cr *cairo.Context, layout pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer((&pango.Layout).Native()))

	C.pango_cairo_show_layout(_arg1, _arg2)
}

// UpdateContext updates a `PangoContext` previously created for use with Cairo
// to match the current transformation and target surface of a Cairo context.
//
// If any layouts have been created for the context, it's necessary to call
// [method@Pango.Layout.context_changed] on those layouts.
func UpdateContext(cr *cairo.Context, context pango.Context) {
	var _arg1 *C.cairo_t      // out
	var _arg2 *C.PangoContext // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = (*C.PangoContext)(unsafe.Pointer((&pango.Context).Native()))

	C.pango_cairo_update_context(_arg1, _arg2)
}

// UpdateLayout updates the private `PangoContext` of a `PangoLayout` created
// with [func@create_layout] to match the current transformation and target
// surface of a Cairo context.
func UpdateLayout(cr *cairo.Context, layout pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(*cairo.Context))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer((&pango.Layout).Native()))

	C.pango_cairo_update_layout(_arg1, _arg2)
}

// Font: `PangoCairoFont` is an interface exported by fonts for use with Cairo.
//
// The actual type of the font will depend on the particular font technology
// Cairo was compiled to use.
type Font interface {
	gextras.Objector

	// ScaledFont gets the `cairo_scaled_font_t` used by @font. The scaled font
	// can be referenced and kept using cairo_scaled_font_reference().
	ScaledFont() *cairo.ScaledFont
}

// FontInterface implements the Font interface.
type FontInterface struct {
	pango.FontClass
}

var _ Font = (*FontInterface)(nil)

func wrapFont(obj *externglib.Object) Font {
	return &FontInterface{
		FontClass: pango.FontClass{
			Object: obj,
		},
	}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFont(obj), nil
}

// ScaledFont gets the `cairo_scaled_font_t` used by @font. The scaled font can
// be referenced and kept using cairo_scaled_font_reference().
func (f *FontInterface) ScaledFont() *cairo.ScaledFont {
	var _arg0 *C.PangoCairoFont      // out
	var _cret *C.cairo_scaled_font_t // in

	_arg0 = (*C.PangoCairoFont)(unsafe.Pointer((&Font).Native()))

	_cret = C.pango_cairo_font_get_scaled_font(_arg0)

	var _scaledFont *cairo.ScaledFont // out

	_scaledFont = (*cairo.ScaledFont)(unsafe.Pointer(*C.cairo_scaled_font_t))

	return _scaledFont
}

// FontMap: `PangoCairoFontMap` is an interface exported by font maps for use
// with Cairo.
//
// The actual type of the font map will depend on the particular font technology
// Cairo was compiled to use.
type FontMap interface {
	gextras.Objector

	// FontType gets the type of Cairo font backend that @fontmap uses.
	FontType() cairo.FontType
	// Resolution gets the resolution for the fontmap.
	//
	// See [method@PangoCairo.FontMap.set_resolution].
	Resolution() float64
	// SetDefault sets a default `PangoCairoFontMap` to use with Cairo.
	//
	// This can be used to change the Cairo font backend that the default
	// fontmap uses for example. The old default font map is unreffed and the
	// new font map referenced.
	//
	// Note that since Pango 1.32.6, the default fontmap is per-thread. This
	// function only changes the default fontmap for the current thread. Default
	// fontmaps of existing threads are not changed. Default fontmaps of any new
	// threads will still be created using [type_func@PangoCairo.FontMap.new].
	//
	// A value of nil for @fontmap will cause the current default font map to be
	// released and a new default font map to be created on demand, using
	// [type_func@PangoCairo.FontMap.new].
	SetDefault()
	// SetResolution sets the resolution for the fontmap.
	//
	// This is a scale factor between points specified in a
	// `PangoFontDescription` and Cairo units. The default value is 96, meaning
	// that a 10 point font will be 13 units high. (10 * 96. / 72. = 13.3).
	SetResolution(dpi float64)
}

// FontMapInterface implements the FontMap interface.
type FontMapInterface struct {
	pango.FontMapClass
}

var _ FontMap = (*FontMapInterface)(nil)

func wrapFontMap(obj *externglib.Object) FontMap {
	return &FontMapInterface{
		FontMapClass: pango.FontMapClass{
			Object: obj,
		},
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontMap(obj), nil
}

// FontType gets the type of Cairo font backend that @fontmap uses.
func (f *FontMapInterface) FontType() cairo.FontType {
	var _arg0 *C.PangoCairoFontMap // out
	var _cret C.cairo_font_type_t  // in

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer((&FontMap).Native()))

	_cret = C.pango_cairo_font_map_get_font_type(_arg0)

	var _fontType cairo.FontType // out

	_fontType = (cairo.FontType)(C.cairo_font_type_t)

	return _fontType
}

// Resolution gets the resolution for the fontmap.
//
// See [method@PangoCairo.FontMap.set_resolution].
func (f *FontMapInterface) Resolution() float64 {
	var _arg0 *C.PangoCairoFontMap // out
	var _cret C.double             // in

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer((&FontMap).Native()))

	_cret = C.pango_cairo_font_map_get_resolution(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetDefault sets a default `PangoCairoFontMap` to use with Cairo.
//
// This can be used to change the Cairo font backend that the default fontmap
// uses for example. The old default font map is unreffed and the new font map
// referenced.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread. This
// function only changes the default fontmap for the current thread. Default
// fontmaps of existing threads are not changed. Default fontmaps of any new
// threads will still be created using [type_func@PangoCairo.FontMap.new].
//
// A value of nil for @fontmap will cause the current default font map to be
// released and a new default font map to be created on demand, using
// [type_func@PangoCairo.FontMap.new].
func (f *FontMapInterface) SetDefault() {
	var _arg0 *C.PangoCairoFontMap // out

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer((&FontMap).Native()))

	C.pango_cairo_font_map_set_default(_arg0)
}

// SetResolution sets the resolution for the fontmap.
//
// This is a scale factor between points specified in a `PangoFontDescription`
// and Cairo units. The default value is 96, meaning that a 10 point font will
// be 13 units high. (10 * 96. / 72. = 13.3).
func (f *FontMapInterface) SetResolution(dpi float64) {
	var _arg0 *C.PangoCairoFontMap // out
	var _arg1 C.double             // out

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer((&FontMap).Native()))
	_arg1 = C.double(dpi)

	C.pango_cairo_font_map_set_resolution(_arg0, _arg1)
}
