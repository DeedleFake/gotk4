// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// ItemizeWithBaseDir: like pango_itemize(), but with an explicitly specified
// base direction.
//
// The base direction is used when computing bidirectional levels. (see
// pango.Context.SetBaseDir()). itemize gets the base direction from the
// PangoContext.
//
// The function takes the following parameters:
//
//    - context: structure holding information that affects the itemization
//      process.
//    - baseDir: base direction to use for bidirectional processing.
//    - text to itemize.
//    - startIndex: first byte in text to process.
//    - length: number of bytes (not characters) to process after start_index.
//      This must be >= 0.
//    - attrs: set of attributes that apply to text.
//    - cachedIter (optional): cached attribute iterator, or NULL.
//
// The function returns the following values:
//
//    - list: GList of pango.Item structures. The items should be freed using
//      pango.Item.Free() probably in combination with g_list_free_full().
//
func ItemizeWithBaseDir(context *Context, baseDir Direction, text string, startIndex, length int, attrs *AttrList, cachedIter *AttrIterator) []*Item {
	var _arg1 *C.PangoContext      // out
	var _arg2 C.PangoDirection     // out
	var _arg3 *C.char              // out
	var _arg4 C.int                // out
	var _arg5 C.int                // out
	var _arg6 *C.PangoAttrList     // out
	var _arg7 *C.PangoAttrIterator // out
	var _cret *C.GList             // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = C.PangoDirection(baseDir)
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.int(startIndex)
	_arg5 = C.int(length)
	_arg6 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attrs)))
	if cachedIter != nil {
		_arg7 = (*C.PangoAttrIterator)(gextras.StructNative(unsafe.Pointer(cachedIter)))
	}

	_cret = C.pango_itemize_with_base_dir(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(context)
	runtime.KeepAlive(baseDir)
	runtime.KeepAlive(text)
	runtime.KeepAlive(startIndex)
	runtime.KeepAlive(length)
	runtime.KeepAlive(attrs)
	runtime.KeepAlive(cachedIter)

	var _list []*Item // out

	_list = make([]*Item, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.PangoItem)(v)
		var dst *Item // out
		dst = (*Item)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_item_free((*C.PangoItem)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}
