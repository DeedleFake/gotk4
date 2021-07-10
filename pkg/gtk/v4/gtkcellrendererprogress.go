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
		{T: externglib.Type(C.gtk_cell_renderer_progress_get_type()), F: marshalCellRendererProgresser},
	})
}

// CellRendererProgresser describes CellRendererProgress's methods.
type CellRendererProgresser interface {
	gextras.Objector

	privateCellRendererProgress()
}

// CellRendererProgress renders numbers as progress bars
//
// CellRendererProgress renders a numeric value as a progress par in a cell.
// Additionally, it can display a text on top of the progress bar.
type CellRendererProgress struct {
	CellRenderer

	Orientable
}

var _ CellRendererProgresser = (*CellRendererProgress)(nil)

func wrapCellRendererProgresser(obj *externglib.Object) CellRendererProgresser {
	return &CellRendererProgress{
		CellRenderer: CellRenderer{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellRendererProgresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRendererProgresser(obj), nil
}

// NewCellRendererProgress creates a new CellRendererProgress.
func NewCellRendererProgress() *CellRendererProgress {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_progress_new()

	var _cellRendererProgress *CellRendererProgress // out

	_cellRendererProgress = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*CellRendererProgress)

	return _cellRendererProgress
}

func (*CellRendererProgress) privateCellRendererProgress() {}
