// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_key_file_get_type()), F: marshalKeyFile},
	})
}

// KeyFileError: error codes returned by key file parsing.
type KeyFileError int

const (
	// KeyFileErrorUnknownEncoding: text being parsed was in an unknown encoding
	KeyFileErrorUnknownEncoding KeyFileError = iota
	// KeyFileErrorParse: document was ill-formed
	KeyFileErrorParse
	// KeyFileErrorNotFound: file was not found
	KeyFileErrorNotFound
	// KeyFileErrorKeyNotFound: requested key was not found
	KeyFileErrorKeyNotFound
	// KeyFileErrorGroupNotFound: requested group was not found
	KeyFileErrorGroupNotFound
	// KeyFileErrorInvalidValue: value could not be parsed
	KeyFileErrorInvalidValue
)

// String returns the name in string for KeyFileError.
func (k KeyFileError) String() string {
	switch k {
	case KeyFileErrorUnknownEncoding:
		return "UnknownEncoding"
	case KeyFileErrorParse:
		return "Parse"
	case KeyFileErrorNotFound:
		return "NotFound"
	case KeyFileErrorKeyNotFound:
		return "KeyNotFound"
	case KeyFileErrorGroupNotFound:
		return "GroupNotFound"
	case KeyFileErrorInvalidValue:
		return "InvalidValue"
	default:
		return fmt.Sprintf("KeyFileError(%d)", k)
	}
}

// KeyFileFlags flags which influence the parsing.
type KeyFileFlags int

const (
	// KeyFileNone: no flags, default behaviour
	KeyFileNone KeyFileFlags = 0b0
	// KeyFileKeepComments: use this flag if you plan to write the (possibly
	// modified) contents of the key file back to a file; otherwise all comments
	// will be lost when the key file is written back.
	KeyFileKeepComments KeyFileFlags = 0b1
	// KeyFileKeepTranslations: use this flag if you plan to write the (possibly
	// modified) contents of the key file back to a file; otherwise only the
	// translations for the current language will be written back.
	KeyFileKeepTranslations KeyFileFlags = 0b10
)

