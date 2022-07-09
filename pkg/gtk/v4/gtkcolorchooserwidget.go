// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeColorChooserWidget returns the GType for the type ColorChooserWidget.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeColorChooserWidget() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ColorChooserWidget").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalColorChooserWidget)
	return gtype
}

// ColorChooserWidget: GtkColorChooserWidget widget lets the user select a
// color.
//
// By default, the chooser presents a predefined palette of colors, plus a small
// number of settable custom colors. It is also possible to select a different
// color with the single-color editor.
//
// To enter the single-color editing mode, use the context menu of any color of
// the palette, or use the '+' button to add a new custom color.
//
// The chooser automatically remembers the last selection, as well as custom
// colors.
//
// To create a GtkColorChooserWidget, use gtk.ColorChooserWidget.New.
//
// To change the initially selected color, use gtk.ColorChooser.SetRGBA(). To
// get the selected color use gtk.ColorChooser.GetRGBA().
//
// The GtkColorChooserWidget is used in the gtk.ColorChooserDialog to provide a
// dialog for selecting colors.
//
//
// CSS names
//
// GtkColorChooserWidget has a single CSS node with name colorchooser.
type ColorChooserWidget struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	ColorChooser
}

var (
	_ Widgetter         = (*ColorChooserWidget)(nil)
	_ coreglib.Objector = (*ColorChooserWidget)(nil)
)

func wrapColorChooserWidget(obj *coreglib.Object) *ColorChooserWidget {
	return &ColorChooserWidget{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		ColorChooser: ColorChooser{
			Object: obj,
		},
	}
}

func marshalColorChooserWidget(p uintptr) (interface{}, error) {
	return wrapColorChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewColorChooserWidget creates a new GtkColorChooserWidget.
//
// The function returns the following values:
//
//    - colorChooserWidget: new GtkColorChooserWidget.
//
func NewColorChooserWidget() *ColorChooserWidget {
	_info := girepository.MustFind("Gtk", "ColorChooserWidget")
	_gret := _info.InvokeClassMethod("new_ColorChooserWidget", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _colorChooserWidget *ColorChooserWidget // out

	_colorChooserWidget = wrapColorChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorChooserWidget
}
