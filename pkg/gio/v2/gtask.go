// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
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
		{T: externglib.Type(C.g_task_get_type()), F: marshalTask},
	})
}

// Task represents and manages a cancellable "task".
//
//
// Asynchronous operations
//
// The most common usage of #GTask is as a Result, to manage data during an
// asynchronous operation. You call g_task_new() in the "start" method, followed
// by g_task_set_task_data() and the like if you need to keep some additional
// data associated with the task, and then pass the task object around through
// your asynchronous operation. Eventually, you will call a method such as
// g_task_return_pointer() or g_task_return_error(), which will save the value
// you give it and then invoke the task's callback function in the
// [thread-default main context][g-main-context-push-thread-default] where it
// was created (waiting until the next iteration of the main loop first, if
// necessary). The caller will pass the #GTask back to the operation's finish
// function (as a Result), and you can use g_task_propagate_pointer() or the
// like to extract the return value.
//
// Here is an example for using GTask as a GAsyncResult:
//
//        static void
//        bake_cake_thread (GTask         *task,
//                          gpointer       source_object,
//                          gpointer       task_data,
//                          GCancellable  *cancellable)
//        {
//          Baker *self = source_object;
//          CakeData *cake_data = task_data;
//          Cake *cake;
//          GError *error = NULL;
//
//          cake = bake_cake (baker, cake_data->radius, cake_data->flavor,
//                            cake_data->frosting, cake_data->message,
//                            &error);
//          if (error)
//            {
//              g_task_return_error (task, error);
//              return;
//            }
//
//          // If the task has already been cancelled, then we don't want to add
//          // the cake to the cake cache. Likewise, we don't  want to have the
//          // task get cancelled in the middle of updating the cache.
//          // g_task_set_return_on_cancel() will return true here if it managed
//          // to disable return-on-cancel, or false if the task was cancelled
//          // before it could.
//          if (g_task_set_return_on_cancel (task, FALSE))
//            {
//              // If the caller cancels at this point, their
//              // GAsyncReadyCallback won't be invoked until we return,
//              // so we don't have to worry that this code will run at
//              // the same time as that code does. But if there were
//              // other functions that might look at the cake cache,
//              // then we'd probably need a GMutex here as well.
//              baker_add_cake_to_cache (baker, cake);
//              g_task_return_pointer (task, cake, g_object_unref);
//            }
//        }
//
//        void
//        baker_bake_cake_async (Baker               *self,
//                               guint                radius,
//                               CakeFlavor           flavor,
//                               CakeFrostingType     frosting,
//                               const char          *message,
//                               GCancellable        *cancellable,
//                               GAsyncReadyCallback  callback,
//                               gpointer             user_data)
//        {
//          CakeData *cake_data;
//          GTask *task;
//
//          cake_data = g_slice_new (CakeData);
//
//          ...
//
//          task = g_task_new (self, cancellable, callback, user_data);
//          g_task_set_task_data (task, cake_data, (GDestroyNotify) cake_data_free);
//          g_task_set_return_on_cancel (task, TRUE);
//          g_task_run_in_thread (task, bake_cake_thread);
//        }
//
//        Cake *
//        baker_bake_cake_sync (Baker               *self,
//                              guint                radius,
//                              CakeFlavor           flavor,
//                              CakeFrostingType     frosting,
//                              const char          *message,
//                              GCancellable        *cancellable,
//                              GError             **error)
//        {
//          CakeData *cake_data;
//          GTask *task;
//          Cake *cake;
//
//          cake_data = g_slice_new (CakeData);
//
//          ...
//
//          task = g_task_new (self, cancellable, NULL, NULL);
//          g_task_set_task_data (task, cake_data, (GDestroyNotify) cake_data_free);
//          g_task_set_return_on_cancel (task, TRUE);
//          g_task_run_in_thread_sync (task, bake_cake_thread);
//
//          cake = g_task_propagate_pointer (task, error);
//          g_object_unref (task);
//          return cake;
//        }
//
//
// Porting from GSimpleAsyncResult
//
// #GTask's API attempts to be simpler than AsyncResult's in several ways: - You
// can save task-specific data with g_task_set_task_data(), and retrieve it
// later with g_task_get_task_data(). This replaces the abuse of
// g_simple_async_result_set_op_res_gpointer() for the same purpose with
// AsyncResult. - In addition to the task data, #GTask also keeps track of the
// [priority][io-priority], #GCancellable, and Context associated with the task,
// so tasks that consist of a chain of simpler asynchronous operations will have
// easy access to those values when starting each sub-task. -
// g_task_return_error_if_cancelled() provides simplified handling for
// cancellation. In addition, cancellation overrides any other #GTask return
// value by default, like AsyncResult does when
// g_simple_async_result_set_check_cancellable() is called. (You can use
// g_task_set_check_cancellable() to turn off that behavior.) On the other hand,
// g_task_run_in_thread() guarantees that it will always run your `task_func`,
// even if the task's #GCancellable is already cancelled before the task gets a
// chance to run; you can start your `task_func` with a
// g_task_return_error_if_cancelled() check if you need the old behavior. - The
// "return" methods (eg, g_task_return_pointer()) automatically cause the task
// to be "completed" as well, and there is no need to worry about the "complete"
// vs "complete in idle" distinction. (#GTask automatically figures out whether
// the task's callback can be invoked directly, or if it needs to be sent to
// another Context, or delayed until the next iteration of the current Context.)
// - The "finish" functions for #GTask based operations are generally much
// simpler than AsyncResult ones, normally consisting of only a single call to
// g_task_propagate_pointer() or the like. Since g_task_propagate_pointer()
// "steals" the return value from the #GTask, it is not necessary to juggle
// pointers around to prevent it from being freed twice. - With AsyncResult, it
// was common to call g_simple_async_result_propagate_error() from the
// `_finish()` wrapper function, and have virtual method implementations only
// deal with successful returns. This behavior is deprecated, because it makes
// it difficult for a subclass to chain to a parent class's async methods.
// Instead, the wrapper function should just be a simple wrapper, and the
// virtual method should call an appropriate `g_task_propagate_` function. Note
// that wrapper methods can now use g_async_result_legacy_propagate_error() to
// do old-style AsyncResult error-returning behavior, and
// g_async_result_is_tagged() to check if a result is tagged as having come from
// the `_async()` wrapper function (for "short-circuit" results, such as when
// passing 0 to g_input_stream_read_async()).
type Task interface {
	gextras.Objector

	// Cancellable gets @task's #GCancellable
	Cancellable() *CancellableClass
	// CheckCancellable gets @task's check-cancellable flag. See
	// g_task_set_check_cancellable() for more details.
	CheckCancellable() bool
	// Completed gets the value of #GTask:completed. This changes from false to
	// true after the task’s callback is invoked, and will return false if
	// called from inside the callback.
	Completed() bool
	// Context gets the Context that @task will return its result in (that is,
	// the context that was the [thread-default main
	// context][g-main-context-push-thread-default] at the point when @task was
	// created).
	//
	// This will always return a non-nil value, even if the task's context is
	// the default Context.
	Context() *glib.MainContext
	// Name gets @task’s name. See g_task_set_name().
	Name() string
	// Priority gets @task's priority
	Priority() int
	// ReturnOnCancel gets @task's return-on-cancel flag. See
	// g_task_set_return_on_cancel() for more details.
	ReturnOnCancel() bool
	// SourceObject gets the source object from @task. Like
	// g_async_result_get_source_object(), but does not ref the object.
	SourceObject() *externglib.Object
	// SourceTag gets @task's source tag. See g_task_set_source_tag().
	SourceTag() interface{}
	// TaskData gets @task's `task_data`.
	TaskData() interface{}
	// HadError tests if @task resulted in an error.
	HadError() bool
	// PropagateBoolean gets the result of @task as a #gboolean.
	//
	// If the task resulted in an error, or was cancelled, then this will
	// instead return false and set @error.
	//
	// Since this method transfers ownership of the return value (or error) to
	// the caller, you may only call it once.
	PropagateBoolean() error
	// PropagateInt gets the result of @task as an integer (#gssize).
	//
	// If the task resulted in an error, or was cancelled, then this will
	// instead return -1 and set @error.
	//
	// Since this method transfers ownership of the return value (or error) to
	// the caller, you may only call it once.
	PropagateInt() (int, error)
	// PropagatePointer gets the result of @task as a pointer, and transfers
	// ownership of that value to the caller.
	//
	// If the task resulted in an error, or was cancelled, then this will
	// instead return nil and set @error.
	//
	// Since this method transfers ownership of the return value (or error) to
	// the caller, you may only call it once.
	PropagatePointer() (interface{}, error)
	// PropagateValue gets the result of @task as a #GValue, and transfers
	// ownership of that value to the caller. As with g_task_return_value(),
	// this is a generic low-level method; g_task_propagate_pointer() and the
	// like will usually be more useful for C code.
	//
	// If the task resulted in an error, or was cancelled, then this will
	// instead set @error and return false.
	//
	// Since this method transfers ownership of the return value (or error) to
	// the caller, you may only call it once.
	PropagateValue() (externglib.Value, error)
	// ReturnBoolean sets @task's result to @result and completes the task (see
	// g_task_return_pointer() for more discussion of exactly what this means).
	ReturnBoolean(result bool)
	// ReturnError sets @task's result to @error (which @task assumes ownership
	// of) and completes the task (see g_task_return_pointer() for more
	// discussion of exactly what this means).
	//
	// Note that since the task takes ownership of @error, and since the task
	// may be completed before returning from g_task_return_error(), you cannot
	// assume that @error is still valid after calling this. Call g_error_copy()
	// on the error if you need to keep a local copy as well.
	//
	// See also g_task_return_new_error().
	ReturnError(err error)
	// ReturnErrorIfCancelled checks if @task's #GCancellable has been
	// cancelled, and if so, sets @task's error accordingly and completes the
	// task (see g_task_return_pointer() for more discussion of exactly what
	// this means).
	ReturnErrorIfCancelled() bool
	// ReturnInt sets @task's result to @result and completes the task (see
	// g_task_return_pointer() for more discussion of exactly what this means).
	ReturnInt(result int)
	// ReturnValue sets @task's result to @result (by copying it) and completes
	// the task.
	//
	// If @result is nil then a #GValue of type TYPE_POINTER with a value of nil
	// will be used for the result.
	//
	// This is a very generic low-level method intended primarily for use by
	// language bindings; for C code, g_task_return_pointer() and the like will
	// normally be much easier to use.
	ReturnValue(result *externglib.Value)
	// SetCheckCancellable sets or clears @task's check-cancellable flag. If
	// this is true (the default), then g_task_propagate_pointer(), etc, and
	// g_task_had_error() will check the task's #GCancellable first, and if it
	// has been cancelled, then they will consider the task to have returned an
	// "Operation was cancelled" error (G_IO_ERROR_CANCELLED), regardless of any
	// other error or return value the task may have had.
	//
	// If @check_cancellable is false, then the #GTask will not check the
	// cancellable itself, and it is up to @task's owner to do this (eg, via
	// g_task_return_error_if_cancelled()).
	//
	// If you are using g_task_set_return_on_cancel() as well, then you must
	// leave check-cancellable set true.
	SetCheckCancellable(checkCancellable bool)
	// SetName sets @task’s name, used in debugging and profiling. The name
	// defaults to nil.
	//
	// The task name should describe in a human readable way what the task does.
	// For example, ‘Open file’ or ‘Connect to network host’. It is used to set
	// the name of the #GSource used for idle completion of the task.
	//
	// This function may only be called before the @task is first used in a
	// thread other than the one it was constructed in.
	SetName(name string)
	// SetPriority sets @task's priority. If you do not call this, it will
	// default to G_PRIORITY_DEFAULT.
	//
	// This will affect the priority of #GSources created with
	// g_task_attach_source() and the scheduling of tasks run in threads, and
	// can also be explicitly retrieved later via g_task_get_priority().
	SetPriority(priority int)
	// SetReturnOnCancel sets or clears @task's return-on-cancel flag. This is
	// only meaningful for tasks run via g_task_run_in_thread() or
	// g_task_run_in_thread_sync().
	//
	// If @return_on_cancel is true, then cancelling @task's #GCancellable will
	// immediately cause it to return, as though the task's ThreadFunc had
	// called g_task_return_error_if_cancelled() and then returned.
	//
	// This allows you to create a cancellable wrapper around an uninterruptible
	// function. The ThreadFunc just needs to be careful that it does not modify
	// any externally-visible state after it has been cancelled. To do that, the
	// thread should call g_task_set_return_on_cancel() again to (atomically)
	// set return-on-cancel false before making externally-visible changes; if
	// the task gets cancelled before the return-on-cancel flag could be
	// changed, g_task_set_return_on_cancel() will indicate this by returning
	// false.
	//
	// You can disable and re-enable this flag multiple times if you wish. If
	// the task's #GCancellable is cancelled while return-on-cancel is false,
	// then calling g_task_set_return_on_cancel() to set it true again will
	// cause the task to be cancelled at that point.
	//
	// If the task's #GCancellable is already cancelled before you call
	// g_task_run_in_thread()/g_task_run_in_thread_sync(), then the ThreadFunc
	// will still be run (for consistency), but the task will also be completed
	// right away.
	SetReturnOnCancel(returnOnCancel bool) bool
	// SetSourceTag sets @task's source tag. You can use this to tag a task
	// return value with a particular pointer (usually a pointer to the function
	// doing the tagging) and then later check it using g_task_get_source_tag()
	// (or g_async_result_is_tagged()) in the task's "finish" function, to
	// figure out if the response came from a particular place.
	SetSourceTag(sourceTag interface{})
}

