// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_TextTagTable_ConnectTagRemoved(gpointer, GtkTextTag*, guintptr);
// extern void _gotk4_gtk3_TextTagTable_ConnectTagChanged(gpointer, GtkTextTag*, gboolean, guintptr);
// extern void _gotk4_gtk3_TextTagTable_ConnectTagAdded(gpointer, GtkTextTag*, guintptr);
// extern void _gotk4_gtk3_TextTagTableForEach(GtkTextTag*, gpointer);
// extern void _gotk4_gtk3_TextTagTableClass_tag_removed(GtkTextTagTable*, GtkTextTag*);
// extern void _gotk4_gtk3_TextTagTableClass_tag_changed(GtkTextTagTable*, GtkTextTag*, gboolean);
// extern void _gotk4_gtk3_TextTagTableClass_tag_added(GtkTextTagTable*, GtkTextTag*);
// void _gotk4_gtk3_TextTagTable_virtual_tag_added(void* fnptr, GtkTextTagTable* arg0, GtkTextTag* arg1) {
//   ((void (*)(GtkTextTagTable*, GtkTextTag*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk3_TextTagTable_virtual_tag_changed(void* fnptr, GtkTextTagTable* arg0, GtkTextTag* arg1, gboolean arg2) {
//   ((void (*)(GtkTextTagTable*, GtkTextTag*, gboolean))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gtk3_TextTagTable_virtual_tag_removed(void* fnptr, GtkTextTagTable* arg0, GtkTextTag* arg1) {
//   ((void (*)(GtkTextTagTable*, GtkTextTag*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeTextTagTable = coreglib.Type(C.gtk_text_tag_table_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTextTagTable, F: marshalTextTagTable},
	})
}

type TextTagTableForEach func(tag *TextTag)

// TextTagTableOverrides contains methods that are overridable.
type TextTagTableOverrides struct {
	// The function takes the following parameters:
	//
	TagAdded func(tag *TextTag)
	// The function takes the following parameters:
	//
	//    - tag
	//    - sizeChanged
	//
	TagChanged func(tag *TextTag, sizeChanged bool)
	// The function takes the following parameters:
	//
	TagRemoved func(tag *TextTag)
}

func defaultTextTagTableOverrides(v *TextTagTable) TextTagTableOverrides {
	return TextTagTableOverrides{
		TagAdded:   v.tagAdded,
		TagChanged: v.tagChanged,
		TagRemoved: v.tagRemoved,
	}
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
//    </object>.
type TextTagTable struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
}

var (
	_ coreglib.Objector = (*TextTagTable)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*TextTagTable, *TextTagTableClass, TextTagTableOverrides](
		GTypeTextTagTable,
		initTextTagTableClass,
		wrapTextTagTable,
		defaultTextTagTableOverrides,
	)
}

