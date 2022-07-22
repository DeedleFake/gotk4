// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_FileChooserWidget_ConnectDesktopFolder(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectDownFolder(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectHomeFolder(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectLocationPopup(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectLocationPopupOnPaste(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectLocationTogglePopup(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectPlacesShortcut(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectQuickBookmark(gpointer, gint, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectRecentShortcut(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectSearchShortcut(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectShowHidden(gpointer, guintptr);
// extern void _gotk4_gtk4_FileChooserWidget_ConnectUpFolder(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeFileChooserWidget = coreglib.Type(C.gtk_file_chooser_widget_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileChooserWidget, F: marshalFileChooserWidget},
	})
}

// FileChooserWidget: GtkFileChooserWidget is a widget for choosing files.
//
// It exposes the gtk.FileChooser interface, and you should use the methods of
// this interface to interact with the widget.
//
//
// CSS nodes
//
// GtkFileChooserWidget has a single CSS node with name filechooser.
type FileChooserWidget struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	FileChooser
}

var (
	_ Widgetter         = (*FileChooserWidget)(nil)
	_ coreglib.Objector = (*FileChooserWidget)(nil)
)

func wrapFileChooserWidget(obj *coreglib.Object) *FileChooserWidget {
	return &FileChooserWidget{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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
		Object: obj,
		FileChooser: FileChooser{
			Object: obj,
		},
	}
}

