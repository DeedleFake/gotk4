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
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_mount_get_type()), F: marshalMounter},
	})
}

// MountOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MountOverrider interface {
	// CanEject checks if mount can be ejected.
	CanEject() bool
	// CanUnmount checks if mount can be unmounted.
	CanUnmount() bool
	Changed()
	// Eject ejects a mount. This is an asynchronous operation, and is finished
	// by calling g_mount_eject_finish() with the mount and Result data returned
	// in the callback.
	//
	// Deprecated: Use g_mount_eject_with_operation() instead.
	Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a mount. If any errors occurred during the
	// operation, error will be set to contain the errors and FALSE will be
	// returned.
	//
	// Deprecated: Use g_mount_eject_with_operation_finish() instead.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperation ejects a mount. This is an asynchronous operation, and
	// is finished by calling g_mount_eject_with_operation_finish() with the
	// mount and Result data returned in the callback.
	EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a mount. If any errors
	// occurred during the operation, error will be set to contain the errors
	// and FALSE will be returned.
	EjectWithOperationFinish(result AsyncResulter) error
	// DefaultLocation gets the default location of mount. The default location
	// of the given mount is a path that reflects the main entry point for the
	// user (e.g. the home directory, or the root of the volume).
	DefaultLocation() Filer
	// Drive gets the drive for the mount.
	//
	// This is a convenience method for getting the #GVolume and then using that
	// object to get the #GDrive.
	Drive() Driver
	// Icon gets the icon for mount.
	Icon() Iconner
	// Name gets the name of mount.
	Name() string
	// Root gets the root directory on mount.
	Root() Filer
	// SortKey gets the sort key for mount, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for mount.
	SymbolicIcon() Iconner
	// UUID gets the UUID for the mount. The reference is typically based on the
	// file system UUID for the mount in question and should be considered an
	// opaque string. Returns NULL if there is no UUID available.
	UUID() string
	// Volume gets the volume for the mount.
	Volume() Volumer
	// GuessContentType tries to guess the type of content stored on mount.
	// Returns one or more textual identifiers of well-known content types
	// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
	// camera memory cards. See the shared-mime-info
	// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
	// specification for more on x-content types.
	//
	// This is an asynchronous operation (see g_mount_guess_content_type_sync()
	// for the synchronous version), and is finished by calling
	// g_mount_guess_content_type_finish() with the mount and Result data
	// returned in the callback.
	GuessContentType(ctx context.Context, forceRescan bool, callback AsyncReadyCallback)
	// GuessContentTypeFinish finishes guessing content types of mount. If any
	// errors occurred during the operation, error will be set to contain the
	// errors and FALSE will be returned. In particular, you may get an
	// G_IO_ERROR_NOT_SUPPORTED if the mount does not support content guessing.
	GuessContentTypeFinish(result AsyncResulter) ([]string, error)
	// GuessContentTypeSync tries to guess the type of content stored on mount.
	// Returns one or more textual identifiers of well-known content types
	// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
	// camera memory cards. See the shared-mime-info
	// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
	// specification for more on x-content types.
	//
	// This is a synchronous operation and as such may block doing IO; see
	// g_mount_guess_content_type() for the asynchronous version.
	GuessContentTypeSync(ctx context.Context, forceRescan bool) ([]string, error)
	PreUnmount()
	// Remount remounts a mount. This is an asynchronous operation, and is
	// finished by calling g_mount_remount_finish() with the mount and Results
	// data returned in the callback.
	//
	// Remounting is useful when some setting affecting the operation of the
	// volume has been changed, as these may need a remount to take affect.
	// While this is semantically equivalent with unmounting and then remounting
	// not all backends might need to actually be unmounted.
	Remount(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// RemountFinish finishes remounting a mount. If any errors occurred during
	// the operation, error will be set to contain the errors and FALSE will be
	// returned.
	RemountFinish(result AsyncResulter) error
	// Unmount unmounts a mount. This is an asynchronous operation, and is
	// finished by calling g_mount_unmount_finish() with the mount and Result
	// data returned in the callback.
	//
	// Deprecated: Use g_mount_unmount_with_operation() instead.
	Unmount(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// UnmountFinish finishes unmounting a mount. If any errors occurred during
	// the operation, error will be set to contain the errors and FALSE will be
	// returned.
	//
	// Deprecated: Use g_mount_unmount_with_operation_finish() instead.
	UnmountFinish(result AsyncResulter) error
	// UnmountWithOperation unmounts a mount. This is an asynchronous operation,
	// and is finished by calling g_mount_unmount_with_operation_finish() with
	// the mount and Result data returned in the callback.
	UnmountWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// UnmountWithOperationFinish finishes unmounting a mount. If any errors
	// occurred during the operation, error will be set to contain the errors
	// and FALSE will be returned.
	UnmountWithOperationFinish(result AsyncResulter) error
	Unmounted()
}

// Mount interface represents user-visible mounts. Note, when porting from
// GnomeVFS, #GMount is the moral equivalent of VFSVolume.
//
// #GMount is a "mounted" filesystem that you can access. Mounted is in quotes
// because it's not the same as a unix mount, it might be a gvfs mount, but you
// can still access the files on it if you use GIO. Might or might not be
// related to a volume object.
//
// Unmounting a #GMount instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To unmount
// a #GMount instance, first call g_mount_unmount_with_operation() with (at
// least) the #GMount instance and a ReadyCallback. The callback will be fired
// when the operation has resolved (either with success or failure), and a
// Result structure will be passed to the callback. That callback should then
// call g_mount_unmount_with_operation_finish() with the #GMount and the Result
// data to see if the operation was completed successfully. If an error is
// present when g_mount_unmount_with_operation_finish() is called, then it will
// be filled with any error information.
type Mount struct {
	*externglib.Object
}

// Mounter describes Mount's abstract methods.
type Mounter interface {
	externglib.Objector

	// CanEject checks if mount can be ejected.
	CanEject() bool
	// CanUnmount checks if mount can be unmounted.
	CanUnmount() bool
	// Eject ejects a mount.
	Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// EjectFinish finishes ejecting a mount.
	EjectFinish(result AsyncResulter) error
	// EjectWithOperation ejects a mount.
	EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// EjectWithOperationFinish finishes ejecting a mount.
	EjectWithOperationFinish(result AsyncResulter) error
	// DefaultLocation gets the default location of mount.
	DefaultLocation() Filer
	// Drive gets the drive for the mount.
	Drive() Driver
	// Icon gets the icon for mount.
	Icon() Iconner
	// Name gets the name of mount.
	Name() string
	// Root gets the root directory on mount.
	Root() Filer
	// SortKey gets the sort key for mount, if any.
	SortKey() string
	// SymbolicIcon gets the symbolic icon for mount.
	SymbolicIcon() Iconner
	// UUID gets the UUID for the mount.
	UUID() string
	// Volume gets the volume for the mount.
	Volume() Volumer
	// GuessContentType tries to guess the type of content stored on mount.
	GuessContentType(ctx context.Context, forceRescan bool, callback AsyncReadyCallback)
	// GuessContentTypeFinish finishes guessing content types of mount.
	GuessContentTypeFinish(result AsyncResulter) ([]string, error)
	// GuessContentTypeSync tries to guess the type of content stored on mount.
	GuessContentTypeSync(ctx context.Context, forceRescan bool) ([]string, error)
	// IsShadowed determines if mount is shadowed.
	IsShadowed() bool
	// Remount remounts a mount.
	Remount(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// RemountFinish finishes remounting a mount.
	RemountFinish(result AsyncResulter) error
	// Shadow increments the shadow count on mount.
	Shadow()
	// Unmount unmounts a mount.
	Unmount(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback)
	// UnmountFinish finishes unmounting a mount.
	UnmountFinish(result AsyncResulter) error
	// UnmountWithOperation unmounts a mount.
	UnmountWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback)
	// UnmountWithOperationFinish finishes unmounting a mount.
	UnmountWithOperationFinish(result AsyncResulter) error
	// Unshadow decrements the shadow count on mount.
	Unshadow()
}

var _ Mounter = (*Mount)(nil)

func wrapMount(obj *externglib.Object) *Mount {
	return &Mount{
		Object: obj,
	}
}

func marshalMounter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMount(obj), nil
}

