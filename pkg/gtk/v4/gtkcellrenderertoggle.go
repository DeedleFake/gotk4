// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_CellRendererToggle_ConnectToggled(gpointer, gchar*, guintptr);
import "C"

// GType values.
var (
	GTypeCellRendererToggle = coreglib.Type(C.gtk_cell_renderer_toggle_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererToggle, F: marshalCellRendererToggle},
	})
}

// CellRendererToggle renders a toggle button in a cell
//
// CellRendererToggle renders a toggle button in a cell. The button is drawn as
// a radio or a checkbutton, depending on the CellRendererToggle:radio property.
// When activated, it emits the CellRendererToggle::toggled signal.
type CellRendererToggle struct {
	_ [0]func() // equal guard
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererToggle)(nil)
)

func wrapCellRendererToggle(obj *coreglib.Object) *CellRendererToggle {
	return &CellRendererToggle{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererToggle(p uintptr) (interface{}, error) {
	return wrapCellRendererToggle(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectToggled signal is emitted when the cell is toggled.
//
// It is the responsibility of the application to update the model with the
// correct value to store at path. Often this is simply the opposite of the
// value currently stored at path.
func (toggle *CellRendererToggle) ConnectToggled(f func(path string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(toggle, "toggled", false, unsafe.Pointer(C._gotk4_gtk4_CellRendererToggle_ConnectToggled), f)
}

// NewCellRendererToggle creates a new CellRendererToggle. Adjust rendering
// parameters using object properties. Object properties can be set globally
// (with g_object_set()). Also, with TreeViewColumn, you can bind a property to
// a value in a TreeModel. For example, you can bind the “active” property on
// the cell renderer to a boolean value in the model, thus causing the check
// button to reflect the state of the model.
//
// The function returns the following values:
//
//    - cellRendererToggle: new cell renderer.
//
func NewCellRendererToggle() *CellRendererToggle {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_toggle_new()

	var _cellRendererToggle *CellRendererToggle // out

	_cellRendererToggle = wrapCellRendererToggle(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererToggle
}

// Activatable returns whether the cell renderer is activatable. See
// gtk_cell_renderer_toggle_set_activatable().
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer is activatable.
//
func (toggle *CellRendererToggle) Activatable() bool {
	var _arg0 *C.GtkCellRendererToggle // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(coreglib.InternObject(toggle).Native()))

	_cret = C.gtk_cell_renderer_toggle_get_activatable(_arg0)
	runtime.KeepAlive(toggle)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Active returns whether the cell renderer is active. See
// gtk_cell_renderer_toggle_set_active().
//
// The function returns the following values:
//
//    - ok: TRUE if the cell renderer is active.
//
func (toggle *CellRendererToggle) Active() bool {
	var _arg0 *C.GtkCellRendererToggle // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(coreglib.InternObject(toggle).Native()))

	_cret = C.gtk_cell_renderer_toggle_get_active(_arg0)
	runtime.KeepAlive(toggle)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Radio returns whether we’re rendering radio toggles rather than checkboxes.
//
// The function returns the following values:
//
//    - ok: TRUE if we’re rendering radio toggles rather than checkboxes.
//
func (toggle *CellRendererToggle) Radio() bool {
	var _arg0 *C.GtkCellRendererToggle // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(coreglib.InternObject(toggle).Native()))

	_cret = C.gtk_cell_renderer_toggle_get_radio(_arg0)
	runtime.KeepAlive(toggle)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActivatable makes the cell renderer activatable.
//
// The function takes the following parameters:
//
//    - setting: value to set.
//
func (toggle *CellRendererToggle) SetActivatable(setting bool) {
	var _arg0 *C.GtkCellRendererToggle // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(coreglib.InternObject(toggle).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_toggle_set_activatable(_arg0, _arg1)
	runtime.KeepAlive(toggle)
	runtime.KeepAlive(setting)
}

// SetActive activates or deactivates a cell renderer.
//
// The function takes the following parameters:
//
//    - setting: value to set.
//
func (toggle *CellRendererToggle) SetActive(setting bool) {
	var _arg0 *C.GtkCellRendererToggle // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(coreglib.InternObject(toggle).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_toggle_set_active(_arg0, _arg1)
	runtime.KeepAlive(toggle)
	runtime.KeepAlive(setting)
}

// SetRadio: if radio is TRUE, the cell renderer renders a radio toggle (i.e. a
// toggle in a group of mutually-exclusive toggles). If FALSE, it renders a
// check toggle (a standalone boolean option). This can be set globally for the
// cell renderer, or changed just before rendering each cell in the model (for
// TreeView, you set up a per-row setting using TreeViewColumn to associate
// model columns with cell renderer properties).
//
// The function takes the following parameters:
//
//    - radio: TRUE to make the toggle look like a radio button.
//
func (toggle *CellRendererToggle) SetRadio(radio bool) {
	var _arg0 *C.GtkCellRendererToggle // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkCellRendererToggle)(unsafe.Pointer(coreglib.InternObject(toggle).Native()))
	if radio {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_toggle_set_radio(_arg0, _arg1)
	runtime.KeepAlive(toggle)
	runtime.KeepAlive(radio)
}
