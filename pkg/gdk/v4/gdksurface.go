// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_surface_get_type()), F: marshalSurfacer},
	})
}

// Surface: GdkSurface is a rectangular region on the screen.
//
// It’s a low-level object, used to implement high-level objects such as
// gtk.Window or gtk.Dialog in GTK.
//
// The surfaces you see in practice are either gdk.Toplevel or gdk.Popup, and
// those interfaces provide much of the required API to interact with these
// surfaces. Other, more specialized surface types exist, but you will rarely
// interact with them directly.
type Surface struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Surface)(nil)

// Surfacer describes Surface's abstract methods.
type Surfacer interface {
	// Beep emits a short beep associated to surface.
	Beep()
	// CreateCairoContext creates a new GdkCairoContext for rendering on
	// surface.
	CreateCairoContext() CairoContexter
	// CreateGLContext creates a new GdkGLContext for the GdkSurface.
	CreateGLContext() (GLContexter, error)
	// CreateSimilarSurface: create a new Cairo surface that is as compatible as
	// possible with the given surface.
	CreateSimilarSurface(content cairo.Content, width int, height int) *cairo.Surface
	// CreateVulkanContext creates a new GdkVulkanContext for rendering on
	// surface.
	CreateVulkanContext() (VulkanContexter, error)
	// Destroy destroys the window system resources associated with surface and
	// decrements surface's reference count.
	Destroy()
	// Cursor retrieves a GdkCursor pointer for the cursor currently set on the
	// GdkSurface.
	Cursor() *Cursor
	// DeviceCursor retrieves a GdkCursor pointer for the device currently set
	// on the specified GdkSurface.
	DeviceCursor(device Devicer) *Cursor
	// DevicePosition obtains the current device position and modifier state.
	DevicePosition(device Devicer) (x float64, y float64, mask ModifierType, ok bool)
	// Display gets the GdkDisplay associated with a GdkSurface.
	Display() *Display
	// FrameClock gets the frame clock for the surface.
	FrameClock() FrameClocker
	// Height returns the height of the given surface.
	Height() int
	// Mapped checks whether the surface has been mapped.
	Mapped() bool
	// ScaleFactor returns the internal scale factor that maps from surface
	// coordinates to the actual device pixels.
	ScaleFactor() int
	// Width returns the width of the given surface.
	Width() int
	// Hide the surface.
	Hide()
	// IsDestroyed: check to see if a surface is destroyed.
	IsDestroyed() bool
	// QueueRender forces a gdk.Surface::render signal emission for surface to
	// be scheduled.
	QueueRender()
	// RequestLayout: request a layout phase from the surface's frame clock.
	RequestLayout()
	// SetCursor sets the default mouse pointer for a GdkSurface.
	SetCursor(cursor *Cursor)
	// SetDeviceCursor sets a specific GdkCursor for a given device when it gets
	// inside surface.
	SetDeviceCursor(device Devicer, cursor *Cursor)
	// SetInputRegion: apply the region to the surface for the purpose of event
	// handling.
	SetInputRegion(region *cairo.Region)
	// SetOpaqueRegion marks a region of the GdkSurface as opaque.
	SetOpaqueRegion(region *cairo.Region)
}

var _ Surfacer = (*Surface)(nil)

func wrapSurface(obj *externglib.Object) *Surface {
	return &Surface{
		Object: obj,
	}
}

func marshalSurfacer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSurface(obj), nil
}

