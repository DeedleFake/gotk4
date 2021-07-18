// Code generated by girgen. DO NOT EDIT.

package gdkpixdata

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixdata.h>
// #include <glib-object.h>
import "C"

// PixdataDumpType: enumeration which is used by gdk_pixdata_to_csource() to
// determine the form of C source to be generated. The three values
// GDK_PIXDATA_DUMP_PIXDATA_STREAM, GDK_PIXDATA_DUMP_PIXDATA_STRUCT and
// GDK_PIXDATA_DUMP_MACROS are mutually exclusive, as are GDK_PIXBUF_DUMP_GTYPES
// and GDK_PIXBUF_DUMP_CTYPES. The remaining elements are optional flags that
// can be freely added.
//
// Deprecated: since version 2.32.
type PixdataDumpType int

const (
	// PixdataDumpPixdataStream: generate pixbuf data stream (a single string
	// containing a serialized Pixdata structure in network byte order).
	PixdataDumpPixdataStream PixdataDumpType = 0b0
	// PixdataDumpPixdataStruct: generate Pixdata structure (needs the Pixdata
	// structure definition from gdk-pixdata.h).
	PixdataDumpPixdataStruct PixdataDumpType = 0b1
	// PixdataDumpMacros: generate <function>*_ROWSTRIDE</function>,
	// <function>*_WIDTH</function>, <function>*_HEIGHT</function>,
	// <function>*_BYTES_PER_PIXEL</function> and
	// <function>*_RLE_PIXEL_DATA</function> or
	// <function>*_PIXEL_DATA</function> macro definitions for the image.
	PixdataDumpMacros PixdataDumpType = 0b10
	// PixdataDumpGTypes: generate GLib data types instead of standard C data
	// types.
	PixdataDumpGTypes PixdataDumpType = 0b0
	// PixdataDumpCtypes: generate standard C data types instead of GLib data
	// types.
	PixdataDumpCtypes PixdataDumpType = 0b100000000
	// PixdataDumpStatic: generate static symbols.
	PixdataDumpStatic PixdataDumpType = 0b1000000000
	// PixdataDumpConst: generate const symbols.
	PixdataDumpConst PixdataDumpType = 0b10000000000
	// PixdataDumpRleDecoder: provide a <function>*_RUN_LENGTH_DECODE(image_buf,
	// rle_data, size, bpp)</function> macro definition to decode run-length
	// encoded image data.
	PixdataDumpRleDecoder PixdataDumpType = 0b10000000000000000
)

