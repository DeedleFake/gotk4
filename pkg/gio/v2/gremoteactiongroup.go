// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_remote_action_group_get_type()), F: marshalRemoteActionGroup},
	})
}

// RemoteActionGroupOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RemoteActionGroupOverrider interface {
	// ActivateActionFull activates the remote action.
	//
	// This is the same as g_action_group_activate_action() except that it
	// allows for provision of "platform data" to be sent along with the
	// activation request. This typically contains details such as the user
	// interaction timestamp or startup notification information.
	//
	// @platform_data must be non-nil and must have the type
	// G_VARIANT_TYPE_VARDICT. If it is floating, it will be consumed.
	ActivateActionFull(actionName string, parameter *glib.Variant, platformData *glib.Variant)
	// ChangeActionStateFull changes the state of a remote action.
	//
	// This is the same as g_action_group_change_action_state() except that it
	// allows for provision of "platform data" to be sent along with the state
	// change request. This typically contains details such as the user
	// interaction timestamp or startup notification information.
	//
	// @platform_data must be non-nil and must have the type
	// G_VARIANT_TYPE_VARDICT. If it is floating, it will be consumed.
	ChangeActionStateFull(actionName string, value *glib.Variant, platformData *glib.Variant)
}

// RemoteActionGroup: the GRemoteActionGroup interface is implemented by Group
// instances that either transmit action invocations to other processes or
// receive action invocations in the local process from other processes.
//
// The interface has `_full` variants of the two methods on Group used to
// activate actions: g_action_group_activate_action() and
// g_action_group_change_action_state(). These variants allow a "platform data"
// #GVariant to be specified: a dictionary providing context for the action
// invocation (for example: timestamps, startup notification IDs, etc).
//
// BusActionGroup implements ActionGroup. This provides a mechanism to send
// platform data for action invocations over D-Bus.
//
// Additionally, g_dbus_connection_export_action_group() will check if the
// exported Group implements ActionGroup and use the `_full` variants of the
// calls if available. This provides a mechanism by which to receive platform
// data for action invocations that arrive by way of D-Bus.
type RemoteActionGroup interface {
	gextras.Objector

	// ActivateActionFull activates the remote action.
	//
	// This is the same as g_action_group_activate_action() except that it
	// allows for provision of "platform data" to be sent along with the
	// activation request. This typically contains details such as the user
	// interaction timestamp or startup notification information.
	//
	// @platform_data must be non-nil and must have the type
	// G_VARIANT_TYPE_VARDICT. If it is floating, it will be consumed.
	ActivateActionFull(actionName string, parameter *glib.Variant, platformData *glib.Variant)
	// ChangeActionStateFull changes the state of a remote action.
	//
	// This is the same as g_action_group_change_action_state() except that it
	// allows for provision of "platform data" to be sent along with the state
	// change request. This typically contains details such as the user
	// interaction timestamp or startup notification information.
	//
	// @platform_data must be non-nil and must have the type
	// G_VARIANT_TYPE_VARDICT. If it is floating, it will be consumed.
	ChangeActionStateFull(actionName string, value *glib.Variant, platformData *glib.Variant)
}

// RemoteActionGroupInterface implements the RemoteActionGroup interface.
type RemoteActionGroupInterface struct {
	ActionGroupInterface
}

var _ RemoteActionGroup = (*RemoteActionGroupInterface)(nil)

func wrapRemoteActionGroup(obj *externglib.Object) RemoteActionGroup {
	return &RemoteActionGroupInterface{
		ActionGroupInterface: ActionGroupInterface{
			Object: obj,
		},
	}
}

func marshalRemoteActionGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRemoteActionGroup(obj), nil
}

// ActivateActionFull activates the remote action.
//
// This is the same as g_action_group_activate_action() except that it allows
// for provision of "platform data" to be sent along with the activation
// request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// @platform_data must be non-nil and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
func (r *RemoteActionGroupInterface) ActivateActionFull(actionName string, parameter *glib.Variant, platformData *glib.Variant) {
	var _arg0 *C.GRemoteActionGroup // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 *C.GVariant           // out

	_arg0 = (*C.GRemoteActionGroup)(unsafe.Pointer((&RemoteActionGroup).Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(unsafe.Pointer(*glib.Variant))
	_arg3 = (*C.GVariant)(unsafe.Pointer(*glib.Variant))

	C.g_remote_action_group_activate_action_full(_arg0, _arg1, _arg2, _arg3)
}

// ChangeActionStateFull changes the state of a remote action.
//
// This is the same as g_action_group_change_action_state() except that it
// allows for provision of "platform data" to be sent along with the state
// change request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// @platform_data must be non-nil and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
func (r *RemoteActionGroupInterface) ChangeActionStateFull(actionName string, value *glib.Variant, platformData *glib.Variant) {
	var _arg0 *C.GRemoteActionGroup // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 *C.GVariant           // out

	_arg0 = (*C.GRemoteActionGroup)(unsafe.Pointer((&RemoteActionGroup).Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(unsafe.Pointer(*glib.Variant))
	_arg3 = (*C.GVariant)(unsafe.Pointer(*glib.Variant))

	C.g_remote_action_group_change_action_state_full(_arg0, _arg1, _arg2, _arg3)
}
