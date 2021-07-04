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
		{T: externglib.Type(C.gtk_expander_accessible_get_type()), F: marshalExpanderAccessible},
	})
}

type ExpanderAccessible interface {
	ContainerAccessible
	atk.Action
}

// expanderAccessible implements the ExpanderAccessible class.
type expanderAccessible struct {
	ContainerAccessible
}

// WrapExpanderAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapExpanderAccessible(obj *externglib.Object) ExpanderAccessible {
	return expanderAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalExpanderAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapExpanderAccessible(obj), nil
}

func (a expanderAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a expanderAccessible) Description(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Description(i)
}

func (a expanderAccessible) Keybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Keybinding(i)
}

func (a expanderAccessible) LocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).LocalizedName(i)
}

func (a expanderAccessible) NActions() int {
	return atk.WrapAction(gextras.InternObject(a)).NActions()
}

func (a expanderAccessible) Name(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Name(i)
}

func (a expanderAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}
