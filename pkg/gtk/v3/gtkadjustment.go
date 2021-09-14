// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_adjustment_get_type()), F: marshalAdjustmenter},
	})
}

// AdjustmentOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AdjustmentOverrider interface {
	// Changed emits a Adjustment::changed signal from the Adjustment. This is
	// typically called by the owner of the Adjustment after it has changed any
	// of the Adjustment properties other than the value.
	//
	// Deprecated: GTK+ emits Adjustment::changed itself whenever any of the
	// properties (other than value) change.
	Changed()
	// ValueChanged emits a Adjustment::value-changed signal from the
	// Adjustment. This is typically called by the owner of the Adjustment after
	// it has changed the Adjustment:value property.
	//
	// Deprecated: GTK+ emits Adjustment::value-changed itself whenever the
	// value changes.
	ValueChanged()
}

// Adjustment object represents a value which has an associated lower and upper
// bound, together with step and page increments, and a page size. It is used
// within several GTK+ widgets, including SpinButton, Viewport, and Range (which
// is a base class for Scrollbar and Scale).
//
// The Adjustment object does not update the value itself. Instead it is left up
// to the owner of the Adjustment to control the value.
type Adjustment struct {
	externglib.InitiallyUnowned
}

func wrapAdjustment(obj *externglib.Object) *Adjustment {
	return &Adjustment{
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalAdjustmenter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAdjustment(obj), nil
}

// NewAdjustment creates a new Adjustment.
func NewAdjustment(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) *Adjustment {
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out
	var _arg3 C.gdouble        // out
	var _arg4 C.gdouble        // out
	var _arg5 C.gdouble        // out
	var _arg6 C.gdouble        // out
	var _cret *C.GtkAdjustment // in

	_arg1 = C.gdouble(value)
	_arg2 = C.gdouble(lower)
	_arg3 = C.gdouble(upper)
	_arg4 = C.gdouble(stepIncrement)
	_arg5 = C.gdouble(pageIncrement)
	_arg6 = C.gdouble(pageSize)

	_cret = C.gtk_adjustment_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(value)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
	runtime.KeepAlive(stepIncrement)
	runtime.KeepAlive(pageIncrement)
	runtime.KeepAlive(pageSize)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// Changed emits a Adjustment::changed signal from the Adjustment. This is
// typically called by the owner of the Adjustment after it has changed any of
// the Adjustment properties other than the value.
//
// Deprecated: GTK+ emits Adjustment::changed itself whenever any of the
// properties (other than value) change.
func (adjustment *Adjustment) Changed() {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_adjustment_changed(_arg0)
	runtime.KeepAlive(adjustment)
}

// ClampPage updates the Adjustment:value property to ensure that the range
// between lower and upper is in the current page (i.e. between Adjustment:value
// and Adjustment:value + Adjustment:page-size). If the range is larger than the
// page size, then only the start of it will be in the current page.
//
// A Adjustment::value-changed signal will be emitted if the value is changed.
func (adjustment *Adjustment) ClampPage(lower float64, upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.gdouble(lower)
	_arg2 = C.gdouble(upper)

	C.gtk_adjustment_clamp_page(_arg0, _arg1, _arg2)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
}

// Configure sets all properties of the adjustment at once.
//
// Use this function to avoid multiple emissions of the Adjustment::changed
// signal. See gtk_adjustment_set_lower() for an alternative way of compressing
// multiple emissions of Adjustment::changed into one.
func (adjustment *Adjustment) Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out
	var _arg2 C.gdouble        // out
	var _arg3 C.gdouble        // out
	var _arg4 C.gdouble        // out
	var _arg5 C.gdouble        // out
	var _arg6 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.gdouble(value)
	_arg2 = C.gdouble(lower)
	_arg3 = C.gdouble(upper)
	_arg4 = C.gdouble(stepIncrement)
	_arg5 = C.gdouble(pageIncrement)
	_arg6 = C.gdouble(pageSize)

	C.gtk_adjustment_configure(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(value)
	runtime.KeepAlive(lower)
	runtime.KeepAlive(upper)
	runtime.KeepAlive(stepIncrement)
	runtime.KeepAlive(pageIncrement)
	runtime.KeepAlive(pageSize)
}

// Lower retrieves the minimum value of the adjustment.
func (adjustment *Adjustment) Lower() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_lower(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MinimumIncrement gets the smaller of step increment and page increment.
func (adjustment *Adjustment) MinimumIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_minimum_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageIncrement retrieves the page increment of the adjustment.
func (adjustment *Adjustment) PageIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_page_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageSize retrieves the page size of the adjustment.
func (adjustment *Adjustment) PageSize() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_page_size(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// StepIncrement retrieves the step increment of the adjustment.
func (adjustment *Adjustment) StepIncrement() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_step_increment(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Upper retrieves the maximum value of the adjustment.
func (adjustment *Adjustment) Upper() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_upper(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Value gets the current value of the adjustment. See
// gtk_adjustment_set_value().
func (adjustment *Adjustment) Value() float64 {
	var _arg0 *C.GtkAdjustment // out
	var _cret C.gdouble        // in

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_adjustment_get_value(_arg0)
	runtime.KeepAlive(adjustment)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetLower sets the minimum value of the adjustment.
//
// When setting multiple adjustment properties via their individual setters,
// multiple Adjustment::changed signals will be emitted. However, since the
// emission of the Adjustment::changed signal is tied to the emission of the
// #GObject::notify signals of the changed properties, it’s possible to compress
// the Adjustment::changed signals into one by calling g_object_freeze_notify()
// and g_object_thaw_notify() around the calls to the individual setters.
//
// Alternatively, using a single g_object_set() for all the properties to
// change, or using gtk_adjustment_configure() has the same effect of
// compressing Adjustment::changed emissions.
func (adjustment *Adjustment) SetLower(lower float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.gdouble(lower)

	C.gtk_adjustment_set_lower(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(lower)
}

// SetPageIncrement sets the page increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
func (adjustment *Adjustment) SetPageIncrement(pageIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.gdouble(pageIncrement)

	C.gtk_adjustment_set_page_increment(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(pageIncrement)
}

// SetPageSize sets the page size of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the GtkAdjustment::changed signal when setting multiple adjustment
// properties.
func (adjustment *Adjustment) SetPageSize(pageSize float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.gdouble(pageSize)

	C.gtk_adjustment_set_page_size(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(pageSize)
}

// SetStepIncrement sets the step increment of the adjustment.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
func (adjustment *Adjustment) SetStepIncrement(stepIncrement float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.gdouble(stepIncrement)

	C.gtk_adjustment_set_step_increment(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(stepIncrement)
}

// SetUpper sets the maximum value of the adjustment.
//
// Note that values will be restricted by upper - page-size if the page-size
// property is nonzero.
//
// See gtk_adjustment_set_lower() about how to compress multiple emissions of
// the Adjustment::changed signal when setting multiple adjustment properties.
func (adjustment *Adjustment) SetUpper(upper float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.gdouble(upper)

	C.gtk_adjustment_set_upper(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(upper)
}

// SetValue sets the Adjustment value. The value is clamped to lie between
// Adjustment:lower and Adjustment:upper.
//
// Note that for adjustments which are used in a Scrollbar, the effective range
// of allowed values goes from Adjustment:lower to Adjustment:upper -
// Adjustment:page-size.
func (adjustment *Adjustment) SetValue(value float64) {
	var _arg0 *C.GtkAdjustment // out
	var _arg1 C.gdouble        // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_adjustment_set_value(_arg0, _arg1)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(value)
}

// ValueChanged emits a Adjustment::value-changed signal from the Adjustment.
// This is typically called by the owner of the Adjustment after it has changed
// the Adjustment:value property.
//
// Deprecated: GTK+ emits Adjustment::value-changed itself whenever the value
// changes.
func (adjustment *Adjustment) ValueChanged() {
	var _arg0 *C.GtkAdjustment // out

	_arg0 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_adjustment_value_changed(_arg0)
	runtime.KeepAlive(adjustment)
}
