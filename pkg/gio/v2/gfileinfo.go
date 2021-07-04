// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_info_get_type()), F: marshalFileInfo},
	})
}

// FileInfo: functionality for manipulating basic metadata for files. Info
// implements methods for getting information that all files should contain, and
// allows for manipulation of extended attributes.
//
// See [GFileAttribute][gio-GFileAttribute] for more information on how GIO
// handles file attributes.
//
// To obtain a Info for a #GFile, use g_file_query_info() (or its async
// variant). To obtain a Info for a file input or output stream, use
// g_file_input_stream_query_info() or g_file_output_stream_query_info() (or
// their async variants).
//
// To change the actual attributes of a file, you should then set the attribute
// in the Info and call g_file_set_attributes_from_info() or
// g_file_set_attributes_async() on a GFile.
//
// However, not all attributes can be changed in the file. For instance, the
// actual size of a file cannot be changed via g_file_info_set_size(). You may
// call g_file_query_settable_attributes() and
// g_file_query_writable_namespaces() to discover the settable attributes of a
// particular file at runtime.
//
// AttributeMatcher allows for searching through a Info for attributes.
type FileInfo interface {
	gextras.Objector

	ClearStatusFileInfo()

	CopyIntoFileInfo(destInfo FileInfo)

	DupFileInfo() FileInfo

	AttributeAsString(attribute string) string

	AttributeBoolean(attribute string) bool

	AttributeByteString(attribute string) string

	AttributeData(attribute string) (FileAttributeType, interface{}, FileAttributeStatus, bool)

	AttributeInt32(attribute string) int32

	AttributeInt64(attribute string) int64

	AttributeObject(attribute string) gextras.Objector

	AttributeStatus(attribute string) FileAttributeStatus

	AttributeString(attribute string) string

	AttributeStringv(attribute string) []string

	AttributeType(attribute string) FileAttributeType

	AttributeUint32(attribute string) uint32

	AttributeUint64(attribute string) uint64

	ContentType() string

	DisplayName() string

	EditName() string

	Etag() string

	FileType() FileType

	Icon() Icon

	IsBackup() bool

	IsHidden() bool

	IsSymlink() bool

	ModificationTime() glib.TimeVal

	Name() string

	Size() int64

	SortOrder() int32

	SymbolicIcon() Icon

	SymlinkTarget() string

	HasAttributeFileInfo(attribute string) bool

	HasNamespaceFileInfo(nameSpace string) bool

	ListAttributesFileInfo(nameSpace string) []string

	RemoveAttributeFileInfo(attribute string)

	SetAttributeFileInfo(attribute string, typ FileAttributeType, valueP interface{})

	SetAttributeBooleanFileInfo(attribute string, attrValue bool)

	SetAttributeByteStringFileInfo(attribute string, attrValue string)

	SetAttributeInt32FileInfo(attribute string, attrValue int32)

	SetAttributeInt64FileInfo(attribute string, attrValue int64)

	SetAttributeMaskFileInfo(mask *FileAttributeMatcher)

	SetAttributeObjectFileInfo(attribute string, attrValue gextras.Objector)

	SetAttributeStatusFileInfo(attribute string, status FileAttributeStatus) bool

	SetAttributeStringFileInfo(attribute string, attrValue string)

	SetAttributeStringvFileInfo(attribute string, attrValue []string)

	SetAttributeUint32FileInfo(attribute string, attrValue uint32)

	SetAttributeUint64FileInfo(attribute string, attrValue uint64)

	SetContentTypeFileInfo(contentType string)

	SetDisplayNameFileInfo(displayName string)

	SetEditNameFileInfo(editName string)

	SetFileTypeFileInfo(typ FileType)

	SetIconFileInfo(icon Icon)

	SetIsHiddenFileInfo(isHidden bool)

	SetIsSymlinkFileInfo(isSymlink bool)

	SetModificationTimeFileInfo(mtime *glib.TimeVal)

	SetNameFileInfo(name string)

	SetSizeFileInfo(size int64)

	SetSortOrderFileInfo(sortOrder int32)

	SetSymbolicIconFileInfo(icon Icon)

	SetSymlinkTargetFileInfo(symlinkTarget string)

	UnsetAttributeMaskFileInfo()
}

