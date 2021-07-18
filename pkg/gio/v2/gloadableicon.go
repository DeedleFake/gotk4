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
	externglib "github.com/gotk3/gotk3/glib"
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
		{T: externglib.Type(C.g_loadable_icon_get_type()), F: marshalLoadableIconer},
	})
}

// LoadableIconOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type LoadableIconOverrider interface {
	// Load loads a loadable icon. For the asynchronous version of this
	// function, see g_loadable_icon_load_async().
	Load(ctx context.Context, size int) (string, InputStreamer, error)
	// LoadAsync loads an icon asynchronously. To finish this function, see
	// g_loadable_icon_load_finish(). For the synchronous, blocking version of
	// this function, see g_loadable_icon_load().
	LoadAsync(ctx context.Context, size int, callback AsyncReadyCallback)
	// LoadFinish finishes an asynchronous icon load started in
	// g_loadable_icon_load_async().
	LoadFinish(res AsyncResulter) (string, InputStreamer, error)
}

// LoadableIcon extends the #GIcon interface and adds the ability to load icons
// from streams.
type LoadableIcon struct {
	Icon
}

var _ gextras.Nativer = (*LoadableIcon)(nil)

// LoadableIconer describes LoadableIcon's abstract methods.
type LoadableIconer interface {
	// Load loads a loadable icon.
	Load(ctx context.Context, size int) (string, InputStreamer, error)
	// LoadAsync loads an icon asynchronously.
	LoadAsync(ctx context.Context, size int, callback AsyncReadyCallback)
	// LoadFinish finishes an asynchronous icon load started in
	// g_loadable_icon_load_async().
	LoadFinish(res AsyncResulter) (string, InputStreamer, error)
}

var _ LoadableIconer = (*LoadableIcon)(nil)

func wrapLoadableIcon(obj *externglib.Object) *LoadableIcon {
	return &LoadableIcon{
		Icon: Icon{
			Object: obj,
		},
	}
}

func marshalLoadableIconer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLoadableIcon(obj), nil
}

// Load loads a loadable icon. For the asynchronous version of this function,
// see g_loadable_icon_load_async().
func (icon *LoadableIcon) Load(ctx context.Context, size int) (string, InputStreamer, error) {
	var _arg0 *C.GLoadableIcon // out
	var _arg3 *C.GCancellable  // out
	var _arg1 C.int            // out
	var _arg2 *C.char          // in
	var _cret *C.GInputStream  // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GLoadableIcon)(unsafe.Pointer(icon.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(size)

	_cret = C.g_loadable_icon_load(_arg0, _arg1, &_arg2, _arg3, &_cerr)

	var _typ string                // out
	var _inputStream InputStreamer // out
	var _goerr error               // out

	_typ = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_inputStream = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(InputStreamer)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _typ, _inputStream, _goerr
}

// LoadAsync loads an icon asynchronously. To finish this function, see
// g_loadable_icon_load_finish(). For the synchronous, blocking version of this
// function, see g_loadable_icon_load().
func (icon *LoadableIcon) LoadAsync(ctx context.Context, size int, callback AsyncReadyCallback) {
	var _arg0 *C.GLoadableIcon      // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.int                 // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GLoadableIcon)(unsafe.Pointer(icon.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(size)
	_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg4 = C.gpointer(gbox.AssignOnce(callback))

	C.g_loadable_icon_load_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// LoadFinish finishes an asynchronous icon load started in
// g_loadable_icon_load_async().
func (icon *LoadableIcon) LoadFinish(res AsyncResulter) (string, InputStreamer, error) {
	var _arg0 *C.GLoadableIcon // out
	var _arg1 *C.GAsyncResult  // out
	var _arg2 *C.char          // in
	var _cret *C.GInputStream  // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GLoadableIcon)(unsafe.Pointer(icon.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((res).(gextras.Nativer).Native()))

	_cret = C.g_loadable_icon_load_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _typ string                // out
	var _inputStream InputStreamer // out
	var _goerr error               // out

	_typ = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_inputStream = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(InputStreamer)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _typ, _inputStream, _goerr
}
