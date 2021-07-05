// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_css_provider_error_get_type()), F: marshalCSSProviderError},
		{T: externglib.Type(C.gtk_css_provider_get_type()), F: marshalCSSProvider},
	})
}

// CSSProviderError: error codes for GTK_CSS_PROVIDER_ERROR.
type CSSProviderError int

const (
	// failed: failed.
	CSSProviderErrorFailed CSSProviderError = iota
	// syntax: syntax error.
	CSSProviderErrorSyntax
	// import: import error.
	CSSProviderErrorImport
	// name: name error.
	CSSProviderErrorName
	// deprecated: deprecation error.
	CSSProviderErrorDeprecated
	// UnknownValue: unknown value.
	CSSProviderErrorUnknownValue
)

func marshalCSSProviderError(p uintptr) (interface{}, error) {
	return CSSProviderError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CSSProvider is an object implementing the StyleProvider interface. It is able
// to parse [CSS-like][css-overview] input in order to style widgets.
//
// An application can make GTK+ parse a specific CSS style sheet by calling
// gtk_css_provider_load_from_file() or gtk_css_provider_load_from_resource()
// and adding the provider with gtk_style_context_add_provider() or
// gtk_style_context_add_provider_for_screen().
//
// In addition, certain files will be read when GTK+ is initialized. First, the
// file `$XDG_CONFIG_HOME/gtk-3.0/gtk.css` is loaded if it exists. Then, GTK+
// loads the first existing file among
// `XDG_DATA_HOME/themes/THEME/gtk-VERSION/gtk.css`,
// `$HOME/.themes/THEME/gtk-VERSION/gtk.css`,
// `$XDG_DATA_DIRS/themes/THEME/gtk-VERSION/gtk.css` and
// `DATADIR/share/themes/THEME/gtk-VERSION/gtk.css`, where `THEME` is the name
// of the current theme (see the Settings:gtk-theme-name setting), `DATADIR` is
// the prefix configured when GTK+ was compiled (unless overridden by the
// `GTK_DATA_PREFIX` environment variable), and `VERSION` is the GTK+ version
// number. If no file is found for the current version, GTK+ tries older
// versions all the way back to 3.0.
//
// In the same way, GTK+ tries to load a gtk-keys.css file for the current key
// theme, as defined by Settings:gtk-key-theme-name.
type CSSProvider interface {
	gextras.Objector

	// AsStyleProvider casts the class to the StyleProvider interface.
	AsStyleProvider() StyleProvider

	// LoadFromDataCSSProvider loads @data into @css_provider, and by doing so
	// clears any previously loaded information.
	LoadFromDataCSSProvider(data []byte) error
	// LoadFromPathCSSProvider loads the data contained in @path into
	// @css_provider, making it clear any previously loaded information.
	LoadFromPathCSSProvider(path string) error
	// LoadFromResourceCSSProvider loads the data contained in the resource at
	// @resource_path into the CssProvider, clearing any previously loaded
	// information.
	//
	// To track errors while loading CSS, connect to the
	// CssProvider::parsing-error signal.
	LoadFromResourceCSSProvider(resourcePath string)
	// String converts the @provider into a string representation in CSS format.
	//
	// Using gtk_css_provider_load_from_data() with the return value from this
	// function on a new provider created with gtk_css_provider_new() will
	// basically create a duplicate of this @provider.
	String() string
}

// cssProvider implements the CSSProvider class.
type cssProvider struct {
	gextras.Objector
}

// WrapCSSProvider wraps a GObject to the right type. It is
// primarily used internally.
func WrapCSSProvider(obj *externglib.Object) CSSProvider {
	return cssProvider{
		Objector: obj,
	}
}

func marshalCSSProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCSSProvider(obj), nil
}

// NewCSSProvider returns a newly created CssProvider.
func NewCSSProvider() CSSProvider {
	var _cret *C.GtkCssProvider // in

	_cret = C.gtk_css_provider_new()

	var _cssProvider CSSProvider // out

	_cssProvider = WrapCSSProvider(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cssProvider
}

func (c cssProvider) LoadFromDataCSSProvider(data []byte) error {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.gchar
	var _arg2 C.gssize
	var _cerr *C.GError // in

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(c.Native()))
	_arg2 = C.gssize(len(data))
	_arg1 = (*C.gchar)(unsafe.Pointer(&data[0]))

	C.gtk_css_provider_load_from_data(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (c cssProvider) LoadFromPathCSSProvider(path string) error {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.gchar          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_css_provider_load_from_path(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (c cssProvider) LoadFromResourceCSSProvider(resourcePath string) {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_css_provider_load_from_resource(_arg0, _arg1)
}

func (p cssProvider) String() string {
	var _arg0 *C.GtkCssProvider // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_css_provider_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (c cssProvider) AsStyleProvider() StyleProvider {
	return WrapStyleProvider(gextras.InternObject(c))
}
