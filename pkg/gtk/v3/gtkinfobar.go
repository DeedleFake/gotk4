// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_InfoBarClass_close(GtkInfoBar*);
// extern void _gotk4_gtk3_InfoBarClass_response(GtkInfoBar*, gint);
// extern void _gotk4_gtk3_InfoBar_ConnectClose(gpointer, guintptr);
// extern void _gotk4_gtk3_InfoBar_ConnectResponse(gpointer, gint, guintptr);
import "C"

// GType values.
var (
	GTypeInfoBar = coreglib.Type(C.gtk_info_bar_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeInfoBar, F: marshalInfoBar},
	})
}

// InfoBarOverrider contains methods that are overridable.
type InfoBarOverrider interface {
	Close()
	// Response emits the “response” signal with the given response_id.
	//
	// The function takes the following parameters:
	//
	//    - responseId: response ID.
	//
	Response(responseId int)
}

// InfoBar is a widget that can be used to show messages to the user without
// showing a dialog. It is often temporarily shown at the top or bottom of a
// document. In contrast to Dialog, which has a action area at the bottom,
// InfoBar has an action area at the side.
//
// The API of InfoBar is very similar to Dialog, allowing you to add buttons to
// the action area with gtk_info_bar_add_button() or
// gtk_info_bar_new_with_buttons(). The sensitivity of action widgets can be
// controlled with gtk_info_bar_set_response_sensitive(). To add widgets to the
// main content area of a InfoBar, use gtk_info_bar_get_content_area() and add
// your widgets to the container.
//
// Similar to MessageDialog, the contents of a InfoBar can by classified as
// error message, warning, informational message, etc, by using
// gtk_info_bar_set_message_type(). GTK+ may use the message type to determine
// how the message is displayed.
//
// A simple example for using a InfoBar:
//
//    GtkWidget *widget, *message_label, *content_area;
//    GtkWidget *grid;
//    GtkInfoBar *bar;
//
//    // set up info bar
//    widget = gtk_info_bar_new ();
//    bar = GTK_INFO_BAR (widget);
//    grid = gtk_grid_new ();
//
//    gtk_widget_set_no_show_all (widget, TRUE);
//    message_label = gtk_label_new ("");
//    content_area = gtk_info_bar_get_content_area (bar);
//    gtk_container_add (GTK_CONTAINER (content_area),
//                       message_label);
//    gtk_info_bar_add_button (bar,
//                             _("_OK"),
//                             GTK_RESPONSE_OK);
//    g_signal_connect (bar,
//                      "response",
//                      G_CALLBACK (gtk_widget_hide),
//                      NULL);
//    gtk_grid_attach (GTK_GRID (grid),
//                     widget,
//                     0, 2, 1, 1);
//
//    // ...
//
//    // show an error message
//    gtk_label_set_text (GTK_LABEL (message_label), "An error occurred!");
//    gtk_info_bar_set_message_type (bar,
//                                   GTK_MESSAGE_ERROR);
//    gtk_widget_show (bar);
//
//
// GtkInfoBar as GtkBuildable
//
// The GtkInfoBar implementation of the GtkBuildable interface exposes the
// content area and action area as internal children with the names
// “content_area” and “action_area”.
//
// GtkInfoBar supports a custom <action-widgets> element, which can contain
// multiple <action-widget> elements. The “response” attribute specifies a
// numeric response, and the content of the element is the id of widget (which
// should be a child of the dialogs action_area).
//
//
// CSS nodes
//
// GtkInfoBar has a single CSS node with name infobar. The node may get one of
// the style classes .info, .warning, .error or .question, depending on the
// message type.
type InfoBar struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*InfoBar)(nil)
	_ coreglib.Objector = (*InfoBar)(nil)
)

func classInitInfoBarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkInfoBarClass)(unsafe.Pointer(gclassPtr))

	if _, ok := goval.(interface{ Close() }); ok {
		pclass.close = (*[0]byte)(C._gotk4_gtk3_InfoBarClass_close)
	}

	if _, ok := goval.(interface{ Response(responseId int) }); ok {
		pclass.response = (*[0]byte)(C._gotk4_gtk3_InfoBarClass_response)
	}
}

