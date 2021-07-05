// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_io_channel_get_type()), F: marshalIOChannel},
	})
}

// IOChannelError: error codes returned by OChannel operations.
type IOChannelError int

const (
	// fbig: file too large.
	IOChannelErrorFbig IOChannelError = iota
	// inval: invalid argument.
	IOChannelErrorInval
	// io: IO error.
	IOChannelErrorIO
	// isdir: file is a directory.
	IOChannelErrorIsdir
	// nospc: no space left on device.
	IOChannelErrorNospc
	// nxio: no such device or address.
	IOChannelErrorNxio
	// overflow: value too large for defined datatype.
	IOChannelErrorOverflow
	// pipe: broken pipe.
	IOChannelErrorPipe
	// failed: some other error.
	IOChannelErrorFailed
)

// IOError is only used by the deprecated functions g_io_channel_read(),
// g_io_channel_write(), and g_io_channel_seek().
type IOError int

const (
	// none: no error
	IOErrorNone IOError = iota
	// again: EAGAIN error occurred
	IOErrorAgain
	// inval: EINVAL error occurred
	IOErrorInval
	// unknown: another error occurred
	IOErrorUnknown
)

// IOStatus statuses returned by most of the OFuncs functions.
type IOStatus int

const (
	// error occurred.
	IOStatusError IOStatus = iota
	// normal: success.
	IOStatusNormal
	// eof: end of file.
	IOStatusEOF
	// again: resource temporarily unavailable.
	IOStatusAgain
)

// SeekType: enumeration specifying the base position for a
// g_io_channel_seek_position() operation.
type SeekType int

const (
	// cur: the current position in the file.
	SeekTypeCur SeekType = iota
	// set: the start of the file.
	SeekTypeSet
	// end: the end of the file.
	SeekTypeEnd
)

// IOFlags specifies properties of a OChannel. Some of the flags can only be
// read with g_io_channel_get_flags(), but not changed with
// g_io_channel_set_flags().
type IOFlags int

const (
	// IOFlagsAppend turns on append mode, corresponds to O_APPEND (see the
	// documentation of the UNIX open() syscall)
	IOFlagsAppend IOFlags = 0b1
	// IOFlagsNonblock turns on nonblocking mode, corresponds to
	// O_NONBLOCK/O_NDELAY (see the documentation of the UNIX open() syscall)
	IOFlagsNonblock IOFlags = 0b10
	// IOFlagsIsReadable indicates that the io channel is readable. This flag
	// cannot be changed.
	IOFlagsIsReadable IOFlags = 0b100
	// IOFlagsIsWritable indicates that the io channel is writable. This flag
	// cannot be changed.
	IOFlagsIsWritable IOFlags = 0b1000
	// IOFlagsIsWriteable: misspelled version of @G_IO_FLAG_IS_WRITABLE that
	// existed before the spelling was fixed in GLib 2.30. It is kept here for
	// compatibility reasons. Deprecated since 2.30
	IOFlagsIsWriteable IOFlags = 0b1000
	// IOFlagsIsSeekable indicates that the io channel is seekable, i.e. that
	// g_io_channel_seek_position() can be used on it. This flag cannot be
	// changed.
	IOFlagsIsSeekable IOFlags = 0b10000
	// IOFlagsMask: the mask that specifies all the valid flags.
	IOFlagsMask IOFlags = 0b11111
	// IOFlagsGetMask: the mask of the flags that are returned from
	// g_io_channel_get_flags()
	IOFlagsGetMask IOFlags = 0b11111
	// IOFlagsSetMask: the mask of the flags that the user can modify with
	// g_io_channel_set_flags()
	IOFlagsSetMask IOFlags = 0b11
)

// IOCreateWatch creates a #GSource that's dispatched when @condition is met for
// the given @channel. For example, if condition is IO_IN, the source will be
// dispatched when there's data available for reading.
//
// The callback function invoked by the #GSource should be added with
// g_source_set_callback(), but it has type OFunc (not Func).
//
// g_io_add_watch() is a simpler interface to this same functionality, for the
// case where you want to add the source to the default main loop context at the
// default priority.
//
// On Windows, polling a #GSource created to watch a channel for a socket puts
// the socket in non-blocking mode. This is a side-effect of the implementation
// and unavoidable.
func IOCreateWatch(channel *IOChannel, condition IOCondition) *Source {
	var _arg1 *C.GIOChannel  // out
	var _arg2 C.GIOCondition // out
	var _cret *C.GSource     // in

	_arg1 = (*C.GIOChannel)(unsafe.Pointer(channel))
	_arg2 = C.GIOCondition(condition)

	_cret = C.g_io_create_watch(_arg1, _arg2)

	var _source *Source // out

	_source = (*Source)(unsafe.Pointer(_cret))
	C.g_source_ref(_cret)
	runtime.SetFinalizer(_source, func(v *Source) {
		C.g_source_unref((*C.GSource)(unsafe.Pointer(v)))
	})

	return _source
}

