// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_ShortcutFunc(GtkWidget*, GVariant*, gpointer);
// extern void callbackDelete(gpointer);
import "C"

// GType values.
var (
	GTypeShortcutActionFlags = coreglib.Type(C.gtk_shortcut_action_flags_get_type())
	GTypeActivateAction      = coreglib.Type(C.gtk_activate_action_get_type())
	GTypeCallbackAction      = coreglib.Type(C.gtk_callback_action_get_type())
	GTypeMnemonicAction      = coreglib.Type(C.gtk_mnemonic_action_get_type())
	GTypeNamedAction         = coreglib.Type(C.gtk_named_action_get_type())
	GTypeNothingAction       = coreglib.Type(C.gtk_nothing_action_get_type())
	GTypeShortcutAction      = coreglib.Type(C.gtk_shortcut_action_get_type())
	GTypeSignalAction        = coreglib.Type(C.gtk_signal_action_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutActionFlags, F: marshalShortcutActionFlags},
		coreglib.TypeMarshaler{T: GTypeActivateAction, F: marshalActivateAction},
		coreglib.TypeMarshaler{T: GTypeCallbackAction, F: marshalCallbackAction},
		coreglib.TypeMarshaler{T: GTypeMnemonicAction, F: marshalMnemonicAction},
		coreglib.TypeMarshaler{T: GTypeNamedAction, F: marshalNamedAction},
		coreglib.TypeMarshaler{T: GTypeNothingAction, F: marshalNothingAction},
		coreglib.TypeMarshaler{T: GTypeShortcutAction, F: marshalShortcutAction},
		coreglib.TypeMarshaler{T: GTypeSignalAction, F: marshalSignalAction},
	})
}

// ShortcutActionFlags: list of flags that can be passed to action activation.
//
// More flags may be added in the future.
type ShortcutActionFlags C.guint

const (
	// ShortcutActionExclusive: action is the only action that can be activated.
	// If this flag is not set, a future activation may select a different
	// action.
	ShortcutActionExclusive ShortcutActionFlags = 0b1
)

func marshalShortcutActionFlags(p uintptr) (interface{}, error) {
	return ShortcutActionFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ShortcutActionFlags.
func (s ShortcutActionFlags) String() string {
	if s == 0 {
		return "ShortcutActionFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(23)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case ShortcutActionExclusive:
			builder.WriteString("Exclusive|")
		default:
			builder.WriteString(fmt.Sprintf("ShortcutActionFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s ShortcutActionFlags) Has(other ShortcutActionFlags) bool {
	return (s & other) == other
}

// ShortcutFunc: prototype for shortcuts based on user callbacks.
type ShortcutFunc func(widget Widgetter, args *glib.Variant) (ok bool)

//export _gotk4_gtk4_ShortcutFunc
func _gotk4_gtk4_ShortcutFunc(arg1 *C.GtkWidget, arg2 *C.GVariant, arg3 C.gpointer) (cret C.gboolean) {
	var fn ShortcutFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ShortcutFunc)
	}

	var _widget Widgetter   // out
	var _args *glib.Variant // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}
	if arg2 != nil {
		_args = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(arg2)))
		C.g_variant_ref(arg2)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_args)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	ok := fn(_widget, _args)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ActivateActionOverrider contains methods that are overridable.
type ActivateActionOverrider interface {
}

// ActivateAction: GtkShortcutAction that calls gtk_widget_activate().
type ActivateAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*ActivateAction)(nil)
)

func initClassActivateAction(gclass unsafe.Pointer, goval any) {
}

func wrapActivateAction(obj *coreglib.Object) *ActivateAction {
	return &ActivateAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalActivateAction(p uintptr) (interface{}, error) {
	return wrapActivateAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ActivateActionGet gets the activate action.
//
// This is an action that calls gtk_widget_activate() on the given widget upon
// activation.
//
// The function returns the following values:
//
//    - activateAction: activate action.
//
func ActivateActionGet() *ActivateAction {
	var _cret *C.GtkShortcutAction // in

	_cret = C.gtk_activate_action_get()

	var _activateAction *ActivateAction // out

	_activateAction = wrapActivateAction(coreglib.Take(unsafe.Pointer(_cret)))

	return _activateAction
}

// CallbackActionOverrider contains methods that are overridable.
type CallbackActionOverrider interface {
}

// CallbackAction: GtkShortcutAction that invokes a callback.
type CallbackAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*CallbackAction)(nil)
)

func initClassCallbackAction(gclass unsafe.Pointer, goval any) {
}

