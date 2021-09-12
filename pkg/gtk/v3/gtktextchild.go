// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_child_anchor_get_type()), F: marshalTextChildAnchorrer},
	})
}

// TextChildAnchor is a spot in the buffer where child widgets can be “anchored”
// (inserted inline, as if they were characters). The anchor can have multiple
// widgets anchored, to allow for multiple views.
type TextChildAnchor struct {
	*externglib.Object
}

func wrapTextChildAnchor(obj *externglib.Object) *TextChildAnchor {
	return &TextChildAnchor{
		Object: obj,
	}
}

func marshalTextChildAnchorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTextChildAnchor(obj), nil
}

// NewTextChildAnchor creates a new TextChildAnchor. Usually you would then
// insert it into a TextBuffer with gtk_text_buffer_insert_child_anchor(). To
// perform the creation and insertion in one step, use the convenience function
// gtk_text_buffer_create_child_anchor().
func NewTextChildAnchor() *TextChildAnchor {
	var _cret *C.GtkTextChildAnchor // in

	_cret = C.gtk_text_child_anchor_new()

	var _textChildAnchor *TextChildAnchor // out

	_textChildAnchor = wrapTextChildAnchor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textChildAnchor
}

// Deleted determines whether a child anchor has been deleted from the buffer.
// Keep in mind that the child anchor will be unreferenced when removed from the
// buffer, so you need to hold your own reference (with g_object_ref()) if you
// plan to use this function — otherwise all deleted child anchors will also be
// finalized.
func (anchor *TextChildAnchor) Deleted() bool {
	var _arg0 *C.GtkTextChildAnchor // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(anchor.Native()))

	_cret = C.gtk_text_child_anchor_get_deleted(_arg0)
	runtime.KeepAlive(anchor)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Widgets gets a list of all widgets anchored at this child anchor. The
// returned list should be freed with g_list_free().
func (anchor *TextChildAnchor) Widgets() []Widgetter {
	var _arg0 *C.GtkTextChildAnchor // out
	var _cret *C.GList              // in

	_arg0 = (*C.GtkTextChildAnchor)(unsafe.Pointer(anchor.Native()))

	_cret = C.gtk_text_child_anchor_get_widgets(_arg0)
	runtime.KeepAlive(anchor)

	var _list []Widgetter // out

	_list = make([]Widgetter, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkWidget)(v)
		var dst Widgetter // out
		dst = (externglib.CastObject(externglib.Take(unsafe.Pointer(src)))).(Widgetter)
		_list = append(_list, dst)
	})

	return _list
}
