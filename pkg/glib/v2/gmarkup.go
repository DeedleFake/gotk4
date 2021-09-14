// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"runtime/cgo"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_markup_parse_context_get_type()), F: marshalMarkupParseContext},
	})
}

// MarkupError: error codes returned by markup parsing.
type MarkupError int

const (
	// MarkupErrorBadUTF8: text being parsed was not valid UTF-8
	MarkupErrorBadUTF8 MarkupError = iota
	// MarkupErrorEmpty: document contained nothing, or only whitespace
	MarkupErrorEmpty
	// MarkupErrorParse: document was ill-formed
	MarkupErrorParse
	// MarkupErrorUnknownElement: error should be set by Parser functions;
	// element wasn't known
	MarkupErrorUnknownElement
	// MarkupErrorUnknownAttribute: error should be set by Parser functions;
	// attribute wasn't known
	MarkupErrorUnknownAttribute
	// MarkupErrorInvalidContent: error should be set by Parser functions;
	// content was invalid
	MarkupErrorInvalidContent
	// MarkupErrorMissingAttribute: error should be set by Parser functions; a
	// required attribute was missing
	MarkupErrorMissingAttribute
)

// String returns the name in string for MarkupError.
func (m MarkupError) String() string {
	switch m {
	case MarkupErrorBadUTF8:
		return "BadUTF8"
	case MarkupErrorEmpty:
		return "Empty"
	case MarkupErrorParse:
		return "Parse"
	case MarkupErrorUnknownElement:
		return "UnknownElement"
	case MarkupErrorUnknownAttribute:
		return "UnknownAttribute"
	case MarkupErrorInvalidContent:
		return "InvalidContent"
	case MarkupErrorMissingAttribute:
		return "MissingAttribute"
	default:
		return fmt.Sprintf("MarkupError(%d)", m)
	}
}

// MarkupCollectType: mixed enumerated type and flags field. You must specify
// one type (string, strdup, boolean, tristate). Additionally, you may
// optionally bitwise OR the type with the flag G_MARKUP_COLLECT_OPTIONAL.
//
// It is likely that this enum will be extended in the future to support other
// types.
type MarkupCollectType int

const (
	// MarkupCollectInvalid: used to terminate the list of attributes to collect
	MarkupCollectInvalid MarkupCollectType = 0b0
	// MarkupCollectString: collect the string pointer directly from the
	// attribute_values[] array. Expects a parameter of type (const char **). If
	// G_MARKUP_COLLECT_OPTIONAL is specified and the attribute isn't present
	// then the pointer will be set to NULL
	MarkupCollectString MarkupCollectType = 0b1
	// MarkupCollectStrdup as with G_MARKUP_COLLECT_STRING, but expects a
	// parameter of type (char **) and g_strdup()s the returned pointer. The
	// pointer must be freed with g_free()
	MarkupCollectStrdup MarkupCollectType = 0b10
	// MarkupCollectBoolean expects a parameter of type (gboolean *) and parses
	// the attribute value as a boolean. Sets FALSE if the attribute isn't
	// present. Valid boolean values consist of (case-insensitive) "false", "f",
	// "no", "n", "0" and "true", "t", "yes", "y", "1"
	MarkupCollectBoolean MarkupCollectType = 0b11
	// MarkupCollectTristate as with G_MARKUP_COLLECT_BOOLEAN, but in the case
	// of a missing attribute a value is set that compares equal to neither
	// FALSE nor TRUE G_MARKUP_COLLECT_OPTIONAL is implied
	MarkupCollectTristate MarkupCollectType = 0b100
	// MarkupCollectOptional: can be bitwise ORed with the other fields. If
	// present, allows the attribute not to appear. A default value is set
	// depending on what value type is used
	MarkupCollectOptional MarkupCollectType = 0b10000000000000000
)

