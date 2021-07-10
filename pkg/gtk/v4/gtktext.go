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
	gextras.Objector

	// ActivatesDefault retrieves the value set by
	// gtk_text_set_activates_default().
	ActivatesDefault() bool
	// Attributes gets the attribute list that was set on the `GtkText` using
	// gtk_text_set_attributes().
	Attributes() *pango.AttrList
	// Buffer: get the `GtkEntryBuffer` object which holds the text for this
	// self.
	Buffer() *EntryBufferClass
	// EnableEmojiCompletion returns whether Emoji completion is enabled for
	// this `GtkText` widget.
	EnableEmojiCompletion() bool
	// ExtraMenu gets the menu model set with gtk_text_set_extra_menu().
	ExtraMenu() *gio.MenuModelClass
	// InputHints gets the input hints of the `GtkText`.
	InputHints() InputHints
	// InputPurpose gets the input purpose of the `GtkText`.
	InputPurpose() InputPurpose
	// InvisibleChar retrieves the character displayed in place of the real
	// characters for entries with visibility set to false.
	//
	// Note that GTK does not compute this value unless it needs it, so the
	// value returned by this function is not very useful unless it has been
	// explicitly set with [method@Gtk.Text.set_invisible_char].
	InvisibleChar() uint32
	// MaxLength retrieves the maximum allowed length of the text in @self.
	//
	// See [method@Gtk.Text.set_max_length].
	//
	// This is equivalent to getting @self's `GtkEntryBuffer` and calling
	// [method@Gtk.EntryBuffer.get_max_length] on it.
	MaxLength() int
	// OverwriteMode gets the value set by gtk_text_set_overwrite_mode().
	OverwriteMode() bool
	// PlaceholderText retrieves the text that will be displayed when @self is
	// empty and unfocused
	PlaceholderText() string
	// PropagateTextWidth returns whether the `GtkText` will grow and shrink
	// with the content.
	PropagateTextWidth() bool
	// Tabs gets the tabstops that were set on the `GtkText` using
	// gtk_text_set_tabs().
	Tabs() *pango.TabArray
	// TextLength retrieves the current length of the text in @self.
	//
	// This is equivalent to getting @self's `GtkEntryBuffer` and calling
	// [method@Gtk.EntryBuffer.get_length] on it.
	TextLength() uint16
	// TruncateMultiline returns whether the `GtkText` will truncate multi-line
	// text that is pasted into the widget
	TruncateMultiline() bool
	// Visibility retrieves whether the text in @self is visible.
	Visibility() bool
	// GrabFocusWithoutSelecting causes @self to have keyboard focus.
	//
	// It behaves like [method@Gtk.Widget.grab_focus], except that it doesn't
	// select the contents of @self. You only want to call this on some special
	// entries which the user usually doesn't want to replace all text in, such
	// as search-as-you-type entries.
	GrabFocusWithoutSelecting() bool
	// SetActivatesDefault: if @activates is true, pressing Enter in the @self
	// will activate the default widget for the window containing @self.
	//
	// This usually means that the dialog containing the `GtkText` will be
	// closed, since the default widget is usually one of the dialog buttons.
	SetActivatesDefault(activates bool)
	// SetAttributes sets attributes that are applied to the text.
	SetAttributes(attrs *pango.AttrList)
	// SetBuffer: set the `GtkEntryBuffer` object which holds the text for this
	// widget.
	SetBuffer(buffer EntryBuffer)
	// SetEnableEmojiCompletion sets whether Emoji completion is enabled.
	//
	// If it is, typing ':', followed by a recognized keyword, will pop up a
	// window with suggested Emojis matching the keyword.
	SetEnableEmojiCompletion(enableEmojiCompletion bool)
	// SetExtraMenu sets a menu model to add when constructing the context menu
	// for @self.
	SetExtraMenu(model gio.MenuModel)
	// SetInvisibleChar sets the character to use in place of the actual text
	// when in “password mode”.
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
	// This is equivalent to getting @self's `GtkEntryBuffer` and calling
	// [method@Gtk.EntryBuffer.set_max_length] on it.
	SetMaxLength(length int)
	// SetOverwriteMode sets whether the text is overwritten when typing in the
	// `GtkText`.
	SetOverwriteMode(overwrite bool)
	// SetPlaceholderText sets text to be displayed in @self when it is empty.
	//
	// This can be used to give a visual hint of the expected contents of the
	// `GtkText`.
	SetPlaceholderText(text string)
	// SetPropagateTextWidth sets whether the `GtkText` should grow and shrink
	// with the content.
	SetPropagateTextWidth(propagateTextWidth bool)
	// SetTabs sets tabstops that are applied to the text.
	SetTabs(tabs *pango.TabArray)
	// SetTruncateMultiline sets whether the `GtkText` should truncate
	// multi-line text that is pasted into the widget.
	SetTruncateMultiline(truncateMultiline bool)
	// SetVisibility sets whether the contents of the `GtkText` are visible or
	// not.
	//
	// When visibility is set to false, characters are displayed as the
	// invisible char, and will also appear that way when the text in the widget
	// is copied to the clipboard.
	//
	// By default, GTK picks the best invisible character available in the
	// current font, but it can be changed with
	// [method@Gtk.Text.set_invisible_char].
	//
	// Note that you probably want to set [property@Gtk.Text:input-purpose] to
	// GTK_INPUT_PURPOSE_PASSWORD or GTK_INPUT_PURPOSE_PIN to inform input
	// methods about the purpose of this self, in addition to setting visibility
	// to false.
	SetVisibility(visible bool)
	// UnsetInvisibleChar unsets the invisible char.
	//
	// After calling this, the default invisible char is used again.
	UnsetInvisibleChar()
}