// String returns the names in string for PixdataDumpType.
func (p PixdataDumpType) String() string {
	if p == 0 {
		return "PixdataDumpType(0)"
	}

	var builder strings.Builder
	builder.Grow(160)

	for p != 0 {
		next := p & (p - 1)
		bit := p - next

		switch bit {
		case PixdataDumpPixdataStream:
			builder.WriteString("PixdataStream|")
		case PixdataDumpPixdataStruct:
			builder.WriteString("PixdataStruct|")
		case PixdataDumpMacros:
			builder.WriteString("Macros|")
		case PixdataDumpCtypes:
			builder.WriteString("Ctypes|")
		case PixdataDumpStatic:
			builder.WriteString("Static|")
		case PixdataDumpConst:
			builder.WriteString("Const|")
		case PixdataDumpRleDecoder:
			builder.WriteString("RleDecoder|")
		default:
			builder.WriteString(fmt.Sprintf("PixdataDumpType(0b%b)|", bit))
		}

		p = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// PixdataType: enumeration containing three sets of flags for a Pixdata struct:
// one for the used colorspace, one for the width of the samples and one for the
// encoding of the pixel data.
//
// Deprecated: since version 2.32.
type PixdataType int

const (
	// PixdataColorTypeRGB: each pixel has red, green and blue samples.
	PixdataColorTypeRGB PixdataType = 0b1
	// PixdataColorTypeRGBA: each pixel has red, green and blue samples and an
	// alpha value.
	PixdataColorTypeRGBA PixdataType = 0b10
	// PixdataColorTypeMask: mask for the colortype flags of the enum.
	PixdataColorTypeMask PixdataType = 0b11111111
	// PixdataSampleWidth8: each sample has 8 bits.
	PixdataSampleWidth8 PixdataType = 0b10000000000000000
	// PixdataSampleWidthMask: mask for the sample width flags of the enum.
	PixdataSampleWidthMask PixdataType = 0b11110000000000000000
	// PixdataEncodingRaw: pixel data is in raw form.
	PixdataEncodingRaw PixdataType = 0b1000000000000000000000000
	// PixdataEncodingRle: pixel data is run-length encoded. Runs may be up to
	// 127 bytes long; their length is stored in a single byte preceding the
	// pixel data for the run. If a run is constant, its length byte has the
	// high bit set and the pixel data consists of a single pixel which must be
	// repeated.
	PixdataEncodingRle PixdataType = 0b10000000000000000000000000
	// PixdataEncodingMask: mask for the encoding flags of the enum.
	PixdataEncodingMask PixdataType = 0b1111000000000000000000000000
)

// String returns the names in string for PixdataType.
func (p PixdataType) String() string {
	if p == 0 {
		return "PixdataType(0)"
	}

	var builder strings.Builder
	builder.Grow(162)

	for p != 0 {
		next := p & (p - 1)
		bit := p - next

		switch bit {
		case PixdataColorTypeRGB:
			builder.WriteString("ColorTypeRGB|")
		case PixdataColorTypeRGBA:
			builder.WriteString("ColorTypeRGBA|")
		case PixdataColorTypeMask:
			builder.WriteString("ColorTypeMask|")
		case PixdataSampleWidth8:
			builder.WriteString("SampleWidth8|")
		case PixdataSampleWidthMask:
			builder.WriteString("SampleWidthMask|")
		case PixdataEncodingRaw:
			builder.WriteString("EncodingRaw|")
		case PixdataEncodingRle:
			builder.WriteString("EncodingRle|")
		case PixdataEncodingMask:
			builder.WriteString("EncodingMask|")
		default:
			builder.WriteString(fmt.Sprintf("PixdataType(0b%b)|", bit))
		}

		p = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// PixbufFromPixdata converts a GdkPixdata to a GdkPixbuf.
//
// If copy_pixels is TRUE or if the pixel data is run-length-encoded, the pixel
// data is copied into newly-allocated memory; otherwise it is reused.
//
// Deprecated: Use GResource instead.
func PixbufFromPixdata(pixdata *Pixdata, copyPixels bool) (*gdkpixbuf.Pixbuf, error) {
	var _arg1 *C.GdkPixdata // out
	var _arg2 C.gboolean    // out
	var _cret *C.GdkPixbuf  // in
	var _cerr *C.GError     // in

	_arg1 = (*C.GdkPixdata)(gextras.StructNative(unsafe.Pointer(pixdata)))
	if copyPixels {
		_arg2 = C.TRUE
	}

	_cret = C.gdk_pixbuf_from_pixdata(_arg1, _arg2, &_cerr)

	var _pixbuf *gdkpixbuf.Pixbuf // out
	var _goerr error              // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

// Pixdata: pixel buffer suitable for serialization and streaming.
//
// Using GdkPixdata, images can be compiled into an application, making it
// unnecessary to refer to external image files at runtime.
//
// GdkPixbuf includes a utility named gdk-pixbuf-csource, which can be used to
// convert image files into GdkPixdata structures suitable for inclusion in C
// sources. To convert the GdkPixdata structures back into a GdkPixbuf, use
// gdk_pixbuf_from_pixdata().
//
// Deprecated: GdkPixdata should not be used any more. GResource should be used
// to save the original compressed images inside the program's binary.
type Pixdata struct {
	nocopy gextras.NoCopy
	native *C.GdkPixdata
}

// Deserialize deserializes (reconstruct) a Pixdata structure from a byte
// stream.
//
// The byte stream consists of a straightforward writeout of the GdkPixdata
// fields in network byte order, plus the pixel_data bytes the structure points
// to.
//
// The pixdata contents are reconstructed byte by byte and are checked for
// validity.
//
// This function may fail with GDK_PIXBUF_ERROR_CORRUPT_IMAGE or
// GDK_PIXBUF_ERROR_UNKNOWN_TYPE.
//
// Deprecated: Use GResource instead.
func (pixdata *Pixdata) Deserialize(stream []byte) error {
	var _arg0 *C.GdkPixdata // out
	var _arg2 *C.guint8
	var _arg1 C.guint
	var _cerr *C.GError // in

	_arg0 = (*C.GdkPixdata)(gextras.StructNative(unsafe.Pointer(pixdata)))
	_arg1 = (C.guint)(len(stream))
	if len(stream) > 0 {
		_arg2 = (*C.guint8)(unsafe.Pointer(&stream[0]))
	}

	C.gdk_pixdata_deserialize(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
