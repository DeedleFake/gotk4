// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_gl_context_get_type()), F: marshalGLContexter},
	})
}

// GLContext: GdkGLContext is an object representing a platform-specific OpenGL
// draw context.
//
// GdkGLContexts are created for a surface using gdk.Surface.CreateGLContext(),
// and the context will match the the characteristics of the surface.
//
// A GdkGLContext is not tied to any particular normal framebuffer. For
// instance, it cannot draw to the surface back buffer. The GDK repaint system
// is in full control of the painting to that. Instead, you can create render
// buffers or textures and use cairo_draw_from_gl in the draw function of your
// widget to draw them. Then GDK will handle the integration of your rendering
// with that of other widgets.
//
// Support for GdkGLContext is platform-specific and context creation can fail,
// returning NULL context.
//
// A GdkGLContext has to be made "current" in order to start using it, otherwise
// any OpenGL call will be ignored.
//
//
// Creating a new OpenGL context
//
// In order to create a new GdkGLContext instance you need a GdkSurface, which
// you typically get during the realize call of a widget.
//
// A GdkGLContext is not realized until either gdk.GLContext.MakeCurrent() or
// gdk.GLContext.Realize() is called. It is possible to specify details of the
// GL context like the OpenGL version to be used, or whether the GL context
// should have extra state validation enabled after calling
// gdk.Surface.CreateGLContext() by calling gdk.GLContext.Realize(). If the
// realization fails you have the option to change the settings of the
// GdkGLContext and try again.
//
//
// Using a GdkGLContext
//
// You will need to make the GdkGLContext the current context before issuing
// OpenGL calls; the system sends OpenGL commands to whichever context is
// current. It is possible to have multiple contexts, so you always need to
// ensure that the one which you want to draw with is the current one before
// issuing commands:
//
//    gdk_gl_context_make_current (context);
//
//
// You can now perform your drawing using OpenGL commands.
//
// You can check which GdkGLContext is the current one by using
// gdk.GLContext().GetCurrent; you can also unset any GdkGLContext that is
// currently set by calling gdk.GLContext().ClearCurrent.
type GLContext struct {
	DrawContext
}

var _ gextras.Nativer = (*GLContext)(nil)

// GLContexter describes GLContext's abstract methods.
type GLContexter interface {
	// DebugEnabled retrieves whether the context is doing extra validations and
	// runtime checking.
	DebugEnabled() bool
	// Display retrieves the display the context is created for
	Display() *Display
	// ForwardCompatible retrieves whether the context is forward-compatible.
	ForwardCompatible() bool
	// RequiredVersion retrieves required OpenGL version.
	RequiredVersion() (major int, minor int)
	// SharedContext retrieves the GdkGLContext that this context share data
	// with.
	SharedContext() GLContexter
	// Surface retrieves the surface used by the context.
	Surface() Surfacer
	// UseES checks whether the context is using an OpenGL or OpenGL ES profile.
	UseES() bool
	// Version retrieves the OpenGL version of the context.
	Version() (major int, minor int)
	// IsLegacy: whether the GdkGLContext is in legacy mode or not.
	IsLegacy() bool
	// MakeCurrent makes the context the current one.
	MakeCurrent()
	// Realize realizes the given GdkGLContext.
	Realize() error
	// SetDebugEnabled sets whether the GdkGLContext should perform extra
	// validations and runtime checking.
	SetDebugEnabled(enabled bool)
	// SetForwardCompatible sets whether the GdkGLContext should be
	// forward-compatible.
	SetForwardCompatible(compatible bool)
	// SetRequiredVersion sets the major and minor version of OpenGL to request.
	SetRequiredVersion(major int, minor int)
	// SetUseES requests that GDK create an OpenGL ES context instead of an
	// OpenGL one.
	SetUseES(useEs int)
}

var _ GLContexter = (*GLContext)(nil)

func wrapGLContext(obj *externglib.Object) *GLContext {
	return &GLContext{
		DrawContext: DrawContext{
			Object: obj,
		},
	}
}

func marshalGLContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGLContext(obj), nil
}

