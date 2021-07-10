// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_entry_icon_position_get_type()), F: marshalEntryIconPosition},
		{T: externglib.Type(C.gtk_entry_get_type()), F: marshalEntry},
	})
}

// EntryIconPosition specifies the side of the entry at which an icon is placed.
type EntryIconPosition int

const (
	// Primary: at the beginning of the entry (depending on the text direction).
	EntryIconPositionPrimary EntryIconPosition = iota
	// Secondary: at the end of the entry (depending on the text direction).
	EntryIconPositionSecondary
)

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	return EntryIconPosition(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EntryOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type EntryOverrider interface {
	Activate()
}

// Entry: `GtkEntry` is a single line text entry widget.
//
// !An example GtkEntry (entry.png)
//
// A fairly large set of key bindings are supported by default. If the entered
// text is longer than the allocation of the widget, the widget will scroll so
// that the cursor position is visible.
//
// When using an entry for passwords and other sensitive information, it can be
// put into “password mode” using [method@Gtk.Entry.set_visibility]. In this
// mode, entered text is displayed using a “invisible” character. By default,
// GTK picks the best invisible character that is available in the current font,
// but it can be changed with [method@Gtk.Entry.set_invisible_char].
//
// `GtkEntry` has the ability to display progress or activity information behind
// the text. To make an entry display such information, use
// [method@Gtk.Entry.set_progress_fraction] or
// [method@Gtk.Entry.set_progress_pulse_step].
//
// Additionally, `GtkEntry` can show icons at either side of the entry. These
// icons can be activatable by clicking, can be set up as drag source and can
// have tooltips. To add an icon, use [method@Gtk.Entry.set_icon_from_gicon] or
// one of the various other functions that set an icon from an icon name or a
// paintable. To trigger an action when the user clicks an icon, connect to the
// [signal@Gtk.Entry::icon-press] signal. To allow DND operations from an icon,
// use [method@Gtk.Entry.set_icon_drag_source]. To set a tooltip on an icon, use
// [method@Gtk.Entry.set_icon_tooltip_text] or the corresponding function for
// markup.
//
// Note that functionality or information that is only available by clicking on
// an icon in an entry may not be accessible at all to users which are not able
// to use a mouse or other pointing device. It is therefore recommended that any
// such functionality should also be available by other means, e.g. via the
// context menu of the entry.
//
//
// CSS nodes
//
// “` entry[.flat][.warning][.error] ├── text[.readonly] ├── image.left ├──
// image.right ╰── [progress[.pulse]] “`
//
// `GtkEntry` has a main node with the name entry. Depending on the properties
// of the entry, the style classes .read-only and .flat may appear. The style
// classes .warning and .error may also be used with entries.
//
// When the entry shows icons, it adds subnodes with the name image and the
// style class .left or .right, depending on where the icon appears.
//
// When the entry shows progress, it adds a subnode with the name progress. The
// node has the style class .pulse when the shown progress is pulsing.
//
// For all the subnodes added to the text node in various situations, see
// [class@Gtk.Text].
//
//
// GtkEntry as GtkBuildable
//
// The `GtkEntry` implementation of the `GtkBuildable` interface supports a
// custom <attributes> element, which supports any number of <attribute>
// elements. The <attribute> element has attributes named “name“, “value“,
// “start“ and “end“ and allows you to specify Attribute values for this label.
//
// An example of a UI definition fragment specifying Pango attributes: “`xml
// <object class="GtkEnry"> <attributes> <attribute name="weight"
// value="PANGO_WEIGHT_BOLD"/> <attribute name="background" value="red"
// start="5" end="10"/> </attributes> </object> “`
//
// The start and end attributes specify the range of characters to which the
// Pango attribute applies. If start and end are not specified, the attribute is
// applied to the whole text. Note that specifying ranges does not make much
// sense with translatable attributes. Use markup embedded in the translatable
// content instead.
//
//
// Accessibility
//
// `GtkEntry` uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type Entry interface {
	gextras.Objector

	// ActivatesDefault retrieves the value set by
	// gtk_entry_set_activates_default().
	ActivatesDefault() bool
	// Alignment gets the value set by gtk_entry_set_alignment().
	//
	// See also: [property@Gtk.Editable:xalign]
	Alignment() float32
	// Attributes gets the attribute list of the `GtkEntry`.
	//
	// See [method@Gtk.Entry.set_attributes].
	Attributes() *pango.AttrList
	// Buffer: get the `GtkEntryBuffer` object which holds the text for this
	// widget.
	Buffer() *EntryBufferClass
	// Completion returns the auxiliary completion object currently in use by
	// @entry.
	Completion() *EntryCompletionClass
	// CurrentIconDragSource returns the index of the icon which is the source
	// of the current DND operation, or -1.
	CurrentIconDragSource() int
	// ExtraMenu gets the menu model set with gtk_entry_set_extra_menu().
	ExtraMenu() *gio.MenuModelClass
	// HasFrame gets the value set by gtk_entry_set_has_frame().
	HasFrame() bool
	// IconAtPos finds the icon at the given position and return its index.
	//
	// The position’s coordinates are relative to the @entry’s top left corner.
	// If @x, @y doesn’t lie inside an icon, -1 is returned. This function is
	// intended for use in a [signal@Gtk.Widget::query-tooltip] signal handler.
	IconAtPos(x int, y int) int
	// InputHints gets the input hints of this `GtkEntry`.
	InputHints() InputHints
	// InputPurpose gets the input purpose of the `GtkEntry`.
	InputPurpose() InputPurpose
	// InvisibleChar retrieves the character displayed in place of the actual
	// text in “password mode”.
	InvisibleChar() uint32
	// MaxLength retrieves the maximum allowed length of the text in @entry.
	//
	// See [method@Gtk.Entry.set_max_length].
	MaxLength() int
	// OverwriteMode gets whether the `GtkEntry` is in overwrite mode.
	OverwriteMode() bool
	// PlaceholderText retrieves the text that will be displayed when @entry is
	// empty and unfocused
	PlaceholderText() string
	// ProgressFraction returns the current fraction of the task that’s been
	// completed.
	//
	// See [method@Gtk.Entry.set_progress_fraction].
	ProgressFraction() float64
	// ProgressPulseStep retrieves the pulse step set with
	// gtk_entry_set_progress_pulse_step().
	ProgressPulseStep() float64
	// Tabs gets the tabstops of the `GtkEntry.
	//
	// See [method@Gtk.Entry.set_tabs].
	Tabs() *pango.TabArray
	// TextLength retrieves the current length of the text in @entry.
	//
	// This is equivalent to getting @entry's `GtkEntryBuffer` and calling
	// [method@Gtk.EntryBuffer.get_length] on it.
	TextLength() uint16
	// Visibility retrieves whether the text in @entry is visible.
	//
	// See [method@Gtk.Entry.set_visibility].
	Visibility() bool
	// GrabFocusWithoutSelecting causes @entry to have keyboard focus.
	//
	// It behaves like [method@Gtk.Widget.grab_focus], except that it doesn't
	// select the contents of the entry. You only want to call this on some
	// special entries which the user usually doesn't want to replace all text
	// in, such as search-as-you-type entries.
	GrabFocusWithoutSelecting() bool
	// ProgressPulse indicates that some progress is made, but you don’t know
	// how much.
	//
	// Causes the entry’s progress indicator to enter “activity mode”, where a
	// block bounces back and forth. Each call to gtk_entry_progress_pulse()
	// causes the block to move by a little bit (the amount of movement per
	// pulse is determined by [method@Gtk.Entry.set_progress_pulse_step]).
	ProgressPulse()
	// ResetImContext: reset the input method context of the entry if needed.
	//
	// This can be necessary in the case where modifying the buffer would
	// confuse on-going input method behavior.
	ResetImContext()
	// SetActivatesDefault sets whether pressing Enter in the @entry will
	// activate the default widget for the window containing the entry.
	//
	// This usually means that the dialog containing the entry will be closed,
	// since the default widget is usually one of the dialog buttons.
	SetActivatesDefault(setting bool)
	// SetAlignment sets the alignment for the contents of the entry.
	//
	// This controls the horizontal positioning of the contents when the
	// displayed text is shorter than the width of the entry.
	//
	// See also: [property@Gtk.Editable:xalign]
	SetAlignment(xalign float32)
	// SetAttributes sets a `PangoAttrList`.
	//
	// The attributes in the list are applied to the entry text.
	//
	// Since the attributes will be applies to text that changes as the user
	// types, it makes most sense to use attributes with unlimited extent.
	SetAttributes(attrs *pango.AttrList)
	// SetBuffer: set the `GtkEntryBuffer` object which holds the text for this
	// widget.
	SetBuffer(buffer EntryBuffer)
	// SetCompletion sets @completion to be the auxiliary completion object to
	// use with @entry.
	//
	// All further configuration of the completion mechanism is done on
	// @completion using the `GtkEntryCompletion` API. Completion is disabled if
	// @completion is set to nil.
	SetCompletion(completion EntryCompletion)
	// SetExtraMenu sets a menu model to add when constructing the context menu
	// for @entry.
	SetExtraMenu(model gio.MenuModel)
	// SetHasFrame sets whether the entry has a beveled frame around it.
	SetHasFrame(setting bool)
	// SetInvisibleChar sets the character to use in place of the actual text in
	// “password mode”.
	//
	// See [method@Gtk.Entry.set_visibility] for how to enable “password mode”.
	//
	// By default, GTK picks the best invisible char available in the current
	// font. If you set the invisible char to 0, then the user will get no
	// feedback at all; there will be no text on the screen as they type.
	SetInvisibleChar(ch uint32)
	// SetMaxLength sets the maximum allowed length of the contents of the
	// widget.
	//
	// If the current contents are longer than the given length, then they will
	// be truncated to fit.
	//
	// This is equivalent to getting @entry's `GtkEntryBuffer` and calling
	// [method@Gtk.EntryBuffer.set_max_length] on it.
	SetMaxLength(max int)
	// SetOverwriteMode sets whether the text is overwritten when typing in the
	// `GtkEntry`.
	SetOverwriteMode(overwrite bool)
	// SetPlaceholderText sets text to be displayed in @entry when it is empty.
	//
	// This can be used to give a visual hint of the expected contents of the
	// `GtkEntry`.
	SetPlaceholderText(text string)
	// SetProgressFraction causes the entry’s progress indicator to “fill in”
	// the given fraction of the bar.
	//
	// The fraction should be between 0.0 and 1.0, inclusive.
	SetProgressFraction(fraction float64)
	// SetProgressPulseStep sets the fraction of total entry width to move the
	// progress bouncing block for each pulse.
	//
	// Use [method@Gtk.Entry.progress_pulse] to pulse the progress.
	SetProgressPulseStep(fraction float64)
	// SetTabs sets a `PangoTabArray`.
	//
	// The tabstops in the array are applied to the entry text.
	SetTabs(tabs *pango.TabArray)
	// SetVisibility sets whether the contents of the entry are visible or not.
	//
	// When visibility is set to false, characters are displayed as the
	// invisible char, and will also appear that way when the text in the entry
	// widget is copied elsewhere.
	//
	// By default, GTK picks the best invisible character available in the
	// current font, but it can be changed with
	// [method@Gtk.Entry.set_invisible_char].
	//
	// Note that you probably want to set [property@Gtk.Entry:input-purpose] to
	// GTK_INPUT_PURPOSE_PASSWORD or GTK_INPUT_PURPOSE_PIN to inform input
	// methods about the purpose of this entry, in addition to setting
	// visibility to false.
	SetVisibility(visible bool)
	// UnsetInvisibleChar unsets the invisible char, so that the default
	// invisible char is used again. See [method@Gtk.Entry.set_invisible_char].
	UnsetInvisibleChar()
}

