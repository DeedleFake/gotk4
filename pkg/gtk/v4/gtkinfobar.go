// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_info_bar_get_type()), F: marshalInfoBar},
	})
}

// InfoBar: `GtkInfoBar` can be show messages to the user without a dialog.
//
// !An example GtkInfoBar (info-bar.png)
//
// It is often temporarily shown at the top or bottom of a document. In contrast
// to [class@Gtk.Dialog], which has an action area at the bottom, `GtkInfoBar`
// has an action area at the side.
//
// The API of `GtkInfoBar` is very similar to `GtkDialog`, allowing you to add
// buttons to the action area with [method@Gtk.InfoBar.add_button] or
// [ctor@Gtk.InfoBar.new_with_buttons]. The sensitivity of action widgets can be
// controlled with [method@Gtk.InfoBar.set_response_sensitive].
//
// To add widgets to the main content area of a `GtkInfoBar`, use
// [method@Gtk.InfoBar.add_child].
//
// Similar to [class@Gtk.MessageDialog], the contents of a `GtkInfoBar` can by
// classified as error message, warning, informational message, etc, by using
// [method@Gtk.InfoBar.set_message_type]. GTK may use the message type to
// determine how the message is displayed.
//
// A simple example for using a `GtkInfoBar`: “`c GtkWidget *message_label;
// GtkWidget *widget; GtkWidget *grid; GtkInfoBar *bar;
//
// // set up info bar widget = gtk_info_bar_new (); bar = GTK_INFO_BAR (widget);
// // grid = gtk_grid_new ();
//
// message_label = gtk_label_new (""); gtk_info_bar_add_child (bar,
// message_label); gtk_info_bar_add_button (bar, _("_OK"), GTK_RESPONSE_OK);
// g_signal_connect (bar, "response", G_CALLBACK (gtk_widget_hide), NULL);
// gtk_grid_attach (GTK_GRID (grid), widget, 0, 2, 1, 1);
//
// // ...
//
// // show an error message gtk_label_set_text (GTK_LABEL (message_label), "An
// // error occurred!"); gtk_info_bar_set_message_type (bar, GTK_MESSAGE_ERROR);
// // gtk_widget_show (bar); “`
//
//
// GtkInfoBar as GtkBuildable
//
// The `GtkInfoBar` implementation of the `GtkBuildable` interface exposes the
// content area and action area as internal children with the names
// “content_area” and “action_area”.
//
// `GtkInfoBar` supports a custom <action-widgets> element, which can contain
// multiple <action-widget> elements. The “response” attribute specifies a
// numeric response, and the content of the element is the id of widget (which
// should be a child of the dialogs @action_area).
//
//
// CSS nodes
//
// `GtkInfoBar` has a single CSS node with name infobar. The node may get one of
// the style classes .info, .warning, .error or .question, depending on the
// message type. If the info bar shows a close button, that button will have the
// .close style class applied.
type InfoBar interface {
	gextras.Objector

	// AddActionWidget: add an activatable widget to the action area of a
	// `GtkInfoBar`.
	//
	// This also connects a signal handler that will emit the
	// [signal@Gtk.InfoBar::response] signal on the message area when the widget
	// is activated. The widget is appended to the end of the message areas
	// action area.
	AddActionWidget(child Widget, responseId int)
	// AddButton adds a button with the given text.
	//
	// Clicking the button will emit the [signal@Gtk.InfoBar::response] signal
	// with the given response_id. The button is appended to the end of the info
	// bars's action area. The button widget is returned, but usually you don't
	// need it.
	AddButton(buttonText string, responseId int) *ButtonClass
	// AddChild adds a widget to the content area of the info bar.
	AddChild(widget Widget)
	// MessageType returns the message type of the message area.
	MessageType() MessageType
	// Revealed returns whether the info bar is currently revealed.
	Revealed() bool
	// ShowCloseButton returns whether the widget will display a standard close
	// button.
	ShowCloseButton() bool
	// RemoveActionWidget removes a widget from the action area of @info_bar.
	//
	// The widget must have been put there by a call to
	// [method@Gtk.InfoBar.add_action_widget] or
	// [method@Gtk.InfoBar.add_button].
	RemoveActionWidget(widget Widget)
	// RemoveChild removes a widget from the content area of the info bar.
	RemoveChild(widget Widget)
	// Response emits the “response” signal with the given @response_id.
	Response(responseId int)
	// SetDefaultResponse sets the last widget in the info bar’s action area
	// with the given response_id as the default widget for the dialog.
	//
	// Pressing “Enter” normally activates the default widget.
	//
	// Note that this function currently requires @info_bar to be added to a
	// widget hierarchy.
	SetDefaultResponse(responseId int)
	// SetResponseSensitive sets the sensitivity of action widgets for
	// @response_id.
	//
	// Calls `gtk_widget_set_sensitive (widget, setting)` for each widget in the
	// info bars’s action area with the given @response_id. A convenient way to
	// sensitize/desensitize buttons.
	SetResponseSensitive(responseId int, setting bool)
	// SetRevealed sets whether the `GtkInfoBar` is revealed.
	//
	// Changing this will make @info_bar reveal or conceal itself via a sliding
	// transition.
	//
	// Note: this does not show or hide @info_bar in the
	// [property@Gtk.Widget:visible] sense, so revealing has no effect if
	// [property@Gtk.Widget:visible] is false.
	SetRevealed(revealed bool)
	// SetShowCloseButton: if true, a standard close button is shown.
	//
	// When clicked it emits the response GTK_RESPONSE_CLOSE.
	SetShowCloseButton(setting bool)
}

