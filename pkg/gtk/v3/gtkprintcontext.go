// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_context_get_type()), F: marshalPrintContexter},
	})
}

// PrintContext encapsulates context information that is required when drawing
// pages for printing, such as the cairo context and important parameters like
// page size and resolution. It also lets you easily create Layout and Context
// objects that match the font metrics of the cairo surface.
//
// GtkPrintContext objects gets passed to the PrintOperation::begin-print,
// PrintOperation::end-print, PrintOperation::request-page-setup and
// PrintOperation::draw-page signals on the PrintOperation.
//
// Using GtkPrintContext in a PrintOperation::draw-page callback
//
//    static void
//    draw_page (GtkPrintOperation *operation,
//    	   GtkPrintContext   *context,
//    	   int                page_nr)
//    {
//      cairo_t *cr;
//      PangoLayout *layout;
//      PangoFontDescription *desc;
//
//      cr = gtk_print_context_get_cairo_context (context);
//
//      // Draw a red rectangle, as wide as the paper (inside the margins)
//      cairo_set_source_rgb (cr, 1.0, 0, 0);
//      cairo_rectangle (cr, 0, 0, gtk_print_context_get_width (context), 50);
//
//      cairo_fill (cr);
//
//      // Draw some lines
//      cairo_move_to (cr, 20, 10);
//      cairo_line_to (cr, 40, 20);
//      cairo_arc (cr, 60, 60, 20, 0, M_PI);
//      cairo_line_to (cr, 80, 20);
//
//      cairo_set_source_rgb (cr, 0, 0, 0);
//      cairo_set_line_width (cr, 5);
//      cairo_set_line_cap (cr, CAIRO_LINE_CAP_ROUND);
//      cairo_set_line_join (cr, CAIRO_LINE_JOIN_ROUND);
//
//      cairo_stroke (cr);
//
//      // Draw some text
//      layout = gtk_print_context_create_pango_layout (context);
//      pango_layout_set_text (layout, "Hello World! Printing is easy", -1);
//      desc = pango_font_description_from_string ("sans 28");
//      pango_layout_set_font_description (layout, desc);
//      pango_font_description_free (desc);
//
//      cairo_move_to (cr, 30, 20);
//      pango_cairo_layout_path (cr, layout);
//
//      // Font Outline
//      cairo_set_source_rgb (cr, 0.93, 1.0, 0.47);
//      cairo_set_line_width (cr, 0.5);
//      cairo_stroke_preserve (cr);
//
//      // Font Fill
//      cairo_set_source_rgb (cr, 0, 0.0, 1.0);
//      cairo_fill (cr);
//
//      g_object_unref (layout);
//    }
//
// Printing support was added in GTK+ 2.10.
type PrintContext struct {
	*externglib.Object
}

var _ gextras.Nativer = (*PrintContext)(nil)

func wrapPrintContext(obj *externglib.Object) *PrintContext {
	return &PrintContext{
		Object: obj,
	}
}

func marshalPrintContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPrintContext(obj), nil
}

// CreatePangoContext creates a new Context that can be used with the
// PrintContext.
func (context *PrintContext) CreatePangoContext() *pango.Context {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoContext    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_create_pango_context(_arg0)

	var _ret *pango.Context // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_ret = &pango.Context{
			Object: obj,
		}
	}

	return _ret
}

// CreatePangoLayout creates a new Layout that is suitable for use with the
// PrintContext.
func (context *PrintContext) CreatePangoLayout() *pango.Layout {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoLayout     // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_create_pango_layout(_arg0)

	var _layout *pango.Layout // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_layout = &pango.Layout{
			Object: obj,
		}
	}

	return _layout
}

// CairoContext obtains the cairo context that is associated with the
// PrintContext.
func (context *PrintContext) CairoContext() *cairo.Context {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_cairo_context(_arg0)

	var _ret *cairo.Context // out

	_ret = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_ret, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})

	return _ret
}

// DPIX obtains the horizontal resolution of the PrintContext, in dots per inch.
func (context *PrintContext) DPIX() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.gdouble          // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_dpi_x(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DPIY obtains the vertical resolution of the PrintContext, in dots per inch.
func (context *PrintContext) DPIY() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.gdouble          // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_dpi_y(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// HardMargins obtains the hardware printer margins of the PrintContext, in
// units.
func (context *PrintContext) HardMargins() (top float64, bottom float64, left float64, right float64, ok bool) {
	var _arg0 *C.GtkPrintContext // out
	var _arg1 C.gdouble          // in
	var _arg2 C.gdouble          // in
	var _arg3 C.gdouble          // in
	var _arg4 C.gdouble          // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_hard_margins(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)

	var _top float64    // out
	var _bottom float64 // out
	var _left float64   // out
	var _right float64  // out
	var _ok bool        // out

	_top = float64(_arg1)
	_bottom = float64(_arg2)
	_left = float64(_arg3)
	_right = float64(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _top, _bottom, _left, _right, _ok
}

// Height obtains the height of the PrintContext, in pixels.
func (context *PrintContext) Height() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.gdouble          // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_height(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageSetup obtains the PageSetup that determines the page dimensions of the
// PrintContext.
func (context *PrintContext) PageSetup() *PageSetup {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.GtkPageSetup    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_page_setup(_arg0)

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(externglib.Take(unsafe.Pointer(_cret)))

	return _pageSetup
}

// PangoFontmap returns a FontMap that is suitable for use with the
// PrintContext.
func (context *PrintContext) PangoFontmap() *pango.FontMap {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoFontMap    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_pango_fontmap(_arg0)

	var _fontMap *pango.FontMap // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_fontMap = &pango.FontMap{
			Object: obj,
		}
	}

	return _fontMap
}

// Width obtains the width of the PrintContext, in pixels.
func (context *PrintContext) Width() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.gdouble          // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_width(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetCairoContext sets a new cairo context on a print context.
//
// This function is intended to be used when implementing an internal print
// preview, it is not needed for printing, since GTK+ itself creates a suitable
// cairo context in that case.
func (context *PrintContext) SetCairoContext(cr *cairo.Context, dpiX float64, dpiY float64) {
	var _arg0 *C.GtkPrintContext // out
	var _arg1 *C.cairo_t         // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.double(dpiX)
	_arg3 = C.double(dpiY)

	C.gtk_print_context_set_cairo_context(_arg0, _arg1, _arg2, _arg3)
}
