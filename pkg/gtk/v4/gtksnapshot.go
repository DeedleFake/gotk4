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
	gextras.Objector

	// AppendCairo creates a new `GskCairoNode` and appends it to the current
	// render node of @snapshot, without changing the current node.
	AppendCairo(bounds *graphene.Rect) *cairo.Context
	// AppendColor creates a new render node drawing the @color into the given
	// @bounds and appends it to the current render node of @snapshot.
	//
	// You should try to avoid calling this function if @color is transparent.
	AppendColor(color *gdk.RGBA, bounds *graphene.Rect)
	// AppendInsetShadow appends an inset shadow into the box given by @outline.
	AppendInsetShadow(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32)
	AppendLayout(layout pango.Layout, color *gdk.RGBA)
	// AppendNode appends @node to the current render node of @snapshot, without
	// changing the current node.
	//
	// If @snapshot does not have a current node yet, @node will become the
	// initial node.
	AppendNode(node gsk.RenderNode)
	// AppendOutsetShadow appends an outset shadow node around the box given by
	// @outline.
	AppendOutsetShadow(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32)
	// AppendTexture creates a new render node drawing the @texture into the
	// given @bounds and appends it to the current render node of @snapshot.
	AppendTexture(texture gdk.Texture, bounds *graphene.Rect)
	// GLShaderPopTexture removes the top element from the stack of render nodes
	// and adds it to the nearest `GskGLShaderNode` below it.
	//
	// This must be called the same number of times as the number of textures is
	// needed for the shader in [method@Gtk.Snapshot.push_gl_shader].
	GLShaderPopTexture()
	// Perspective applies a perspective projection transform.
	//
	// See [method@Gsk.Transform.perspective] for a discussion on the details.
	Perspective(depth float32)
	// Pop removes the top element from the stack of render nodes, and appends
	// it to the node underneath it.
	Pop()
	// PushBlur blurs an image.
	//
	// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
	PushBlur(radius float64)
	// PushClip clips an image to a rectangle.
	//
	// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
	PushClip(bounds *graphene.Rect)
	// PushColorMatrix modifies the colors of an image by applying an affine
	// transformation in RGB space.
	//
	// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
	PushColorMatrix(colorMatrix *graphene.Matrix, colorOffset *graphene.Vec4)
	// PushCrossFade snapshots a cross-fade operation between two images with
	// the given @progress.
	//
	// Until the first call to [method@Gtk.Snapshot.pop], the start image will
	// be snapshot. After that call, the end image will be recorded until the
	// second call to [method@Gtk.Snapshot.pop].
	//
	// Calling this function requires two subsequent calls to
	// [method@Gtk.Snapshot.pop].
	PushCrossFade(progress float64)
	// PushOpacity modifies the opacity of an image.
	//
	// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
	PushOpacity(opacity float64)
	// PushRepeat creates a node that repeats the child node.
	//
	// The child is recorded until the next call to [method@Gtk.Snapshot.pop].
	PushRepeat(bounds *graphene.Rect, childBounds *graphene.Rect)
	// PushRoundedClip clips an image to a rounded rectangle.
	//
	// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
	PushRoundedClip(bounds *gsk.RoundedRect)
	// PushShadow applies a shadow to an image.
	//
	// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
	PushShadow(shadow *gsk.Shadow, nShadows uint)
	// RenderBackground creates a render node for the CSS background according
	// to @context, and appends it to the current node of @snapshot, without
	// changing the current node.
	RenderBackground(context StyleContext, x float64, y float64, width float64, height float64)
	// RenderFocus creates a render node for the focus outline according to
	// @context, and appends it to the current node of @snapshot, without
	// changing the current node.
	RenderFocus(context StyleContext, x float64, y float64, width float64, height float64)
	// RenderFrame creates a render node for the CSS border according to
	// @context, and appends it to the current node of @snapshot, without
	// changing the current node.
	RenderFrame(context StyleContext, x float64, y float64, width float64, height float64)
	// RenderLayout creates a render node for rendering @layout according to the
	// style information in @context, and appends it to the current node of
	// @snapshot, without changing the current node.
	RenderLayout(context StyleContext, x float64, y float64, layout pango.Layout)
	// Restore restores @snapshot to the state saved by a preceding call to
	// gtk_snapshot_save() and removes that state from the stack of saved
	// states.
	Restore()
	// Rotate rotates @@snapshot's coordinate system by @angle degrees in 2D
	// space - or in 3D speak, rotates around the Z axis.
	//
	// To rotate around other axes, use [method@Gsk.Transform.rotate_3d].
	Rotate(angle float32)
	// Rotate3D rotates @snapshot's coordinate system by @angle degrees around
	// @axis.
	//
	// For a rotation in 2D space, use [method@Gsk.Transform.rotate].
	Rotate3D(angle float32, axis *graphene.Vec3)
	// Save makes a copy of the current state of @snapshot and saves it on an
	// internal stack.
	//
	// When [method@Gtk.Snapshot.restore] is called, @snapshot will be restored
	// to the saved state. Multiple calls to gtk_snapshot_save() and
	// gtk_snapshot_restore() can be nested; each call to gtk_snapshot_restore()
	// restores the state from the matching paired gtk_snapshot_save().
	//
	// It is necessary to clear all saved states with corresponding calls to
	// gtk_snapshot_restore().
	Save()
	// Scale scales @snapshot's coordinate system in 2-dimensional space by the
	// given factors.
	//
	// Use [method@Gtk.Snapshot.scale_3d] to scale in all 3 dimensions.
	Scale(factorX float32, factorY float32)
	// Scale3D scales @snapshot's coordinate system by the given factors.
	Scale3D(factorX float32, factorY float32, factorZ float32)
	// ToNode returns the render node that was constructed by @snapshot.
	//
	// After calling this function, it is no longer possible to add more nodes
	// to @snapshot. The only function that should be called after this is
	// g_object_unref().
	ToNode() *gsk.RenderNodeClass
	// Transform transforms @snapshot's coordinate system with the given
	// @transform.
	Transform(transform *gsk.Transform)
	// TransformMatrix transforms @snapshot's coordinate system with the given
	// @matrix.
	TransformMatrix(matrix *graphene.Matrix)
	// Translate translates @snapshot's coordinate system by @point in
	// 2-dimensional space.
	Translate(point *graphene.Point)
	// Translate3D translates @snapshot's coordinate system by @point.
	Translate3D(point *graphene.Point3D)
}

