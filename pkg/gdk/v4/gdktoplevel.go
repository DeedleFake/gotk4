// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"strings"
	"unsafe"

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
		{T: externglib.Type(C.gdk_fullscreen_mode_get_type()), F: marshalFullscreenMode},
		{T: externglib.Type(C.gdk_surface_edge_get_type()), F: marshalSurfaceEdge},
		{T: externglib.Type(C.gdk_toplevel_state_get_type()), F: marshalToplevelState},
		{T: externglib.Type(C.gdk_toplevel_get_type()), F: marshalTopleveller},
	})
}

// FullscreenMode indicates which monitor a surface should span over when in
// fullscreen mode.
type FullscreenMode int

const (
	// FullscreenOnCurrentMonitor: fullscreen on current monitor only.
	FullscreenOnCurrentMonitor FullscreenMode = iota
	// FullscreenOnAllMonitors: span across all monitors when fullscreen.
	FullscreenOnAllMonitors
)

func marshalFullscreenMode(p uintptr) (interface{}, error) {
	return FullscreenMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for FullscreenMode.
func (f FullscreenMode) String() string {
	switch f {
	case FullscreenOnCurrentMonitor:
		return "CurrentMonitor"
	case FullscreenOnAllMonitors:
		return "AllMonitors"
	default:
		return fmt.Sprintf("FullscreenMode(%d)", f)
	}
}

// SurfaceEdge determines a surface edge or corner.
type SurfaceEdge int

const (
	// SurfaceEdgeNorthWest: top left corner.
	SurfaceEdgeNorthWest SurfaceEdge = iota
	// SurfaceEdgeNorth: top edge.
	SurfaceEdgeNorth
	// SurfaceEdgeNorthEast: top right corner.
	SurfaceEdgeNorthEast
	// SurfaceEdgeWest: left edge.
	SurfaceEdgeWest
	// SurfaceEdgeEast: right edge.
	SurfaceEdgeEast
	// SurfaceEdgeSouthWest: lower left corner.
	SurfaceEdgeSouthWest
	// SurfaceEdgeSouth: lower edge.
	SurfaceEdgeSouth
	// SurfaceEdgeSouthEast: lower right corner.
	SurfaceEdgeSouthEast
)

func marshalSurfaceEdge(p uintptr) (interface{}, error) {
	return SurfaceEdge(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for SurfaceEdge.
func (s SurfaceEdge) String() string {
	switch s {
	case SurfaceEdgeNorthWest:
		return "NorthWest"
	case SurfaceEdgeNorth:
		return "North"
	case SurfaceEdgeNorthEast:
		return "NorthEast"
	case SurfaceEdgeWest:
		return "West"
	case SurfaceEdgeEast:
		return "East"
	case SurfaceEdgeSouthWest:
		return "SouthWest"
	case SurfaceEdgeSouth:
		return "South"
	case SurfaceEdgeSouthEast:
		return "SouthEast"
	default:
		return fmt.Sprintf("SurfaceEdge(%d)", s)
	}
}

// ToplevelState specifies the state of a toplevel surface.
//
// On platforms that support information about individual edges, the
// GDK_TOPLEVEL_STATE_TILED state will be set whenever any of the individual
// tiled states is set. On platforms that lack that support, the tiled state
// will give an indication of tiledness without any of the per-edge states being
// set.
type ToplevelState int

const (
	// ToplevelStateMinimized: surface is minimized
	ToplevelStateMinimized ToplevelState = 0b1
	// ToplevelStateMaximized: surface is maximized
	ToplevelStateMaximized ToplevelState = 0b10
	// ToplevelStateSticky: surface is sticky
	ToplevelStateSticky ToplevelState = 0b100
	// ToplevelStateFullscreen: surface is maximized without decorations
	ToplevelStateFullscreen ToplevelState = 0b1000
	// ToplevelStateAbove: surface is kept above other surfaces
	ToplevelStateAbove ToplevelState = 0b10000
	// ToplevelStateBelow: surface is kept below other surfaces
	ToplevelStateBelow ToplevelState = 0b100000
	// ToplevelStateFocused: surface is presented as focused (with active
	// decorations)
	ToplevelStateFocused ToplevelState = 0b1000000
	// ToplevelStateTiled: surface is in a tiled state
	ToplevelStateTiled ToplevelState = 0b10000000
	// ToplevelStateTopTiled: whether the top edge is tiled
	ToplevelStateTopTiled ToplevelState = 0b100000000
	// ToplevelStateTopResizable: whether the top edge is resizable
	ToplevelStateTopResizable ToplevelState = 0b1000000000
	// ToplevelStateRightTiled: whether the right edge is tiled
	ToplevelStateRightTiled ToplevelState = 0b10000000000
	// ToplevelStateRightResizable: whether the right edge is resizable
	ToplevelStateRightResizable ToplevelState = 0b100000000000
	// ToplevelStateBottomTiled: whether the bottom edge is tiled
	ToplevelStateBottomTiled ToplevelState = 0b1000000000000
	// ToplevelStateBottomResizable: whether the bottom edge is resizable
	ToplevelStateBottomResizable ToplevelState = 0b10000000000000
	// ToplevelStateLeftTiled: whether the left edge is tiled
	ToplevelStateLeftTiled ToplevelState = 0b100000000000000
	// ToplevelStateLeftResizable: whether the left edge is resizable
	ToplevelStateLeftResizable ToplevelState = 0b1000000000000000
)

func marshalToplevelState(p uintptr) (interface{}, error) {
	return ToplevelState(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for ToplevelState.
func (t ToplevelState) String() string {
	if t == 0 {
		return "ToplevelState(0)"
	}

	var builder strings.Builder
	builder.Grow(371)

	for t != 0 {
		next := t & (t - 1)
		bit := t - next

		switch bit {
		case ToplevelStateMinimized:
			builder.WriteString("Minimized|")
		case ToplevelStateMaximized:
			builder.WriteString("Maximized|")
		case ToplevelStateSticky:
			builder.WriteString("Sticky|")
		case ToplevelStateFullscreen:
			builder.WriteString("Fullscreen|")
		case ToplevelStateAbove:
			builder.WriteString("Above|")
		case ToplevelStateBelow:
			builder.WriteString("Below|")
		case ToplevelStateFocused:
			builder.WriteString("Focused|")
		case ToplevelStateTiled:
			builder.WriteString("Tiled|")
		case ToplevelStateTopTiled:
			builder.WriteString("TopTiled|")
		case ToplevelStateTopResizable:
			builder.WriteString("TopResizable|")
		case ToplevelStateRightTiled:
			builder.WriteString("RightTiled|")
		case ToplevelStateRightResizable:
			builder.WriteString("RightResizable|")
		case ToplevelStateBottomTiled:
			builder.WriteString("BottomTiled|")
		case ToplevelStateBottomResizable:
			builder.WriteString("BottomResizable|")
		case ToplevelStateLeftTiled:
			builder.WriteString("LeftTiled|")
		case ToplevelStateLeftResizable:
			builder.WriteString("LeftResizable|")
		default:
			builder.WriteString(fmt.Sprintf("ToplevelState(0b%b)|", bit))
		}

		t = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Toplevel: GdkToplevel is a freestanding toplevel surface.
//
// The GdkToplevel interface provides useful APIs for interacting with the
// windowing system, such as controlling maximization and size of the surface,
// setting icons and transient parents for dialogs.
type Toplevel struct {
	Surface
}

var _ gextras.Nativer = (*Toplevel)(nil)

// Topleveller describes Toplevel's abstract methods.
type Topleveller interface {
	// BeginMove begins an interactive move operation.
	BeginMove(device Devicer, button int, x float64, y float64, timestamp uint32)
	// BeginResize begins an interactive resize operation.
	BeginResize(edge SurfaceEdge, device Devicer, button int, x float64, y float64, timestamp uint32)
	// Focus sets keyboard focus to surface.
	Focus(timestamp uint32)
	// State gets the bitwise or of the currently active surface state flags,
	// from the GdkToplevelState enumeration.
	State() ToplevelState
	// InhibitSystemShortcuts requests that the toplevel inhibit the system
	// shortcuts.
	InhibitSystemShortcuts(event Eventer)
	// Lower asks to lower the toplevel below other windows.
	Lower() bool
	// Minimize asks to minimize the toplevel.
	Minimize() bool
	// Present toplevel after having processed the GdkToplevelLayout rules.
	Present(layout *ToplevelLayout)
	// RestoreSystemShortcuts: restore default system keyboard shortcuts which
	// were previously inhibited.
	RestoreSystemShortcuts()
	// SetDecorated sets the toplevel to be decorated.
	SetDecorated(decorated bool)
	// SetDeletable sets the toplevel to be deletable.
	SetDeletable(deletable bool)
	// SetIconList sets a list of icons for the surface.
	SetIconList(surfaces []Texturer)
	// SetModal sets the toplevel to be modal.
	SetModal(modal bool)
	// SetStartupID sets the startup notification ID.
	SetStartupID(startupId string)
	// SetTitle sets the title of a toplevel surface.
	SetTitle(title string)
	// SetTransientFor sets a transient-for parent.
	SetTransientFor(parent Surfacer)
	// ShowWindowMenu asks the windowing system to show the window menu.
	ShowWindowMenu(event Eventer) bool
	// SupportsEdgeConstraints returns whether the desktop environment supports
	// tiled window states.
	SupportsEdgeConstraints() bool
}

var _ Topleveller = (*Toplevel)(nil)

func wrapToplevel(obj *externglib.Object) *Toplevel {
	return &Toplevel{
		Surface: Surface{
			Object: obj,
		},
	}
}

func marshalTopleveller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToplevel(obj), nil
}

// BeginMove begins an interactive move operation.
//
// You might use this function to implement draggable titlebars.
func (toplevel *Toplevel) BeginMove(device Devicer, button int, x float64, y float64, timestamp uint32) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GdkDevice   // out
	var _arg2 C.int          // out
	var _arg3 C.double       // out
	var _arg4 C.double       // out
	var _arg5 C.guint32      // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer((device).(gextras.Nativer).Native()))
	_arg2 = C.int(button)
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.guint32(timestamp)

	C.gdk_toplevel_begin_move(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// BeginResize begins an interactive resize operation.
//
// You might use this function to implement a “window resize grip.”
func (toplevel *Toplevel) BeginResize(edge SurfaceEdge, device Devicer, button int, x float64, y float64, timestamp uint32) {
	var _arg0 *C.GdkToplevel   // out
	var _arg1 C.GdkSurfaceEdge // out
	var _arg2 *C.GdkDevice     // out
	var _arg3 C.int            // out
	var _arg4 C.double         // out
	var _arg5 C.double         // out
	var _arg6 C.guint32        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = C.GdkSurfaceEdge(edge)
	_arg2 = (*C.GdkDevice)(unsafe.Pointer((device).(gextras.Nativer).Native()))
	_arg3 = C.int(button)
	_arg4 = C.double(x)
	_arg5 = C.double(y)
	_arg6 = C.guint32(timestamp)

	C.gdk_toplevel_begin_resize(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// Focus sets keyboard focus to surface.
//
// In most cases, gtk.Window.PresentWithTime() should be used on a gtk.Window,
// rather than calling this function.
func (toplevel *Toplevel) Focus(timestamp uint32) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 C.guint32      // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = C.guint32(timestamp)

	C.gdk_toplevel_focus(_arg0, _arg1)
}

// State gets the bitwise or of the currently active surface state flags, from
// the GdkToplevelState enumeration.
func (toplevel *Toplevel) State() ToplevelState {
	var _arg0 *C.GdkToplevel     // out
	var _cret C.GdkToplevelState // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))

	_cret = C.gdk_toplevel_get_state(_arg0)

	var _toplevelState ToplevelState // out

	_toplevelState = ToplevelState(_cret)

	return _toplevelState
}

// InhibitSystemShortcuts requests that the toplevel inhibit the system
// shortcuts.
//
// This is asking the desktop environment/windowing system to let all keyboard
// events reach the surface, as long as it is focused, instead of triggering
// system actions.
//
// If granted, the rerouting remains active until the default shortcuts
// processing is restored with gdk.Toplevel.RestoreSystemShortcuts(), or the
// request is revoked by the desktop environment, windowing system or the user.
//
// A typical use case for this API is remote desktop or virtual machine viewers
// which need to inhibit the default system keyboard shortcuts so that the
// remote session or virtual host gets those instead of the local environment.
//
// The windowing system or desktop environment may ask the user to grant or deny
// the request or even choose to ignore the request entirely.
//
// The caller can be notified whenever the request is granted or revoked by
// listening to the gdk.Toplevel:shortcuts-inhibited property.
func (toplevel *Toplevel) InhibitSystemShortcuts(event Eventer) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GdkEvent    // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer((event).(gextras.Nativer).Native()))

	C.gdk_toplevel_inhibit_system_shortcuts(_arg0, _arg1)
}

// Lower asks to lower the toplevel below other windows.
//
// The windowing system may choose to ignore the request.
func (toplevel *Toplevel) Lower() bool {
	var _arg0 *C.GdkToplevel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))

	_cret = C.gdk_toplevel_lower(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Minimize asks to minimize the toplevel.
//
// The windowing system may choose to ignore the request.
func (toplevel *Toplevel) Minimize() bool {
	var _arg0 *C.GdkToplevel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))

	_cret = C.gdk_toplevel_minimize(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Present toplevel after having processed the GdkToplevelLayout rules.
//
// If the toplevel was previously not showing, it will be showed, otherwise it
// will change layout according to layout.
//
// GDK may emit the gdk.Toplevel::compute-size signal to let the user of this
// toplevel compute the preferred size of the toplevel surface.
//
// Presenting is asynchronous and the specified layout parameters are not
// guaranteed to be respected.
func (toplevel *Toplevel) Present(layout *ToplevelLayout) {
	var _arg0 *C.GdkToplevel       // out
	var _arg1 *C.GdkToplevelLayout // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.GdkToplevelLayout)(gextras.StructNative(unsafe.Pointer(layout)))

	C.gdk_toplevel_present(_arg0, _arg1)
}

// RestoreSystemShortcuts: restore default system keyboard shortcuts which were
// previously inhibited.
//
// This undoes the effect of gdk.Toplevel.InhibitSystemShortcuts().
func (toplevel *Toplevel) RestoreSystemShortcuts() {
	var _arg0 *C.GdkToplevel // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))

	C.gdk_toplevel_restore_system_shortcuts(_arg0)
}

// SetDecorated sets the toplevel to be decorated.
//
// Setting decorated to FALSE hints the desktop environment that the surface has
// its own, client-side decorations and does not need to have window decorations
// added.
func (toplevel *Toplevel) SetDecorated(decorated bool) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	if decorated {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_set_decorated(_arg0, _arg1)
}

// SetDeletable sets the toplevel to be deletable.
//
// Setting deletable to TRUE hints the desktop environment that it should offer
// the user a way to close the surface.
func (toplevel *Toplevel) SetDeletable(deletable bool) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	if deletable {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_set_deletable(_arg0, _arg1)
}

// SetIconList sets a list of icons for the surface.
//
// One of these will be used to represent the surface in iconic form. The icon
// may be shown in window lists or task bars. Which icon size is shown depends
// on the window manager. The window manager can scale the icon but setting
// several size icons can give better image quality.
//
// Note that some platforms don't support surface icons.
func (toplevel *Toplevel) SetIconList(surfaces []Texturer) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GList       // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	for i := len(surfaces) - 1; i >= 0; i-- {
		src := surfaces[i]
		var dst *C.GdkTexture // out
		dst = (*C.GdkTexture)(unsafe.Pointer((src).(gextras.Nativer).Native()))
		_arg1 = C.g_list_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_list_free(_arg1)

	C.gdk_toplevel_set_icon_list(_arg0, _arg1)
}

// SetModal sets the toplevel to be modal.
//
// The application can use this hint to tell the window manager that a certain
// surface has modal behaviour. The window manager can use this information to
// handle modal surfaces in a special way.
//
// You should only use this on surfaces for which you have previously called
// gdk.Toplevel.SetTransientFor().
func (toplevel *Toplevel) SetModal(modal bool) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_set_modal(_arg0, _arg1)
}

// SetStartupID sets the startup notification ID.
//
// When using GTK, typically you should use gtk.Window.SetStartupID() instead of
// this low-level function.
func (toplevel *Toplevel) SetStartupID(startupId string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_toplevel_set_startup_id(_arg0, _arg1)
}

// SetTitle sets the title of a toplevel surface.
//
// The title maybe be displayed in the titlebar, in lists of windows, etc.
func (toplevel *Toplevel) SetTitle(title string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_toplevel_set_title(_arg0, _arg1)
}

// SetTransientFor sets a transient-for parent.
//
// Indicates to the window manager that surface is a transient dialog associated
// with the application surface parent. This allows the window manager to do
// things like center surface on parent and keep surface above parent.
//
// See gtk.Window.SetTransientFor() if you’re using gtk.Window or gtk.Dialog.
func (toplevel *Toplevel) SetTransientFor(parent Surfacer) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GdkSurface  // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer((parent).(gextras.Nativer).Native()))

	C.gdk_toplevel_set_transient_for(_arg0, _arg1)
}

// ShowWindowMenu asks the windowing system to show the window menu.
//
// The window menu is the menu shown when right-clicking the titlebar on
// traditional windows managed by the window manager. This is useful for windows
// using client-side decorations, activating it with a right-click on the window
// decorations.
func (toplevel *Toplevel) ShowWindowMenu(event Eventer) bool {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GdkEvent    // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer((event).(gextras.Nativer).Native()))

	_cret = C.gdk_toplevel_show_window_menu(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsEdgeConstraints returns whether the desktop environment supports
// tiled window states.
func (toplevel *Toplevel) SupportsEdgeConstraints() bool {
	var _arg0 *C.GdkToplevel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(toplevel.Native()))

	_cret = C.gdk_toplevel_supports_edge_constraints(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
