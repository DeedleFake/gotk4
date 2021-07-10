// Code generated by girgen. DO NOT EDIT.

package gdk

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
import "C"

// ThreadsEnter: this function marks the beginning of a critical section in
// which GDK and GTK+ functions can be called safely and without causing race
// conditions. Only one thread at a time can be in such a critial section.
//
// Deprecated: All GDK and GTK+ calls should be made from the main thread.
func ThreadsEnter() {
	C.gdk_threads_enter()
}

// ThreadsInit initializes GDK so that it can be used from multiple threads in
// conjunction with gdk_threads_enter() and gdk_threads_leave().
//
// This call must be made before any use of the main loop from GTK+; to be safe,
// call it before gtk_init().
//
// Deprecated: All GDK and GTK+ calls should be made from the main thread.
func ThreadsInit() {
	C.gdk_threads_init()
}

// ThreadsLeave leaves a critical region begun with gdk_threads_enter().
//
// Deprecated: All GDK and GTK+ calls should be made from the main thread.
func ThreadsLeave() {
	C.gdk_threads_leave()
}
