// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

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
		{T: externglib.Type(C.gtk_css_section_type_get_type()), F: marshalCSSSectionType},
		{T: externglib.Type(C.gtk_css_section_get_type()), F: marshalCSSSection},
	})
}

// CSSSectionType: the different types of sections indicate parts of a CSS
// document as parsed by GTK’s CSS parser. They are oriented towards the CSS
// Grammar (http://www.w3.org/TR/CSS21/grammar.html), but may contain
// extensions.
//
// More types might be added in the future as the parser incorporates more
// features.
type CSSSectionType int

const (
	// Document: the section describes a complete document. This section time is
	// the only one where gtk_css_section_get_parent() might return nil.
	CSSSectionTypeDocument CSSSectionType = iota
	// Import: the section defines an import rule.
	CSSSectionTypeImport
	// ColorDefinition: the section defines a color. This is a GTK extension to
	// CSS.
	CSSSectionTypeColorDefinition
	// BindingSet: the section defines a binding set. This is a GTK extension to
	// CSS.
	CSSSectionTypeBindingSet
	// Ruleset: the section defines a CSS ruleset.
	CSSSectionTypeRuleset
	// Selector: the section defines a CSS selector.
	CSSSectionTypeSelector
	// Declaration: the section defines the declaration of a CSS variable.
	CSSSectionTypeDeclaration
	// Value: the section defines the value of a CSS declaration.
	CSSSectionTypeValue
	// Keyframes: the section defines keyframes. See [CSS
	// Animations](http://dev.w3.org/csswg/css3-animations/#keyframes) for
	// details. Since 3.6
	CSSSectionTypeKeyframes
)

func marshalCSSSectionType(p uintptr) (interface{}, error) {
	return CSSSectionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CSSSection defines a part of a CSS document. Because sections are nested into
// one another, you can use gtk_css_section_get_parent() to get the containing
// region.
type CSSSection struct {
	native C.GtkCssSection
}

// WrapCSSSection wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCSSSection(ptr unsafe.Pointer) *CSSSection {
	return (*CSSSection)(ptr)
}

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*CSSSection)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CSSSection) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// EndLine returns the line in the CSS document where this section end. The line
// number is 0-indexed, so the first line of the document will return 0. This
// value may change in future invocations of this function if @section is not
// yet parsed completely. This will for example happen in the
// GtkCssProvider::parsing-error signal. The end position and line may be
// identical to the start position and line for sections which failed to parse
// anything successfully.
func (s *CSSSection) EndLine() uint {
	var _arg0 *C.GtkCssSection // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(*CSSSection))

	_cret = C.gtk_css_section_get_end_line(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// EndPosition returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_end_line(). This value may change in future
// invocations of this function if @section is not yet parsed completely. This
// will for example happen in the GtkCssProvider::parsing-error signal. The end
// position and line may be identical to the start position and line for
// sections which failed to parse anything successfully.
func (s *CSSSection) EndPosition() uint {
	var _arg0 *C.GtkCssSection // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(*CSSSection))

	_cret = C.gtk_css_section_get_end_position(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Parent gets the parent section for the given @section. The parent section is
// the section that contains this @section. A special case are sections of type
// K_CSS_SECTION_DOCUMENT. Their parent will either be nil if they are the
// original CSS document that was loaded by gtk_css_provider_load_from_file() or
// a section of type K_CSS_SECTION_IMPORT if it was loaded with an import rule
// from a different file.
func (s *CSSSection) Parent() *CSSSection {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GtkCssSection // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(*CSSSection))

	_cret = C.gtk_css_section_get_parent(_arg0)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(unsafe.Pointer(*C.GtkCssSection))
	C.gtk_css_section_ref(_cret)
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.gtk_css_section_unref((*C.GtkCssSection)(unsafe.Pointer(v)))
	})

	return _cssSection
}

// SectionType gets the type of information that @section describes.
func (s *CSSSection) SectionType() CSSSectionType {
	var _arg0 *C.GtkCssSection    // out
	var _cret C.GtkCssSectionType // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(*CSSSection))

	_cret = C.gtk_css_section_get_section_type(_arg0)

	var _cssSectionType CSSSectionType // out

	_cssSectionType = (CSSSectionType)(C.GtkCssSectionType)

	return _cssSectionType
}

// StartLine returns the line in the CSS document where this section starts. The
// line number is 0-indexed, so the first line of the document will return 0.
func (s *CSSSection) StartLine() uint {
	var _arg0 *C.GtkCssSection // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(*CSSSection))

	_cret = C.gtk_css_section_get_start_line(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// StartPosition returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_start_line().
func (s *CSSSection) StartPosition() uint {
	var _arg0 *C.GtkCssSection // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(*CSSSection))

	_cret = C.gtk_css_section_get_start_position(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Ref increments the reference count on @section.
func (s *CSSSection) ref() *CSSSection {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GtkCssSection // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(*CSSSection))

	_cret = C.gtk_css_section_ref(_arg0)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(unsafe.Pointer(*C.GtkCssSection))
	C.gtk_css_section_ref(_cret)
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.gtk_css_section_unref((*C.GtkCssSection)(unsafe.Pointer(v)))
	})

	return _cssSection
}

// Unref decrements the reference count on @section, freeing the structure if
// the reference count reaches 0.
func (s *CSSSection) unref() {
	var _arg0 *C.GtkCssSection // out

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(*CSSSection))

	C.gtk_css_section_unref(_arg0)
}
