// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
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
// extern void _gotk4_gtk3_ToggleButtonClass_toggled(GtkToggleButton*);
// extern void _gotk4_gtk3_ToggleButton_ConnectToggled(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeToggleButton = coreglib.Type(C.gtk_toggle_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeToggleButton, F: marshalToggleButton},
	})
}

// ToggleButtonOverrider contains methods that are overridable.
type ToggleButtonOverrider interface {
	// Toggled emits the ToggleButton::toggled signal on the ToggleButton. There
	// is no good reason for an application ever to call this function.
	Toggled()
}

// ToggleButton is a Button which will remain “pressed-in” when clicked.
// Clicking again will cause the toggle button to return to its normal state.
//
// A toggle button is created by calling either gtk_toggle_button_new() or
// gtk_toggle_button_new_with_label(). If using the former, it is advisable to
// pack a widget, (such as a Label and/or a Image), into the toggle button’s
// container. (See Button for more information).
//
// The state of a ToggleButton can be set specifically using
// gtk_toggle_button_set_active(), and retrieved using
// gtk_toggle_button_get_active().
//
// To simply switch the state of a toggle button, use
// gtk_toggle_button_toggled().
//
//
// CSS nodes
//
// GtkToggleButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .toggle style class.
//
// Creating two ToggleButton widgets.
//
//    static void output_state (GtkToggleButton *source, gpointer user_data) {
//      printf ("Active: d\n", gtk_toggle_button_get_active (source));
//    }
//
//    void make_toggles (void) {
//      GtkWidget *window, *toggle1, *toggle2;
//      GtkWidget *box;
//      const char *text;
//
//      window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
//
//      text = "Hi, I’m a toggle button.";
//      toggle1 = gtk_toggle_button_new_with_label (text);
//
//      // Makes this toggle button invisible
//      gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle1),
//                                  TRUE);
//
//      g_signal_connect (toggle1, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_container_add (GTK_CONTAINER (box), toggle1);
//
//      text = "Hi, I’m a toggle button.";
//      toggle2 = gtk_toggle_button_new_with_label (text);
//      gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle2),
//                                  FALSE);
//      g_signal_connect (toggle2, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_container_add (GTK_CONTAINER (box), toggle2);
//
//      gtk_container_add (GTK_CONTAINER (window), box);
//      gtk_widget_show_all (window);
//    }.
type ToggleButton struct {
	_ [0]func() // equal guard
	Button
}

var (
	_ Binner            = (*ToggleButton)(nil)
	_ coreglib.Objector = (*ToggleButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:        GTypeToggleButton,
		GoType:       reflect.TypeOf((*ToggleButton)(nil)),
		InitClass:    initClassToggleButton,
		ClassSize:    uint16(unsafe.Sizeof(C.GtkToggleButton{})),
		InstanceSize: uint16(unsafe.Sizeof(C.GtkToggleButtonClass{})),
	})
}

func initClassToggleButton(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GtkToggleButtonClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface{ Toggled() }); ok {
		pclass.toggled = (*[0]byte)(C._gotk4_gtk3_ToggleButtonClass_toggled)
	}
}

//export _gotk4_gtk3_ToggleButtonClass_toggled
func _gotk4_gtk3_ToggleButtonClass_toggled(arg0 *C.GtkToggleButton) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface{ Toggled() })

	iface.Toggled()
}

