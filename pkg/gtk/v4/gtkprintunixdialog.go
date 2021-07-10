// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_unix_dialog_get_type()), F: marshalPrintUnixDialog},
	})
}

// PrintUnixDialog: `GtkPrintUnixDialog` implements a print dialog for platforms
// which don’t provide a native print dialog, like Unix.
//
// !An example GtkPrintUnixDialog (printdialog.png)
//
// It can be used very much like any other GTK dialog, at the cost of the
// portability offered by the high-level printing API with
// [class@Gtk.PrintOperation].
//
// In order to print something with `GtkPrintUnixDialog`, you need to use
// [method@Gtk.PrintUnixDialog.get_selected_printer] to obtain a
// [class@Gtk.Printer] object and use it to construct a [class@Gtk.PrintJob]
// using [ctor@Gtk.PrintJob.new].
//
// `GtkPrintUnixDialog` uses the following response values:
//
// - GTK_RESPONSE_OK: for the “Print” button - GTK_RESPONSE_APPLY: for the
// “Preview” button - GTK_RESPONSE_CANCEL: for the “Cancel” button
//
//
// GtkPrintUnixDialog as GtkBuildable
//
// The `GtkPrintUnixDialog` implementation of the `GtkBuildable` interface
// exposes its @notebook internal children with the name “notebook”.
//
// An example of a `GtkPrintUnixDialog` UI definition fragment:
//
// “`xml <object class="GtkPrintUnixDialog" id="dialog1"> <child
// internal-child="notebook"> <object class="GtkNotebook" id="notebook"> <child>
// <object type="GtkNotebookPage"> <property name="tab_expand">False</property>
// <property name="tab_fill">False</property> <property name="tab"> <object
// class="GtkLabel" id="tablabel"> <property name="label">Tab label</property>
// </object> </property> <property name="child"> <object class="GtkLabel"
// id="tabcontent"> <property name="label">Content on notebook tab</property>
// </object> </property> </object> </child> </object> </child> </object> “`
//
//
// CSS nodes
//
// `GtkPrintUnixDialog` has a single CSS node with name window. The style
// classes dialog and print are added.
type PrintUnixDialog interface {
	gextras.Objector

	// AddCustomTab adds a custom tab to the print dialog.
	AddCustomTab(child Widget, tabLabel Widget)
	// CurrentPage gets the current page of the `GtkPrintUnixDialog`.
	CurrentPage() int
	// EmbedPageSetup gets whether to embed the page setup.
	EmbedPageSetup() bool
	// HasSelection gets whether there is a selection.
	HasSelection() bool
	// ManualCapabilities gets the capabilities that have been set on this
	// `GtkPrintUnixDialog`.
	ManualCapabilities() PrintCapabilities
	// PageSetup gets the page setup that is used by the `GtkPrintUnixDialog`.
	PageSetup() *PageSetupClass
	// PageSetupSet gets whether a page setup was set by the user.
	PageSetupSet() bool
	// SelectedPrinter gets the currently selected printer.
	SelectedPrinter() *PrinterClass
	// Settings gets a new `GtkPrintSettings` object that represents the current
	// values in the print dialog.
	//
	// Note that this creates a new object, and you need to unref it if don’t
	// want to keep it.
	Settings() *PrintSettingsClass
	// SupportSelection gets whether the print dialog allows user to print a
	// selection.
	SupportSelection() bool
	// SetCurrentPage sets the current page number.
	//
	// If @current_page is not -1, this enables the current page choice for the
	// range of pages to print.
	SetCurrentPage(currentPage int)
	// SetEmbedPageSetup: embed page size combo box and orientation combo box
	// into page setup page.
	SetEmbedPageSetup(embed bool)
	// SetHasSelection sets whether a selection exists.
	SetHasSelection(hasSelection bool)
	// SetPageSetup sets the page setup of the `GtkPrintUnixDialog`.
	SetPageSetup(pageSetup PageSetup)
	// SetSettings sets the `GtkPrintSettings` for the `GtkPrintUnixDialog`.
	//
	// Typically, this is used to restore saved print settings from a previous
	// print operation before the print dialog is shown.
	SetSettings(settings PrintSettings)
	// SetSupportSelection sets whether the print dialog allows user to print a
	// selection.
	SetSupportSelection(supportSelection bool)
}

// PrintUnixDialogClass implements the PrintUnixDialog interface.
type PrintUnixDialogClass struct {
	*externglib.Object
	DialogClass
	AccessibleIface
	BuildableIface
	ConstraintTargetIface
	NativeIface
	RootIface
	ShortcutManagerIface
}

var _ PrintUnixDialog = (*PrintUnixDialogClass)(nil)

