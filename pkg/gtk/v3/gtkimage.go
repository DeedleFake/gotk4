// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_image_type_get_type()), F: marshalImageType},
		{T: externglib.Type(C.gtk_image_get_type()), F: marshalImager},
	})
}

// ImageType describes the image data representation used by a Image. If you
// want to get the image from the widget, you can only get the currently-stored
// representation. e.g. if the gtk_image_get_storage_type() returns
// K_IMAGE_PIXBUF, then you can call gtk_image_get_pixbuf() but not
// gtk_image_get_stock(). For empty images, you can request any storage type
// (call any of the "get" functions), but they will all return nil values.
type ImageType int

const (
	// Empty: there is no image displayed by the widget
	ImageTypeEmpty ImageType = iota
	// Pixbuf: the widget contains a Pixbuf
	ImageTypePixbuf
	// Stock: the widget contains a [stock item name][gtkstock]
	ImageTypeStock
	// IconSet: the widget contains a IconSet
	ImageTypeIconSet
	// Animation: the widget contains a PixbufAnimation
	ImageTypeAnimation
	// IconName: the widget contains a named icon. This image type was added in
	// GTK+ 2.6
	ImageTypeIconName
	// GIcon: the widget contains a #GIcon. This image type was added in GTK+
	// 2.14
	ImageTypeGIcon
	// Surface: the widget contains a #cairo_surface_t. This image type was
	// added in GTK+ 3.10
	ImageTypeSurface
)

