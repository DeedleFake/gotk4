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
		{T: externglib.Type(C.gtk_spin_button_update_policy_get_type()), F: marshalSpinButtonUpdatePolicy},
		{T: externglib.Type(C.gtk_spin_type_get_type()), F: marshalSpinType},
		{T: externglib.Type(C.gtk_spin_button_get_type()), F: marshalSpinButton},
	})
}

// SpinButtonUpdatePolicy determines whether the spin button displays values
// outside the adjustment bounds.
//
// See [method@Gtk.SpinButton.set_update_policy].
type SpinButtonUpdatePolicy int

const (
	// Always: when refreshing your SpinButton, the value is always displayed
	SpinButtonUpdatePolicyAlways SpinButtonUpdatePolicy = iota
	// IfValid: when refreshing your SpinButton, the value is only displayed if
	// it is valid within the bounds of the spin button's adjustment
	SpinButtonUpdatePolicyIfValid
)

func marshalSpinButtonUpdatePolicy(p uintptr) (interface{}, error) {
	return SpinButtonUpdatePolicy(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SpinType: the values of the GtkSpinType enumeration are used to specify the
// change to make in gtk_spin_button_spin().
type SpinType int

const (
	// StepForward: increment by the adjustments step increment.
	SpinTypeStepForward SpinType = iota
	// StepBackward: decrement by the adjustments step increment.
	SpinTypeStepBackward
	// PageForward: increment by the adjustments page increment.
	SpinTypePageForward
	// PageBackward: decrement by the adjustments page increment.
	SpinTypePageBackward
	// Home: go to the adjustments lower bound.
	SpinTypeHome
	// End: go to the adjustments upper bound.
	SpinTypeEnd
	// UserDefined: change by a specified amount.
	SpinTypeUserDefined
)

func marshalSpinType(p uintptr) (interface{}, error) {
	return SpinType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SpinButton: `GtkSpinButton` is an ideal way to allow the user to set the
// value of some attribute.
//
// !An example GtkSpinButton (spinbutton.png)
//
// Rather than having to directly type a number into a `GtkEntry`,
// `GtkSpinButton` allows the user to click on one of two arrows to increment or
// decrement the displayed value. A value can still be typed in, with the bonus
// that it can be checked to ensure it is in a given range.
//
// The main properties of a `GtkSpinButton` are through an adjustment. See the
// [class@Gtk.Adjustment] documentation for more details about an adjustment's
// properties.
//
// Note that `GtkSpinButton` will by default make its entry large enough to
// accommodate the lower and upper bounds of the adjustment. If this is not
// desired, the automatic sizing can be turned off by explicitly setting
// [property@Gtk.Editable:width-chars] to a value != -1.
//
//
// Using a GtkSpinButton to get an integer
//
// “`c // Provides a function to retrieve an integer value from a GtkSpinButton
// // and creates a spin button to model percentage values.
//
// int grab_int_value (GtkSpinButton *button, gpointer user_data) { return
// gtk_spin_button_get_value_as_int (button); }
//
// void create_integer_spin_button (void) {
//
//    GtkWidget *window, *button;
//    GtkAdjustment *adjustment;
//
//    adjustment = gtk_adjustment_new (50.0, 0.0, 100.0, 1.0, 5.0, 0.0);
//
//    window = gtk_window_new ();
//
//    // creates the spinbutton, with no decimal places
//    button = gtk_spin_button_new (adjustment, 1.0, 0);
//    gtk_window_set_child (GTK_WINDOW (window), button);
//
//    gtk_widget_show (window);
//
// } “`
//
//
// Using a GtkSpinButton to get a floating point value
//
// “`c // Provides a function to retrieve a floating point value from a //
// GtkSpinButton, and creates a high precision spin button.
//
// float grab_float_value (GtkSpinButton *button, gpointer user_data) { return
// gtk_spin_button_get_value (button); }
//
// void create_floating_spin_button (void) { GtkWidget *window, *button;
// GtkAdjustment *adjustment;
//
//    adjustment = gtk_adjustment_new (2.500, 0.0, 5.0, 0.001, 0.1, 0.0);
//
//    window = gtk_window_new ();
//
//    // creates the spinbutton, with three decimal places
//    button = gtk_spin_button_new (adjustment, 0.001, 3);
//    gtk_window_set_child (GTK_WINDOW (window), button);
//
//    gtk_widget_show (window);
//
// } “`
//
//
// CSS nodes
//
// “` spinbutton.horizontal ├── text │ ├── undershoot.left │ ╰──
// undershoot.right ├── button.down ╰── button.up “`
//
// “` spinbutton.vertical ├── button.up ├── text │ ├── undershoot.left │ ╰──
// undershoot.right ╰── button.down “`
//
// `GtkSpinButton`s main CSS node has the name spinbutton. It creates subnodes
// for the entry and the two buttons, with these names. The button nodes have
// the style classes .up and .down. The `GtkText` subnodes (if present) are put
// below the text node. The orientation of the spin button is reflected in the
// .vertical or .horizontal style class on the main node.
//
//
// Accessiblity
//
// `GtkSpinButton` uses the GTK_ACCESSIBLE_ROLE_SPIN_BUTTON role.
type SpinButton interface {
	gextras.Objector

	// Configure changes the properties of an existing spin button.
	//
	// The adjustment, climb rate, and number of decimal places are updated
	// accordingly.
	Configure(adjustment Adjustment, climbRate float64, digits uint)
	// Adjustment: get the adjustment associated with a `GtkSpinButton`.
	Adjustment() *AdjustmentClass
	// ClimbRate returns the acceleration rate for repeated changes.
	ClimbRate() float64
	// Digits fetches the precision of @spin_button.
	Digits() uint
	// Increments gets the current step and page the increments used by
	// @spin_button.
	//
	// See [method@Gtk.SpinButton.set_increments].
	Increments() (step float64, page float64)
	// Numeric returns whether non-numeric text can be typed into the spin
	// button.
	Numeric() bool
	// Range gets the range allowed for @spin_button.
	//
	// See [method@Gtk.SpinButton.set_range].
	Range() (min float64, max float64)
	// SnapToTicks returns whether the values are corrected to the nearest step.
	SnapToTicks() bool
	// UpdatePolicy gets the update behavior of a spin button.
	//
	// See [method@Gtk.SpinButton.set_update_policy].
	UpdatePolicy() SpinButtonUpdatePolicy
	// Value: get the value in the @spin_button.
	Value() float64
	// ValueAsInt: get the value @spin_button represented as an integer.
	ValueAsInt() int
	// Wrap returns whether the spin button’s value wraps around to the opposite
	// limit when the upper or lower limit of the range is exceeded.
	Wrap() bool
	// SetAdjustment replaces the `GtkAdjustment` associated with @spin_button.
	SetAdjustment(adjustment Adjustment)
	// SetClimbRate sets the acceleration rate for repeated changes when you
	// hold down a button or key.
	SetClimbRate(climbRate float64)
	// SetDigits: set the precision to be displayed by @spin_button.
	//
	// Up to 20 digit precision is allowed.
	SetDigits(digits uint)
	// SetIncrements sets the step and page increments for spin_button.
	//
	// This affects how quickly the value changes when the spin button’s arrows
	// are activated.
	SetIncrements(step float64, page float64)
	// SetNumeric sets the flag that determines if non-numeric text can be typed
	// into the spin button.
	SetNumeric(numeric bool)
	// SetRange sets the minimum and maximum allowable values for @spin_button.
	//
	// If the current value is outside this range, it will be adjusted to fit
	// within the range, otherwise it will remain unchanged.
	SetRange(min float64, max float64)
	// SetSnapToTicks sets the policy as to whether values are corrected to the
	// nearest step increment when a spin button is activated after providing an
	// invalid value.
	SetSnapToTicks(snapToTicks bool)
	// SetValue sets the value of @spin_button.
	SetValue(value float64)
	// SetWrap sets the flag that determines if a spin button value wraps around
	// to the opposite limit when the upper or lower limit of the range is
	// exceeded.
	SetWrap(wrap bool)
	// Update: manually force an update of the spin button.
	Update()
}

// SpinButtonClass implements the SpinButton interface.
type SpinButtonClass struct {
	*externglib.Object
	WidgetClass
	AccessibleInterface
	BuildableInterface
	CellEditableInterface
	ConstraintTargetInterface
	EditableInterface
	OrientableInterface
}

var _ SpinButton = (*SpinButtonClass)(nil)

func wrapSpinButton(obj *externglib.Object) SpinButton {
	return &SpinButtonClass{
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
		CellEditableInterface: CellEditableInterface{
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
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
		EditableInterface: EditableInterface{
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
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalSpinButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSpinButton(obj), nil
}

// NewSpinButton creates a new `GtkSpinButton`.
func NewSpinButton(adjustment Adjustment, climbRate float64, digits uint) *SpinButtonClass {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 C.double         // out
	var _arg3 C.guint          // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((&Adjustment).Native()))
	_arg2 = C.double(climbRate)
	_arg3 = C.guint(digits)

	_cret = C.gtk_spin_button_new(_arg1, _arg2, _arg3)

	var _spinButton *SpinButtonClass // out

	_spinButton = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*SpinButtonClass)

	return _spinButton
}

// NewSpinButtonWithRange creates a new `GtkSpinButton` with the given
// properties.
//
// This is a convenience constructor that allows creation of a numeric
// `GtkSpinButton` without manually creating an adjustment. The value is
// initially set to the minimum value and a page increment of 10 * @step is the
// default. The precision of the spin button is equivalent to the precision of
// @step.
//
// Note that the way in which the precision is derived works best if @step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// [method@Gtk.SpinButton.set_digits] to correct it.
func NewSpinButtonWithRange(min float64, max float64, step float64) *SpinButtonClass {
	var _arg1 C.double     // out
	var _arg2 C.double     // out
	var _arg3 C.double     // out
	var _cret *C.GtkWidget // in

	_arg1 = C.double(min)
	_arg2 = C.double(max)
	_arg3 = C.double(step)

	_cret = C.gtk_spin_button_new_with_range(_arg1, _arg2, _arg3)

	var _spinButton *SpinButtonClass // out

	_spinButton = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*SpinButtonClass)

	return _spinButton
}

// Configure changes the properties of an existing spin button.
//
// The adjustment, climb rate, and number of decimal places are updated
// accordingly.
func (s *SpinButtonClass) Configure(adjustment Adjustment, climbRate float64, digits uint) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 *C.GtkAdjustment // out
	var _arg2 C.double         // out
	var _arg3 C.guint          // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((&Adjustment).Native()))
	_arg2 = C.double(climbRate)
	_arg3 = C.guint(digits)

	C.gtk_spin_button_configure(_arg0, _arg1, _arg2, _arg3)
}

// Adjustment: get the adjustment associated with a `GtkSpinButton`.
func (s *SpinButtonClass) Adjustment() *AdjustmentClass {
	var _arg0 *C.GtkSpinButton // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_adjustment(_arg0)

	var _adjustment *AdjustmentClass // out

	_adjustment = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*AdjustmentClass)

	return _adjustment
}

// ClimbRate returns the acceleration rate for repeated changes.
func (s *SpinButtonClass) ClimbRate() float64 {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.double         // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_climb_rate(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Digits fetches the precision of @spin_button.
func (s *SpinButtonClass) Digits() uint {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_digits(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Increments gets the current step and page the increments used by
// @spin_button.
//
// See [method@Gtk.SpinButton.set_increments].
func (s *SpinButtonClass) Increments() (step float64, page float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // in
	var _arg2 C.double         // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	C.gtk_spin_button_get_increments(_arg0, &_arg1, &_arg2)

	var _step float64 // out
	var _page float64 // out

	_step = float64(_arg1)
	_page = float64(_arg2)

	return _step, _page
}

// Numeric returns whether non-numeric text can be typed into the spin button.
func (s *SpinButtonClass) Numeric() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_numeric(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Range gets the range allowed for @spin_button.
//
// See [method@Gtk.SpinButton.set_range].
func (s *SpinButtonClass) Range() (min float64, max float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // in
	var _arg2 C.double         // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	C.gtk_spin_button_get_range(_arg0, &_arg1, &_arg2)

	var _min float64 // out
	var _max float64 // out

	_min = float64(_arg1)
	_max = float64(_arg2)

	return _min, _max
}

// SnapToTicks returns whether the values are corrected to the nearest step.
func (s *SpinButtonClass) SnapToTicks() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_snap_to_ticks(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UpdatePolicy gets the update behavior of a spin button.
//
// See [method@Gtk.SpinButton.set_update_policy].
func (s *SpinButtonClass) UpdatePolicy() SpinButtonUpdatePolicy {
	var _arg0 *C.GtkSpinButton            // out
	var _cret C.GtkSpinButtonUpdatePolicy // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_update_policy(_arg0)

	var _spinButtonUpdatePolicy SpinButtonUpdatePolicy // out

	_spinButtonUpdatePolicy = (SpinButtonUpdatePolicy)(C.GtkSpinButtonUpdatePolicy)

	return _spinButtonUpdatePolicy
}

// Value: get the value in the @spin_button.
func (s *SpinButtonClass) Value() float64 {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.double         // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ValueAsInt: get the value @spin_button represented as an integer.
func (s *SpinButtonClass) ValueAsInt() int {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.int            // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_value_as_int(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Wrap returns whether the spin button’s value wraps around to the opposite
// limit when the upper or lower limit of the range is exceeded.
func (s *SpinButtonClass) Wrap() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	_cret = C.gtk_spin_button_get_wrap(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAdjustment replaces the `GtkAdjustment` associated with @spin_button.
func (s *SpinButtonClass) SetAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer((&Adjustment).Native()))

	C.gtk_spin_button_set_adjustment(_arg0, _arg1)
}

// SetClimbRate sets the acceleration rate for repeated changes when you hold
// down a button or key.
func (s *SpinButtonClass) SetClimbRate(climbRate float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	_arg1 = C.double(climbRate)

	C.gtk_spin_button_set_climb_rate(_arg0, _arg1)
}

// SetDigits: set the precision to be displayed by @spin_button.
//
// Up to 20 digit precision is allowed.
func (s *SpinButtonClass) SetDigits(digits uint) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	_arg1 = C.guint(digits)

	C.gtk_spin_button_set_digits(_arg0, _arg1)
}

// SetIncrements sets the step and page increments for spin_button.
//
// This affects how quickly the value changes when the spin button’s arrows are
// activated.
func (s *SpinButtonClass) SetIncrements(step float64, page float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // out
	var _arg2 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	_arg1 = C.double(step)
	_arg2 = C.double(page)

	C.gtk_spin_button_set_increments(_arg0, _arg1, _arg2)
}

// SetNumeric sets the flag that determines if non-numeric text can be typed
// into the spin button.
func (s *SpinButtonClass) SetNumeric(numeric bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	if numeric {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_numeric(_arg0, _arg1)
}

// SetRange sets the minimum and maximum allowable values for @spin_button.
//
// If the current value is outside this range, it will be adjusted to fit within
// the range, otherwise it will remain unchanged.
func (s *SpinButtonClass) SetRange(min float64, max float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // out
	var _arg2 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	_arg1 = C.double(min)
	_arg2 = C.double(max)

	C.gtk_spin_button_set_range(_arg0, _arg1, _arg2)
}

// SetSnapToTicks sets the policy as to whether values are corrected to the
// nearest step increment when a spin button is activated after providing an
// invalid value.
func (s *SpinButtonClass) SetSnapToTicks(snapToTicks bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	if snapToTicks {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_snap_to_ticks(_arg0, _arg1)
}

// SetValue sets the value of @spin_button.
func (s *SpinButtonClass) SetValue(value float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	_arg1 = C.double(value)

	C.gtk_spin_button_set_value(_arg0, _arg1)
}

// SetWrap sets the flag that determines if a spin button value wraps around to
// the opposite limit when the upper or lower limit of the range is exceeded.
func (s *SpinButtonClass) SetWrap(wrap bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))
	if wrap {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_wrap(_arg0, _arg1)
}

// Update: manually force an update of the spin button.
func (s *SpinButtonClass) Update() {
	var _arg0 *C.GtkSpinButton // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer((&SpinButton).Native()))

	C.gtk_spin_button_update(_arg0)
}
