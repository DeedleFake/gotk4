// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_crossing_mode_get_type()), F: marshalCrossingMode},
		{T: externglib.Type(C.gdk_event_type_get_type()), F: marshalEventType},
		{T: externglib.Type(C.gdk_key_match_get_type()), F: marshalKeyMatch},
		{T: externglib.Type(C.gdk_notify_type_get_type()), F: marshalNotifyType},
		{T: externglib.Type(C.gdk_scroll_direction_get_type()), F: marshalScrollDirection},
		{T: externglib.Type(C.gdk_touchpad_gesture_phase_get_type()), F: marshalTouchpadGesturePhase},
		{T: externglib.Type(C.gdk_button_event_get_type()), F: marshalButtonEvent},
		{T: externglib.Type(C.gdk_crossing_event_get_type()), F: marshalCrossingEvent},
		{T: externglib.Type(C.gdk_dnd_event_get_type()), F: marshalDNDEvent},
		{T: externglib.Type(C.gdk_delete_event_get_type()), F: marshalDeleteEvent},
		{T: externglib.Type(C.gdk_event_get_type()), F: marshalEvent},
		{T: externglib.Type(C.gdk_focus_event_get_type()), F: marshalFocusEvent},
		{T: externglib.Type(C.gdk_grab_broken_event_get_type()), F: marshalGrabBrokenEvent},
		{T: externglib.Type(C.gdk_key_event_get_type()), F: marshalKeyEvent},
		{T: externglib.Type(C.gdk_motion_event_get_type()), F: marshalMotionEvent},
		{T: externglib.Type(C.gdk_pad_event_get_type()), F: marshalPadEvent},
		{T: externglib.Type(C.gdk_proximity_event_get_type()), F: marshalProximityEvent},
		{T: externglib.Type(C.gdk_scroll_event_get_type()), F: marshalScrollEvent},
		{T: externglib.Type(C.gdk_touch_event_get_type()), F: marshalTouchEvent},
		{T: externglib.Type(C.gdk_touchpad_event_get_type()), F: marshalTouchpadEvent},
		{T: externglib.Type(C.gdk_event_sequence_get_type()), F: marshalEventSequence},
	})
}

// CrossingMode specifies the crossing mode for enter and leave events.
type CrossingMode int

const (
	// normal: crossing because of pointer motion.
	CrossingModeNormal CrossingMode = 0
	// grab: crossing because a grab is activated.
	CrossingModeGrab CrossingMode = 1
	// ungrab: crossing because a grab is deactivated.
	CrossingModeUngrab CrossingMode = 2
	// GTKGrab: crossing because a GTK grab is activated.
	CrossingModeGTKGrab CrossingMode = 3
	// GTKUngrab: crossing because a GTK grab is deactivated.
	CrossingModeGTKUngrab CrossingMode = 4
	// StateChanged: crossing because a GTK widget changed state (e.g.
	// sensitivity).
	CrossingModeStateChanged CrossingMode = 5
	// TouchBegin: crossing because a touch sequence has begun, this event is
	// synthetic as the pointer might have not left the surface.
	CrossingModeTouchBegin CrossingMode = 6
	// TouchEnd: crossing because a touch sequence has ended, this event is
	// synthetic as the pointer might have not left the surface.
	CrossingModeTouchEnd CrossingMode = 7
	// DeviceSwitch: crossing because of a device switch (i.e. a mouse taking
	// control of the pointer after a touch device), this event is synthetic as
	// the pointer didn’t leave the surface.
	CrossingModeDeviceSwitch CrossingMode = 8
)

