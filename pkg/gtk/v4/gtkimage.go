// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_image_type_get_type()), F: marshalImageType},
		{T: externglib.Type(C.gtk_image_get_type()), F: marshalImager},
	})
}

// ImageType describes the image data representation used by a gtk.Image.
//
// If you want to get the image from the widget, you can only get the
// currently-stored representation; for instance, if the
// gtk_image_get_storage_type() returns GTK_IMAGE_PAINTABLE, then you can call
// gtk_image_get_paintable().
//
// For empty images, you can request any storage type (call any of the "get"
// functions), but they will all return NULL values.
type ImageType int

const (
	// Empty: there is no image displayed by the widget
	ImageTypeEmpty ImageType = iota
	// IconName: widget contains a named icon
	ImageTypeIconName
	// GIcon: widget contains a #GIcon
	ImageTypeGIcon
	// Paintable: widget contains a Paintable
	ImageTypePaintable
)

func marshalImageType(p uintptr) (interface{}, error) {
	return ImageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Image: GtkImage widget displays an image.
//
// !An example GtkImage (image.png)
//
// Various kinds of object can be displayed as an image; most typically, you
// would load a GdkTexture from a file, using the convenience function
// gtk.Image.NewFromFile, for instance:
//
//    GtkWidget *image = gtk_image_new_from_file ("myfile.png");
//
//
// If the file isn’t loaded successfully, the image will contain a “broken
// image” icon similar to that used in many web browsers.
//
// If you want to handle errors in loading the file yourself, for example by
// displaying an error message, then load the image with
// gdk.Texture.NewFromFile, then create the GtkImage with
// gtk.Image.NewFromPaintable.
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of GResource inside GIO, for
// details. In this case, gtk.Image:resource, gtk.Image.NewFromResource, and
// gtk.Image.SetFromResource() should be used.
//
// GtkImage displays its image as an icon, with a size that is determined by the
// application. See gtk.Picture if you want to show an image at is actual size.
//
//
// CSS nodes
//
// GtkImage has a single CSS node with the name image. The style classes
// .normal-icons or .large-icons may appear, depending on the
// gtk.Image:icon-size property.
//
//
// Accessibility
//
// GtkImage uses the GTK_ACCESSIBLE_ROLE_IMG role.
type Image struct {
	Widget
}

var _ gextras.Nativer = (*Image)(nil)

func wrapImage(obj *externglib.Object) *Image {
	return &Image{
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

func marshalImager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapImage(obj), nil
}

// NewImage creates a new empty GtkImage widget.
func NewImage() *Image {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_image_new()

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromFile creates a new GtkImage displaying the file filename.
//
// If the file isn’t found or can’t be loaded, the resulting GtkImage will
// display a “broken image” icon. This function never returns NULL, it always
// returns a valid GtkImage widget.
//
// If you need to detect failures to load the file, use gdk.Texture.NewFromFile
// to load the file yourself, then create the GtkImage from the texture.
//
// The storage type (see gtk.Image.GetStorageType()) of the returned image is
// not defined, it will be whatever is appropriate for displaying the file.
func NewImageFromFile(filename string) *Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))

	_cret = C.gtk_image_new_from_file(_arg1)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromGIcon creates a GtkImage displaying an icon from the current icon
// theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
func NewImageFromGIcon(icon gio.Iconner) *Image {
	var _arg1 *C.GIcon     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GIcon)(unsafe.Pointer((icon).(gextras.Nativer).Native()))

	_cret = C.gtk_image_new_from_gicon(_arg1)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromIconName creates a GtkImage displaying an icon from the current
// icon theme.
//
// If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
func NewImageFromIconName(iconName string) *Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))

	_cret = C.gtk_image_new_from_icon_name(_arg1)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromPaintable creates a new GtkImage displaying paintable.
//
// The GtkImage does not assume a reference to the paintable; you still need to
// unref it if you own references. GtkImage will add its own reference rather
// than adopting yours.
//
// The GtkImage will track changes to the paintable and update its size and
// contents in response to it.
func NewImageFromPaintable(paintable gdk.Paintabler) *Image {
	var _arg1 *C.GdkPaintable // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GdkPaintable)(unsafe.Pointer((paintable).(gextras.Nativer).Native()))

	_cret = C.gtk_image_new_from_paintable(_arg1)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromPixbuf creates a new GtkImage displaying pixbuf.
//
// The GtkImage does not assume a reference to the pixbuf; you still need to
// unref it if you own references. GtkImage will add its own reference rather
// than adopting yours.
//
// This is a helper for gtk.Image.NewFromPaintable, and you can't get back the
// exact pixbuf once this is called, only a texture.
//
// Note that this function just creates an GtkImage from the pixbuf. The
// GtkImage created will not react to state changes. Should you want that, you
// should use gtk.Image.NewFromIconName.
func NewImageFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Image {
	var _arg1 *C.GdkPixbuf // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gtk_image_new_from_pixbuf(_arg1)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromResource creates a new GtkImage displaying the resource file
// resource_path.
//
// If the file isn’t found or can’t be loaded, the resulting GtkImage will
// display a “broken image” icon. This function never returns NULL, it always
// returns a valid GtkImage widget.
//
// If you need to detect failures to load the file, use
// gdkpixbuf.Pixbuf.NewFromFile to load the file yourself, then create the
// GtkImage from the pixbuf.
//
// The storage type (see gtk.Image.GetStorageType()) of the returned image is
// not defined, it will be whatever is appropriate for displaying the file.
func NewImageFromResource(resourcePath string) *Image {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))

	_cret = C.gtk_image_new_from_resource(_arg1)

	var _image *Image // out

	_image = wrapImage(externglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// Clear resets the image to be empty.
func (image *Image) Clear() {
	var _arg0 *C.GtkImage // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	C.gtk_image_clear(_arg0)
}

// GIcon gets the GIcon being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk.Image.GetStorageType()). The caller of this function does not own a
// reference to the returned GIcon.
func (image *Image) GIcon() *gio.Icon {
	var _arg0 *C.GtkImage // out
	var _cret *C.GIcon    // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_gicon(_arg0)

	var _icon *gio.Icon // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_icon = &gio.Icon{
			Object: obj,
		}
	}

	return _icon
}

