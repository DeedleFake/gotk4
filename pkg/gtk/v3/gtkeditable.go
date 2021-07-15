// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_editable_get_type()), F: marshalEditabler},
	})
}

// EditableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type EditableOverrider interface {
	Changed()
	// DeleteText deletes a sequence of characters. The characters that are
	// deleted are those characters at positions from start_pos up to, but not
	// including end_pos. If end_pos is negative, then the characters deleted
	// are those from start_pos to the end of the text.
	//
	// Note that the positions are specified in characters, not bytes.
	DeleteText(startPos int, endPos int)
	// DoDeleteText deletes a sequence of characters. The characters that are
	// deleted are those characters at positions from start_pos up to, but not
	// including end_pos. If end_pos is negative, then the characters deleted
	// are those from start_pos to the end of the text.
	//
	// Note that the positions are specified in characters, not bytes.
	DoDeleteText(startPos int, endPos int)
	// Chars retrieves a sequence of characters. The characters that are
	// retrieved are those characters at positions from start_pos up to, but not
	// including end_pos. If end_pos is negative, then the characters retrieved
	// are those characters from start_pos to the end of the text.
	//
	// Note that positions are specified in characters, not bytes.
	Chars(startPos int, endPos int) string
	// Position retrieves the current position of the cursor relative to the
	// start of the content of the editable.
	//
	// Note that this position is in characters, not in bytes.
	Position() int
	// SelectionBounds retrieves the selection bound of the editable. start_pos
	// will be filled with the start of the selection and end_pos with end. If
	// no text was selected both will be identical and FALSE will be returned.
	//
	// Note that positions are specified in characters, not bytes.
	SelectionBounds() (startPos int, endPos int, ok bool)
	// SetPosition sets the cursor position in the editable to the given value.
	//
	// The cursor is displayed before the character with the given (base 0)
	// index in the contents of the editable. The value must be less than or
	// equal to the number of characters in the editable. A value of -1
	// indicates that the position should be set after the last character of the
	// editable. Note that position is in characters, not in bytes.
	SetPosition(position int)
	// SetSelectionBounds selects a region of text. The characters that are
	// selected are those characters at positions from start_pos up to, but not
	// including end_pos. If end_pos is negative, then the characters selected
	// are those characters from start_pos to the end of the text.
	//
	// Note that positions are specified in characters, not bytes.
	SetSelectionBounds(startPos int, endPos int)
}

// Editable interface is an interface which should be implemented by text
// editing widgets, such as Entry and SpinButton. It contains functions for
// generically manipulating an editable widget, a large number of action signals
// used for key bindings, and several signals that an application can connect to
// to modify the behavior of a widget.
//
// As an example of the latter usage, by connecting the following handler to
// Editable::insert-text, an application can convert all entry into a widget
// into uppercase.
//
// Forcing entry to uppercase.
//
//    #include <ctype.h>;
//
//    void
//    insert_text_handler (GtkEditable *editable,
//                         const gchar *text,
//                         gint         length,
//                         gint        *position,
//                         gpointer     data)
//    {
//      gchar *result = g_utf8_strup (text, length);
//
//      g_signal_handlers_block_by_func (editable,
//                                   (gpointer) insert_text_handler, data);
//      gtk_editable_insert_text (editable, result, length, position);
//      g_signal_handlers_unblock_by_func (editable,
//                                         (gpointer) insert_text_handler, data);
//
//      g_signal_stop_emission_by_name (editable, "insert_text");
//
//      g_free (result);
//    }
type Editable struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Editable)(nil)

// Editabler describes Editable's abstract methods.
type Editabler interface {
	// CopyClipboard copies the contents of the currently selected content in
	// the editable and puts it on the clipboard.
	CopyClipboard()
	// CutClipboard removes the contents of the currently selected content in
	// the editable and puts it on the clipboard.
	CutClipboard()
	// DeleteSelection deletes the currently selected text of the editable.
	DeleteSelection()
	// DeleteText deletes a sequence of characters.
	DeleteText(startPos int, endPos int)
	// Chars retrieves a sequence of characters.
	Chars(startPos int, endPos int) string
	// Editable retrieves whether editable is editable.
	Editable() bool
	// Position retrieves the current position of the cursor relative to the
	// start of the content of the editable.
	Position() int
	// SelectionBounds retrieves the selection bound of the editable.
	SelectionBounds() (startPos int, endPos int, ok bool)
	// PasteClipboard pastes the content of the clipboard to the current
	// position of the cursor in the editable.
	PasteClipboard()
	// SelectRegion selects a region of text.
	SelectRegion(startPos int, endPos int)
	// SetEditable determines if the user can edit the text in the editable
	// widget or not.
	SetEditable(isEditable bool)
	// SetPosition sets the cursor position in the editable to the given value.
	SetPosition(position int)
}