// TextClass implements the Text interface.
type TextClass struct {
	*externglib.Object
	WidgetClass
	AccessibleIface
	BuildableIface
	ConstraintTargetIface
	EditableIface
}

var _ Text = (*TextClass)(nil)

func wrapText(obj *externglib.Object) Text {
	return &TextClass{
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

func marshalText(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapText(obj), nil
}

// NewText creates a new `GtkText`.
func NewText() *TextClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_text_new()

	var _text *TextClass // out

	_text = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*TextClass)

	return _text
}

// NewTextWithBuffer creates a new `GtkText` with the specified text buffer.
func NewTextWithBuffer(buffer EntryBuffer) *TextClass {
	var _arg1 *C.GtkEntryBuffer // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_text_new_with_buffer(_arg1)

	var _text *TextClass // out

	_text = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*TextClass)

	return _text
}

// ActivatesDefault retrieves the value set by gtk_text_set_activates_default().
func (self *TextClass) ActivatesDefault() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_activates_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Attributes gets the attribute list that was set on the `GtkText` using
// gtk_text_set_attributes().
func (self *TextClass) Attributes() *pango.AttrList {
	var _arg0 *C.GtkText       // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_attributes(_arg0)

	var _attrList *pango.AttrList // out

	_attrList = (*pango.AttrList)(unsafe.Pointer(_cret))
	C.pango_attr_list_ref(_cret)
	runtime.SetFinalizer(_attrList, func(v *pango.AttrList) {
		C.pango_attr_list_unref((*C.PangoAttrList)(unsafe.Pointer(v)))
	})

	return _attrList
}

// Buffer: get the `GtkEntryBuffer` object which holds the text for this self.
func (self *TextClass) Buffer() *EntryBufferClass {
	var _arg0 *C.GtkText        // out
	var _cret *C.GtkEntryBuffer // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_buffer(_arg0)

	var _entryBuffer *EntryBufferClass // out

	_entryBuffer = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*EntryBufferClass)

	return _entryBuffer
}

// EnableEmojiCompletion returns whether Emoji completion is enabled for this
// `GtkText` widget.
func (self *TextClass) EnableEmojiCompletion() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_enable_emoji_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ExtraMenu gets the menu model set with gtk_text_set_extra_menu().
func (self *TextClass) ExtraMenu() *gio.MenuModelClass {
	var _arg0 *C.GtkText    // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_extra_menu(_arg0)

	var _menuModel *gio.MenuModelClass // out

	_menuModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.MenuModelClass)

	return _menuModel
}

