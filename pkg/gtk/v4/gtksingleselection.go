// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_single_selection_get_type()), F: marshalSingleSelectioner},
	})
}

// SingleSelection: GtkSingleSelection is a GtkSelectionModel that allows
// selecting a single item.
//
// Note that the selection is *persistent* -- if the selected item is removed
// and re-added in the same ::items-changed emission, it stays selected. In
// particular, this means that changing the sort order of an underlying sort
// model will preserve the selection.
type SingleSelection struct {
	*externglib.Object

	SelectionModel
}

func wrapSingleSelection(obj *externglib.Object) *SingleSelection {
	return &SingleSelection{
		Object: obj,
		SelectionModel: SelectionModel{
			ListModel: gio.ListModel{
				Object: obj,
			},
		},
	}
}

func marshalSingleSelectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSingleSelection(obj), nil
}

// NewSingleSelection creates a new selection to handle model.
func NewSingleSelection(model gio.ListModeller) *SingleSelection {
	var _arg1 *C.GListModel         // out
	var _cret *C.GtkSingleSelection // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}

	_cret = C.gtk_single_selection_new(_arg1)
	runtime.KeepAlive(model)

	var _singleSelection *SingleSelection // out

	_singleSelection = wrapSingleSelection(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _singleSelection
}

// Autoselect checks if autoselect has been enabled or disabled via
// gtk_single_selection_set_autoselect().
func (self *SingleSelection) Autoselect() bool {
	var _arg0 *C.GtkSingleSelection // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_autoselect(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanUnselect: if TRUE, gtk_selection_model_unselect_item() is supported and
// allows unselecting the selected item.
func (self *SingleSelection) CanUnselect() bool {
	var _arg0 *C.GtkSingleSelection // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_can_unselect(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Model gets the model that self is wrapping.
func (self *SingleSelection) Model() gio.ListModeller {
	var _arg0 *C.GtkSingleSelection // out
	var _cret *C.GListModel         // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	_listModel = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// Selected gets the position of the selected item.
//
// If no item is selected, GTK_INVALID_LIST_POSITION is returned.
func (self *SingleSelection) Selected() uint {
	var _arg0 *C.GtkSingleSelection // out
	var _cret C.guint               // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_selected(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SelectedItem gets the selected item.
//
// If no item is selected, NULL is returned.
func (self *SingleSelection) SelectedItem() *externglib.Object {
	var _arg0 *C.GtkSingleSelection // out
	var _cret C.gpointer            // in

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_single_selection_get_selected_item(_arg0)
	runtime.KeepAlive(self)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// SetAutoselect enables or disables autoselect.
//
// If autoselect is TRUE, self will enforce that an item is always selected. It
// will select a new item when the currently selected item is deleted and it
// will disallow unselecting the current item.
func (self *SingleSelection) SetAutoselect(autoselect bool) {
	var _arg0 *C.GtkSingleSelection // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))
	if autoselect {
		_arg1 = C.TRUE
	}

	C.gtk_single_selection_set_autoselect(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(autoselect)
}

// SetCanUnselect: if TRUE, unselecting the current item via
// gtk_selection_model_unselect_item() is supported.
//
// Note that setting gtk.SingleSelection:autoselect will cause unselecting to
// not work, so it practically makes no sense to set both at the same time the
// same time.
func (self *SingleSelection) SetCanUnselect(canUnselect bool) {
	var _arg0 *C.GtkSingleSelection // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))
	if canUnselect {
		_arg1 = C.TRUE
	}

	C.gtk_single_selection_set_can_unselect(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canUnselect)
}

// SetModel sets the model that self should wrap.
//
// If model is NULL, self will be empty.
func (self *SingleSelection) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkSingleSelection // out
	var _arg1 *C.GListModel         // out

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_single_selection_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetSelected selects the item at the given position.
//
// If the list does not have an item at position or GTK_INVALID_LIST_POSITION is
// given, the behavior depends on the value of the
// gtk.SingleSelection:autoselect property: If it is set, no change will occur
// and the old item will stay selected. If it is unset, the selection will be
// unset and no item will be selected.
func (self *SingleSelection) SetSelected(position uint) {
	var _arg0 *C.GtkSingleSelection // out
	var _arg1 C.guint               // out

	_arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	C.gtk_single_selection_set_selected(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}
