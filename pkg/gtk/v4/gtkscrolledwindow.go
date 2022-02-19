// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_ScrolledWindow_ConnectScrollChild(gpointer, GtkScrollType, gboolean, guintptr);
// extern void _gotk4_gtk4_ScrolledWindow_ConnectEdgeOvershot(gpointer, GtkPositionType, guintptr);
// extern void _gotk4_gtk4_ScrolledWindow_ConnectEdgeReached(gpointer, GtkPositionType, guintptr);
// extern void _gotk4_gtk4_ScrolledWindow_ConnectMoveFocusOut(gpointer, GtkDirectionType, guintptr);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_corner_type_get_type()), F: marshalCornerType},
		{T: externglib.Type(C.gtk_policy_type_get_type()), F: marshalPolicyType},
		{T: externglib.Type(C.gtk_scrolled_window_get_type()), F: marshalScrolledWindower},
	})
}

// CornerType specifies which corner a child widget should be placed in when
// packed into a GtkScrolledWindow.
//
// This is effectively the opposite of where the scroll bars are placed.
type CornerType C.gint

const (
	// CornerTopLeft: place the scrollbars on the right and bottom of the widget
	// (default behaviour).
	CornerTopLeft CornerType = iota
	// CornerBottomLeft: place the scrollbars on the top and right of the
	// widget.
	CornerBottomLeft
	// CornerTopRight: place the scrollbars on the left and bottom of the
	// widget.
	CornerTopRight
	// CornerBottomRight: place the scrollbars on the top and left of the
	// widget.
	CornerBottomRight
)

