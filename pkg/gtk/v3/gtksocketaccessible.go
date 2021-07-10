// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_socket_accessible_get_type()), F: marshalSocketAccessible},
	})
}

type SocketAccessible interface {
	gextras.Objector

	Embed(path string)
}

// SocketAccessibleClass implements the SocketAccessible interface.
type SocketAccessibleClass struct {
	ContainerAccessibleClass
}

var _ SocketAccessible = (*SocketAccessibleClass)(nil)

func wrapSocketAccessible(obj *externglib.Object) SocketAccessible {
	return &SocketAccessibleClass{
		ContainerAccessibleClass: ContainerAccessibleClass{
			WidgetAccessibleClass: WidgetAccessibleClass{
				AccessibleClass: AccessibleClass{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalSocketAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketAccessible(obj), nil
}

func (socket *SocketAccessibleClass) Embed(path string) {
	var _arg0 *C.GtkSocketAccessible // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkSocketAccessible)(unsafe.Pointer(socket.Native()))
	_arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_socket_accessible_embed(_arg0, _arg1)
}
