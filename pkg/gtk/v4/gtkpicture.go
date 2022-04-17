// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// glib.Type values for gtkpicture.go.
var GTypePicture = externglib.Type(C.gtk_picture_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePicture, F: marshalPicture},
	})
}

// PictureOverrider contains methods that are overridable.
type PictureOverrider interface {
}

// Picture: GtkPicture widget displays a GdkPaintable.
//
// !An example GtkPicture (picture.png)
//
// Many convenience functions are provided to make pictures simple to use. For
// example, if you want to load an image from a file, and then display it,
// there’s a convenience function to do this:
//
//    GtkWidget *widget = gtk_picture_new_for_filename ("myfile.png");
//
//
// If the file isn’t loaded successfully, the picture will contain a “broken
// image” icon similar to that used in many web browsers. If you want to handle
// errors in loading the file yourself, for example by displaying an error
// message, then load the image with gdk.Texture.NewFromFile, then create the
// GtkPicture with gtk.Picture.NewForPaintable.
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of GResource for details. In this
// case, gtk.Picture.NewForResource and gtk.Picture.SetResource() should be
// used.
//
// GtkPicture displays an image at its natural size. See gtk.Image if you want
// to display a fixed-size image, such as an icon.
//
//
// Sizing the paintable
//
// You can influence how the paintable is displayed inside the GtkPicture. By
// turning off gtk.Picture:keep-aspect-ratio you can allow the paintable to get
// stretched. gtk.Picture:can-shrink can be unset to make sure that paintables
// are never made smaller than their ideal size - but be careful if you do not
// know the size of the paintable in use (like when displaying user-loaded
// images). This can easily cause the picture to grow larger than the screen.
// And gtkwidget:halign and gtkwidget:valign can be used to make sure the
// paintable doesn't fill all available space but is instead displayed at its
// original size.
//
//
// CSS nodes
//
// GtkPicture has a single CSS node with the name picture.
//
//
// Accessibility
//
// GtkPicture uses the GTK_ACCESSIBLE_ROLE_IMG role.
type Picture struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Picture)(nil)
)