// fileInfo implements the FileInfo class.
type fileInfo struct {
	gextras.Objector
}

// WrapFileInfo wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileInfo(obj *externglib.Object) FileInfo {
	return fileInfo{
		Objector: obj,
	}
}

func marshalFileInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileInfo(obj), nil
}

func NewFileInfo() FileInfo {
	var _cret *C.GFileInfo // in

	_cret = C.g_file_info_new()

	var _fileInfo FileInfo // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)

	return _fileInfo
}

func (i fileInfo) ClearStatusFileInfo() {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	C.g_file_info_clear_status(_arg0)
}

func (s fileInfo) CopyIntoFileInfo(destInfo FileInfo) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GFileInfo)(unsafe.Pointer(destInfo.Native()))

	C.g_file_info_copy_into(_arg0, _arg1)
}

func (o fileInfo) DupFileInfo() FileInfo {
	var _arg0 *C.GFileInfo // out
	var _cret *C.GFileInfo // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(o.Native()))

	_cret = C.g_file_info_dup(_arg0)

	var _fileInfo FileInfo // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)

	return _fileInfo
}

func (i fileInfo) AttributeAsString(attribute string) string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_as_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (i fileInfo) AttributeBoolean(attribute string) bool {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_boolean(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i fileInfo) AttributeByteString(attribute string) string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_byte_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i fileInfo) AttributeData(attribute string) (FileAttributeType, interface{}, FileAttributeStatus, bool) {
	var _arg0 *C.GFileInfo           // out
	var _arg1 *C.char                // out
	var _arg2 C.GFileAttributeType   // in
	var _arg3 C.gpointer             // in
	var _arg4 C.GFileAttributeStatus // in
	var _cret C.gboolean             // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_data(_arg0, _arg1, &_arg2, &_arg3, &_arg4)

	var _typ FileAttributeType      // out
	var _valuePp interface{}        // out
	var _status FileAttributeStatus // out
	var _ok bool                    // out

	_typ = FileAttributeType(_arg2)
	_valuePp = box.Get(uintptr(_arg3))
	_status = FileAttributeStatus(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _typ, _valuePp, _status, _ok
}

func (i fileInfo) AttributeInt32(attribute string) int32 {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret C.gint32     // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_int32(_arg0, _arg1)

	var _gint32 int32 // out

	_gint32 = int32(_cret)

	return _gint32
}

func (i fileInfo) AttributeInt64(attribute string) int64 {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret C.gint64     // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_int64(_arg0, _arg1)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

func (i fileInfo) AttributeObject(attribute string) gextras.Objector {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret *C.GObject   // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_object(_arg0, _arg1)

	var _object gextras.Objector // out

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gextras.Objector)

	return _object
}

func (i fileInfo) AttributeStatus(attribute string) FileAttributeStatus {
	var _arg0 *C.GFileInfo           // out
	var _arg1 *C.char                // out
	var _cret C.GFileAttributeStatus // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_status(_arg0, _arg1)

	var _fileAttributeStatus FileAttributeStatus // out

	_fileAttributeStatus = FileAttributeStatus(_cret)

	return _fileAttributeStatus
}

func (i fileInfo) AttributeString(attribute string) string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i fileInfo) AttributeStringv(attribute string) []string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret **C.char

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_stringv(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

func (i fileInfo) AttributeType(attribute string) FileAttributeType {
	var _arg0 *C.GFileInfo         // out
	var _arg1 *C.char              // out
	var _cret C.GFileAttributeType // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_type(_arg0, _arg1)

	var _fileAttributeType FileAttributeType // out

	_fileAttributeType = FileAttributeType(_cret)

	return _fileAttributeType
}

func (i fileInfo) AttributeUint32(attribute string) uint32 {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret C.guint32    // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_uint32(_arg0, _arg1)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (i fileInfo) AttributeUint64(attribute string) uint64 {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret C.guint64    // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_get_attribute_uint64(_arg0, _arg1)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

func (i fileInfo) ContentType() string {
	var _arg0 *C.GFileInfo // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_content_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i fileInfo) DisplayName() string {
	var _arg0 *C.GFileInfo // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_display_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i fileInfo) EditName() string {
	var _arg0 *C.GFileInfo // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_edit_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i fileInfo) Etag() string {
	var _arg0 *C.GFileInfo // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_etag(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i fileInfo) FileType() FileType {
	var _arg0 *C.GFileInfo // out
	var _cret C.GFileType  // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_file_type(_arg0)

	var _fileType FileType // out

	_fileType = FileType(_cret)

	return _fileType
}

func (i fileInfo) Icon() Icon {
	var _arg0 *C.GFileInfo // out
	var _cret *C.GIcon     // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

func (i fileInfo) IsBackup() bool {
	var _arg0 *C.GFileInfo // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_is_backup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i fileInfo) IsHidden() bool {
	var _arg0 *C.GFileInfo // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_is_hidden(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i fileInfo) IsSymlink() bool {
	var _arg0 *C.GFileInfo // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_is_symlink(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i fileInfo) ModificationTime() glib.TimeVal {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.GTimeVal   // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	C.g_file_info_get_modification_time(_arg0, &_arg1)

	var _result glib.TimeVal // out

	{
		var refTmpIn *C.GTimeVal
		var refTmpOut *glib.TimeVal

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*glib.TimeVal)(unsafe.Pointer(refTmpIn))

		_result = *refTmpOut
	}

	return _result
}

func (i fileInfo) Name() string {
	var _arg0 *C.GFileInfo // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_name(_arg0)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

func (i fileInfo) Size() int64 {
	var _arg0 *C.GFileInfo // out
	var _cret C.goffset    // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_size(_arg0)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

func (i fileInfo) SortOrder() int32 {
	var _arg0 *C.GFileInfo // out
	var _cret C.gint32     // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_sort_order(_arg0)

	var _gint32 int32 // out

	_gint32 = int32(_cret)

	return _gint32
}

func (i fileInfo) SymbolicIcon() Icon {
	var _arg0 *C.GFileInfo // out
	var _cret *C.GIcon     // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_symbolic_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

func (i fileInfo) SymlinkTarget() string {
	var _arg0 *C.GFileInfo // out
	var _cret *C.char      // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_info_get_symlink_target(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (i fileInfo) HasAttributeFileInfo(attribute string) bool {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_has_attribute(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i fileInfo) HasNamespaceFileInfo(nameSpace string) bool {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(nameSpace))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_has_namespace(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i fileInfo) ListAttributesFileInfo(nameSpace string) []string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _cret **C.char

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(nameSpace))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_info_list_attributes(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (i fileInfo) RemoveAttributeFileInfo(attribute string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_remove_attribute(_arg0, _arg1)
}

func (i fileInfo) SetAttributeFileInfo(attribute string, typ FileAttributeType, valueP interface{}) {
	var _arg0 *C.GFileInfo         // out
	var _arg1 *C.char              // out
	var _arg2 C.GFileAttributeType // out
	var _arg3 C.gpointer           // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GFileAttributeType(typ)
	_arg3 = C.gpointer(box.Assign(unsafe.Pointer(valueP)))

	C.g_file_info_set_attribute(_arg0, _arg1, _arg2, _arg3)
}

func (i fileInfo) SetAttributeBooleanFileInfo(attribute string, attrValue bool) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.gboolean   // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	if attrValue {
		_arg2 = C.TRUE
	}

	C.g_file_info_set_attribute_boolean(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetAttributeByteStringFileInfo(attribute string, attrValue string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(attrValue))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_file_info_set_attribute_byte_string(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetAttributeInt32FileInfo(attribute string, attrValue int32) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.gint32     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint32(attrValue)

	C.g_file_info_set_attribute_int32(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetAttributeInt64FileInfo(attribute string, attrValue int64) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.gint64     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint64(attrValue)

	C.g_file_info_set_attribute_int64(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetAttributeMaskFileInfo(mask *FileAttributeMatcher) {
	var _arg0 *C.GFileInfo             // out
	var _arg1 *C.GFileAttributeMatcher // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(mask.Native()))

	C.g_file_info_set_attribute_mask(_arg0, _arg1)
}

func (i fileInfo) SetAttributeObjectFileInfo(attribute string, attrValue gextras.Objector) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 *C.GObject   // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GObject)(unsafe.Pointer(attrValue.Native()))

	C.g_file_info_set_attribute_object(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetAttributeStatusFileInfo(attribute string, status FileAttributeStatus) bool {
	var _arg0 *C.GFileInfo           // out
	var _arg1 *C.char                // out
	var _arg2 C.GFileAttributeStatus // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GFileAttributeStatus(status)

	_cret = C.g_file_info_set_attribute_status(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i fileInfo) SetAttributeStringFileInfo(attribute string, attrValue string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(attrValue))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_file_info_set_attribute_string(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetAttributeStringvFileInfo(attribute string, attrValue []string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 **C.char

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.char)(C.malloc(C.ulong(len(attrValue)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(attrValue))
		for i := range attrValue {
			out[i] = (*C.char)(C.CString(attrValue[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.g_file_info_set_attribute_stringv(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetAttributeUint32FileInfo(attribute string, attrValue uint32) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.guint32    // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint32(attrValue)

	C.g_file_info_set_attribute_uint32(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetAttributeUint64FileInfo(attribute string, attrValue uint64) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.guint64    // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint64(attrValue)

	C.g_file_info_set_attribute_uint64(_arg0, _arg1, _arg2)
}

func (i fileInfo) SetContentTypeFileInfo(contentType string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_content_type(_arg0, _arg1)
}

func (i fileInfo) SetDisplayNameFileInfo(displayName string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(displayName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_display_name(_arg0, _arg1)
}

func (i fileInfo) SetEditNameFileInfo(editName string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(editName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_edit_name(_arg0, _arg1)
}

func (i fileInfo) SetFileTypeFileInfo(typ FileType) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.GFileType  // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = C.GFileType(typ)

	C.g_file_info_set_file_type(_arg0, _arg1)
}

func (i fileInfo) SetIconFileInfo(icon Icon) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.GIcon     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.g_file_info_set_icon(_arg0, _arg1)
}

func (i fileInfo) SetIsHiddenFileInfo(isHidden bool) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	if isHidden {
		_arg1 = C.TRUE
	}

	C.g_file_info_set_is_hidden(_arg0, _arg1)
}

func (i fileInfo) SetIsSymlinkFileInfo(isSymlink bool) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	if isSymlink {
		_arg1 = C.TRUE
	}

	C.g_file_info_set_is_symlink(_arg0, _arg1)
}

func (i fileInfo) SetModificationTimeFileInfo(mtime *glib.TimeVal) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.GTimeVal  // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GTimeVal)(unsafe.Pointer(mtime.Native()))

	C.g_file_info_set_modification_time(_arg0, _arg1)
}

func (i fileInfo) SetNameFileInfo(name string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_name(_arg0, _arg1)
}

func (i fileInfo) SetSizeFileInfo(size int64) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.goffset    // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = C.goffset(size)

	C.g_file_info_set_size(_arg0, _arg1)
}

func (i fileInfo) SetSortOrderFileInfo(sortOrder int32) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.gint32     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = C.gint32(sortOrder)

	C.g_file_info_set_sort_order(_arg0, _arg1)
}

func (i fileInfo) SetSymbolicIconFileInfo(icon Icon) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.GIcon     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.g_file_info_set_symbolic_icon(_arg0, _arg1)
}

func (i fileInfo) SetSymlinkTargetFileInfo(symlinkTarget string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(symlinkTarget))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_symlink_target(_arg0, _arg1)
}

func (i fileInfo) UnsetAttributeMaskFileInfo() {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	C.g_file_info_unset_attribute_mask(_arg0)
}
