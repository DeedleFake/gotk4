// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_PrintSettingsFunc(char*, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_settings_get_type()), F: marshalPrintSettings},
	})
}

type PrintSettingsFunc func(key string, value string)

//export gotk4_PrintSettingsFunc
func gotk4_PrintSettingsFunc(arg0 *C.char, arg1 *C.char, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var key string   // out
	var value string // out

	key = C.GoString(arg0)
	value = C.GoString(arg1)

	fn := v.(PrintSettingsFunc)
	fn(key, value)
}

// PrintSettings: a `GtkPrintSettings` object represents the settings of a print
// dialog in a system-independent way.
//
// The main use for this object is that once you’ve printed you can get a
// settings object that represents the settings the user chose, and the next
// time you print you can pass that object in so that the user doesn’t have to
// re-set all his settings.
//
// Its also possible to enumerate the settings so that you can easily save the
// settings for the next time your app runs, or even store them in a document.
// The predefined keys try to use shared values as much as possible so that
// moving such a document between systems still works.
type PrintSettings interface {
	gextras.Objector

	CopyPrintSettings() PrintSettings

	ForeachPrintSettings(fn PrintSettingsFunc)

	GetPrintSettings(key string) string

	Bool(key string) bool

	Collate() bool

	DefaultSource() string

	Dither() string

	Double(key string) float64

	DoubleWithDefault(key string, def float64) float64

	Duplex() PrintDuplex

	Finishings() string

	Int(key string) int

	IntWithDefault(key string, def int) int

	Length(key string, unit Unit) float64

	MediaType() string

	NCopies() int

	NumberUp() int

	NumberUpLayout() NumberUpLayout

	Orientation() PageOrientation

	OutputBin() string

	PageSet() PageSet

	PaperHeight(unit Unit) float64

	PaperSize() *PaperSize

	PaperWidth(unit Unit) float64

	PrintPages() PrintPages

	Printer() string

	PrinterLpi() float64

	Quality() PrintQuality

	Resolution() int

	ResolutionX() int

	ResolutionY() int

	Reverse() bool

	Scale() float64

	UseColor() bool

	HasKeyPrintSettings(key string) bool

	LoadFilePrintSettings(fileName string) error

	LoadKeyFilePrintSettings(keyFile *glib.KeyFile, groupName string) error

	SetPrintSettings(key string, value string)

	SetBoolPrintSettings(key string, value bool)

	SetCollatePrintSettings(collate bool)

	SetDefaultSourcePrintSettings(defaultSource string)

	SetDitherPrintSettings(dither string)

	SetDoublePrintSettings(key string, value float64)

	SetDuplexPrintSettings(duplex PrintDuplex)

	SetFinishingsPrintSettings(finishings string)

	SetIntPrintSettings(key string, value int)

	SetLengthPrintSettings(key string, value float64, unit Unit)

	SetMediaTypePrintSettings(mediaType string)

	SetNCopiesPrintSettings(numCopies int)

	SetNumberUpPrintSettings(numberUp int)

	SetNumberUpLayoutPrintSettings(numberUpLayout NumberUpLayout)

	SetOrientationPrintSettings(orientation PageOrientation)

	SetOutputBinPrintSettings(outputBin string)

	SetPageRangesPrintSettings(pageRanges []PageRange)

	SetPageSetPrintSettings(pageSet PageSet)

	SetPaperHeightPrintSettings(height float64, unit Unit)

	SetPaperSizePrintSettings(paperSize *PaperSize)

	SetPaperWidthPrintSettings(width float64, unit Unit)

	SetPrintPagesPrintSettings(pages PrintPages)

	SetPrinterPrintSettings(printer string)

	SetPrinterLpiPrintSettings(lpi float64)

	SetQualityPrintSettings(quality PrintQuality)

	SetResolutionPrintSettings(resolution int)

	SetResolutionXYPrintSettings(resolutionX int, resolutionY int)

	SetReversePrintSettings(reverse bool)

	SetScalePrintSettings(scale float64)

	SetUseColorPrintSettings(useColor bool)

	ToFilePrintSettings(fileName string) error

	ToGVariantPrintSettings() *glib.Variant

	ToKeyFilePrintSettings(keyFile *glib.KeyFile, groupName string)

	UnsetPrintSettings(key string)
}