//export _gotk4_gtk3_InfoBarClass_close
func _gotk4_gtk3_InfoBarClass_close(arg0 *C.GtkInfoBar) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Close() })

	iface.Close()
}

//export _gotk4_gtk3_InfoBarClass_response
func _gotk4_gtk3_InfoBarClass_response(arg0 *C.GtkInfoBar, arg1 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Response(responseId int) })

	var _responseId int // out

	_responseId = int(arg1)

	iface.Response(_responseId)
}

func wrapInfoBar(obj *coreglib.Object) *InfoBar {
	return &InfoBar{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalInfoBar(p uintptr) (interface{}, error) {
	return wrapInfoBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_InfoBar_ConnectClose
func _gotk4_gtk3_InfoBar_ConnectClose(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClose signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user uses a keybinding to dismiss the info bar.
//
// The default binding for this signal is the Escape key.
func (infoBar *InfoBar) ConnectClose(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(infoBar, "close", false, unsafe.Pointer(C._gotk4_gtk3_InfoBar_ConnectClose), f)
}

//export _gotk4_gtk3_InfoBar_ConnectResponse
func _gotk4_gtk3_InfoBar_ConnectResponse(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(responseId int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(responseId int))
	}

	var _responseId int // out

	_responseId = int(arg1)

	f(_responseId)
}

// ConnectResponse is emitted when an action widget is clicked or the
// application programmer calls gtk_dialog_response(). The response_id depends
// on which action widget was clicked.
func (infoBar *InfoBar) ConnectResponse(f func(responseId int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(infoBar, "response", false, unsafe.Pointer(C._gotk4_gtk3_InfoBar_ConnectResponse), f)
}

// NewInfoBar creates a new InfoBar object.
//
// The function returns the following values:
//
//    - infoBar: new InfoBar object.
//
func NewInfoBar() *InfoBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_info_bar_new()

	var _infoBar *InfoBar // out

	_infoBar = wrapInfoBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _infoBar
}

// AddActionWidget: add an activatable widget to the action area of a InfoBar,
// connecting a signal handler that will emit the InfoBar::response signal on
// the message area when the widget is activated. The widget is appended to the
// end of the message areas action area.
//
// The function takes the following parameters:
//
//    - child: activatable widget.
//    - responseId: response ID for child.
//
func (infoBar *InfoBar) AddActionWidget(child Widgetter, responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	_arg2 = C.gint(responseId)

	C.gtk_info_bar_add_action_widget(_arg0, _arg1, _arg2)
	runtime.KeepAlive(infoBar)
	runtime.KeepAlive(child)
	runtime.KeepAlive(responseId)
}

// AddButton adds a button with the given text and sets things up so that
// clicking the button will emit the “response” signal with the given
// response_id. The button is appended to the end of the info bars's action
// area. The button widget is returned, but usually you don't need it.
//
// The function takes the following parameters:
//
//    - buttonText: text of button.
//    - responseId: response ID for the button.
//
// The function returns the following values:
//
//    - button widget that was added.
//
func (infoBar *InfoBar) AddButton(buttonText string, responseId int) *Button {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gint        // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(buttonText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(responseId)

	_cret = C.gtk_info_bar_add_button(_arg0, _arg1, _arg2)
	runtime.KeepAlive(infoBar)
	runtime.KeepAlive(buttonText)
	runtime.KeepAlive(responseId)

	var _button *Button // out

	_button = wrapButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _button
}

// ActionArea returns the action area of info_bar.
//
// The function returns the following values:
//
//    - box: action area.
//
func (infoBar *InfoBar) ActionArea() *Box {
	var _arg0 *C.GtkInfoBar // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))

	_cret = C.gtk_info_bar_get_action_area(_arg0)
	runtime.KeepAlive(infoBar)

	var _box *Box // out

	_box = wrapBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// ContentArea returns the content area of info_bar.
//
// The function returns the following values:
//
//    - box: content area.
//
func (infoBar *InfoBar) ContentArea() *Box {
	var _arg0 *C.GtkInfoBar // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))

	_cret = C.gtk_info_bar_get_content_area(_arg0)
	runtime.KeepAlive(infoBar)

	var _box *Box // out

	_box = wrapBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// MessageType returns the message type of the message area.
//
// The function returns the following values:
//
//    - messageType: message type of the message area.
//
func (infoBar *InfoBar) MessageType() MessageType {
	var _arg0 *C.GtkInfoBar    // out
	var _cret C.GtkMessageType // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))

	_cret = C.gtk_info_bar_get_message_type(_arg0)
	runtime.KeepAlive(infoBar)

	var _messageType MessageType // out

	_messageType = MessageType(_cret)

	return _messageType
}

// The function returns the following values:
//
//    - ok: current value of the GtkInfoBar:revealed property.
//
func (infoBar *InfoBar) Revealed() bool {
	var _arg0 *C.GtkInfoBar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))

	_cret = C.gtk_info_bar_get_revealed(_arg0)
	runtime.KeepAlive(infoBar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether the widget will display a standard close
// button.
//
// The function returns the following values:
//
//    - ok: TRUE if the widget displays standard close button.
//
func (infoBar *InfoBar) ShowCloseButton() bool {
	var _arg0 *C.GtkInfoBar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))

	_cret = C.gtk_info_bar_get_show_close_button(_arg0)
	runtime.KeepAlive(infoBar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Response emits the “response” signal with the given response_id.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//
func (infoBar *InfoBar) Response(responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))
	_arg1 = C.gint(responseId)

	C.gtk_info_bar_response(_arg0, _arg1)
	runtime.KeepAlive(infoBar)
	runtime.KeepAlive(responseId)
}

// SetDefaultResponse sets the last widget in the info bar’s action area with
// the given response_id as the default widget for the dialog. Pressing “Enter”
// normally activates the default widget.
//
// Note that this function currently requires info_bar to be added to a widget
// hierarchy.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//
func (infoBar *InfoBar) SetDefaultResponse(responseId int) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))
	_arg1 = C.gint(responseId)

	C.gtk_info_bar_set_default_response(_arg0, _arg1)
	runtime.KeepAlive(infoBar)
	runtime.KeepAlive(responseId)
}

