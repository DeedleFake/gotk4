// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_emblem_get_type()), F: marshalEmblem},
	})
}

// Emblem is an implementation of #GIcon that supports having an emblem, which
// is an icon with additional properties. It can than be added to a Icon.
//
// Currently, only metainformation about the emblem's origin is supported. More
// may be added in the future.
type Emblem interface {
	Icon

	GetIcon() Icon

	Origin() EmblemOrigin
}

// emblem implements the Emblem class.
type emblem struct {
	gextras.Objector
}

// WrapEmblem wraps a GObject to the right type. It is
// primarily used internally.
func WrapEmblem(obj *externglib.Object) Emblem {
	return emblem{
		Objector: obj,
	}
}

func marshalEmblem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEmblem(obj), nil
}

func NewEmblem(icon Icon) Emblem {
	var _arg1 *C.GIcon   // out
	var _cret *C.GEmblem // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.g_emblem_new(_arg1)

	var _emblem Emblem // out

	_emblem = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Emblem)

	return _emblem
}

func NewEmblemWithOrigin(icon Icon, origin EmblemOrigin) Emblem {
	var _arg1 *C.GIcon        // out
	var _arg2 C.GEmblemOrigin // out
	var _cret *C.GEmblem      // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg2 = C.GEmblemOrigin(origin)

	_cret = C.g_emblem_new_with_origin(_arg1, _arg2)

	var _emblem Emblem // out

	_emblem = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Emblem)

	return _emblem
}

func (e emblem) GetIcon() Icon {
	var _arg0 *C.GEmblem // out
	var _cret *C.GIcon   // in

	_arg0 = (*C.GEmblem)(unsafe.Pointer(e.Native()))

	_cret = C.g_emblem_get_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

func (e emblem) Origin() EmblemOrigin {
	var _arg0 *C.GEmblem      // out
	var _cret C.GEmblemOrigin // in

	_arg0 = (*C.GEmblem)(unsafe.Pointer(e.Native()))

	_cret = C.g_emblem_get_origin(_arg0)

	var _emblemOrigin EmblemOrigin // out

	_emblemOrigin = EmblemOrigin(_cret)

	return _emblemOrigin
}

func (i emblem) Equal(icon2 Icon) bool {
	return WrapIcon(gextras.InternObject(i)).Equal(icon2)
}

func (i emblem) Serialize() *glib.Variant {
	return WrapIcon(gextras.InternObject(i)).Serialize()
}

func (i emblem) String() string {
	return WrapIcon(gextras.InternObject(i)).String()
}
