// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// Sequence: the #GSequence struct is an opaque data type representing a
// [sequence][glib-Sequences] data type.
type Sequence struct {
	native C.GSequence
}

// WrapSequence wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSequence(ptr unsafe.Pointer) *Sequence {
	if ptr == nil {
		return nil
	}

	return (*Sequence)(ptr)
}

func marshalSequence(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSequence(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *Sequence) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Append adds a new item to the end of @seq.
func (s *Sequence) Append(data interface{}) *SequenceIter {
	var _arg0 *C.GSequence
	var _arg1 C.gpointer

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))
	_arg1 = C.gpointer(data)

	var _cret *C.GSequenceIter

	cret = C.g_sequence_append(_arg0, _arg1)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// Foreach calls @func for each item in the sequence passing @user_data to the
// function. @func must not modify the sequence itself.
func (s *Sequence) Foreach() {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	C.g_sequence_foreach(_arg0)
}

// Free frees the memory allocated for @seq. If @seq has a data destroy function
// associated with it, that function is called on all items in @seq.
func (s *Sequence) Free() {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	C.g_sequence_free(_arg0)
}

// BeginIter returns the begin iterator for @seq.
func (s *Sequence) BeginIter() *SequenceIter {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_get_begin_iter(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// EndIter returns the end iterator for @seg
func (s *Sequence) EndIter() *SequenceIter {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_get_end_iter(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// IterAtPos returns the iterator at position @pos. If @pos is negative or
// larger than the number of items in @seq, the end iterator is returned.
func (s *Sequence) IterAtPos(pos int) *SequenceIter {
	var _arg0 *C.GSequence
	var _arg1 C.gint

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(pos)

	var _cret *C.GSequenceIter

	cret = C.g_sequence_get_iter_at_pos(_arg0, _arg1)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// Length returns the positive length (>= 0) of @seq. Note that this method is
// O(h) where `h' is the height of the tree. It is thus more efficient to use
// g_sequence_is_empty() when comparing the length to zero.
func (s *Sequence) Length() int {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret C.gint

	cret = C.g_sequence_get_length(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// InsertSorted inserts @data into @seq using @cmp_func to determine the new
// position. The sequence must already be sorted according to @cmp_func;
// otherwise the new position of @data is undefined.
//
// @cmp_func is called with two items of the @seq, and @cmp_data. It should
// return 0 if the items are equal, a negative value if the first item comes
// before the second, and a positive value if the second item comes before the
// first.
//
// Note that when adding a large amount of data to a #GSequence, it is more
// efficient to do unsorted insertions and then call g_sequence_sort() or
// g_sequence_sort_iter().
func (s *Sequence) InsertSorted() *SequenceIter {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_insert_sorted(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// InsertSortedIter: like g_sequence_insert_sorted(), but uses a IterCompareFunc
// instead of a DataFunc as the compare function.
//
// @iter_cmp is called with two iterators pointing into @seq. It should return 0
// if the iterators are equal, a negative value if the first iterator comes
// before the second, and a positive value if the second iterator comes before
// the first.
//
// Note that when adding a large amount of data to a #GSequence, it is more
// efficient to do unsorted insertions and then call g_sequence_sort() or
// g_sequence_sort_iter().
func (s *Sequence) InsertSortedIter() *SequenceIter {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_insert_sorted_iter(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// IsEmpty returns true if the sequence contains zero items.
//
// This function is functionally identical to checking the result of
// g_sequence_get_length() being equal to zero. However this function is
// implemented in O(1) running time.
func (s *Sequence) IsEmpty() bool {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	cret = C.g_sequence_is_empty(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Lookup returns an iterator pointing to the position of the first item found
// equal to @data according to @cmp_func and @cmp_data. If more than one item is
// equal, it is not guaranteed that it is the first which is returned. In that
// case, you can use g_sequence_iter_next() and g_sequence_iter_prev() to get
// others.
//
// @cmp_func is called with two items of the @seq, and @cmp_data. It should
// return 0 if the items are equal, a negative value if the first item comes
// before the second, and a positive value if the second item comes before the
// first.
//
// This function will fail if the data contained in the sequence is unsorted.
func (s *Sequence) Lookup() *SequenceIter {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_lookup(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// LookupIter: like g_sequence_lookup(), but uses a IterCompareFunc instead of a
// DataFunc as the compare function.
//
// @iter_cmp is called with two iterators pointing into @seq. It should return 0
// if the iterators are equal, a negative value if the first iterator comes
// before the second, and a positive value if the second iterator comes before
// the first.
//
// This function will fail if the data contained in the sequence is unsorted.
func (s *Sequence) LookupIter() *SequenceIter {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_lookup_iter(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// Prepend adds a new item to the front of @seq
func (s *Sequence) Prepend(data interface{}) *SequenceIter {
	var _arg0 *C.GSequence
	var _arg1 C.gpointer

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))
	_arg1 = C.gpointer(data)

	var _cret *C.GSequenceIter

	cret = C.g_sequence_prepend(_arg0, _arg1)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// Search returns an iterator pointing to the position where @data would be
// inserted according to @cmp_func and @cmp_data.
//
// @cmp_func is called with two items of the @seq, and @cmp_data. It should
// return 0 if the items are equal, a negative value if the first item comes
// before the second, and a positive value if the second item comes before the
// first.
//
// If you are simply searching for an existing element of the sequence, consider
// using g_sequence_lookup().
//
// This function will fail if the data contained in the sequence is unsorted.
func (s *Sequence) Search() *SequenceIter {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_search(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// SearchIter: like g_sequence_search(), but uses a IterCompareFunc instead of a
// DataFunc as the compare function.
//
// @iter_cmp is called with two iterators pointing into @seq. It should return 0
// if the iterators are equal, a negative value if the first iterator comes
// before the second, and a positive value if the second iterator comes before
// the first.
//
// If you are simply searching for an existing element of the sequence, consider
// using g_sequence_lookup_iter().
//
// This function will fail if the data contained in the sequence is unsorted.
func (s *Sequence) SearchIter() *SequenceIter {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_search_iter(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// Sort sorts @seq using @cmp_func.
//
// @cmp_func is passed two items of @seq and should return 0 if they are equal,
// a negative value if the first comes before the second, and a positive value
// if the second comes before the first.
func (s *Sequence) Sort() {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	C.g_sequence_sort(_arg0)
}

// SortIter: like g_sequence_sort(), but uses a IterCompareFunc instead of a
// DataFunc as the compare function
//
// @cmp_func is called with two iterators pointing into @seq. It should return 0
// if the iterators are equal, a negative value if the first iterator comes
// before the second, and a positive value if the second iterator comes before
// the first.
func (s *Sequence) SortIter() {
	var _arg0 *C.GSequence

	_arg0 = (*C.GSequence)(unsafe.Pointer(s.Native()))

	C.g_sequence_sort_iter(_arg0)
}

// SequenceIter: the Iter struct is an opaque data type representing an iterator
// pointing into a #GSequence.
type SequenceIter struct {
	native C.GSequenceIter
}

// WrapSequenceIter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSequenceIter(ptr unsafe.Pointer) *SequenceIter {
	if ptr == nil {
		return nil
	}

	return (*SequenceIter)(ptr)
}

func marshalSequenceIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSequenceIter(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *SequenceIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Compare returns a negative number if @a comes before @b, 0 if they are equal,
// and a positive number if @a comes after @b.
//
// The @a and @b iterators must point into the same sequence.
func (a *SequenceIter) Compare(b *SequenceIter) int {
	var _arg0 *C.GSequenceIter
	var _arg1 *C.GSequenceIter

	_arg0 = (*C.GSequenceIter)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GSequenceIter)(unsafe.Pointer(b.Native()))

	var _cret C.gint

	cret = C.g_sequence_iter_compare(_arg0, _arg1)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Position returns the position of @iter
func (i *SequenceIter) Position() int {
	var _arg0 *C.GSequenceIter

	_arg0 = (*C.GSequenceIter)(unsafe.Pointer(i.Native()))

	var _cret C.gint

	cret = C.g_sequence_iter_get_position(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Sequence returns the #GSequence that @iter points into.
func (i *SequenceIter) Sequence() *Sequence {
	var _arg0 *C.GSequenceIter

	_arg0 = (*C.GSequenceIter)(unsafe.Pointer(i.Native()))

	var _cret *C.GSequence

	cret = C.g_sequence_iter_get_sequence(_arg0)

	var _sequence *Sequence

	_sequence = WrapSequence(unsafe.Pointer(_cret))

	return _sequence
}

// IsBegin returns whether @iter is the begin iterator
func (i *SequenceIter) IsBegin() bool {
	var _arg0 *C.GSequenceIter

	_arg0 = (*C.GSequenceIter)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.g_sequence_iter_is_begin(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsEnd returns whether @iter is the end iterator
func (i *SequenceIter) IsEnd() bool {
	var _arg0 *C.GSequenceIter

	_arg0 = (*C.GSequenceIter)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	cret = C.g_sequence_iter_is_end(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Move returns the Iter which is @delta positions away from @iter. If @iter is
// closer than -@delta positions to the beginning of the sequence, the begin
// iterator is returned. If @iter is closer than @delta positions to the end of
// the sequence, the end iterator is returned.
func (i *SequenceIter) Move(delta int) *SequenceIter {
	var _arg0 *C.GSequenceIter
	var _arg1 C.gint

	_arg0 = (*C.GSequenceIter)(unsafe.Pointer(i.Native()))
	_arg1 = C.gint(delta)

	var _cret *C.GSequenceIter

	cret = C.g_sequence_iter_move(_arg0, _arg1)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// Next returns an iterator pointing to the next position after @iter. If @iter
// is the end iterator, the end iterator is returned.
func (i *SequenceIter) Next() *SequenceIter {
	var _arg0 *C.GSequenceIter

	_arg0 = (*C.GSequenceIter)(unsafe.Pointer(i.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_iter_next(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}

// Prev returns an iterator pointing to the previous position before @iter. If
// @iter is the begin iterator, the begin iterator is returned.
func (i *SequenceIter) Prev() *SequenceIter {
	var _arg0 *C.GSequenceIter

	_arg0 = (*C.GSequenceIter)(unsafe.Pointer(i.Native()))

	var _cret *C.GSequenceIter

	cret = C.g_sequence_iter_prev(_arg0)

	var _sequenceIter *SequenceIter

	_sequenceIter = WrapSequenceIter(unsafe.Pointer(_cret))

	return _sequenceIter
}
