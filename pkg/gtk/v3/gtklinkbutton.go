// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_link_button_get_type()), F: marshalLinkButton},
	})
}

// LinkButton: a GtkLinkButton is a Button with a hyperlink, similar to the one
// used by web browsers, which triggers an action when clicked. It is useful to
// show quick links to resources.
//
// A link button is created by calling either gtk_link_button_new() or
// gtk_link_button_new_with_label(). If using the former, the URI you pass to
// the constructor is used as a label for the widget.
//
// The URI bound to a GtkLinkButton can be set specifically using
// gtk_link_button_set_uri(), and retrieved using gtk_link_button_get_uri().
//
// By default, GtkLinkButton calls gtk_show_uri_on_window() when the button is
// clicked. This behaviour can be overridden by connecting to the
// LinkButton::activate-link signal and returning true from the signal handler.
//
//
// CSS nodes
//
// GtkLinkButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .link style class.
type LinkButton interface {
	Button

	URI() string

	Visited() bool

	SetURILinkButton(uri string)

	SetVisitedLinkButton(visited bool)
}

// linkButton implements the LinkButton class.
type linkButton struct {
	Button
}

// WrapLinkButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapLinkButton(obj *externglib.Object) LinkButton {
	return linkButton{
		Button: WrapButton(obj),
	}
}

func marshalLinkButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLinkButton(obj), nil
}

func NewLinkButton(uri string) LinkButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_link_button_new(_arg1)

	var _linkButton LinkButton // out

	_linkButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(LinkButton)

	return _linkButton
}

func NewLinkButtonWithLabel(uri string, label string) LinkButton {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_link_button_new_with_label(_arg1, _arg2)

	var _linkButton LinkButton // out

	_linkButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(LinkButton)

	return _linkButton
}

func (l linkButton) URI() string {
	var _arg0 *C.GtkLinkButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_link_button_get_uri(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (l linkButton) Visited() bool {
	var _arg0 *C.GtkLinkButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_link_button_get_visited(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (l linkButton) SetURILinkButton(uri string) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_link_button_set_uri(_arg0, _arg1)
}

func (l linkButton) SetVisitedLinkButton(visited bool) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))
	if visited {
		_arg1 = C.TRUE
	}

	C.gtk_link_button_set_visited(_arg0, _arg1)
}

func (b linkButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b linkButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b linkButton) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b linkButton) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b linkButton) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b linkButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b linkButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b linkButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b linkButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b linkButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a linkButton) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a linkButton) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a linkButton) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a linkButton) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a linkButton) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (b linkButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b linkButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b linkButton) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b linkButton) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b linkButton) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b linkButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b linkButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b linkButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b linkButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b linkButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a linkButton) DoSetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).DoSetRelatedAction(action)
}

func (a linkButton) RelatedAction() Action {
	return WrapActivatable(gextras.InternObject(a)).RelatedAction()
}

func (a linkButton) UseActionAppearance() bool {
	return WrapActivatable(gextras.InternObject(a)).UseActionAppearance()
}

func (a linkButton) SetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).SetRelatedAction(action)
}

func (a linkButton) SetUseActionAppearance(useAppearance bool) {
	WrapActivatable(gextras.InternObject(a)).SetUseActionAppearance(useAppearance)
}

func (a linkButton) SyncActionProperties(action Action) {
	WrapActivatable(gextras.InternObject(a)).SyncActionProperties(action)
}
