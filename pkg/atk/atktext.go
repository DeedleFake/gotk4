// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_text_attribute_get_type()), F: marshalTextAttribute},
		{T: externglib.Type(C.atk_text_boundary_get_type()), F: marshalTextBoundary},
		{T: externglib.Type(C.atk_text_clip_type_get_type()), F: marshalTextClipType},
		{T: externglib.Type(C.atk_text_granularity_get_type()), F: marshalTextGranularity},
		{T: externglib.Type(C.atk_text_get_type()), F: marshalText},
		{T: externglib.Type(C.atk_text_range_get_type()), F: marshalTextRange},
	})
}

// TextAttribute describes the text attributes supported
type TextAttribute int

const (
	// Invalid: invalid attribute, like bad spelling or grammar.
	TextAttributeInvalid TextAttribute = iota
	// LeftMargin: the pixel width of the left margin
	TextAttributeLeftMargin
	// RightMargin: the pixel width of the right margin
	TextAttributeRightMargin
	// Indent: the number of pixels that the text is indented
	TextAttributeIndent
	// Invisible: either "true" or "false" indicating whether text is visible or
	// not
	TextAttributeInvisible
	// Editable: either "true" or "false" indicating whether text is editable or
	// not
	TextAttributeEditable
	// PixelsAboveLines pixels of blank space to leave above each
	// newline-terminated line.
	TextAttributePixelsAboveLines
	// PixelsBelowLines pixels of blank space to leave below each
	// newline-terminated line.
	TextAttributePixelsBelowLines
	// PixelsInsideWrap pixels of blank space to leave between wrapped lines
	// inside the same newline-terminated line (paragraph).
	TextAttributePixelsInsideWrap
	// BgFullHeight: "true" or "false" whether to make the background color for
	// each character the height of the highest font used on the current line,
	// or the height of the font used for the current character.
	TextAttributeBgFullHeight
	// Rise: number of pixels that the characters are risen above the baseline.
	// See also ATK_TEXT_ATTR_TEXT_POSITION.
	TextAttributeRise
	// Underline: "none", "single", "double", "low", or "error"
	TextAttributeUnderline
	// Strikethrough: "true" or "false" whether the text is strikethrough
	TextAttributeStrikethrough
	// Size: the size of the characters in points. eg: 10
	TextAttributeSize
	// Scale: the scale of the characters. The value is a string representation
	// of a double
	TextAttributeScale
	// Weight: the weight of the characters.
	TextAttributeWeight
	// Language: the language used
	TextAttributeLanguage
	// FamilyName: the font family name
	TextAttributeFamilyName
	// BgColor: the background color. The value is an RGB value of the format
	// "u,u,u"
	TextAttributeBgColor
	// FgColor: the foreground color. The value is an RGB value of the format
	// "u,u,u"
	TextAttributeFgColor
	// BgStipple: "true" if a Bitmap is set for stippling the background color.
	TextAttributeBgStipple
	// FgStipple: "true" if a Bitmap is set for stippling the foreground color.
	TextAttributeFgStipple
	// WrapMode: the wrap mode of the text, if any. Values are "none", "char",
	// "word", or "word_char".
	TextAttributeWrapMode
	// Direction: the direction of the text, if set. Values are "none", "ltr" or
	// "rtl"
	TextAttributeDirection
	// Justification: the justification of the text, if set. Values are "left",
	// "right", "center" or "fill"
	TextAttributeJustification
	// Stretch: the stretch of the text, if set. Values are "ultra_condensed",
	// "extra_condensed", "condensed", "semi_condensed", "normal",
	// "semi_expanded", "expanded", "extra_expanded" or "ultra_expanded"
	TextAttributeStretch
	// Variant: the capitalization variant of the text, if set. Values are
	// "normal" or "small_caps"
	TextAttributeVariant
	// Style: the slant style of the text, if set. Values are "normal",
	// "oblique" or "italic"
	TextAttributeStyle
	// TextPosition: the vertical position with respect to the baseline. Values
	// are "baseline", "super", or "sub". Note that a super or sub text
	// attribute refers to position with respect to the baseline of the prior
	// character.
	TextAttributeTextPosition
	// LastDefined: not a valid text attribute, used for finding end of
	// enumeration
	TextAttributeLastDefined
)

