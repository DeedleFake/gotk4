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
		{T: externglib.Type(C.gtk_tree_store_get_type()), F: marshalTreeStore},
	})
}

// TreeStore: tree-like data structure that can be used with the GtkTreeView
//
// The TreeStore object is a list model for use with a TreeView widget. It
// implements the TreeModel interface, and consequently, can use all of the
// methods available there. It also implements the TreeSortable interface so it
// can be sorted by the view. Finally, it also implements the tree [drag and
// drop][gtk3-GtkTreeView-drag-and-drop] interfaces.
//
//
// GtkTreeStore as GtkBuildable
//
// The GtkTreeStore implementation of the Buildable interface allows to specify
// the model columns with a <columns> element that may contain multiple <column>
// elements, each specifying one model column. The “type” attribute specifies
// the data type for the column.
//
// An example of a UI Definition fragment for a tree store:
//
//    <object class="GtkTreeStore">
//      <columns>
//        <column type="gchararray"/>
//        <column type="gchararray"/>
//        <column type="gint"/>
//      </columns>
//    </object>
type TreeStore interface {
	gextras.Objector

	// Append appends a new row to @tree_store. If @parent is non-nil, then it
	// will append the new row after the last child of @parent, otherwise it
	// will append a row to the top level. @iter will be changed to point to
	// this new row. The row will be empty after this function is called. To
	// fill in values, you need to call gtk_tree_store_set() or
	// gtk_tree_store_set_value().
	Append(parent *TreeIter) TreeIter
	// Clear removes all rows from @tree_store
	Clear()
	// Insert creates a new row at @position. If parent is non-nil, then the row
	// will be made a child of @parent. Otherwise, the row will be created at
	// the toplevel. If @position is -1 or is larger than the number of rows at
	// that level, then the new row will be inserted to the end of the list.
	// @iter will be changed to point to this new row. The row will be empty
	// after this function is called. To fill in values, you need to call
	// gtk_tree_store_set() or gtk_tree_store_set_value().
	Insert(parent *TreeIter, position int) TreeIter
	// InsertAfter inserts a new row after @sibling. If @sibling is nil, then
	// the row will be prepended to @parent ’s children. If @parent and @sibling
	// are nil, then the row will be prepended to the toplevel. If both @sibling
	// and @parent are set, then @parent must be the parent of @sibling. When
	// @sibling is set, @parent is optional.
	//
	// @iter will be changed to point to this new row. The row will be empty
	// after this function is called. To fill in values, you need to call
	// gtk_tree_store_set() or gtk_tree_store_set_value().
	InsertAfter(parent *TreeIter, sibling *TreeIter) TreeIter
	// InsertBefore inserts a new row before @sibling. If @sibling is nil, then
	// the row will be appended to @parent ’s children. If @parent and @sibling
	// are nil, then the row will be appended to the toplevel. If both @sibling
	// and @parent are set, then @parent must be the parent of @sibling. When
	// @sibling is set, @parent is optional.
	//
	// @iter will be changed to point to this new row. The row will be empty
	// after this function is called. To fill in values, you need to call
	// gtk_tree_store_set() or gtk_tree_store_set_value().
	InsertBefore(parent *TreeIter, sibling *TreeIter) TreeIter
	// InsertWithValuesv: variant of gtk_tree_store_insert_with_values() which
	// takes the columns and values as two arrays, instead of varargs. This
	// function is mainly intended for language bindings.
	InsertWithValuesv(parent *TreeIter, position int, columns []int, values []externglib.Value) TreeIter
	// IsAncestor returns true if @iter is an ancestor of @descendant. That is,
	// @iter is the parent (or grandparent or great-grandparent) of @descendant.
	IsAncestor(iter *TreeIter, descendant *TreeIter) bool
	// IterDepth returns the depth of @iter. This will be 0 for anything on the
	// root level, 1 for anything down a level, etc.
	IterDepth(iter *TreeIter) int
	// IterIsValid: WARNING: This function is slow. Only use it for debugging
	// and/or testing purposes.
	//
	// Checks if the given iter is a valid iter for this TreeStore.
	IterIsValid(iter *TreeIter) bool
	// MoveAfter moves @iter in @tree_store to the position after @position.
	// @iter and @position should be in the same level. Note that this function
	// only works with unsorted stores. If @position is nil, @iter will be moved
	// to the start of the level.
	MoveAfter(iter *TreeIter, position *TreeIter)
	// MoveBefore moves @iter in @tree_store to the position before @position.
	// @iter and @position should be in the same level. Note that this function
	// only works with unsorted stores. If @position is nil, @iter will be moved
	// to the end of the level.
	MoveBefore(iter *TreeIter, position *TreeIter)
	// Prepend prepends a new row to @tree_store. If @parent is non-nil, then it
	// will prepend the new row before the first child of @parent, otherwise it
	// will prepend a row to the top level. @iter will be changed to point to
	// this new row. The row will be empty after this function is called. To
	// fill in values, you need to call gtk_tree_store_set() or
	// gtk_tree_store_set_value().
	Prepend(parent *TreeIter) TreeIter
	// Remove removes @iter from @tree_store. After being removed, @iter is set
	// to the next valid row at that level, or invalidated if it previously
	// pointed to the last one.
	Remove(iter *TreeIter) bool
	// SetColumnTypes: this function is meant primarily for #GObjects that
	// inherit from TreeStore, and should only be used when constructing a new
	// TreeStore. It will not function after a row has been added, or a method
	// on the TreeModel interface is called.
	SetColumnTypes(types []externglib.Type)
	// SetValue sets the data in the cell specified by @iter and @column. The
	// type of @value must be convertible to the type of the column.
	SetValue(iter *TreeIter, column int, value *externglib.Value)
	// SetValuesv: variant of gtk_tree_store_set_valist() which takes the
	// columns and values as two arrays, instead of varargs. This function is
	// mainly intended for language bindings or in case the number of columns to
	// change is not known until run-time.
	SetValuesv(iter *TreeIter, columns []int, values []externglib.Value)
	// Swap swaps @a and @b in the same level of @tree_store. Note that this
	// function only works with unsorted stores.
	Swap(a *TreeIter, b *TreeIter)
}

