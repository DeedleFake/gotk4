// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_CellRendererTextClass_edited(GtkCellRendererText*, gchar*, gchar*);
// extern void _gotk4_gtk3_CellRendererText_ConnectEdited(gpointer, gchar*, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeCellRendererText = coreglib.Type(C.gtk_cell_renderer_text_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererText, F: marshalCellRendererText},
	})
}

// CellRendererTextOverrider contains methods that are overridable.
type CellRendererTextOverrider interface {
	// The function takes the following parameters:
	//
	//    - path
	//    - newText
	//
	Edited(path, newText string)
}

// CellRendererText renders a given text in its cell, using the font, color and
// style information provided by its properties. The text will be ellipsized if
// it is too long and the CellRendererText:ellipsize property allows it.
//
// If the CellRenderer:mode is GTK_CELL_RENDERER_MODE_EDITABLE, the
// CellRendererText allows to edit its text using an entry.
type CellRendererText struct {
	_ [0]func() // equal guard
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererText)(nil)
)

func classInitCellRendererTexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCellRendererTextClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ Edited(path, newText string) }); ok {
		pclass.edited = (*[0]byte)(C._gotk4_gtk3_CellRendererTextClass_edited)
	}
}

//export _gotk4_gtk3_CellRendererTextClass_edited
func _gotk4_gtk3_CellRendererTextClass_edited(arg0 *C.GtkCellRendererText, arg1 *C.gchar, arg2 *C.gchar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Edited(path, newText string) })

	var _path string    // out
	var _newText string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_newText = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	iface.Edited(_path, _newText)
}

func wrapCellRendererText(obj *coreglib.Object) *CellRendererText {
	return &CellRendererText{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererText(p uintptr) (interface{}, error) {
	return wrapCellRendererText(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_CellRendererText_ConnectEdited
func _gotk4_gtk3_CellRendererText_ConnectEdited(arg0 C.gpointer, arg1 *C.gchar, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(path, newText string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path, newText string))
	}

	var _path string    // out
	var _newText string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_newText = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_path, _newText)
}

// ConnectEdited: this signal is emitted after renderer has been edited.
//
// It is the responsibility of the application to update the model and store
// new_text at the position indicated by path.
func (renderer *CellRendererText) ConnectEdited(f func(path, newText string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(renderer, "edited", false, unsafe.Pointer(C._gotk4_gtk3_CellRendererText_ConnectEdited), f)
}

// NewCellRendererText creates a new CellRendererText. Adjust how text is drawn
// using object properties. Object properties can be set globally (with
// g_object_set()). Also, with TreeViewColumn, you can bind a property to a
// value in a TreeModel. For example, you can bind the “text” property on the
// cell renderer to a string value in the model, thus rendering a different
// string in each row of the TreeView.
//
// The function returns the following values:
//
//    - cellRendererText: new cell renderer.
//
func NewCellRendererText() *CellRendererText {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_text_new()

	var _cellRendererText *CellRendererText // out

	_cellRendererText = wrapCellRendererText(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererText
}

// SetFixedHeightFromFont sets the height of a renderer to explicitly be
// determined by the “font” and “y_pad” property set on it. Further changes in
// these properties do not affect the height, so they must be accompanied by a
// subsequent call to this function. Using this function is unflexible, and
// should really only be used if calculating the size of a cell is too slow (ie,
// a massive number of cells displayed). If number_of_rows is -1, then the fixed
// height is unset, and the height is determined by the properties again.
//
// The function takes the following parameters:
//
//    - numberOfRows: number of rows of text each cell renderer is allocated, or
//      -1.
//
func (renderer *CellRendererText) SetFixedHeightFromFont(numberOfRows int) {
	var _arg0 *C.GtkCellRendererText // out
	var _arg1 C.gint                 // out

	_arg0 = (*C.GtkCellRendererText)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	_arg1 = C.gint(numberOfRows)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(numberOfRows)
}
