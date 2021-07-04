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
		{T: externglib.Type(C.gtk_menu_bar_get_type()), F: marshalMenuBar},
	})
}

// MenuBar: the MenuBar is a subclass of MenuShell which contains one or more
// MenuItems. The result is a standard menu bar which can hold many menu items.
//
//
// CSS nodes
//
// GtkMenuBar has a single CSS node with name menubar.
type MenuBar interface {
	MenuShell

	ChildPackDirection() PackDirection

	PackDirection() PackDirection

	SetChildPackDirectionMenuBar(childPackDir PackDirection)

	SetPackDirectionMenuBar(packDir PackDirection)
}

// menuBar implements the MenuBar class.
type menuBar struct {
	MenuShell
}

// WrapMenuBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuBar(obj *externglib.Object) MenuBar {
	return menuBar{
		MenuShell: WrapMenuShell(obj),
	}
}

func marshalMenuBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuBar(obj), nil
}

func NewMenuBar() MenuBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_bar_new()

	var _menuBar MenuBar // out

	_menuBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MenuBar)

	return _menuBar
}

func NewMenuBarFromModel(model gio.MenuModel) MenuBar {
	var _arg1 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_menu_bar_new_from_model(_arg1)

	var _menuBar MenuBar // out

	_menuBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MenuBar)

	return _menuBar
}

func (m menuBar) ChildPackDirection() PackDirection {
	var _arg0 *C.GtkMenuBar      // out
	var _cret C.GtkPackDirection // in

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_bar_get_child_pack_direction(_arg0)

	var _packDirection PackDirection // out

	_packDirection = PackDirection(_cret)

	return _packDirection
}

func (m menuBar) PackDirection() PackDirection {
	var _arg0 *C.GtkMenuBar      // out
	var _cret C.GtkPackDirection // in

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_bar_get_pack_direction(_arg0)

	var _packDirection PackDirection // out

	_packDirection = PackDirection(_cret)

	return _packDirection
}

func (m menuBar) SetChildPackDirectionMenuBar(childPackDir PackDirection) {
	var _arg0 *C.GtkMenuBar      // out
	var _arg1 C.GtkPackDirection // out

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(m.Native()))
	_arg1 = C.GtkPackDirection(childPackDir)

	C.gtk_menu_bar_set_child_pack_direction(_arg0, _arg1)
}

func (m menuBar) SetPackDirectionMenuBar(packDir PackDirection) {
	var _arg0 *C.GtkMenuBar      // out
	var _arg1 C.GtkPackDirection // out

	_arg0 = (*C.GtkMenuBar)(unsafe.Pointer(m.Native()))
	_arg1 = C.GtkPackDirection(packDir)

	C.gtk_menu_bar_set_pack_direction(_arg0, _arg1)
}

func (b menuBar) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b menuBar) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b menuBar) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b menuBar) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b menuBar) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b menuBar) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b menuBar) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b menuBar) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b menuBar) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b menuBar) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
