// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	_ "runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

func init() {
	girepository.Require("Gdk", "4.0", girepository.LoadFlagLazy)
}

// The function returns the following values:
//
func GLErrorQuark() glib.Quark {
	_info := girepository.MustFind("Gdk", "quark")
	_gret := _info.Invoke(nil, nil)
	_cret := *(*C.guint32)(unsafe.Pointer(&_gret))

	var _quark glib.Quark // out

	_quark = uint32(*(*C.guint32)(unsafe.Pointer(&_cret)))

	return _quark
}

// The function returns the following values:
//
func VulkanErrorQuark() glib.Quark {
	_info := girepository.MustFind("Gdk", "quark")
	_gret := _info.Invoke(nil, nil)
	_cret := *(*C.guint32)(unsafe.Pointer(&_gret))

	var _quark glib.Quark // out

	_quark = uint32(*(*C.guint32)(unsafe.Pointer(&_cret)))

	return _quark
}
