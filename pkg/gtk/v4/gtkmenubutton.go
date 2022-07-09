// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_MenuButtonCreatePopupFunc(void*, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// GTypeMenuButton returns the GType for the type MenuButton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeMenuButton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "MenuButton").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalMenuButton)
	return gtype
}

// MenuButtonCreatePopupFunc: user-provided callback function to create a popup
// for a GtkMenuButton on demand.
//
// This function is called when the popup of menu_button is shown, but none has
// been provided via gtk.MenuButton.SetPopover() or
// gtk.MenuButton.SetMenuModel().
type MenuButtonCreatePopupFunc func(menuButton *MenuButton)

//export _gotk4_gtk4_MenuButtonCreatePopupFunc
func _gotk4_gtk4_MenuButtonCreatePopupFunc(arg1 *C.void, arg2 C.gpointer) {
	var fn MenuButtonCreatePopupFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(MenuButtonCreatePopupFunc)
	}

	var _menuButton *MenuButton // out

	_menuButton = wrapMenuButton(coreglib.Take(unsafe.Pointer(arg1)))

	fn(_menuButton)
}

// MenuButton: GtkMenuButton widget is used to display a popup when clicked.
//
// !An example GtkMenuButton (menu-button.png)
//
// This popup can be provided either as a GtkPopover or as an abstract
// GMenuModel.
//
// The GtkMenuButton widget can show either an icon (set with the
// gtk.MenuButton:icon-name property) or a label (set with the
// gtk.MenuButton:label property). If neither is explicitly set, a gtk.Image is
// automatically created, using an arrow image oriented according to
// gtk.MenuButton:direction or the generic “open-menu-symbolic” icon if the
// direction is not set.
//
// The positioning of the popup is determined by the gtk.MenuButton:direction
// property of the menu button.
//
// For menus, the gtk.Widget:halign and gtk.Widget:valign properties of the menu
// are also taken into account. For example, when the direction is
// GTK_ARROW_DOWN and the horizontal alignment is GTK_ALIGN_START, the menu will
// be positioned below the button, with the starting edge (depending on the text
// direction) of the menu aligned with the starting edge of the button. If there
// is not enough space below the button, the menu is popped up above the button
// instead. If the alignment would move part of the menu offscreen, it is
// “pushed in”.
//
// | | start | center | end | | - | --- | --- | --- | | **down** | !
// (down-start.png) | ! (down-center.png) | ! (down-end.png) | | **up** | !
// (up-start.png) | ! (up-center.png) | ! (up-end.png) | | **left** | !
// (left-start.png) | ! (left-center.png) | ! (left-end.png) | | **right** | !
// (right-start.png) | ! (right-center.png) | ! (right-end.png) |
//
// CSS nodes
//
//    menubutton
//    ╰── button.toggle
//        ╰── <content>
//             ╰── [arrow]
//
//
// GtkMenuButton has a single CSS node with name menubutton which contains a
// button node with a .toggle style class.
//
// Inside the toggle button content, there is an arrow node for the indicator,
// which will carry one of the .none, .up, .down, .left or .right style classes
// to indicate the direction that the menu will appear in. The CSS is expected
// to provide a suitable image for each of these cases using the
// -gtk-icon-source property.
//
// Optionally, the menubutton node can carry the .circular style class to
// request a round appearance.
//
//
// Accessibility
//
// GtkMenuButton uses the K_ACCESSIBLE_ROLE_BUTTON role.
type MenuButton struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*MenuButton)(nil)
)

