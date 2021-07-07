// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_shortcut_controller_get_type()), F: marshalShortcutController},
	})
}

// ShortcutController: `GtkShortcutController` is an event controller that
// manages shortcuts.
//
// Most common shortcuts are using this controller implicitly, e.g. by adding a
// mnemonic underline to a `GtkLabel`, or by installing a key binding using
// gtk_widget_class_add_binding(), or by adding accelerators to global actions
// using gtk_application_set_accels_for_action().
//
// But it is possible to create your own shortcut controller, and add shortcuts
// to it.
//
// `GtkShortcutController` implements `GListModel` for querying the shortcuts
// that have been added to it.
//
//
// GtkShortcutController as a GtkBuildable
//
// `GtkShortcutControllers` can be creates in ui files to set up shortcuts in
// the same place as the widgets.
//
// An example of a UI definition fragment with `GtkShortcutController`: “`xml
// <object class='GtkButton'> <child> <object class='GtkShortcutController'>
// <property name='scope'>managed</property> <child> <object
// class='GtkShortcut'> <property
// name='trigger'>&amp;lt;Control&amp;gt;k</property> <property
// name='action'>activate</property> </object> </child> </object> </child>
// </object> “`
//
// This example creates a [class@Gtk.ActivateAction] for triggering the
// `activate` signal of the `GtkButton`. See
// [ctor@Gtk.ShortcutAction.parse_string] for the syntax for other kinds of
// `GtkShortcutAction`. See [ctor@Gtk.ShortcutTrigger.parse_string] to learn
// more about the syntax for triggers.
type ShortcutController interface {
	EventController

	// AsEventController casts the class to the EventController interface.
	AsEventController() EventController
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// GetCurrentEvent returns the event that is currently being handled by the
	// controller, and nil at other times.
	//
	// This method is inherited from EventController
	GetCurrentEvent() gdk.Event
	// GetCurrentEventDevice returns the device of the event that is currently
	// being handled by the controller, and nil otherwise.
	//
	// This method is inherited from EventController
	GetCurrentEventDevice() gdk.Device
	// GetCurrentEventState returns the modifier state of the event that is
	// currently being handled by the controller, and 0 otherwise.
	//
	// This method is inherited from EventController
	GetCurrentEventState() gdk.ModifierType
	// GetCurrentEventTime returns the timestamp of the event that is currently
	// being handled by the controller, and 0 otherwise.
	//
	// This method is inherited from EventController
	GetCurrentEventTime() uint32
	// GetName gets the name of @controller.
	//
	// This method is inherited from EventController
	GetName() string
	// GetPropagationLimit gets the propagation limit of the event controller.
	//
	// This method is inherited from EventController
	GetPropagationLimit() PropagationLimit
	// GetPropagationPhase gets the propagation phase at which @controller
	// handles events.
	//
	// This method is inherited from EventController
	GetPropagationPhase() PropagationPhase
	// GetWidget returns the Widget this controller relates to.
	//
	// This method is inherited from EventController
	GetWidget() Widget
	// Reset resets the @controller to a clean state.
	//
	// This method is inherited from EventController
	Reset()
	// SetName sets a name on the controller that can be used for debugging.
	//
	// This method is inherited from EventController
	SetName(name string)
	// SetPropagationLimit sets the event propagation limit on the event
	// controller.
	//
	// If the limit is set to GTK_LIMIT_SAME_NATIVE, the controller won't handle
	// events that are targeted at widgets on a different surface, such as
	// popovers.
	//
	// This method is inherited from EventController
	SetPropagationLimit(limit PropagationLimit)
	// SetPropagationPhase sets the propagation phase at which a controller
	// handles events.
	//
	// If @phase is GTK_PHASE_NONE, no automatic event handling will be
	// performed, but other additional gesture maintenance will.
	//
	// This method is inherited from EventController
	SetPropagationPhase(phase PropagationPhase)
	// GetBuildableID gets the ID of the @buildable object.
	//
	// `GtkBuilder` sets the name based on the ID attribute of the <object> tag
	// used to construct the @buildable.
	//
	// This method is inherited from Buildable
	GetBuildableID() string

	// AddShortcut adds @shortcut to the list of shortcuts handled by @self.
	//
	// If this controller uses an external shortcut list, this function does
	// nothing.
	AddShortcut(shortcut Shortcut)
	// MnemonicsModifiers gets the mnemonics modifiers for when this controller
	// activates its shortcuts.
	MnemonicsModifiers() gdk.ModifierType
	// Scope gets the scope for when this controller activates its shortcuts.
	// See gtk_shortcut_controller_set_scope() for details.
	Scope() ShortcutScope
	// RemoveShortcut removes @shortcut from the list of shortcuts handled by
	// @self.
	//
	// If @shortcut had not been added to @controller or this controller uses an
	// external shortcut list, this function does nothing.
	RemoveShortcut(shortcut Shortcut)
	// SetMnemonicsModifiers sets the controller to have the given
	// @mnemonics_modifiers.
	//
	// The mnemonics modifiers determines which modifiers need to be pressed to
	// allow activation of shortcuts with mnemonics triggers.
	//
	// GTK normally uses the Alt modifier for mnemonics, except in PopoverMenus,
	// where mnemonics can be triggered without any modifiers. It should be very
	// rarely necessary to change this, and doing so is likely to interfere with
	// other shortcuts.
	//
	// This value is only relevant for local shortcut controllers. Global and
	// managed shortcut controllers will have their shortcuts activated from
	// other places which have their own modifiers for activating mnemonics.
	SetMnemonicsModifiers(modifiers gdk.ModifierType)
	// SetScope sets the controller to have the given @scope.
	//
	// The scope allows shortcuts to be activated outside of the normal event
	// propagation. In particular, it allows installing global keyboard
	// shortcuts that can be activated even when a widget does not have focus.
	//
	// With GTK_SHORTCUT_SCOPE_LOCAL, shortcuts will only be activated when the
	// widget has focus.
	SetScope(scope ShortcutScope)
}

