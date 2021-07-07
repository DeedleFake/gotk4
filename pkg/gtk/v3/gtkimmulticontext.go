// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_im_multicontext_get_type()), F: marshalIMMulticontext},
	})
}

type IMMulticontext interface {
	IMContext

	// AsIMContext casts the class to the IMContext interface.
	AsIMContext() IMContext

	// DeleteSurrounding asks the widget that the input context is attached to
	// to delete characters around the cursor position by emitting the
	// GtkIMContext::delete_surrounding signal. Note that @offset and @n_chars
	// are in characters not in bytes which differs from the usage other places
	// in IMContext.
	//
	// In order to use this function, you should first call
	// gtk_im_context_get_surrounding() to get the current context, and call
	// this function immediately afterwards to make sure that you know what you
	// are deleting. You should also account for the fact that even if the
	// signal was handled, the input context might not have deleted all the
	// characters that were requested to be deleted.
	//
	// This function is used by an input method that wants to make subsitutions
	// in the existing text in response to new input. It is not useful for
	// applications.
	//
	// This method is inherited from IMContext
	DeleteSurrounding(offset int, nChars int) bool
	// FilterKeypress: allow an input method to internally handle key press and
	// release events. If this function returns true, then no further processing
	// should be done for this key event.
	//
	// This method is inherited from IMContext
	FilterKeypress(event *gdk.EventKey) bool
	// FocusIn: notify the input method that the widget to which this input
	// context corresponds has gained focus. The input method may, for example,
	// change the displayed feedback to reflect this change.
	//
	// This method is inherited from IMContext
	FocusIn()
	// FocusOut: notify the input method that the widget to which this input
	// context corresponds has lost focus. The input method may, for example,
	// change the displayed feedback or reset the contexts state to reflect this
	// change.
	//
	// This method is inherited from IMContext
	FocusOut()
	// GetPreeditString: retrieve the current preedit string for the input
	// context, and a list of attributes to apply to the string. This string
	// should be displayed inserted at the insertion point.
	//
	// This method is inherited from IMContext
	GetPreeditString() (string, *pango.AttrList, int)
	// GetSurrounding retrieves context around the insertion point. Input
	// methods typically want context in order to constrain input text based on
	// existing text; this is important for languages such as Thai where only
	// some sequences of characters are allowed.
	//
	// This function is implemented by emitting the
	// GtkIMContext::retrieve_surrounding signal on the input method; in
	// response to this signal, a widget should provide as much context as is
	// available, up to an entire paragraph, by calling
	// gtk_im_context_set_surrounding(). Note that there is no obligation for a
	// widget to respond to the ::retrieve_surrounding signal, so input methods
	// must be prepared to function without context.
	//
	// This method is inherited from IMContext
	GetSurrounding() (string, int, bool)
	// Reset: notify the input method that a change such as a change in cursor
	// position has been made. This will typically cause the input method to
	// clear the preedit state.
	//
	// This method is inherited from IMContext
	Reset()
	// SetClientWindow: set the client window for the input context; this is the
	// Window in which the input appears. This window is used in order to
	// correctly position status windows, and may also be used for purposes
	// internal to the input method.
	//
	// This method is inherited from IMContext
	SetClientWindow(window gdk.Window)
	// SetCursorLocation: notify the input method that a change in cursor
	// position has been made. The location is relative to the client window.
	//
	// This method is inherited from IMContext
	SetCursorLocation(area *gdk.Rectangle)
	// SetSurrounding sets surrounding context around the insertion point and
	// preedit string. This function is expected to be called in response to the
	// GtkIMContext::retrieve_surrounding signal, and will likely have no effect
	// if called at other times.
	//
	// This method is inherited from IMContext
	SetSurrounding(text string, len int, cursorIndex int)
	// SetUsePreedit sets whether the IM context should use the preedit string
	// to display feedback. If @use_preedit is FALSE (default is TRUE), then the
	// IM context may use some other method to display feedback, such as
	// displaying it in a child of the root window.
	//
	// This method is inherited from IMContext
	SetUsePreedit(usePreedit bool)

	// AppendMenuitems: add menuitems for various available input methods to a
	// menu; the menuitems, when selected, will switch the input method for the
	// context and the global default input method.
	//
	// Deprecated: since version 3.10.
	AppendMenuitems(menushell MenuShell)
	// ContextID gets the id of the currently active slave of the @context.
	ContextID() string
	// SetContextID sets the context id for @context.
	//
	// This causes the currently active slave of @context to be replaced by the
	// slave corresponding to the new context id.
	SetContextID(contextId string)
}