// TreeStoreClass implements the TreeStore interface.
type TreeStoreClass struct {
	*externglib.Object
	*externglib.Object
	BuildableIface
	TreeDragDestIface
	TreeDragSourceIface
	TreeModelIface
	TreeSortableIface
}

var _ TreeStore = (*TreeStoreClass)(nil)

func wrapTreeStore(obj *externglib.Object) TreeStore {
	return &TreeStoreClass{
		Object: obj,
		Object: obj,
		BuildableIface: BuildableIface{
			Object: obj,
		},
		TreeDragDestIface: TreeDragDestIface{
			Object: obj,
		},
		TreeDragSourceIface: TreeDragSourceIface{
			Object: obj,
		},
		TreeModelIface: TreeModelIface{
			Object: obj,
		},
		TreeSortableIface: TreeSortableIface{
			TreeModelIface: TreeModelIface{
				Object: obj,
			},
		},
	}
}

func marshalTreeStore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTreeStore(obj), nil
}

// NewTreeStoreV: non vararg creation function. Used primarily by language
// bindings.
func NewTreeStoreV(types []externglib.Type) *TreeStoreClass {
	var _arg2 *C.GType
	var _arg1 C.int
	var _cret *C.GtkTreeStore // in

	_arg1 = C.int(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	_cret = C.gtk_tree_store_newv(_arg1, _arg2)

	var _treeStore *TreeStoreClass // out

	_treeStore = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*TreeStoreClass)

	return _treeStore
}

// Append appends a new row to @tree_store. If @parent is non-nil, then it will
// append the new row after the last child of @parent, otherwise it will append
// a row to the top level. @iter will be changed to point to this new row. The
// row will be empty after this function is called. To fill in values, you need
// to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (treeStore *TreeStoreClass) Append(parent *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))

	C.gtk_tree_store_append(_arg0, &_arg1, _arg2)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// Clear removes all rows from @tree_store
