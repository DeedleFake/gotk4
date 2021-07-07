// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_single_get_type()), F: marshalGestureSingle},
	})
}

// GestureSingle is a subclass of Gesture, optimized (although not restricted)
// for dealing with mouse and single-touch gestures. Under interaction, these
// gestures stick to the first interacting sequence, which is accessible through
// gtk_gesture_single_get_current_sequence() while the gesture is being
// interacted with.
//
// By default gestures react to both GDK_BUTTON_PRIMARY and touch events,
// gtk_gesture_single_set_touch_only() can be used to change the touch behavior.
// Callers may also specify a different mouse button number to interact with
// through gtk_gesture_single_set_button(), or react to any mouse button by
// setting 0. While the gesture is active, the button being currently pressed
// can be known through gtk_gesture_single_get_current_button().
type GestureSingle interface {
	Gesture

	// AsGesture casts the class to the Gesture interface.
	AsGesture() Gesture

	// GetBoundingBox: if there are touch sequences being currently handled by
	// @gesture, this function returns true and fills in @rect with the bounding
	// box containing all active touches. Otherwise, false will be returned.
	//
	// Note: This function will yield unexpected results on touchpad gestures.
	// Since there is no correlation between physical and pixel distances, these
	// will look as if constrained in an infinitely small area, @rect width and
	// height will thus be 0 regardless of the number of touchpoints.
	//
	// This method is inherited from Gesture
	GetBoundingBox() (gdk.Rectangle, bool)
	// GetBoundingBoxCenter: if there are touch sequences being currently
	// handled by @gesture, this function returns true and fills in @x and @y
	// with the center of the bounding box containing all active touches.
	// Otherwise, false will be returned.
	//
	// This method is inherited from Gesture
	GetBoundingBoxCenter() (x float64, y float64, ok bool)
	// GetDevice returns the master Device that is currently operating on
	// @gesture, or nil if the gesture is not being interacted.
	//
	// This method is inherited from Gesture
	GetDevice() gdk.Device
	// GetLastUpdatedSequence returns the EventSequence that was last updated on
	// @gesture.
	//
	// This method is inherited from Gesture
	GetLastUpdatedSequence() *gdk.EventSequence
	// GetPoint: if @sequence is currently being interpreted by @gesture, this
	// function returns true and fills in @x and @y with the last coordinates
	// stored for that event sequence. The coordinates are always relative to
	// the widget allocation.
	//
	// This method is inherited from Gesture
	GetPoint(sequence *gdk.EventSequence) (x float64, y float64, ok bool)
	// GetSequenceState returns the @sequence state, as seen by @gesture.
	//
	// This method is inherited from Gesture
	GetSequenceState(sequence *gdk.EventSequence) EventSequenceState
	// GetWindow returns the user-defined window that receives the events
	// handled by @gesture. See gtk_gesture_set_window() for more information.
	//
	// This method is inherited from Gesture
	GetWindow() gdk.Window
	// Group adds @gesture to the same group than @group_gesture. Gestures are
	// by default isolated in their own groups.
	//
	// When gestures are grouped, the state of EventSequences is kept in sync
	// for all of those, so calling gtk_gesture_set_sequence_state(), on one
	// will transfer the same value to the others.
	//
	// Groups also perform an "implicit grabbing" of sequences, if a
	// EventSequence state is set to K_EVENT_SEQUENCE_CLAIMED on one group,
	// every other gesture group attached to the same Widget will switch the
	// state for that sequence to K_EVENT_SEQUENCE_DENIED.
	//
	// This method is inherited from Gesture
	Group(gesture Gesture)
	// HandlesSequence returns true if @gesture is currently handling events
	// corresponding to @sequence.
	//
	// This method is inherited from Gesture
	HandlesSequence(sequence *gdk.EventSequence) bool
	// IsActive returns true if the gesture is currently active. A gesture is
	// active meanwhile there are touch sequences interacting with it.
	//
	// This method is inherited from Gesture
	IsActive() bool
	// IsGroupedWith returns true if both gestures pertain to the same group.
	//
	// This method is inherited from Gesture
	IsGroupedWith(other Gesture) bool
	// IsRecognized returns true if the gesture is currently recognized. A
	// gesture is recognized if there are as many interacting touch sequences as
	// required by @gesture, and Gesture::check returned true for the sequences
	// being currently interpreted.
	//
	// This method is inherited from Gesture
	IsRecognized() bool
	// SetSequenceState sets the state of @sequence in @gesture. Sequences start
	// in state K_EVENT_SEQUENCE_NONE, and whenever they change state, they can
	// never go back to that state. Likewise, sequences in state
	// K_EVENT_SEQUENCE_DENIED cannot turn back to a not denied state. With
	// these rules, the lifetime of an event sequence is constrained to the next
	// four:
	//
	// * None * None → Denied * None → Claimed * None → Claimed → Denied
	//
	// Note: Due to event handling ordering, it may be unsafe to set the state
	// on another gesture within a Gesture::begin signal handler, as the
	// callback might be executed before the other gesture knows about the
	// sequence. A safe way to perform this could be:
	//
	//    static void
	//    first_gesture_begin_cb (GtkGesture       *first_gesture,
	//                            GdkEventSequence *sequence,
	//                            gpointer          user_data)
	//    {
	//      gtk_gesture_set_sequence_state (first_gesture, sequence, GTK_EVENT_SEQUENCE_CLAIMED);
	//      gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
	//    }
	//
	//    static void
	//    second_gesture_begin_cb (GtkGesture       *second_gesture,
	//                             GdkEventSequence *sequence,
	//                             gpointer          user_data)
	//    {
	//      if (gtk_gesture_get_sequence_state (first_gesture, sequence) == GTK_EVENT_SEQUENCE_CLAIMED)
	//        gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
	//    }
	//
	// If both gestures are in the same group, just set the state on the gesture
	// emitting the event, the sequence will be already be initialized to the
	// group's global state when the second gesture processes the event.
	//
	// This method is inherited from Gesture
	SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool
	// SetState sets the state of all sequences that @gesture is currently
	// interacting with. See gtk_gesture_set_sequence_state() for more details
	// on sequence states.
	//
	// This method is inherited from Gesture
	SetState(state EventSequenceState) bool
	// SetWindow sets a specific window to receive events about, so @gesture
	// will effectively handle only events targeting @window, or a child of it.
	// @window must pertain to gtk_event_controller_get_widget().
	//
	// This method is inherited from Gesture
	SetWindow(window gdk.Window)
	// Ungroup separates @gesture into an isolated group.
	//
	// This method is inherited from Gesture
	Ungroup()
	// GetPropagationPhase gets the propagation phase at which @controller
	// handles events.
	//
	// This method is inherited from EventController
	GetPropagationPhase() PropagationPhase
	// GetWidget returns the Widget this controller relates to.
	//
	// This method is inherited from EventController
	GetWidget() Widget
	// Reset resets the @controller to a clean state. Every interaction the
	// controller did through EventController::handle-event will be dropped at
	// this point.
	//
	// This method is inherited from EventController
	Reset()
	// SetPropagationPhase sets the propagation phase at which a controller
	// handles events.
	//
	// If @phase is GTK_PHASE_NONE, no automatic event handling will be
	// performed, but other additional gesture maintenance will. In that phase,
	// the events can be managed by calling gtk_event_controller_handle_event().
	//
	// This method is inherited from EventController
	SetPropagationPhase(phase PropagationPhase)

	// Button returns the button number @gesture listens for, or 0 if @gesture
	// reacts to any button press.
	Button() uint
	// CurrentButton returns the button number currently interacting with
	// @gesture, or 0 if there is none.
	CurrentButton() uint
	// CurrentSequence returns the event sequence currently interacting with
	// @gesture. This is only meaningful if gtk_gesture_is_active() returns
	// true.
	CurrentSequence() *gdk.EventSequence
	// Exclusive gets whether a gesture is exclusive. For more information, see
	// gtk_gesture_single_set_exclusive().
	Exclusive() bool
	// TouchOnly returns true if the gesture is only triggered by touch events.
	TouchOnly() bool
	// SetButton sets the button number @gesture listens to. If non-0, every
	// button press from a different button number will be ignored. Touch events
	// implicitly match with button 1.
	SetButton(button uint)
	// SetExclusive sets whether @gesture is exclusive. An exclusive gesture
	// will only handle pointer and "pointer emulated" touch events, so at any
	// given time, there is only one sequence able to interact with those.
	SetExclusive(exclusive bool)
	// SetTouchOnly: if @touch_only is true, @gesture will only handle events of
	// type K_TOUCH_BEGIN, K_TOUCH_UPDATE or K_TOUCH_END. If false, mouse events
	// will be handled too.
	SetTouchOnly(touchOnly bool)
}