// printSettings implements the PrintSettings class.
type printSettings struct {
	gextras.Objector
}

// WrapPrintSettings wraps a GObject to the right type. It is
// primarily used internally.
func WrapPrintSettings(obj *externglib.Object) PrintSettings {
	return printSettings{
		Objector: obj,
	}
}

func marshalPrintSettings(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPrintSettings(obj), nil
}

func NewPrintSettings() PrintSettings {
	var _cret *C.GtkPrintSettings // in

	_cret = C.gtk_print_settings_new()

	var _printSettings PrintSettings // out

	_printSettings = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PrintSettings)

	return _printSettings
}

func NewPrintSettingsFromFile(fileName string) (PrintSettings, error) {
	var _arg1 *C.char             // out
	var _cret *C.GtkPrintSettings // in
	var _cerr *C.GError           // in

	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_new_from_file(_arg1, &_cerr)

	var _printSettings PrintSettings // out
	var _goerr error                 // out

	_printSettings = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PrintSettings)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _printSettings, _goerr
}

func NewPrintSettingsFromGVariant(variant *glib.Variant) PrintSettings {
	var _arg1 *C.GVariant         // out
	var _cret *C.GtkPrintSettings // in

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant.Native()))

	_cret = C.gtk_print_settings_new_from_gvariant(_arg1)

	var _printSettings PrintSettings // out

	_printSettings = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PrintSettings)

	return _printSettings
}

func NewPrintSettingsFromKeyFile(keyFile *glib.KeyFile, groupName string) (PrintSettings, error) {
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out
	var _cret *C.GtkPrintSettings // in
	var _cerr *C.GError           // in

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile.Native()))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_print_settings_new_from_key_file(_arg1, _arg2, &_cerr)

	var _printSettings PrintSettings // out
	var _goerr error                 // out

	_printSettings = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PrintSettings)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _printSettings, _goerr
}

func (o printSettings) CopyPrintSettings() PrintSettings {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GtkPrintSettings // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_settings_copy(_arg0)

	var _printSettings PrintSettings // out

	_printSettings = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PrintSettings)

	return _printSettings
}

func (s printSettings) ForeachPrintSettings(fn PrintSettingsFunc) {
	var _arg0 *C.GtkPrintSettings    // out
	var _arg1 C.GtkPrintSettingsFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*[0]byte)(C.gotk4_PrintSettingsFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_print_settings_foreach(_arg0, _arg1, _arg2)
}

func (s printSettings) GetPrintSettings(key string) string {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) Bool(key string) bool {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_bool(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) Collate() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_collate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) DefaultSource() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_default_source(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) Dither() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_dither(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) Double(key string) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_double(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) DoubleWithDefault(key string, def float64) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(def)

	_cret = C.gtk_print_settings_get_double_with_default(_arg0, _arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) Duplex() PrintDuplex {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintDuplex    // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_duplex(_arg0)

	var _printDuplex PrintDuplex // out

	_printDuplex = PrintDuplex(_cret)

	return _printDuplex
}

