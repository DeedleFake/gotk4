// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_widget_get_type()), F: marshalFileChooserWidgeter},
	})
}

// FileChooserWidget: GtkFileChooserWidget is a widget for choosing files.
//
// It exposes the gtk.FileChooser interface, and you should use the methods of
// this interface to interact with the widget.
//
//
// CSS nodes
//
// GtkFileChooserWidget has a single CSS node with name filechooser.
type FileChooserWidget struct {
	Widget

	FileChooser
}

var _ gextras.Nativer = (*FileChooserWidget)(nil)

func wrapFileChooserWidget(obj *externglib.Object) *FileChooserWidget {
	return &FileChooserWidget{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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
		FileChooser: FileChooser{
			Object: obj,
		},
	}
}

func marshalFileChooserWidgeter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileChooserWidget(obj), nil
}

// NewFileChooserWidget creates a new GtkFileChooserWidget.
//
// This is a file chooser widget that can be embedded in custom windows, and it
// is the same widget that is used by GtkFileChooserDialog.
func NewFileChooserWidget(action FileChooserAction) *FileChooserWidget {
	var _arg1 C.GtkFileChooserAction // out
	var _cret *C.GtkWidget           // in

	_arg1 = C.GtkFileChooserAction(action)

	_cret = C.gtk_file_chooser_widget_new(_arg1)

	var _fileChooserWidget *FileChooserWidget // out

	_fileChooserWidget = wrapFileChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _fileChooserWidget
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *FileChooserWidget) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

func (*FileChooserWidget) privateFileChooserWidget() {}
