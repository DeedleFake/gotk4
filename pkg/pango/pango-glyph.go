// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"runtime"
	"strings"
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
		{T: externglib.Type(C.pango_shape_flags_get_type()), F: marshalShapeFlags},
		{T: externglib.Type(C.pango_glyph_string_get_type()), F: marshalGlyphString},
	})
}

// GlyphUnit: PangoGlyphUnit type is used to store dimensions within Pango.
//
// Dimensions are stored in 1/PANGO_SCALE of a device unit. (A device unit might
// be a pixel for screen display, or a point on a printer.) PANGO_SCALE is
// currently 1024, and may change in the future (unlikely though), but you
// should not depend on its exact value. The PANGO_PIXELS() macro can be used to
// convert from glyph units into device units with correct rounding.
type GlyphUnit = int32

// ShapeFlags flags influencing the shaping process.
//
// PangoShapeFlags can be passed to pango_shape_with_flags().
type ShapeFlags int

const (
	// ShapeNone: default value.
	ShapeNone ShapeFlags = 0b0
	// ShapeRoundPositions: round glyph positions and widths to whole device
	// units. This option should be set if the target renderer can't do subpixel
	// positioning of glyphs.
	ShapeRoundPositions ShapeFlags = 0b1
)

func marshalShapeFlags(p uintptr) (interface{}, error) {
	return ShapeFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for ShapeFlags.
func (s ShapeFlags) String() string {
	if s == 0 {
		return "ShapeFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(29)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case ShapeNone:
			builder.WriteString("None|")
		case ShapeRoundPositions:
			builder.WriteString("RoundPositions|")
		default:
			builder.WriteString(fmt.Sprintf("ShapeFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Shape: convert the characters in text into glyphs.
//
// Given a segment of text and the corresponding PangoAnalysis structure
// returned from itemize, convert the characters into glyphs. You may also pass
// in only a substring of the item from itemize.
//
// It is recommended that you use shape_full instead, since that API allows for
// shaping interaction happening across text item boundaries.
//
// Note that the extra attributes in the analyis that is returned from itemize
// have indices that are relative to the entire paragraph, so you need to
// subtract the item offset from their indices before calling shape.
func Shape(text string, length int, analysis *Analysis, glyphs *GlyphString) {
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.PangoAnalysis    // out
	var _arg4 *C.PangoGlyphString // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg4 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_shape(_arg1, _arg2, _arg3, _arg4)
}

// ShapeFull: convert the characters in text into glyphs.
//
// Given a segment of text and the corresponding PangoAnalysis structure
// returned from itemize, convert the characters into glyphs. You may also pass
// in only a substring of the item from itemize.
//
// This is similar to shape, except it also can optionally take the full
// paragraph text as input, which will then be used to perform certain
// cross-item shaping interactions. If you have access to the broader text of
// which item_text is part of, provide the broader text as paragraph_text. If
// paragraph_text is NULL, item text is used instead.
//
// Note that the extra attributes in the analyis that is returned from itemize
// have indices that are relative to the entire paragraph, so you do not pass
// the full paragraph text as paragraph_text, you need to subtract the item
// offset from their indices before calling shape_full.
func ShapeFull(itemText string, itemLength int, paragraphText string, paragraphLength int, analysis *Analysis, glyphs *GlyphString) {
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.char             // out
	var _arg4 C.int               // out
	var _arg5 *C.PangoAnalysis    // out
	var _arg6 *C.PangoGlyphString // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(itemText)))
	_arg2 = C.int(itemLength)
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(paragraphText)))
	_arg4 = C.int(paragraphLength)
	_arg5 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg6 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_shape_full(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// ShapeWithFlags: convert the characters in text into glyphs.
//
// Given a segment of text and the corresponding PangoAnalysis structure
// returned from itemize, convert the characters into glyphs. You may also pass
// in only a substring of the item from itemize.
//
// This is similar to shape_full, except it also takes flags that can influence
// the shaping process.
//
// Note that the extra attributes in the analyis that is returned from itemize
// have indices that are relative to the entire paragraph, so you do not pass
// the full paragraph text as paragraph_text, you need to subtract the item
// offset from their indices before calling shape_with_flags.
func ShapeWithFlags(itemText string, itemLength int, paragraphText string, paragraphLength int, analysis *Analysis, glyphs *GlyphString, flags ShapeFlags) {
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.char             // out
	var _arg4 C.int               // out
	var _arg5 *C.PangoAnalysis    // out
	var _arg6 *C.PangoGlyphString // out
	var _arg7 C.PangoShapeFlags   // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(itemText)))
	_arg2 = C.int(itemLength)
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(paragraphText)))
	_arg4 = C.int(paragraphLength)
	_arg5 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg6 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg7 = C.PangoShapeFlags(flags)

	C.pango_shape_with_flags(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// GlyphGeometry: PangoGlyphGeometry structure contains width and positioning
// information for a single glyph.
type GlyphGeometry struct {
	nocopy gextras.NoCopy
	native *C.PangoGlyphGeometry
}

// GlyphInfo: PangoGlyphInfo structure represents a single glyph with
// positioning information and visual attributes.
type GlyphInfo struct {
	nocopy gextras.NoCopy
	native *C.PangoGlyphInfo
}

// GlyphString: PangoGlyphString is used to store strings of glyphs with
// geometry and visual attribute information.
//
// The storage for the glyph information is owned by the structure which
// simplifies memory management.
type GlyphString struct {
	nocopy gextras.NoCopy
	native *C.PangoGlyphString
}

func marshalGlyphString(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &GlyphString{native: (*C.PangoGlyphString)(unsafe.Pointer(b))}, nil
}

// NewGlyphString constructs a struct GlyphString.
func NewGlyphString() *GlyphString {
	var _cret *C.PangoGlyphString // in

	_cret = C.pango_glyph_string_new()

	var _glyphString *GlyphString // out

	_glyphString = (*GlyphString)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_glyphString, func(v *GlyphString) {
		C.pango_glyph_string_free((*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _glyphString
}

// Copy a glyph string and associated storage.
func (_string *GlyphString) Copy() *GlyphString {
	var _arg0 *C.PangoGlyphString // out
	var _cret *C.PangoGlyphString // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(_string)))

	_cret = C.pango_glyph_string_copy(_arg0)

	var _glyphString *GlyphString // out

	_glyphString = (*GlyphString)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_glyphString, func(v *GlyphString) {
		C.pango_glyph_string_free((*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _glyphString
}

// Extents: compute the logical and ink extents of a glyph string.
//
// See the documentation for pango.Font.GetGlyphExtents() for details about the
// interpretation of the rectangles.
//
// Examples of logical (red) and ink (green) rects:
//
// ! (rects1.png) ! (rects2.png)
func (glyphs *GlyphString) Extents(font Fonter) (inkRect Rectangle, logicalRect Rectangle) {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 *C.PangoFont        // out
	var _arg2 C.PangoRectangle    // in
	var _arg3 C.PangoRectangle    // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg1 = (*C.PangoFont)(unsafe.Pointer((font).(gextras.Nativer).Native()))

	C.pango_glyph_string_extents(_arg0, _arg1, &_arg2, &_arg3)

	var _inkRect Rectangle     // out
	var _logicalRect Rectangle // out

	_inkRect = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_logicalRect = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _inkRect, _logicalRect
}

// ExtentsRange computes the extents of a sub-portion of a glyph string.
//
// The extents are relative to the start of the glyph string range (the origin
// of their coordinate system is at the start of the range, not at the start of
// the entire glyph string).
func (glyphs *GlyphString) ExtentsRange(start int, end int, font Fonter) (inkRect Rectangle, logicalRect Rectangle) {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 C.int               // out
	var _arg2 C.int               // out
	var _arg3 *C.PangoFont        // out
	var _arg4 C.PangoRectangle    // in
	var _arg5 C.PangoRectangle    // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg1 = C.int(start)
	_arg2 = C.int(end)
	_arg3 = (*C.PangoFont)(unsafe.Pointer((font).(gextras.Nativer).Native()))

	C.pango_glyph_string_extents_range(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5)

	var _inkRect Rectangle     // out
	var _logicalRect Rectangle // out

	_inkRect = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))
	_logicalRect = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg5))))

	return _inkRect, _logicalRect
}

// Free: free a glyph string and associated storage.
func (_string *GlyphString) free() {
	var _arg0 *C.PangoGlyphString // out

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(_string)))

	C.pango_glyph_string_free(_arg0)
}

// Width computes the logical width of the glyph string.
//
// This can also be computed using pango.GlyphString.Extents(). However, since
// this only computes the width, it's much faster. This is in fact only a
// convenience function that computes the sum of geometry.width for each glyph
// in the glyphs.
func (glyphs *GlyphString) Width() int {
	var _arg0 *C.PangoGlyphString // out
	var _cret C.int               // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	_cret = C.pango_glyph_string_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IndexToX converts from character position to x position.
//
// The X position is measured from the left edge of the run. Character positions
// are computed by dividing up each cluster into equal portions.
func (glyphs *GlyphString) IndexToX(text string, length int, analysis *Analysis, index_ int, trailing bool) int {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.PangoAnalysis    // out
	var _arg4 C.int               // out
	var _arg5 C.gboolean          // out
	var _arg6 C.int               // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg4 = C.int(index_)
	if trailing {
		_arg5 = C.TRUE
	}

	C.pango_glyph_string_index_to_x(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_arg6)

	var _xPos int // out

	_xPos = int(_arg6)

	return _xPos
}

// SetSize: resize a glyph string to the given length.
func (_string *GlyphString) SetSize(newLen int) {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 C.gint              // out

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(_string)))
	_arg1 = C.gint(newLen)

	C.pango_glyph_string_set_size(_arg0, _arg1)
}

// XToIndex: convert from x offset to character position.
//
// Character positions are computed by dividing up each cluster into equal
// portions. In scripts where positioning within a cluster is not allowed (such
// as Thai), the returned value may not be a valid cursor position; the caller
// must combine the result with the logical attributes for the text to compute
// the valid cursor position.
func (glyphs *GlyphString) XToIndex(text string, length int, analysis *Analysis, xPos int) (index_ int, trailing int) {
	var _arg0 *C.PangoGlyphString // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _arg3 *C.PangoAnalysis    // out
	var _arg4 C.int               // out
	var _arg5 C.int               // in
	var _arg6 C.int               // in

	_arg0 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg4 = C.int(xPos)

	C.pango_glyph_string_x_to_index(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_arg6)

	var _index_ int   // out
	var _trailing int // out

	_index_ = int(_arg5)
	_trailing = int(_arg6)

	return _index_, _trailing
}

// GlyphVisAttr: PangoGlyphVisAttr structure communicates information between
// the shaping and rendering phases.
//
// Currently, it contains only cluster start information. yMore attributes may
// be added in the future.
type GlyphVisAttr struct {
	nocopy gextras.NoCopy
	native *C.PangoGlyphVisAttr
}
