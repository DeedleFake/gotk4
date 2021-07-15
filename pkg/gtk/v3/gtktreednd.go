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
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_drag_dest_get_type()), F: marshalTreeDragDester},
		{T: externglib.Type(C.gtk_tree_drag_source_get_type()), F: marshalTreeDragSourcer},
	})
}

// TreeGetRowDragData obtains a tree_model and path from selection data of
// target type GTK_TREE_MODEL_ROW. Normally called from a drag_data_received
// handler. This function can only be used if selection_data originates from the
// same process that’s calling this function, because a pointer to the tree
// model is being passed around. If you aren’t in the same process, then you'll
// get memory corruption. In the TreeDragDest drag_data_received handler, you
// can assume that selection data of type GTK_TREE_MODEL_ROW is in from the
// current process. The returned path must be freed with gtk_tree_path_free().
func TreeGetRowDragData(selectionData *SelectionData) (*TreeModel, *TreePath, bool) {
	var _arg1 *C.GtkSelectionData // out
	var _arg2 *C.GtkTreeModel     // in
	var _arg3 *C.GtkTreePath      // in
	var _cret C.gboolean          // in

	_arg1 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))

	_cret = C.gtk_tree_get_row_drag_data(_arg1, &_arg2, &_arg3)

	var _treeModel *TreeModel // out
	var _path *TreePath       // out
	var _ok bool              // out

	_treeModel = wrapTreeModel(externglib.Take(unsafe.Pointer(_arg2)))
	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(_arg3)))
	runtime.SetFinalizer(_path, func(v *TreePath) {
		C.gtk_tree_path_free((*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(v))))
	})
	if _cret != 0 {
		_ok = true
	}

	return _treeModel, _path, _ok
}

// TreeSetRowDragData sets selection data of target type GTK_TREE_MODEL_ROW.
// Normally used in a drag_data_get handler.
func TreeSetRowDragData(selectionData *SelectionData, treeModel TreeModeller, path *TreePath) bool {
	var _arg1 *C.GtkSelectionData // out
	var _arg2 *C.GtkTreeModel     // out
	var _arg3 *C.GtkTreePath      // out
	var _cret C.gboolean          // in

	_arg1 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))
	_arg2 = (*C.GtkTreeModel)(unsafe.Pointer((treeModel).(gextras.Nativer).Native()))
	_arg3 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_set_row_drag_data(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragDestOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TreeDragDestOverrider interface {
	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from selection_data. If dest is
	// outside the tree so that inserting before it is impossible, FALSE will be
	// returned. Also, FALSE may be returned if the new row is not created for
	// some model-specific reason. Should robustly handle a dest no longer found
	// in the model!
	DragDataReceived(dest *TreePath, selectionData *SelectionData) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path. i.e., can we drop the data in
	// selection_data at that location. dest_path does not have to exist; the
	// return value will almost certainly be FALSE if the parent of dest_path
	// doesn’t exist, though.
	RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool
}

type TreeDragDest struct {
	*externglib.Object
}

var _ gextras.Nativer = (*TreeDragDest)(nil)

// TreeDragDester describes TreeDragDest's abstract methods.
type TreeDragDester interface {
	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from selection_data.
	DragDataReceived(dest *TreePath, selectionData *SelectionData) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path.
	RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool
}

var _ TreeDragDester = (*TreeDragDest)(nil)

func wrapTreeDragDest(obj *externglib.Object) *TreeDragDest {
	return &TreeDragDest{
		Object: obj,
	}
}

func marshalTreeDragDester(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeDragDest(obj), nil
}

// DragDataReceived asks the TreeDragDest to insert a row before the path dest,
// deriving the contents of the row from selection_data. If dest is outside the
// tree so that inserting before it is impossible, FALSE will be returned. Also,
// FALSE may be returned if the new row is not created for some model-specific
// reason. Should robustly handle a dest no longer found in the model!
func (dragDest *TreeDragDest) DragDataReceived(dest *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragDest  // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkSelectionData // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(dragDest.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(dest)))
	_arg2 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))

	_cret = C.gtk_tree_drag_dest_drag_data_received(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowDropPossible determines whether a drop is possible before the given
// dest_path, at the same depth as dest_path. i.e., can we drop the data in
// selection_data at that location. dest_path does not have to exist; the return
// value will almost certainly be FALSE if the parent of dest_path doesn’t
// exist, though.
func (dragDest *TreeDragDest) RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragDest  // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkSelectionData // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(dragDest.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(destPath)))
	_arg2 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))

	_cret = C.gtk_tree_drag_dest_row_drop_possible(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragSourceOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TreeDragSourceOverrider interface {
	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop. Returns FALSE if the
	// deletion fails because path no longer exists, or for some model-specific
	// reason. Should robustly handle a path no longer found in the model!
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to fill in selection_data with a
	// representation of the row at path. selection_data->target gives the
	// required type of the data. Should robustly handle a path no longer found
	// in the model!
	DragDataGet(path *TreePath, selectionData *SelectionData) bool
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation. If the source doesn’t implement this
	// interface, the row is assumed draggable.
	RowDraggable(path *TreePath) bool
}

type TreeDragSource struct {
	*externglib.Object
}

var _ gextras.Nativer = (*TreeDragSource)(nil)

// TreeDragSourcer describes TreeDragSource's abstract methods.
type TreeDragSourcer interface {
	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop.
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to fill in selection_data with a
	// representation of the row at path.
	DragDataGet(path *TreePath, selectionData *SelectionData) bool
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation.
	RowDraggable(path *TreePath) bool
}

var _ TreeDragSourcer = (*TreeDragSource)(nil)

func wrapTreeDragSource(obj *externglib.Object) *TreeDragSource {
	return &TreeDragSource{
		Object: obj,
	}
}

func marshalTreeDragSourcer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeDragSource(obj), nil
}

// DragDataDelete asks the TreeDragSource to delete the row at path, because it
// was moved somewhere else via drag-and-drop. Returns FALSE if the deletion
// fails because path no longer exists, or for some model-specific reason.
// Should robustly handle a path no longer found in the model!
func (dragSource *TreeDragSource) DragDataDelete(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_drag_data_delete(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragDataGet asks the TreeDragSource to fill in selection_data with a
// representation of the row at path. selection_data->target gives the required
// type of the data. Should robustly handle a path no longer found in the model!
func (dragSource *TreeDragSource) DragDataGet(path *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _arg2 *C.GtkSelectionData  // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))
	_arg2 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))

	_cret = C.gtk_tree_drag_source_drag_data_get(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowDraggable asks the TreeDragSource whether a particular row can be used as
// the source of a DND operation. If the source doesn’t implement this
// interface, the row is assumed draggable.
func (dragSource *TreeDragSource) RowDraggable(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_row_draggable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
