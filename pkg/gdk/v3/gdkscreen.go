// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_screen_get_type()), F: marshalScreener},
	})
}

// Screen objects are the GDK representation of the screen on which windows can
// be displayed and on which the pointer moves. X originally identified screens
// with physical screens, but nowadays it is more common to have a single Screen
// which combines several physical monitors (see gdk_screen_get_n_monitors()).
//
// GdkScreen is used throughout GDK and GTK+ to specify which screen the top
// level windows are to be displayed on. it is also used to query the screen
// specification and default settings such as the default visual
// (gdk_screen_get_system_visual()), the dimensions of the physical monitors
// (gdk_screen_get_monitor_geometry()), etc.
type Screen struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Screen)(nil)

func wrapScreen(obj *externglib.Object) *Screen {
	return &Screen{
		Object: obj,
	}
}

func marshalScreener(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScreen(obj), nil
}

// ActiveWindow returns the screen’s currently active window.
//
// On X11, this is done by inspecting the _NET_ACTIVE_WINDOW property on the
// root window, as described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec). If there is no currently
// currently active window, or the window manager does not support the
// _NET_ACTIVE_WINDOW hint, this function returns NULL.
//
// On other platforms, this function may return NULL, depending on whether it is
// implementable on that platform.
//
// The returned window should be unrefed using g_object_unref() when no longer
// needed.
//
// Deprecated: since version 3.22.
func (screen *Screen) ActiveWindow() Windower {
	var _arg0 *C.GdkScreen // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_active_window(_arg0)

	var _window Windower // out

	_window = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Windower)

	return _window
}

