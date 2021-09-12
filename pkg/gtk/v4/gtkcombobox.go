// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_combo_box_get_type()), F: marshalComboBoxer},
	})
}

// ComboBoxOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ComboBoxOverrider interface {
	Changed()
	FormatEntryText(path string) string
}

// ComboBox: GtkComboBox is a widget that allows the user to choose from a list
// of valid choices.
//
// !An example GtkComboBox (combo-box.png)
//
// The GtkComboBox displays the selected choice; when activated, the GtkComboBox
// displays a popup which allows the user to make a new choice.
//
// The GtkComboBox uses the model-view pattern; the list of valid choices is
// specified in the form of a tree model, and the display of the choices can be
// adapted to the data in the model by using cell renderers, as you would in a
// tree view. This is possible since GtkComboBox implements the gtk.CellLayout
// interface. The tree model holding the valid choices is not restricted to a
// flat list, it can be a real tree, and the popup will reflect the tree
// structure.
//
// To allow the user to enter values not in the model, the
// gtk.ComboBox:has-entry property allows the GtkComboBox to contain a
// gtk.Entry. This entry can be accessed by calling gtk.ComboBox.GetChild() on
// the combo box.
//
// For a simple list of textual choices, the model-view API of GtkComboBox can
// be a bit overwhelming. In this case, gtk.ComboBoxText offers a simple
// alternative. Both GtkComboBox and GtkComboBoxText can contain an entry.
//
// CSS nodes
//
//    combobox
//    ├── box.linked
//    │   ╰── button.combo
//    │       ╰── box
//    │           ├── cellview
//    │           ╰── arrow
//    ╰── window.popup
//
//
// A normal combobox contains a box with the .linked class, a button with the
// .combo class and inside those buttons, there are a cellview and an arrow.
//
//    combobox
//    ├── box.linked
//    │   ├── entry.combo
//    │   ╰── button.combo
//    │       ╰── box
//    │           ╰── arrow
//    ╰── window.popup
//
//
// A GtkComboBox with an entry has a single CSS node with name combobox. It
// contains a box with the .linked class. That box contains an entry and a
// button, both with the .combo class added. The button also contains another
// node with name arrow.
//
//
// Accessibility
//
// GtkComboBox uses the GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type ComboBox struct {
	Widget

	CellEditable
	CellLayout
	*externglib.Object
}