// EntryClass implements the Entry interface.
type EntryClass struct {
	*externglib.Object
	WidgetClass
	AccessibleIface
	BuildableIface
	CellEditableIface
	ConstraintTargetIface
	EditableIface
}

var _ Entry = (*EntryClass)(nil)

func wrapEntry(obj *externglib.Object) Entry {
	return &EntryClass{
		Object: obj,
		WidgetClass: WidgetClass{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			AccessibleIface: AccessibleIface{
				Object: obj,
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
			ConstraintTargetIface: ConstraintTargetIface{
				Object: obj,
			},
		},
		AccessibleIface: AccessibleIface{
			Object: obj,
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
		CellEditableIface: CellEditableIface{
			Object: obj,
			WidgetClass: WidgetClass{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				AccessibleIface: AccessibleIface{
					Object: obj,
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
				ConstraintTargetIface: ConstraintTargetIface{
					Object: obj,
				},
			},
		},
		ConstraintTargetIface: ConstraintTargetIface{
			Object: obj,
		},
		EditableIface: EditableIface{
			Object: obj,
			WidgetClass: WidgetClass{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				AccessibleIface: AccessibleIface{
					Object: obj,
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
				ConstraintTargetIface: ConstraintTargetIface{
					Object: obj,
				},
			},
		},
	}
}

func marshalEntry(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEntry(obj), nil
}

// NewEntry creates a new entry.
func NewEntry() *EntryClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_entry_new()

	var _entry *EntryClass // out

	_entry = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*EntryClass)

	return _entry
}

// NewEntryWithBuffer creates a new entry with the specified text buffer.
func NewEntryWithBuffer(buffer EntryBuffer) *EntryClass {
	var _arg1 *C.GtkEntryBuffer // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_entry_new_with_buffer(_arg1)

	var _entry *EntryClass // out

	_entry = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*EntryClass)

	return _entry
}

// ActivatesDefault retrieves the value set by
// gtk_entry_set_activates_default().
func (entry *EntryClass) ActivatesDefault() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_activates_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Alignment gets the value set by gtk_entry_set_alignment().
//
// See also: [property@Gtk.Editable:xalign]
func (entry *EntryClass) Alignment() float32 {
	var _arg0 *C.GtkEntry // out
	var _cret C.float     // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_alignment(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Attributes gets the attribute list of the `GtkEntry`.
//
// See [method@Gtk.Entry.set_attributes].
func (entry *EntryClass) Attributes() *pango.AttrList {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_attributes(_arg0)

	var _attrList *pango.AttrList // out

	_attrList = (*pango.AttrList)(unsafe.Pointer(_cret))
	C.pango_attr_list_ref(_cret)
	runtime.SetFinalizer(_attrList, func(v *pango.AttrList) {
		C.pango_attr_list_unref((*C.PangoAttrList)(unsafe.Pointer(v)))
	})

	return _attrList
}

// Buffer: get the `GtkEntryBuffer` object which holds the text for this widget.
func (entry *EntryClass) Buffer() *EntryBufferClass {
	var _arg0 *C.GtkEntry       // out
	var _cret *C.GtkEntryBuffer // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_buffer(_arg0)

	var _entryBuffer *EntryBufferClass // out

	_entryBuffer = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*EntryBufferClass)

	return _entryBuffer
}

// Completion returns the auxiliary completion object currently in use by
// @entry.
func (entry *EntryClass) Completion() *EntryCompletionClass {
	var _arg0 *C.GtkEntry           // out
	var _cret *C.GtkEntryCompletion // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_completion(_arg0)

	var _entryCompletion *EntryCompletionClass // out

	_entryCompletion = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*EntryCompletionClass)

	return _entryCompletion
}