func (treeStore *TreeStoreClass) Clear() {
	var _arg0 *C.GtkTreeStore // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))

	C.gtk_tree_store_clear(_arg0)
}

// Insert creates a new row at @position. If parent is non-nil, then the row
// will be made a child of @parent. Otherwise, the row will be created at the
// toplevel. If @position is -1 or is larger than the number of rows at that
// level, then the new row will be inserted to the end of the list. @iter will
// be changed to point to this new row. The row will be empty after this
// function is called. To fill in values, you need to call gtk_tree_store_set()
// or gtk_tree_store_set_value().
func (treeStore *TreeStoreClass) Insert(parent *TreeIter, position int) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 C.int           // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))
	_arg3 = C.int(position)

	C.gtk_tree_store_insert(_arg0, &_arg1, _arg2, _arg3)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// InsertAfter inserts a new row after @sibling. If @sibling is nil, then the
// row will be prepended to @parent ’s children. If @parent and @sibling are
// nil, then the row will be prepended to the toplevel. If both @sibling and
// @parent are set, then @parent must be the parent of @sibling. When @sibling
// is set, @parent is optional.
//
// @iter will be changed to point to this new row. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
func (treeStore *TreeStoreClass) InsertAfter(parent *TreeIter, sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))
	_arg3 = (*C.GtkTreeIter)(unsafe.Pointer(sibling))

	C.gtk_tree_store_insert_after(_arg0, &_arg1, _arg2, _arg3)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// InsertBefore inserts a new row before @sibling. If @sibling is nil, then the
// row will be appended to @parent ’s children. If @parent and @sibling are nil,
// then the row will be appended to the toplevel. If both @sibling and @parent
// are set, then @parent must be the parent of @sibling. When @sibling is set,
// @parent is optional.
//
// @iter will be changed to point to this new row. The row will be empty after
// this function is called. To fill in values, you need to call
// gtk_tree_store_set() or gtk_tree_store_set_value().
func (treeStore *TreeStoreClass) InsertBefore(parent *TreeIter, sibling *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))
	_arg3 = (*C.GtkTreeIter)(unsafe.Pointer(sibling))

	C.gtk_tree_store_insert_before(_arg0, &_arg1, _arg2, _arg3)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// InsertWithValuesv: variant of gtk_tree_store_insert_with_values() which takes