// CanEject checks if mount can be ejected.
func (mount *Mount) CanEject() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_can_eject(_arg0)
	runtime.KeepAlive(mount)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanUnmount checks if mount can be unmounted.
func (mount *Mount) CanUnmount() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_can_unmount(_arg0)
	runtime.KeepAlive(mount)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Eject ejects a mount. This is an asynchronous operation, and is finished by
// calling g_mount_eject_finish() with the mount and Result data returned in the
// callback.
//
// Deprecated: Use g_mount_eject_with_operation() instead.
func (mount *Mount) Eject(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_eject(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// EjectFinish finishes ejecting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// Deprecated: Use g_mount_eject_with_operation_finish() instead.
func (mount *Mount) EjectFinish(result AsyncResulter) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_eject_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// EjectWithOperation ejects a mount. This is an asynchronous operation, and is
// finished by calling g_mount_eject_with_operation_finish() with the mount and
// Result data returned in the callback.
func (mount *Mount) EjectWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_eject_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// EjectWithOperationFinish finishes ejecting a mount. If any errors occurred
// during the operation, error will be set to contain the errors and FALSE will
// be returned.
func (mount *Mount) EjectWithOperationFinish(result AsyncResulter) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_eject_with_operation_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// DefaultLocation gets the default location of mount. The default location of
// the given mount is a path that reflects the main entry point for the user
// (e.g. the home directory, or the root of the volume).
func (mount *Mount) DefaultLocation() Filer {
	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_default_location(_arg0)
	runtime.KeepAlive(mount)

	var _file Filer // out

	_file = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Filer)

	return _file
}

// Drive gets the drive for the mount.
//
// This is a convenience method for getting the #GVolume and then using that
// object to get the #GDrive.
func (mount *Mount) Drive() Driver {
	var _arg0 *C.GMount // out
	var _cret *C.GDrive // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_drive(_arg0)
	runtime.KeepAlive(mount)

	var _drive Driver // out

	if _cret != nil {
		_drive = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Driver)
	}

	return _drive
}

