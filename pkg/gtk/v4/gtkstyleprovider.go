// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_style_provider_get_type()), F: marshalStyleProviderer},
	})
}

// STYLE_PROVIDER_PRIORITY_APPLICATION: priority that can be used when adding a
// GtkStyleProvider for application-specific style information.
const STYLE_PROVIDER_PRIORITY_APPLICATION = 600

// STYLE_PROVIDER_PRIORITY_FALLBACK: priority used for default style information
// that is used in the absence of themes.
//
// Note that this is not very useful for providing default styling for custom
// style classes - themes are likely to override styling provided at this
// priority with catch-all * {...} rules.
const STYLE_PROVIDER_PRIORITY_FALLBACK = 1

// STYLE_PROVIDER_PRIORITY_SETTINGS: priority used for style information
// provided via GtkSettings.
//
// This priority is higher than K_STYLE_PROVIDER_PRIORITY_THEME to let settings
// override themes.
const STYLE_PROVIDER_PRIORITY_SETTINGS = 400

// STYLE_PROVIDER_PRIORITY_THEME: priority used for style information provided
// by themes.
const STYLE_PROVIDER_PRIORITY_THEME = 200

// STYLE_PROVIDER_PRIORITY_USER: priority used for the style information from
// $XDG_CONFIG_HOME/gtk-4.0/gtk.css.
//
// You should not use priorities higher than this, to give the user the last
// word.
const STYLE_PROVIDER_PRIORITY_USER = 800

// StyleProvider: GtkStyleProvider is an interface for style information used by
// GtkStyleContext.
//
// See gtk.StyleContext.AddProvider() and
// gtk.StyleContext().AddProviderForDisplay for adding GtkStyleProviders.
//
// GTK uses the GtkStyleProvider implementation for CSS in gtk.CSSProvider.
type StyleProvider struct {
	*externglib.Object
}

// StyleProviderer describes StyleProvider's abstract methods.
type StyleProviderer interface {
	externglib.Objector

	privateStyleProvider()
}

var _ StyleProviderer = (*StyleProvider)(nil)

func wrapStyleProvider(obj *externglib.Object) *StyleProvider {
	return &StyleProvider{
		Object: obj,
	}
}

func marshalStyleProviderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStyleProvider(obj), nil
}

func (*StyleProvider) privateStyleProvider() {}