// the columns and values as two arrays, instead of varargs. This function is
// mainly intended for language bindings.
func (treeStore *TreeStoreClass) InsertWithValuesv(parent *TreeIter, position int, columns []int, values []externglib.Value) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 C.int           // out
	var _arg4 *C.int
	var _arg6 C.int
	var _arg5 *C.GValue

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))
	_arg3 = C.int(position)
	_arg6 = C.int(len(columns))
	_arg4 = (*C.int)(unsafe.Pointer(&columns[0]))
	_arg6 = C.int(len(values))
	_arg5 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg5))
	{
		out := unsafe.Slice(_arg5, len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer(&(&values[i]).GValue))
		}
	}

	C.gtk_tree_store_insert_with_valuesv(_arg0, &_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// IsAncestor returns true if @iter is an ancestor of @descendant. That is,
// @iter is the parent (or grandparent or great-grandparent) of @descendant.
func (treeStore *TreeStoreClass) IsAncestor(iter *TreeIter, descendant *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(descendant))

	_cret = C.gtk_tree_store_is_ancestor(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IterDepth returns the depth of @iter. This will be 0 for anything on the root
// level, 1 for anything down a level, etc.
func (treeStore *TreeStoreClass) IterDepth(iter *TreeIter) int {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.int           // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_iter_depth(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IterIsValid: WARNING: This function is slow. Only use it for debugging and/or
// testing purposes.
//
// Checks if the given iter is a valid iter for this TreeStore.
func (treeStore *TreeStoreClass) IterIsValid(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_iter_is_valid(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MoveAfter moves @iter in @tree_store to the position after @position. @iter
// and @position should be in the same level. Note that this function only works
// with unsorted stores. If @position is nil, @iter will be moved to the start
// of the level.
func (treeStore *TreeStoreClass) MoveAfter(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_tree_store_move_after(_arg0, _arg1, _arg2)
}

// MoveBefore moves @iter in @tree_store to the position before @position. @iter
// and @position should be in the same level. Note that this function only works
// with unsorted stores. If @position is nil, @iter will be moved to the end of
// the level.
func (treeStore *TreeStoreClass) MoveBefore(iter *TreeIter, position *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_tree_store_move_before(_arg0, _arg1, _arg2)
}

// Prepend prepends a new row to @tree_store. If @parent is non-nil, then it
// will prepend the new row before the first child of @parent, otherwise it will
// prepend a row to the top level. @iter will be changed to point to this new
// row. The row will be empty after this function is called. To fill in values,
// you need to call gtk_tree_store_set() or gtk_tree_store_set_value().
func (treeStore *TreeStoreClass) Prepend(parent *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 C.GtkTreeIter   // in
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(parent))

	C.gtk_tree_store_prepend(_arg0, &_arg1, _arg2)

	var _iter TreeIter // out

	_iter = *(*TreeIter)(unsafe.Pointer((&_arg1)))

	return _iter
}

// Remove removes @iter from @tree_store. After being removed, @iter is set to
// the next valid row at that level, or invalidated if it previously pointed to
// the last one.
func (treeStore *TreeStoreClass) Remove(iter *TreeIter) bool {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))

	_cret = C.gtk_tree_store_remove(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetColumnTypes: this function is meant primarily for #GObjects that inherit
// from TreeStore, and should only be used when constructing a new TreeStore. It
// will not function after a row has been added, or a method on the TreeModel
// interface is called.
func (treeStore *TreeStoreClass) SetColumnTypes(types []externglib.Type) {
	var _arg0 *C.GtkTreeStore // out
	var _arg2 *C.GType
	var _arg1 C.int

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = C.int(len(types))
	_arg2 = (*C.GType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(types))
		for i := range types {
			out[i] = (C.GType)(types[i])
		}
	}

	C.gtk_tree_store_set_column_types(_arg0, _arg1, _arg2)
}

// SetValue sets the data in the cell specified by @iter and @column. The type
// of @value must be convertible to the type of the column.
func (treeStore *TreeStoreClass) SetValue(iter *TreeIter, column int, value *externglib.Value) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 C.int           // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg2 = C.int(column)
	_arg3 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gtk_tree_store_set_value(_arg0, _arg1, _arg2, _arg3)
}

// SetValuesv: variant of gtk_tree_store_set_valist() which takes the columns
// and values as two arrays, instead of varargs. This function is mainly
// intended for language bindings or in case the number of columns to change is
// not known until run-time.
func (treeStore *TreeStoreClass) SetValuesv(iter *TreeIter, columns []int, values []externglib.Value) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.int
	var _arg4 C.int
	var _arg3 *C.GValue

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter))
	_arg4 = C.int(len(columns))
	_arg2 = (*C.int)(unsafe.Pointer(&columns[0]))
	_arg4 = C.int(len(values))
	_arg3 = (*C.GValue)(C.malloc(C.ulong(len(values)) * C.ulong(C.sizeof_GValue)))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice(_arg3, len(values))
		for i := range values {
			out[i] = *(*C.GValue)(unsafe.Pointer(&(&values[i]).GValue))
		}
	}

	C.gtk_tree_store_set_valuesv(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// Swap swaps @a and @b in the same level of @tree_store. Note that this
// function only works with unsorted stores.
func (treeStore *TreeStoreClass) Swap(a *TreeIter, b *TreeIter) {
	var _arg0 *C.GtkTreeStore // out
	var _arg1 *C.GtkTreeIter  // out
	var _arg2 *C.GtkTreeIter  // out

	_arg0 = (*C.GtkTreeStore)(unsafe.Pointer(treeStore.Native()))
	_arg1 = (*C.GtkTreeIter)(unsafe.Pointer(a))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(b))

	C.gtk_tree_store_swap(_arg0, _arg1, _arg2)
}
