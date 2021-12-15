// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_get_type()), F: marshalAppChooserer},
	})
}

// AppChooser: GtkAppChooser is an interface for widgets which allow the user to
// choose an application.
//
// The main objects that implement this interface are gtk.AppChooserWidget,
// gtk.AppChooserDialog and gtk.AppChooserButton.
//
// Applications are represented by GIO GAppInfo objects here. GIO has a concept
// of recommended and fallback applications for a given content type.
// Recommended applications are those that claim to handle the content type
// itself, while fallback also includes applications that handle a more generic
// content type. GIO also knows the default and last-used application for a
// given content type. The GtkAppChooserWidget provides detailed control over
// whether the shown list of applications should include default, recommended or
// fallback applications.
//
// To obtain the application that has been selected in a GtkAppChooser, use
// gtk.AppChooser.GetAppInfo().
type AppChooser struct {
	Widget
}

var (
	_ Widgetter = (*AppChooser)(nil)
)

// AppChooserer describes AppChooser's interface methods.
type AppChooserer interface {
	externglib.Objector

	// AppInfo returns the currently selected application.
	AppInfo() gio.AppInfor
	// ContentType returns the content type for which the GtkAppChooser shows
	// applications.
	ContentType() string
	// Refresh reloads the list of applications.
	Refresh()
}

var _ AppChooserer = (*AppChooser)(nil)

func wrapAppChooser(obj *externglib.Object) *AppChooser {
	return &AppChooser{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalAppChooserer(p uintptr) (interface{}, error) {
	return wrapAppChooser(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AppInfo returns the currently selected application.
func (self *AppChooser) AppInfo() gio.AppInfor {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.GAppInfo      // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_get_app_info(_arg0)
	runtime.KeepAlive(self)

	var _appInfo gio.AppInfor // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(gio.AppInfor)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.AppInfor")
			}
			_appInfo = rv
		}
	}

	return _appInfo
}

// ContentType returns the content type for which the GtkAppChooser shows
// applications.
func (self *AppChooser) ContentType() string {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_get_content_type(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Refresh reloads the list of applications.
func (self *AppChooser) Refresh() {
	var _arg0 *C.GtkAppChooser // out

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	C.gtk_app_chooser_refresh(_arg0)
	runtime.KeepAlive(self)
}
