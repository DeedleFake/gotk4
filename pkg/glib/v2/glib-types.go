// Code generated by girgen. DO NOT EDIT.

package glib

import (
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "glib.h"
import "C"

type Type = uint

func StrvGetType() externglib.Type {
	var _cret C.GType // in

	_cret = C.g_strv_get_type()

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

func VariantGetGType() externglib.Type {
	var _cret C.GType // in

	_cret = C.g_variant_get_gtype()

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}
