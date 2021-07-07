// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_input_stream_get_type()), F: marshalFileInputStream},
	})
}

// FileInputStreamOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileInputStreamOverrider interface {
	CanSeek() bool
	// QueryInfo queries a file input stream the given @attributes. This
	// function blocks while querying the stream. For the asynchronous
	// (non-blocking) version of this function, see
	// g_file_input_stream_query_info_async(). While the stream is blocked, the
	// stream will set the pending flag internally, and any other operations on
	// the stream will fail with G_IO_ERROR_PENDING.
	QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error)
	// QueryInfoAsync queries the stream information asynchronously. When the
	// operation is finished @callback will be called. You can then call
	// g_file_input_stream_query_info_finish() to get the result of the
	// operation.
	//
	// For the synchronous version of this function, see
	// g_file_input_stream_query_info().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// QueryInfoFinish finishes an asynchronous info query operation.
	QueryInfoFinish(result AsyncResult) (FileInfo, error)
	Seek(offset int64, typ glib.SeekType, cancellable Cancellable) error
	Tell() int64
}

// FileInputStream provides input streams that take their content from a file.
//
// GFileInputStream implements #GSeekable, which allows the input stream to jump
// to arbitrary positions in the file, provided the filesystem of the file
// allows it. To find the position of a file input stream, use
// g_seekable_tell(). To find out if a file input stream supports seeking, use
// g_seekable_can_seek(). To position a file input stream, use
// g_seekable_seek().
type FileInputStream interface {
	InputStream

	// AsInputStream casts the class to the InputStream interface.
	AsInputStream() InputStream
	// AsSeekable casts the class to the Seekable interface.
	AsSeekable() Seekable

	// ClearPending clears the pending flag on @stream.
	//
	// This method is inherited from InputStream
	ClearPending()
	// Close closes the stream, releasing resources related to it.
	//
	// Once the stream is closed, all other operations will return
	// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an
	// error.
	//
	// Streams will be automatically closed when the last reference is dropped,
	// but you might want to call this function to make sure resources are
	// released as early as possible.
	//
	// Some streams might keep the backing store of the stream (e.g. a file
	// descriptor) open after the stream is closed. See the documentation for
	// the individual stream for details.
	//
	// On failure the first error that happened will be reported, but the close
	// operation will finish as much as possible. A stream that failed to close
	// will still return G_IO_ERROR_CLOSED for all operations. Still, it is
	// important to check and report the error to the user.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	// Cancelling a close will still leave the stream closed, but some streams
	// can use a faster close that doesn't block to e.g. check errors.
	//
	// This method is inherited from InputStream
	Close(cancellable Cancellable) error
	// CloseAsync requests an asynchronous closes of the stream, releasing
	// resources related to it. When the operation is finished @callback will be
	// called. You can then call g_input_stream_close_finish() to get the result
	// of the operation.
	//
	// For behaviour details see g_input_stream_close().
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	//
	// This method is inherited from InputStream
	CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// CloseFinish finishes closing a stream asynchronously, started from
	// g_input_stream_close_async().
	//
	// This method is inherited from InputStream
	CloseFinish(result AsyncResult) error
	// HasPending checks if an input stream has pending actions.
	//
	// This method is inherited from InputStream
	HasPending() bool
	// IsClosed checks if an input stream is closed.
	//
	// This method is inherited from InputStream
	IsClosed() bool
	// ReadAllFinish finishes an asynchronous stream read operation started with
	// g_input_stream_read_all_async().
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_read will be set to the number of bytes that were successfully
	// read before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_input_stream_read_async().
	//
	// This method is inherited from InputStream
	ReadAllFinish(result AsyncResult) (uint, error)
	// ReadBytesAsync: request an asynchronous read of @count bytes from the
	// stream into a new #GBytes. When the operation is finished @callback will
	// be called. You can then call g_input_stream_read_bytes_finish() to get
	// the result of the operation.
	//
	// During an async request no other sync and async calls are allowed on
	// @stream, and will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the new #GBytes will be passed to the callback. It is not an
	// error if this is smaller than the requested size, as it can happen e.g.
	// near the end of a file, but generally we try to read as many bytes as
	// requested. Zero is returned on end of file (or if @count is zero), but
	// never otherwise.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// This method is inherited from InputStream
	ReadBytesAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// ReadFinish finishes an asynchronous stream read operation.
	//
	// This method is inherited from InputStream
	ReadFinish(result AsyncResult) (int, error)
	// SetPending sets @stream to have actions pending. If the pending flag is
	// already set or @stream is closed, it will return false and set @error.
	//
	// This method is inherited from InputStream
	SetPending() error
	// Skip tries to skip @count bytes from the stream. Will block during the
	// operation.
	//
	// This is identical to g_input_stream_read(), from a behaviour standpoint,
	// but the bytes that are skipped are not returned to the user. Some streams
	// have an implementation that is more efficient than reading the data.
	//
	// This function is optional for inherited classes, as the default
	// implementation emulates it using read.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// This method is inherited from InputStream
	Skip(count uint, cancellable Cancellable) (int, error)
	// SkipAsync: request an asynchronous skip of @count bytes from the stream.
	// When the operation is finished @callback will be called. You can then
	// call g_input_stream_skip_finish() to get the result of the operation.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes skipped will be passed to the callback.
	// It is not an error if this is not the same as the requested size, as it
	// can happen e.g. near the end of a file, but generally we try to skip as
	// many bytes as requested. Zero is returned on end of file (or if @count is
	// zero), but never otherwise.
	//
	// Any outstanding i/o request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one, you must override all.
	//
	// This method is inherited from InputStream
	SkipAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// SkipFinish finishes a stream skip operation.
	//
	// This method is inherited from InputStream
	SkipFinish(result AsyncResult) (int, error)
	// CanSeek tests if the stream supports the Iface.
	//
	// This method is inherited from Seekable
	CanSeek() bool
	// CanTruncate tests if the length of the stream can be adjusted with
	// g_seekable_truncate().
	//
	// This method is inherited from Seekable
	CanTruncate() bool
	// Seek seeks in the stream by the given @offset, modified by @type.
	//
	// Attempting to seek past the end of the stream will have different results
	// depending on if the stream is fixed-sized or resizable. If the stream is
	// resizable then seeking past the end and then writing will result in zeros
	// filling the empty space. Seeking past the end of a resizable stream and
	// reading will result in EOF. Seeking past the end of a fixed-sized stream
	// will fail.
	//
	// Any operation that would result in a negative offset will fail.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	//
	// This method is inherited from Seekable
	Seek(offset int64, typ glib.SeekType, cancellable Cancellable) error
	// Tell tells the current position within the stream.
	//
	// This method is inherited from Seekable
	Tell() int64
	// Truncate sets the length of the stream to @offset. If the stream was
	// previously larger than @offset, the extra data is discarded. If the
	// stream was previously shorter than @offset, it is extended with NUL
	// ('\0') bytes.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned. If an
	// operation was partially finished when the operation was cancelled the
	// partial result will be returned, without an error.
	//
	// This method is inherited from Seekable
	Truncate(offset int64, cancellable Cancellable) error

	// QueryInfo queries a file input stream the given @attributes. This
	// function blocks while querying the stream. For the asynchronous
	// (non-blocking) version of this function, see
	// g_file_input_stream_query_info_async(). While the stream is blocked, the
	// stream will set the pending flag internally, and any other operations on
	// the stream will fail with G_IO_ERROR_PENDING.
	QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error)
	// QueryInfoAsync queries the stream information asynchronously. When the
	// operation is finished @callback will be called. You can then call
	// g_file_input_stream_query_info_finish() to get the result of the
	// operation.
	//
	// For the synchronous version of this function, see
	// g_file_input_stream_query_info().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// QueryInfoFinish finishes an asynchronous info query operation.
	QueryInfoFinish(result AsyncResult) (FileInfo, error)
}

