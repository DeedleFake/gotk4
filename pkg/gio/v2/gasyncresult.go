// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GObject* _gotk4_gio2_AsyncResultIface_get_source_object(GAsyncResult*);
// extern gboolean _gotk4_gio2_AsyncResultIface_is_tagged(GAsyncResult*, gpointer);
// extern gpointer _gotk4_gio2_AsyncResultIface_get_user_data(GAsyncResult*);
import "C"

// GType values.
var (
	GTypeAsyncResult = coreglib.Type(C.g_async_result_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAsyncResult, F: marshalAsyncResult},
	})
}

// AsyncResultOverrider contains methods that are overridable.
type AsyncResultOverrider interface {
	// SourceObject gets the source object from a Result.
	//
	// The function returns the following values:
	//
	//    - object (optional): new reference to the source object for the res, or
	//      NULL if there is none.
	//
	SourceObject() *coreglib.Object
	// UserData gets the user data from a Result.
	//
	// The function returns the following values:
	//
	//    - gpointer (optional): user data for res.
	//
	UserData() unsafe.Pointer
	// IsTagged checks if res has the given source_tag (generally a function
	// pointer indicating the function res was created by).
	//
	// The function takes the following parameters:
	//
	//    - sourceTag (optional): application-defined tag.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if res has the indicated source_tag, FALSE if not.
	//
	IsTagged(sourceTag unsafe.Pointer) bool
}

// AsyncResult provides a base class for implementing asynchronous function
// results.
//
// Asynchronous operations are broken up into two separate operations which are
// chained together by a ReadyCallback. To begin an asynchronous operation,
// provide a ReadyCallback to the asynchronous function. This callback will be
// triggered when the operation has completed, and must be run in a later
// iteration of the [thread-default main
// context][g-main-context-push-thread-default] from where the operation was
// initiated. It will be passed a Result instance filled with the details of the
// operation's success or failure, the object the asynchronous function was
// started for and any error codes returned. The asynchronous callback function
// is then expected to call the corresponding "_finish()" function, passing the
// object the function was called for, the Result instance, and (optionally) an
// error to grab any error conditions that may have occurred.
//
// The "_finish()" function for an operation takes the generic result (of type
// Result) and returns the specific result that the operation in question yields
// (e.g. a Enumerator for a "enumerate children" operation). If the result or
// error status of the operation is not needed, there is no need to call the
// "_finish()" function; GIO will take care of cleaning up the result and error
// information after the ReadyCallback returns. You can pass NULL for the
// ReadyCallback if you don't need to take any action at all after the operation
// completes. Applications may also take a reference to the Result and call
// "_finish()" later; however, the "_finish()" function may be called at most
// once.
//
// Example of a typical asynchronous operation flow:
//
//    void _theoretical_frobnitz_async (Theoretical         *t,
//                                      GCancellable        *c,
//                                      GAsyncReadyCallback  cb,
//                                      gpointer             u);
//
//    gboolean _theoretical_frobnitz_finish (Theoretical   *t,
//                                           GAsyncResult  *res,
//                                           GError       **e);
//
//    static void
//    frobnitz_result_func (GObject      *source_object,
//    		 GAsyncResult *res,
//    		 gpointer      user_data)
//    {
//      gboolean success = FALSE;
//
//      success = _theoretical_frobnitz_finish (source_object, res, NULL);
//
//      if (success)
//        g_printf ("Hurray!\n");
//      else
//        g_printf ("Uh oh!\n");
//
//      ...
//
//    }
//
//    int main (int argc, void *argv[])
//    {
//       ...
//
//       _theoretical_frobnitz_async (theoretical_data,
//                                    NULL,
//                                    frobnitz_result_func,
//                                    NULL);
//
//       ...
//    }
//
// The callback for an asynchronous operation is called only once, and is always
// called, even in the case of a cancelled operation. On cancellation the result
// is a G_IO_ERROR_CANCELLED error.
//
// I/O Priority
//
// Many I/O-related asynchronous operations have a priority parameter, which is
// used in certain cases to determine the order in which operations are
// executed. They are not used to determine system-wide I/O scheduling.
// Priorities are integers, with lower numbers indicating higher priority. It is
// recommended to choose priorities between G_PRIORITY_LOW and G_PRIORITY_HIGH,
// with G_PRIORITY_DEFAULT as a default.
//
// AsyncResult wraps an interface. This means the user can get the
// underlying type by calling Cast().
type AsyncResult struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*AsyncResult)(nil)
)

// AsyncResulter describes AsyncResult's interface methods.
type AsyncResulter interface {
	coreglib.Objector

	// SourceObject gets the source object from a Result.
	SourceObject() *coreglib.Object
	// UserData gets the user data from a Result.
	UserData() unsafe.Pointer
	// IsTagged checks if res has the given source_tag (generally a function
	// pointer indicating the function res was created by).
	IsTagged(sourceTag unsafe.Pointer) bool
	// LegacyPropagateError: if res is a AsyncResult, this is equivalent to
	// g_simple_async_result_propagate_error().
	LegacyPropagateError() error
}

