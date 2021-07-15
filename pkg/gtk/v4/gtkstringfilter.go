// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_string_filter_match_mode_get_type()), F: marshalStringFilterMatchMode},
		{T: externglib.Type(C.gtk_string_filter_get_type()), F: marshalStringFilterer},
	})
}

// StringFilterMatchMode specifies how search strings are matched inside text.
type StringFilterMatchMode int

const (
	// Exact: search string and text must match exactly.
	StringFilterMatchModeExact StringFilterMatchMode = iota
	// Substring: search string must be contained as a substring inside the
	// text.
	StringFilterMatchModeSubstring
	// Prefix: text must begin with the search string.
	StringFilterMatchModePrefix
)

func marshalStringFilterMatchMode(p uintptr) (interface{}, error) {
	return StringFilterMatchMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// StringFilter: GtkStringFilter determines whether to include items by
// comparing strings to a fixed search term.
//
// The strings are obtained from the items by evaluating a GtkExpression set
// with gtk.StringFilter.SetExpression(), and they are compared against a search
// term set with gtk.StringFilter.SetSearch().
//
// GtkStringFilter has several different modes of comparison - it can match the
// whole string, just a prefix, or any substring. Use
// gtk.StringFilter.SetMatchMode() choose a mode.
//
// It is also possible to make case-insensitive comparisons, with
// gtk.StringFilter.SetIgnoreCase().
type StringFilter struct {
	Filter
}

var _ gextras.Nativer = (*StringFilter)(nil)

func wrapStringFilter(obj *externglib.Object) *StringFilter {
	return &StringFilter{
		Filter: Filter{
			Object: obj,
		},
	}
}

func marshalStringFilterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStringFilter(obj), nil
}

// NewStringFilter creates a new string filter.
//
// You will want to set up the filter by providing a string to search for and by
// providing a property to look up on the item.
func NewStringFilter(expression Expressioner) *StringFilter {
	var _arg1 *C.GtkExpression   // out
	var _cret *C.GtkStringFilter // in

	_arg1 = (*C.GtkExpression)(unsafe.Pointer((expression).(gextras.Nativer).Native()))

	_cret = C.gtk_string_filter_new(_arg1)

	var _stringFilter *StringFilter // out

	_stringFilter = wrapStringFilter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stringFilter
}

// Expression gets the expression that the string filter uses to obtain strings
// from items.
func (self *StringFilter) Expression() *Expression {
	var _arg0 *C.GtkStringFilter // out
	var _cret *C.GtkExpression   // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_string_filter_get_expression(_arg0)

	var _expression *Expression // out

	_expression = wrapExpression(externglib.Take(unsafe.Pointer(_cret)))

	return _expression
}

// IgnoreCase returns whether the filter ignores case differences.
func (self *StringFilter) IgnoreCase() bool {
	var _arg0 *C.GtkStringFilter // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_string_filter_get_ignore_case(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MatchMode returns the match mode that the filter is using.
func (self *StringFilter) MatchMode() StringFilterMatchMode {
	var _arg0 *C.GtkStringFilter         // out
	var _cret C.GtkStringFilterMatchMode // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_string_filter_get_match_mode(_arg0)

	var _stringFilterMatchMode StringFilterMatchMode // out

	_stringFilterMatchMode = StringFilterMatchMode(_cret)

	return _stringFilterMatchMode
}

// Search gets the search term.
func (self *StringFilter) Search() string {
	var _arg0 *C.GtkStringFilter // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_string_filter_get_search(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetExpression sets the expression that the string filter uses to obtain
// strings from items.
//
// The expression must have a value type of G_TYPE_STRING.
func (self *StringFilter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 *C.GtkExpression   // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkExpression)(unsafe.Pointer((expression).(gextras.Nativer).Native()))

	C.gtk_string_filter_set_expression(_arg0, _arg1)
}

// SetIgnoreCase sets whether the filter ignores case differences.
func (self *StringFilter) SetIgnoreCase(ignoreCase bool) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(self.Native()))
	if ignoreCase {
		_arg1 = C.TRUE
	}

	C.gtk_string_filter_set_ignore_case(_arg0, _arg1)
}

// SetMatchMode sets the match mode for the filter.
func (self *StringFilter) SetMatchMode(mode StringFilterMatchMode) {
	var _arg0 *C.GtkStringFilter         // out
	var _arg1 C.GtkStringFilterMatchMode // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkStringFilterMatchMode(mode)

	C.gtk_string_filter_set_match_mode(_arg0, _arg1)
}

// SetSearch sets the string to search for.
func (self *StringFilter) SetSearch(search string) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(search)))

	C.gtk_string_filter_set_search(_arg0, _arg1)
}
