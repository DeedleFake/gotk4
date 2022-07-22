// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gdk4_Surface_ConnectEvent(gpointer, gpointer*, guintptr);
// extern gboolean _gotk4_gdk4_Surface_ConnectRender(gpointer, cairo_region_t*, guintptr);
// extern void _gotk4_gdk4_Surface_ConnectEnterMonitor(gpointer, GdkMonitor*, guintptr);
// extern void _gotk4_gdk4_Surface_ConnectLayout(gpointer, gint, gint, guintptr);
// extern void _gotk4_gdk4_Surface_ConnectLeaveMonitor(gpointer, GdkMonitor*, guintptr);
import "C"

// GType values.
var (
	GTypeSurface = coreglib.Type(C.gdk_surface_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSurface, F: marshalSurface},
	})
}

// SurfaceOverrider contains methods that are overridable.
type SurfaceOverrider interface {
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
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Surface)(nil)
)

// Surfacer describes types inherited from class Surface.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Surfacer interface {
	coreglib.Objector
	baseSurface() *Surface
}

var _ Surfacer = (*Surface)(nil)

func initClassSurface(gclass unsafe.Pointer, goval any) {
}

func wrapSurface(obj *coreglib.Object) *Surface {
	return &Surface{
		Object: obj,
	}
}

func marshalSurface(p uintptr) (interface{}, error) {
	return wrapSurface(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (surface *Surface) baseSurface() *Surface {
	return surface
}

// BaseSurface returns the underlying base object.
func BaseSurface(obj Surfacer) *Surface {
	return obj.baseSurface()
}

//export _gotk4_gdk4_Surface_ConnectEnterMonitor
func _gotk4_gdk4_Surface_ConnectEnterMonitor(arg0 C.gpointer, arg1 *C.GdkMonitor, arg2 C.guintptr) {
	var f func(monitor *Monitor)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(monitor *Monitor))
	}

	var _monitor *Monitor // out

	_monitor = wrapMonitor(coreglib.Take(unsafe.Pointer(arg1)))

	f(_monitor)
}

// ConnectEnterMonitor is emitted when surface starts being present on the
// monitor.
func (surface *Surface) ConnectEnterMonitor(f func(monitor *Monitor)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(surface, "enter-monitor", false, unsafe.Pointer(C._gotk4_gdk4_Surface_ConnectEnterMonitor), f)
}

//export _gotk4_gdk4_Surface_ConnectEvent
func _gotk4_gdk4_Surface_ConnectEvent(arg0 C.gpointer, arg1 *C.gpointer, arg2 C.guintptr) (cret C.gboolean) {
	var f func(event Eventer) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(event Eventer) (ok bool))
	}

	var _event Eventer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Eventer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Eventer)
			return ok
		})
		rv, ok := casted.(Eventer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Eventer")
		}
		_event = rv
	}

	ok := f(_event)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectEvent is emitted when GDK receives an input event for surface.
func (surface *Surface) ConnectEvent(f func(event Eventer) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(surface, "event", false, unsafe.Pointer(C._gotk4_gdk4_Surface_ConnectEvent), f)
}

//export _gotk4_gdk4_Surface_ConnectLayout
func _gotk4_gdk4_Surface_ConnectLayout(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(width, height int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(width, height int))
	}

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	f(_width, _height)
}

// ConnectLayout is emitted when the size of surface is changed, or when
// relayout should be performed.
//
// Surface size is reported in ”application pixels”, not ”device pixels” (see
// gdk_surface_get_scale_factor()).
func (surface *Surface) ConnectLayout(f func(width, height int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(surface, "layout", false, unsafe.Pointer(C._gotk4_gdk4_Surface_ConnectLayout), f)
}

//export _gotk4_gdk4_Surface_ConnectLeaveMonitor
func _gotk4_gdk4_Surface_ConnectLeaveMonitor(arg0 C.gpointer, arg1 *C.GdkMonitor, arg2 C.guintptr) {
	var f func(monitor *Monitor)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(monitor *Monitor))
	}

	var _monitor *Monitor // out

	_monitor = wrapMonitor(coreglib.Take(unsafe.Pointer(arg1)))

	f(_monitor)
}

// ConnectLeaveMonitor is emitted when surface stops being present on the
// monitor.
func (surface *Surface) ConnectLeaveMonitor(f func(monitor *Monitor)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(surface, "leave-monitor", false, unsafe.Pointer(C._gotk4_gdk4_Surface_ConnectLeaveMonitor), f)
}

