// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_CellAccessibleClass_update_cache(GtkCellAccessible*, gboolean);
import "C"

// GType values.
var (
	GTypeCellAccessible = coreglib.Type(C.gtk_cell_accessible_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellAccessible, F: marshalCellAccessible},
	})
}

// CellAccessibleOverrider contains methods that are overridable.
type CellAccessibleOverrider interface {
	// The function takes the following parameters:
	//
	UpdateCache(emitSignal bool)
}

type CellAccessible struct {
	_ [0]func() // equal guard
	Accessible

	*coreglib.Object
	atk.Action
	atk.Component
	atk.ObjectClass
	atk.TableCell
}

var (
	_ coreglib.Objector = (*CellAccessible)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:        GTypeCellAccessible,
		GoType:       reflect.TypeOf((*CellAccessible)(nil)),
		InitClass:    initClassCellAccessible,
		ClassSize:    uint16(unsafe.Sizeof(C.GtkCellAccessible{})),
		InstanceSize: uint16(unsafe.Sizeof(C.GtkCellAccessibleClass{})),
	})
}

func initClassCellAccessible(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GtkCellAccessibleClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface{ UpdateCache(emitSignal bool) }); ok {
		pclass.update_cache = (*[0]byte)(C._gotk4_gtk3_CellAccessibleClass_update_cache)
	}
}

//export _gotk4_gtk3_CellAccessibleClass_update_cache
func _gotk4_gtk3_CellAccessibleClass_update_cache(arg0 *C.GtkCellAccessible, arg1 C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface{ UpdateCache(emitSignal bool) })

	var _emitSignal bool // out

	if arg1 != 0 {
		_emitSignal = true
	}

	iface.UpdateCache(_emitSignal)
}

func wrapCellAccessible(obj *coreglib.Object) *CellAccessible {
	return &CellAccessible{
		Accessible: Accessible{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Object: obj,
		Action: atk.Action{
			Object: obj,
		},
		Component: atk.Component{
			Object: obj,
		},
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		TableCell: atk.TableCell{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
	}
}

func marshalCellAccessible(p uintptr) (interface{}, error) {
	return wrapCellAccessible(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