// InputHints gets the input hints of the `GtkText`.
func (self *TextClass) InputHints() InputHints {
	var _arg0 *C.GtkText      // out
	var _cret C.GtkInputHints // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_input_hints(_arg0)

	var _inputHints InputHints // out

	_inputHints = (InputHints)(_cret)

	return _inputHints
}

// InputPurpose gets the input purpose of the `GtkText`.
func (self *TextClass) InputPurpose() InputPurpose {
	var _arg0 *C.GtkText        // out
	var _cret C.GtkInputPurpose // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_input_purpose(_arg0)

	var _inputPurpose InputPurpose // out

	_inputPurpose = (InputPurpose)(_cret)

	return _inputPurpose
}

// InvisibleChar retrieves the character displayed in place of the real
// characters for entries with visibility set to false.
//
// Note that GTK does not compute this value unless it needs it, so the value
// returned by this function is not very useful unless it has been explicitly
// set with [method@Gtk.Text.set_invisible_char].
func (self *TextClass) InvisibleChar() uint32 {
	var _arg0 *C.GtkText // out
	var _cret C.gunichar // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_invisible_char(_arg0)

	var _gunichar uint32 // out

	_gunichar = uint32(_cret)

	return _gunichar
}

// MaxLength retrieves the maximum allowed length of the text in @self.
//
// See [method@Gtk.Text.set_max_length].
//
// This is equivalent to getting @self's `GtkEntryBuffer` and calling
// [method@Gtk.EntryBuffer.get_max_length] on it.
func (self *TextClass) MaxLength() int {
	var _arg0 *C.GtkText // out
	var _cret C.int      // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_max_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// OverwriteMode gets the value set by gtk_text_set_overwrite_mode().
func (self *TextClass) OverwriteMode() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_overwrite_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PlaceholderText retrieves the text that will be displayed when @self is empty
// and unfocused
func (self *TextClass) PlaceholderText() string {
	var _arg0 *C.GtkText // out
	var _cret *C.char    // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_placeholder_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PropagateTextWidth returns whether the `GtkText` will grow and shrink with
// the content.
func (self *TextClass) PropagateTextWidth() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_propagate_text_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Tabs gets the tabstops that were set on the `GtkText` using
// gtk_text_set_tabs().
func (self *TextClass) Tabs() *pango.TabArray {
	var _arg0 *C.GtkText       // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_tabs(_arg0)

	var _tabArray *pango.TabArray // out

	_tabArray = (*pango.TabArray)(unsafe.Pointer(_cret))

	return _tabArray
}

// TextLength retrieves the current length of the text in @self.
//
// This is equivalent to getting @self's `GtkEntryBuffer` and calling
// [method@Gtk.EntryBuffer.get_length] on it.
func (self *TextClass) TextLength() uint16 {
	var _arg0 *C.GtkText // out
	var _cret C.guint16  // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_text_length(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// TruncateMultiline returns whether the `GtkText` will truncate multi-line text
// that is pasted into the widget
func (self *TextClass) TruncateMultiline() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_truncate_multiline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visibility retrieves whether the text in @self is visible.
func (self *TextClass) Visibility() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_get_visibility(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GrabFocusWithoutSelecting causes @self to have keyboard focus.
//
// It behaves like [method@Gtk.Widget.grab_focus], except that it doesn't select
// the contents of @self. You only want to call this on some special entries
// which the user usually doesn't want to replace all text in, such as
// search-as-you-type entries.
func (self *TextClass) GrabFocusWithoutSelecting() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_text_grab_focus_without_selecting(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatesDefault: if @activates is true, pressing Enter in the @self will
// activate the default widget for the window containing @self.
//
// This usually means that the dialog containing the `GtkText` will be closed,
// since the default widget is usually one of the dialog buttons.
func (self *TextClass) SetActivatesDefault(activates bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	if activates {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_activates_default(_arg0, _arg1)
}

// SetAttributes sets attributes that are applied to the text.
func (self *TextClass) SetAttributes(attrs *pango.AttrList) {
	var _arg0 *C.GtkText       // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs))

	C.gtk_text_set_attributes(_arg0, _arg1)
}

// SetBuffer: set the `GtkEntryBuffer` object which holds the text for this
// widget.
func (self *TextClass) SetBuffer(buffer EntryBuffer) {
	var _arg0 *C.GtkText        // out
	var _arg1 *C.GtkEntryBuffer // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_text_set_buffer(_arg0, _arg1)
}

// SetEnableEmojiCompletion sets whether Emoji completion is enabled.
//
// If it is, typing ':', followed by a recognized keyword, will pop up a window
// with suggested Emojis matching the keyword.
func (self *TextClass) SetEnableEmojiCompletion(enableEmojiCompletion bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	if enableEmojiCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_enable_emoji_completion(_arg0, _arg1)
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// @self.
func (self *TextClass) SetExtraMenu(model gio.MenuModel) {
	var _arg0 *C.GtkText    // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_text_set_extra_menu(_arg0, _arg1)
}

// SetInvisibleChar sets the character to use in place of the actual text when
// in “password mode”.
//
// By default, GTK picks the best invisible char available in the current font.
// If you set the invisible char to 0, then the user will get no feedback at
// all; there will be no text on the screen as they type.
func (self *TextClass) SetInvisibleChar(ch uint32) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gunichar // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	_arg1 = C.gunichar(ch)

	C.gtk_text_set_invisible_char(_arg0, _arg1)
}

// SetMaxLength sets the maximum allowed length of the contents of the widget.
//
// If the current contents are longer than the given length, then they will be
// truncated to fit.
//
// This is equivalent to getting @self's `GtkEntryBuffer` and calling
// [method@Gtk.EntryBuffer.set_max_length] on it.
func (self *TextClass) SetMaxLength(length int) {
	var _arg0 *C.GtkText // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(length)

	C.gtk_text_set_max_length(_arg0, _arg1)
}

// SetOverwriteMode sets whether the text is overwritten when typing in the
// `GtkText`.
func (self *TextClass) SetOverwriteMode(overwrite bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	if overwrite {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_overwrite_mode(_arg0, _arg1)
}

// SetPlaceholderText sets text to be displayed in @self when it is empty.
//
// This can be used to give a visual hint of the expected contents of the
// `GtkText`.
func (self *TextClass) SetPlaceholderText(text string) {
	var _arg0 *C.GtkText // out
	var _arg1 *C.char    // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_text_set_placeholder_text(_arg0, _arg1)
}

// SetPropagateTextWidth sets whether the `GtkText` should grow and shrink with
// the content.
func (self *TextClass) SetPropagateTextWidth(propagateTextWidth bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	if propagateTextWidth {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_propagate_text_width(_arg0, _arg1)
}

// SetTabs sets tabstops that are applied to the text.
func (self *TextClass) SetTabs(tabs *pango.TabArray) {
	var _arg0 *C.GtkText       // out
	var _arg1 *C.PangoTabArray // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.PangoTabArray)(unsafe.Pointer(tabs))

	C.gtk_text_set_tabs(_arg0, _arg1)
}

// SetTruncateMultiline sets whether the `GtkText` should truncate multi-line
// text that is pasted into the widget.
func (self *TextClass) SetTruncateMultiline(truncateMultiline bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	if truncateMultiline {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_truncate_multiline(_arg0, _arg1)
}

// SetVisibility sets whether the contents of the `GtkText` are visible or not.
//
// When visibility is set to false, characters are displayed as the invisible
// char, and will also appear that way when the text in the widget is copied to
// the clipboard.
//
// By default, GTK picks the best invisible character available in the current
// font, but it can be changed with [method@Gtk.Text.set_invisible_char].
//
// Note that you probably want to set [property@Gtk.Text:input-purpose] to
// GTK_INPUT_PURPOSE_PASSWORD or GTK_INPUT_PURPOSE_PIN to inform input methods
// about the purpose of this self, in addition to setting visibility to false.
func (self *TextClass) SetVisibility(visible bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_visibility(_arg0, _arg1)
}

// UnsetInvisibleChar unsets the invisible char.
//
// After calling this, the default invisible char is used again.
func (self *TextClass) UnsetInvisibleChar() {
	var _arg0 *C.GtkText // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(self.Native()))

	C.gtk_text_unset_invisible_char(_arg0)
}
