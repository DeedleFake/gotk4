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
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_editable_get_type()), F: marshalCellEditabler},
	})
}

// CellEditableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellEditableOverrider interface {
	// EditingDone emits the CellEditable::editing-done signal.
	EditingDone()
	// RemoveWidget emits the CellEditable::remove-widget signal.
	RemoveWidget()
}

// CellEditable interface must be implemented for widgets to be usable to edit
// the contents of a TreeView cell. It provides a way to specify how temporary
// widgets should be configured for editing, get the new value, etc.
type CellEditable struct {
	Widget
}

var _ gextras.Nativer = (*CellEditable)(nil)

// CellEditabler describes CellEditable's abstract methods.
type CellEditabler interface {
	// EditingDone emits the CellEditable::editing-done signal.
	EditingDone()
	// RemoveWidget emits the CellEditable::remove-widget signal.
	RemoveWidget()
}

var _ CellEditabler = (*CellEditable)(nil)

func wrapCellEditable(obj *externglib.Object) *CellEditable {
	return &CellEditable{
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
	}
}

func marshalCellEditabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellEditable(obj), nil
}

// EditingDone emits the CellEditable::editing-done signal.
func (cellEditable *CellEditable) EditingDone() {
	var _arg0 *C.GtkCellEditable // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(cellEditable.Native()))

	C.gtk_cell_editable_editing_done(_arg0)
}

// RemoveWidget emits the CellEditable::remove-widget signal.
func (cellEditable *CellEditable) RemoveWidget() {
	var _arg0 *C.GtkCellEditable // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(cellEditable.Native()))

	C.gtk_cell_editable_remove_widget(_arg0)
}
