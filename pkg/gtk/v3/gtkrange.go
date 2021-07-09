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
		{T: externglib.Type(C.gtk_range_get_type()), F: marshalRange},
	})
}

// RangeOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RangeOverrider interface {
	AdjustBounds(newValue float64)
	RangeBorder(border_ *Border)
	ValueChanged()
}

// Range is the common base class for widgets which visualize an adjustment, e.g
// Scale or Scrollbar.
//
// Apart from signals for monitoring the parameters of the adjustment, Range
// provides properties and methods for influencing the sensitivity of the
// “steppers”. It also provides properties and methods for setting a “fill
// level” on range widgets. See gtk_range_set_fill_level().
type Range interface {
	gextras.Objector

	// Adjustment: get the Adjustment which is the “model” object for Range. See
	// gtk_range_set_adjustment() for details. The return value does not have a
	// reference added, so should not be unreferenced.
	Adjustment() *AdjustmentClass
	// FillLevel gets the current position of the fill level indicator.
	FillLevel() float64
	// Flippable gets the value set by gtk_range_set_flippable().
	Flippable() bool
	// Inverted gets the value set by gtk_range_set_inverted().
	Inverted() bool
	// LowerStepperSensitivity gets the sensitivity policy for the stepper that
	// points to the 'lower' end of the GtkRange’s adjustment.
	LowerStepperSensitivity() SensitivityType
	// MinSliderSize: this function is useful mainly for Range subclasses.
	//
	// See gtk_range_set_min_slider_size().
	//
	// Deprecated: since version 3.20.
	MinSliderSize() int
	// RestrictToFillLevel gets whether the range is restricted to the fill
	// level.
	RestrictToFillLevel() bool
	// RoundDigits gets the number of digits to round the value to when it
	// changes. See Range::change-value.
	RoundDigits() int
	// ShowFillLevel gets whether the range displays the fill level graphically.
	ShowFillLevel() bool
	// SliderRange: this function returns sliders range along the long
	// dimension, in widget->window coordinates.
	//
	// This function is useful mainly for Range subclasses.
	SliderRange() (sliderStart int, sliderEnd int)
	// SliderSizeFixed: this function is useful mainly for Range subclasses.
	//
	// See gtk_range_set_slider_size_fixed().
	SliderSizeFixed() bool
	// UpperStepperSensitivity gets the sensitivity policy for the stepper that
	// points to the 'upper' end of the GtkRange’s adjustment.
	UpperStepperSensitivity() SensitivityType
	// Value gets the current value of the range.
	Value() float64
	// SetAdjustment sets the adjustment to be used as the “model” object for
	// this range widget. The adjustment indicates the current range value, the
	// minimum and maximum range values, the step/page increments used for
	// keybindings and scrolling, and the page size. The page size is normally 0
	// for Scale and nonzero for Scrollbar, and indicates the size of the
	// visible area of the widget being scrolled. The page size affects the size
	// of the scrollbar slider.
	SetAdjustment(adjustment Adjustment)
	// SetFillLevel: set the new position of the fill level indicator.
	//
	// The “fill level” is probably best described by its most prominent use
	// case, which is an indicator for the amount of pre-buffering in a
	// streaming media player. In that use case, the value of the range would
	// indicate the current play position, and the fill level would be the
	// position up to which the file/stream has been downloaded.
	//
	// This amount of prebuffering can be displayed on the range’s trough and is
	// themeable separately from the trough. To enable fill level display, use
	// gtk_range_set_show_fill_level(). The range defaults to not showing the
	// fill level.
	//
	// Additionally, it’s possible to restrict the range’s slider position to
	// values which are smaller than the fill level. This is controller by
	// gtk_range_set_restrict_to_fill_level() and is by default enabled.
	SetFillLevel(fillLevel float64)
	// SetFlippable: if a range is flippable, it will switch its direction if it
	// is horizontal and its direction is GTK_TEXT_DIR_RTL.
	//
	// See gtk_widget_get_direction().
	SetFlippable(flippable bool)
	// SetIncrements sets the step and page sizes for the range. The step size
	// is used when the user clicks the Scrollbar arrows or moves Scale via
	// arrow keys. The page size is used for example when moving via Page Up or
	// Page Down keys.
	SetIncrements(step float64, page float64)
	// SetInverted ranges normally move from lower to higher values as the
	// slider moves from top to bottom or left to right. Inverted ranges have
	// higher values at the top or on the right rather than on the bottom or
	// left.
	SetInverted(setting bool)
	// SetMinSliderSize sets the minimum size of the range’s slider.
	//
	// This function is useful mainly for Range subclasses.
	//
	// Deprecated: since version 3.20.
	SetMinSliderSize(minSize int)
	// SetRange sets the allowable values in the Range, and clamps the range
	// value to be between @min and @max. (If the range has a non-zero page
	// size, it is clamped between @min and @max - page-size.)
	SetRange(min float64, max float64)
	// SetRestrictToFillLevel sets whether the slider is restricted to the fill
	// level. See gtk_range_set_fill_level() for a general description of the
	// fill level concept.
	SetRestrictToFillLevel(restrictToFillLevel bool)
	// SetRoundDigits sets the number of digits to round the value to when it
	// changes. See Range::change-value.
	SetRoundDigits(roundDigits int)
	// SetShowFillLevel sets whether a graphical fill level is show on the
	// trough. See gtk_range_set_fill_level() for a general description of the
	// fill level concept.
	SetShowFillLevel(showFillLevel bool)
	// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
	// size that depends on its adjustment’s page size.
	//
	// This function is useful mainly for Range subclasses.
	SetSliderSizeFixed(sizeFixed bool)
	// SetValue sets the current value of the range; if the value is outside the
	// minimum or maximum range values, it will be clamped to fit inside them.
	// The range emits the Range::value-changed signal if the value changes.
	SetValue(value float64)
}