func marshalCornerType(p uintptr) (interface{}, error) {
	return CornerType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CornerType.
func (c CornerType) String() string {
	switch c {
	case CornerTopLeft:
		return "TopLeft"
	case CornerBottomLeft:
		return "BottomLeft"
	case CornerTopRight:
		return "TopRight"
	case CornerBottomRight:
		return "BottomRight"
	default:
		return fmt.Sprintf("CornerType(%d)", c)
	}
}

// PolicyType determines how the size should be computed to achieve the one of
// the visibility mode for the scrollbars.
type PolicyType C.gint

const (
	// PolicyAlways: scrollbar is always visible. The view size is independent
	// of the content.
	PolicyAlways PolicyType = iota
	// PolicyAutomatic: scrollbar will appear and disappear as necessary. For
	// example, when all of a GtkTreeView can not be seen.
	PolicyAutomatic
	// PolicyNever: scrollbar should never appear. In this mode the content
	// determines the size.
	PolicyNever
	// PolicyExternal: don't show a scrollbar, but don't force the size to
	// follow the content. This can be used e.g. to make multiple scrolled
	// windows share a scrollbar.
	PolicyExternal
)

func marshalPolicyType(p uintptr) (interface{}, error) {
	return PolicyType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PolicyType.
func (p PolicyType) String() string {
	switch p {
	case PolicyAlways:
		return "Always"
	case PolicyAutomatic:
		return "Automatic"
	case PolicyNever:
		return "Never"
	case PolicyExternal:
		return "External"
	default:
		return fmt.Sprintf("PolicyType(%d)", p)
	}
}

// ScrolledWindow: GtkScrolledWindow is a container that makes its child
// scrollable.
//
// It does so using either internally added scrollbars or externally associated
// adjustments, and optionally draws a frame around the child.
//
// Widgets with native scrolling support, i.e. those whose classes implement the
// gtk.Scrollable interface, are added directly. For other types of widget, the
// class gtk.Viewport acts as an adaptor, giving scrollability to other widgets.
// gtk.ScrolledWindow.SetChild() intelligently accounts for whether or not the
// added child is a GtkScrollable. If it isn’t, then it wraps the child in a
// GtkViewport. Therefore, you can just add any child widget and not worry about
// the details.
//
// If gtk.ScrolledWindow.SetChild() has added a GtkViewport for you, you can
// remove both your added child widget from the GtkViewport, and the GtkViewport
// from the GtkScrolledWindow, like this:
//
//    GtkWidget *scrolled_window = gtk_scrolled_window_new ();
//    GtkWidget *child_widget = gtk_button_new ();
//
//    // GtkButton is not a GtkScrollable, so GtkScrolledWindow will automatically
//    // add a GtkViewport.
//    gtk_box_append (GTK_BOX (scrolled_window), child_widget);
//
//    // Either of these will result in child_widget being unparented:
//    gtk_box_remove (GTK_BOX (scrolled_window), child_widget);
//    // or
//    gtk_box_remove (GTK_BOX (scrolled_window),
//                          gtk_bin_get_child (GTK_BIN (scrolled_window)));
//
//
// Unless gtk.ScrolledWindow:hscrollbar-policy and
// gtk.ScrolledWindow:vscrollbar-policy are GTK_POLICY_NEVER or
// GTK_POLICY_EXTERNAL, GtkScrolledWindow adds internal GtkScrollbar widgets
// around its child. The scroll position of the child, and if applicable the
// scrollbars, is controlled by the gtk.ScrolledWindow:hadjustment and
// gtk.ScrolledWindow:vadjustment that are associated with the
// GtkScrolledWindow. See the docs on gtk.Scrollbar for the details, but note
// that the “step_increment” and “page_increment” fields are only effective if
// the policy causes scrollbars to be present.
//
// If a GtkScrolledWindow doesn’t behave quite as you would like, or doesn’t
// have exactly the right layout, it’s very possible to set up your own
// scrolling with GtkScrollbar and for example a GtkGrid.
//
//
// Touch support
//
// GtkScrolledWindow has built-in support for touch devices. When a touchscreen
// is used, swiping will move the scrolled window, and will expose 'kinetic'
// behavior. This can be turned off with the
// gtk.ScrolledWindow:kinetic-scrolling property if it is undesired.
//
// GtkScrolledWindow also displays visual 'overshoot' indication when the
// content is pulled beyond the end, and this situation can be captured with the
// gtk.ScrolledWindow::edge-overshot signal.
//
// If no mouse device is present, the scrollbars will overlaid as narrow,
// auto-hiding indicators over the content. If traditional scrollbars are
// desired although no mouse is present, this behaviour can be turned off with
// the gtk.ScrolledWindow:overlay-scrolling property.
//
//
// CSS nodes
//
// GtkScrolledWindow has a main CSS node with name scrolledwindow. It gets a
// .frame style class added when gtk.ScrolledWindow:has-frame is TRUE.
//
// It uses subnodes with names overshoot and undershoot to draw the overflow and
// underflow indications. These nodes get the .left, .right, .top or .bottom
// style class added depending on where the indication is drawn.
//
// GtkScrolledWindow also sets the positional style classes (.left, .right,
// .top, .bottom) and style classes related to overlay scrolling
// (.overlay-indicator, .dragging, .hovering) on its scrollbars.
//
// If both scrollbars are visible, the area where they meet is drawn with a
// subnode named junction.
//
//
// Accessibility
//
// GtkScrolledWindow uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type ScrolledWindow struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*ScrolledWindow)(nil)
)

func wrapScrolledWindow(obj *externglib.Object) *ScrolledWindow {
	return &ScrolledWindow{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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
	}
}

func marshalScrolledWindower(p uintptr) (interface{}, error) {
	return wrapScrolledWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_ScrolledWindow_ConnectEdgeOvershot
func _gotk4_gtk4_ScrolledWindow_ConnectEdgeOvershot(arg0 C.gpointer, arg1 C.GtkPositionType, arg2 C.guintptr) {
	var f func(pos PositionType)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pos PositionType))
	}

	var _pos PositionType // out

	_pos = PositionType(arg1)

	f(_pos)
}

// ConnectEdgeOvershot: emitted whenever user initiated scrolling makes the
// scrolled window firmly surpass the limits defined by the adjustment in that
// orientation.
//
// A similar behavior without edge resistance is provided by the
// gtk.ScrolledWindow::edge-reached signal.
//
// Note: The pos argument is LTR/RTL aware, so callers should be aware too if
// intending to provide behavior on horizontal edges.
func (scrolledWindow *ScrolledWindow) ConnectEdgeOvershot(f func(pos PositionType)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(scrolledWindow, "edge-overshot", false, unsafe.Pointer(C._gotk4_gtk4_ScrolledWindow_ConnectEdgeOvershot), f)
}

