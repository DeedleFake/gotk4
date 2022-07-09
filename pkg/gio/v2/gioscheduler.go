// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// IOSchedulerCancelAllJobs cancels all cancellable I/O jobs.
//
// A job is cancellable if a #GCancellable was passed into
// g_io_scheduler_push_job().
//
// Deprecated: You should never call this function, since you don't know how
// other libraries in your program might be making use of gioscheduler.
func IOSchedulerCancelAllJobs() {
	_info := girepository.MustFind("Gio", "io_scheduler_cancel_all_jobs")
	_info.Invoke(nil, nil)
}
