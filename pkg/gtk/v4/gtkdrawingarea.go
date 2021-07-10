// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drawing_area_get_type()), F: marshalDrawingAreaer},
	})
}

// DrawingAreaDrawFunc: whenever @drawing_area needs to redraw, this function
// will be called.
//
// This function should exclusively redraw the contents of the drawing area and
// must not call any widget functions that cause changes.
type DrawingAreaDrawFunc func(drawingArea *DrawingArea, cr *cairo.Context, width int, height int, userData interface{})

//export gotk4_DrawingAreaDrawFunc
func gotk4_DrawingAreaDrawFunc(arg0 *C.GtkDrawingArea, arg1 *C.cairo_t, arg2 C.int, arg3 C.int, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var drawingArea *DrawingArea // out
	var cr *cairo.Context        // out
	var width int                // out
	var height int               // out
	var userData interface{}     // out

	drawingArea = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*DrawingArea)
	cr = (*cairo.Context)(unsafe.Pointer(arg1))
	width = int(arg2)
	height = int(arg3)
	userData = box.Get(uintptr(arg4))

	fn := v.(DrawingAreaDrawFunc)
	fn(drawingArea, cr, width, height, userData)
}

// DrawingAreaerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DrawingAreaerOverrider interface {
	Resize(width int, height int)
}

// DrawingAreaer describes DrawingArea's methods.
type DrawingAreaer interface {
	gextras.Objector

	ContentHeight() int
	ContentWidth() int
	SetContentHeight(height int)
	SetContentWidth(width int)
}

// DrawingArea: `GtkDrawingArea` is a widget that allows drawing with cairo.
//
// !An example GtkDrawingArea (drawingarea.png)
//
// It’s essentially a blank widget; you can draw on it. After creating a drawing
// area, the application may want to connect to:
//
// - The [signal@Gtk.Widget::realize] signal to take any necessary actions when
// the widget is instantiated on a particular display. (Create GDK resources in
// response to this signal.)
//
// - The [signal@Gtk.DrawingArea::resize] signal to take any necessary actions
// when the widget changes size.
//
// - Call [method@Gtk.DrawingArea.set_draw_func] to handle redrawing the
// contents of the widget.
//
// The following code portion demonstrates using a drawing area to display a
// circle in the normal widget foreground color.
//
//
// Simple GtkDrawingArea usage
//
// “`c static void draw_function (GtkDrawingArea *area, cairo_t *cr, int width,
// int height, gpointer data) { GdkRGBA color; GtkStyleContext *context;
//
//    context = gtk_widget_get_style_context (GTK_WIDGET (area));
//
//    cairo_arc (cr,
//               width / 2.0, height / 2.0,
//               MIN (width, height) / 2.0,
//               0, 2 * G_PI);
//
//    gtk_style_context_get_color (context,
//                                 &color);
//    gdk_cairo_set_source_rgba (cr, &color);
//
//    cairo_fill (cr);
//
// }
//
// int main (int argc, char **argv) { gtk_init ();
//
//    GtkWidget *area = gtk_drawing_area_new ();
//    gtk_drawing_area_set_content_width (GTK_DRAWING_AREA (area), 100);
//    gtk_drawing_area_set_content_height (GTK_DRAWING_AREA (area), 100);
//    gtk_drawing_area_set_draw_func (GTK_DRAWING_AREA (area),
//                                    draw_function,
//                                    NULL, NULL);
//    return 0;
//
// } “`
//
// The draw function is normally called when a drawing area first comes
// onscreen, or when it’s covered by another window and then uncovered. You can
// also force a redraw by adding to the “damage region” of the drawing area’s
// window using [method@Gtk.Widget.queue_draw]. This will cause the drawing area
// to call the draw function again.
//
// The available routines for drawing are documented on the [GDK Drawing
// Primitives][gdk4-Cairo-Interaction] page and the cairo documentation.
//
// To receive mouse events on a drawing area, you will need to use event
// controllers. To receive keyboard events, you will need to set the “can-focus”
// property on the drawing area, and you should probably draw some user-visible
// indication that the drawing area is focused.
//
// If you need more complex control over your widget, you should consider
// creating your own `GtkWidget` subclass.
type DrawingArea struct {
	*externglib.Object

	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ DrawingAreaer = (*DrawingArea)(nil)

func wrapDrawingAreaer(obj *externglib.Object) DrawingAreaer {
	return &DrawingArea{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
	}
}

func marshalDrawingAreaer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrawingAreaer(obj), nil
}

// NewDrawingArea creates a new drawing area.
func NewDrawingArea() *DrawingArea {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_drawing_area_new()

	var _drawingArea *DrawingArea // out

	_drawingArea = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*DrawingArea)

	return _drawingArea
}

// ContentHeight retrieves the content height of the `GtkDrawingArea`.
func (self *DrawingArea) ContentHeight() int {
	var _arg0 *C.GtkDrawingArea // out
	var _cret C.int             // in

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drawing_area_get_content_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ContentWidth retrieves the content width of the `GtkDrawingArea`.
func (self *DrawingArea) ContentWidth() int {
	var _arg0 *C.GtkDrawingArea // out
	var _cret C.int             // in

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drawing_area_get_content_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetContentHeight sets the desired height of the contents of the drawing area.
//
// Note that because widgets may be allocated larger sizes than they requested,
// it is possible that the actual height passed to your draw function is larger
// than the height set here. You can use [method@Gtk.Widget.set_valign] to avoid
// that.
//
// If the height is set to 0 (the default), the drawing area may disappear.
func (self *DrawingArea) SetContentHeight(height int) {
	var _arg0 *C.GtkDrawingArea // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(height)

	C.gtk_drawing_area_set_content_height(_arg0, _arg1)
}

// SetContentWidth sets the desired width of the contents of the drawing area.
//
// Note that because widgets may be allocated larger sizes than they requested,
// it is possible that the actual width passed to your draw function is larger
// than the width set here. You can use [method@Gtk.Widget.set_halign] to avoid
// that.
//
// If the width is set to 0 (the default), the drawing area may disappear.
func (self *DrawingArea) SetContentWidth(width int) {
	var _arg0 *C.GtkDrawingArea // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(width)

	C.gtk_drawing_area_set_content_width(_arg0, _arg1)
}
