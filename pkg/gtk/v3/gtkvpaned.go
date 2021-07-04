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
		{T: externglib.Type(C.gtk_vpaned_get_type()), F: marshalVPaned},
	})
}

// VPaned: the VPaned widget is a container widget with two children arranged
// vertically. The division between the two panes is adjustable by the user by
// dragging a handle. See Paned for details.
//
// GtkVPaned has been deprecated, use Paned instead.
type VPaned interface {
	Paned
}

// vPaned implements the VPaned class.
type vPaned struct {
	Paned
}

// WrapVPaned wraps a GObject to the right type. It is
// primarily used internally.
func WrapVPaned(obj *externglib.Object) VPaned {
	return vPaned{
		Paned: WrapPaned(obj),
	}
}

func marshalVPaned(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVPaned(obj), nil
}

func NewVPaned() VPaned {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vpaned_new()

	var _vPaned VPaned // out

	_vPaned = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(VPaned)

	return _vPaned
}

func (b vPaned) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b vPaned) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b vPaned) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b vPaned) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b vPaned) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b vPaned) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b vPaned) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b vPaned) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b vPaned) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b vPaned) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o vPaned) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o vPaned) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