func wrapCallbackAction(obj *coreglib.Object) *CallbackAction {
	return &CallbackAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalCallbackAction(p uintptr) (interface{}, error) {
	return wrapCallbackAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCallbackAction: create a custom action that calls the given callback when
// activated.
//
// The function takes the following parameters:
//
//    - callback (optional) to call.
//
// The function returns the following values:
//
//    - callbackAction: new shortcut action.
//
func NewCallbackAction(callback ShortcutFunc) *CallbackAction {
	var _arg1 C.GtkShortcutFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify
	var _cret *C.GtkShortcutAction // in

	if callback != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_ShortcutFunc)
		_arg2 = C.gpointer(gbox.Assign(callback))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.gtk_callback_action_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(callback)

	var _callbackAction *CallbackAction // out

	_callbackAction = wrapCallbackAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _callbackAction
}

// MnemonicActionOverrider contains methods that are overridable.
type MnemonicActionOverrider interface {
}

// MnemonicAction: GtkShortcutAction that calls gtk_widget_mnemonic_activate().
type MnemonicAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*MnemonicAction)(nil)
)

func initClassMnemonicAction(gclass unsafe.Pointer, goval any) {
}

func wrapMnemonicAction(obj *coreglib.Object) *MnemonicAction {
	return &MnemonicAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalMnemonicAction(p uintptr) (interface{}, error) {
	return wrapMnemonicAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// MnemonicActionGet gets the mnemonic action.
//
// This is an action that calls gtk_widget_mnemonic_activate() on the given
// widget upon activation.
//
// The function returns the following values:
//
//    - mnemonicAction: mnemonic action.
//
func MnemonicActionGet() *MnemonicAction {
	var _cret *C.GtkShortcutAction // in

	_cret = C.gtk_mnemonic_action_get()

	var _mnemonicAction *MnemonicAction // out

	_mnemonicAction = wrapMnemonicAction(coreglib.Take(unsafe.Pointer(_cret)))

	return _mnemonicAction
}

// NamedActionOverrider contains methods that are overridable.
type NamedActionOverrider interface {
}

// NamedAction: GtkShortcutAction that activates an action by name.
type NamedAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*NamedAction)(nil)
)

func initClassNamedAction(gclass unsafe.Pointer, goval any) {
}

func wrapNamedAction(obj *coreglib.Object) *NamedAction {
	return &NamedAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalNamedAction(p uintptr) (interface{}, error) {
	return wrapNamedAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNamedAction creates an action that when activated, activates the named
// action on the widget.
//
// It also passes the given arguments to it.
//
// See gtk.Widget.InsertActionGroup() for how to add actions to widgets.
//
// The function takes the following parameters:
//
//    - name: detailed name of the action.
//
// The function returns the following values:
//
//    - namedAction: new GtkShortcutAction.
//
func NewNamedAction(name string) *NamedAction {
	var _arg1 *C.char              // out
	var _cret *C.GtkShortcutAction // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_named_action_new(_arg1)
	runtime.KeepAlive(name)

	var _namedAction *NamedAction // out

	_namedAction = wrapNamedAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _namedAction
}

// ActionName returns the name of the action that will be activated.
//
// The function returns the following values:
//
//    - utf8: name of the action to activate.
//
func (self *NamedAction) ActionName() string {
	var _arg0 *C.GtkNamedAction // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkNamedAction)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_named_action_get_action_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NothingActionOverrider contains methods that are overridable.
type NothingActionOverrider interface {
}

// NothingAction: GtkShortcutAction that does nothing.
type NothingAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*NothingAction)(nil)
)

func initClassNothingAction(gclass unsafe.Pointer, goval any) {
}

func wrapNothingAction(obj *coreglib.Object) *NothingAction {
	return &NothingAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalNothingAction(p uintptr) (interface{}, error) {
	return wrapNothingAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NothingActionGet gets the nothing action.
//
// This is an action that does nothing and where activating it always fails.
//
// The function returns the following values:
//
//    - nothingAction: nothing action.
//
func NothingActionGet() *NothingAction {
	var _cret *C.GtkShortcutAction // in

	_cret = C.gtk_nothing_action_get()

	var _nothingAction *NothingAction // out

	_nothingAction = wrapNothingAction(coreglib.Take(unsafe.Pointer(_cret)))

	return _nothingAction
}

// ShortcutActionOverrider contains methods that are overridable.
type ShortcutActionOverrider interface {
}

// ShortcutAction: GtkShortcutAction encodes an action that can be triggered by
// a keyboard shortcut.
//
// GtkShortcutActions contain functions that allow easy presentation to end
// users as well as being printed for debugging.
//
// All GtkShortcutActions are immutable, you can only specify their properties
// during construction. If you want to change a action, you have to replace it
// with a new one. If you need to pass arguments to an action, these are
// specified by the higher-level GtkShortcut object.
//
// To activate a GtkShortcutAction manually, gtk.ShortcutAction.Activate() can
// be called.
//
// GTK provides various actions:
//
//    - gtk.MnemonicAction: a shortcut action that calls
//      gtk_widget_mnemonic_activate()
//    - gtk.CallbackAction: a shortcut action that invokes
//      a given callback
//    - gtk.SignalAction: a shortcut action that emits a
//      given signal
//    - gtk.ActivateAction: a shortcut action that calls
//      gtk_widget_activate()
//    - gtk.NamedAction: a shortcut action that calls
//      gtk_widget_activate_action()
//    - gtk.NothingAction: a shortcut action that does nothing.
type ShortcutAction struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ShortcutAction)(nil)
)

// ShortcutActioner describes types inherited from class ShortcutAction.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type ShortcutActioner interface {
	coreglib.Objector
	baseShortcutAction() *ShortcutAction
}

