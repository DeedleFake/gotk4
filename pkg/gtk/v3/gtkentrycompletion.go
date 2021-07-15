// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk3_EntryCompletionMatchFunc(GtkEntryCompletion*, gchar*, GtkTreeIter*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_entry_completion_get_type()), F: marshalEntryCompletioner},
	})
}

// EntryCompletionMatchFunc: function which decides whether the row indicated by
// iter matches a given key, and should be displayed as a possible completion
// for key. Note that key is normalized and case-folded (see g_utf8_normalize()
// and g_utf8_casefold()). If this is not appropriate, match functions have
// access to the unmodified key via gtk_entry_get_text (GTK_ENTRY
// (gtk_entry_completion_get_entry ())).
type EntryCompletionMatchFunc func(completion *EntryCompletion, key string, iter *TreeIter) (ok bool)

//export _gotk4_gtk3_EntryCompletionMatchFunc
func _gotk4_gtk3_EntryCompletionMatchFunc(arg0 *C.GtkEntryCompletion, arg1 *C.gchar, arg2 *C.GtkTreeIter, arg3 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var completion *EntryCompletion // out
	var key string                  // out
	var iter *TreeIter              // out

	completion = wrapEntryCompletion(externglib.Take(unsafe.Pointer(arg0)))
	key = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	defer C.free(unsafe.Pointer(arg1))
	iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	runtime.SetFinalizer(iter, func(v *TreeIter) {
		C.gtk_tree_iter_free((*C.GtkTreeIter)(gextras.StructNative(unsafe.Pointer(v))))
	})

	fn := v.(EntryCompletionMatchFunc)
	ok := fn(completion, key, iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// EntryCompletionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type EntryCompletionOverrider interface {
	ActionActivated(index_ int)
	CursorOnMatch(model TreeModeller, iter *TreeIter) bool
	InsertPrefix(prefix string) bool
	MatchSelected(model TreeModeller, iter *TreeIter) bool
	NoMatches()
}

// EntryCompletion is an auxiliary object to be used in conjunction with Entry
// to provide the completion functionality. It implements the CellLayout
// interface, to allow the user to add extra cells to the TreeView with
// completion matches.
//
// “Completion functionality” means that when the user modifies the text in the
// entry, EntryCompletion checks which rows in the model match the current
// content of the entry, and displays a list of matches. By default, the
// matching is done by comparing the entry text case-insensitively against the
// text column of the model (see gtk_entry_completion_set_text_column()), but
// this can be overridden with a custom match function (see
// gtk_entry_completion_set_match_func()).
//
// When the user selects a completion, the content of the entry is updated. By
// default, the content of the entry is replaced by the text column of the
// model, but this can be overridden by connecting to the
// EntryCompletion::match-selected signal and updating the entry in the signal
// handler. Note that you should return TRUE from the signal handler to suppress
// the default behaviour.
//
// To add completion functionality to an entry, use gtk_entry_set_completion().
//
// In addition to regular completion matches, which will be inserted into the
// entry when they are selected, EntryCompletion also allows to display
// “actions” in the popup window. Their appearance is similar to menuitems, to
// differentiate them clearly from completion strings. When an action is
// selected, the EntryCompletion::action-activated signal is emitted.
//
// GtkEntryCompletion uses a TreeModelFilter model to represent the subset of
// the entire model that is currently matching. While the GtkEntryCompletion
// signals EntryCompletion::match-selected and EntryCompletion::cursor-on-match
// take the original model and an iter pointing to that model as arguments,
// other callbacks and signals (such as CellLayoutDataFuncs or
// CellArea::apply-attributes) will generally take the filter model as argument.
// As long as you are only calling gtk_tree_model_get(), this will make no
// difference to you. If for some reason, you need the original model, use
// gtk_tree_model_filter_get_model(). Don’t forget to use
// gtk_tree_model_filter_convert_iter_to_child_iter() to obtain a matching iter.
type EntryCompletion struct {
	*externglib.Object

	Buildable
	CellLayout
}

var _ gextras.Nativer = (*EntryCompletion)(nil)

func wrapEntryCompletion(obj *externglib.Object) *EntryCompletion {
	return &EntryCompletion{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
		CellLayout: CellLayout{
			Object: obj,
		},
	}
}

func marshalEntryCompletioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEntryCompletion(obj), nil
}