//export _gotk4_gdk4_Surface_ConnectRender
func _gotk4_gdk4_Surface_ConnectRender(arg0 C.gpointer, arg1 *C.cairo_region_t, arg2 C.guintptr) (cret C.gboolean) {
	var f func(region *cairo.Region) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(region *cairo.Region) (ok bool))
	}

	var _region *cairo.Region // out

	{
		_pp := &struct{ p unsafe.Pointer }{unsafe.Pointer(arg1)}
		_region = (*cairo.Region)(unsafe.Pointer(_pp))
	}
	C.cairo_region_reference(arg1)
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.cairo_region_destroy((*C.cairo_region_t)(unsafe.Pointer(v.Native())))
	})

	ok := f(_region)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectRender is emitted when part of the surface needs to be redrawn.
func (surface *Surface) ConnectRender(f func(region *cairo.Region) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(surface, "render", false, unsafe.Pointer(C._gotk4_gdk4_Surface_ConnectRender), f)
}

// NewSurfacePopup: create a new popup surface.
//
// The surface will be attached to parent and can be positioned relative to it
// using gdk.Popup.Present().
//
// The function takes the following parameters:
//
//    - parent surface to attach the surface to.
//    - autohide: whether to hide the surface on outside clicks.
//
// The function returns the following values:
//
//    - surface: new GdkSurface.
//
func NewSurfacePopup(parent Surfacer, autohide bool) *Surface {
	var _arg1 *C.GdkSurface // out
	var _arg2 C.gboolean    // out
	var _cret *C.GdkSurface // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	if autohide {
		_arg2 = C.TRUE
	}

	_cret = C.gdk_surface_new_popup(_arg1, _arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(autohide)

	var _surface *Surface // out

	_surface = wrapSurface(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _surface
}

// NewSurfaceToplevel creates a new toplevel surface.
//
// The function takes the following parameters:
//
//    - display to create the surface on.
//
// The function returns the following values:
//
//    - surface: new GdkSurface.
//
func NewSurfaceToplevel(display *Display) *Surface {
	var _arg1 *C.GdkDisplay // out
	var _cret *C.GdkSurface // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gdk_surface_new_toplevel(_arg1)
	runtime.KeepAlive(display)

	var _surface *Surface // out

	_surface = wrapSurface(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _surface
}

// Beep emits a short beep associated to surface.
//
// If the display of surface does not support per-surface beeps, emits a short
// beep on the display just as gdk.Display.Beep().
func (surface *Surface) Beep() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	C.gdk_surface_beep(_arg0)
	runtime.KeepAlive(surface)
}

// CreateCairoContext creates a new GdkCairoContext for rendering on surface.
//
// The function returns the following values:
//
//    - cairoContext: newly created GdkCairoContext.
//
func (surface *Surface) CreateCairoContext() CairoContexter {
	var _arg0 *C.GdkSurface      // out
	var _cret *C.GdkCairoContext // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_create_cairo_context(_arg0)
	runtime.KeepAlive(surface)

	var _cairoContext CairoContexter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.CairoContexter is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CairoContexter)
			return ok
		})
		rv, ok := casted.(CairoContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.CairoContexter")
		}
		_cairoContext = rv
	}

	return _cairoContext
}

// CreateGLContext creates a new GdkGLContext for the GdkSurface.
//
// The context is disconnected from any particular surface or surface. If the
// creation of the GdkGLContext failed, error will be set. Before using the
// returned GdkGLContext, you will need to call gdk.GLContext.MakeCurrent() or
// gdk.GLContext.Realize().
//
// The function returns the following values:
//
//    - glContext: newly created GdkGLContext, or NULL on error.
//
func (surface *Surface) CreateGLContext() (GLContexter, error) {
	var _arg0 *C.GdkSurface   // out
	var _cret *C.GdkGLContext // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_create_gl_context(_arg0, &_cerr)
	runtime.KeepAlive(surface)

	var _glContext GLContexter // out
	var _goerr error           // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.GLContexter is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(GLContexter)
			return ok
		})
		rv, ok := casted.(GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.GLContexter")
		}
		_glContext = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

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
//
// The function takes the following parameters:
//
//    - content for the new surface.
//    - width of the new surface.
//    - height of the new surface.
//
// The function returns the following values:
//
//    - ret: pointer to the newly allocated surface. The caller owns the surface
//      and should call cairo_surface_destroy() when done with it.
//
func (surface *Surface) CreateSimilarSurface(content cairo.Content, width, height int) *cairo.Surface {
	var _arg0 *C.GdkSurface      // out
	var _arg1 C.cairo_content_t  // out
	var _arg2 C.int              // out
	var _arg3 C.int              // out
	var _cret *C.cairo_surface_t // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))
	_arg1 = C.cairo_content_t(content)
	_arg2 = C.int(width)
	_arg3 = C.int(height)

	_cret = C.gdk_surface_create_similar_surface(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(content)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ret *cairo.Surface // out

	_ret = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_ret, func(v *cairo.Surface) {
		C.cairo_surface_destroy((*C.cairo_surface_t)(unsafe.Pointer(v.Native())))
	})

	return _ret
}

