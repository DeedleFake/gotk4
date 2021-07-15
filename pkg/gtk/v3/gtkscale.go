// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_scale_get_type()), F: marshalScaler},
	})
}

// ScaleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScaleOverrider interface {
	DrawValue()
	FormatValue(value float64) string
	// LayoutOffsets obtains the coordinates where the scale will draw the
	// Layout representing the text in the scale. Remember when using the Layout
	// function you need to convert to and from pixels using PANGO_PIXELS() or
	// NGO_SCALE.
	//
	// If the Scale:draw-value property is FALSE, the return values are
	// undefined.
	LayoutOffsets() (x int, y int)
}

// Scale is a slider control used to select a numeric value. To use it, you’ll
// probably want to investigate the methods on its base class, Range, in
// addition to the methods for GtkScale itself. To set the value of a scale, you
// would normally use gtk_range_set_value(). To detect changes to the value, you
// would normally use the Range::value-changed signal.
//
// Note that using the same upper and lower bounds for the Scale (through the
// Range methods) will hide the slider itself. This is useful for applications
// that want to show an undeterminate value on the scale, without changing the
// layout of the application (such as movie or music players).
//
//
// GtkScale as GtkBuildable
//
// GtkScale supports a custom <marks> element, which can contain multiple <mark>
// elements. The “value” and “position” attributes have the same meaning as
// gtk_scale_add_mark() parameters of the same name. If the element is not
// empty, its content is taken as the markup to show at the mark. It can be
// translated with the usual ”translatable” and “context” attributes.
//
// CSS nodes
//
//    scale[.fine-tune][.marks-before][.marks-after]
//    ├── marks.top
//    │   ├── mark
//    │   ┊    ├── [label]
//    │   ┊    ╰── indicator
//    ┊   ┊
//    │   ╰── mark
//    ├── [value]
//    ├── contents
//    │   ╰── trough
//    │       ├── slider
//    │       ├── [highlight]
//    │       ╰── [fill]
//    ╰── marks.bottom
//        ├── mark
//        ┊    ├── indicator
//        ┊    ╰── [label]
//        ╰── mark
//
// GtkScale has a main CSS node with name scale and a subnode for its contents,
// with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scale is in
// 'fine-tuning' mode.
//
// If the scale has an origin (see gtk_scale_set_has_origin()), there is a
// subnode with name highlight below the trough node that is used for rendering
// the highlighted part of the trough.
//
// If the scale is showing a fill level (see gtk_range_set_show_fill_level()),
// there is a subnode with name fill below the trough node that is used for
// rendering the filled in part of the trough.
//
// If marks are present, there is a marks subnode before or after the contents
// node, below which each mark gets a node with name mark. The marks nodes get
// either the .top or .bottom style class.
//
// The mark node has a subnode named indicator. If the mark has text, it also
// has a subnode named label. When the mark is either above or left of the
// scale, the label subnode is the first when present. Otherwise, the indicator
// subnode is the first.
//
// The main CSS node gets the 'marks-before' and/or 'marks-after' style classes
// added depending on what marks are present.
//
// If the scale is displaying the value (see Scale:draw-value), there is subnode
// with name value.
type Scale struct {
	Range
}

var _ gextras.Nativer = (*Scale)(nil)

func wrapScale(obj *externglib.Object) *Scale {
	return &Scale{
		Range: Range{
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
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalScaler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScale(obj), nil
}

// NewScale creates a new Scale.
func NewScale(orientation Orientation, adjustment *Adjustment) *Scale {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_scale_new(_arg1, _arg2)

	var _scale *Scale // out

	_scale = wrapScale(externglib.Take(unsafe.Pointer(_cret)))

	return _scale
}

// NewScaleWithRange creates a new scale widget with the given orientation that
// lets the user input a number between min and max (including min and max) with
// the increment step. step must be nonzero; it’s the distance the slider moves
// when using the arrow keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk_scale_set_digits() to correct it.
func NewScaleWithRange(orientation Orientation, min float64, max float64, step float64) *Scale {
	var _arg1 C.GtkOrientation // out
	var _arg2 C.gdouble        // out
	var _arg3 C.gdouble        // out
	var _arg4 C.gdouble        // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.gdouble(min)
	_arg3 = C.gdouble(max)
	_arg4 = C.gdouble(step)

	_cret = C.gtk_scale_new_with_range(_arg1, _arg2, _arg3, _arg4)

	var _scale *Scale // out

	_scale = wrapScale(externglib.Take(unsafe.Pointer(_cret)))

	return _scale
}

// AddMark adds a mark at value.
//
// A mark is indicated visually by drawing a tick mark next to the scale, and
// GTK+ makes it easy for the user to position the scale exactly at the marks
// value.
//
// If markup is not NULL, text is shown next to the tick mark.
//
// To remove marks from a scale, use gtk_scale_clear_marks().
func (scale *Scale) AddMark(value float64, position PositionType, markup string) {
	var _arg0 *C.GtkScale       // out
	var _arg1 C.gdouble         // out
	var _arg2 C.GtkPositionType // out
	var _arg3 *C.gchar          // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))
	_arg1 = C.gdouble(value)
	_arg2 = C.GtkPositionType(position)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))

	C.gtk_scale_add_mark(_arg0, _arg1, _arg2, _arg3)
}