var _ Editabler = (*Editable)(nil)

func wrapEditable(obj *externglib.Object) *Editable {
	return &Editable{
		Object: obj,
	}
}

func marshalEditabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEditable(obj), nil
}

// CopyClipboard copies the contents of the currently selected content in the
// editable and puts it on the clipboard.
func (editable *Editable) CopyClipboard() {
	var _arg0 *C.GtkEditable // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	C.gtk_editable_copy_clipboard(_arg0)
}

// CutClipboard removes the contents of the currently selected content in the
// editable and puts it on the clipboard.
func (editable *Editable) CutClipboard() {
	var _arg0 *C.GtkEditable // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	C.gtk_editable_cut_clipboard(_arg0)
}

// DeleteSelection deletes the currently selected text of the editable. This
// call doesn’t do anything if there is no selected text.
func (editable *Editable) DeleteSelection() {
	var _arg0 *C.GtkEditable // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	C.gtk_editable_delete_selection(_arg0)
}

// DeleteText deletes a sequence of characters. The characters that are deleted
// are those characters at positions from start_pos up to, but not including
// end_pos. If end_pos is negative, then the characters deleted are those from
// start_pos to the end of the text.
//
// Note that the positions are specified in characters, not bytes.
func (editable *Editable) DeleteText(startPos int, endPos int) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.gint         // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	C.gtk_editable_delete_text(_arg0, _arg1, _arg2)
}

// Chars retrieves a sequence of characters. The characters that are retrieved
// are those characters at positions from start_pos up to, but not including
// end_pos. If end_pos is negative, then the characters retrieved are those
// characters from start_pos to the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (editable *Editable) Chars(startPos int, endPos int) string {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.gint         // out
	var _arg2 C.gint         // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	_cret = C.gtk_editable_get_chars(_arg0, _arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Editable retrieves whether editable is editable. See
// gtk_editable_set_editable().
func (editable *Editable) Editable() bool {
	var _arg0 *C.GtkEditable // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_editable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Position retrieves the current position of the cursor relative to the start
// of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (editable *Editable) Position() int {
	var _arg0 *C.GtkEditable // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_position(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SelectionBounds retrieves the selection bound of the editable. start_pos will
// be filled with the start of the selection and end_pos with end. If no text
// was selected both will be identical and FALSE will be returned.
//
// Note that positions are specified in characters, not bytes.
func (editable *Editable) SelectionBounds() (startPos int, endPos int, ok bool) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.gint         // in
	var _arg2 C.gint         // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_selection_bounds(_arg0, &_arg1, &_arg2)

	var _startPos int // out
	var _endPos int   // out
	var _ok bool      // out

	_startPos = int(_arg1)
	_endPos = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _startPos, _endPos, _ok
}

// PasteClipboard pastes the content of the clipboard to the current position of
// the cursor in the editable.
func (editable *Editable) PasteClipboard() {
	var _arg0 *C.GtkEditable // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	C.gtk_editable_paste_clipboard(_arg0)
}

// SelectRegion selects a region of text. The characters that are selected are
// those characters at positions from start_pos up to, but not including
// end_pos. If end_pos is negative, then the characters selected are those
// characters from start_pos to the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (editable *Editable) SelectRegion(startPos int, endPos int) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.gint         // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.gint(startPos)
	_arg2 = C.gint(endPos)

	C.gtk_editable_select_region(_arg0, _arg1, _arg2)
}

// SetEditable determines if the user can edit the text in the editable widget
// or not.
func (editable *Editable) SetEditable(isEditable bool) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	if isEditable {
		_arg1 = C.TRUE
	}

	C.gtk_editable_set_editable(_arg0, _arg1)
}

// SetPosition sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0) index in
// the contents of the editable. The value must be less than or equal to the
// number of characters in the editable. A value of -1 indicates that the
// position should be set after the last character of the editable. Note that
// position is in characters, not in bytes.
func (editable *Editable) SetPosition(position int) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.gint(position)

	C.gtk_editable_set_position(_arg0, _arg1)
}
