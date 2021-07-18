// Code generated by girgen. DO NOT EDIT.

package pango

import (
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
		{T: externglib.Type(C.pango_language_get_type()), F: marshalLanguage},
	})
}

// Language: PangoLanguage structure is used to represent a language.
//
// PangoLanguage pointers can be efficiently copied and compared with each
// other.
type Language struct {
	nocopy gextras.NoCopy
	native *C.PangoLanguage
}

func marshalLanguage(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Language{native: (*C.PangoLanguage)(unsafe.Pointer(b))}, nil
}

// SampleString: get a string that is representative of the characters needed to
// render a particular language.
//
// The sample text may be a pangram, but is not necessarily. It is chosen to be
// demonstrative of normal text in the language, as well as exposing font
// feature requirements unique to the language. It is suitable for use as sample
// text in a font selection dialog.
//
// If language is NULL, the default language as found by
// pango.Language.GetDefault is used.
//
// If Pango does not have a sample string for language, the classic "The quick
// brown fox..." is returned. This can be detected by comparing the returned
// pointer value to that returned for (non-existent) language code "xx". That
// is, compare to:
//
//    pango_language_get_sample_string (pango_language_from_string ("xx"))
func (language *Language) SampleString() string {
	var _arg0 *C.PangoLanguage // out
	var _cret *C.char          // in

	_arg0 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_language_get_sample_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IncludesScript determines if script is one of the scripts used to write
// language. The returned value is conservative; if nothing is known about the
// language tag language, TRUE will be returned, since, as far as Pango knows,
// script might be used to write language.
//
// This routine is used in Pango's itemization process when determining if a
// supplied language tag is relevant to a particular section of text. It
// probably is not useful for applications in most circumstances.
//
// This function uses pango.Language.GetScripts() internally.
func (language *Language) IncludesScript(script Script) bool {
	var _arg0 *C.PangoLanguage // out
	var _arg1 C.PangoScript    // out
	var _cret C.gboolean       // in

	_arg0 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))
	_arg1 = C.PangoScript(script)

	_cret = C.pango_language_includes_script(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Matches checks if a language tag matches one of the elements in a list of
// language ranges.
//
// A language tag is considered to match a range in the list if the range is
// '*', the range is exactly the tag, or the range is a prefix of the tag, and
// the character after it in the tag is '-'.
func (language *Language) Matches(rangeList string) bool {
	var _arg0 *C.PangoLanguage // out
	var _arg1 *C.char          // out
	var _cret C.gboolean       // in

	_arg0 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(rangeList)))

	_cret = C.pango_language_matches(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String gets the RFC-3066 format string representing the given language tag.
func (language *Language) String() string {
	var _arg0 *C.PangoLanguage // out
	var _cret *C.char          // in

	_arg0 = (*C.PangoLanguage)(gextras.StructNative(unsafe.Pointer(language)))

	_cret = C.pango_language_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LanguageFromString: convert a language tag to a PangoLanguage.
//
// The language tag must be in a RFC-3066 format. PangoLanguage pointers can be
// efficiently copied (copy the pointer) and compared with other language tags
// (compare the pointer.)
//
// This function first canonicalizes the string by converting it to lowercase,
// mapping '_' to '-', and stripping all characters other than letters and '-'.
//
// Use pango.Language.GetDefault if you want to get the PangoLanguage for the
// current locale of the process.
func LanguageFromString(language string) *Language {
	var _arg1 *C.char          // out
	var _cret *C.PangoLanguage // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(language)))

	_cret = C.pango_language_from_string(_arg1)

	var _ret *Language // out

	_ret = (*Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _ret
}

// LanguageGetDefault returns the PangoLanguage for the current locale of the
// process.
//
// On Unix systems, this is the return value is derived from setlocale
// (LC_CTYPE, NULL), and the user can affect this through the environment
// variables LC_ALL, LC_CTYPE or LANG (checked in that order). The locale string
// typically is in the form lang_COUNTRY, where lang is an ISO-639 language
// code, and COUNTRY is an ISO-3166 country code. For instance, sv_FI for
// Swedish as written in Finland or pt_BR for Portuguese as written in Brazil.
//
// On Windows, the C library does not use any such environment variables, and
// setting them won't affect the behavior of functions like ctime(). The user
// sets the locale through the Regional Options in the Control Panel. The C
// library (in the setlocale() function) does not use country and language
// codes, but country and language names spelled out in English. However, this
// function does check the above environment variables, and does return a
// Unix-style locale string based on either said environment variables or the
// thread's current locale.
//
// Your application should call setlocale(LC_ALL, "") for the user settings to
// take effect. GTK does this in its initialization functions automatically (by
// calling gtk_set_locale()). See the setlocale() manpage for more details.
//
// Note that the default language can change over the life of an application.
func LanguageGetDefault() *Language {
	var _cret *C.PangoLanguage // in

	_cret = C.pango_language_get_default()

	var _language *Language // out

	_language = (*Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _language
}