var _ AsyncResulter = (*AsyncResult)(nil)

func ifaceInitAsyncResulter(gifacePtr, data C.gpointer) {
	iface := (*C.GAsyncResultIface)(unsafe.Pointer(gifacePtr))
	iface.get_source_object = (*[0]byte)(C._gotk4_gio2_AsyncResultIface_get_source_object)
	iface.get_user_data = (*[0]byte)(C._gotk4_gio2_AsyncResultIface_get_user_data)
	iface.is_tagged = (*[0]byte)(C._gotk4_gio2_AsyncResultIface_is_tagged)
}

//export _gotk4_gio2_AsyncResultIface_get_source_object
func _gotk4_gio2_AsyncResultIface_get_source_object(arg0 *C.GAsyncResult) (cret *C.GObject) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(AsyncResultOverrider)

	object := iface.SourceObject()

	if object != nil {
		cret = (*C.GObject)(unsafe.Pointer(object.Native()))
		C.g_object_ref(C.gpointer(object.Native()))
	}

	return cret
}

//export _gotk4_gio2_AsyncResultIface_get_user_data
func _gotk4_gio2_AsyncResultIface_get_user_data(arg0 *C.GAsyncResult) (cret C.gpointer) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(AsyncResultOverrider)

	gpointer := iface.UserData()

	cret = (C.gpointer)(unsafe.Pointer(gpointer))

	return cret
}

//export _gotk4_gio2_AsyncResultIface_is_tagged
func _gotk4_gio2_AsyncResultIface_is_tagged(arg0 *C.GAsyncResult, arg1 C.gpointer) (cret C.gboolean) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(AsyncResultOverrider)

	var _sourceTag unsafe.Pointer // out

	_sourceTag = (unsafe.Pointer)(unsafe.Pointer(arg1))

	ok := iface.IsTagged(_sourceTag)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapAsyncResult(obj *coreglib.Object) *AsyncResult {
	return &AsyncResult{
		Object: obj,
	}
}

func marshalAsyncResult(p uintptr) (interface{}, error) {
	return wrapAsyncResult(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// SourceObject gets the source object from a Result.
//
// The function returns the following values:
//
//    - object (optional): new reference to the source object for the res, or
//      NULL if there is none.
//
func (res *AsyncResult) SourceObject() *coreglib.Object {
	var _arg0 *C.GAsyncResult // out
	var _cret *C.GObject      // in

	_arg0 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_cret = C.g_async_result_get_source_object(_arg0)
	runtime.KeepAlive(res)

	var _object *coreglib.Object // out

	if _cret != nil {
		_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))
	}

	return _object
}

// UserData gets the user data from a Result.
//
// The function returns the following values:
//
//    - gpointer (optional): user data for res.
//
func (res *AsyncResult) UserData() unsafe.Pointer {
	var _arg0 *C.GAsyncResult // out
	var _cret C.gpointer      // in

	_arg0 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	_cret = C.g_async_result_get_user_data(_arg0)
	runtime.KeepAlive(res)

	var _gpointer unsafe.Pointer // out

	_gpointer = (unsafe.Pointer)(unsafe.Pointer(_cret))

	return _gpointer
}

// IsTagged checks if res has the given source_tag (generally a function pointer
// indicating the function res was created by).
//
// The function takes the following parameters:
//
//    - sourceTag (optional): application-defined tag.
//
// The function returns the following values:
//
//    - ok: TRUE if res has the indicated source_tag, FALSE if not.
//
func (res *AsyncResult) IsTagged(sourceTag unsafe.Pointer) bool {
	var _arg0 *C.GAsyncResult // out
	var _arg1 C.gpointer      // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(sourceTag))

	_cret = C.g_async_result_is_tagged(_arg0, _arg1)
	runtime.KeepAlive(res)
	runtime.KeepAlive(sourceTag)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LegacyPropagateError: if res is a AsyncResult, this is equivalent to
// g_simple_async_result_propagate_error(). Otherwise it returns FALSE.
//
// This can be used for legacy error handling in async *_finish() wrapper
// functions that traditionally handled AsyncResult error returns themselves
// rather than calling into the virtual method. This should not be used in new
// code; Result errors that are set by virtual methods should also be extracted
// by virtual methods, to enable subclasses to chain up correctly.
func (res *AsyncResult) LegacyPropagateError() error {
	var _arg0 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(res).Native()))

	C.g_async_result_legacy_propagate_error(_arg0, &_cerr)
	runtime.KeepAlive(res)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
