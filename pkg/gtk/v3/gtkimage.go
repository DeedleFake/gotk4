// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeImageType = coreglib.Type(C.gtk_image_type_get_type())
	GTypeImage     = coreglib.Type(C.gtk_image_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeImageType, F: marshalImageType},
		coreglib.TypeMarshaler{T: GTypeImage, F: marshalImage},
	})
}

// ImageType describes the image data representation used by a Image. If you
// want to get the image from the widget, you can only get the currently-stored
// representation. e.g. if the gtk_image_get_storage_type() returns
// K_IMAGE_PIXBUF, then you can call gtk_image_get_pixbuf() but not
// gtk_image_get_stock(). For empty images, you can request any storage type
// (call any of the "get" functions), but they will all return NULL values.
type ImageType C.gint

const (
	// ImageEmpty: there is no image displayed by the widget.
	ImageEmpty ImageType = iota
	// ImagePixbuf: widget contains a Pixbuf.
	ImagePixbuf
	// ImageStock: widget contains a [stock item name][gtkstock].
	ImageStock
	// ImageIconSet: widget contains a IconSet.
	ImageIconSet
	// ImageAnimation: widget contains a PixbufAnimation.
	ImageAnimation
	// ImageIconName: widget contains a named icon. This image type was added in
	// GTK+ 2.6.
	ImageIconName
	// ImageGIcon: widget contains a #GIcon. This image type was added in GTK+
	// 2.14.
	ImageGIcon
	// ImageSurface: widget contains a #cairo_surface_t. This image type was
	// added in GTK+ 3.10.
	ImageSurface
)

func marshalImageType(p uintptr) (interface{}, error) {
	return ImageType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ImageType.
func (i ImageType) String() string {
	switch i {
	case ImageEmpty:
		return "Empty"
	case ImagePixbuf:
		return "Pixbuf"
	case ImageStock:
		return "Stock"
	case ImageIconSet:
		return "IconSet"
	case ImageAnimation:
		return "Animation"
	case ImageIconName:
		return "IconName"
	case ImageGIcon:
		return "GIcon"
	case ImageSurface:
		return "Surface"
	default:
		return fmt.Sprintf("ImageType(%d)", i)
	}
}

// ImageOverrider contains methods that are overridable.
type ImageOverrider interface {
}

// Image widget displays an image. Various kinds of object can be displayed as
// an image; most typically, you would load a Pixbuf ("pixel buffer") from a
// file, and then display that. There’s a convenience function to do this,
// gtk_image_new_from_file(), used as follows:
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
	_ [0]func() // equal guard
	Misc
}

var (
	_ Miscer = (*Image)(nil)
)

func initClassImage(gclass unsafe.Pointer, goval any) {
}