func wrapPrintUnixDialog(obj *externglib.Object) PrintUnixDialog {
	return &PrintUnixDialogClass{
		Object: obj,
		DialogClass: DialogClass{
			Object: obj,
			WindowClass: WindowClass{
				Object: obj,
				WidgetClass: WidgetClass{
					Object: obj,
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					AccessibleIface: AccessibleIface{
						Object: obj,
					},
					BuildableIface: BuildableIface{
						Object: obj,
					},
					ConstraintTargetIface: ConstraintTargetIface{
						Object: obj,
					},
				},
				AccessibleIface: AccessibleIface{
					Object: obj,
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
				ConstraintTargetIface: ConstraintTargetIface{
					Object: obj,
				},
				NativeIface: NativeIface{
					Object: obj,
					WidgetClass: WidgetClass{
						Object: obj,
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						AccessibleIface: AccessibleIface{
							Object: obj,
						},
						BuildableIface: BuildableIface{
							Object: obj,
						},
						ConstraintTargetIface: ConstraintTargetIface{
							Object: obj,
						},
					},
				},
				RootIface: RootIface{
					Object: obj,
					NativeIface: NativeIface{
						Object: obj,
						WidgetClass: WidgetClass{
							Object: obj,
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							AccessibleIface: AccessibleIface{
								Object: obj,
							},
							BuildableIface: BuildableIface{
								Object: obj,
							},
							ConstraintTargetIface: ConstraintTargetIface{
								Object: obj,
							},
						},
					},
					WidgetClass: WidgetClass{
						Object: obj,
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						AccessibleIface: AccessibleIface{
							Object: obj,
						},
						BuildableIface: BuildableIface{
							Object: obj,
						},
						ConstraintTargetIface: ConstraintTargetIface{
							Object: obj,
						},
					},
				},
				ShortcutManagerIface: ShortcutManagerIface{
					Object: obj,
				},
			},
			AccessibleIface: AccessibleIface{
				Object: obj,
			},
			BuildableIface: BuildableIface{
				Object: obj,
			},
			ConstraintTargetIface: ConstraintTargetIface{
				Object: obj,
			},
			NativeIface: NativeIface{
				Object: obj,
				WidgetClass: WidgetClass{
					Object: obj,
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					AccessibleIface: AccessibleIface{
						Object: obj,
					},
					BuildableIface: BuildableIface{
						Object: obj,
					},
					ConstraintTargetIface: ConstraintTargetIface{
						Object: obj,
					},
				},
			},
			RootIface: RootIface{
				Object: obj,
				NativeIface: NativeIface{
					Object: obj,
					WidgetClass: WidgetClass{
						Object: obj,
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						AccessibleIface: AccessibleIface{
							Object: obj,
						},
						BuildableIface: BuildableIface{
							Object: obj,
						},
						ConstraintTargetIface: ConstraintTargetIface{
							Object: obj,
						},
					},
				},
				WidgetClass: WidgetClass{
					Object: obj,
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					AccessibleIface: AccessibleIface{
						Object: obj,
					},
					BuildableIface: BuildableIface{
						Object: obj,
					},
					ConstraintTargetIface: ConstraintTargetIface{
						Object: obj,
					},
				},
			},
			ShortcutManagerIface: ShortcutManagerIface{
				Object: obj,
			},
		},
		AccessibleIface: AccessibleIface{
			Object: obj,
		},
		BuildableIface: BuildableIface{
			Object: obj,
		},
		ConstraintTargetIface: ConstraintTargetIface{
			Object: obj,
		},
		NativeIface: NativeIface{
			Object: obj,
			WidgetClass: WidgetClass{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				AccessibleIface: AccessibleIface{
					Object: obj,
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
				ConstraintTargetIface: ConstraintTargetIface{
					Object: obj,
				},
			},
		},
		RootIface: RootIface{
			Object: obj,
			NativeIface: NativeIface{
				Object: obj,
				WidgetClass: WidgetClass{
					Object: obj,
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					AccessibleIface: AccessibleIface{
						Object: obj,
					},
					BuildableIface: BuildableIface{
						Object: obj,
					},
					ConstraintTargetIface: ConstraintTargetIface{
						Object: obj,
					},
				},
			},
			WidgetClass: WidgetClass{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				AccessibleIface: AccessibleIface{
					Object: obj,
				},
				BuildableIface: BuildableIface{
					Object: obj,
				},
				ConstraintTargetIface: ConstraintTargetIface{
					Object: obj,
				},
			},
		},
		ShortcutManagerIface: ShortcutManagerIface{
			Object: obj,
		},
	}
}

func marshalPrintUnixDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPrintUnixDialog(obj), nil
}

// NewPrintUnixDialog creates a new `GtkPrintUnixDialog`.
func NewPrintUnixDialog(title string, parent Window) *PrintUnixDialogClass {
	var _arg1 *C.char      // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	_cret = C.gtk_print_unix_dialog_new(_arg1, _arg2)

	var _printUnixDialog *PrintUnixDialogClass // out

	_printUnixDialog = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*PrintUnixDialogClass)

	return _printUnixDialog
}

