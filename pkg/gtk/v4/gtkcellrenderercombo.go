// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_combo_get_type()), F: marshalCellRendererCombo},
	})
}

// CellRendererCombo renders a combobox in a cell
//
// CellRendererCombo renders text in a cell like CellRendererText from which it
// is derived. But while CellRendererText offers a simple entry to edit the
// text, CellRendererCombo offers a ComboBox widget to edit the text. The values
// to display in the combo box are taken from the tree model specified in the
// CellRendererCombo:model property.
//
// The combo cell renderer takes care of adding a text cell renderer to the
// combo box and sets it to display the column specified by its
// CellRendererCombo:text-column property. Further properties of the combo box
// can be set in a handler for the CellRenderer::editing-started signal.
type CellRendererCombo interface {
	gextras.Objector

	privateCellRendererComboClass()
}

// CellRendererComboClass implements the CellRendererCombo interface.
type CellRendererComboClass struct {
	CellRendererTextClass
}

var _ CellRendererCombo = (*CellRendererComboClass)(nil)

func wrapCellRendererCombo(obj *externglib.Object) CellRendererCombo {
	return &CellRendererComboClass{
		CellRendererTextClass: CellRendererTextClass{
			CellRendererClass: CellRendererClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			},
		},
	}
}

func marshalCellRendererCombo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRendererCombo(obj), nil
}

// NewCellRendererCombo creates a new CellRendererCombo. Adjust how text is
// drawn using object properties. Object properties can be set globally (with
// g_object_set()). Also, with TreeViewColumn, you can bind a property to a
// value in a TreeModel. For example, you can bind the “text” property on the
// cell renderer to a string value in the model, thus rendering a different
// string in each row of the TreeView.
func NewCellRendererCombo() *CellRendererComboClass {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_combo_new()

	var _cellRendererCombo *CellRendererComboClass // out

	_cellRendererCombo = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*CellRendererComboClass)

	return _cellRendererCombo
}

func (*CellRendererComboClass) privateCellRendererComboClass() {}
