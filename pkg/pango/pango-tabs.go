// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "pango.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_tab_align_get_type()), F: marshalTabAlign},
		{T: externglib.Type(C.pango_tab_array_get_type()), F: marshalTabArray},
	})
}

// TabAlign: PangoTabAlign specifies where a tab stop appears relative to the
// text.
type TabAlign int

const (
	// TabLeft: tab stop appears to the left of the text.
	TabLeft TabAlign = iota
)

func marshalTabAlign(p uintptr) (interface{}, error) {
	return TabAlign(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for TabAlign.
func (t TabAlign) String() string {
	switch t {
	case TabLeft:
		return "Left"
	default:
		return fmt.Sprintf("TabAlign(%d)", t)
	}
}

// TabArray: PangoTabArray contains an array of tab stops.
//
// PangoTabArray can be used to set tab stops in a PangoLayout. Each tab stop
// has an alignment and a position.
//
// An instance of this type is always passed by reference.
type TabArray struct {
	*tabArray
}

// tabArray is the struct that's finalized.
type tabArray struct {
	native *C.PangoTabArray
}

func marshalTabArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &TabArray{&tabArray{(*C.PangoTabArray)(unsafe.Pointer(b))}}, nil
}

// NewTabArray constructs a struct TabArray.
func NewTabArray(initialSize int, positionsInPixels bool) *TabArray {
	var _arg1 C.gint           // out
	var _arg2 C.gboolean       // out
	var _cret *C.PangoTabArray // in

	_arg1 = C.gint(initialSize)
	if positionsInPixels {
		_arg2 = C.TRUE
	}

	_cret = C.pango_tab_array_new(_arg1, _arg2)
	runtime.KeepAlive(initialSize)
	runtime.KeepAlive(positionsInPixels)

	var _tabArray *TabArray // out

	_tabArray = (*TabArray)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_tabArray)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_tab_array_free((*C.PangoTabArray)(intern.C))
		},
	)

	return _tabArray
}

// Copy copies a PangoTabArray.
func (src *TabArray) Copy() *TabArray {
	var _arg0 *C.PangoTabArray // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.PangoTabArray)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.pango_tab_array_copy(_arg0)
	runtime.KeepAlive(src)

	var _tabArray *TabArray // out

	_tabArray = (*TabArray)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_tabArray)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_tab_array_free((*C.PangoTabArray)(intern.C))
		},
	)

	return _tabArray
}

// PositionsInPixels returns TRUE if the tab positions are in pixels, FALSE if
// they are in Pango units.
func (tabArray *TabArray) PositionsInPixels() bool {
	var _arg0 *C.PangoTabArray // out
	var _cret C.gboolean       // in

	_arg0 = (*C.PangoTabArray)(gextras.StructNative(unsafe.Pointer(tabArray)))

	_cret = C.pango_tab_array_get_positions_in_pixels(_arg0)
	runtime.KeepAlive(tabArray)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Size gets the number of tab stops in tab_array.
func (tabArray *TabArray) Size() int {
	var _arg0 *C.PangoTabArray // out
	var _cret C.gint           // in

	_arg0 = (*C.PangoTabArray)(gextras.StructNative(unsafe.Pointer(tabArray)))

	_cret = C.pango_tab_array_get_size(_arg0)
	runtime.KeepAlive(tabArray)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Tab gets the alignment and position of a tab stop.
func (tabArray *TabArray) Tab(tabIndex int) (TabAlign, int) {
	var _arg0 *C.PangoTabArray // out
	var _arg1 C.gint           // out
	var _arg2 C.PangoTabAlign  // in
	var _arg3 C.gint           // in

	_arg0 = (*C.PangoTabArray)(gextras.StructNative(unsafe.Pointer(tabArray)))
	_arg1 = C.gint(tabIndex)

	C.pango_tab_array_get_tab(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(tabArray)
	runtime.KeepAlive(tabIndex)

	var _alignment TabAlign // out
	var _location int       // out

	_alignment = TabAlign(_arg2)
	_location = int(_arg3)

	return _alignment, _location
}

// Resize resizes a tab array.
//
// You must subsequently initialize any tabs that were added as a result of
// growing the array.
func (tabArray *TabArray) Resize(newSize int) {
	var _arg0 *C.PangoTabArray // out
	var _arg1 C.gint           // out

	_arg0 = (*C.PangoTabArray)(gextras.StructNative(unsafe.Pointer(tabArray)))
	_arg1 = C.gint(newSize)

	C.pango_tab_array_resize(_arg0, _arg1)
	runtime.KeepAlive(tabArray)
	runtime.KeepAlive(newSize)
}

// SetTab sets the alignment and location of a tab stop.
//
// alignment must always be PANGO_TAB_LEFT in the current implementation.
func (tabArray *TabArray) SetTab(tabIndex int, alignment TabAlign, location int) {
	var _arg0 *C.PangoTabArray // out
	var _arg1 C.gint           // out
	var _arg2 C.PangoTabAlign  // out
	var _arg3 C.gint           // out

	_arg0 = (*C.PangoTabArray)(gextras.StructNative(unsafe.Pointer(tabArray)))
	_arg1 = C.gint(tabIndex)
	_arg2 = C.PangoTabAlign(alignment)
	_arg3 = C.gint(location)

	C.pango_tab_array_set_tab(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(tabArray)
	runtime.KeepAlive(tabIndex)
	runtime.KeepAlive(alignment)
	runtime.KeepAlive(location)
}
