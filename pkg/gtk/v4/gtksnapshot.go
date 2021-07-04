// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/graphene"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_snapshot_get_type()), F: marshalSnapshot},
	})
}

// Snapshot: `GtkSnapshot` assists in creating `GskRenderNodes` for widgets.
//
// It functions in a similar way to a cairo context, and maintains a stack of
// render nodes and their associated transformations.
//
// The node at the top of the stack is the the one that gtk_snapshot_append_…
// functions operate on. Use the gtk_snapshot_push_… functions and
// gtk_snapshot_pop() to change the current node.
//
// The typical way to obtain a `GtkSnapshot` object is as an argument to the
// GtkWidgetClass.snapshot() vfunc. If you need to create your own
// `GtkSnapshot`, use [ctor@Gtk.Snapshot.new].
type Snapshot interface {
	gdk.Snapshot

	AppendBorderSnapshot(outline *gsk.RoundedRect, borderWidth [4]float32, borderColor [4]gdk.RGBA)

	AppendCairoSnapshot(bounds *graphene.Rect) *cairo.Context

	AppendColorSnapshot(color *gdk.RGBA, bounds *graphene.Rect)

	AppendConicGradientSnapshot(bounds *graphene.Rect, center *graphene.Point, rotation float32, stops []gsk.ColorStop)

	AppendInsetShadowSnapshot(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32)

	AppendLayoutSnapshot(layout pango.Layout, color *gdk.RGBA)

	AppendLinearGradientSnapshot(bounds *graphene.Rect, startPoint *graphene.Point, endPoint *graphene.Point, stops []gsk.ColorStop)

	AppendNodeSnapshot(node gsk.RenderNode)

	AppendOutsetShadowSnapshot(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32)

	AppendRadialGradientSnapshot(bounds *graphene.Rect, center *graphene.Point, hradius float32, vradius float32, start float32, end float32, stops []gsk.ColorStop)

	AppendRepeatingLinearGradientSnapshot(bounds *graphene.Rect, startPoint *graphene.Point, endPoint *graphene.Point, stops []gsk.ColorStop)

	AppendRepeatingRadialGradientSnapshot(bounds *graphene.Rect, center *graphene.Point, hradius float32, vradius float32, start float32, end float32, stops []gsk.ColorStop)

	AppendTextureSnapshot(texture gdk.Texture, bounds *graphene.Rect)

	GLShaderPopTextureSnapshot()

	PerspectiveSnapshot(depth float32)

	PopSnapshot()

	PushBlendSnapshot(blendMode gsk.BlendMode)

	PushBlurSnapshot(radius float64)

	PushClipSnapshot(bounds *graphene.Rect)

	PushColorMatrixSnapshot(colorMatrix *graphene.Matrix, colorOffset *graphene.Vec4)

	PushCrossFadeSnapshot(progress float64)

	PushOpacitySnapshot(opacity float64)

	PushRepeatSnapshot(bounds *graphene.Rect, childBounds *graphene.Rect)

	PushRoundedClipSnapshot(bounds *gsk.RoundedRect)

	PushShadowSnapshot(shadow *gsk.Shadow, nShadows uint)

	RenderBackgroundSnapshot(context StyleContext, x float64, y float64, width float64, height float64)

	RenderFocusSnapshot(context StyleContext, x float64, y float64, width float64, height float64)

	RenderFrameSnapshot(context StyleContext, x float64, y float64, width float64, height float64)

	RenderInsertionCursorSnapshot(context StyleContext, x float64, y float64, layout pango.Layout, index int, direction pango.Direction)

	RenderLayoutSnapshot(context StyleContext, x float64, y float64, layout pango.Layout)

	RestoreSnapshot()

	RotateSnapshot(angle float32)

	Rotate3DSnapshot(angle float32, axis *graphene.Vec3)

	SaveSnapshot()

	ScaleSnapshot(factorX float32, factorY float32)

	Scale3DSnapshot(factorX float32, factorY float32, factorZ float32)

	ToNodeSnapshot() gsk.RenderNode

	TransformSnapshot(transform *gsk.Transform)

	TransformMatrixSnapshot(matrix *graphene.Matrix)

	TranslateSnapshot(point *graphene.Point)

	Translate3DSnapshot(point *graphene.Point3D)
}