func marshalTextAttribute(p uintptr) (interface{}, error) {
	return TextAttribute(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TextBoundary: text boundary types used for specifying boundaries for regions
// of text. This enumeration is deprecated since 2.9.4 and should not be used.
// Use AtkTextGranularity with #atk_text_get_string_at_offset instead.
type TextBoundary int

const (
	// Char: boundary is the boundary between characters (including non-printing
	// characters)
	TextBoundaryChar TextBoundary = iota
	// WordStart: boundary is the start (i.e. first character) of a word.
	TextBoundaryWordStart
	// WordEnd: boundary is the end (i.e. last character) of a word.
	TextBoundaryWordEnd
	// SentenceStart: boundary is the first character in a sentence.
	TextBoundarySentenceStart
	// SentenceEnd: boundary is the last (terminal) character in a sentence; in
	// languages which use "sentence stop" punctuation such as English, the
	// boundary is thus the '.', '?', or similar terminal punctuation character.
	TextBoundarySentenceEnd
	// LineStart: boundary is the initial character of the content or a
	// character immediately following a newline, linefeed, or return character.
	TextBoundaryLineStart
	// LineEnd: boundary is the linefeed, or return character.
	TextBoundaryLineEnd
)

func marshalTextBoundary(p uintptr) (interface{}, error) {
	return TextBoundary(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TextClipType describes the type of clipping required.
type TextClipType int

const (
	// None: no clipping to be done
	TextClipTypeNone TextClipType = iota
	// Min: text clipped by min coordinate is omitted
	TextClipTypeMin
	// Max: text clipped by max coordinate is omitted
	TextClipTypeMax
	// Both: only text fully within mix/max bound is retained
	TextClipTypeBoth
)

func marshalTextClipType(p uintptr) (interface{}, error) {
	return TextClipType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TextGranularity: text granularity types used for specifying the granularity
// of the region of text we are interested in.
type TextGranularity int

const (
	// Char: granularity is defined by the boundaries between characters
	// (including non-printing characters)
	TextGranularityChar TextGranularity = iota
	// Word: granularity is defined by the boundaries of a word, starting at the
	// beginning of the current word and finishing at the beginning of the
	// following one, if present.
	TextGranularityWord
	// Sentence: granularity is defined by the boundaries of a sentence,
	// starting at the beginning of the current sentence and finishing at the
	// beginning of the following one, if present.
	TextGranularitySentence
	// Line: granularity is defined by the boundaries of a line, starting at the
	// beginning of the current line and finishing at the beginning of the
	// following one, if present.
	TextGranularityLine
	// Paragraph: granularity is defined by the boundaries of a paragraph,
	// starting at the beginning of the current paragraph and finishing at the
	// beginning of the following one, if present.
	TextGranularityParagraph
)

func marshalTextGranularity(p uintptr) (interface{}, error) {
	return TextGranularity(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TextOverrider interface {
	// AddSelection adds a selection bounded by the specified offsets.
	AddSelection(startOffset int, endOffset int) bool
	// CaretOffset gets the offset of the position of the caret (cursor).
	CaretOffset() int
	// CharacterAtOffset gets the specified text.
	CharacterAtOffset(offset int) uint32
	// CharacterCount gets the character count.
	CharacterCount() int
	// NSelections gets the number of selected regions.
	NSelections() int
	// Selection gets the text from the specified selection.
	Selection(selectionNum int) (startOffset int, endOffset int, utf8 string)
	// Text gets the specified text.
	Text(startOffset int, endOffset int) string
	// RemoveSelection removes the specified selection.
	RemoveSelection(selectionNum int) bool
	// SetCaretOffset sets the caret (cursor) position to the specified @offset.
	//
	// In the case of rich-text content, this method should either grab focus or
	// move the sequential focus navigation starting point (if the application
	// supports this concept) as if the user had clicked on the new caret
	// position. Typically, this means that the target of this operation is the
	// node containing the new caret position or one of its ancestors. In other
	// words, after this method is called, if the user advances focus, it should
	// move to the first focusable node following the new caret position.
	//
	// Calling this method should also scroll the application viewport in a way
	// that matches the behavior of the application's typical caret motion or
	// tab navigation as closely as possible. This also means that if the
	// application's caret motion or focus navigation does not trigger a scroll
	// operation, this method should not trigger one either. If the application
	// does not have a caret motion or focus navigation operation, this method
	// should try to scroll the new caret position into view while minimizing
	// unnecessary scroll motion.
	SetCaretOffset(offset int) bool
	// SetSelection changes the start and end offset of the specified selection.
	SetSelection(selectionNum int, startOffset int, endOffset int) bool
	TextAttributesChanged()
	TextCaretMoved(location int)
	TextChanged(position int, length int)
	TextSelectionChanged()
}

// Text should be implemented by Objects on behalf of widgets that have text
// content which is either attributed or otherwise non-trivial. Objects whose
// text content is simple, unattributed, and very brief may expose that content
// via #atk_object_get_name instead; however if the text is editable,
// multi-line, typically longer than three or four words, attributed,
// selectable, or if the object already uses the 'name' ATK property for other
// information, the Text interface should be used to expose the text content. In
// the case of editable text content, EditableText (a subtype of the Text
// interface) should be implemented instead.
//
//    Text provides not only traversal facilities and change
//
// notification for text content, but also caret tracking and glyph bounding box
// calculations. Note that the text strings are exposed as UTF-8, and are
// therefore potentially multi-byte, and caret-to-byte offset mapping makes no
// assumptions about the character length; also bounding box glyph-to-offset
// mapping may be complex for languages which use ligatures.
type Text interface {
	gextras.Objector

	// AddSelection adds a selection bounded by the specified offsets.
	AddSelection(startOffset int, endOffset int) bool
	// CaretOffset gets the offset of the position of the caret (cursor).
	CaretOffset() int
	// CharacterAtOffset gets the specified text.
	CharacterAtOffset(offset int) uint32
	// CharacterCount gets the character count.
	CharacterCount() int
	// NSelections gets the number of selected regions.
	NSelections() int
	// Selection gets the text from the specified selection.
	Selection(selectionNum int) (startOffset int, endOffset int, utf8 string)
	// Text gets the specified text.
	Text(startOffset int, endOffset int) string
	// RemoveSelection removes the specified selection.
	RemoveSelection(selectionNum int) bool
	// SetCaretOffset sets the caret (cursor) position to the specified @offset.
	//
	// In the case of rich-text content, this method should either grab focus or
	// move the sequential focus navigation starting point (if the application
	// supports this concept) as if the user had clicked on the new caret
	// position. Typically, this means that the target of this operation is the
	// node containing the new caret position or one of its ancestors. In other
	// words, after this method is called, if the user advances focus, it should
	// move to the first focusable node following the new caret position.
	//
	// Calling this method should also scroll the application viewport in a way
	// that matches the behavior of the application's typical caret motion or
	// tab navigation as closely as possible. This also means that if the
	// application's caret motion or focus navigation does not trigger a scroll
	// operation, this method should not trigger one either. If the application
	// does not have a caret motion or focus navigation operation, this method
	// should try to scroll the new caret position into view while minimizing
	// unnecessary scroll motion.
	SetCaretOffset(offset int) bool
	// SetSelection changes the start and end offset of the specified selection.
	SetSelection(selectionNum int, startOffset int, endOffset int) bool
}

// TextIface implements the Text interface.
type TextIface struct {
	*externglib.Object
}

var _ Text = (*TextIface)(nil)

func wrapText(obj *externglib.Object) Text {
	return &TextIface{
		Object: obj,
	}
}

func marshalText(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapText(obj), nil
}

// AddSelection adds a selection bounded by the specified offsets.
func (text *TextIface) AddSelection(startOffset int, endOffset int) bool {
	var _arg0 *C.AtkText // out
	var _arg1 C.gint     // out
	var _arg2 C.gint     // out
	var _cret C.gboolean // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(startOffset)
	_arg2 = C.gint(endOffset)

	_cret = C.atk_text_add_selection(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CaretOffset gets the offset of the position of the caret (cursor).
func (text *TextIface) CaretOffset() int {
	var _arg0 *C.AtkText // out
	var _cret C.gint     // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))

	_cret = C.atk_text_get_caret_offset(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// CharacterAtOffset gets the specified text.
func (text *TextIface) CharacterAtOffset(offset int) uint32 {
	var _arg0 *C.AtkText // out
	var _arg1 C.gint     // out
	var _cret C.gunichar // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(offset)

	_cret = C.atk_text_get_character_at_offset(_arg0, _arg1)

	var _gunichar uint32 // out

	_gunichar = uint32(_cret)

	return _gunichar
}

// CharacterCount gets the character count.
func (text *TextIface) CharacterCount() int {
	var _arg0 *C.AtkText // out
	var _cret C.gint     // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))

	_cret = C.atk_text_get_character_count(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NSelections gets the number of selected regions.
func (text *TextIface) NSelections() int {
	var _arg0 *C.AtkText // out
	var _cret C.gint     // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))

	_cret = C.atk_text_get_n_selections(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Selection gets the text from the specified selection.
func (text *TextIface) Selection(selectionNum int) (startOffset int, endOffset int, utf8 string) {
	var _arg0 *C.AtkText // out
	var _arg1 C.gint     // out
	var _arg2 C.gint     // in
	var _arg3 C.gint     // in
	var _cret *C.gchar   // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(selectionNum)

	_cret = C.atk_text_get_selection(_arg0, _arg1, &_arg2, &_arg3)

	var _startOffset int // out
	var _endOffset int   // out
	var _utf8 string     // out

	_startOffset = int(_arg2)
	_endOffset = int(_arg3)
	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _startOffset, _endOffset, _utf8
}

// Text gets the specified text.
func (text *TextIface) Text(startOffset int, endOffset int) string {
	var _arg0 *C.AtkText // out
	var _arg1 C.gint     // out
	var _arg2 C.gint     // out
	var _cret *C.gchar   // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(startOffset)
	_arg2 = C.gint(endOffset)

	_cret = C.atk_text_get_text(_arg0, _arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RemoveSelection removes the specified selection.
func (text *TextIface) RemoveSelection(selectionNum int) bool {
	var _arg0 *C.AtkText // out
	var _arg1 C.gint     // out
	var _cret C.gboolean // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(selectionNum)

	_cret = C.atk_text_remove_selection(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCaretOffset sets the caret (cursor) position to the specified @offset.
//
// In the case of rich-text content, this method should either grab focus or
// move the sequential focus navigation starting point (if the application
// supports this concept) as if the user had clicked on the new caret position.
// Typically, this means that the target of this operation is the node
// containing the new caret position or one of its ancestors. In other words,
// after this method is called, if the user advances focus, it should move to
// the first focusable node following the new caret position.
//
// Calling this method should also scroll the application viewport in a way that
// matches the behavior of the application's typical caret motion or tab
// navigation as closely as possible. This also means that if the application's
// caret motion or focus navigation does not trigger a scroll operation, this
// method should not trigger one either. If the application does not have a
// caret motion or focus navigation operation, this method should try to scroll
// the new caret position into view while minimizing unnecessary scroll motion.
func (text *TextIface) SetCaretOffset(offset int) bool {
	var _arg0 *C.AtkText // out
	var _arg1 C.gint     // out
	var _cret C.gboolean // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(offset)

	_cret = C.atk_text_set_caret_offset(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSelection changes the start and end offset of the specified selection.
func (text *TextIface) SetSelection(selectionNum int, startOffset int, endOffset int) bool {
	var _arg0 *C.AtkText // out
	var _arg1 C.gint     // out
	var _arg2 C.gint     // out
	var _arg3 C.gint     // out
	var _cret C.gboolean // in

	_arg0 = (*C.AtkText)(unsafe.Pointer(text.Native()))
	_arg1 = C.gint(selectionNum)
	_arg2 = C.gint(startOffset)
	_arg3 = C.gint(endOffset)

	_cret = C.atk_text_set_selection(_arg0, _arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TextRange: structure used to describe a text range.
type TextRange struct {
	native C.AtkTextRange
}

// WrapTextRange wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTextRange(ptr unsafe.Pointer) *TextRange {
	return (*TextRange)(ptr)
}

func marshalTextRange(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*TextRange)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TextRange) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// TextRectangle: structure used to store a rectangle used by AtkText.
type TextRectangle struct {
	native C.AtkTextRectangle
}

// WrapTextRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTextRectangle(ptr unsafe.Pointer) *TextRectangle {
	return (*TextRectangle)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TextRectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