// ClearMarks removes any marks that have been added with gtk_scale_add_mark().
func (scale *Scale) ClearMarks() {
	var _arg0 *C.GtkScale // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	C.gtk_scale_clear_marks(_arg0)
}

// Digits gets the number of decimal places that are displayed in the value.
func (scale *Scale) Digits() int {
	var _arg0 *C.GtkScale // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_digits(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// DrawValue returns whether the current value is displayed as a string next to
// the slider.
func (scale *Scale) DrawValue() bool {
	var _arg0 *C.GtkScale // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_draw_value(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasOrigin returns whether the scale has an origin.
func (scale *Scale) HasOrigin() bool {
	var _arg0 *C.GtkScale // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_has_origin(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Layout gets the Layout used to display the scale. The returned object is
// owned by the scale so does not need to be freed by the caller.
func (scale *Scale) Layout() *pango.Layout {
	var _arg0 *C.GtkScale    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_layout(_arg0)

	var _layout *pango.Layout // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_layout = &pango.Layout{
			Object: obj,
		}
	}

	return _layout
}

// LayoutOffsets obtains the coordinates where the scale will draw the Layout
// representing the text in the scale. Remember when using the Layout function
// you need to convert to and from pixels using PANGO_PIXELS() or NGO_SCALE.
//
// If the Scale:draw-value property is FALSE, the return values are undefined.
func (scale *Scale) LayoutOffsets() (x int, y int) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	C.gtk_scale_get_layout_offsets(_arg0, &_arg1, &_arg2)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

// ValuePos gets the position in which the current value is displayed.
func (scale *Scale) ValuePos() PositionType {
	var _arg0 *C.GtkScale       // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))

	_cret = C.gtk_scale_get_value_pos(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// SetDigits sets the number of decimal places that are displayed in the value.
// Also causes the value of the adjustment to be rounded to this number of
// digits, so the retrieved value matches the displayed one, if Scale:draw-value
// is TRUE when the value changes. If you want to enforce rounding the value
// when Scale:draw-value is FALSE, you can set Range:round-digits instead.
//
// Note that rounding to a small number of digits can interfere with the smooth
// autoscrolling that is built into Scale. As an alternative, you can use the
// Scale::format-value signal to format the displayed value yourself.
func (scale *Scale) SetDigits(digits int) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))
	_arg1 = C.gint(digits)

	C.gtk_scale_set_digits(_arg0, _arg1)
}

// SetDrawValue specifies whether the current value is displayed as a string
// next to the slider.
func (scale *Scale) SetDrawValue(drawValue bool) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))
	if drawValue {
		_arg1 = C.TRUE
	}

	C.gtk_scale_set_draw_value(_arg0, _arg1)
}

// SetHasOrigin: if Scale:has-origin is set to TRUE (the default), the scale
// will highlight the part of the trough between the origin (bottom or left
// side) and the current value.
func (scale *Scale) SetHasOrigin(hasOrigin bool) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))
	if hasOrigin {
		_arg1 = C.TRUE
	}

	C.gtk_scale_set_has_origin(_arg0, _arg1)
}

// SetValuePos sets the position in which the current value is displayed.
func (scale *Scale) SetValuePos(pos PositionType) {
	var _arg0 *C.GtkScale       // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(scale.Native()))
	_arg1 = C.GtkPositionType(pos)

	C.gtk_scale_set_value_pos(_arg0, _arg1)
}
