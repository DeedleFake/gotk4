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
		{T: externglib.Type(C.gtk_action_get_type()), F: marshalAction},
	})
}

// Action: > In GTK+ 3.10, GtkAction has been deprecated. Use #GAction >
// instead, and associate actions with Actionable widgets. Use > Model for
// creating menus with gtk_menu_new_from_model().
//
// Actions represent operations that the user can be perform, along with some
// information how it should be presented in the interface. Each action provides
// methods to create icons, menu items and toolbar items representing itself.
//
// As well as the callback that is called when the action gets activated, the
// following also gets associated with the action:
//
// - a name (not translated, for path lookup)
//
// - a label (translated, for display)
//
// - an accelerator
//
// - whether label indicates a stock id
//
// - a tooltip (optional, translated)
//
// - a toolbar label (optional, shorter than label)
//
//    The action will also have some state information:
//
// - visible (shown/hidden)
//
// - sensitive (enabled/disabled)
//
// Apart from regular actions, there are [toggle actions][GtkToggleAction],
// which can be toggled between two states and [radio actions][GtkRadioAction],
// of which only one in a group can be in the “active” state. Other actions can
// be implemented as Action subclasses.
//
// Each action can have one or more proxy widgets. To act as an action proxy,
// widget needs to implement Activatable interface. Proxies mirror the state of
// the action and should change when the action’s state changes. Properties that
// are always mirrored by proxies are Action:sensitive and Action:visible.
// Action:gicon, Action:icon-name, Action:label, Action:short-label and
// Action:stock-id properties are only mirorred if proxy widget has
// Activatable:use-action-appearance property set to true.
//
// When the proxy is activated, it should activate its action.
type Action interface {
	Buildable

	ActivateAction()

	BlockActivateAction()

	ConnectAcceleratorAction()

	CreateIconAction(iconSize int) Widget

	CreateMenuAction() Widget

	CreateMenuItemAction() Widget

	CreateToolItemAction() Widget

	DisconnectAcceleratorAction()

	AccelPath() string

	AlwaysShowImage() bool

	IconName() string

	IsImportant() bool

	Label() string

	GetName() string

	Sensitive() bool

	ShortLabel() string

	StockID() string

	Tooltip() string

	Visible() bool

	VisibleHorizontal() bool

	VisibleVertical() bool

	IsSensitiveAction() bool

	IsVisibleAction() bool

	SetAccelGroupAction(accelGroup AccelGroup)

	SetAccelPathAction(accelPath string)

	SetAlwaysShowImageAction(alwaysShow bool)

	SetIconNameAction(iconName string)

	SetIsImportantAction(isImportant bool)

	SetLabelAction(label string)

	SetSensitiveAction(sensitive bool)

	SetShortLabelAction(shortLabel string)

	SetStockIDAction(stockId string)

	SetTooltipAction(tooltip string)

	SetVisibleAction(visible bool)

	SetVisibleHorizontalAction(visibleHorizontal bool)

	SetVisibleVerticalAction(visibleVertical bool)

	UnblockActivateAction()
}

// action implements the Action class.
type action struct {
	gextras.Objector
}

// WrapAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapAction(obj *externglib.Object) Action {
	return action{
		Objector: obj,
	}
}

func marshalAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAction(obj), nil
}

func NewAction(name string, label string, tooltip string, stockId string) Action {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out
	var _arg4 *C.gchar     // out
	var _cret *C.GtkAction // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))

	_cret = C.gtk_action_new(_arg1, _arg2, _arg3, _arg4)

	var _action Action // out

	_action = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Action)

	return _action
}

func (a action) ActivateAction() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_activate(_arg0)
}

func (a action) BlockActivateAction() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_block_activate(_arg0)
}

func (a action) ConnectAcceleratorAction() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_connect_accelerator(_arg0)
}

func (a action) CreateIconAction(iconSize int) Widget {
	var _arg0 *C.GtkAction  // out
	var _arg1 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	_cret = C.gtk_action_create_icon(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a action) CreateMenuAction() Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_create_menu(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a action) CreateMenuItemAction() Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_create_menu_item(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a action) CreateToolItemAction() Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_create_tool_item(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a action) DisconnectAcceleratorAction() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_disconnect_accelerator(_arg0)
}

func (a action) AccelPath() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_accel_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) AlwaysShowImage() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_always_show_image(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) IconName() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) IsImportant() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_is_important(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) Label() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) GetName() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) Sensitive() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) ShortLabel() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_short_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) StockID() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_stock_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) Tooltip() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_tooltip(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a action) Visible() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) VisibleHorizontal() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_visible_horizontal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) VisibleVertical() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_get_visible_vertical(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) IsSensitiveAction() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_is_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) IsVisibleAction() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_is_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a action) SetAccelGroupAction(accelGroup AccelGroup) {
	var _arg0 *C.GtkAction     // out
	var _arg1 *C.GtkAccelGroup // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_action_set_accel_group(_arg0, _arg1)
}

func (a action) SetAccelPathAction(accelPath string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_accel_path(_arg0, _arg1)
}

func (a action) SetAlwaysShowImageAction(alwaysShow bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if alwaysShow {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_always_show_image(_arg0, _arg1)
}

func (a action) SetIconNameAction(iconName string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_icon_name(_arg0, _arg1)
}

func (a action) SetIsImportantAction(isImportant bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if isImportant {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_is_important(_arg0, _arg1)
}

func (a action) SetLabelAction(label string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_label(_arg0, _arg1)
}

func (a action) SetSensitiveAction(sensitive bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_sensitive(_arg0, _arg1)
}

func (a action) SetShortLabelAction(shortLabel string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(shortLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_short_label(_arg0, _arg1)
}

func (a action) SetStockIDAction(stockId string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_stock_id(_arg0, _arg1)
}

func (a action) SetTooltipAction(tooltip string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_tooltip(_arg0, _arg1)
}

func (a action) SetVisibleAction(visible bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible(_arg0, _arg1)
}

func (a action) SetVisibleHorizontalAction(visibleHorizontal bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if visibleHorizontal {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible_horizontal(_arg0, _arg1)
}

func (a action) SetVisibleVerticalAction(visibleVertical bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))
	if visibleVertical {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible_vertical(_arg0, _arg1)
}

func (a action) UnblockActivateAction() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(a.Native()))

	C.gtk_action_unblock_activate(_arg0)
}

func (b action) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b action) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b action) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b action) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b action) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b action) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b action) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b action) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b action) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b action) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
