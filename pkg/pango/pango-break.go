// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
import "C"

// Break determines possible line, word, and character breaks for a string of
// Unicode text with a single analysis.
//
// For most purposes you may want to use pango_get_log_attrs().
//
// Deprecated: Use pango_default_break() and pango_tailor_break().
func Break(text string, length int, analysis *Analysis, attrs []LogAttr) {
	var _arg1 *C.gchar         // out
	var _arg2 C.int            // out
	var _arg3 *C.PangoAnalysis // out
	var _arg4 *C.PangoLogAttr
	var _arg5 C.int

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg5 = (C.int)(len(attrs))
	if len(attrs) > 0 {
		_arg4 = (*C.PangoLogAttr)(unsafe.Pointer(&attrs[0]))
	}

	C.pango_break(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// DefaultBreak: this is the default break algorithm.
//
// It applies Unicode rules without language-specific tailoring, therefore the
// analyis argument is unused and can be NULL.
//
// See pango_tailor_break() for language-specific breaks.
func DefaultBreak(text string, length int, analysis *Analysis, attrs *LogAttr, attrsLen int) {
	var _arg1 *C.gchar         // out
	var _arg2 C.int            // out
	var _arg3 *C.PangoAnalysis // out
	var _arg4 *C.PangoLogAttr  // out
	var _arg5 C.int            // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg4 = (*C.PangoLogAttr)(gextras.StructNative(unsafe.Pointer(attrs)))
	_arg5 = C.int(attrsLen)

	C.pango_default_break(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// FindParagraphBoundary locates a paragraph boundary in text.
//
// A boundary is caused by delimiter characters, such as a newline, carriage
// return, carriage return-newline pair, or Unicode paragraph separator
// character. The index of the run of delimiters is returned in
// paragraph_delimiter_index. The index of the start of the paragrap (index
// after all delimiters) is stored in next_paragraph_start.
//
// If no delimiters are found, both paragraph_delimiter_index and
// next_paragraph_start are filled with the length of text (an index one off the
// end).
func FindParagraphBoundary(text string, length int) (paragraphDelimiterIndex int, nextParagraphStart int) {
	var _arg1 *C.gchar // out
	var _arg2 C.gint   // out
	var _arg3 C.gint   // in
	var _arg4 C.gint   // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	_arg2 = C.gint(length)

	C.pango_find_paragraph_boundary(_arg1, _arg2, &_arg3, &_arg4)

	var _paragraphDelimiterIndex int // out
	var _nextParagraphStart int      // out

	_paragraphDelimiterIndex = int(_arg3)
	_nextParagraphStart = int(_arg4)

	return _paragraphDelimiterIndex, _nextParagraphStart
}

// GetLogAttrs computes a PangoLogAttr for each character in text.
//
// The log_attrs array must have one PangoLogAttr for each position in text; if
// text contains N characters, it has N+1 positions, including the last position
// at the end of the text. text should be an entire paragraph; logical
// attributes can't be computed without context (for example you need to see
// spaces on either side of a word to know the word is a word).
func GetLogAttrs(text string, length int, level int, language *Language, logAttrs []LogAttr) {
	var _arg1 *C.char          // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out
	var _arg4 *C.PangoLanguage // out
	var _arg5 *C.PangoLogAttr
	var _arg6 C.int

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	_arg2 = C.int(length)
	_arg3 = C.int(level)
	_arg4 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))
	_arg6 = (C.int)(len(logAttrs))
	if len(logAttrs) > 0 {
		_arg5 = (*C.PangoLogAttr)(unsafe.Pointer(&logAttrs[0]))
	}

	C.pango_get_log_attrs(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// TailorBreak: apply language-specific tailoring to the breaks in log_attrs.
//
// The line breaks are assumed to have been produced by default_break.
//
// If offset is not -1, it is used to apply attributes from analysis that are
// relevant to line breaking.
func TailorBreak(text string, length int, analysis *Analysis, offset int, logAttrs []LogAttr) {
	var _arg1 *C.char          // out
	var _arg2 C.int            // out
	var _arg3 *C.PangoAnalysis // out
	var _arg4 C.int            // out
	var _arg5 *C.PangoLogAttr
	var _arg6 C.int

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg4 = C.int(offset)
	_arg6 = (C.int)(len(logAttrs))
	if len(logAttrs) > 0 {
		_arg5 = (*C.PangoLogAttr)(unsafe.Pointer(&logAttrs[0]))
	}

	C.pango_tailor_break(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// LogAttr: PangoLogAttr structure stores information about the attributes of a
// single character.
type LogAttr struct {
	nocopy gextras.NoCopy
	native *C.PangoLogAttr
}
