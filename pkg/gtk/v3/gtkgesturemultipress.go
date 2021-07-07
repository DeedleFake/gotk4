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
		{T: externglib.Type(C.gtk_gesture_multi_press_get_type()), F: marshalGestureMultiPress},
	})
}

// GestureMultiPress is a Gesture implementation able to recognize multiple
// clicks on a nearby zone, which can be listened for through the
// GestureMultiPress::pressed signal. Whenever time or distance between clicks
// exceed the GTK+ defaults, GestureMultiPress::stopped is emitted, and the
// click counter is reset.
//
// Callers may also restrict the area that is considered valid for a >1
// touch/button press through gtk_gesture_multi_press_set_area(), so any click
// happening outside that area is considered to be a first click of its own.
type GestureMultiPress interface {
	GestureSingle

	// AsGestureSingle casts the class to the GestureSingle interface.
	AsGestureSingle() GestureSingle

	// GetButton returns the button number @gesture listens for, or 0 if
	// @gesture reacts to any button press.
	//
	// This method is inherited from GestureSingle
	GetButton() uint
	// GetCurrentButton returns the button number currently interacting with
	// @gesture, or 0 if there is none.
	//
	// This method is inherited from GestureSingle
	GetCurrentButton() uint
	// GetCurrentSequence returns the event sequence currently interacting with
	// @gesture. This is only meaningful if gtk_gesture_is_active() returns
	// true.
	//
	// This method is inherited from GestureSingle
	GetCurrentSequence() *gdk.EventSequence
	// GetExclusive gets whether a gesture is exclusive. For more information,
	// see gtk_gesture_single_set_exclusive().
	//
	// This method is inherited from GestureSingle
	GetExclusive() bool
	// GetTouchOnly returns true if the gesture is only triggered by touch
	// events.
	//
	// This method is inherited from GestureSingle
	GetTouchOnly() bool
	// SetButton sets the button number @gesture listens to. If non-0, every
	// button press from a different button number will be ignored. Touch events
	// implicitly match with button 1.
	//
	// This method is inherited from GestureSingle
	SetButton(button uint)
	// SetExclusive sets whether @gesture is exclusive. An exclusive gesture
	// will only handle pointer and "pointer emulated" touch events, so at any
	// given time, there is only one sequence able to interact with those.
	//
	// This method is inherited from GestureSingle
	SetExclusive(exclusive bool)
	// SetTouchOnly: if @touch_only is true, @gesture will only handle events of
	// type K_TOUCH_BEGIN, K_TOUCH_UPDATE or K_TOUCH_END. If false, mouse events
	// will be handled too.
	//
	// This method is inherited from GestureSingle
	SetTouchOnly(touchOnly bool)
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

	// Area: if an area was set through gtk_gesture_multi_press_set_area(), this
	// function will return true and fill in @rect with the press area. See
	// gtk_gesture_multi_press_set_area() for more details on what the press
	// area represents.
	Area() (gdk.Rectangle, bool)
	// SetArea: if @rect is non-nil, the press area will be checked to be
	// confined within the rectangle, otherwise the button count will be reset
	// so the press is seen as being the first one. If @rect is nil, the area
	// will be reset to an unrestricted state.
	//
	// Note: The rectangle is only used to determine whether any non-first click
	// falls within the expected area. This is not akin to an input shape.
	SetArea(rect *gdk.Rectangle)
}

// gestureMultiPress implements the GestureMultiPress interface.
type gestureMultiPress struct {
	*externglib.Object
}

var _ GestureMultiPress = (*gestureMultiPress)(nil)

// WrapGestureMultiPress wraps a GObject to a type that implements
// interface GestureMultiPress. It is primarily used internally.
func WrapGestureMultiPress(obj *externglib.Object) GestureMultiPress {
	return gestureMultiPress{obj}
}

func marshalGestureMultiPress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureMultiPress(obj), nil
}

// NewGestureMultiPress returns a newly created Gesture that recognizes single
// and multiple presses.
func NewGestureMultiPress(widget Widget) GestureMultiPress {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_multi_press_new(_arg1)

	var _gestureMultiPress GestureMultiPress // out

	_gestureMultiPress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(GestureMultiPress)

	return _gestureMultiPress
}

func (g gestureMultiPress) AsGestureSingle() GestureSingle {
	return WrapGestureSingle(gextras.InternObject(g))
}

func (g gestureMultiPress) GetButton() uint {
	return WrapGestureSingle(gextras.InternObject(g)).GetButton()
}

func (g gestureMultiPress) GetCurrentButton() uint {
	return WrapGestureSingle(gextras.InternObject(g)).GetCurrentButton()
}