// SnapshotClass implements the Snapshot interface.
type SnapshotClass struct {
	gdk.SnapshotClass
}

var _ Snapshot = (*SnapshotClass)(nil)

func wrapSnapshot(obj *externglib.Object) Snapshot {
	return &SnapshotClass{
		SnapshotClass: gdk.SnapshotClass{
			Object: obj,
		},
	}
}

func marshalSnapshot(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSnapshot(obj), nil
}

// NewSnapshot creates a new `GtkSnapshot`.
func NewSnapshot() *SnapshotClass {
	var _cret *C.GtkSnapshot // in

	_cret = C.gtk_snapshot_new()

	var _snapshot *SnapshotClass // out

	_snapshot = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*SnapshotClass)

	return _snapshot
}

// AppendCairo creates a new `GskCairoNode` and appends it to the current render
// node of @snapshot, without changing the current node.
func (s *SnapshotClass) AppendCairo(bounds *graphene.Rect) *cairo.Context {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(*graphene.Rect))

	_cret = C.gtk_snapshot_append_cairo(_arg0, _arg1)

	var _context *cairo.Context // out

	_context = (*cairo.Context)(unsafe.Pointer(*C.cairo_t))
	runtime.SetFinalizer(_context, func(v *cairo.Context) {
		C.free(unsafe.Pointer(v))
	})

	return _context
}

// AppendColor creates a new render node drawing the @color into the given
// @bounds and appends it to the current render node of @snapshot.
//
// You should try to avoid calling this function if @color is transparent.
func (s *SnapshotClass) AppendColor(color *gdk.RGBA, bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GdkRGBA         // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(*gdk.RGBA))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(*graphene.Rect))

	C.gtk_snapshot_append_color(_arg0, _arg1, _arg2)
}

