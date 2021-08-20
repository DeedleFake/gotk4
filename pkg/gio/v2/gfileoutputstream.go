// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_output_stream_get_type()), F: marshalFileOutputStreamer},
	})
}

// FileOutputStreamOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileOutputStreamOverrider interface {
	CanSeek() bool
	CanTruncate() bool
	// ETag gets the entity tag for the file when it has been written. This must
	// be called after the stream has been written and closed, as the etag can
	// change while writing.
	ETag() string
	// QueryInfo queries a file output stream for the given attributes. This
	// function blocks while querying the stream. For the asynchronous version
	// of this function, see g_file_output_stream_query_info_async(). While the
	// stream is blocked, the stream will set the pending flag internally, and
	// any other operations on the stream will fail with G_IO_ERROR_PENDING.
	//
	// Can fail if the stream was already closed (with error being set to
	// G_IO_ERROR_CLOSED), the stream has pending operations (with error being
	// set to G_IO_ERROR_PENDING), or if querying info is not supported for the
	// stream's interface (with error being set to G_IO_ERROR_NOT_SUPPORTED). In
	// all cases of failure, NULL will be returned.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and NULL will
	// be returned.
	QueryInfo(ctx context.Context, attributes string) (*FileInfo, error)
	// QueryInfoAsync: asynchronously queries the stream for a Info. When
	// completed, callback will be called with a Result which can be used to
	// finish the operation with g_file_output_stream_query_info_finish().
	//
	// For the synchronous version of this function, see
	// g_file_output_stream_query_info().
	QueryInfoAsync(ctx context.Context, attributes string, ioPriority int, callback AsyncReadyCallback)
	// QueryInfoFinish finalizes the asynchronous query started by
	// g_file_output_stream_query_info_async().
	QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
	Seek(ctx context.Context, offset int64, typ glib.SeekType) error
	Tell() int64
	TruncateFn(ctx context.Context, size int64) error
}

// FileOutputStream provides output streams that write their content to a file.
//
// GFileOutputStream implements #GSeekable, which allows the output stream to
// jump to arbitrary positions in the file and to truncate the file, provided
// the filesystem of the file supports these operations.
//
// To find the position of a file output stream, use g_seekable_tell(). To find
// out if a file output stream supports seeking, use g_seekable_can_seek().To
// position a file output stream, use g_seekable_seek(). To find out if a file
// output stream supports truncating, use g_seekable_can_truncate(). To truncate
// a file output stream, use g_seekable_truncate().
type FileOutputStream struct {
	OutputStream

	Seekable
	*externglib.Object
}

func wrapFileOutputStream(obj *externglib.Object) *FileOutputStream {
	return &FileOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
		Seekable: Seekable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalFileOutputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileOutputStream(obj), nil
}

// ETag gets the entity tag for the file when it has been written. This must be
// called after the stream has been written and closed, as the etag can change
// while writing.
func (stream *FileOutputStream) ETag() string {
	var _arg0 *C.GFileOutputStream // out
	var _cret *C.char              // in

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_file_output_stream_get_etag(_arg0)
	runtime.KeepAlive(stream)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// QueryInfo queries a file output stream for the given attributes. This
// function blocks while querying the stream. For the asynchronous version of
// this function, see g_file_output_stream_query_info_async(). While the stream
// is blocked, the stream will set the pending flag internally, and any other
// operations on the stream will fail with G_IO_ERROR_PENDING.
//
// Can fail if the stream was already closed (with error being set to
// G_IO_ERROR_CLOSED), the stream has pending operations (with error being set
// to G_IO_ERROR_PENDING), or if querying info is not supported for the stream's
// interface (with error being set to G_IO_ERROR_NOT_SUPPORTED). In all cases of
// failure, NULL will be returned.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be set, and NULL will be returned.
func (stream *FileOutputStream) QueryInfo(ctx context.Context, attributes string) (*FileInfo, error) {
	var _arg0 *C.GFileOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 *C.char              // out
	var _cret *C.GFileInfo         // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_output_stream_query_info(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(attributes)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = wrapFileInfo(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _fileInfo, _goerr
}

// QueryInfoAsync: asynchronously queries the stream for a Info. When completed,
// callback will be called with a Result which can be used to finish the
// operation with g_file_output_stream_query_info_finish().
//
// For the synchronous version of this function, see
// g_file_output_stream_query_info().
func (stream *FileOutputStream) QueryInfoAsync(ctx context.Context, attributes string, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GFileOutputStream  // out
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_file_output_stream_query_info_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// QueryInfoFinish finalizes the asynchronous query started by
// g_file_output_stream_query_info_async().
func (stream *FileOutputStream) QueryInfoFinish(result AsyncResulter) (*FileInfo, error) {
	var _arg0 *C.GFileOutputStream // out
	var _arg1 *C.GAsyncResult      // out
	var _cret *C.GFileInfo         // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_file_output_stream_query_info_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = wrapFileInfo(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _fileInfo, _goerr
}
