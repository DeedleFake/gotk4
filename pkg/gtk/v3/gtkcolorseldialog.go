// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_selection_dialog_get_type()), F: marshalColorSelectionDialogger},
	})
}

type ColorSelectionDialog struct {
	Dialog
}

func wrapColorSelectionDialog(obj *externglib.Object) *ColorSelectionDialog {
	return &ColorSelectionDialog{
		Dialog: Dialog{
			Window: Window{
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
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalColorSelectionDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorSelectionDialog(obj), nil
}

// NewColorSelectionDialog creates a new ColorSelectionDialog.
func NewColorSelectionDialog(title string) *ColorSelectionDialog {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_color_selection_dialog_new(_arg1)
	runtime.KeepAlive(title)

	var _colorSelectionDialog *ColorSelectionDialog // out

	_colorSelectionDialog = wrapColorSelectionDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _colorSelectionDialog
}

// ColorSelection retrieves the ColorSelection widget embedded in the dialog.
func (colorsel *ColorSelectionDialog) ColorSelection() Widgetter {
	var _arg0 *C.GtkColorSelectionDialog // out
	var _cret *C.GtkWidget               // in

	_arg0 = (*C.GtkColorSelectionDialog)(unsafe.Pointer(colorsel.Native()))

	_cret = C.gtk_color_selection_dialog_get_color_selection(_arg0)
	runtime.KeepAlive(colorsel)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}