// AddCustomTab adds a custom tab to the print dialog.
func (dialog *PrintUnixDialogClass) AddCustomTab(child Widget, tabLabel Widget) {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _arg1 *C.GtkWidget          // out
	var _arg2 *C.GtkWidget          // out

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	C.gtk_print_unix_dialog_add_custom_tab(_arg0, _arg1, _arg2)
}

// CurrentPage gets the current page of the `GtkPrintUnixDialog`.
func (dialog *PrintUnixDialogClass) CurrentPage() int {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_current_page(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// EmbedPageSetup gets whether to embed the page setup.
func (dialog *PrintUnixDialogClass) EmbedPageSetup() bool {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_embed_page_setup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasSelection gets whether there is a selection.
func (dialog *PrintUnixDialogClass) HasSelection() bool {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_has_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ManualCapabilities gets the capabilities that have been set on this
// `GtkPrintUnixDialog`.
func (dialog *PrintUnixDialogClass) ManualCapabilities() PrintCapabilities {
	var _arg0 *C.GtkPrintUnixDialog  // out
	var _cret C.GtkPrintCapabilities // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_manual_capabilities(_arg0)

	var _printCapabilities PrintCapabilities // out

	_printCapabilities = (PrintCapabilities)(_cret)

	return _printCapabilities
}

// PageSetup gets the page setup that is used by the `GtkPrintUnixDialog`.
func (dialog *PrintUnixDialogClass) PageSetup() *PageSetupClass {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _cret *C.GtkPageSetup       // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_page_setup(_arg0)

	var _pageSetup *PageSetupClass // out

	_pageSetup = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*PageSetupClass)

	return _pageSetup
}

// PageSetupSet gets whether a page setup was set by the user.
func (dialog *PrintUnixDialogClass) PageSetupSet() bool {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_page_setup_set(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectedPrinter gets the currently selected printer.
func (dialog *PrintUnixDialogClass) SelectedPrinter() *PrinterClass {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _cret *C.GtkPrinter         // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_selected_printer(_arg0)

	var _printer *PrinterClass // out

	_printer = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*PrinterClass)

	return _printer
}

// Settings gets a new `GtkPrintSettings` object that represents the current
// values in the print dialog.
//
// Note that this creates a new object, and you need to unref it if don’t want
// to keep it.
func (dialog *PrintUnixDialogClass) Settings() *PrintSettingsClass {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _cret *C.GtkPrintSettings   // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_settings(_arg0)

	var _printSettings *PrintSettingsClass // out

	_printSettings = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PrintSettingsClass)

	return _printSettings
}

// SupportSelection gets whether the print dialog allows user to print a
// selection.
func (dialog *PrintUnixDialogClass) SupportSelection() bool {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_print_unix_dialog_get_support_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetCurrentPage sets the current page number.
//
// If @current_page is not -1, this enables the current page choice for the
// range of pages to print.
func (dialog *PrintUnixDialogClass) SetCurrentPage(currentPage int) {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.int(currentPage)

	C.gtk_print_unix_dialog_set_current_page(_arg0, _arg1)
}

// SetEmbedPageSetup: embed page size combo box and orientation combo box into
// page setup page.
func (dialog *PrintUnixDialogClass) SetEmbedPageSetup(embed bool) {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))
	if embed {
		_arg1 = C.TRUE
	}

	C.gtk_print_unix_dialog_set_embed_page_setup(_arg0, _arg1)
}

// SetHasSelection sets whether a selection exists.
func (dialog *PrintUnixDialogClass) SetHasSelection(hasSelection bool) {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))
	if hasSelection {
		_arg1 = C.TRUE
	}

	C.gtk_print_unix_dialog_set_has_selection(_arg0, _arg1)
}

// SetPageSetup sets the page setup of the `GtkPrintUnixDialog`.
func (dialog *PrintUnixDialogClass) SetPageSetup(pageSetup PageSetup) {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _arg1 *C.GtkPageSetup       // out

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.GtkPageSetup)(unsafe.Pointer(pageSetup.Native()))

	C.gtk_print_unix_dialog_set_page_setup(_arg0, _arg1)
}

// SetSettings sets the `GtkPrintSettings` for the `GtkPrintUnixDialog`.
//
// Typically, this is used to restore saved print settings from a previous print
// operation before the print dialog is shown.
func (dialog *PrintUnixDialogClass) SetSettings(settings PrintSettings) {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _arg1 *C.GtkPrintSettings   // out

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	C.gtk_print_unix_dialog_set_settings(_arg0, _arg1)
}

// SetSupportSelection sets whether the print dialog allows user to print a
// selection.
func (dialog *PrintUnixDialogClass) SetSupportSelection(supportSelection bool) {
	var _arg0 *C.GtkPrintUnixDialog // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkPrintUnixDialog)(unsafe.Pointer(dialog.Native()))
	if supportSelection {
		_arg1 = C.TRUE
	}

	C.gtk_print_unix_dialog_set_support_selection(_arg0, _arg1)
}
