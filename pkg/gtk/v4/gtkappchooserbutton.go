// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_AppChooserButton_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk4_AppChooserButton_ConnectCustomItemActivated(gpointer, void*, guintptr);
import "C"

// GTypeAppChooserButton returns the GType for the type AppChooserButton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAppChooserButton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "AppChooserButton").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAppChooserButton)
	return gtype
}

// AppChooserButton: GtkAppChooserButton lets the user select an application.
//
// !An example GtkAppChooserButton (appchooserbutton.png)
//
// Initially, a GtkAppChooserButton selects the first application in its list,
// which will either be the most-recently used application or, if
// gtk.AppChooserButton:show-default-item is TRUE, the default application.
//
// The list of applications shown in a GtkAppChooserButton includes the
// recommended applications for the given content type. When
// gtk.AppChooserButton:show-default-item is set, the default application is
// also included. To let the user chooser other applications, you can set the
// gtk.AppChooserButton:show-dialog-item property, which allows to open a full
// gtk.AppChooserDialog.
//
// It is possible to add custom items to the list, using
// gtk.AppChooserButton.AppendCustomItem(). These items cause the
// gtk.AppChooserButton::custom-item-activated signal to be emitted when they
// are selected.
//
// To track changes in the selected application, use the
// gtk.AppChooserButton::changed signal.
//
//
// CSS nodes
//
// GtkAppChooserButton has a single CSS node with the name “appchooserbutton”.
type AppChooserButton struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	AppChooser
}

var (
	_ Widgetter         = (*AppChooserButton)(nil)
	_ coreglib.Objector = (*AppChooserButton)(nil)
)

func wrapAppChooserButton(obj *coreglib.Object) *AppChooserButton {
	return &AppChooserButton{
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
		Object: obj,
		AppChooser: AppChooser{
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
		},
	}
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	return wrapAppChooserButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_AppChooserButton_ConnectChanged
func _gotk4_gtk4_AppChooserButton_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectChanged is emitted when the active application changes.
func (self *AppChooserButton) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "changed", false, unsafe.Pointer(C._gotk4_gtk4_AppChooserButton_ConnectChanged), f)
}

//export _gotk4_gtk4_AppChooserButton_ConnectCustomItemActivated
func _gotk4_gtk4_AppChooserButton_ConnectCustomItemActivated(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(itemName string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(itemName string))
	}

	var _itemName string // out

	_itemName = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_itemName)
}

// ConnectCustomItemActivated is emitted when a custom item is activated.
//
// Use gtk.AppChooserButton.AppendCustomItem(), to add custom items.
func (self *AppChooserButton) ConnectCustomItemActivated(f func(itemName string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "custom-item-activated", false, unsafe.Pointer(C._gotk4_gtk4_AppChooserButton_ConnectCustomItemActivated), f)
}

// NewAppChooserButton creates a new GtkAppChooserButton for applications that
// can handle content of the given type.
//
// The function takes the following parameters:
//
//    - contentType: content type to show applications for.
//
// The function returns the following values:
//
//    - appChooserButton: newly created GtkAppChooserButton.
//
func NewAppChooserButton(contentType string) *AppChooserButton {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_args[0]))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_gret := _info.InvokeClassMethod("new_AppChooserButton", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(contentType)

	var _appChooserButton *AppChooserButton // out

	_appChooserButton = wrapAppChooserButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _appChooserButton
}

// AppendCustomItem appends a custom item to the list of applications that is
// shown in the popup.
//
// The item name must be unique per-widget. Clients can use the provided name as
// a detail for the gtk.AppChooserButton::custom-item-activated signal, to add a
// callback for the activation of a particular custom item in the list.
//
// See also gtk.AppChooserButton.AppendSeparator().
//
// The function takes the following parameters:
//
//    - name of the custom item.
//    - label for the custom item.
//    - icon for the custom item.
//
func (self *AppChooserButton) AppendCustomItem(name, label string, icon gio.Iconner) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_args[2]))
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_info.InvokeClassMethod("append_custom_item", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
	runtime.KeepAlive(label)
	runtime.KeepAlive(icon)
}

// AppendSeparator appends a separator to the list of applications that is shown
// in the popup.
func (self *AppChooserButton) AppendSeparator() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_info.InvokeClassMethod("append_separator", _args[:], nil)

	runtime.KeepAlive(self)
}

// Heading returns the text to display at the top of the dialog.
//
// The function returns the following values:
//
//    - utf8 (optional): text to display at the top of the dialog, or NULL, in
//      which case a default text is displayed.
//
func (self *AppChooserButton) Heading() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_gret := _info.InvokeClassMethod("get_heading", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Modal gets whether the dialog is modal.
//
// The function returns the following values:
//
//    - ok: TRUE if the dialog is modal.
//
func (self *AppChooserButton) Modal() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_gret := _info.InvokeClassMethod("get_modal", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowDefaultItem returns whether the dropdown menu should show the default
// application at the top.
//
// The function returns the following values:
//
//    - ok: value of gtk.AppChooserButton:show-default-item.
//
func (self *AppChooserButton) ShowDefaultItem() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_gret := _info.InvokeClassMethod("get_show_default_item", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowDialogItem returns whether the dropdown menu shows an item for a
// GtkAppChooserDialog.
//
// The function returns the following values:
//
//    - ok: value of gtk.AppChooserButton:show-dialog-item.
//
func (self *AppChooserButton) ShowDialogItem() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_gret := _info.InvokeClassMethod("get_show_dialog_item", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveCustomItem selects a custom item.
//
// See gtk.AppChooserButton.AppendCustomItem().
//
// Use gtk.AppChooser.Refresh() to bring the selection to its initial state.
//
// The function takes the following parameters:
//
//    - name of the custom item.
//
func (self *AppChooserButton) SetActiveCustomItem(name string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_info.InvokeClassMethod("set_active_custom_item", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetHeading sets the text to display at the top of the dialog.
//
// If the heading is not set, the dialog displays a default text.
//
// The function takes the following parameters:
//
//    - heading: string containing Pango markup.
//
func (self *AppChooserButton) SetHeading(heading string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(heading)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_info.InvokeClassMethod("set_heading", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(heading)
}

// SetModal sets whether the dialog should be modal.
//
// The function takes the following parameters:
//
//    - modal: TRUE to make the dialog modal.
//
func (self *AppChooserButton) SetModal(modal bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if modal {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_info.InvokeClassMethod("set_modal", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(modal)
}

// SetShowDefaultItem sets whether the dropdown menu of this button should show
// the default application for the given content type at top.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserButton:show-default-item.
//
func (self *AppChooserButton) SetShowDefaultItem(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_info.InvokeClassMethod("set_show_default_item", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetShowDialogItem sets whether the dropdown menu of this button should show
// an entry to trigger a GtkAppChooserDialog.
//
// The function takes the following parameters:
//
//    - setting: new value for gtk.AppChooserButton:show-dialog-item.
//
func (self *AppChooserButton) SetShowDialogItem(setting bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if setting {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "AppChooserButton")
	_info.InvokeClassMethod("set_show_dialog_item", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}
