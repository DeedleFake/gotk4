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
		{T: externglib.Type(C.gtk_media_controls_get_type()), F: marshalMediaControlser},
	})
}

// MediaControls: GtkMediaControls is a widget to show controls for a video.
//
// !An example GtkMediaControls (media-controls.png)
//
// Usually, GtkMediaControls is used as part of gtk.Video.
type MediaControls struct {
	Widget
}

var _ gextras.Nativer = (*MediaControls)(nil)

func wrapMediaControls(obj *externglib.Object) *MediaControls {
	return &MediaControls{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalMediaControlser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMediaControls(obj), nil
}

// NewMediaControls creates a new GtkMediaControls managing the stream passed to
// it.
func NewMediaControls(stream MediaStreamer) *MediaControls {
	var _arg1 *C.GtkMediaStream // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer((stream).(gextras.Nativer).Native()))

	_cret = C.gtk_media_controls_new(_arg1)

	var _mediaControls *MediaControls // out

	_mediaControls = wrapMediaControls(externglib.Take(unsafe.Pointer(_cret)))

	return _mediaControls
}

// MediaStream gets the media stream managed by controls or NULL if none.
func (controls *MediaControls) MediaStream() *MediaStream {
	var _arg0 *C.GtkMediaControls // out
	var _cret *C.GtkMediaStream   // in

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(controls.Native()))

	_cret = C.gtk_media_controls_get_media_stream(_arg0)

	var _mediaStream *MediaStream // out

	_mediaStream = wrapMediaStream(externglib.Take(unsafe.Pointer(_cret)))

	return _mediaStream
}

// SetMediaStream sets the stream that is controlled by controls.
func (controls *MediaControls) SetMediaStream(stream MediaStreamer) {
	var _arg0 *C.GtkMediaControls // out
	var _arg1 *C.GtkMediaStream   // out

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(controls.Native()))
	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer((stream).(gextras.Nativer).Native()))

	C.gtk_media_controls_set_media_stream(_arg0, _arg1)
}