func classInitPicturer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapPicture(obj *externglib.Object) *Picture {
	return &Picture{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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

func marshalPicture(p uintptr) (interface{}, error) {
	return wrapPicture(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPicture creates a new empty GtkPicture widget.
//
// The function returns the following values:
//
//    - picture: newly created GtkPicture widget.
//
func NewPicture() *Picture {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_picture_new()

	var _picture *Picture // out

	_picture = wrapPicture(externglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForFile creates a new GtkPicture displaying the given file.
//
// If the file isn’t found or can’t be loaded, the resulting GtkPicture is
// empty.
//
// If you need to detect failures to load the file, use gdk.Texture.NewFromFile
// to load the file yourself, then create the GtkPicture from the texture.
//
// The function takes the following parameters:
//
//    - file (optional): GFile.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForFile(file gio.Filer) *Picture {
	var _arg1 *C.GFile     // out
	var _cret *C.GtkWidget // in

	if file != nil {
		_arg1 = (*C.GFile)(unsafe.Pointer(externglib.InternObject(file).Native()))
	}

	_cret = C.gtk_picture_new_for_file(_arg1)
	runtime.KeepAlive(file)

	var _picture *Picture // out

	_picture = wrapPicture(externglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForFilename creates a new GtkPicture displaying the file filename.
//
// This is a utility function that calls gtk.Picture.NewForFile. See that
// function for details.
//
// The function takes the following parameters:
//
//    - filename (optional): filename.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForFilename(filename string) *Picture {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	if filename != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_picture_new_for_filename(_arg1)
	runtime.KeepAlive(filename)

	var _picture *Picture // out

	_picture = wrapPicture(externglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForPaintable creates a new GtkPicture displaying paintable.
//
// The GtkPicture will track changes to the paintable and update its size and
// contents in response to it.
//
// The function takes the following parameters:
//
//    - paintable (optional): GdkPaintable, or NULL.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForPaintable(paintable gdk.Paintabler) *Picture {
	var _arg1 *C.GdkPaintable // out
	var _cret *C.GtkWidget    // in

	if paintable != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(externglib.InternObject(paintable).Native()))
	}

	_cret = C.gtk_picture_new_for_paintable(_arg1)
	runtime.KeepAlive(paintable)

	var _picture *Picture // out

	_picture = wrapPicture(externglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForPixbuf creates a new GtkPicture displaying pixbuf.
//
// This is a utility function that calls gtk.Picture.NewForPaintable, See that
// function for details.
//
// The pixbuf must not be modified after passing it to this function.
//
// The function takes the following parameters:
//
//    - pixbuf (optional): GdkPixbuf, or NULL.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Picture {
	var _arg1 *C.GdkPixbuf // out
	var _cret *C.GtkWidget // in

	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(externglib.InternObject(pixbuf).Native()))
	}

	_cret = C.gtk_picture_new_for_pixbuf(_arg1)
	runtime.KeepAlive(pixbuf)

	var _picture *Picture // out

	_picture = wrapPicture(externglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// NewPictureForResource creates a new GtkPicture displaying the resource at
// resource_path.
//
// This is a utility function that calls gtk.Picture.NewForFile. See that
// function for details.
//
// The function takes the following parameters:
//
//    - resourcePath (optional): resource path to play back.
//
// The function returns the following values:
//
//    - picture: new GtkPicture.
//
func NewPictureForResource(resourcePath string) *Picture {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	if resourcePath != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_picture_new_for_resource(_arg1)
	runtime.KeepAlive(resourcePath)

	var _picture *Picture // out

	_picture = wrapPicture(externglib.Take(unsafe.Pointer(_cret)))

	return _picture
}

// AlternativeText gets the alternative textual description of the picture.
//
// The returned string will be NULL if the picture cannot be described
// textually.
//
// The function returns the following values:
//
//    - utf8 (optional): alternative textual description of self.
//
func (self *Picture) AlternativeText() string {
	var _arg0 *C.GtkPicture // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_picture_get_alternative_text(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CanShrink returns whether the GtkPicture respects its contents size.
//
// The function returns the following values:
//
//    - ok: TRUE if the picture can be made smaller than its contents.
//
func (self *Picture) CanShrink() bool {
	var _arg0 *C.GtkPicture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_picture_get_can_shrink(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// File gets the GFile currently displayed if self is displaying a file.
//
// If self is not displaying a file, for example when gtk.Picture.SetPaintable()
// was used, then NULL is returned.
//
// The function returns the following values:
//
//    - file (optional): GFile displayed by self.
//
func (self *Picture) File() gio.Filer {
	var _arg0 *C.GtkPicture // out
	var _cret *C.GFile      // in

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_picture_get_file(_arg0)
	runtime.KeepAlive(self)

	var _file gio.Filer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gio.Filer)
				return ok
			})
			rv, ok := casted.(gio.Filer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
			}
			_file = rv
		}
	}

	return _file
}

// KeepAspectRatio returns whether the GtkPicture preserves its contents aspect
// ratio.
//
// The function returns the following values:
//
//    - ok: TRUE if the self tries to keep the contents' aspect ratio.
//
func (self *Picture) KeepAspectRatio() bool {
	var _arg0 *C.GtkPicture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_picture_get_keep_aspect_ratio(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Paintable gets the GdkPaintable being displayed by the GtkPicture.
//
// The function returns the following values:
//
//    - paintable (optional): displayed paintable, or NULL if the picture is
//      empty.
//
func (self *Picture) Paintable() gdk.Paintabler {
	var _arg0 *C.GtkPicture   // out
	var _cret *C.GdkPaintable // in

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_picture_get_paintable(_arg0)
	runtime.KeepAlive(self)

	var _paintable gdk.Paintabler // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gdk.Paintabler)
				return ok
			})
			rv, ok := casted.(gdk.Paintabler)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Paintabler")
			}
			_paintable = rv
		}
	}

	return _paintable
}

// SetAlternativeText sets an alternative textual description for the picture
// contents.
//
// It is equivalent to the "alt" attribute for images on websites.
//
// This text will be made available to accessibility tools.
//
// If the picture cannot be described textually, set this property to NULL.
//
// The function takes the following parameters:
//
//    - alternativeText (optional): textual description of the contents.
//
func (self *Picture) SetAlternativeText(alternativeText string) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if alternativeText != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(alternativeText)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_picture_set_alternative_text(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(alternativeText)
}

// SetCanShrink: if set to TRUE, the self can be made smaller than its contents.
//
// The contents will then be scaled down when rendering.
//
// If you want to still force a minimum size manually, consider using
// gtk.Widget.SetSizeRequest().
//
// Also of note is that a similar function for growing does not exist because
// the grow behavior can be controlled via gtk.Widget.SetHAlign() and
// gtk.Widget.SetVAlign().
//
// The function takes the following parameters:
//
//    - canShrink: if self can be made smaller than its contents.
//
func (self *Picture) SetCanShrink(canShrink bool) {
	var _arg0 *C.GtkPicture // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if canShrink {
		_arg1 = C.TRUE
	}

	C.gtk_picture_set_can_shrink(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canShrink)
}

