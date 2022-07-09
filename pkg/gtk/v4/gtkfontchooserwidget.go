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

// GTypeFontChooserWidget returns the GType for the type FontChooserWidget.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontChooserWidget() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "FontChooserWidget").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontChooserWidget)
	return gtype
}

// FontChooserWidget: GtkFontChooserWidget widget lets the user select a font.
//
// It is used in the GtkFontChooserDialog widget to provide a dialog for
// selecting fonts.
//
// To set the font which is initially selected, use gtk.FontChooser.SetFont() or
// gtk.FontChooser.SetFontDesc().
//
// To get the selected font use gtk.FontChooser.GetFont() or
// gtk.FontChooser.GetFontDesc().
//
// To change the text which is shown in the preview area, use
// gtk.FontChooser.SetPreviewText().
//
//
// CSS nodes
//
// GtkFontChooserWidget has a single CSS node with name fontchooser.
type FontChooserWidget struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	FontChooser
}

var (
	_ Widgetter         = (*FontChooserWidget)(nil)
	_ coreglib.Objector = (*FontChooserWidget)(nil)
)

func wrapFontChooserWidget(obj *coreglib.Object) *FontChooserWidget {
	return &FontChooserWidget{
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
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontChooserWidget(p uintptr) (interface{}, error) {
	return wrapFontChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFontChooserWidget creates a new GtkFontChooserWidget.
//
// The function returns the following values:
//
//    - fontChooserWidget: new GtkFontChooserWidget.
//
func NewFontChooserWidget() *FontChooserWidget {
	_info := girepository.MustFind("Gtk", "FontChooserWidget")
	_gret := _info.InvokeClassMethod("new_FontChooserWidget", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _fontChooserWidget *FontChooserWidget // out

	_fontChooserWidget = wrapFontChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserWidget
}
