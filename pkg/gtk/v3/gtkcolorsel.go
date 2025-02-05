// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_ColorSelection_ConnectColorChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_ColorSelectionClass_color_changed(GtkColorSelection*);
// void _gotk4_gtk3_ColorSelection_virtual_color_changed(void* fnptr, GtkColorSelection* arg0) {
//   ((void (*)(GtkColorSelection*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeColorSelection = coreglib.Type(C.gtk_color_selection_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorSelection, F: marshalColorSelection},
	})
}

// ColorSelectionOverrides contains methods that are overridable.
type ColorSelectionOverrides struct {
	ColorChanged func()
}

func defaultColorSelectionOverrides(v *ColorSelection) ColorSelectionOverrides {
	return ColorSelectionOverrides{
		ColorChanged: v.colorChanged,
	}
}

type ColorSelection struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*ColorSelection)(nil)
	_ coreglib.Objector = (*ColorSelection)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ColorSelection, *ColorSelectionClass, ColorSelectionOverrides](
		GTypeColorSelection,
		initColorSelectionClass,
		wrapColorSelection,
		defaultColorSelectionOverrides,
	)
}

func initColorSelectionClass(gclass unsafe.Pointer, overrides ColorSelectionOverrides, classInitFunc func(*ColorSelectionClass)) {
	pclass := (*C.GtkColorSelectionClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeColorSelection))))

	if overrides.ColorChanged != nil {
		pclass.color_changed = (*[0]byte)(C._gotk4_gtk3_ColorSelectionClass_color_changed)
	}

	if classInitFunc != nil {
		class := (*ColorSelectionClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapColorSelection(obj *coreglib.Object) *ColorSelection {
	return &ColorSelection{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalColorSelection(p uintptr) (interface{}, error) {
	return wrapColorSelection(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectColorChanged: this signal is emitted when the color changes in the
// ColorSelection according to its update policy.
func (colorsel *ColorSelection) ConnectColorChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(colorsel, "color-changed", false, unsafe.Pointer(C._gotk4_gtk3_ColorSelection_ConnectColorChanged), f)
}

// NewColorSelection creates a new GtkColorSelection.
//
// The function returns the following values:
//
//    - colorSelection: new ColorSelection.
//
func NewColorSelection() *ColorSelection {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_selection_new()

	var _colorSelection *ColorSelection // out

	_colorSelection = wrapColorSelection(coreglib.Take(unsafe.Pointer(_cret)))

	return _colorSelection
}

// CurrentAlpha returns the current alpha value.
//
// The function returns the following values:
//
//    - guint16: integer between 0 and 65535.
//
func (colorsel *ColorSelection) CurrentAlpha() uint16 {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.guint16            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_cret = C.gtk_color_selection_get_current_alpha(_arg0)
	runtime.KeepAlive(colorsel)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// CurrentColor sets color to be the current color in the GtkColorSelection
// widget.
//
// Deprecated: Use gtk_color_selection_get_current_rgba() instead.
//
// The function returns the following values:
//
//    - color to fill in with the current color.
//
func (colorsel *ColorSelection) CurrentColor() *gdk.Color {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkColor           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	C.gtk_color_selection_get_current_color(_arg0, &_arg1)
	runtime.KeepAlive(colorsel)

	var _color *gdk.Color // out

	_color = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _color
}

// CurrentRGBA sets rgba to be the current color in the GtkColorSelection
// widget.
//
// The function returns the following values:
//
//    - rgba to fill in with the current color.
//
func (colorsel *ColorSelection) CurrentRGBA() *gdk.RGBA {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkRGBA            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	C.gtk_color_selection_get_current_rgba(_arg0, &_arg1)
	runtime.KeepAlive(colorsel)

	var _rgba *gdk.RGBA // out

	_rgba = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _rgba
}

// HasOpacityControl determines whether the colorsel has an opacity control.
//
// The function returns the following values:
//
//    - ok: TRUE if the colorsel has an opacity control, FALSE if it does't.
//
func (colorsel *ColorSelection) HasOpacityControl() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_cret = C.gtk_color_selection_get_has_opacity_control(_arg0)
	runtime.KeepAlive(colorsel)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasPalette determines whether the color selector has a color palette.
//
// The function returns the following values:
//
//    - ok: TRUE if the selector has a palette, FALSE if it hasn't.
//
func (colorsel *ColorSelection) HasPalette() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_cret = C.gtk_color_selection_get_has_palette(_arg0)
	runtime.KeepAlive(colorsel)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PreviousAlpha returns the previous alpha value.
//
// The function returns the following values:
//
//    - guint16: integer between 0 and 65535.
//
func (colorsel *ColorSelection) PreviousAlpha() uint16 {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.guint16            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_cret = C.gtk_color_selection_get_previous_alpha(_arg0)
	runtime.KeepAlive(colorsel)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// PreviousColor fills color in with the original color value.
//
// Deprecated: Use gtk_color_selection_get_previous_rgba() instead.
//
// The function returns the following values:
//
//    - color to fill in with the original color value.
//
func (colorsel *ColorSelection) PreviousColor() *gdk.Color {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkColor           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	C.gtk_color_selection_get_previous_color(_arg0, &_arg1)
	runtime.KeepAlive(colorsel)

	var _color *gdk.Color // out

	_color = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _color
}

// PreviousRGBA fills rgba in with the original color value.
//
// The function returns the following values:
//
//    - rgba to fill in with the original color value.
//
func (colorsel *ColorSelection) PreviousRGBA() *gdk.RGBA {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkRGBA            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	C.gtk_color_selection_get_previous_rgba(_arg0, &_arg1)
	runtime.KeepAlive(colorsel)

	var _rgba *gdk.RGBA // out

	_rgba = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _rgba
}

// IsAdjusting gets the current state of the colorsel.
//
// The function returns the following values:
//
//    - ok: TRUE if the user is currently dragging a color around, and FALSE if
//      the selection has stopped.
//
func (colorsel *ColorSelection) IsAdjusting() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))

	_cret = C.gtk_color_selection_is_adjusting(_arg0)
	runtime.KeepAlive(colorsel)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCurrentAlpha sets the current opacity to be alpha.
//
// The first time this is called, it will also set the original opacity to be
// alpha too.
//
// The function takes the following parameters:
//
//    - alpha: integer between 0 and 65535.
//
func (colorsel *ColorSelection) SetCurrentAlpha(alpha uint16) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.guint16            // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_selection_set_current_alpha(_arg0, _arg1)
	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(alpha)
}

// SetCurrentColor sets the current color to be color.
//
// The first time this is called, it will also set the original color to be
// color too.
//
// Deprecated: Use gtk_color_selection_set_current_rgba() instead.
//
// The function takes the following parameters:
//
//    - color to set the current color with.
//
func (colorsel *ColorSelection) SetCurrentColor(color *gdk.Color) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkColor          // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	C.gtk_color_selection_set_current_color(_arg0, _arg1)
	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(color)
}

// SetCurrentRGBA sets the current color to be rgba.
//
// The first time this is called, it will also set the original color to be rgba
// too.
//
// The function takes the following parameters:
//
//    - rgba to set the current color with.
//
func (colorsel *ColorSelection) SetCurrentRGBA(rgba *gdk.RGBA) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkRGBA           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	C.gtk_color_selection_set_current_rgba(_arg0, _arg1)
	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(rgba)
}

// SetHasOpacityControl sets the colorsel to use or not use opacity.
//
// The function takes the following parameters:
//
//    - hasOpacity: TRUE if colorsel can set the opacity, FALSE otherwise.
//
func (colorsel *ColorSelection) SetHasOpacityControl(hasOpacity bool) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	if hasOpacity {
		_arg1 = C.TRUE
	}

	C.gtk_color_selection_set_has_opacity_control(_arg0, _arg1)
	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(hasOpacity)
}

// SetHasPalette shows and hides the palette based upon the value of
// has_palette.
//
// The function takes the following parameters:
//
//    - hasPalette: TRUE if palette is to be visible, FALSE otherwise.
//
func (colorsel *ColorSelection) SetHasPalette(hasPalette bool) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	if hasPalette {
		_arg1 = C.TRUE
	}

	C.gtk_color_selection_set_has_palette(_arg0, _arg1)
	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(hasPalette)
}

