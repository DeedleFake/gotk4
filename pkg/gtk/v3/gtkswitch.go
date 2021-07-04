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
		{T: externglib.Type(C.gtk_switch_get_type()), F: marshalSwitch},
	})
}

// Switch is a widget that has two states: on or off. The user can control which
// state should be active by clicking the empty area, or by dragging the handle.
//
// GtkSwitch can also handle situations where the underlying state changes with
// a delay. See Switch::state-set for details.
//
// CSS nodes
//
//    switch
//    ╰── slider
//
// GtkSwitch has two css nodes, the main node with the name switch and a subnode
// named slider. Neither of them is using any style classes.
type Switch interface {
	Actionable
	Activatable

	Active() bool

	State() bool

	SetActiveSwitch(isActive bool)

	SetStateSwitch(state bool)
}

// _switch implements the Switch class.
type _switch struct {
	Widget
}

// WrapSwitch wraps a GObject to the right type. It is
// primarily used internally.
func WrapSwitch(obj *externglib.Object) Switch {
	return _switch{
		Widget: WrapWidget(obj),
	}
}

func marshalSwitch(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSwitch(obj), nil
}

func NewSwitch() Switch {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_switch_new()

	var __switch Switch // out

	__switch = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Switch)

	return __switch
}

func (s _switch) Active() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_switch_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s _switch) State() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_switch_get_state(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s _switch) SetActiveSwitch(isActive bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_active(_arg0, _arg1)
}

func (s _switch) SetStateSwitch(state bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(s.Native()))
	if state {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_state(_arg0, _arg1)
}

func (a _switch) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a _switch) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a _switch) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a _switch) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a _switch) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (b _switch) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b _switch) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b _switch) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b _switch) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b _switch) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b _switch) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b _switch) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b _switch) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b _switch) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b _switch) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a _switch) DoSetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).DoSetRelatedAction(action)
}

func (a _switch) RelatedAction() Action {
	return WrapActivatable(gextras.InternObject(a)).RelatedAction()
}

func (a _switch) UseActionAppearance() bool {
	return WrapActivatable(gextras.InternObject(a)).UseActionAppearance()
}

func (a _switch) SetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).SetRelatedAction(action)
}

func (a _switch) SetUseActionAppearance(useAppearance bool) {
	WrapActivatable(gextras.InternObject(a)).SetUseActionAppearance(useAppearance)
}

func (a _switch) SyncActionProperties(action Action) {
	WrapActivatable(gextras.InternObject(a)).SyncActionProperties(action)
}
