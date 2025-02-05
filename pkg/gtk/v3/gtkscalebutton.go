// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_ScaleButton_ConnectValueChanged(gpointer, gdouble, guintptr);
// extern void _gotk4_gtk3_ScaleButton_ConnectPopup(gpointer, guintptr);
// extern void _gotk4_gtk3_ScaleButton_ConnectPopdown(gpointer, guintptr);
// extern void _gotk4_gtk3_ScaleButtonClass_value_changed(GtkScaleButton*, gdouble);
// void _gotk4_gtk3_ScaleButton_virtual_value_changed(void* fnptr, GtkScaleButton* arg0, gdouble arg1) {
//   ((void (*)(GtkScaleButton*, gdouble))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeScaleButton = coreglib.Type(C.gtk_scale_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScaleButton, F: marshalScaleButton},
	})
}

// ScaleButtonOverrides contains methods that are overridable.
type ScaleButtonOverrides struct {
	// The function takes the following parameters:
	//
	ValueChanged func(value float64)
}

func defaultScaleButtonOverrides(v *ScaleButton) ScaleButtonOverrides {
	return ScaleButtonOverrides{
		ValueChanged: v.valueChanged,
	}
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
	_ [0]func() // equal guard
	Button

	*coreglib.Object
	Orientable
}

var (
	_ coreglib.Objector = (*ScaleButton)(nil)
	_ Binner            = (*ScaleButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ScaleButton, *ScaleButtonClass, ScaleButtonOverrides](
		GTypeScaleButton,
		initScaleButtonClass,
		wrapScaleButton,
		defaultScaleButtonOverrides,
	)
}

