// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_FontButton_ConnectFontSet(gpointer, guintptr);
import "C"

// GTypeFontButton returns the GType for the type FontButton.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFontButton() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "FontButton").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFontButton)
	return gtype
}

// FontButton: GtkFontButton allows to open a font chooser dialog to change the
// font.
//
// !An example GtkFontButton (font-button.png)
//
// It is suitable widget for selecting a font in a preference dialog.
//
// CSS nodes
//
//    fontbutton
//    ╰── button.font
//        ╰── [content]
//
//
// GtkFontButton has a single CSS node with name fontbutton which contains a
// button node with the .font style class.
type FontButton struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	FontChooser
}

var (
	_ Widgetter         = (*FontButton)(nil)
	_ coreglib.Objector = (*FontButton)(nil)
)

func wrapFontButton(obj *coreglib.Object) *FontButton {
	return &FontButton{
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
		FontChooser: FontChooser{
			Object: obj,
		},
	}
}

func marshalFontButton(p uintptr) (interface{}, error) {
	return wrapFontButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_FontButton_ConnectFontSet
func _gotk4_gtk4_FontButton_ConnectFontSet(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectFontSet is emitted when the user selects a font.
//
// When handling this signal, use gtk.FontChooser.GetFont() to find out which
// font was just selected.
//
// Note that this signal is only emitted when the user changes the font. If you
// need to react to programmatic font changes as well, use the notify::font
// signal.
func (fontButton *FontButton) ConnectFontSet(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(fontButton, "font-set", false, unsafe.Pointer(C._gotk4_gtk4_FontButton_ConnectFontSet), f)
}

// NewFontButton creates a new font picker widget.
//
// The function returns the following values:
//
//    - fontButton: new font picker widget.
//
func NewFontButton() *FontButton {
	_info := girepository.MustFind("Gtk", "FontButton")
	_gret := _info.InvokeClassMethod("new_FontButton", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _fontButton *FontButton // out

	_fontButton = wrapFontButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _fontButton
}

// NewFontButtonWithFont creates a new font picker widget showing the given
// font.
//
// The function takes the following parameters:
//
//    - fontname: name of font to display in font chooser dialog.
//
// The function returns the following values:
//
//    - fontButton: new font picker widget.
//
func NewFontButtonWithFont(fontname string) *FontButton {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(fontname)))
	defer C.free(unsafe.Pointer(_args[0]))

	_info := girepository.MustFind("Gtk", "FontButton")
	_gret := _info.InvokeClassMethod("new_FontButton_with_font", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontname)

	var _fontButton *FontButton // out

	_fontButton = wrapFontButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _fontButton
}

// Modal gets whether the dialog is modal.
//
// The function returns the following values:
//
//    - ok: TRUE if the dialog is modal.
//
func (fontButton *FontButton) Modal() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_info := girepository.MustFind("Gtk", "FontButton")
	_gret := _info.InvokeClassMethod("get_modal", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the font chooser dialog.
//
// The function returns the following values:
//
//    - utf8: internal copy of the title string which must not be freed.
//
func (fontButton *FontButton) Title() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_info := girepository.MustFind("Gtk", "FontButton")
	_gret := _info.InvokeClassMethod("get_title", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseFont returns whether the selected font is used in the label.
//
// The function returns the following values:
//
//    - ok: whether the selected font is used in the label.
//
func (fontButton *FontButton) UseFont() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_info := girepository.MustFind("Gtk", "FontButton")
	_gret := _info.InvokeClassMethod("get_use_font", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// UseSize returns whether the selected size is used in the label.
//
// The function returns the following values:
//
//    - ok: whether the selected size is used in the label.
//
func (fontButton *FontButton) UseSize() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))

	_info := girepository.MustFind("Gtk", "FontButton")
	_gret := _info.InvokeClassMethod("get_use_size", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(fontButton)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetModal sets whether the dialog should be modal.
//
// The function takes the following parameters:
//
//    - modal: TRUE to make the dialog modal.
//
func (fontButton *FontButton) SetModal(modal bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	if modal {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "FontButton")
	_info.InvokeClassMethod("set_modal", _args[:], nil)

	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(modal)
}

// SetTitle sets the title for the font chooser dialog.
//
// The function takes the following parameters:
//
//    - title: string containing the font chooser dialog title.
//
func (fontButton *FontButton) SetTitle(title string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "FontButton")
	_info.InvokeClassMethod("set_title", _args[:], nil)

	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(title)
}

// SetUseFont: if use_font is TRUE, the font name will be written using the
// selected font.
//
// The function takes the following parameters:
//
//    - useFont: if TRUE, font name will be written using font chosen.
//
func (fontButton *FontButton) SetUseFont(useFont bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	if useFont {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "FontButton")
	_info.InvokeClassMethod("set_use_font", _args[:], nil)

	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(useFont)
}

// SetUseSize: if use_size is TRUE, the font name will be written using the
// selected size.
//
// The function takes the following parameters:
//
//    - useSize: if TRUE, font name will be written using the selected size.
//
func (fontButton *FontButton) SetUseSize(useSize bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fontButton).Native()))
	if useSize {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "FontButton")
	_info.InvokeClassMethod("set_use_size", _args[:], nil)

	runtime.KeepAlive(fontButton)
	runtime.KeepAlive(useSize)
}
