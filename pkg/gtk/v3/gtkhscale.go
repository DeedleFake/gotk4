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
import "C"

// GType values.
var (
	GTypeHScale = coreglib.Type(C.gtk_hscale_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHScale, F: marshalHScale},
	})
}

// HScaleOverrides contains methods that are overridable.
type HScaleOverrides struct {
}

func defaultHScaleOverrides(v *HScale) HScaleOverrides {
	return HScaleOverrides{}
}

// HScale widget is used to allow the user to select a value using a horizontal
// slider. To create one, use gtk_hscale_new_with_range().
//
// The position to show the current value, and the number of decimal places
// shown can be set using the parent Scale class’s functions.
//
// GtkHScale has been deprecated, use Scale instead.
type HScale struct {
	_ [0]func() // equal guard
	Scale
}

var (
	_ Ranger = (*HScale)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*HScale, *HScaleClass, HScaleOverrides](
		GTypeHScale,
		initHScaleClass,
		wrapHScale,
		defaultHScaleOverrides,
	)
}

func initHScaleClass(gclass unsafe.Pointer, overrides HScaleOverrides, classInitFunc func(*HScaleClass)) {
	if classInitFunc != nil {
		class := (*HScaleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapHScale(obj *coreglib.Object) *HScale {
	return &HScale{
		Scale: Scale{
			Range: Range{
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
				Object: obj,
				Orientable: Orientable{
					Object: obj,
				},
			},
		},
	}
}

func marshalHScale(p uintptr) (interface{}, error) {
	return wrapHScale(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHScale creates a new HScale.
//
// Deprecated: Use gtk_scale_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function takes the following parameters:
//
//    - adjustment (optional) which sets the range of the scale.
//
// The function returns the following values:
//
//    - hScale: new HScale.
//
func NewHScale(adjustment *Adjustment) *HScale {
	var _arg1 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_cret = C.gtk_hscale_new(_arg1)
	runtime.KeepAlive(adjustment)

	var _hScale *HScale // out

	_hScale = wrapHScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _hScale
}

// NewHScaleWithRange creates a new horizontal scale widget that lets the user
// input a number between min and max (including min and max) with the increment
// step. step must be nonzero; it’s the distance the slider moves when using the
// arrow keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk_scale_set_digits() to correct it.
//
// Deprecated: Use gtk_scale_new_with_range() with GTK_ORIENTATION_HORIZONTAL
// instead.
//
// The function takes the following parameters:
//
//    - min: minimum value.
//    - max: maximum value.
//    - step increment (tick size) used with keyboard shortcuts.
//
// The function returns the following values:
//
//    - hScale: new HScale.
//
func NewHScaleWithRange(min, max, step float64) *HScale {
	var _arg1 C.gdouble    // out
	var _arg2 C.gdouble    // out
	var _arg3 C.gdouble    // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)
	_arg3 = C.gdouble(step)

	_cret = C.gtk_hscale_new_with_range(_arg1, _arg2, _arg3)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _hScale *HScale // out

	_hScale = wrapHScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _hScale
}

// HScaleClass: instance of this type is always passed by reference.
type HScaleClass struct {
	*hScaleClass
}

// hScaleClass is the struct that's finalized.
type hScaleClass struct {
	native *C.GtkHScaleClass
}

func (h *HScaleClass) ParentClass() *ScaleClass {
	valptr := &h.native.parent_class
	var _v *ScaleClass // out
	_v = (*ScaleClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
