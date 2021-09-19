// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_action_get_type()), F: marshalFileChooserAction},
		{T: externglib.Type(C.gtk_file_chooser_error_get_type()), F: marshalFileChooserError},
		{T: externglib.Type(C.gtk_file_chooser_get_type()), F: marshalFileChooserer},
	})
}

// FileChooserAction describes whether a GtkFileChooser is being used to open
// existing files or to save to a possibly new file.
type FileChooserAction int

const (
	// FileChooserActionOpen indicates open mode. The file chooser will only let
	// the user pick an existing file.
	FileChooserActionOpen FileChooserAction = iota
	// FileChooserActionSave indicates save mode. The file chooser will let the
	// user pick an existing file, or type in a new filename.
	FileChooserActionSave
	// FileChooserActionSelectFolder indicates an Open mode for selecting
	// folders. The file chooser will let the user pick an existing folder.
	FileChooserActionSelectFolder
)

func marshalFileChooserAction(p uintptr) (interface{}, error) {
	return FileChooserAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for FileChooserAction.
func (f FileChooserAction) String() string {
	switch f {
	case FileChooserActionOpen:
		return "Open"
	case FileChooserActionSave:
		return "Save"
	case FileChooserActionSelectFolder:
		return "SelectFolder"
	default:
		return fmt.Sprintf("FileChooserAction(%d)", f)
	}
}

// FileChooserError: these identify the various errors that can occur while
// calling GtkFileChooser functions.
type FileChooserError int

const (
	// FileChooserErrorNonexistent indicates that a file does not exist.
	FileChooserErrorNonexistent FileChooserError = iota
	// FileChooserErrorBadFilename indicates a malformed filename.
	FileChooserErrorBadFilename
	// FileChooserErrorAlreadyExists indicates a duplicate path (e.g. when
	// adding a bookmark).
	FileChooserErrorAlreadyExists
	// FileChooserErrorIncompleteHostname indicates an incomplete hostname (e.g.
	// "http://foo" without a slash after that).
	FileChooserErrorIncompleteHostname
)

func marshalFileChooserError(p uintptr) (interface{}, error) {
	return FileChooserError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for FileChooserError.
func (f FileChooserError) String() string {
	switch f {
	case FileChooserErrorNonexistent:
		return "Nonexistent"
	case FileChooserErrorBadFilename:
		return "BadFilename"
	case FileChooserErrorAlreadyExists:
		return "AlreadyExists"
	case FileChooserErrorIncompleteHostname:
		return "IncompleteHostname"
	default:
		return fmt.Sprintf("FileChooserError(%d)", f)
	}
}

// FileChooserErrorQuark registers an error quark for GtkFileChooser errors.
func FileChooserErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_file_chooser_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// FileChooser: GtkFileChooser is an interface that can be implemented by file
// selection widgets.
//
// In GTK, the main objects that implement this interface are
// gtk.FileChooserWidget and gtk.FileChooserDialog.
//
// You do not need to write an object that implements the GtkFileChooser
// interface unless you are trying to adapt an existing file selector to expose
// a standard programming interface.
//
// GtkFileChooser allows for shortcuts to various places in the filesystem. In
// the default implementation these are displayed in the left pane. It may be a
// bit confusing at first that these shortcuts come from various sources and in
// various flavours, so lets explain the terminology here:
//
// - Bookmarks: are created by the user, by dragging folders from the right pane
// to the left pane, or by using the “Add”. Bookmarks can be renamed and deleted
// by the user.
//
// - Shortcuts: can be provided by the application. For example, a Paint program
// may want to add a shortcut for a Clipart folder. Shortcuts cannot be modified
// by the user.
//
// - Volumes: are provided by the underlying filesystem abstraction. They are
// the “roots” of the filesystem.
//
//
// File Names and Encodings
//
// When the user is finished selecting files in a GtkFileChooser, your program
// can get the selected filenames as GFiles.
//
//
// Adding options
//
// You can add extra widgets to a file chooser to provide options that are not
// present in the default design, by using gtk.FileChooser.AddChoice(). Each
// choice has an identifier and a user visible label; additionally, each choice
// can have multiple options. If a choice has no option, it will be rendered as
// a check button with the given label; if a choice has options, it will be
// rendered as a combo box.
type FileChooser struct {
	*externglib.Object
}

// FileChooserer describes FileChooser's abstract methods.
type FileChooserer interface {
	externglib.Objector

	// AddChoice adds a 'choice' to the file chooser.
	AddChoice(id string, label string, options []string, optionLabels []string)
	// AddFilter adds filter to the list of filters that the user can select
	// between.
	AddFilter(filter *FileFilter)
	// AddShortcutFolder adds a folder to be displayed with the shortcut folders
	// in a file chooser.
	AddShortcutFolder(folder gio.Filer) error
	// Action gets the type of operation that the file chooser is performing.
	Action() FileChooserAction
	// Choice gets the currently selected option in the 'choice' with the given
	// ID.
	Choice(id string) string
	// CreateFolders gets whether file chooser will offer to create new folders.
	CreateFolders() bool
	// CurrentFolder gets the current folder of chooser as #GFile.
	CurrentFolder() gio.Filer
	// CurrentName gets the current name in the file selector, as entered by the
	// user.
	CurrentName() string
	// File gets the GFile for the currently selected file in the file selector.
	File() gio.Filer
	// Files lists all the selected files and subfolders in the current folder
	// of chooser as #GFile.
	Files() gio.ListModeller
	// Filter gets the current filter.
	Filter() *FileFilter
	// Filters gets the current set of user-selectable filters, as a list model.
	Filters() gio.ListModeller
	// SelectMultiple gets whether multiple files can be selected in the file
	// chooser.
	SelectMultiple() bool
	// ShortcutFolders queries the list of shortcut folders in the file chooser.
	ShortcutFolders() gio.ListModeller
	// RemoveChoice removes a 'choice' that has been added with
	// gtk_file_chooser_add_choice().
	RemoveChoice(id string)
	// RemoveFilter removes filter from the list of filters that the user can
	// select between.
	RemoveFilter(filter *FileFilter)
	// RemoveShortcutFolder removes a folder from the shortcut folders in a file
	// chooser.
	RemoveShortcutFolder(folder gio.Filer) error
	// SetAction sets the type of operation that the chooser is performing.
	SetAction(action FileChooserAction)
	// SetChoice selects an option in a 'choice' that has been added with
	// gtk_file_chooser_add_choice().
	SetChoice(id string, option string)
	// SetCreateFolders sets whether file chooser will offer to create new
	// folders.
	SetCreateFolders(createFolders bool)
	// SetCurrentFolder sets the current folder for chooser from a #GFile.
	SetCurrentFolder(file gio.Filer) error
	// SetCurrentName sets the current name in the file selector, as if entered
	// by the user.
	SetCurrentName(name string)
	// SetFile sets file as the current filename for the file chooser.
	SetFile(file gio.Filer) error
	// SetFilter sets the current filter.
	SetFilter(filter *FileFilter)
	// SetSelectMultiple sets whether multiple files can be selected in the file
	// chooser.
	SetSelectMultiple(selectMultiple bool)
}

var _ FileChooserer = (*FileChooser)(nil)

func wrapFileChooser(obj *externglib.Object) *FileChooser {
	return &FileChooser{
		Object: obj,
	}
}

func marshalFileChooserer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileChooser(obj), nil
}

// AddChoice adds a 'choice' to the file chooser.
//
// This is typically implemented as a combobox or, for boolean choices, as a
// checkbutton. You can select a value using gtk.FileChooser.SetChoice() before
// the dialog is shown, and you can obtain the user-selected value in the
// gtk.Dialog::response signal handler using gtk.FileChooser.GetChoice().
func (chooser *FileChooser) AddChoice(id string, label string, options []string, optionLabels []string) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out
	var _arg2 *C.char           // out
	var _arg3 **C.char          // out
	var _arg4 **C.char          // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))
	{
		_arg3 = (**C.char)(C.malloc(C.ulong(len(options)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(options)+1)
			var zero *C.char
			out[len(options)] = zero
			for i := range options {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(options[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	{
		_arg4 = (**C.char)(C.malloc(C.ulong(len(optionLabels)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg4))
		{
			out := unsafe.Slice(_arg4, len(optionLabels)+1)
			var zero *C.char
			out[len(optionLabels)] = zero
			for i := range optionLabels {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(optionLabels[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_file_chooser_add_choice(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(id)
	runtime.KeepAlive(label)
	runtime.KeepAlive(options)
	runtime.KeepAlive(optionLabels)
}

// AddFilter adds filter to the list of filters that the user can select
// between.
//
// When a filter is selected, only files that are passed by that filter are
// displayed.
//
// Note that the chooser takes ownership of the filter if it is floating, so you
// have to ref and sink it if you want to keep a reference.
func (chooser *FileChooser) AddFilter(filter *FileFilter) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GtkFileFilter  // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_file_chooser_add_filter(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// AddShortcutFolder adds a folder to be displayed with the shortcut folders in
// a file chooser.
func (chooser *FileChooser) AddShortcutFolder(folder gio.Filer) error {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GFile          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(folder.Native()))

	C.gtk_file_chooser_add_shortcut_folder(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(folder)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Action gets the type of operation that the file chooser is performing.
func (chooser *FileChooser) Action() FileChooserAction {
	var _arg0 *C.GtkFileChooser      // out
	var _cret C.GtkFileChooserAction // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_action(_arg0)
	runtime.KeepAlive(chooser)

	var _fileChooserAction FileChooserAction // out

	_fileChooserAction = FileChooserAction(_cret)

	return _fileChooserAction
}

// Choice gets the currently selected option in the 'choice' with the given ID.
func (chooser *FileChooser) Choice(id string) string {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_file_chooser_get_choice(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(id)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// CreateFolders gets whether file chooser will offer to create new folders.
func (chooser *FileChooser) CreateFolders() bool {
	var _arg0 *C.GtkFileChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_create_folders(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CurrentFolder gets the current folder of chooser as #GFile.
func (chooser *FileChooser) CurrentFolder() gio.Filer {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GFile          // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_current_folder(_arg0)
	runtime.KeepAlive(chooser)

	var _file gio.Filer // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(gio.Filer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Filer")
		}
		_file = rv
	}

	return _file
}

// CurrentName gets the current name in the file selector, as entered by the
// user.
//
// This is meant to be used in save dialogs, to get the currently typed filename
// when the file itself does not exist yet.
func (chooser *FileChooser) CurrentName() string {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_current_name(_arg0)
	runtime.KeepAlive(chooser)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// File gets the GFile for the currently selected file in the file selector.
//
// If multiple files are selected, one of the files will be returned at random.
//
// If the file chooser is in folder mode, this function returns the selected
// folder.
func (chooser *FileChooser) File() gio.Filer {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GFile          // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_file(_arg0)
	runtime.KeepAlive(chooser)

	var _file gio.Filer // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(gio.Filer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Filer")
		}
		_file = rv
	}

	return _file
}

// Files lists all the selected files and subfolders in the current folder of
// chooser as #GFile.
func (chooser *FileChooser) Files() gio.ListModeller {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GListModel     // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_files(_arg0)
	runtime.KeepAlive(chooser)

	var _listModel gio.ListModeller // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(gio.ListModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.ListModeller")
		}
		_listModel = rv
	}

	return _listModel
}

// Filter gets the current filter.
func (chooser *FileChooser) Filter() *FileFilter {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GtkFileFilter  // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_filter(_arg0)
	runtime.KeepAlive(chooser)

	var _fileFilter *FileFilter // out

	if _cret != nil {
		_fileFilter = wrapFileFilter(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _fileFilter
}

// Filters gets the current set of user-selectable filters, as a list model.
//
// See gtk.FileChooser.AddFilter() and gtk.FileChooser.RemoveFilter() for
// changing individual filters.
//
// You should not modify the returned list model. Future changes to chooser may
// or may not affect the returned model.
func (chooser *FileChooser) Filters() gio.ListModeller {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GListModel     // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_filters(_arg0)
	runtime.KeepAlive(chooser)

	var _listModel gio.ListModeller // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(gio.ListModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.ListModeller")
		}
		_listModel = rv
	}

	return _listModel
}

// SelectMultiple gets whether multiple files can be selected in the file
// chooser.
func (chooser *FileChooser) SelectMultiple() bool {
	var _arg0 *C.GtkFileChooser // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_select_multiple(_arg0)
	runtime.KeepAlive(chooser)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShortcutFolders queries the list of shortcut folders in the file chooser.
//
// You should not modify the returned list model. Future changes to chooser may
// or may not affect the returned model.
func (chooser *FileChooser) ShortcutFolders() gio.ListModeller {
	var _arg0 *C.GtkFileChooser // out
	var _cret *C.GListModel     // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))

	_cret = C.gtk_file_chooser_get_shortcut_folders(_arg0)
	runtime.KeepAlive(chooser)

	var _listModel gio.ListModeller // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(gio.ListModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.ListModeller")
		}
		_listModel = rv
	}

	return _listModel
}

// RemoveChoice removes a 'choice' that has been added with
// gtk_file_chooser_add_choice().
func (chooser *FileChooser) RemoveChoice(id string) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_remove_choice(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(id)
}

// RemoveFilter removes filter from the list of filters that the user can select
// between.
func (chooser *FileChooser) RemoveFilter(filter *FileFilter) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GtkFileFilter  // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_file_chooser_remove_filter(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// RemoveShortcutFolder removes a folder from the shortcut folders in a file
// chooser.
func (chooser *FileChooser) RemoveShortcutFolder(folder gio.Filer) error {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GFile          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(folder.Native()))

	C.gtk_file_chooser_remove_shortcut_folder(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(folder)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetAction sets the type of operation that the chooser is performing.
//
// The user interface is adapted to suit the selected action.
//
// For example, an option to create a new folder might be shown if the action is
// GTK_FILE_CHOOSER_ACTION_SAVE but not if the action is
// GTK_FILE_CHOOSER_ACTION_OPEN.
func (chooser *FileChooser) SetAction(action FileChooserAction) {
	var _arg0 *C.GtkFileChooser      // out
	var _arg1 C.GtkFileChooserAction // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = C.GtkFileChooserAction(action)

	C.gtk_file_chooser_set_action(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(action)
}

// SetChoice selects an option in a 'choice' that has been added with
// gtk_file_chooser_add_choice().
//
// For a boolean choice, the possible options are "true" and "false".
func (chooser *FileChooser) SetChoice(id string, option string) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out
	var _arg2 *C.char           // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(option)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_file_chooser_set_choice(_arg0, _arg1, _arg2)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(id)
	runtime.KeepAlive(option)
}

// SetCreateFolders sets whether file chooser will offer to create new folders.
//
// This is only relevant if the action is not set to be
// GTK_FILE_CHOOSER_ACTION_OPEN.
func (chooser *FileChooser) SetCreateFolders(createFolders bool) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	if createFolders {
		_arg1 = C.TRUE
	}

	C.gtk_file_chooser_set_create_folders(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(createFolders)
}

// SetCurrentFolder sets the current folder for chooser from a #GFile.
func (chooser *FileChooser) SetCurrentFolder(file gio.Filer) error {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GFile          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_file_chooser_set_current_folder(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(file)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetCurrentName sets the current name in the file selector, as if entered by
// the user.
//
// Note that the name passed in here is a UTF-8 string rather than a filename.
// This function is meant for such uses as a suggested name in a “Save As...”
// dialog. You can pass “Untitled.doc” or a similarly suitable suggestion for
// the name.
//
// If you want to preselect a particular existing file, you should use
// gtk.FileChooser.SetFile() instead.
//
// Please see the documentation for those functions for an example of using
// gtk.FileChooser.SetCurrentName() as well.
func (chooser *FileChooser) SetCurrentName(name string) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_set_current_name(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(name)
}

// SetFile sets file as the current filename for the file chooser.
//
// This includes changing to the file’s parent folder and actually selecting the
// file in list. If the chooser is in GTK_FILE_CHOOSER_ACTION_SAVE mode, the
// file’s base name will also appear in the dialog’s file name entry.
//
// If the file name isn’t in the current folder of chooser, then the current
// folder of chooser will be changed to the folder containing filename.
//
// Note that the file must exist, or nothing will be done except for the
// directory change.
//
// If you are implementing a save dialog, you should use this function if you
// already have a file name to which the user may save; for example, when the
// user opens an existing file and then does “Save As…”. If you don’t have a
// file name already — for example, if the user just created a new file and is
// saving it for the first time, do not call this function.
//
// Instead, use something similar to this:
//
//    static void
//    prepare_file_chooser (GtkFileChooser *chooser,
//                          GFile          *existing_file)
//    {
//      gboolean document_is_new = (existing_file == NULL);
//
//      if (document_is_new)
//        {
//          GFile *default_file_for_saving = g_file_new_for_path ("./out.txt");
//          // the user just created a new document
//          gtk_file_chooser_set_current_folder (chooser, default_file_for_saving, NULL);
//          gtk_file_chooser_set_current_name (chooser, "Untitled document");
//          g_object_unref (default_file_for_saving);
//        }
//      else
//        {
//          // the user edited an existing document
//          gtk_file_chooser_set_file (chooser, existing_file, NULL);
//        }
//    }
func (chooser *FileChooser) SetFile(file gio.Filer) error {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GFile          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_file_chooser_set_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(file)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetFilter sets the current filter.
//
// Only the files that pass the filter will be displayed. If the user-selectable
// list of filters is non-empty, then the filter should be one of the filters in
// that list.
//
// Setting the current filter when the list of filters is empty is useful if you
// want to restrict the displayed set of files without letting the user change
// it.
func (chooser *FileChooser) SetFilter(filter *FileFilter) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 *C.GtkFileFilter  // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	_arg1 = (*C.GtkFileFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_file_chooser_set_filter(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(filter)
}

// SetSelectMultiple sets whether multiple files can be selected in the file
// chooser.
//
// This is only relevant if the action is set to be GTK_FILE_CHOOSER_ACTION_OPEN
// or GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
func (chooser *FileChooser) SetSelectMultiple(selectMultiple bool) {
	var _arg0 *C.GtkFileChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFileChooser)(unsafe.Pointer(chooser.Native()))
	if selectMultiple {
		_arg1 = C.TRUE
	}

	C.gtk_file_chooser_set_select_multiple(_arg0, _arg1)
	runtime.KeepAlive(chooser)
	runtime.KeepAlive(selectMultiple)
}