func wrapImage(obj *coreglib.Object) *Image {
	return &Image{
		Misc: Misc{
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
	}
}

func marshalImage(p uintptr) (interface{}, error) {
	return wrapImage(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewImage creates a new empty Image widget.
//
// The function returns the following values:
//
//    - image: newly created Image widget.
//
func NewImage() *Image {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_image_new()

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

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
//
// The function takes the following parameters:
//
//    - animation: animation.
//
// The function returns the following values:
//
//    - image: new Image widget.
//
func NewImageFromAnimation(animation *gdkpixbuf.PixbufAnimation) *Image {
	var _arg1 *C.GdkPixbufAnimation // out
	var _cret *C.GtkWidget          // in

	_arg1 = (*C.GdkPixbufAnimation)(unsafe.Pointer(coreglib.InternObject(animation).Native()))

	_cret = C.gtk_image_new_from_animation(_arg1)
	runtime.KeepAlive(animation)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromFile creates a new Image displaying the file filename. If the
// file isn’t found or can’t be loaded, the resulting Image will display a
// “broken image” icon. This function never returns NULL, it always returns a
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
//
// The function takes the following parameters:
//
//    - filename: filename.
//
// The function returns the following values:
//
//    - image: new Image.
//
func NewImageFromFile(filename string) *Image {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_file(_arg1)
	runtime.KeepAlive(filename)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromGIcon creates a Image displaying an icon from the current icon
// theme. If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
//
// The function takes the following parameters:
//
//    - icon: icon.
//    - size: stock icon size (IconSize).
//
// The function returns the following values:
//
//    - image: new Image displaying the themed icon.
//
func NewImageFromGIcon(icon gio.Iconner, size int) *Image {
	var _arg1 *C.GIcon      // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_image_new_from_gicon(_arg1, _arg2)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(size)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromIconName creates a Image displaying an icon from the current icon
// theme. If the icon name isn’t known, a “broken image” icon will be displayed
// instead. If the current icon theme is changed, the icon will be updated
// appropriately.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name or NULL.
//    - size: stock icon size (IconSize).
//
// The function returns the following values:
//
//    - image: new Image displaying the themed icon.
//
func NewImageFromIconName(iconName string, size int) *Image {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_image_new_from_icon_name(_arg1, _arg2)
	runtime.KeepAlive(iconName)
	runtime.KeepAlive(size)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

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
//
// The function takes the following parameters:
//
//    - iconSet: IconSet.
//    - size: stock icon size (IconSize).
//
// The function returns the following values:
//
//    - image: new Image.
//
func NewImageFromIconSet(iconSet *IconSet, size int) *Image {
	var _arg1 *C.GtkIconSet // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GtkIconSet)(gextras.StructNative(unsafe.Pointer(iconSet)))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_image_new_from_icon_set(_arg1, _arg2)
	runtime.KeepAlive(iconSet)
	runtime.KeepAlive(size)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromPixbuf creates a new Image displaying pixbuf. The Image does not
// assume a reference to the pixbuf; you still need to unref it if you own
// references. Image will add its own reference rather than adopting yours.
//
// Note that this function just creates an Image from the pixbuf. The Image
// created will not react to state changes. Should you want that, you should use
// gtk_image_new_from_icon_name().
//
// The function takes the following parameters:
//
//    - pixbuf (optional) or NULL.
//
// The function returns the following values:
//
//    - image: new Image.
//
func NewImageFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Image {
	var _arg1 *C.GdkPixbuf // out
	var _cret *C.GtkWidget // in

	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	_cret = C.gtk_image_new_from_pixbuf(_arg1)
	runtime.KeepAlive(pixbuf)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromResource creates a new Image displaying the resource file
// resource_path. If the file isn’t found or can’t be loaded, the resulting
// Image will display a “broken image” icon. This function never returns NULL,
// it always returns a valid Image widget.
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
//
// The function takes the following parameters:
//
//    - resourcePath: resource path.
//
// The function returns the following values:
//
//    - image: new Image.
//
func NewImageFromResource(resourcePath string) *Image {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_new_from_resource(_arg1)
	runtime.KeepAlive(resourcePath)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromStock creates a Image displaying a stock icon. Sample stock icon
// names are K_STOCK_OPEN, K_STOCK_QUIT. Sample stock sizes are
// K_ICON_SIZE_MENU, K_ICON_SIZE_SMALL_TOOLBAR. If the stock icon name isn’t
// known, the image will be empty. You can register your own stock icon names,
// see gtk_icon_factory_add_default() and gtk_icon_factory_add().
//
// Deprecated: Use gtk_image_new_from_icon_name() instead.
//
// The function takes the following parameters:
//
//    - stockId: stock icon name.
//    - size: stock icon size (IconSize).
//
// The function returns the following values:
//
//    - image: new Image displaying the stock icon.
//
func NewImageFromStock(stockId string, size int) *Image {
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_image_new_from_stock(_arg1, _arg2)
	runtime.KeepAlive(stockId)
	runtime.KeepAlive(size)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// NewImageFromSurface creates a new Image displaying surface. The Image does
// not assume a reference to the surface; you still need to unref it if you own
// references. Image will add its own reference rather than adopting yours.
//
// The function takes the following parameters:
//
//    - surface (optional) or NULL.
//
// The function returns the following values:
//
//    - image: new Image.
//
func NewImageFromSurface(surface *cairo.Surface) *Image {
	var _arg1 *C.cairo_surface_t // out
	var _cret *C.GtkWidget       // in

	if surface != nil {
		_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))
	}

	_cret = C.gtk_image_new_from_surface(_arg1)
	runtime.KeepAlive(surface)

	var _image *Image // out

	_image = wrapImage(coreglib.Take(unsafe.Pointer(_cret)))

	return _image
}

// Clear resets the image to be empty.
func (image *Image) Clear() {
	var _arg0 *C.GtkImage // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	C.gtk_image_clear(_arg0)
	runtime.KeepAlive(image)
}

// Animation gets the PixbufAnimation being displayed by the Image. The storage
// type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ANIMATION (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned animation.
//
// The function returns the following values:
//
//    - pixbufAnimation (optional): displayed animation, or NULL if the image is
//      empty.
//
func (image *Image) Animation() *gdkpixbuf.PixbufAnimation {
	var _arg0 *C.GtkImage           // out
	var _cret *C.GdkPixbufAnimation // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_cret = C.gtk_image_get_animation(_arg0)
	runtime.KeepAlive(image)

	var _pixbufAnimation *gdkpixbuf.PixbufAnimation // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_pixbufAnimation = &gdkpixbuf.PixbufAnimation{
				Object: obj,
			}
		}
	}

	return _pixbufAnimation
}

// GIcon gets the #GIcon and size being displayed by the Image. The storage type
// of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned #GIcon.
//
// The function returns the following values:
//
//    - gicon (optional): place to store a #GIcon, or NULL.
//    - size (optional): place to store an icon size (IconSize), or NULL.
//
func (image *Image) GIcon() (*gio.Icon, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GIcon      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	C.gtk_image_get_gicon(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(image)

	var _gicon *gio.Icon // out
	var _size int        // out

	if _arg1 != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_arg1))
			_gicon = &gio.Icon{
				Object: obj,
			}
		}
	}
	_size = int(_arg2)

	return _gicon, _size
}

// IconName gets the icon name and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME (see
// gtk_image_get_storage_type()). The returned string is owned by the Image and
// should not be freed.
//
// The function returns the following values:
//
//    - iconName (optional): place to store an icon name, or NULL.
//    - size (optional): place to store an icon size (IconSize), or NULL.
//
func (image *Image) IconName() (string, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	C.gtk_image_get_icon_name(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(image)

	var _iconName string // out
	var _size int        // out

	if _arg1 != nil {
		_iconName = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	}
	_size = int(_arg2)

	return _iconName, _size
}

// IconSet gets the icon set and size being displayed by the Image. The storage
// type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_SET (see
// gtk_image_get_storage_type()).
//
// Deprecated: Use gtk_image_get_icon_name() instead.
//
// The function returns the following values:
//
//    - iconSet (optional): location to store a IconSet, or NULL.
//    - size (optional): location to store a stock icon size (IconSize), or NULL.
//
func (image *Image) IconSet() (*IconSet, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GtkIconSet // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	C.gtk_image_get_icon_set(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(image)

	var _iconSet *IconSet // out
	var _size int         // out

	if _arg1 != nil {
		_iconSet = (*IconSet)(gextras.NewStructNative(unsafe.Pointer(_arg1)))
		C.gtk_icon_set_ref(_arg1)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_iconSet)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}
	_size = int(_arg2)

	return _iconSet, _size
}

// Pixbuf gets the Pixbuf being displayed by the Image. The storage type of the
// image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned pixbuf.
//
// The function returns the following values:
//
//    - pixbuf (optional): displayed pixbuf, or NULL if the image is empty.
//
func (image *Image) Pixbuf() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkImage  // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_cret = C.gtk_image_get_pixbuf(_arg0)
	runtime.KeepAlive(image)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// PixelSize gets the pixel size used for named icons.
//
// The function returns the following values:
//
//    - gint: pixel size used for named icons.
//
func (image *Image) PixelSize() int {
	var _arg0 *C.GtkImage // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_cret = C.gtk_image_get_pixel_size(_arg0)
	runtime.KeepAlive(image)

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
//
// The function returns the following values:
//
//    - stockId (optional): place to store a stock icon name, or NULL.
//    - size (optional): place to store a stock icon size (IconSize), or NULL.
//
func (image *Image) Stock() (string, int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // in
	var _arg2 C.GtkIconSize // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	C.gtk_image_get_stock(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(image)

	var _stockId string // out
	var _size int       // out

	if _arg1 != nil {
		_stockId = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	}
	_size = int(_arg2)

	return _stockId, _size
}

// StorageType gets the type of representation being used by the Image to store
// image data. If the Image has no image data, the return value will be
// GTK_IMAGE_EMPTY.
//
// The function returns the following values:
//
//    - imageType: image representation being used.
//
func (image *Image) StorageType() ImageType {
	var _arg0 *C.GtkImage    // out
	var _cret C.GtkImageType // in

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))

	_cret = C.gtk_image_get_storage_type(_arg0)
	runtime.KeepAlive(image)

	var _imageType ImageType // out

	_imageType = ImageType(_cret)

	return _imageType
}

// SetFromAnimation causes the Image to display the given animation (or display
// nothing, if you set the animation to NULL).
//
// The function takes the following parameters:
//
//    - animation: PixbufAnimation.
//
func (image *Image) SetFromAnimation(animation *gdkpixbuf.PixbufAnimation) {
	var _arg0 *C.GtkImage           // out
	var _arg1 *C.GdkPixbufAnimation // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	_arg1 = (*C.GdkPixbufAnimation)(unsafe.Pointer(coreglib.InternObject(animation).Native()))

	C.gtk_image_set_from_animation(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(animation)
}

// SetFromFile: see gtk_image_new_from_file() for details.
//
// The function takes the following parameters:
//
//    - filename (optional) or NULL.
//
func (image *Image) SetFromFile(filename string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if filename != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_image_set_from_file(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(filename)
}

// SetFromGIcon: see gtk_image_new_from_gicon() for details.
//
// The function takes the following parameters:
//
//    - icon: icon.
//    - size: icon size (IconSize).
//
func (image *Image) SetFromGIcon(icon gio.Iconner, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GIcon      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))
	_arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_gicon(_arg0, _arg1, _arg2)
	runtime.KeepAlive(image)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(size)
}

// SetFromIconName: see gtk_image_new_from_icon_name() for details.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name or NULL.
//    - size: icon size (IconSize).
//
func (image *Image) SetFromIconName(iconName string, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_icon_name(_arg0, _arg1, _arg2)
	runtime.KeepAlive(image)
	runtime.KeepAlive(iconName)
	runtime.KeepAlive(size)
}

// SetFromIconSet: see gtk_image_new_from_icon_set() for details.
//
// Deprecated: Use gtk_image_set_from_icon_name() instead.
//
// The function takes the following parameters:
//
//    - iconSet: IconSet.
//    - size: stock icon size (IconSize).
//
func (image *Image) SetFromIconSet(iconSet *IconSet, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.GtkIconSet // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	_arg1 = (*C.GtkIconSet)(gextras.StructNative(unsafe.Pointer(iconSet)))
	_arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_icon_set(_arg0, _arg1, _arg2)
	runtime.KeepAlive(image)
	runtime.KeepAlive(iconSet)
	runtime.KeepAlive(size)
}

// SetFromPixbuf: see gtk_image_new_from_pixbuf() for details.
//
// The function takes the following parameters:
//
//    - pixbuf (optional) or NULL.
//
func (image *Image) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkImage  // out
	var _arg1 *C.GdkPixbuf // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	C.gtk_image_set_from_pixbuf(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(pixbuf)
}

// SetFromResource: see gtk_image_new_from_resource() for details.
//
// The function takes the following parameters:
//
//    - resourcePath (optional): resource path or NULL.
//
func (image *Image) SetFromResource(resourcePath string) {
	var _arg0 *C.GtkImage // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if resourcePath != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_image_set_from_resource(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(resourcePath)
}

// SetFromStock: see gtk_image_new_from_stock() for details.
//
// Deprecated: Use gtk_image_set_from_icon_name() instead.
//
// The function takes the following parameters:
//
//    - stockId: stock icon name.
//    - size: stock icon size (IconSize).
//
func (image *Image) SetFromStock(stockId string, size int) {
	var _arg0 *C.GtkImage   // out
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_stock(_arg0, _arg1, _arg2)
	runtime.KeepAlive(image)
	runtime.KeepAlive(stockId)
	runtime.KeepAlive(size)
}

// SetFromSurface: see gtk_image_new_from_surface() for details.
//
// The function takes the following parameters:
//
//    - surface (optional): cairo_surface_t or NULL.
//
func (image *Image) SetFromSurface(surface *cairo.Surface) {
	var _arg0 *C.GtkImage        // out
	var _arg1 *C.cairo_surface_t // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	if surface != nil {
		_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))
	}

	C.gtk_image_set_from_surface(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(surface)
}

// SetPixelSize sets the pixel size to use for named icons. If the pixel size is
// set to a value != -1, it is used instead of the icon size set by
// gtk_image_set_from_icon_name().
//
// The function takes the following parameters:
//
//    - pixelSize: new pixel size.
//
func (image *Image) SetPixelSize(pixelSize int) {
	var _arg0 *C.GtkImage // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkImage)(unsafe.Pointer(coreglib.InternObject(image).Native()))
	_arg1 = C.gint(pixelSize)

	C.gtk_image_set_pixel_size(_arg0, _arg1)
	runtime.KeepAlive(image)
	runtime.KeepAlive(pixelSize)
}
