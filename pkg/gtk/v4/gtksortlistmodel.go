// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_sort_list_model_get_type()), F: marshalSortListModel},
	})
}

// SortListModel: `GtkSortListModel` is a list model that sorts the elements of
// the underlying model according to a `GtkSorter`.
//
// The model can be set up to do incremental sorting, so that sorting long lists
// doesn't block the UI. See [method@Gtk.SortListModel.set_incremental] for
// details.
//
// `GtkSortListModel` is a generic model and because of that it cannot take
// advantage of any external knowledge when sorting. If you run into performance
// issues with `GtkSortListModel`, it is strongly recommended that you write
// your own sorting list model.
type SortListModel interface {
	gextras.Objector

	// Incremental returns whether incremental sorting is enabled.
	//
	// See [method@Gtk.SortListModel.set_incremental].
	Incremental() bool
	// Pending estimates progress of an ongoing sorting operation.
	//
	// The estimate is the number of items that would still need to be sorted to
	// finish the sorting operation if this was a linear algorithm. So this
	// number is not related to how many items are already correctly sorted.
	//
	// If you want to estimate the progress, you can use code like this: “`c
	// pending = gtk_sort_list_model_get_pending (self); model =
	// gtk_sort_list_model_get_model (self); progress = 1.0 - pending / (double)
	// MAX (1, g_list_model_get_n_items (model)); “`
	//
	// If no sort operation is ongoing - in particular when
	// [property@Gtk.SortListModel:incremental] is false - this function returns
	// 0.
	Pending() uint
	// Sorter gets the sorter that is used to sort @self.
	Sorter() *SorterClass
	// SetIncremental sets the sort model to do an incremental sort.
	//
	// When incremental sorting is enabled, the `GtkSortListModel` will not do a
	// complete sort immediately, but will instead queue an idle handler that
	// incrementally sorts the items towards their correct position. This of
	// course means that items do not instantly appear in the right place. It
	// also means that the total sorting time is a lot slower.
	//
	// When your filter blocks the UI while sorting, you might consider turning
	// this on. Depending on your model and sorters, this may become interesting
	// around 10,000 to 100,000 items.
	//
	// By default, incremental sorting is disabled.
	//
	// See [method@Gtk.SortListModel.get_pending] for progress information about
	// an ongoing incremental sorting operation.
	SetIncremental(incremental bool)
	// SetSorter sets a new sorter on @self.
	SetSorter(sorter Sorter)
}

// SortListModelClass implements the SortListModel interface.
type SortListModelClass struct {
	*externglib.Object
}

var _ SortListModel = (*SortListModelClass)(nil)

func wrapSortListModel(obj *externglib.Object) SortListModel {
	return &SortListModelClass{
		Object: obj,
	}
}

func marshalSortListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSortListModel(obj), nil
}

// Incremental returns whether incremental sorting is enabled.
//
// See [method@Gtk.SortListModel.set_incremental].
func (s *SortListModelClass) Incremental() bool {
	var _arg0 *C.GtkSortListModel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer((&SortListModel).Native()))

	_cret = C.gtk_sort_list_model_get_incremental(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Pending estimates progress of an ongoing sorting operation.
//
// The estimate is the number of items that would still need to be sorted to
// finish the sorting operation if this was a linear algorithm. So this number
// is not related to how many items are already correctly sorted.
//
// If you want to estimate the progress, you can use code like this: “`c pending
// = gtk_sort_list_model_get_pending (self); model =
// gtk_sort_list_model_get_model (self); progress = 1.0 - pending / (double) MAX
// (1, g_list_model_get_n_items (model)); “`
//
// If no sort operation is ongoing - in particular when
// [property@Gtk.SortListModel:incremental] is false - this function returns 0.
func (s *SortListModelClass) Pending() uint {
	var _arg0 *C.GtkSortListModel // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer((&SortListModel).Native()))

	_cret = C.gtk_sort_list_model_get_pending(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Sorter gets the sorter that is used to sort @self.
func (s *SortListModelClass) Sorter() *SorterClass {
	var _arg0 *C.GtkSortListModel // out
	var _cret *C.GtkSorter        // in

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer((&SortListModel).Native()))

	_cret = C.gtk_sort_list_model_get_sorter(_arg0)

	var _sorter *SorterClass // out

	_sorter = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*SorterClass)

	return _sorter
}

// SetIncremental sets the sort model to do an incremental sort.
//
// When incremental sorting is enabled, the `GtkSortListModel` will not do a
// complete sort immediately, but will instead queue an idle handler that
// incrementally sorts the items towards their correct position. This of course
// means that items do not instantly appear in the right place. It also means
// that the total sorting time is a lot slower.
//
// When your filter blocks the UI while sorting, you might consider turning this
// on. Depending on your model and sorters, this may become interesting around
// 10,000 to 100,000 items.
//
// By default, incremental sorting is disabled.
//
// See [method@Gtk.SortListModel.get_pending] for progress information about an
// ongoing incremental sorting operation.
func (s *SortListModelClass) SetIncremental(incremental bool) {
	var _arg0 *C.GtkSortListModel // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer((&SortListModel).Native()))
	if incremental {
		_arg1 = C.TRUE
	}

	C.gtk_sort_list_model_set_incremental(_arg0, _arg1)
}

// SetSorter sets a new sorter on @self.
func (s *SortListModelClass) SetSorter(sorter Sorter) {
	var _arg0 *C.GtkSortListModel // out
	var _arg1 *C.GtkSorter        // out

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer((&SortListModel).Native()))
	_arg1 = (*C.GtkSorter)(unsafe.Pointer((&Sorter).Native()))

	C.gtk_sort_list_model_set_sorter(_arg0, _arg1)
}
