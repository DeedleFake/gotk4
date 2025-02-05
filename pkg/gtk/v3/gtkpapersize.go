// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypePaperSize = coreglib.Type(C.gtk_paper_size_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePaperSize, F: marshalPaperSize},
	})
}

// PAPER_NAME_A3: name for the A3 paper size.
const PAPER_NAME_A3 = "iso_a3"

// PAPER_NAME_A4: name for the A4 paper size.
const PAPER_NAME_A4 = "iso_a4"

// PAPER_NAME_A5: name for the A5 paper size.
const PAPER_NAME_A5 = "iso_a5"

// PAPER_NAME_B5: name for the B5 paper size.
const PAPER_NAME_B5 = "iso_b5"

// PAPER_NAME_EXECUTIVE: name for the Executive paper size.
const PAPER_NAME_EXECUTIVE = "na_executive"

// PAPER_NAME_LEGAL: name for the Legal paper size.
const PAPER_NAME_LEGAL = "na_legal"

// PAPER_NAME_LETTER: name for the Letter paper size.
const PAPER_NAME_LETTER = "na_letter"

// PaperSize handles paper sizes. It uses the standard called PWG 5101.1-2002
// PWG: Standard for Media Standardized Names
// (http://www.pwg.org/standards.html) to name the paper sizes (and to get the
// data for the page sizes). In addition to standard paper sizes, GtkPaperSize
// allows to construct custom paper sizes with arbitrary dimensions.
//
// The PaperSize object stores not only the dimensions (width and height) of a
// paper size and its name, it also provides default [print
// margins][print-margins].
//
// Printing support has been added in GTK+ 2.10.
//
// An instance of this type is always passed by reference.
type PaperSize struct {
	*paperSize
}

// paperSize is the struct that's finalized.
type paperSize struct {
	native *C.GtkPaperSize
}

func marshalPaperSize(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &PaperSize{&paperSize{(*C.GtkPaperSize)(b)}}, nil
}

