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
		{T: externglib.Type(C.gtk_app_chooser_widget_get_type()), F: marshalAppChooserWidget},
	})
}

// AppChooserWidget: `GtkAppChooserWidget` is a widget for selecting
// applications.
//
// It is the main building block for [class@Gtk.AppChooserDialog]. Most
// applications only need to use the latter; but you can use this widget as part
// of a larger widget if you have special needs.
//
// `GtkAppChooserWidget` offers detailed control over what applications are
// shown, using the [property@Gtk.AppChooserWidget:show-default],
// [property@Gtk.AppChooserWidget:show-recommended],
// [property@Gtk.AppChooserWidget:show-fallback],
// [property@Gtk.AppChooserWidget:show-other] and
// [property@Gtk.AppChooserWidget:show-all] properties. See the
// [iface@Gtk.AppChooser] documentation for more information about these groups
// of applications.
//
// To keep track of the selected application, use the
// [signal@Gtk.AppChooserWidget::application-selected] and
// [signal@Gtk.AppChooserWidget::application-activated] signals.
//
//
// CSS nodes
//
// `GtkAppChooserWidget` has a single CSS node with name appchooser.
type AppChooserWidget interface {
	AppChooser

	DefaultText() string

	ShowAll() bool

	ShowDefault() bool

	ShowFallback() bool

	ShowOther() bool

	ShowRecommended() bool

	SetDefaultTextAppChooserWidget(text string)

	SetShowAllAppChooserWidget(setting bool)

	SetShowDefaultAppChooserWidget(setting bool)

	SetShowFallbackAppChooserWidget(setting bool)

	SetShowOtherAppChooserWidget(setting bool)

	SetShowRecommendedAppChooserWidget(setting bool)
}

// appChooserWidget implements the AppChooserWidget class.
type appChooserWidget struct {
	Widget
}

// WrapAppChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserWidget(obj *externglib.Object) AppChooserWidget {
	return appChooserWidget{
		Widget: WrapWidget(obj),
	}
}

func marshalAppChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserWidget(obj), nil
}

func NewAppChooserWidget(contentType string) AppChooserWidget {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_app_chooser_widget_new(_arg1)

	var _appChooserWidget AppChooserWidget // out

	_appChooserWidget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(AppChooserWidget)

	return _appChooserWidget
}

func (s appChooserWidget) DefaultText() string {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_widget_get_default_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s appChooserWidget) ShowAll() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_all(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s appChooserWidget) ShowDefault() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s appChooserWidget) ShowFallback() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_fallback(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s appChooserWidget) ShowOther() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_other(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s appChooserWidget) ShowRecommended() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_recommended(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s appChooserWidget) SetDefaultTextAppChooserWidget(text string) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_widget_set_default_text(_arg0, _arg1)
}

func (s appChooserWidget) SetShowAllAppChooserWidget(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_all(_arg0, _arg1)
}

func (s appChooserWidget) SetShowDefaultAppChooserWidget(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_default(_arg0, _arg1)
}

func (s appChooserWidget) SetShowFallbackAppChooserWidget(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_fallback(_arg0, _arg1)
}

func (s appChooserWidget) SetShowOtherAppChooserWidget(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_other(_arg0, _arg1)
}

func (s appChooserWidget) SetShowRecommendedAppChooserWidget(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_recommended(_arg0, _arg1)
}

func (s appChooserWidget) ContentType() string {
	return WrapAppChooser(gextras.InternObject(s)).ContentType()
}

func (s appChooserWidget) Refresh() {
	WrapAppChooser(gextras.InternObject(s)).Refresh()
}

func (s appChooserWidget) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s appChooserWidget) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s appChooserWidget) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s appChooserWidget) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s appChooserWidget) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s appChooserWidget) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s appChooserWidget) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b appChooserWidget) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
