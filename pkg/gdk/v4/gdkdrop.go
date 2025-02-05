// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// extern void _gotk4_gdk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// GType values.
var (
	GTypeDrop = coreglib.Type(C.gdk_drop_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDrop, F: marshalDrop},
	})
}

// Drop: GdkDrop object represents the target of an ongoing DND operation.
//
// Possible drop sites get informed about the status of the ongoing drag
// operation with events of type GDK_DRAG_ENTER, GDK_DRAG_LEAVE, GDK_DRAG_MOTION
// and GDK_DROP_START. The GdkDrop object can be obtained from these gdk.Event
// types using gdk.DNDEvent.GetDrop().
//
// The actual data transfer is initiated from the target side via an async read,
// using one of the GdkDrop methods for this purpose: gdk.Drop.ReadAsync() or
// gdk.Drop.ReadValueAsync().
//
// GTK provides a higher level abstraction based on top of these functions, and
// so they are not normally needed in GTK applications. See the "Drag and Drop"
// section of the GTK documentation for more information.
type Drop struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Drop)(nil)
)

// Dropper describes types inherited from class Drop.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Dropper interface {
	coreglib.Objector
	baseDrop() *Drop
}

var _ Dropper = (*Drop)(nil)

func wrapDrop(obj *coreglib.Object) *Drop {
	return &Drop{
		Object: obj,
	}
}

func marshalDrop(p uintptr) (interface{}, error) {
	return wrapDrop(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *Drop) baseDrop() *Drop {
	return self
}

// BaseDrop returns the underlying base object.
func BaseDrop(obj Dropper) *Drop {
	return obj.baseDrop()
}

// Finish ends the drag operation after a drop.
//
// The action must be a single action selected from the actions available via
// gdk.Drop.GetActions().
//
// The function takes the following parameters:
//
//    - action performed by the destination or 0 if the drop failed.
//
func (self *Drop) Finish(action DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GdkDragAction(action)

	C.gdk_drop_finish(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(action)
}

// Actions returns the possible actions for this GdkDrop.
//
// If this value contains multiple actions - i.e. gdk.DragAction().IsUnique
// returns FALSE for the result - gdk.Drop.Finish() must choose the action to
// use when accepting the drop. This will only happen if you passed
// GDK_ACTION_ASK as one of the possible actions in gdk.Drop.Status().
// GDK_ACTION_ASK itself will not be included in the actions returned by this
// function.
//
// This value may change over the lifetime of the gdk.Drop both as a response to
// source side actions as well as to calls to gdk.Drop.Status() or
// gdk.Drop.Finish(). The source side will not change this value anymore once a
// drop has started.
//
// The function returns the following values:
//
//    - dragAction: possible GdkDragActions.
//
func (self *Drop) Actions() DragAction {
	var _arg0 *C.GdkDrop      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gdk_drop_get_actions(_arg0)
	runtime.KeepAlive(self)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// Device returns the GdkDevice performing the drop.
//
// The function returns the following values:
//
//    - device: GdkDevice performing the drop.
//
func (self *Drop) Device() Devicer {
	var _arg0 *C.GdkDrop   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gdk_drop_get_device(_arg0)
	runtime.KeepAlive(self)

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	return _device
}

// Display gets the GdkDisplay that self was created for.
//
// The function returns the following values:
//
//    - display: GdkDisplay.
//
func (self *Drop) Display() *Display {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gdk_drop_get_display(_arg0)
	runtime.KeepAlive(self)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Drag: if this is an in-app drag-and-drop operation, returns the GdkDrag that
// corresponds to this drop.
//
// If it is not, NULL is returned.
//
// The function returns the following values:
//
//    - drag (optional): corresponding GdkDrag.
//
func (self *Drop) Drag() Dragger {
	var _arg0 *C.GdkDrop // out
	var _cret *C.GdkDrag // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gdk_drop_get_drag(_arg0)
	runtime.KeepAlive(self)

	var _drag Dragger // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Dragger)
				return ok
			})
			rv, ok := casted.(Dragger)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dragger")
			}
			_drag = rv
		}
	}

	return _drag
}

