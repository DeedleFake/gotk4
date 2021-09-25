// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

// IsZeroWidth checks if a character that should not be normally rendered.
//
// This includes all Unicode characters with "ZERO WIDTH" in their name, as well
// as *bidi* formatting characters, and a few other ones. This is totally
// different from g_unichar_iszerowidth() and is at best misnamed.
func IsZeroWidth(ch uint32) bool {
	var _arg1 C.gunichar // out
	var _cret C.gboolean // in

	_arg1 = C.gunichar(ch)

	_cret = C.pango_is_zero_width(_arg1)
	runtime.KeepAlive(ch)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ParseEnum parses an enum type and stores the result in value.
//
// If str does not match the nick name of any of the possible values for the
// enum and is not an integer, FALSE is returned, a warning is issued if warn is
// TRUE, and a string representing the list of possible values is stored in
// possible_values. The list is slash-separated, eg. "none/start/middle/end". If
// failed and possible_values is not NULL, returned string should be freed using
// g_free().
//
// Deprecated: since version 1.38.
func ParseEnum(typ externglib.Type, str string, warn bool) (int, string, bool) {
	var _arg1 C.GType    // out
	var _arg2 *C.char    // out
	var _arg3 C.int      // in
	var _arg4 C.gboolean // out
	var _arg5 *C.char    // in
	var _cret C.gboolean // in

	_arg1 = C.GType(typ)
	if str != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(str)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if warn {
		_arg4 = C.TRUE
	}

	_cret = C.pango_parse_enum(_arg1, _arg2, &_arg3, _arg4, &_arg5)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _value int             // out
	var _possibleValues string // out
	var _ok bool               // out

	_value = int(_arg3)
	if _arg5 != nil {
		_possibleValues = C.GoString((*C.gchar)(unsafe.Pointer(_arg5)))
		defer C.free(unsafe.Pointer(_arg5))
	}
	if _cret != 0 {
		_ok = true
	}

	return _value, _possibleValues, _ok
}

// ParseStretch parses a font stretch.
//
// The allowed values are "ultra_condensed", "extra_condensed", "condensed",
// "semi_condensed", "normal", "semi_expanded", "expanded", "extra_expanded" and
// "ultra_expanded". Case variations are ignored and the '_' characters may be
// omitted.
func ParseStretch(str string, warn bool) (Stretch, bool) {
	var _arg1 *C.char        // out
	var _arg2 C.PangoStretch // in
	var _arg3 C.gboolean     // out
	var _cret C.gboolean     // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	if warn {
		_arg3 = C.TRUE
	}

	_cret = C.pango_parse_stretch(_arg1, &_arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _stretch Stretch // out
	var _ok bool         // out

	_stretch = Stretch(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _stretch, _ok
}

// ParseStyle parses a font style.
//
// The allowed values are "normal", "italic" and "oblique", case variations
// being ignored.
func ParseStyle(str string, warn bool) (Style, bool) {
	var _arg1 *C.char      // out
	var _arg2 C.PangoStyle // in
	var _arg3 C.gboolean   // out
	var _cret C.gboolean   // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	if warn {
		_arg3 = C.TRUE
	}

	_cret = C.pango_parse_style(_arg1, &_arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _style Style // out
	var _ok bool     // out

	_style = Style(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _style, _ok
}

// ParseVariant parses a font variant.
//
// The allowed values are "normal" and "smallcaps" or "small_caps", case
// variations being ignored.
func ParseVariant(str string, warn bool) (Variant, bool) {
	var _arg1 *C.char        // out
	var _arg2 C.PangoVariant // in
	var _arg3 C.gboolean     // out
	var _cret C.gboolean     // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	if warn {
		_arg3 = C.TRUE
	}

	_cret = C.pango_parse_variant(_arg1, &_arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _variant Variant // out
	var _ok bool         // out

	_variant = Variant(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _variant, _ok
}

// ParseWeight parses a font weight.
//
// The allowed values are "heavy", "ultrabold", "bold", "normal", "light",
// "ultraleight" and integers. Case variations are ignored.
func ParseWeight(str string, warn bool) (Weight, bool) {
	var _arg1 *C.char       // out
	var _arg2 C.PangoWeight // in
	var _arg3 C.gboolean    // out
	var _cret C.gboolean    // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))
	if warn {
		_arg3 = C.TRUE
	}

	_cret = C.pango_parse_weight(_arg1, &_arg2, _arg3)
	runtime.KeepAlive(str)
	runtime.KeepAlive(warn)

	var _weight Weight // out
	var _ok bool       // out

	_weight = Weight(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _weight, _ok
}

// SplitFileList splits a G_SEARCHPATH_SEPARATOR-separated list of files,
// stripping white space and substituting ~/ with $HOME/.
//
// Deprecated: since version 1.38.
func SplitFileList(str string) []string {
	var _arg1 *C.char  // out
	var _cret **C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_split_file_list(_arg1)
	runtime.KeepAlive(str)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// TrimString trims leading and trailing whitespace from a string.
//
// Deprecated: since version 1.38.
func TrimString(str string) string {
	var _arg1 *C.char // out
	var _cret *C.char // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_trim_string(_arg1)
	runtime.KeepAlive(str)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Version returns the encoded version of Pango available at run-time.
//
// This is similar to the macro PANGO_VERSION except that the macro returns the
// encoded version available at compile-time. A version number can be encoded
// into an integer using PANGO_VERSION_ENCODE().
func Version() int {
	var _cret C.int // in

	_cret = C.pango_version()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// VersionCheck checks that the Pango library in use is compatible with the
// given version.
//
// Generally you would pass in the constants PANGO_VERSION_MAJOR,
// PANGO_VERSION_MINOR, PANGO_VERSION_MICRO as the three arguments to this
// function; that produces a check that the library in use at run-time is
// compatible with the version of Pango the application or module was compiled
// against.
//
// Compatibility is defined by two things: first the version of the running
// library is newer than the version
// required_major.required_minor.required_micro. Second the running library must
// be binary compatible with the version
// required_major.required_minor.required_micro (same major version.)
//
// For compile-time version checking use PANGO_VERSION_CHECK().
func VersionCheck(requiredMajor int, requiredMinor int, requiredMicro int) string {
	var _arg1 C.int   // out
	var _arg2 C.int   // out
	var _arg3 C.int   // out
	var _cret *C.char // in

	_arg1 = C.int(requiredMajor)
	_arg2 = C.int(requiredMinor)
	_arg3 = C.int(requiredMicro)

	_cret = C.pango_version_check(_arg1, _arg2, _arg3)
	runtime.KeepAlive(requiredMajor)
	runtime.KeepAlive(requiredMinor)
	runtime.KeepAlive(requiredMicro)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// VersionString returns the version of Pango available at run-time.
//
// This is similar to the macro PANGO_VERSION_STRING except that the macro
// returns the version available at compile-time.
func VersionString() string {
	var _cret *C.char // in

	_cret = C.pango_version_string()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
