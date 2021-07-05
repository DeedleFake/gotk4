// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_state_type_get_type()), F: marshalStateType},
	})
}

type State = uint64

// StateType: the possible types of states of an object
type StateType int

const (
	// invalid indicates an invalid state - probably an error condition.
	StateTypeInvalid StateType = iota
	// active indicates a window is currently the active window, or an object is
	// the active subelement within a container or table. ATK_STATE_ACTIVE
	// should not be used for objects which have ATK_STATE_FOCUSABLE or
	// ATK_STATE_SELECTABLE: Those objects should use ATK_STATE_FOCUSED and
	// ATK_STATE_SELECTED respectively. ATK_STATE_ACTIVE is a means to indicate
	// that an object which is not focusable and not selectable is the
	// currently-active item within its parent container.
	StateTypeActive
	// armed indicates that the object is 'armed', i.e. will be activated by if
	// a pointer button-release event occurs within its bounds. Buttons often
	// enter this state when a pointer click occurs within their bounds, as a
	// precursor to activation. ATK_STATE_ARMED has been deprecated since
	// ATK-2.16 and should not be used in newly-written code.
	StateTypeArmed
	// busy indicates the current object is busy, i.e. onscreen representation
	// is in the process of changing, or the object is temporarily unavailable
	// for interaction due to activity already in progress. This state may be
	// used by implementors of Document to indicate that content loading is
	// underway. It also may indicate other 'pending' conditions; clients may
	// wish to interrogate this object when the ATK_STATE_BUSY flag is removed.
	StateTypeBusy
	// checked indicates this object is currently checked, for instance a
	// checkbox is 'non-empty'.
	StateTypeChecked
	// defunct indicates that this object no longer has a valid backing widget
	// (for instance, if its peer object has been destroyed)
	StateTypeDefunct
	// editable indicates that this object can contain text, and that the user
	// can change the textual contents of this object by editing those contents
	// directly. For an object which is expected to be editable due to its type,
	// but which cannot be edited due to the application or platform preventing
	// the user from doing so, that object's StateSet should lack
	// ATK_STATE_EDITABLE and should contain ATK_STATE_READ_ONLY.
	StateTypeEditable
	// enabled indicates that this object is enabled, i.e. that it currently
	// reflects some application state. Objects that are "greyed out" may lack
	// this state, and may lack the STATE_SENSITIVE if direct user interaction
	// cannot cause them to acquire STATE_ENABLED. See also: ATK_STATE_SENSITIVE
	StateTypeEnabled
	// expandable indicates this object allows progressive disclosure of its
	// children
	StateTypeExpandable
	// expanded indicates this object its expanded - see ATK_STATE_EXPANDABLE
	// above
	StateTypeExpanded
	// focusable indicates this object can accept keyboard focus, which means
	// all events resulting from typing on the keyboard will normally be passed
	// to it when it has focus
	StateTypeFocusable
	// focused indicates this object currently has the keyboard focus
	StateTypeFocused
	// horizontal indicates the orientation of this object is horizontal; used,
	// for instance, by objects of ATK_ROLE_SCROLL_BAR. For objects where
	// vertical/horizontal orientation is especially meaningful.
	StateTypeHorizontal
	// iconified indicates this object is minimized and is represented only by
	// an icon
	StateTypeIconified
	// modal indicates something must be done with this object before the user
	// can interact with an object in a different window
	StateTypeModal
	// MultiLine indicates this (text) object can contain multiple lines of text
	StateTypeMultiLine
	// multiselectable indicates this object allows more than one of its
	// children to be selected at the same time, or in the case of text objects,
	// that the object supports non-contiguous text selections.
	StateTypeMultiselectable
	// opaque indicates this object paints every pixel within its rectangular
	// region.
	StateTypeOpaque
	// pressed indicates this object is currently pressed.
	StateTypePressed
	// resizable indicates the size of this object is not fixed
	StateTypeResizable
	// selectable indicates this object is the child of an object that allows
	// its children to be selected and that this child is one of those children
	// that can be selected
	StateTypeSelectable
	// selected indicates this object is the child of an object that allows its
	// children to be selected and that this child is one of those children that
	// has been selected
	StateTypeSelected
	// sensitive indicates this object is sensitive, e.g. to user interaction.
	// STATE_SENSITIVE usually accompanies STATE_ENABLED for user-actionable
	// controls, but may be found in the absence of STATE_ENABLED if the current
	// visible state of the control is "disconnected" from the application
	// state. In such cases, direct user interaction can often result in the
	// object gaining STATE_SENSITIVE, for instance if a user makes an explicit
	// selection using an object whose current state is ambiguous or undefined.
	// @see STATE_ENABLED, STATE_INDETERMINATE.
	StateTypeSensitive
	// showing indicates this object, the object's parent, the object's parent's
	// parent, and so on, are all 'shown' to the end-user, i.e. subject to
	// "exposure" if blocking or obscuring objects do not interpose between this
	// object and the top of the window stack.
	StateTypeShowing
	// SingleLine indicates this (text) object can contain only a single line of
	// text
	StateTypeSingleLine
	// stale indicates that the information returned for this object may no
	// longer be synchronized with the application state. This is implied if the
	// object has STATE_TRANSIENT, and can also occur towards the end of the
	// object peer's lifecycle. It can also be used to indicate that the index
	// associated with this object has changed since the user accessed the
	// object (in lieu of "index-in-parent-changed" events).
	StateTypeStale
	// transient indicates this object is transient, i.e. a snapshot which may
	// not emit events when its state changes. Data from objects with
	// ATK_STATE_TRANSIENT should not be cached, since there may be no
	// notification given when the cached data becomes obsolete.
	StateTypeTransient
	// vertical indicates the orientation of this object is vertical
	StateTypeVertical
	// visible indicates this object is visible, e.g. has been explicitly marked
	// for exposure to the user. **note**: ATK_STATE_VISIBLE is no guarantee
	// that the object is actually unobscured on the screen, only that it is
	// 'potentially' visible, barring obstruction, being scrolled or clipped out
	// of the field of view, or having an ancestor container that has not yet
	// made visible. A widget is potentially onscreen if it has both
	// ATK_STATE_VISIBLE and ATK_STATE_SHOWING. The absence of ATK_STATE_VISIBLE
	// and ATK_STATE_SHOWING is semantically equivalent to saying that an object
	// is 'hidden'. See also ATK_STATE_TRUNCATED, which applies if an object
	// with ATK_STATE_VISIBLE and ATK_STATE_SHOWING set lies within a viewport
	// which means that its contents are clipped, e.g. a truncated spreadsheet
	// cell or an image within a scrolling viewport. Mostly useful for
	// screen-review and magnification algorithms.
	StateTypeVisible
	// ManagesDescendants indicates that "active-descendant-changed" event is
	// sent when children become 'active' (i.e. are selected or navigated to
	// onscreen). Used to prevent need to enumerate all children in very large
	// containers, like tables. The presence of STATE_MANAGES_DESCENDANTS is an
	// indication to the client. that the children should not, and need not, be
	// enumerated by the client. Objects implementing this state are expected to
	// provide relevant state notifications to listening clients, for instance
	// notifications of visibility changes and activation of their contained
	// child objects, without the client having previously requested references
	// to those children.
	StateTypeManagesDescendants
	// indeterminate indicates that the value, or some other quantifiable
	// property, of this AtkObject cannot be fully determined. In the case of a
	// large data set in which the total number of items in that set is unknown
	// (e.g. 1 of 999+), implementors should expose the currently-known set size
	// (999) along with this state. In the case of a check box, this state
	// should be used to indicate that the check box is a tri-state check box
	// which is currently neither checked nor unchecked.
	StateTypeIndeterminate
	// truncated indicates that an object is truncated, e.g. a text value in a
	// speradsheet cell.
	StateTypeTruncated
	// required indicates that explicit user interaction with an object is
	// required by the user interface, e.g. a required field in a "web-form"
	// interface.
	StateTypeRequired
	// InvalidEntry indicates that the object has encountered an error condition
	// due to failure of input validation. For instance, a form control may
	// acquire this state in response to invalid or malformed user input.
	StateTypeInvalidEntry
	// SupportsAutocompletion indicates that the object in question implements
	// some form of ¨typeahead¨ or pre-selection behavior whereby entering the
	// first character of one or more sub-elements causes those elements to
	// scroll into view or become selected. Subsequent character input may
	// narrow the selection further as long as one or more sub-elements match
	// the string. This state is normally only useful and encountered on objects
	// that implement Selection. In some cases the typeahead behavior may result
	// in full or partial ¨completion¨ of the data in the input field, in which
	// case these input events may trigger text-changed events from the AtkText
	// interface. This state supplants @ATK_ROLE_AUTOCOMPLETE.
	StateTypeSupportsAutocompletion
	// SelectableText indicates that the object in question supports text
	// selection. It should only be exposed on objects which implement the Text
	// interface, in order to distinguish this state from @ATK_STATE_SELECTABLE,
	// which infers that the object in question is a selectable child of an
	// object which implements Selection. While similar, text selection and
	// subelement selection are distinct operations.
	StateTypeSelectableText
	// default indicates that the object is the "default" active component, i.e.
	// the object which is activated by an end-user press of the "Enter" or
	// "Return" key. Typically a "close" or "submit" button.
	StateTypeDefault
	// animated indicates that the object changes its appearance dynamically as
	// an inherent part of its presentation. This state may come and go if an
	// object is only temporarily animated on the way to a 'final' onscreen
	// presentation. **note**: some applications, notably content viewers, may
	// not be able to detect all kinds of animated content. Therefore the
	// absence of this state should not be taken as definitive evidence that the
	// object's visual representation is static; this state is advisory.
	StateTypeAnimated
	// visited indicates that the object (typically a hyperlink) has already
	// been 'activated', and/or its backing data has already been downloaded,
	// rendered, or otherwise "visited".
	StateTypeVisited
	// checkable indicates this object has the potential to be checked, such as
	// a checkbox or toggle-able table cell. @Since: ATK-2.12
	StateTypeCheckable
	// HasPopup indicates that the object has a popup context menu or sub-level
	// menu which may or may not be showing. This means that activation renders
	// conditional content. Note that ordinary tooltips are not considered
	// popups in this context. @Since: ATK-2.12
	StateTypeHasPopup
	// HasTooltip indicates this object has a tooltip. @Since: ATK-2.16
	StateTypeHasTooltip
	// ReadOnly indicates that a widget which is ENABLED and SENSITIVE has a
	// value which can be read, but not modified, by the user. Note that this
	// state should only be applied to widget types whose value is normally
	// directly user modifiable, such as check boxes, radio buttons, spin
	// buttons, text input fields, and combo boxes, as a means to convey that
	// the expected interaction with that widget is not possible. When the
	// expected interaction with a widget does not include modification by the
	// user, as is the case with labels and containers, ATK_STATE_READ_ONLY
	// should not be applied. See also ATK_STATE_EDITABLE. @Since: ATK-2-16
	StateTypeReadOnly
	// LastDefined: not a valid state, used for finding end of enumeration
	StateTypeLastDefined
)

func marshalStateType(p uintptr) (interface{}, error) {
	return StateType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}