func wrapMenuButton(obj *coreglib.Object) *MenuButton {
	return &MenuButton{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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

func marshalMenuButton(p uintptr) (interface{}, error) {
	return wrapMenuButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMenuButton creates a new GtkMenuButton widget with downwards-pointing
// arrow as the only child.
//
// You can replace the child widget with another GtkWidget should you wish to.
//
// The function returns the following values:
//
//    - menuButton: newly created GtkMenuButton.
//
func NewMenuButton() *MenuButton {
	_info := girepository.MustFind("Gtk", "MenuButton")
	_gret := _info.InvokeClassMethod("new_MenuButton", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _menuButton *MenuButton // out

	_menuButton = wrapMenuButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _menuButton
}

// HasFrame returns whether the button has a frame.
//
// The function returns the following values:
//
//    - ok: TRUE if the button has a frame.
//
func (menuButton *MenuButton) HasFrame() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_gret := _info.InvokeClassMethod("get_has_frame", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuButton)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IconName gets the name of the icon shown in the button.
//
// The function returns the following values:
//
//    - utf8: name of the icon shown in the button.
//
func (menuButton *MenuButton) IconName() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_gret := _info.InvokeClassMethod("get_icon_name", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Label gets the label shown in the button.
//
// The function returns the following values:
//
//    - utf8: label shown in the button.
//
func (menuButton *MenuButton) Label() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_gret := _info.InvokeClassMethod("get_label", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// MenuModel returns the GMenuModel used to generate the popup.
//
// The function returns the following values:
//
//    - menuModel (optional): GMenuModel or NULL.
//
func (menuButton *MenuButton) MenuModel() gio.MenuModeller {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_gret := _info.InvokeClassMethod("get_menu_model", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuButton)

	var _menuModel gio.MenuModeller // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.MenuModeller)
				return ok
			})
			rv, ok := casted.(gio.MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// Popover returns the GtkPopover that pops out of the button.
//
// If the button is not using a GtkPopover, this function returns NULL.
//
// The function returns the following values:
//
//    - popover (optional): GtkPopover or NULL.
//
func (menuButton *MenuButton) Popover() *Popover {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_gret := _info.InvokeClassMethod("get_popover", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuButton)

	var _popover *Popover // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_popover = wrapPopover(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _popover
}

// UseUnderline returns whether an embedded underline in the text indicates a
// mnemonic.
//
// The function returns the following values:
//
//    - ok: TRUE whether an embedded underline in the text indicates the mnemonic
//      accelerator keys.
//
func (menuButton *MenuButton) UseUnderline() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_gret := _info.InvokeClassMethod("get_use_underline", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(menuButton)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Popdown dismiss the menu.
func (menuButton *MenuButton) Popdown() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("popdown", _args[:], nil)

	runtime.KeepAlive(menuButton)
}

// Popup: pop up the menu.
func (menuButton *MenuButton) Popup() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("popup", _args[:], nil)

	runtime.KeepAlive(menuButton)
}

// SetCreatePopupFunc sets func to be called when a popup is about to be shown.
//
// func should use one of
//
//    - gtk.MenuButton.SetPopover()
//    - gtk.MenuButton.SetMenuModel()
//
// to set a popup for menu_button. If func is non-NULL, menu_button will always
// be sensitive.
//
// Using this function will not reset the menu widget attached to menu_button.
// Instead, this can be done manually in func.
//
// The function takes the following parameters:
//
//    - fn (optional): function to call when a popup is about to be shown, but
//      none has been provided via other means, or NULL to reset to default
//      behavior.
//
func (menuButton *MenuButton) SetCreatePopupFunc(fn MenuButtonCreatePopupFunc) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))
	if fn != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk4_MenuButtonCreatePopupFunc)
		_args[2] = C.gpointer(gbox.Assign(fn))
		_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("set_create_popup_func", _args[:], nil)

	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(fn)
}

// SetHasFrame sets the style of the button.
//
// The function takes the following parameters:
//
//    - hasFrame: whether the button should have a visible frame.
//
func (menuButton *MenuButton) SetHasFrame(hasFrame bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))
	if hasFrame {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("set_has_frame", _args[:], nil)

	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(hasFrame)
}

// SetIconName sets the name of an icon to show inside the menu button.
//
// The function takes the following parameters:
//
//    - iconName: icon name.
//
func (menuButton *MenuButton) SetIconName(iconName string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("set_icon_name", _args[:], nil)

	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(iconName)
}

// SetLabel sets the label to show inside the menu button.
//
// The function takes the following parameters:
//
//    - label: label.
//
func (menuButton *MenuButton) SetLabel(label string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("set_label", _args[:], nil)

	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(label)
}

// SetMenuModel sets the GMenuModel from which the popup will be constructed.
//
// If menu_model is NULL, the button is disabled.
//
// A gtk.Popover will be created from the menu model with
// gtk.PopoverMenu.NewFromModel. Actions will be connected as documented for
// this function.
//
// If gtk.MenuButton:popover is already set, it will be dissociated from the
// menu_button, and the property is set to NULL.
//
// The function takes the following parameters:
//
//    - menuModel (optional): GMenuModel, or NULL to unset and disable the
//      button.
//
func (menuButton *MenuButton) SetMenuModel(menuModel gio.MenuModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))
	if menuModel != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuModel).Native()))
	}

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("set_menu_model", _args[:], nil)

	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(menuModel)
}

// SetPopover sets the GtkPopover that will be popped up when the menu_button is
// clicked.
//
// If popover is NULL, the button is disabled.
//
// If gtk.MenuButton:menu-model is set, the menu model is dissociated from the
// menu_button, and the property is set to NULL.
//
// The function takes the following parameters:
//
//    - popover (optional): GtkPopover, or NULL to unset and disable the button.
//
func (menuButton *MenuButton) SetPopover(popover Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))
	if popover != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	}

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("set_popover", _args[:], nil)

	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(popover)
}

// SetUseUnderline: if true, an underline in the text indicates a mnemonic.
//
// The function takes the following parameters:
//
//    - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (menuButton *MenuButton) SetUseUnderline(useUnderline bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuButton).Native()))
	if useUnderline {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "MenuButton")
	_info.InvokeClassMethod("set_use_underline", _args[:], nil)

	runtime.KeepAlive(menuButton)
	runtime.KeepAlive(useUnderline)
}