// DebugEnabled retrieves whether the context is doing extra validations and
// runtime checking.
//
// See gdk.GLContext.SetDebugEnabled().
func (context *GLContext) DebugEnabled() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_gl_context_get_debug_enabled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Display retrieves the display the context is created for
func (context *GLContext) Display() *Display {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_gl_context_get_display(_arg0)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// ForwardCompatible retrieves whether the context is forward-compatible.
//
// See gdk.GLContext.SetForwardCompatible().
func (context *GLContext) ForwardCompatible() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_gl_context_get_forward_compatible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RequiredVersion retrieves required OpenGL version.
//
// See gdk.GLContext.SetRequiredVersion().
func (context *GLContext) RequiredVersion() (major int, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // in
	var _arg2 C.int           // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	C.gdk_gl_context_get_required_version(_arg0, &_arg1, &_arg2)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

// SharedContext retrieves the GdkGLContext that this context share data with.
func (context *GLContext) SharedContext() GLContexter {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkGLContext // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_gl_context_get_shared_context(_arg0)

	var _glContext GLContexter // out

	_glContext = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(GLContexter)

	return _glContext
}

// Surface retrieves the surface used by the context.
func (context *GLContext) Surface() Surfacer {
	var _arg0 *C.GdkGLContext // out
	var _cret *C.GdkSurface   // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_gl_context_get_surface(_arg0)

	var _surface Surfacer // out

	_surface = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Surfacer)

	return _surface
}

// UseES checks whether the context is using an OpenGL or OpenGL ES profile.
func (context *GLContext) UseES() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_gl_context_get_use_es(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Version retrieves the OpenGL version of the context.
//
// The context must be realized prior to calling this function.
func (context *GLContext) Version() (major int, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // in
	var _arg2 C.int           // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	C.gdk_gl_context_get_version(_arg0, &_arg1, &_arg2)

	var _major int // out
	var _minor int // out

	_major = int(_arg1)
	_minor = int(_arg2)

	return _major, _minor
}

// IsLegacy: whether the GdkGLContext is in legacy mode or not.
//
// The GdkGLContext must be realized before calling this function.
//
// When realizing a GL context, GDK will try to use the OpenGL 3.2 core profile;
// this profile removes all the OpenGL API that was deprecated prior to the 3.2
// version of the specification. If the realization is successful, this function
// will return FALSE.
//
// If the underlying OpenGL implementation does not support core profiles, GDK
// will fall back to a pre-3.2 compatibility profile, and this function will
// return TRUE.
//
// You can use the value returned by this function to decide which kind of
// OpenGL API to use, or whether to do extension discovery, or what kind of
// shader programs to load.
func (context *GLContext) IsLegacy() bool {
	var _arg0 *C.GdkGLContext // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_gl_context_is_legacy(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MakeCurrent makes the context the current one.
func (context *GLContext) MakeCurrent() {
	var _arg0 *C.GdkGLContext // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	C.gdk_gl_context_make_current(_arg0)
}

// Realize realizes the given GdkGLContext.
//
// It is safe to call this function on a realized GdkGLContext.
func (context *GLContext) Realize() error {
	var _arg0 *C.GdkGLContext // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))

	C.gdk_gl_context_realize(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetDebugEnabled sets whether the GdkGLContext should perform extra
// validations and runtime checking.
//
// This is useful during development, but has additional overhead.
//
// The GdkGLContext must not be realized or made current prior to calling this
// function.
func (context *GLContext) SetDebugEnabled(enabled bool) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.gdk_gl_context_set_debug_enabled(_arg0, _arg1)
}

// SetForwardCompatible sets whether the GdkGLContext should be
// forward-compatible.
//
// Forward-compatible contexts must not support OpenGL functionality that has
// been marked as deprecated in the requested version; non-forward compatible
// contexts, on the other hand, must support both deprecated and non deprecated
// functionality.
//
// The GdkGLContext must not be realized or made current prior to calling this
// function.
func (context *GLContext) SetForwardCompatible(compatible bool) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))
	if compatible {
		_arg1 = C.TRUE
	}

	C.gdk_gl_context_set_forward_compatible(_arg0, _arg1)
}

// SetRequiredVersion sets the major and minor version of OpenGL to request.
//
// Setting major and minor to zero will use the default values.
//
// The GdkGLContext must not be realized or made current prior to calling this
// function.
func (context *GLContext) SetRequiredVersion(major int, minor int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.int(major)
	_arg2 = C.int(minor)

	C.gdk_gl_context_set_required_version(_arg0, _arg1, _arg2)
}

// SetUseES requests that GDK create an OpenGL ES context instead of an OpenGL
// one.
//
// Not all platforms support OpenGL ES.
//
// The context must not have been realized.
//
// By default, GDK will attempt to automatically detect whether the underlying
// GL implementation is OpenGL or OpenGL ES once the context is realized.
//
// You should check the return value of gdk.GLContext.GetUseES() after calling
// gdk.GLContext.Realize() to decide whether to use the OpenGL or OpenGL ES API,
// extensions, or shaders.
func (context *GLContext) SetUseES(useEs int) {
	var _arg0 *C.GdkGLContext // out
	var _arg1 C.int           // out

	_arg0 = (*C.GdkGLContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.int(useEs)

	C.gdk_gl_context_set_use_es(_arg0, _arg1)
}

// GLContextClearCurrent clears the current GdkGLContext.
//
// Any OpenGL call after this function returns will be ignored until
// gdk.GLContext.MakeCurrent() is called.
func GLContextClearCurrent() {
	C.gdk_gl_context_clear_current()
}

// GLContextGetCurrent retrieves the current GdkGLContext.
func GLContextGetCurrent() GLContexter {
	var _cret *C.GdkGLContext // in

	_cret = C.gdk_gl_context_get_current()

	var _glContext GLContexter // out

	_glContext = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(GLContexter)

	return _glContext
}