// Formats returns the GdkContentFormats that the drop offers the data to be
// read in.
//
// The function returns the following values:
//
//    - contentFormats: possible GdkContentFormats.
//
func (self *Drop) Formats() *ContentFormats {
	var _arg0 *C.GdkDrop           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gdk_drop_get_formats(_arg0)
	runtime.KeepAlive(self)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// Surface returns the GdkSurface performing the drop.
//
// The function returns the following values:
//
//    - surface: GdkSurface performing the drop.
//
func (self *Drop) Surface() Surfacer {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gdk_drop_get_surface(_arg0)
	runtime.KeepAlive(self)

	var _surface Surfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Surfacer)
			return ok
		})
		rv, ok := casted.(Surfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
		}
		_surface = rv
	}

	return _surface
}

// ReadAsync: asynchronously read the dropped data from a GdkDrop in a format
// that complies with one of the mime types.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional GCancellable object, NULL to ignore.
//    - mimeTypes: pointer to an array of mime types.
//    - ioPriority: i/O priority for the read operation.
//    - callback (optional): GAsyncReadyCallback to call when the request is
//      satisfied.
//
func (self *Drop) ReadAsync(ctx context.Context, mimeTypes []string, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkDrop            // out
	var _arg3 *C.GCancellable       // out
	var _arg1 **C.char              // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	{
		_arg1 = (**C.char)(C.calloc(C.size_t((len(mimeTypes) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(mimeTypes)+1)
			var zero *C.char
			out[len(mimeTypes)] = zero
			for i := range mimeTypes {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(mimeTypes[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_drop_read_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(mimeTypes)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ReadFinish finishes an async drop read operation.
//
// Note that you must not use blocking read calls on the returned stream in the
// GTK thread, since some platforms might require communication with GTK to
// complete the data transfer. You can use async APIs such as
// g_input_stream_read_bytes_async().
//
// See gdk.Drop.ReadAsync().
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
// The function returns the following values:
//
//    - outMimeType: return location for the used mime type.
//    - inputStream (optional): GInputStream, or NULL.
//
func (self *Drop) ReadFinish(result gio.AsyncResulter) (string, gio.InputStreamer, error) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.char         // in
	var _cret *C.GInputStream // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.gdk_drop_read_finish(_arg0, _arg1, &_arg2, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _outMimeType string            // out
	var _inputStream gio.InputStreamer // out
	var _goerr error                   // out

	_outMimeType = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.InputStreamer)
				return ok
			})
			rv, ok := casted.(gio.InputStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
			}
			_inputStream = rv
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _outMimeType, _inputStream, _goerr
}

// ReadValueAsync: asynchronously request the drag operation's contents
// converted to the given type.
//
// When the operation is finished callback will be called. You must then call
// gdk.Drop.ReadValueFinish() to get the resulting GValue.
//
// For local drag'n'drop operations that are available in the given GType, the
// value will be copied directly. Otherwise, GDK will try to use
// gdk.ContentDeserializeAsync() to convert the data.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional GCancellable object, NULL to ignore.
//    - typ: GType to read.
//    - ioPriority: i/O priority of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (self *Drop) ReadValueAsync(ctx context.Context, typ coreglib.Type, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkDrop            // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GType               // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GType(typ)
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_drop_read_value_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ReadValueFinish finishes an async drop read.
//
// See gdk.Drop.ReadValueAsync().
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
// The function returns the following values:
//
//    - value: GValue containing the result.
//
func (self *Drop) ReadValueFinish(result gio.AsyncResulter) (*coreglib.Value, error) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GValue       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.gdk_drop_read_value_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _value *coreglib.Value // out
	var _goerr error           // out

	_value = coreglib.ValueFromNative(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _value, _goerr
}

// Status selects all actions that are potentially supported by the destination.
//
// When calling this function, do not restrict the passed in actions to the ones
// provided by gdk.Drop.GetActions(). Those actions may change in the future,
// even depending on the actions you provide here.
//
// The preferred action is a hint to the drag'n'drop mechanism about which
// action to use when multiple actions are possible.
//
// This function should be called by drag destinations in response to
// GDK_DRAG_ENTER or GDK_DRAG_MOTION events. If the destination does not yet
// know the exact actions it supports, it should set any possible actions first
// and then later call this function again.
//
// The function takes the following parameters:
//
//    - actions: supported actions of the destination, or 0 to indicate that a
//      drop will not be accepted.
//    - preferred: unique action that's a member of actions indicating the
//      preferred action.
//
func (self *Drop) Status(actions, preferred DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out
	var _arg2 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GdkDragAction(actions)
	_arg2 = C.GdkDragAction(preferred)

	C.gdk_drop_status(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(preferred)
}
