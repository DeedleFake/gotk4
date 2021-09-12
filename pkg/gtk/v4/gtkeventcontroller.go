// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_get_type()), F: marshalEventControllerer},
	})
}

// EventController: GtkEventController is the base class for event controllers.
//
// These are ancillary objects associated to widgets, which react to GdkEvents,
// and possibly trigger actions as a consequence.
//
// Event controllers are added to a widget with gtk.Widget.AddController(). It
// is rarely necessary to explicitly remove a controller with
// gtk.Widget.RemoveController().
//
// See the chapter of input handling (input-handling.html) for an overview of
// the basic concepts, such as the capture and bubble phases of even
// propagation.
type EventController struct {
	*externglib.Object
}

// EventControllerer describes EventController's abstract methods.
type EventControllerer interface {
	externglib.Objector

	// CurrentEvent returns the event that is currently being handled by the
	// controller, and NULL at other times.
	CurrentEvent() gdk.Eventer
	// CurrentEventDevice returns the device of the event that is currently
	// being handled by the controller, and NULL otherwise.
	CurrentEventDevice() gdk.Devicer
	// CurrentEventState returns the modifier state of the event that is
	// currently being handled by the controller, and 0 otherwise.
	CurrentEventState() gdk.ModifierType
	// CurrentEventTime returns the timestamp of the event that is currently
	// being handled by the controller, and 0 otherwise.
	CurrentEventTime() uint32
	// Name gets the name of controller.
	Name() string
	// PropagationLimit gets the propagation limit of the event controller.
	PropagationLimit() PropagationLimit
	// PropagationPhase gets the propagation phase at which controller handles
	// events.
	PropagationPhase() PropagationPhase
	// Widget returns the Widget this controller relates to.
	Widget() Widgetter
	// Reset resets the controller to a clean state.
	Reset()
	// SetName sets a name on the controller that can be used for debugging.
	SetName(name string)
	// SetPropagationLimit sets the event propagation limit on the event
	// controller.
	SetPropagationLimit(limit PropagationLimit)
	// SetPropagationPhase sets the propagation phase at which a controller
	// handles events.
	SetPropagationPhase(phase PropagationPhase)
}

var _ EventControllerer = (*EventController)(nil)

func wrapEventController(obj *externglib.Object) *EventController {
	return &EventController{
		Object: obj,
	}
}

func marshalEventControllerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventController(obj), nil
}

// CurrentEvent returns the event that is currently being handled by the
// controller, and NULL at other times.
func (controller *EventController) CurrentEvent() gdk.Eventer {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GdkEvent           // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_current_event(_arg0)
	runtime.KeepAlive(controller)

	var _event gdk.Eventer // out

	if _cret != nil {
		_event = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gdk.Eventer)
	}

	return _event
}

// CurrentEventDevice returns the device of the event that is currently being
// handled by the controller, and NULL otherwise.
func (controller *EventController) CurrentEventDevice() gdk.Devicer {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GdkDevice          // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_current_event_device(_arg0)
	runtime.KeepAlive(controller)

	var _device gdk.Devicer // out

	if _cret != nil {
		_device = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gdk.Devicer)
	}

	return _device
}

// CurrentEventState returns the modifier state of the event that is currently
// being handled by the controller, and 0 otherwise.
func (controller *EventController) CurrentEventState() gdk.ModifierType {
	var _arg0 *C.GtkEventController // out
	var _cret C.GdkModifierType     // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_current_event_state(_arg0)
	runtime.KeepAlive(controller)

	var _modifierType gdk.ModifierType // out

	_modifierType = gdk.ModifierType(_cret)

	return _modifierType
}

// CurrentEventTime returns the timestamp of the event that is currently being
// handled by the controller, and 0 otherwise.
func (controller *EventController) CurrentEventTime() uint32 {
	var _arg0 *C.GtkEventController // out
	var _cret C.guint32             // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_current_event_time(_arg0)
	runtime.KeepAlive(controller)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Name gets the name of controller.
func (controller *EventController) Name() string {
	var _arg0 *C.GtkEventController // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_name(_arg0)
	runtime.KeepAlive(controller)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PropagationLimit gets the propagation limit of the event controller.
func (controller *EventController) PropagationLimit() PropagationLimit {
	var _arg0 *C.GtkEventController // out
	var _cret C.GtkPropagationLimit // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_propagation_limit(_arg0)
	runtime.KeepAlive(controller)

	var _propagationLimit PropagationLimit // out

	_propagationLimit = PropagationLimit(_cret)

	return _propagationLimit
}

// PropagationPhase gets the propagation phase at which controller handles
// events.
func (controller *EventController) PropagationPhase() PropagationPhase {
	var _arg0 *C.GtkEventController // out
	var _cret C.GtkPropagationPhase // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_propagation_phase(_arg0)
	runtime.KeepAlive(controller)

	var _propagationPhase PropagationPhase // out

	_propagationPhase = PropagationPhase(_cret)

	return _propagationPhase
}

// Widget returns the Widget this controller relates to.
func (controller *EventController) Widget() Widgetter {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_widget(_arg0)
	runtime.KeepAlive(controller)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// Reset resets the controller to a clean state.
func (controller *EventController) Reset() {
	var _arg0 *C.GtkEventController // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	C.gtk_event_controller_reset(_arg0)
	runtime.KeepAlive(controller)
}

// SetName sets a name on the controller that can be used for debugging.
func (controller *EventController) SetName(name string) {
	var _arg0 *C.GtkEventController // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_event_controller_set_name(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(name)
}

// SetPropagationLimit sets the event propagation limit on the event controller.
//
// If the limit is set to GTK_LIMIT_SAME_NATIVE, the controller won't handle
// events that are targeted at widgets on a different surface, such as popovers.
func (controller *EventController) SetPropagationLimit(limit PropagationLimit) {
	var _arg0 *C.GtkEventController // out
	var _arg1 C.GtkPropagationLimit // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))
	_arg1 = C.GtkPropagationLimit(limit)

	C.gtk_event_controller_set_propagation_limit(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(limit)
}

// SetPropagationPhase sets the propagation phase at which a controller handles
// events.
//
// If phase is GTK_PHASE_NONE, no automatic event handling will be performed,
// but other additional gesture maintenance will.
func (controller *EventController) SetPropagationPhase(phase PropagationPhase) {
	var _arg0 *C.GtkEventController // out
	var _arg1 C.GtkPropagationPhase // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))
	_arg1 = C.GtkPropagationPhase(phase)

	C.gtk_event_controller_set_propagation_phase(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(phase)
}
