// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_text_get_type()), F: marshalText},
	})
}

// Text: the `GtkText` widget is a single-line text entry widget.
//
// `GtkText` is the common implementation of single-line text editing that is
// shared between `GtkEntry`, `GtkPasswordEntry, `GtkSpinButton` and other
// widgets. In all of these, `GtkText` is used as the delegate for the
// [iface@Gtk.Editable] implementation.
//
// A fairly large set of key bindings are supported by default. If the entered
// text is longer than the allocation of the widget, the widget will scroll so
// that the cursor position is visible.
//
// When using an entry for passwords and other sensitive information, it can be
// put into “password mode” using [method@Gtk.Text.set_visibility]. In this
// mode, entered text is displayed using a “invisible” character. By default,
// GTK picks the best invisible character that is available in the current font,
// but it can be changed with [method@Gtk.Text.set_invisible_char].
//
// If you are looking to add icons or progress display in an entry, look at
// `GtkEntry`. There other alternatives for more specialized use cases, such as
// `GtkSearchEntry`.
//
// If you need multi-line editable text, look at `GtkTextView`.
//
//
// CSS nodes
//
// “` text[.read-only] ├── placeholder ├── undershoot.left ├── undershoot.right
// ├── [selection] ├── [block-cursor] ╰── [window.popup] “`
//
// `GtkText` has a main node with the name text. Depending on the properties of
// the widget, the .read-only style class may appear.
//
// When the entry has a selection, it adds a subnode with the name selection.
//
// When the entry is in overwrite mode, it adds a subnode with the name
// block-cursor that determines how the block cursor is drawn.
//
// The CSS node for a context menu is added as a subnode below text as well.
//
// The undershoot nodes are used to draw the underflow indication when content
// is scrolled out of view. These nodes get the .left and .right style classes
// added depending on where the indication is drawn.
//
// When touch is used and touch selection handles are shown, they are using CSS
// nodes with name cursor-handle. They get the .top or .bottom style class
// depending on where they are shown in relation to the selection. If there is
// just a single handle for the text cursor, it gets the style class
// .insertion-cursor.
//
//
// Accessibility
//
// `GtkText` uses the GTK_ACCESSIBLE_ROLE_NONE role, which causes it to be
// skipped for accessibility. This is because `GtkText` is expected to be used
// as a delegate for a `GtkEditable` implementation that will be represented to
// accessibility.
type Text interface {
	Editable

	ActivatesDefault() bool

	Attributes() *pango.AttrList

	Buffer() EntryBuffer

	EnableEmojiCompletion() bool

	ExtraMenu() gio.MenuModel

	InputHints() InputHints

	InputPurpose() InputPurpose

	InvisibleChar() uint32

	MaxLength() int

	OverwriteMode() bool

	PlaceholderText() string

	PropagateTextWidth() bool

	Tabs() *pango.TabArray

	TextLength() uint16

	TruncateMultiline() bool

	Visibility() bool

	GrabFocusWithoutSelectingText() bool

	SetActivatesDefaultText(activates bool)

	SetAttributesText(attrs *pango.AttrList)

	SetBufferText(buffer EntryBuffer)

	SetEnableEmojiCompletionText(enableEmojiCompletion bool)

	SetExtraMenuText(model gio.MenuModel)

	SetInputHintsText(hints InputHints)

	SetInputPurposeText(purpose InputPurpose)

	SetInvisibleCharText(ch uint32)

	SetMaxLengthText(length int)

	SetOverwriteModeText(overwrite bool)

	SetPlaceholderTextText(text string)

	SetPropagateTextWidthText(propagateTextWidth bool)

	SetTabsText(tabs *pango.TabArray)

	SetTruncateMultilineText(truncateMultiline bool)

	SetVisibilityText(visible bool)

	UnsetInvisibleCharText()
}

// text implements the Text class.
type text struct {
	Widget
}

// WrapText wraps a GObject to the right type. It is
// primarily used internally.
func WrapText(obj *externglib.Object) Text {
	return text{
		Widget: WrapWidget(obj),
	}
}

func marshalText(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapText(obj), nil
}

func NewText() Text {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_text_new()

	var _text Text // out

	_text = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Text)

	return _text
}