// SetPreviousAlpha sets the “previous” alpha to be alpha.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that alpha change.
//
// The function takes the following parameters:
//
//    - alpha: integer between 0 and 65535.
//
func (colorsel *ColorSelection) SetPreviousAlpha(alpha uint16) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.guint16            // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_selection_set_previous_alpha(_arg0, _arg1)
	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(alpha)
}

// SetPreviousColor sets the “previous” color to be color.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that color change. Calling
// gtk_color_selection_set_current_color() will also set this color the first
// time it is called.
//
// Deprecated: Use gtk_color_selection_set_previous_rgba() instead.
//
// The function takes the following parameters:
//
//    - color to set the previous color with.
//
func (colorsel *ColorSelection) SetPreviousColor(color *gdk.Color) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkColor          // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	_arg1 = (*C.GdkColor)(gextras.StructNative(unsafe.Pointer(color)))

	C.gtk_color_selection_set_previous_color(_arg0, _arg1)
	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(color)
}

// SetPreviousRGBA sets the “previous” color to be rgba.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that color change. Calling
// gtk_color_selection_set_current_rgba() will also set this color the first
// time it is called.
//
// The function takes the following parameters:
//
//    - rgba to set the previous color with.
//
func (colorsel *ColorSelection) SetPreviousRGBA(rgba *gdk.RGBA) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkRGBA           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorsel).Native()))
	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	C.gtk_color_selection_set_previous_rgba(_arg0, _arg1)
	runtime.KeepAlive(colorsel)
	runtime.KeepAlive(rgba)
}