// AppendInsetShadow appends an inset shadow into the box given by @outline.
func (s *SnapshotClass) AppendInsetShadow(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out
	var _arg2 *C.GdkRGBA        // out
	var _arg3 C.float           // out
	var _arg4 C.float           // out
	var _arg5 C.float           // out
	var _arg6 C.float           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GskRoundedRect)(unsafe.Pointer(*gsk.RoundedRect))
	_arg2 = (*C.GdkRGBA)(unsafe.Pointer(*gdk.RGBA))
	_arg3 = C.float(dx)
	_arg4 = C.float(dy)
	_arg5 = C.float(spread)
	_arg6 = C.float(blurRadius)

	C.gtk_snapshot_append_inset_shadow(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (s *SnapshotClass) AppendLayout(layout pango.Layout, color *gdk.RGBA) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 *C.PangoLayout // out
	var _arg2 *C.GdkRGBA     // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.PangoLayout)(unsafe.Pointer((&pango.Layout).Native()))
	_arg2 = (*C.GdkRGBA)(unsafe.Pointer(*gdk.RGBA))

	C.gtk_snapshot_append_layout(_arg0, _arg1, _arg2)
}

// AppendNode appends @node to the current render node of @snapshot, without
// changing the current node.
//
// If @snapshot does not have a current node yet, @node will become the initial
// node.
func (s *SnapshotClass) AppendNode(node gsk.RenderNode) {
	var _arg0 *C.GtkSnapshot   // out
	var _arg1 *C.GskRenderNode // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer((&gsk.RenderNode).Native()))

	C.gtk_snapshot_append_node(_arg0, _arg1)
}

// AppendOutsetShadow appends an outset shadow node around the box given by
// @outline.
func (s *SnapshotClass) AppendOutsetShadow(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out
	var _arg2 *C.GdkRGBA        // out
	var _arg3 C.float           // out
	var _arg4 C.float           // out
	var _arg5 C.float           // out
	var _arg6 C.float           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GskRoundedRect)(unsafe.Pointer(*gsk.RoundedRect))
	_arg2 = (*C.GdkRGBA)(unsafe.Pointer(*gdk.RGBA))
	_arg3 = C.float(dx)
	_arg4 = C.float(dy)
	_arg5 = C.float(spread)
	_arg6 = C.float(blurRadius)

	C.gtk_snapshot_append_outset_shadow(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// AppendTexture creates a new render node drawing the @texture into the given
// @bounds and appends it to the current render node of @snapshot.
func (s *SnapshotClass) AppendTexture(texture gdk.Texture, bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GdkTexture      // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GdkTexture)(unsafe.Pointer((&gdk.Texture).Native()))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(*graphene.Rect))

	C.gtk_snapshot_append_texture(_arg0, _arg1, _arg2)
}

// GLShaderPopTexture removes the top element from the stack of render nodes and
// adds it to the nearest `GskGLShaderNode` below it.
//
// This must be called the same number of times as the number of textures is
// needed for the shader in [method@Gtk.Snapshot.push_gl_shader].
func (s *SnapshotClass) GLShaderPopTexture() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))

	C.gtk_snapshot_gl_shader_pop_texture(_arg0)
}

// Perspective applies a perspective projection transform.
//
// See [method@Gsk.Transform.perspective] for a discussion on the details.
func (s *SnapshotClass) Perspective(depth float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = C.float(depth)

	C.gtk_snapshot_perspective(_arg0, _arg1)
}

// Pop removes the top element from the stack of render nodes, and appends it to
// the node underneath it.
func (s *SnapshotClass) Pop() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))

	C.gtk_snapshot_pop(_arg0)
}

// PushBlur blurs an image.
//
// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
func (s *SnapshotClass) PushBlur(radius float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = C.double(radius)

	C.gtk_snapshot_push_blur(_arg0, _arg1)
}

// PushClip clips an image to a rectangle.
//
// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
func (s *SnapshotClass) PushClip(bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(*graphene.Rect))

	C.gtk_snapshot_push_clip(_arg0, _arg1)
}