func (s printSettings) Finishings() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_finishings(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) Int(key string) int {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_int(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) IntWithDefault(key string, def int) int {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(def)

	_cret = C.gtk_print_settings_get_int_with_default(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) Length(key string, unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_length(_arg0, _arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) MediaType() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_media_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) NCopies() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_n_copies(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) NumberUp() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_number_up(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) NumberUpLayout() NumberUpLayout {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkNumberUpLayout // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_number_up_layout(_arg0)

	var _numberUpLayout NumberUpLayout // out

	_numberUpLayout = NumberUpLayout(_cret)

	return _numberUpLayout
}

func (s printSettings) Orientation() PageOrientation {
	var _arg0 *C.GtkPrintSettings  // out
	var _cret C.GtkPageOrientation // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_orientation(_arg0)

	var _pageOrientation PageOrientation // out

	_pageOrientation = PageOrientation(_cret)

	return _pageOrientation
}

func (s printSettings) OutputBin() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_output_bin(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) PageSet() PageSet {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPageSet        // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_page_set(_arg0)

	var _pageSet PageSet // out

	_pageSet = PageSet(_cret)

	return _pageSet
}

func (s printSettings) PaperHeight(unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_paper_height(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) PaperSize() *PaperSize {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GtkPaperSize     // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_paper_size(_arg0)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_paperSize, func(v **PaperSize) {
		C.free(unsafe.Pointer(v))
	})

	return _paperSize
}

func (s printSettings) PaperWidth(unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_paper_width(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) PrintPages() PrintPages {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintPages     // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_print_pages(_arg0)

	var _printPages PrintPages // out

	_printPages = PrintPages(_cret)

	return _printPages
}

func (s printSettings) Printer() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_printer(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s printSettings) PrinterLpi() float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_printer_lpi(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) Quality() PrintQuality {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintQuality   // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_quality(_arg0)

	var _printQuality PrintQuality // out

	_printQuality = PrintQuality(_cret)

	return _printQuality
}

func (s printSettings) Resolution() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_resolution(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) ResolutionX() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_resolution_x(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) ResolutionY() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_resolution_y(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s printSettings) Reverse() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_reverse(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) Scale() float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_scale(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s printSettings) UseColor() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_get_use_color(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) HasKeyPrintSettings(key string) bool {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_has_key(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s printSettings) LoadFilePrintSettings(fileName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_load_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s printSettings) LoadKeyFilePrintSettings(keyFile *glib.KeyFile, groupName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile.Native()))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_load_key_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s printSettings) SetPrintSettings(key string, value string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_set(_arg0, _arg1, _arg2)
}

func (s printSettings) SetBoolPrintSettings(key string, value bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	if value {
		_arg2 = C.TRUE
	}

	C.gtk_print_settings_set_bool(_arg0, _arg1, _arg2)
}

func (s printSettings) SetCollatePrintSettings(collate bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	if collate {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_collate(_arg0, _arg1)
}

func (s printSettings) SetDefaultSourcePrintSettings(defaultSource string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(defaultSource))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_default_source(_arg0, _arg1)
}

func (s printSettings) SetDitherPrintSettings(dither string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(dither))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_dither(_arg0, _arg1)
}

func (s printSettings) SetDoublePrintSettings(key string, value float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)

	C.gtk_print_settings_set_double(_arg0, _arg1, _arg2)
}

func (s printSettings) SetDuplexPrintSettings(duplex PrintDuplex) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintDuplex    // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPrintDuplex(duplex)

	C.gtk_print_settings_set_duplex(_arg0, _arg1)
}

func (s printSettings) SetFinishingsPrintSettings(finishings string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(finishings))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_finishings(_arg0, _arg1)
}

func (s printSettings) SetIntPrintSettings(key string, value int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(value)

	C.gtk_print_settings_set_int(_arg0, _arg1, _arg2)
}

func (s printSettings) SetLengthPrintSettings(key string, value float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out
	var _arg3 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)
	_arg3 = C.GtkUnit(unit)

	C.gtk_print_settings_set_length(_arg0, _arg1, _arg2, _arg3)
}

func (s printSettings) SetMediaTypePrintSettings(mediaType string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(mediaType))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_media_type(_arg0, _arg1)
}

func (s printSettings) SetNCopiesPrintSettings(numCopies int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(numCopies)

	C.gtk_print_settings_set_n_copies(_arg0, _arg1)
}

func (s printSettings) SetNumberUpPrintSettings(numberUp int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(numberUp)

	C.gtk_print_settings_set_number_up(_arg0, _arg1)
}

func (s printSettings) SetNumberUpLayoutPrintSettings(numberUpLayout NumberUpLayout) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkNumberUpLayout // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkNumberUpLayout(numberUpLayout)

	C.gtk_print_settings_set_number_up_layout(_arg0, _arg1)
}

