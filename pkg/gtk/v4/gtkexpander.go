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
		{T: externglib.Type(C.gtk_expander_get_type()), F: marshalExpanderrer},
	})
}

// Expanderrer describes Expander's methods.
type Expanderrer interface {
	gextras.Objector

	Child() *Widget
	Expanded() bool
	Label() string
	LabelWidget() *Widget
	ResizeToplevel() bool
	UseMarkup() bool
	UseUnderline() bool
	SetChild(child Widgetter)
	SetExpanded(expanded bool)
	SetLabel(label string)
	SetLabelWidget(labelWidget Widgetter)
	SetResizeToplevel(resizeToplevel bool)
	SetUseMarkup(useMarkup bool)
	SetUseUnderline(useUnderline bool)
}

// Expander: `GtkExpander` allows the user to reveal its child by clicking on an
// expander triangle.
//
// !An example GtkExpander (expander.png)
//
// This is similar to the triangles used in a `GtkTreeView`.
//
// Normally you use an expander as you would use a frame; you create the child
// widget and use [method@Gtk.Expander.set_child] to add it to the expander.
// When the expander is toggled, it will take care of showing and hiding the
// child automatically.
//
//
// Special Usage
//
// There are situations in which you may prefer to show and hide the expanded
// widget yourself, such as when you want to actually create the widget at
// expansion time. In this case, create a `GtkExpander` but do not add a child
// to it. The expander widget has an [property@Gtk.Expander:expanded[ property
// which can be used to monitor its expansion state. You should watch this
// property with a signal connection as follows:
//
// “`c static void expander_callback (GObject *object, GParamSpec *param_spec,
// gpointer user_data) { GtkExpander *expander;
//
//    expander = GTK_EXPANDER (object);
//
//    if (gtk_expander_get_expanded (expander))
//      {
//        // Show or create widgets
//      }
//    else
//      {
//        // Hide or destroy widgets
//      }
//
// }
//
// static void create_expander (void) { GtkWidget *expander =
// gtk_expander_new_with_mnemonic ("_More Options"); g_signal_connect (expander,
// "notify::expanded", G_CALLBACK (expander_callback), NULL);
//
//    // ...
//
// } “`
//
//
// GtkExpander as GtkBuildable
//
// The `GtkExpander` implementation of the `GtkBuildable` interface supports
// placing a child in the label position by specifying “label” as the “type”
// attribute of a <child> element. A normal content child can be specified
// without specifying a <child> type attribute.
//
// An example of a UI definition fragment with GtkExpander:
//
// “`xml <object class="GtkExpander"> <child type="label"> <object
// class="GtkLabel" id="expander-label"/> </child> <child> <object
// class="GtkEntry" id="expander-content"/> </child> </object> “`
//
//
// CSS nodes
//
// “` expander ╰── box ├── title │ ├── arrow │ ╰── <label widget> ╰── <child> “`
//
// `GtkExpander` has three CSS nodes, the main node with the name expander, a
// subnode with name title and node below it with name arrow. The arrow of an
// expander that is showing its child gets the :checked pseudoclass added to it.
//
//
// Accessibility
//
// `GtkExpander` uses the GTK_ACCESSIBLE_ROLE_BUTTON role.
type Expander struct {
	*externglib.Object

	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ Expanderrer = (*Expander)(nil)

func wrapExpanderrer(obj *externglib.Object) Expanderrer {
	return &Expander{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
	}
}

func marshalExpanderrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapExpanderrer(obj), nil
}

// NewExpander creates a new expander using @label as the text of the label.
func NewExpander(label string) *Expander {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_expander_new(_arg1)

	var _expander *Expander // out

	_expander = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Expander)

	return _expander
}

// NewExpanderWithMnemonic creates a new expander using @label as the text of
// the label.
//
// If characters in @label are preceded by an underscore, they are underlined.
// If you need a literal underscore character in a label, use “__” (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic.
//
// Pressing Alt and that key activates the button.
func NewExpanderWithMnemonic(label string) *Expander {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_expander_new_with_mnemonic(_arg1)

	var _expander *Expander // out

	_expander = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Expander)

	return _expander
}

// Child gets the child widget of @expander.
func (expander *Expander) Child() *Widget {
	var _arg0 *C.GtkExpander // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_child(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// Expanded queries a Expander and returns its current state.
//
// Returns true if the child widget is revealed.
func (expander *Expander) Expanded() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_expanded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Label fetches the text from a label widget.
//
// This is including any embedded underlines indicating mnemonics and Pango
// markup, as set by [method@Gtk.Expander.set_label]. If the label text has not
// been set the return value will be nil. This will be the case if you create an
// empty button with gtk_button_new() to use as a container.
func (expander *Expander) Label() string {
	var _arg0 *C.GtkExpander // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// LabelWidget retrieves the label widget for the frame.
func (expander *Expander) LabelWidget() *Widget {
	var _arg0 *C.GtkExpander // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_label_widget(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// ResizeToplevel returns whether the expander will resize the toplevel widget
// containing the expander upon resizing and collpasing.
func (expander *Expander) ResizeToplevel() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_resize_toplevel(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseMarkup returns whether the label’s text is interpreted as Pango markup.
func (expander *Expander) UseMarkup() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_use_markup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline returns whether an underline in the text indicates a mnemonic.
func (expander *Expander) UseUnderline() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetChild sets the child widget of @expander.
func (expander *Expander) SetChild(child Widgetter) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_expander_set_child(_arg0, _arg1)
}

// SetExpanded sets the state of the expander.
//
// Set to true, if you want the child widget to be revealed, and false if you
// want the child widget to be hidden.
func (expander *Expander) SetExpanded(expanded bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if expanded {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_expanded(_arg0, _arg1)
}

// SetLabel sets the text of the label of the expander to @label.
//
// This will also clear any previously set labels.
func (expander *Expander) SetLabel(label string) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_expander_set_label(_arg0, _arg1)
}

// SetLabelWidget: set the label widget for the expander.
//
// This is the widget that will appear embedded alongside the expander arrow.
func (expander *Expander) SetLabelWidget(labelWidget Widgetter) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_expander_set_label_widget(_arg0, _arg1)
}

// SetResizeToplevel sets whether the expander will resize the toplevel widget
// containing the expander upon resizing and collpasing.
func (expander *Expander) SetResizeToplevel(resizeToplevel bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if resizeToplevel {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_resize_toplevel(_arg0, _arg1)
}

// SetUseMarkup sets whether the text of the label contains Pango markup.
func (expander *Expander) SetUseMarkup(useMarkup bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if useMarkup {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_markup(_arg0, _arg1)
}

// SetUseUnderline: if true, an underline in the text indicates a mnemonic.
func (expander *Expander) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_underline(_arg0, _arg1)
}
