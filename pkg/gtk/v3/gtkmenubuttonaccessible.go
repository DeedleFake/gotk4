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
		{T: externglib.Type(C.gtk_menu_button_accessible_get_type()), F: marshalMenuButtonAccessible},
	})
}

type MenuButtonAccessible interface {
	ToggleButtonAccessible
}

// menuButtonAccessible implements the MenuButtonAccessible class.
type menuButtonAccessible struct {
	ToggleButtonAccessible
}

// WrapMenuButtonAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuButtonAccessible(obj *externglib.Object) MenuButtonAccessible {
	return menuButtonAccessible{
		ToggleButtonAccessible: WrapToggleButtonAccessible(obj),
	}
}

func marshalMenuButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuButtonAccessible(obj), nil
}

func (a menuButtonAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a menuButtonAccessible) Description(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Description(i)
}

func (a menuButtonAccessible) Keybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Keybinding(i)
}

func (a menuButtonAccessible) LocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).LocalizedName(i)
}

func (a menuButtonAccessible) NActions() int {
	return atk.WrapAction(gextras.InternObject(a)).NActions()
}

func (a menuButtonAccessible) Name(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Name(i)
}

func (a menuButtonAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}

func (i menuButtonAccessible) ImageDescription() string {
	return atk.WrapImage(gextras.InternObject(i)).ImageDescription()
}

func (i menuButtonAccessible) ImageLocale() string {
	return atk.WrapImage(gextras.InternObject(i)).ImageLocale()
}

func (i menuButtonAccessible) ImagePosition(coordType atk.CoordType) (x int, y int) {
	return atk.WrapImage(gextras.InternObject(i)).ImagePosition(coordType)
}

func (i menuButtonAccessible) ImageSize() (width int, height int) {
	return atk.WrapImage(gextras.InternObject(i)).ImageSize()
}

func (i menuButtonAccessible) SetImageDescription(description string) bool {
	return atk.WrapImage(gextras.InternObject(i)).SetImageDescription(description)
}