// NewEntryCompletion creates a new EntryCompletion object.
func NewEntryCompletion() *EntryCompletion {
	var _cret *C.GtkEntryCompletion // in

	_cret = C.gtk_entry_completion_new()

	var _entryCompletion *EntryCompletion // out

	_entryCompletion = wrapEntryCompletion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryCompletion
}

// NewEntryCompletionWithArea creates a new EntryCompletion object using the
// specified area to layout cells in the underlying TreeViewColumn for the
// drop-down menu.
func NewEntryCompletionWithArea(area CellAreaer) *EntryCompletion {
	var _arg1 *C.GtkCellArea        // out
	var _cret *C.GtkEntryCompletion // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer((area).(gextras.Nativer).Native()))

	_cret = C.gtk_entry_completion_new_with_area(_arg1)

	var _entryCompletion *EntryCompletion // out

	_entryCompletion = wrapEntryCompletion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryCompletion
}

// Complete requests a completion operation, or in other words a refiltering of
// the current list with completions, using the current key. The completion list
// view will be updated accordingly.
func (completion *EntryCompletion) Complete() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	C.gtk_entry_completion_complete(_arg0)
}

// ComputePrefix computes the common prefix that is shared by all rows in
// completion that start with key. If no row matches key, NULL will be returned.
// Note that a text column must have been set for this function to work, see
// gtk_entry_completion_set_text_column() for details.
func (completion *EntryCompletion) ComputePrefix(key string) string {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.char               // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))

	_cret = C.gtk_entry_completion_compute_prefix(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DeleteAction deletes the action at index_ from completion’s action list.
//
// Note that index_ is a relative position and the position of an action may
// have changed since it was inserted.
func (completion *EntryCompletion) DeleteAction(index_ int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = C.gint(index_)

	C.gtk_entry_completion_delete_action(_arg0, _arg1)
}

// CompletionPrefix: get the original text entered by the user that triggered
// the completion or NULL if there’s no completion ongoing.
func (completion *EntryCompletion) CompletionPrefix() string {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_completion_prefix(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Entry gets the entry completion has been attached to.
func (completion *EntryCompletion) Entry() *Widget {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_entry(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// InlineCompletion returns whether the common prefix of the possible
// completions should be automatically inserted in the entry.
func (completion *EntryCompletion) InlineCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_inline_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InlineSelection returns TRUE if inline-selection mode is turned on.
func (completion *EntryCompletion) InlineSelection() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_inline_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MinimumKeyLength returns the minimum key length as set for completion.
func (completion *EntryCompletion) MinimumKeyLength() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gint                // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_minimum_key_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Model returns the model the EntryCompletion is using as data source. Returns
// NULL if the model is unset.
func (completion *EntryCompletion) Model() *TreeModel {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkTreeModel       // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_model(_arg0)

	var _treeModel *TreeModel // out

	_treeModel = wrapTreeModel(externglib.Take(unsafe.Pointer(_cret)))

	return _treeModel
}

// PopupCompletion returns whether the completions should be presented in a
// popup window.
func (completion *EntryCompletion) PopupCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_popup_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSetWidth returns whether the completion popup window will be resized to
// the width of the entry.
func (completion *EntryCompletion) PopupSetWidth() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_popup_set_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSingleMatch returns whether the completion popup window will appear even
// if there is only a single match.
func (completion *EntryCompletion) PopupSingleMatch() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_popup_single_match(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TextColumn returns the column in the model of completion to get strings from.
func (completion *EntryCompletion) TextColumn() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gint                // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_text_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InsertActionMarkup inserts an action in completion’s action item list at
// position index_ with markup markup.
func (completion *EntryCompletion) InsertActionMarkup(index_ int, markup string) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out
	var _arg2 *C.gchar              // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = C.gint(index_)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))

	C.gtk_entry_completion_insert_action_markup(_arg0, _arg1, _arg2)
}

// InsertActionText inserts an action in completion’s action item list at
// position index_ with text text. If you want the action item to have markup,
// use gtk_entry_completion_insert_action_markup().
//
// Note that index_ is a relative position in the list of actions and the
// position of an action can change when deleting a different action.
func (completion *EntryCompletion) InsertActionText(index_ int, text string) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out
	var _arg2 *C.gchar              // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = C.gint(index_)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))

	C.gtk_entry_completion_insert_action_text(_arg0, _arg1, _arg2)
}

// InsertPrefix requests a prefix insertion.
func (completion *EntryCompletion) InsertPrefix() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	C.gtk_entry_completion_insert_prefix(_arg0)
}

// SetInlineCompletion sets whether the common prefix of the possible
// completions should be automatically inserted in the entry.
func (completion *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if inlineCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_completion(_arg0, _arg1)
}

// SetInlineSelection sets whether it is possible to cycle through the possible
// completions inside the entry.
func (completion *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if inlineSelection {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_selection(_arg0, _arg1)
}

// SetMatchFunc sets the match function for completion to be func. The match
// function is used to determine if a row should or should not be in the
// completion list.
func (completion *EntryCompletion) SetMatchFunc(fn EntryCompletionMatchFunc) {
	var _arg0 *C.GtkEntryCompletion         // out
	var _arg1 C.GtkEntryCompletionMatchFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_EntryCompletionMatchFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_entry_completion_set_match_func(_arg0, _arg1, _arg2, _arg3)
}

// SetMinimumKeyLength requires the length of the search key for completion to
// be at least length. This is useful for long lists, where completing using a
// small key takes a lot of time and will come up with meaningless results
// anyway (ie, a too large dataset).
func (completion *EntryCompletion) SetMinimumKeyLength(length int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = C.gint(length)

	C.gtk_entry_completion_set_minimum_key_length(_arg0, _arg1)
}

// SetModel sets the model for a EntryCompletion. If completion already has a
// model set, it will remove it before setting the new model. If model is NULL,
// then it will unset the model.
func (completion *EntryCompletion) SetModel(model TreeModeller) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.GtkTreeModel       // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer((model).(gextras.Nativer).Native()))

	C.gtk_entry_completion_set_model(_arg0, _arg1)
}

// SetPopupCompletion sets whether the completions should be presented in a
// popup window.
func (completion *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if popupCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_completion(_arg0, _arg1)
}

// SetPopupSetWidth sets whether the completion popup window will be resized to
// be the same width as the entry.
func (completion *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if popupSetWidth {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_set_width(_arg0, _arg1)
}

// SetPopupSingleMatch sets whether the completion popup window will appear even
// if there is only a single match. You may want to set this to FALSE if you are
// using [inline completion][GtkEntryCompletion--inline-completion].
func (completion *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if popupSingleMatch {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_single_match(_arg0, _arg1)
}

// SetTextColumn: convenience function for setting up the most used case of this
// code: a completion list with just strings. This function will set up
// completion to have a list displaying all (and just) strings in the completion
// list, and to get those strings from column in the model of completion.
//
// This functions creates and adds a CellRendererText for the selected column.
// If you need to set the text column, but don't want the cell renderer, use
// g_object_set() to set the EntryCompletion:text-column property directly.
func (completion *EntryCompletion) SetTextColumn(column int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gint                // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = C.gint(column)

	C.gtk_entry_completion_set_text_column(_arg0, _arg1)
}