// SetMessageType sets the message type of the message area.
//
// GTK+ uses this type to determine how the message is displayed.
//
// The function takes the following parameters:
//
//    - messageType: MessageType.
//
func (infoBar *InfoBar) SetMessageType(messageType MessageType) {
	var _arg0 *C.GtkInfoBar    // out
	var _arg1 C.GtkMessageType // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))
	_arg1 = C.GtkMessageType(messageType)

	C.gtk_info_bar_set_message_type(_arg0, _arg1)
	runtime.KeepAlive(infoBar)
	runtime.KeepAlive(messageType)
}

// SetResponseSensitive calls gtk_widget_set_sensitive (widget, setting) for
// each widget in the info bars’s action area with the given response_id. A
// convenient way to sensitize/desensitize dialog buttons.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//    - setting: TRUE for sensitive.
//
func (infoBar *InfoBar) SetResponseSensitive(responseId int, setting bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gint        // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))
	_arg1 = C.gint(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_info_bar_set_response_sensitive(_arg0, _arg1, _arg2)
	runtime.KeepAlive(infoBar)
	runtime.KeepAlive(responseId)
	runtime.KeepAlive(setting)
}

// SetRevealed sets the GtkInfoBar:revealed property to revealed. This will
// cause info_bar to show up with a slide-in transition.
//
// Note that this property does not automatically show info_bar and thus won’t
// have any effect if it is invisible.
//
// The function takes the following parameters:
//
//    - revealed: new value of the property.
//
func (infoBar *InfoBar) SetRevealed(revealed bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))
	if revealed {
		_arg1 = C.TRUE
	}

	C.gtk_info_bar_set_revealed(_arg0, _arg1)
	runtime.KeepAlive(infoBar)
	runtime.KeepAlive(revealed)
}

// SetShowCloseButton: if true, a standard close button is shown. When clicked
// it emits the response GTK_RESPONSE_CLOSE.
//
// The function takes the following parameters:
//
//    - setting: TRUE to include a close button.
//
func (infoBar *InfoBar) SetShowCloseButton(setting bool) {
	var _arg0 *C.GtkInfoBar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkInfoBar)(unsafe.Pointer(coreglib.InternObject(infoBar).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_info_bar_set_show_close_button(_arg0, _arg1)
	runtime.KeepAlive(infoBar)
	runtime.KeepAlive(setting)
}
