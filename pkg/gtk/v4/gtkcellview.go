// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeCellView = coreglib.Type(C.gtk_cell_view_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellView, F: marshalCellView},
	})
}

// CellView: widget displaying a single row of a GtkTreeModel
//
// A CellView displays a single row of a TreeModel using a CellArea and
// CellAreaContext. A CellAreaContext can be provided to the CellView at
// construction time in order to keep the cellview in context of a group of cell
// views, this ensures that the renderers displayed will be properly aligned
// with each other (like the aligned cells in the menus of ComboBox).
//
// CellView is Orientable in order to decide in which orientation the underlying
// CellAreaContext should be allocated. Taking the ComboBox menu as an example,
// cellviews should be oriented horizontally if the menus are listed
// top-to-bottom and thus all share the same width but may have separate
// individual heights (left-to-right menus should be allocated vertically since
// they all share the same height but may have variable widths).
//
//
// CSS nodes
//
// GtkCellView has a single CSS node with name cellview.
type CellView struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	CellLayout
	Orientable
}

var (
	_ Widgetter         = (*CellView)(nil)
	_ coreglib.Objector = (*CellView)(nil)
)

func wrapCellView(obj *coreglib.Object) *CellView {
	return &CellView{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		CellLayout: CellLayout{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCellView(p uintptr) (interface{}, error) {
	return wrapCellView(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCellView creates a new CellView widget.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellView() *CellView {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_cell_view_new()

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithContext creates a new CellView widget with a specific CellArea
// to layout cells and a specific CellAreaContext.
//
// Specifying the same context for a handful of cells lets the underlying area
// synchronize the geometry for those cells, in this way alignments with
// cellviews for other rows are possible.
//
// The function takes the following parameters:
//
//    - area to layout cells.
//    - context in which to calculate cell geometry.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellViewWithContext(area CellAreaer, context *CellAreaContext) *CellView {
	var _arg1 *C.GtkCellArea        // out
	var _arg2 *C.GtkCellAreaContext // out
	var _cret *C.GtkWidget          // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(coreglib.InternObject(area).Native()))
	_arg2 = (*C.GtkCellAreaContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gtk_cell_view_new_with_context(_arg1, _arg2)
	runtime.KeepAlive(area)
	runtime.KeepAlive(context)

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithMarkup creates a new CellView widget, adds a CellRendererText
// to it, and makes it show markup. The text can be marked up with the [Pango
// text markup language][PangoMarkupFormat].
//
// The function takes the following parameters:
//
//    - markup: text to display in the cell view.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellViewWithMarkup(markup string) *CellView {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_cell_view_new_with_markup(_arg1)
	runtime.KeepAlive(markup)

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithText creates a new CellView widget, adds a CellRendererText to
// it, and makes it show text.
//
// The function takes the following parameters:
//
//    - text to display in the cell view.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellViewWithText(text string) *CellView {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_cell_view_new_with_text(_arg1)
	runtime.KeepAlive(text)

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithTexture creates a new CellView widget, adds a
// CellRendererPixbuf to it, and makes it show texture.
//
// The function takes the following parameters:
//
//    - texture: image to display in the cell view.
//
// The function returns the following values:
//
//    - cellView: newly created CellView widget.
//
func NewCellViewWithTexture(texture gdk.Texturer) *CellView {
	var _arg1 *C.GdkTexture // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GdkTexture)(unsafe.Pointer(coreglib.InternObject(texture).Native()))

	_cret = C.gtk_cell_view_new_with_texture(_arg1)
	runtime.KeepAlive(texture)

	var _cellView *CellView // out

	_cellView = wrapCellView(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// DisplayedRow returns a TreePath referring to the currently displayed row. If
// no row is currently displayed, NULL is returned.
//
// The function returns the following values:
//
//    - treePath (optional): currently displayed row or NULL.
//
func (cellView *CellView) DisplayedRow() *TreePath {
	var _arg0 *C.GtkCellView // out
	var _cret *C.GtkTreePath // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_cret = C.gtk_cell_view_get_displayed_row(_arg0)
	runtime.KeepAlive(cellView)

	var _treePath *TreePath // out

	if _cret != nil {
		_treePath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_treePath)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_tree_path_free((*C.GtkTreePath)(intern.C))
			},
		)
	}

	return _treePath
}

// DrawSensitive gets whether cell_view is configured to draw all of its cells
// in a sensitive state.
//
// The function returns the following values:
//
//    - ok: whether cell_view draws all of its cells in a sensitive state.
//
func (cellView *CellView) DrawSensitive() bool {
	var _arg0 *C.GtkCellView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_cret = C.gtk_cell_view_get_draw_sensitive(_arg0)
	runtime.KeepAlive(cellView)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FitModel gets whether cell_view is configured to request space to fit the
// entire TreeModel.
//
// The function returns the following values:
//
//    - ok: whether cell_view requests space to fit the entire TreeModel.
//
func (cellView *CellView) FitModel() bool {
	var _arg0 *C.GtkCellView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_cret = C.gtk_cell_view_get_fit_model(_arg0)
	runtime.KeepAlive(cellView)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Model returns the model for cell_view. If no model is used NULL is returned.
//
// The function returns the following values:
//
//    - treeModel (optional) used or NULL.
//
func (cellView *CellView) Model() *TreeModel {
	var _arg0 *C.GtkCellView  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))

	_cret = C.gtk_cell_view_get_model(_arg0)
	runtime.KeepAlive(cellView)

	var _treeModel *TreeModel // out

	if _cret != nil {
		_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _treeModel
}

// SetDisplayedRow sets the row of the model that is currently displayed by the
// CellView. If the path is unset, then the contents of the cellview “stick” at
// their last value; this is not normally a desired result, but may be a needed
// intermediate state if say, the model for the CellView becomes temporarily
// empty.
//
// The function takes the following parameters:
//
//    - path (optional) or NULL to unset.
//
func (cellView *CellView) SetDisplayedRow(path *TreePath) {
	var _arg0 *C.GtkCellView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if path != nil {
		_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))
	}

	C.gtk_cell_view_set_displayed_row(_arg0, _arg1)
	runtime.KeepAlive(cellView)
	runtime.KeepAlive(path)
}

// SetDrawSensitive sets whether cell_view should draw all of its cells in a
// sensitive state, this is used by ComboBox menus to ensure that rows with
// insensitive cells that contain children appear sensitive in the parent menu
// item.
//
// The function takes the following parameters:
//
//    - drawSensitive: whether to draw all cells in a sensitive state.
//
func (cellView *CellView) SetDrawSensitive(drawSensitive bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if drawSensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_draw_sensitive(_arg0, _arg1)
	runtime.KeepAlive(cellView)
	runtime.KeepAlive(drawSensitive)
}

// SetFitModel sets whether cell_view should request space to fit the entire
// TreeModel.
//
// This is used by ComboBox to ensure that the cell view displayed on the combo
// box’s button always gets enough space and does not resize when selection
// changes.
//
// The function takes the following parameters:
//
//    - fitModel: whether cell_view should request space for the whole model.
//
func (cellView *CellView) SetFitModel(fitModel bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if fitModel {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_fit_model(_arg0, _arg1)
	runtime.KeepAlive(cellView)
	runtime.KeepAlive(fitModel)
}

// SetModel sets the model for cell_view. If cell_view already has a model set,
// it will remove it before setting the new model. If model is NULL, then it
// will unset the old model.
//
// The function takes the following parameters:
//
//    - model (optional): TreeModel.
//
func (cellView *CellView) SetModel(model TreeModeller) {
	var _arg0 *C.GtkCellView  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(coreglib.InternObject(cellView).Native()))
	if model != nil {
		_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.gtk_cell_view_set_model(_arg0, _arg1)
	runtime.KeepAlive(cellView)
	runtime.KeepAlive(model)
}