// CurrentIconDragSource returns the index of the icon which is the source of
// the current DND operation, or -1.
func (entry *EntryClass) CurrentIconDragSource() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.int       // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_current_icon_drag_source(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ExtraMenu gets the menu model set with gtk_entry_set_extra_menu().
func (entry *EntryClass) ExtraMenu() *gio.MenuModelClass {
	var _arg0 *C.GtkEntry   // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_extra_menu(_arg0)

	var _menuModel *gio.MenuModelClass // out

	_menuModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.MenuModelClass)

	return _menuModel
}

// HasFrame gets the value set by gtk_entry_set_has_frame().
func (entry *EntryClass) HasFrame() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_has_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconAtPos finds the icon at the given position and return its index.
//
// The position’s coordinates are relative to the @entry’s top left corner. If
// @x, @y doesn’t lie inside an icon, -1 is returned. This function is intended
// for use in a [signal@Gtk.Widget::query-tooltip] signal handler.
func (entry *EntryClass) IconAtPos(x int, y int) int {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.int       // out
	var _arg2 C.int       // out
	var _cret C.int       // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gtk_entry_get_icon_at_pos(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InputHints gets the input hints of this `GtkEntry`.
func (entry *EntryClass) InputHints() InputHints {
	var _arg0 *C.GtkEntry     // out
	var _cret C.GtkInputHints // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_input_hints(_arg0)

	var _inputHints InputHints // out

	_inputHints = (InputHints)(_cret)

	return _inputHints
}

// InputPurpose gets the input purpose of the `GtkEntry`.
func (entry *EntryClass) InputPurpose() InputPurpose {
	var _arg0 *C.GtkEntry       // out
	var _cret C.GtkInputPurpose // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_input_purpose(_arg0)

	var _inputPurpose InputPurpose // out

	_inputPurpose = (InputPurpose)(_cret)

	return _inputPurpose
}

// InvisibleChar retrieves the character displayed in place of the actual text
// in “password mode”.
func (entry *EntryClass) InvisibleChar() uint32 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gunichar  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_invisible_char(_arg0)

	var _gunichar uint32 // out

	_gunichar = uint32(_cret)

	return _gunichar
}

