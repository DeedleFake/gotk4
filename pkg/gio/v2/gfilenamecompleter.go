// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_filename_completer_get_type()), F: marshalFilenameCompleterer},
	})
}

// FilenameCompleterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FilenameCompleterOverrider interface {
	GotCompletionData()
}

// FilenameCompleter completes partial file and directory names given a partial
// string by looking in the file system for clues. Can return a list of possible
// completion strings for widget implementations.
type FilenameCompleter struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*FilenameCompleter)(nil)
)

func wrapFilenameCompleter(obj *externglib.Object) *FilenameCompleter {
	return &FilenameCompleter{
		Object: obj,
	}
}

func marshalFilenameCompleterer(p uintptr) (interface{}, error) {
	return wrapFilenameCompleter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectGotCompletionData: emitted when the file name completion information
// comes available.
func (completer *FilenameCompleter) ConnectGotCompletionData(f func()) externglib.SignalHandle {
	return completer.Connect("got-completion-data", f)
}

// NewFilenameCompleter creates a new filename completer.
//
// The function returns the following values:
//
//    - filenameCompleter: Completer.
//
func NewFilenameCompleter() *FilenameCompleter {
	var _cret *C.GFilenameCompleter // in

	_cret = C.g_filename_completer_new()

	var _filenameCompleter *FilenameCompleter // out

	_filenameCompleter = wrapFilenameCompleter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _filenameCompleter
}

// CompletionSuffix obtains a completion for initial_text from completer.
//
// The function takes the following parameters:
//
//    - initialText: text to be completed.
//
// The function returns the following values:
//
//    - utf8 (optional): completed string, or NULL if no completion exists. This
//      string is not owned by GIO, so remember to g_free() it when finished.
//
func (completer *FilenameCompleter) CompletionSuffix(initialText string) string {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(completer.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(initialText)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_filename_completer_get_completion_suffix(_arg0, _arg1)
	runtime.KeepAlive(completer)
	runtime.KeepAlive(initialText)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Completions gets an array of completion strings for a given initial text.
//
// The function takes the following parameters:
//
//    - initialText: text to be completed.
//
// The function returns the following values:
//
//    - utf8s: array of strings with possible completions for initial_text. This
//      array must be freed by g_strfreev() when finished.
//
func (completer *FilenameCompleter) Completions(initialText string) []string {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 *C.char               // out
	var _cret **C.char              // in

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(completer.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(initialText)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_filename_completer_get_completions(_arg0, _arg1)
	runtime.KeepAlive(completer)
	runtime.KeepAlive(initialText)

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

// SetDirsOnly: if dirs_only is TRUE, completer will only complete directory
// names, and not file names.
//
// The function takes the following parameters:
//
//    - dirsOnly: #gboolean.
//
func (completer *FilenameCompleter) SetDirsOnly(dirsOnly bool) {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(completer.Native()))
	if dirsOnly {
		_arg1 = C.TRUE
	}

	C.g_filename_completer_set_dirs_only(_arg0, _arg1)
	runtime.KeepAlive(completer)
	runtime.KeepAlive(dirsOnly)
}