// IOChannel: data structure representing an IO Channel. The fields should be
// considered private and should only be accessed with the following functions.
type IOChannel struct {
	native C.GIOChannel
}

// WrapIOChannel wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIOChannel(ptr unsafe.Pointer) *IOChannel {
	return (*IOChannel)(ptr)
}

func marshalIOChannel(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*IOChannel)(unsafe.Pointer(b)), nil
}

// NewIOChannelFile constructs a struct IOChannel.
func NewIOChannelFile(filename string, mode string) (*IOChannel, error) {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _cret *C.GIOChannel // in
	var _cerr *C.GError     // in

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(mode))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_io_channel_new_file(_arg1, _arg2, &_cerr)

	var _ioChannel *IOChannel // out
	var _goerr error          // out

	_ioChannel = (*IOChannel)(unsafe.Pointer(_cret))
	C.g_io_channel_ref(_cret)
	runtime.SetFinalizer(_ioChannel, func(v *IOChannel) {
		C.g_io_channel_unref((*C.GIOChannel)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioChannel, _goerr
}

// NewIOChannelUnix constructs a struct IOChannel.
func NewIOChannelUnix(fd int) *IOChannel {
	var _arg1 C.int         // out
	var _cret *C.GIOChannel // in

	_arg1 = C.int(fd)

	_cret = C.g_io_channel_unix_new(_arg1)

	var _ioChannel *IOChannel // out

	_ioChannel = (*IOChannel)(unsafe.Pointer(_cret))
	C.g_io_channel_ref(_cret)
	runtime.SetFinalizer(_ioChannel, func(v *IOChannel) {
		C.g_io_channel_unref((*C.GIOChannel)(unsafe.Pointer(v)))
	})

	return _ioChannel
}

// Native returns the underlying C source pointer.
func (i *IOChannel) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// Close an IO channel. Any pending data to be written will be flushed, ignoring
// errors. The channel will not be freed until the last reference is dropped
// using g_io_channel_unref().
//
// Deprecated: since version 2.2.
func (c *IOChannel) Close() {
	var _arg0 *C.GIOChannel // out

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	C.g_io_channel_close(_arg0)
}

// Flush flushes the write buffer for the GIOChannel.
func (c *IOChannel) Flush() (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_flush(_arg0, &_cerr)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStatus, _goerr
}

// BufferCondition: this function returns a OCondition depending on whether
// there is data to be read/space to write data in the internal buffers in the
// OChannel. Only the flags G_IO_IN and G_IO_OUT may be set.
func (c *IOChannel) BufferCondition() IOCondition {
	var _arg0 *C.GIOChannel  // out
	var _cret C.GIOCondition // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_get_buffer_condition(_arg0)

	var _ioCondition IOCondition // out

	_ioCondition = IOCondition(_cret)

	return _ioCondition
}

// BufferSize gets the buffer size.
func (c *IOChannel) BufferSize() uint {
	var _arg0 *C.GIOChannel // out
	var _cret C.gsize       // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_get_buffer_size(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Buffered returns whether @channel is buffered.
func (c *IOChannel) Buffered() bool {
	var _arg0 *C.GIOChannel // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_get_buffered(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CloseOnUnref returns whether the file/socket/whatever associated with
// @channel will be closed when @channel receives its final unref and is
// destroyed. The default value of this is true for channels created by
// g_io_channel_new_file (), and false for all other channels.
func (c *IOChannel) CloseOnUnref() bool {
	var _arg0 *C.GIOChannel // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_get_close_on_unref(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Encoding gets the encoding for the input/output of the channel. The internal
// encoding is always UTF-8. The encoding nil makes the channel safe for binary
// data.
func (c *IOChannel) Encoding() string {
	var _arg0 *C.GIOChannel // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_get_encoding(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Flags gets the current flags for a OChannel, including read-only flags such
// as G_IO_FLAG_IS_READABLE.
//
// The values of the flags G_IO_FLAG_IS_READABLE and G_IO_FLAG_IS_WRITABLE are
// cached for internal use by the channel when it is created. If they should
// change at some later point (e.g. partial shutdown of a socket with the UNIX
// shutdown() function), the user should immediately call
// g_io_channel_get_flags() to update the internal values of these flags.
func (c *IOChannel) Flags() IOFlags {
	var _arg0 *C.GIOChannel // out
	var _cret C.GIOFlags    // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_get_flags(_arg0)

	var _ioFlags IOFlags // out

	_ioFlags = IOFlags(_cret)

	return _ioFlags
}

// LineTerm: this returns the string that OChannel uses to determine where in
// the file a line break occurs. A value of nil indicates autodetection.
func (c *IOChannel) LineTerm(length *int) string {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gint       // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = (*C.gint)(unsafe.Pointer(length))

	_cret = C.g_io_channel_get_line_term(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Init initializes a OChannel struct.
//
// This is called by each of the above functions when creating a OChannel, and
// so is not often needed by the application programmer (unless you are creating
// a new type of OChannel).
func (c *IOChannel) Init() {
	var _arg0 *C.GIOChannel // out

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	C.g_io_channel_init(_arg0)
}

// Read reads data from a OChannel.
//
// Deprecated: since version 2.2.
func (c *IOChannel) Read(buf string, count uint, bytesRead *uint) IOError {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gsize       // out
	var _arg3 *C.gsize      // out
	var _cret C.GIOError    // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = (*C.gchar)(C.CString(buf))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gsize(count)
	_arg3 = (*C.gsize)(unsafe.Pointer(bytesRead))

	_cret = C.g_io_channel_read(_arg0, _arg1, _arg2, _arg3)

	var _ioError IOError // out

	_ioError = IOError(_cret)

	return _ioError
}

// ReadLine reads a line, including the terminating character(s), from a
// OChannel into a newly-allocated string. @str_return will contain allocated
// memory if the return is G_IO_STATUS_NORMAL.
func (c *IOChannel) ReadLine() (strReturn string, length uint, terminatorPos uint, ioStatus IOStatus, goerr error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // in
	var _arg2 C.gsize       // in
	var _arg3 C.gsize       // in
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_read_line(_arg0, &_arg1, &_arg2, &_arg3, &_cerr)

	var _strReturn string   // out
	var _length uint        // out
	var _terminatorPos uint // out
	var _ioStatus IOStatus  // out
	var _goerr error        // out

	_strReturn = C.GoString(_arg1)
	defer C.free(unsafe.Pointer(_arg1))
	_length = uint(_arg2)
	_terminatorPos = uint(_arg3)
	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _strReturn, _length, _terminatorPos, _ioStatus, _goerr
}

// ReadToEnd reads all the remaining data from the file.
func (c *IOChannel) ReadToEnd() ([]byte, IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar
	var _arg2 C.gsize     // in
	var _cret C.GIOStatus // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_read_to_end(_arg0, &_arg1, &_arg2, &_cerr)

	var _strReturn []byte
	var _ioStatus IOStatus // out
	var _goerr error       // out

	_strReturn = unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), _arg2)
	runtime.SetFinalizer(&_strReturn, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _strReturn, _ioStatus, _goerr
}

// ReadUnichar reads a Unicode character from @channel. This function cannot be
// called on a channel with nil encoding.
func (c *IOChannel) ReadUnichar() (uint32, IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gunichar    // in
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_read_unichar(_arg0, &_arg1, &_cerr)

	var _thechar uint32    // out
	var _ioStatus IOStatus // out
	var _goerr error       // out

	_thechar = uint32(_arg1)
	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _thechar, _ioStatus, _goerr
}

// Ref increments the reference count of a OChannel.
func (c *IOChannel) Ref() *IOChannel {
	var _arg0 *C.GIOChannel // out
	var _cret *C.GIOChannel // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_ref(_arg0)

	var _ioChannel *IOChannel // out

	_ioChannel = (*IOChannel)(unsafe.Pointer(_cret))
	C.g_io_channel_ref(_cret)
	runtime.SetFinalizer(_ioChannel, func(v *IOChannel) {
		C.g_io_channel_unref((*C.GIOChannel)(unsafe.Pointer(v)))
	})

	return _ioChannel
}

// Seek sets the current position in the OChannel, similar to the standard
// library function fseek().
//
// Deprecated: since version 2.2.
func (c *IOChannel) Seek(offset int64, typ SeekType) IOError {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gint64      // out
	var _arg2 C.GSeekType   // out
	var _cret C.GIOError    // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = C.gint64(offset)
	_arg2 = C.GSeekType(typ)

	_cret = C.g_io_channel_seek(_arg0, _arg1, _arg2)

	var _ioError IOError // out

	_ioError = IOError(_cret)

	return _ioError
}

// SeekPosition: replacement for g_io_channel_seek() with the new API.
func (c *IOChannel) SeekPosition(offset int64, typ SeekType) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gint64      // out
	var _arg2 C.GSeekType   // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = C.gint64(offset)
	_arg2 = C.GSeekType(typ)

	_cret = C.g_io_channel_seek_position(_arg0, _arg1, _arg2, &_cerr)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStatus, _goerr
}

// SetBufferSize sets the buffer size.
func (c *IOChannel) SetBufferSize(size uint) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gsize       // out

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = C.gsize(size)

	C.g_io_channel_set_buffer_size(_arg0, _arg1)
}

// SetBuffered: the buffering state can only be set if the channel's encoding is
// nil. For any other encoding, the channel must be buffered.
//
// A buffered channel can only be set unbuffered if the channel's internal
// buffers have been flushed. Newly created channels or channels which have
// returned G_IO_STATUS_EOF not require such a flush. For write-only channels, a
// call to g_io_channel_flush () is sufficient. For all other channels, the
// buffers may be flushed by a call to g_io_channel_seek_position (). This
// includes the possibility of seeking with seek type G_SEEK_CUR and an offset
// of zero. Note that this means that socket-based channels cannot be set
// unbuffered once they have had data read from them.
//
// On unbuffered channels, it is safe to mix read and write calls from the new
// and old APIs, if this is necessary for maintaining old code.
//
// The default state of the channel is buffered.
func (c *IOChannel) SetBuffered(buffered bool) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	if buffered {
		_arg1 = C.TRUE
	}

	C.g_io_channel_set_buffered(_arg0, _arg1)
}

// SetCloseOnUnref: whether to close the channel on the final unref of the
// OChannel data structure. The default value of this is true for channels
// created by g_io_channel_new_file (), and false for all other channels.
//
// Setting this flag to true for a channel you have already closed can cause
// problems when the final reference to the OChannel is dropped.
func (c *IOChannel) SetCloseOnUnref(doClose bool) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	if doClose {
		_arg1 = C.TRUE
	}

	C.g_io_channel_set_close_on_unref(_arg0, _arg1)
}

// SetEncoding sets the encoding for the input/output of the channel. The
// internal encoding is always UTF-8. The default encoding for the external file
// is UTF-8.
//
// The encoding nil is safe to use with binary data.
//
// The encoding can only be set if one of the following conditions is true:
//
// - The channel was just created, and has not been written to or read from yet.
//
// - The channel is write-only.
//
// - The channel is a file, and the file pointer was just repositioned by a call
// to g_io_channel_seek_position(). (This flushes all the internal buffers.)
//
// - The current encoding is nil or UTF-8.
//
// - One of the (new API) read functions has just returned G_IO_STATUS_EOF (or,
// in the case of g_io_channel_read_to_end(), G_IO_STATUS_NORMAL).
//
// - One of the functions g_io_channel_read_chars() or
// g_io_channel_read_unichar() has returned G_IO_STATUS_AGAIN or
// G_IO_STATUS_ERROR. This may be useful in the case of
// G_CONVERT_ERROR_ILLEGAL_SEQUENCE. Returning one of these statuses from
// g_io_channel_read_line(), g_io_channel_read_line_string(), or
// g_io_channel_read_to_end() does not guarantee that the encoding can be
// changed.
//
// Channels which do not meet one of the above conditions cannot call
// g_io_channel_seek_position() with an offset of G_SEEK_CUR, and, if they are
// "seekable", cannot call g_io_channel_write_chars() after calling one of the
// API "read" functions.
func (c *IOChannel) SetEncoding(encoding string) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = (*C.gchar)(C.CString(encoding))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_io_channel_set_encoding(_arg0, _arg1, &_cerr)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStatus, _goerr
}

// SetFlags sets the (writeable) flags in @channel to (@flags &
// G_IO_FLAG_SET_MASK).
func (c *IOChannel) SetFlags(flags IOFlags) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.GIOFlags    // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = C.GIOFlags(flags)

	_cret = C.g_io_channel_set_flags(_arg0, _arg1, &_cerr)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStatus, _goerr
}

// SetLineTerm: this sets the string that OChannel uses to determine where in
// the file a line break occurs.
func (c *IOChannel) SetLineTerm(lineTerm string, length int) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = (*C.gchar)(C.CString(lineTerm))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)

	C.g_io_channel_set_line_term(_arg0, _arg1, _arg2)
}