// SetFile makes self load and display file.
//
// See gtk.Picture.NewForFile for details.
//
// The function takes the following parameters:
//
//    - file (optional): GFile or NULL.
//
func (self *Picture) SetFile(file gio.Filer) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.GFile      // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if file != nil {
		_arg1 = (*C.GFile)(unsafe.Pointer(externglib.InternObject(file).Native()))
	}

	C.gtk_picture_set_file(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(file)
}

// SetFilename makes self load and display the given filename.
//
// This is a utility function that calls gtk.Picture.SetFile().
//
// The function takes the following parameters:
//
//    - filename (optional) to play.
//
func (self *Picture) SetFilename(filename string) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if filename != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_picture_set_filename(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(filename)
}

// SetKeepAspectRatio: if set to TRUE, the self will render its contents
// according to their aspect ratio.
//
// That means that empty space may show up at the top/bottom or left/right of
// self.
//
// If set to FALSE or if the contents provide no aspect ratio, the contents will
// be stretched over the picture's whole area.
//
// The function takes the following parameters:
//
//    - keepAspectRatio: whether to keep aspect ratio.
//
func (self *Picture) SetKeepAspectRatio(keepAspectRatio bool) {
	var _arg0 *C.GtkPicture // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if keepAspectRatio {
		_arg1 = C.TRUE
	}

	C.gtk_picture_set_keep_aspect_ratio(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(keepAspectRatio)
}

// SetPaintable makes self display the given paintable.
//
// If paintable is NULL, nothing will be displayed.
//
// See gtk.Picture.NewForPaintable for details.
//
// The function takes the following parameters:
//
//    - paintable (optional): GdkPaintable or NULL.
//
func (self *Picture) SetPaintable(paintable gdk.Paintabler) {
	var _arg0 *C.GtkPicture   // out
	var _arg1 *C.GdkPaintable // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if paintable != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(externglib.InternObject(paintable).Native()))
	}

	C.gtk_picture_set_paintable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(paintable)
}

// SetPixbuf sets a GtkPicture to show a GdkPixbuf.
//
// See gtk.Picture.NewForPixbuf for details.
//
// This is a utility function that calls gtk.Picture.SetPaintable().
//
// The function takes the following parameters:
//
//    - pixbuf (optional): GdkPixbuf or NULL.
//
func (self *Picture) SetPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.GdkPixbuf  // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(externglib.InternObject(pixbuf).Native()))
	}

	C.gtk_picture_set_pixbuf(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(pixbuf)
}

// SetResource makes self load and display the resource at the given
// resource_path.
//
// This is a utility function that calls gtk.Picture.SetFile().
//
// The function takes the following parameters:
//
//    - resourcePath (optional): resource to set.
//
func (self *Picture) SetResource(resourcePath string) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if resourcePath != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_picture_set_resource(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(resourcePath)
}
