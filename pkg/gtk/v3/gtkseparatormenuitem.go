// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeSeparatorMenuItem = coreglib.Type(C.gtk_separator_menu_item_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSeparatorMenuItem, F: marshalSeparatorMenuItem},
	})
}

// SeparatorMenuItemOverrides contains methods that are overridable.
type SeparatorMenuItemOverrides struct {
}

func defaultSeparatorMenuItemOverrides(v *SeparatorMenuItem) SeparatorMenuItemOverrides {
	return SeparatorMenuItemOverrides{}
}

// SeparatorMenuItem is a separator used to group items within a menu. It
// displays a horizontal line with a shadow to make it appear sunken into the
// interface.
//
//
// CSS nodes
//
// GtkSeparatorMenuItem has a single CSS node with name separator.
type SeparatorMenuItem struct {
	_ [0]func() // equal guard
	MenuItem
}

var (
	_ Binner            = (*SeparatorMenuItem)(nil)
	_ coreglib.Objector = (*SeparatorMenuItem)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SeparatorMenuItem, *SeparatorMenuItemClass, SeparatorMenuItemOverrides](
		GTypeSeparatorMenuItem,
		initSeparatorMenuItemClass,
		wrapSeparatorMenuItem,
		defaultSeparatorMenuItemOverrides,
	)
}

func initSeparatorMenuItemClass(gclass unsafe.Pointer, overrides SeparatorMenuItemOverrides, classInitFunc func(*SeparatorMenuItemClass)) {
	if classInitFunc != nil {
		class := (*SeparatorMenuItemClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSeparatorMenuItem(obj *coreglib.Object) *SeparatorMenuItem {
	return &SeparatorMenuItem{
		MenuItem: MenuItem{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
			Object: obj,
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalSeparatorMenuItem(p uintptr) (interface{}, error) {
	return wrapSeparatorMenuItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSeparatorMenuItem creates a new SeparatorMenuItem.
//
// The function returns the following values:
//
//    - separatorMenuItem: new SeparatorMenuItem.
//
func NewSeparatorMenuItem() *SeparatorMenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_separator_menu_item_new()

	var _separatorMenuItem *SeparatorMenuItem // out

	_separatorMenuItem = wrapSeparatorMenuItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _separatorMenuItem
}

// SeparatorMenuItemClass: instance of this type is always passed by reference.
type SeparatorMenuItemClass struct {
	*separatorMenuItemClass
}

// separatorMenuItemClass is the struct that's finalized.
type separatorMenuItemClass struct {
	native *C.GtkSeparatorMenuItemClass
}

// ParentClass: parent class.
func (s *SeparatorMenuItemClass) ParentClass() *MenuItemClass {
	valptr := &s.native.parent_class
	var _v *MenuItemClass // out
	_v = (*MenuItemClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
