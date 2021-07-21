// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_context_get_type()), F: marshalContexter},
	})
}

// Itemize breaks a piece of text into segments with consistent directional
// level and font.
//
// Each byte of text will be contained in exactly one of the items in the
// returned list; the generated list of items will be in logical order (the
// start offsets of the items are ascending).
//
// cached_iter should be an iterator over attrs currently positioned at a range
// before or containing start_index; cached_iter will be advanced to the range
// covering the position just after start_index + length. (i.e. if itemizing in
// a loop, just keep passing in the same cached_iter).
func Itemize(context *Context, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) *externglib.List {
	var _arg1 *C.PangoContext      // out
	var _arg2 *C.char              // out
	var _arg3 C.int                // out
	var _arg4 C.int                // out
	var _arg5 *C.PangoAttrList     // out
	var _arg6 *C.PangoAttrIterator // out
	var _cret *C.GList             // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	_arg3 = C.int(startIndex)
	_arg4 = C.int(length)
	_arg5 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attrs)))
	_arg6 = (*C.PangoAttrIterator)(gextras.StructNative(unsafe.Pointer(cachedIter)))

	_cret = C.pango_itemize(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.PangoItem)(_p)
		var dst Item // out
		dst = *(*Item)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(&dst, func(v *Item) {
			C.pango_item_free((*C.PangoItem)(gextras.StructNative(unsafe.Pointer(v))))
		})
		return dst
	})
	_list.AttachFinalizer(func(v uintptr) {
		C.pango_item_free((*C.PangoItem)(unsafe.Pointer(v)))
	})

	return _list
}

// ItemizeWithBaseDir: like pango_itemize(), but with an explicitly specified
// base direction.
//
// The base direction is used when computing bidirectional levels. (see
// pango.Context.SetBaseDir()). itemize gets the base direction from the
// PangoContext.
func ItemizeWithBaseDir(context *Context, baseDir Direction, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) *externglib.List {
	var _arg1 *C.PangoContext      // out
	var _arg2 C.PangoDirection     // out
	var _arg3 *C.char              // out
	var _arg4 C.int                // out
	var _arg5 C.int                // out
	var _arg6 *C.PangoAttrList     // out
	var _arg7 *C.PangoAttrIterator // out
	var _cret *C.GList             // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.PangoDirection(baseDir)
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(text)))
	_arg4 = C.int(startIndex)
	_arg5 = C.int(length)
	_arg6 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attrs)))
	_arg7 = (*C.PangoAttrIterator)(gextras.StructNative(unsafe.Pointer(cachedIter)))

	_cret = C.pango_itemize_with_base_dir(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.PangoItem)(_p)
		var dst Item // out
		dst = *(*Item)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(&dst, func(v *Item) {
			C.pango_item_free((*C.PangoItem)(gextras.StructNative(unsafe.Pointer(v))))
		})
		return dst
	})
	_list.AttachFinalizer(func(v uintptr) {
		C.pango_item_free((*C.PangoItem)(unsafe.Pointer(v)))
	})

	return _list
}

// Context: PangoContext stores global information used to control the
// itemization process.
//
// The information stored by `PangoContext includes the fontmap used to look up
// fonts, and default values such as the default language, default gravity, or
// default font.
//
// To obtain a PangoContext, use pango.FontMap.CreateContext().
type Context struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Context)(nil)

func wrapContext(obj *externglib.Object) *Context {
	return &Context{
		Object: obj,
	}
}

func marshalContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapContext(obj), nil
}

// NewContext creates a new PangoContext initialized to default values.
//
// This function is not particularly useful as it should always be followed by a
// pango.Context.SetFontMap() call, and the function
// pango.FontMap.CreateContext() does these two steps together and hence users
// are recommended to use that.
//
// If you are using Pango as part of a higher-level system, that system may have
// it's own way of create a PangoContext. For instance, the GTK toolkit has,
// among others, gtk_widget_get_pango_context(). Use those instead.
func NewContext() *Context {
	var _cret *C.PangoContext // in

	_cret = C.pango_context_new()

	var _context *Context // out

	_context = wrapContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _context
}

// Changed forces a change in the context, which will cause any PangoLayout
// using this context to re-layout.
//
// This function is only useful when implementing a new backend for Pango,
// something applications won't do. Backends should call this function if they
// have attached extra data to the context and such data is changed.
func (context *Context) Changed() {
	var _arg0 *C.PangoContext // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	C.pango_context_changed(_arg0)
}

// BaseDir retrieves the base direction for the context.
//
// See pango.Context.SetBaseDir().
func (context *Context) BaseDir() Direction {
	var _arg0 *C.PangoContext  // out
	var _cret C.PangoDirection // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_base_dir(_arg0)

	var _direction Direction // out

	_direction = Direction(_cret)

	return _direction
}

