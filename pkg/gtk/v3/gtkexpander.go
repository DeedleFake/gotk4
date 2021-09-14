// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_expander_get_type()), F: marshalExpanderer},
	})
}

// ExpanderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ExpanderOverrider interface {
	Activate()
}

// Expander allows the user to hide or show its child by clicking on an expander
// triangle similar to the triangles used in a TreeView.
//
// Normally you use an expander as you would use any other descendant of Bin;
// you create the child widget and use gtk_container_add() to add it to the
// expander. When the expander is toggled, it will take care of showing and
// hiding the child automatically.
//
//
// Special Usage
//
// There are situations in which you may prefer to show and hide the expanded
// widget yourself, such as when you want to actually create the widget at
// expansion time. In this case, create a Expander but do not add a child to it.
// The expander widget has an Expander:expanded property which can be used to
// monitor its expansion state. You should watch this property with a signal
// connection as follows:
//
//    expander
//    ├── title
//    │   ├── arrow
//    │   ╰── <label widget>
//    ╰── <child>
//
// GtkExpander has three CSS nodes, the main node with the name expander, a
// subnode with name title and node below it with name arrow. The arrow of an
// expander that is showing its child gets the :checked pseudoclass added to it.
type Expander struct {
	Bin
}

func wrapExpander(obj *externglib.Object) *Expander {
	return &Expander{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
		},
	}
}

func marshalExpanderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapExpander(obj), nil
}

// NewExpander creates a new expander using label as the text of the label.
func NewExpander(label string) *Expander {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_expander_new(_arg1)
	runtime.KeepAlive(label)

	var _expander *Expander // out

	_expander = wrapExpander(externglib.Take(unsafe.Pointer(_cret)))

	return _expander
}

// NewExpanderWithMnemonic creates a new expander using label as the text of the
// label. If characters in label are preceded by an underscore, they are
// underlined. If you need a literal underscore character in a label, use “__”
// (two underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. Pressing Alt and that key activates the
// button.
func NewExpanderWithMnemonic(label string) *Expander {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_expander_new_with_mnemonic(_arg1)
	runtime.KeepAlive(label)

	var _expander *Expander // out

	_expander = wrapExpander(externglib.Take(unsafe.Pointer(_cret)))

	return _expander
}

// Expanded queries a Expander and returns its current state. Returns TRUE if
// the child widget is revealed.
//
// See gtk_expander_set_expanded().
func (expander *Expander) Expanded() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_expanded(_arg0)
	runtime.KeepAlive(expander)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Label fetches the text from a label widget including any embedded underlines
// indicating mnemonics and Pango markup, as set by gtk_expander_set_label(). If
// the label text has not been set the return value will be NULL. This will be
// the case if you create an empty button with gtk_button_new() to use as a
// container.
//
// Note that this function behaved differently in versions prior to 2.14 and
// used to return the label text stripped of embedded underlines indicating
// mnemonics and Pango markup. This problem can be avoided by fetching the label
// text directly from the label widget.
func (expander *Expander) Label() string {
	var _arg0 *C.GtkExpander // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_label(_arg0)
	runtime.KeepAlive(expander)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LabelFill returns whether the label widget will fill all available horizontal
// space allocated to expander.
func (expander *Expander) LabelFill() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_label_fill(_arg0)
	runtime.KeepAlive(expander)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LabelWidget retrieves the label widget for the frame. See
// gtk_expander_set_label_widget().
func (expander *Expander) LabelWidget() Widgetter {
	var _arg0 *C.GtkExpander // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_label_widget(_arg0)
	runtime.KeepAlive(expander)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// ResizeToplevel returns whether the expander will resize the toplevel widget
// containing the expander upon resizing and collpasing.
func (expander *Expander) ResizeToplevel() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_resize_toplevel(_arg0)
	runtime.KeepAlive(expander)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Spacing gets the value set by gtk_expander_set_spacing().
//
// Deprecated: Use margins on the child instead.
func (expander *Expander) Spacing() int {
	var _arg0 *C.GtkExpander // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_spacing(_arg0)
	runtime.KeepAlive(expander)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UseMarkup returns whether the label’s text is interpreted as marked up with
// the [Pango text markup language][PangoMarkupFormat]. See
// gtk_expander_set_use_markup().
func (expander *Expander) UseMarkup() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_use_markup(_arg0)
	runtime.KeepAlive(expander)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline returns whether an embedded underline in the expander label
// indicates a mnemonic. See gtk_expander_set_use_underline().
func (expander *Expander) UseUnderline() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))

	_cret = C.gtk_expander_get_use_underline(_arg0)
	runtime.KeepAlive(expander)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpanded sets the state of the expander. Set to TRUE, if you want the
// child widget to be revealed, and FALSE if you want the child widget to be
// hidden.
func (expander *Expander) SetExpanded(expanded bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if expanded {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_expanded(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(expanded)
}

// SetLabel sets the text of the label of the expander to label.
//
// This will also clear any previously set labels.
func (expander *Expander) SetLabel(label string) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_expander_set_label(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(label)
}

// SetLabelFill sets whether the label widget should fill all available
// horizontal space allocated to expander.
//
// Note that this function has no effect since 3.20.
func (expander *Expander) SetLabelFill(labelFill bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if labelFill {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_label_fill(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(labelFill)
}

// SetLabelWidget: set the label widget for the expander. This is the widget
// that will appear embedded alongside the expander arrow.
func (expander *Expander) SetLabelWidget(labelWidget Widgetter) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if labelWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))
	}

	C.gtk_expander_set_label_widget(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(labelWidget)
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
	runtime.KeepAlive(expander)
	runtime.KeepAlive(resizeToplevel)
}

// SetSpacing sets the spacing field of expander, which is the number of pixels
// to place between expander and the child.
//
// Deprecated: Use margins on the child instead.
func (expander *Expander) SetSpacing(spacing int) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	_arg1 = C.gint(spacing)

	C.gtk_expander_set_spacing(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(spacing)
}

// SetUseMarkup sets whether the text of the label contains markup in [Pango’s
// text markup language][PangoMarkupFormat]. See gtk_label_set_markup().
func (expander *Expander) SetUseMarkup(useMarkup bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if useMarkup {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_markup(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(useMarkup)
}

// SetUseUnderline: if true, an underline in the text of the expander label
// indicates the next character should be used for the mnemonic accelerator key.
func (expander *Expander) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(expander.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(useUnderline)
}
