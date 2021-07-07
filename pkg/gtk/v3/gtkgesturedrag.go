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
		{T: externglib.Type(C.gtk_gesture_drag_get_type()), F: marshalGestureDrag},
	})
}

// GestureDrag is a Gesture implementation that recognizes drag operations. The
// drag operation itself can be tracked throught the GestureDrag::drag-begin,
// GestureDrag::drag-update and GestureDrag::drag-end signals, or the relevant
// coordinates be extracted through gtk_gesture_drag_get_offset() and
// gtk_gesture_drag_get_start_point().
type GestureDrag interface {
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

	// Offset: if the @gesture is active, this function returns true and fills
	// in @x and @y with the coordinates of the current point, as an offset to
	// the starting drag point.
	Offset() (x float64, y float64, ok bool)
	// StartPoint: if the @gesture is active, this function returns true and
	// fills in @x and @y with the drag start coordinates, in window-relative
	// coordinates.
	StartPoint() (x float64, y float64, ok bool)
}

// gestureDrag implements the GestureDrag interface.
type gestureDrag struct {
	*externglib.Object
}

var _ GestureDrag = (*gestureDrag)(nil)

// WrapGestureDrag wraps a GObject to a type that implements
// interface GestureDrag. It is primarily used internally.
func WrapGestureDrag(obj *externglib.Object) GestureDrag {
	return gestureDrag{obj}
}

func marshalGestureDrag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureDrag(obj), nil
}

// NewGestureDrag returns a newly created Gesture that recognizes drags.
func NewGestureDrag(widget Widget) GestureDrag {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_drag_new(_arg1)

	var _gestureDrag GestureDrag // out

	_gestureDrag = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(GestureDrag)

	return _gestureDrag
}

func (g gestureDrag) AsGestureSingle() GestureSingle {
	return WrapGestureSingle(gextras.InternObject(g))
}

func (g gestureDrag) GetButton() uint {
	return WrapGestureSingle(gextras.InternObject(g)).GetButton()
}

func (g gestureDrag) GetCurrentButton() uint {
	return WrapGestureSingle(gextras.InternObject(g)).GetCurrentButton()
}

func (g gestureDrag) GetCurrentSequence() *gdk.EventSequence {
	return WrapGestureSingle(gextras.InternObject(g)).GetCurrentSequence()
}

func (g gestureDrag) GetExclusive() bool {
	return WrapGestureSingle(gextras.InternObject(g)).GetExclusive()
}

func (g gestureDrag) GetTouchOnly() bool {
	return WrapGestureSingle(gextras.InternObject(g)).GetTouchOnly()
}

func (g gestureDrag) SetButton(button uint) {
	WrapGestureSingle(gextras.InternObject(g)).SetButton(button)
}

func (g gestureDrag) SetExclusive(exclusive bool) {
	WrapGestureSingle(gextras.InternObject(g)).SetExclusive(exclusive)
}

func (g gestureDrag) SetTouchOnly(touchOnly bool) {
	WrapGestureSingle(gextras.InternObject(g)).SetTouchOnly(touchOnly)
}

func (g gestureDrag) GetBoundingBox() (gdk.Rectangle, bool) {
	return WrapGesture(gextras.InternObject(g)).GetBoundingBox()
}

func (g gestureDrag) GetBoundingBoxCenter() (x float64, y float64, ok bool) {
	return WrapGesture(gextras.InternObject(g)).GetBoundingBoxCenter()
}

func (g gestureDrag) GetDevice() gdk.Device {
	return WrapGesture(gextras.InternObject(g)).GetDevice()
}

func (g gestureDrag) GetLastUpdatedSequence() *gdk.EventSequence {
	return WrapGesture(gextras.InternObject(g)).GetLastUpdatedSequence()
}

func (g gestureDrag) GetPoint(sequence *gdk.EventSequence) (x float64, y float64, ok bool) {
	return WrapGesture(gextras.InternObject(g)).GetPoint(sequence)
}

func (g gestureDrag) GetSequenceState(sequence *gdk.EventSequence) EventSequenceState {
	return WrapGesture(gextras.InternObject(g)).GetSequenceState(sequence)
}

func (g gestureDrag) GetWindow() gdk.Window {
	return WrapGesture(gextras.InternObject(g)).GetWindow()
}

func (g gestureDrag) Group(gesture Gesture) {
	WrapGesture(gextras.InternObject(g)).Group(gesture)
}

func (g gestureDrag) HandlesSequence(sequence *gdk.EventSequence) bool {
	return WrapGesture(gextras.InternObject(g)).HandlesSequence(sequence)
}

func (g gestureDrag) IsActive() bool {
	return WrapGesture(gextras.InternObject(g)).IsActive()
}

func (g gestureDrag) IsGroupedWith(other Gesture) bool {
	return WrapGesture(gextras.InternObject(g)).IsGroupedWith(other)
}

func (g gestureDrag) IsRecognized() bool {
	return WrapGesture(gextras.InternObject(g)).IsRecognized()
}

func (g gestureDrag) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	return WrapGesture(gextras.InternObject(g)).SetSequenceState(sequence, state)
}

func (g gestureDrag) SetState(state EventSequenceState) bool {
	return WrapGesture(gextras.InternObject(g)).SetState(state)
}

func (g gestureDrag) SetWindow(window gdk.Window) {
	WrapGesture(gextras.InternObject(g)).SetWindow(window)
}

func (g gestureDrag) Ungroup() {
	WrapGesture(gextras.InternObject(g)).Ungroup()
}

func (c gestureDrag) GetPropagationPhase() PropagationPhase {
	return WrapEventController(gextras.InternObject(c)).GetPropagationPhase()
}

func (c gestureDrag) GetWidget() Widget {
	return WrapEventController(gextras.InternObject(c)).GetWidget()
}

func (c gestureDrag) Reset() {
	WrapEventController(gextras.InternObject(c)).Reset()
}

func (c gestureDrag) SetPropagationPhase(phase PropagationPhase) {
	WrapEventController(gextras.InternObject(c)).SetPropagationPhase(phase)
}

func (g gestureDrag) Offset() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out
	var _arg1 C.gdouble         // in
	var _arg2 C.gdouble         // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_drag_get_offset(_arg0, &_arg1, &_arg2)

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

func (g gestureDrag) StartPoint() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out
	var _arg1 C.gdouble         // in
	var _arg2 C.gdouble         // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_drag_get_start_point(_arg0, &_arg1, &_arg2)

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