// TaskClass implements the Task interface.
type TaskClass struct {
	*externglib.Object
	AsyncResultIface
}

var _ Task = (*TaskClass)(nil)

func wrapTask(obj *externglib.Object) Task {
	return &TaskClass{
		Object: obj,
		AsyncResultIface: AsyncResultIface{
			Object: obj,
		},
	}
}

func marshalTask(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTask(obj), nil
}

// NewTask creates a #GTask acting on @source_object, which will eventually be
// used to invoke @callback in the current [thread-default main
// context][g-main-context-push-thread-default].
//
// Call this in the "start" method of your asynchronous method, and pass the
// #GTask around throughout the asynchronous operation. You can use
// g_task_set_task_data() to attach task-specific data to the object, which you
// can retrieve later via g_task_get_task_data().
//
// By default, if @cancellable is cancelled, then the return value of the task
// will always be G_IO_ERROR_CANCELLED, even if the task had already completed
// before the cancellation. This allows for simplified handling in cases where
// cancellation may imply that other objects that the task depends on have been
// destroyed. If you do not want this behavior, you can use
// g_task_set_check_cancellable() to change it.
func NewTask(sourceObject gextras.Objector, cancellable Cancellable, callback AsyncReadyCallback) *TaskClass {
	var _arg1 C.gpointer            // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer
	var _cret *C.GTask // in

	_arg1 = (C.gpointer)(unsafe.Pointer(sourceObject.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	_cret = C.g_task_new(_arg1, _arg2, _arg3, _arg4)

	var _task *TaskClass // out

	_task = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TaskClass)

	return _task
}

// Cancellable gets @task's #GCancellable
func (task *TaskClass) Cancellable() *CancellableClass {
	var _arg0 *C.GTask        // out
	var _cret *C.GCancellable // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_cancellable(_arg0)

	var _cancellable *CancellableClass // out

	_cancellable = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*CancellableClass)

	return _cancellable
}

