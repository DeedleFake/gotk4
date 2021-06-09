// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
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
	LayoutManager
}

var _ CenterLayout = (*centerLayout)(nil)

// WrapCenterLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapCenterLayout(obj *externglib.Object) CenterLayout {
	return CenterLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalCenterLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCenterLayout(obj), nil
}

// NewCenterLayout constructs a class CenterLayout.
func NewCenterLayout() CenterLayout {
	var _cret C.GtkCenterLayout

	cret = C.gtk_center_layout_new()

	var _centerLayout CenterLayout

	_centerLayout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(CenterLayout)

	return _centerLayout
}

// BaselinePosition returns the baseline position of the layout.
func (s centerLayout) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterLayout

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	var _cret C.GtkBaselinePosition

	cret = C.gtk_center_layout_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// CenterWidget returns the center widget of the layout.
func (s centerLayout) CenterWidget() Widget {
	var _arg0 *C.GtkCenterLayout

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_center_layout_get_center_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// EndWidget returns the end widget of the layout.
func (s centerLayout) EndWidget() Widget {
	var _arg0 *C.GtkCenterLayout

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_center_layout_get_end_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Orientation gets the current orienration of the layout manager.
func (s centerLayout) Orientation() Orientation {
	var _arg0 *C.GtkCenterLayout

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	var _cret C.GtkOrientation

	cret = C.gtk_center_layout_get_orientation(_arg0)

	var _orientation Orientation

	_orientation = Orientation(_cret)

	return _orientation
}

// StartWidget returns the start widget fo the layout.
func (s centerLayout) StartWidget() Widget {
	var _arg0 *C.GtkCenterLayout

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_center_layout_get_start_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// SetBaselinePosition sets the new baseline position of @self
func (s centerLayout) SetBaselinePosition(baselinePosition BaselinePosition) {
	var _arg0 *C.GtkCenterLayout
	var _arg1 C.GtkBaselinePosition

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkBaselinePosition)(baselinePosition)

	C.gtk_center_layout_set_baseline_position(_arg0, _arg1)
}

// SetCenterWidget sets the new center widget of @self.
//
// To remove the existing center widget, pass nil.
func (s centerLayout) SetCenterWidget(widget Widget) {
	var _arg0 *C.GtkCenterLayout
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_center_widget(_arg0, _arg1)
}

// SetEndWidget sets the new end widget of @self.
//
// To remove the existing center widget, pass nil.
func (s centerLayout) SetEndWidget(widget Widget) {
	var _arg0 *C.GtkCenterLayout
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_end_widget(_arg0, _arg1)
}

// SetOrientation sets the orientation of @self.
func (s centerLayout) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkCenterLayout
	var _arg1 C.GtkOrientation

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkOrientation)(orientation)

	C.gtk_center_layout_set_orientation(_arg0, _arg1)
}

// SetStartWidget sets the new start widget of @self.
//
// To remove the existing start widget, pass nil.
func (s centerLayout) SetStartWidget(widget Widget) {
	var _arg0 *C.GtkCenterLayout
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkCenterLayout)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_center_layout_set_start_widget(_arg0, _arg1)
}
