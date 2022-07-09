// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_StatusbarClass_text_popped(void*, guint, void*);
// extern void _gotk4_gtk3_StatusbarClass_text_pushed(void*, guint, void*);
// extern void _gotk4_gtk3_Statusbar_ConnectTextPopped(gpointer, guint, void*, guintptr);
// extern void _gotk4_gtk3_Statusbar_ConnectTextPushed(gpointer, guint, void*, guintptr);
import "C"

// GTypeStatusbar returns the GType for the type Statusbar.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeStatusbar() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Statusbar").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalStatusbar)
	return gtype
}

// StatusbarOverrider contains methods that are overridable.
type StatusbarOverrider interface {
	// The function takes the following parameters:
	//
	//    - contextId
	//    - text
	//
	TextPopped(contextId uint32, text string)
	// The function takes the following parameters:
	//
	//    - contextId
	//    - text
	//
	TextPushed(contextId uint32, text string)
}

// Statusbar is usually placed along the bottom of an application's main Window.
// It may provide a regular commentary of the application's status (as is
// usually the case in a web browser, for example), or may be used to simply
// output a message when the status changes, (when an upload is complete in an
// FTP client, for example).
//
// Status bars in GTK+ maintain a stack of messages. The message at the top of
// the each bar’s stack is the one that will currently be displayed.
//
// Any messages added to a statusbar’s stack must specify a context id that is
// used to uniquely identify the source of a message. This context id can be
// generated by gtk_statusbar_get_context_id(), given a message and the
// statusbar that it will be added to. Note that messages are stored in a stack,
// and when choosing which message to display, the stack structure is adhered
// to, regardless of the context identifier of a message.
//
// One could say that a statusbar maintains one stack of messages for display
// purposes, but allows multiple message producers to maintain sub-stacks of the
// messages they produced (via context ids).
//
// Status bars are created using gtk_statusbar_new().
//
// Messages are added to the bar’s stack with gtk_statusbar_push().
//
// The message at the top of the stack can be removed using gtk_statusbar_pop().
// A message can be removed from anywhere in the stack if its message id was
// recorded at the time it was added. This is done using gtk_statusbar_remove().
//
//
// CSS node
//
// GtkStatusbar has a single CSS node with name statusbar.
type Statusbar struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*Statusbar)(nil)
	_ coreglib.Objector = (*Statusbar)(nil)
)

func classInitStatusbarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "StatusbarClass")

	if _, ok := goval.(interface {
		TextPopped(contextId uint32, text string)
	}); ok {
		o := pclass.StructFieldOffset("text_popped")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_StatusbarClass_text_popped)
	}

	if _, ok := goval.(interface {
		TextPushed(contextId uint32, text string)
	}); ok {
		o := pclass.StructFieldOffset("text_pushed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_StatusbarClass_text_pushed)
	}
}

//export _gotk4_gtk3_StatusbarClass_text_popped
func _gotk4_gtk3_StatusbarClass_text_popped(arg0 *C.void, arg1 C.guint, arg2 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		TextPopped(contextId uint32, text string)
	})

	var _contextId uint32 // out
	var _text string      // out

	_contextId = uint32(arg1)
	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	iface.TextPopped(_contextId, _text)
}

//export _gotk4_gtk3_StatusbarClass_text_pushed
func _gotk4_gtk3_StatusbarClass_text_pushed(arg0 *C.void, arg1 C.guint, arg2 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		TextPushed(contextId uint32, text string)
	})

	var _contextId uint32 // out
	var _text string      // out

	_contextId = uint32(arg1)
	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	iface.TextPushed(_contextId, _text)
}

func wrapStatusbar(obj *coreglib.Object) *Statusbar {
	return &Statusbar{
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

func marshalStatusbar(p uintptr) (interface{}, error) {
	return wrapStatusbar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Statusbar_ConnectTextPopped
func _gotk4_gtk3_Statusbar_ConnectTextPopped(arg0 C.gpointer, arg1 C.guint, arg2 *C.void, arg3 C.guintptr) {
	var f func(contextId uint32, text string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(contextId uint32, text string))
	}

	var _contextId uint32 // out
	var _text string      // out

	_contextId = uint32(arg1)
	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_contextId, _text)
}

// ConnectTextPopped is emitted whenever a new message is popped off a
// statusbar's stack.
func (statusbar *Statusbar) ConnectTextPopped(f func(contextId uint32, text string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(statusbar, "text-popped", false, unsafe.Pointer(C._gotk4_gtk3_Statusbar_ConnectTextPopped), f)
}

//export _gotk4_gtk3_Statusbar_ConnectTextPushed
func _gotk4_gtk3_Statusbar_ConnectTextPushed(arg0 C.gpointer, arg1 C.guint, arg2 *C.void, arg3 C.guintptr) {
	var f func(contextId uint32, text string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(contextId uint32, text string))
	}

	var _contextId uint32 // out
	var _text string      // out

	_contextId = uint32(arg1)
	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_contextId, _text)
}