// IconName gets the icon name and size being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME
// (see gtk.Image.GetStorageType()). The returned string is owned by the
// GtkImage and should not be freed.
func (image *Image) IconName() string {
	var _arg0 *C.GtkImage // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IconSize gets the icon size used by the image when rendering icons.
func (image *Image) IconSize() IconSize {
	var _arg0 *C.GtkImage   // out
	var _cret C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_icon_size(_arg0)

	var _iconSize IconSize // out

	_iconSize = IconSize(_cret)

	return _iconSize
}

// Paintable gets the image GdkPaintable being displayed by the GtkImage.
//
// The storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PAINTABLE
// (see gtk.Image.GetStorageType()). The caller of this function does not own a
// reference to the returned paintable.
func (image *Image) Paintable() *gdk.Paintable {
	var _arg0 *C.GtkImage     // out
	var _cret *C.GdkPaintable // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_paintable(_arg0)

	var _paintable *gdk.Paintable // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_paintable = &gdk.Paintable{
			Object: obj,
		}
	}

	return _paintable
}

// PixelSize gets the pixel size used for named icons.
func (image *Image) PixelSize() int {
	var _arg0 *C.GtkImage // out
	var _cret C.int       // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_pixel_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// StorageType gets the type of representation being used by the GtkImage to
// store image data.
//
// If the GtkImage has no image data, the return value will be GTK_IMAGE_EMPTY.
func (image *Image) StorageType() ImageType {
	var _arg0 *C.GtkImage    // out
	var _cret C.GtkImageType // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_storage_type(_arg0)

	var _imageType ImageType // out

	_imageType = ImageType(_cret)

	return _imageType
}

// SetFromFile sets a GtkImage to show a file.
//
// See gtk.Image.NewFromFile for details.
func (image *Image) SetFromFile(filename string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))

	C.gtk_image_set_from_file(_arg0, _arg1)
}

// SetFromGIcon sets a GtkImage to show a GIcon.
//
// See gtk.Image.NewFromGIcon for details.
func (image *Image) SetFromGIcon(icon gio.Iconner) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.GIcon    // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer((icon).(gextras.Nativer).Native()))

	C.gtk_image_set_from_gicon(_arg0, _arg1)
}

// SetFromIconName sets a GtkImage to show a named icon.
//
// See gtk.Image.NewFromIconName for details.
func (image *Image) SetFromIconName(iconName string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))

	C.gtk_image_set_from_icon_name(_arg0, _arg1)
}

// SetFromPaintable sets a GtkImage to show a GdkPaintable.
//
// See gtk.Image.NewFromPaintable for details.
func (image *Image) SetFromPaintable(paintable gdk.Paintabler) {
	var _arg0 *C.GtkImage     // out
	var _arg1 *C.GdkPaintable // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.GdkPaintable)(unsafe.Pointer((paintable).(gextras.Nativer).Native()))

	C.gtk_image_set_from_paintable(_arg0, _arg1)
}

// SetFromPixbuf sets a GtkImage to show a GdkPixbuf.
//
// See gtk.Image.NewFromPixbuf for details.
//
// Note: This is a helper for gtk.Image.SetFromPaintable(), and you can't get
// back the exact pixbuf once this is called, only a paintable.
func (image *Image) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkImage  // out
	var _arg1 *C.GdkPixbuf // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_image_set_from_pixbuf(_arg0, _arg1)
}

// SetFromResource sets a GtkImage to show a resource.
//
// See gtk.Image.NewFromResource for details.
func (image *Image) SetFromResource(resourcePath string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))

	C.gtk_image_set_from_resource(_arg0, _arg1)
}

// SetIconSize suggests an icon size to the theme for named icons.
func (image *Image) SetIconSize(iconSize IconSize) {
	var _arg0 *C.GtkImage   // out
	var _arg1 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	C.gtk_image_set_icon_size(_arg0, _arg1)
}

// SetPixelSize sets the pixel size to use for named icons.
//
// If the pixel size is set to a value != -1, it is used instead of the icon
// size set by gtk.Image.SetFromIconName().
func (image *Image) SetPixelSize(pixelSize int) {
	var _arg0 *C.GtkImage // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = C.int(pixelSize)

	C.gtk_image_set_pixel_size(_arg0, _arg1)
}
