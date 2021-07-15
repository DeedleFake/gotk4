// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_progress_get_type()), F: marshalCellRendererProgresser},
	})
}

// CellRendererProgress renders a numeric value as a progress par in a cell.
// Additionally, it can display a text on top of the progress bar.
//
// The CellRendererProgress cell renderer was added in GTK+ 2.6.
type CellRendererProgress struct {
	CellRenderer

	Orientable
}

var _ gextras.Nativer = (*CellRendererProgress)(nil)

func wrapCellRendererProgress(obj *externglib.Object) *CellRendererProgress {
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
	return wrapCellRendererProgress(obj), nil
}

// NewCellRendererProgress creates a new CellRendererProgress.
func NewCellRendererProgress() *CellRendererProgress {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_progress_new()

	var _cellRendererProgress *CellRendererProgress // out

	_cellRendererProgress = wrapCellRendererProgress(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererProgress
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *CellRendererProgress) Native() uintptr {
	return v.CellRenderer.InitiallyUnowned.Object.Native()
}

func (*CellRendererProgress) privateCellRendererProgress() {}
