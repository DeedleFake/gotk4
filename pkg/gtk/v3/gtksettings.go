// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeSettings returns the GType for the type Settings.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSettings() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Settings").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSettings)
	return gtype
}

// SettingsOverrider contains methods that are overridable.
type SettingsOverrider interface {
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

func classInitSettingser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "Settings")
	_info.InvokeClassMethod("reset_property", _args[:], nil)

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
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.gdouble)(unsafe.Pointer(&_args[2])) = C.gdouble(vDouble)
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(origin)))
	defer C.free(unsafe.Pointer(_args[3]))

	_info := girepository.MustFind("Gtk", "Settings")
	_info.InvokeClassMethod("set_double_property", _args[:], nil)

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
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.glong)(unsafe.Pointer(&_args[2])) = C.glong(vLong)
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(origin)))
	defer C.free(unsafe.Pointer(_args[3]))

	_info := girepository.MustFind("Gtk", "Settings")
	_info.InvokeClassMethod("set_long_property", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(svalue)))

	_info := girepository.MustFind("Gtk", "Settings")
	_info.InvokeClassMethod("set_property_value", _args[:], nil)

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
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(vString)))
	defer C.free(unsafe.Pointer(_args[2]))
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(C.CString(origin)))
	defer C.free(unsafe.Pointer(_args[3]))

	_info := girepository.MustFind("Gtk", "Settings")
	_info.InvokeClassMethod("set_string_property", _args[:], nil)

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
	_info := girepository.MustFind("Gtk", "get_default")
	_gret := _info.Invoke(nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _settings *Settings // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_settings = wrapSettings(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _settings
}

// SettingsGetForScreen gets the Settings object for screen, creating it if
// necessary.
//
// The function takes the following parameters:
//
//    - screen: Screen.
//
// The function returns the following values:
//
//    - settings Settings object.
//
func SettingsGetForScreen(screen *gdk.Screen) *Settings {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_info := girepository.MustFind("Gtk", "get_for_screen")
	_gret := _info.Invoke(_args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(screen)

	var _settings *Settings // out

	_settings = wrapSettings(coreglib.Take(unsafe.Pointer(_cret)))

	return _settings
}

// SettingsValue: instance of this type is always passed by reference.
type SettingsValue struct {
	*settingsValue
}

// settingsValue is the struct that's finalized.
type settingsValue struct {
	native unsafe.Pointer
}

// Origin should be something like “filename:linenumber” for rc files, or e.g.
// “XProperty” for other sources.
func (s *SettingsValue) Origin() string {
	offset := girepository.MustFind("Gtk", "SettingsValue").StructFieldOffset("origin")
	valptr := (*uintptr)(unsafe.Add(s.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return v
}
