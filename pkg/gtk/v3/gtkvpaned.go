// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_vpaned_get_type()), F: marshalVPaned},
	})
}

// VPaned: the VPaned widget is a container widget with two children arranged
// vertically. The division between the two panes is adjustable by the user by
// dragging a handle. See Paned for details.
//
// GtkVPaned has been deprecated, use Paned instead.
type VPaned interface {
	gextras.Objector

	privateVPanedClass()
}

// VPanedClass implements the VPaned interface.
type VPanedClass struct {
	*externglib.Object
	PanedClass
	BuildableIface
	OrientableIface
}

var _ VPaned = (*VPanedClass)(nil)

func wrapVPaned(obj *externglib.Object) VPaned {
	return &VPanedClass{
		Object: obj,
		PanedClass: PanedClass{
			Object: obj,
			ContainerClass: ContainerClass{
				Object: obj,
				WidgetClass: WidgetClass{
					Object: obj,
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					BuildableIface: BuildableIface{
						Object: obj,
					},
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
			OrientableIface: OrientableIface{
				Object: obj,
			},
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
		OrientableIface: OrientableIface{
			Object: obj,
		},
	}
}

func marshalVPaned(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVPaned(obj), nil
}

// NewVPaned: create a new VPaned
//
// Deprecated: Use gtk_paned_new() with GTK_ORIENTATION_VERTICAL instead.
func NewVPaned() *VPanedClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vpaned_new()

	var _vPaned *VPanedClass // out

	_vPaned = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*VPanedClass)

	return _vPaned
}

func (*VPanedClass) privateVPanedClass() {}