// snapshot implements the Snapshot class.
type snapshot struct {
	gdk.Snapshot
}

// WrapSnapshot wraps a GObject to the right type. It is
// primarily used internally.
func WrapSnapshot(obj *externglib.Object) Snapshot {
	return snapshot{
		Snapshot: gdk.WrapSnapshot(obj),
	}
}

func marshalSnapshot(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSnapshot(obj), nil
}

func NewSnapshot() Snapshot {
	var _cret *C.GtkSnapshot // in

	_cret = C.gtk_snapshot_new()

	var _snapshot Snapshot // out

	_snapshot = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Snapshot)

	return _snapshot
}

func (s snapshot) AppendBorderSnapshot(outline *gsk.RoundedRect, borderWidth [4]float32, borderColor [4]gdk.RGBA) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out
	var _arg2 *C.float
	var _arg3 *C.GdkRGBA

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GskRoundedRect)(unsafe.Pointer(outline.Native()))
	_arg2 = (*C.float)(unsafe.Pointer(&borderWidth))
	_arg3 = (*C.GdkRGBA)(unsafe.Pointer(&borderColor))

	C.gtk_snapshot_append_border(_arg0, _arg1, _arg2, _arg3)
}

func (s snapshot) AppendCairoSnapshot(bounds *graphene.Rect) *cairo.Context {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))

	_cret = C.gtk_snapshot_append_cairo(_arg0, _arg1)

	var _context *cairo.Context // out

	_context = (*cairo.Context)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_context, func(v **cairo.Context) {
		C.free(unsafe.Pointer(v))
	})

	return _context
}

func (s snapshot) AppendColorSnapshot(color *gdk.RGBA, bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GdkRGBA         // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))

	C.gtk_snapshot_append_color(_arg0, _arg1, _arg2)
}

func (s snapshot) AppendConicGradientSnapshot(bounds *graphene.Rect, center *graphene.Point, rotation float32, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 C.float             // out
	var _arg4 *C.GskColorStop
	var _arg5 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	_arg2 = (*C.graphene_point_t)(unsafe.Pointer(center.Native()))
	_arg3 = C.float(rotation)
	_arg5 = C.gsize(len(stops))
	_arg4 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))

	C.gtk_snapshot_append_conic_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s snapshot) AppendInsetShadowSnapshot(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out
	var _arg2 *C.GdkRGBA        // out
	var _arg3 C.float           // out
	var _arg4 C.float           // out
	var _arg5 C.float           // out
	var _arg6 C.float           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GskRoundedRect)(unsafe.Pointer(outline.Native()))
	_arg2 = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))
	_arg3 = C.float(dx)
	_arg4 = C.float(dy)
	_arg5 = C.float(spread)
	_arg6 = C.float(blurRadius)

	C.gtk_snapshot_append_inset_shadow(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (s snapshot) AppendLayoutSnapshot(layout pango.Layout, color *gdk.RGBA) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 *C.PangoLayout // out
	var _arg2 *C.GdkRGBA     // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	_arg2 = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))

	C.gtk_snapshot_append_layout(_arg0, _arg1, _arg2)
}

func (s snapshot) AppendLinearGradientSnapshot(bounds *graphene.Rect, startPoint *graphene.Point, endPoint *graphene.Point, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 *C.graphene_point_t // out
	var _arg4 *C.GskColorStop
	var _arg5 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	_arg2 = (*C.graphene_point_t)(unsafe.Pointer(startPoint.Native()))
	_arg3 = (*C.graphene_point_t)(unsafe.Pointer(endPoint.Native()))
	_arg5 = C.gsize(len(stops))
	_arg4 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))

	C.gtk_snapshot_append_linear_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s snapshot) AppendNodeSnapshot(node gsk.RenderNode) {
	var _arg0 *C.GtkSnapshot   // out
	var _arg1 *C.GskRenderNode // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer(node.Native()))

	C.gtk_snapshot_append_node(_arg0, _arg1)
}