// String returns the names in string for KeyFileFlags.
func (k KeyFileFlags) String() string {
	if k == 0 {
		return "KeyFileFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(55)

	for k != 0 {
		next := k & (k - 1)
		bit := k - next

		switch bit {
		case KeyFileNone:
			builder.WriteString("None|")
		case KeyFileKeepComments:
			builder.WriteString("KeepComments|")
		case KeyFileKeepTranslations:
			builder.WriteString("KeepTranslations|")
		default:
			builder.WriteString(fmt.Sprintf("KeyFileFlags(0b%b)|", bit))
		}

		k = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// KeyFile struct contains only private data and should not be accessed
// directly.
type KeyFile struct {
	nocopy gextras.NoCopy
	native *C.GKeyFile
}

func marshalKeyFile(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &KeyFile{native: (*C.GKeyFile)(unsafe.Pointer(b))}, nil
}

// NewKeyFile constructs a struct KeyFile.
func NewKeyFile() *KeyFile {
	var _cret *C.GKeyFile // in

	_cret = C.g_key_file_new()

	var _keyFile *KeyFile // out

	_keyFile = (*KeyFile)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_keyFile, func(v *KeyFile) {
		C.free(gextras.StructNative(unsafe.Pointer(v)))
	})

	return _keyFile
}

// Boolean returns the value associated with key under group_name as a boolean.
//
// If key cannot be found then FALSE is returned and error is set to
// KEY_FILE_ERROR_KEY_NOT_FOUND. Likewise, if the value associated with key
// cannot be interpreted as a boolean then FALSE is returned and error is set to
// KEY_FILE_ERROR_INVALID_VALUE.
func (keyFile *KeyFile) Boolean(groupName string, key string) error {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_key_file_get_boolean(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Comment retrieves a comment above key from group_name. If key is NULL then
// comment will be read from above group_name. If both key and group_name are
// NULL, then comment will be read from above the first group in the file.
//
// Note that the returned string does not include the '#' comment markers, but
// does include any whitespace after them (on each line). It includes the line
// breaks between lines, but does not include the final line break.
func (keyFile *KeyFile) Comment(groupName string, key string) (string, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cret *C.gchar    // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_key_file_get_comment(_arg0, _arg1, _arg2, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

// Double returns the value associated with key under group_name as a double. If
// group_name is NULL, the start_group is used.
//
// If key cannot be found then 0.0 is returned and error is set to
// KEY_FILE_ERROR_KEY_NOT_FOUND. Likewise, if the value associated with key
// cannot be interpreted as a double then 0.0 is returned and error is set to
// KEY_FILE_ERROR_INVALID_VALUE.
func (keyFile *KeyFile) Double(groupName string, key string) (float64, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cret C.gdouble   // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_key_file_get_double(_arg0, _arg1, _arg2, &_cerr)

	var _gdouble float64 // out
	var _goerr error     // out

	_gdouble = float64(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gdouble, _goerr
}

// Groups returns all groups in the key file loaded with key_file. The array of
// returned groups will be NULL-terminated, so length may optionally be NULL.
func (keyFile *KeyFile) Groups() (uint, []string) {
	var _arg0 *C.GKeyFile // out
	var _arg1 C.gsize     // in
	var _cret **C.gchar   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))

	_cret = C.g_key_file_get_groups(_arg0, &_arg1)

	var _length uint    // out
	var _utf8s []string // out

	_length = uint(_arg1)
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _length, _utf8s
}

// Int64 returns the value associated with key under group_name as a signed
// 64-bit integer. This is similar to g_key_file_get_integer() but can return
// 64-bit results without truncation.
func (keyFile *KeyFile) Int64(groupName string, key string) (int64, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cret C.gint64    // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_key_file_get_int64(_arg0, _arg1, _arg2, &_cerr)

	var _gint64 int64 // out
	var _goerr error  // out

	_gint64 = int64(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint64, _goerr
}

// Integer returns the value associated with key under group_name as an integer.
//
// If key cannot be found then 0 is returned and error is set to
// KEY_FILE_ERROR_KEY_NOT_FOUND. Likewise, if the value associated with key
// cannot be interpreted as an integer, or is out of range for a #gint, then 0
// is returned and error is set to KEY_FILE_ERROR_INVALID_VALUE.
func (keyFile *KeyFile) Integer(groupName string, key string) (int, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cret C.gint      // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_key_file_get_integer(_arg0, _arg1, _arg2, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

// Keys returns all keys for the group name group_name. The array of returned
// keys will be NULL-terminated, so length may optionally be NULL. In the event
// that the group_name cannot be found, NULL is returned and error is set to
// KEY_FILE_ERROR_GROUP_NOT_FOUND.
func (keyFile *KeyFile) Keys(groupName string) (uint, []string, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 C.gsize     // in
	var _cret **C.gchar   // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_key_file_get_keys(_arg0, _arg1, &_arg2, &_cerr)

	var _length uint    // out
	var _utf8s []string // out
	var _goerr error    // out

	_length = uint(_arg2)
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _utf8s, _goerr
}

// LocaleForKey returns the actual locale which the result of
// g_key_file_get_locale_string() or g_key_file_get_locale_string_list() came
// from.
//
// If calling g_key_file_get_locale_string() or
// g_key_file_get_locale_string_list() with exactly the same key_file,
// group_name, key and locale, the result of those functions will have
// originally been tagged with the locale that is the result of this function.
func (keyFile *KeyFile) LocaleForKey(groupName string, key string, locale string) string {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gchar    // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(locale)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_key_file_get_locale_for_key(_arg0, _arg1, _arg2, _arg3)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// LocaleString returns the value associated with key under group_name
// translated in the given locale if available. If locale is NULL then the
// current locale is assumed.
//
// If locale is to be non-NULL, or if the current locale will change over the
// lifetime of the File, it must be loaded with G_KEY_FILE_KEEP_TRANSLATIONS in
// order to load strings for all locales.
//
// If key cannot be found then NULL is returned and error is set to
// KEY_FILE_ERROR_KEY_NOT_FOUND. If the value associated with key cannot be
// interpreted or no suitable translation can be found then the untranslated
// value is returned.
func (keyFile *KeyFile) LocaleString(groupName string, key string, locale string) (string, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gchar    // out
	var _cret *C.gchar    // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(locale)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_key_file_get_locale_string(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

// StartGroup returns the name of the start group of the file.
func (keyFile *KeyFile) StartGroup() string {
	var _arg0 *C.GKeyFile // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))

	_cret = C.g_key_file_get_start_group(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// String returns the string value associated with key under group_name. Unlike
// g_key_file_get_value(), this function handles escape sequences like \s.
//
// In the event the key cannot be found, NULL is returned and error is set to
// KEY_FILE_ERROR_KEY_NOT_FOUND. In the event that the group_name cannot be
// found, NULL is returned and error is set to KEY_FILE_ERROR_GROUP_NOT_FOUND.
func (keyFile *KeyFile) String(groupName string, key string) (string, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cret *C.gchar    // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_key_file_get_string(_arg0, _arg1, _arg2, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

// Uint64 returns the value associated with key under group_name as an unsigned
// 64-bit integer. This is similar to g_key_file_get_integer() but can return
// large positive results without truncation.
func (keyFile *KeyFile) Uint64(groupName string, key string) (uint64, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cret C.guint64   // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_key_file_get_uint64(_arg0, _arg1, _arg2, &_cerr)

	var _guint64 uint64 // out
	var _goerr error    // out

	_guint64 = uint64(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint64, _goerr
}

// Value returns the raw value associated with key under group_name. Use
// g_key_file_get_string() to retrieve an unescaped UTF-8 string.
//
// In the event the key cannot be found, NULL is returned and error is set to
// KEY_FILE_ERROR_KEY_NOT_FOUND. In the event that the group_name cannot be
// found, NULL is returned and error is set to KEY_FILE_ERROR_GROUP_NOT_FOUND.
func (keyFile *KeyFile) Value(groupName string, key string) (string, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cret *C.gchar    // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_key_file_get_value(_arg0, _arg1, _arg2, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

// HasGroup looks whether the key file has the group group_name.
func (keyFile *KeyFile) HasGroup(groupName string) bool {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_key_file_has_group(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LoadFromData loads a key file from memory into an empty File structure. If
// the object cannot be created then error is set to a FileError.
func (keyFile *KeyFile) LoadFromData(data string, length uint, flags KeyFileFlags) error {
	var _arg0 *C.GKeyFile     // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gsize         // out
	var _arg3 C.GKeyFileFlags // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(data)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gsize(length)
	_arg3 = C.GKeyFileFlags(flags)

	C.g_key_file_load_from_data(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// LoadFromDataDirs: this function looks for a key file named file in the paths
// returned from g_get_user_data_dir() and g_get_system_data_dirs(), loads the
// file into key_file and returns the file's full path in full_path. If the file
// could not be loaded then an error is set to either a Error or FileError.
func (keyFile *KeyFile) LoadFromDataDirs(file string, flags KeyFileFlags) (string, error) {
	var _arg0 *C.GKeyFile     // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // in
	var _arg3 C.GKeyFileFlags // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(file)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg3 = C.GKeyFileFlags(flags)

	C.g_key_file_load_from_data_dirs(_arg0, _arg1, &_arg2, _arg3, &_cerr)

	var _fullPath string // out
	var _goerr error     // out

	_fullPath = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fullPath, _goerr
}

// LoadFromDirs: this function looks for a key file named file in the paths
// specified in search_dirs, loads the file into key_file and returns the file's
// full path in full_path.
//
// If the file could not be found in any of the search_dirs,
// G_KEY_FILE_ERROR_NOT_FOUND is returned. If the file is found but the OS
// returns an error when opening or reading the file, a G_FILE_ERROR is
// returned. If there is a problem parsing the file, a G_KEY_FILE_ERROR is
// returned.
func (keyFile *KeyFile) LoadFromDirs(file string, searchDirs []string, flags KeyFileFlags) (string, error) {
	var _arg0 *C.GKeyFile     // out
	var _arg1 *C.gchar        // out
	var _arg2 **C.gchar       // out
	var _arg3 *C.gchar        // in
	var _arg4 C.GKeyFileFlags // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(file)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(searchDirs)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(searchDirs)+1)
			var zero *C.gchar
			out[len(searchDirs)] = zero
			for i := range searchDirs {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(searchDirs[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg4 = C.GKeyFileFlags(flags)

	C.g_key_file_load_from_dirs(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)

	var _fullPath string // out
	var _goerr error     // out

	_fullPath = C.GoString((*C.gchar)(unsafe.Pointer(_arg3)))
	defer C.free(unsafe.Pointer(_arg3))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fullPath, _goerr
}

// LoadFromFile loads a key file into an empty File structure.
//
// If the OS returns an error when opening or reading the file, a G_FILE_ERROR
// is returned. If there is a problem parsing the file, a G_KEY_FILE_ERROR is
// returned.
//
// This function will never return a G_KEY_FILE_ERROR_NOT_FOUND error. If the
// file is not found, G_FILE_ERROR_NOENT is returned.
func (keyFile *KeyFile) LoadFromFile(file string, flags KeyFileFlags) error {
	var _arg0 *C.GKeyFile     // out
	var _arg1 *C.gchar        // out
	var _arg2 C.GKeyFileFlags // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(file)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GKeyFileFlags(flags)

	C.g_key_file_load_from_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// RemoveComment removes a comment above key from group_name. If key is NULL
// then comment will be removed above group_name. If both key and group_name are
// NULL, then comment will be removed above the first group in the file.
func (keyFile *KeyFile) RemoveComment(groupName string, key string) error {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_key_file_remove_comment(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// RemoveGroup removes the specified group, group_name, from the key file.
func (keyFile *KeyFile) RemoveGroup(groupName string) error {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_key_file_remove_group(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// RemoveKey removes key in group_name from the key file.
func (keyFile *KeyFile) RemoveKey(groupName string, key string) error {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_key_file_remove_key(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SaveToFile writes the contents of key_file to filename using
// g_file_set_contents(). If you need stricter guarantees about durability of
// the written file than are provided by g_file_set_contents(), use
// g_file_set_contents_full() with the return value of g_key_file_to_data().
//
// This function can fail for any of the reasons that g_file_set_contents() may
// fail.
func (keyFile *KeyFile) SaveToFile(filename string) error {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_key_file_save_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetBoolean associates a new boolean value with key under group_name. If key
// cannot be found then it is created.
func (keyFile *KeyFile) SetBoolean(groupName string, key string, value bool) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 C.gboolean  // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	if value {
		_arg3 = C.TRUE
	}

	C.g_key_file_set_boolean(_arg0, _arg1, _arg2, _arg3)
}

// SetBooleanList associates a list of boolean values with key under group_name.
// If key cannot be found then it is created. If group_name is NULL, the
// start_group is used.
func (keyFile *KeyFile) SetBooleanList(groupName string, key string, list []bool) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gboolean // out
	var _arg4 C.gsize

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg4 = (C.gsize)(len(list))
	if len(list) > 0 {
		_arg3 = (*C.gboolean)(unsafe.Pointer(&list[0]))
	}

	C.g_key_file_set_boolean_list(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetComment places a comment above key from group_name.
//
// If key is NULL then comment will be written above group_name. If both key and
// group_name are NULL, then comment will be written above the first group in
// the file.
//
// Note that this function prepends a '#' comment marker to each line of
// comment.
func (keyFile *KeyFile) SetComment(groupName string, key string, comment string) error {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gchar    // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(comment)))
	defer C.free(unsafe.Pointer(_arg3))

	C.g_key_file_set_comment(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetDouble associates a new double value with key under group_name. If key
// cannot be found then it is created.
func (keyFile *KeyFile) SetDouble(groupName string, key string, value float64) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 C.gdouble   // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gdouble(value)

	C.g_key_file_set_double(_arg0, _arg1, _arg2, _arg3)
}

// SetDoubleList associates a list of double values with key under group_name.
// If key cannot be found then it is created.
func (keyFile *KeyFile) SetDoubleList(groupName string, key string, list []float64) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gdouble  // out
	var _arg4 C.gsize

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg4 = (C.gsize)(len(list))
	if len(list) > 0 {
		_arg3 = (*C.gdouble)(unsafe.Pointer(&list[0]))
	}

	C.g_key_file_set_double_list(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetInt64 associates a new integer value with key under group_name. If key
// cannot be found then it is created.
func (keyFile *KeyFile) SetInt64(groupName string, key string, value int64) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 C.gint64    // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint64(value)

	C.g_key_file_set_int64(_arg0, _arg1, _arg2, _arg3)
}

// SetInteger associates a new integer value with key under group_name. If key
// cannot be found then it is created.
func (keyFile *KeyFile) SetInteger(groupName string, key string, value int) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 C.gint      // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(value)

	C.g_key_file_set_integer(_arg0, _arg1, _arg2, _arg3)
}

// SetIntegerList associates a list of integer values with key under group_name.
// If key cannot be found then it is created.
func (keyFile *KeyFile) SetIntegerList(groupName string, key string, list []int) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gint     // out
	var _arg4 C.gsize

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg4 = (C.gsize)(len(list))
	if len(list) > 0 {
		_arg3 = (*C.gint)(unsafe.Pointer(&list[0]))
	}

	C.g_key_file_set_integer_list(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetListSeparator sets the character which is used to separate values in
// lists. Typically ';' or ',' are used as separators. The default list
// separator is ';'.
func (keyFile *KeyFile) SetListSeparator(separator byte) {
	var _arg0 *C.GKeyFile // out
	var _arg1 C.gchar     // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = C.gchar(separator)

	C.g_key_file_set_list_separator(_arg0, _arg1)
}

// SetLocaleString associates a string value for key and locale under
// group_name. If the translation for key cannot be found then it is created.
func (keyFile *KeyFile) SetLocaleString(groupName string, key string, locale string, _string string) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gchar    // out
	var _arg4 *C.gchar    // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(locale)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg4))

	C.g_key_file_set_locale_string(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetLocaleStringList associates a list of string values for key and locale
// under group_name. If the translation for key cannot be found then it is
// created.
func (keyFile *KeyFile) SetLocaleStringList(groupName string, key string, locale string, list []string) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gchar    // out
	var _arg4 **C.gchar   // out
	var _arg5 C.gsize

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(locale)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg5 = (C.gsize)(len(list))
	_arg4 = (**C.gchar)(C.malloc(C.ulong(len(list)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice((**C.gchar)(_arg4), len(list))
		for i := range list {
			out[i] = (*C.gchar)(unsafe.Pointer(C.CString(list[i])))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.g_key_file_set_locale_string_list(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// SetString associates a new string value with key under group_name. If key
// cannot be found then it is created. If group_name cannot be found then it is
// created. Unlike g_key_file_set_value(), this function handles characters that
// need escaping, such as newlines.
func (keyFile *KeyFile) SetString(groupName string, key string, _string string) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gchar    // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg3))

	C.g_key_file_set_string(_arg0, _arg1, _arg2, _arg3)
}

// SetStringList associates a list of string values for key under group_name. If
// key cannot be found then it is created. If group_name cannot be found then it
// is created.
func (keyFile *KeyFile) SetStringList(groupName string, key string, list []string) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 **C.gchar   // out
	var _arg4 C.gsize

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg4 = (C.gsize)(len(list))
	_arg3 = (**C.gchar)(C.malloc(C.ulong(len(list)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((**C.gchar)(_arg3), len(list))
		for i := range list {
			out[i] = (*C.gchar)(unsafe.Pointer(C.CString(list[i])))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.g_key_file_set_string_list(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetUint64 associates a new integer value with key under group_name. If key
// cannot be found then it is created.
func (keyFile *KeyFile) SetUint64(groupName string, key string, value uint64) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 C.guint64   // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint64(value)

	C.g_key_file_set_uint64(_arg0, _arg1, _arg2, _arg3)
}

// SetValue associates a new value with key under group_name.
//
// If key cannot be found then it is created. If group_name cannot be found then
// it is created. To set an UTF-8 string which may contain characters that need
// escaping (such as newlines or spaces), use g_key_file_set_string().
func (keyFile *KeyFile) SetValue(groupName string, key string, value string) {
	var _arg0 *C.GKeyFile // out
	var _arg1 *C.gchar    // out
	var _arg2 *C.gchar    // out
	var _arg3 *C.gchar    // out

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg3))

	C.g_key_file_set_value(_arg0, _arg1, _arg2, _arg3)
}

// ToData: this function outputs key_file as a string.
//
// Note that this function never reports an error, so it is safe to pass NULL as
// error.
func (keyFile *KeyFile) ToData() (uint, string, error) {
	var _arg0 *C.GKeyFile // out
	var _arg1 C.gsize     // in
	var _cret *C.gchar    // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))

	_cret = C.g_key_file_to_data(_arg0, &_arg1, &_cerr)

	var _length uint // out
	var _utf8 string // out
	var _goerr error // out

	_length = uint(_arg1)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _length, _utf8, _goerr
}
