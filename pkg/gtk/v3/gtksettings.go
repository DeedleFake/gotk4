// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeSettings = coreglib.Type(C.gtk_settings_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSettings, F: marshalSettings},
	})
}

// SettingsOverrides contains methods that are overridable.
type SettingsOverrides struct {
}

func defaultSettingsOverrides(v *Settings) SettingsOverrides {
	return SettingsOverrides{}
}

// Settings provide a mechanism to share global settings between applications.
//
// On the X window system, this sharing is realized by an XSettings
// (http://www.freedesktop.org/wiki/Specifications/xsettings-spec) manager that
// is usually part of the desktop environment, along with utilities that let the
// user change these settings. In the absence of an Xsettings manager, GTK+
// reads default values for settings from settings.ini files in /etc/gtk-3.0,
// $XDG_CONFIG_DIRS/gtk-3.0 and $XDG_CONFIG_HOME/gtk-3.0. These files must be
// valid key files (see File), and have a section called Settings. Themes can
// also provide default values for settings by installing a settings.ini file
// next to their gtk.css file.
//
// Applications can override system-wide settings by setting the property of the
// GtkSettings object with g_object_set(). This should be restricted to special
// cases though; GtkSettings are not meant as an application configuration
// facility. When doing so, you need to be aware that settings that are specific
// to individual widgets may not be available before the widget type has been
// realized at least once. The following example demonstrates a way to do this:
//
//      gtk_init (&argc, &argv);
//
//      // make sure the type is realized
//      g_type_class_unref (g_type_class_ref (GTK_TYPE_IMAGE_MENU_ITEM));
//
//      g_object_set (gtk_settings_get_default (), "gtk-enable-animations", FALSE, NULL);
//
// There is one GtkSettings instance per screen. It can be obtained with
// gtk_settings_get_for_screen(), but in many cases, it is more convenient to
// use gtk_widget_get_settings(). gtk_settings_get_default() returns the
// GtkSettings instance for the default screen.
type Settings struct {
	_ [0]func() // equal guard
	*coreglib.Object

	StyleProvider
}

var (
	_ coreglib.Objector = (*Settings)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Settings, *SettingsClass, SettingsOverrides](
		GTypeSettings,
		initSettingsClass,
		wrapSettings,
		defaultSettingsOverrides,
	)
}

func initSettingsClass(gclass unsafe.Pointer, overrides SettingsOverrides, classInitFunc func(*SettingsClass)) {
	if classInitFunc != nil {
		class := (*SettingsClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSettings(obj *coreglib.Object) *Settings {
	return &Settings{
		Object: obj,
		StyleProvider: StyleProvider{
			Object: obj,
		},
	}
}

func marshalSettings(p uintptr) (interface{}, error) {
	return wrapSettings(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ResetProperty undoes the effect of calling g_object_set() to install an
// application-specific value for a setting. After this call, the setting will
// again follow the session-wide value for this setting.
//
// The function takes the following parameters:
//
//    - name of the setting to reset.
//
func (settings *Settings) ResetProperty(name string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_settings_reset_property(_arg0, _arg1)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(name)
}

// SetDoubleProperty: deprecated: Use g_object_set() instead.
//
// The function takes the following parameters:
//
//    - name
//    - vDouble
//    - origin
//
func (settings *Settings) SetDoubleProperty(name string, vDouble float64, origin string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.gchar       // out
	var _arg2 C.gdouble      // out
	var _arg3 *C.gchar       // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(vDouble)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(origin)))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_settings_set_double_property(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(name)
	runtime.KeepAlive(vDouble)
	runtime.KeepAlive(origin)
}

// SetLongProperty: deprecated: Use g_object_set() instead.
//
// The function takes the following parameters:
//
//    - name
//    - vLong
//    - origin
//
func (settings *Settings) SetLongProperty(name string, vLong int32, origin string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.gchar       // out
	var _arg2 C.glong        // out
	var _arg3 *C.gchar       // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.glong(vLong)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(origin)))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_settings_set_long_property(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(name)
	runtime.KeepAlive(vLong)
	runtime.KeepAlive(origin)
}

// SetPropertyValue: deprecated: Use g_object_set() instead.
//
// The function takes the following parameters:
//
//    - name
//    - svalue
//
func (settings *Settings) SetPropertyValue(name string, svalue *SettingsValue) {
	var _arg0 *C.GtkSettings      // out
	var _arg1 *C.gchar            // out
	var _arg2 *C.GtkSettingsValue // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkSettingsValue)(gextras.StructNative(unsafe.Pointer(svalue)))

	C.gtk_settings_set_property_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(name)
	runtime.KeepAlive(svalue)
}

// SetStringProperty: deprecated: Use g_object_set() instead.
//
// The function takes the following parameters:
//
//    - name
//    - vString
//    - origin
//
func (settings *Settings) SetStringProperty(name, vString, origin string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.gchar       // out
	var _arg2 *C.gchar       // out
	var _arg3 *C.gchar       // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(vString)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(origin)))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_settings_set_string_property(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(name)
	runtime.KeepAlive(vString)
	runtime.KeepAlive(origin)
}

// SettingsGetDefault gets the Settings object for the default GDK screen,
// creating it if necessary. See gtk_settings_get_for_screen().
//
// The function returns the following values:
//
//    - settings (optional) object. If there is no default screen, then returns
//      NULL.
//
func SettingsGetDefault() *Settings {
	var _cret *C.GtkSettings // in

	_cret = C.gtk_settings_get_default()

	var _settings *Settings // out

	if _cret != nil {
		_settings = wrapSettings(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _settings
}

// SettingsClass: instance of this type is always passed by reference.
type SettingsClass struct {
	*settingsClass
}

// settingsClass is the struct that's finalized.
type settingsClass struct {
	native *C.GtkSettingsClass
}

// SettingsValue: instance of this type is always passed by reference.
type SettingsValue struct {
	*settingsValue
}

// settingsValue is the struct that's finalized.
type settingsValue struct {
	native *C.GtkSettingsValue
}

// Origin should be something like “filename:linenumber” for rc files, or e.g.
// “XProperty” for other sources.
func (s *SettingsValue) Origin() string {
	valptr := &s.native.origin
	var _v string // out
	_v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return _v
}
