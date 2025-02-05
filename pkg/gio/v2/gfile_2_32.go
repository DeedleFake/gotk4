// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// NewFileTmp opens a file in the preferred directory for temporary files (as
// returned by g_get_tmp_dir()) and returns a #GFile and IOStream pointing to
// it.
//
// tmpl should be a string in the GLib file name encoding containing a sequence
// of six 'X' characters, and containing no directory components. If it is NULL,
// a default template is used.
//
// Unlike the other #GFile constructors, this will return NULL if a temporary
// file could not be created.
//
// The function takes the following parameters:
//
//    - tmpl (optional): template for the file name, as in g_file_open_tmp(), or
//      NULL for a default template.
//
// The function returns the following values:
//
//    - iostream: on return, a IOStream for the created file.
//    - file: new #GFile. Free the returned object with g_object_unref().
//
func NewFileTmp(tmpl string) (*FileIOStream, *File, error) {
	var _arg1 *C.char          // out
	var _arg2 *C.GFileIOStream // in
	var _cret *C.GFile         // in
	var _cerr *C.GError        // in

	if tmpl != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(tmpl)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.g_file_new_tmp(_arg1, &_arg2, &_cerr)
	runtime.KeepAlive(tmpl)

	var _iostream *FileIOStream // out
	var _file *File             // out
	var _goerr error            // out

	_iostream = wrapFileIOStream(coreglib.AssumeOwnership(unsafe.Pointer(_arg2)))
	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _iostream, _file, _goerr
}