// gestureSingle implements the GestureSingle interface.
type gestureSingle struct {
	*externglib.Object
}

var _ GestureSingle = (*gestureSingle)(nil)

// WrapGestureSingle wraps a GObject to a type that implements
// interface GestureSingle. It is primarily used internally.
func WrapGestureSingle(obj *externglib.Object) GestureSingle {
	return gestureSingle{obj}
}

func marshalGestureSingle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureSingle(obj), nil
}

func (g gestureSingle) AsGesture() Gesture {
	return WrapGesture(gextras.InternObject(g))
}

func (g gestureSingle) GetBoundingBox() (gdk.Rectangle, bool) {
	return WrapGesture(gextras.InternObject(g)).GetBoundingBox()
}

func (g gestureSingle) GetBoundingBoxCenter() (x float64, y float64, ok bool) {
	return WrapGesture(gextras.InternObject(g)).GetBoundingBoxCenter()
}

func (g gestureSingle) GetDevice() gdk.Device {
	return WrapGesture(gextras.InternObject(g)).GetDevice()
}

func (g gestureSingle) GetLastUpdatedSequence() *gdk.EventSequence {
	return WrapGesture(gextras.InternObject(g)).GetLastUpdatedSequence()
}

func (g gestureSingle) GetPoint(sequence *gdk.EventSequence) (x float64, y float64, ok bool) {
	return WrapGesture(gextras.InternObject(g)).GetPoint(sequence)
}