func (s snapshot) AppendOutsetShadowSnapshot(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out
	var _arg2 *C.GdkRGBA        // out
	var _arg3 C.float           // out
	var _arg4 C.float           // out
	var _arg5 C.float           // out
	var _arg6 C.float           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GskRoundedRect)(unsafe.Pointer(outline.Native()))
	_arg2 = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))
	_arg3 = C.float(dx)
	_arg4 = C.float(dy)
	_arg5 = C.float(spread)
	_arg6 = C.float(blurRadius)

	C.gtk_snapshot_append_outset_shadow(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (s snapshot) AppendRadialGradientSnapshot(bounds *graphene.Rect, center *graphene.Point, hradius float32, vradius float32, start float32, end float32, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 C.float             // out
	var _arg4 C.float             // out
	var _arg5 C.float             // out
	var _arg6 C.float             // out
	var _arg7 *C.GskColorStop
	var _arg8 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	_arg2 = (*C.graphene_point_t)(unsafe.Pointer(center.Native()))
	_arg3 = C.float(hradius)
	_arg4 = C.float(vradius)
	_arg5 = C.float(start)
	_arg6 = C.float(end)
	_arg8 = C.gsize(len(stops))
	_arg7 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))

	C.gtk_snapshot_append_radial_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

func (s snapshot) AppendRepeatingLinearGradientSnapshot(bounds *graphene.Rect, startPoint *graphene.Point, endPoint *graphene.Point, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 *C.graphene_point_t // out
	var _arg4 *C.GskColorStop
	var _arg5 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	_arg2 = (*C.graphene_point_t)(unsafe.Pointer(startPoint.Native()))
	_arg3 = (*C.graphene_point_t)(unsafe.Pointer(endPoint.Native()))
	_arg5 = C.gsize(len(stops))
	_arg4 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))

	C.gtk_snapshot_append_repeating_linear_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s snapshot) AppendRepeatingRadialGradientSnapshot(bounds *graphene.Rect, center *graphene.Point, hradius float32, vradius float32, start float32, end float32, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 C.float             // out
	var _arg4 C.float             // out
	var _arg5 C.float             // out
	var _arg6 C.float             // out
	var _arg7 *C.GskColorStop
	var _arg8 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	_arg2 = (*C.graphene_point_t)(unsafe.Pointer(center.Native()))
	_arg3 = C.float(hradius)
	_arg4 = C.float(vradius)
	_arg5 = C.float(start)
	_arg6 = C.float(end)
	_arg8 = C.gsize(len(stops))
	_arg7 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))

	C.gtk_snapshot_append_repeating_radial_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

func (s snapshot) AppendTextureSnapshot(texture gdk.Texture, bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GdkTexture      // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))

	C.gtk_snapshot_append_texture(_arg0, _arg1, _arg2)
}

func (s snapshot) GLShaderPopTextureSnapshot() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))

	C.gtk_snapshot_gl_shader_pop_texture(_arg0)
}

func (s snapshot) PerspectiveSnapshot(depth float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(depth)

	C.gtk_snapshot_perspective(_arg0, _arg1)
}

func (s snapshot) PopSnapshot() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))

	C.gtk_snapshot_pop(_arg0)
}

func (s snapshot) PushBlendSnapshot(blendMode gsk.BlendMode) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.GskBlendMode // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.GskBlendMode(blendMode)

	C.gtk_snapshot_push_blend(_arg0, _arg1)
}

func (s snapshot) PushBlurSnapshot(radius float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(radius)

	C.gtk_snapshot_push_blur(_arg0, _arg1)
}

func (s snapshot) PushClipSnapshot(bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))

	C.gtk_snapshot_push_clip(_arg0, _arg1)
}

func (s snapshot) PushColorMatrixSnapshot(colorMatrix *graphene.Matrix, colorOffset *graphene.Vec4) {
	var _arg0 *C.GtkSnapshot       // out
	var _arg1 *C.graphene_matrix_t // out
	var _arg2 *C.graphene_vec4_t   // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(colorMatrix.Native()))
	_arg2 = (*C.graphene_vec4_t)(unsafe.Pointer(colorOffset.Native()))

	C.gtk_snapshot_push_color_matrix(_arg0, _arg1, _arg2)
}

func (s snapshot) PushCrossFadeSnapshot(progress float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(progress)

	C.gtk_snapshot_push_cross_fade(_arg0, _arg1)
}

func (s snapshot) PushOpacitySnapshot(opacity float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(opacity)

	C.gtk_snapshot_push_opacity(_arg0, _arg1)
}