func wrapToggleButton(obj *coreglib.Object) *ToggleButton {
	return &ToggleButton{
		Button: Button{
			Bin: Bin{
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
			},
			Object: obj,
			Actionable: Actionable{
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
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalToggleButton(p uintptr) (interface{}, error) {
	return wrapToggleButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ToggleButton_ConnectToggled
func _gotk4_gtk3_ToggleButton_ConnectToggled(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectToggled: should be connected if you wish to perform an action whenever
// the ToggleButton's state is changed.
func (toggleButton *ToggleButton) ConnectToggled(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(toggleButton, "toggled", false, unsafe.Pointer(C._gotk4_gtk3_ToggleButton_ConnectToggled), f)
}

// NewToggleButton creates a new toggle button. A widget should be packed into
// the button, as in gtk_button_new().
//
// The function returns the following values:
//
//    - toggleButton: new toggle button.
//
func NewToggleButton() *ToggleButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_toggle_button_new()

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// NewToggleButtonWithLabel creates a new toggle button with a text label.
//
// The function takes the following parameters:
//
//    - label: string containing the message to be placed in the toggle button.
//
// The function returns the following values:
//
//    - toggleButton: new toggle button.
//
func NewToggleButtonWithLabel(label string) *ToggleButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_label(_arg1)
	runtime.KeepAlive(label)

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// NewToggleButtonWithMnemonic creates a new ToggleButton containing a label.
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in label indicate the mnemonic for the button.
//
// The function takes the following parameters:
//
//    - label: text of the button, with an underscore in front of the mnemonic
//      character.
//
// The function returns the following values:
//
//    - toggleButton: new ToggleButton.
//
func NewToggleButtonWithMnemonic(label string) *ToggleButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_mnemonic(_arg1)
	runtime.KeepAlive(label)

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// Active queries a ToggleButton and returns its current state. Returns TRUE if
// the toggle button is pressed in and FALSE if it is raised.
//
// The function returns the following values:
//
//    - ok: #gboolean value.
//
func (toggleButton *ToggleButton) Active() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(coreglib.InternObject(toggleButton).Native()))

	_cret = C.gtk_toggle_button_get_active(_arg0)
	runtime.KeepAlive(toggleButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Inconsistent gets the value set by gtk_toggle_button_set_inconsistent().
//
// The function returns the following values:
//
//    - ok: TRUE if the button is displayed as inconsistent, FALSE otherwise.
//
func (toggleButton *ToggleButton) Inconsistent() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(coreglib.InternObject(toggleButton).Native()))

	_cret = C.gtk_toggle_button_get_inconsistent(_arg0)
	runtime.KeepAlive(toggleButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Mode retrieves whether the button is displayed as a separate indicator and
// label. See gtk_toggle_button_set_mode().
//
// The function returns the following values:
//
//    - ok: TRUE if the togglebutton is drawn as a separate indicator and label.
//
func (toggleButton *ToggleButton) Mode() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(coreglib.InternObject(toggleButton).Native()))

	_cret = C.gtk_toggle_button_get_mode(_arg0)
	runtime.KeepAlive(toggleButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the status of the toggle button. Set to TRUE if you want the
// GtkToggleButton to be “pressed in”, and FALSE to raise it. This action causes
// the ToggleButton::toggled signal and the Button::clicked signal to be
// emitted.
//
// The function takes the following parameters:
//
//    - isActive: TRUE or FALSE.
//
func (toggleButton *ToggleButton) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(coreglib.InternObject(toggleButton).Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_active(_arg0, _arg1)
	runtime.KeepAlive(toggleButton)
	runtime.KeepAlive(isActive)
}

// SetInconsistent: if the user has selected a range of elements (such as some
// text or spreadsheet cells) that are affected by a toggle button, and the
// current values in that range are inconsistent, you may want to display the
// toggle in an “in between” state. This function turns on “in between” display.
// Normally you would turn off the inconsistent state again if the user toggles
// the toggle button. This has to be done manually,
// gtk_toggle_button_set_inconsistent() only affects visual appearance, it
// doesn’t affect the semantics of the button.
//
// The function takes the following parameters:
//
//    - setting: TRUE if state is inconsistent.
//
func (toggleButton *ToggleButton) SetInconsistent(setting bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(coreglib.InternObject(toggleButton).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_inconsistent(_arg0, _arg1)
	runtime.KeepAlive(toggleButton)
	runtime.KeepAlive(setting)
}

// SetMode sets whether the button is displayed as a separate indicator and
// label. You can call this function on a checkbutton or a radiobutton with
// draw_indicator = FALSE to make the button look like a normal button.
//
// This can be used to create linked strip of buttons that work like a
// StackSwitcher.
//
// This function only affects instances of classes like CheckButton and
// RadioButton that derive from ToggleButton, not instances of ToggleButton
// itself.
//
// The function takes the following parameters:
//
//    - drawIndicator: if TRUE, draw the button as a separate indicator and
//      label; if FALSE, draw the button like a normal button.
//
func (toggleButton *ToggleButton) SetMode(drawIndicator bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(coreglib.InternObject(toggleButton).Native()))
	if drawIndicator {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_mode(_arg0, _arg1)
	runtime.KeepAlive(toggleButton)
	runtime.KeepAlive(drawIndicator)
}

// Toggled emits the ToggleButton::toggled signal on the ToggleButton. There is
// no good reason for an application ever to call this function.
func (toggleButton *ToggleButton) Toggled() {
	var _arg0 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(coreglib.InternObject(toggleButton).Native()))

	C.gtk_toggle_button_toggled(_arg0)
	runtime.KeepAlive(toggleButton)
}
