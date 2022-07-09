// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gio2_OutputStreamClass_close_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_OutputStreamClass_close_fn(void*, void*, GError**);
// extern gboolean _gotk4_gio2_OutputStreamClass_flush(void*, void*, GError**);
// extern gboolean _gotk4_gio2_OutputStreamClass_flush_finish(void*, void*, GError**);
// extern gboolean _gotk4_gio2_OutputStreamClass_writev_finish(void*, void*, void*, GError**);
// extern gssize _gotk4_gio2_OutputStreamClass_splice_finish(void*, void*, GError**);
// extern gssize _gotk4_gio2_OutputStreamClass_write_finish(void*, void*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
import "C"

// GTypeOutputStream returns the GType for the type OutputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeOutputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "OutputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalOutputStream)
	return gtype
}

// OutputStreamOverrider contains methods that are overridable.
type OutputStreamOverrider interface {
	// CloseFinish closes an output stream.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	CloseFinish(result AsyncResulter) error
	// The function takes the following parameters:
	//
	CloseFn(ctx context.Context) error
	// Flush forces a write of all user-space buffered data for the given
	// stream. Will block during the operation. Closing the stream will
	// implicitly cause a flush.
	//
	// This function is optional for inherited classes.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional cancellable object.
	//
	Flush(ctx context.Context) error
	// FlushFinish finishes flushing an output stream.
	//
	// The function takes the following parameters:
	//
	//    - result: GAsyncResult.
	//
	FlushFinish(result AsyncResulter) error
	// SpliceFinish finishes an asynchronous stream splice operation.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - gssize of the number of bytes spliced. Note that if the number of
	//      bytes spliced is greater than G_MAXSSIZE, then that will be returned,
	//      and there is no way to determine the actual number of bytes spliced.
	//
	SpliceFinish(result AsyncResulter) (int, error)
	// WriteFinish finishes a stream write operation.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - gssize containing the number of bytes written to the stream.
	//
	WriteFinish(result AsyncResulter) (int, error)
	// WritevFinish finishes a stream writev operation.
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - bytesWritten (optional): location to store the number of bytes that
	//      were written to the stream.
	//
	WritevFinish(result AsyncResulter) (uint, error)
}

// OutputStream has functions to write to a stream (g_output_stream_write()), to
// close a stream (g_output_stream_close()) and to flush pending writes
// (g_output_stream_flush()).
//
// To copy the content of an input stream to an output stream without manually
// handling the reads and writes, use g_output_stream_splice().
//
// See the documentation for OStream for details of thread safety of streaming
// APIs.
//
// All of these functions have async variants too.
type OutputStream struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*OutputStream)(nil)
)

// OutputStreamer describes types inherited from class OutputStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type OutputStreamer interface {
	coreglib.Objector
	baseOutputStream() *OutputStream
}

var _ OutputStreamer = (*OutputStream)(nil)

func classInitOutputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gio", "OutputStreamClass")

	if _, ok := goval.(interface {
		CloseFinish(result AsyncResulter) error
	}); ok {
		o := pclass.StructFieldOffset("close_finish")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_OutputStreamClass_close_finish)
	}

	if _, ok := goval.(interface {
		CloseFn(ctx context.Context) error
	}); ok {
		o := pclass.StructFieldOffset("close_fn")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_OutputStreamClass_close_fn)
	}

	if _, ok := goval.(interface {
		Flush(ctx context.Context) error
	}); ok {
		o := pclass.StructFieldOffset("flush")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_OutputStreamClass_flush)
	}

	if _, ok := goval.(interface {
		FlushFinish(result AsyncResulter) error
	}); ok {
		o := pclass.StructFieldOffset("flush_finish")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_OutputStreamClass_flush_finish)
	}

	if _, ok := goval.(interface {
		SpliceFinish(result AsyncResulter) (int, error)
	}); ok {
		o := pclass.StructFieldOffset("splice_finish")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_OutputStreamClass_splice_finish)
	}

	if _, ok := goval.(interface {
		WriteFinish(result AsyncResulter) (int, error)
	}); ok {
		o := pclass.StructFieldOffset("write_finish")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_OutputStreamClass_write_finish)
	}

	if _, ok := goval.(interface {
		WritevFinish(result AsyncResulter) (uint, error)
	}); ok {
		o := pclass.StructFieldOffset("writev_finish")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_OutputStreamClass_writev_finish)
	}
}

//export _gotk4_gio2_OutputStreamClass_close_finish
func _gotk4_gio2_OutputStreamClass_close_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CloseFinish(result AsyncResulter) error
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.CloseFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_OutputStreamClass_close_fn
func _gotk4_gio2_OutputStreamClass_close_fn(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CloseFn(ctx context.Context) error
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.CloseFn(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_OutputStreamClass_flush
func _gotk4_gio2_OutputStreamClass_flush(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Flush(ctx context.Context) error
	})

	var _cancellable context.Context // out

	if arg1 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg1))
	}

	_goerr := iface.Flush(_cancellable)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_OutputStreamClass_flush_finish