// ConnectTextPushed is emitted whenever a new message gets pushed onto a
// statusbar's stack.
func (statusbar *Statusbar) ConnectTextPushed(f func(contextId uint32, text string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(statusbar, "text-pushed", false, unsafe.Pointer(C._gotk4_gtk3_Statusbar_ConnectTextPushed), f)
}

// NewStatusbar creates a new Statusbar ready for messages.
//
// The function returns the following values:
//
//    - statusbar: new Statusbar.
//
func NewStatusbar() *Statusbar {
	_info := girepository.MustFind("Gtk", "Statusbar")
	_gret := _info.InvokeClassMethod("new_Statusbar", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _statusbar *Statusbar // out

	_statusbar = wrapStatusbar(coreglib.Take(unsafe.Pointer(_cret)))

	return _statusbar
}

// ContextID returns a new context identifier, given a description of the actual
// context. Note that the description is not shown in the UI.
//
// The function takes the following parameters:
//
//    - contextDescription: textual description of what context the new message
//      is being used in.
//
// The function returns the following values:
//
//    - guint: integer id.
//
func (statusbar *Statusbar) ContextID(contextDescription string) uint32 {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(statusbar).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(contextDescription)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "Statusbar")
	_gret := _info.InvokeClassMethod("get_context_id", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextDescription)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// MessageArea retrieves the box containing the label widget.
//
// The function returns the following values:
//
//    - box: Box.
//
func (statusbar *Statusbar) MessageArea() *Box {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(statusbar).Native()))

	_info := girepository.MustFind("Gtk", "Statusbar")
	_gret := _info.InvokeClassMethod("get_message_area", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(statusbar)

	var _box *Box // out

	_box = wrapBox(coreglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// Pop removes the first message in the Statusbar’s stack with the given context
// id.
//
// Note that this may not change the displayed message, if the message at the
// top of the stack has a different context id.
//
// The function takes the following parameters:
//
//    - contextId: context identifier.
//
func (statusbar *Statusbar) Pop(contextId uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(statusbar).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(contextId)

	_info := girepository.MustFind("Gtk", "Statusbar")
	_info.InvokeClassMethod("pop", _args[:], nil)

	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextId)
}

// Push pushes a new message onto a statusbar’s stack.
//
// The function takes the following parameters:
//
//    - contextId message’s context id, as returned by
//      gtk_statusbar_get_context_id().
//    - text: message to add to the statusbar.
//
// The function returns the following values:
//
//    - guint: message id that can be used with gtk_statusbar_remove().
//
func (statusbar *Statusbar) Push(contextId uint32, text string) uint32 {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(statusbar).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(contextId)
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[2]))

	_info := girepository.MustFind("Gtk", "Statusbar")
	_gret := _info.InvokeClassMethod("push", _args[:], nil)
	_cret := *(*C.guint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextId)
	runtime.KeepAlive(text)

	var _guint uint32 // out

	_guint = uint32(*(*C.guint)(unsafe.Pointer(&_cret)))

	return _guint
}

// Remove forces the removal of a message from a statusbar’s stack. The exact
// context_id and message_id must be specified.
//
// The function takes the following parameters:
//
//    - contextId: context identifier.
//    - messageId: message identifier, as returned by gtk_statusbar_push().
//
func (statusbar *Statusbar) Remove(contextId, messageId uint32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(statusbar).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(contextId)
	*(*C.guint)(unsafe.Pointer(&_args[2])) = C.guint(messageId)

	_info := girepository.MustFind("Gtk", "Statusbar")
	_info.InvokeClassMethod("remove", _args[:], nil)

	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextId)
	runtime.KeepAlive(messageId)
}

// RemoveAll forces the removal of all messages from a statusbar's stack with
// the exact context_id.
//
// The function takes the following parameters:
//
//    - contextId: context identifier.
//
func (statusbar *Statusbar) RemoveAll(contextId uint32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(statusbar).Native()))
	*(*C.guint)(unsafe.Pointer(&_args[1])) = C.guint(contextId)

	_info := girepository.MustFind("Gtk", "Statusbar")
	_info.InvokeClassMethod("remove_all", _args[:], nil)

	runtime.KeepAlive(statusbar)
	runtime.KeepAlive(contextId)
}