func (s printSettings) SetOrientationPrintSettings(orientation PageOrientation) {
	var _arg0 *C.GtkPrintSettings  // out
	var _arg1 C.GtkPageOrientation // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPageOrientation(orientation)

	C.gtk_print_settings_set_orientation(_arg0, _arg1)
}

func (s printSettings) SetOutputBinPrintSettings(outputBin string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(outputBin))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_output_bin(_arg0, _arg1)
}

func (s printSettings) SetPageRangesPrintSettings(pageRanges []PageRange) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GtkPageRange
	var _arg2 C.int

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg2 = C.int(len(pageRanges))
	_arg1 = (*C.GtkPageRange)(unsafe.Pointer(&pageRanges[0]))

	C.gtk_print_settings_set_page_ranges(_arg0, _arg1, _arg2)
}

func (s printSettings) SetPageSetPrintSettings(pageSet PageSet) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPageSet        // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPageSet(pageSet)

	C.gtk_print_settings_set_page_set(_arg0, _arg1)
}

func (s printSettings) SetPaperHeightPrintSettings(height float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out
	var _arg2 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(height)
	_arg2 = C.GtkUnit(unit)

	C.gtk_print_settings_set_paper_height(_arg0, _arg1, _arg2)
}

func (s printSettings) SetPaperSizePrintSettings(paperSize *PaperSize) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GtkPaperSize     // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(paperSize.Native()))

	C.gtk_print_settings_set_paper_size(_arg0, _arg1)
}

func (s printSettings) SetPaperWidthPrintSettings(width float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out
	var _arg2 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(width)
	_arg2 = C.GtkUnit(unit)

	C.gtk_print_settings_set_paper_width(_arg0, _arg1, _arg2)
}

func (s printSettings) SetPrintPagesPrintSettings(pages PrintPages) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintPages     // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPrintPages(pages)

	C.gtk_print_settings_set_print_pages(_arg0, _arg1)
}

func (s printSettings) SetPrinterPrintSettings(printer string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(printer))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_printer(_arg0, _arg1)
}

func (s printSettings) SetPrinterLpiPrintSettings(lpi float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(lpi)

	C.gtk_print_settings_set_printer_lpi(_arg0, _arg1)
}

func (s printSettings) SetQualityPrintSettings(quality PrintQuality) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintQuality   // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPrintQuality(quality)

	C.gtk_print_settings_set_quality(_arg0, _arg1)
}

func (s printSettings) SetResolutionPrintSettings(resolution int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(resolution)

	C.gtk_print_settings_set_resolution(_arg0, _arg1)
}

func (s printSettings) SetResolutionXYPrintSettings(resolutionX int, resolutionY int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out
	var _arg2 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(resolutionX)
	_arg2 = C.int(resolutionY)

	C.gtk_print_settings_set_resolution_xy(_arg0, _arg1, _arg2)
}

func (s printSettings) SetReversePrintSettings(reverse bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	if reverse {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_reverse(_arg0, _arg1)
}

func (s printSettings) SetScalePrintSettings(scale float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = C.double(scale)

	C.gtk_print_settings_set_scale(_arg0, _arg1)
}

func (s printSettings) SetUseColorPrintSettings(useColor bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	if useColor {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_use_color(_arg0, _arg1)
}

func (s printSettings) ToFilePrintSettings(fileName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s printSettings) ToGVariantPrintSettings() *glib.Variant {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GVariant         // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_print_settings_to_gvariant(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))

	return _variant
}

func (s printSettings) ToKeyFilePrintSettings(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile.Native()))
	_arg2 = (*C.char)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_print_settings_to_key_file(_arg0, _arg1, _arg2)
}

func (s printSettings) UnsetPrintSettings(key string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_unset(_arg0, _arg1)
}

// PageRange: a range of pages to print.
//
// See also [method@Gtk.PrintSettings.set_page_ranges].
type PageRange C.GtkPageRange

// WrapPageRange wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPageRange(ptr unsafe.Pointer) *PageRange {
	return (*PageRange)(ptr)
}

// Native returns the underlying C source pointer.
func (p *PageRange) Native() unsafe.Pointer {
	return unsafe.Pointer(p)
}