func marshalImageType(p uintptr) (interface{}, error) {
	return ImageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Imager describes Image's methods.
type Imager interface {
	gextras.Objector

	Clear()
	Animation() *gdkpixbuf.PixbufAnimation
	GIcon() (*gio.Icon, int)
	IconName() (string, int)
	IconSet() (*IconSet, int)
	Pixbuf() *gdkpixbuf.Pixbuf
	PixelSize() int
	Stock() (string, int)
	StorageType() ImageType
	SetFromAnimation(animation gdkpixbuf.PixbufAnimationer)
	SetFromFile(filename string)
	SetFromGIcon(icon gio.Iconner, size int)
	SetFromIconName(iconName string, size int)
	SetFromIconSet(iconSet *IconSet, size int)
	SetFromPixbuf(pixbuf gdkpixbuf.Pixbuffer)
	SetFromResource(resourcePath string)
	SetFromStock(stockId string, size int)
	SetFromSurface(surface *cairo.Surface)
	SetPixelSize(pixelSize int)
}

// Image: the Image widget displays an image. Various kinds of object can be
// displayed as an image; most typically, you would load a Pixbuf ("pixel
// buffer") from a file, and then display that. There’s a convenience function
// to do this, gtk_image_new_from_file(), used as follows:
//
//      static gboolean
//      button_press_callback (GtkWidget      *event_box,
//                             GdkEventButton *event,
//                             gpointer        data)
//      {
//        g_print ("Event box clicked at coordinates f,f\n",
//                 event->x, event->y);
//
//        // Returning TRUE means we handled the event, so the signal
//        // emission should be stopped (don’t call any further callbacks
//        // that may be connected). Return FALSE to continue invoking callbacks.
//        return TRUE;
//      }
//
//      static GtkWidget*
//      create_image (void)
//      {
//        GtkWidget *image;
//        GtkWidget *event_box;
//
//        image = gtk_image_new_from_file ("myfile.png");
//
//        event_box = gtk_event_box_new ();
//
//        gtk_container_add (GTK_CONTAINER (event_box), image);
//
//        g_signal_connect (G_OBJECT (event_box),
//                          "button_press_event",
//                          G_CALLBACK (button_press_callback),
//                          image);
//
//        return image;
//      }
//
// When handling events on the event box, keep in mind that coordinates in the
// image may be different from event box coordinates due to the alignment and
// padding settings on the image (see Misc). The simplest way to solve this is
// to set the alignment to 0.0 (left/top), and set the padding to zero. Then the
// origin of the image will be the same as the origin of the event box.
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. GTK+ comes with a program to avoid this, called
// “gdk-pixbuf-csource”. This library allows you to convert an image into a C
// variable declaration, which can then be loaded into a Pixbuf using
// gdk_pixbuf_new_from_inline().
//
//
// CSS nodes
//
// GtkImage has a single CSS node with the name image. The style classes may
// appear on image CSS nodes: .icon-dropshadow, .lowres-icon.
type Image struct {
	*externglib.Object

	Misc
	atk.ImplementorIface
	Buildable
}

var _ Imager = (*Image)(nil)

func wrapImager(obj *externglib.Object) Imager {
	return &Image{
		Object: obj,
		Misc: Misc{
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
		},
		ImplementorIface: atk.ImplementorIface{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalImager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapImager(obj), nil
}

// NewImage creates a new empty Image widget.
func NewImage() *Image {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_image_new()

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromAnimation creates a Image displaying the given animation. The
// Image does not assume a reference to the animation; you still need to unref
// it if you own references. Image will add its own reference rather than
// adopting yours.
//
// Note that the animation frames are shown using a timeout with
// PRIORITY_DEFAULT. When using animations to indicate busyness, keep in mind
// that the animation will only be shown if the main loop is not busy with
// something that has a higher priority.
func NewImageFromAnimation(animation gdkpixbuf.PixbufAnimationer) *Image {
	var _arg1 *C.GdkPixbufAnimation // out
	var _cret *C.GtkWidget          // in

	_arg1 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	_cret = C.gtk_image_new_from_animation(_arg1)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromFile creates a new Image displaying the file @filename. If the
// file isn’t found or can’t be loaded, the resulting Image will display a
// “broken image” icon. This function never returns nil, it always returns a
// valid Image widget.
//
// If the file contains an animation, the image will contain an animation.
//
// If you need to detect failures to load the file, use
// gdk_pixbuf_new_from_file() to load the file yourself, then create the Image
// from the pixbuf. (Or for animations, use
// gdk_pixbuf_animation_new_from_file()).
//
// The storage type (gtk_image_get_storage_type()) of the returned image is not
// defined, it will be whatever is appropriate for displaying the file.
func NewImageFromFile(filename string) *Image {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_file(_arg1)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromGIcon creates a Image displaying an icon from the current icon
// theme. If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
func NewImageFromGIcon(icon gio.Iconner, size int) *Image {
	var _arg1 *C.GIcon      // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_image_new_from_gicon(_arg1, _arg2)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromIconName creates a Image displaying an icon from the current icon
// theme. If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
func NewImageFromIconName(iconName string, size int) *Image {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_image_new_from_icon_name(_arg1, _arg2)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromIconSet creates a Image displaying an icon set. Sample stock
// sizes are K_ICON_SIZE_MENU, K_ICON_SIZE_SMALL_TOOLBAR. Instead of using this
// function, usually it’s better to create a IconFactory, put your icon sets in
// the icon factory, add the icon factory to the list of default factories with
// gtk_icon_factory_add_default(), and then use gtk_image_new_from_stock(). This
// will allow themes to override the icon you ship with your application.
//
// The Image does not assume a reference to the icon set; you still need to
// unref it if you own references. Image will add its own reference rather than
// adopting yours.
//
// Deprecated: Use gtk_image_new_from_icon_name() instead.
func NewImageFromIconSet(iconSet *IconSet, size int) *Image {
	var _arg1 *C.GtkIconSet // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GtkIconSet)(unsafe.Pointer(iconSet))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_image_new_from_icon_set(_arg1, _arg2)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromPixbuf creates a new Image displaying @pixbuf. The Image does not
// assume a reference to the pixbuf; you still need to unref it if you own
// references. Image will add its own reference rather than adopting yours.
//
// Note that this function just creates an Image from the pixbuf. The Image
// created will not react to state changes. Should you want that, you should use
// gtk_image_new_from_icon_name().
func NewImageFromPixbuf(pixbuf gdkpixbuf.Pixbuffer) *Image {
	var _arg1 *C.GdkPixbuf // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gtk_image_new_from_pixbuf(_arg1)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromResource creates a new Image displaying the resource file
// @resource_path. If the file isn’t found or can’t be loaded, the resulting
// Image will display a “broken image” icon. This function never returns nil, it
// always returns a valid Image widget.
//
// If the file contains an animation, the image will contain an animation.
//
// If you need to detect failures to load the file, use
// gdk_pixbuf_new_from_file() to load the file yourself, then create the Image
// from the pixbuf. (Or for animations, use
// gdk_pixbuf_animation_new_from_file()).
//
// The storage type (gtk_image_get_storage_type()) of the returned image is not
// defined, it will be whatever is appropriate for displaying the file.
func NewImageFromResource(resourcePath string) *Image {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_resource(_arg1)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromStock creates a Image displaying a stock icon. Sample stock icon
// names are K_STOCK_OPEN, K_STOCK_QUIT. Sample stock sizes are
// K_ICON_SIZE_MENU, K_ICON_SIZE_SMALL_TOOLBAR. If the stock icon name isn’t
// known, the image will be empty. You can register your own stock icon names,
// see gtk_icon_factory_add_default() and gtk_icon_factory_add().
//
// Deprecated: Use gtk_image_new_from_icon_name() instead.
func NewImageFromStock(stockId string, size int) *Image {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_image_new_from_stock(_arg1, _arg2)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// NewImageFromSurface creates a new Image displaying @surface. The Image does
// not assume a reference to the surface; you still need to unref it if you own
// references. Image will add its own reference rather than adopting yours.
func NewImageFromSurface(surface *cairo.Surface) *Image {
	var _arg1 *C.cairo_surface_t // out
	var _cret *C.GtkWidget       // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface))

	_cret = C.gtk_image_new_from_surface(_arg1)

	var _image *Image // out

	_image = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Image)

	return _image
}

// Clear resets the image to be empty.
func (image *Image) Clear() {
	var _arg0 *C.GtkImage // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	C.gtk_image_clear(_arg0)
}

// Animation gets the PixbufAnimation being displayed by the Image. The storage
// type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ANIMATION (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned animation.
func (image *Image) Animation() *gdkpixbuf.PixbufAnimation {
	var _arg0 *C.GtkImage           // out
	var _cret *C.GdkPixbufAnimation // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_animation(_arg0)

	var _pixbufAnimation *gdkpixbuf.PixbufAnimation // out

	_pixbufAnimation = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdkpixbuf.PixbufAnimation)

	return _pixbufAnimation
}

// GIcon gets the #GIcon and size being displayed by the Image. The storage type
// of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned #GIcon.
func (image *Image) GIcon() (*gio.Icon, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GIcon      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	C.gtk_image_get_gicon(_arg0, &_arg1, &_arg2)

	var _gicon *gio.Icon // out
	var _size int        // out

	_gicon = (gextras.CastObject(externglib.Take(unsafe.Pointer(_arg1)))).(*gio.Icon)
	_size = int(_arg2)

	return _gicon, _size
}

// IconName gets the icon name and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME (see
// gtk_image_get_storage_type()). The returned string is owned by the Image and
// should not be freed.
func (image *Image) IconName() (string, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	C.gtk_image_get_icon_name(_arg0, &_arg1, &_arg2)

	var _iconName string // out
	var _size int        // out

	_iconName = C.GoString(_arg1)
	_size = int(_arg2)

	return _iconName, _size
}

// IconSet gets the icon set and size being displayed by the Image. The storage
// type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_SET (see
// gtk_image_get_storage_type()).
//
// Deprecated: Use gtk_image_get_icon_name() instead.
func (image *Image) IconSet() (*IconSet, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GtkIconSet // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	C.gtk_image_get_icon_set(_arg0, &_arg1, &_arg2)

	var _iconSet *IconSet // out
	var _size int         // out

	_iconSet = (*IconSet)(unsafe.Pointer(_arg1))
	C.gtk_icon_set_ref(_arg1)
	runtime.SetFinalizer(_iconSet, func(v *IconSet) {
		C.gtk_icon_set_unref((*C.GtkIconSet)(unsafe.Pointer(v)))
	})
	_size = int(_arg2)

	return _iconSet, _size
}

// Pixbuf gets the Pixbuf being displayed by the Image. The storage type of the
// image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned pixbuf.
func (image *Image) Pixbuf() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkImage  // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_pixbuf(_arg0)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	_pixbuf = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdkpixbuf.Pixbuf)

	return _pixbuf
}

// PixelSize gets the pixel size used for named icons.
func (image *Image) PixelSize() int {
	var _arg0 *C.GtkImage // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_pixel_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Stock gets the stock icon name and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_STOCK (see
// gtk_image_get_storage_type()). The returned string is owned by the Image and
// should not be freed.
//
// Deprecated: Use gtk_image_get_icon_name() instead.
func (image *Image) Stock() (string, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	C.gtk_image_get_stock(_arg0, &_arg1, &_arg2)

	var _stockId string // out
	var _size int       // out

	_stockId = C.GoString(_arg1)
	_size = int(_arg2)

	return _stockId, _size
}

// StorageType gets the type of representation being used by the Image to store
// image data. If the Image has no image data, the return value will be
// GTK_IMAGE_EMPTY.
func (image *Image) StorageType() ImageType {
	var _arg0 *C.GtkImage    // out
	var _cret C.GtkImageType // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))

	_cret = C.gtk_image_get_storage_type(_arg0)

	var _imageType ImageType // out

	_imageType = (ImageType)(_cret)

	return _imageType
}

// SetFromAnimation causes the Image to display the given animation (or display
// nothing, if you set the animation to nil).
func (image *Image) SetFromAnimation(animation gdkpixbuf.PixbufAnimationer) {
	var _arg0 *C.GtkImage           // out
	var _arg1 *C.GdkPixbufAnimation // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	C.gtk_image_set_from_animation(_arg0, _arg1)
}

// SetFromFile: see gtk_image_new_from_file() for details.
func (image *Image) SetFromFile(filename string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_image_set_from_file(_arg0, _arg1)
}

// SetFromGIcon: see gtk_image_new_from_gicon() for details.
func (image *Image) SetFromGIcon(icon gio.Iconner, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GIcon      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_gicon(_arg0, _arg1, _arg2)
}

// SetFromIconName: see gtk_image_new_from_icon_name() for details.
func (image *Image) SetFromIconName(iconName string, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_icon_name(_arg0, _arg1, _arg2)
}

// SetFromIconSet: see gtk_image_new_from_icon_set() for details.
//
// Deprecated: Use gtk_image_set_from_icon_name() instead.
func (image *Image) SetFromIconSet(iconSet *IconSet, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GtkIconSet // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.GtkIconSet)(unsafe.Pointer(iconSet))
	_arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_icon_set(_arg0, _arg1, _arg2)
}

// SetFromPixbuf: see gtk_image_new_from_pixbuf() for details.
func (image *Image) SetFromPixbuf(pixbuf gdkpixbuf.Pixbuffer) {
	var _arg0 *C.GtkImage  // out
	var _arg1 *C.GdkPixbuf // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_image_set_from_pixbuf(_arg0, _arg1)
}

// SetFromResource: see gtk_image_new_from_resource() for details.
func (image *Image) SetFromResource(resourcePath string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_image_set_from_resource(_arg0, _arg1)
}

// SetFromStock: see gtk_image_new_from_stock() for details.
//
// Deprecated: Use gtk_image_set_from_icon_name() instead.
func (image *Image) SetFromStock(stockId string, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_stock(_arg0, _arg1, _arg2)
}

// SetFromSurface: see gtk_image_new_from_surface() for details.
func (image *Image) SetFromSurface(surface *cairo.Surface) {
	var _arg0 *C.GtkImage        // out
	var _arg1 *C.cairo_surface_t // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface))

	C.gtk_image_set_from_surface(_arg0, _arg1)
}

// SetPixelSize sets the pixel size to use for named icons. If the pixel size is
// set to a value != -1, it is used instead of the icon size set by
// gtk_image_set_from_icon_name().
func (image *Image) SetPixelSize(pixelSize int) {
	var _arg0 *C.GtkImage // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(image.Native()))
	_arg1 = C.gint(pixelSize)

	C.gtk_image_set_pixel_size(_arg0, _arg1)
}
