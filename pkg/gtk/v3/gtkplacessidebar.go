// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_places_open_flags_get_type()), F: marshalPlacesOpenFlags},
		{T: externglib.Type(C.gtk_places_sidebar_get_type()), F: marshalPlacesSidebarrer},
	})
}

// PlacesOpenFlags: these flags serve two purposes. First, the application can
// call gtk_places_sidebar_set_open_flags() using these flags as a bitmask. This
// tells the sidebar that the application is able to open folders selected from
// the sidebar in various ways, for example, in new tabs or in new windows in
// addition to the normal mode.
//
// Second, when one of these values gets passed back to the application in the
// PlacesSidebar::open-location signal, it means that the application should
// open the selected location in the normal way, in a new tab, or in a new
// window. The sidebar takes care of determining the desired way to open the
// location, based on the modifier keys that the user is pressing at the time
// the selection is made.
//
// If the application never calls gtk_places_sidebar_set_open_flags(), then the
// sidebar will only use K_PLACES_OPEN_NORMAL in the
// PlacesSidebar::open-location signal. This is the default mode of operation.
type PlacesOpenFlags C.guint

const (
	// PlacesOpenNormal: this is the default mode that PlacesSidebar uses if no
	// other flags are specified. It indicates that the calling application
	// should open the selected location in the normal way, for example, in the
	// folder view beside the sidebar.
	PlacesOpenNormal PlacesOpenFlags = 0b1
	// PlacesOpenNewTab: when passed to gtk_places_sidebar_set_open_flags(),
	// this indicates that the application can open folders selected from the
	// sidebar in new tabs. This value will be passed to the
	// PlacesSidebar::open-location signal when the user selects that a location
	// be opened in a new tab instead of in the standard fashion.
	PlacesOpenNewTab PlacesOpenFlags = 0b10
	// PlacesOpenNewWindow: similar to GTK_PLACES_OPEN_NEW_TAB, but indicates
	// that the application can open folders in new windows.
	PlacesOpenNewWindow PlacesOpenFlags = 0b100
)

