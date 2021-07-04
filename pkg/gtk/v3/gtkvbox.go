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
		{T: externglib.Type(C.gtk_vbox_get_type()), F: marshalVBox},
	})
}

// VBox: a VBox is a container that organizes child widgets into a single
// column.
//
// Use the Box packing interface to determine the arrangement, spacing, height,
// and alignment of VBox children.
//
// All children are allocated the same width.
//
// GtkVBox has been deprecated. You can use Box with a Orientable:orientation
// set to GTK_ORIENTATION_VERTICAL instead when calling gtk_box_new(), which is
// a very quick and easy change.
//
// If you have derived your own classes from GtkVBox, you can change the
// inheritance to derive directly from Box, and set the Orientable:orientation
// property to GTK_ORIENTATION_VERTICAL in your instance init function, with a
// call like:
//
//    gtk_orientable_set_orientation (GTK_ORIENTABLE (object),
//                                    GTK_ORIENTATION_VERTICAL);
//
// If you have a grid-like layout composed of nested boxes, and you don’t need
// first-child or last-child styling, the recommendation is to switch to Grid.
// For more information about migrating to Grid, see [Migrating from other
// containers to GtkGrid][gtk-migrating-GtkGrid].
type VBox interface {
	Box
}

// vBox implements the VBox class.
type vBox struct {
	Box
}

// WrapVBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapVBox(obj *externglib.Object) VBox {
	return vBox{
		Box: WrapBox(obj),
	}
}

func marshalVBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVBox(obj), nil
}

func NewVBox(homogeneous bool, spacing int) VBox {
	var _arg1 C.gboolean   // out
	var _arg2 C.gint       // out
	var _cret *C.GtkWidget // in

	if homogeneous {
		_arg1 = C.TRUE
	}
	_arg2 = C.gint(spacing)

	_cret = C.gtk_vbox_new(_arg1, _arg2)

	var _vBox VBox // out

	_vBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(VBox)

	return _vBox
}

func (b vBox) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b vBox) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b vBox) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b vBox) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b vBox) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b vBox) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b vBox) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b vBox) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b vBox) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b vBox) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o vBox) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o vBox) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
