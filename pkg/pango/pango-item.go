// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_item_get_type()), F: marshalItem},
	})
}

// Analysis: the `PangoAnalysis` structure stores information about the
// properties of a segment of text.
type Analysis struct {
	native C.PangoAnalysis
}

// WrapAnalysis wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAnalysis(ptr unsafe.Pointer) *Analysis {
	return (*Analysis)(ptr)
}

// Native returns the underlying C source pointer.
func (a *Analysis) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Item: the `PangoItem` structure stores information about a segment of text.
//
// You typically obtain `PangoItems` by itemizing a piece of text with
// [func@itemize].
type Item struct {
	native C.PangoItem
}

// WrapItem wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapItem(ptr unsafe.Pointer) *Item {
	return (*Item)(ptr)
}

func marshalItem(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Item)(unsafe.Pointer(b)), nil
}

// NewItem constructs a struct Item.
func NewItem() *Item {
	var _cret *C.PangoItem // in

	_cret = C.pango_item_new()

	var _item *Item // out

	_item = (*Item)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_item, func(v *Item) {
		C.free(unsafe.Pointer(v))
	})

	return _item
}

// Native returns the underlying C source pointer.
func (i *Item) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// ApplyAttrs: add attributes to a `PangoItem`.
//
// The idea is that you have attributes that don't affect itemization, such as
// font features, so you filter them out using [method@Pango.AttrList.filter],
// itemize your text, then reapply the attributes to the resulting items using
// this function.
//
// The @iter should be positioned before the range of the item, and will be
// advanced past it. This function is meant to be called in a loop over the
// items resulting from itemization, while passing the iter to each call.
func (item *Item) ApplyAttrs(iter *AttrIterator) {
	var _arg0 *C.PangoItem         // out
	var _arg1 *C.PangoAttrIterator // out

	_arg0 = (*C.PangoItem)(unsafe.Pointer(item))
	_arg1 = (*C.PangoAttrIterator)(unsafe.Pointer(iter))

	C.pango_item_apply_attrs(_arg0, _arg1)
}

// Copy an existing `PangoItem` structure.
func (item *Item) Copy() *Item {
	var _arg0 *C.PangoItem // out
	var _cret *C.PangoItem // in

	_arg0 = (*C.PangoItem)(unsafe.Pointer(item))

	_cret = C.pango_item_copy(_arg0)

	var _ret *Item // out

	_ret = (*Item)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *Item) {
		C.free(unsafe.Pointer(v))
	})

	return _ret
}

// Free: free a `PangoItem` and all associated memory.
func (item *Item) free() {
	var _arg0 *C.PangoItem // out

	_arg0 = (*C.PangoItem)(unsafe.Pointer(item))

	C.pango_item_free(_arg0)
}

// Split modifies @orig to cover only the text after @split_index, and returns a
// new item that covers the text before @split_index that used to be in @orig.
//
// You can think of @split_index as the length of the returned item.
// @split_index may not be 0, and it may not be greater than or equal to the
// length of @orig (that is, there must be at least one byte assigned to each
// item, you can't create a zero-length item). @split_offset is the length of
// the first item in chars, and must be provided because the text used to
// generate the item isn't available, so `pango_item_split()` can't count the
// char length of the split items itself.
func (orig *Item) Split(splitIndex int, splitOffset int) *Item {
	var _arg0 *C.PangoItem // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out
	var _cret *C.PangoItem // in

	_arg0 = (*C.PangoItem)(unsafe.Pointer(orig))
	_arg1 = C.int(splitIndex)
	_arg2 = C.int(splitOffset)

	_cret = C.pango_item_split(_arg0, _arg1, _arg2)

	var _item *Item // out

	_item = (*Item)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_item, func(v *Item) {
		C.free(unsafe.Pointer(v))
	})

	return _item
}
