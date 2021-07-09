// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scale_button_get_type()), F: marshalScaleButton},
	})
}

// ScaleButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScaleButtonOverrider interface {
	ValueChanged(value float64)
}

// ScaleButton: `GtkScaleButton` provides a button which pops up a scale widget.
//
// This kind of widget is commonly used for volume controls in multimedia
// applications, and GTK provides a [class@Gtk.VolumeButton] subclass that is
// tailored for this use case.
//
//
// CSS nodes
//
// `GtkScaleButton` has a single CSS node with name button. To differentiate it
// from a plain `GtkButton`, it gets the .scale style class.
type ScaleButton interface {
	gextras.Objector

	// Adjustment gets the `GtkAdjustment` associated with the
	// `GtkScaleButton`’s scale.
	//
	// See [method@Gtk.Range.get_adjustment] for details.
	Adjustment() *AdjustmentClass
	// MinusButton retrieves the minus button of the `GtkScaleButton`.
	MinusButton() *ButtonClass
	// PlusButton retrieves the plus button of the `GtkScaleButton.`
	PlusButton() *ButtonClass
	// Popup retrieves the popup of the `GtkScaleButton`.
	Popup() *WidgetClass
	// Value gets the current value of the scale button.
	Value() float64
	// SetAdjustment sets the `GtkAdjustment` to be used as a model for the
	// `GtkScaleButton`’s scale.
	//
	// See [method@Gtk.Range.set_adjustment] for details.
	SetAdjustment(adjustment Adjustment)
	// SetIcons sets the icons to be used by the scale button.
	SetIcons(icons []string)
	// SetValue sets the current value of the scale.
	//
	// If the value is outside the minimum or maximum range values, it will be
	// clamped to fit inside them.
	//
	// The scale button emits the [signal@Gtk.ScaleButton::value-changed] signal
	// if the value changes.
	SetValue(value float64)
}

// ScaleButtonClass implements the ScaleButton interface.
type ScaleButtonClass struct {
	*externglib.Object
	WidgetClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
	OrientableInterface
}

var _ ScaleButton = (*ScaleButtonClass)(nil)

func wrapScaleButton(obj *externglib.Object) ScaleButton {
	return &ScaleButtonClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalScaleButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScaleButton(obj), nil
}

// NewScaleButton creates a `GtkScaleButton`.
//
// The new scale button has a range between @min and @max, with a stepping of
// @step.
func NewScaleButton(min float64, max float64, step float64, icons []string) *ScaleButtonClass {
	var _arg1 C.double // out
	var _arg2 C.double // out
	var _arg3 C.double // out
	var _arg4 **C.char
	var _cret *C.GtkWidget // in

	_arg1 = C.double(min)
	_arg2 = C.double(max)
	_arg3 = C.double(step)
	_arg4 = (**C.char)(C.malloc(C.ulong(len(icons)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice(_arg4, len(icons))
		for i := range icons {
			out[i] = (*C.char)(C.CString(icons[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	_cret = C.gtk_scale_button_new(_arg1, _arg2, _arg3, _arg4)

	var _scaleButton *ScaleButtonClass // out

	_scaleButton = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ScaleButtonClass)

	return _scaleButton
}

// Adjustment gets the `GtkAdjustment` associated with the `GtkScaleButton`’s
// scale.
//
// See [method@Gtk.Range.get_adjustment] for details.
func (b *ScaleButtonClass) Adjustment() *AdjustmentClass {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer((&ScaleButton).Native()))

	_cret = C.gtk_scale_button_get_adjustment(_arg0)

	var _adjustment *AdjustmentClass // out

	_adjustment = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*AdjustmentClass)

	return _adjustment
}

// MinusButton retrieves the minus button of the `GtkScaleButton`.
func (b *ScaleButtonClass) MinusButton() *ButtonClass {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer((&ScaleButton).Native()))

	_cret = C.gtk_scale_button_get_minus_button(_arg0)

	var _ret *ButtonClass // out

	_ret = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ButtonClass)

	return _ret
}

// PlusButton retrieves the plus button of the `GtkScaleButton.`
func (b *ScaleButtonClass) PlusButton() *ButtonClass {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer((&ScaleButton).Native()))

	_cret = C.gtk_scale_button_get_plus_button(_arg0)

	var _ret *ButtonClass // out

	_ret = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ButtonClass)

	return _ret
}

// Popup retrieves the popup of the `GtkScaleButton`.
func (b *ScaleButtonClass) Popup() *WidgetClass {
	var _arg0 *C.GtkScaleButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer((&ScaleButton).Native()))

	_cret = C.gtk_scale_button_get_popup(_arg0)

	var _widget *WidgetClass // out

	_widget = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*WidgetClass)

	return _widget
}

// Value gets the current value of the scale button.
func (b *ScaleButtonClass) Value() float64 {
	var _arg0 *C.GtkScaleButton // out
	var _cret C.double          // in

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer((&ScaleButton).Native()))

	_cret = C.gtk_scale_button_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the `GtkAdjustment` to be used as a model for the
// `GtkScaleButton`’s scale.
//
// See [method@Gtk.Range.set_adjustment] for details.
func (b *ScaleButtonClass) SetAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 *C.GtkAdjustment  // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer((&ScaleButton).Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((&Adjustment).Native()))

	C.gtk_scale_button_set_adjustment(_arg0, _arg1)
}

// SetIcons sets the icons to be used by the scale button.
func (b *ScaleButtonClass) SetIcons(icons []string) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 **C.char

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer((&ScaleButton).Native()))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(icons)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(icons))
		for i := range icons {
			out[i] = (*C.char)(C.CString(icons[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_scale_button_set_icons(_arg0, _arg1)
}

// SetValue sets the current value of the scale.
//
// If the value is outside the minimum or maximum range values, it will be
// clamped to fit inside them.
//
// The scale button emits the [signal@Gtk.ScaleButton::value-changed] signal if
// the value changes.
func (b *ScaleButtonClass) SetValue(value float64) {
	var _arg0 *C.GtkScaleButton // out
	var _arg1 C.double          // out

	_arg0 = (*C.GtkScaleButton)(unsafe.Pointer((&ScaleButton).Native()))
	_arg1 = C.double(value)

	C.gtk_scale_button_set_value(_arg0, _arg1)
}
