// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_hypertext_get_type()), F: marshalHypertexter},
	})
}

// HypertextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type HypertextOverrider interface {
	// Link gets the link in this hypertext document at index link_index
	Link(linkIndex int) *Hyperlink
	// LinkIndex gets the index into the array of hyperlinks that is associated
	// with the character specified by char_index.
	LinkIndex(charIndex int) int
	// NLinks gets the number of links within this hypertext document.
	NLinks() int
	LinkSelected(linkIndex int)
}

// Hypertext: interface used for objects which implement linking between
// multiple resource or content locations, or multiple 'markers' within a single
// document. A Hypertext instance is associated with one or more Hyperlinks,
// which are associated with particular offsets within the Hypertext's included
// content. While this interface is derived from Text, there is no requirement
// that Hypertext instances have textual content; they may implement Image as
// well, and Hyperlinks need not have non-zero text offsets.
type Hypertext struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Hypertext)(nil)

// Hypertexter describes Hypertext's abstract methods.
type Hypertexter interface {
	// Link gets the link in this hypertext document at index link_index
	Link(linkIndex int) *Hyperlink
	// LinkIndex gets the index into the array of hyperlinks that is associated
	// with the character specified by char_index.
	LinkIndex(charIndex int) int
	// NLinks gets the number of links within this hypertext document.
	NLinks() int
}

var _ Hypertexter = (*Hypertext)(nil)

func wrapHypertext(obj *externglib.Object) *Hypertext {
	return &Hypertext{
		Object: obj,
	}
}

func marshalHypertexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHypertext(obj), nil
}

// Link gets the link in this hypertext document at index link_index
func (hypertext *Hypertext) Link(linkIndex int) *Hyperlink {
	var _arg0 *C.AtkHypertext // out
	var _arg1 C.gint          // out
	var _cret *C.AtkHyperlink // in

	_arg0 = (*C.AtkHypertext)(unsafe.Pointer(hypertext.Native()))
	_arg1 = C.gint(linkIndex)

	_cret = C.atk_hypertext_get_link(_arg0, _arg1)

	var _hyperlink *Hyperlink // out

	_hyperlink = wrapHyperlink(externglib.Take(unsafe.Pointer(_cret)))

	return _hyperlink
}

// LinkIndex gets the index into the array of hyperlinks that is associated with
// the character specified by char_index.
func (hypertext *Hypertext) LinkIndex(charIndex int) int {
	var _arg0 *C.AtkHypertext // out
	var _arg1 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkHypertext)(unsafe.Pointer(hypertext.Native()))
	_arg1 = C.gint(charIndex)

	_cret = C.atk_hypertext_get_link_index(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NLinks gets the number of links within this hypertext document.
func (hypertext *Hypertext) NLinks() int {
	var _arg0 *C.AtkHypertext // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkHypertext)(unsafe.Pointer(hypertext.Native()))

	_cret = C.atk_hypertext_get_n_links(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
