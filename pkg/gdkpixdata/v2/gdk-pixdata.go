// Code generated by girgen. DO NOT EDIT.

package gdkpixdata

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk-pixbuf/gdk-pixdata.h>
import "C"

// PixdataDumpType: an enumeration which is used by gdk_pixdata_to_csource() to
// determine the form of C source to be generated. The three values
// @GDK_PIXDATA_DUMP_PIXDATA_STREAM, @GDK_PIXDATA_DUMP_PIXDATA_STRUCT and
// @GDK_PIXDATA_DUMP_MACROS are mutually exclusive, as are
// @GDK_PIXBUF_DUMP_GTYPES and @GDK_PIXBUF_DUMP_CTYPES. The remaining elements
// are optional flags that can be freely added.
type PixdataDumpType int

const (
	// PixdataDumpTypePixdataStream: generate pixbuf data stream (a single
	// string containing a serialized Pixdata structure in network byte order).
	PixdataDumpTypePixdataStream PixdataDumpType = 0
	// PixdataDumpTypePixdataStruct: generate Pixdata structure (needs the
	// Pixdata structure definition from gdk-pixdata.h).
	PixdataDumpTypePixdataStruct PixdataDumpType = 1
	// PixdataDumpTypeMacros: generate <function>*_ROWSTRIDE</function>,
	// <function>*_WIDTH</function>, <function>*_HEIGHT</function>,
	// <function>*_BYTES_PER_PIXEL</function> and
	// <function>*_RLE_PIXEL_DATA</function> or
	// <function>*_PIXEL_DATA</function> macro definitions for the image.
	PixdataDumpTypeMacros PixdataDumpType = 2
	// PixdataDumpTypeGTypes: generate GLib data types instead of standard C
	// data types.
	PixdataDumpTypeGTypes PixdataDumpType = 0
	// PixdataDumpTypeCtypes: generate standard C data types instead of GLib
	// data types.
	PixdataDumpTypeCtypes PixdataDumpType = 256
	// PixdataDumpTypeStatic: generate static symbols.
	PixdataDumpTypeStatic PixdataDumpType = 512
	// PixdataDumpTypeConst: generate const symbols.
	PixdataDumpTypeConst PixdataDumpType = 1024
	// PixdataDumpTypeRleDecoder: provide a
	// <function>*_RUN_LENGTH_DECODE(image_buf, rle_data, size, bpp)</function>
	// macro definition to decode run-length encoded image data.
	PixdataDumpTypeRleDecoder PixdataDumpType = 65536
)

// PixdataType: an enumeration containing three sets of flags for a Pixdata
// struct: one for the used colorspace, one for the width of the samples and one
// for the encoding of the pixel data.
type PixdataType int

const (
	// PixdataTypeColorTypeRGB: each pixel has red, green and blue samples.
	PixdataTypeColorTypeRGB PixdataType = 1
	// PixdataTypeColorTypeRGBA: each pixel has red, green and blue samples and
	// an alpha value.
	PixdataTypeColorTypeRGBA PixdataType = 2
	// PixdataTypeColorTypeMask: mask for the colortype flags of the enum.
	PixdataTypeColorTypeMask PixdataType = 255
	// PixdataTypeSampleWidth8: each sample has 8 bits.
	PixdataTypeSampleWidth8 PixdataType = 65536
	// PixdataTypeSampleWidthMask: mask for the sample width flags of the enum.
	PixdataTypeSampleWidthMask PixdataType = 983040
	// PixdataTypeEncodingRaw: the pixel data is in raw form.
	PixdataTypeEncodingRaw PixdataType = 16777216
	// PixdataTypeEncodingRle: the pixel data is run-length encoded. Runs may be
	// up to 127 bytes long; their length is stored in a single byte preceding
	// the pixel data for the run. If a run is constant, its length byte has the
	// high bit set and the pixel data consists of a single pixel which must be
	// repeated.
	PixdataTypeEncodingRle PixdataType = 33554432
	// PixdataTypeEncodingMask: mask for the encoding flags of the enum.
	PixdataTypeEncodingMask PixdataType = 251658240
)