func NewTextWithBuffer(buffer EntryBuffer) Text {
	var _arg1 *C.GtkEntryBuffer // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_text_new_with_buffer(_arg1)

	var _text Text // out

	_text = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Text)

	return _text
}

func (s text) ActivatesDefault() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_activates_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s text) Attributes() *pango.AttrList {
	var _arg0 *C.GtkText       // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_attributes(_arg0)

	var _attrList *pango.AttrList // out

	_attrList = (*pango.AttrList)(unsafe.Pointer(_cret))

	return _attrList
}

func (s text) Buffer() EntryBuffer {
	var _arg0 *C.GtkText        // out
	var _cret *C.GtkEntryBuffer // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_buffer(_arg0)

	var _entryBuffer EntryBuffer // out

	_entryBuffer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(EntryBuffer)

	return _entryBuffer
}

func (s text) EnableEmojiCompletion() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_enable_emoji_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s text) ExtraMenu() gio.MenuModel {
	var _arg0 *C.GtkText    // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_extra_menu(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.MenuModel)

	return _menuModel
}

func (s text) InputHints() InputHints {
	var _arg0 *C.GtkText      // out
	var _cret C.GtkInputHints // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_input_hints(_arg0)

	var _inputHints InputHints // out

	_inputHints = InputHints(_cret)

	return _inputHints
}

func (s text) InputPurpose() InputPurpose {
	var _arg0 *C.GtkText        // out
	var _cret C.GtkInputPurpose // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_input_purpose(_arg0)

	var _inputPurpose InputPurpose // out

	_inputPurpose = InputPurpose(_cret)

	return _inputPurpose
}

func (s text) InvisibleChar() uint32 {
	var _arg0 *C.GtkText // out
	var _cret C.gunichar // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_invisible_char(_arg0)

	var _gunichar uint32 // out

	_gunichar = uint32(_cret)

	return _gunichar
}

func (s text) MaxLength() int {
	var _arg0 *C.GtkText // out
	var _cret C.int      // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_max_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s text) OverwriteMode() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_overwrite_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s text) PlaceholderText() string {
	var _arg0 *C.GtkText // out
	var _cret *C.char    // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_placeholder_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s text) PropagateTextWidth() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_propagate_text_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s text) Tabs() *pango.TabArray {
	var _arg0 *C.GtkText       // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_tabs(_arg0)

	var _tabArray *pango.TabArray // out

	_tabArray = (*pango.TabArray)(unsafe.Pointer(_cret))

	return _tabArray
}

func (s text) TextLength() uint16 {
	var _arg0 *C.GtkText // out
	var _cret C.guint16  // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_text_length(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (s text) TruncateMultiline() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_truncate_multiline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s text) Visibility() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_get_visibility(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s text) GrabFocusWithoutSelectingText() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_text_grab_focus_without_selecting(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s text) SetActivatesDefaultText(activates bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	if activates {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_activates_default(_arg0, _arg1)
}

func (s text) SetAttributesText(attrs *pango.AttrList) {
	var _arg0 *C.GtkText       // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs.Native()))

	C.gtk_text_set_attributes(_arg0, _arg1)
}

func (s text) SetBufferText(buffer EntryBuffer) {
	var _arg0 *C.GtkText        // out
	var _arg1 *C.GtkEntryBuffer // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_text_set_buffer(_arg0, _arg1)
}

func (s text) SetEnableEmojiCompletionText(enableEmojiCompletion bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	if enableEmojiCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_enable_emoji_completion(_arg0, _arg1)
}

func (s text) SetExtraMenuText(model gio.MenuModel) {
	var _arg0 *C.GtkText    // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_text_set_extra_menu(_arg0, _arg1)
}

func (s text) SetInputHintsText(hints InputHints) {
	var _arg0 *C.GtkText      // out
	var _arg1 C.GtkInputHints // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkInputHints(hints)

	C.gtk_text_set_input_hints(_arg0, _arg1)
}

func (s text) SetInputPurposeText(purpose InputPurpose) {
	var _arg0 *C.GtkText        // out
	var _arg1 C.GtkInputPurpose // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkInputPurpose(purpose)

	C.gtk_text_set_input_purpose(_arg0, _arg1)
}

func (s text) SetInvisibleCharText(ch uint32) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gunichar // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = C.gunichar(ch)

	C.gtk_text_set_invisible_char(_arg0, _arg1)
}

