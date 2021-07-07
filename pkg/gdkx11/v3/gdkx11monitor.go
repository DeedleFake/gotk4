// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_monitor_get_type()), F: marshalX11Monitor},
	})
}

type X11Monitor interface {
	gdk.Monitor

	// AsMonitor casts the class to the gdk.Monitor interface.
	AsMonitor() gdk.Monitor

	// GetDisplay gets the display that this monitor belongs to.
	//
	// This method is inherited from gdk.Monitor
	GetDisplay() gdk.Display
	// GetGeometry retrieves the size and position of an individual monitor
	// within the display coordinate space. The returned geometry is in
	// ”application pixels”, not in ”device pixels” (see
	// gdk_monitor_get_scale_factor()).
	//
	// This method is inherited from gdk.Monitor
	GetGeometry() gdk.Rectangle
	// GetHeightMm gets the height in millimeters of the monitor.
	//
	// This method is inherited from gdk.Monitor
	GetHeightMm() int
	// GetManufacturer gets the name or PNP ID of the monitor's manufacturer, if
	// available.
	//
	// Note that this value might also vary depending on actual display backend.
	//
	// PNP ID registry is located at https://uefi.org/pnp_id_list
	//
	// This method is inherited from gdk.Monitor
	GetManufacturer() string
	// GetModel gets the a string identifying the monitor model, if available.
	//
	// This method is inherited from gdk.Monitor
	GetModel() string
	// GetRefreshRate gets the refresh rate of the monitor, if available.
	//
	// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as
	// 60000.
	//
	// This method is inherited from gdk.Monitor
	GetRefreshRate() int
	// GetScaleFactor gets the internal scale factor that maps from monitor
	// coordinates to the actual device pixels. On traditional systems this is
	// 1, but on very high density outputs this can be a higher value (often 2).
	//
	// This can be used if you want to create pixel based data for a particular
	// monitor, but most of the time you’re drawing to a window where it is
	// better to use gdk_window_get_scale_factor() instead.
	//
	// This method is inherited from gdk.Monitor
	GetScaleFactor() int
	// GetSubpixelLayout gets information about the layout of red, green and
	// blue primaries for each pixel in this monitor, if available.
	//
	// This method is inherited from gdk.Monitor
	GetSubpixelLayout() gdk.SubpixelLayout
	// GetWidthMm gets the width in millimeters of the monitor.
	//
	// This method is inherited from gdk.Monitor
	GetWidthMm() int
	// GetWorkarea retrieves the size and position of the “work area” on a
	// monitor within the display coordinate space. The returned geometry is in
	// ”application pixels”, not in ”device pixels” (see
	// gdk_monitor_get_scale_factor()).
	//
	// The work area should be considered when positioning menus and similar
	// popups, to avoid placing them below panels, docks or other desktop
	// components.
	//
	// Note that not all backends may have a concept of workarea. This function
	// will return the monitor geometry if a workarea is not available, or does
	// not apply.
	//
	// This method is inherited from gdk.Monitor
	GetWorkarea() gdk.Rectangle
	// IsPrimary gets whether this monitor should be considered primary (see
	// gdk_display_get_primary_monitor()).
	//
	// This method is inherited from gdk.Monitor
	IsPrimary() bool
}

// x11Monitor implements the X11Monitor interface.
type x11Monitor struct {
	*externglib.Object
}

var _ X11Monitor = (*x11Monitor)(nil)

// WrapX11Monitor wraps a GObject to a type that implements
// interface X11Monitor. It is primarily used internally.
func WrapX11Monitor(obj *externglib.Object) X11Monitor {
	return x11Monitor{obj}
}

func marshalX11Monitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Monitor(obj), nil
}

func (x x11Monitor) AsMonitor() gdk.Monitor {
	return gdk.WrapMonitor(gextras.InternObject(x))
}

func (m x11Monitor) GetDisplay() gdk.Display {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetDisplay()
}

func (m x11Monitor) GetGeometry() gdk.Rectangle {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetGeometry()
}

func (m x11Monitor) GetHeightMm() int {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetHeightMm()
}

func (m x11Monitor) GetManufacturer() string {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetManufacturer()
}

func (m x11Monitor) GetModel() string {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetModel()
}

func (m x11Monitor) GetRefreshRate() int {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetRefreshRate()
}

func (m x11Monitor) GetScaleFactor() int {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetScaleFactor()
}

func (m x11Monitor) GetSubpixelLayout() gdk.SubpixelLayout {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetSubpixelLayout()
}

func (m x11Monitor) GetWidthMm() int {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetWidthMm()
}

func (m x11Monitor) GetWorkarea() gdk.Rectangle {
	return gdk.WrapMonitor(gextras.InternObject(m)).GetWorkarea()
}

func (m x11Monitor) IsPrimary() bool {
	return gdk.WrapMonitor(gextras.InternObject(m)).IsPrimary()
}
