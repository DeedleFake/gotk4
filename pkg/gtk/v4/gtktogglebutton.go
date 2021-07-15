// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_toggle_button_get_type()), F: marshalToggleButtonner},
	})
}

// ToggleButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToggleButtonOverrider interface {
	// Toggled emits the ::toggled signal on the GtkToggleButton.
	//
	// There is no good reason for an application ever to call this function.
	Toggled()
}

// ToggleButton: GtkToggleButton is a button which remains “pressed-in” when
// clicked.
//
// Clicking again will cause the toggle button to return to its normal state.
//
// A toggle button is created by calling either gtk.ToggleButton.New or
// gtk.ToggleButton.NewWithLabel. If using the former, it is advisable to pack a
// widget, (such as a GtkLabel and/or a GtkImage), into the toggle button’s
// container. (See gtk.Button for more information).
//
// The state of a GtkToggleButton can be set specifically using
// gtk.ToggleButton.SetActive(), and retrieved using
// gtk.ToggleButton.GetActive().
//
// To simply switch the state of a toggle button, use
// gtk.ToggleButton.Toggled().
//
//
// Grouping
//
// Toggle buttons can be grouped together, to form mutually exclusive groups -
// only one of the buttons can be toggled at a time, and toggling another one
// will switch the currently toggled one off.
//
// To add a GtkToggleButton to a group, use gtk.ToggleButton.SetGroup().
//
//
// CSS nodes
//
// GtkToggleButton has a single CSS node with name button. To differentiate it
// from a plain GtkButton, it gets the .toggle style class.
//
// Creating two GtkToggleButton widgets.
//
//    static void output_state (GtkToggleButton *source, gpointer user_data)
//    {
//      printf ("Active: d\n", gtk_toggle_button_get_active (source));
//    }
//
//    void make_toggles (void)
//    {
//      GtkWidget *window, *toggle1, *toggle2;
//      GtkWidget *box;
//      const char *text;
//
//      window = gtk_window_new ();
//      box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
//
//      text = "Hi, I’m a toggle button.";
//      toggle1 = gtk_toggle_button_new_with_label (text);
//
//      g_signal_connect (toggle1, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_box_append (GTK_BOX (box), toggle1);
//
//      text = "Hi, I’m a toggle button.";
//      toggle2 = gtk_toggle_button_new_with_label (text);
//      g_signal_connect (toggle2, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_box_append (GTK_BOX (box), toggle2);
//
//      gtk_window_set_child (GTK_WINDOW (window), box);
//      gtk_widget_show (window);
//    }
type ToggleButton struct {
	Button
}

var _ gextras.Nativer = (*ToggleButton)(nil)

func wrapToggleButton(obj *externglib.Object) *ToggleButton {
	return &ToggleButton{
		Button: Button{
			Widget: Widget{
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
			Actionable: Actionable{
				Widget: Widget{
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
			},
		},
	}
}

func marshalToggleButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapToggleButton(obj), nil
}

// NewToggleButton creates a new toggle button.
//
// A widget should be packed into the button, as in gtk.Button.New.
func NewToggleButton() *ToggleButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_toggle_button_new()

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// NewToggleButtonWithLabel creates a new toggle button with a text label.
func NewToggleButtonWithLabel(label string) *ToggleButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))

	_cret = C.gtk_toggle_button_new_with_label(_arg1)

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// NewToggleButtonWithMnemonic creates a new GtkToggleButton containing a label.
//
// The label will be created using gtk.Label.NewWithMnemonic, so underscores in
// label indicate the mnemonic for the button.
func NewToggleButtonWithMnemonic(label string) *ToggleButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))

	_cret = C.gtk_toggle_button_new_with_mnemonic(_arg1)

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// Active queries a GtkToggleButton and returns its current state.
//
// Returns TRUE if the toggle button is pressed in and FALSE if it is raised.
func (toggleButton *ToggleButton) Active() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))

	_cret = C.gtk_toggle_button_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the status of the toggle button.
//
// Set to TRUE if you want the GtkToggleButton to be “pressed in”, and FALSE to
// raise it.
//
// If the status of the button changes, this action causes the
// gtktogglebutton::toggled signal to be emitted.
func (toggleButton *ToggleButton) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_active(_arg0, _arg1)
}

// SetGroup adds self to the group of group.
//
// In a group of multiple toggle buttons, only one button can be active at a
// time.
//
// Setting up groups in a cycle leads to undefined behavior.
//
// Note that the same effect can be achieved via the gtk.Actionable API, by
// using the same action with parameter type and state type 's' for all buttons
// in the group, and giving each button its own target value.
func (toggleButton *ToggleButton) SetGroup(group *ToggleButton) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))
	_arg1 = (*C.GtkToggleButton)(unsafe.Pointer(group.Native()))

	C.gtk_toggle_button_set_group(_arg0, _arg1)
}

// Toggled emits the ::toggled signal on the GtkToggleButton.
//
// There is no good reason for an application ever to call this function.
func (toggleButton *ToggleButton) Toggled() {
	var _arg0 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))

	C.gtk_toggle_button_toggled(_arg0)
}