func marshalCrossingMode(p uintptr) (interface{}, error) {
	return CrossingMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EventType specifies the type of the event.
type EventType int

const (
	// delete: the window manager has requested that the toplevel surface be
	// hidden or destroyed, usually when the user clicks on a special icon in
	// the title bar.
	EventTypeDelete EventType = 0
	// MotionNotify: the pointer (usually a mouse) has moved.
	EventTypeMotionNotify EventType = 1
	// ButtonPress: a mouse button has been pressed.
	EventTypeButtonPress EventType = 2
	// ButtonRelease: a mouse button has been released.
	EventTypeButtonRelease EventType = 3
	// KeyPress: a key has been pressed.
	EventTypeKeyPress EventType = 4
	// KeyRelease: a key has been released.
	EventTypeKeyRelease EventType = 5
	// EnterNotify: the pointer has entered the surface.
	EventTypeEnterNotify EventType = 6
	// LeaveNotify: the pointer has left the surface.
	EventTypeLeaveNotify EventType = 7
	// FocusChange: the keyboard focus has entered or left the surface.
	EventTypeFocusChange EventType = 8
	// ProximityIn: an input device has moved into contact with a sensing
	// surface (e.g. a touchscreen or graphics tablet).
	EventTypeProximityIn EventType = 9
	// ProximityOut: an input device has moved out of contact with a sensing
	// surface.
	EventTypeProximityOut EventType = 10
	// DragEnter: the mouse has entered the surface while a drag is in progress.
	EventTypeDragEnter EventType = 11
	// DragLeave: the mouse has left the surface while a drag is in progress.
	EventTypeDragLeave EventType = 12
	// DragMotion: the mouse has moved in the surface while a drag is in
	// progress.
	EventTypeDragMotion EventType = 13
	// DropStart: a drop operation onto the surface has started.
	EventTypeDropStart EventType = 14
	// scroll: the scroll wheel was turned
	EventTypeScroll EventType = 15
	// GrabBroken: a pointer or keyboard grab was broken.
	EventTypeGrabBroken EventType = 16
	// TouchBegin: a new touch event sequence has just started.
	EventTypeTouchBegin EventType = 17
	// TouchUpdate: a touch event sequence has been updated.
	EventTypeTouchUpdate EventType = 18
	// TouchEnd: a touch event sequence has finished.
	EventTypeTouchEnd EventType = 19
	// TouchCancel: a touch event sequence has been canceled.
	EventTypeTouchCancel EventType = 20
	// TouchpadSwipe: a touchpad swipe gesture event, the current state is
	// determined by its phase field.
	EventTypeTouchpadSwipe EventType = 21
	// TouchpadPinch: a touchpad pinch gesture event, the current state is
	// determined by its phase field.
	EventTypeTouchpadPinch EventType = 22
	// PadButtonPress: a tablet pad button press event.
	EventTypePadButtonPress EventType = 23
	// PadButtonRelease: a tablet pad button release event.
	EventTypePadButtonRelease EventType = 24
	// PadRing: a tablet pad axis event from a "ring".
	EventTypePadRing EventType = 25
	// PadStrip: a tablet pad axis event from a "strip".
	EventTypePadStrip EventType = 26
	// PadGroupMode: a tablet pad group mode change.
	EventTypePadGroupMode EventType = 27
	// EventLast marks the end of the GdkEventType enumeration.
	EventTypeEventLast EventType = 28
)

func marshalEventType(p uintptr) (interface{}, error) {
	return EventType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// KeyMatch describes how well an event matches a given keyval and modifiers.
//
// `GdkKeyMatch` values are returned by [method@Gdk.KeyEvent.matches].
type KeyMatch int

const (
	// none: the key event does not match
	KeyMatchNone KeyMatch = 0
	// partial: the key event matches if keyboard state (specifically, the
	// currently active group) is ignored
	KeyMatchPartial KeyMatch = 1
	// exact: the key event matches
	KeyMatchExact KeyMatch = 2
)

func marshalKeyMatch(p uintptr) (interface{}, error) {
	return KeyMatch(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// NotifyType specifies the kind of crossing for enter and leave events.
//
// See the X11 protocol specification of LeaveNotify for full details of
// crossing event generation.
type NotifyType int

const (
	// ancestor: the surface is entered from an ancestor or left towards an
	// ancestor.
	NotifyTypeAncestor NotifyType = 0
	// virtual: the pointer moves between an ancestor and an inferior of the
	// surface.
	NotifyTypeVirtual NotifyType = 1
	// inferior: the surface is entered from an inferior or left towards an
	// inferior.
	NotifyTypeInferior NotifyType = 2
	// nonlinear: the surface is entered from or left towards a surface which is
	// neither an ancestor nor an inferior.
	NotifyTypeNonlinear NotifyType = 3
	// NonlinearVirtual: the pointer moves between two surfaces which are not
	// ancestors of each other and the surface is part of the ancestor chain
	// between one of these surfaces and their least common ancestor.
	NotifyTypeNonlinearVirtual NotifyType = 4
	// unknown: an unknown type of enter/leave event occurred.
	NotifyTypeUnknown NotifyType = 5
)

func marshalNotifyType(p uintptr) (interface{}, error) {
	return NotifyType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScrollDirection specifies the direction for scroll events.
type ScrollDirection int

const (
	// up: the surface is scrolled up.
	ScrollDirectionUp ScrollDirection = 0
	// down: the surface is scrolled down.
	ScrollDirectionDown ScrollDirection = 1
	// left: the surface is scrolled to the left.
	ScrollDirectionLeft ScrollDirection = 2
	// right: the surface is scrolled to the right.
	ScrollDirectionRight ScrollDirection = 3
	// smooth: the scrolling is determined by the delta values in scroll events.
	// See gdk_scroll_event_get_deltas()
	ScrollDirectionSmooth ScrollDirection = 4
)

func marshalScrollDirection(p uintptr) (interface{}, error) {
	return ScrollDirection(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TouchpadGesturePhase specifies the current state of a touchpad gesture.
//
// All gestures are guaranteed to begin with an event with phase
// GDK_TOUCHPAD_GESTURE_PHASE_BEGIN, followed by 0 or several events with phase
// GDK_TOUCHPAD_GESTURE_PHASE_UPDATE.
//
// A finished gesture may have 2 possible outcomes, an event with phase
// GDK_TOUCHPAD_GESTURE_PHASE_END will be emitted when the gesture is considered
// successful, this should be used as the hint to perform any permanent changes.
//
// Cancelled gestures may be so for a variety of reasons, due to hardware or the
// compositor, or due to the gesture recognition layers hinting the gesture did
// not finish resolutely (eg. a 3rd finger being added during a pinch gesture).
// In these cases, the last event will report the phase
// GDK_TOUCHPAD_GESTURE_PHASE_CANCEL, this should be used as a hint to undo any
// visible/permanent changes that were done throughout the progress of the
// gesture.
type TouchpadGesturePhase int

const (
	// begin: the gesture has begun.
	TouchpadGesturePhaseBegin TouchpadGesturePhase = 0
	// update: the gesture has been updated.
	TouchpadGesturePhaseUpdate TouchpadGesturePhase = 1
	// end: the gesture was finished, changes should be permanently applied.
	TouchpadGesturePhaseEnd TouchpadGesturePhase = 2
	// cancel: the gesture was cancelled, all changes should be undone.
	TouchpadGesturePhaseCancel TouchpadGesturePhase = 3
)

func marshalTouchpadGesturePhase(p uintptr) (interface{}, error) {
	return TouchpadGesturePhase(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EventsGetAngle returns the relative angle from @event1 to @event2.
//
// The relative angle is the angle between the X axis and the line through both
// events' positions. The rotation direction for positive angles is from the
// positive X axis towards the positive Y axis.
//
// This assumes that both events have X/Y information. If not, this function
// returns false.
func EventsGetAngle(event1 Event, event2 Event) (float64, bool) {
	var _arg1 *C.GdkEvent // out
	var _arg2 *C.GdkEvent // out
	var _arg3 C.double    // in
	var _cret C.gboolean  // in

	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event1.Native()))
	_arg2 = (*C.GdkEvent)(unsafe.Pointer(event2.Native()))

	_cret = C.gdk_events_get_angle(_arg1, _arg2, &_arg3)

	var _angle float64 // out
	var _ok bool       // out

	_angle = float64(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _angle, _ok
}

// EventsGetCenter returns the point halfway between the events' positions.
//
// This assumes that both events have X/Y information. If not, this function
// returns false.
func EventsGetCenter(event1 Event, event2 Event) (x float64, y float64, ok bool) {
	var _arg1 *C.GdkEvent // out
	var _arg2 *C.GdkEvent // out
	var _arg3 C.double    // in
	var _arg4 C.double    // in
	var _cret C.gboolean  // in

	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event1.Native()))
	_arg2 = (*C.GdkEvent)(unsafe.Pointer(event2.Native()))

	_cret = C.gdk_events_get_center(_arg1, _arg2, &_arg3, &_arg4)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg3)
	_y = float64(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// EventsGetDistance returns the distance between the event locations.
//
// This assumes that both events have X/Y information. If not, this function
// returns false.
func EventsGetDistance(event1 Event, event2 Event) (float64, bool) {
	var _arg1 *C.GdkEvent // out
	var _arg2 *C.GdkEvent // out
	var _arg3 C.double    // in
	var _cret C.gboolean  // in

	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event1.Native()))
	_arg2 = (*C.GdkEvent)(unsafe.Pointer(event2.Native()))

	_cret = C.gdk_events_get_distance(_arg1, _arg2, &_arg3)

	var _distance float64 // out
	var _ok bool          // out

	_distance = float64(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _distance, _ok
}

// ButtonEvent: an event related to a button on a pointer device.
type ButtonEvent interface {
	Event

	Button() uint
}

// buttonEvent implements the ButtonEvent class.
type buttonEvent struct {
	Event
}

// WrapButtonEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapButtonEvent(obj *externglib.Object) ButtonEvent {
	return buttonEvent{
		Event: WrapEvent(obj),
	}
}

func marshalButtonEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapButtonEvent(obj), nil
}

func (e buttonEvent) Button() uint {
	var _arg0 *C.GdkEvent // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_button_event_get_button(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CrossingEvent: an event caused by a pointing device moving between surfaces.
type CrossingEvent interface {
	Event

	Detail() NotifyType

	Focus() bool

	Mode() CrossingMode
}

// crossingEvent implements the CrossingEvent class.
type crossingEvent struct {
	Event
}

// WrapCrossingEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapCrossingEvent(obj *externglib.Object) CrossingEvent {
	return crossingEvent{
		Event: WrapEvent(obj),
	}
}

func marshalCrossingEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCrossingEvent(obj), nil
}

func (e crossingEvent) Detail() NotifyType {
	var _arg0 *C.GdkEvent     // out
	var _cret C.GdkNotifyType // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_crossing_event_get_detail(_arg0)

	var _notifyType NotifyType // out

	_notifyType = NotifyType(_cret)

	return _notifyType
}

func (e crossingEvent) Focus() bool {
	var _arg0 *C.GdkEvent // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_crossing_event_get_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e crossingEvent) Mode() CrossingMode {
	var _arg0 *C.GdkEvent       // out
	var _cret C.GdkCrossingMode // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_crossing_event_get_mode(_arg0)

	var _crossingMode CrossingMode // out

	_crossingMode = CrossingMode(_cret)

	return _crossingMode
}

// DNDEvent: an event related to drag and drop operations.
type DNDEvent interface {
	Event

	Drop() Drop
}

// dndEvent implements the DNDEvent class.
type dndEvent struct {
	Event
}

// WrapDNDEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapDNDEvent(obj *externglib.Object) DNDEvent {
	return dndEvent{
		Event: WrapEvent(obj),
	}
}

func marshalDNDEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDNDEvent(obj), nil
}

func (e dndEvent) Drop() Drop {
	var _arg0 *C.GdkEvent // out
	var _cret *C.GdkDrop  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_dnd_event_get_drop(_arg0)

	var _drop Drop // out

	_drop = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Drop)

	return _drop
}

