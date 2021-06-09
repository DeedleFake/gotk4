// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_css_provider_get_type()), F: marshalCSSProvider},
	})
}

// CSSProvider: `GtkCssProvider` is an object implementing the
// `GtkStyleProvider` interface for CSS.
//
// It is able to parse CSS-like input in order to style widgets.
//
// An application can make GTK parse a specific CSS style sheet by calling
// [method@Gtk.CssProvider.load_from_file] or
// [method@Gtk.CssProvider.load_from_resource] and adding the provider with
// [method@Gtk.StyleContext.add_provider] or
// [func@Gtk.StyleContext.add_provider_for_display].
//
// In addition, certain files will be read when GTK is initialized. First, the
// file `$XDG_CONFIG_HOME/gtk-4.0/gtk.css` is loaded if it exists. Then, GTK
// loads the first existing file among
// `XDG_DATA_HOME/themes/THEME/gtk-VERSION/gtk-VARIANT.css`,
// `$HOME/.themes/THEME/gtk-VERSION/gtk-VARIANT.css`,
// `$XDG_DATA_DIRS/themes/THEME/gtk-VERSION/gtk-VARIANT.css` and
// `DATADIR/share/themes/THEME/gtk-VERSION/gtk-VARIANT.css`, where `THEME` is
// the name of the current theme (see the [property@Gtk.Settings:gtk-theme-name]
// setting), `VARIANT` is the variant to load (see the
// [property@Gtk.Settings:gtk-application-prefer-dark-theme] setting), `DATADIR`
// is the prefix configured when GTK was compiled (unless overridden by the
// `GTK_DATA_PREFIX` environment variable), and `VERSION` is the GTK version
// number. If no file is found for the current version, GTK tries older versions
// all the way back to 4.0.
//
// To track errors while loading CSS, connect to the
// [signal@Gtk.CssProvider::parsing-error] signal.
type CSSProvider interface {
	gextras.Objector
	StyleProvider

	// LoadFromData loads @data into @css_provider.
	//
	// This clears any previously loaded information.
	LoadFromData()
	// LoadFromFile loads the data contained in @file into @css_provider.
	//
	// This clears any previously loaded information.
	LoadFromFile(file gio.File)
	// LoadFromPath loads the data contained in @path into @css_provider.
	//
	// This clears any previously loaded information.
	LoadFromPath(path string)
	// LoadFromResource loads the data contained in the resource at
	// @resource_path into the @css_provider.
	//
	// This clears any previously loaded information.
	LoadFromResource(resourcePath string)
	// LoadNamed loads a theme from the usual theme paths.
	//
	// The actual process of finding the theme might change between releases,
	// but it is guaranteed that this function uses the same mechanism to load
	// the theme that GTK uses for loading its own theme.
	LoadNamed(name string, variant string)
	// String converts the @provider into a string representation in CSS format.
	//
	// Using [method@Gtk.CssProvider.load_from_data] with the return value from
	// this function on a new provider created with [ctor@Gtk.CssProvider.new]
	// will basically create a duplicate of this @provider.
	String() string
}

// cssProvider implements the CSSProvider interface.
type cssProvider struct {
	gextras.Objector
	StyleProvider
}

var _ CSSProvider = (*cssProvider)(nil)

// WrapCSSProvider wraps a GObject to the right type. It is
// primarily used internally.
func WrapCSSProvider(obj *externglib.Object) CSSProvider {
	return CSSProvider{
		Objector:      obj,
		StyleProvider: WrapStyleProvider(obj),
	}
}

func marshalCSSProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCSSProvider(obj), nil
}

// NewCSSProvider constructs a class CSSProvider.
func NewCSSProvider() CSSProvider {
	var _cret C.GtkCssProvider

	cret = C.gtk_css_provider_new()

	var _cssProvider CSSProvider

	_cssProvider = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(CSSProvider)

	return _cssProvider
}

// LoadFromData loads @data into @css_provider.
//
// This clears any previously loaded information.
func (c cssProvider) LoadFromData() {
	var _arg0 *C.GtkCssProvider

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(c.Native()))

	C.gtk_css_provider_load_from_data(_arg0)
}

// LoadFromFile loads the data contained in @file into @css_provider.
//
// This clears any previously loaded information.
func (c cssProvider) LoadFromFile(file gio.File) {
	var _arg0 *C.GtkCssProvider
	var _arg1 *C.GFile

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_css_provider_load_from_file(_arg0, _arg1)
}

// LoadFromPath loads the data contained in @path into @css_provider.
//
// This clears any previously loaded information.
func (c cssProvider) LoadFromPath(path string) {
	var _arg0 *C.GtkCssProvider
	var _arg1 *C.char

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_css_provider_load_from_path(_arg0, _arg1)
}

// LoadFromResource loads the data contained in the resource at
// @resource_path into the @css_provider.
//
// This clears any previously loaded information.
func (c cssProvider) LoadFromResource(resourcePath string) {
	var _arg0 *C.GtkCssProvider
	var _arg1 *C.char

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_css_provider_load_from_resource(_arg0, _arg1)
}

// LoadNamed loads a theme from the usual theme paths.
//
// The actual process of finding the theme might change between releases,
// but it is guaranteed that this function uses the same mechanism to load
// the theme that GTK uses for loading its own theme.
func (p cssProvider) LoadNamed(name string, variant string) {
	var _arg0 *C.GtkCssProvider
	var _arg1 *C.char
	var _arg2 *C.char

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(variant))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_css_provider_load_named(_arg0, _arg1, _arg2)
}

// String converts the @provider into a string representation in CSS format.
//
// Using [method@Gtk.CssProvider.load_from_data] with the return value from
// this function on a new provider created with [ctor@Gtk.CssProvider.new]
// will basically create a duplicate of this @provider.
func (p cssProvider) String() string {
	var _arg0 *C.GtkCssProvider

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(p.Native()))

	var _cret *C.char

	cret = C.gtk_css_provider_to_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
