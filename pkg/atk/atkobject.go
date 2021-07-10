// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.atk_layer_get_type()), F: marshalLayer},
		{T: externglib.Type(C.atk_role_get_type()), F: marshalRole},
		{T: externglib.Type(C.atk_implementor_get_type()), F: marshalImplementorIfacer},
		{T: externglib.Type(C.atk_object_get_type()), F: marshalObjectClasser},
	})
}

// AttributeSet: this is a singly-linked list (a List) of Attribute. It is used
// by atk_text_get_run_attributes(), atk_text_get_default_attributes(),
// atk_editable_text_set_run_attributes(), atk_document_get_attributes() and
// atk_object_get_attributes()
type AttributeSet = externglib.SList

// Layer describes the layer of a component
//
// These enumerated "layer values" are used when determining which UI rendering
// layer a component is drawn into, which can help in making determinations of
// when components occlude one another.
type Layer int

const (
	// Invalid: the object does not have a layer
	LayerInvalid Layer = iota
	// Background: this layer is reserved for the desktop background
	LayerBackground
	// Canvas: this layer is used for Canvas components
	LayerCanvas
	// Widget: this layer is normally used for components
	LayerWidget
	// MDI: this layer is used for layered components
	LayerMDI
	// Popup: this layer is used for popup components, such as menus
	LayerPopup
	// Overlay: this layer is reserved for future use.
	LayerOverlay
	// Window: this layer is used for toplevel windows.
	LayerWindow
)