// DeleteEvent: an event related to closing a top-level surface.
type DeleteEvent interface {
	Event
}

// deleteEvent implements the DeleteEvent class.
type deleteEvent struct {
	Event
}

// WrapDeleteEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapDeleteEvent(obj *externglib.Object) DeleteEvent {
	return deleteEvent{
		Event: WrapEvent(obj),
	}
}

func marshalDeleteEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDeleteEvent(obj), nil
}

// Event `GdkEvent`s are immutable data structures, created by GDK to represent
// windowing system events.
//
// In GTK applications the events are handled automatically by toplevel widgets
// and passed on to the event controllers of appropriate widgets, so using
// `GdkEvent` and its related API is rarely needed.
type Event interface {
	gextras.Objector

	Axes() ([]float64, bool)

	Axis(axisUse AxisUse) (float64, bool)

	Device() Device

	DeviceTool() DeviceTool

	Display() Display

	EventSequence() *EventSequence

	EventType() EventType

	ModifierState() ModifierType

	PointerEmulated() bool

	Position() (x float64, y float64, ok bool)

	Seat() Seat

	Surface() Surface

	Time() uint32

	RefEvent() Event

	TriggersContextMenuEvent() bool

	UnrefEvent()
}

// event implements the Event class.
type event struct {
	gextras.Objector
}

// WrapEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapEvent(obj *externglib.Object) Event {
	return event{
		Objector: obj,
	}
}

func marshalEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEvent(obj), nil
}

func (e event) Axes() ([]float64, bool) {
	var _arg0 *C.GdkEvent // out
	var _arg1 *C.double
	var _arg2 C.guint    // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_axes(_arg0, &_arg1, &_arg2)

	var _axes []float64
	var _ok bool // out

	_axes = make([]float64, _arg2)
	copy(_axes, unsafe.Slice((*float64)(unsafe.Pointer(_arg1)), _arg2))
	if _cret != 0 {
		_ok = true
	}

	return _axes, _ok
}

func (e event) Axis(axisUse AxisUse) (float64, bool) {
	var _arg0 *C.GdkEvent  // out
	var _arg1 C.GdkAxisUse // out
	var _arg2 C.double     // in
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))
	_arg1 = C.GdkAxisUse(axisUse)

	_cret = C.gdk_event_get_axis(_arg0, _arg1, &_arg2)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

func (e event) Device() Device {
	var _arg0 *C.GdkEvent  // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_device(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}

func (e event) DeviceTool() DeviceTool {
	var _arg0 *C.GdkEvent      // out
	var _cret *C.GdkDeviceTool // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_device_tool(_arg0)

	var _deviceTool DeviceTool // out

	_deviceTool = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(DeviceTool)

	return _deviceTool
}

func (e event) Display() Display {
	var _arg0 *C.GdkEvent   // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (e event) EventSequence() *EventSequence {
	var _arg0 *C.GdkEvent         // out
	var _cret *C.GdkEventSequence // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_event_sequence(_arg0)

	var _eventSequence *EventSequence // out

	_eventSequence = (*EventSequence)(unsafe.Pointer(_cret))

	return _eventSequence
}

func (e event) EventType() EventType {
	var _arg0 *C.GdkEvent    // out
	var _cret C.GdkEventType // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_event_type(_arg0)

	var _eventType EventType // out

	_eventType = EventType(_cret)

	return _eventType
}

func (e event) ModifierState() ModifierType {
	var _arg0 *C.GdkEvent       // out
	var _cret C.GdkModifierType // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_modifier_state(_arg0)

	var _modifierType ModifierType // out

	_modifierType = ModifierType(_cret)

	return _modifierType
}

func (e event) PointerEmulated() bool {
	var _arg0 *C.GdkEvent // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_pointer_emulated(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e event) Position() (x float64, y float64, ok bool) {
	var _arg0 *C.GdkEvent // out
	var _arg1 C.double    // in
	var _arg2 C.double    // in
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_position(_arg0, &_arg1, &_arg2)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg1)
	_y = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

func (e event) Seat() Seat {
	var _arg0 *C.GdkEvent // out
	var _cret *C.GdkSeat  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_seat(_arg0)

	var _seat Seat // out

	_seat = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Seat)

	return _seat
}

