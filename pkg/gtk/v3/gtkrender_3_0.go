// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// RenderActivity renders an activity indicator (such as in Spinner). The state
// GTK_STATE_FLAG_CHECKED determines whether there is activity going on.
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//
func RenderActivity(context *StyleContext, cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_activity(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// RenderArrow renders an arrow pointing to angle.
//
// Typical arrow rendering at 0, 1⁄2 π;, π; and 3⁄2 π:
//
// ! (arrows.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - angle: arrow angle from 0 to 2 * G_PI, being 0 the arrow pointing to the
//      north.
//    - x: x origin of the render area.
//    - y: y origin of the render area.
//    - size: square side for render area.
//
func RenderArrow(context *StyleContext, cr *cairo.Context, angle, x, y, size float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(angle)
	_arg4 = C.gdouble(x)
	_arg5 = C.gdouble(y)
	_arg6 = C.gdouble(size)

	C.gtk_render_arrow(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(angle)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(size)
}

// RenderCheck renders a checkmark (as in a CheckButton).
//
// The GTK_STATE_FLAG_CHECKED state determines whether the check is on or off,
// and GTK_STATE_FLAG_INCONSISTENT determines whether it should be marked as
// undefined.
//
// Typical checkmark rendering:
//
// ! (checks.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//
func RenderCheck(context *StyleContext, cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_check(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// RenderExpander renders an expander (as used in TreeView and Expander) in the
// area defined by x, y, width, height. The state GTK_STATE_FLAG_CHECKED
// determines whether the expander is collapsed or expanded.
//
// Typical expander rendering:
//
// ! (expanders.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//
func RenderExpander(context *StyleContext, cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_expander(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// RenderExtension renders a extension (as in a Notebook tab) in the rectangle
// defined by x, y, width, height. The side where the extension connects to is
// defined by gap_side.
//
// Typical extension rendering:
//
// ! (extensions.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//    - gapSide: side where the gap is.
//
func RenderExtension(context *StyleContext, cr *cairo.Context, x, y, width, height float64, gapSide PositionType) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out
	var _arg7 C.GtkPositionType  // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)
	_arg7 = C.GtkPositionType(gapSide)

	C.gtk_render_extension(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(gapSide)
}

// RenderFocus renders a focus indicator on the rectangle determined by x, y,
// width, height.
//
// Typical focus rendering:
//
// ! (focus.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//
func RenderFocus(context *StyleContext, cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_focus(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// RenderFrame renders a frame around the rectangle defined by x, y, width,
// height.
//
// Examples of frame rendering, showing the effect of border-image,
// border-color, border-width, border-radius and junctions:
//
// ! (frames.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//
func RenderFrame(context *StyleContext, cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_frame(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// RenderFrameGap renders a frame around the rectangle defined by (x, y, width,
// height), leaving a gap on one side. xy0_gap and xy1_gap will mean X
// coordinates for GTK_POS_TOP and GTK_POS_BOTTOM gap sides, and Y coordinates
// for GTK_POS_LEFT and GTK_POS_RIGHT.
//
// Typical rendering of a frame with a gap:
//
// ! (frame-gap.png)
//
// Deprecated: Use gtk_render_frame() instead. Themes can create gaps by
// omitting borders via CSS.
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//    - gapSide: side where the gap is.
//    - xy0Gap: initial coordinate (X or Y depending on gap_side) for the gap.
//    - xy1Gap: end coordinate (X or Y depending on gap_side) for the gap.
//
func RenderFrameGap(context *StyleContext, cr *cairo.Context, x, y, width, height float64, gapSide PositionType, xy0Gap, xy1Gap float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out
	var _arg7 C.GtkPositionType  // out
	var _arg8 C.gdouble          // out
	var _arg9 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)
	_arg7 = C.GtkPositionType(gapSide)
	_arg8 = C.gdouble(xy0Gap)
	_arg9 = C.gdouble(xy1Gap)

	C.gtk_render_frame_gap(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(gapSide)
	runtime.KeepAlive(xy0Gap)
	runtime.KeepAlive(xy1Gap)
}

// RenderHandle renders a handle (as in HandleBox, Paned and Window’s resize
// grip), in the rectangle determined by x, y, width, height.
//
// Handles rendered for the paned and grip classes:
//
// ! (handles.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//
func RenderHandle(context *StyleContext, cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_handle(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// RenderIconPixbuf renders the icon specified by source at the given size,
// returning the result in a pixbuf.
//
// Deprecated: Use gtk_icon_theme_load_icon() instead.
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - source specifying the icon to render.
//    - size (IconSize) to render the icon at. A size of (GtkIconSize) -1 means
//      render at the size of the source and don’t scale.
//
// The function returns the following values:
//
//    - pixbuf: newly-created Pixbuf containing the rendered icon.
//
func RenderIconPixbuf(context *StyleContext, source *IconSource, size int) *gdkpixbuf.Pixbuf {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.GtkIconSource   // out
	var _arg3 C.GtkIconSize      // out
	var _cret *C.GdkPixbuf       // in

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.GtkIconSource)(gextras.StructNative(unsafe.Pointer(source)))
	_arg3 = C.GtkIconSize(size)

	_cret = C.gtk_render_icon_pixbuf(_arg1, _arg2, _arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(source)
	runtime.KeepAlive(size)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// RenderLayout renders layout on the coordinates x, y.
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin.
//    - y: y origin.
//    - layout to render.
//
func RenderLayout(context *StyleContext, cr *cairo.Context, x, y float64, layout *pango.Layout) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 *C.PangoLayout     // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = (*C.PangoLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	C.gtk_render_layout(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(layout)
}

// RenderLine renders a line from (x0, y0) to (x1, y1).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x0: x coordinate for the origin of the line.
//    - y0: y coordinate for the origin of the line.
//    - x1: x coordinate for the end of the line.
//    - y1: y coordinate for the end of the line.
//
func RenderLine(context *StyleContext, cr *cairo.Context, x0, y0, x1, y1 float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x0)
	_arg4 = C.gdouble(y0)
	_arg5 = C.gdouble(x1)
	_arg6 = C.gdouble(y1)

	C.gtk_render_line(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x0)
	runtime.KeepAlive(y0)
	runtime.KeepAlive(x1)
	runtime.KeepAlive(y1)
}

// RenderOption renders an option mark (as in a RadioButton), the
// GTK_STATE_FLAG_CHECKED state will determine whether the option is on or off,
// and GTK_STATE_FLAG_INCONSISTENT whether it should be marked as undefined.
//
// Typical option mark rendering:
//
// ! (options.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//
func RenderOption(context *StyleContext, cr *cairo.Context, x, y, width, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_option(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// RenderSlider renders a slider (as in Scale) in the rectangle defined by x, y,
// width, height. orientation defines whether the slider is vertical or
// horizontal.
//
// Typical slider rendering:
//
// ! (sliders.png).
//
// The function takes the following parameters:
//
//    - context: StyleContext.
//    - cr: #cairo_t.
//    - x: x origin of the rectangle.
//    - y: y origin of the rectangle.
//    - width: rectangle width.
//    - height: rectangle height.
//    - orientation of the slider.
//
func RenderSlider(context *StyleContext, cr *cairo.Context, x, y, width, height float64, orientation Orientation) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out
	var _arg7 C.GtkOrientation   // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)
	_arg7 = C.GtkOrientation(orientation)

	C.gtk_render_slider(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(orientation)
}
