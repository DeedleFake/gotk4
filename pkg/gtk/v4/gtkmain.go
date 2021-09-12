// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

// PRIORITY_RESIZE: use this priority for functionality related to size
// allocation.
//
// It is used internally by GTK+ to compute the sizes of widgets. This priority
// is higher than GDK_PRIORITY_REDRAW to avoid resizing a widget which was just
// redrawn.
const PRIORITY_RESIZE = 110

// DisableSetlocale prevents gtk_init and gtk_init_check from automatically
// calling setlocale (LC_ALL, "").
//
// You would want to use this function if you wanted to set the locale for your
// program to something other than the user’s locale, or if you wanted to set
// different values for different locale categories.
//
// Most programs should not need to call this function.
func DisableSetlocale() {
	C.gtk_disable_setlocale()
}

// GetDefaultLanguage returns the Language for the default language currently in
// effect. (Note that this can change over the life of an application.) The
// default language is derived from the current locale. It determines, for
// example, whether GTK uses the right-to-left or left-to-right text direction.
//
// This function is equivalent to pango_language_get_default(). See that
// function for details.
func GetDefaultLanguage() *pango.Language {
	var _cret *C.PangoLanguage // in

	_cret = C.gtk_get_default_language()

	var _language *pango.Language // out

	_language = (*pango.Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _language
}

// GetLocaleDirection: get the direction of the current locale. This is the
// expected reading direction for text and UI.
//
// This function depends on the current locale being set with setlocale() and
// will default to setting the GTK_TEXT_DIR_LTR direction otherwise.
// GTK_TEXT_DIR_NONE will never be returned.
//
// GTK sets the default text direction according to the locale during
// gtk_init(), and you should normally use gtk_widget_get_direction() or
// gtk_widget_get_default_direction() to obtain the current direction.
//
// This function is only needed rare cases when the locale is changed after GTK
// has already been initialized. In this case, you can use it to update the
// default text direction as follows:
//
//    setlocale (LC_ALL, new_locale);
//    direction = gtk_get_locale_direction ();
//    gtk_widget_set_default_direction (direction);
func GetLocaleDirection() TextDirection {
	var _cret C.GtkTextDirection // in

	_cret = C.gtk_get_locale_direction()

	var _textDirection TextDirection // out

	_textDirection = TextDirection(_cret)

	return _textDirection
}

// Init: call this function before using any other GTK functions in your GUI
// applications. It will initialize everything needed to operate the toolkit and
// parses some standard command line options.
//
// If you are using Application, you don't have to call gtk_init() or
// gtk_init_check(); the #GApplication::startup handler does it for you.
//
// This function will terminate your program if it was unable to initialize the
// windowing system for some reason. If you want your program to fall back to a
// textual interface you want to call gtk_init_check() instead.
//
// GTK calls signal (SIGPIPE, SIG_IGN) during initialization, to ignore SIGPIPE
// signals, since these are almost never wanted in graphical applications. If
// you do need to handle SIGPIPE for some reason, reset the handler after
// gtk_init(), but notice that other libraries (e.g. libdbus or gvfs) might do
// similar things.
func Init() {
	C.gtk_init()
}

// InitCheck: this function does the same work as gtk_init() with only a single
// change: It does not terminate the program if the windowing system can’t be
// initialized. Instead it returns FALSE on failure.
//
// This way the application can fall back to some other means of communication
// with the user - for example a curses or command line interface.
func InitCheck() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_init_check()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsInitialized: use this function to check if GTK has been initialized with
// gtk_init() or gtk_init_check().
func IsInitialized() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_is_initialized()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
