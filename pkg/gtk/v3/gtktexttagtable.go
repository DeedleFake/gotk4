// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
//
// void gotk4_TextTagTableForeach(GtkTextTag*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_tag_table_get_type()), F: marshalTextTagTable},
	})
}

type TextTagTableForeach func(tag *TextTagClass, data interface{})

//export gotk4_TextTagTableForeach
func gotk4_TextTagTableForeach(arg0 *C.GtkTextTag, arg1 C.gpointer) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var tag *TextTagClass // out
	var data interface{}  // out

	tag = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*TextTagClass)
	data = box.Get(uintptr(arg1))

	fn := v.(TextTagTableForeach)
	fn(tag, data)
}

// TextTagTableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TextTagTableOverrider interface {
	TagAdded(tag TextTag)
	TagChanged(tag TextTag, sizeChanged bool)
	TagRemoved(tag TextTag)
}

// TextTagTable: you may wish to begin by reading the [text widget conceptual
// overview][TextWidget] which gives an overview of all the objects and data
// types related to the text widget and how they work together.
//
//
// GtkTextTagTables as GtkBuildable
//
// The GtkTextTagTable implementation of the GtkBuildable interface supports
// adding tags by specifying “tag” as the “type” attribute of a <child> element.
//
// An example of a UI definition fragment specifying tags:
//
//    <object class="GtkTextTagTable">
//     <child type="tag">
//       <object class="GtkTextTag"/>
//     </child>
//    </object>
type TextTagTable interface {
	gextras.Objector

	// Add a tag to the table. The tag is assigned the highest priority in the
	// table.
	//
	// @tag must not be in a tag table already, and may not have the same name
	// as an already-added tag.
	Add(tag TextTag) bool
	// Foreach calls @func on each tag in @table, with user data @data. Note
	// that the table may not be modified while iterating over it (you can’t
	// add/remove tags).
	Foreach(fn TextTagTableForeach)
	// Size returns the size of the table (number of tags)
	Size() int
	// Lookup: look up a named tag.
	Lookup(name string) *TextTagClass
	// Remove a tag from the table. If a TextBuffer has @table as its tag table,
	// the tag is removed from the buffer. The table’s reference to the tag is
	// removed, so the tag will end up destroyed if you don’t have a reference
	// to it.
	Remove(tag TextTag)
}

// TextTagTableClass implements the TextTagTable interface.
type TextTagTableClass struct {
	*externglib.Object
	BuildableIface
}

var _ TextTagTable = (*TextTagTableClass)(nil)

func wrapTextTagTable(obj *externglib.Object) TextTagTable {
	return &TextTagTableClass{
		Object: obj,
		BuildableIface: BuildableIface{
			Object: obj,
		},
	}
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTextTagTable(obj), nil
}

// NewTextTagTable creates a new TextTagTable. The table contains no tags by
// default.
func NewTextTagTable() *TextTagTableClass {
	var _cret *C.GtkTextTagTable // in

	_cret = C.gtk_text_tag_table_new()

	var _textTagTable *TextTagTableClass // out

	_textTagTable = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TextTagTableClass)

	return _textTagTable
}

// Add a tag to the table. The tag is assigned the highest priority in the
// table.
//
// @tag must not be in a tag table already, and may not have the same name as an
// already-added tag.
func (table *TextTagTableClass) Add(tag TextTag) bool {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	_cret = C.gtk_text_tag_table_add(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Foreach calls @func on each tag in @table, with user data @data. Note that
// the table may not be modified while iterating over it (you can’t add/remove
// tags).
func (table *TextTagTableClass) Foreach(fn TextTagTableForeach) {
	var _arg0 *C.GtkTextTagTable       // out
	var _arg1 C.GtkTextTagTableForeach // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*[0]byte)(C.gotk4_TextTagTableForeach)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_text_tag_table_foreach(_arg0, _arg1, _arg2)
}

// Size returns the size of the table (number of tags)
func (table *TextTagTableClass) Size() int {
	var _arg0 *C.GtkTextTagTable // out
	var _cret C.gint             // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))

	_cret = C.gtk_text_tag_table_get_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Lookup: look up a named tag.
func (table *TextTagTableClass) Lookup(name string) *TextTagClass {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.gchar           // out
	var _cret *C.GtkTextTag      // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_text_tag_table_lookup(_arg0, _arg1)

	var _textTag *TextTagClass // out

	_textTag = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*TextTagClass)

	return _textTag
}

// Remove a tag from the table. If a TextBuffer has @table as its tag table, the
// tag is removed from the buffer. The table’s reference to the tag is removed,
// so the tag will end up destroyed if you don’t have a reference to it.
func (table *TextTagTableClass) Remove(tag TextTag) {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	C.gtk_text_tag_table_remove(_arg0, _arg1)
}