// CheckCancellable gets @task's check-cancellable flag. See
// g_task_set_check_cancellable() for more details.
func (task *TaskClass) CheckCancellable() bool {
	var _arg0 *C.GTask   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_check_cancellable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Completed gets the value of #GTask:completed. This changes from false to true
// after the task’s callback is invoked, and will return false if called from
// inside the callback.
func (task *TaskClass) Completed() bool {
	var _arg0 *C.GTask   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_completed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Context gets the Context that @task will return its result in (that is, the
// context that was the [thread-default main
// context][g-main-context-push-thread-default] at the point when @task was
// created).
//
// This will always return a non-nil value, even if the task's context is the
// default Context.
func (task *TaskClass) Context() *glib.MainContext {
	var _arg0 *C.GTask        // out
	var _cret *C.GMainContext // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_context(_arg0)

	var _mainContext *glib.MainContext // out

	_mainContext = (*glib.MainContext)(unsafe.Pointer(_cret))
	C.g_main_context_ref(_cret)
	runtime.SetFinalizer(_mainContext, func(v *glib.MainContext) {
		C.g_main_context_unref((*C.GMainContext)(unsafe.Pointer(v)))
	})

	return _mainContext
}

// Name gets @task’s name. See g_task_set_name().
func (task *TaskClass) Name() string {
	var _arg0 *C.GTask // out
	var _cret *C.gchar // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Priority gets @task's priority
func (task *TaskClass) Priority() int {
	var _arg0 *C.GTask // out
	var _cret C.gint   // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_priority(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ReturnOnCancel gets @task's return-on-cancel flag. See
// g_task_set_return_on_cancel() for more details.
func (task *TaskClass) ReturnOnCancel() bool {
	var _arg0 *C.GTask   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_return_on_cancel(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SourceObject gets the source object from @task. Like
// g_async_result_get_source_object(), but does not ref the object.
func (task *TaskClass) SourceObject() *externglib.Object {
	var _arg0 *C.GTask   // out
	var _cret C.gpointer // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_source_object(_arg0)

	var _object *externglib.Object // out

	_object = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*externglib.Object)

	return _object
}

// SourceTag gets @task's source tag. See g_task_set_source_tag().
func (task *TaskClass) SourceTag() interface{} {
	var _arg0 *C.GTask   // out
	var _cret C.gpointer // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_source_tag(_arg0)

	var _gpointer interface{} // out

	_gpointer = box.Get(uintptr(_cret))

	return _gpointer
}

// TaskData gets @task's `task_data`.
func (task *TaskClass) TaskData() interface{} {
	var _arg0 *C.GTask   // out
	var _cret C.gpointer // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_get_task_data(_arg0)

	var _gpointer interface{} // out

	_gpointer = box.Get(uintptr(_cret))

	return _gpointer
}

// HadError tests if @task resulted in an error.
func (task *TaskClass) HadError() bool {
	var _arg0 *C.GTask   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_had_error(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PropagateBoolean gets the result of @task as a #gboolean.
//
// If the task resulted in an error, or was cancelled, then this will instead
// return false and set @error.
//
// Since this method transfers ownership of the return value (or error) to the
// caller, you may only call it once.
func (task *TaskClass) PropagateBoolean() error {
	var _arg0 *C.GTask  // out
	var _cerr *C.GError // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	C.g_task_propagate_boolean(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PropagateInt gets the result of @task as an integer (#gssize).
//
// If the task resulted in an error, or was cancelled, then this will instead
// return -1 and set @error.
//
// Since this method transfers ownership of the return value (or error) to the
// caller, you may only call it once.
func (task *TaskClass) PropagateInt() (int, error) {
	var _arg0 *C.GTask  // out
	var _cret C.gssize  // in
	var _cerr *C.GError // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_propagate_int(_arg0, &_cerr)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gssize, _goerr
}

// PropagatePointer gets the result of @task as a pointer, and transfers
// ownership of that value to the caller.
//
// If the task resulted in an error, or was cancelled, then this will instead
// return nil and set @error.
//
// Since this method transfers ownership of the return value (or error) to the
// caller, you may only call it once.
func (task *TaskClass) PropagatePointer() (interface{}, error) {
	var _arg0 *C.GTask   // out
	var _cret C.gpointer // in
	var _cerr *C.GError  // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_propagate_pointer(_arg0, &_cerr)

	var _gpointer interface{} // out
	var _goerr error          // out

	_gpointer = box.Get(uintptr(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gpointer, _goerr
}

// PropagateValue gets the result of @task as a #GValue, and transfers ownership
// of that value to the caller. As with g_task_return_value(), this is a generic
// low-level method; g_task_propagate_pointer() and the like will usually be
// more useful for C code.
//
// If the task resulted in an error, or was cancelled, then this will instead
// set @error and return false.
//
// Since this method transfers ownership of the return value (or error) to the
// caller, you may only call it once.
func (task *TaskClass) PropagateValue() (externglib.Value, error) {
	var _arg0 *C.GTask  // out
	var _arg1 C.GValue  // in
	var _cerr *C.GError // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	C.g_task_propagate_value(_arg0, &_arg1, &_cerr)

	var _value externglib.Value // out
	var _goerr error            // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg1)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _value, _goerr
}

// ReturnBoolean sets @task's result to @result and completes the task (see
// g_task_return_pointer() for more discussion of exactly what this means).
func (task *TaskClass) ReturnBoolean(result bool) {
	var _arg0 *C.GTask   // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	if result {
		_arg1 = C.TRUE
	}

	C.g_task_return_boolean(_arg0, _arg1)
}

// ReturnError sets @task's result to @error (which @task assumes ownership of)
// and completes the task (see g_task_return_pointer() for more discussion of
// exactly what this means).
//
// Note that since the task takes ownership of @error, and since the task may be
// completed before returning from g_task_return_error(), you cannot assume that
// @error is still valid after calling this. Call g_error_copy() on the error if
// you need to keep a local copy as well.
//
// See also g_task_return_new_error().
func (task *TaskClass) ReturnError(err error) {
	var _arg0 *C.GTask  // out
	var _arg1 *C.GError // out

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	_arg1 = (*C.GError)(gerror.New(err))

	C.g_task_return_error(_arg0, _arg1)
}

// ReturnErrorIfCancelled checks if @task's #GCancellable has been cancelled,
// and if so, sets @task's error accordingly and completes the task (see
// g_task_return_pointer() for more discussion of exactly what this means).
func (task *TaskClass) ReturnErrorIfCancelled() bool {
	var _arg0 *C.GTask   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))

	_cret = C.g_task_return_error_if_cancelled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReturnInt sets @task's result to @result and completes the task (see
// g_task_return_pointer() for more discussion of exactly what this means).
func (task *TaskClass) ReturnInt(result int) {
	var _arg0 *C.GTask // out
	var _arg1 C.gssize // out

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	_arg1 = C.gssize(result)

	C.g_task_return_int(_arg0, _arg1)
}

// ReturnValue sets @task's result to @result (by copying it) and completes the
// task.
//
// If @result is nil then a #GValue of type TYPE_POINTER with a value of nil
// will be used for the result.
//
// This is a very generic low-level method intended primarily for use by
// language bindings; for C code, g_task_return_pointer() and the like will
// normally be much easier to use.
func (task *TaskClass) ReturnValue(result *externglib.Value) {
	var _arg0 *C.GTask  // out
	var _arg1 *C.GValue // out

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(&result.GValue))

	C.g_task_return_value(_arg0, _arg1)
}

// SetCheckCancellable sets or clears @task's check-cancellable flag. If this is
// true (the default), then g_task_propagate_pointer(), etc, and
// g_task_had_error() will check the task's #GCancellable first, and if it has
// been cancelled, then they will consider the task to have returned an
// "Operation was cancelled" error (G_IO_ERROR_CANCELLED), regardless of any
// other error or return value the task may have had.
//
// If @check_cancellable is false, then the #GTask will not check the
// cancellable itself, and it is up to @task's owner to do this (eg, via
// g_task_return_error_if_cancelled()).
//
// If you are using g_task_set_return_on_cancel() as well, then you must leave
// check-cancellable set true.
func (task *TaskClass) SetCheckCancellable(checkCancellable bool) {
	var _arg0 *C.GTask   // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	if checkCancellable {
		_arg1 = C.TRUE
	}

	C.g_task_set_check_cancellable(_arg0, _arg1)
}

// SetName sets @task’s name, used in debugging and profiling. The name defaults
// to nil.
//
// The task name should describe in a human readable way what the task does. For
// example, ‘Open file’ or ‘Connect to network host’. It is used to set the name
// of the #GSource used for idle completion of the task.
//
// This function may only be called before the @task is first used in a thread
// other than the one it was constructed in.
func (task *TaskClass) SetName(name string) {
	var _arg0 *C.GTask // out
	var _arg1 *C.gchar // out

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_task_set_name(_arg0, _arg1)
}

// SetPriority sets @task's priority. If you do not call this, it will default
// to G_PRIORITY_DEFAULT.
//
// This will affect the priority of #GSources created with
// g_task_attach_source() and the scheduling of tasks run in threads, and can
// also be explicitly retrieved later via g_task_get_priority().
func (task *TaskClass) SetPriority(priority int) {
	var _arg0 *C.GTask // out
	var _arg1 C.gint   // out

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	_arg1 = C.gint(priority)

	C.g_task_set_priority(_arg0, _arg1)
}

// SetReturnOnCancel sets or clears @task's return-on-cancel flag. This is only
// meaningful for tasks run via g_task_run_in_thread() or
// g_task_run_in_thread_sync().
//
// If @return_on_cancel is true, then cancelling @task's #GCancellable will
// immediately cause it to return, as though the task's ThreadFunc had called
// g_task_return_error_if_cancelled() and then returned.
//
// This allows you to create a cancellable wrapper around an uninterruptible
// function. The ThreadFunc just needs to be careful that it does not modify any
// externally-visible state after it has been cancelled. To do that, the thread
// should call g_task_set_return_on_cancel() again to (atomically) set
// return-on-cancel false before making externally-visible changes; if the task
// gets cancelled before the return-on-cancel flag could be changed,
// g_task_set_return_on_cancel() will indicate this by returning false.
//
// You can disable and re-enable this flag multiple times if you wish. If the
// task's #GCancellable is cancelled while return-on-cancel is false, then
// calling g_task_set_return_on_cancel() to set it true again will cause the
// task to be cancelled at that point.
//
// If the task's #GCancellable is already cancelled before you call
// g_task_run_in_thread()/g_task_run_in_thread_sync(), then the ThreadFunc will
// still be run (for consistency), but the task will also be completed right
// away.
func (task *TaskClass) SetReturnOnCancel(returnOnCancel bool) bool {
	var _arg0 *C.GTask   // out
	var _arg1 C.gboolean // out
	var _cret C.gboolean // in

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	if returnOnCancel {
		_arg1 = C.TRUE
	}

	_cret = C.g_task_set_return_on_cancel(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSourceTag sets @task's source tag. You can use this to tag a task return
// value with a particular pointer (usually a pointer to the function doing the
// tagging) and then later check it using g_task_get_source_tag() (or
// g_async_result_is_tagged()) in the task's "finish" function, to figure out if
// the response came from a particular place.
func (task *TaskClass) SetSourceTag(sourceTag interface{}) {
	var _arg0 *C.GTask   // out
	var _arg1 C.gpointer // out

	_arg0 = (*C.GTask)(unsafe.Pointer(task.Native()))
	_arg1 = (C.gpointer)(box.Assign(sourceTag))

	C.g_task_set_source_tag(_arg0, _arg1)
}
