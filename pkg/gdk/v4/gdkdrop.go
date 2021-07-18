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
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drop_get_type()), F: marshalDropper},
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
	*externglib.Object
}

var _ gextras.Nativer = (*Drop)(nil)

// Dropper describes Drop's abstract methods.
type Dropper interface {
	// Finish ends the drag operation after a drop.
	Finish(action DragAction)
	// Actions returns the possible actions for this GdkDrop.
	Actions() DragAction
	// Device returns the GdkDevice performing the drop.
	Device() Devicer
	// Display gets the GdkDisplay that self was created for.
	Display() *Display
	// Drag: if this is an in-app drag-and-drop operation, returns the GdkDrag
	// that corresponds to this drop.
	Drag() Dragger
	// Formats returns the GdkContentFormats that the drop offers the data to be
	// read in.
	Formats() *ContentFormats
	// Surface returns the GdkSurface performing the drop.
	Surface() Surfacer
	// ReadAsync: asynchronously read the dropped data from a GdkDrop in a
	// format that complies with one of the mime types.
	ReadAsync(ctx context.Context, mimeTypes []string, ioPriority int, callback gio.AsyncReadyCallback)
	// ReadFinish finishes an async drop read operation.
	ReadFinish(result gio.AsyncResulter) (string, gio.InputStreamer, error)
	// ReadValueAsync: asynchronously request the drag operation's contents
	// converted to the given type.
	ReadValueAsync(ctx context.Context, typ externglib.Type, ioPriority int, callback gio.AsyncReadyCallback)
	// ReadValueFinish finishes an async drop read.
	ReadValueFinish(result gio.AsyncResulter) (*externglib.Value, error)
	// Status selects all actions that are potentially supported by the
	// destination.
	Status(actions DragAction, preferred DragAction)
}

var _ Dropper = (*Drop)(nil)

func wrapDrop(obj *externglib.Object) *Drop {
	return &Drop{
		Object: obj,
	}
}

func marshalDropper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrop(obj), nil
}

// Finish ends the drag operation after a drop.
//
// The action must be a single action selected from the actions available via
// gdk.Drop.GetActions().
func (self *Drop) Finish(action DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))
	_arg1 = C.GdkDragAction(action)

	C.gdk_drop_finish(_arg0, _arg1)
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
func (self *Drop) Actions() DragAction {
	var _arg0 *C.GdkDrop      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_drop_get_actions(_arg0)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// Device returns the GdkDevice performing the drop.
func (self *Drop) Device() Devicer {
	var _arg0 *C.GdkDrop   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_drop_get_device(_arg0)

	var _device Devicer // out

	_device = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Devicer)

	return _device
}

// Display gets the GdkDisplay that self was created for.
func (self *Drop) Display() *Display {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_drop_get_display(_arg0)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Drag: if this is an in-app drag-and-drop operation, returns the GdkDrag that
// corresponds to this drop.
//
// If it is not, NULL is returned.
func (self *Drop) Drag() Dragger {
	var _arg0 *C.GdkDrop // out
	var _cret *C.GdkDrag // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_drop_get_drag(_arg0)

	var _drag Dragger // out

	_drag = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Dragger)

	return _drag
}

// Formats returns the GdkContentFormats that the drop offers the data to be
// read in.
func (self *Drop) Formats() *ContentFormats {
	var _arg0 *C.GdkDrop           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_drop_get_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _contentFormats
}

// Surface returns the GdkSurface performing the drop.
func (self *Drop) Surface() Surfacer {
	var _arg0 *C.GdkDrop    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))

	_cret = C.gdk_drop_get_surface(_arg0)

	var _surface Surfacer // out

	_surface = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Surfacer)

	return _surface
}

// ReadAsync: asynchronously read the dropped data from a GdkDrop in a format
// that complies with one of the mime types.
func (self *Drop) ReadAsync(ctx context.Context, mimeTypes []string, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkDrop      // out
	var _arg3 *C.GCancellable // out
	var _arg1 **C.char
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	{
		_arg1 = (**C.char)(C.malloc(C.ulong(len(mimeTypes)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(_arg1, len(mimeTypes)+1)
			var zero *C.char
			out[len(mimeTypes)] = zero
			for i := range mimeTypes {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(mimeTypes[i])))
			}
		}
	}
	_arg2 = C.int(ioPriority)
	_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.AssignOnce(callback))

	C.gdk_drop_read_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// ReadFinish finishes an async drop read operation.
//
// Note that you must not use blocking read calls on the returned stream in the
// GTK thread, since some platforms might require communication with GTK to
// complete the data transfer. You can use async APIs such as
// g_input_stream_read_bytes_async().
//
// See gdk.Drop.ReadAsync().
func (self *Drop) ReadFinish(result gio.AsyncResulter) (string, gio.InputStreamer, error) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.char         // in
	var _cret *C.GInputStream // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.gdk_drop_read_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _outMimeType string            // out
	var _inputStream gio.InputStreamer // out
	var _goerr error                   // out

	_outMimeType = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_inputStream = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gio.InputStreamer)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

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
func (self *Drop) ReadValueAsync(ctx context.Context, typ externglib.Type, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkDrop            // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GType               // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GType(typ)
	_arg2 = C.int(ioPriority)
	_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.AssignOnce(callback))

	C.gdk_drop_read_value_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// ReadValueFinish finishes an async drop read.
//
// See gdk.Drop.ReadValueAsync().
func (self *Drop) ReadValueFinish(result gio.AsyncResulter) (*externglib.Value, error) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GValue       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	_cret = C.gdk_drop_read_value_finish(_arg0, _arg1, &_cerr)

	var _value *externglib.Value // out
	var _goerr error             // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

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
func (self *Drop) Status(actions DragAction, preferred DragAction) {
	var _arg0 *C.GdkDrop      // out
	var _arg1 C.GdkDragAction // out
	var _arg2 C.GdkDragAction // out

	_arg0 = (*C.GdkDrop)(unsafe.Pointer(self.Native()))
	_arg1 = C.GdkDragAction(actions)
	_arg2 = C.GdkDragAction(preferred)

	C.gdk_drop_status(_arg0, _arg1, _arg2)
}