func (e event) Surface() Surface {
	var _arg0 *C.GdkEvent   // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Surface)

	return _surface
}

func (e event) Time() uint32 {
	var _arg0 *C.GdkEvent // out
	var _cret C.guint32   // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_get_time(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (e event) RefEvent() Event {
	var _arg0 *C.GdkEvent // out
	var _cret *C.GdkEvent // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_ref(_arg0)

	var _ret Event // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Event)

	return _ret
}

func (e event) TriggersContextMenuEvent() bool {
	var _arg0 *C.GdkEvent // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_event_triggers_context_menu(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e event) UnrefEvent() {
	var _arg0 *C.GdkEvent // out

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	C.gdk_event_unref(_arg0)
}

// FocusEvent: an event related to a keyboard focus change.
type FocusEvent interface {
	Event

	In() bool
}

// focusEvent implements the FocusEvent class.
type focusEvent struct {
	Event
}

// WrapFocusEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapFocusEvent(obj *externglib.Object) FocusEvent {
	return focusEvent{
		Event: WrapEvent(obj),
	}
}

func marshalFocusEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFocusEvent(obj), nil
}

func (e focusEvent) In() bool {
	var _arg0 *C.GdkEvent // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_focus_event_get_in(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GrabBrokenEvent: an event related to a broken windowing system grab.
type GrabBrokenEvent interface {
	Event

	GrabSurface() Surface

	Implicit() bool
}

// grabBrokenEvent implements the GrabBrokenEvent class.
type grabBrokenEvent struct {
	Event
}

// WrapGrabBrokenEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapGrabBrokenEvent(obj *externglib.Object) GrabBrokenEvent {
	return grabBrokenEvent{
		Event: WrapEvent(obj),
	}
}

func marshalGrabBrokenEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGrabBrokenEvent(obj), nil
}

func (e grabBrokenEvent) GrabSurface() Surface {
	var _arg0 *C.GdkEvent   // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_grab_broken_event_get_grab_surface(_arg0)

	var _surface Surface // out

	_surface = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Surface)

	return _surface
}

func (e grabBrokenEvent) Implicit() bool {
	var _arg0 *C.GdkEvent // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_grab_broken_event_get_implicit(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyEvent: an event related to a key-based device.
type KeyEvent interface {
	Event

	ConsumedModifiers() ModifierType

	Keycode() uint

	Keyval() uint

	Layout() uint

	Level() uint

	Match() (uint, ModifierType, bool)

	IsModifierKeyEvent() bool

	MatchesKeyEvent(keyval uint, modifiers ModifierType) KeyMatch
}

// keyEvent implements the KeyEvent class.
type keyEvent struct {
	Event
}

// WrapKeyEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapKeyEvent(obj *externglib.Object) KeyEvent {
	return keyEvent{
		Event: WrapEvent(obj),
	}
}

func marshalKeyEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapKeyEvent(obj), nil
}

func (e keyEvent) ConsumedModifiers() ModifierType {
	var _arg0 *C.GdkEvent       // out
	var _cret C.GdkModifierType // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_key_event_get_consumed_modifiers(_arg0)

	var _modifierType ModifierType // out

	_modifierType = ModifierType(_cret)

	return _modifierType
}

