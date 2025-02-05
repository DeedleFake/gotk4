// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_Text_ConnectToggleOverwrite(gpointer, guintptr);
// extern void _gotk4_gtk4_Text_ConnectPreeditChanged(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk4_Text_ConnectPasteClipboard(gpointer, guintptr);
// extern void _gotk4_gtk4_Text_ConnectMoveCursor(gpointer, GtkMovementStep, gint, gboolean, guintptr);
// extern void _gotk4_gtk4_Text_ConnectInsertEmoji(gpointer, guintptr);
// extern void _gotk4_gtk4_Text_ConnectInsertAtCursor(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk4_Text_ConnectDeleteFromCursor(gpointer, GtkDeleteType, gint, guintptr);
// extern void _gotk4_gtk4_Text_ConnectCutClipboard(gpointer, guintptr);
// extern void _gotk4_gtk4_Text_ConnectCopyClipboard(gpointer, guintptr);
// extern void _gotk4_gtk4_Text_ConnectBackspace(gpointer, guintptr);
// extern void _gotk4_gtk4_Text_ConnectActivate(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeText = coreglib.Type(C.gtk_text_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeText, F: marshalText},
	})
}

// Text: GtkText widget is a single-line text entry widget.
//
// GtkText is the common implementation of single-line text editing that is
// shared between GtkEntry, GtkPasswordEntry, GtkSpinButton and other widgets.
// In all of these, GtkText` is used as the delegate for the gtk.Editable
// implementation.
//
// A fairly large set of key bindings are supported by default. If the entered
// text is longer than the allocation of the widget, the widget will scroll so
// that the cursor position is visible.
//
// When using an entry for passwords and other sensitive information, it can be
// put into “password mode” using gtk.Text.SetVisibility(). In this mode,
// entered text is displayed using a “invisible” character. By default, GTK
// picks the best invisible character that is available in the current font, but
// it can be changed with gtk.Text.SetInvisibleChar().
//
// If you are looking to add icons or progress display in an entry, look at
// GtkEntry. There other alternatives for more specialized use cases, such as
// GtkSearchEntry.
//
// If you need multi-line editable text, look at GtkTextView.
//
// CSS nodes
//
//    text[.read-only]
//    ├── placeholder
//    ├── undershoot.left
//    ├── undershoot.right
//    ├── [selection]
//    ├── [block-cursor]
//    ╰── [window.popup]
//
//
// GtkText has a main node with the name text. Depending on the properties of
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
// GtkText uses the GTK_ACCESSIBLE_ROLE_NONE role, which causes it to be skipped
// for accessibility. This is because GtkText is expected to be used as a
// delegate for a GtkEditable implementation that will be represented to
// accessibility.
type Text struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Editable
}

var (
	_ Widgetter         = (*Text)(nil)
	_ coreglib.Objector = (*Text)(nil)
)

func wrapText(obj *coreglib.Object) *Text {
	return &Text{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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
		Object: obj,
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
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
		},
	}
}

func marshalText(p uintptr) (interface{}, error) {
	return wrapText(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate is emitted when the user hits the Enter key.
//
// The default bindings for this signal are all forms of the <kbd>Enter</kbd>
// key.
func (self *Text) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "activate", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectActivate), f)
}

// ConnectBackspace is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are <kbd>Backspace</kbd> and
// <kbd>Shift</kbd>-<kbd>Backspace</kbd>.
func (self *Text) ConnectBackspace(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "backspace", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectBackspace), f)
}

// ConnectCopyClipboard is emitted to copy the selection to the clipboard.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are <kbd>Ctrl</kbd>-<kbd>c</kbd> and
// <kbd>Ctrl</kbd>-<kbd>Insert</kbd>.
func (self *Text) ConnectCopyClipboard(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "copy-clipboard", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectCopyClipboard), f)
}

// ConnectCutClipboard is emitted to cut the selection to the clipboard.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are <kbd>Ctrl</kbd>-<kbd>x</kbd> and
// <kbd>Shift</kbd>-<kbd>Delete</kbd>.
func (self *Text) ConnectCutClipboard(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "cut-clipboard", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectCutClipboard), f)
}

// ConnectDeleteFromCursor is emitted when the user initiates a text deletion.
//
// This is a keybinding signal (class.SignalAction.html).
//
// If the type is GTK_DELETE_CHARS, GTK deletes the selection if there is one,
// otherwise it deletes the requested number of characters.
//
// The default bindings for this signal are <kbd>Delete</kbd> for deleting a
// character and <kbd>Ctrl</kbd>-<kbd>Delete</kbd> for deleting a word.
func (self *Text) ConnectDeleteFromCursor(f func(typ DeleteType, count int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "delete-from-cursor", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectDeleteFromCursor), f)
}