func initTextTagTableClass(gclass unsafe.Pointer, overrides TextTagTableOverrides, classInitFunc func(*TextTagTableClass)) {
	pclass := (*C.GtkTextTagTableClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeTextTagTable))))

	if overrides.TagAdded != nil {
		pclass.tag_added = (*[0]byte)(C._gotk4_gtk3_TextTagTableClass_tag_added)
	}

	if overrides.TagChanged != nil {
		pclass.tag_changed = (*[0]byte)(C._gotk4_gtk3_TextTagTableClass_tag_changed)
	}

	if overrides.TagRemoved != nil {
		pclass.tag_removed = (*[0]byte)(C._gotk4_gtk3_TextTagTableClass_tag_removed)
	}

	if classInitFunc != nil {
		class := (*TextTagTableClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTextTagTable(obj *coreglib.Object) *TextTagTable {
	return &TextTagTable{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalTextTagTable(p uintptr) (interface{}, error) {
	return wrapTextTagTable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (table *TextTagTable) ConnectTagAdded(f func(tag *TextTag)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(table, "tag-added", false, unsafe.Pointer(C._gotk4_gtk3_TextTagTable_ConnectTagAdded), f)
}

func (table *TextTagTable) ConnectTagChanged(f func(tag *TextTag, sizeChanged bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(table, "tag-changed", false, unsafe.Pointer(C._gotk4_gtk3_TextTagTable_ConnectTagChanged), f)
}

func (table *TextTagTable) ConnectTagRemoved(f func(tag *TextTag)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(table, "tag-removed", false, unsafe.Pointer(C._gotk4_gtk3_TextTagTable_ConnectTagRemoved), f)
}

// NewTextTagTable creates a new TextTagTable. The table contains no tags by
// default.
//
// The function returns the following values:
//
//    - textTagTable: new TextTagTable.
//
func NewTextTagTable() *TextTagTable {
	var _cret *C.GtkTextTagTable // in

	_cret = C.gtk_text_tag_table_new()

	var _textTagTable *TextTagTable // out

	_textTagTable = wrapTextTagTable(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textTagTable
}

// Add a tag to the table. The tag is assigned the highest priority in the
// table.
//
// tag must not be in a tag table already, and may not have the same name as an
// already-added tag.
//
// The function takes the following parameters:
//
//    - tag: TextTag.
//
// The function returns the following values:
//
//    - ok: TRUE on success.
//
func (table *TextTagTable) Add(tag *TextTag) bool {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(coreglib.InternObject(tag).Native()))

	_cret = C.gtk_text_tag_table_add(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ForEach calls func on each tag in table, with user data data. Note that the
// table may not be modified while iterating over it (you can’t add/remove
// tags).
//
// The function takes the following parameters:
//
//    - fn: function to call on each tag.
//
func (table *TextTagTable) ForEach(fn TextTagTableForEach) {
	var _arg0 *C.GtkTextTagTable       // out
	var _arg1 C.GtkTextTagTableForeach // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_TextTagTableForEach)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.gtk_text_tag_table_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(table)
	runtime.KeepAlive(fn)
}

// Size returns the size of the table (number of tags).
//
// The function returns the following values:
//
//    - gint: number of tags in table.
//
func (table *TextTagTable) Size() int {
	var _arg0 *C.GtkTextTagTable // out
	var _cret C.gint             // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))

	_cret = C.gtk_text_tag_table_get_size(_arg0)
	runtime.KeepAlive(table)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Lookup: look up a named tag.
//
// The function takes the following parameters:
//
//    - name of a tag.
//
// The function returns the following values:
//
//    - textTag (optional): tag, or NULL if none by that name is in the table.
//
func (table *TextTagTable) Lookup(name string) *TextTag {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.gchar           // out
	var _cret *C.GtkTextTag      // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_text_tag_table_lookup(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(name)

	var _textTag *TextTag // out

	if _cret != nil {
		_textTag = wrapTextTag(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _textTag
}

// Remove a tag from the table. If a TextBuffer has table as its tag table, the
// tag is removed from the buffer. The table’s reference to the tag is removed,
// so the tag will end up destroyed if you don’t have a reference to it.
//
// The function takes the following parameters:
//
//    - tag: TextTag.
//
func (table *TextTagTable) Remove(tag *TextTag) {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(coreglib.InternObject(tag).Native()))

	C.gtk_text_tag_table_remove(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)
}

// The function takes the following parameters:
//
func (table *TextTagTable) tagAdded(tag *TextTag) {
	gclass := (*C.GtkTextTagTableClass)(coreglib.PeekParentClass(table))
	fnarg := gclass.tag_added

	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(coreglib.InternObject(tag).Native()))

	C._gotk4_gtk3_TextTagTable_virtual_tag_added(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)
}

// The function takes the following parameters:
//
//    - tag
//    - sizeChanged
//
func (table *TextTagTable) tagChanged(tag *TextTag, sizeChanged bool) {
	gclass := (*C.GtkTextTagTableClass)(coreglib.PeekParentClass(table))
	fnarg := gclass.tag_changed

	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out
	var _arg2 C.gboolean         // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(coreglib.InternObject(tag).Native()))
	if sizeChanged {
		_arg2 = C.TRUE
	}

	C._gotk4_gtk3_TextTagTable_virtual_tag_changed(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)
	runtime.KeepAlive(sizeChanged)
}

// The function takes the following parameters:
//
func (table *TextTagTable) tagRemoved(tag *TextTag) {
	gclass := (*C.GtkTextTagTableClass)(coreglib.PeekParentClass(table))
	fnarg := gclass.tag_removed

	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(coreglib.InternObject(table).Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(coreglib.InternObject(tag).Native()))

	C._gotk4_gtk3_TextTagTable_virtual_tag_removed(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)
}

// TextTagTableClass: instance of this type is always passed by reference.
type TextTagTableClass struct {
	*textTagTableClass
}

// textTagTableClass is the struct that's finalized.
type textTagTableClass struct {
	native *C.GtkTextTagTableClass
}