// Shutdown: close an IO channel. Any pending data to be written will be flushed
// if @flush is true. The channel will not be freed until the last reference is
// dropped using g_io_channel_unref().
func (c *IOChannel) Shutdown(flush bool) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gboolean    // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	if flush {
		_arg1 = C.TRUE
	}

	_cret = C.g_io_channel_shutdown(_arg0, _arg1, &_cerr)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStatus, _goerr
}

// UnixGetFd returns the file descriptor of the OChannel.
//
// On Windows this function returns the file descriptor or socket of the
// OChannel.
func (c *IOChannel) UnixGetFd() int {
	var _arg0 *C.GIOChannel // out
	var _cret C.gint        // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	_cret = C.g_io_channel_unix_get_fd(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Unref decrements the reference count of a OChannel.
func (c *IOChannel) Unref() {
	var _arg0 *C.GIOChannel // out

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))

	C.g_io_channel_unref(_arg0)
}

// Write writes data to a OChannel.
//
// Deprecated: since version 2.2.
func (c *IOChannel) Write(buf string, count uint, bytesWritten *uint) IOError {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gsize       // out
	var _arg3 *C.gsize      // out
	var _cret C.GIOError    // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = (*C.gchar)(C.CString(buf))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gsize(count)
	_arg3 = (*C.gsize)(unsafe.Pointer(bytesWritten))

	_cret = C.g_io_channel_write(_arg0, _arg1, _arg2, _arg3)

	var _ioError IOError // out

	_ioError = IOError(_cret)

	return _ioError
}