func (g gestureMultiPress) GetCurrentSequence() *gdk.EventSequence {
	return WrapGestureSingle(gextras.InternObject(g)).GetCurrentSequence()
}

func (g gestureMultiPress) GetExclusive() bool {
	return WrapGestureSingle(gextras.InternObject(g)).GetExclusive()
}

func (g gestureMultiPress) GetTouchOnly() bool {
	return WrapGestureSingle(gextras.InternObject(g)).GetTouchOnly()
}

func (g gestureMultiPress) SetButton(button uint) {
	WrapGestureSingle(gextras.InternObject(g)).SetButton(button)
}

func (g gestureMultiPress) SetExclusive(exclusive bool) {
	WrapGestureSingle(gextras.InternObject(g)).SetExclusive(exclusive)
}

func (g gestureMultiPress) SetTouchOnly(touchOnly bool) {
	WrapGestureSingle(gextras.InternObject(g)).SetTouchOnly(touchOnly)
}

func (g gestureMultiPress) GetBoundingBox() (gdk.Rectangle, bool) {
	return WrapGesture(gextras.InternObject(g)).GetBoundingBox()
}

func (g gestureMultiPress) GetBoundingBoxCenter() (x float64, y float64, ok bool) {
	return WrapGesture(gextras.InternObject(g)).GetBoundingBoxCenter()
}

func (g gestureMultiPress) GetDevice() gdk.Device {
	return WrapGesture(gextras.InternObject(g)).GetDevice()
}

func (g gestureMultiPress) GetLastUpdatedSequence() *gdk.EventSequence {
	return WrapGesture(gextras.InternObject(g)).GetLastUpdatedSequence()
}

func (g gestureMultiPress) GetPoint(sequence *gdk.EventSequence) (x float64, y float64, ok bool) {
	return WrapGesture(gextras.InternObject(g)).GetPoint(sequence)
}

func (g gestureMultiPress) GetSequenceState(sequence *gdk.EventSequence) EventSequenceState {
	return WrapGesture(gextras.InternObject(g)).GetSequenceState(sequence)
}

func (g gestureMultiPress) GetWindow() gdk.Window {
	return WrapGesture(gextras.InternObject(g)).GetWindow()
}

func (g gestureMultiPress) Group(gesture Gesture) {
	WrapGesture(gextras.InternObject(g)).Group(gesture)
}

func (g gestureMultiPress) HandlesSequence(sequence *gdk.EventSequence) bool {
	return WrapGesture(gextras.InternObject(g)).HandlesSequence(sequence)
}

func (g gestureMultiPress) IsActive() bool {
	return WrapGesture(gextras.InternObject(g)).IsActive()
}

func (g gestureMultiPress) IsGroupedWith(other Gesture) bool {
	return WrapGesture(gextras.InternObject(g)).IsGroupedWith(other)
}

func (g gestureMultiPress) IsRecognized() bool {
	return WrapGesture(gextras.InternObject(g)).IsRecognized()
}

func (g gestureMultiPress) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	return WrapGesture(gextras.InternObject(g)).SetSequenceState(sequence, state)
}

func (g gestureMultiPress) SetState(state EventSequenceState) bool {
	return WrapGesture(gextras.InternObject(g)).SetState(state)
}

func (g gestureMultiPress) SetWindow(window gdk.Window) {
	WrapGesture(gextras.InternObject(g)).SetWindow(window)
}

func (g gestureMultiPress) Ungroup() {
	WrapGesture(gextras.InternObject(g)).Ungroup()
}

func (c gestureMultiPress) GetPropagationPhase() PropagationPhase {
	return WrapEventController(gextras.InternObject(c)).GetPropagationPhase()
}

func (c gestureMultiPress) GetWidget() Widget {
	return WrapEventController(gextras.InternObject(c)).GetWidget()
}

func (c gestureMultiPress) Reset() {
	WrapEventController(gextras.InternObject(c)).Reset()
}

func (c gestureMultiPress) SetPropagationPhase(phase PropagationPhase) {
	WrapEventController(gextras.InternObject(c)).SetPropagationPhase(phase)
}

func (g gestureMultiPress) Area() (gdk.Rectangle, bool) {
	var _arg0 *C.GtkGestureMultiPress // out
	var _arg1 C.GdkRectangle          // in
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkGestureMultiPress)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_multi_press_get_area(_arg0, &_arg1)

	var _rect gdk.Rectangle // out
	var _ok bool            // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_rect = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

func (g gestureMultiPress) SetArea(rect *gdk.Rectangle) {
	var _arg0 *C.GtkGestureMultiPress // out
	var _arg1 *C.GdkRectangle         // out

	_arg0 = (*C.GtkGestureMultiPress)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect))

	C.gtk_gesture_multi_press_set_area(_arg0, _arg1)
}