// NewSurfacePopup: create a new popup surface.
//
// The surface will be attached to parent and can be positioned relative to it
// using gdk.Popup.Present().
func NewSurfacePopup(parent Surfacer, autohide bool) *Surface {
	var _arg1 *C.GdkSurface // out
	var _arg2 C.gboolean    // out
	var _cret *C.GdkSurface // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer((parent).(gextras.Nativer).Native()))
	if autohide {
		_arg2 = C.TRUE
	}

	_cret = C.gdk_surface_new_popup(_arg1, _arg2)

	var _surface *Surface // out

	_surface = wrapSurface(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _surface
}

// NewSurfaceToplevel creates a new toplevel surface.
func NewSurfaceToplevel(display *Display) *Surface {
	var _arg1 *C.GdkDisplay // out
	var _cret *C.GdkSurface // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_surface_new_toplevel(_arg1)

	var _surface *Surface // out

	_surface = wrapSurface(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _surface
}

// Beep emits a short beep associated to surface.
//
// If the display of surface does not support per-surface beeps, emits a short
// beep on the display just as gdk.Display.Beep().
func (surface *Surface) Beep() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gdk_surface_beep(_arg0)
}

// CreateCairoContext creates a new GdkCairoContext for rendering on surface.
func (surface *Surface) CreateCairoContext() CairoContexter {
	var _arg0 *C.GdkSurface      // out
	var _cret *C.GdkCairoContext // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_create_cairo_context(_arg0)

	var _cairoContext CairoContexter // out

	_cairoContext = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(CairoContexter)

	return _cairoContext
}

// CreateGLContext creates a new GdkGLContext for the GdkSurface.
//
// The context is disconnected from any particular surface or surface. If the
// creation of the GdkGLContext failed, error will be set. Before using the
// returned GdkGLContext, you will need to call gdk.GLContext.MakeCurrent() or
// gdk.GLContext.Realize().
func (surface *Surface) CreateGLContext() (GLContexter, error) {
	var _arg0 *C.GdkSurface   // out
	var _cret *C.GdkGLContext // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_create_gl_context(_arg0, &_cerr)

	var _glContext GLContexter // out
	var _goerr error           // out

	_glContext = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(GLContexter)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _glContext, _goerr
}

// CreateSimilarSurface: create a new Cairo surface that is as compatible as
// possible with the given surface.
//
// For example the new surface will have the same fallback resolution and font
// options as surface. Generally, the new surface will also use the same backend
// as surface, unless that is not possible for some reason. The type of the
// returned surface may be examined with cairo_surface_get_type().
//
// Initially the surface contents are all 0 (transparent if contents have
// transparency, black otherwise.)
//
// This function always returns a valid pointer, but it will return a pointer to
// a “nil” surface if other is already in an error state or any other error
// occurs.
func (surface *Surface) CreateSimilarSurface(content cairo.Content, width int, height int) *cairo.Surface {
	var _arg0 *C.GdkSurface      // out
	var _arg1 C.cairo_content_t  // out
	var _arg2 C.int              // out
	var _arg3 C.int              // out
	var _cret *C.cairo_surface_t // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = C.cairo_content_t(content)
	_arg2 = C.int(width)
	_arg3 = C.int(height)

	_cret = C.gdk_surface_create_similar_surface(_arg0, _arg1, _arg2, _arg3)

	var _ret *cairo.Surface // out

	_ret = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
	C.cairo_surface_reference(_cret)
	runtime.SetFinalizer(_ret, func(v *cairo.Surface) {
		C.cairo_surface_destroy((*C.cairo_surface_t)(unsafe.Pointer(v.Native())))
	})

	return _ret
}

// CreateVulkanContext creates a new GdkVulkanContext for rendering on surface.
//
// If the creation of the GdkVulkanContext failed, error will be set.
func (surface *Surface) CreateVulkanContext() (VulkanContexter, error) {
	var _arg0 *C.GdkSurface       // out
	var _cret *C.GdkVulkanContext // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_create_vulkan_context(_arg0, &_cerr)

	var _vulkanContext VulkanContexter // out
	var _goerr error                   // out

	_vulkanContext = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(VulkanContexter)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _vulkanContext, _goerr
}

// Destroy destroys the window system resources associated with surface and
// decrements surface's reference count.
//
// The window system resources for all children of surface are also destroyed,
// but the children’s reference counts are not decremented.
//
// Note that a surface will not be destroyed automatically when its reference
// count reaches zero. You must call this function yourself before that happens.
func (surface *Surface) Destroy() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gdk_surface_destroy(_arg0)
}

// Cursor retrieves a GdkCursor pointer for the cursor currently set on the
// GdkSurface.
//
// If the return value is NULL then there is no custom cursor set on the
// surface, and it is using the cursor for its parent surface.
func (surface *Surface) Cursor() *Cursor {
	var _arg0 *C.GdkSurface // out
	var _cret *C.GdkCursor  // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_get_cursor(_arg0)

	var _cursor *Cursor // out

	_cursor = wrapCursor(externglib.Take(unsafe.Pointer(_cret)))

	return _cursor
}

// DeviceCursor retrieves a GdkCursor pointer for the device currently set on
// the specified GdkSurface.
//
// If the return value is NULL then there is no custom cursor set on the
// specified surface, and it is using the cursor for its parent surface.
func (surface *Surface) DeviceCursor(device Devicer) *Cursor {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkDevice  // out
	var _cret *C.GdkCursor  // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((device).(gextras.Nativer).Native()))

	_cret = C.gdk_surface_get_device_cursor(_arg0, _arg1)

	var _cursor *Cursor // out

	_cursor = wrapCursor(externglib.Take(unsafe.Pointer(_cret)))

	return _cursor
}

// DevicePosition obtains the current device position and modifier state.
//
// The position is given in coordinates relative to the upper left corner of
// surface.
func (surface *Surface) DevicePosition(device Devicer) (x float64, y float64, mask ModifierType, ok bool) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.GdkDevice      // out
	var _arg2 C.double          // in
	var _arg3 C.double          // in
	var _arg4 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((device).(gextras.Nativer).Native()))

	_cret = C.gdk_surface_get_device_position(_arg0, _arg1, &_arg2, &_arg3, &_arg4)

	var _x float64         // out
	var _y float64         // out
	var _mask ModifierType // out
	var _ok bool           // out

	_x = float64(_arg2)
	_y = float64(_arg3)
	_mask = ModifierType(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _mask, _ok
}

