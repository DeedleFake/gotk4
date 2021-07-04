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
		{T: externglib.Type(C.gtk_overlay_get_type()), F: marshalOverlay},
	})
}

// Overlay: gtkOverlay is a container which contains a single main child, on top
// of which it can place “overlay” widgets. The position of each overlay widget
// is determined by its Widget:halign and Widget:valign properties. E.g. a
// widget with both alignments set to GTK_ALIGN_START will be placed at the top
// left corner of the GtkOverlay container, whereas an overlay with halign set
// to GTK_ALIGN_CENTER and valign set to GTK_ALIGN_END will be placed a the
// bottom edge of the GtkOverlay, horizontally centered. The position can be
// adjusted by setting the margin properties of the child to non-zero values.
//
// More complicated placement of overlays is possible by connecting to the
// Overlay::get-child-position signal.
//
// An overlay’s minimum and natural sizes are those of its main child. The sizes
// of overlay children are not considered when measuring these preferred sizes.
//
//
// GtkOverlay as GtkBuildable
//
// The GtkOverlay implementation of the GtkBuildable interface supports placing
// a child as an overlay by specifying “overlay” as the “type” attribute of a
// `<child>` element.
//
//
// CSS nodes
//
// GtkOverlay has a single CSS node with the name “overlay”. Overlay children
// whose alignments cause them to be positioned at an edge get the style classes
// “.left”, “.right”, “.top”, and/or “.bottom” according to their position.
type Overlay interface {
	Bin

	AddOverlayOverlay(widget Widget)

	OverlayPassThrough(widget Widget) bool

	ReorderOverlayOverlay(child Widget, index_ int)

	SetOverlayPassThroughOverlay(widget Widget, passThrough bool)
}

// overlay implements the Overlay class.
type overlay struct {
	Bin
}

// WrapOverlay wraps a GObject to the right type. It is
// primarily used internally.
func WrapOverlay(obj *externglib.Object) Overlay {
	return overlay{
		Bin: WrapBin(obj),
	}
}

func marshalOverlay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapOverlay(obj), nil
}

func NewOverlay() Overlay {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_overlay_new()

	var _overlay Overlay // out

	_overlay = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Overlay)

	return _overlay
}

func (o overlay) AddOverlayOverlay(widget Widget) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_overlay_add_overlay(_arg0, _arg1)
}

func (o overlay) OverlayPassThrough(widget Widget) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_overlay_get_overlay_pass_through(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o overlay) ReorderOverlayOverlay(child Widget, index_ int) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(index_)

	C.gtk_overlay_reorder_overlay(_arg0, _arg1, _arg2)
}

func (o overlay) SetOverlayPassThroughOverlay(widget Widget, passThrough bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	if passThrough {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_overlay_pass_through(_arg0, _arg1, _arg2)
}

func (b overlay) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b overlay) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b overlay) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b overlay) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b overlay) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b overlay) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b overlay) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b overlay) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b overlay) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b overlay) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
