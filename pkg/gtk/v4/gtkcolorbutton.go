// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_ColorButton_ConnectColorSet(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeColorButton = coreglib.Type(C.gtk_color_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorButton, F: marshalColorButton},
	})
}

// ColorButton: GtkColorButton allows to open a color chooser dialog to change
// the color.
//
// !An example GtkColorButton (color-button.png)
//
// It is suitable widget for selecting a color in a preference dialog.
//
// CSS nodes
//
//    colorbutton
//    ╰── button.color
//        ╰── [content]
//
//
// GtkColorButton has a single CSS node with name colorbutton which contains a
// button node. To differentiate it from a plain GtkButton, it gets the .color
// style class.
type ColorButton struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	ColorChooser
}

var (
	_ Widgetter         = (*ColorButton)(nil)
	_ coreglib.Objector = (*ColorButton)(nil)
)

func wrapColorButton(obj *coreglib.Object) *ColorButton {
	return &ColorButton{
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

func marshalColorButton(p uintptr) (interface{}, error) {
	return wrapColorButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_ColorButton_ConnectColorSet
func _gotk4_gtk4_ColorButton_ConnectColorSet(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectColorSet is emitted when the user selects a color.
//
// When handling this signal, use gtk.ColorChooser.GetRGBA() to find out which
// color was just selected.
//
// Note that this signal is only emitted when the user changes the color. If you
// need to react to programmatic color changes as well, use the notify::color
// signal.
func (button *ColorButton) ConnectColorSet(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "color-set", false, unsafe.Pointer(C._gotk4_gtk4_ColorButton_ConnectColorSet), f)
}

// NewColorButton creates a new color button.
//
// This returns a widget in the form of a small button containing a swatch
// representing the current selected color. When the button is clicked, a color
// chooser dialog will open, allowing the user to select a color. The swatch
// will be updated to reflect the new color when the user finishes.
//
// The function returns the following values:
//
//    - colorButton: new color button.
//
func NewColorButton() *ColorButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_button_new()

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// NewColorButtonWithRGBA creates a new color button showing the given color.
//
// The function takes the following parameters:
//
//    - rgba: GdkRGBA to set the current color with.
//
// The function returns the following values:
//
//    - colorButton: new color button.
//
func NewColorButtonWithRGBA(rgba *gdk.RGBA) *ColorButton {
	var _arg1 *C.GdkRGBA   // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	_cret = C.gtk_color_button_new_with_rgba(_arg1)
	runtime.KeepAlive(rgba)

	var _colorButton *ColorButton // out

	_colorButton = wrapColorButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorButton
}

// Modal gets whether the dialog is modal.
//
// The function returns the following values:
//
//    - ok: TRUE if the dialog is modal.
//
func (button *ColorButton) Modal() bool {
	var _arg0 *C.GtkColorButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_color_button_get_modal(_arg0)
	runtime.KeepAlive(button)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the title of the color chooser dialog.
//
// The function returns the following values:
//
//    - utf8: internal string, do not free the return value.
//
func (button *ColorButton) Title() string {
	var _arg0 *C.GtkColorButton // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_color_button_get_title(_arg0)
	runtime.KeepAlive(button)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetModal sets whether the dialog should be modal.
//
// The function takes the following parameters:
//
//    - modal: TRUE to make the dialog modal.
//
func (button *ColorButton) SetModal(modal bool) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_color_button_set_modal(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(modal)
}

// SetTitle sets the title for the color chooser dialog.
//
// The function takes the following parameters:
//
//    - title: string containing new window title.
//
func (button *ColorButton) SetTitle(title string) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_color_button_set_title(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(title)
}