// imMulticontext implements the IMMulticontext interface.
type imMulticontext struct {
	*externglib.Object
}

var _ IMMulticontext = (*imMulticontext)(nil)

// WrapIMMulticontext wraps a GObject to a type that implements
// interface IMMulticontext. It is primarily used internally.
func WrapIMMulticontext(obj *externglib.Object) IMMulticontext {
	return imMulticontext{obj}
}

func marshalIMMulticontext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMMulticontext(obj), nil
}

// NewIMMulticontext creates a new IMMulticontext.
func NewIMMulticontext() IMMulticontext {
	var _cret *C.GtkIMContext // in

	_cret = C.gtk_im_multicontext_new()

	var _imMulticontext IMMulticontext // out

	_imMulticontext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(IMMulticontext)

	return _imMulticontext
}

func (i imMulticontext) AsIMContext() IMContext {
	return WrapIMContext(gextras.InternObject(i))
}

func (c imMulticontext) DeleteSurrounding(offset int, nChars int) bool {
	return WrapIMContext(gextras.InternObject(c)).DeleteSurrounding(offset, nChars)
}

func (c imMulticontext) FilterKeypress(event *gdk.EventKey) bool {
	return WrapIMContext(gextras.InternObject(c)).FilterKeypress(event)
}

func (c imMulticontext) FocusIn() {
	WrapIMContext(gextras.InternObject(c)).FocusIn()
}

func (c imMulticontext) FocusOut() {
	WrapIMContext(gextras.InternObject(c)).FocusOut()
}

func (c imMulticontext) GetPreeditString() (string, *pango.AttrList, int) {
	return WrapIMContext(gextras.InternObject(c)).GetPreeditString()
}

func (c imMulticontext) GetSurrounding() (string, int, bool) {
	return WrapIMContext(gextras.InternObject(c)).GetSurrounding()
}

func (c imMulticontext) Reset() {
	WrapIMContext(gextras.InternObject(c)).Reset()
}

func (c imMulticontext) SetClientWindow(window gdk.Window) {
	WrapIMContext(gextras.InternObject(c)).SetClientWindow(window)
}

func (c imMulticontext) SetCursorLocation(area *gdk.Rectangle) {
	WrapIMContext(gextras.InternObject(c)).SetCursorLocation(area)
}

func (c imMulticontext) SetSurrounding(text string, len int, cursorIndex int) {
	WrapIMContext(gextras.InternObject(c)).SetSurrounding(text, len, cursorIndex)
}

func (c imMulticontext) SetUsePreedit(usePreedit bool) {
	WrapIMContext(gextras.InternObject(c)).SetUsePreedit(usePreedit)
}

func (c imMulticontext) AppendMenuitems(menushell MenuShell) {
	var _arg0 *C.GtkIMMulticontext // out
	var _arg1 *C.GtkMenuShell      // out

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkMenuShell)(unsafe.Pointer(menushell.Native()))

	C.gtk_im_multicontext_append_menuitems(_arg0, _arg1)
}

func (c imMulticontext) ContextID() string {
	var _arg0 *C.GtkIMMulticontext // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_im_multicontext_get_context_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (c imMulticontext) SetContextID(contextId string) {
	var _arg0 *C.GtkIMMulticontext // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(contextId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_im_multicontext_set_context_id(_arg0, _arg1)
}