// RangeClass implements the Range interface.
type RangeClass struct {
	*externglib.Object
	WidgetClass
	BuildableInterface
	OrientableInterface
}

var _ Range = (*RangeClass)(nil)

func wrapRange(obj *externglib.Object) Range {
	return &RangeClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalRange(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRange(obj), nil
}

// Adjustment: get the Adjustment which is the “model” object for Range. See
// gtk_range_set_adjustment() for details. The return value does not have a
// reference added, so should not be unreferenced.
func (r *RangeClass) Adjustment() *AdjustmentClass {
	var _arg0 *C.GtkRange      // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_adjustment(_arg0)

	var _adjustment *AdjustmentClass // out

	_adjustment = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*AdjustmentClass)

	return _adjustment
}

// FillLevel gets the current position of the fill level indicator.
func (r *RangeClass) FillLevel() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_fill_level(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Flippable gets the value set by gtk_range_set_flippable().
func (r *RangeClass) Flippable() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_flippable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Inverted gets the value set by gtk_range_set_inverted().
func (r *RangeClass) Inverted() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LowerStepperSensitivity gets the sensitivity policy for the stepper that
// points to the 'lower' end of the GtkRange’s adjustment.
func (r *RangeClass) LowerStepperSensitivity() SensitivityType {
	var _arg0 *C.GtkRange          // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_lower_stepper_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = (SensitivityType)(C.GtkSensitivityType)

	return _sensitivityType
}

// MinSliderSize: this function is useful mainly for Range subclasses.
//
// See gtk_range_set_min_slider_size().
//
// Deprecated: since version 3.20.
func (r *RangeClass) MinSliderSize() int {
	var _arg0 *C.GtkRange // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_min_slider_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RestrictToFillLevel gets whether the range is restricted to the fill level.
func (r *RangeClass) RestrictToFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_restrict_to_fill_level(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RoundDigits gets the number of digits to round the value to when it changes.
// See Range::change-value.
func (r *RangeClass) RoundDigits() int {
	var _arg0 *C.GtkRange // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_round_digits(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ShowFillLevel gets whether the range displays the fill level graphically.
func (r *RangeClass) ShowFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_show_fill_level(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SliderRange: this function returns sliders range along the long dimension, in
// widget->window coordinates.
//
// This function is useful mainly for Range subclasses.
func (r *RangeClass) SliderRange() (sliderStart int, sliderEnd int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	C.gtk_range_get_slider_range(_arg0, &_arg1, &_arg2)

	var _sliderStart int // out
	var _sliderEnd int   // out

	_sliderStart = int(_arg1)
	_sliderEnd = int(_arg2)

	return _sliderStart, _sliderEnd
}

// SliderSizeFixed: this function is useful mainly for Range subclasses.
//
// See gtk_range_set_slider_size_fixed().
func (r *RangeClass) SliderSizeFixed() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_slider_size_fixed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UpperStepperSensitivity gets the sensitivity policy for the stepper that
// points to the 'upper' end of the GtkRange’s adjustment.
func (r *RangeClass) UpperStepperSensitivity() SensitivityType {
	var _arg0 *C.GtkRange          // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_upper_stepper_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = (SensitivityType)(C.GtkSensitivityType)

	return _sensitivityType
}

// Value gets the current value of the range.
func (r *RangeClass) Value() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))

	_cret = C.gtk_range_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the adjustment to be used as the “model” object for this
// range widget. The adjustment indicates the current range value, the minimum
// and maximum range values, the step/page increments used for keybindings and
// scrolling, and the page size. The page size is normally 0 for Scale and
// nonzero for Scrollbar, and indicates the size of the visible area of the
// widget being scrolled. The page size affects the size of the scrollbar
// slider.
func (r *RangeClass) SetAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkRange      // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((&Adjustment).Native()))

	C.gtk_range_set_adjustment(_arg0, _arg1)
}

// SetFillLevel: set the new position of the fill level indicator.
//
// The “fill level” is probably best described by its most prominent use case,
// which is an indicator for the amount of pre-buffering in a streaming media
// player. In that use case, the value of the range would indicate the current
// play position, and the fill level would be the position up to which the
// file/stream has been downloaded.
//
// This amount of prebuffering can be displayed on the range’s trough and is
// themeable separately from the trough. To enable fill level display, use
// gtk_range_set_show_fill_level(). The range defaults to not showing the fill
// level.
//
// Additionally, it’s possible to restrict the range’s slider position to values
// which are smaller than the fill level. This is controller by
// gtk_range_set_restrict_to_fill_level() and is by default enabled.
func (r *RangeClass) SetFillLevel(fillLevel float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	_arg1 = C.gdouble(fillLevel)

	C.gtk_range_set_fill_level(_arg0, _arg1)
}

// SetFlippable: if a range is flippable, it will switch its direction if it is
// horizontal and its direction is GTK_TEXT_DIR_RTL.
//
// See gtk_widget_get_direction().
func (r *RangeClass) SetFlippable(flippable bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	if flippable {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_flippable(_arg0, _arg1)
}

// SetIncrements sets the step and page sizes for the range. The step size is
// used when the user clicks the Scrollbar arrows or moves Scale via arrow keys.
// The page size is used for example when moving via Page Up or Page Down keys.
func (r *RangeClass) SetIncrements(step float64, page float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	_arg1 = C.gdouble(step)
	_arg2 = C.gdouble(page)

	C.gtk_range_set_increments(_arg0, _arg1, _arg2)
}

// SetInverted ranges normally move from lower to higher values as the slider
// moves from top to bottom or left to right. Inverted ranges have higher values
// at the top or on the right rather than on the bottom or left.
func (r *RangeClass) SetInverted(setting bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_inverted(_arg0, _arg1)
}

// SetMinSliderSize sets the minimum size of the range’s slider.
//
// This function is useful mainly for Range subclasses.
//
// Deprecated: since version 3.20.
func (r *RangeClass) SetMinSliderSize(minSize int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	_arg1 = C.gint(minSize)

	C.gtk_range_set_min_slider_size(_arg0, _arg1)
}

// SetRange sets the allowable values in the Range, and clamps the range value
// to be between @min and @max. (If the range has a non-zero page size, it is
// clamped between @min and @max - page-size.)
func (r *RangeClass) SetRange(min float64, max float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)

	C.gtk_range_set_range(_arg0, _arg1, _arg2)
}

// SetRestrictToFillLevel sets whether the slider is restricted to the fill
// level. See gtk_range_set_fill_level() for a general description of the fill
// level concept.
func (r *RangeClass) SetRestrictToFillLevel(restrictToFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	if restrictToFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_restrict_to_fill_level(_arg0, _arg1)
}

// SetRoundDigits sets the number of digits to round the value to when it
// changes. See Range::change-value.
func (r *RangeClass) SetRoundDigits(roundDigits int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	_arg1 = C.gint(roundDigits)

	C.gtk_range_set_round_digits(_arg0, _arg1)
}

// SetShowFillLevel sets whether a graphical fill level is show on the trough.
// See gtk_range_set_fill_level() for a general description of the fill level
// concept.
func (r *RangeClass) SetShowFillLevel(showFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	if showFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_show_fill_level(_arg0, _arg1)
}

// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
// size that depends on its adjustment’s page size.
//
// This function is useful mainly for Range subclasses.
func (r *RangeClass) SetSliderSizeFixed(sizeFixed bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	if sizeFixed {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_slider_size_fixed(_arg0, _arg1)
}

// SetValue sets the current value of the range; if the value is outside the
// minimum or maximum range values, it will be clamped to fit inside them. The
// range emits the Range::value-changed signal if the value changes.
func (r *RangeClass) SetValue(value float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer((&Range).Native()))
	_arg1 = C.gdouble(value)

	C.gtk_range_set_value(_arg0, _arg1)
}
