// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_tree_drag_dest_get_type()), F: marshalTreeDragDest},
		{T: externglib.Type(C.gtk_tree_drag_source_get_type()), F: marshalTreeDragSource},
	})
}

// TreeGetRowDragData obtains a @tree_model and @path from selection data of
// target type GTK_TREE_MODEL_ROW. Normally called from a drag_data_received
// handler. This function can only be used if @selection_data originates from
// the same process that’s calling this function, because a pointer to the tree
// model is being passed around. If you aren’t in the same process, then you'll
// get memory corruption. In the TreeDragDest drag_data_received handler, you
// can assume that selection data of type GTK_TREE_MODEL_ROW is in from the
// current process. The returned path must be freed with gtk_tree_path_free().
func TreeGetRowDragData(selectionData *SelectionData) (TreeModel, *TreePath, bool) {
	var _arg1 *C.GtkSelectionData // out
	var _arg2 *C.GtkTreeModel     // in
	var _arg3 *C.GtkTreePath      // in
	var _cret C.gboolean          // in

	_arg1 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))

	_cret = C.gtk_tree_get_row_drag_data(_arg1, &_arg2, &_arg3)

	var _treeModel TreeModel // out
	var _path *TreePath      // out
	var _ok bool             // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg2))).(TreeModel)
	_path = (*TreePath)(unsafe.Pointer(_arg3))
	runtime.SetFinalizer(&_path, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})
	if _cret != 0 {
		_ok = true
	}

	return _treeModel, _path, _ok
}

// TreeSetRowDragData sets selection data of target type GTK_TREE_MODEL_ROW.
// Normally used in a drag_data_get handler.
func TreeSetRowDragData(selectionData *SelectionData, treeModel TreeModel, path *TreePath) bool {
	var _arg1 *C.GtkSelectionData // out
	var _arg2 *C.GtkTreeModel     // out
	var _arg3 *C.GtkTreePath      // out
	var _cret C.gboolean          // in

	_arg1 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))
	_arg2 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	_arg3 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_tree_set_row_drag_data(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

type TreeDragDest interface {
	gextras.Objector

	// DragDataReceived determines whether a drop is possible before the given
	// @dest_path, at the same depth as @dest_path. i.e., can we drop the data
	// in @selection_data at that location. @dest_path does not have to exist;
	// the return value will almost certainly be false if the parent of
	// @dest_path doesn’t exist, though.
	DragDataReceived(dest *TreePath, selectionData *SelectionData) bool
	// RowDropPossible determines whether a drop is possible before the given
	// @dest_path, at the same depth as @dest_path. i.e., can we drop the data
	// in @selection_data at that location. @dest_path does not have to exist;
	// the return value will almost certainly be false if the parent of
	// @dest_path doesn’t exist, though.
	RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool
}

// treeDragDest implements the TreeDragDest interface.
type treeDragDest struct {
	gextras.Objector
}

var _ TreeDragDest = (*treeDragDest)(nil)

// WrapTreeDragDest wraps a GObject to a type that implements
// interface TreeDragDest. It is primarily used internally.
func WrapTreeDragDest(obj *externglib.Object) TreeDragDest {
	return treeDragDest{
		Objector: obj,
	}
}

func marshalTreeDragDest(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeDragDest(obj), nil
}

func (d treeDragDest) DragDataReceived(dest *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragDest  // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkSelectionData // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(dest.Native()))
	_arg2 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))

	_cret = C.gtk_tree_drag_dest_drag_data_received(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d treeDragDest) RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragDest  // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkSelectionData // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(destPath.Native()))
	_arg2 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))

	_cret = C.gtk_tree_drag_dest_row_drop_possible(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

type TreeDragSource interface {
	gextras.Objector

	// DragDataDelete asks the TreeDragSource whether a particular row can be
	// used as the source of a DND operation. If the source doesn’t implement
	// this interface, the row is assumed draggable.
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation. If the source doesn’t implement this
	// interface, the row is assumed draggable.
	DragDataGet(path *TreePath, selectionData *SelectionData) bool
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation. If the source doesn’t implement this
	// interface, the row is assumed draggable.
	RowDraggable(path *TreePath) bool
}

// treeDragSource implements the TreeDragSource interface.
type treeDragSource struct {
	gextras.Objector
}

var _ TreeDragSource = (*treeDragSource)(nil)

// WrapTreeDragSource wraps a GObject to a type that implements
// interface TreeDragSource. It is primarily used internally.
func WrapTreeDragSource(obj *externglib.Object) TreeDragSource {
	return treeDragSource{
		Objector: obj,
	}
}

func marshalTreeDragSource(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeDragSource(obj), nil
}

func (d treeDragSource) DragDataDelete(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_tree_drag_source_drag_data_delete(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d treeDragSource) DragDataGet(path *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _arg2 *C.GtkSelectionData  // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))
	_arg2 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))

	_cret = C.gtk_tree_drag_source_drag_data_get(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d treeDragSource) RowDraggable(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_tree_drag_source_row_draggable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
