// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_tool_button_get_type()), F: marshalToolButton},
	})
}

// ToolButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToolButtonOverrider interface {
	Clicked()
}

// ToolButton are ToolItems containing buttons.
//
// Use gtk_tool_button_new() to create a new ToolButton.
//
// The label of a ToolButton is determined by the properties
// ToolButton:label-widget, ToolButton:label, and ToolButton:stock-id. If
// ToolButton:label-widget is non-nil, then that widget is used as the label.
// Otherwise, if ToolButton:label is non-nil, that string is used as the label.
// Otherwise, if ToolButton:stock-id is non-nil, the label is determined by the
// stock item. Otherwise, the button does not have a label.
//
// The icon of a ToolButton is determined by the properties
// ToolButton:icon-widget and ToolButton:stock-id. If ToolButton:icon-widget is
// non-nil, then that widget is used as the icon. Otherwise, if
// ToolButton:stock-id is non-nil, the icon is determined by the stock item.
// Otherwise, the button does not have a icon.
//
//
// CSS nodes
//
// GtkToolButton has a single CSS node with name toolbutton.
type ToolButton interface {
	gextras.Objector

	// IconName returns the name of the themed icon for the tool button, see
	// gtk_tool_button_set_icon_name().
	IconName() string
	// IconWidget: return the widget used as icon widget on @button. See
	// gtk_tool_button_set_icon_widget().
	IconWidget() *WidgetClass
	// Label returns the label used by the tool button, or nil if the tool
	// button doesn’t have a label. or uses a the label from a stock item. The
	// returned string is owned by GTK+, and must not be modified or freed.
	Label() string
	// LabelWidget returns the widget used as label on @button. See
	// gtk_tool_button_set_label_widget().
	LabelWidget() *WidgetClass
	// StockID returns the name of the stock item. See
	// gtk_tool_button_set_stock_id(). The returned string is owned by GTK+ and
	// must not be freed or modifed.
	//
	// Deprecated: Use gtk_tool_button_get_icon_name() instead.
	StockID() string
	// UseUnderline returns whether underscores in the label property are used
	// as mnemonics on menu items on the overflow menu. See
	// gtk_tool_button_set_use_underline().
	UseUnderline() bool
	// SetIconName sets the icon for the tool button from a named themed icon.
	// See the docs for IconTheme for more details. The ToolButton:icon-name
	// property only has an effect if not overridden by non-nil
	// ToolButton:label-widget, ToolButton:icon-widget and ToolButton:stock-id
	// properties.
	SetIconName(iconName string)
	// SetIconWidget sets @icon as the widget used as icon on @button. If
	// @icon_widget is nil the icon is determined by the ToolButton:stock-id
	// property. If the ToolButton:stock-id property is also nil, @button will
	// not have an icon.
	SetIconWidget(iconWidget Widget)
	// SetLabel sets @label as the label used for the tool button. The
	// ToolButton:label property only has an effect if not overridden by a
	// non-nil ToolButton:label-widget property. If both the
	// ToolButton:label-widget and ToolButton:label properties are nil, the
	// label is determined by the ToolButton:stock-id property. If the
	// ToolButton:stock-id property is also nil, @button will not have a label.
	SetLabel(label string)
	// SetLabelWidget sets @label_widget as the widget that will be used as the
	// label for @button. If @label_widget is nil the ToolButton:label property
	// is used as label. If ToolButton:label is also nil, the label in the stock
	// item determined by the ToolButton:stock-id property is used as label. If
	// ToolButton:stock-id is also nil, @button does not have a label.
	SetLabelWidget(labelWidget Widget)
	// SetStockID sets the name of the stock item. See
	// gtk_tool_button_new_from_stock(). The stock_id property only has an
	// effect if not overridden by non-nil ToolButton:label-widget and
	// ToolButton:icon-widget properties.
	//
	// Deprecated: Use gtk_tool_button_set_icon_name() instead.
	SetStockID(stockId string)
	// SetUseUnderline: if set, an underline in the label property indicates
	// that the next character should be used for the mnemonic accelerator key
	// in the overflow menu. For example, if the label property is “_Open” and
	// @use_underline is true, the label on the tool button will be “Open” and
	// the item on the overflow menu will have an underlined “O”.
	//
	// Labels shown on tool buttons never have mnemonics on them; this property
	// only affects the menu item on the overflow menu.
	SetUseUnderline(useUnderline bool)
}

// ToolButtonClass implements the ToolButton interface.
type ToolButtonClass struct {
	*externglib.Object
	ToolItemClass
	ActionableIface
	ActivatableIface
	BuildableIface
}

var _ ToolButton = (*ToolButtonClass)(nil)

func wrapToolButton(obj *externglib.Object) ToolButton {
	return &ToolButtonClass{
		Object: obj,
		ToolItemClass: ToolItemClass{
			Object: obj,
			BinClass: BinClass{
				Object: obj,
				ContainerClass: ContainerClass{
					Object: obj,
					WidgetClass: WidgetClass{
						Object: obj,
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						BuildableIface: BuildableIface{
							Object: obj,
						},
					},
					BuildableIface: BuildableIface{
						Object: obj,
					},
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
			},
			ActivatableIface: ActivatableIface{
				Object: obj,
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
		},
		ActionableIface: ActionableIface{
			Object: obj,
			WidgetClass: WidgetClass{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
			},
		},
		ActivatableIface: ActivatableIface{
			Object: obj,
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
	}
}

func marshalToolButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToolButton(obj), nil
}

// NewToolButton creates a new ToolButton using @icon_widget as contents and
// @label as label.
func NewToolButton(iconWidget Widget, label string) *ToolButtonClass {
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(iconWidget.Native()))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_tool_button_new(_arg1, _arg2)

	var _toolButton *ToolButtonClass // out

	_toolButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ToolButtonClass)

	return _toolButton
}