// PixbufFromPixdata converts a `GdkPixdata` to a `GdkPixbuf`.
//
// If `copy_pixels` is `TRUE` or if the pixel data is run-length-encoded, the
// pixel data is copied into newly-allocated memory; otherwise it is reused.
func PixbufFromPixdata(pixdata *Pixdata, copyPixels bool) (gdkpixbuf.Pixbuf, error) {
	var _arg1 *C.GdkPixdata
	var _arg2 C.gboolean

	_arg1 = (*C.GdkPixdata)(unsafe.Pointer(pixdata.Native()))
	if copyPixels {
		_arg2 = C.gboolean(1)
	}

	var _cret *C.GdkPixbuf
	var _cerr *C.GError

	cret = C.gdk_pixbuf_from_pixdata(_arg1, _arg2, _cerr)

	var _pixbuf gdkpixbuf.Pixbuf
	var _goerr error

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(gdkpixbuf.Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

// Pixdata: a pixel buffer suitable for serialization and streaming.
//
// Using `GdkPixdata`, images can be compiled into an application, making it
// unnecessary to refer to external image files at runtime.
//
// `GdkPixbuf` includes a utility named `gdk-pixbuf-csource`, which can be used
// to convert image files into `GdkPixdata` structures suitable for inclusion in
// C sources. To convert the `GdkPixdata` structures back into a `GdkPixbuf`,
// use `gdk_pixbuf_from_pixdata()`.
type Pixdata struct {
	native C.GdkPixdata
}

// WrapPixdata wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPixdata(ptr unsafe.Pointer) *Pixdata {
	if ptr == nil {
		return nil
	}

	return (*Pixdata)(ptr)
}

func marshalPixdata(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPixdata(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *Pixdata) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Magic gets the field inside the struct.
func (p *Pixdata) Magic() uint32 {
	var v uint32
	v = (uint32)(p.native.magic)
	return v
}

// Length gets the field inside the struct.
func (p *Pixdata) Length() int32 {
	var v int32
	v = (int32)(p.native.length)
	return v
}

// PixdataType gets the field inside the struct.
func (p *Pixdata) PixdataType() uint32 {
	var v uint32
	v = (uint32)(p.native.pixdata_type)
	return v
}

// Rowstride gets the field inside the struct.
func (p *Pixdata) Rowstride() uint32 {
	var v uint32
	v = (uint32)(p.native.rowstride)
	return v
}

// Width gets the field inside the struct.
func (p *Pixdata) Width() uint32 {
	var v uint32
	v = (uint32)(p.native.width)
	return v
}

// Height gets the field inside the struct.
func (p *Pixdata) Height() uint32 {
	var v uint32
	v = (uint32)(p.native.height)
	return v
}

// Serialize serializes a Pixdata structure into a byte stream. The byte stream
// consists of a straightforward writeout of the Pixdata fields in network byte
// order, plus the @pixel_data bytes the structure points to.
func (p *Pixdata) Serialize() []byte {
	var _arg0 *C.GdkPixdata

	_arg0 = (*C.GdkPixdata)(unsafe.Pointer(p.Native()))

	var _cret *C.guint8
	var _arg1 *C.guint

	cret = C.gdk_pixdata_serialize(_arg0)

	var _guint8s []byte

	ptr.SetSlice(unsafe.Pointer(&_guint8s), unsafe.Pointer(_cret), int(_arg1))
	runtime.SetFinalizer(&_guint8s, func(v *[]byte) {
		C.free(ptr.Slice(unsafe.Pointer(v)))
	})

	return _guint8s
}

// ToCsource generates C source code suitable for compiling images directly into
// programs.
//
// GdkPixbuf ships with a program called `gdk-pixbuf-csource`, which offers a
// command line interface to this function.
func (p *Pixdata) ToCsource(name string, dumpType PixdataDumpType) *glib.String {
	var _arg0 *C.GdkPixdata
	var _arg1 *C.gchar
	var _arg2 C.GdkPixdataDumpType

	_arg0 = (*C.GdkPixdata)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GdkPixdataDumpType)(dumpType)

	var _cret *C.GString

	cret = C.gdk_pixdata_to_csource(_arg0, _arg1, _arg2)

	var _string *glib.String

	_string = glib.WrapString(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_string, func(v *glib.String) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _string
}
