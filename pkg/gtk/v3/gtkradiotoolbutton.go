// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_tool_button_get_type()), F: marshalRadioToolButtonner},
	})
}

// RadioToolButton is a ToolItem that contains a radio button, that is, a button
// that is part of a group of toggle buttons where only one button can be active
// at a time.
//
// Use gtk_radio_tool_button_new() to create a new GtkRadioToolButton. Use
// gtk_radio_tool_button_new_from_widget() to create a new GtkRadioToolButton
// that is part of the same group as an existing GtkRadioToolButton.
//
//
// CSS nodes
//
// GtkRadioToolButton has a single CSS node with name toolbutton.
type RadioToolButton struct {
	ToggleToolButton
}

var (
	_ externglib.Objector = (*RadioToolButton)(nil)
	_ Binner              = (*RadioToolButton)(nil)
)

func wrapRadioToolButton(obj *externglib.Object) *RadioToolButton {
	return &RadioToolButton{
		ToggleToolButton: ToggleToolButton{
			ToolButton: ToolButton{
				ToolItem: ToolItem{
					Bin: Bin{
						Container: Container{
							Widget: Widget{
								InitiallyUnowned: externglib.InitiallyUnowned{
									Object: obj,
								},
								Object: obj,
								ImplementorIface: atk.ImplementorIface{
									Object: obj,
								},
								Buildable: Buildable{
									Object: obj,
								},
							},
						},
					},
					Object: obj,
					Activatable: Activatable{
						Object: obj,
					},
				},
				Object: obj,
				Actionable: Actionable{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalRadioToolButtonner(p uintptr) (interface{}, error) {
	return wrapRadioToolButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRadioToolButton creates a new RadioToolButton, adding it to group.
//
// The function takes the following parameters:
//
//    - group: an existing radio button group, or NULL if you are creating a
//    new group.
//
func NewRadioToolButton(group []RadioButton) *RadioToolButton {
	var _arg1 *C.GSList      // out
	var _cret *C.GtkToolItem // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioButton // out
			dst = (*C.GtkRadioButton)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}

	_cret = C.gtk_radio_tool_button_new(_arg1)
	runtime.KeepAlive(group)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// NewRadioToolButtonFromStock creates a new RadioToolButton, adding it to
// group. The new RadioToolButton will contain an icon and label from the stock
// item indicated by stock_id.
//
// Deprecated: Use gtk_radio_tool_button_new() instead.
//
// The function takes the following parameters:
//
//    - group: existing radio button group, or NULL if you are creating a new
//    group.
//    - stockId: name of a stock item.
//
func NewRadioToolButtonFromStock(group []RadioButton, stockId string) *RadioToolButton {
	var _arg1 *C.GSList      // out
	var _arg2 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioButton // out
			dst = (*C.GtkRadioButton)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_tool_button_new_from_stock(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(stockId)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// NewRadioToolButtonFromWidget creates a new RadioToolButton adding it to the
// same group as gruup.
//
// The function takes the following parameters:
//
//    - group: existing RadioToolButton, or NULL.
//
func NewRadioToolButtonFromWidget(group *RadioToolButton) *RadioToolButton {
	var _arg1 *C.GtkRadioToolButton // out
	var _cret *C.GtkToolItem        // in

	if group != nil {
		_arg1 = (*C.GtkRadioToolButton)(unsafe.Pointer(group.Native()))
	}

	_cret = C.gtk_radio_tool_button_new_from_widget(_arg1)
	runtime.KeepAlive(group)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// NewRadioToolButtonWithStockFromWidget creates a new RadioToolButton adding it
// to the same group as group. The new RadioToolButton will contain an icon and
// label from the stock item indicated by stock_id.
//
// Deprecated: gtk_radio_tool_button_new_from_widget.
//
// The function takes the following parameters:
//
//    - group: existing RadioToolButton.
//    - stockId: name of a stock item.
//
func NewRadioToolButtonWithStockFromWidget(group *RadioToolButton, stockId string) *RadioToolButton {
	var _arg1 *C.GtkRadioToolButton // out
	var _arg2 *C.gchar              // out
	var _cret *C.GtkToolItem        // in

	if group != nil {
		_arg1 = (*C.GtkRadioToolButton)(unsafe.Pointer(group.Native()))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_tool_button_new_with_stock_from_widget(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(stockId)

	var _radioToolButton *RadioToolButton // out

	_radioToolButton = wrapRadioToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioToolButton
}

// Group returns the radio button group button belongs to.
func (button *RadioToolButton) Group() *gextras.SList[RadioButton] {
	var _arg0 *C.GtkRadioToolButton // out
	var _cret *C.GSList             // in

	_arg0 = (*C.GtkRadioToolButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_radio_tool_button_get_group(_arg0)
	runtime.KeepAlive(button)

	var _sList *gextras.SList[RadioButton] // out

	_sList = gextras.NewSList[RadioButton](
		unsafe.Pointer(_cret),
		gextras.ListOpts[RadioButton]{
			Convert: func(ptr unsafe.Pointer) RadioButton {
				src := *(**C.GtkRadioButton)(ptr)
				var dst RadioButton // out
				dst = *wrapRadioButton(externglib.Take(unsafe.Pointer(src)))
				return dst
			},
		},
		false,
	)

	return _sList
}

// SetGroup adds button to group, removing it from the group it belonged to
// before.
//
// The function takes the following parameters:
//
//    - group: existing radio button group, or NULL.
//
func (button *RadioToolButton) SetGroup(group []RadioButton) {
	var _arg0 *C.GtkRadioToolButton // out
	var _arg1 *C.GSList             // out

	_arg0 = (*C.GtkRadioToolButton)(unsafe.Pointer(button.Native()))
	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioButton // out
			dst = (*C.GtkRadioButton)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}

	C.gtk_radio_tool_button_set_group(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(group)
}
