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
	GTypeHBox = coreglib.Type(C.gtk_hbox_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHBox, F: marshalHBox},
	})
}

// HBoxOverrides contains methods that are overridable.
type HBoxOverrides struct {
}

func defaultHBoxOverrides(v *HBox) HBoxOverrides {
	return HBoxOverrides{}
}

// HBox is a container that organizes child widgets into a single row.
//
// Use the Box packing interface to determine the arrangement, spacing, width,
// and alignment of HBox children.
//
// All children are allocated the same height.
//
// GtkHBox has been deprecated. You can use Box instead, which is a very quick
// and easy change. If you have derived your own classes from GtkHBox, you can
// simply change the inheritance to derive directly from Box. No further changes
// are needed, since the default value of the Orientable:orientation property is
// GTK_ORIENTATION_HORIZONTAL.
//
// If you have a grid-like layout composed of nested boxes, and you don’t need
// first-child or last-child styling, the recommendation is to switch to Grid.
// For more information about migrating to Grid, see [Migrating from other
// containers to GtkGrid][gtk-migrating-GtkGrid].
type HBox struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*HBox)(nil)
	_ coreglib.Objector = (*HBox)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*HBox, *HBoxClass, HBoxOverrides](
		GTypeHBox,
		initHBoxClass,
		wrapHBox,
		defaultHBoxOverrides,
	)
}

func initHBoxClass(gclass unsafe.Pointer, overrides HBoxOverrides, classInitFunc func(*HBoxClass)) {
	if classInitFunc != nil {
		class := (*HBoxClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapHBox(obj *coreglib.Object) *HBox {
	return &HBox{
		Box: Box{
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
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalHBox(p uintptr) (interface{}, error) {
	return wrapHBox(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHBox creates a new HBox.
//
// Deprecated: You can use gtk_box_new() with GTK_ORIENTATION_HORIZONTAL
// instead, which is a quick and easy change. But the recommendation is to
// switch to Grid, since Box is going to go away eventually. See [Migrating from
// other containers to GtkGrid][gtk-migrating-GtkGrid].
//
// The function takes the following parameters:
//
//    - homogeneous: TRUE if all children are to be given equal space allotments.
//    - spacing: number of pixels to place by default between children.
//
// The function returns the following values:
//
//    - hBox: new HBox.
//
func NewHBox(homogeneous bool, spacing int) *HBox {
	var _arg1 C.gboolean   // out
	var _arg2 C.gint       // out
	var _cret *C.GtkWidget // in

	if homogeneous {
		_arg1 = C.TRUE
	}
	_arg2 = C.gint(spacing)

	_cret = C.gtk_hbox_new(_arg1, _arg2)
	runtime.KeepAlive(homogeneous)
	runtime.KeepAlive(spacing)

	var _hBox *HBox // out

	_hBox = wrapHBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _hBox
}

// HBoxClass: instance of this type is always passed by reference.
type HBoxClass struct {
	*hBoxClass
}

// hBoxClass is the struct that's finalized.
type hBoxClass struct {
	native *C.GtkHBoxClass
}

func (h *HBoxClass) ParentClass() *BoxClass {
	valptr := &h.native.parent_class
	var _v *BoxClass // out
	_v = (*BoxClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