func (s text) SetMaxLengthText(length int) {
	var _arg0 *C.GtkText // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(length)

	C.gtk_text_set_max_length(_arg0, _arg1)
}

func (s text) SetOverwriteModeText(overwrite bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	if overwrite {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_overwrite_mode(_arg0, _arg1)
}

func (s text) SetPlaceholderTextText(text string) {
	var _arg0 *C.GtkText // out
	var _arg1 *C.char    // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_text_set_placeholder_text(_arg0, _arg1)
}

func (s text) SetPropagateTextWidthText(propagateTextWidth bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	if propagateTextWidth {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_propagate_text_width(_arg0, _arg1)
}

func (s text) SetTabsText(tabs *pango.TabArray) {
	var _arg0 *C.GtkText       // out
	var _arg1 *C.PangoTabArray // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.PangoTabArray)(unsafe.Pointer(tabs.Native()))

	C.gtk_text_set_tabs(_arg0, _arg1)
}

func (s text) SetTruncateMultilineText(truncateMultiline bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	if truncateMultiline {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_truncate_multiline(_arg0, _arg1)
}

func (s text) SetVisibilityText(visible bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_visibility(_arg0, _arg1)
}

func (s text) UnsetInvisibleCharText() {
	var _arg0 *C.GtkText // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(s.Native()))

	C.gtk_text_unset_invisible_char(_arg0)
}

func (e text) DeleteSelection() {
	WrapEditable(gextras.InternObject(e)).DeleteSelection()
}

func (e text) DeleteText(startPos int, endPos int) {
	WrapEditable(gextras.InternObject(e)).DeleteText(startPos, endPos)
}

func (e text) FinishDelegate() {
	WrapEditable(gextras.InternObject(e)).FinishDelegate()
}

func (e text) Alignment() float32 {
	return WrapEditable(gextras.InternObject(e)).Alignment()
}

func (e text) Chars(startPos int, endPos int) string {
	return WrapEditable(gextras.InternObject(e)).Chars(startPos, endPos)
}

func (e text) Delegate() Editable {
	return WrapEditable(gextras.InternObject(e)).Delegate()
}

func (e text) Editable() bool {
	return WrapEditable(gextras.InternObject(e)).Editable()
}

func (e text) EnableUndo() bool {
	return WrapEditable(gextras.InternObject(e)).EnableUndo()
}

func (e text) MaxWidthChars() int {
	return WrapEditable(gextras.InternObject(e)).MaxWidthChars()
}

func (e text) Position() int {
	return WrapEditable(gextras.InternObject(e)).Position()
}

func (e text) SelectionBounds() (startPos int, endPos int, ok bool) {
	return WrapEditable(gextras.InternObject(e)).SelectionBounds()
}

func (e text) Text() string {
	return WrapEditable(gextras.InternObject(e)).Text()
}

func (e text) WidthChars() int {
	return WrapEditable(gextras.InternObject(e)).WidthChars()
}

func (e text) InitDelegate() {
	WrapEditable(gextras.InternObject(e)).InitDelegate()
}

func (e text) SelectRegion(startPos int, endPos int) {
	WrapEditable(gextras.InternObject(e)).SelectRegion(startPos, endPos)
}

func (e text) SetAlignment(xalign float32) {
	WrapEditable(gextras.InternObject(e)).SetAlignment(xalign)
}

func (e text) SetEditable(isEditable bool) {
	WrapEditable(gextras.InternObject(e)).SetEditable(isEditable)
}

func (e text) SetEnableUndo(enableUndo bool) {
	WrapEditable(gextras.InternObject(e)).SetEnableUndo(enableUndo)
}

func (e text) SetMaxWidthChars(nChars int) {
	WrapEditable(gextras.InternObject(e)).SetMaxWidthChars(nChars)
}

func (e text) SetPosition(position int) {
	WrapEditable(gextras.InternObject(e)).SetPosition(position)
}

func (e text) SetText(text string) {
	WrapEditable(gextras.InternObject(e)).SetText(text)
}

func (e text) SetWidthChars(nChars int) {
	WrapEditable(gextras.InternObject(e)).SetWidthChars(nChars)
}

func (s text) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s text) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s text) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s text) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s text) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s text) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s text) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b text) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
