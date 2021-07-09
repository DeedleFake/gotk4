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
		{T: externglib.Type(C.gtk_stack_switcher_get_type()), F: marshalStackSwitcher},
	})
}

// StackSwitcher: the GtkStackSwitcher widget acts as a controller for a Stack;
// it shows a row of buttons to switch between the various pages of the
// associated stack widget.
//
// All the content for the buttons comes from the child properties of the Stack;
// the button visibility in a StackSwitcher widget is controlled by the
// visibility of the child in the Stack.
//
// It is possible to associate multiple StackSwitcher widgets with the same
// Stack widget.
//
// The GtkStackSwitcher widget was added in 3.10.
//
//
// CSS nodes
//
// GtkStackSwitcher has a single CSS node named stackswitcher and style class
// .stack-switcher.
//
// When circumstances require it, GtkStackSwitcher adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSwitcher interface {
	gextras.Objector

	// Stack retrieves the stack. See gtk_stack_switcher_set_stack().
	Stack() *StackClass
	// SetStack sets the stack to control.
	SetStack(stack Stack)
}

// StackSwitcherClass implements the StackSwitcher interface.
type StackSwitcherClass struct {
	*externglib.Object
	BoxClass
	BuildableInterface
	OrientableInterface
}

var _ StackSwitcher = (*StackSwitcherClass)(nil)

func wrapStackSwitcher(obj *externglib.Object) StackSwitcher {
	return &StackSwitcherClass{
		Object: obj,
		BoxClass: BoxClass{
			Object: obj,
			ContainerClass: ContainerClass{
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
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			OrientableInterface: OrientableInterface{
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

func marshalStackSwitcher(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStackSwitcher(obj), nil
}

// NewStackSwitcher: create a new StackSwitcher.
func NewStackSwitcher() *StackSwitcherClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_switcher_new()

	var _stackSwitcher *StackSwitcherClass // out

	_stackSwitcher = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*StackSwitcherClass)

	return _stackSwitcher
}

// Stack retrieves the stack. See gtk_stack_switcher_set_stack().
func (s *StackSwitcherClass) Stack() *StackClass {
	var _arg0 *C.GtkStackSwitcher // out
	var _cret *C.GtkStack         // in

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer((&StackSwitcher).Native()))

	_cret = C.gtk_stack_switcher_get_stack(_arg0)

	var _stack *StackClass // out

	_stack = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*StackClass)

	return _stack
}

// SetStack sets the stack to control.
func (s *StackSwitcherClass) SetStack(stack Stack) {
	var _arg0 *C.GtkStackSwitcher // out
	var _arg1 *C.GtkStack         // out

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer((&StackSwitcher).Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer((&Stack).Native()))

	C.gtk_stack_switcher_set_stack(_arg0, _arg1)
}
