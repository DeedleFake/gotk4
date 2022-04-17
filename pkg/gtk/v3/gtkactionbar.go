// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkactionbar.go.
var GTypeActionBar = externglib.Type(C.gtk_action_bar_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeActionBar, F: marshalActionBar},
	})
}

// ActionBarOverrider contains methods that are overridable.
type ActionBarOverrider interface {
}

// ActionBar is designed to present contextual actions. It is expected to be
// displayed below the content and expand horizontally to fill the area.
//
// It allows placing children at the start or the end. In addition, it contains
// an internal centered box which is centered with respect to the full width of
// the box, even if the children at either side take up different amounts of
// space.
//
//
// CSS nodes
//
// GtkActionBar has a single CSS node with name actionbar.
type ActionBar struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*ActionBar)(nil)
)

func classInitActionBarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapActionBar(obj *externglib.Object) *ActionBar {
	return &ActionBar{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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
		},
	}
}

func marshalActionBar(p uintptr) (interface{}, error) {
	return wrapActionBar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewActionBar creates a new ActionBar widget.
//
// The function returns the following values:
//
//    - actionBar: new ActionBar.
//
func NewActionBar() *ActionBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_action_bar_new()

	var _actionBar *ActionBar // out

	_actionBar = wrapActionBar(externglib.Take(unsafe.Pointer(_cret)))

	return _actionBar
}

// CenterWidget retrieves the center bar widget of the bar.
//
// The function returns the following values:
//
//    - widget (optional): center Widget or NULL.
//
func (actionBar *ActionBar) CenterWidget() Widgetter {
	var _arg0 *C.GtkActionBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(externglib.InternObject(actionBar).Native()))

	_cret = C.gtk_action_bar_get_center_widget(_arg0)
	runtime.KeepAlive(actionBar)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// PackEnd adds child to action_bar, packed with reference to the end of the
// action_bar.
//
// The function takes the following parameters:
//
//    - child to be added to action_bar.
//
func (actionBar *ActionBar) PackEnd(child Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(externglib.InternObject(actionBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))

	C.gtk_action_bar_pack_end(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// PackStart adds child to action_bar, packed with reference to the start of the
// action_bar.
//
// The function takes the following parameters:
//
//    - child to be added to action_bar.
//
func (actionBar *ActionBar) PackStart(child Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(externglib.InternObject(actionBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))

	C.gtk_action_bar_pack_start(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// SetCenterWidget sets the center widget for the ActionBar.
//
// The function takes the following parameters:
//
//    - centerWidget (optional): widget to use for the center.
//
func (actionBar *ActionBar) SetCenterWidget(centerWidget Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(externglib.InternObject(actionBar).Native()))
	if centerWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(centerWidget).Native()))
	}

	C.gtk_action_bar_set_center_widget(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(centerWidget)
}
