// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_pollfd_get_type()), F: marshalPollFD},
	})
}

// Poll polls @fds, as with the poll() system call, but portably. (On systems
// that don't have poll(), it is emulated using select().) This is used
// internally by Context, but it can be called directly if you need to block
// until a file descriptor is ready, but don't want to run the full main loop.
//
// Each element of @fds is a FD describing a single file descriptor to poll. The
// @fd field indicates the file descriptor, and the @events field indicates the
// events to poll for. On return, the @revents fields will be filled with the
// events that actually occurred.
//
// On POSIX systems, the file descriptors in @fds can be any sort of file
// descriptor, but the situation is much more complicated on Windows. If you
// need to use g_poll() in code that has to run on Windows, the easiest solution
// is to construct all of your FDs with g_io_channel_win32_make_pollfd().
func Poll(fds *PollFD, nfds uint, timeout int) int {
	var _arg1 *C.GPollFD // out
	var _arg2 C.guint    // out
	var _arg3 C.gint     // out
	var _cret C.gint     // in

	_arg1 = (*C.GPollFD)(unsafe.Pointer(fds))
	_arg2 = C.guint(nfds)
	_arg3 = C.gint(timeout)

	_cret = C.g_poll(_arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PollFD represents a file descriptor, which events to poll for, and which
// events occurred.
type PollFD struct {
	native C.GPollFD
}

// WrapPollFD wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPollFD(ptr unsafe.Pointer) *PollFD {
	return (*PollFD)(ptr)
}

func marshalPollFD(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*PollFD)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PollFD) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Fd: the file descriptor to poll (or a HANDLE on Win32)
func (p *PollFD) Fd() int {
	var v int // out
	v = int(p.native.fd)
	return v
}

// Events: bitwise combination from OCondition, specifying which events should
// be polled for. Typically for reading from a file descriptor you would use
// G_IO_IN | G_IO_HUP | G_IO_ERR, and for writing you would use G_IO_OUT |
// G_IO_ERR.
func (p *PollFD) Events() uint16 {
	var v uint16 // out
	v = uint16(p.native.events)
	return v
}

// Revents: bitwise combination of flags from OCondition, returned from the
// poll() function to indicate which events occurred.
func (p *PollFD) Revents() uint16 {
	var v uint16 // out
	v = uint16(p.native.revents)
	return v
}