// CreateVulkanContext creates a new GdkVulkanContext for rendering on surface.
//
// If the creation of the GdkVulkanContext failed, error will be set.
//
// The function returns the following values:
//
//    - vulkanContext: newly created GdkVulkanContext, or NULL on error.
//
func (surface *Surface) CreateVulkanContext() (VulkanContexter, error) {
	var _arg0 *C.GdkSurface       // out
	var _cret *C.GdkVulkanContext // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_create_vulkan_context(_arg0, &_cerr)
	runtime.KeepAlive(surface)

	var _vulkanContext VulkanContexter // out
	var _goerr error                   // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.VulkanContexter is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(VulkanContexter)
			return ok
		})
		rv, ok := casted.(VulkanContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.VulkanContexter")
		}
		_vulkanContext = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

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

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	C.gdk_surface_destroy(_arg0)
	runtime.KeepAlive(surface)
}

// Cursor retrieves a GdkCursor pointer for the cursor currently set on the
// GdkSurface.
//
// If the return value is NULL then there is no custom cursor set on the
// surface, and it is using the cursor for its parent surface.
//
// The function returns the following values:
//
//    - cursor (optional): GdkCursor, or NULL. The returned object is owned by
//      the GdkSurface and should not be unreferenced directly. Use
//      gdk.Surface.SetCursor() to unset the cursor of the surface.
//
func (surface *Surface) Cursor() *Cursor {
	var _arg0 *C.GdkSurface // out
	var _cret *C.GdkCursor  // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_get_cursor(_arg0)
	runtime.KeepAlive(surface)

	var _cursor *Cursor // out

	if _cret != nil {
		_cursor = wrapCursor(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _cursor
}

// DeviceCursor retrieves a GdkCursor pointer for the device currently set on
// the specified GdkSurface.
//
// If the return value is NULL then there is no custom cursor set on the
// specified surface, and it is using the cursor for its parent surface.
//
// The function takes the following parameters:
//
//    - device: pointer GdkDevice.
//
// The function returns the following values:
//
//    - cursor (optional): GdkCursor, or NULL. The returned object is owned by
//      the GdkSurface and should not be unreferenced directly. Use
//      gdk.Surface.SetCursor() to unset the cursor of the surface.
//
func (surface *Surface) DeviceCursor(device Devicer) *Cursor {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkDevice  // out
	var _cret *C.GdkCursor  // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_cret = C.gdk_surface_get_device_cursor(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(device)

	var _cursor *Cursor // out

	if _cret != nil {
		_cursor = wrapCursor(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _cursor
}

// DevicePosition obtains the current device position and modifier state.
//
// The position is given in coordinates relative to the upper left corner of
// surface.
//
// The function takes the following parameters:
//
//    - device: pointer GdkDevice to query to.
//
// The function returns the following values:
//
//    - x (optional): return locatio for the X coordinate of device, or NULL.
//    - y (optional): return location for the Y coordinate of device, or NULL.
//    - mask (optional): return location for the modifier mask, or NULL.
//    - ok: TRUE if the device is over the surface.
//
func (surface *Surface) DevicePosition(device Devicer) (x, y float64, mask ModifierType, ok bool) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.GdkDevice      // out
	var _arg2 C.double          // in
	var _arg3 C.double          // in
	var _arg4 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))

	_cret = C.gdk_surface_get_device_position(_arg0, _arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(device)

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
//
// The function returns the following values:
//
//    - display: GdkDisplay associated with surface.
//
func (surface *Surface) Display() *Display {
	var _arg0 *C.GdkSurface // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_get_display(_arg0)
	runtime.KeepAlive(surface)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// FrameClock gets the frame clock for the surface.
//
// The frame clock for a surface never changes unless the surface is reparented
// to a new toplevel surface.
//
// The function returns the following values:
//
//    - frameClock: frame clock.
//
func (surface *Surface) FrameClock() FrameClocker {
	var _arg0 *C.GdkSurface    // out
	var _cret *C.GdkFrameClock // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_get_frame_clock(_arg0)
	runtime.KeepAlive(surface)

	var _frameClock FrameClocker // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.FrameClocker is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(FrameClocker)
			return ok
		})
		rv, ok := casted.(FrameClocker)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.FrameClocker")
		}
		_frameClock = rv
	}

	return _frameClock
}

