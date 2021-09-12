// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_tag_get_type()), F: marshalTextTagger},
	})
}

// TextTag: tag that can be applied to text contained in a GtkTextBuffer.
//
// You may wish to begin by reading the text widget conceptual overview
// (section-text-widget.html), which gives an overview of all the objects and
// data types related to the text widget and how they work together.
//
// Tags should be in the gtk.TextTagTable for a given GtkTextBuffer before using
// them with that buffer.
//
// gtk.TextBuffer.CreateTag() is the best way to create tags. See “gtk4-demo”
// for numerous examples.
//
// For each property of GtkTextTag, there is a “set” property, e.g. “font-set”
// corresponds to “font”. These “set” properties reflect whether a property has
// been set or not.
//
// They are maintained by GTK and you should not set them independently.
type TextTag struct {
	*externglib.Object
}

func wrapTextTag(obj *externglib.Object) *TextTag {
	return &TextTag{
		Object: obj,
	}
}

func marshalTextTagger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTextTag(obj), nil
}

// NewTextTag creates a GtkTextTag.
func NewTextTag(name string) *TextTag {
	var _arg1 *C.char       // out
	var _cret *C.GtkTextTag // in

	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_text_tag_new(_arg1)
	runtime.KeepAlive(name)

	var _textTag *TextTag // out

	_textTag = wrapTextTag(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textTag
}

// Changed emits the gtk.TextTagTable::tag-changed signal on the GtkTextTagTable
// where the tag is included.
//
// The signal is already emitted when setting a GtkTextTag property. This
// function is useful for a GtkTextTag subclass.
func (tag *TextTag) Changed(sizeChanged bool) {
	var _arg0 *C.GtkTextTag // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))
	if sizeChanged {
		_arg1 = C.TRUE
	}

	C.gtk_text_tag_changed(_arg0, _arg1)
	runtime.KeepAlive(tag)
	runtime.KeepAlive(sizeChanged)
}

// Priority: get the tag priority.
func (tag *TextTag) Priority() int {
	var _arg0 *C.GtkTextTag // out
	var _cret C.int         // in

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	_cret = C.gtk_text_tag_get_priority(_arg0)
	runtime.KeepAlive(tag)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetPriority sets the priority of a GtkTextTag.
//
// Valid priorities start at 0 and go to one less than
// gtk.TextTagTable.GetSize(). Each tag in a table has a unique priority;
// setting the priority of one tag shifts the priorities of all the other tags
// in the table to maintain a unique priority for each tag.
//
// Higher priority tags “win” if two tags both set the same text attribute. When
// adding a tag to a tag table, it will be assigned the highest priority in the
// table by default; so normally the precedence of a set of tags is the order in
// which they were added to the table, or created with
// gtk.TextBuffer.CreateTag(), which adds the tag to the buffer’s table
// automatically.
func (tag *TextTag) SetPriority(priority int) {
	var _arg0 *C.GtkTextTag // out
	var _arg1 C.int         // out

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))
	_arg1 = C.int(priority)

	C.gtk_text_tag_set_priority(_arg0, _arg1)
	runtime.KeepAlive(tag)
	runtime.KeepAlive(priority)
}
