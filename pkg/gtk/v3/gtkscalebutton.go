// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scale_button_get_type()), F: marshalScaleButtoner},
	})
}

// ScaleButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScaleButtonOverrider interface {
	ValueChanged(value float64)
}

// ScaleButton provides a button which pops up a scale widget. This kind of
// widget is commonly used for volume controls in multimedia applications, and
// GTK+ provides a VolumeButton subclass that is tailored for this use case.
//
//
// CSS nodes
//
// GtkScaleButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .scale style class.
//
// The popup widget that contains the scale has a .scale-popup style class.
type ScaleButton struct {
	Button

	Orientable
}

var _ gextras.Nativer = (*ScaleButton)(nil)

func wrapScaleButton(obj *externglib.Object) *ScaleButton {
	return &ScaleButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalScaleButtoner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScaleButton(obj), nil
}

// NewScaleButton creates a ScaleButton, with a range between min and max, with
// a stepping of step.
func NewScaleButton(size int, min float64, max float64, step float64, icons []string) *ScaleButton {
	var _arg1 C.GtkIconSize // out
	var _arg2 C.gdouble     // out
	var _arg3 C.gdouble     // out
	var _arg4 C.gdouble     // out
	var _arg5 **C.gchar
	var _cret *C.GtkWidget // in

	_arg1 = C.GtkIconSize(size)
	_arg2 = C.gdouble(min)
	_arg3 = C.gdouble(max)
	_arg4 = C.gdouble(step)
	{
		_arg5 = (**C.gchar)(C.malloc(C.ulong(len(icons)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(_arg5, len(icons)+1)
			var zero *C.gchar
			out[len(icons)] = zero
			for i := range icons {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(icons[i])))
			}
		}
	}

	_cret = C.gtk_scale_button_new(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _scaleButton *ScaleButton // out

	_scaleButton = wrapScaleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _scaleButton
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *ScaleButton) Native() uintptr {
	return v.Button.Bin.Container.Widget.InitiallyUnowned.Object.Native()
}

// Adjustment gets the Adjustment associated with the ScaleButton’s scale. See
// gtk_range_get_adjustment() for details.
func (button *ScaleButton) Adjustment() *Adjustment {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_adjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// MinusButton retrieves the minus button of the ScaleButton.
func (button *ScaleButton) MinusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_minus_button(_arg0)

	var _ret *Button // out

	_ret = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// PlusButton retrieves the plus button of the ScaleButton.
func (button *ScaleButton) PlusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_plus_button(_arg0)

	var _ret *Button // out

	_ret = wrapButton(externglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// Popup retrieves the popup of the ScaleButton.
func (button *ScaleButton) Popup() Widgeter {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_popup(_arg0)

	var _widget Widgeter // out

	_widget = (*gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgeter)

	return _widget
}

// Value gets the current value of the scale button.
func (button *ScaleButton) Value() float64 {
	var _arg0 *C.GtkScaleButton // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_scale_button_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the Adjustment to be used as a model for the ScaleButton’s
// scale. See gtk_range_set_adjustment() for details.
func (button *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 *C.GtkAdjustment  // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_scale_button_set_adjustment(_arg0, _arg1)
}

// SetIcons sets the icons to be used by the scale button. For details, see the
// ScaleButton:icons property.
func (button *ScaleButton) SetIcons(icons []string) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 **C.gchar

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	{
		_arg1 = (**C.gchar)(C.malloc(C.ulong(len(icons)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(_arg1, len(icons)+1)
			var zero *C.gchar
			out[len(icons)] = zero
			for i := range icons {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(icons[i])))
			}
		}
	}

	C.gtk_scale_button_set_icons(_arg0, _arg1)
}

// SetValue sets the current value of the scale; if the value is outside the
// minimum or maximum range values, it will be clamped to fit inside them. The
// scale button emits the ScaleButton::value-changed signal if the value
// changes.
func (button *ScaleButton) SetValue(value float64) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_scale_button_set_value(_arg0, _arg1)
}
