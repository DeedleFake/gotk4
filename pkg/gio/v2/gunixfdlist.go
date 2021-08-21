// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_unix_fd_list_get_type()), F: marshalUnixFDLister},
	})
}

// UnixFDList contains a list of file descriptors. It owns the file descriptors
// that it contains, closing them when finalized.
//
// It may be wrapped in a FDMessage and sent over a #GSocket in the
// G_SOCKET_FAMILY_UNIX family by using g_socket_send_message() and received
// using g_socket_receive_message().
//
// Note that <gio/gunixfdlist.h> belongs to the UNIX-specific GIO interfaces,
// thus you have to use the gio-unix-2.0.pc pkg-config file when using it.
type UnixFDList struct {
	*externglib.Object
}

func wrapUnixFDList(obj *externglib.Object) *UnixFDList {
	return &UnixFDList{
		Object: obj,
	}
}

func marshalUnixFDLister(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapUnixFDList(obj), nil
}

// NewUnixFDList creates a new FDList containing no file descriptors.
func NewUnixFDList() *UnixFDList {
	var _cret *C.GUnixFDList // in

	_cret = C.g_unix_fd_list_new()

	var _unixFDList *UnixFDList // out

	_unixFDList = wrapUnixFDList(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixFDList
}

// NewUnixFDListFromArray creates a new FDList containing the file descriptors
// given in fds. The file descriptors become the property of the new list and
// may no longer be used by the caller. The array itself is owned by the caller.
//
// Each file descriptor in the array should be set to close-on-exec.
//
// If n_fds is -1 then fds must be terminated with -1.
func NewUnixFDListFromArray(fds []int) *UnixFDList {
	var _arg1 *C.gint // out
	var _arg2 C.gint
	var _cret *C.GUnixFDList // in

	_arg2 = (C.gint)(len(fds))
	if len(fds) > 0 {
		_arg1 = (*C.gint)(unsafe.Pointer(&fds[0]))
	}

	_cret = C.g_unix_fd_list_new_from_array(_arg1, _arg2)
	runtime.KeepAlive(fds)

	var _unixFDList *UnixFDList // out

	_unixFDList = wrapUnixFDList(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixFDList
}

// Append adds a file descriptor to list.
//
// The file descriptor is duplicated using dup(). You keep your copy of the
// descriptor and the copy contained in list will be closed when list is
// finalized.
//
// A possible cause of failure is exceeding the per-process or system-wide file
// descriptor limit.
//
// The index of the file descriptor in the list is returned. If you use this
// index with g_unix_fd_list_get() then you will receive back a duplicated copy
// of the same file descriptor.
func (list *UnixFDList) Append(fd int) (int, error) {
	var _arg0 *C.GUnixFDList // out
	var _arg1 C.gint         // out
	var _cret C.gint         // in
	var _cerr *C.GError      // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(list.Native()))
	_arg1 = C.gint(fd)

	_cret = C.g_unix_fd_list_append(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(list)
	runtime.KeepAlive(fd)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// Get gets a file descriptor out of list.
//
// index_ specifies the index of the file descriptor to get. It is a programmer
// error for index_ to be out of range; see g_unix_fd_list_get_length().
//
// The file descriptor is duplicated using dup() and set as close-on-exec before
// being returned. You must call close() on it when you are done.
//
// A possible cause of failure is exceeding the per-process or system-wide file
// descriptor limit.
func (list *UnixFDList) Get(index_ int) (int, error) {
	var _arg0 *C.GUnixFDList // out
	var _arg1 C.gint         // out
	var _cret C.gint         // in
	var _cerr *C.GError      // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(list.Native()))
	_arg1 = C.gint(index_)

	_cret = C.g_unix_fd_list_get(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(list)
	runtime.KeepAlive(index_)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// Length gets the length of list (ie: the number of file descriptors contained
// within).
func (list *UnixFDList) Length() int {
	var _arg0 *C.GUnixFDList // out
	var _cret C.gint         // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(list.Native()))

	_cret = C.g_unix_fd_list_get_length(_arg0)
	runtime.KeepAlive(list)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PeekFds returns the array of file descriptors that is contained in this
// object.
//
// After this call, the descriptors remain the property of list. The caller must
// not close them and must not free the array. The array is valid only until
// list is changed in any way.
//
// If length is non-NULL then it is set to the number of file descriptors in the
// returned array. The returned array is also terminated with -1.
//
// This function never returns NULL. In case there are no file descriptors
// contained in list, an empty array is returned.
func (list *UnixFDList) PeekFds() []int {
	var _arg0 *C.GUnixFDList // out
	var _cret *C.gint        // in
	var _arg1 C.gint         // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(list.Native()))

	_cret = C.g_unix_fd_list_peek_fds(_arg0, &_arg1)
	runtime.KeepAlive(list)

	var _gints []int // out

	_gints = make([]int, _arg1)
	copy(_gints, unsafe.Slice((*int)(unsafe.Pointer(_cret)), _arg1))

	return _gints
}

// StealFds returns the array of file descriptors that is contained in this
// object.
//
// After this call, the descriptors are no longer contained in list. Further
// calls will return an empty list (unless more descriptors have been added).
//
// The return result of this function must be freed with g_free(). The caller is
// also responsible for closing all of the file descriptors. The file
// descriptors in the array are set to close-on-exec.
//
// If length is non-NULL then it is set to the number of file descriptors in the
// returned array. The returned array is also terminated with -1.
//
// This function never returns NULL. In case there are no file descriptors
// contained in list, an empty array is returned.
func (list *UnixFDList) StealFds() []int {
	var _arg0 *C.GUnixFDList // out
	var _cret *C.gint        // in
	var _arg1 C.gint         // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(list.Native()))

	_cret = C.g_unix_fd_list_steal_fds(_arg0, &_arg1)
	runtime.KeepAlive(list)

	var _gints []int // out

	defer C.free(unsafe.Pointer(_cret))
	_gints = make([]int, _arg1)
	copy(_gints, unsafe.Slice((*int)(unsafe.Pointer(_cret)), _arg1))

	return _gints
}
