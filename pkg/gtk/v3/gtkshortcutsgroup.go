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
		{T: externglib.Type(C.gtk_shortcuts_group_get_type()), F: marshalShortcutsGroup},
	})
}

// ShortcutsGroup represents a group of related keyboard shortcuts or gestures.
// The group has a title. It may optionally be associated with a view of the
// application, which can be used to show only relevant shortcuts depending on
// the application context.
//
// This widget is only meant to be used with ShortcutsWindow.
type ShortcutsGroup interface {
	gextras.Objector

	privateShortcutsGroupClass()
}

// ShortcutsGroupClass implements the ShortcutsGroup interface.
type ShortcutsGroupClass struct {
	*externglib.Object
	BoxClass
	BuildableIface
	OrientableIface
}

var _ ShortcutsGroup = (*ShortcutsGroupClass)(nil)

func wrapShortcutsGroup(obj *externglib.Object) ShortcutsGroup {
	return &ShortcutsGroupClass{
		Object: obj,
		BoxClass: BoxClass{
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

func marshalShortcutsGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutsGroup(obj), nil
}

func (*ShortcutsGroupClass) privateShortcutsGroupClass() {}