// PushColorMatrix modifies the colors of an image by applying an affine
// transformation in RGB space.
//
// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
func (s *SnapshotClass) PushColorMatrix(colorMatrix *graphene.Matrix, colorOffset *graphene.Vec4) {
	var _arg0 *C.GtkSnapshot       // out
	var _arg1 *C.graphene_matrix_t // out
	var _arg2 *C.graphene_vec4_t   // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(*graphene.Matrix))
	_arg2 = (*C.graphene_vec4_t)(unsafe.Pointer(*graphene.Vec4))

	C.gtk_snapshot_push_color_matrix(_arg0, _arg1, _arg2)
}

// PushCrossFade snapshots a cross-fade operation between two images with the
// given @progress.
//
// Until the first call to [method@Gtk.Snapshot.pop], the start image will be
// snapshot. After that call, the end image will be recorded until the second
// call to [method@Gtk.Snapshot.pop].
//
// Calling this function requires two subsequent calls to
// [method@Gtk.Snapshot.pop].
func (s *SnapshotClass) PushCrossFade(progress float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = C.double(progress)

	C.gtk_snapshot_push_cross_fade(_arg0, _arg1)
}

// PushOpacity modifies the opacity of an image.
//
// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
func (s *SnapshotClass) PushOpacity(opacity float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = C.double(opacity)

	C.gtk_snapshot_push_opacity(_arg0, _arg1)
}

// PushRepeat creates a node that repeats the child node.
//
// The child is recorded until the next call to [method@Gtk.Snapshot.pop].
func (s *SnapshotClass) PushRepeat(bounds *graphene.Rect, childBounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.graphene_rect_t)(unsafe.Pointer(*graphene.Rect))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(*graphene.Rect))

	C.gtk_snapshot_push_repeat(_arg0, _arg1, _arg2)
}

// PushRoundedClip clips an image to a rounded rectangle.
//
// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
func (s *SnapshotClass) PushRoundedClip(bounds *gsk.RoundedRect) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GskRoundedRect)(unsafe.Pointer(*gsk.RoundedRect))

	C.gtk_snapshot_push_rounded_clip(_arg0, _arg1)
}

// PushShadow applies a shadow to an image.
//
// The image is recorded until the next call to [method@Gtk.Snapshot.pop].
func (s *SnapshotClass) PushShadow(shadow *gsk.Shadow, nShadows uint) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 *C.GskShadow   // out
	var _arg2 C.gsize        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GskShadow)(unsafe.Pointer(*gsk.Shadow))
	_arg2 = C.gsize(nShadows)

	C.gtk_snapshot_push_shadow(_arg0, _arg1, _arg2)
}

