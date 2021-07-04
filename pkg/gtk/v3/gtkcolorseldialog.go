// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_selection_dialog_get_type()), F: marshalColorSelectionDialog},
	})
}

type ColorSelectionDialog interface {
	Dialog

	ColorSelection() Widget
}

// colorSelectionDialog implements the ColorSelectionDialog class.
type colorSelectionDialog struct {
	Dialog
}

// WrapColorSelectionDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorSelectionDialog(obj *externglib.Object) ColorSelectionDialog {
	return colorSelectionDialog{
		Dialog: WrapDialog(obj),
	}
}

func marshalColorSelectionDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorSelectionDialog(obj), nil
}

func NewColorSelectionDialog(title string) ColorSelectionDialog {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_color_selection_dialog_new(_arg1)

	var _colorSelectionDialog ColorSelectionDialog // out

	_colorSelectionDialog = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ColorSelectionDialog)

	return _colorSelectionDialog
}

func (c colorSelectionDialog) ColorSelection() Widget {
	var _arg0 *C.GtkColorSelectionDialog // out
	var _cret *C.GtkWidget               // in

	_arg0 = (*C.GtkColorSelectionDialog)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_color_selection_dialog_get_color_selection(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (b colorSelectionDialog) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b colorSelectionDialog) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b colorSelectionDialog) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b colorSelectionDialog) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b colorSelectionDialog) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b colorSelectionDialog) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b colorSelectionDialog) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b colorSelectionDialog) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b colorSelectionDialog) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b colorSelectionDialog) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
