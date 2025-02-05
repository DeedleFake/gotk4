// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// IconThemeAddBuiltinIcon registers a built-in icon for icon theme lookups. The
// idea of built-in icons is to allow an application or library that uses themed
// icons to function requiring files to be present in the file system. For
// instance, the default images for all of GTK+’s stock icons are registered as
// built-icons.
//
// In general, if you use gtk_icon_theme_add_builtin_icon() you should also
// install the icon in the icon theme, so that the icon is generally available.
//
// This function will generally be used with pixbufs loaded via
// gdk_pixbuf_new_from_inline().
//
// Deprecated: Use gtk_icon_theme_add_resource_path() to add
// application-specific icons to the icon theme.
//
// The function takes the following parameters:
//
//    - iconName: name of the icon to register.
//    - size in pixels at which to register the icon (different images can be
//      registered for the same icon name at different sizes.).
//    - pixbuf that contains the image to use for icon_name.
//
func IconThemeAddBuiltinIcon(iconName string, size int, pixbuf *gdkpixbuf.Pixbuf) {
	var _arg1 *C.gchar     // out
	var _arg2 C.gint       // out
	var _arg3 *C.GdkPixbuf // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(size)
	_arg3 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))

	C.gtk_icon_theme_add_builtin_icon(_arg1, _arg2, _arg3)
	runtime.KeepAlive(iconName)
	runtime.KeepAlive(size)
	runtime.KeepAlive(pixbuf)
}

// IconThemeGetDefault gets the icon theme for the default screen. See
// gtk_icon_theme_get_for_screen().
//
// The function returns the following values:
//
//    - iconTheme: unique IconTheme associated with the default screen. This icon
//      theme is associated with the screen and can be used as long as the screen
//      is open. Do not ref or unref it.
//
func IconThemeGetDefault() *IconTheme {
	var _cret *C.GtkIconTheme // in

	_cret = C.gtk_icon_theme_get_default()

	var _iconTheme *IconTheme // out

	_iconTheme = wrapIconTheme(coreglib.Take(unsafe.Pointer(_cret)))

	return _iconTheme
}

// IconThemeGetForScreen gets the icon theme object associated with screen; if
// this function has not previously been called for the given screen, a new icon
// theme object will be created and associated with the screen. Icon theme
// objects are fairly expensive to create, so using this function is usually a
// better choice than calling than gtk_icon_theme_new() and setting the screen
// yourself; by using this function a single icon theme object will be shared
// between users.
//
// The function takes the following parameters:
//
//    - screen: Screen.
//
// The function returns the following values:
//
//    - iconTheme: unique IconTheme associated with the given screen. This icon
//      theme is associated with the screen and can be used as long as the screen
//      is open. Do not ref or unref it.
//
func IconThemeGetForScreen(screen *gdk.Screen) *IconTheme {
	var _arg1 *C.GdkScreen    // out
	var _cret *C.GtkIconTheme // in

	_arg1 = (*C.GdkScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))

	_cret = C.gtk_icon_theme_get_for_screen(_arg1)
	runtime.KeepAlive(screen)

	var _iconTheme *IconTheme // out

	_iconTheme = wrapIconTheme(coreglib.Take(unsafe.Pointer(_cret)))

	return _iconTheme
}