// BaseGravity retrieves the base gravity for the context.
//
// See pango.Context.SetBaseGravity().
func (context *Context) BaseGravity() Gravity {
	var _arg0 *C.PangoContext // out
	var _cret C.PangoGravity  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_base_gravity(_arg0)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// FontDescription: retrieve the default font description for the context.
func (context *Context) FontDescription() *FontDescription {
	var _arg0 *C.PangoContext         // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_font_description(_arg0)

	var _fontDescription *FontDescription // out

	_fontDescription = (*FontDescription)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _fontDescription
}

// FontMap gets the PangoFontMap used to look up fonts for this context.
func (context *Context) FontMap() FontMapper {
	var _arg0 *C.PangoContext // out
	var _cret *C.PangoFontMap // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_font_map(_arg0)

	var _fontMap FontMapper // out

	_fontMap = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(FontMapper)

	return _fontMap
}

// Gravity retrieves the gravity for the context.
//
// This is similar to pango.Context.GetBaseGravity(), except for when the base
// gravity is PANGO_GRAVITY_AUTO for which pango.Gravity.GetForMatrix is used to
// return the gravity from the current context matrix.
func (context *Context) Gravity() Gravity {
	var _arg0 *C.PangoContext // out
	var _cret C.PangoGravity  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_gravity(_arg0)

	var _gravity Gravity // out

	_gravity = Gravity(_cret)

	return _gravity
}

// GravityHint retrieves the gravity hint for the context.
//
// See pango.Context.SetGravityHint() for details.
func (context *Context) GravityHint() GravityHint {
	var _arg0 *C.PangoContext    // out
	var _cret C.PangoGravityHint // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_gravity_hint(_arg0)

	var _gravityHint GravityHint // out

	_gravityHint = GravityHint(_cret)

	return _gravityHint
}

// Language retrieves the global language tag for the context.
func (context *Context) Language() *Language {
	var _arg0 *C.PangoContext  // out
	var _cret *C.PangoLanguage // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_language(_arg0)

	var _language *Language // out

	_language = (*Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_language, func(v *Language) {
		C.free(gextras.StructNative(unsafe.Pointer(v)))
	})

	return _language
}

// Matrix gets the transformation matrix that will be applied when rendering
// with this context.
//
// See pango.Context.SetMatrix().
func (context *Context) Matrix() *Matrix {
	var _arg0 *C.PangoContext // out
	var _cret *C.PangoMatrix  // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_matrix(_arg0)

	var _matrix *Matrix // out

	_matrix = (*Matrix)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _matrix
}

// Metrics: get overall metric information for a particular font description.
//
// Since the metrics may be substantially different for different scripts, a
// language tag can be provided to indicate that the metrics should be retrieved
// that correspond to the script(s) used by that language.
//
// The PangoFontDescription is interpreted in the same way as by itemize, and
// the family name may be a comma separated list of names. If characters from
// multiple of these families would be used to render the string, then the
// returned fonts would be a composite of the metrics for the fonts loaded for
// the individual families.
func (context *Context) Metrics(desc *FontDescription, language *Language) *FontMetrics {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 *C.PangoLanguage        // out
	var _cret *C.PangoFontMetrics     // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))
	_arg2 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_context_get_metrics(_arg0, _arg1, _arg2)

	var _fontMetrics *FontMetrics // out

	_fontMetrics = (*FontMetrics)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.pango_font_metrics_ref(_cret)
	runtime.SetFinalizer(_fontMetrics, func(v *FontMetrics) {
		C.pango_font_metrics_unref((*C.PangoFontMetrics)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _fontMetrics
}

// RoundGlyphPositions returns whether font rendering with this context should
// round glyph positions and widths.
func (context *Context) RoundGlyphPositions() bool {
	var _arg0 *C.PangoContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_round_glyph_positions(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Serial returns the current serial number of context.
//
// The serial number is initialized to an small number larger than zero when a
// new context is created and is increased whenever the context is changed using
// any of the setter functions, or the PangoFontMap it uses to find fonts has
// changed. The serial may wrap, but will never have the value 0. Since it can
// wrap, never compare it with "less than", always use "not equals".
//
// This can be used to automatically detect changes to a PangoContext, and is
// only useful when implementing objects that need update when their
// PangoContext changes, like PangoLayout.
func (context *Context) Serial() uint {
	var _arg0 *C.PangoContext // out
	var _cret C.guint         // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_context_get_serial(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ListFamilies: list all families for a context.
func (context *Context) ListFamilies() []FontFamilier {
	var _arg0 *C.PangoContext     // out
	var _arg1 **C.PangoFontFamily // in
	var _arg2 C.int               // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	C.pango_context_list_families(_arg0, &_arg1, &_arg2)

	var _families []FontFamilier // out

	defer C.free(unsafe.Pointer(_arg1))
	{
		src := unsafe.Slice(_arg1, _arg2)
		_families = make([]FontFamilier, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_families[i] = (gextras.CastObject(externglib.Take(unsafe.Pointer(src[i])))).(FontFamilier)
		}
	}

	return _families
}

// LoadFont loads the font in one of the fontmaps in the context that is the
// closest match for desc.
func (context *Context) LoadFont(desc *FontDescription) Fonter {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _cret *C.PangoFont            // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))

	_cret = C.pango_context_load_font(_arg0, _arg1)

	var _font Fonter // out

	_font = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Fonter)

	return _font
}

// LoadFontset: load a set of fonts in the context that can be used to render a
// font matching desc.
func (context *Context) LoadFontset(desc *FontDescription, language *Language) Fontsetter {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 *C.PangoLanguage        // out
	var _cret *C.PangoFontset         // in

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))
	_arg2 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_context_load_fontset(_arg0, _arg1, _arg2)

	var _fontset Fontsetter // out

	_fontset = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Fontsetter)

	return _fontset
}

// SetBaseDir sets the base direction for the context.
//
// The base direction is used in applying the Unicode bidirectional algorithm;
// if the direction is PANGO_DIRECTION_LTR or PANGO_DIRECTION_RTL, then the
// value will be used as the paragraph direction in the Unicode bidirectional
// algorithm. A value of PANGO_DIRECTION_WEAK_LTR or PANGO_DIRECTION_WEAK_RTL is
// used only for paragraphs that do not contain any strong characters
// themselves.
func (context *Context) SetBaseDir(direction Direction) {
	var _arg0 *C.PangoContext  // out
	var _arg1 C.PangoDirection // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.PangoDirection(direction)

	C.pango_context_set_base_dir(_arg0, _arg1)
}

// SetBaseGravity sets the base gravity for the context.
//
// The base gravity is used in laying vertical text out.
func (context *Context) SetBaseGravity(gravity Gravity) {
	var _arg0 *C.PangoContext // out
	var _arg1 C.PangoGravity  // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.PangoGravity(gravity)

	C.pango_context_set_base_gravity(_arg0, _arg1)
}

// SetFontDescription: set the default font description for the context
func (context *Context) SetFontDescription(desc *FontDescription) {
	var _arg0 *C.PangoContext         // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontDescription)(gextras.StructNative(unsafe.Pointer(desc)))

	C.pango_context_set_font_description(_arg0, _arg1)
}

// SetFontMap sets the font map to be searched when fonts are looked-up in this
// context.
//
// This is only for internal use by Pango backends, a PangoContext obtained via
// one of the recommended methods should already have a suitable font map.
func (context *Context) SetFontMap(fontMap FontMapper) {
	var _arg0 *C.PangoContext // out
	var _arg1 *C.PangoFontMap // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoFontMap)(unsafe.Pointer((fontMap).(gextras.Nativer).Native()))

	C.pango_context_set_font_map(_arg0, _arg1)
}

// SetGravityHint sets the gravity hint for the context.
//
// The gravity hint is used in laying vertical text out, and is only relevant if
// gravity of the context as returned by pango.Context.GetGravity() is set to
// PANGO_GRAVITY_EAST or PANGO_GRAVITY_WEST.
func (context *Context) SetGravityHint(hint GravityHint) {
	var _arg0 *C.PangoContext    // out
	var _arg1 C.PangoGravityHint // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.PangoGravityHint(hint)

	C.pango_context_set_gravity_hint(_arg0, _arg1)
}

// SetLanguage sets the global language tag for the context.
//
// The default language for the locale of the running process can be found using
// pango.Language.GetDefault.
func (context *Context) SetLanguage(language *Language) {
	var _arg0 *C.PangoContext  // out
	var _arg1 *C.PangoLanguage // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	C.pango_context_set_language(_arg0, _arg1)
}

// SetMatrix sets the transformation matrix that will be applied when rendering
// with this context.
//
// Note that reported metrics are in the user space coordinates before the
// application of the matrix, not device-space coordinates after the application
// of the matrix. So, they don't scale with the matrix, though they may change
// slightly for different matrices, depending on how the text is fit to the
// pixel grid.
func (context *Context) SetMatrix(matrix *Matrix) {
	var _arg0 *C.PangoContext // out
	var _arg1 *C.PangoMatrix  // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))

	C.pango_context_set_matrix(_arg0, _arg1)
}

// SetRoundGlyphPositions sets whether font rendering with this context should
// round glyph positions and widths to integral positions, in device units.
//
// This is useful when the renderer can't handle subpixel positioning of glyphs.
//
// The default value is to round glyph positions, to remain compatible with
// previous Pango behavior.
func (context *Context) SetRoundGlyphPositions(roundPositions bool) {
	var _arg0 *C.PangoContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	if roundPositions {
		_arg1 = C.TRUE
	}

	C.pango_context_set_round_glyph_positions(_arg0, _arg1)
}
