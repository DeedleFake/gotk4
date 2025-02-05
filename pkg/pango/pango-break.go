// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// Break determines possible line, word, and character breaks for a string of
// Unicode text with a single analysis.
//
// For most purposes you may want to use pango_get_log_attrs().
//
// Deprecated: Use pango_default_break() and pango_tailor_break().
//
// The function takes the following parameters:
//
//    - text to process. Must be valid UTF-8.
//    - length of text in bytes (may be -1 if text is nul-terminated).
//    - analysis structure from pango_itemize().
//    - attrs: array to store character information in.
//
func Break(text string, length int, analysis *Analysis, attrs []LogAttr) {
	var _arg1 *C.gchar         // out
	var _arg2 C.int            // out
	var _arg3 *C.PangoAnalysis // out
	var _arg4 *C.PangoLogAttr  // out
	var _arg5 C.int

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	_arg5 = (C.int)(len(attrs))
	_arg4 = (*C.PangoLogAttr)(C.calloc(C.size_t(len(attrs)), C.size_t(C.sizeof_PangoLogAttr)))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice((*C.PangoLogAttr)(_arg4), len(attrs))
		for i := range attrs {
			out[i] = *(*C.PangoLogAttr)(gextras.StructNative(unsafe.Pointer((&attrs[i]))))
		}
	}

	C.pango_break(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(attrs)
}

// DefaultBreak: this is the default break algorithm.
//
// It applies Unicode rules without language-specific tailoring, therefore the
// analyis argument is unused and can be NULL.
//
// See pango_tailor_break() for language-specific breaks.
//
// The function takes the following parameters:
//
//    - text to break. Must be valid UTF-8.
//    - length of text in bytes (may be -1 if text is nul-terminated).
//    - analysis (optional) for the text.
//    - attrs: logical attributes to fill in.
//    - attrsLen: size of the array passed as attrs.
//
func DefaultBreak(text string, length int, analysis *Analysis, attrs *LogAttr, attrsLen int) {
	var _arg1 *C.gchar         // out
	var _arg2 C.int            // out
	var _arg3 *C.PangoAnalysis // out
	var _arg4 *C.PangoLogAttr  // out
	var _arg5 C.int            // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	if analysis != nil {
		_arg3 = (*C.PangoAnalysis)(gextras.StructNative(unsafe.Pointer(analysis)))
	}
	_arg4 = (*C.PangoLogAttr)(gextras.StructNative(unsafe.Pointer(attrs)))
	_arg5 = C.int(attrsLen)

	C.pango_default_break(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(attrs)
	runtime.KeepAlive(attrsLen)
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
//
// The function takes the following parameters:
//
//    - text: UTF-8 text.
//    - length of text in bytes, or -1 if nul-terminated.
//
// The function returns the following values:
//
//    - paragraphDelimiterIndex: return location for index of delimiter.
//    - nextParagraphStart: return location for start of next paragraph.
//
func FindParagraphBoundary(text string, length int) (paragraphDelimiterIndex, nextParagraphStart int) {
	var _arg1 *C.gchar // out
	var _arg2 C.gint   // out
	var _arg3 C.gint   // in
	var _arg4 C.gint   // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)

	C.pango_find_paragraph_boundary(_arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)

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
//
// The function takes the following parameters:
//
//    - text to process. Must be valid UTF-8.
//    - length in bytes of text.
//    - level: embedding level, or -1 if unknown.
//    - language tag.
//    - logAttrs: array with one PangoLogAttr per character in text, plus one
//      extra, to be filled in.
//
func GetLogAttrs(text string, length, level int, language *Language, logAttrs []LogAttr) {
	var _arg1 *C.char          // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out
	var _arg4 *C.PangoLanguage // out
	var _arg5 *C.PangoLogAttr  // out
	var _arg6 C.int

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	_arg3 = C.int(level)
	_arg4 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))
	_arg6 = (C.int)(len(logAttrs))
	_arg5 = (*C.PangoLogAttr)(C.calloc(C.size_t(len(logAttrs)), C.size_t(C.sizeof_PangoLogAttr)))
	defer C.free(unsafe.Pointer(_arg5))
	{
		out := unsafe.Slice((*C.PangoLogAttr)(_arg5), len(logAttrs))
		for i := range logAttrs {
			out[i] = *(*C.PangoLogAttr)(gextras.StructNative(unsafe.Pointer((&logAttrs[i]))))
		}
	}

	C.pango_get_log_attrs(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
	runtime.KeepAlive(level)
	runtime.KeepAlive(language)
	runtime.KeepAlive(logAttrs)
}

// LogAttr: PangoLogAttr structure stores information about the attributes of a
// single character.
//
// An instance of this type is always passed by reference.
type LogAttr struct {
	*logAttr
}

// logAttr is the struct that's finalized.
type logAttr struct {
	native *C.PangoLogAttr
}
