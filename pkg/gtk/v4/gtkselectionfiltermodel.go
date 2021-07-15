// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_selection_filter_model_get_type()), F: marshalSelectionFilterModeller},
	})
}

// SelectionFilterModel: GtkSelectionFilterModel is a list model that presents
// the selection from a GtkSelectionModel.
type SelectionFilterModel struct {
	*externglib.Object

	gio.ListModel
}

var _ gextras.Nativer = (*SelectionFilterModel)(nil)

func wrapSelectionFilterModel(obj *externglib.Object) *SelectionFilterModel {
	return &SelectionFilterModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalSelectionFilterModeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSelectionFilterModel(obj), nil
}

// NewSelectionFilterModel creates a new GtkSelectionFilterModel that will
// include the selected items from the underlying selection model.
func NewSelectionFilterModel(model SelectionModeller) *SelectionFilterModel {
	var _arg1 *C.GtkSelectionModel       // out
	var _cret *C.GtkSelectionFilterModel // in

	_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	_cret = C.gtk_selection_filter_model_new(_arg1)

	var _selectionFilterModel *SelectionFilterModel // out

	_selectionFilterModel = wrapSelectionFilterModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _selectionFilterModel
}

// Model gets the model currently filtered or NULL if none.
func (self *SelectionFilterModel) Model() *SelectionModel {
	var _arg0 *C.GtkSelectionFilterModel // out
	var _cret *C.GtkSelectionModel       // in

	_arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_selection_filter_model_get_model(_arg0)

	var _selectionModel *SelectionModel // out

	_selectionModel = wrapSelectionModel(externglib.Take(unsafe.Pointer(_cret)))

	return _selectionModel
}

// SetModel sets the model to be filtered.
//
// Note that GTK makes no effort to ensure that model conforms to the item type
// of self. It assumes that the caller knows what they are doing and have set up
// an appropriate filter to ensure that item types match.
func (self *SelectionFilterModel) SetModel(model SelectionModeller) {
	var _arg0 *C.GtkSelectionFilterModel // out
	var _arg1 *C.GtkSelectionModel       // out

	_arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	C.gtk_selection_filter_model_set_model(_arg0, _arg1)
}
