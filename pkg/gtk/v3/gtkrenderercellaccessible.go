// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_renderer_cell_accessible_get_type()), F: marshalRendererCellAccessibler},
	})
}

type RendererCellAccessible struct {
	CellAccessible
}

func wrapRendererCellAccessible(obj *externglib.Object) *RendererCellAccessible {
	return &RendererCellAccessible{
		CellAccessible: CellAccessible{
			Accessible: Accessible{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
			Action: atk.Action{
				Object: obj,
			},
			Component: atk.Component{
				Object: obj,
			},
			TableCell: atk.TableCell{
				ObjectClass: atk.ObjectClass{
					Object: obj,
				},
			},
			Object: obj,
		},
	}
}

func marshalRendererCellAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRendererCellAccessible(obj), nil
}

func NewRendererCellAccessible(renderer CellRendererer) *RendererCellAccessible {
	var _arg1 *C.GtkCellRenderer // out
	var _cret *C.AtkObject       // in

	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	_cret = C.gtk_renderer_cell_accessible_new(_arg1)
	runtime.KeepAlive(renderer)

	var _rendererCellAccessible *RendererCellAccessible // out

	_rendererCellAccessible = wrapRendererCellAccessible(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _rendererCellAccessible
}

func (*RendererCellAccessible) privateRendererCellAccessible() {}
