// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// TargetTableNewFromList: this function creates an TargetEntry array that
// contains the same targets as the passed list. The returned table is newly
// allocated and should be freed using gtk_target_table_free() when no longer
// needed.
//
// The function takes the following parameters:
//
//    - list: TargetList.
//
// The function returns the following values:
//
//    - targetEntrys: new table.
//
func TargetTableNewFromList(list *TargetList) []TargetEntry {
	var _arg1 *C.GtkTargetList  // out
	var _cret *C.GtkTargetEntry // in
	var _arg2 C.gint            // in

	_arg1 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(list)))

	_cret = C.gtk_target_table_new_from_list(_arg1, &_arg2)
	runtime.KeepAlive(list)

	var _targetEntrys []TargetEntry // out

	defer C.free(unsafe.Pointer(_cret))
	{
		src := unsafe.Slice((*C.GtkTargetEntry)(_cret), _arg2)
		_targetEntrys = make([]TargetEntry, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_targetEntrys[i] = *(*TargetEntry)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(&_targetEntrys[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.gtk_target_entry_free((*C.GtkTargetEntry)(intern.C))
				},
			)
		}
	}

	return _targetEntrys
}
