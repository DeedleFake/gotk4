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
		{T: externglib.Type(C.gtk_center_layout_get_type()), F: marshalCenterLayout},
	})
}

// CenterLayout: `GtkCenterLayout` is a layout manager that manages up to three
// children.
//
// The start widget is allocated at the start of the layout (left in
// left-to-right locales and right in right-to-left ones), and the end widget at
// the end.
//
// The center widget is centered regarding the full width of the layout's.
type CenterLayout interface {
	LayoutManager

	// AsLayoutManager casts the class to the LayoutManager interface.
	AsLayoutManager() LayoutManager

	// Allocate assigns the given @width, @height, and @baseline to a @widget,
	// and computes the position and sizes of the children of the @widget using
	// the layout management policy of @manager.
	//
	// This method is inherited from LayoutManager
	Allocate(widget Widget, width int, height int, baseline int)
	// GetLayoutChild retrieves a `GtkLayoutChild` instance for the
	// `GtkLayoutManager`, creating one if necessary.
	//
	// The @child widget must be a child of the widget using @manager.
	//
	// The `GtkLayoutChild` instance is owned by the `GtkLayoutManager`, and is
	// guaranteed to exist as long as @child is a child of the `GtkWidget` using
	// the given `GtkLayoutManager`.
	//
	// This method is inherited from LayoutManager
	GetLayoutChild(child Widget) LayoutChild
	// GetRequestMode retrieves the request mode of @manager.
	//
	// This method is inherited from LayoutManager
	GetRequestMode() SizeRequestMode
	// GetWidget retrieves the `GtkWidget` using the given `GtkLayoutManager`.
	//
	// This method is inherited from LayoutManager
	GetWidget() Widget
	// LayoutChanged queues a resize on the `GtkWidget` using @manager, if any.
	//
	// This function should be called by subclasses of `GtkLayoutManager` in
	// response to changes to their layout management policies.
	//
	// This method is inherited from LayoutManager
	LayoutChanged()
	// Measure measures the size of the @widget using @manager, for the given
	// @orientation and size.
	//
	// See the [class@Gtk.Widget] documentation on layout management for more
	// details.
	//
	// This method is inherited from LayoutManager
	Measure(widget Widget, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int)

	// BaselinePosition returns the baseline position of the layout.
	BaselinePosition() BaselinePosition
	// CenterWidget returns the center widget of the layout.
	CenterWidget() Widget
	// EndWidget returns the end widget of the layout.
	EndWidget() Widget
	// Orientation gets the current orienration of the layout manager.
	Orientation() Orientation
	// StartWidget returns the start widget fo the layout.
	StartWidget() Widget
	// SetBaselinePosition sets the new baseline position of @self
	SetBaselinePosition(baselinePosition BaselinePosition)
	// SetCenterWidget sets the new center widget of @self.
	//
	// To remove the existing center widget, pass nil.
	SetCenterWidget(widget Widget)
	// SetEndWidget sets the new end widget of @self.
	//
	// To remove the existing center widget, pass nil.
	SetEndWidget(widget Widget)
	// SetOrientation sets the orientation of @self.
	SetOrientation(orientation Orientation)
	// SetStartWidget sets the new start widget of @self.
	//
	// To remove the existing start widget, pass nil.
	SetStartWidget(widget Widget)
}

// centerLayout implements the CenterLayout interface.
type centerLayout struct {
	*externglib.Object
}

var _ CenterLayout = (*centerLayout)(nil)

// WrapCenterLayout wraps a GObject to a type that implements
// interface CenterLayout. It is primarily used internally.
func WrapCenterLayout(obj *externglib.Object) CenterLayout {
	return centerLayout{obj}
}

func marshalCenterLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCenterLayout(obj), nil
}

// NewCenterLayout creates a new `GtkCenterLayout`.
func NewCenterLayout() CenterLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_center_layout_new()

	var _centerLayout CenterLayout // out

	_centerLayout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(CenterLayout)

	return _centerLayout
}

func (c centerLayout) AsLayoutManager() LayoutManager {
	return WrapLayoutManager(gextras.InternObject(c))
}

func (m centerLayout) Allocate(widget Widget, width int, height int, baseline int) {
	WrapLayoutManager(gextras.InternObject(m)).Allocate(widget, width, height, baseline)
}

func (m centerLayout) GetLayoutChild(child Widget) LayoutChild {
	return WrapLayoutManager(gextras.InternObject(m)).GetLayoutChild(child)
}

func (m centerLayout) GetRequestMode() SizeRequestMode {
	return WrapLayoutManager(gextras.InternObject(m)).GetRequestMode()
}

func (m centerLayout) GetWidget() Widget {
	return WrapLayoutManager(gextras.InternObject(m)).GetWidget()
}

func (m centerLayout) LayoutChanged() {
	WrapLayoutManager(gextras.InternObject(m)).LayoutChanged()
}

func (m centerLayout) Measure(widget Widget, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int) {
	return WrapLayoutManager(gextras.InternObject(m)).Measure(widget, orientation, forSize)
}

func (s centerLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterLayout    // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

func (s centerLayout) CenterWidget() Widget {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_center_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerLayout) EndWidget() Widget {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_end_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerLayout) Orientation() Orientation {
	var _arg0 *C.GtkCenterLayout // out
	var _cret C.GtkOrientation   // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_orientation(_arg0)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

func (s centerLayout) StartWidget() Widget {
	var _arg0 *C.GtkCenterLayout // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_center_layout_get_start_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s centerLayout) SetBaselinePosition(baselinePosition BaselinePosition) {
	var _arg0 *C.GtkCenterLayout    // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkBaselinePosition(baselinePosition)

	C.gtk_center_layout_set_baseline_position(_arg0, _arg1)
}

func (s centerLayout) SetCenterWidget(widget Widget) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_center_widget(_arg0, _arg1)
}

func (s centerLayout) SetEndWidget(widget Widget) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_end_widget(_arg0, _arg1)
}

func (s centerLayout) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 C.GtkOrientation   // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_center_layout_set_orientation(_arg0, _arg1)
}

func (s centerLayout) SetStartWidget(widget Widget) {
	var _arg0 *C.GtkCenterLayout // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_start_widget(_arg0, _arg1)
}