var _ ShortcutActioner = (*ShortcutAction)(nil)

func initClassShortcutAction(gclass unsafe.Pointer, goval any) {
}

func wrapShortcutAction(obj *coreglib.Object) *ShortcutAction {
	return &ShortcutAction{
		Object: obj,
	}
}

func marshalShortcutAction(p uintptr) (interface{}, error) {
	return wrapShortcutAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *ShortcutAction) baseShortcutAction() *ShortcutAction {
	return self
}

// BaseShortcutAction returns the underlying base object.
func BaseShortcutAction(obj ShortcutActioner) *ShortcutAction {
	return obj.baseShortcutAction()
}

// NewShortcutActionParseString tries to parse the given string into an action.
//
// On success, the parsed action is returned. When parsing failed, NULL is
// returned.
//
// The accepted strings are:
//
// - nothing, for GtkNothingAction
//
// - activate, for GtkActivateAction
//
// - mnemonic-activate, for GtkMnemonicAction
//
// - action(NAME), for a GtkNamedAction for the action named NAME
//
// - signal(NAME), for a GtkSignalAction for the signal NAME.
//
// The function takes the following parameters:
//
//    - str: string to parse.
//
// The function returns the following values:
//
//    - shortcutAction (optional): new GtkShortcutAction or NULL on error.
//
func NewShortcutActionParseString(str string) *ShortcutAction {
	var _arg1 *C.char              // out
	var _cret *C.GtkShortcutAction // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_shortcut_action_parse_string(_arg1)
	runtime.KeepAlive(str)

	var _shortcutAction *ShortcutAction // out

	if _cret != nil {
		_shortcutAction = wrapShortcutAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _shortcutAction
}

// Activate activates the action on the widget with the given args.
//
// Note that some actions ignore the passed in flags, widget or args.
//
// Activation of an action can fail for various reasons. If the action is not
// supported by the widget, if the args don't match the action or if the
// activation otherwise had no effect, FALSE will be returned.
//
// The function takes the following parameters:
//
//    - flags to activate with.
//    - widget: target of the activation.
//    - args (optional) arguments to pass.
//
// The function returns the following values:
//
//    - ok: TRUE if this action was activated successfully.
//
func (self *ShortcutAction) Activate(flags ShortcutActionFlags, widget Widgetter, args *glib.Variant) bool {
	var _arg0 *C.GtkShortcutAction     // out
	var _arg1 C.GtkShortcutActionFlags // out
	var _arg2 *C.GtkWidget             // out
	var _arg3 *C.GVariant              // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkShortcutAction)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkShortcutActionFlags(flags)
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	if args != nil {
		_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(args)))
	}

	_cret = C.gtk_shortcut_action_activate(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(args)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String prints the given action into a human-readable string.
//
// This is a small wrapper around gtk.ShortcutAction.Print() to help when
// debugging.
//
// The function returns the following values:
//
//    - utf8: new string.
//
func (self *ShortcutAction) String() string {
	var _arg0 *C.GtkShortcutAction // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkShortcutAction)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_shortcut_action_to_string(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SignalActionOverrider contains methods that are overridable.
type SignalActionOverrider interface {
}

// SignalAction: GtkShortcutAction that emits a signal.
//
// Signals that are used in this way are referred to as keybinding signals, and
// they are expected to be defined with the G_SIGNAL_ACTION flag.
type SignalAction struct {
	_ [0]func() // equal guard
	ShortcutAction
}

var (
	_ ShortcutActioner = (*SignalAction)(nil)
)

func initClassSignalAction(gclass unsafe.Pointer, goval any) {
}

func wrapSignalAction(obj *coreglib.Object) *SignalAction {
	return &SignalAction{
		ShortcutAction: ShortcutAction{
			Object: obj,
		},
	}
}

func marshalSignalAction(p uintptr) (interface{}, error) {
	return wrapSignalAction(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSignalAction creates an action that when activated, emits the given action
// signal on the provided widget.
//
// It will also unpack the args into arguments passed to the signal.
//
// The function takes the following parameters:
//
//    - signalName: name of the signal to emit.
//
// The function returns the following values:
//
//    - signalAction: new GtkShortcutAction.
//
func NewSignalAction(signalName string) *SignalAction {
	var _arg1 *C.char              // out
	var _cret *C.GtkShortcutAction // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(signalName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_signal_action_new(_arg1)
	runtime.KeepAlive(signalName)

	var _signalAction *SignalAction // out

	_signalAction = wrapSignalAction(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _signalAction
}

// SignalName returns the name of the signal that will be emitted.
//
// The function returns the following values:
//
//    - utf8: name of the signal to emit.
//
func (self *SignalAction) SignalName() string {
	var _arg0 *C.GtkSignalAction // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkSignalAction)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_signal_action_get_signal_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