// WriteChars: replacement for g_io_channel_write() with the new API.
//
// On seekable channels with encodings other than nil or UTF-8, generic mixing
// of reading and writing is not allowed. A call to g_io_channel_write_chars ()
// may only be made on a channel from which data has been read in the cases
// described in the documentation for g_io_channel_set_encoding ().
func (c *IOChannel) WriteChars(buf []byte, count int) (uint, IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 *C.gchar
	var _arg2 C.gssize    // out
	var _arg3 C.gsize     // in
	var _cret C.GIOStatus // in
	var _cerr *C.GError   // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = (*C.gchar)(unsafe.Pointer(&buf[0]))
	_arg2 = C.gssize(count)

	_cret = C.g_io_channel_write_chars(_arg0, _arg1, _arg2, &_arg3, &_cerr)

	var _bytesWritten uint // out
	var _ioStatus IOStatus // out
	var _goerr error       // out

	_bytesWritten = uint(_arg3)
	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _bytesWritten, _ioStatus, _goerr
}

// WriteUnichar writes a Unicode character to @channel. This function cannot be
// called on a channel with nil encoding.
func (c *IOChannel) WriteUnichar(thechar uint32) (IOStatus, error) {
	var _arg0 *C.GIOChannel // out
	var _arg1 C.gunichar    // out
	var _cret C.GIOStatus   // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GIOChannel)(unsafe.Pointer(c))
	_arg1 = C.gunichar(thechar)

	_cret = C.g_io_channel_write_unichar(_arg0, _arg1, &_cerr)

	var _ioStatus IOStatus // out
	var _goerr error       // out

	_ioStatus = IOStatus(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _ioStatus, _goerr
}

// IOFuncs: table of functions used to handle different types of OChannel in a
// generic way.
type IOFuncs struct {
	native C.GIOFuncs
}

// WrapIOFuncs wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIOFuncs(ptr unsafe.Pointer) *IOFuncs {
	return (*IOFuncs)(ptr)
}

// Native returns the underlying C source pointer.
func (i *IOFuncs) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}