// fileInputStream implements the FileInputStream interface.
type fileInputStream struct {
	*externglib.Object
}

var _ FileInputStream = (*fileInputStream)(nil)

// WrapFileInputStream wraps a GObject to a type that implements
// interface FileInputStream. It is primarily used internally.
func WrapFileInputStream(obj *externglib.Object) FileInputStream {
	return fileInputStream{obj}
}

func marshalFileInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileInputStream(obj), nil
}

func (f fileInputStream) AsInputStream() InputStream {
	return WrapInputStream(gextras.InternObject(f))
}

func (f fileInputStream) AsSeekable() Seekable {
	return WrapSeekable(gextras.InternObject(f))
}

func (s fileInputStream) ClearPending() {
	WrapInputStream(gextras.InternObject(s)).ClearPending()
}

func (s fileInputStream) Close(cancellable Cancellable) error {
	return WrapInputStream(gextras.InternObject(s)).Close(cancellable)
}

func (s fileInputStream) CloseAsync(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapInputStream(gextras.InternObject(s)).CloseAsync(ioPriority, cancellable, callback)
}

func (s fileInputStream) CloseFinish(result AsyncResult) error {
	return WrapInputStream(gextras.InternObject(s)).CloseFinish(result)
}

func (s fileInputStream) HasPending() bool {
	return WrapInputStream(gextras.InternObject(s)).HasPending()
}