func _gotk4_gio2_OutputStreamClass_flush_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		FlushFinish(result AsyncResulter) error
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := iface.FlushFinish(_result)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_OutputStreamClass_splice_finish
func _gotk4_gio2_OutputStreamClass_splice_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		SpliceFinish(result AsyncResulter) (int, error)
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	gssize, _goerr := iface.SpliceFinish(_result)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_OutputStreamClass_write_finish
func _gotk4_gio2_OutputStreamClass_write_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret C.gssize) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		WriteFinish(result AsyncResulter) (int, error)
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	gssize, _goerr := iface.WriteFinish(_result)

	cret = C.gssize(gssize)
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_OutputStreamClass_writev_finish
func _gotk4_gio2_OutputStreamClass_writev_finish(arg0 *C.void, arg1 *C.void, arg2 *C.void, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		WritevFinish(result AsyncResulter) (uint, error)
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	bytesWritten, _goerr := iface.WritevFinish(_result)

	*arg2 = (*C.void)(unsafe.Pointer(bytesWritten))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapOutputStream(obj *coreglib.Object) *OutputStream {
	return &OutputStream{
		Object: obj,
	}
}

func marshalOutputStream(p uintptr) (interface{}, error) {
	return wrapOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (stream *OutputStream) baseOutputStream() *OutputStream {
	return stream
}

// BaseOutputStream returns the underlying base object.
func BaseOutputStream(obj OutputStreamer) *OutputStream {
	return obj.baseOutputStream()
}

// ClearPending clears the pending flag on stream.
func (stream *OutputStream) ClearPending() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("clear_pending", _args[:], nil)

	runtime.KeepAlive(stream)
}

// Close closes the stream, releasing resources related to it.
//
// Once the stream is closed, all other operations will return
// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an error.
//
// Closing a stream will automatically flush any outstanding buffers in the
// stream.
//
// Streams will be automatically closed when the last reference is dropped, but
// you might want to call this function to make sure resources are released as
// early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file
// descriptor) open after the stream is closed. See the documentation for the
// individual stream for details.
//
// On failure the first error that happened will be reported, but the close
// operation will finish as much as possible. A stream that failed to close will
// still return G_IO_ERROR_CLOSED for all operations. Still, it is important to
// check and report the error to the user, otherwise there might be a loss of
// data as all data might not be written.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned. Cancelling a close will
// still leave the stream closed, but there some streams can use a faster close
// that doesn't block to e.g. check errors. On cancellation (as with any error)
// there is no guarantee that all written data will reach the target.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellable object.
//
func (stream *OutputStream) Close(ctx context.Context) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("close", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CloseAsync requests an asynchronous close of the stream, releasing resources
// related to it. When the operation is finished callback will be called. You
// can then call g_output_stream_close_finish() to get the result of the
// operation.
//
// For behaviour details see g_output_stream_close().
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellable object.
//    - ioPriority: io priority of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *OutputStream) CloseAsync(ctx context.Context, ioPriority int32, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[4] = C.gpointer(gbox.AssignOnce(callback))
	}

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("close_async", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// CloseFinish closes an output stream.
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (stream *OutputStream) CloseFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("close_finish", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Flush forces a write of all user-space buffered data for the given stream.
// Will block during the operation. Closing the stream will implicitly cause a
// flush.
//
// This function is optional for inherited classes.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellable object.
//
func (stream *OutputStream) Flush(ctx context.Context) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[1] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("flush", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// FlushAsync forces an asynchronous write of all user-space buffered data for
// the given stream. For behaviour details see g_output_stream_flush().
//
// When the operation is finished callback will be called. You can then call
// g_output_stream_flush_finish() to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - ioPriority: io priority of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *OutputStream) FlushAsync(ctx context.Context, ioPriority int32, callback AsyncReadyCallback) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[3])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[4] = C.gpointer(gbox.AssignOnce(callback))
	}

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("flush_async", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// FlushFinish finishes flushing an output stream.
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
func (stream *OutputStream) FlushFinish(result AsyncResulter) error {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("flush_finish", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// HasPending checks if an output stream has pending actions.
//
// The function returns the following values:
//
//    - ok: TRUE if stream has pending actions.
//
func (stream *OutputStream) HasPending() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_gret := _info.InvokeClassMethod("has_pending", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsClosed checks if an output stream has already been closed.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is closed. FALSE otherwise.
//
func (stream *OutputStream) IsClosed() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_gret := _info.InvokeClassMethod("is_closed", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsClosing checks if an output stream is being closed. This can be used inside
// e.g. a flush implementation to see if the flush (or other i/o operation) is
// called from within the closing operation.
//
// The function returns the following values:
//
//    - ok: TRUE if stream is being closed. FALSE otherwise.
//
func (stream *OutputStream) IsClosing() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_gret := _info.InvokeClassMethod("is_closing", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetPending sets stream to have actions pending. If the pending flag is
// already set or stream is closed, it will return FALSE and set error.
func (stream *OutputStream) SetPending() error {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("set_pending", _args[:], nil)

	runtime.KeepAlive(stream)

	var _goerr error // out

	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SpliceFinish finishes an asynchronous stream splice operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - gssize of the number of bytes spliced. Note that if the number of bytes
//      spliced is greater than G_MAXSSIZE, then that will be returned, and there
//      is no way to determine the actual number of bytes spliced.
//
func (stream *OutputStream) SpliceFinish(result AsyncResulter) (int, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_gret := _info.InvokeClassMethod("splice_finish", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// WriteAllFinish finishes an asynchronous stream write operation started with
// g_output_stream_write_all_async().
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns FALSE (and sets error) then bytes_written
// will be set to the number of bytes that were successfully written before the
// error was encountered. This functionality is only available from C. If you
// need it from another language then you must write your own loop around
// g_output_stream_write_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - bytesWritten (optional): location to store the number of bytes that was
//      written to the stream.
//
func (stream *OutputStream) WriteAllFinish(result AsyncResulter) (uint, error) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("write_all_finish", _args[:], _outs[:])

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _bytesWritten uint // out
	var _goerr error       // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_bytesWritten = *(*uint)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesWritten, _goerr
}

// WriteBytes: wrapper function for g_output_stream_write() which takes a
// #GBytes as input. This can be more convenient for use by language bindings or
// in other cases where the refcounted nature of #GBytes is helpful over a bare
// pointer interface.
//
// However, note that this function may still perform partial writes, just like
// g_output_stream_write(). If that occurs, to continue writing, you will need
// to create a new #GBytes containing just the remaining bytes, using
// g_bytes_new_from_bytes(). Passing the same #GBytes instance multiple times
// potentially can result in duplicated data in the output stream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellable object.
//    - bytes to write.
//
// The function returns the following values:
//
//    - gssize: number of bytes written, or -1 on error.
//
func (stream *OutputStream) WriteBytes(ctx context.Context, bytes *glib.Bytes) (int, error) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(bytes)))

	_info := girepository.MustFind("Gio", "OutputStream")
	_gret := _info.InvokeClassMethod("write_bytes", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(bytes)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// WriteBytesAsync: this function is similar to g_output_stream_write_async(),
// but takes a #GBytes as input. Due to the refcounted nature of #GBytes, this
// allows the stream to avoid taking a copy of the data.
//
// However, note that this function may still perform partial writes, just like
// g_output_stream_write_async(). If that occurs, to continue writing, you will
// need to create a new #GBytes containing just the remaining bytes, using
// g_bytes_new_from_bytes(). Passing the same #GBytes instance multiple times
// potentially can result in duplicated data in the output stream.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_write_bytes().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - bytes to write.
//    - ioPriority: io priority of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *OutputStream) WriteBytesAsync(ctx context.Context, bytes *glib.Bytes, ioPriority int32, callback AsyncReadyCallback) {
	var _args [6]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(bytes)))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[4])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[5] = C.gpointer(gbox.AssignOnce(callback))
	}

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("write_bytes_async", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(bytes)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// WriteBytesFinish finishes a stream write-from-#GBytes operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - gssize containing the number of bytes written to the stream.
//
func (stream *OutputStream) WriteBytesFinish(result AsyncResulter) (int, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_gret := _info.InvokeClassMethod("write_bytes_finish", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// WriteFinish finishes a stream write operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - gssize containing the number of bytes written to the stream.
//
func (stream *OutputStream) WriteFinish(result AsyncResulter) (int, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_gret := _info.InvokeClassMethod("write_finish", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(*(*C.gssize)(unsafe.Pointer(&_cret)))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// WritevAllFinish finishes an asynchronous stream write operation started with
// g_output_stream_writev_all_async().
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns FALSE (and sets error) then bytes_written
// will be set to the number of bytes that were successfully written before the
// error was encountered. This functionality is only available from C. If you
// need it from another language then you must write your own loop around
// g_output_stream_writev_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - bytesWritten (optional): location to store the number of bytes that were
//      written to the stream.
//
func (stream *OutputStream) WritevAllFinish(result AsyncResulter) (uint, error) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("writev_all_finish", _args[:], _outs[:])

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _bytesWritten uint // out
	var _goerr error       // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_bytesWritten = *(*uint)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesWritten, _goerr
}

// WritevFinish finishes a stream writev operation.
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - bytesWritten (optional): location to store the number of bytes that were
//      written to the stream.
//
func (stream *OutputStream) WritevFinish(result AsyncResulter) (uint, error) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "OutputStream")
	_info.InvokeClassMethod("writev_finish", _args[:], _outs[:])

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _bytesWritten uint // out
	var _goerr error       // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_bytesWritten = *(*uint)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesWritten, _goerr
}