func wrapComboBox(obj *externglib.Object) *ComboBox {
	return &ComboBox{
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
		CellEditable: CellEditable{
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
		},
		CellLayout: CellLayout{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalComboBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapComboBox(obj), nil
}

// NewComboBox creates a new empty GtkComboBox.
func NewComboBox() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new()

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithEntry creates a new empty GtkComboBox with an entry.
func NewComboBoxWithEntry() *ComboBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_new_with_entry()

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModel creates a new GtkComboBox with a model.
func NewComboBoxWithModel(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_combo_box_new_with_model(_arg1)
	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// NewComboBoxWithModelAndEntry creates a new empty GtkComboBox with an entry
// and a model.
func NewComboBoxWithModelAndEntry(model TreeModeller) *ComboBox {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_combo_box_new_with_model_and_entry(_arg1)
	runtime.KeepAlive(model)

	var _comboBox *ComboBox // out

	_comboBox = wrapComboBox(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBox
}

// Active returns the index of the currently active item.
//
// If the model is a non-flat treemodel, and the active item is not an immediate
// child of the root of the tree, this function returns
// gtk_tree_path_get_indices (path)[0], where path is the gtk.TreePath of the
// active item.
func (comboBox *ComboBox) Active() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ActiveID returns the ID of the active row of combo_box.
//
// This value is taken from the active row and the column specified by the
// gtk.ComboBox:id-column property of combo_box (see
// gtk.ComboBox.SetIDColumn()).
//
// The returned value is an interned string which means that you can compare the
// pointer by value to other interned strings and that you must not free it.
//
// If the gtk.ComboBox:id-column property of combo_box is not set, or if no row
// is active, or if the active row has a NULL ID value, then NULL is returned.
func (comboBox *ComboBox) ActiveID() string {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active_id(_arg0)
	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ActiveIter sets iter to point to the currently active item.
//
// If no item is active, iter is left unchanged.
func (comboBox *ComboBox) ActiveIter() (TreeIter, bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.GtkTreeIter  // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_active_iter(_arg0, &_arg1)
	runtime.KeepAlive(comboBox)

	var _iter TreeIter // out
	var _ok bool       // out

	_iter = *(*TreeIter)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _iter, _ok
}

// ButtonSensitivity returns whether the combo box sets the dropdown button
// sensitive or not when there are no items in the model.
func (comboBox *ComboBox) ButtonSensitivity() SensitivityType {
	var _arg0 *C.GtkComboBox       // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_button_sensitivity(_arg0)
	runtime.KeepAlive(comboBox)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

// Child gets the child widget of combo_box.
func (comboBox *ComboBox) Child() Widgetter {
	var _arg0 *C.GtkComboBox // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_child(_arg0)
	runtime.KeepAlive(comboBox)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// EntryTextColumn returns the column which combo_box is using to get the
// strings from to display in the internal entry.
func (comboBox *ComboBox) EntryTextColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_entry_text_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// HasEntry returns whether the combo box has an entry.
func (comboBox *ComboBox) HasEntry() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_has_entry(_arg0)
	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IDColumn returns the column which combo_box is using to get string IDs for
// values from.
func (comboBox *ComboBox) IDColumn() int {
	var _arg0 *C.GtkComboBox // out
	var _cret C.int          // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_id_column(_arg0)
	runtime.KeepAlive(comboBox)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Model returns the GtkTreeModel of combo_box.
func (comboBox *ComboBox) Model() TreeModeller {
	var _arg0 *C.GtkComboBox  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_model(_arg0)
	runtime.KeepAlive(comboBox)

	var _treeModel TreeModeller // out

	if _cret != nil {
		_treeModel = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(TreeModeller)
	}

	return _treeModel
}

// PopupFixedWidth gets whether the popup uses a fixed width.
func (comboBox *ComboBox) PopupFixedWidth() bool {
	var _arg0 *C.GtkComboBox // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_get_popup_fixed_width(_arg0)
	runtime.KeepAlive(comboBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown hides the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
func (comboBox *ComboBox) Popdown() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	C.gtk_combo_box_popdown(_arg0)
	runtime.KeepAlive(comboBox)
}

// Popup pops up the menu or dropdown list of combo_box.
//
// This function is mostly intended for use by accessibility technologies;
// applications should have little use for it.
//
// Before calling this, combo_box must be mapped, or nothing will happen.
func (comboBox *ComboBox) Popup() {
	var _arg0 *C.GtkComboBox // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))

	C.gtk_combo_box_popup(_arg0)
	runtime.KeepAlive(comboBox)
}

// PopupForDevice pops up the menu of combo_box.
//
// Note that currently this does not do anything with the device, as it was
// previously only used for list-mode combo boxes, and those were removed in GTK
// 4. However, it is retained in case similar functionality is added back later.
func (comboBox *ComboBox) PopupForDevice(device gdk.Devicer) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GdkDevice   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gtk_combo_box_popup_for_device(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(device)
}

// SetActive sets the active item of combo_box to be the item at index.
func (comboBox *ComboBox) SetActive(index_ int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(index_)

	C.gtk_combo_box_set_active(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(index_)
}

// SetActiveID changes the active row of combo_box to the one that has an ID
// equal to active_id.
//
// If active_id is NULL, the active row is unset. Rows having a NULL ID string
// cannot be made active by this function.
//
// If the gtk.ComboBox:id-column property of combo_box is unset or if no row has
// the given ID then the function does nothing and returns FALSE.
func (comboBox *ComboBox) SetActiveID(activeId string) bool {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if activeId != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(activeId)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_combo_box_set_active_id(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(activeId)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActiveIter sets the current active item to be the one referenced by iter.
//
// If iter is NULL, the active item is unset.
func (comboBox *ComboBox) SetActiveIter(iter *TreeIter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkTreeIter // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if iter != nil {
		_arg1 = (*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(iter)))
	}

	C.gtk_combo_box_set_active_iter(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(iter)
}

// SetButtonSensitivity sets whether the dropdown button of the combo box should
// update its sensitivity depending on the model contents.
func (comboBox *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	var _arg0 *C.GtkComboBox       // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_combo_box_set_button_sensitivity(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(sensitivity)
}

// SetChild sets the child widget of combo_box.
func (comboBox *ComboBox) SetChild(child Widgetter) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_combo_box_set_child(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(child)
}

// SetEntryTextColumn sets the model column which combo_box should use to get
// strings from to be text_column.
//
// The column text_column in the model of combo_box must be of type
// G_TYPE_STRING.
//
// This is only relevant if combo_box has been created with
// gtk.ComboBox:has-entry as TRUE.
func (comboBox *ComboBox) SetEntryTextColumn(textColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(textColumn)

	C.gtk_combo_box_set_entry_text_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(textColumn)
}

// SetIDColumn sets the model column which combo_box should use to get string
// IDs for values from.
//
// The column id_column in the model of combo_box must be of type G_TYPE_STRING.
func (comboBox *ComboBox) SetIDColumn(idColumn int) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(idColumn)

	C.gtk_combo_box_set_id_column(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(idColumn)
}

// SetModel sets the model used by combo_box to be model.
//
// Will unset a previously set model (if applicable). If model is NULL, then it
// will unset the model.
//
// Note that this function does not clear the cell renderers, you have to call
// gtk.CellLayout.Clear() yourself if you need to set up different cell
// renderers for the new model.
func (comboBox *ComboBox) SetModel(model TreeModeller) {
	var _arg0 *C.GtkComboBox  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if model != nil {
		_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_combo_box_set_model(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(model)
}

// SetPopupFixedWidth specifies whether the popup’s width should be a fixed
// width.
//
// If fixed is TRUE, the popup's width is set to match the allocated width of
// the combo box.
func (comboBox *ComboBox) SetPopupFixedWidth(fixed bool) {
	var _arg0 *C.GtkComboBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if fixed {
		_arg1 = C.TRUE
	}

	C.gtk_combo_box_set_popup_fixed_width(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fixed)
}

// SetRowSeparatorFunc sets the row separator function, which is used to
// determine whether a row should be drawn as a separator.
//
// If the row separator function is NULL, no separators are drawn. This is the
// default value.
func (comboBox *ComboBox) SetRowSeparatorFunc(fn TreeViewRowSeparatorFunc) {
	var _arg0 *C.GtkComboBox                // out
	var _arg1 C.GtkTreeViewRowSeparatorFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkComboBox)(unsafe.Pointer(comboBox.Native()))
	if fn != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_TreeViewRowSeparatorFunc)
		_arg2 = C.gpointer(gbox.Assign(fn))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_combo_box_set_row_separator_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(fn)
}
