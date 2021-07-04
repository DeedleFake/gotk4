// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_selection_get_type()), F: marshalSelection},
	})
}

// Selection: Selection should be implemented by UI components with children
// which are exposed by #atk_object_ref_child and #atk_object_get_n_children, if
// the use of the parent UI component ordinarily involves selection of one or
// more of the objects corresponding to those Object children - for example,
// selectable lists.
//
// Note that other types of "selection" (for instance text selection) are
// accomplished a other ATK interfaces - Selection is limited to the
// selection/deselection of children.
type Selection interface {
	gextras.Objector

	// AddSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	AddSelection(i int) bool
	// ClearSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	ClearSelection() bool
	// SelectionCount causes every child of the object to be selected if the
	// object supports multiple selections.
	SelectionCount() int
	// IsChildSelected causes every child of the object to be selected if the
	// object supports multiple selections.
	IsChildSelected(i int) bool
	// RefSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	RefSelection(i int) Object
	// RemoveSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	RemoveSelection(i int) bool
	// SelectAllSelection causes every child of the object to be selected if the
	// object supports multiple selections.
	SelectAllSelection() bool
}

// selection implements the Selection interface.
type selection struct {
	gextras.Objector
}

var _ Selection = (*selection)(nil)

// WrapSelection wraps a GObject to a type that implements
// interface Selection. It is primarily used internally.
func WrapSelection(obj *externglib.Object) Selection {
	return selection{
		Objector: obj,
	}
}

func marshalSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSelection(obj), nil
}

func (s selection) AddSelection(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_add_selection(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s selection) ClearSelection() bool {
	var _arg0 *C.AtkSelection // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(s.Native()))

	_cret = C.atk_selection_clear_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s selection) SelectionCount() int {
	var _arg0 *C.AtkSelection // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(s.Native()))

	_cret = C.atk_selection_get_selection_count(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s selection) IsChildSelected(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_is_child_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s selection) RefSelection(i int) Object {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret *C.AtkObject    // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_ref_selection(_arg0, _arg1)

	var _object Object // out

	_object = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Object)

	return _object
}

func (s selection) RemoveSelection(i int) bool {
	var _arg0 *C.AtkSelection // out
	var _arg1 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_selection_remove_selection(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s selection) SelectAllSelection() bool {
	var _arg0 *C.AtkSelection // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkSelection)(unsafe.Pointer(s.Native()))

	_cret = C.atk_selection_select_all_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
