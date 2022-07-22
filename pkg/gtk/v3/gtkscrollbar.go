// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
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
	GTypeScrollbar = coreglib.Type(C.gtk_scrollbar_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScrollbar, F: marshalScrollbar},
	})
}

// ScrollbarOverrider contains methods that are overridable.
type ScrollbarOverrider interface {
}

// Scrollbar widget is a horizontal or vertical scrollbar, depending on the
// value of the Orientable:orientation property.
//
// Its position and movement are controlled by the adjustment that is passed to
// or created by gtk_scrollbar_new(). See Adjustment for more details. The
// Adjustment:value field sets the position of the thumb and must be between
// Adjustment:lower and Adjustment:upper - Adjustment:page-size. The
// Adjustment:page-size represents the size of the visible scrollable area. The
// fields Adjustment:step-increment and Adjustment:page-increment fields are
// added to or subtracted from the Adjustment:value when the user asks to move
// by a step (using e.g. the cursor arrow keys or, if present, the stepper
// buttons) or by a page (using e.g. the Page Down/Up keys).
//
// CSS nodes
//
//    scrollbar[.fine-tune]
//    ╰── contents
//        ├── [button.up]
//        ├── [button.down]
//        ├── trough
//        │   ╰── slider
//        ├── [button.up]
//        ╰── [button.down]
//
// GtkScrollbar has a main CSS node with name scrollbar and a subnode for its
// contents, with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scrollbar is in
// 'fine-tuning' mode.
//
// If steppers are enabled, they are represented by up to four additional
// subnodes with name button. These get the style classes .up and .down to
// indicate in which direction they are moving.
//
// Other style classes that may be added to scrollbars inside ScrolledWindow
// include the positional classes (.left, .right, .top, .bottom) and style
// classes related to overlay scrolling (.overlay-indicator, .dragging,
// .hovering).
type Scrollbar struct {
	_ [0]func() // equal guard
	Range
}

var (
	_ Ranger = (*Scrollbar)(nil)
)

func initClassScrollbar(gclass unsafe.Pointer, goval any) {
}

func wrapScrollbar(obj *coreglib.Object) *Scrollbar {
	return &Scrollbar{
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
	}
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	return wrapScrollbar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewScrollbar creates a new scrollbar with the given orientation.
//
// The function takes the following parameters:
//
//    - orientation scrollbar’s orientation.
//    - adjustment (optional) to use, or NULL to create a new adjustment.
//
// The function returns the following values:
//
//    - scrollbar: new Scrollbar.
//
func NewScrollbar(orientation Orientation, adjustment *Adjustment) *Scrollbar {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	if adjustment != nil {
		_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_cret = C.gtk_scrollbar_new(_arg1, _arg2)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(adjustment)

	var _scrollbar *Scrollbar // out

	_scrollbar = wrapScrollbar(coreglib.Take(unsafe.Pointer(_cret)))

	return _scrollbar
}