// ConnectInsertAtCursor is emitted when the user initiates the insertion of a
// fixed string at the cursor.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This signal has no default bindings.
func (self *Text) ConnectInsertAtCursor(f func(str string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "insert-at-cursor", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectInsertAtCursor), f)
}

// ConnectInsertEmoji is emitted to present the Emoji chooser for the self.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are <kbd>Ctrl</kbd>-<kbd>.</kbd> and
// <kbd>Ctrl</kbd>-<kbd>;</kbd>.
func (self *Text) ConnectInsertEmoji(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "insert-emoji", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectInsertEmoji), f)
}

// ConnectMoveCursor is emitted when the user initiates a cursor movement.
//
// If the cursor is not visible in self, this signal causes the viewport to be
// moved instead.
//
// This is a keybinding signal (class.SignalAction.html).
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control the cursor programmatically.
//
// The default bindings for this signal come in two variants, the variant with
// the <kbd>Shift</kbd> modifier extends the selection, the variant without it
// does not. There are too many key combinations to list them all here.
//
// - <kbd>←</kbd>, <kbd>→</kbd>, <kbd>↑</kbd>, <kbd>↓</kbd> move by individual
// characters/lines
//
// - <kbd>Ctrl</kbd>-<kbd>→</kbd>, etc. move by words/paragraphs
//
// - <kbd>Home</kbd>, <kbd>End</kbd> move to the ends of the buffer.
func (self *Text) ConnectMoveCursor(f func(step MovementStep, count int, extend bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "move-cursor", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectMoveCursor), f)
}

// ConnectPasteClipboard is emitted to paste the contents of the clipboard.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal are <kbd>Ctrl</kbd>-<kbd>v</kbd> and
// <kbd>Shift</kbd>-<kbd>Insert</kbd>.
func (self *Text) ConnectPasteClipboard(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "paste-clipboard", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectPasteClipboard), f)
}

// ConnectPreeditChanged is emitted when the preedit text changes.
//
// If an input method is used, the typed text will not immediately be committed
// to the buffer. So if you are interested in the text, connect to this signal.
func (self *Text) ConnectPreeditChanged(f func(preedit string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "preedit-changed", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectPreeditChanged), f)
}

// ConnectToggleOverwrite is emitted to toggle the overwrite mode of the
// GtkText.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default bindings for this signal is <kbd>Insert</kbd>.
func (self *Text) ConnectToggleOverwrite(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "toggle-overwrite", false, unsafe.Pointer(C._gotk4_gtk4_Text_ConnectToggleOverwrite), f)
}

// NewText creates a new GtkText.
//
// The function returns the following values:
//
//    - text: new GtkText.
//
func NewText() *Text {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_text_new()

	var _text *Text // out

	_text = wrapText(coreglib.Take(unsafe.Pointer(_cret)))

	return _text
}

// NewTextWithBuffer creates a new GtkText with the specified text buffer.
//
// The function takes the following parameters:
//
//    - buffer to use for the new GtkText.
//
// The function returns the following values:
//
//    - text: new GtkText.
//
func NewTextWithBuffer(buffer *EntryBuffer) *Text {
	var _arg1 *C.GtkEntryBuffer // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_cret = C.gtk_text_new_with_buffer(_arg1)
	runtime.KeepAlive(buffer)

	var _text *Text // out

	_text = wrapText(coreglib.Take(unsafe.Pointer(_cret)))

	return _text
}

// ActivatesDefault retrieves the value set by gtk_text_set_activates_default().
//
// The function returns the following values:
//
//    - ok: TRUE if the GtkText will activate the default widget.
//
func (self *Text) ActivatesDefault() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_activates_default(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Attributes gets the attribute list that was set on the GtkText using
// gtk_text_set_attributes().
//
// The function returns the following values:
//
//    - attrList (optional): attribute list, or NULL if none was set.
//
func (self *Text) Attributes() *pango.AttrList {
	var _arg0 *C.GtkText       // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_attributes(_arg0)
	runtime.KeepAlive(self)

	var _attrList *pango.AttrList // out

	if _cret != nil {
		_attrList = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.pango_attr_list_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_attrList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_attr_list_unref((*C.PangoAttrList)(intern.C))
			},
		)
	}

	return _attrList
}