func (s snapshot) PushRepeatSnapshot(bounds *graphene.Rect, childBounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(bounds.Native()))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(childBounds.Native()))

	C.gtk_snapshot_push_repeat(_arg0, _arg1, _arg2)
}

func (s snapshot) PushRoundedClipSnapshot(bounds *gsk.RoundedRect) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GskRoundedRect)(unsafe.Pointer(bounds.Native()))

	C.gtk_snapshot_push_rounded_clip(_arg0, _arg1)
}

func (s snapshot) PushShadowSnapshot(shadow *gsk.Shadow, nShadows uint) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 *C.GskShadow   // out
	var _arg2 C.gsize        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GskShadow)(unsafe.Pointer(shadow.Native()))
	_arg2 = C.gsize(nShadows)

	C.gtk_snapshot_push_shadow(_arg0, _arg1, _arg2)
}

func (s snapshot) RenderBackgroundSnapshot(context StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_background(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s snapshot) RenderFocusSnapshot(context StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_focus(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s snapshot) RenderFrameSnapshot(context StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_frame(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s snapshot) RenderInsertionCursorSnapshot(context StyleContext, x float64, y float64, layout pango.Layout, index int, direction pango.Direction) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 *C.PangoLayout     // out
	var _arg5 C.int              // out
	var _arg6 C.PangoDirection   // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	_arg5 = C.int(index)
	_arg6 = C.PangoDirection(direction)

	C.gtk_snapshot_render_insertion_cursor(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (s snapshot) RenderLayoutSnapshot(context StyleContext, x float64, y float64, layout pango.Layout) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 *C.PangoLayout     // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.gtk_snapshot_render_layout(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (s snapshot) RestoreSnapshot() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))

	C.gtk_snapshot_restore(_arg0)
}

func (s snapshot) RotateSnapshot(angle float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(angle)

	C.gtk_snapshot_rotate(_arg0, _arg1)
}

func (s snapshot) Rotate3DSnapshot(angle float32, axis *graphene.Vec3) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 C.float            // out
	var _arg2 *C.graphene_vec3_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(angle)
	_arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(axis.Native()))

	C.gtk_snapshot_rotate_3d(_arg0, _arg1, _arg2)
}

func (s snapshot) SaveSnapshot() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))

	C.gtk_snapshot_save(_arg0)
}

func (s snapshot) ScaleSnapshot(factorX float32, factorY float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out
	var _arg2 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(factorX)
	_arg2 = C.float(factorY)

	C.gtk_snapshot_scale(_arg0, _arg1, _arg2)
}

func (s snapshot) Scale3DSnapshot(factorX float32, factorY float32, factorZ float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out
	var _arg2 C.float        // out
	var _arg3 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(factorX)
	_arg2 = C.float(factorY)
	_arg3 = C.float(factorZ)

	C.gtk_snapshot_scale_3d(_arg0, _arg1, _arg2, _arg3)
}

func (s snapshot) ToNodeSnapshot() gsk.RenderNode {
	var _arg0 *C.GtkSnapshot   // out
	var _cret *C.GskRenderNode // in

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_snapshot_to_node(_arg0)

	var _renderNode gsk.RenderNode // out

	_renderNode = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gsk.RenderNode)

	return _renderNode
}

func (s snapshot) TransformSnapshot(transform *gsk.Transform) {
	var _arg0 *C.GtkSnapshot  // out
	var _arg1 *C.GskTransform // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GskTransform)(unsafe.Pointer(transform.Native()))

	C.gtk_snapshot_transform(_arg0, _arg1)
}

func (s snapshot) TransformMatrixSnapshot(matrix *graphene.Matrix) {
	var _arg0 *C.GtkSnapshot       // out
	var _arg1 *C.graphene_matrix_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(matrix.Native()))

	C.gtk_snapshot_transform_matrix(_arg0, _arg1)
}

func (s snapshot) TranslateSnapshot(point *graphene.Point) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_point_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(point.Native()))

	C.gtk_snapshot_translate(_arg0, _arg1)
}

func (s snapshot) Translate3DSnapshot(point *graphene.Point3D) {
	var _arg0 *C.GtkSnapshot        // out
	var _arg1 *C.graphene_point3d_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	C.gtk_snapshot_translate_3d(_arg0, _arg1)
}
