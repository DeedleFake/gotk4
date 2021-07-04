// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_settings_get_type()), F: marshalSettings},
	})
}

// Settings: `GtkSettings` provides a mechanism to share global settings between
// applications.
//
// On the X window system, this sharing is realized by an XSettings
// (http://www.freedesktop.org/wiki/Specifications/xsettings-spec) manager that
// is usually part of the desktop environment, along with utilities that let the
// user change these settings.
//
// On Wayland, the settings are obtained either via a settings portal, or by
// reading desktop settings from DConf.
//
// In the absence of these sharing mechanisms, GTK reads default values for
// settings from `settings.ini` files in `/etc/gtk-4.0`,
// `$XDG_CONFIG_DIRS/gtk-4.0` and `$XDG_CONFIG_HOME/gtk-4.0`. These files must
// be valid key files (see `GKeyFile`), and have a section called Settings.
// Themes can also provide default values for settings by installing a
// `settings.ini` file next to their `gtk.css` file.
//
// Applications can override system-wide settings by setting the property of the
// `GtkSettings` object with g_object_set(). This should be restricted to
// special cases though; `GtkSettings` are not meant as an application
// configuration facility.
//
// There is one `GtkSettings` instance per display. It can be obtained with
// [type_func@GtkSettings.get_for_display], but in many cases, it is more
// convenient to use [method@Gtk.Widget.get_settings].
type Settings interface {
	StyleProvider

	ResetPropertySettings(name string)
}

// settings implements the Settings class.
type settings struct {
	gextras.Objector
}

// WrapSettings wraps a GObject to the right type. It is
// primarily used internally.
func WrapSettings(obj *externglib.Object) Settings {
	return settings{
		Objector: obj,
	}
}

func marshalSettings(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSettings(obj), nil
}

func (s settings) ResetPropertySettings(name string) {
	var _arg0 *C.GtkSettings // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_settings_reset_property(_arg0, _arg1)
}
