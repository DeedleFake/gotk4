// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gdk.h"
import "C"

// TestRenderSync retrieves a pixel from window to force the windowing system to
// carry out any pending rendering commands.
//
// This function is intended to be used to synchronize with rendering pipelines,
// to benchmark windowing system rendering operations.
func TestRenderSync(window Windower) {
	var _arg1 *C.GdkWindow // out

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	C.gdk_test_render_sync(_arg1)
	runtime.KeepAlive(window)
}

// TestSimulateButton: this function is intended to be used in GTK+ test
// programs. It will warp the mouse pointer to the given (x,y) coordinates
// within window and simulate a button press or release event. Because the mouse
// pointer needs to be warped to the target location, use of this function
// outside of test programs that run in their own virtual windowing system (e.g.
// Xvfb) is not recommended.
//
// Also, gdk_test_simulate_button() is a fairly low level function, for most
// testing purposes, gtk_test_widget_click() is the right function to call which
// will generate a button press event followed by its accompanying button
// release event.
func TestSimulateButton(window Windower, x int, y int, button uint, modifiers ModifierType, buttonPressrelease EventType) bool {
	var _arg1 *C.GdkWindow      // out
	var _arg2 C.gint            // out
	var _arg3 C.gint            // out
	var _arg4 C.guint           // out
	var _arg5 C.GdkModifierType // out
	var _arg6 C.GdkEventType    // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)
	_arg4 = C.guint(button)
	_arg5 = C.GdkModifierType(modifiers)
	_arg6 = C.GdkEventType(buttonPressrelease)

	_cret = C.gdk_test_simulate_button(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(window)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(button)
	runtime.KeepAlive(modifiers)
	runtime.KeepAlive(buttonPressrelease)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TestSimulateKey: this function is intended to be used in GTK+ test programs.
// If (x,y) are > (-1,-1), it will warp the mouse pointer to the given (x,y)
// coordinates within window and simulate a key press or release event.
//
// When the mouse pointer is warped to the target location, use of this function
// outside of test programs that run in their own virtual windowing system (e.g.
// Xvfb) is not recommended. If (x,y) are passed as (-1,-1), the mouse pointer
// will not be warped and window origin will be used as mouse pointer location
// for the event.
//
// Also, gdk_test_simulate_key() is a fairly low level function, for most
// testing purposes, gtk_test_widget_send_key() is the right function to call
// which will generate a key press event followed by its accompanying key
// release event.
func TestSimulateKey(window Windower, x int, y int, keyval uint, modifiers ModifierType, keyPressrelease EventType) bool {
	var _arg1 *C.GdkWindow      // out
	var _arg2 C.gint            // out
	var _arg3 C.gint            // out
	var _arg4 C.guint           // out
	var _arg5 C.GdkModifierType // out
	var _arg6 C.GdkEventType    // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)
	_arg4 = C.guint(keyval)
	_arg5 = C.GdkModifierType(modifiers)
	_arg6 = C.GdkEventType(keyPressrelease)

	_cret = C.gdk_test_simulate_key(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(window)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)
	runtime.KeepAlive(keyPressrelease)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