// NewToolButtonFromStock creates a new ToolButton containing the image and text
// from a stock item. Some stock ids have preprocessor macros like K_STOCK_OK
// and K_STOCK_APPLY.
//
// It is an error if @stock_id is not a name of a stock item.
//
// Deprecated: Use gtk_tool_button_new() together with
// gtk_image_new_from_icon_name() instead.
func NewToolButtonFromStock(stockId string) *ToolButtonClass {
	var _arg1 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_tool_button_new_from_stock(_arg1)

	var _toolButton *ToolButtonClass // out

	_toolButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ToolButtonClass)

	return _toolButton
}

// IconName returns the name of the themed icon for the tool button, see
// gtk_tool_button_set_icon_name().
func (button *ToolButtonClass) IconName() string {
	var _arg0 *C.GtkToolButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_tool_button_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IconWidget: return the widget used as icon widget on @button. See
// gtk_tool_button_set_icon_widget().
func (button *ToolButtonClass) IconWidget() *WidgetClass {
	var _arg0 *C.GtkToolButton // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_tool_button_get_icon_widget(_arg0)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// Label returns the label used by the tool button, or nil if the tool button
// doesn’t have a label. or uses a the label from a stock item. The returned
// string is owned by GTK+, and must not be modified or freed.
func (button *ToolButtonClass) Label() string {
	var _arg0 *C.GtkToolButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_tool_button_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// LabelWidget returns the widget used as label on @button. See
// gtk_tool_button_set_label_widget().
func (button *ToolButtonClass) LabelWidget() *WidgetClass {
	var _arg0 *C.GtkToolButton // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_tool_button_get_label_widget(_arg0)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// StockID returns the name of the stock item. See
// gtk_tool_button_set_stock_id(). The returned string is owned by GTK+ and must
// not be freed or modifed.
//
// Deprecated: Use gtk_tool_button_get_icon_name() instead.
func (button *ToolButtonClass) StockID() string {
	var _arg0 *C.GtkToolButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_tool_button_get_stock_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// UseUnderline returns whether underscores in the label property are used as
// mnemonics on menu items on the overflow menu. See
// gtk_tool_button_set_use_underline().
func (button *ToolButtonClass) UseUnderline() bool {
	var _arg0 *C.GtkToolButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_tool_button_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetIconName sets the icon for the tool button from a named themed icon. See
// the docs for IconTheme for more details. The ToolButton:icon-name property
// only has an effect if not overridden by non-nil ToolButton:label-widget,
// ToolButton:icon-widget and ToolButton:stock-id properties.
func (button *ToolButtonClass) SetIconName(iconName string) {
	var _arg0 *C.GtkToolButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_button_set_icon_name(_arg0, _arg1)
}

// SetIconWidget sets @icon as the widget used as icon on @button. If
// @icon_widget is nil the icon is determined by the ToolButton:stock-id
// property. If the ToolButton:stock-id property is also nil, @button will not
// have an icon.
func (button *ToolButtonClass) SetIconWidget(iconWidget Widget) {
	var _arg0 *C.GtkToolButton // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(iconWidget.Native()))

	C.gtk_tool_button_set_icon_widget(_arg0, _arg1)
}

// SetLabel sets @label as the label used for the tool button. The
// ToolButton:label property only has an effect if not overridden by a non-nil
// ToolButton:label-widget property. If both the ToolButton:label-widget and
// ToolButton:label properties are nil, the label is determined by the
// ToolButton:stock-id property. If the ToolButton:stock-id property is also
// nil, @button will not have a label.
func (button *ToolButtonClass) SetLabel(label string) {
	var _arg0 *C.GtkToolButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_button_set_label(_arg0, _arg1)
}

// SetLabelWidget sets @label_widget as the widget that will be used as the
// label for @button. If @label_widget is nil the ToolButton:label property is
// used as label. If ToolButton:label is also nil, the label in the stock item
// determined by the ToolButton:stock-id property is used as label. If
// ToolButton:stock-id is also nil, @button does not have a label.
func (button *ToolButtonClass) SetLabelWidget(labelWidget Widget) {
	var _arg0 *C.GtkToolButton // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_tool_button_set_label_widget(_arg0, _arg1)
}

// SetStockID sets the name of the stock item. See
// gtk_tool_button_new_from_stock(). The stock_id property only has an effect if
// not overridden by non-nil ToolButton:label-widget and ToolButton:icon-widget
// properties.
//
// Deprecated: Use gtk_tool_button_set_icon_name() instead.
func (button *ToolButtonClass) SetStockID(stockId string) {
	var _arg0 *C.GtkToolButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_button_set_stock_id(_arg0, _arg1)
}

// SetUseUnderline: if set, an underline in the label property indicates that
// the next character should be used for the mnemonic accelerator key in the
// overflow menu. For example, if the label property is “_Open” and
// @use_underline is true, the label on the tool button will be “Open” and the
// item on the overflow menu will have an underlined “O”.
//
// Labels shown on tool buttons never have mnemonics on them; this property only
// affects the menu item on the overflow menu.
func (button *ToolButtonClass) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkToolButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkToolButton)(unsafe.Pointer(button.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_tool_button_set_use_underline(_arg0, _arg1)
}
