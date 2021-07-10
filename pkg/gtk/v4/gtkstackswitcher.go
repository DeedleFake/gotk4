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
		{T: externglib.Type(C.gtk_stack_switcher_get_type()), F: marshalStackSwitcher},
	})
}

// StackSwitcher: the `GtkStackSwitcher` shows a row of buttons to switch
// between `GtkStack` pages.
//
// !An example GtkStackSwitcher (stackswitcher.png)
//
// It acts as a controller for the associated `GtkStack`.
//
// All the content for the buttons comes from the properties of the stacks
// [class@Gtk.StackPage] objects; the button visibility in a `GtkStackSwitcher`
// widget is controlled by the visibility of the child in the `GtkStack`.
//
// It is possible to associate multiple `GtkStackSwitcher` widgets with the same
// `GtkStack` widget.
//
//
// CSS nodes
//
// `GtkStackSwitcher` has a single CSS node named stackswitcher and style class
// .stack-switcher.
//
// When circumstances require it, `GtkStackSwitcher` adds the .needs-attention
// style class to the widgets representing the stack pages.
//
//
// Accessibility
//
// `GtkStackSwitcher` uses the GTK_ACCESSIBLE_ROLE_TAB_LIST role and uses the
// GTK_ACCESSIBLE_ROLE_TAB for its buttons.
type StackSwitcher interface {
	gextras.Objector

	// Stack retrieves the stack.
	Stack() *StackClass
	// SetStack sets the stack to control.
	SetStack(stack Stack)
}

// StackSwitcherClass implements the StackSwitcher interface.
type StackSwitcherClass struct {
	*externglib.Object
	WidgetClass
	AccessibleIface
	BuildableIface
	ConstraintTargetIface
}

var _ StackSwitcher = (*StackSwitcherClass)(nil)

func wrapStackSwitcher(obj *externglib.Object) StackSwitcher {
	return &StackSwitcherClass{
		Object: obj,
		WidgetClass: WidgetClass{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			AccessibleIface: AccessibleIface{
				Object: obj,
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
			ConstraintTargetIface: ConstraintTargetIface{
				Object: obj,
			},
		},
		AccessibleIface: AccessibleIface{
			Object: obj,
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
		ConstraintTargetIface: ConstraintTargetIface{
			Object: obj,
		},
	}
}

func marshalStackSwitcher(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStackSwitcher(obj), nil
}

// NewStackSwitcher: create a new `GtkStackSwitcher`.
func NewStackSwitcher() *StackSwitcherClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_switcher_new()

	var _stackSwitcher *StackSwitcherClass // out

	_stackSwitcher = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*StackSwitcherClass)

	return _stackSwitcher
}

// Stack retrieves the stack.
func (switcher *StackSwitcherClass) Stack() *StackClass {
	var _arg0 *C.GtkStackSwitcher // out
	var _cret *C.GtkStack         // in

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(switcher.Native()))

	_cret = C.gtk_stack_switcher_get_stack(_arg0)

	var _stack *StackClass // out

	_stack = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*StackClass)

	return _stack
}

// SetStack sets the stack to control.
func (switcher *StackSwitcherClass) SetStack(stack Stack) {
	var _arg0 *C.GtkStackSwitcher // out
	var _arg1 *C.GtkStack         // out

	_arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(switcher.Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_switcher_set_stack(_arg0, _arg1)
}