func (g gestureSingle) GetSequenceState(sequence *gdk.EventSequence) EventSequenceState {
	return WrapGesture(gextras.InternObject(g)).GetSequenceState(sequence)
}

func (g gestureSingle) GetWindow() gdk.Window {
	return WrapGesture(gextras.InternObject(g)).GetWindow()
}

func (g gestureSingle) Group(gesture Gesture) {
	WrapGesture(gextras.InternObject(g)).Group(gesture)
}

func (g gestureSingle) HandlesSequence(sequence *gdk.EventSequence) bool {
	return WrapGesture(gextras.InternObject(g)).HandlesSequence(sequence)
}

func (g gestureSingle) IsActive() bool {
	return WrapGesture(gextras.InternObject(g)).IsActive()
}

func (g gestureSingle) IsGroupedWith(other Gesture) bool {
	return WrapGesture(gextras.InternObject(g)).IsGroupedWith(other)
}

func (g gestureSingle) IsRecognized() bool {
	return WrapGesture(gextras.InternObject(g)).IsRecognized()
}

func (g gestureSingle) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	return WrapGesture(gextras.InternObject(g)).SetSequenceState(sequence, state)
}

func (g gestureSingle) SetState(state EventSequenceState) bool {
	return WrapGesture(gextras.InternObject(g)).SetState(state)
}

func (g gestureSingle) SetWindow(window gdk.Window) {
	WrapGesture(gextras.InternObject(g)).SetWindow(window)
}

func (g gestureSingle) Ungroup() {
	WrapGesture(gextras.InternObject(g)).Ungroup()
}

func (c gestureSingle) GetPropagationPhase() PropagationPhase {
	return WrapEventController(gextras.InternObject(c)).GetPropagationPhase()
}

func (c gestureSingle) GetWidget() Widget {
	return WrapEventController(gextras.InternObject(c)).GetWidget()
}

func (c gestureSingle) Reset() {
	WrapEventController(gextras.InternObject(c)).Reset()
}

func (c gestureSingle) SetPropagationPhase(phase PropagationPhase) {
	WrapEventController(gextras.InternObject(c)).SetPropagationPhase(phase)
}

func (g gestureSingle) Button() uint {
	var _arg0 *C.GtkGestureSingle // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_single_get_button(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (g gestureSingle) CurrentButton() uint {
	var _arg0 *C.GtkGestureSingle // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_single_get_current_button(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (g gestureSingle) CurrentSequence() *gdk.EventSequence {
	var _arg0 *C.GtkGestureSingle // out
	var _cret *C.GdkEventSequence // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_single_get_current_sequence(_arg0)

	var _eventSequence *gdk.EventSequence // out

	_eventSequence = (*gdk.EventSequence)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_eventSequence, func(v *gdk.EventSequence) {
		C.free(unsafe.Pointer(v))
	})

	return _eventSequence
}

func (g gestureSingle) Exclusive() bool {
	var _arg0 *C.GtkGestureSingle // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_single_get_exclusive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (g gestureSingle) TouchOnly() bool {
	var _arg0 *C.GtkGestureSingle // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_single_get_touch_only(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (g gestureSingle) SetButton(button uint) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.guint             // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))
	_arg1 = C.guint(button)

	C.gtk_gesture_single_set_button(_arg0, _arg1)
}

func (g gestureSingle) SetExclusive(exclusive bool) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))
	if exclusive {
		_arg1 = C.TRUE
	}

	C.gtk_gesture_single_set_exclusive(_arg0, _arg1)
}

func (g gestureSingle) SetTouchOnly(touchOnly bool) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))
	if touchOnly {
		_arg1 = C.TRUE
	}

	C.gtk_gesture_single_set_touch_only(_arg0, _arg1)
}
