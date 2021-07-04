// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_color_chooser_dialog_get_type()), F: marshalColorChooserDialog},
	})
}

// ColorChooserDialog: the ColorChooserDialog widget is a dialog for choosing a
// color. It implements the ColorChooser interface.
type ColorChooserDialog interface {
	Dialog
	ColorChooser
}

// colorChooserDialog implements the ColorChooserDialog class.
type colorChooserDialog struct {
	Dialog
}

// WrapColorChooserDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorChooserDialog(obj *externglib.Object) ColorChooserDialog {
	return colorChooserDialog{
		Dialog: WrapDialog(obj),
	}
}

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorChooserDialog(obj), nil
}

func NewColorChooserDialog(title string, parent Window) ColorChooserDialog {
	var _arg1 *C.gchar     // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	_cret = C.gtk_color_chooser_dialog_new(_arg1, _arg2)

	var _colorChooserDialog ColorChooserDialog // out

	_colorChooserDialog = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ColorChooserDialog)

	return _colorChooserDialog
}

func (b colorChooserDialog) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b colorChooserDialog) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b colorChooserDialog) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b colorChooserDialog) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b colorChooserDialog) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b colorChooserDialog) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b colorChooserDialog) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b colorChooserDialog) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b colorChooserDialog) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b colorChooserDialog) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (c colorChooserDialog) AddPalette(orientation Orientation, colorsPerLine int, colors []gdk.RGBA) {
	WrapColorChooser(gextras.InternObject(c)).AddPalette(orientation, colorsPerLine, colors)
}

func (c colorChooserDialog) RGBA() gdk.RGBA {
	return WrapColorChooser(gextras.InternObject(c)).RGBA()
}

func (c colorChooserDialog) UseAlpha() bool {
	return WrapColorChooser(gextras.InternObject(c)).UseAlpha()
}

func (c colorChooserDialog) SetRGBA(color *gdk.RGBA) {
	WrapColorChooser(gextras.InternObject(c)).SetRGBA(color)
}

func (c colorChooserDialog) SetUseAlpha(useAlpha bool) {
	WrapColorChooser(gextras.InternObject(c)).SetUseAlpha(useAlpha)
}