func marshalLayer(p uintptr) (interface{}, error) {
	return Layer(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Role describes the role of an object
//
// These are the built-in enumerated roles that UI components can have in ATK.
// Other roles may be added at runtime, so an AtkRole >= ATK_ROLE_LAST_DEFINED
// is not necessarily an error.
type Role int

const (
	// Invalid: invalid role
	RoleInvalid Role = iota
	// AcceleratorLabel: label which represents an accelerator
	RoleAcceleratorLabel
	// Alert: object which is an alert to the user. Assistive Technologies
	// typically respond to ATK_ROLE_ALERT by reading the entire onscreen
	// contents of containers advertising this role. Should be used for warning
	// dialogs, etc.
	RoleAlert
	// Animation: object which is an animated image
	RoleAnimation
	// Arrow in one of the four cardinal directions
	RoleArrow
	// Calendar: object that displays a calendar and allows the user to select a
	// date
	RoleCalendar
	// Canvas: object that can be drawn into and is used to trap events
	RoleCanvas
	// CheckBox: choice that can be checked or unchecked and provides a separate
	// indicator for the current state
	RoleCheckBox
	// CheckMenuItem: menu item with a check box
	RoleCheckMenuItem
	// ColorChooser: specialized dialog that lets the user choose a color
	RoleColorChooser
	// ColumnHeader: the header for a column of data
	RoleColumnHeader
	// ComboBox: collapsible list of choices the user can select from
	RoleComboBox
	// DateEditor: object whose purpose is to allow a user to edit a date
	RoleDateEditor
	// DesktopIcon: inconifed internal frame within a DESKTOP_PANE
	RoleDesktopIcon
	// DesktopFrame: pane that supports internal frames and iconified versions
	// of those internal frames
	RoleDesktopFrame
	// Dial: object whose purpose is to allow a user to set a value
	RoleDial
	// Dialog: top level window with title bar and a border
	RoleDialog
	// DirectoryPane: pane that allows the user to navigate through and select
	// the contents of a directory
	RoleDirectoryPane
	// DrawingArea: object used for drawing custom user interface elements
	RoleDrawingArea
	// FileChooser: specialized dialog that lets the user choose a file
	RoleFileChooser
	// Filler: object that fills up space in a user interface
	RoleFiller
	// FontChooser: specialized dialog that lets the user choose a font
	RoleFontChooser
	// Frame: top level window with a title bar, border, menubar, etc.
	RoleFrame
	// GlassPane: pane that is guaranteed to be painted on top of all panes
	// beneath it
	RoleGlassPane
	// HtmlContainer: document container for HTML, whose children represent the
	// document content
	RoleHtmlContainer
	// Icon: small fixed size picture, typically used to decorate components
	RoleIcon
	// Image: object whose primary purpose is to display an image
	RoleImage
	// InternalFrame: frame-like object that is clipped by a desktop pane
	RoleInternalFrame
	// Label: object used to present an icon or short string in an interface
	RoleLabel
	// LayeredPane: specialized pane that allows its children to be drawn in
	// layers, providing a form of stacking order
	RoleLayeredPane
	// List: object that presents a list of objects to the user and allows the
	// user to select one or more of them
	RoleList
	// ListItem: object that represents an element of a list
	RoleListItem
	// Menu: object usually found inside a menu bar that contains a list of
	// actions the user can choose from
	RoleMenu
	// MenuBar: object usually drawn at the top of the primary dialog box of an
	// application that contains a list of menus the user can choose from
	RoleMenuBar
	// MenuItem: object usually contained in a menu that presents an action the
	// user can choose
	RoleMenuItem
	// OptionPane: specialized pane whose primary use is inside a DIALOG
	RoleOptionPane
	// PageTab: object that is a child of a page tab list
	RolePageTab
	// PageTabList: object that presents a series of panels (or page tabs), one
	// at a time, through some mechanism provided by the object
	RolePageTabList
	// Panel: generic container that is often used to group objects
	RolePanel
	// PasswordText: text object uses for passwords, or other places where the
	// text content is not shown visibly to the user
	RolePasswordText
	// PopupMenu: temporary window that is usually used to offer the user a list
	// of choices, and then hides when the user selects one of those choices
	RolePopupMenu
	// ProgressBar: object used to indicate how much of a task has been
	// completed
	RoleProgressBar
	// PushButton: object the user can manipulate to tell the application to do
	// something
	RolePushButton
	// RadioButton: specialized check box that will cause other radio buttons in
	// the same group to become unchecked when this one is checked
	RoleRadioButton
	// RadioMenuItem: check menu item which belongs to a group. At each instant
	// exactly one of the radio menu items from a group is selected
	RoleRadioMenuItem
	// RootPane: specialized pane that has a glass pane and a layered pane as
	// its children
	RoleRootPane
	// RowHeader: the header for a row of data
	RoleRowHeader
	// ScrollBar: object usually used to allow a user to incrementally view a
	// large amount of data.
	RoleScrollBar
	// ScrollPane: object that allows a user to incrementally view a large
	// amount of information
	RoleScrollPane
	// Separator: object usually contained in a menu to provide a visible and
	// logical separation of the contents in a menu
	RoleSeparator
	// Slider: object that allows the user to select from a bounded range
	RoleSlider
	// SplitPane: specialized panel that presents two other panels at the same
	// time
	RoleSplitPane
	// SpinButton: object used to get an integer or floating point number from
	// the user
	RoleSpinButton
	// Statusbar: object which reports messages of minor importance to the user
	RoleStatusbar
	// Table: object used to represent information in terms of rows and columns
	RoleTable
	// TableCell: cell in a table
	RoleTableCell
	// TableColumnHeader: the header for a column of a table
	RoleTableColumnHeader
	// TableRowHeader: the header for a row of a table
	RoleTableRowHeader
	// TearOffMenuItem: menu item used to tear off and reattach its menu
	RoleTearOffMenuItem
	// Terminal: object that represents an accessible terminal. (Since: 0.6)
	RoleTerminal
	// Text: interactive widget that supports multiple lines of text and
	// optionally accepts user input, but whose purpose is not to solicit user
	// input. Thus ATK_ROLE_TEXT is appropriate for the text view in a plain
	// text editor but inappropriate for an input field in a dialog box or web
	// form. For widgets whose purpose is to solicit input from the user, see
	// ATK_ROLE_ENTRY and ATK_ROLE_PASSWORD_TEXT. For generic objects which
	// display a brief amount of textual information, see ATK_ROLE_STATIC.
	RoleText
	// ToggleButton: specialized push button that can be checked or unchecked,
	// but does not provide a separate indicator for the current state
	RoleToggleButton
	// ToolBar: bar or palette usually composed of push buttons or toggle
	// buttons
	RoleToolBar
	// ToolTip: object that provides information about another object
	RoleToolTip
	// Tree: object used to represent hierarchical information to the user
	RoleTree
	// TreeTable: object capable of expanding and collapsing rows as well as
	// showing multiple columns of data. (Since: 0.7)
	RoleTreeTable
	// Unknown: the object contains some Accessible information, but its role is
	// not known
	RoleUnknown
	// Viewport: object usually used in a scroll pane
	RoleViewport
	// Window: top level window with no title or border.
	RoleWindow
	// Header: object that serves as a document header. (Since: 1.1.1)
	RoleHeader
	// Footer: object that serves as a document footer. (Since: 1.1.1)
	RoleFooter
	// Paragraph: object which is contains a paragraph of text content. (Since:
	// 1.1.1)
	RoleParagraph
	// Ruler: object which describes margins and tab stops, etc. for text
	// objects which it controls (should have CONTROLLER_FOR relation to such).
	// (Since: 1.1.1)
	RoleRuler
	// Application: the object is an application object, which may contain
	// @ATK_ROLE_FRAME objects or other types of accessibles. The root
	// accessible of any application's ATK hierarchy should have
	// ATK_ROLE_APPLICATION. (Since: 1.1.4)
	RoleApplication
	// Autocomplete: the object is a dialog or list containing items for
	// insertion into an entry widget, for instance a list of words for
	// completion of a text entry. (Since: 1.3)
	RoleAutocomplete
	// EditBar: the object is an editable text object in a toolbar. (Since: 1.5)
	RoleEditBar
	// Embedded: the object is an embedded container within a document or panel.
	// This role is a grouping "hint" indicating that the contained objects
	// share a context. (Since: 1.7.2)
	RoleEmbedded
	// Entry: the object is a component whose textual content may be entered or
	// modified by the user, provided @ATK_STATE_EDITABLE is present. (Since:
	// 1.11)
	RoleEntry
	// Chart: the object is a graphical depiction of quantitative data. It may
	// contain multiple subelements whose attributes and/or description may be
	// queried to obtain both the quantitative data and information about how
	// the data is being presented. The LABELLED_BY relation is particularly
	// important in interpreting objects of this type, as is the
	// accessible-description property. (Since: 1.11)
	RoleChart
	// Caption: the object contains descriptive information, usually textual,
	// about another user interface element such as a table, chart, or image.
	// (Since: 1.11)
	RoleCaption
	// DocumentFrame: the object is a visual frame or container which contains a
	// view of document content. Document frames may occur within another
	// Document instance, in which case the second document may be said to be
	// embedded in the containing instance. HTML frames are often
	// ROLE_DOCUMENT_FRAME. Either this object, or a singleton descendant,
	// should implement the Document interface. (Since: 1.11)
	RoleDocumentFrame
	// Heading: the object serves as a heading for content which follows it in a
	// document. The 'heading level' of the heading, if availabe, may be
	// obtained by querying the object's attributes.
	RoleHeading
	// Page: the object is a containing instance which encapsulates a page of
	// information. @ATK_ROLE_PAGE is used in documents and content which
	// support a paginated navigation model. (Since: 1.11)
	RolePage
	// Section: the object is a containing instance of document content which
	// constitutes a particular 'logical' section of the document. The type of
	// content within a section, and the nature of the section division itself,
	// may be obtained by querying the object's attributes. Sections may be
	// nested. (Since: 1.11)
	RoleSection
	// RedundantObject: the object is redundant with another object in the
	// hierarchy, and is exposed for purely technical reasons. Objects of this
	// role should normally be ignored by clients. (Since: 1.11)
	RoleRedundantObject
	// Form: the object is a container for form controls, for instance as part
	// of a web form or user-input form within a document. This role is
	// primarily a tag/convenience for clients when navigating complex
	// documents, it is not expected that ordinary GUI containers will always
	// have ATK_ROLE_FORM. (Since: 1.12.0)
	RoleForm
	// Link: the object is a hypertext anchor, i.e. a "link" in a hypertext
	// document. Such objects are distinct from 'inline' content which may also
	// use the Hypertext/Hyperlink interfaces to indicate the range/location
	// within a text object where an inline or embedded object lies. (Since:
	// 1.12.1)
	RoleLink
	// InputMethodWindow: the object is a window or similar viewport which is
	// used to allow composition or input of a 'complex character', in other
	// words it is an "input method window." (Since: 1.12.1)
	RoleInputMethodWindow
	// TableRow: row in a table. (Since: 2.1.0)
	RoleTableRow
	// TreeItem: object that represents an element of a tree. (Since: 2.1.0)
	RoleTreeItem
	// DocumentSpreadsheet: document frame which contains a spreadsheet. (Since:
	// 2.1.0)
	RoleDocumentSpreadsheet
	// DocumentPresentation: document frame which contains a presentation or
	// slide content. (Since: 2.1.0)
	RoleDocumentPresentation
	// DocumentText: document frame which contains textual content, such as
	// found in a word processing application. (Since: 2.1.0)
	RoleDocumentText
	// DocumentWeb: document frame which contains HTML or other markup suitable
	// for display in a web browser. (Since: 2.1.0)
	RoleDocumentWeb
	// DocumentEmail: document frame which contains email content to be
	// displayed or composed either in plain text or HTML. (Since: 2.1.0)
	RoleDocumentEmail
	// Comment: object found within a document and designed to present a
	// comment, note, or other annotation. In some cases, this object might not
	// be visible until activated. (Since: 2.1.0)
	RoleComment
	// ListBox: non-collapsible list of choices the user can select from.
	// (Since: 2.1.0)
	RoleListBox
	// Grouping: group of related widgets. This group typically has a label.
	// (Since: 2.1.0)
	RoleGrouping
	// ImageMap: image map object. Usually a graphic with multiple hotspots,
	// where each hotspot can be activated resulting in the loading of another
	// document or section of a document. (Since: 2.1.0)
	RoleImageMap
	// Notification: transitory object designed to present a message to the
	// user, typically at the desktop level rather than inside a particular
	// application. (Since: 2.1.0)
	RoleNotification
	// InfoBar: object designed to present a message to the user within an
	// existing window. (Since: 2.1.0)
	RoleInfoBar
	// LevelBar: bar that serves as a level indicator to, for instance, show the
	// strength of a password or the state of a battery. (Since: 2.7.3)
	RoleLevelBar
	// TitleBar: bar that serves as the title of a window or a dialog. (Since:
	// 2.12)
	RoleTitleBar
	// BlockQuote: object which contains a text section that is quoted from
	// another source. (Since: 2.12)
	RoleBlockQuote
	// Audio: object which represents an audio element. (Since: 2.12)
	RoleAudio
	// Video: object which represents a video element. (Since: 2.12)
	RoleVideo
	// Definition of a term or concept. (Since: 2.12)
	RoleDefinition
	// Article: section of a page that consists of a composition that forms an
	// independent part of a document, page, or site. Examples: A blog entry, a
	// news story, a forum post. (Since: 2.12)
	RoleArticle
	// Landmark: region of a web page intended as a navigational landmark. This
	// is designed to allow Assistive Technologies to provide quick navigation
	// among key regions within a document. (Since: 2.12)
	RoleLandmark
	// Log: text widget or container holding log content, such as chat history
	// and error logs. In this role there is a relationship between the arrival
	// of new items in the log and the reading order. The log contains a
	// meaningful sequence and new information is added only to the end of the
	// log, not at arbitrary points. (Since: 2.12)
	RoleLog
	// Marquee: container where non-essential information changes frequently.
	// Common usages of marquee include stock tickers and ad banners. The
	// primary difference between a marquee and a log is that logs usually have
	// a meaningful order or sequence of important content changes. (Since:
	// 2.12)
	RoleMarquee
	// Math: text widget or container that holds a mathematical expression.
	// (Since: 2.12)
	RoleMath
	// Rating: widget whose purpose is to display a rating, such as the number
	// of stars associated with a song in a media player. Objects of this role
	// should also implement AtkValue. (Since: 2.12)
	RoleRating
	// Timer: object containing a numerical counter which indicates an amount of
	// elapsed time from a start point, or the time remaining until an end
	// point. (Since: 2.12)
	RoleTimer
	// DescriptionList: object that represents a list of term-value groups. A
	// term-value group represents a individual description and consist of one
	// or more names (ATK_ROLE_DESCRIPTION_TERM) followed by one or more values
	// (ATK_ROLE_DESCRIPTION_VALUE). For each list, there should not be more
	// than one group with the same term name. (Since: 2.12)
	RoleDescriptionList
	// DescriptionTerm: object that represents a term or phrase with a
	// corresponding definition. (Since: 2.12)
	RoleDescriptionTerm
	// DescriptionValue: object that represents the description, definition or
	// value of a term. (Since: 2.12)
	RoleDescriptionValue
	// Static: generic non-container object whose purpose is to display a brief
	// amount of information to the user and whose role is known by the
	// implementor but lacks semantic value for the user. Examples in which
	// ATK_ROLE_STATIC is appropriate include the message displayed in a message
	// box and an image used as an alternative means to display text.
	// ATK_ROLE_STATIC should not be applied to widgets which are traditionally
	// interactive, objects which display a significant amount of content, or
	// any object which has an accessible relation pointing to another object.
	// Implementors should expose the displayed information through the
	// accessible name of the object. If doing so seems inappropriate, it may
	// indicate that a different role should be used. For labels which describe
	// another widget, see ATK_ROLE_LABEL. For text views, see ATK_ROLE_TEXT.
	// For generic containers, see ATK_ROLE_PANEL. For objects whose role is not
	// known by the implementor, see ATK_ROLE_UNKNOWN. (Since: 2.16)
	RoleStatic
	// MathFraction: object that represents a mathematical fraction. (Since:
	// 2.16)
	RoleMathFraction
	// MathRoot: object that represents a mathematical expression displayed with
	// a radical. (Since: 2.16)
	RoleMathRoot
	// Subscript: object that contains text that is displayed as a subscript.
	// (Since: 2.16)
	RoleSubscript
	// Superscript: object that contains text that is displayed as a
	// superscript. (Since: 2.16)
	RoleSuperscript
	// Footnote: object that contains the text of a footnote. (Since: 2.26)
	RoleFootnote
	// ContentDeletion: content previously deleted or proposed to be deleted,
	// e.g. in revision history or a content view providing suggestions from
	// reviewers. (Since: 2.34)
	RoleContentDeletion
	// ContentInsertion: content previously inserted or proposed to be inserted,
	// e.g. in revision history or a content view providing suggestions from
	// reviewers. (Since: 2.34)
	RoleContentInsertion
	// Mark: run of content that is marked or highlighted, such as for reference
	// purposes, or to call it out as having a special purpose. If the marked
	// content has an associated section in the document elaborating on the
	// reason for the mark, then ATK_RELATION_DETAILS should be used on the mark
	// to point to that associated section. In addition, the reciprocal relation
	// ATK_RELATION_DETAILS_FOR should be used on the associated content section
	// to point back to the mark. (Since: 2.36)
	RoleMark
	// Suggestion: container for content that is called out as a proposed change
	// from the current version of the document, such as by a reviewer of the
	// content. This role should include either ATK_ROLE_CONTENT_DELETION and/or
	// ATK_ROLE_CONTENT_INSERTION children, in any order, to indicate what the
	// actual change is. (Since: 2.36)
	RoleSuggestion
	// LastDefined: not a valid role, used for finding end of the enumeration
	RoleLastDefined
)

func marshalRole(p uintptr) (interface{}, error) {
	return Role(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Function is a function definition used for padding which has been added to
// class and interface structures to allow for expansion in the future.
type Function func(userData interface{}) (ok bool)

//export gotk4_Function
func gotk4_Function(arg0 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	var userData interface{} // out

	userData = box.Get(uintptr(arg0))

	fn := v.(Function)
	ok := fn(userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ImplementorIfacer describes ImplementorIface's methods.
type ImplementorIfacer interface {
	gextras.Objector

	privateImplementorIface()
}

// ImplementorIface: the AtkImplementor interface is implemented by objects for
// which AtkObject peers may be obtained via calls to
// iface->(ref_accessible)(implementor);
type ImplementorIface struct {
	*externglib.Object
}

var _ ImplementorIfacer = (*ImplementorIface)(nil)

func wrapImplementorIfacer(obj *externglib.Object) ImplementorIfacer {
	return &ImplementorIface{
		Object: obj,
	}
}

func marshalImplementorIfacer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapImplementorIfacer(obj), nil
}

func (*ImplementorIface) privateImplementorIface() {}

// ObjectClasserOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ObjectClasserOverrider interface {
	ActiveDescendantChanged(child interface{})
	ChildrenChanged(changeIndex uint, changedChild interface{})
	FocusEvent(focusIn bool)
	// Description gets the accessible description of the accessible.
	Description() string
	// IndexInParent gets the 0-based index of this accessible in its parent;
	// returns -1 if the accessible does not have an accessible parent.
	IndexInParent() int
	// Layer gets the layer of the accessible.
	//
	// Deprecated: Use atk_component_get_layer instead.
	Layer() Layer
	// MDIZOrder gets the zorder of the accessible. The value G_MININT will be
	// returned if the layer of the accessible is not ATK_LAYER_MDI.
	//
	// Deprecated: Use atk_component_get_mdi_zorder instead.
	MDIZOrder() int
	NChildren() int
	// Name gets the accessible name of the accessible.
	Name() string
	// ObjectLocale gets a UTF-8 string indicating the POSIX-style LC_MESSAGES
	// locale of @accessible.
	ObjectLocale() string
	// Parent gets the accessible parent of the accessible. By default this is
	// the one assigned with atk_object_set_parent(), but it is assumed that ATK
	// implementors have ways to get the parent of the object without the need
	// of assigning it manually with atk_object_set_parent(), and will return it
	// with this method.
	//
	// If you are only interested on the parent assigned with
	// atk_object_set_parent(), use atk_object_peek_parent().
	Parent() *ObjectClass
	// Role gets the role of the accessible.
	Role() Role
	// Initialize: this function is called when implementing subclasses of
	// Object. It does initialization required for the new object. It is
	// intended that this function should called only in the ..._new() functions
	// used to create an instance of a subclass of Object
	Initialize(data interface{})
	PropertyChange(values *PropertyValues)
	// RefRelationSet gets the RelationSet associated with the object.
	RefRelationSet() *RelationSet
	// RefStateSet gets a reference to the state set of the accessible; the
	// caller must unreference it when it is no longer needed.
	RefStateSet() *StateSet
	// RemovePropertyChangeHandler removes a property change handler.
	//
	// Deprecated: See atk_object_connect_property_change_handler().
	RemovePropertyChangeHandler(handlerId uint)
	// SetDescription sets the accessible description of the accessible. You
	// can't set the description to NULL. This is reserved for the initial
	// value. In this aspect NULL is similar to ATK_ROLE_UNKNOWN. If you want to
	// set the name to a empty value you can use "".
	SetDescription(description string)
	// SetName sets the accessible name of the accessible. You can't set the
	// name to NULL. This is reserved for the initial value. In this aspect NULL
	// is similar to ATK_ROLE_UNKNOWN. If you want to set the name to a empty
	// value you can use "".
	SetName(name string)
	// SetParent sets the accessible parent of the accessible. @parent can be
	// NULL.
	SetParent(parent ObjectClasser)
	StateChange(name string, stateSet bool)
	VisibleDataChanged()
}

// ObjectClasser describes ObjectClass's methods.
type ObjectClasser interface {
	gextras.Objector

	AccessibleID() string
	Description() string
	IndexInParent() int
	Layer() Layer
	MDIZOrder() int
	NAccessibleChildren() int
	Name() string
	ObjectLocale() string
	Parent() *ObjectClass
	Role() Role
	Initialize(data interface{})
	PeekParent() *ObjectClass
	RefAccessibleChild(i int) *ObjectClass
	RefRelationSet() *RelationSet
	RefStateSet() *StateSet
	RemovePropertyChangeHandler(handlerId uint)
	SetAccessibleID(name string)
	SetDescription(description string)
	SetName(name string)
	SetParent(parent ObjectClasser)
}

// Object: this class is the primary class for accessibility support via the
// Accessibility ToolKit (ATK). Objects which are instances of Object (or
// instances of AtkObject-derived types) are queried for properties which relate
// basic (and generic) properties of a UI component such as name and
// description. Instances of Object may also be queried as to whether they
// implement other ATK interfaces (e.g. Action, Component, etc.), as appropriate
// to the role which a given UI component plays in a user interface.
//
// All UI components in an application which provide useful information or
// services to the user must provide corresponding Object instances on request
// (in GTK+, for instance, usually on a call to #gtk_widget_get_accessible ()),
// either via ATK support built into the toolkit for the widget class or
// ancestor class, or in the case of custom widgets, if the inherited Object
// implementation is insufficient, via instances of a new Object subclass.
//
// See also: ObjectFactory, Registry. (GTK+ users see also Accessible).
type ObjectClass struct {
	*externglib.Object
}

var _ ObjectClasser = (*ObjectClass)(nil)

func wrapObjectClasser(obj *externglib.Object) ObjectClasser {
	return &ObjectClass{
		Object: obj,
	}
}

func marshalObjectClasser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapObjectClasser(obj), nil
}

// AccessibleID gets the accessible id of the accessible.
func (accessible *ObjectClass) AccessibleID() string {
	var _arg0 *C.AtkObject // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_accessible_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Description gets the accessible description of the accessible.
func (accessible *ObjectClass) Description() string {
	var _arg0 *C.AtkObject // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IndexInParent gets the 0-based index of this accessible in its parent;
// returns -1 if the accessible does not have an accessible parent.
func (accessible *ObjectClass) IndexInParent() int {
	var _arg0 *C.AtkObject // out
	var _cret C.gint       // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_index_in_parent(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Layer gets the layer of the accessible.
//
// Deprecated: Use atk_component_get_layer instead.
func (accessible *ObjectClass) Layer() Layer {
	var _arg0 *C.AtkObject // out
	var _cret C.AtkLayer   // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_layer(_arg0)

	var _layer Layer // out

	_layer = (Layer)(_cret)

	return _layer
}

// MDIZOrder gets the zorder of the accessible. The value G_MININT will be
// returned if the layer of the accessible is not ATK_LAYER_MDI.
//
// Deprecated: Use atk_component_get_mdi_zorder instead.
func (accessible *ObjectClass) MDIZOrder() int {
	var _arg0 *C.AtkObject // out
	var _cret C.gint       // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_mdi_zorder(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NAccessibleChildren gets the number of accessible children of the accessible.
func (accessible *ObjectClass) NAccessibleChildren() int {
	var _arg0 *C.AtkObject // out
	var _cret C.gint       // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_n_accessible_children(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Name gets the accessible name of the accessible.
func (accessible *ObjectClass) Name() string {
	var _arg0 *C.AtkObject // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ObjectLocale gets a UTF-8 string indicating the POSIX-style LC_MESSAGES
// locale of @accessible.
func (accessible *ObjectClass) ObjectLocale() string {
	var _arg0 *C.AtkObject // out
	var _cret *C.gchar     // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_object_locale(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Parent gets the accessible parent of the accessible. By default this is the
// one assigned with atk_object_set_parent(), but it is assumed that ATK
// implementors have ways to get the parent of the object without the need of
// assigning it manually with atk_object_set_parent(), and will return it with
// this method.
//
// If you are only interested on the parent assigned with
// atk_object_set_parent(), use atk_object_peek_parent().
func (accessible *ObjectClass) Parent() *ObjectClass {
	var _arg0 *C.AtkObject // out
	var _cret *C.AtkObject // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_parent(_arg0)

	var _object *ObjectClass // out

	_object = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ObjectClass)

	return _object
}

// Role gets the role of the accessible.
func (accessible *ObjectClass) Role() Role {
	var _arg0 *C.AtkObject // out
	var _cret C.AtkRole    // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_get_role(_arg0)

	var _role Role // out

	_role = (Role)(_cret)

	return _role
}

// Initialize: this function is called when implementing subclasses of Object.
// It does initialization required for the new object. It is intended that this
// function should called only in the ..._new() functions used to create an
// instance of a subclass of Object
func (accessible *ObjectClass) Initialize(data interface{}) {
	var _arg0 *C.AtkObject // out
	var _arg1 C.gpointer   // out

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))
	_arg1 = (C.gpointer)(box.Assign(data))

	C.atk_object_initialize(_arg0, _arg1)
}

// PeekParent gets the accessible parent of the accessible, if it has been
// manually assigned with atk_object_set_parent. Otherwise, this function
// returns nil.
//
// This method is intended as an utility for ATK implementors, and not to be
// exposed to accessible tools. See atk_object_get_parent() for further
// reference.
func (accessible *ObjectClass) PeekParent() *ObjectClass {
	var _arg0 *C.AtkObject // out
	var _cret *C.AtkObject // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_peek_parent(_arg0)

	var _object *ObjectClass // out

	_object = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ObjectClass)

	return _object
}

// RefAccessibleChild gets a reference to the specified accessible child of the
// object. The accessible children are 0-based so the first accessible child is
// at index 0, the second at index 1 and so on.
func (accessible *ObjectClass) RefAccessibleChild(i int) *ObjectClass {
	var _arg0 *C.AtkObject // out
	var _arg1 C.gint       // out
	var _cret *C.AtkObject // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_object_ref_accessible_child(_arg0, _arg1)

	var _object *ObjectClass // out

	_object = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ObjectClass)

	return _object
}

// RefRelationSet gets the RelationSet associated with the object.
func (accessible *ObjectClass) RefRelationSet() *RelationSet {
	var _arg0 *C.AtkObject      // out
	var _cret *C.AtkRelationSet // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_ref_relation_set(_arg0)

	var _relationSet *RelationSet // out

	_relationSet = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*RelationSet)

	return _relationSet
}

// RefStateSet gets a reference to the state set of the accessible; the caller
// must unreference it when it is no longer needed.
func (accessible *ObjectClass) RefStateSet() *StateSet {
	var _arg0 *C.AtkObject   // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))

	_cret = C.atk_object_ref_state_set(_arg0)

	var _stateSet *StateSet // out

	_stateSet = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*StateSet)

	return _stateSet
}

// RemovePropertyChangeHandler removes a property change handler.
//
// Deprecated: See atk_object_connect_property_change_handler().
func (accessible *ObjectClass) RemovePropertyChangeHandler(handlerId uint) {
	var _arg0 *C.AtkObject // out
	var _arg1 C.guint      // out

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))
	_arg1 = C.guint(handlerId)

	C.atk_object_remove_property_change_handler(_arg0, _arg1)
}

// SetAccessibleID sets the accessible ID of the accessible. This is not meant
// to be presented to the user, but to be an ID which is stable over application
// development. Typically, this is the gtkbuilder ID. Such an ID will be
// available for instance to identify a given well-known accessible object for
// tailored screen reading, or for automatic regression testing.
func (accessible *ObjectClass) SetAccessibleID(name string) {
	var _arg0 *C.AtkObject // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.atk_object_set_accessible_id(_arg0, _arg1)
}

// SetDescription sets the accessible description of the accessible. You can't
// set the description to NULL. This is reserved for the initial value. In this
// aspect NULL is similar to ATK_ROLE_UNKNOWN. If you want to set the name to a
// empty value you can use "".
func (accessible *ObjectClass) SetDescription(description string) {
	var _arg0 *C.AtkObject // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))
	_arg1 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(_arg1))

	C.atk_object_set_description(_arg0, _arg1)
}