// String returns the names in string for MarkupCollectType.
func (m MarkupCollectType) String() string {
	if m == 0 {
		return "MarkupCollectType(0)"
	}

	var builder strings.Builder
	builder.Grow(125)

	for m != 0 {
		next := m & (m - 1)
		bit := m - next

		switch bit {
		case MarkupCollectInvalid:
			builder.WriteString("Invalid|")
		case MarkupCollectString:
			builder.WriteString("String|")
		case MarkupCollectStrdup:
			builder.WriteString("Strdup|")
		case MarkupCollectBoolean:
			builder.WriteString("Boolean|")
		case MarkupCollectTristate:
			builder.WriteString("Tristate|")
		case MarkupCollectOptional:
			builder.WriteString("Optional|")
		default:
			builder.WriteString(fmt.Sprintf("MarkupCollectType(0b%b)|", bit))
		}

		m = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if m contains other.
func (m MarkupCollectType) Has(other MarkupCollectType) bool {
	return (m & other) == other
}

// MarkupParseFlags flags that affect the behaviour of the parser.
type MarkupParseFlags int

const (
	// MarkupDoNotUseThisUnsupportedFlag: flag you should not use
	MarkupDoNotUseThisUnsupportedFlag MarkupParseFlags = 0b1
	// MarkupTreatCdataAsText: when this flag is set, CDATA marked sections are
	// not passed literally to the passthrough function of the parser. Instead,
	// the content of the section (without the <![CDATA[ and ]]>) is passed to
	// the text function. This flag was added in GLib 2.12
	MarkupTreatCdataAsText MarkupParseFlags = 0b10
	// MarkupPrefixErrorPosition: normally errors caught by GMarkup itself have
	// line/column information prefixed to them to let the caller know the
	// location of the error. When this flag is set the location information is
	// also prefixed to errors generated by the Parser implementation functions
	MarkupPrefixErrorPosition MarkupParseFlags = 0b100
	// MarkupIgnoreQualified: ignore (don't report) qualified attributes and
	// tags, along with their contents. A qualified attribute or tag is one that
	// contains ':' in its name (ie: is in another namespace). Since: 2.40.
	MarkupIgnoreQualified MarkupParseFlags = 0b1000
)

// String returns the names in string for MarkupParseFlags.
func (m MarkupParseFlags) String() string {
	if m == 0 {
		return "MarkupParseFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(104)

	for m != 0 {
		next := m & (m - 1)
		bit := m - next

		switch bit {
		case MarkupDoNotUseThisUnsupportedFlag:
			builder.WriteString("DoNotUseThisUnsupportedFlag|")
		case MarkupTreatCdataAsText:
			builder.WriteString("TreatCdataAsText|")
		case MarkupPrefixErrorPosition:
			builder.WriteString("PrefixErrorPosition|")
		case MarkupIgnoreQualified:
			builder.WriteString("IgnoreQualified|")
		default:
			builder.WriteString(fmt.Sprintf("MarkupParseFlags(0b%b)|", bit))
		}

		m = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if m contains other.
func (m MarkupParseFlags) Has(other MarkupParseFlags) bool {
	return (m & other) == other
}

// MarkupEscapeText escapes text so that the markup parser will parse it
// verbatim. Less than, greater than, ampersand, etc. are replaced with the
// corresponding entities. This function would typically be used when writing
// out a file to be parsed with the markup parser.
//
// Note that this function doesn't protect whitespace and line endings from
// being processed according to the XML rules for normalization of line endings
// and attribute values.
//
// Note also that this function will produce character references in the range
// of &#x1; ... &#x1f; for all control sequences except for tabstop, newline and
// carriage return. The character references in this range are not valid XML
// 1.0, but they are valid XML 1.1 and will be accepted by the GMarkup parser.
func MarkupEscapeText(text string, length int) string {
	var _arg1 *C.gchar // out
	var _arg2 C.gssize // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.g_markup_escape_text(_arg1, _arg2)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// MarkupParseContext: parse context is used to parse a stream of bytes that you
// expect to contain marked-up text.
//
// See g_markup_parse_context_new(), Parser, and so on for more details.
//
// An instance of this type is always passed by reference.
type MarkupParseContext struct {
	*markupParseContext
}

// markupParseContext is the struct that's finalized.
type markupParseContext struct {
	native *C.GMarkupParseContext
}

func marshalMarkupParseContext(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &MarkupParseContext{&markupParseContext{(*C.GMarkupParseContext)(unsafe.Pointer(b))}}, nil
}

// EndParse signals to the ParseContext that all data has been fed into the
// parse context with g_markup_parse_context_parse().
//
// This function reports an error if the document isn't complete, for example if
// elements are still open.
func (context *MarkupParseContext) EndParse() error {
	var _arg0 *C.GMarkupParseContext // out
	var _cerr *C.GError              // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	C.g_markup_parse_context_end_parse(_arg0, &_cerr)
	runtime.KeepAlive(context)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Element retrieves the name of the currently open element.
//
// If called from the start_element or end_element handlers this will give the
// element_name as passed to those functions. For the parent elements, see
// g_markup_parse_context_get_element_stack().
func (context *MarkupParseContext) Element() string {
	var _arg0 *C.GMarkupParseContext // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	_cret = C.g_markup_parse_context_get_element(_arg0)
	runtime.KeepAlive(context)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Position retrieves the current line number and the number of the character on
// that line. Intended for use in error messages; there are no strict semantics
// for what constitutes the "current" line number other than "the best number we
// could come up with for error messages."
func (context *MarkupParseContext) Position() (lineNumber int, charNumber int) {
	var _arg0 *C.GMarkupParseContext // out
	var _arg1 C.gint                 // in
	var _arg2 C.gint                 // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	C.g_markup_parse_context_get_position(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

	var _lineNumber int // out
	var _charNumber int // out

	_lineNumber = int(_arg1)
	_charNumber = int(_arg2)

	return _lineNumber, _charNumber
}

// UserData returns the user_data associated with context.
//
// This will either be the user_data that was provided to
// g_markup_parse_context_new() or to the most recent call of
// g_markup_parse_context_push().
func (context *MarkupParseContext) UserData() cgo.Handle {
	var _arg0 *C.GMarkupParseContext // out
	var _cret C.gpointer             // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	_cret = C.g_markup_parse_context_get_user_data(_arg0)
	runtime.KeepAlive(context)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// Parse: feed some data to the ParseContext.
//
// The data need not be valid UTF-8; an error will be signaled if it's invalid.
// The data need not be an entire document; you can feed a document into the
// parser incrementally, via multiple calls to this function. Typically, as you
// receive data from a network connection or file, you feed each received chunk
// of data into this function, aborting the process if an error occurs. Once an
// error is reported, no further data may be fed to the ParseContext; all errors
// are fatal.
func (context *MarkupParseContext) Parse(text string, textLen int) error {
	var _arg0 *C.GMarkupParseContext // out
	var _arg1 *C.gchar               // out
	var _arg2 C.gssize               // out
	var _cerr *C.GError              // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(textLen)

	C.g_markup_parse_context_parse(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(context)
	runtime.KeepAlive(text)
	runtime.KeepAlive(textLen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Pop completes the process of a temporary sub-parser redirection.
//
// This function exists to collect the user_data allocated by a matching call to
// g_markup_parse_context_push(). It must be called in the end_element handler
// corresponding to the start_element handler during which
// g_markup_parse_context_push() was called. You must not call this function
// from the error callback -- the user_data is provided directly to the callback
// in that case.
//
// This function is not intended to be directly called by users interested in
// invoking subparsers. Instead, it is intended to be used by the subparsers
// themselves to implement a higher-level interface.
func (context *MarkupParseContext) Pop() cgo.Handle {
	var _arg0 *C.GMarkupParseContext // out
	var _cret C.gpointer             // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	_cret = C.g_markup_parse_context_pop(_arg0)
	runtime.KeepAlive(context)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// Push: temporarily redirects markup data to a sub-parser.
//
// This function may only be called from the start_element handler of a Parser.
// It must be matched with a corresponding call to g_markup_parse_context_pop()
// in the matching end_element handler (except in the case that the parser
// aborts due to an error).
//
// All tags, text and other data between the matching tags is redirected to the
// subparser given by parser. user_data is used as the user_data for that
// parser. user_data is also passed to the error callback in the event that an
// error occurs. This includes errors that occur in subparsers of the subparser.
//
// The end tag matching the start tag for which this call was made is handled by
// the previous parser (which is given its own user_data) which is why
// g_markup_parse_context_pop() is provided to allow "one last access" to the
// user_data provided to this function. In the case of error, the user_data
// provided here is passed directly to the error callback of the subparser and
// g_markup_parse_context_pop() should not be called. In either case, if
// user_data was allocated then it ought to be freed from both of these
// locations.
//
// This function is not intended to be directly called by users interested in
// invoking subparsers. Instead, it is intended to be used by the subparsers
// themselves to implement a higher-level interface.
//
// As an example, see the following implementation of a simple parser that
// counts the number of tags encountered.
//
//    static void start_element (context, element_name, ...)
//    {
//      if (strcmp (element_name, "count-these") == 0)
//        start_counting (context);
//
//      // else, handle other tags...
//    }
//
//    static void end_element (context, element_name, ...)
//    {
//      if (strcmp (element_name, "count-these") == 0)
//        g_print ("Counted d tags\n", end_counting (context));
//
//      // else, handle other tags...
//    }
func (context *MarkupParseContext) Push(parser *MarkupParser, userData cgo.Handle) {
	var _arg0 *C.GMarkupParseContext // out
	var _arg1 *C.GMarkupParser       // out
	var _arg2 C.gpointer             // out

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))
	_arg1 = (*C.GMarkupParser)(gextras.StructNative(unsafe.Pointer(parser)))
	_arg2 = (C.gpointer)(unsafe.Pointer(userData))

	C.g_markup_parse_context_push(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(parser)
	runtime.KeepAlive(userData)
}

// MarkupParser: any of the fields in Parser can be NULL, in which case they
// will be ignored. Except for the error function, any of these callbacks can
// set an error; in particular the G_MARKUP_ERROR_UNKNOWN_ELEMENT,
// G_MARKUP_ERROR_UNKNOWN_ATTRIBUTE, and G_MARKUP_ERROR_INVALID_CONTENT errors
// are intended to be set from these callbacks. If you set an error from a
// callback, g_markup_parse_context_parse() will report that error back to its
// caller.
//
// An instance of this type is always passed by reference.
type MarkupParser struct {
	*markupParser
}

// markupParser is the struct that's finalized.
type markupParser struct {
	native *C.GMarkupParser
}