// Buffer: get the GtkEntryBuffer object which holds the text for this self.
//
// The function returns the following values:
//
//    - entryBuffer: GtkEntryBuffer object.
//
func (self *Text) Buffer() *EntryBuffer {
	var _arg0 *C.GtkText        // out
	var _cret *C.GtkEntryBuffer // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_buffer(_arg0)
	runtime.KeepAlive(self)

	var _entryBuffer *EntryBuffer // out

	_entryBuffer = wrapEntryBuffer(coreglib.Take(unsafe.Pointer(_cret)))

	return _entryBuffer
}

// EnableEmojiCompletion returns whether Emoji completion is enabled for this
// GtkText widget.
//
// The function returns the following values:
//
//    - ok: TRUE if Emoji completion is enabled.
//
func (self *Text) EnableEmojiCompletion() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_enable_emoji_completion(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ExtraMenu gets the menu model set with gtk_text_set_extra_menu().
//
// The function returns the following values:
//
//    - menuModel (optional): menu model.
//
func (self *Text) ExtraMenu() gio.MenuModeller {
	var _arg0 *C.GtkText    // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_extra_menu(_arg0)
	runtime.KeepAlive(self)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.MenuModeller)
				return ok
			})
			rv, ok := casted.(gio.MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// InputHints gets the input hints of the GtkText.
//
// The function returns the following values:
//
func (self *Text) InputHints() InputHints {
	var _arg0 *C.GtkText      // out
	var _cret C.GtkInputHints // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_input_hints(_arg0)
	runtime.KeepAlive(self)

	var _inputHints InputHints // out

	_inputHints = InputHints(_cret)

	return _inputHints
}

// InputPurpose gets the input purpose of the GtkText.
//
// The function returns the following values:
//
func (self *Text) InputPurpose() InputPurpose {
	var _arg0 *C.GtkText        // out
	var _cret C.GtkInputPurpose // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_input_purpose(_arg0)
	runtime.KeepAlive(self)

	var _inputPurpose InputPurpose // out

	_inputPurpose = InputPurpose(_cret)

	return _inputPurpose
}

// InvisibleChar retrieves the character displayed in place of the real
// characters for entries with visibility set to false.
//
// Note that GTK does not compute this value unless it needs it, so the value
// returned by this function is not very useful unless it has been explicitly
// set with gtk.Text.SetInvisibleChar().
//
// The function returns the following values:
//
//    - gunichar: current invisible char, or 0, if text does not show invisible
//      text at all.
//
func (self *Text) InvisibleChar() uint32 {
	var _arg0 *C.GtkText // out
	var _cret C.gunichar // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_invisible_char(_arg0)
	runtime.KeepAlive(self)

	var _gunichar uint32 // out

	_gunichar = uint32(_cret)

	return _gunichar
}

// MaxLength retrieves the maximum allowed length of the text in self.
//
// See gtk.Text.SetMaxLength().
//
// This is equivalent to getting self's GtkEntryBuffer and calling
// gtk.EntryBuffer.GetMaxLength() on it.
//
// The function returns the following values:
//
//    - gint: maximum allowed number of characters in GtkText, or 0 if there is
//      no maximum.
//
func (self *Text) MaxLength() int {
	var _arg0 *C.GtkText // out
	var _cret C.int      // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_max_length(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// OverwriteMode gets the value set by gtk_text_set_overwrite_mode().
//
// The function returns the following values:
//
//    - ok: whether the text is overwritten when typing.
//
func (self *Text) OverwriteMode() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_overwrite_mode(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PlaceholderText retrieves the text that will be displayed when self is empty
// and unfocused.
//
// The function returns the following values:
//
//    - utf8 (optional): pointer to the placeholder text as a string. This string
//      points to internally allocated storage in the widget and must not be
//      freed, modified or stored. If no placeholder text has been set, NULL will
//      be returned.
//
func (self *Text) PlaceholderText() string {
	var _arg0 *C.GtkText // out
	var _cret *C.char    // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_placeholder_text(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// PropagateTextWidth returns whether the GtkText will grow and shrink with the
// content.
//
// The function returns the following values:
//
//    - ok: TRUE if self will propagate the text width.
//
func (self *Text) PropagateTextWidth() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_propagate_text_width(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Tabs gets the tabstops that were set on the GtkText using
// gtk_text_set_tabs().
//
// The function returns the following values:
//
//    - tabArray (optional): tabstops, or NULL if none was set.
//
func (self *Text) Tabs() *pango.TabArray {
	var _arg0 *C.GtkText       // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_tabs(_arg0)
	runtime.KeepAlive(self)

	var _tabArray *pango.TabArray // out

	if _cret != nil {
		_tabArray = (*pango.TabArray)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _tabArray
}

// TextLength retrieves the current length of the text in self.
//
// This is equivalent to getting self's GtkEntryBuffer and calling
// gtk.EntryBuffer.GetLength() on it.
//
// The function returns the following values:
//
//    - guint16: current number of characters in GtkText, or 0 if there are none.
//
func (self *Text) TextLength() uint16 {
	var _arg0 *C.GtkText // out
	var _cret C.guint16  // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_text_length(_arg0)
	runtime.KeepAlive(self)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// TruncateMultiline returns whether the GtkText will truncate multi-line text
// that is pasted into the widget.
//
// The function returns the following values:
//
//    - ok: TRUE if self will truncate multi-line text.
//
func (self *Text) TruncateMultiline() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_truncate_multiline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visibility retrieves whether the text in self is visible.
//
// The function returns the following values:
//
//    - ok: TRUE if the text is currently visible.
//
func (self *Text) Visibility() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_get_visibility(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GrabFocusWithoutSelecting causes self to have keyboard focus.
//
// It behaves like gtk.Widget.GrabFocus(), except that it doesn't select the
// contents of self. You only want to call this on some special entries which
// the user usually doesn't want to replace all text in, such as
// search-as-you-type entries.
//
// The function returns the following values:
//
//    - ok: TRUE if focus is now inside self.
//
func (self *Text) GrabFocusWithoutSelecting() bool {
	var _arg0 *C.GtkText // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_text_grab_focus_without_selecting(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatesDefault: if activates is TRUE, pressing Enter in the self will
// activate the default widget for the window containing self.
//
// This usually means that the dialog containing the GtkText will be closed,
// since the default widget is usually one of the dialog buttons.
//
// The function takes the following parameters:
//
//    - activates: TRUE to activate window’s default widget on Enter keypress.
//
func (self *Text) SetActivatesDefault(activates bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if activates {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_activates_default(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(activates)
}

// SetAttributes sets attributes that are applied to the text.
//
// The function takes the following parameters:
//
//    - attrs (optional): PangoAttrList or NULL to unset.
//
func (self *Text) SetAttributes(attrs *pango.AttrList) {
	var _arg0 *C.GtkText       // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if attrs != nil {
		_arg1 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attrs)))
	}

	C.gtk_text_set_attributes(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(attrs)
}

// SetBuffer: set the GtkEntryBuffer object which holds the text for this
// widget.
//
// The function takes the following parameters:
//
//    - buffer: GtkEntryBuffer.
//
func (self *Text) SetBuffer(buffer *EntryBuffer) {
	var _arg0 *C.GtkText        // out
	var _arg1 *C.GtkEntryBuffer // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	C.gtk_text_set_buffer(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(buffer)
}

// SetEnableEmojiCompletion sets whether Emoji completion is enabled.
//
// If it is, typing ':', followed by a recognized keyword, will pop up a window
// with suggested Emojis matching the keyword.
//
// The function takes the following parameters:
//
//    - enableEmojiCompletion: TRUE to enable Emoji completion.
//
func (self *Text) SetEnableEmojiCompletion(enableEmojiCompletion bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if enableEmojiCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_enable_emoji_completion(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableEmojiCompletion)
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// self.
//
// The function takes the following parameters:
//
//    - model (optional): GMenuModel.
//
func (self *Text) SetExtraMenu(model gio.MenuModeller) {
	var _arg0 *C.GtkText    // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.gtk_text_set_extra_menu(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetInputHints sets input hints that allow input methods to fine-tune their
// behaviour.
//
// The function takes the following parameters:
//
//    - hints: hints.
//
func (self *Text) SetInputHints(hints InputHints) {
	var _arg0 *C.GtkText      // out
	var _arg1 C.GtkInputHints // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkInputHints(hints)

	C.gtk_text_set_input_hints(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(hints)
}

// SetInputPurpose sets the input purpose of the GtkText.
//
// This can be used by on-screen keyboards and other input methods to adjust
// their behaviour.
//
// The function takes the following parameters:
//
//    - purpose: purpose.
//
func (self *Text) SetInputPurpose(purpose InputPurpose) {
	var _arg0 *C.GtkText        // out
	var _arg1 C.GtkInputPurpose // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkInputPurpose(purpose)

	C.gtk_text_set_input_purpose(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(purpose)
}

// SetInvisibleChar sets the character to use in place of the actual text when
// in “password mode”.
//
// By default, GTK picks the best invisible char available in the current font.
// If you set the invisible char to 0, then the user will get no feedback at
// all; there will be no text on the screen as they type.
//
// The function takes the following parameters:
//
//    - ch: unicode character.
//
func (self *Text) SetInvisibleChar(ch uint32) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gunichar // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.gunichar(ch)

	C.gtk_text_set_invisible_char(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ch)
}

// SetMaxLength sets the maximum allowed length of the contents of the widget.
//
// If the current contents are longer than the given length, then they will be
// truncated to fit.
//
// This is equivalent to getting self's GtkEntryBuffer and calling
// gtk.EntryBuffer.SetMaxLength() on it.
//
// The function takes the following parameters:
//
//    - length: maximum length of the GtkText, or 0 for no maximum. (other than
//      the maximum length of entries.) The value passed in will be clamped to
//      the range 0-65536.
//
func (self *Text) SetMaxLength(length int) {
	var _arg0 *C.GtkText // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(length)

	C.gtk_text_set_max_length(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(length)
}

// SetOverwriteMode sets whether the text is overwritten when typing in the
// GtkText.
//
// The function takes the following parameters:
//
//    - overwrite: new value.
//
func (self *Text) SetOverwriteMode(overwrite bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if overwrite {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_overwrite_mode(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(overwrite)
}

// SetPlaceholderText sets text to be displayed in self when it is empty.
//
// This can be used to give a visual hint of the expected contents of the
// GtkText.
//
// The function takes the following parameters:
//
//    - text (optional): string to be displayed when self is empty and unfocused,
//      or NULL.
//
func (self *Text) SetPlaceholderText(text string) {
	var _arg0 *C.GtkText // out
	var _arg1 *C.char    // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if text != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_text_set_placeholder_text(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(text)
}

// SetPropagateTextWidth sets whether the GtkText should grow and shrink with
// the content.
//
// The function takes the following parameters:
//
//    - propagateTextWidth: TRUE to propagate the text width.
//
func (self *Text) SetPropagateTextWidth(propagateTextWidth bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if propagateTextWidth {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_propagate_text_width(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(propagateTextWidth)
}

// SetTabs sets tabstops that are applied to the text.
//
// The function takes the following parameters:
//
//    - tabs (optional): PangoTabArray.
//
func (self *Text) SetTabs(tabs *pango.TabArray) {
	var _arg0 *C.GtkText       // out
	var _arg1 *C.PangoTabArray // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if tabs != nil {
		_arg1 = (*C.PangoTabArray)(gextras.StructNative(unsafe.Pointer(tabs)))
	}

	C.gtk_text_set_tabs(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(tabs)
}

// SetTruncateMultiline sets whether the GtkText should truncate multi-line text
// that is pasted into the widget.
//
// The function takes the following parameters:
//
//    - truncateMultiline: TRUE to truncate multi-line text.
//
func (self *Text) SetTruncateMultiline(truncateMultiline bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if truncateMultiline {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_truncate_multiline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(truncateMultiline)
}

// SetVisibility sets whether the contents of the GtkText are visible or not.
//
// When visibility is set to FALSE, characters are displayed as the invisible
// char, and will also appear that way when the text in the widget is copied to
// the clipboard.
//
// By default, GTK picks the best invisible character available in the current
// font, but it can be changed with gtk.Text.SetInvisibleChar().
//
// Note that you probably want to set gtk.Text:input-purpose to
// GTK_INPUT_PURPOSE_PASSWORD or GTK_INPUT_PURPOSE_PIN to inform input methods
// about the purpose of this self, in addition to setting visibility to FALSE.
//
// The function takes the following parameters:
//
//    - visible: TRUE if the contents of the GtkText are displayed as plaintext.
//
func (self *Text) SetVisibility(visible bool) {
	var _arg0 *C.GtkText // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_text_set_visibility(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(visible)
}

// UnsetInvisibleChar unsets the invisible char.
//
// After calling this, the default invisible char is used again.
func (self *Text) UnsetInvisibleChar() {
	var _arg0 *C.GtkText // out

	_arg0 = (*C.GtkText)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gtk_text_unset_invisible_char(_arg0)
	runtime.KeepAlive(self)
}
