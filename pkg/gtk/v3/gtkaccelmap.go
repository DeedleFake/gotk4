// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_AccelMap_ConnectChanged(gpointer, gchar*, guint, GdkModifierType, guintptr);
import "C"

// GType values.
var (
	GTypeAccelMap = coreglib.Type(C.gtk_accel_map_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAccelMap, F: marshalAccelMap},
	})
}

// AccelMap: accelerator maps are used to define runtime configurable
// accelerators. Functions for manipulating them are are usually used by higher
// level convenience mechanisms like UIManager and are thus considered
// “low-level”. You’ll want to use them if you’re manually creating menus that
// should have user-configurable accelerators.
//
// An accelerator is uniquely defined by:
//
// - accelerator path
//
// - accelerator key
//
// - accelerator modifiers
//
// The accelerator path must consist of
// “<WINDOWTYPE>/Category1/Category2/.../Action”, where WINDOWTYPE should be a
// unique application-specific identifier that corresponds to the kind of window
// the accelerator is being used in, e.g. “Gimp-Image”, “Abiword-Document” or
// “Gnumeric-Settings”. The “Category1/.../Action” portion is most appropriately
// chosen by the action the accelerator triggers, i.e. for accelerators on menu
// items, choose the item’s menu path, e.g. “File/Save As”, “Image/View/Zoom” or
// “Edit/Select All”. So a full valid accelerator path may look like:
// “<Gimp-Toolbox>/File/Dialogs/Tool Options...”.
//
// All accelerators are stored inside one global AccelMap that can be obtained
// using gtk_accel_map_get(). See [Monitoring changes][monitoring-changes] for
// additional details.
//
//
// Manipulating accelerators
//
// New accelerators can be added using gtk_accel_map_add_entry(). To search for
// specific accelerator, use gtk_accel_map_lookup_entry(). Modifications of
// existing accelerators should be done using gtk_accel_map_change_entry().
//
// In order to avoid having some accelerators changed, they can be locked using
// gtk_accel_map_lock_path(). Unlocking is done using
// gtk_accel_map_unlock_path().
//
//
// Saving and loading accelerator maps
//
// Accelerator maps can be saved to and loaded from some external resource. For
// simple saving and loading from file, gtk_accel_map_save() and
// gtk_accel_map_load() are provided. Saving and loading can also be done by
// providing file descriptor to gtk_accel_map_save_fd() and
// gtk_accel_map_load_fd().
//
//
// Monitoring changes
//
// AccelMap object is only useful for monitoring changes of accelerators. By
// connecting to AccelMap::changed signal, one can monitor changes of all
// accelerators. It is also possible to monitor only single accelerator path by
// using it as a detail of the AccelMap::changed signal.
type AccelMap struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*AccelMap)(nil)
)

func wrapAccelMap(obj *coreglib.Object) *AccelMap {
	return &AccelMap{
		Object: obj,
	}
}

func marshalAccelMap(p uintptr) (interface{}, error) {
	return wrapAccelMap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged notifies of a change in the global accelerator map. The path
// is also used as the detail for the signal, so it is possible to connect to
// changed::accel_path.
func (v *AccelMap) ConnectChanged(f func(accelPath string, accelKey uint, accelMods gdk.ModifierType)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk3_AccelMap_ConnectChanged), f)
}

// AccelMapAddEntry registers a new accelerator with the global accelerator map.
// This function should only be called once per accel_path with the canonical
// accel_key and accel_mods for this path. To change the accelerator during
// runtime programatically, use gtk_accel_map_change_entry().
//
// Set accel_key and accel_mods to 0 to request a removal of the accelerator.
//
// Note that accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
//
// The function takes the following parameters:
//
//    - accelPath: valid accelerator path.
//    - accelKey: accelerator key.
//    - accelMods: accelerator modifiers.
//
func AccelMapAddEntry(accelPath string, accelKey uint, accelMods gdk.ModifierType) {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(accelKey)
	_arg3 = C.GdkModifierType(accelMods)

	C.gtk_accel_map_add_entry(_arg1, _arg2, _arg3)
	runtime.KeepAlive(accelPath)
	runtime.KeepAlive(accelKey)
	runtime.KeepAlive(accelMods)
}

