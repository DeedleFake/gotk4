// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_spinner_get_type()), F: marshalCellRendererSpinnerer},
	})
}

// CellRendererSpinner renders a spinning animation in a cell, very similar to
// Spinner. It can often be used as an alternative to a CellRendererProgress for
// displaying indefinite activity, instead of actual progress.
//
// To start the animation in a cell, set the CellRendererSpinner:active property
// to TRUE and increment the CellRendererSpinner:pulse property at regular
// intervals. The usual way to set the cell renderer properties for each cell is
// to bind them to columns in your tree model using e.g.
// gtk_tree_view_column_add_attribute().
type CellRendererSpinner struct {
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererSpinner)(nil)
)

func wrapCellRendererSpinner(obj *externglib.Object) *CellRendererSpinner {
	return &CellRendererSpinner{
		CellRenderer: CellRenderer{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererSpinnerer(p uintptr) (interface{}, error) {
	return wrapCellRendererSpinner(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellRendererSpinner returns a new cell renderer which will show a spinner
// to indicate activity.
//
// The function returns the following values:
//
//    - cellRendererSpinner: new CellRenderer.
//
func NewCellRendererSpinner() *CellRendererSpinner {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_spinner_new()

	var _cellRendererSpinner *CellRendererSpinner // out

	_cellRendererSpinner = wrapCellRendererSpinner(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererSpinner
}