//export _gotk4_gtk4_ScrolledWindow_ConnectEdgeReached
func _gotk4_gtk4_ScrolledWindow_ConnectEdgeReached(arg0 C.gpointer, arg1 C.GtkPositionType, arg2 C.guintptr) {
	var f func(pos PositionType)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pos PositionType))
	}

	var _pos PositionType // out

	_pos = PositionType(arg1)

	f(_pos)
}

// ConnectEdgeReached: emitted whenever user-initiated scrolling makes the
// scrolled window exactly reach the lower or upper limits defined by the
// adjustment in that orientation.
//
// A similar behavior with edge resistance is provided by the
// gtk.ScrolledWindow::edge-overshot signal.
//
// Note: The pos argument is LTR/RTL aware, so callers should be aware too if
// intending to provide behavior on horizontal edges.
func (scrolledWindow *ScrolledWindow) ConnectEdgeReached(f func(pos PositionType)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(scrolledWindow, "edge-reached", false, unsafe.Pointer(C._gotk4_gtk4_ScrolledWindow_ConnectEdgeReached), f)
}

//export _gotk4_gtk4_ScrolledWindow_ConnectMoveFocusOut
func _gotk4_gtk4_ScrolledWindow_ConnectMoveFocusOut(arg0 C.gpointer, arg1 C.GtkDirectionType, arg2 C.guintptr) {
	var f func(directionType DirectionType)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(directionType DirectionType))
	}

	var _directionType DirectionType // out

	_directionType = DirectionType(arg1)

	f(_directionType)
}

// ConnectMoveFocusOut: emitted when focus is moved away from the scrolled
// window by a keybinding.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are Ctrl + Tab to move forward and Ctrl
// + Shift + Tab to move backward.
func (scrolledWindow *ScrolledWindow) ConnectMoveFocusOut(f func(directionType DirectionType)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(scrolledWindow, "move-focus-out", false, unsafe.Pointer(C._gotk4_gtk4_ScrolledWindow_ConnectMoveFocusOut), f)
}

//export _gotk4_gtk4_ScrolledWindow_ConnectScrollChild
func _gotk4_gtk4_ScrolledWindow_ConnectScrollChild(arg0 C.gpointer, arg1 C.GtkScrollType, arg2 C.gboolean, arg3 C.guintptr) (cret C.gboolean) {
	var f func(scroll ScrollType, horizontal bool) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(scroll ScrollType, horizontal bool) (ok bool))
	}

	var _scroll ScrollType // out
	var _horizontal bool   // out

	_scroll = ScrollType(arg1)
	if arg2 != 0 {
		_horizontal = true
	}

	ok := f(_scroll, _horizontal)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectScrollChild: emitted when a keybinding that scrolls is pressed.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The horizontal or vertical adjustment is updated which triggers a signal that
// the scrolled window’s child may listen to and scroll itself.
func (scrolledWindow *ScrolledWindow) ConnectScrollChild(f func(scroll ScrollType, horizontal bool) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(scrolledWindow, "scroll-child", false, unsafe.Pointer(C._gotk4_gtk4_ScrolledWindow_ConnectScrollChild), f)
}

// NewScrolledWindow creates a new scrolled window.
//
// The function returns the following values:
//
//    - scrolledWindow: new scrolled window.
//
func NewScrolledWindow() *ScrolledWindow {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_scrolled_window_new()

	var _scrolledWindow *ScrolledWindow // out

	_scrolledWindow = wrapScrolledWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _scrolledWindow
}

