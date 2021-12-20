// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_down_get_type()), F: marshalDropDowner},
	})
}

// DropDown: GtkDropDown is a widget that allows the user to choose an item from
// a list of options.
//
// !An example GtkDropDown (drop-down.png)
//
// The GtkDropDown displays the selected choice.
//
// The options are given to GtkDropDown in the form of GListModel and how the
// individual options are represented is determined by a gtk.ListItemFactory.
// The default factory displays simple strings.
//
// GtkDropDown knows how to obtain strings from the items in a gtk.StringList;
// for other models, you have to provide an expression to find the strings via
// gtk.DropDown.SetExpression().
//
// GtkDropDown can optionally allow search in the popup, which is useful if the
// list of options is long. To enable the search entry, use
// gtk.DropDown.SetEnableSearch().
//
//
// CSS nodes
//
// GtkDropDown has a single CSS node with name dropdown, with the button and
// popover nodes as children.
//
//
// Accessibility
//
// GtkDropDown uses the GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type DropDown struct {
	Widget
}

var (
	_ Widgetter = (*DropDown)(nil)
)

func wrapDropDown(obj *externglib.Object) *DropDown {
	return &DropDown{
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

func marshalDropDowner(p uintptr) (interface{}, error) {
	return wrapDropDown(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewDropDown creates a new GtkDropDown.
//
// You may want to call gtk.DropDown.SetFactory() to set up a way to map its
// items to widgets.
//
// The function takes the following parameters:
//
//    - model (optional) to use or NULL for none.
//    - expression (optional) to use or NULL for none.
//
// The function returns the following values:
//
//    - dropDown: new GtkDropDown.
//
func NewDropDown(model gio.ListModeller, expression Expressioner) *DropDown {
	var _arg1 *C.GListModel    // out
	var _arg2 *C.GtkExpression // out
	var _cret *C.GtkWidget     // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}
	if expression != nil {
		_arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
		C.g_object_ref(C.gpointer(expression.Native()))
	}

	_cret = C.gtk_drop_down_new(_arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(expression)

	var _dropDown *DropDown // out

	_dropDown = wrapDropDown(externglib.Take(unsafe.Pointer(_cret)))

	return _dropDown
}

// NewDropDownFromStrings creates a new GtkDropDown that is populated with the
// strings.
//
// The function takes the following parameters:
//
//    - strings to put in the dropdown.
//
// The function returns the following values:
//
//    - dropDown: new GtkDropDown.
//
func NewDropDownFromStrings(strings []string) *DropDown {
	var _arg1 **C.char     // out
	var _cret *C.GtkWidget // in

	{
		_arg1 = (**C.char)(C.calloc(C.size_t((len(strings) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(strings)+1)
			var zero *C.char
			out[len(strings)] = zero
			for i := range strings {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(strings[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gtk_drop_down_new_from_strings(_arg1)
	runtime.KeepAlive(strings)

	var _dropDown *DropDown // out

	_dropDown = wrapDropDown(externglib.Take(unsafe.Pointer(_cret)))

	return _dropDown
}

// EnableSearch returns whether search is enabled.
//
// The function returns the following values:
//
//    - ok: TRUE if the popup includes a search entry.
//
func (self *DropDown) EnableSearch() bool {
	var _arg0 *C.GtkDropDown // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_down_get_enable_search(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Expression gets the expression set that is used to obtain strings from items.
//
// See gtk.DropDown.SetExpression().
//
// The function returns the following values:
//
//    - expression (optional): GtkExpression or NULL.
//
func (self *DropDown) Expression() Expressioner {
	var _arg0 *C.GtkDropDown   // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_down_get_expression(_arg0)
	runtime.KeepAlive(self)

	var _expression Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Expressioner)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// Factory gets the factory that's currently used to populate list items.
//
// The factory returned by this function is always used for the item in the
// button. It is also used for items in the popup if gtk.DropDown:list-factory
// is not set.
//
// The function returns the following values:
//
//    - listItemFactory (optional): factory in use.
//
func (self *DropDown) Factory() *ListItemFactory {
	var _arg0 *C.GtkDropDown        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_down_get_factory(_arg0)
	runtime.KeepAlive(self)

	var _listItemFactory *ListItemFactory // out

	if _cret != nil {
		_listItemFactory = wrapListItemFactory(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _listItemFactory
}

// ListFactory gets the factory that's currently used to populate list items in
// the popup.
//
// The function returns the following values:
//
//    - listItemFactory (optional): factory in use.
//
func (self *DropDown) ListFactory() *ListItemFactory {
	var _arg0 *C.GtkDropDown        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_down_get_list_factory(_arg0)
	runtime.KeepAlive(self)

	var _listItemFactory *ListItemFactory // out

	if _cret != nil {
		_listItemFactory = wrapListItemFactory(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _listItemFactory
}

// Model gets the model that provides the displayed items.
//
// The function returns the following values:
//
//    - listModel (optional): model in use.
//
func (self *DropDown) Model() gio.ListModeller {
	var _arg0 *C.GtkDropDown // out
	var _cret *C.GListModel  // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_down_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(gio.ListModeller)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gio.ListModeller")
			}
			_listModel = rv
		}
	}

	return _listModel
}

// Selected gets the position of the selected item.
//
// The function returns the following values:
//
//    - guint: position of the selected item, or GTK_INVALID_LIST_POSITION if not
//      item is selected.
//
func (self *DropDown) Selected() uint {
	var _arg0 *C.GtkDropDown // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_down_get_selected(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SelectedItem gets the selected item. If no item is selected, NULL is
// returned.
//
// The function returns the following values:
//
//    - object (optional): selected item.
//
func (self *DropDown) SelectedItem() *externglib.Object {
	var _arg0 *C.GtkDropDown // out
	var _cret C.gpointer     // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_down_get_selected_item(_arg0)
	runtime.KeepAlive(self)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// SetEnableSearch sets whether a search entry will be shown in the popup that
// allows to search for items in the list.
//
// Note that gtk.DropDown:expression must be set for search to work.
//
// The function takes the following parameters:
//
//    - enableSearch: whether to enable search.
//
func (self *DropDown) SetEnableSearch(enableSearch bool) {
	var _arg0 *C.GtkDropDown // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))
	if enableSearch {
		_arg1 = C.TRUE
	}

	C.gtk_drop_down_set_enable_search(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableSearch)
}

// SetExpression sets the expression that gets evaluated to obtain strings from
// items.
//
// This is used for search in the popup. The expression must have a value type
// of G_TYPE_STRING.
//
// The function takes the following parameters:
//
//    - expression (optional): GtkExpression, or NULL.
//
func (self *DropDown) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkDropDown   // out
	var _arg1 *C.GtkExpression // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
	}

	C.gtk_drop_down_set_expression(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetFactory sets the GtkListItemFactory to use for populating list items.
//
// The function takes the following parameters:
//
//    - factory (optional) to use or NULL for none.
//
func (self *DropDown) SetFactory(factory *ListItemFactory) {
	var _arg0 *C.GtkDropDown        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))
	if factory != nil {
		_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))
	}

	C.gtk_drop_down_set_factory(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(factory)
}

// SetListFactory sets the GtkListItemFactory to use for populating list items
// in the popup.
//
// The function takes the following parameters:
//
//    - factory (optional) to use or NULL for none.
//
func (self *DropDown) SetListFactory(factory *ListItemFactory) {
	var _arg0 *C.GtkDropDown        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))
	if factory != nil {
		_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))
	}

	C.gtk_drop_down_set_list_factory(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(factory)
}

// SetModel sets the GListModel to use.
//
// The function takes the following parameters:
//
//    - model (optional) to use or NULL for none.
//
func (self *DropDown) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkDropDown // out
	var _arg1 *C.GListModel  // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_drop_down_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetSelected selects the item at the given position.
//
// The function takes the following parameters:
//
//    - position of the item to select, or GTK_INVALID_LIST_POSITION.
//
func (self *DropDown) SetSelected(position uint) {
	var _arg0 *C.GtkDropDown // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	C.gtk_drop_down_set_selected(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}