func (e keyEvent) Keycode() uint {
	var _arg0 *C.GdkEvent // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_key_event_get_keycode(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (e keyEvent) Keyval() uint {
	var _arg0 *C.GdkEvent // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_key_event_get_keyval(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (e keyEvent) Layout() uint {
	var _arg0 *C.GdkEvent // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_key_event_get_layout(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (e keyEvent) Level() uint {
	var _arg0 *C.GdkEvent // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_key_event_get_level(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (e keyEvent) Match() (uint, ModifierType, bool) {
	var _arg0 *C.GdkEvent       // out
	var _arg1 C.guint           // in
	var _arg2 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_key_event_get_match(_arg0, &_arg1, &_arg2)

	var _keyval uint            // out
	var _modifiers ModifierType // out
	var _ok bool                // out

	_keyval = uint(_arg1)
	_modifiers = ModifierType(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _modifiers, _ok
}

func (e keyEvent) IsModifierKeyEvent() bool {
	var _arg0 *C.GdkEvent // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_key_event_is_modifier(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e keyEvent) MatchesKeyEvent(keyval uint, modifiers ModifierType) KeyMatch {
	var _arg0 *C.GdkEvent       // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _cret C.GdkKeyMatch     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))
	_arg1 = C.guint(keyval)
	_arg2 = C.GdkModifierType(modifiers)

	_cret = C.gdk_key_event_matches(_arg0, _arg1, _arg2)

	var _keyMatch KeyMatch // out

	_keyMatch = KeyMatch(_cret)

	return _keyMatch
}

// MotionEvent: an event related to a pointer or touch device motion.
type MotionEvent interface {
	Event
}

// motionEvent implements the MotionEvent class.
type motionEvent struct {
	Event
}

// WrapMotionEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapMotionEvent(obj *externglib.Object) MotionEvent {
	return motionEvent{
		Event: WrapEvent(obj),
	}
}

func marshalMotionEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMotionEvent(obj), nil
}

// PadEvent: an event related to a pad-based device.
type PadEvent interface {
	Event

	AxisValue() (uint, float64)

	Button() uint

	GroupMode() (group uint, mode uint)
}

// padEvent implements the PadEvent class.
type padEvent struct {
	Event
}

// WrapPadEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapPadEvent(obj *externglib.Object) PadEvent {
	return padEvent{
		Event: WrapEvent(obj),
	}
}

func marshalPadEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPadEvent(obj), nil
}

func (e padEvent) AxisValue() (uint, float64) {
	var _arg0 *C.GdkEvent // out
	var _arg1 C.guint     // in
	var _arg2 C.double    // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	C.gdk_pad_event_get_axis_value(_arg0, &_arg1, &_arg2)

	var _index uint    // out
	var _value float64 // out

	_index = uint(_arg1)
	_value = float64(_arg2)

	return _index, _value
}

func (e padEvent) Button() uint {
	var _arg0 *C.GdkEvent // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_pad_event_get_button(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (e padEvent) GroupMode() (group uint, mode uint) {
	var _arg0 *C.GdkEvent // out
	var _arg1 C.guint     // in
	var _arg2 C.guint     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	C.gdk_pad_event_get_group_mode(_arg0, &_arg1, &_arg2)

	var _group uint // out
	var _mode uint  // out

	_group = uint(_arg1)
	_mode = uint(_arg2)

	return _group, _mode
}

// ProximityEvent: an event related to the proximity of a tool to a device.
type ProximityEvent interface {
	Event
}

// proximityEvent implements the ProximityEvent class.
type proximityEvent struct {
	Event
}

// WrapProximityEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapProximityEvent(obj *externglib.Object) ProximityEvent {
	return proximityEvent{
		Event: WrapEvent(obj),
	}
}

func marshalProximityEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProximityEvent(obj), nil
}

// ScrollEvent: an event related to a scrolling motion.
type ScrollEvent interface {
	Event

	Deltas() (deltaX float64, deltaY float64)

	Direction() ScrollDirection

	IsStopScrollEvent() bool
}

// scrollEvent implements the ScrollEvent class.
type scrollEvent struct {
	Event
}

// WrapScrollEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapScrollEvent(obj *externglib.Object) ScrollEvent {
	return scrollEvent{
		Event: WrapEvent(obj),
	}
}

func marshalScrollEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrollEvent(obj), nil
}

func (e scrollEvent) Deltas() (deltaX float64, deltaY float64) {
	var _arg0 *C.GdkEvent // out
	var _arg1 C.double    // in
	var _arg2 C.double    // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	C.gdk_scroll_event_get_deltas(_arg0, &_arg1, &_arg2)

	var _deltaX float64 // out
	var _deltaY float64 // out

	_deltaX = float64(_arg1)
	_deltaY = float64(_arg2)

	return _deltaX, _deltaY
}

