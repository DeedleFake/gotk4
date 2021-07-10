// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_progress_bar_get_type()), F: marshalProgressBarrer},
	})
}

// ProgressBarrer describes ProgressBar's methods.
type ProgressBarrer interface {
	gextras.Objector

	Ellipsize() pango.EllipsizeMode
	Fraction() float64
	Inverted() bool
	PulseStep() float64
	ShowText() bool
	Text() string
	Pulse()
	SetFraction(fraction float64)
	SetInverted(inverted bool)
	SetPulseStep(fraction float64)
	SetShowText(showText bool)
	SetText(text string)
}

// ProgressBar: the ProgressBar is typically used to display the progress of a
// long running operation. It provides a visual clue that processing is
// underway. The GtkProgressBar can be used in two different modes: percentage
// mode and activity mode.
//
// When an application can determine how much work needs to take place (e.g.
// read a fixed number of bytes from a file) and can monitor its progress, it
// can use the GtkProgressBar in percentage mode and the user sees a growing bar
// indicating the percentage of the work that has been completed. In this mode,
// the application is required to call gtk_progress_bar_set_fraction()
// periodically to update the progress bar.
//
// When an application has no accurate way of knowing the amount of work to do,
// it can use the ProgressBar in activity mode, which shows activity by a block
// moving back and forth within the progress area. In this mode, the application
// is required to call gtk_progress_bar_pulse() periodically to update the
// progress bar.
//
// There is quite a bit of flexibility provided to control the appearance of the
// ProgressBar. Functions are provided to control the orientation of the bar,
// optional text can be displayed along with the bar, and the step size used in
// activity mode can be set.
//
// CSS nodes
//
//    progressbar[.osd]
//    ├── [text]
//    ╰── trough[.empty][.full]
//        ╰── progress[.pulse]
//
// GtkProgressBar has a main CSS node with name progressbar and subnodes with
// names text and trough, of which the latter has a subnode named progress. The
// text subnode is only present if text is shown. The progress subnode has the
// style class .pulse when in activity mode. It gets the style classes .left,
// .right, .top or .bottom added when the progress 'touches' the corresponding
// end of the GtkProgressBar. The .osd class on the progressbar node is for use
// in overlays like the one Epiphany has for page loading progress.
type ProgressBar struct {
	*externglib.Object

	Widget
	atk.ImplementorIface
	Buildable
	Orientable
}

var _ ProgressBarrer = (*ProgressBar)(nil)

func wrapProgressBarrer(obj *externglib.Object) ProgressBarrer {
	return &ProgressBar{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
		ImplementorIface: atk.ImplementorIface{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalProgressBarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapProgressBarrer(obj), nil
}

// NewProgressBar creates a new ProgressBar.
func NewProgressBar() *ProgressBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_progress_bar_new()

	var _progressBar *ProgressBar // out

	_progressBar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ProgressBar)

	return _progressBar
}

// Ellipsize returns the ellipsizing position of the progress bar. See
// gtk_progress_bar_set_ellipsize().
func (pbar *ProgressBar) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkProgressBar    // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))

	_cret = C.gtk_progress_bar_get_ellipsize(_arg0)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = (pango.EllipsizeMode)(_cret)

	return _ellipsizeMode
}

// Fraction returns the current fraction of the task that’s been completed.
func (pbar *ProgressBar) Fraction() float64 {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))

	_cret = C.gtk_progress_bar_get_fraction(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Inverted gets the value set by gtk_progress_bar_set_inverted().
func (pbar *ProgressBar) Inverted() bool {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))

	_cret = C.gtk_progress_bar_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PulseStep retrieves the pulse step set with
// gtk_progress_bar_set_pulse_step().
func (pbar *ProgressBar) PulseStep() float64 {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))

	_cret = C.gtk_progress_bar_get_pulse_step(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ShowText gets the value of the ProgressBar:show-text property. See
// gtk_progress_bar_set_show_text().
func (pbar *ProgressBar) ShowText() bool {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))

	_cret = C.gtk_progress_bar_get_show_text(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Text retrieves the text that is displayed with the progress bar, if any,
// otherwise nil. The return value is a reference to the text, not a copy of it,
// so will become invalid if you change the text in the progress bar.
func (pbar *ProgressBar) Text() string {
	var _arg0 *C.GtkProgressBar // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))

	_cret = C.gtk_progress_bar_get_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Pulse indicates that some progress has been made, but you don’t know how
// much. Causes the progress bar to enter “activity mode,” where a block bounces
// back and forth. Each call to gtk_progress_bar_pulse() causes the block to
// move by a little bit (the amount of movement per pulse is determined by
// gtk_progress_bar_set_pulse_step()).
func (pbar *ProgressBar) Pulse() {
	var _arg0 *C.GtkProgressBar // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))

	C.gtk_progress_bar_pulse(_arg0)
}

// SetFraction causes the progress bar to “fill in” the given fraction of the
// bar. The fraction should be between 0.0 and 1.0, inclusive.
func (pbar *ProgressBar) SetFraction(fraction float64) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))
	_arg1 = C.gdouble(fraction)

	C.gtk_progress_bar_set_fraction(_arg0, _arg1)
}

// SetInverted progress bars normally grow from top to bottom or left to right.
// Inverted progress bars grow in the opposite direction.
func (pbar *ProgressBar) SetInverted(inverted bool) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.gtk_progress_bar_set_inverted(_arg0, _arg1)
}

// SetPulseStep sets the fraction of total progress bar length to move the
// bouncing block for each call to gtk_progress_bar_pulse().
func (pbar *ProgressBar) SetPulseStep(fraction float64) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))
	_arg1 = C.gdouble(fraction)

	C.gtk_progress_bar_set_pulse_step(_arg0, _arg1)
}

// SetShowText sets whether the progress bar will show text next to the bar. The
// shown text is either the value of the ProgressBar:text property or, if that
// is nil, the ProgressBar:fraction value, as a percentage.
//
// To make a progress bar that is styled and sized suitably for containing text
// (even if the actual text is blank), set ProgressBar:show-text to true and
// ProgressBar:text to the empty string (not nil).
func (pbar *ProgressBar) SetShowText(showText bool) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))
	if showText {
		_arg1 = C.TRUE
	}

	C.gtk_progress_bar_set_show_text(_arg0, _arg1)
}

// SetText causes the given @text to appear next to the progress bar.
//
// If @text is nil and ProgressBar:show-text is true, the current value of
// ProgressBar:fraction will be displayed as a percentage.
//
// If @text is non-nil and ProgressBar:show-text is true, the text will be
// displayed. In this case, it will not display the progress percentage. If
// @text is the empty string, the progress bar will still be styled and sized
// suitably for containing text, as long as ProgressBar:show-text is true.
func (pbar *ProgressBar) SetText(text string) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(pbar.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_progress_bar_set_text(_arg0, _arg1)
}