// MaxLength retrieves the maximum allowed length of the text in @entry.
//
// See [method@Gtk.Entry.set_max_length].
func (entry *EntryClass) MaxLength() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.int       // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_max_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// OverwriteMode gets whether the `GtkEntry` is in overwrite mode.
func (entry *EntryClass) OverwriteMode() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_overwrite_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PlaceholderText retrieves the text that will be displayed when @entry is
// empty and unfocused
func (entry *EntryClass) PlaceholderText() string {
	var _arg0 *C.GtkEntry // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_placeholder_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ProgressFraction returns the current fraction of the task that’s been
// completed.
//
// See [method@Gtk.Entry.set_progress_fraction].
func (entry *EntryClass) ProgressFraction() float64 {
	var _arg0 *C.GtkEntry // out
	var _cret C.double    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_progress_fraction(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ProgressPulseStep retrieves the pulse step set with
// gtk_entry_set_progress_pulse_step().
func (entry *EntryClass) ProgressPulseStep() float64 {
	var _arg0 *C.GtkEntry // out
	var _cret C.double    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_progress_pulse_step(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Tabs gets the tabstops of the `GtkEntry.
//
// See [method@Gtk.Entry.set_tabs].
func (entry *EntryClass) Tabs() *pango.TabArray {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_tabs(_arg0)

	var _tabArray *pango.TabArray // out

	_tabArray = (*pango.TabArray)(unsafe.Pointer(_cret))

	return _tabArray
}

// TextLength retrieves the current length of the text in @entry.
//
// This is equivalent to getting @entry's `GtkEntryBuffer` and calling
// [method@Gtk.EntryBuffer.get_length] on it.
func (entry *EntryClass) TextLength() uint16 {
	var _arg0 *C.GtkEntry // out
	var _cret C.guint16   // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_text_length(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Visibility retrieves whether the text in @entry is visible.
//
// See [method@Gtk.Entry.set_visibility].
func (entry *EntryClass) Visibility() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_visibility(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GrabFocusWithoutSelecting causes @entry to have keyboard focus.
//
// It behaves like [method@Gtk.Widget.grab_focus], except that it doesn't select
// the contents of the entry. You only want to call this on some special entries
// which the user usually doesn't want to replace all text in, such as
// search-as-you-type entries.
func (entry *EntryClass) GrabFocusWithoutSelecting() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_grab_focus_without_selecting(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ProgressPulse indicates that some progress is made, but you don’t know how
// much.
//
// Causes the entry’s progress indicator to enter “activity mode”, where a block
// bounces back and forth. Each call to gtk_entry_progress_pulse() causes the
// block to move by a little bit (the amount of movement per pulse is determined
// by [method@Gtk.Entry.set_progress_pulse_step]).
func (entry *EntryClass) ProgressPulse() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_progress_pulse(_arg0)
}

// ResetImContext: reset the input method context of the entry if needed.
//
// This can be necessary in the case where modifying the buffer would confuse
// on-going input method behavior.
func (entry *EntryClass) ResetImContext() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_reset_im_context(_arg0)
}

// SetActivatesDefault sets whether pressing Enter in the @entry will activate
// the default widget for the window containing the entry.
//
// This usually means that the dialog containing the entry will be closed, since
// the default widget is usually one of the dialog buttons.
func (entry *EntryClass) SetActivatesDefault(setting bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_activates_default(_arg0, _arg1)
}

// SetAlignment sets the alignment for the contents of the entry.
//
// This controls the horizontal positioning of the contents when the displayed
// text is shorter than the width of the entry.
//
// See also: [property@Gtk.Editable:xalign]
func (entry *EntryClass) SetAlignment(xalign float32) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.float     // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.float(xalign)

	C.gtk_entry_set_alignment(_arg0, _arg1)
}

// SetAttributes sets a `PangoAttrList`.
//
// The attributes in the list are applied to the entry text.
//
// Since the attributes will be applies to text that changes as the user types,
// it makes most sense to use attributes with unlimited extent.
func (entry *EntryClass) SetAttributes(attrs *pango.AttrList) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs))

	C.gtk_entry_set_attributes(_arg0, _arg1)
}

// SetBuffer: set the `GtkEntryBuffer` object which holds the text for this
// widget.
func (entry *EntryClass) SetBuffer(buffer EntryBuffer) {
	var _arg0 *C.GtkEntry       // out
	var _arg1 *C.GtkEntryBuffer // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_entry_set_buffer(_arg0, _arg1)
}

// SetCompletion sets @completion to be the auxiliary completion object to use
// with @entry.
//
// All further configuration of the completion mechanism is done on @completion
// using the `GtkEntryCompletion` API. Completion is disabled if @completion is
// set to nil.
func (entry *EntryClass) SetCompletion(completion EntryCompletion) {
	var _arg0 *C.GtkEntry           // out
	var _arg1 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	C.gtk_entry_set_completion(_arg0, _arg1)
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// @entry.
func (entry *EntryClass) SetExtraMenu(model gio.MenuModel) {
	var _arg0 *C.GtkEntry   // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_entry_set_extra_menu(_arg0, _arg1)
}

// SetHasFrame sets whether the entry has a beveled frame around it.
func (entry *EntryClass) SetHasFrame(setting bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_has_frame(_arg0, _arg1)
}

// SetInvisibleChar sets the character to use in place of the actual text in
// “password mode”.
//
// See [method@Gtk.Entry.set_visibility] for how to enable “password mode”.
//
// By default, GTK picks the best invisible char available in the current font.
// If you set the invisible char to 0, then the user will get no feedback at
// all; there will be no text on the screen as they type.
func (entry *EntryClass) SetInvisibleChar(ch uint32) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gunichar  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gunichar(ch)

	C.gtk_entry_set_invisible_char(_arg0, _arg1)
}

// SetMaxLength sets the maximum allowed length of the contents of the widget.
//
// If the current contents are longer than the given length, then they will be
// truncated to fit.
//
// This is equivalent to getting @entry's `GtkEntryBuffer` and calling
// [method@Gtk.EntryBuffer.set_max_length] on it.
func (entry *EntryClass) SetMaxLength(max int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.int(max)

	C.gtk_entry_set_max_length(_arg0, _arg1)
}

// SetOverwriteMode sets whether the text is overwritten when typing in the
// `GtkEntry`.
func (entry *EntryClass) SetOverwriteMode(overwrite bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	if overwrite {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_overwrite_mode(_arg0, _arg1)
}

// SetPlaceholderText sets text to be displayed in @entry when it is empty.
//
// This can be used to give a visual hint of the expected contents of the
// `GtkEntry`.
func (entry *EntryClass) SetPlaceholderText(text string) {
	var _arg0 *C.GtkEntry // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_entry_set_placeholder_text(_arg0, _arg1)
}

// SetProgressFraction causes the entry’s progress indicator to “fill in” the
// given fraction of the bar.
//
// The fraction should be between 0.0 and 1.0, inclusive.
func (entry *EntryClass) SetProgressFraction(fraction float64) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.double    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.double(fraction)

	C.gtk_entry_set_progress_fraction(_arg0, _arg1)
}

// SetProgressPulseStep sets the fraction of total entry width to move the
// progress bouncing block for each pulse.
//
// Use [method@Gtk.Entry.progress_pulse] to pulse the progress.
func (entry *EntryClass) SetProgressPulseStep(fraction float64) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.double    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.double(fraction)

	C.gtk_entry_set_progress_pulse_step(_arg0, _arg1)
}

// SetTabs sets a `PangoTabArray`.
//
// The tabstops in the array are applied to the entry text.
func (entry *EntryClass) SetTabs(tabs *pango.TabArray) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.PangoTabArray // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.PangoTabArray)(unsafe.Pointer(tabs))

	C.gtk_entry_set_tabs(_arg0, _arg1)
}

// SetVisibility sets whether the contents of the entry are visible or not.
//
// When visibility is set to false, characters are displayed as the invisible
// char, and will also appear that way when the text in the entry widget is
// copied elsewhere.
//
// By default, GTK picks the best invisible character available in the current
// font, but it can be changed with [method@Gtk.Entry.set_invisible_char].
//
// Note that you probably want to set [property@Gtk.Entry:input-purpose] to
// GTK_INPUT_PURPOSE_PASSWORD or GTK_INPUT_PURPOSE_PIN to inform input methods
// about the purpose of this entry, in addition to setting visibility to false.
func (entry *EntryClass) SetVisibility(visible bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_visibility(_arg0, _arg1)
}

// UnsetInvisibleChar unsets the invisible char, so that the default invisible
// char is used again. See [method@Gtk.Entry.set_invisible_char].
func (entry *EntryClass) UnsetInvisibleChar() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_unset_invisible_char(_arg0)
}
