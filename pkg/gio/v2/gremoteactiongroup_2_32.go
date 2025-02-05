// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// void _gotk4_gio2_RemoteActionGroup_virtual_activate_action_full(void* fnptr, GRemoteActionGroup* arg0, gchar* arg1, GVariant* arg2, GVariant* arg3) {
//   ((void (*)(GRemoteActionGroup*, gchar*, GVariant*, GVariant*))(fnptr))(arg0, arg1, arg2, arg3);
// };
// void _gotk4_gio2_RemoteActionGroup_virtual_change_action_state_full(void* fnptr, GRemoteActionGroup* arg0, gchar* arg1, GVariant* arg2, GVariant* arg3) {
//   ((void (*)(GRemoteActionGroup*, gchar*, GVariant*, GVariant*))(fnptr))(arg0, arg1, arg2, arg3);
// };
import "C"

// GType values.
var (
	GTypeRemoteActionGroup = coreglib.Type(C.g_remote_action_group_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRemoteActionGroup, F: marshalRemoteActionGroup},
	})
}

// RemoteActionGroup interface is implemented by Group instances that either
// transmit action invocations to other processes or receive action invocations
// in the local process from other processes.
//
// The interface has _full variants of the two methods on Group used to activate
// actions: g_action_group_activate_action() and
// g_action_group_change_action_state(). These variants allow a "platform data"
// #GVariant to be specified: a dictionary providing context for the action
// invocation (for example: timestamps, startup notification IDs, etc).
//
// BusActionGroup implements ActionGroup. This provides a mechanism to send
// platform data for action invocations over D-Bus.
//
// Additionally, g_dbus_connection_export_action_group() will check if the
// exported Group implements ActionGroup and use the _full variants of the calls
// if available. This provides a mechanism by which to receive platform data for
// action invocations that arrive by way of D-Bus.
//
// RemoteActionGroup wraps an interface. This means the user can get the
// underlying type by calling Cast().
type RemoteActionGroup struct {
	_ [0]func() // equal guard
	ActionGroup
}

var ()

// RemoteActionGrouper describes RemoteActionGroup's interface methods.
type RemoteActionGrouper interface {
	coreglib.Objector

	// ActivateActionFull activates the remote action.
	ActivateActionFull(actionName string, parameter, platformData *glib.Variant)
	// ChangeActionStateFull changes the state of a remote action.
	ChangeActionStateFull(actionName string, value, platformData *glib.Variant)
}

var _ RemoteActionGrouper = (*RemoteActionGroup)(nil)

func wrapRemoteActionGroup(obj *coreglib.Object) *RemoteActionGroup {
	return &RemoteActionGroup{
		ActionGroup: ActionGroup{
			Object: obj,
		},
	}
}

func marshalRemoteActionGroup(p uintptr) (interface{}, error) {
	return wrapRemoteActionGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ActivateActionFull activates the remote action.
//
// This is the same as g_action_group_activate_action() except that it allows
// for provision of "platform data" to be sent along with the activation
// request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// platform_data must be non-NULL and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to activate.
//    - parameter (optional): optional parameter to the activation.
//    - platformData: platform data to send.
//
func (remote *RemoteActionGroup) ActivateActionFull(actionName string, parameter, platformData *glib.Variant) {
	var _arg0 *C.GRemoteActionGroup // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 *C.GVariant           // out

	_arg0 = (*C.GRemoteActionGroup)(unsafe.Pointer(coreglib.InternObject(remote).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameter != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameter)))
	}
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(platformData)))

	C.g_remote_action_group_activate_action_full(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(remote)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(parameter)
	runtime.KeepAlive(platformData)
}

// ChangeActionStateFull changes the state of a remote action.
//
// This is the same as g_action_group_change_action_state() except that it
// allows for provision of "platform data" to be sent along with the state
// change request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// platform_data must be non-NULL and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to change the state of.
//    - value: new requested value for the state.
//    - platformData: platform data to send.
//
func (remote *RemoteActionGroup) ChangeActionStateFull(actionName string, value, platformData *glib.Variant) {
	var _arg0 *C.GRemoteActionGroup // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 *C.GVariant           // out

	_arg0 = (*C.GRemoteActionGroup)(unsafe.Pointer(coreglib.InternObject(remote).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(platformData)))

	C.g_remote_action_group_change_action_state_full(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(remote)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(value)
	runtime.KeepAlive(platformData)
}

// activateActionFull activates the remote action.
//
// This is the same as g_action_group_activate_action() except that it allows
// for provision of "platform data" to be sent along with the activation
// request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// platform_data must be non-NULL and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to activate.
//    - parameter (optional): optional parameter to the activation.
//    - platformData: platform data to send.
//
func (remote *RemoteActionGroup) activateActionFull(actionName string, parameter, platformData *glib.Variant) {
	gclass := (*C.GRemoteActionGroupInterface)(coreglib.PeekParentClass(remote))
	fnarg := gclass.activate_action_full

	var _arg0 *C.GRemoteActionGroup // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 *C.GVariant           // out

	_arg0 = (*C.GRemoteActionGroup)(unsafe.Pointer(coreglib.InternObject(remote).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameter != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(parameter)))
	}
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(platformData)))

	C._gotk4_gio2_RemoteActionGroup_virtual_activate_action_full(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(remote)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(parameter)
	runtime.KeepAlive(platformData)
}

// changeActionStateFull changes the state of a remote action.
//
// This is the same as g_action_group_change_action_state() except that it
// allows for provision of "platform data" to be sent along with the state
// change request. This typically contains details such as the user interaction
// timestamp or startup notification information.
//
// platform_data must be non-NULL and must have the type G_VARIANT_TYPE_VARDICT.
// If it is floating, it will be consumed.
//
// The function takes the following parameters:
//
//    - actionName: name of the action to change the state of.
//    - value: new requested value for the state.
//    - platformData: platform data to send.
//
func (remote *RemoteActionGroup) changeActionStateFull(actionName string, value, platformData *glib.Variant) {
	gclass := (*C.GRemoteActionGroupInterface)(coreglib.PeekParentClass(remote))
	fnarg := gclass.change_action_state_full

	var _arg0 *C.GRemoteActionGroup // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.GVariant           // out
	var _arg3 *C.GVariant           // out

	_arg0 = (*C.GRemoteActionGroup)(unsafe.Pointer(coreglib.InternObject(remote).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(platformData)))

	C._gotk4_gio2_RemoteActionGroup_virtual_change_action_state_full(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(remote)
	runtime.KeepAlive(actionName)
	runtime.KeepAlive(value)
	runtime.KeepAlive(platformData)
}

// RemoteActionGroupInterface: virtual function table for ActionGroup.
//
// An instance of this type is always passed by reference.
type RemoteActionGroupInterface struct {
	*remoteActionGroupInterface
}

// remoteActionGroupInterface is the struct that's finalized.
type remoteActionGroupInterface struct {
	native *C.GRemoteActionGroupInterface
}