// SetName sets the accessible name of the accessible. You can't set the name to
// NULL. This is reserved for the initial value. In this aspect NULL is similar
// to ATK_ROLE_UNKNOWN. If you want to set the name to a empty value you can use
// "".
func (accessible *ObjectClass) SetName(name string) {
	var _arg0 *C.AtkObject // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.atk_object_set_name(_arg0, _arg1)
}

// SetParent sets the accessible parent of the accessible. @parent can be NULL.
func (accessible *ObjectClass) SetParent(parent ObjectClasser) {
	var _arg0 *C.AtkObject // out
	var _arg1 *C.AtkObject // out

	_arg0 = (*C.AtkObject)(unsafe.Pointer(accessible.Native()))
	_arg1 = (*C.AtkObject)(unsafe.Pointer(parent.Native()))

	C.atk_object_set_parent(_arg0, _arg1)
}

// Attribute is a string name/value pair representing a generic attribute. This
// can be used to expose additional information from an accessible object as a
// whole (see atk_object_get_attributes()) or an document (see
// atk_document_get_attributes()). In the case of text attributes (see
// atk_text_get_default_attributes()), TextAttribute enum defines all the
// possible text attribute names. You can use atk_text_attribute_get_name() to
// get the string name from the enum value. See also
// atk_text_attribute_for_name() and atk_text_attribute_get_value() for more
// information.
//
// A string name/value pair representing a generic attribute.
type Attribute struct {
	native C.AtkAttribute
}

// Native returns the underlying C source pointer.
func (a *Attribute) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// PropertyValues: note: @old_value field of PropertyValues will not contain a
// valid value. This is a field defined with the purpose of contain the previous
// value of the property, but is not used anymore.
type PropertyValues struct {
	native C.AtkPropertyValues
}

// Native returns the underlying C source pointer.
func (p *PropertyValues) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}
