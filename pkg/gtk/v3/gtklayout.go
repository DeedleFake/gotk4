// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeLayout = coreglib.Type(C.gtk_layout_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLayout, F: marshalLayout},
	})
}

// LayoutOverrides contains methods that are overridable.
type LayoutOverrides struct {
}

func defaultLayoutOverrides(v *Layout) LayoutOverrides {
	return LayoutOverrides{}
}

// Layout is similar to DrawingArea in that it’s a “blank slate” and doesn’t do
// anything except paint a blank background by default. It’s different in that
// it supports scrolling natively due to implementing Scrollable, and can
// contain child widgets since it’s a Container.
//
// If you just want to draw, a DrawingArea is a better choice since it has lower
// overhead. If you just need to position child widgets at specific points, then
// Fixed provides that functionality on its own.
//
// When handling expose events on a Layout, you must draw to the Window returned
// by gtk_layout_get_bin_window(), rather than to the one returned by
// gtk_widget_get_window() as you would for a DrawingArea.
type Layout struct {
	_ [0]func() // equal guard
	Container

	*coreglib.Object
	Scrollable
}

var (
	_ Containerer       = (*Layout)(nil)
	_ coreglib.Objector = (*Layout)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Layout, *LayoutClass, LayoutOverrides](
		GTypeLayout,
		initLayoutClass,
		wrapLayout,
		defaultLayoutOverrides,
	)
}

func initLayoutClass(gclass unsafe.Pointer, overrides LayoutOverrides, classInitFunc func(*LayoutClass)) {
	if classInitFunc != nil {
		class := (*LayoutClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapLayout(obj *coreglib.Object) *Layout {
	return &Layout{
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
		Object: obj,
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalLayout(p uintptr) (interface{}, error) {
	return wrapLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewLayout creates a new Layout. Unless you have a specific adjustment you’d
// like the layout to use for scrolling, pass NULL for hadjustment and
// vadjustment.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): horizontal scroll adjustment, or NULL.
//    - vadjustment (optional): vertical scroll adjustment, or NULL.
//
// The function returns the following values:
//
//    - layout: new Layout.
//
func NewLayout(hadjustment, vadjustment *Adjustment) *Layout {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	if hadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}
	if vadjustment != nil {
		_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	_cret = C.gtk_layout_new(_arg1, _arg2)
	runtime.KeepAlive(hadjustment)
	runtime.KeepAlive(vadjustment)

	var _layout *Layout // out

	_layout = wrapLayout(coreglib.Take(unsafe.Pointer(_cret)))

	return _layout
}

// BinWindow: retrieve the bin window of the layout used for drawing operations.
//
// The function returns the following values:
//
//    - window: Window.
//
func (layout *Layout) BinWindow() gdk.Windower {
	var _arg0 *C.GtkLayout // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	_cret = C.gtk_layout_get_bin_window(_arg0)
	runtime.KeepAlive(layout)

	var _window gdk.Windower // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gdk.Windower)
			return ok
		})
		rv, ok := casted.(gdk.Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_window = rv
	}

	return _window
}

// HAdjustment: this function should only be called after the layout has been
// placed in a ScrolledWindow or otherwise configured for scrolling. It returns
// the Adjustment used for communication between the horizontal scrollbar and
// layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
//
// The function returns the following values:
//
//    - adjustment: horizontal scroll adjustment.
//
func (layout *Layout) HAdjustment() *Adjustment {
	var _arg0 *C.GtkLayout     // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	_cret = C.gtk_layout_get_hadjustment(_arg0)
	runtime.KeepAlive(layout)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// Size gets the size that has been set on the layout, and that determines the
// total extents of the layout’s scrollbar area. See gtk_layout_set_size ().
//
// The function returns the following values:
//
//    - width (optional): location to store the width set on layout, or NULL.
//    - height (optional): location to store the height set on layout, or NULL.
//
func (layout *Layout) Size() (width, height uint) {
	var _arg0 *C.GtkLayout // out
	var _arg1 C.guint      // in
	var _arg2 C.guint      // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	C.gtk_layout_get_size(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(layout)

	var _width uint  // out
	var _height uint // out

	_width = uint(_arg1)
	_height = uint(_arg2)

	return _width, _height
}

// VAdjustment: this function should only be called after the layout has been
// placed in a ScrolledWindow or otherwise configured for scrolling. It returns
// the Adjustment used for communication between the vertical scrollbar and
// layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
//
// The function returns the following values:
//
//    - adjustment: vertical scroll adjustment.
//
func (layout *Layout) VAdjustment() *Adjustment {
	var _arg0 *C.GtkLayout     // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))

	_cret = C.gtk_layout_get_vadjustment(_arg0)
	runtime.KeepAlive(layout)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// Move moves a current child of layout to a new position.
//
// The function takes the following parameters:
//
//    - childWidget: current child of layout.
//    - x: x position to move to.
//    - y: y position to move to.
//
func (layout *Layout) Move(childWidget Widgetter, x, y int) {
	var _arg0 *C.GtkLayout // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(childWidget).Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_layout_move(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(childWidget)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// Put adds child_widget to layout, at position (x,y). layout becomes the new
// parent container of child_widget.
//
// The function takes the following parameters:
//
//    - childWidget: child widget.
//    - x: x position of child widget.
//    - y: y position of child widget.
//
func (layout *Layout) Put(childWidget Widgetter, x, y int) {
	var _arg0 *C.GtkLayout // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(childWidget).Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_layout_put(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(childWidget)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// SetHAdjustment sets the horizontal scroll adjustment for the layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_set_hadjustment().
//
// The function takes the following parameters:
//
//    - adjustment (optional): new scroll adjustment.
//
func (layout *Layout) SetHAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkLayout     // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	C.gtk_layout_set_hadjustment(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(adjustment)
}

// SetSize sets the size of the scrollable area of the layout.
//
// The function takes the following parameters:
//
//    - width of entire scrollable area.
//    - height of entire scrollable area.
//
func (layout *Layout) SetSize(width, height uint) {
	var _arg0 *C.GtkLayout // out
	var _arg1 C.guint      // out
	var _arg2 C.guint      // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	_arg1 = C.guint(width)
	_arg2 = C.guint(height)

	C.gtk_layout_set_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetVAdjustment sets the vertical scroll adjustment for the layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_set_vadjustment().
//
// The function takes the following parameters:
//
//    - adjustment (optional): new scroll adjustment.
//
func (layout *Layout) SetVAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkLayout     // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(coreglib.InternObject(layout).Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	C.gtk_layout_set_vadjustment(_arg0, _arg1)
	runtime.KeepAlive(layout)
	runtime.KeepAlive(adjustment)
}

// LayoutClass: instance of this type is always passed by reference.
type LayoutClass struct {
	*layoutClass
}

// layoutClass is the struct that's finalized.
type layoutClass struct {
	native *C.GtkLayoutClass
}

func (l *LayoutClass) ParentClass() *ContainerClass {
	valptr := &l.native.parent_class
	var _v *ContainerClass // out
	_v = (*ContainerClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
