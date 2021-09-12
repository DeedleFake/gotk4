// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_icon_theme_error_get_type()), F: marshalIconThemeError},
		{T: externglib.Type(C.gtk_icon_lookup_flags_get_type()), F: marshalIconLookupFlags},
		{T: externglib.Type(C.gtk_icon_paintable_get_type()), F: marshalIconPaintabler},
		{T: externglib.Type(C.gtk_icon_theme_get_type()), F: marshalIconThemer},
	})
}

// IconThemeError: error codes for GtkIconTheme operations.
type IconThemeError int

const (
	// IconThemeNotFound: icon specified does not exist in the theme
	IconThemeNotFound IconThemeError = iota
	// IconThemeFailed: unspecified error occurred.
	IconThemeFailed
)

func marshalIconThemeError(p uintptr) (interface{}, error) {
	return IconThemeError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for IconThemeError.
func (i IconThemeError) String() string {
	switch i {
	case IconThemeNotFound:
		return "NotFound"
	case IconThemeFailed:
		return "Failed"
	default:
		return fmt.Sprintf("IconThemeError(%d)", i)
	}
}

// IconLookupFlags: used to specify options for gtk_icon_theme_lookup_icon().
type IconLookupFlags int

const (
	// IconLookupForceRegular: try to always load regular icons, even when
	// symbolic icon names are given
	IconLookupForceRegular IconLookupFlags = 0b1
	// IconLookupForceSymbolic: try to always load symbolic icons, even when
	// regular icon names are given
	IconLookupForceSymbolic IconLookupFlags = 0b10
	// IconLookupPreload starts loading the texture in the background so it is
	// ready when later needed.
	IconLookupPreload IconLookupFlags = 0b100
)

func marshalIconLookupFlags(p uintptr) (interface{}, error) {
	return IconLookupFlags(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for IconLookupFlags.
func (i IconLookupFlags) String() string {
	if i == 0 {
		return "IconLookupFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(64)

	for i != 0 {
		next := i & (i - 1)
		bit := i - next

		switch bit {
		case IconLookupForceRegular:
			builder.WriteString("ForceRegular|")
		case IconLookupForceSymbolic:
			builder.WriteString("ForceSymbolic|")
		case IconLookupPreload:
			builder.WriteString("Preload|")
		default:
			builder.WriteString(fmt.Sprintf("IconLookupFlags(0b%b)|", bit))
		}

		i = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if i contains other.
func (i IconLookupFlags) Has(other IconLookupFlags) bool {
	return (i & other) == other
}

// IconPaintable contains information found when looking up an icon in
// GtkIconTheme.
//
// GtkIconPaintable implements GdkPaintable.
type IconPaintable struct {
	*externglib.Object

	gdk.Paintable
}

func wrapIconPaintable(obj *externglib.Object) *IconPaintable {
	return &IconPaintable{
		Object: obj,
		Paintable: gdk.Paintable{
			Object: obj,
		},
	}
}

func marshalIconPaintabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIconPaintable(obj), nil
}

// NewIconPaintableForFile creates a GtkIconPaintable for a file with a given
// size and scale.
//
// The icon can then be rendered by using it as a GdkPaintable.
func NewIconPaintableForFile(file gio.Filer, size int, scale int) *IconPaintable {
	var _arg1 *C.GFile            // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _cret *C.GtkIconPaintable // in

	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))
	_arg2 = C.int(size)
	_arg3 = C.int(scale)

	_cret = C.gtk_icon_paintable_new_for_file(_arg1, _arg2, _arg3)
	runtime.KeepAlive(file)
	runtime.KeepAlive(size)
	runtime.KeepAlive(scale)

	var _iconPaintable *IconPaintable // out

	_iconPaintable = wrapIconPaintable(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _iconPaintable
}

// File gets the GFile that was used to load the icon.
//
// Returns NULL if the icon was not loaded from a file.
func (self *IconPaintable) File() gio.Filer {
	var _arg0 *C.GtkIconPaintable // out
	var _cret *C.GFile            // in

	_arg0 = (*C.GtkIconPaintable)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_paintable_get_file(_arg0)
	runtime.KeepAlive(self)

	var _file gio.Filer // out

	if _cret != nil {
		_file = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gio.Filer)
	}

	return _file
}

// IconName: get the icon name being used for this icon.
//
// When an icon looked up in the icon theme was not available, the icon theme
// may use fallback icons - either those specified to
// gtk_icon_theme_lookup_icon() or the always-available "image-missing". The
// icon chosen is returned by this function.
//
// If the icon was created without an icon theme, this function returns NULL.
func (self *IconPaintable) IconName() string {
	var _arg0 *C.GtkIconPaintable // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkIconPaintable)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_paintable_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _filename string // out

	if _cret != nil {
		_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _filename
}

// IsSymbolic checks if the icon is symbolic or not.
//
// This currently uses only the file name and not the file contents for
// determining this. This behaviour may change in the future.
//
// Note that to render a symbolic GtkIconPaintable properly (with recoloring),
// you have to set its icon name on a GtkImage.
func (self *IconPaintable) IsSymbolic() bool {
	var _arg0 *C.GtkIconPaintable // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkIconPaintable)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_paintable_is_symbolic(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconTheme: GtkIconTheme provides a facility for loading themed icons.
//
// The main reason for using a name rather than simply providing a filename is
// to allow different icons to be used depending on what “icon theme” is
// selected by the user. The operation of icon themes on Linux and Unix follows
// the Icon Theme Specification
// (http://www.freedesktop.org/Standards/icon-theme-spec) There is a fallback
// icon theme, named hicolor, where applications should install their icons, but
// additional icon themes can be installed as operating system vendors and users
// choose.
//
// In many cases, named themes are used indirectly, via gtk.Image rather than
// directly, but looking up icons directly is also simple. The GtkIconTheme
// object acts as a database of all the icons in the current theme. You can
// create new GtkIconTheme objects, but it’s much more efficient to use the
// standard icon theme of the GtkWidget so that the icon information is shared
// with other people looking up icons.
//
//    GtkIconTheme *icon_theme;
//    GtkIconPaintable *icon;
//    GdkPaintable *paintable;
//
//    icon_theme = gtk_icon_theme_get_for_display (gtk_widget_get_display (my_widget));
//    icon = gtk_icon_theme_lookup_icon (icon_theme,
//                                       "my-icon-name", // icon name
//                                       48, // icon size
//                                       1,  // scale
//                                       0,  // flags);
//    paintable = GDK_PAINTABLE (icon);
//    // Use the paintable
//    g_object_unref (icon);
type IconTheme struct {
	*externglib.Object
}

func wrapIconTheme(obj *externglib.Object) *IconTheme {
	return &IconTheme{
		Object: obj,
	}
}

func marshalIconThemer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapIconTheme(obj), nil
}

// NewIconTheme creates a new icon theme object.
//
// Icon theme objects are used to lookup up an icon by name in a particular icon
// theme. Usually, you’ll want to use gtk.IconTheme().GetForDisplay rather than
// creating a new icon theme object for scratch.
func NewIconTheme() *IconTheme {
	var _cret *C.GtkIconTheme // in

	_cret = C.gtk_icon_theme_new()

	var _iconTheme *IconTheme // out

	_iconTheme = wrapIconTheme(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _iconTheme
}

// AddResourcePath adds a resource path that will be looked at when looking for
// icons, similar to search paths.
//
// See gtk.IconTheme.SetResourcePath().
//
// This function should be used to make application-specific icons available as
// part of the icon theme.
func (self *IconTheme) AddResourcePath(path string) {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_add_resource_path(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(path)
}

// AddSearchPath appends a directory to the search path.
//
// See gtk.IconTheme.SetSearchPath().
func (self *IconTheme) AddSearchPath(path string) {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_add_search_path(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(path)
}

// Display returns the display that the GtkIconTheme object was created for.
func (self *IconTheme) Display() *gdk.Display {
	var _arg0 *C.GtkIconTheme // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_display(_arg0)
	runtime.KeepAlive(self)

	var _display *gdk.Display // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_display = &gdk.Display{
				Object: obj,
			}
		}
	}

	return _display
}

// IconNames lists the names of icons in the current icon theme.
func (self *IconTheme) IconNames() []string {
	var _arg0 *C.GtkIconTheme // out
	var _cret **C.char        // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_icon_names(_arg0)
	runtime.KeepAlive(self)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// IconSizes returns an array of integers describing the sizes at which the icon
// is available without scaling.
//
// A size of -1 means that the icon is available in a scalable format. The array
// is zero-terminated.
func (self *IconTheme) IconSizes(iconName string) []int {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out
	var _cret *C.int          // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_icon_theme_get_icon_sizes(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)

	var _gints []int // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z C.int
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_gints = make([]int, i)
		for i := range src {
			_gints[i] = int(src[i])
		}
	}

	return _gints
}

// ResourcePath gets the current resource path.
//
// See gtk.IconTheme.SetResourcePath().
func (self *IconTheme) ResourcePath() []string {
	var _arg0 *C.GtkIconTheme // out
	var _cret **C.char        // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_resource_path(_arg0)
	runtime.KeepAlive(self)

	var _utf8s []string // out

	if _cret != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.char
			for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}

// SearchPath gets the current search path.
//
// See gtk.IconTheme.SetSearchPath().
func (self *IconTheme) SearchPath() []string {
	var _arg0 *C.GtkIconTheme // out
	var _cret **C.char        // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_search_path(_arg0)
	runtime.KeepAlive(self)

	var _filenames []string // out

	if _cret != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.char
			for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_filenames = make([]string, i)
			for i := range src {
				_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _filenames
}

// ThemeName gets the current icon theme name.
//
// Returns (transfer full): the current icon theme name,
func (self *IconTheme) ThemeName() string {
	var _arg0 *C.GtkIconTheme // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_icon_theme_get_theme_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// HasGIcon checks whether an icon theme includes an icon for a particular
// GIcon.
func (self *IconTheme) HasGIcon(gicon gio.Iconner) bool {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.GIcon        // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(gicon.Native()))

	_cret = C.gtk_icon_theme_has_gicon(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(gicon)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasIcon checks whether an icon theme includes an icon for a particular name.
func (self *IconTheme) HasIcon(iconName string) bool {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_icon_theme_has_icon(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LookupByGIcon looks up a icon for a desired size and window scale.
//
// The icon can then be rendered by using it as a GdkPaintable, or you can get
// information such as the filename and size.
func (self *IconTheme) LookupByGIcon(icon gio.Iconner, size int, scale int, direction TextDirection, flags IconLookupFlags) *IconPaintable {
	var _arg0 *C.GtkIconTheme      // out
	var _arg1 *C.GIcon             // out
	var _arg2 C.int                // out
	var _arg3 C.int                // out
	var _arg4 C.GtkTextDirection   // out
	var _arg5 C.GtkIconLookupFlags // out
	var _cret *C.GtkIconPaintable  // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	_arg2 = C.int(size)
	_arg3 = C.int(scale)
	_arg4 = C.GtkTextDirection(direction)
	_arg5 = C.GtkIconLookupFlags(flags)

	_cret = C.gtk_icon_theme_lookup_by_gicon(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(self)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(size)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(direction)
	runtime.KeepAlive(flags)

	var _iconPaintable *IconPaintable // out

	_iconPaintable = wrapIconPaintable(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _iconPaintable
}

// LookupIcon looks up a named icon for a desired size and window scale,
// returning a GtkIconPaintable.
//
// The icon can then be rendered by using it as a GdkPaintable, or you can get
// information such as the filename and size.
//
// If the available icon_name is not available and fallbacks are provided, they
// will be tried in order.
//
// If no matching icon is found, then a paintable that renders the "missing
// icon" icon is returned. If you need to do something else for missing icons
// you need to use gtk.IconTheme.HasIcon().
//
// Note that you probably want to listen for icon theme changes and update the
// icon. This is usually done by overriding the GtkWidgetClass.css-changed()
// function.
func (self *IconTheme) LookupIcon(iconName string, fallbacks []string, size int, scale int, direction TextDirection, flags IconLookupFlags) *IconPaintable {
	var _arg0 *C.GtkIconTheme      // out
	var _arg1 *C.char              // out
	var _arg2 **C.char             // out
	var _arg3 C.int                // out
	var _arg4 C.int                // out
	var _arg5 C.GtkTextDirection   // out
	var _arg6 C.GtkIconLookupFlags // out
	var _cret *C.GtkIconPaintable  // in

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.char)(C.malloc(C.ulong(len(fallbacks)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(fallbacks)+1)
			var zero *C.char
			out[len(fallbacks)] = zero
			for i := range fallbacks {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(fallbacks[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg3 = C.int(size)
	_arg4 = C.int(scale)
	_arg5 = C.GtkTextDirection(direction)
	_arg6 = C.GtkIconLookupFlags(flags)

	_cret = C.gtk_icon_theme_lookup_icon(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
	runtime.KeepAlive(fallbacks)
	runtime.KeepAlive(size)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(direction)
	runtime.KeepAlive(flags)

	var _iconPaintable *IconPaintable // out

	_iconPaintable = wrapIconPaintable(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _iconPaintable
}

// SetSearchPath sets the search path for the icon theme object.
//
// When looking for an icon theme, GTK will search for a subdirectory of one or
// more of the directories in path with the same name as the icon theme
// containing an index.theme file. (Themes from multiple of the path elements
// are combined to allow themes to be extended by adding icons in the user’s
// home directory.)
//
// In addition if an icon found isn’t found either in the current icon theme or
// the default icon theme, and an image file with the right name is found
// directly in one of the elements of path, then that image will be used for the
// icon name. (This is legacy feature, and new icons should be put into the
// fallback icon theme, which is called hicolor, rather than directly on the
// icon path.)
func (self *IconTheme) SetSearchPath(path []string) {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 **C.char        // out

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	{
		_arg1 = (**C.char)(C.malloc(C.ulong(len(path)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(path)+1)
			var zero *C.char
			out[len(path)] = zero
			for i := range path {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(path[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_icon_theme_set_search_path(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(path)
}

// SetThemeName sets the name of the icon theme that the GtkIconTheme object
// uses overriding system configuration.
//
// This function cannot be called on the icon theme objects returned from
// gtk.IconTheme.GetForDisplay.
func (self *IconTheme) SetThemeName(themeName string) {
	var _arg0 *C.GtkIconTheme // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(self.Native()))
	if themeName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(themeName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_icon_theme_set_theme_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(themeName)
}

// IconThemeGetForDisplay gets the icon theme object associated with display.
//
// If this function has not previously been called for the given display, a new
// icon theme object will be created and associated with the display. Icon theme
// objects are fairly expensive to create, so using this function is usually a
// better choice than calling gtk.IconTheme.New and setting the display
// yourself; by using this function a single icon theme object will be shared
// between users.
func IconThemeGetForDisplay(display *gdk.Display) *IconTheme {
	var _arg1 *C.GdkDisplay   // out
	var _cret *C.GtkIconTheme // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gtk_icon_theme_get_for_display(_arg1)
	runtime.KeepAlive(display)

	var _iconTheme *IconTheme // out

	_iconTheme = wrapIconTheme(externglib.Take(unsafe.Pointer(_cret)))

	return _iconTheme
}
