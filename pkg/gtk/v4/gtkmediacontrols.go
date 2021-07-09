// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_media_controls_get_type()), F: marshalMediaControls},
	})
}

// MediaControls: `GtkMediaControls` is a widget to show controls for a video.
//
// !An example GtkMediaControls (media-controls.png)
//
// Usually, `GtkMediaControls` is used as part of [class@Gtk.Video].
type MediaControls interface {
	gextras.Objector

	// MediaStream gets the media stream managed by @controls or nil if none.
	MediaStream() *MediaStreamClass
	// SetMediaStream sets the stream that is controlled by @controls.
	SetMediaStream(stream MediaStream)
}

// MediaControlsClass implements the MediaControls interface.
type MediaControlsClass struct {
	*externglib.Object
	WidgetClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
}

var _ MediaControls = (*MediaControlsClass)(nil)

func wrapMediaControls(obj *externglib.Object) MediaControls {
	return &MediaControlsClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
	}
}

func marshalMediaControls(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMediaControls(obj), nil
}

// NewMediaControls creates a new `GtkMediaControls` managing the @stream passed
// to it.
func NewMediaControls(stream MediaStream) *MediaControlsClass {
	var _arg1 *C.GtkMediaStream // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer((&MediaStream).Native()))

	_cret = C.gtk_media_controls_new(_arg1)

	var _mediaControls *MediaControlsClass // out

	_mediaControls = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*MediaControlsClass)

	return _mediaControls
}

// MediaStream gets the media stream managed by @controls or nil if none.
func (c *MediaControlsClass) MediaStream() *MediaStreamClass {
	var _arg0 *C.GtkMediaControls // out
	var _cret *C.GtkMediaStream   // in

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer((&MediaControls).Native()))

	_cret = C.gtk_media_controls_get_media_stream(_arg0)

	var _mediaStream *MediaStreamClass // out

	_mediaStream = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*MediaStreamClass)

	return _mediaStream
}

// SetMediaStream sets the stream that is controlled by @controls.
func (c *MediaControlsClass) SetMediaStream(stream MediaStream) {
	var _arg0 *C.GtkMediaControls // out
	var _arg1 *C.GtkMediaStream   // out

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer((&MediaControls).Native()))
	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer((&MediaStream).Native()))

	C.gtk_media_controls_set_media_stream(_arg0, _arg1)
}
