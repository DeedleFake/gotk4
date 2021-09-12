// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_view_get_type()), F: marshalListViewer},
	})
}

// ListView: GtkListView presents a large dynamic list of items.
//
// GtkListView uses its factory to generate one row widget for each visible item
// and shows them in a linear display, either vertically or horizontally.
//
// The gtk.ListView:show-separators property offers a simple way to display
// separators between the rows.
//
// GtkListView allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// gtk.ListView:enable-rubberband.
//
// If you need multiple columns with headers, see gtk.ColumnView.
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
// An example of using GtkListView:
//
//    static void
//    setup_listitem_cb (GtkListItemFactory *factory,
//                       GtkListItem        *list_item)
//    {
//      GtkWidget *image;
//
//      image = gtk_image_new ();
//      gtk_image_set_icon_size (GTK_IMAGE (image), GTK_ICON_SIZE_LARGE);
//      gtk_list_item_set_child (list_item, image);
//    }
//
//    static void
//    bind_listitem_cb (GtkListItemFactory *factory,
//                      GtkListItem        *list_item)
//    {
//      GtkWidget *image;
//      GAppInfo *app_info;
//
//      image = gtk_list_item_get_child (list_item);
//      app_info = gtk_list_item_get_item (list_item);
//      gtk_image_set_from_gicon (GTK_IMAGE (image), g_app_info_get_icon (app_info));
//    }
//
//    static void
//    activate_cb (GtkListView  *list,
//                 guint         position,
//                 gpointer      unused)
//    {
//      GAppInfo *app_info;
//
//      app_info = g_list_model_get_item (G_LIST_MODEL (gtk_list_view_get_model (list)), position);
//      g_app_info_launch (app_info, NULL, NULL, NULL);
//      g_object_unref (app_info);
//    }
//
//    ...
//
//      model = create_application_list ();
//
//      factory = gtk_signal_list_item_factory_new ();
//      g_signal_connect (factory, "setup", G_CALLBACK (setup_listitem_cb), NULL);
//      g_signal_connect (factory, "bind", G_CALLBACK (bind_listitem_cb), NULL);
//
//      list = gtk_list_view_new (GTK_SELECTION_MODEL (gtk_single_selection_new (model)), factory);
//
//      g_signal_connect (list, "activate", G_CALLBACK (activate_cb), NULL);
//
//      gtk_scrolled_window_set_child (GTK_SCROLLED_WINDOW (sw), list);
//
//
// CSS nodes
//
//    listview[.separators][.rich-list][.navigation-sidebar][.data-table]
//    ├── row
//    │
//    ├── row
//    │
//    ┊
//    ╰── [rubberband]
//
//
// GtkListView uses a single CSS node named listview. It may carry the
// .separators style class, when GtkListView:show-separators property is set.
// Each child widget uses a single CSS node named row. For rubberband selection,
// a node with name rubberband is used.
//
// The main listview node may also carry style classes to select the style of
// list presentation (ListContainers.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// GtkListView uses the GTK_ACCESSIBLE_ROLE_LIST role, and the list items use
// the GTK_ACCESSIBLE_ROLE_LIST_ITEM role.
type ListView struct {
	ListBase
}

func wrapListView(obj *externglib.Object) *ListView {
	return &ListView{
		ListBase: ListBase{
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
				Object: obj,
			},
			Orientable: Orientable{
				Object: obj,
			},
			Scrollable: Scrollable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalListViewer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapListView(obj), nil
}

// NewListView creates a new GtkListView that uses the given factory for mapping
// items to widgets.
//
// The function takes ownership of the arguments, so you can write code like
//
//    list_view = gtk_list_view_new (create_model (),
//      gtk_builder_list_item_factory_new_from_resource ("/resource.ui"));
func NewListView(model SelectionModeller, factory *ListItemFactory) *ListView {
	var _arg1 *C.GtkSelectionModel  // out
	var _arg2 *C.GtkListItemFactory // out
	var _cret *C.GtkWidget          // in

	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}
	if factory != nil {
		_arg2 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))
		C.g_object_ref(C.gpointer(factory.Native()))
	}

	_cret = C.gtk_list_view_new(_arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(factory)

	var _listView *ListView // out

	_listView = wrapListView(externglib.Take(unsafe.Pointer(_cret)))

	return _listView
}

// EnableRubberband returns whether rows can be selected by dragging with the
// mouse.
func (self *ListView) EnableRubberband() bool {
	var _arg0 *C.GtkListView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_view_get_enable_rubberband(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Factory gets the factory that's currently used to populate list items.
func (self *ListView) Factory() *ListItemFactory {
	var _arg0 *C.GtkListView        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_view_get_factory(_arg0)
	runtime.KeepAlive(self)

	var _listItemFactory *ListItemFactory // out

	if _cret != nil {
		_listItemFactory = wrapListItemFactory(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _listItemFactory
}

// Model gets the model that's currently used to read the items displayed.
func (self *ListView) Model() SelectionModeller {
	var _arg0 *C.GtkListView       // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_view_get_model(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel SelectionModeller // out

	if _cret != nil {
		_selectionModel = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(SelectionModeller)
	}

	return _selectionModel
}

// ShowSeparators returns whether the list box should show separators between
// rows.
func (self *ListView) ShowSeparators() bool {
	var _arg0 *C.GtkListView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_view_get_show_separators(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SingleClickActivate returns whether rows will be activated on single click
// and selected on hover.
func (self *ListView) SingleClickActivate() bool {
	var _arg0 *C.GtkListView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_list_view_get_single_click_activate(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetEnableRubberband sets whether selections can be changed by dragging with
// the mouse.
func (self *ListView) SetEnableRubberband(enableRubberband bool) {
	var _arg0 *C.GtkListView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))
	if enableRubberband {
		_arg1 = C.TRUE
	}

	C.gtk_list_view_set_enable_rubberband(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableRubberband)
}

// SetFactory sets the GtkListItemFactory to use for populating list items.
func (self *ListView) SetFactory(factory *ListItemFactory) {
	var _arg0 *C.GtkListView        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))
	if factory != nil {
		_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))
	}

	C.gtk_list_view_set_factory(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(factory)
}

// SetModel sets the model to use.
//
// This must be a gtk.SelectionModel to use.
func (self *ListView) SetModel(model SelectionModeller) {
	var _arg0 *C.GtkListView       // out
	var _arg1 *C.GtkSelectionModel // out

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_list_view_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetShowSeparators sets whether the list box should show separators between
// rows.
func (self *ListView) SetShowSeparators(showSeparators bool) {
	var _arg0 *C.GtkListView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))
	if showSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_list_view_set_show_separators(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(showSeparators)
}

// SetSingleClickActivate sets whether rows should be activated on single click
// and selected on hover.
func (self *ListView) SetSingleClickActivate(singleClickActivate bool) {
	var _arg0 *C.GtkListView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkListView)(unsafe.Pointer(self.Native()))
	if singleClickActivate {
		_arg1 = C.TRUE
	}

	C.gtk_list_view_set_single_click_activate(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(singleClickActivate)
}
