// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_pad_action_type_get_type()), F: marshalPadActionType},
		{T: externglib.Type(C.gtk_pad_controller_get_type()), F: marshalPadControllerer},
	})
}

// PadActionType: type of a pad action.
type PadActionType int

const (
	// Button: action is triggered by a pad button
	PadActionTypeButton PadActionType = iota
	// Ring: action is triggered by a pad ring
	PadActionTypeRing
	// Strip: action is triggered by a pad strip
	PadActionTypeStrip
)

func marshalPadActionType(p uintptr) (interface{}, error) {
	return PadActionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PadController is an event controller for the pads found in drawing tablets
// (The collection of buttons and tactile sensors often found around the
// stylus-sensitive area).
//
// These buttons and sensors have no implicit meaning, and by default they
// perform no action, this event controller is provided to map those to #GAction
// objects, thus letting the application give those a more semantic meaning.
//
// Buttons and sensors are not constrained to triggering a single action, some
// GDK_SOURCE_TABLET_PAD devices feature multiple "modes", all these input
// elements have one current mode, which may determine the final action being
// triggered. Pad devices often divide buttons and sensors into groups, all
// elements in a group share the same current mode, but different groups may
// have different modes. See gdk_device_pad_get_n_groups() and
// gdk_device_pad_get_group_n_modes().
//
// Each of the actions that a given button/strip/ring performs for a given mode
// is defined by PadActionEntry, it contains an action name that will be looked
// up in the given Group and activated whenever the specified input element and
// mode are triggered.
//
// A simple example of PadController usage, assigning button 1 in all modes and
// pad devices to an "invert-selection" action:
//
//      GtkPadActionEntry *pad_actions[] = {
//        { GTK_PAD_ACTION_BUTTON, 1, -1, "Invert selection", "pad-actions.invert-selection" },
//        …
//      };
//
//      …
//      action_group = g_simple_action_group_new ();
//      action = g_simple_action_new ("pad-actions.invert-selection", NULL);
//      g_signal_connect (action, "activate", on_invert_selection_activated, NULL);
//      g_action_map_add_action (G_ACTION_MAP (action_group), action);
//      …
//      pad_controller = gtk_pad_controller_new (window, action_group, NULL);
//
// The actions belonging to rings/strips will be activated with a parameter of
// type G_VARIANT_TYPE_DOUBLE bearing the value of the given axis, it is
// required that those are made stateful and accepting this Type.
type PadController struct {
	EventController
}

var _ gextras.Nativer = (*PadController)(nil)

func wrapPadController(obj *externglib.Object) *PadController {
	return &PadController{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalPadControllerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPadController(obj), nil
}

// NewPadController creates a new PadController that will associate events from
// pad to actions. A NULL pad may be provided so the controller manages all pad
// devices generically, it is discouraged to mix PadController objects with NULL
// and non-NULL pad argument on the same window, as execution order is not
// guaranteed.
//
// The PadController is created with no mapped actions. In order to map pad
// events to actions, use gtk_pad_controller_set_action_entries() or
// gtk_pad_controller_set_action().
func NewPadController(window *Window, group gio.ActionGrouper, pad gdk.Devicer) *PadController {
	var _arg1 *C.GtkWindow        // out
	var _arg2 *C.GActionGroup     // out
	var _arg3 *C.GdkDevice        // out
	var _cret *C.GtkPadController // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = (*C.GActionGroup)(unsafe.Pointer((group).(gextras.Nativer).Native()))
	_arg3 = (*C.GdkDevice)(unsafe.Pointer((pad).(gextras.Nativer).Native()))

	_cret = C.gtk_pad_controller_new(_arg1, _arg2, _arg3)

	var _padController *PadController // out

	_padController = wrapPadController(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _padController
}

// SetAction adds an individual action to controller. This action will only be
// activated if the given button/ring/strip number in index is interacted while
// the current mode is mode. -1 may be used for simple cases, so the action is
// triggered on all modes.
//
// The given label should be considered user-visible, so internationalization
// rules apply. Some windowing systems may be able to use those for user
// feedback.
func (controller *PadController) SetAction(typ PadActionType, index int, mode int, label string, actionName string) {
	var _arg0 *C.GtkPadController // out
	var _arg1 C.GtkPadActionType  // out
	var _arg2 C.gint              // out
	var _arg3 C.gint              // out
	var _arg4 *C.gchar            // out
	var _arg5 *C.gchar            // out

	_arg0 = (*C.GtkPadController)(unsafe.Pointer(controller.Native()))
	_arg1 = C.GtkPadActionType(typ)
	_arg2 = C.gint(index)
	_arg3 = C.gint(mode)
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))

	C.gtk_pad_controller_set_action(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// SetActionEntries: this is a convenience function to add a group of action
// entries on controller. See PadActionEntry and
// gtk_pad_controller_set_action().
func (controller *PadController) SetActionEntries(entries []PadActionEntry) {
	var _arg0 *C.GtkPadController // out
	var _arg1 *C.GtkPadActionEntry
	var _arg2 C.gint

	_arg0 = (*C.GtkPadController)(unsafe.Pointer(controller.Native()))
	_arg2 = (C.gint)(len(entries))
	if len(entries) > 0 {
		_arg1 = (*C.GtkPadActionEntry)(unsafe.Pointer(&entries[0]))
	}

	C.gtk_pad_controller_set_action_entries(_arg0, _arg1, _arg2)
}

// PadActionEntry: struct defining a pad action entry.
type PadActionEntry struct {
	nocopy gextras.NoCopy
	native *C.GtkPadActionEntry
}

// Type: type of pad feature that will trigger this action entry.
func (p *PadActionEntry) Type() PadActionType {
	var v PadActionType // out
	v = PadActionType(p.native._type)
	return v
}

// Index: 0-indexed button/ring/strip number that will trigger this action
// entry.
func (p *PadActionEntry) Index() int {
	var v int // out
	v = int(p.native.index)
	return v
}

// Mode: mode that will trigger this action entry, or -1 for all modes.
func (p *PadActionEntry) Mode() int {
	var v int // out
	v = int(p.native.mode)
	return v
}

// Label: human readable description of this action entry, this string should be
// deemed user-visible.
func (p *PadActionEntry) Label() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(p.native.label)))
	return v
}

// ActionName: action name that will be activated in the Group.
func (p *PadActionEntry) ActionName() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(p.native.action_name)))
	return v
}
