// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_check_button_get_type()), F: marshalCheckButtonner},
	})
}

// CheckButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CheckButtonOverrider interface {
	DrawIndicator(cr *cairo.Context)
}

// CheckButton places a discrete ToggleButton next to a widget, (usually a
// Label). See the section on ToggleButton widgets for more information about
// toggle/check buttons.
//
// The important signal ( ToggleButton::toggled ) is also inherited from
// ToggleButton.
//
// CSS nodes
//
//    button.check
//    ├── check
//    ╰── <child>
//
// A GtkCheckButton without indicator changes the name of its main node to
// button and adds a .check style class to it. The subnode is invisible in this
// case.
type CheckButton struct {
	ToggleButton
}

var _ gextras.Nativer = (*CheckButton)(nil)

func wrapCheckButton(obj *externglib.Object) *CheckButton {
	return &CheckButton{
		ToggleButton: ToggleButton{
			Button: Button{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
						},
					},
				},
				Actionable: Actionable{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
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
		},
	}
}

func marshalCheckButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCheckButton(obj), nil
}

// NewCheckButton creates a new CheckButton.
func NewCheckButton() *CheckButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_check_button_new()

	var _checkButton *CheckButton // out

	_checkButton = wrapCheckButton(externglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}

// NewCheckButtonWithLabel creates a new CheckButton with a Label to the right
// of it.
func NewCheckButtonWithLabel(label string) *CheckButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))

	_cret = C.gtk_check_button_new_with_label(_arg1)

	var _checkButton *CheckButton // out

	_checkButton = wrapCheckButton(externglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}

// NewCheckButtonWithMnemonic creates a new CheckButton containing a label. The
// label will be created using gtk_label_new_with_mnemonic(), so underscores in
// label indicate the mnemonic for the check button.
func NewCheckButtonWithMnemonic(label string) *CheckButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))

	_cret = C.gtk_check_button_new_with_mnemonic(_arg1)

	var _checkButton *CheckButton // out

	_checkButton = wrapCheckButton(externglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}

func (*CheckButton) privateCheckButton() {}