// NewPaperSize constructs a struct PaperSize.
func NewPaperSize(name string) *PaperSize {
	var _arg1 *C.gchar        // out
	var _cret *C.GtkPaperSize // in

	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_paper_size_new(_arg1)
	runtime.KeepAlive(name)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// NewPaperSizeCustom constructs a struct PaperSize.
func NewPaperSizeCustom(name string, displayName string, width float64, height float64, unit Unit) *PaperSize {
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // out
	var _arg3 C.gdouble       // out
	var _arg4 C.gdouble       // out
	var _arg5 C.GtkUnit       // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(displayName)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gdouble(width)
	_arg4 = C.gdouble(height)
	_arg5 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_new_custom(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(name)
	runtime.KeepAlive(displayName)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(unit)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// NewPaperSizeFromGVariant constructs a struct PaperSize.
func NewPaperSizeFromGVariant(variant *glib.Variant) *PaperSize {
	var _arg1 *C.GVariant     // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))

	_cret = C.gtk_paper_size_new_from_gvariant(_arg1)
	runtime.KeepAlive(variant)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// NewPaperSizeFromIPP constructs a struct PaperSize.
func NewPaperSizeFromIPP(ippName string, width float64, height float64) *PaperSize {
	var _arg1 *C.gchar        // out
	var _arg2 C.gdouble       // out
	var _arg3 C.gdouble       // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(ippName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(width)
	_arg3 = C.gdouble(height)

	_cret = C.gtk_paper_size_new_from_ipp(_arg1, _arg2, _arg3)
	runtime.KeepAlive(ippName)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// NewPaperSizeFromKeyFile constructs a struct PaperSize.
func NewPaperSizeFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PaperSize, error) {
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out
	var _cret *C.GtkPaperSize // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_paper_size_new_from_key_file(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)

	var _paperSize *PaperSize // out
	var _goerr error          // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _paperSize, _goerr
}

// NewPaperSizeFromPPD constructs a struct PaperSize.
func NewPaperSizeFromPPD(ppdName string, ppdDisplayName string, width float64, height float64) *PaperSize {
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // out
	var _arg3 C.gdouble       // out
	var _arg4 C.gdouble       // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(ppdName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(ppdDisplayName)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gdouble(width)
	_arg4 = C.gdouble(height)

	_cret = C.gtk_paper_size_new_from_ppd(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(ppdName)
	runtime.KeepAlive(ppdDisplayName)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// Copy copies an existing PaperSize.
//
// The function returns the following values:
//
//    - paperSize: copy of other.
//
func (other *PaperSize) Copy() *PaperSize {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.GtkPaperSize // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = C.gtk_paper_size_copy(_arg0)
	runtime.KeepAlive(other)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// DefaultBottomMargin gets the default bottom margin for the PaperSize.
//
// The function takes the following parameters:
//
//    - unit for the return value, not GTK_UNIT_NONE.
//
// The function returns the following values:
//
//    - gdouble: default bottom margin.
//
func (size *PaperSize) DefaultBottomMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_bottom_margin(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultLeftMargin gets the default left margin for the PaperSize.
//
// The function takes the following parameters:
//
//    - unit for the return value, not GTK_UNIT_NONE.
//
// The function returns the following values:
//
//    - gdouble: default left margin.
//
func (size *PaperSize) DefaultLeftMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_left_margin(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultRightMargin gets the default right margin for the PaperSize.
//
// The function takes the following parameters:
//
//    - unit for the return value, not GTK_UNIT_NONE.
//
// The function returns the following values:
//
//    - gdouble: default right margin.
//
func (size *PaperSize) DefaultRightMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_right_margin(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultTopMargin gets the default top margin for the PaperSize.
//
// The function takes the following parameters:
//
//    - unit for the return value, not GTK_UNIT_NONE.
//
// The function returns the following values:
//
//    - gdouble: default top margin.
//
func (size *PaperSize) DefaultTopMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_top_margin(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DisplayName gets the human-readable name of the PaperSize.
//
// The function returns the following values:
//
//    - utf8: human-readable name of size.
//
func (size *PaperSize) DisplayName() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_get_display_name(_arg0)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Height gets the paper height of the PaperSize, in units of unit.
//
// The function takes the following parameters:
//
//    - unit for the return value, not GTK_UNIT_NONE.
//
// The function returns the following values:
//
//    - gdouble: paper height.
//
func (size *PaperSize) Height(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_height(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Name gets the name of the PaperSize.
//
// The function returns the following values:
//
//    - utf8: name of size.
//
func (size *PaperSize) Name() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_get_name(_arg0)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PPDName gets the PPD name of the PaperSize, which may be NULL.
//
// The function returns the following values:
//
//    - utf8: PPD name of size.
//
func (size *PaperSize) PPDName() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_get_ppd_name(_arg0)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Width gets the paper width of the PaperSize, in units of unit.
//
// The function takes the following parameters:
//
//    - unit for the return value, not GTK_UNIT_NONE.
//
// The function returns the following values:
//
//    - gdouble: paper width.
//
func (size *PaperSize) Width(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.gdouble       // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_width(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// IsCustom returns TRUE if size is not a standard paper size.
//
// The function returns the following values:
//
//    - ok: whether size is a custom paper size.
//
func (size *PaperSize) IsCustom() bool {
	var _arg0 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_is_custom(_arg0)
	runtime.KeepAlive(size)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEqual compares two PaperSize objects.
//
// The function takes the following parameters:
//
//    - size2: another PaperSize object.
//
// The function returns the following values:
//
//    - ok: TRUE, if size1 and size2 represent the same paper size.
//
func (size1 *PaperSize) IsEqual(size2 *PaperSize) bool {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size1)))
	_arg1 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size2)))

	_cret = C.gtk_paper_size_is_equal(_arg0, _arg1)
	runtime.KeepAlive(size1)
	runtime.KeepAlive(size2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsIPP returns TRUE if size is an IPP standard paper size.
//
// The function returns the following values:
//
//    - ok: whether size is not an IPP custom paper size.
//
func (size *PaperSize) IsIPP() bool {
	var _arg0 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_is_ipp(_arg0)
	runtime.KeepAlive(size)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSize changes the dimensions of a size to width x height.
//
// The function takes the following parameters:
//
//    - width: new width in units of unit.
//    - height: new height in units of unit.
//    - unit for width and height.
//
func (size *PaperSize) SetSize(width float64, height float64, unit Unit) {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.gdouble       // out
	var _arg2 C.gdouble       // out
	var _arg3 C.GtkUnit       // out

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.gdouble(width)
	_arg2 = C.gdouble(height)
	_arg3 = C.GtkUnit(unit)

	C.gtk_paper_size_set_size(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(size)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(unit)
}

// ToGVariant: serialize a paper size to an a{sv} variant.
//
// The function returns the following values:
//
//    - variant: new, floating, #GVariant.
//
func (paperSize *PaperSize) ToGVariant() *glib.Variant {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(paperSize)))

	_cret = C.gtk_paper_size_to_gvariant(_arg0)
	runtime.KeepAlive(paperSize)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// ToKeyFile: this function adds the paper size from size to key_file.
//
// The function takes the following parameters:
//
//    - keyFile to save the paper size to.
//    - groupName: group to add the settings to in key_file.
//
func (size *PaperSize) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_paper_size_to_key_file(_arg0, _arg1, _arg2)
	runtime.KeepAlive(size)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)
}