func marshalFileChooserWidget(p uintptr) (interface{}, error) {
	return wrapFileChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_FileChooserWidget_ConnectDesktopFolder
func _gotk4_gtk4_FileChooserWidget_ConnectDesktopFolder(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDesktopFolder is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show the user's Desktop folder in the
// file list.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>D</kbd>.
func (v *FileChooserWidget) ConnectDesktopFolder(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "desktop-folder", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectDesktopFolder), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectDownFolder
func _gotk4_gtk4_FileChooserWidget_ConnectDownFolder(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDownFolder is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser go to a child of the current folder in
// the file hierarchy. The subfolder that will be used is displayed in the path
// bar widget of the file chooser. For example, if the path bar is showing
// "/foo/bar/baz", with bar currently displayed, then this will cause the file
// chooser to switch to the "baz" subfolder.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>Down</kbd>.
func (v *FileChooserWidget) ConnectDownFolder(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "down-folder", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectDownFolder), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectHomeFolder
func _gotk4_gtk4_FileChooserWidget_ConnectHomeFolder(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectHomeFolder is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show the user's home folder in the file
// list.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>Home</kbd>.
func (v *FileChooserWidget) ConnectHomeFolder(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "home-folder", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectHomeFolder), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectLocationPopup
func _gotk4_gtk4_FileChooserWidget_ConnectLocationPopup(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(path string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(path string))
	}

	var _path string // out

	_path = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_path)
}

// ConnectLocationPopup is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show a "Location" prompt which the user
// can use to manually type the name of the file he wishes to select.
//
// The default bindings for this signal are <kbd>Control</kbd>-<kbd>L</kbd> with
// a path string of "" (the empty string). It is also bound to <kbd>/</kbd> with
// a path string of "/" (a slash): this lets you type / and immediately type a
// path name. On Unix systems, this is bound to <kbd>~</kbd> (tilde) with a path
// string of "~" itself for access to home directories.
func (v *FileChooserWidget) ConnectLocationPopup(f func(path string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "location-popup", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectLocationPopup), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectLocationPopupOnPaste
func _gotk4_gtk4_FileChooserWidget_ConnectLocationPopupOnPaste(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectLocationPopupOnPaste is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show a "Location" prompt when the user
// pastes into a GtkFileChooserWidget.
//
// The default binding for this signal is <kbd>Control</kbd>-<kbd>V</kbd>.
func (v *FileChooserWidget) ConnectLocationPopupOnPaste(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "location-popup-on-paste", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectLocationPopupOnPaste), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectLocationTogglePopup
func _gotk4_gtk4_FileChooserWidget_ConnectLocationTogglePopup(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectLocationTogglePopup is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to toggle the visibility of a "Location" prompt which the user
// can use to manually type the name of the file he wishes to select.
//
// The default binding for this signal is <kbd>Control</kbd>-<kbd>L</kbd>.
func (v *FileChooserWidget) ConnectLocationTogglePopup(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "location-toggle-popup", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectLocationTogglePopup), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectPlacesShortcut
func _gotk4_gtk4_FileChooserWidget_ConnectPlacesShortcut(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectPlacesShortcut is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to move the focus to the places sidebar.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>P</kbd>.
func (v *FileChooserWidget) ConnectPlacesShortcut(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "places-shortcut", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectPlacesShortcut), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectQuickBookmark
func _gotk4_gtk4_FileChooserWidget_ConnectQuickBookmark(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(bookmarkIndex int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(bookmarkIndex int))
	}

	var _bookmarkIndex int // out

	_bookmarkIndex = int(arg1)

	f(_bookmarkIndex)
}

// ConnectQuickBookmark is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser switch to the bookmark specified in the
// bookmark_index parameter. For example, if you have three bookmarks, you can
// pass 0, 1, 2 to this signal to switch to each of them, respectively.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>1</kbd>,
// <kbd>Alt</kbd>-<kbd>2</kbd>, etc. until <kbd>Alt</kbd>-<kbd>0</kbd>. Note
// that in the default binding, that <kbd>Alt</kbd>-<kbd>1</kbd> is actually
// defined to switch to the bookmark at index 0, and so on successively.
func (v *FileChooserWidget) ConnectQuickBookmark(f func(bookmarkIndex int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "quick-bookmark", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectQuickBookmark), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectRecentShortcut
func _gotk4_gtk4_FileChooserWidget_ConnectRecentShortcut(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectRecentShortcut is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show the Recent location.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>R</kbd>.
func (v *FileChooserWidget) ConnectRecentShortcut(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "recent-shortcut", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectRecentShortcut), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectSearchShortcut
func _gotk4_gtk4_FileChooserWidget_ConnectSearchShortcut(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectSearchShortcut is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser show the search entry.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>S</kbd>.
func (v *FileChooserWidget) ConnectSearchShortcut(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "search-shortcut", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectSearchShortcut), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectShowHidden
func _gotk4_gtk4_FileChooserWidget_ConnectShowHidden(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectShowHidden is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser display hidden files.
//
// The default binding for this signal is <kbd>Control</kbd>-<kbd>H</kbd>.
func (v *FileChooserWidget) ConnectShowHidden(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "show-hidden", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectShowHidden), f)
}

//export _gotk4_gtk4_FileChooserWidget_ConnectUpFolder
func _gotk4_gtk4_FileChooserWidget_ConnectUpFolder(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectUpFolder is emitted when the user asks for it.
//
// This is a keybinding signal (class.SignalAction.html).
//
// This is used to make the file chooser go to the parent of the current folder
// in the file hierarchy.
//
// The default binding for this signal is <kbd>Alt</kbd>-<kbd>Up</kbd>.
func (v *FileChooserWidget) ConnectUpFolder(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "up-folder", false, unsafe.Pointer(C._gotk4_gtk4_FileChooserWidget_ConnectUpFolder), f)
}

// NewFileChooserWidget creates a new GtkFileChooserWidget.
//
// This is a file chooser widget that can be embedded in custom windows, and it
// is the same widget that is used by GtkFileChooserDialog.
//
// The function takes the following parameters:
//
//    - action: open or save mode for the widget.
//
// The function returns the following values:
//
//    - fileChooserWidget: new GtkFileChooserWidget.
//
func NewFileChooserWidget(action FileChooserAction) *FileChooserWidget {
	var _arg1 C.GtkFileChooserAction // out
	var _cret *C.GtkWidget           // in

	_arg1 = C.GtkFileChooserAction(action)

	_cret = C.gtk_file_chooser_widget_new(_arg1)
	runtime.KeepAlive(action)

	var _fileChooserWidget *FileChooserWidget // out

	_fileChooserWidget = wrapFileChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _fileChooserWidget
}