// shortcutController implements the ShortcutController interface.
type shortcutController struct {
	*externglib.Object
}

var _ ShortcutController = (*shortcutController)(nil)

// WrapShortcutController wraps a GObject to a type that implements
// interface ShortcutController. It is primarily used internally.
func WrapShortcutController(obj *externglib.Object) ShortcutController {
	return shortcutController{obj}
}

func marshalShortcutController(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutController(obj), nil
}

// NewShortcutController creates a new shortcut controller.
func NewShortcutController() ShortcutController {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_shortcut_controller_new()

	var _shortcutController ShortcutController // out

	_shortcutController = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(ShortcutController)

	return _shortcutController
}

func (s shortcutController) AsEventController() EventController {
	return WrapEventController(gextras.InternObject(s))
}

func (s shortcutController) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(s))
}

func (c shortcutController) GetCurrentEvent() gdk.Event {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEvent()
}

func (c shortcutController) GetCurrentEventDevice() gdk.Device {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventDevice()
}

func (c shortcutController) GetCurrentEventState() gdk.ModifierType {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventState()
}

func (c shortcutController) GetCurrentEventTime() uint32 {
	return WrapEventController(gextras.InternObject(c)).GetCurrentEventTime()
}

func (c shortcutController) GetName() string {
	return WrapEventController(gextras.InternObject(c)).GetName()
}

func (c shortcutController) GetPropagationLimit() PropagationLimit {
	return WrapEventController(gextras.InternObject(c)).GetPropagationLimit()
}

func (c shortcutController) GetPropagationPhase() PropagationPhase {
	return WrapEventController(gextras.InternObject(c)).GetPropagationPhase()
}

func (c shortcutController) GetWidget() Widget {
	return WrapEventController(gextras.InternObject(c)).GetWidget()
}

func (c shortcutController) Reset() {
	WrapEventController(gextras.InternObject(c)).Reset()
}

func (c shortcutController) SetName(name string) {
	WrapEventController(gextras.InternObject(c)).SetName(name)
}

func (c shortcutController) SetPropagationLimit(limit PropagationLimit) {
	WrapEventController(gextras.InternObject(c)).SetPropagationLimit(limit)
}

func (c shortcutController) SetPropagationPhase(phase PropagationPhase) {
	WrapEventController(gextras.InternObject(c)).SetPropagationPhase(phase)
}

func (b shortcutController) GetBuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).GetBuildableID()
}

func (s shortcutController) AddShortcut(shortcut Shortcut) {
	var _arg0 *C.GtkShortcutController // out
	var _arg1 *C.GtkShortcut           // out

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkShortcut)(unsafe.Pointer(shortcut.Native()))

	C.gtk_shortcut_controller_add_shortcut(_arg0, _arg1)
}

func (s shortcutController) MnemonicsModifiers() gdk.ModifierType {
	var _arg0 *C.GtkShortcutController // out
	var _cret C.GdkModifierType        // in

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_shortcut_controller_get_mnemonics_modifiers(_arg0)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

func (s shortcutController) Scope() ShortcutScope {
	var _arg0 *C.GtkShortcutController // out
	var _cret C.GtkShortcutScope       // in

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_shortcut_controller_get_scope(_arg0)

	var _shortcutScope ShortcutScope // out

	_shortcutScope = ShortcutScope(_cret)

	return _shortcutScope
}

func (s shortcutController) RemoveShortcut(shortcut Shortcut) {
	var _arg0 *C.GtkShortcutController // out
	var _arg1 *C.GtkShortcut           // out

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkShortcut)(unsafe.Pointer(shortcut.Native()))

	C.gtk_shortcut_controller_remove_shortcut(_arg0, _arg1)
}

func (s shortcutController) SetMnemonicsModifiers(modifiers gdk.ModifierType) {
	var _arg0 *C.GtkShortcutController // out
	var _arg1 C.GdkModifierType        // out

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkModifierType(modifiers)

	C.gtk_shortcut_controller_set_mnemonics_modifiers(_arg0, _arg1)
}

func (s shortcutController) SetScope(scope ShortcutScope) {
	var _arg0 *C.GtkShortcutController // out
	var _arg1 C.GtkShortcutScope       // out

	_arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkShortcutScope(scope)

	C.gtk_shortcut_controller_set_scope(_arg0, _arg1)
}