func (s fileInputStream) IsClosed() bool {
	return WrapInputStream(gextras.InternObject(s)).IsClosed()
}

func (s fileInputStream) ReadAllFinish(result AsyncResult) (uint, error) {
	return WrapInputStream(gextras.InternObject(s)).ReadAllFinish(result)
}

func (s fileInputStream) ReadBytesAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapInputStream(gextras.InternObject(s)).ReadBytesAsync(count, ioPriority, cancellable, callback)
}

func (s fileInputStream) ReadFinish(result AsyncResult) (int, error) {
	return WrapInputStream(gextras.InternObject(s)).ReadFinish(result)
}

func (s fileInputStream) SetPending() error {
	return WrapInputStream(gextras.InternObject(s)).SetPending()
}

func (s fileInputStream) Skip(count uint, cancellable Cancellable) (int, error) {
	return WrapInputStream(gextras.InternObject(s)).Skip(count, cancellable)
}

func (s fileInputStream) SkipAsync(count uint, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	WrapInputStream(gextras.InternObject(s)).SkipAsync(count, ioPriority, cancellable, callback)
}

func (s fileInputStream) SkipFinish(result AsyncResult) (int, error) {
	return WrapInputStream(gextras.InternObject(s)).SkipFinish(result)
}

func (s fileInputStream) CanSeek() bool {
	return WrapSeekable(gextras.InternObject(s)).CanSeek()
}

func (s fileInputStream) CanTruncate() bool {
	return WrapSeekable(gextras.InternObject(s)).CanTruncate()
}

func (s fileInputStream) Seek(offset int64, typ glib.SeekType, cancellable Cancellable) error {
	return WrapSeekable(gextras.InternObject(s)).Seek(offset, typ, cancellable)
}

func (s fileInputStream) Tell() int64 {
	return WrapSeekable(gextras.InternObject(s)).Tell()
}

func (s fileInputStream) Truncate(offset int64, cancellable Cancellable) error {
	return WrapSeekable(gextras.InternObject(s)).Truncate(offset, cancellable)
}

func (s fileInputStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	var _arg0 *C.GFileInputStream // out
	var _arg1 *C.char             // out
	var _arg2 *C.GCancellable     // out
	var _cret *C.GFileInfo        // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_file_input_stream_query_info(_arg0, _arg1, _arg2, &_cerr)

	var _fileInfo FileInfo // out
	var _goerr error       // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}

func (s fileInputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GFileInputStream   // out
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_file_input_stream_query_info_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (s fileInputStream) QueryInfoFinish(result AsyncResult) (FileInfo, error) {
	var _arg0 *C.GFileInputStream // out
	var _arg1 *C.GAsyncResult     // out
	var _cret *C.GFileInfo        // in
	var _cerr *C.GError           // in

	_arg0 = (*C.GFileInputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_file_input_stream_query_info_finish(_arg0, _arg1, &_cerr)

	var _fileInfo FileInfo // out
	var _goerr error       // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _fileInfo, _goerr
}