// Display gets the display to which the screen belongs.
func (screen *Screen) Display() *Display {
	var _arg0 *C.GdkScreen  // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_display(_arg0)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// FontOptions gets any options previously set with
// gdk_screen_set_font_options().
func (screen *Screen) FontOptions() *cairo.FontOptions {
	var _arg0 *C.GdkScreen            // out
	var _cret *C.cairo_font_options_t // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_font_options(_arg0)

	var _fontOptions *cairo.FontOptions // out

	_fontOptions = (*cairo.FontOptions)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _fontOptions
}

// Height gets the height of screen in pixels. The returned size is in
// ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// Deprecated: Use per-monitor information instead.
func (screen *Screen) Height() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// HeightMm returns the height of screen in millimeters.
//
// Note that this value is somewhat ill-defined when the screen has multiple
// monitors of different resolution. It is recommended to use the monitor
// dimensions instead.
//
// Deprecated: Use per-monitor information instead.
func (screen *Screen) HeightMm() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_height_mm(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MonitorAtPoint returns the monitor number in which the point (x,y) is
// located.
//
// Deprecated: Use gdk_display_get_monitor_at_point() instead.
func (screen *Screen) MonitorAtPoint(x int, y int) int {
	var _arg0 *C.GdkScreen // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gdk_screen_get_monitor_at_point(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MonitorAtWindow returns the number of the monitor in which the largest area
// of the bounding rectangle of window resides.
//
// Deprecated: Use gdk_display_get_monitor_at_window() instead.
func (screen *Screen) MonitorAtWindow(window Windower) int {
	var _arg0 *C.GdkScreen // out
	var _arg1 *C.GdkWindow // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer((window).(gextras.Nativer).Native()))

	_cret = C.gdk_screen_get_monitor_at_window(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MonitorGeometry retrieves the Rectangle representing the size and position of
// the individual monitor within the entire screen area. The returned geometry
// is in ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// Monitor numbers start at 0. To obtain the number of monitors of screen, use
// gdk_screen_get_n_monitors().
//
// Note that the size of the entire screen area can be retrieved via
// gdk_screen_get_width() and gdk_screen_get_height().
//
// Deprecated: Use gdk_monitor_get_geometry() instead.
func (screen *Screen) MonitorGeometry(monitorNum int) Rectangle {
	var _arg0 *C.GdkScreen   // out
	var _arg1 C.gint         // out
	var _arg2 C.GdkRectangle // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = C.gint(monitorNum)

	C.gdk_screen_get_monitor_geometry(_arg0, _arg1, &_arg2)

	var _dest Rectangle // out

	_dest = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _dest
}

// MonitorHeightMm gets the height in millimeters of the specified monitor.
//
// Deprecated: Use gdk_monitor_get_height_mm() instead.
func (screen *Screen) MonitorHeightMm(monitorNum int) int {
	var _arg0 *C.GdkScreen // out
	var _arg1 C.gint       // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = C.gint(monitorNum)

	_cret = C.gdk_screen_get_monitor_height_mm(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MonitorPlugName returns the output name of the specified monitor. Usually
// something like VGA, DVI, or TV, not the actual product name of the display
// device.
//
// Deprecated: Use gdk_monitor_get_model() instead.
func (screen *Screen) MonitorPlugName(monitorNum int) string {
	var _arg0 *C.GdkScreen // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = C.gint(monitorNum)

	_cret = C.gdk_screen_get_monitor_plug_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// MonitorScaleFactor returns the internal scale factor that maps from monitor
// coordinates to the actual device pixels. On traditional systems this is 1,
// but on very high density outputs this can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a window where it is better
// to use gdk_window_get_scale_factor() instead.
//
// Deprecated: Use gdk_monitor_get_scale_factor() instead.
func (screen *Screen) MonitorScaleFactor(monitorNum int) int {
	var _arg0 *C.GdkScreen // out
	var _arg1 C.gint       // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = C.gint(monitorNum)

	_cret = C.gdk_screen_get_monitor_scale_factor(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MonitorWidthMm gets the width in millimeters of the specified monitor, if
// available.
//
// Deprecated: Use gdk_monitor_get_width_mm() instead.
func (screen *Screen) MonitorWidthMm(monitorNum int) int {
	var _arg0 *C.GdkScreen // out
	var _arg1 C.gint       // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = C.gint(monitorNum)

	_cret = C.gdk_screen_get_monitor_width_mm(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MonitorWorkarea retrieves the Rectangle representing the size and position of
// the “work area” on a monitor within the entire screen area. The returned
// geometry is in ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// The work area should be considered when positioning menus and similar popups,
// to avoid placing them below panels, docks or other desktop components.
//
// Note that not all backends may have a concept of workarea. This function will
// return the monitor geometry if a workarea is not available, or does not
// apply.
//
// Monitor numbers start at 0. To obtain the number of monitors of screen, use
// gdk_screen_get_n_monitors().
//
// Deprecated: Use gdk_monitor_get_workarea() instead.
func (screen *Screen) MonitorWorkarea(monitorNum int) Rectangle {
	var _arg0 *C.GdkScreen   // out
	var _arg1 C.gint         // out
	var _arg2 C.GdkRectangle // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = C.gint(monitorNum)

	C.gdk_screen_get_monitor_workarea(_arg0, _arg1, &_arg2)

	var _dest Rectangle // out

	_dest = *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _dest
}

// NMonitors returns the number of monitors which screen consists of.
//
// Deprecated: Use gdk_display_get_n_monitors() instead.
func (screen *Screen) NMonitors() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_n_monitors(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Number gets the index of screen among the screens in the display to which it
// belongs. (See gdk_screen_get_display())
//
// Deprecated: since version 3.22.
func (screen *Screen) Number() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_number(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PrimaryMonitor gets the primary monitor for screen. The primary monitor is
// considered the monitor where the “main desktop” lives. While normal
// application windows typically allow the window manager to place the windows,
// specialized desktop applications such as panels should place themselves on
// the primary monitor.
//
// If no primary monitor is configured by the user, the return value will be 0,
// defaulting to the first monitor.
//
// Deprecated: Use gdk_display_get_primary_monitor() instead.
func (screen *Screen) PrimaryMonitor() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_primary_monitor(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Resolution gets the resolution for font handling on the screen; see
// gdk_screen_set_resolution() for full details.
func (screen *Screen) Resolution() float64 {
	var _arg0 *C.GdkScreen // out
	var _cret C.gdouble    // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_resolution(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RGBAVisual gets a visual to use for creating windows with an alpha channel.
// The windowing system on which GTK+ is running may not support this
// capability, in which case NULL will be returned. Even if a non-NULL value is
// returned, its possible that the window’s alpha channel won’t be honored when
// displaying the window on the screen: in particular, for X an appropriate
// windowing manager and compositing manager must be running to provide
// appropriate display.
//
// This functionality is not implemented in the Windows backend.
//
// For setting an overall opacity for a top-level window, see
// gdk_window_set_opacity().
func (screen *Screen) RGBAVisual() *Visual {
	var _arg0 *C.GdkScreen // out
	var _cret *C.GdkVisual // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_rgba_visual(_arg0)

	var _visual *Visual // out

	_visual = wrapVisual(externglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// RootWindow gets the root window of screen.
func (screen *Screen) RootWindow() Windower {
	var _arg0 *C.GdkScreen // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_root_window(_arg0)

	var _window Windower // out

	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Windower)

	return _window
}

// Setting retrieves a desktop-wide setting such as double-click time for the
// Screen screen.
//
// FIXME needs a list of valid settings here, or a link to more information.
func (screen *Screen) Setting(name string, value *externglib.Value) bool {
	var _arg0 *C.GdkScreen // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.GValue    // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	_arg2 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gdk_screen_get_setting(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SystemVisual: get the system’s default visual for screen. This is the visual
// for the root window of the display. The return value should not be freed.
func (screen *Screen) SystemVisual() *Visual {
	var _arg0 *C.GdkScreen // out
	var _cret *C.GdkVisual // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_system_visual(_arg0)

	var _visual *Visual // out

	_visual = wrapVisual(externglib.Take(unsafe.Pointer(_cret)))

	return _visual
}

// ToplevelWindows obtains a list of all toplevel windows known to GDK on the
// screen screen. A toplevel window is a child of the root window (see
// gdk_get_default_root_window()).
//
// The returned list should be freed with g_list_free(), but its elements need
// not be freed.
func (screen *Screen) ToplevelWindows() *externglib.List {
	var _arg0 *C.GdkScreen // out
	var _cret *C.GList     // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_toplevel_windows(_arg0)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.GdkWindow)(_p)
		var dst Windower // out
		dst = (gextras.CastObject(externglib.Take(unsafe.Pointer(src)))).(Windower)
		return dst
	})
	_list.AttachFinalizer(nil)

	return _list
}

// Width gets the width of screen in pixels. The returned size is in
// ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// Deprecated: Use per-monitor information instead.
func (screen *Screen) Width() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WidthMm gets the width of screen in millimeters.
//
// Note that this value is somewhat ill-defined when the screen has multiple
// monitors of different resolution. It is recommended to use the monitor
// dimensions instead.
//
// Deprecated: Use per-monitor information instead.
func (screen *Screen) WidthMm() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_width_mm(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WindowStack returns a #GList of Windows representing the current window
// stack.
//
// On X11, this is done by inspecting the _NET_CLIENT_LIST_STACKING property on
// the root window, as described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec). If the window manager does
// not support the _NET_CLIENT_LIST_STACKING hint, this function returns NULL.
//
// On other platforms, this function may return NULL, depending on whether it is
// implementable on that platform.
//
// The returned list is newly allocated and owns references to the windows it
// contains, so it should be freed using g_list_free() and its windows unrefed
// using g_object_unref() when no longer needed.
func (screen *Screen) WindowStack() *externglib.List {
	var _arg0 *C.GdkScreen // out
	var _cret *C.GList     // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_get_window_stack(_arg0)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.GdkWindow)(_p)
		var dst Windower // out
		dst = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(Windower)
		return dst
	})
	_list.AttachFinalizer(func(v uintptr) {
		C.g_object_unref(C.gpointer(uintptr(unsafe.Pointer(v))))
	})

	return _list
}

// IsComposited returns whether windows with an RGBA visual can reasonably be
// expected to have their alpha channel drawn correctly on the screen.
//
// On X11 this function returns whether a compositing manager is compositing
// screen.
func (screen *Screen) IsComposited() bool {
	var _arg0 *C.GdkScreen // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_is_composited(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListVisuals lists the available visuals for the specified screen. A visual
// describes a hardware image data format. For example, a visual might support
// 24-bit color, or 8-bit color, and might expect pixels to be in a certain
// format.
//
// Call g_list_free() on the return value when you’re finished with it.
func (screen *Screen) ListVisuals() *externglib.List {
	var _arg0 *C.GdkScreen // out
	var _cret *C.GList     // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_list_visuals(_arg0)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.GdkVisual)(_p)
		var dst Visual // out
		dst = *wrapVisual(externglib.Take(unsafe.Pointer(src)))
		return dst
	})
	_list.AttachFinalizer(nil)

	return _list
}

// MakeDisplayName determines the name to pass to gdk_display_open() to get a
// Display with this screen as the default screen.
//
// Deprecated: since version 3.22.
func (screen *Screen) MakeDisplayName() string {
	var _arg0 *C.GdkScreen // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_screen_make_display_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetFontOptions sets the default font options for the screen. These options
// will be set on any Context’s newly created with
// gdk_pango_context_get_for_screen(). Changing the default set of font options
// does not affect contexts that have already been created.
func (screen *Screen) SetFontOptions(options *cairo.FontOptions) {
	var _arg0 *C.GdkScreen            // out
	var _arg1 *C.cairo_font_options_t // out

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = (*C.cairo_font_options_t)(gextras.StructNative(unsafe.Pointer(options)))

	C.gdk_screen_set_font_options(_arg0, _arg1)
}

// SetResolution sets the resolution for font handling on the screen. This is a
// scale factor between points specified in a FontDescription and cairo units.
// The default value is 96, meaning that a 10 point font will be 13 units high.
// (10 * 96. / 72. = 13.3).
func (screen *Screen) SetResolution(dpi float64) {
	var _arg0 *C.GdkScreen // out
	var _arg1 C.gdouble    // out

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg1 = C.gdouble(dpi)

	C.gdk_screen_set_resolution(_arg0, _arg1)
}

// ScreenGetDefault gets the default screen for the default display. (See
// gdk_display_get_default ()).
func ScreenGetDefault() *Screen {
	var _cret *C.GdkScreen // in

	_cret = C.gdk_screen_get_default()

	var _screen *Screen // out

	_screen = wrapScreen(externglib.Take(unsafe.Pointer(_cret)))

	return _screen
}