func marshalPlacesOpenFlags(p uintptr) (interface{}, error) {
	return PlacesOpenFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for PlacesOpenFlags.
func (p PlacesOpenFlags) String() string {
	if p == 0 {
		return "PlacesOpenFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(53)

	for p != 0 {
		next := p & (p - 1)
		bit := p - next

		switch bit {
		case PlacesOpenNormal:
			builder.WriteString("Normal|")
		case PlacesOpenNewTab:
			builder.WriteString("NewTab|")
		case PlacesOpenNewWindow:
			builder.WriteString("NewWindow|")
		default:
			builder.WriteString(fmt.Sprintf("PlacesOpenFlags(0b%b)|", bit))
		}

		p = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if p contains other.
func (p PlacesOpenFlags) Has(other PlacesOpenFlags) bool {
	return (p & other) == other
}

// PlacesSidebar is a widget that displays a list of frequently-used places in
// the file system: the user’s home directory, the user’s bookmarks, and volumes
// and drives. This widget is used as a sidebar in FileChooser and may be used
// by file managers and similar programs.
//
// The places sidebar displays drives and volumes, and will automatically mount
// or unmount them when the user selects them.
//
// Applications can hook to various signals in the places sidebar to customize
// its behavior. For example, they can add extra commands to the context menu of
// the sidebar.
//
// While bookmarks are completely in control of the user, the places sidebar
// also allows individual applications to provide extra shortcut folders that
// are unique to each application. For example, a Paint program may want to add
// a shortcut for a Clipart folder. You can do this with
// gtk_places_sidebar_add_shortcut().
//
// To make use of the places sidebar, an application at least needs to connect
// to the PlacesSidebar::open-location signal. This is emitted when the user
// selects in the sidebar a location to open. The application should also call
// gtk_places_sidebar_set_location() when it changes the currently-viewed
// location.
//
//
// CSS nodes
//
// GtkPlacesSidebar uses a single CSS node with name placessidebar and style
// class .sidebar.
//
// Among the children of the places sidebar, the following style classes can be
// used:
//
// - .sidebar-new-bookmark-row for the 'Add new bookmark' row
//
// - .sidebar-placeholder-row for a row that is a placeholder
//
// - .has-open-popup when a popup is open for a row.
type PlacesSidebar struct {
	_ [0]func() // equal guard
	ScrolledWindow
}

var (
	_ Binner = (*PlacesSidebar)(nil)
)

func wrapPlacesSidebar(obj *externglib.Object) *PlacesSidebar {
	return &PlacesSidebar{
		ScrolledWindow: ScrolledWindow{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
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
			},
		},
	}
}

func marshalPlacesSidebarrer(p uintptr) (interface{}, error) {
	return wrapPlacesSidebar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectDragActionAsk places sidebar emits this signal when it needs to ask
// the application to pop up a menu to ask the user for which drag action to
// perform.
func (sidebar *PlacesSidebar) ConnectDragActionAsk(f func(actions int) int) externglib.SignalHandle {
	return sidebar.Connect("drag-action-ask", f)
}

// ConnectMount places sidebar emits this signal when it starts a new operation
// because the user clicked on some location that needs mounting. In this way
// the application using the PlacesSidebar can track the progress of the
// operation and, for example, show a notification.
func (sidebar *PlacesSidebar) ConnectMount(f func(mountOperation gio.MountOperation)) externglib.SignalHandle {
	return sidebar.Connect("mount", f)
}

// ConnectOpenLocation places sidebar emits this signal when the user selects a
// location in it. The calling application should display the contents of that
// location; for example, a file manager should show a list of files in the
// specified location.
func (sidebar *PlacesSidebar) ConnectOpenLocation(f func(location gio.Filer, openFlags PlacesOpenFlags)) externglib.SignalHandle {
	return sidebar.Connect("open-location", f)
}

// ConnectPopulatePopup places sidebar emits this signal when the user invokes a
// contextual popup on one of its items. In the signal handler, the application
// may add extra items to the menu as appropriate. For example, a file manager
// may want to add a "Properties" command to the menu.
//
// It is not necessary to store the selected_item for each menu item; during
// their callbacks, the application can use gtk_places_sidebar_get_location() to
// get the file to which the item refers.
//
// The selected_item argument may be NULL in case the selection refers to a
// volume. In this case, selected_volume will be non-NULL. In this case, the
// calling application will have to g_object_ref() the selected_volume and keep
// it around to use it in the callback.
//
// The container and all its contents are destroyed after the user dismisses the
// popup. The popup is re-created (and thus, this signal is emitted) every time
// the user activates the contextual menu.
//
// Before 3.18, the container always was a Menu, and you were expected to add
// your items as MenuItems. Since 3.18, the popup may be implemented as a
// Popover, in which case container will be something else, e.g. a Box, to which
// you may add ModelButtons or other widgets, such as Entries, SpinButtons, etc.
// If your application can deal with this situation, you can set
// PlacesSidebar::populate-all to TRUE to request that this signal is emitted
// for populating popovers as well.
func (sidebar *PlacesSidebar) ConnectPopulatePopup(f func(container Widgetter, selectedItem gio.Filer, selectedVolume gio.Volumer)) externglib.SignalHandle {
	return sidebar.Connect("populate-popup", f)
}

// ConnectShowConnectToServer places sidebar emits this signal when it needs the
// calling application to present an way to connect directly to a network
// server. For example, the application may bring up a dialog box asking for a
// URL like "sftp://ftp.example.com". It is up to the application to create the
// corresponding mount by using, for example, g_file_mount_enclosing_volume().
func (sidebar *PlacesSidebar) ConnectShowConnectToServer(f func()) externglib.SignalHandle {
	return sidebar.Connect("show-connect-to-server", f)
}

// ConnectShowEnterLocation places sidebar emits this signal when it needs the
// calling application to present an way to directly enter a location. For
// example, the application may bring up a dialog box asking for a URL like
// "http://http.example.com".
func (sidebar *PlacesSidebar) ConnectShowEnterLocation(f func()) externglib.SignalHandle {
	return sidebar.Connect("show-enter-location", f)
}

// ConnectShowErrorMessage places sidebar emits this signal when it needs the
// calling application to present an error message. Most of these messages refer
// to mounting or unmounting media, for example, when a drive cannot be started
// for some reason.
func (sidebar *PlacesSidebar) ConnectShowErrorMessage(f func(primary, secondary string)) externglib.SignalHandle {
	return sidebar.Connect("show-error-message", f)
}

// ConnectShowOtherLocations places sidebar emits this signal when it needs the
// calling application to present a way to show other locations e.g. drives and
// network access points. For example, the application may bring up a page
// showing persistent volumes and discovered network addresses.
func (sidebar *PlacesSidebar) ConnectShowOtherLocations(f func()) externglib.SignalHandle {
	return sidebar.Connect("show-other-locations", f)
}

// ConnectShowOtherLocationsWithFlags places sidebar emits this signal when it
// needs the calling application to present a way to show other locations e.g.
// drives and network access points. For example, the application may bring up a
// page showing persistent volumes and discovered network addresses.
func (sidebar *PlacesSidebar) ConnectShowOtherLocationsWithFlags(f func(openFlags PlacesOpenFlags)) externglib.SignalHandle {
	return sidebar.Connect("show-other-locations-with-flags", f)
}

// ConnectShowStarredLocation places sidebar emits this signal when it needs the
// calling application to present a way to show the starred files. In GNOME,
// starred files are implemented by setting the nao:predefined-tag-favorite tag
// in the tracker database.
func (sidebar *PlacesSidebar) ConnectShowStarredLocation(f func(openFlags PlacesOpenFlags)) externglib.SignalHandle {
	return sidebar.Connect("show-starred-location", f)
}

// ConnectUnmount places sidebar emits this signal when it starts a new
// operation because the user for example ejected some drive or unmounted a
// mount. In this way the application using the PlacesSidebar can track the
// progress of the operation and, for example, show a notification.
func (sidebar *PlacesSidebar) ConnectUnmount(f func(mountOperation gio.MountOperation)) externglib.SignalHandle {
	return sidebar.Connect("unmount", f)
}

// NewPlacesSidebar creates a new PlacesSidebar widget.
//
// The application should connect to at least the PlacesSidebar::open-location
// signal to be notified when the user makes a selection in the sidebar.
//
// The function returns the following values:
//
//    - placesSidebar: newly created PlacesSidebar.
//
func NewPlacesSidebar() *PlacesSidebar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_places_sidebar_new()

	var _placesSidebar *PlacesSidebar // out

	_placesSidebar = wrapPlacesSidebar(externglib.Take(unsafe.Pointer(_cret)))

	return _placesSidebar
}

// AddShortcut applications may want to present some folders in the places
// sidebar if they could be immediately useful to users. For example, a drawing
// program could add a “/usr/share/clipart” location when the sidebar is being
// used in an “Insert Clipart” dialog box.
//
// This function adds the specified location to a special place for immutable
// shortcuts. The shortcuts are application-specific; they are not shared across
// applications, and they are not persistent. If this function is called
// multiple times with different locations, then they are added to the sidebar’s
// list in the same order as the function is called.
//
// The function takes the following parameters:
//
//    - location to add as an application-specific shortcut.
//
func (sidebar *PlacesSidebar) AddShortcut(location gio.Filer) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 *C.GFile            // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(location.Native()))

	C.gtk_places_sidebar_add_shortcut(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(location)
}

// LocalOnly returns the value previously set with
// gtk_places_sidebar_set_local_only().
//
// The function returns the following values:
//
//    - ok: TRUE if the sidebar will only show local files.
//
func (sidebar *PlacesSidebar) LocalOnly() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_local_only(_arg0)
	runtime.KeepAlive(sidebar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Location gets the currently selected location in the sidebar. This can be
// NULL when nothing is selected, for example, when
// gtk_places_sidebar_set_location() has been called with a location that is not
// among the sidebar’s list of places to show.
//
// You can use this function to get the selection in the sidebar. Also, if you
// connect to the PlacesSidebar::populate-popup signal, you can use this
// function to get the location that is being referred to during the callbacks
// for your menu items.
//
// The function returns the following values:
//
//    - file (optional) with the selected location, or NULL if nothing is
//      visually selected.
//
func (sidebar *PlacesSidebar) Location() gio.Filer {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret *C.GFile            // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_location(_arg0)
	runtime.KeepAlive(sidebar)

	var _file gio.Filer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
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

// NthBookmark: this function queries the bookmarks added by the user to the
// places sidebar, and returns one of them. This function is used by FileChooser
// to implement the “Alt-1”, “Alt-2”, etc. shortcuts, which activate the
// cooresponding bookmark.
//
// The function takes the following parameters:
//
//    - n: index of the bookmark to query.
//
// The function returns the following values:
//
//    - file (optional): bookmark specified by the index n, or NULL if no such
//      index exist. Note that the indices start at 0, even though the file
//      chooser starts them with the keyboard shortcut "Alt-1".
//
func (sidebar *PlacesSidebar) NthBookmark(n int) gio.Filer {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gint              // out
	var _cret *C.GFile            // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = C.gint(n)

	_cret = C.gtk_places_sidebar_get_nth_bookmark(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(n)

	var _file gio.Filer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
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

// OpenFlags gets the open flags.
//
// The function returns the following values:
//
//    - placesOpenFlags of sidebar.
//
func (sidebar *PlacesSidebar) OpenFlags() PlacesOpenFlags {
	var _arg0 *C.GtkPlacesSidebar  // out
	var _cret C.GtkPlacesOpenFlags // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_open_flags(_arg0)
	runtime.KeepAlive(sidebar)

	var _placesOpenFlags PlacesOpenFlags // out

	_placesOpenFlags = PlacesOpenFlags(_cret)

	return _placesOpenFlags
}

// ShowConnectToServer returns the value previously set with
// gtk_places_sidebar_set_show_connect_to_server()
//
// Deprecated: It is recommended to group this functionality with the drives and
// network location under the new 'Other Location' item.
//
// The function returns the following values:
//
//    - ok: TRUE if the sidebar will display a “Connect to Server” item.
//
func (sidebar *PlacesSidebar) ShowConnectToServer() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_connect_to_server(_arg0)
	runtime.KeepAlive(sidebar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDesktop returns the value previously set with
// gtk_places_sidebar_set_show_desktop().
//
// The function returns the following values:
//
//    - ok: TRUE if the sidebar will display a builtin shortcut to the desktop
//      folder.
//
func (sidebar *PlacesSidebar) ShowDesktop() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_desktop(_arg0)
	runtime.KeepAlive(sidebar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowEnterLocation returns the value previously set with
// gtk_places_sidebar_set_show_enter_location().
//
// The function returns the following values:
//
//    - ok: TRUE if the sidebar will display an “Enter Location” item.
//
func (sidebar *PlacesSidebar) ShowEnterLocation() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_enter_location(_arg0)
	runtime.KeepAlive(sidebar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowOtherLocations returns the value previously set with
// gtk_places_sidebar_set_show_other_locations().
//
// The function returns the following values:
//
//    - ok: TRUE if the sidebar will display an “Other Locations” item.
//
func (sidebar *PlacesSidebar) ShowOtherLocations() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_other_locations(_arg0)
	runtime.KeepAlive(sidebar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRecent returns the value previously set with
// gtk_places_sidebar_set_show_recent().
//
// The function returns the following values:
//
//    - ok: TRUE if the sidebar will display a builtin shortcut for recent files.
//
func (sidebar *PlacesSidebar) ShowRecent() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_recent(_arg0)
	runtime.KeepAlive(sidebar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowStarredLocation returns the value previously set with
// gtk_places_sidebar_set_show_starred_location().
//
// The function returns the following values:
//
//    - ok: TRUE if the sidebar will display a Starred item.
//
func (sidebar *PlacesSidebar) ShowStarredLocation() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_starred_location(_arg0)
	runtime.KeepAlive(sidebar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowTrash returns the value previously set with
// gtk_places_sidebar_set_show_trash().
//
// The function returns the following values:
//
//    - ok: TRUE if the sidebar will display a “Trash” item.
//
func (sidebar *PlacesSidebar) ShowTrash() bool {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_get_show_trash(_arg0)
	runtime.KeepAlive(sidebar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListShortcuts gets the list of shortcuts.
//
// The function returns the following values:
//
//    - sList: A List of #GFile of the locations that have been added as
//      application-specific shortcuts with gtk_places_sidebar_add_shortcut(). To
//      free this list, you can use
//
//         g_slist_free_full (list, (GDestroyNotify) g_object_unref);.
//
func (sidebar *PlacesSidebar) ListShortcuts() []gio.Filer {
	var _arg0 *C.GtkPlacesSidebar // out
	var _cret *C.GSList           // in

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))

	_cret = C.gtk_places_sidebar_list_shortcuts(_arg0)
	runtime.KeepAlive(sidebar)

	var _sList []gio.Filer // out

	_sList = make([]gio.Filer, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GFile)(v)
		var dst gio.Filer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gio.Filer is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gio.Filer)
				return ok
			})
			rv, ok := casted.(gio.Filer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
			}
			dst = rv
		}
		_sList = append(_sList, dst)
	})

	return _sList
}

// RemoveShortcut removes an application-specific shortcut that has been
// previously been inserted with gtk_places_sidebar_add_shortcut(). If the
// location is not a shortcut in the sidebar, then nothing is done.
//
// The function takes the following parameters:
//
//    - location to remove.
//
func (sidebar *PlacesSidebar) RemoveShortcut(location gio.Filer) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 *C.GFile            // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(location.Native()))

	C.gtk_places_sidebar_remove_shortcut(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(location)
}

// SetDropTargetsVisible: make the GtkPlacesSidebar show drop targets, so it can
// show the available drop targets and a "new bookmark" row. This improves the
// Drag-and-Drop experience of the user and allows applications to show all
// available drop targets at once.
//
// This needs to be called when the application is aware of an ongoing drag that
// might target the sidebar. The drop-targets-visible state will be unset
// automatically if the drag finishes in the GtkPlacesSidebar. You only need to
// unset the state when the drag ends on some other widget on your application.
//
// The function takes the following parameters:
//
//    - visible: whether to show the valid targets or not.
//    - context: drag context used to ask the source about the action that wants
//      to perform, so hints are more accurate.
//
func (sidebar *PlacesSidebar) SetDropTargetsVisible(visible bool, context *gdk.DragContext) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out
	var _arg2 *C.GdkDragContext   // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if visible {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	C.gtk_places_sidebar_set_drop_targets_visible(_arg0, _arg1, _arg2)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(visible)
	runtime.KeepAlive(context)
}

// SetLocalOnly sets whether the sidebar should only show local files.
//
// The function takes the following parameters:
//
//    - localOnly: whether to show only local files.
//
func (sidebar *PlacesSidebar) SetLocalOnly(localOnly bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if localOnly {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_local_only(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(localOnly)
}

// SetLocation sets the location that is being shown in the widgets surrounding
// the sidebar, for example, in a folder view in a file manager. In turn, the
// sidebar will highlight that location if it is being shown in the list of
// places, or it will unhighlight everything if the location is not among the
// places in the list.
//
// The function takes the following parameters:
//
//    - location (optional) to select, or NULL for no current path.
//
func (sidebar *PlacesSidebar) SetLocation(location gio.Filer) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 *C.GFile            // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if location != nil {
		_arg1 = (*C.GFile)(unsafe.Pointer(location.Native()))
	}

	C.gtk_places_sidebar_set_location(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(location)
}

// SetOpenFlags sets the way in which the calling application can open new
// locations from the places sidebar. For example, some applications only open
// locations “directly” into their main view, while others may support opening
// locations in a new notebook tab or a new window.
//
// This function is used to tell the places sidebar about the ways in which the
// application can open new locations, so that the sidebar can display (or not)
// the “Open in new tab” and “Open in new window” menu items as appropriate.
//
// When the PlacesSidebar::open-location signal is emitted, its flags argument
// will be set to one of the flags that was passed in
// gtk_places_sidebar_set_open_flags().
//
// Passing 0 for flags will cause K_PLACES_OPEN_NORMAL to always be sent to
// callbacks for the “open-location” signal.
//
// The function takes the following parameters:
//
//    - flags: bitmask of modes in which the calling application can open
//      locations.
//
func (sidebar *PlacesSidebar) SetOpenFlags(flags PlacesOpenFlags) {
	var _arg0 *C.GtkPlacesSidebar  // out
	var _arg1 C.GtkPlacesOpenFlags // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	_arg1 = C.GtkPlacesOpenFlags(flags)

	C.gtk_places_sidebar_set_open_flags(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(flags)
}

// SetShowConnectToServer sets whether the sidebar should show an item for
// connecting to a network server; this is off by default. An application may
// want to turn this on if it implements a way for the user to connect to
// network servers directly.
//
// If you enable this, you should connect to the
// PlacesSidebar::show-connect-to-server signal.
//
// Deprecated: It is recommended to group this functionality with the drives and
// network location under the new 'Other Location' item.
//
// The function takes the following parameters:
//
//    - showConnectToServer: whether to show an item for the Connect to Server
//      command.
//
func (sidebar *PlacesSidebar) SetShowConnectToServer(showConnectToServer bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showConnectToServer {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_connect_to_server(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(showConnectToServer)
}

// SetShowDesktop sets whether the sidebar should show an item for the Desktop
// folder. The default value for this option is determined by the desktop
// environment and the user’s configuration, but this function can be used to
// override it on a per-application basis.
//
// The function takes the following parameters:
//
//    - showDesktop: whether to show an item for the Desktop folder.
//
func (sidebar *PlacesSidebar) SetShowDesktop(showDesktop bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showDesktop {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_desktop(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(showDesktop)
}

// SetShowEnterLocation sets whether the sidebar should show an item for
// entering a location; this is off by default. An application may want to turn
// this on if manually entering URLs is an expected user action.
//
// If you enable this, you should connect to the
// PlacesSidebar::show-enter-location signal.
//
// The function takes the following parameters:
//
//    - showEnterLocation: whether to show an item to enter a location.
//
func (sidebar *PlacesSidebar) SetShowEnterLocation(showEnterLocation bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showEnterLocation {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_enter_location(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(showEnterLocation)
}

// SetShowOtherLocations sets whether the sidebar should show an item for the
// application to show an Other Locations view; this is off by default. When set
// to TRUE, persistent devices such as hard drives are hidden, otherwise they
// are shown in the sidebar. An application may want to turn this on if it
// implements a way for the user to see and interact with drives and network
// servers directly.
//
// If you enable this, you should connect to the
// PlacesSidebar::show-other-locations signal.
//
// The function takes the following parameters:
//
//    - showOtherLocations: whether to show an item for the Other Locations view.
//
func (sidebar *PlacesSidebar) SetShowOtherLocations(showOtherLocations bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showOtherLocations {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_other_locations(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(showOtherLocations)
}

// SetShowRecent sets whether the sidebar should show an item for recent files.
// The default value for this option is determined by the desktop environment,
// but this function can be used to override it on a per-application basis.
//
// The function takes the following parameters:
//
//    - showRecent: whether to show an item for recent files.
//
func (sidebar *PlacesSidebar) SetShowRecent(showRecent bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showRecent {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_recent(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(showRecent)
}

// SetShowStarredLocation: if you enable this, you should connect to the
// PlacesSidebar::show-starred-location signal.
//
// The function takes the following parameters:
//
//    - showStarredLocation: whether to show an item for Starred files.
//
func (sidebar *PlacesSidebar) SetShowStarredLocation(showStarredLocation bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showStarredLocation {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_starred_location(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(showStarredLocation)
}

// SetShowTrash sets whether the sidebar should show an item for the Trash
// location.
//
// The function takes the following parameters:
//
//    - showTrash: whether to show an item for the Trash location.
//
func (sidebar *PlacesSidebar) SetShowTrash(showTrash bool) {
	var _arg0 *C.GtkPlacesSidebar // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar.Native()))
	if showTrash {
		_arg1 = C.TRUE
	}

	C.gtk_places_sidebar_set_show_trash(_arg0, _arg1)
	runtime.KeepAlive(sidebar)
	runtime.KeepAlive(showTrash)
}
