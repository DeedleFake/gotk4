// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeDragIcon = coreglib.Type(C.gtk_drag_icon_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDragIcon, F: marshalDragIcon},
	})
}

// DragIconOverrides contains methods that are overridable.
type DragIconOverrides struct {
}

func defaultDragIconOverrides(v *DragIcon) DragIconOverrides {
	return DragIconOverrides{}
}

// DragIcon: GtkDragIcon is a GtkRoot implementation for drag icons.
//
// A drag icon moves with the pointer during a Drag-and-Drop operation and is
// destroyed when the drag ends.
//
// To set up a drag icon and associate it with an ongoing drag operation, use
// gtk.DragIcon().GetForDrag to get the icon for a drag. You can then use it
// like any other widget and use gtk.DragIcon.SetChild() to set whatever widget
// should be used for the drag icon.
//
// Keep in mind that drag icons do not allow user input.
type DragIcon struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Root
}

var (
	_ Widgetter         = (*DragIcon)(nil)
	_ coreglib.Objector = (*DragIcon)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DragIcon, *DragIconClass, DragIconOverrides](
		GTypeDragIcon,
		initDragIconClass,
		wrapDragIcon,
		defaultDragIconOverrides,
	)
}

func initDragIconClass(gclass unsafe.Pointer, overrides DragIconOverrides, classInitFunc func(*DragIconClass)) {
	if classInitFunc != nil {
		class := (*DragIconClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDragIcon(obj *coreglib.Object) *DragIcon {
	return &DragIcon{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Root: Root{
			NativeSurface: NativeSurface{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					Accessible: Accessible{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					ConstraintTarget: ConstraintTarget{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalDragIcon(p uintptr) (interface{}, error) {
	return wrapDragIcon(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Child gets the widget currently used as drag icon.
//
// The function returns the following values:
//
//    - widget (optional): drag icon or NULL if none.
//
func (self *DragIcon) Child() Widgetter {
	var _arg0 *C.GtkDragIcon // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkDragIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_drag_icon_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
	}

	return _widget
}

// SetChild sets the widget to display as the drag icon.
//
// The function takes the following parameters:
//
//    - child (optional): GtkWidget or NULL.
//
func (self *DragIcon) SetChild(child Widgetter) {
	var _arg0 *C.GtkDragIcon // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkDragIcon)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.gtk_drag_icon_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// DragIconCreateWidgetForValue creates a widget that can be used as a drag icon
// for the given value.
//
// Supported types include strings, GdkRGBA and GtkTextBuffer. If GTK does not
// know how to create a widget for a given value, it will return NULL.
//
// This method is used to set the default drag icon on drag'n'drop operations
// started by GtkDragSource, so you don't need to set a drag icon using this
// function there.
//
// The function takes the following parameters:
//
//    - value: GValue.
//
// The function returns the following values:
//
//    - widget (optional): new GtkWidget for displaying value as a drag icon.
//
func DragIconCreateWidgetForValue(value *coreglib.Value) Widgetter {
	var _arg1 *C.GValue    // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gtk_drag_icon_create_widget_for_value(_arg1)
	runtime.KeepAlive(value)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
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
	}

	return _widget
}

// DragIconGetForDrag gets the GtkDragIcon in use with drag.
//
// If no drag icon exists yet, a new one will be created and shown.
//
// The function takes the following parameters:
//
//    - drag: GdkDrag.
//
// The function returns the following values:
//
//    - widget: GtkDragIcon.
//
func DragIconGetForDrag(drag gdk.Dragger) Widgetter {
	var _arg1 *C.GdkDrag   // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gtk_drag_icon_get_for_drag(_arg1)
	runtime.KeepAlive(drag)

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

// DragIconSetFromPaintable creates a GtkDragIcon that shows paintable, and
// associates it with the drag operation.
//
// The hotspot position on the paintable is aligned with the hotspot of the
// cursor.
//
// The function takes the following parameters:
//
//    - drag: GdkDrag.
//    - paintable: GdkPaintable to display.
//    - hotX: x coordinate of the hotspot.
//    - hotY: y coordinate of the hotspot.
//
func DragIconSetFromPaintable(drag gdk.Dragger, paintable gdk.Paintabler, hotX, hotY int) {
	var _arg1 *C.GdkDrag      // out
	var _arg2 *C.GdkPaintable // out
	var _arg3 C.int           // out
	var _arg4 C.int           // out

	_arg1 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	_arg2 = (*C.GdkPaintable)(unsafe.Pointer(coreglib.InternObject(paintable).Native()))
	_arg3 = C.int(hotX)
	_arg4 = C.int(hotY)

	C.gtk_drag_icon_set_from_paintable(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(drag)
	runtime.KeepAlive(paintable)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// DragIconClass: instance of this type is always passed by reference.
type DragIconClass struct {
	*dragIconClass
}

// dragIconClass is the struct that's finalized.
type dragIconClass struct {
	native *C.GtkDragIconClass
}

func (d *DragIconClass) ParentClass() *WidgetClass {
	valptr := &d.native.parent_class
	var _v *WidgetClass // out
	_v = (*WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