// RenderBackground creates a render node for the CSS background according to
// @context, and appends it to the current node of @snapshot, without changing
// the current node.
func (s *SnapshotClass) RenderBackground(context StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer((&StyleContext).Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_background(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderFocus creates a render node for the focus outline according to
// @context, and appends it to the current node of @snapshot, without changing
// the current node.
func (s *SnapshotClass) RenderFocus(context StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer((&StyleContext).Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_focus(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderFrame creates a render node for the CSS border according to @context,
// and appends it to the current node of @snapshot, without changing the current
// node.
func (s *SnapshotClass) RenderFrame(context StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer((&StyleContext).Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_frame(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderLayout creates a render node for rendering @layout according to the
// style information in @context, and appends it to the current node of
// @snapshot, without changing the current node.
func (s *SnapshotClass) RenderLayout(context StyleContext, x float64, y float64, layout pango.Layout) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 *C.PangoLayout     // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer((&StyleContext).Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = (*C.PangoLayout)(unsafe.Pointer((&pango.Layout).Native()))

	C.gtk_snapshot_render_layout(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// Restore restores @snapshot to the state saved by a preceding call to
// gtk_snapshot_save() and removes that state from the stack of saved states.
func (s *SnapshotClass) Restore() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))

	C.gtk_snapshot_restore(_arg0)
}

// Rotate rotates @@snapshot's coordinate system by @angle degrees in 2D space -
// or in 3D speak, rotates around the Z axis.
//
// To rotate around other axes, use [method@Gsk.Transform.rotate_3d].
func (s *SnapshotClass) Rotate(angle float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = C.float(angle)

	C.gtk_snapshot_rotate(_arg0, _arg1)
}

// Rotate3D rotates @snapshot's coordinate system by @angle degrees around
// @axis.
//
// For a rotation in 2D space, use [method@Gsk.Transform.rotate].
func (s *SnapshotClass) Rotate3D(angle float32, axis *graphene.Vec3) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 C.float            // out
	var _arg2 *C.graphene_vec3_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = C.float(angle)
	_arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(*graphene.Vec3))

	C.gtk_snapshot_rotate_3d(_arg0, _arg1, _arg2)
}

// Save makes a copy of the current state of @snapshot and saves it on an
// internal stack.
//
// When [method@Gtk.Snapshot.restore] is called, @snapshot will be restored to
// the saved state. Multiple calls to gtk_snapshot_save() and
// gtk_snapshot_restore() can be nested; each call to gtk_snapshot_restore()
// restores the state from the matching paired gtk_snapshot_save().
//
// It is necessary to clear all saved states with corresponding calls to
// gtk_snapshot_restore().
func (s *SnapshotClass) Save() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))

	C.gtk_snapshot_save(_arg0)
}

// Scale scales @snapshot's coordinate system in 2-dimensional space by the
// given factors.
//
// Use [method@Gtk.Snapshot.scale_3d] to scale in all 3 dimensions.
func (s *SnapshotClass) Scale(factorX float32, factorY float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out
	var _arg2 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = C.float(factorX)
	_arg2 = C.float(factorY)

	C.gtk_snapshot_scale(_arg0, _arg1, _arg2)
}

// Scale3D scales @snapshot's coordinate system by the given factors.
func (s *SnapshotClass) Scale3D(factorX float32, factorY float32, factorZ float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out
	var _arg2 C.float        // out
	var _arg3 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = C.float(factorX)
	_arg2 = C.float(factorY)
	_arg3 = C.float(factorZ)

	C.gtk_snapshot_scale_3d(_arg0, _arg1, _arg2, _arg3)
}

// ToNode returns the render node that was constructed by @snapshot.
//
// After calling this function, it is no longer possible to add more nodes to
// @snapshot. The only function that should be called after this is
// g_object_unref().
func (s *SnapshotClass) ToNode() *gsk.RenderNodeClass {
	var _arg0 *C.GtkSnapshot   // out
	var _cret *C.GskRenderNode // in

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))

	_cret = C.gtk_snapshot_to_node(_arg0)

	var _renderNode *gsk.RenderNodeClass // out

	_renderNode = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*gsk.RenderNodeClass)

	return _renderNode
}

// Transform transforms @snapshot's coordinate system with the given @transform.
func (s *SnapshotClass) Transform(transform *gsk.Transform) {
	var _arg0 *C.GtkSnapshot  // out
	var _arg1 *C.GskTransform // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.GskTransform)(unsafe.Pointer(*gsk.Transform))

	C.gtk_snapshot_transform(_arg0, _arg1)
}

// TransformMatrix transforms @snapshot's coordinate system with the given
// @matrix.
func (s *SnapshotClass) TransformMatrix(matrix *graphene.Matrix) {
	var _arg0 *C.GtkSnapshot       // out
	var _arg1 *C.graphene_matrix_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(*graphene.Matrix))

	C.gtk_snapshot_transform_matrix(_arg0, _arg1)
}

// Translate translates @snapshot's coordinate system by @point in 2-dimensional
// space.
func (s *SnapshotClass) Translate(point *graphene.Point) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_point_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(*graphene.Point))

	C.gtk_snapshot_translate(_arg0, _arg1)
}

// Translate3D translates @snapshot's coordinate system by @point.
func (s *SnapshotClass) Translate3D(point *graphene.Point3D) {
	var _arg0 *C.GtkSnapshot        // out
	var _arg1 *C.graphene_point3d_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer((&Snapshot).Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(*graphene.Point3D))

	C.gtk_snapshot_translate_3d(_arg0, _arg1)
}