// Height returns the height of the given surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels” (see
// gdk.Surface.GetScaleFactor()).
//
// The function returns the following values:
//
//    - gint: height of surface.
//
func (surface *Surface) Height() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_get_height(_arg0)
	runtime.KeepAlive(surface)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Mapped checks whether the surface has been mapped.
//
// A surface is mapped with gdk.Toplevel.Present() or gdk.Popup.Present().
//
// The function returns the following values:
//
//    - ok: TRUE if the surface is mapped.
//
func (surface *Surface) Mapped() bool {
	var _arg0 *C.GdkSurface // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_get_mapped(_arg0)
	runtime.KeepAlive(surface)

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
//
// The function returns the following values:
//
//    - gint: scale factor.
//
func (surface *Surface) ScaleFactor() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_get_scale_factor(_arg0)
	runtime.KeepAlive(surface)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Width returns the width of the given surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels” (see
// gdk.Surface.GetScaleFactor()).
//
// The function returns the following values:
//
//    - gint: width of surface.
//
func (surface *Surface) Width() int {
	var _arg0 *C.GdkSurface // out
	var _cret C.int         // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_get_width(_arg0)
	runtime.KeepAlive(surface)

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

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	C.gdk_surface_hide(_arg0)
	runtime.KeepAlive(surface)
}

// IsDestroyed: check to see if a surface is destroyed.
//
// The function returns the following values:
//
//    - ok: TRUE if the surface is destroyed.
//
func (surface *Surface) IsDestroyed() bool {
	var _arg0 *C.GdkSurface // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	_cret = C.gdk_surface_is_destroyed(_arg0)
	runtime.KeepAlive(surface)

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

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	C.gdk_surface_queue_render(_arg0)
	runtime.KeepAlive(surface)
}

// RequestLayout: request a layout phase from the surface's frame clock.
//
// See gdk.FrameClock.RequestPhase().
func (surface *Surface) RequestLayout() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))

	C.gdk_surface_request_layout(_arg0)
	runtime.KeepAlive(surface)
}

// SetCursor sets the default mouse pointer for a GdkSurface.
//
// Passing NULL for the cursor argument means that surface will use the cursor
// of its parent surface. Most surfaces should use this default. Note that
// cursor must be for the same display as surface.
//
// Use gdk.Cursor.NewFromName or gdk.Cursor.NewFromTexture to create the cursor.
// To make the cursor invisible, use GDK_BLANK_CURSOR.
//
// The function takes the following parameters:
//
//    - cursor (optional): GdkCursor.
//
func (surface *Surface) SetCursor(cursor *Cursor) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkCursor  // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))
	if cursor != nil {
		_arg1 = (*C.GdkCursor)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))
	}

	C.gdk_surface_set_cursor(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(cursor)
}

// SetDeviceCursor sets a specific GdkCursor for a given device when it gets
// inside surface.
//
// Passing NULL for the cursor argument means that surface will use the cursor
// of its parent surface. Most surfaces should use this default.
//
// Use gdk.Cursor.NewFromName or gdk.Cursor.NewFromTexture to create the cursor.
// To make the cursor invisible, use GDK_BLANK_CURSOR.
//
// The function takes the following parameters:
//
//    - device: pointer GdkDevice.
//    - cursor: GdkCursor.
//
func (surface *Surface) SetDeviceCursor(device Devicer, cursor *Cursor) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkDevice  // out
	var _arg2 *C.GdkCursor  // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	_arg2 = (*C.GdkCursor)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))

	C.gdk_surface_set_device_cursor(_arg0, _arg1, _arg2)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(device)
	runtime.KeepAlive(cursor)
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
//
// The function takes the following parameters:
//
//    - region of surface to be reactive.
//
func (surface *Surface) SetInputRegion(region *cairo.Region) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_surface_set_input_region(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(region)
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
//
// The function takes the following parameters:
//
//    - region (optional): region, or NULL.
//
func (surface *Surface) SetOpaqueRegion(region *cairo.Region) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))
	if region != nil {
		_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))
	}

	C.gdk_surface_set_opaque_region(_arg0, _arg1)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(region)
}