// InfoBarClass implements the InfoBar interface.
type InfoBarClass struct {
	*externglib.Object
	WidgetClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
}

var _ InfoBar = (*InfoBarClass)(nil)

func wrapInfoBar(obj *externglib.Object) InfoBar {
	return &InfoBarClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
	}
}

func marshalInfoBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInfoBar(obj), nil
}

// NewInfoBar creates a new `GtkInfoBar` object.
func NewInfoBar() *InfoBarClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_info_bar_new()

	var _infoBar *InfoBarClass // out

	_infoBar = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*InfoBarClass)

	return _infoBar
}

// AddActionWidget: add an activatable widget to the action area of a
// `GtkInfoBar`.
//
// This also connects a signal handler that will emit the
// [signal@Gtk.InfoBar::response] signal on the message area when the widget is
// activated. The widget is appended to the end of the message areas action
// area.
func (i *InfoBarClass) AddActionWidget(child Widget, responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))
	_arg2 = C.int(responseId)

	C.gtk_info_bar_add_action_widget(_arg0, _arg1, _arg2)
}

// AddButton adds a button with the given text.
//
// Clicking the button will emit the [signal@Gtk.InfoBar::response] signal with
// the given response_id. The button is appended to the end of the info bars's
// action area. The button widget is returned, but usually you don't need it.
func (i *InfoBarClass) AddButton(buttonText string, responseId int) *ButtonClass {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.char       // out
	var _arg2 C.int         // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	_arg1 = (*C.char)(C.CString(buttonText))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(responseId)

	_cret = C.gtk_info_bar_add_button(_arg0, _arg1, _arg2)

	var _button *ButtonClass // out

	_button = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ButtonClass)

	return _button
}

// AddChild adds a widget to the content area of the info bar.
func (i *InfoBarClass) AddChild(widget Widget) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	C.gtk_info_bar_add_child(_arg0, _arg1)
}

// MessageType returns the message type of the message area.
func (i *InfoBarClass) MessageType() MessageType {
	var _arg0 *C.GtkInfoBar    // out
	var _cret C.GtkMessageType // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))

	_cret = C.gtk_info_bar_get_message_type(_arg0)

	var _messageType MessageType // out

	_messageType = (MessageType)(C.GtkMessageType)

	return _messageType
}

// Revealed returns whether the info bar is currently revealed.
func (i *InfoBarClass) Revealed() bool {
	var _arg0 *C.GtkInfoBar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))

	_cret = C.gtk_info_bar_get_revealed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether the widget will display a standard close
// button.
func (i *InfoBarClass) ShowCloseButton() bool {
	var _arg0 *C.GtkInfoBar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))

	_cret = C.gtk_info_bar_get_show_close_button(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveActionWidget removes a widget from the action area of @info_bar.
//
// The widget must have been put there by a call to
// [method@Gtk.InfoBar.add_action_widget] or [method@Gtk.InfoBar.add_button].
func (i *InfoBarClass) RemoveActionWidget(widget Widget) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	C.gtk_info_bar_remove_action_widget(_arg0, _arg1)
}

// RemoveChild removes a widget from the content area of the info bar.
func (i *InfoBarClass) RemoveChild(widget Widget) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	C.gtk_info_bar_remove_child(_arg0, _arg1)
}

// Response emits the “response” signal with the given @response_id.
func (i *InfoBarClass) Response(responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.int         // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	_arg1 = C.int(responseId)

	C.gtk_info_bar_response(_arg0, _arg1)
}

// SetDefaultResponse sets the last widget in the info bar’s action area with
// the given response_id as the default widget for the dialog.
//
// Pressing “Enter” normally activates the default widget.
//
// Note that this function currently requires @info_bar to be added to a widget
// hierarchy.
func (i *InfoBarClass) SetDefaultResponse(responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.int         // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	_arg1 = C.int(responseId)

	C.gtk_info_bar_set_default_response(_arg0, _arg1)
}

// SetResponseSensitive sets the sensitivity of action widgets for @response_id.
//
// Calls `gtk_widget_set_sensitive (widget, setting)` for each widget in the
// info bars’s action area with the given @response_id. A convenient way to
// sensitize/desensitize buttons.
func (i *InfoBarClass) SetResponseSensitive(responseId int, setting bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.int         // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	_arg1 = C.int(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_info_bar_set_response_sensitive(_arg0, _arg1, _arg2)
}

// SetRevealed sets whether the `GtkInfoBar` is revealed.
//
// Changing this will make @info_bar reveal or conceal itself via a sliding
// transition.
//
// Note: this does not show or hide @info_bar in the
// [property@Gtk.Widget:visible] sense, so revealing has no effect if
// [property@Gtk.Widget:visible] is false.
func (i *InfoBarClass) SetRevealed(revealed bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	if revealed {
		_arg1 = C.TRUE
	}

	C.gtk_info_bar_set_revealed(_arg0, _arg1)
}

// SetShowCloseButton: if true, a standard close button is shown.
//
// When clicked it emits the response GTK_RESPONSE_CLOSE.
func (i *InfoBarClass) SetShowCloseButton(setting bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer((&InfoBar).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_info_bar_set_show_close_button(_arg0, _arg1)
}