// Icon gets the icon for mount.
func (mount *Mount) Icon() Iconner {
	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_icon(_arg0)
	runtime.KeepAlive(mount)

	var _icon Iconner // out

	_icon = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Iconner)

	return _icon
}

// Name gets the name of mount.
func (mount *Mount) Name() string {
	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_name(_arg0)
	runtime.KeepAlive(mount)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Root gets the root directory on mount.
func (mount *Mount) Root() Filer {
	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_root(_arg0)
	runtime.KeepAlive(mount)

	var _file Filer // out

	_file = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Filer)

	return _file
}

// SortKey gets the sort key for mount, if any.
func (mount *Mount) SortKey() string {
	var _arg0 *C.GMount // out
	var _cret *C.gchar  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_sort_key(_arg0)
	runtime.KeepAlive(mount)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SymbolicIcon gets the symbolic icon for mount.
func (mount *Mount) SymbolicIcon() Iconner {
	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_symbolic_icon(_arg0)
	runtime.KeepAlive(mount)

	var _icon Iconner // out

	_icon = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Iconner)

	return _icon
}

// UUID gets the UUID for the mount. The reference is typically based on the
// file system UUID for the mount in question and should be considered an opaque
// string. Returns NULL if there is no UUID available.
func (mount *Mount) UUID() string {
	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_uuid(_arg0)
	runtime.KeepAlive(mount)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Volume gets the volume for the mount.
func (mount *Mount) Volume() Volumer {
	var _arg0 *C.GMount  // out
	var _cret *C.GVolume // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_get_volume(_arg0)
	runtime.KeepAlive(mount)

	var _volume Volumer // out

	if _cret != nil {
		_volume = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(Volumer)
	}

	return _volume
}

// GuessContentType tries to guess the type of content stored on mount. Returns
// one or more textual identifiers of well-known content types (typically
// prefixed with "x-content/"), e.g. x-content/image-dcf for camera memory
// cards. See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This is an asynchronous operation (see g_mount_guess_content_type_sync() for
// the synchronous version), and is finished by calling
// g_mount_guess_content_type_finish() with the mount and Result data returned
// in the callback.
func (mount *Mount) GuessContentType(ctx context.Context, forceRescan bool, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.gboolean            // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if forceRescan {
		_arg1 = C.TRUE
	}
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_guess_content_type(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(forceRescan)
	runtime.KeepAlive(callback)
}

// GuessContentTypeFinish finishes guessing content types of mount. If any
// errors occurred during the operation, error will be set to contain the errors
// and FALSE will be returned. In particular, you may get an
// G_IO_ERROR_NOT_SUPPORTED if the mount does not support content guessing.
func (mount *Mount) GuessContentTypeFinish(result AsyncResulter) ([]string, error) {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cret **C.gchar       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_mount_guess_content_type_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _utf8s []string // out
	var _goerr error    // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8s, _goerr
}

// GuessContentTypeSync tries to guess the type of content stored on mount.
// Returns one or more textual identifiers of well-known content types
// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for camera
// memory cards. See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This is a synchronous operation and as such may block doing IO; see
// g_mount_guess_content_type() for the asynchronous version.
func (mount *Mount) GuessContentTypeSync(ctx context.Context, forceRescan bool) ([]string, error) {
	var _arg0 *C.GMount       // out
	var _arg2 *C.GCancellable // out
	var _arg1 C.gboolean      // out
	var _cret **C.gchar       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if forceRescan {
		_arg1 = C.TRUE
	}

	_cret = C.g_mount_guess_content_type_sync(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(forceRescan)

	var _utf8s []string // out
	var _goerr error    // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8s, _goerr
}

// IsShadowed determines if mount is shadowed. Applications or libraries should
// avoid displaying mount in the user interface if it is shadowed.
//
// A mount is said to be shadowed if there exists one or more user visible
// objects (currently #GMount objects) with a root that is inside the root of
// mount.
//
// One application of shadow mounts is when exposing a single file system that
// is used to address several logical volumes. In this situation, a Monitor
// implementation would create two #GVolume objects (for example, one for the
// camera functionality of the device and one for a SD card reader on the
// device) with activation URIs gphoto2://[usb:001,002]/store1/ and
// gphoto2://[usb:001,002]/store2/. When the underlying mount (with root
// gphoto2://[usb:001,002]/) is mounted, said Monitor implementation would
// create two #GMount objects (each with their root matching the corresponding
// volume activation root) that would shadow the original mount.
//
// The proxy monitor in GVfs 2.26 and later, automatically creates and manage
// shadow mounts (and shadows the underlying mount) if the activation root on a
// #GVolume is set.
func (mount *Mount) IsShadowed() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	_cret = C.g_mount_is_shadowed(_arg0)
	runtime.KeepAlive(mount)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remount remounts a mount. This is an asynchronous operation, and is finished
// by calling g_mount_remount_finish() with the mount and Results data returned
// in the callback.
//
// Remounting is useful when some setting affecting the operation of the volume
// has been changed, as these may need a remount to take affect. While this is
// semantically equivalent with unmounting and then remounting not all backends
// might need to actually be unmounted.
func (mount *Mount) Remount(ctx context.Context, flags MountMountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountMountFlags    // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountMountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_remount(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// RemountFinish finishes remounting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
func (mount *Mount) RemountFinish(result AsyncResulter) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_remount_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Shadow increments the shadow count on mount. Usually used by Monitor
// implementations when creating a shadow mount for mount, see
// g_mount_is_shadowed() for more information. The caller will need to emit the
// #GMount::changed signal on mount manually.
func (mount *Mount) Shadow() {
	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	C.g_mount_shadow(_arg0)
	runtime.KeepAlive(mount)
}

// Unmount unmounts a mount. This is an asynchronous operation, and is finished
// by calling g_mount_unmount_finish() with the mount and Result data returned
// in the callback.
//
// Deprecated: Use g_mount_unmount_with_operation() instead.
func (mount *Mount) Unmount(ctx context.Context, flags MountUnmountFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_unmount(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// UnmountFinish finishes unmounting a mount. If any errors occurred during the
// operation, error will be set to contain the errors and FALSE will be
// returned.
//
// Deprecated: Use g_mount_unmount_with_operation_finish() instead.
func (mount *Mount) UnmountFinish(result AsyncResulter) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_unmount_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// UnmountWithOperation unmounts a mount. This is an asynchronous operation, and
// is finished by calling g_mount_unmount_with_operation_finish() with the mount
// and Result data returned in the callback.
func (mount *Mount) UnmountWithOperation(ctx context.Context, flags MountUnmountFlags, mountOperation *MountOperation, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GMountUnmountFlags(flags)
	if mountOperation != nil {
		_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_mount_unmount_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mountOperation)
	runtime.KeepAlive(callback)
}

// UnmountWithOperationFinish finishes unmounting a mount. If any errors
// occurred during the operation, error will be set to contain the errors and
// FALSE will be returned.
func (mount *Mount) UnmountWithOperationFinish(result AsyncResulter) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_unmount_with_operation_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(mount)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Unshadow decrements the shadow count on mount. Usually used by Monitor
// implementations when destroying a shadow mount for mount, see
// g_mount_is_shadowed() for more information. The caller will need to emit the
// #GMount::changed signal on mount manually.
func (mount *Mount) Unshadow() {
	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(mount.Native()))

	C.g_mount_unshadow(_arg0)
	runtime.KeepAlive(mount)
}
