// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperation},
	})
}

// MountOperation: `GtkMountOperation` is an implementation of
// `GMountOperation`.
//
// The functions and objects described here make working with GTK and GIO more
// convenient.
//
// `GtkMountOperation` is needed when mounting volumes: It is an implementation
// of `GMountOperation` that can be used with GIO functions for mounting volumes
// such as g_file_mount_enclosing_volume(), g_file_mount_mountable(),
// g_volume_mount(), g_mount_unmount_with_operation() and others.
//
// When necessary, `GtkMountOperation` shows dialogs to let the user enter
// passwords, ask questions or show processes blocking unmount.
type MountOperation interface {
	gextras.Objector

	// Display gets the display on which windows of the `GtkMountOperation` will
	// be shown.
	Display() *gdk.DisplayClass
	// Parent gets the transient parent used by the `GtkMountOperation`.
	Parent() *WindowClass
	// IsShowing returns whether the `GtkMountOperation` is currently displaying
	// a window.
	IsShowing() bool
	// SetDisplay sets the display to show windows of the `GtkMountOperation`
	// on.
	SetDisplay(display gdk.Display)
	// SetParent sets the transient parent for windows shown by the
	// `GtkMountOperation`.
	SetParent(parent Window)
}

// MountOperationClass implements the MountOperation interface.
type MountOperationClass struct {
	gio.MountOperationClass
}

var _ MountOperation = (*MountOperationClass)(nil)

func wrapMountOperation(obj *externglib.Object) MountOperation {
	return &MountOperationClass{
		MountOperationClass: gio.MountOperationClass{
			Object: obj,
		},
	}
}

func marshalMountOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMountOperation(obj), nil
}

// NewMountOperation creates a new `GtkMountOperation`.
func NewMountOperation(parent Window) *MountOperationClass {
	var _arg1 *C.GtkWindow       // out
	var _cret *C.GMountOperation // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer((&Window).Native()))

	_cret = C.gtk_mount_operation_new(_arg1)

	var _mountOperation *MountOperationClass // out

	_mountOperation = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*MountOperationClass)

	return _mountOperation
}

// Display gets the display on which windows of the `GtkMountOperation` will be
// shown.
func (o *MountOperationClass) Display() *gdk.DisplayClass {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GdkDisplay        // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer((&MountOperation).Native()))

	_cret = C.gtk_mount_operation_get_display(_arg0)

	var _display *gdk.DisplayClass // out

	_display = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*gdk.DisplayClass)

	return _display
}

// Parent gets the transient parent used by the `GtkMountOperation`.
func (o *MountOperationClass) Parent() *WindowClass {
	var _arg0 *C.GtkMountOperation // out
	var _cret *C.GtkWindow         // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer((&MountOperation).Native()))

	_cret = C.gtk_mount_operation_get_parent(_arg0)

	var _window *WindowClass // out

	_window = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*WindowClass)

	return _window
}

// IsShowing returns whether the `GtkMountOperation` is currently displaying a
// window.
func (o *MountOperationClass) IsShowing() bool {
	var _arg0 *C.GtkMountOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer((&MountOperation).Native()))

	_cret = C.gtk_mount_operation_is_showing(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDisplay sets the display to show windows of the `GtkMountOperation` on.
func (o *MountOperationClass) SetDisplay(display gdk.Display) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GdkDisplay        // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer((&MountOperation).Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer((&gdk.Display).Native()))

	C.gtk_mount_operation_set_display(_arg0, _arg1)
}

// SetParent sets the transient parent for windows shown by the
// `GtkMountOperation`.
func (o *MountOperationClass) SetParent(parent Window) {
	var _arg0 *C.GtkMountOperation // out
	var _arg1 *C.GtkWindow         // out

	_arg0 = (*C.GtkMountOperation)(unsafe.Pointer((&MountOperation).Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer((&Window).Native()))

	C.gtk_mount_operation_set_parent(_arg0, _arg1)
}