// AccelMapAddFilter adds a filter to the global list of accel path filters.
//
// Accel map entries whose accel path matches one of the filters are skipped by
// gtk_accel_map_foreach().
//
// This function is intended for GTK+ modules that create their own menus, but
// don’t want them to be saved into the applications accelerator map dump.
//
// The function takes the following parameters:
//
//    - filterPattern: pattern (see Spec).
//
func AccelMapAddFilter(filterPattern string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filterPattern)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_accel_map_add_filter(_arg1)
	runtime.KeepAlive(filterPattern)
}

// AccelMapChangeEntry changes the accel_key and accel_mods currently associated
// with accel_path. Due to conflicts with other accelerators, a change may not
// always be possible, replace indicates whether other accelerators may be
// deleted to resolve such conflicts. A change will only occur if all conflicts
// could be resolved (which might not be the case if conflicting accelerators
// are locked). Successful changes are indicated by a TRUE return value.
//
// Note that accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
//
// The function takes the following parameters:
//
//    - accelPath: valid accelerator path.
//    - accelKey: new accelerator key.
//    - accelMods: new accelerator modifiers.
//    - replace: TRUE if other accelerators may be deleted upon conflicts.
//
// The function returns the following values:
//
//    - ok: TRUE if the accelerator could be changed, FALSE otherwise.
//
func AccelMapChangeEntry(accelPath string, accelKey uint, accelMods gdk.ModifierType, replace bool) bool {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _arg4 C.gboolean        // out
	var _cret C.gboolean        // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(accelKey)
	_arg3 = C.GdkModifierType(accelMods)
	if replace {
		_arg4 = C.TRUE
	}

	_cret = C.gtk_accel_map_change_entry(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(accelPath)
	runtime.KeepAlive(accelKey)
	runtime.KeepAlive(accelMods)
	runtime.KeepAlive(replace)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AccelMapLoad parses a file previously saved with gtk_accel_map_save() for
// accelerator specifications, and propagates them accordingly.
//
// The function takes the following parameters:
//
//    - fileName: file containing accelerator specifications, in the GLib file
//      name encoding.
//
func AccelMapLoad(fileName string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_accel_map_load(_arg1)
	runtime.KeepAlive(fileName)
}

// AccelMapLoadFd: filedescriptor variant of gtk_accel_map_load().
//
// Note that the file descriptor will not be closed by this function.
//
// The function takes the following parameters:
//
//    - fd: valid readable file descriptor.
//
func AccelMapLoadFd(fd int) {
	var _arg1 C.gint // out

	_arg1 = C.gint(fd)

	C.gtk_accel_map_load_fd(_arg1)
	runtime.KeepAlive(fd)
}

// AccelMapLoadScanner variant of gtk_accel_map_load().
//
// The function takes the following parameters:
//
//    - scanner which has already been provided with an input file.
//
func AccelMapLoadScanner(scanner *glib.Scanner) {
	var _arg1 *C.GScanner // out

	_arg1 = (*C.GScanner)(gextras.StructNative(unsafe.Pointer(scanner)))

	C.gtk_accel_map_load_scanner(_arg1)
	runtime.KeepAlive(scanner)
}

// AccelMapLookupEntry looks up the accelerator entry for accel_path and fills
// in key.
//
// The function takes the following parameters:
//
//    - accelPath: valid accelerator path.
//
// The function returns the following values:
//
//    - key (optional): accelerator key to be filled in (optional).
//    - ok: TRUE if accel_path is known, FALSE otherwise.
//
func AccelMapLookupEntry(accelPath string) (*AccelKey, bool) {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkAccelKey // in
	var _cret C.gboolean    // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_accel_map_lookup_entry(_arg1, &_arg2)
	runtime.KeepAlive(accelPath)

	var _key *AccelKey // out
	var _ok bool       // out

	_key = (*AccelKey)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _key, _ok
}

// AccelMapSave saves current accelerator specifications (accelerator path, key
// and modifiers) to file_name. The file is written in a format suitable to be
// read back in by gtk_accel_map_load().
//
// The function takes the following parameters:
//
//    - fileName: name of the file to contain accelerator specifications, in the
//      GLib file name encoding.
//
func AccelMapSave(fileName string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_accel_map_save(_arg1)
	runtime.KeepAlive(fileName)
}

// AccelMapSaveFd: filedescriptor variant of gtk_accel_map_save().
//
// Note that the file descriptor will not be closed by this function.
//
// The function takes the following parameters:
//
//    - fd: valid writable file descriptor.
//
func AccelMapSaveFd(fd int) {
	var _arg1 C.gint // out

	_arg1 = C.gint(fd)

	C.gtk_accel_map_save_fd(_arg1)
	runtime.KeepAlive(fd)
}