// Child gets the child widget of scrolled_window.
//
// The function returns the following values:
//
//    - widget (optional): child widget of scrolled_window.
//
func (scrolledWindow *ScrolledWindow) Child() Widgetter {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_child(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// HAdjustment returns the horizontal scrollbar’s adjustment.
//
// This is the adjustment used to connect the horizontal scrollbar to the child
// widget’s horizontal scroll functionality.
//
// The function returns the following values:
//
//    - adjustment: horizontal GtkAdjustment.
//
func (scrolledWindow *ScrolledWindow) HAdjustment() *Adjustment {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkAdjustment     // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_hadjustment(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// HasFrame gets whether the scrolled window draws a frame.
//
// The function returns the following values:
//
//    - ok: TRUE if the scrolled_window has a frame.
//
func (scrolledWindow *ScrolledWindow) HasFrame() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_has_frame(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HScrollbar returns the horizontal scrollbar of scrolled_window.
//
// The function returns the following values:
//
//    - widget: horizontal scrollbar of the scrolled window.
//
func (scrolledWindow *ScrolledWindow) HScrollbar() Widgetter {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_hscrollbar(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// KineticScrolling returns the specified kinetic scrolling behavior.
//
// The function returns the following values:
//
//    - ok: scrolling behavior flags.
//
func (scrolledWindow *ScrolledWindow) KineticScrolling() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_kinetic_scrolling(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxContentHeight returns the maximum content height set.
//
// The function returns the following values:
//
//    - gint: maximum content height, or -1.
//
func (scrolledWindow *ScrolledWindow) MaxContentHeight() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.int                // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_max_content_height(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MaxContentWidth returns the maximum content width set.
//
// The function returns the following values:
//
//    - gint: maximum content width, or -1.
//
func (scrolledWindow *ScrolledWindow) MaxContentWidth() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.int                // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_max_content_width(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MinContentHeight gets the minimal content height of scrolled_window.
//
// The function returns the following values:
//
//    - gint: minimal content height.
//
func (scrolledWindow *ScrolledWindow) MinContentHeight() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.int                // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_min_content_height(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MinContentWidth gets the minimum content width of scrolled_window.
//
// The function returns the following values:
//
//    - gint: minimum content width.
//
func (scrolledWindow *ScrolledWindow) MinContentWidth() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.int                // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_min_content_width(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// OverlayScrolling returns whether overlay scrolling is enabled for this
// scrolled window.
//
// The function returns the following values:
//
//    - ok: TRUE if overlay scrolling is enabled.
//
func (scrolledWindow *ScrolledWindow) OverlayScrolling() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_overlay_scrolling(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Placement gets the placement of the contents with respect to the scrollbars.
//
// The function returns the following values:
//
//    - cornerType: current placement value.
//
func (scrolledWindow *ScrolledWindow) Placement() CornerType {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.GtkCornerType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_placement(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _cornerType CornerType // out

	_cornerType = CornerType(_cret)

	return _cornerType
}

// Policy retrieves the current policy values for the horizontal and vertical
// scrollbars.
//
// See gtk.ScrolledWindow.SetPolicy().
//
// The function returns the following values:
//
//    - hscrollbarPolicy (optional): location to store the policy for the
//      horizontal scrollbar, or NULL.
//    - vscrollbarPolicy (optional): location to store the policy for the
//      vertical scrollbar, or NULL.
//
func (scrolledWindow *ScrolledWindow) Policy() (hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkPolicyType      // in
	var _arg2 C.GtkPolicyType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	C.gtk_scrolled_window_get_policy(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(scrolledWindow)

	var _hscrollbarPolicy PolicyType // out
	var _vscrollbarPolicy PolicyType // out

	_hscrollbarPolicy = PolicyType(_arg1)
	_vscrollbarPolicy = PolicyType(_arg2)

	return _hscrollbarPolicy, _vscrollbarPolicy
}

// PropagateNaturalHeight reports whether the natural height of the child will
// be calculated and propagated through the scrolled window’s requested natural
// height.
//
// The function returns the following values:
//
//    - ok: whether natural height propagation is enabled.
//
func (scrolledWindow *ScrolledWindow) PropagateNaturalHeight() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_propagate_natural_height(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PropagateNaturalWidth reports whether the natural width of the child will be
// calculated and propagated through the scrolled window’s requested natural
// width.
//
// The function returns the following values:
//
//    - ok: whether natural width propagation is enabled.
//
func (scrolledWindow *ScrolledWindow) PropagateNaturalWidth() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_propagate_natural_width(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VAdjustment returns the vertical scrollbar’s adjustment.
//
// This is the adjustment used to connect the vertical scrollbar to the child
// widget’s vertical scroll functionality.
//
// The function returns the following values:
//
//    - adjustment: vertical GtkAdjustment.
//
func (scrolledWindow *ScrolledWindow) VAdjustment() *Adjustment {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkAdjustment     // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_vadjustment(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// VScrollbar returns the vertical scrollbar of scrolled_window.
//
// The function returns the following values:
//
//    - widget: vertical scrollbar of the scrolled window.
//
func (scrolledWindow *ScrolledWindow) VScrollbar() Widgetter {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_vscrollbar(_arg0)
	runtime.KeepAlive(scrolledWindow)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SetChild sets the child widget of scrolled_window.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (scrolledWindow *ScrolledWindow) SetChild(child Widgetter) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_scrolled_window_set_child(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(child)
}

// SetHAdjustment sets the GtkAdjustment for the horizontal scrollbar.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): GtkAdjustment to use, or NULL to create a new
//      one.
//
func (scrolledWindow *ScrolledWindow) SetHAdjustment(hadjustment *Adjustment) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkAdjustment     // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if hadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	}

	C.gtk_scrolled_window_set_hadjustment(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(hadjustment)
}

// SetHasFrame changes the frame drawn around the contents of scrolled_window.
//
// The function takes the following parameters:
//
//    - hasFrame: whether to draw a frame around scrolled window contents.
//
func (scrolledWindow *ScrolledWindow) SetHasFrame(hasFrame bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if hasFrame {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_has_frame(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(hasFrame)
}

// SetKineticScrolling turns kinetic scrolling on or off.
//
// Kinetic scrolling only applies to devices with source GDK_SOURCE_TOUCHSCREEN.
//
// The function takes the following parameters:
//
//    - kineticScrolling: TRUE to enable kinetic scrolling.
//
func (scrolledWindow *ScrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if kineticScrolling {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_kinetic_scrolling(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(kineticScrolling)
}

// SetMaxContentHeight sets the maximum height that scrolled_window should keep
// visible.
//
// The scrolled_window will grow up to this height before it starts scrolling
// the content.
//
// It is a programming error to set the maximum content height to a value
// smaller than gtk.ScrolledWindow:min-content-height.
//
// The function takes the following parameters:
//
//    - height: maximum content height.
//
func (scrolledWindow *ScrolledWindow) SetMaxContentHeight(height int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.int                // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.int(height)

	C.gtk_scrolled_window_set_max_content_height(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(height)
}

// SetMaxContentWidth sets the maximum width that scrolled_window should keep
// visible.
//
// The scrolled_window will grow up to this width before it starts scrolling the
// content.
//
// It is a programming error to set the maximum content width to a value smaller
// than gtk.ScrolledWindow:min-content-width.
//
// The function takes the following parameters:
//
//    - width: maximum content width.
//
func (scrolledWindow *ScrolledWindow) SetMaxContentWidth(width int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.int                // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.int(width)

	C.gtk_scrolled_window_set_max_content_width(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(width)
}

// SetMinContentHeight sets the minimum height that scrolled_window should keep
// visible.
//
// Note that this can and (usually will) be smaller than the minimum size of the
// content.
//
// It is a programming error to set the minimum content height to a value
// greater than gtk.ScrolledWindow:max-content-height.
//
// The function takes the following parameters:
//
//    - height: minimal content height.
//
func (scrolledWindow *ScrolledWindow) SetMinContentHeight(height int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.int                // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.int(height)

	C.gtk_scrolled_window_set_min_content_height(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(height)
}

// SetMinContentWidth sets the minimum width that scrolled_window should keep
// visible.
//
// Note that this can and (usually will) be smaller than the minimum size of the
// content.
//
// It is a programming error to set the minimum content width to a value greater
// than gtk.ScrolledWindow:max-content-width.
//
// The function takes the following parameters:
//
//    - width: minimal content width.
//
func (scrolledWindow *ScrolledWindow) SetMinContentWidth(width int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.int                // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.int(width)

	C.gtk_scrolled_window_set_min_content_width(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(width)
}

// SetOverlayScrolling enables or disables overlay scrolling for this scrolled
// window.
//
// The function takes the following parameters:
//
//    - overlayScrolling: whether to enable overlay scrolling.
//
func (scrolledWindow *ScrolledWindow) SetOverlayScrolling(overlayScrolling bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if overlayScrolling {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_overlay_scrolling(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(overlayScrolling)
}

// SetPlacement sets the placement of the contents with respect to the
// scrollbars for the scrolled window.
//
// The default is GTK_CORNER_TOP_LEFT, meaning the child is in the top left,
// with the scrollbars underneath and to the right. Other values in
// gtk.CornerType are GTK_CORNER_TOP_RIGHT, GTK_CORNER_BOTTOM_LEFT, and
// GTK_CORNER_BOTTOM_RIGHT.
//
// See also gtk.ScrolledWindow.GetPlacement() and
// gtk.ScrolledWindow.UnsetPlacement().
//
// The function takes the following parameters:
//
//    - windowPlacement: position of the child window.
//
func (scrolledWindow *ScrolledWindow) SetPlacement(windowPlacement CornerType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkCornerType      // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.GtkCornerType(windowPlacement)

	C.gtk_scrolled_window_set_placement(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(windowPlacement)
}

// SetPolicy sets the scrollbar policy for the horizontal and vertical
// scrollbars.
//
// The policy determines when the scrollbar should appear; it is a value from
// the gtk.PolicyType enumeration. If GTK_POLICY_ALWAYS, the scrollbar is always
// present; if GTK_POLICY_NEVER, the scrollbar is never present; if
// GTK_POLICY_AUTOMATIC, the scrollbar is present only if needed (that is, if
// the slider part of the bar would be smaller than the trough — the display is
// larger than the page size).
//
// The function takes the following parameters:
//
//    - hscrollbarPolicy: policy for horizontal bar.
//    - vscrollbarPolicy: policy for vertical bar.
//
func (scrolledWindow *ScrolledWindow) SetPolicy(hscrollbarPolicy, vscrollbarPolicy PolicyType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkPolicyType      // out
	var _arg2 C.GtkPolicyType      // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.GtkPolicyType(hscrollbarPolicy)
	_arg2 = C.GtkPolicyType(vscrollbarPolicy)

	C.gtk_scrolled_window_set_policy(_arg0, _arg1, _arg2)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(hscrollbarPolicy)
	runtime.KeepAlive(vscrollbarPolicy)
}

// SetPropagateNaturalHeight sets whether the natural height of the child should
// be calculated and propagated through the scrolled window’s requested natural
// height.
//
// The function takes the following parameters:
//
//    - propagate: whether to propagate natural height.
//
func (scrolledWindow *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if propagate {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_propagate_natural_height(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(propagate)
}

// SetPropagateNaturalWidth sets whether the natural width of the child should
// be calculated and propagated through the scrolled window’s requested natural
// width.
//
// The function takes the following parameters:
//
//    - propagate: whether to propagate natural width.
//
func (scrolledWindow *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if propagate {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_propagate_natural_width(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(propagate)
}

// SetVAdjustment sets the GtkAdjustment for the vertical scrollbar.
//
// The function takes the following parameters:
//
//    - vadjustment (optional): GtkAdjustment to use, or NULL to create a new
//      one.
//
func (scrolledWindow *ScrolledWindow) SetVAdjustment(vadjustment *Adjustment) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkAdjustment     // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if vadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))
	}

	C.gtk_scrolled_window_set_vadjustment(_arg0, _arg1)
	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(vadjustment)
}

// UnsetPlacement unsets the placement of the contents with respect to the
// scrollbars.
//
// If no window placement is set for a scrolled window, it defaults to
// GTK_CORNER_TOP_LEFT.
func (scrolledWindow *ScrolledWindow) UnsetPlacement() {
	var _arg0 *C.GtkScrolledWindow // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	C.gtk_scrolled_window_unset_placement(_arg0)
	runtime.KeepAlive(scrolledWindow)
}
