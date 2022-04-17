// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// NewNotification creates a new #GNotification with title as its title.
//
// After populating notification with more details, it can be sent to the
// desktop shell with g_application_send_notification(). Changing any properties
// after this call will not have any effect until resending notification.
//
// The function takes the following parameters:
//
//    - title of the notification.
//
// The function returns the following values:
//
//    - notification: new #GNotification instance.
//
func NewNotification(title string) *Notification {
	var _arg1 *C.gchar         // out
	var _cret *C.GNotification // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_notification_new(_arg1)
	runtime.KeepAlive(title)

	var _notification *Notification // out

	_notification = wrapNotification(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _notification
}

// AddButton adds a button to notification that activates the action in
// detailed_action when clicked. That action must be an application-wide action
// (starting with "app."). If detailed_action contains a target, the action will
// be activated with that target as its parameter.
//
// See g_action_parse_detailed_name() for a description of the format for
// detailed_action.
//
// The function takes the following parameters:
//
//    - label of the button.
//    - detailedAction: detailed action name.
//
func (notification *Notification) AddButton(label, detailedAction string) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_notification_add_button(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(label)
	runtime.KeepAlive(detailedAction)
}

// AddButtonWithTarget adds a button to notification that activates action when
// clicked. action must be an application-wide action (it must start with
// "app.").
//
// If target is non-NULL, action will be activated with target as its parameter.
//
// The function takes the following parameters:
//
//    - label of the button.
//    - action name.
//    - target (optional) to use as action's parameter, or NULL.
//
func (notification *Notification) AddButtonWithTarget(label, action string, target *glib.Variant) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.gchar         // out
	var _arg3 *C.GVariant      // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(action)))
	defer C.free(unsafe.Pointer(_arg2))
	if target != nil {
		_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(target)))
	}

	C.g_notification_add_button_with_target_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(label)
	runtime.KeepAlive(action)
	runtime.KeepAlive(target)
}

// SetBody sets the body of notification to body.
//
// The function takes the following parameters:
//
//    - body (optional): new body for notification, or NULL.
//
func (notification *Notification) SetBody(body string) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	if body != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(body)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.g_notification_set_body(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(body)
}

// SetDefaultAction sets the default action of notification to detailed_action.
// This action is activated when the notification is clicked on.
//
// The action in detailed_action must be an application-wide action (it must
// start with "app."). If detailed_action contains a target, the given action
// will be activated with that target as its parameter. See
// g_action_parse_detailed_name() for a description of the format for
// detailed_action.
//
// When no default action is set, the application that the notification was sent
// on is activated.
//
// The function takes the following parameters:
//
//    - detailedAction: detailed action name.
//
func (notification *Notification) SetDefaultAction(detailedAction string) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(detailedAction)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_notification_set_default_action(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(detailedAction)
}

// SetDefaultActionAndTarget sets the default action of notification to action.
// This action is activated when the notification is clicked on. It must be an
// application-wide action (start with "app.").
//
// If target is non-NULL, action will be activated with target as its parameter.
//
// When no default action is set, the application that the notification was sent
// on is activated.
//
// The function takes the following parameters:
//
//    - action name.
//    - target (optional) to use as action's parameter, or NULL.
//
func (notification *Notification) SetDefaultActionAndTarget(action string, target *glib.Variant) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.GVariant      // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(action)))
	defer C.free(unsafe.Pointer(_arg1))
	if target != nil {
		_arg2 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(target)))
	}

	C.g_notification_set_default_action_and_target_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(action)
	runtime.KeepAlive(target)
}

// SetIcon sets the icon of notification to icon.
//
// The function takes the following parameters:
//
//    - icon to be shown in notification, as a #GIcon.
//
func (notification *Notification) SetIcon(icon Iconner) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.GIcon         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(externglib.InternObject(icon).Native()))

	C.g_notification_set_icon(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(icon)
}

// SetPriority sets the priority of notification to priority. See Priority for
// possible values.
//
// The function takes the following parameters:
//
//    - priority: Priority.
//
func (notification *Notification) SetPriority(priority NotificationPriority) {
	var _arg0 *C.GNotification        // out
	var _arg1 C.GNotificationPriority // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	_arg1 = C.GNotificationPriority(priority)

	C.g_notification_set_priority(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(priority)
}

// SetTitle sets the title of notification to title.
//
// The function takes the following parameters:
//
//    - title: new title for notification.
//
func (notification *Notification) SetTitle(title string) {
	var _arg0 *C.GNotification // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_notification_set_title(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(title)
}

// SetUrgent: deprecated in favor of g_notification_set_priority().
//
// Deprecated: Since 2.42, this has been deprecated in favour of
// g_notification_set_priority().
//
// The function takes the following parameters:
//
//    - urgent: TRUE if notification is urgent.
//
func (notification *Notification) SetUrgent(urgent bool) {
	var _arg0 *C.GNotification // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GNotification)(unsafe.Pointer(externglib.InternObject(notification).Native()))
	if urgent {
		_arg1 = C.TRUE
	}

	C.g_notification_set_urgent(_arg0, _arg1)
	runtime.KeepAlive(notification)
	runtime.KeepAlive(urgent)
}
