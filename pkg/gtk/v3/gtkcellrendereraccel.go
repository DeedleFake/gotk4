// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_accel_mode_get_type()), F: marshalCellRendererAccelMode},
		{T: externglib.Type(C.gtk_cell_renderer_accel_get_type()), F: marshalCellRendererAccel},
	})
}

// CellRendererAccelMode determines if the edited accelerators are GTK+
// accelerators. If they are, consumed modifiers are suppressed, only
// accelerators accepted by GTK+ are allowed, and the accelerators are rendered
// in the same way as they are in menus.
type CellRendererAccelMode int

const (
	// GTK: GTK+ accelerators mode
	CellRendererAccelModeGTK CellRendererAccelMode = iota
	// Other: other accelerator mode
	CellRendererAccelModeOther
)

func marshalCellRendererAccelMode(p uintptr) (interface{}, error) {
	return CellRendererAccelMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CellRendererAccelOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellRendererAccelOverrider interface {
	AccelCleared(pathString string)
}

// CellRendererAccel displays a keyboard accelerator (i.e. a key combination
// like `Control + a`). If the cell renderer is editable, the accelerator can be
// changed by simply typing the new combination.
//
// The CellRendererAccel cell renderer was added in GTK+ 2.10.
type CellRendererAccel interface {
	gextras.Objector

	privateCellRendererAccelClass()
}

// CellRendererAccelClass implements the CellRendererAccel interface.
type CellRendererAccelClass struct {
	CellRendererTextClass
}

var _ CellRendererAccel = (*CellRendererAccelClass)(nil)

func wrapCellRendererAccel(obj *externglib.Object) CellRendererAccel {
	return &CellRendererAccelClass{
		CellRendererTextClass: CellRendererTextClass{
			CellRendererClass: CellRendererClass{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCellRendererAccel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRendererAccel(obj), nil
}

// NewCellRendererAccel creates a new CellRendererAccel.
func NewCellRendererAccel() *CellRendererAccelClass {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_accel_new()

	var _cellRendererAccel *CellRendererAccelClass // out

	_cellRendererAccel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*CellRendererAccelClass)

	return _cellRendererAccel
}

func (*CellRendererAccelClass) privateCellRendererAccelClass() {}