func (colorSelection *ColorSelection) colorChanged() {
	gclass := (*C.GtkColorSelectionClass)(coreglib.PeekParentClass(colorSelection))
	fnarg := gclass.color_changed

	var _arg0 *C.GtkColorSelection // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(coreglib.InternObject(colorSelection).Native()))

	C._gotk4_gtk3_ColorSelection_virtual_color_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(colorSelection)
}

// ColorSelectionPaletteFromString parses a color palette string; the string is
// a colon-separated list of color names readable by gdk_color_parse().
//
// The function takes the following parameters:
//
//    - str: string encoding a color palette.
//
// The function returns the following values:
//
//    - colors: return location for allocated array of Color.
//    - ok: TRUE if a palette was successfully parsed.
//
func ColorSelectionPaletteFromString(str string) ([]gdk.Color, bool) {
	var _arg1 *C.gchar    // out
	var _arg2 *C.GdkColor // in
	var _arg3 C.gint      // in
	var _cret C.gboolean  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_color_selection_palette_from_string(_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(str)

	var _colors []gdk.Color // out
	var _ok bool            // out

	defer C.free(unsafe.Pointer(_arg2))
	{
		src := unsafe.Slice((*C.GdkColor)(_arg2), _arg3)
		_colors = make([]gdk.Color, _arg3)
		for i := 0; i < int(_arg3); i++ {
			_colors[i] = *(*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(&_colors[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.gdk_color_free((*C.GdkColor)(intern.C))
				},
			)
		}
	}
	if _cret != 0 {
		_ok = true
	}

	return _colors, _ok
}

// ColorSelectionPaletteToString encodes a palette as a string, useful for
// persistent storage.
//
// The function takes the following parameters:
//
//    - colors: array of colors.
//
// The function returns the following values:
//
//    - utf8: allocated string encoding the palette.
//
func ColorSelectionPaletteToString(colors []gdk.Color) string {
	var _arg1 *C.GdkColor // out
	var _arg2 C.gint
	var _cret *C.gchar // in

	_arg2 = (C.gint)(len(colors))
	_arg1 = (*C.GdkColor)(C.calloc(C.size_t(len(colors)), C.size_t(C.sizeof_GdkColor)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GdkColor)(_arg1), len(colors))
		for i := range colors {
			out[i] = *(*C.GdkColor)(gextras.StructNative(unsafe.Pointer((&colors[i]))))
		}
	}

	_cret = C.gtk_color_selection_palette_to_string(_arg1, _arg2)
	runtime.KeepAlive(colors)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ColorSelectionClass: instance of this type is always passed by reference.
type ColorSelectionClass struct {
	*colorSelectionClass
}

// colorSelectionClass is the struct that's finalized.
type colorSelectionClass struct {
	native *C.GtkColorSelectionClass
}

// ParentClass: parent class.
func (c *ColorSelectionClass) ParentClass() *BoxClass {
	valptr := &c.native.parent_class
	var _v *BoxClass // out
	_v = (*BoxClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