func (e scrollEvent) Direction() ScrollDirection {
	var _arg0 *C.GdkEvent          // out
	var _cret C.GdkScrollDirection // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_scroll_event_get_direction(_arg0)

	var _scrollDirection ScrollDirection // out

	_scrollDirection = ScrollDirection(_cret)

	return _scrollDirection
}

func (e scrollEvent) IsStopScrollEvent() bool {
	var _arg0 *C.GdkEvent // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_scroll_event_is_stop(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TouchEvent: an event related to a touch-based device.
type TouchEvent interface {
	Event

	EmulatingPointer() bool
}

// touchEvent implements the TouchEvent class.
type touchEvent struct {
	Event
}

// WrapTouchEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapTouchEvent(obj *externglib.Object) TouchEvent {
	return touchEvent{
		Event: WrapEvent(obj),
	}
}

func marshalTouchEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTouchEvent(obj), nil
}

func (e touchEvent) EmulatingPointer() bool {
	var _arg0 *C.GdkEvent // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_touch_event_get_emulating_pointer(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TouchpadEvent: an event related to a gesture on a touchpad device.
//
// Unlike touchscreens, where the windowing system sends basic sequences of
// begin, update, end events, and leaves gesture recognition to the clients,
// touchpad gestures are typically processed by the system, resulting in these
// events.
type TouchpadEvent interface {
	Event

	Deltas() (dx float64, dy float64)

	GesturePhase() TouchpadGesturePhase

	NFingers() uint

	PinchAngleDelta() float64

	PinchScale() float64
}

// touchpadEvent implements the TouchpadEvent class.
type touchpadEvent struct {
	Event
}

// WrapTouchpadEvent wraps a GObject to the right type. It is
// primarily used internally.
func WrapTouchpadEvent(obj *externglib.Object) TouchpadEvent {
	return touchpadEvent{
		Event: WrapEvent(obj),
	}
}

func marshalTouchpadEvent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTouchpadEvent(obj), nil
}

func (e touchpadEvent) Deltas() (dx float64, dy float64) {
	var _arg0 *C.GdkEvent // out
	var _arg1 C.double    // in
	var _arg2 C.double    // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	C.gdk_touchpad_event_get_deltas(_arg0, &_arg1, &_arg2)

	var _dx float64 // out
	var _dy float64 // out

	_dx = float64(_arg1)
	_dy = float64(_arg2)

	return _dx, _dy
}

func (e touchpadEvent) GesturePhase() TouchpadGesturePhase {
	var _arg0 *C.GdkEvent               // out
	var _cret C.GdkTouchpadGesturePhase // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_touchpad_event_get_gesture_phase(_arg0)

	var _touchpadGesturePhase TouchpadGesturePhase // out

	_touchpadGesturePhase = TouchpadGesturePhase(_cret)

	return _touchpadGesturePhase
}

func (e touchpadEvent) NFingers() uint {
	var _arg0 *C.GdkEvent // out
	var _cret C.guint     // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_touchpad_event_get_n_fingers(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (e touchpadEvent) PinchAngleDelta() float64 {
	var _arg0 *C.GdkEvent // out
	var _cret C.double    // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_touchpad_event_get_pinch_angle_delta(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (e touchpadEvent) PinchScale() float64 {
	var _arg0 *C.GdkEvent // out
	var _cret C.double    // in

	_arg0 = (*C.GdkEvent)(unsafe.Pointer(e.Native()))

	_cret = C.gdk_touchpad_event_get_pinch_scale(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// EventSequence: `GdkEventSequence` is an opaque type representing a sequence
// of related touch events.
type EventSequence C.GdkEventSequence

// WrapEventSequence wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapEventSequence(ptr unsafe.Pointer) *EventSequence {
	return (*EventSequence)(ptr)
}

func marshalEventSequence(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*EventSequence)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *EventSequence) Native() unsafe.Pointer {
	return unsafe.Pointer(e)
}
