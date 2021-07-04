// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_file_filter_flags_get_type()), F: marshalFileFilterFlags},
		{T: externglib.Type(C.gtk_file_filter_get_type()), F: marshalFileFilter},
	})
}

// FileFilterFlags: these flags indicate what parts of a FileFilterInfo struct
// are filled or need to be filled.
type FileFilterFlags int

const (
	// FileFilterFlagsFilename: the filename of the file being tested
	FileFilterFlagsFilename FileFilterFlags = 0b1
	// FileFilterFlagsURI: the URI for the file being tested
	FileFilterFlagsURI FileFilterFlags = 0b10
	// FileFilterFlagsDisplayName: the string that will be used to display the
	// file in the file chooser
	FileFilterFlagsDisplayName FileFilterFlags = 0b100
	// FileFilterFlagsMIMEType: the mime type of the file
	FileFilterFlagsMIMEType FileFilterFlags = 0b1000
)

func marshalFileFilterFlags(p uintptr) (interface{}, error) {
	return FileFilterFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FileFilterFunc: the type of function that is used with custom filters, see
// gtk_file_filter_add_custom().
type FileFilterFunc func(filterInfo *FileFilterInfo, ok bool)

//export gotk4_FileFilterFunc
func gotk4_FileFilterFunc(arg0 *C.GtkFileFilterInfo, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var filterInfo *FileFilterInfo // out

	filterInfo = (*FileFilterInfo)(unsafe.Pointer(arg0))

	fn := v.(FileFilterFunc)
	ok := fn(filterInfo)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FileFilter: a GtkFileFilter can be used to restrict the files being shown in
// a FileChooser. Files can be filtered based on their name (with
// gtk_file_filter_add_pattern()), on their mime type (with
// gtk_file_filter_add_mime_type()), or by a custom filter function (with
// gtk_file_filter_add_custom()).
//
// Filtering by mime types handles aliasing and subclassing of mime types; e.g.
// a filter for text/plain also matches a file with mime type application/rtf,
// since application/rtf is a subclass of text/plain. Note that FileFilter
// allows wildcards for the subtype of a mime type, so you can e.g. filter for
// image/\*.
//
// Normally, filters are used by adding them to a FileChooser, see
// gtk_file_chooser_add_filter(), but it is also possible to manually use a
// filter on a file with gtk_file_filter_filter().
//
//
// GtkFileFilter as GtkBuildable
//
// The GtkFileFilter implementation of the GtkBuildable interface supports
// adding rules using the <mime-types>, <patterns> and <applications> elements
// and listing the rules within. Specifying a <mime-type> or <pattern> has the
// same effect as as calling gtk_file_filter_add_mime_type() or
// gtk_file_filter_add_pattern().
//
// An example of a UI definition fragment specifying GtkFileFilter rules:
//
//    <object class="GtkFileFilter">
//      <mime-types>
//        <mime-type>text/plain</mime-type>
//        <mime-type>image/ *</mime-type>
//      </mime-types>
//      <patterns>
//        <pattern>*.txt</pattern>
//        <pattern>*.png</pattern>
//      </patterns>
//    </object>
type FileFilter interface {
	Buildable

	AddMIMETypeFileFilter(mimeType string)

	AddPatternFileFilter(pattern string)

	AddPixbufFormatsFileFilter()

	FilterFileFilter(filterInfo *FileFilterInfo) bool

	GetName() string

	Needed() FileFilterFlags

	SetNameFileFilter(name string)

	ToGVariantFileFilter() *glib.Variant
}

// fileFilter implements the FileFilter class.
type fileFilter struct {
	gextras.Objector
}

// WrapFileFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileFilter(obj *externglib.Object) FileFilter {
	return fileFilter{
		Objector: obj,
	}
}

func marshalFileFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileFilter(obj), nil
}

func NewFileFilter() FileFilter {
	var _cret *C.GtkFileFilter // in

	_cret = C.gtk_file_filter_new()

	var _fileFilter FileFilter // out

	_fileFilter = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FileFilter)

	return _fileFilter
}

func NewFileFilterFromGVariant(variant *glib.Variant) FileFilter {
	var _arg1 *C.GVariant      // out
	var _cret *C.GtkFileFilter // in

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant.Native()))

	_cret = C.gtk_file_filter_new_from_gvariant(_arg1)

	var _fileFilter FileFilter // out

	_fileFilter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileFilter)

	return _fileFilter
}

func (f fileFilter) AddMIMETypeFileFilter(mimeType string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_filter_add_mime_type(_arg0, _arg1)
}

func (f fileFilter) AddPatternFileFilter(pattern string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(pattern))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_filter_add_pattern(_arg0, _arg1)
}

func (f fileFilter) AddPixbufFormatsFileFilter() {
	var _arg0 *C.GtkFileFilter // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))

	C.gtk_file_filter_add_pixbuf_formats(_arg0)
}

func (f fileFilter) FilterFileFilter(filterInfo *FileFilterInfo) bool {
	var _arg0 *C.GtkFileFilter     // out
	var _arg1 *C.GtkFileFilterInfo // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkFileFilterInfo)(unsafe.Pointer(filterInfo.Native()))

	_cret = C.gtk_file_filter_filter(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f fileFilter) GetName() string {
	var _arg0 *C.GtkFileFilter // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_file_filter_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (f fileFilter) Needed() FileFilterFlags {
	var _arg0 *C.GtkFileFilter     // out
	var _cret C.GtkFileFilterFlags // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_file_filter_get_needed(_arg0)

	var _fileFilterFlags FileFilterFlags // out

	_fileFilterFlags = FileFilterFlags(_cret)

	return _fileFilterFlags
}

func (f fileFilter) SetNameFileFilter(name string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_filter_set_name(_arg0, _arg1)
}

func (f fileFilter) ToGVariantFileFilter() *glib.Variant {
	var _arg0 *C.GtkFileFilter // out
	var _cret *C.GVariant      // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_file_filter_to_gvariant(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))

	return _variant
}

func (b fileFilter) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b fileFilter) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b fileFilter) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b fileFilter) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b fileFilter) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b fileFilter) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b fileFilter) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b fileFilter) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b fileFilter) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b fileFilter) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

// FileFilterInfo: a FileFilterInfo-struct is used to pass information about the
// tested file to gtk_file_filter_filter().
type FileFilterInfo C.GtkFileFilterInfo

// WrapFileFilterInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileFilterInfo(ptr unsafe.Pointer) *FileFilterInfo {
	return (*FileFilterInfo)(ptr)
}

// Native returns the underlying C source pointer.
func (f *FileFilterInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(f)
}
