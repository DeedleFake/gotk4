// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// GtkWidget* _gotk4_gtk_dialog_new2(const gchar* title, GtkWindow* parent, GtkDialogFlags flags) {
// 	return gtk_dialog_new_with_buttons(title, parent, flags, NULL, NULL);
// }
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_response_type_get_type()), F: marshalResponseType},
		{T: externglib.Type(C.gtk_dialog_flags_get_type()), F: marshalDialogFlags},
		{T: externglib.Type(C.gtk_dialog_get_type()), F: marshalDialogger},
	})
}

// ResponseType: predefined values for use as response ids in
// gtk_dialog_add_button(). All predefined values are negative; GTK+ leaves
// values of 0 or greater for application-defined response ids.
type ResponseType C.gint

const (
	// ResponseNone: returned if an action widget has no response id, or if the
	// dialog gets programmatically hidden or destroyed.
	ResponseNone ResponseType = -1
	// ResponseReject: generic response id, not used by GTK+ dialogs.
	ResponseReject ResponseType = -2
	// ResponseAccept: generic response id, not used by GTK+ dialogs.
	ResponseAccept ResponseType = -3
	// ResponseDeleteEvent: returned if the dialog is deleted.
	ResponseDeleteEvent ResponseType = -4
	// ResponseOK: returned by OK buttons in GTK+ dialogs.
	ResponseOK ResponseType = -5
	// ResponseCancel: returned by Cancel buttons in GTK+ dialogs.
	ResponseCancel ResponseType = -6
	// ResponseClose: returned by Close buttons in GTK+ dialogs.
	ResponseClose ResponseType = -7
	// ResponseYes: returned by Yes buttons in GTK+ dialogs.
	ResponseYes ResponseType = -8
	// ResponseNo: returned by No buttons in GTK+ dialogs.
	ResponseNo ResponseType = -9
	// ResponseApply: returned by Apply buttons in GTK+ dialogs.
	ResponseApply ResponseType = -10
	// ResponseHelp: returned by Help buttons in GTK+ dialogs.
	ResponseHelp ResponseType = -11
)

