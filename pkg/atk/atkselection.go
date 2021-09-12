// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "atk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_selection_get_type()), F: marshalSelectioner},
	})
}

// SelectionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SelectionOverrider interface {
	// AddSelection adds the specified accessible child of the object to the
	// object's selection.
	AddSelection(i int) bool
	// ClearSelection clears the selection in the object so that no children in
	// the object are selected.
	ClearSelection() bool
	// SelectionCount gets the number of accessible children currently selected.
	// Note: callers should not rely on NULL or on a zero value for indication
	// of whether AtkSelectionIface is implemented, they should use type
	// checking/interface checking macros or the atk_get_accessible_value()
	// convenience method.
	SelectionCount() int
	// IsChildSelected determines if the current child of this object is
	// selected Note: callers should not rely on NULL or on a zero value for
	// indication of whether AtkSelectionIface is implemented, they should use
	// type checking/interface checking macros or the atk_get_accessible_value()
	// convenience method.
	IsChildSelected(i int) bool
	// RefSelection gets a reference to the accessible object representing the
	// specified selected child of the object. Note: callers should not rely on
	// NULL or on a zero value for indication of whether AtkSelectionIface is
	// implemented, they should use type checking/interface checking macros or
	// the atk_get_accessible_value() convenience method.
	RefSelection(i int) *ObjectClass
	// RemoveSelection removes the specified child of the object from the
	// object's selection.
	RemoveSelection(i int) bool
	// SelectAllSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	SelectAllSelection() bool
	SelectionChanged()
}

// Selection should be implemented by UI components with children which are
// exposed by #atk_object_ref_child and #atk_object_get_n_children, if the use
// of the parent UI component ordinarily involves selection of one or more of
// the objects corresponding to those Object children - for example, selectable
// lists.
//
// Note that other types of "selection" (for instance text selection) are
// accomplished a other ATK interfaces - Selection is limited to the
// selection/deselection of children.
type Selection struct {
	*externglib.Object
}

// Selectioner describes Selection's abstract methods.
type Selectioner interface {
	externglib.Objector

	// AddSelection adds the specified accessible child of the object to the
	// object's selection.
	AddSelection(i int) bool
	// ClearSelection clears the selection in the object so that no children in
	// the object are selected.
	ClearSelection() bool
	// SelectionCount gets the number of accessible children currently selected.
	SelectionCount() int
	// IsChildSelected determines if the current child of this object is
	// selected Note: callers should not rely on NULL or on a zero value for
	// indication of whether AtkSelectionIface is implemented, they should use
	// type checking/interface checking macros or the atk_get_accessible_value()
	// convenience method.
	IsChildSelected(i int) bool
	// RefSelection gets a reference to the accessible object representing the
	// specified selected child of the object.
	RefSelection(i int) *ObjectClass
	// RemoveSelection removes the specified child of the object from the
	// object's selection.
	RemoveSelection(i int) bool
	// SelectAllSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	SelectAllSelection() bool
}

var _ Selectioner = (*Selection)(nil)

func wrapSelection(obj *externglib.Object) *Selection {
	return &Selection{
		Object: obj,
	}
}

func marshalSelectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSelection(obj), nil
}

// AddSelection adds the specified accessible child of the object to the
// object's selection.
func (selection *Selection) AddSelection(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_add_selection(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(i)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ClearSelection clears the selection in the object so that no children in the
// object are selected.
func (selection *Selection) ClearSelection() bool {
	var _arg0 *C.AtkSelection // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(selection.Native()))

	_cret = C.atk_selection_clear_selection(_arg0)
	runtime.KeepAlive(selection)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionCount gets the number of accessible children currently selected.
// Note: callers should not rely on NULL or on a zero value for indication of
// whether AtkSelectionIface is implemented, they should use type
// checking/interface checking macros or the atk_get_accessible_value()
// convenience method.
func (selection *Selection) SelectionCount() int {
	var _arg0 *C.AtkSelection // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(selection.Native()))

	_cret = C.atk_selection_get_selection_count(_arg0)
	runtime.KeepAlive(selection)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsChildSelected determines if the current child of this object is selected
// Note: callers should not rely on NULL or on a zero value for indication of
// whether AtkSelectionIface is implemented, they should use type
// checking/interface checking macros or the atk_get_accessible_value()
// convenience method.
func (selection *Selection) IsChildSelected(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_is_child_selected(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(i)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RefSelection gets a reference to the accessible object representing the
// specified selected child of the object. Note: callers should not rely on NULL
// or on a zero value for indication of whether AtkSelectionIface is
// implemented, they should use type checking/interface checking macros or the
// atk_get_accessible_value() convenience method.
func (selection *Selection) RefSelection(i int) *ObjectClass {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret *C.AtkObject    // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_ref_selection(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(i)

	var _object *ObjectClass // out

	if _cret != nil {
		_object = wrapObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _object
}

// RemoveSelection removes the specified child of the object from the object's
// selection.
func (selection *Selection) RemoveSelection(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(selection.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_remove_selection(_arg0, _arg1)
	runtime.KeepAlive(selection)
	runtime.KeepAlive(i)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectAllSelection causes every child of the object to be selected if the
// object supports multiple selections.
func (selection *Selection) SelectAllSelection() bool {
	var _arg0 *C.AtkSelection // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(selection.Native()))

	_cret = C.atk_selection_select_all_selection(_arg0)
	runtime.KeepAlive(selection)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
