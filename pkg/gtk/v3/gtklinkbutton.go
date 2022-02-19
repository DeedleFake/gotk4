// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_LinkButtonClass_activate_link(GtkLinkButton*);
// extern gboolean _gotk4_gtk3_LinkButton_ConnectActivateLink(gpointer, guintptr);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_link_button_get_type()), F: marshalLinkButtonner},
	})
}

// LinkButtonOverrider contains methods that are overridable.
type LinkButtonOverrider interface {
	// The function returns the following values:
	//
	ActivateLink() bool
}

// LinkButton is a Button with a hyperlink, similar to the one used by web
// browsers, which triggers an action when clicked. It is useful to show quick
// links to resources.
//
// A link button is created by calling either gtk_link_button_new() or
// gtk_link_button_new_with_label(). If using the former, the URI you pass to
// the constructor is used as a label for the widget.
//
// The URI bound to a GtkLinkButton can be set specifically using
// gtk_link_button_set_uri(), and retrieved using gtk_link_button_get_uri().
//
// By default, GtkLinkButton calls gtk_show_uri_on_window() when the button is
// clicked. This behaviour can be overridden by connecting to the
// LinkButton::activate-link signal and returning TRUE from the signal handler.
//
//
// CSS nodes
//
// GtkLinkButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .link style class.
type LinkButton struct {
	_ [0]func() // equal guard
	Button
}

var (
	_ Binner              = (*LinkButton)(nil)
	_ externglib.Objector = (*LinkButton)(nil)
)

func classInitLinkButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkLinkButtonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkLinkButtonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ActivateLink() bool }); ok {
		pclass.activate_link = (*[0]byte)(C._gotk4_gtk3_LinkButtonClass_activate_link)
	}
}

//export _gotk4_gtk3_LinkButtonClass_activate_link
func _gotk4_gtk3_LinkButtonClass_activate_link(arg0 *C.GtkLinkButton) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ActivateLink() bool })

	ok := iface.ActivateLink()

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapLinkButton(obj *externglib.Object) *LinkButton {
	return &LinkButton{
		Button: Button{
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
			Object: obj,
			Actionable: Actionable{
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
			Activatable: Activatable{
				Object: obj,
			},
		},
	}
}

func marshalLinkButtonner(p uintptr) (interface{}, error) {
	return wrapLinkButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_LinkButton_ConnectActivateLink
func _gotk4_gtk3_LinkButton_ConnectActivateLink(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectActivateLink signal is emitted each time the LinkButton has been
// clicked.
//
// The default handler will call gtk_show_uri_on_window() with the URI stored
// inside the LinkButton:uri property.
//
// To override the default behavior, you can connect to the ::activate-link
// signal and stop the propagation of the signal by returning TRUE from your
// handler.
func (linkButton *LinkButton) ConnectActivateLink(f func() (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(linkButton, "activate-link", false, unsafe.Pointer(C._gotk4_gtk3_LinkButton_ConnectActivateLink), f)
}

// NewLinkButton creates a new LinkButton with the URI as its text.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
// The function returns the following values:
//
//    - linkButton: new link button widget.
//
func NewLinkButton(uri string) *LinkButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_link_button_new(_arg1)
	runtime.KeepAlive(uri)

	var _linkButton *LinkButton // out

	_linkButton = wrapLinkButton(externglib.Take(unsafe.Pointer(_cret)))

	return _linkButton
}

// NewLinkButtonWithLabel creates a new LinkButton containing a label.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//    - label (optional): text of the button.
//
// The function returns the following values:
//
//    - linkButton: new link button widget.
//
func NewLinkButtonWithLabel(uri, label string) *LinkButton {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_link_button_new_with_label(_arg1, _arg2)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(label)

	var _linkButton *LinkButton // out

	_linkButton = wrapLinkButton(externglib.Take(unsafe.Pointer(_cret)))

	return _linkButton
}

// URI retrieves the URI set using gtk_link_button_set_uri().
//
// The function returns the following values:
//
//    - utf8: valid URI. The returned string is owned by the link button and
//      should not be modified or freed.
//
func (linkButton *LinkButton) URI() string {
	var _arg0 *C.GtkLinkButton // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))

	_cret = C.gtk_link_button_get_uri(_arg0)
	runtime.KeepAlive(linkButton)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Visited retrieves the “visited” state of the URI where the LinkButton points.
// The button becomes visited when it is clicked. If the URI is changed on the
// button, the “visited” state is unset again.
//
// The state may also be changed using gtk_link_button_set_visited().
//
// The function returns the following values:
//
//    - ok: TRUE if the link has been visited, FALSE otherwise.
//
func (linkButton *LinkButton) Visited() bool {
	var _arg0 *C.GtkLinkButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))

	_cret = C.gtk_link_button_get_visited(_arg0)
	runtime.KeepAlive(linkButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetURI sets uri as the URI where the LinkButton points. As a side-effect this
// unsets the “visited” state of the button.
//
// The function takes the following parameters:
//
//    - uri: valid URI.
//
func (linkButton *LinkButton) SetURI(uri string) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_link_button_set_uri(_arg0, _arg1)
	runtime.KeepAlive(linkButton)
	runtime.KeepAlive(uri)
}

// SetVisited sets the “visited” state of the URI where the LinkButton points.
// See gtk_link_button_get_visited() for more details.
//
// The function takes the following parameters:
//
//    - visited: new “visited” state.
//
func (linkButton *LinkButton) SetVisited(visited bool) {
	var _arg0 *C.GtkLinkButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkLinkButton)(unsafe.Pointer(linkButton.Native()))
	if visited {
		_arg1 = C.TRUE
	}

	C.gtk_link_button_set_visited(_arg0, _arg1)
	runtime.KeepAlive(linkButton)
	runtime.KeepAlive(visited)
}
