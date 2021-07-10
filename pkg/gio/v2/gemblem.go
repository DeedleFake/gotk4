// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_emblem_get_type()), F: marshalEmblemmer},
	})
}

// Emblemmer describes Emblem's methods.
type Emblemmer interface {
	gextras.Objector

	GetIcon() *Icon
	Origin() EmblemOrigin
}

// Emblem is an implementation of #GIcon that supports having an emblem, which
// is an icon with additional properties. It can than be added to a Icon.
//
// Currently, only metainformation about the emblem's origin is supported. More
// may be added in the future.
type Emblem struct {
	*externglib.Object

	Icon
}

var _ Emblemmer = (*Emblem)(nil)

func wrapEmblemmer(obj *externglib.Object) Emblemmer {
	return &Emblem{
		Object: obj,
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalEmblemmer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEmblemmer(obj), nil
}

// NewEmblem creates a new emblem for @icon.
func NewEmblem(icon Iconner) *Emblem {
	var _arg1 *C.GIcon   // out
	var _cret *C.GEmblem // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.g_emblem_new(_arg1)

	var _emblem *Emblem // out

	_emblem = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Emblem)

	return _emblem
}

// GetIcon gives back the icon from @emblem.
func (emblem *Emblem) GetIcon() *Icon {
	var _arg0 *C.GEmblem // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GEmblem)(unsafe.Pointer(emblem.Native()))

	_cret = C.g_emblem_get_icon(_arg0)

	var _icon *Icon // out

	_icon = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Icon)

	return _icon
}

// Origin gets the origin of the emblem.
func (emblem *Emblem) Origin() EmblemOrigin {
	var _arg0 *C.GEmblem      // out
	var _cret C.GEmblemOrigin // in

	_arg0 = (*C.GEmblem)(unsafe.Pointer(emblem.Native()))

	_cret = C.g_emblem_get_origin(_arg0)

	var _emblemOrigin EmblemOrigin // out

	_emblemOrigin = (EmblemOrigin)(_cret)

	return _emblemOrigin
}