func initScaleButtonClass(gclass unsafe.Pointer, overrides ScaleButtonOverrides, classInitFunc func(*ScaleButtonClass)) {
	pclass := (*C.GtkScaleButtonClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeScaleButton))))

	if overrides.ValueChanged != nil {
		pclass.value_changed = (*[0]byte)(C._gotk4_gtk3_ScaleButtonClass_value_changed)
	}

	if classInitFunc != nil {
		class := (*ScaleButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapScaleButton(obj *coreglib.Object) *ScaleButton {
	return &ScaleButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
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
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
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
			Activatable: Activatable{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalScaleButton(p uintptr) (interface{}, error) {
	return wrapScaleButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectPopdown signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popdown the scale widget.
//
// The default binding for this signal is Escape.
func (button *ScaleButton) ConnectPopdown(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "popdown", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectPopdown), f)
}

// ConnectPopup signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted to popup the scale widget.
//
// The default bindings for this signal are Space, Enter and Return.
func (button *ScaleButton) ConnectPopup(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "popup", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectPopup), f)
}

// ConnectValueChanged signal is emitted when the value field has changed.
func (button *ScaleButton) ConnectValueChanged(f func(value float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "value-changed", false, unsafe.Pointer(C._gotk4_gtk3_ScaleButton_ConnectValueChanged), f)
}

// NewScaleButton creates a ScaleButton, with a range between min and max, with
// a stepping of step.
//
// The function takes the following parameters:
//
//    - size: stock icon size (IconSize).
//    - min: minimum value of the scale (usually 0).
//    - max: maximum value of the scale (usually 100).
//    - step: stepping of value when a scroll-wheel event, or up/down arrow event
//      occurs (usually 2).
//    - icons (optional): NULL-terminated array of icon names, or NULL if you
//      want to set the list later with gtk_scale_button_set_icons().
//
// The function returns the following values:
//
//    - scaleButton: new ScaleButton.
//
func NewScaleButton(size int, min, max, step float64, icons []string) *ScaleButton {
	var _arg1 C.GtkIconSize // out
	var _arg2 C.gdouble     // out
	var _arg3 C.gdouble     // out
	var _arg4 C.gdouble     // out
	var _arg5 **C.gchar     // out
	var _cret *C.GtkWidget  // in

	_arg1 = C.GtkIconSize(size)
	_arg2 = C.gdouble(min)
	_arg3 = C.gdouble(max)
	_arg4 = C.gdouble(step)
	{
		_arg5 = (**C.gchar)(C.calloc(C.size_t((len(icons) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg5))
		{
			out := unsafe.Slice(_arg5, len(icons)+1)
			var zero *C.gchar
			out[len(icons)] = zero
			for i := range icons {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(icons[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gtk_scale_button_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(size)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)
	runtime.KeepAlive(icons)

	var _scaleButton *ScaleButton // out

	_scaleButton = wrapScaleButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _scaleButton
}

// Adjustment gets the Adjustment associated with the ScaleButton’s scale. See
// gtk_range_get_adjustment() for details.
//
// The function returns the following values:
//
//    - adjustment associated with the scale.
//
func (button *ScaleButton) Adjustment() *Adjustment {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_adjustment(_arg0)
	runtime.KeepAlive(button)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// MinusButton retrieves the minus button of the ScaleButton.
//
// The function returns the following values:
//
//    - ret minus button of the ScaleButton as a Button.
//
func (button *ScaleButton) MinusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_minus_button(_arg0)
	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// PlusButton retrieves the plus button of the ScaleButton.
//
// The function returns the following values:
//
//    - ret plus button of the ScaleButton as a Button.
//
func (button *ScaleButton) PlusButton() *Button {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_plus_button(_arg0)
	runtime.KeepAlive(button)

	var _ret *Button // out

	_ret = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _ret
}

// Popup retrieves the popup of the ScaleButton.
//
// The function returns the following values:
//
//    - widget: popup of the ScaleButton.
//
func (button *ScaleButton) Popup() Widgetter {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_popup(_arg0)
	runtime.KeepAlive(button)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// Value gets the current value of the scale button.
//
// The function returns the following values:
//
//    - gdouble: current value of the scale button.
//
func (button *ScaleButton) Value() float64 {
	var _arg0 *C.GtkScaleButton // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_scale_button_get_value(_arg0)
	runtime.KeepAlive(button)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the Adjustment to be used as a model for the ScaleButton’s
// scale. See gtk_range_set_adjustment() for details.
//
// The function takes the following parameters:
//
//    - adjustment: Adjustment.
//
func (button *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 *C.GtkAdjustment  // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))

	C.gtk_scale_button_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(adjustment)
}

// SetIcons sets the icons to be used by the scale button. For details, see the
// ScaleButton:icons property.
//
// The function takes the following parameters:
//
//    - icons: NULL-terminated array of icon names.
//
func (button *ScaleButton) SetIcons(icons []string) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 **C.gchar         // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(icons) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(icons)+1)
			var zero *C.gchar
			out[len(icons)] = zero
			for i := range icons {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(icons[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_scale_button_set_icons(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(icons)
}

// SetValue sets the current value of the scale; if the value is outside the
// minimum or maximum range values, it will be clamped to fit inside them. The
// scale button emits the ScaleButton::value-changed signal if the value
// changes.
//
// The function takes the following parameters:
//
//    - value: new value of the scale button.
//
func (button *ScaleButton) SetValue(value float64) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = C.gdouble(value)

	C.gtk_scale_button_set_value(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(value)
}

// The function takes the following parameters:
//
func (button *ScaleButton) valueChanged(value float64) {
	gclass := (*C.GtkScaleButtonClass)(coreglib.PeekParentClass(button))
	fnarg := gclass.value_changed

	var _arg0 *C.GtkScaleButton // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = C.gdouble(value)

	C._gotk4_gtk3_ScaleButton_virtual_value_changed(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(value)
}

// ScaleButtonClass: instance of this type is always passed by reference.
type ScaleButtonClass struct {
	*scaleButtonClass
}

// scaleButtonClass is the struct that's finalized.
type scaleButtonClass struct {
	native *C.GtkScaleButtonClass
}

func (s *ScaleButtonClass) ParentClass() *ButtonClass {
	valptr := &s.native.parent_class
	var _v *ButtonClass // out
	_v = (*ButtonClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