func marshalResponseType(p uintptr) (interface{}, error) {
	return ResponseType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ResponseType.
func (r ResponseType) String() string {
	switch r {
	case ResponseNone:
		return "None"
	case ResponseReject:
		return "Reject"
	case ResponseAccept:
		return "Accept"
	case ResponseDeleteEvent:
		return "DeleteEvent"
	case ResponseOK:
		return "OK"
	case ResponseCancel:
		return "Cancel"
	case ResponseClose:
		return "Close"
	case ResponseYes:
		return "Yes"
	case ResponseNo:
		return "No"
	case ResponseApply:
		return "Apply"
	case ResponseHelp:
		return "Help"
	default:
		return fmt.Sprintf("ResponseType(%d)", r)
	}
}

// DialogFlags flags used to influence dialog construction.
type DialogFlags C.guint

const (
	// DialogModal: make the constructed dialog modal, see
	// gtk_window_set_modal().
	DialogModal DialogFlags = 0b1
	// DialogDestroyWithParent: destroy the dialog when its parent is destroyed,
	// see gtk_window_set_destroy_with_parent().
	DialogDestroyWithParent DialogFlags = 0b10
	// DialogUseHeaderBar: create dialog with actions in header bar instead of
	// action area. Since 3.12.
	DialogUseHeaderBar DialogFlags = 0b100
)

func marshalDialogFlags(p uintptr) (interface{}, error) {
	return DialogFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for DialogFlags.
func (d DialogFlags) String() string {
	if d == 0 {
		return "DialogFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(54)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case DialogModal:
			builder.WriteString("Modal|")
		case DialogDestroyWithParent:
			builder.WriteString("DestroyWithParent|")
		case DialogUseHeaderBar:
			builder.WriteString("UseHeaderBar|")
		default:
			builder.WriteString(fmt.Sprintf("DialogFlags(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DialogFlags) Has(other DialogFlags) bool {
	return (d & other) == other
}

// AlternativeDialogButtonOrder returns TRUE if dialogs are expected to use an
// alternative button order on the screen screen. See
// gtk_dialog_set_alternative_button_order() for more details about alternative
// button order.
//
// If you need to use this function, you should probably connect to the
// ::notify:gtk-alternative-button-order signal on the Settings object
// associated to screen, in order to be notified if the button order setting
// changes.
//
// Deprecated: Deprecated.
//
// The function takes the following parameters:
//
//    - screen (optional) or NULL to use the default screen.
//
// The function returns the following values:
//
//    - ok: whether the alternative button order should be used.
//
func AlternativeDialogButtonOrder(screen *gdk.Screen) bool {
	var _arg1 *C.GdkScreen // out
	var _cret C.gboolean   // in

	if screen != nil {
		_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	}

	_cret = C.gtk_alternative_dialog_button_order(_arg1)
	runtime.KeepAlive(screen)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DialogOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DialogOverrider interface {
	Close()
	// Response emits the Dialog::response signal with the given response ID.
	// Used to indicate that the user has responded to the dialog in some way;
	// typically either you or gtk_dialog_run() will be monitoring the
	// ::response signal and take appropriate action.
	//
	// The function takes the following parameters:
	//
	//    - responseId: response ID.
	//
	Response(responseId int)
}

// Dialog boxes are a convenient way to prompt the user for a small amount of
// input, e.g. to display a message, ask a question, or anything else that does
// not require extensive effort on the user’s part.
//
// GTK+ treats a dialog as a window split vertically. The top section is a VBox,
// and is where widgets such as a Label or a Entry should be packed. The bottom
// area is known as the “action area”. This is generally used for packing
// buttons into the dialog which may perform functions such as cancel, ok, or
// apply.
//
// Dialog boxes are created with a call to gtk_dialog_new() or
// gtk_dialog_new_with_buttons(). gtk_dialog_new_with_buttons() is recommended;
// it allows you to set the dialog title, some convenient flags, and add simple
// buttons.
//
// If “dialog” is a newly created dialog, the two primary areas of the window
// can be accessed through gtk_dialog_get_content_area() and
// gtk_dialog_get_action_area(), as can be seen from the example below.
//
// A “modal” dialog (that is, one which freezes the rest of the application from
// user input), can be created by calling gtk_window_set_modal() on the dialog.
// Use the GTK_WINDOW() macro to cast the widget returned from gtk_dialog_new()
// into a Window. When using gtk_dialog_new_with_buttons() you can also pass the
// K_DIALOG_MODAL flag to make a dialog modal.
//
// If you add buttons to Dialog using gtk_dialog_new_with_buttons(),
// gtk_dialog_add_button(), gtk_dialog_add_buttons(), or
// gtk_dialog_add_action_widget(), clicking the button will emit a signal called
// Dialog::response with a response ID that you specified. GTK+ will never
// assign a meaning to positive response IDs; these are entirely user-defined.
// But for convenience, you can use the response IDs in the ResponseType
// enumeration (these all have values less than zero). If a dialog receives a
// delete event, the Dialog::response signal will be emitted with a response ID
// of K_RESPONSE_DELETE_EVENT.
//
// If you want to block waiting for a dialog to return before returning control
// flow to your code, you can call gtk_dialog_run(). This function enters a
// recursive main loop and waits for the user to respond to the dialog,
// returning the response ID corresponding to the button the user clicked.
//
// For the simple dialog in the following example, in reality you’d probably use
// MessageDialog to save yourself some effort. But you’d need to create the
// dialog contents manually if you had more than a simple message in the dialog.
//
// An example for simple GtkDialog usage:
//
//    // Function to open a dialog box with a message
//    void
//    quick_message (GtkWindow *parent, gchar *message)
//    {
//     GtkWidget *dialog, *label, *content_area;
//     GtkDialogFlags flags;
//
//     // Create the widgets
//     flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//     dialog = gtk_dialog_new_with_buttons ("Message",
//                                           parent,
//                                           flags,
//                                           _("_OK"),
//                                           GTK_RESPONSE_NONE,
//                                           NULL);
//     content_area = gtk_dialog_get_content_area (GTK_DIALOG (dialog));
//     label = gtk_label_new (message);
//
//     // Ensure that the dialog box is destroyed when the user responds
//
//     g_signal_connect_swapped (dialog,
//                               "response",
//                               G_CALLBACK (gtk_widget_destroy),
//                               dialog);
//
//     // Add the label, and show everything we’ve added
//
//     gtk_container_add (GTK_CONTAINER (content_area), label);
//     gtk_widget_show_all (dialog);
//    }
//
//
// GtkDialog as GtkBuildable
//
// The GtkDialog implementation of the Buildable interface exposes the vbox and
// action_area as internal children with the names “vbox” and “action_area”.
//
// GtkDialog supports a custom <action-widgets> element, which can contain
// multiple <action-widget> elements. The “response” attribute specifies a
// numeric response, and the content of the element is the id of widget (which
// should be a child of the dialogs action_area). To mark a response as default,
// set the “default“ attribute of the <action-widget> element to true.
//
// GtkDialog supports adding action widgets by specifying “action“ as the “type“
// attribute of a <child> element. The widget will be added either to the action
// area or the headerbar of the dialog, depending on the “use-header-bar“
// property. The response id has to be associated with the action widget using
// the <action-widgets> element.
//
// An example of a Dialog UI definition fragment:
//
//    <object class="GtkDialog" id="dialog1">
//      <child type="action">
//        <object class="GtkButton" id="button_cancel"/>
//      </child>
//      <child type="action">
//        <object class="GtkButton" id="button_ok">
//          <property name="can-default">True</property>
//        </object>
//      </child>
//      <action-widgets>
//        <action-widget response="cancel">button_cancel</action-widget>
//        <action-widget response="ok" default="true">button_ok</action-widget>
//      </action-widgets>
//    </object>.
type Dialog struct {
	Window
}

var (
	_ Binner = (*Dialog)(nil)
)

func wrapDialog(obj *externglib.Object) *Dialog {
	return &Dialog{
		Window: Window{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
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
			},
		},
	}
}

func marshalDialogger(p uintptr) (interface{}, error) {
	return wrapDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClose signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user uses a keybinding to close the dialog.
//
// The default binding for this signal is the Escape key.
func (dialog *Dialog) ConnectClose(f func()) externglib.SignalHandle {
	return dialog.Connect("close", f)
}

// ConnectResponse: emitted when an action widget is clicked, the dialog
// receives a delete event, or the application programmer calls
// gtk_dialog_response(). On a delete event, the response ID is
// K_RESPONSE_DELETE_EVENT. Otherwise, it depends on which action widget was
// clicked.
func (dialog *Dialog) ConnectResponse(f func(responseId int)) externglib.SignalHandle {
	return dialog.Connect("response", f)
}

// NewDialog creates a new dialog box.
//
// Widgets should not be packed into this Window directly, but into the vbox and
// action_area, as described above.
//
// The function returns the following values:
//
//    - dialog: new dialog as a Widget.
//
func NewDialog() *Dialog {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_dialog_new()

	var _dialog *Dialog // out

	_dialog = wrapDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _dialog
}

// AddActionWidget adds an activatable widget to the action area of a Dialog,
// connecting a signal handler that will emit the Dialog::response signal on the
// dialog when the widget is activated. The widget is appended to the end of the
// dialog’s action area. If you want to add a non-activatable widget, simply
// pack it into the action_area field of the Dialog struct.
//
// The function takes the following parameters:
//
//    - child: activatable widget.
//    - responseId: response ID for child.
//
func (dialog *Dialog) AddActionWidget(child Widgetter, responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(responseId)

	C.gtk_dialog_add_action_widget(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(child)
	runtime.KeepAlive(responseId)
}

// AddButton adds a button with the given text and sets things up so that
// clicking the button will emit the Dialog::response signal with the given
// response_id. The button is appended to the end of the dialog’s action area.
// The button widget is returned, but usually you don’t need it.
//
// The function takes the following parameters:
//
//    - buttonText: text of button.
//    - responseId: response ID for the button.
//
// The function returns the following values:
//
//    - widget widget that was added.
//
func (dialog *Dialog) AddButton(buttonText string, responseId int) Widgetter {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gint       // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(buttonText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(responseId)

	_cret = C.gtk_dialog_add_button(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(buttonText)
	runtime.KeepAlive(responseId)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// ActionArea returns the action area of dialog.
//
// Deprecated: Direct access to the action area is discouraged; use
// gtk_dialog_add_button(), etc.
//
// The function returns the following values:
//
//    - box: action area.
//
func (dialog *Dialog) ActionArea() *Box {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_dialog_get_action_area(_arg0)
	runtime.KeepAlive(dialog)

	var _box *Box // out

	_box = wrapBox(externglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// ContentArea returns the content area of dialog.
//
// The function returns the following values:
//
//    - box: content area Box.
//
func (dialog *Dialog) ContentArea() *Box {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_dialog_get_content_area(_arg0)
	runtime.KeepAlive(dialog)

	var _box *Box // out

	_box = wrapBox(externglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// HeaderBar returns the header bar of dialog. Note that the headerbar is only
// used by the dialog if the Dialog:use-header-bar property is TRUE.
//
// The function returns the following values:
//
//    - headerBar: header bar.
//
func (dialog *Dialog) HeaderBar() *HeaderBar {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_dialog_get_header_bar(_arg0)
	runtime.KeepAlive(dialog)

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(externglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// ResponseForWidget gets the response id of a widget in the action area of a
// dialog.
//
// The function takes the following parameters:
//
//    - widget in the action area of dialog.
//
// The function returns the following values:
//
//    - gint: response id of widget, or GTK_RESPONSE_NONE if widget doesn’t have
//      a response id set.
//
func (dialog *Dialog) ResponseForWidget(widget Widgetter) int {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _cret C.gint       // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_dialog_get_response_for_widget(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(widget)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WidgetForResponse gets the widget button that uses the given response ID in
// the action area of a dialog.
//
// The function takes the following parameters:
//
//    - responseId: response ID used by the dialog widget.
//
// The function returns the following values:
//
//    - widget (optional) button that uses the given response_id, or NULL.
//
func (dialog *Dialog) WidgetForResponse(responseId int) Widgetter {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.gint(responseId)

	_cret = C.gtk_dialog_get_widget_for_response(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Response emits the Dialog::response signal with the given response ID. Used
// to indicate that the user has responded to the dialog in some way; typically
// either you or gtk_dialog_run() will be monitoring the ::response signal and
// take appropriate action.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//
func (dialog *Dialog) Response(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.gint(responseId)

	C.gtk_dialog_response(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)
}

// Run blocks in a recursive main loop until the dialog either emits the
// Dialog::response signal, or is destroyed. If the dialog is destroyed during
// the call to gtk_dialog_run(), gtk_dialog_run() returns K_RESPONSE_NONE.
// Otherwise, it returns the response ID from the ::response signal emission.
//
// Before entering the recursive main loop, gtk_dialog_run() calls
// gtk_widget_show() on the dialog for you. Note that you still need to show any
// children of the dialog yourself.
//
// During gtk_dialog_run(), the default behavior of Widget::delete-event is
// disabled; if the dialog receives ::delete_event, it will not be destroyed as
// windows usually are, and gtk_dialog_run() will return
// K_RESPONSE_DELETE_EVENT. Also, during gtk_dialog_run() the dialog will be
// modal. You can force gtk_dialog_run() to return at any time by calling
// gtk_dialog_response() to emit the ::response signal. Destroying the dialog
// during gtk_dialog_run() is a very bad idea, because your post-run code won’t
// know whether the dialog was destroyed or not.
//
// After gtk_dialog_run() returns, you are responsible for hiding or destroying
// the dialog if you wish to do so.
//
// Typical usage of this function might be:
//
//      GtkWidget *dialog = gtk_dialog_new ();
//      // Set up dialog...
//
//      int result = gtk_dialog_run (GTK_DIALOG (dialog));
//      switch (result)
//        {
//          case GTK_RESPONSE_ACCEPT:
//             // do_application_specific_something ();
//             break;
//          default:
//             // do_nothing_since_dialog_was_cancelled ();
//             break;
//        }
//      gtk_widget_destroy (dialog);
//
// Note that even though the recursive main loop gives the effect of a modal
// dialog (it prevents the user from interacting with other windows in the same
// window group while the dialog is run), callbacks such as timeouts, IO channel
// watches, DND drops, etc, will be triggered during a gtk_dialog_run() call.
//
// The function returns the following values:
//
//    - gint: response ID.
//
func (dialog *Dialog) Run() int {
	var _arg0 *C.GtkDialog // out
	var _cret C.gint       // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_dialog_run(_arg0)
	runtime.KeepAlive(dialog)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetAlternativeButtonOrderFromArray sets an alternative button order. If the
// Settings:gtk-alternative-button-order setting is set to TRUE, the dialog
// buttons are reordered according to the order of the response ids in
// new_order.
//
// See gtk_dialog_set_alternative_button_order() for more information.
//
// This function is for use by language bindings.
//
// Deprecated: Deprecated.
//
// The function takes the following parameters:
//
//    - newOrder: array of response ids of dialog’s buttons.
//
func (dialog *Dialog) SetAlternativeButtonOrderFromArray(newOrder []int) {
	var _arg0 *C.GtkDialog // out
	var _arg2 *C.gint      // out
	var _arg1 C.gint

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (C.gint)(len(newOrder))
	_arg2 = (*C.gint)(C.calloc(C.size_t(len(newOrder)), C.size_t(C.sizeof_gint)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((*C.gint)(_arg2), len(newOrder))
		for i := range newOrder {
			out[i] = C.gint(newOrder[i])
		}
	}

	C.gtk_dialog_set_alternative_button_order_from_array(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(newOrder)
}

// SetDefaultResponse sets the last widget in the dialog’s action area with the
// given response_id as the default widget for the dialog. Pressing “Enter”
// normally activates the default widget.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//
func (dialog *Dialog) SetDefaultResponse(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.gint(responseId)

	C.gtk_dialog_set_default_response(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)
}

// SetResponseSensitive calls gtk_widget_set_sensitive (widget, setting) for
// each widget in the dialog’s action area with the given response_id. A
// convenient way to sensitize/desensitize dialog buttons.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//    - setting: TRUE for sensitive.
//
func (dialog *Dialog) SetResponseSensitive(responseId int, setting bool) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.gint       // out
	var _arg2 C.gboolean   // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.gint(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_dialog_set_response_sensitive(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)
	runtime.KeepAlive(setting)
}

// NewDialogWithFlags is a slightly more advanced version of NewDialog,
// allowing the user to construct a new dialog with the given
// constructor-only dialog flags.
//
// It is a wrapper around Gtk.Dialog.new_with_buttons in C.
func NewDialogWithFlags(title string, parent *Window, flags DialogFlags) *Dialog {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	w := C._gotk4_gtk_dialog_new2(
		(*C.gchar)(unsafe.Pointer(ctitle)),
		(*C.GtkWindow)(unsafe.Pointer(parent.Native())),
		(C.GtkDialogFlags)(flags),
	)
	runtime.KeepAlive(parent)

	return wrapDialog(externglib.Take(unsafe.Pointer(w)))
}
