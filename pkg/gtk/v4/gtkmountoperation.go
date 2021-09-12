// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperationer},
	})
}

// MountOperation: GtkMountOperation is an implementation of GMountOperation.
//
// The functions and objects described here make working with GTK and GIO more
// convenient.
//
// GtkMountOperation is needed when mounting volumes: It is an implementation of
// GMountOperation that can be used with GIO functions for mounting volumes such
// as g_file_mount_enclosing_volume(), g_file_mount_mountable(),
// g_volume_mount(), g_mount_unmount_with_operation() and others.
//
// When necessary, GtkMountOperation shows dialogs to let the user enter
// passwords, ask questions or show processes blocking unmount.
type MountOperation struct {
	gio.MountOperation
}

func wrapMountOperation(obj *externglib.Object) *MountOperation {
	return &MountOperation{
		MountOperation: gio.MountOperation{
			Object: obj,
		},
	}
}

func marshalMountOperationer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMountOperation(obj), nil
}

// NewMountOperation creates a new GtkMountOperation.
func NewMountOperation(parent *Window) *MountOperation {
	var _arg1 *C.GtkWindow       // out
	var _cret *C.GMountOperation // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}

	_cret = C.gtk_mount_operation_new(_arg1)
	runtime.KeepAlive(parent)

	var _mountOperation *MountOperation // out

	_mountOperation = wrapMountOperation(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mountOperation
}

// Display gets the display on which windows of the GtkMountOperation will be
// shown.
func (op *MountOperation) Display() *gdk.Display {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_mount_operation_get_display(_arg0)
	runtime.KeepAlive(op)

	var _display *gdk.Display // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// Parent gets the transient parent used by the GtkMountOperation.
func (op *MountOperation) Parent() *Window {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GtkWindow         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_mount_operation_get_parent(_arg0)
	runtime.KeepAlive(op)

	var _window *Window // out

	_window = wrapWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _window
}

// IsShowing returns whether the GtkMountOperation is currently displaying a
// window.
func (op *MountOperation) IsShowing() bool {
	var _arg0 *C.GtkMountOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_mount_operation_is_showing(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDisplay sets the display to show windows of the GtkMountOperation on.
func (op *MountOperation) SetDisplay(display *gdk.Display) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GdkDisplay        // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gtk_mount_operation_set_display(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(display)
}

// SetParent sets the transient parent for windows shown by the
// GtkMountOperation.
func (op *MountOperation) SetParent(parent *Window) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GtkWindow         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer(op.Native()))
	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}

	C.gtk_mount_operation_set_parent(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(parent)
}