// Display gets the GdkDisplay associated with a GdkSurface.
func (surface *Surface) Display() *Display {
	var _arg0 *C.GdkSurface // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_get_display(_arg0)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// FrameClock gets the frame clock for the surface.
//
// The frame clock for a surface never changes unless the surface is reparented
// to a new toplevel surface.
func (surface *Surface) FrameClock() FrameClocker {
	var _arg0 *C.GdkSurface    // out
	var _cret *C.GdkFrameClock // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_get_frame_clock(_arg0)

	var _frameClock FrameClocker // out

	_frameClock = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(FrameClocker)

	return _frameClock
}

// Height returns the height of the given surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels” (see
// gdk.Surface.GetScaleFactor()).
func (surface *Surface) Height() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_get_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Mapped checks whether the surface has been mapped.
//
// A surface is mapped with gdk.Toplevel.Present() or gdk.Popup.Present().
func (surface *Surface) Mapped() bool {
	var _arg0 *C.GdkSurface // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_get_mapped(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ScaleFactor returns the internal scale factor that maps from surface
// coordinates to the actual device pixels.
//
// On traditional systems this is 1, but on very high density outputs this can
// be a higher value (often 2). A higher value means that drawing is
// automatically scaled up to a higher resolution, so any code doing drawing
// will automatically look nicer. However, if you are supplying pixel-based data
// the scale value can be used to determine whether to use a pixel resource with
// higher resolution data.
//
// The scale of a surface may change during runtime.
func (surface *Surface) ScaleFactor() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_get_scale_factor(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Width returns the width of the given surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels” (see
// gdk.Surface.GetScaleFactor()).
func (surface *Surface) Width() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Hide the surface.
//
// For toplevel surfaces, withdraws them, so they will no longer be known to the
// window manager; for all surfaces, unmaps them, so they won’t be displayed.
// Normally done automatically as part of gtk.Widget.Hide().
func (surface *Surface) Hide() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gdk_surface_hide(_arg0)
}

// IsDestroyed: check to see if a surface is destroyed.
func (surface *Surface) IsDestroyed() bool {
	var _arg0 *C.GdkSurface // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_surface_is_destroyed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// QueueRender forces a gdk.Surface::render signal emission for surface to be
// scheduled.
//
// This function is useful for implementations that track invalid regions on
// their own.
func (surface *Surface) QueueRender() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gdk_surface_queue_render(_arg0)
}

// RequestLayout: request a layout phase from the surface's frame clock.
//
// See gdk.FrameClock.RequestPhase().
func (surface *Surface) RequestLayout() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gdk_surface_request_layout(_arg0)
}

// SetCursor sets the default mouse pointer for a GdkSurface.
//
// Passing NULL for the cursor argument means that surface will use the cursor
// of its parent surface. Most surfaces should use this default. Note that
// cursor must be for the same display as surface.
//
// Use gdk.Cursor.NewFromName or gdk.Cursor.NewFromTexture to create the cursor.
// To make the cursor invisible, use GDK_BLANK_CURSOR.
func (surface *Surface) SetCursor(cursor *Cursor) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkCursor  // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	C.gdk_surface_set_cursor(_arg0, _arg1)
}

// SetDeviceCursor sets a specific GdkCursor for a given device when it gets
// inside surface.
//
// Passing NULL for the cursor argument means that surface will use the cursor
// of its parent surface. Most surfaces should use this default.
//
// Use gdk.Cursor.NewFromName or gdk.Cursor.NewFromTexture to create the cursor.
// To make the cursor invisible, use GDK_BLANK_CURSOR.
func (surface *Surface) SetDeviceCursor(device Devicer, cursor *Cursor) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkDevice  // out
	var _arg2 *C.GdkCursor  // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((device).(gextras.Nativer).Native()))
	_arg2 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	C.gdk_surface_set_device_cursor(_arg0, _arg1, _arg2)
}

// SetInputRegion: apply the region to the surface for the purpose of event
// handling.
//
// Mouse events which happen while the pointer position corresponds to an unset
// bit in the mask will be passed on the surface below surface.
//
// An input region is typically used with RGBA surfaces. The alpha channel of
// the surface defines which pixels are invisible and allows for nicely
// antialiased borders, and the input region controls where the surface is
// “clickable”.
//
// Use gdk.Display.SupportsInputShapes() to find out if a particular backend
// supports input regions.
func (surface *Surface) SetInputRegion(region *cairo.Region) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_surface_set_input_region(_arg0, _arg1)
}

// SetOpaqueRegion marks a region of the GdkSurface as opaque.
//
// For optimisation purposes, compositing window managers may like to not draw
// obscured regions of surfaces, or turn off blending during for these regions.
// With RGB windows with no transparency, this is just the shape of the window,
// but with ARGB32 windows, the compositor does not know what regions of the
// window are transparent or not.
//
// This function only works for toplevel surfaces.
//
// GTK will update this property automatically if the surface background is
// opaque, as we know where the opaque regions are. If your surface background
// is not opaque, please update this property in your WidgetClass.css_changed()
// handler.
func (surface *Surface) SetOpaqueRegion(region *cairo.Region) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_surface_set_opaque_region(_arg0, _arg1)
}
