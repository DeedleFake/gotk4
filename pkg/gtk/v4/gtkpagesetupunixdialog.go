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
		{T: externglib.Type(C.gtk_page_setup_unix_dialog_get_type()), F: marshalPageSetupUnixDialog},
	})
}

// PageSetupUnixDialog: `GtkPageSetupUnixDialog` implements a page setup dialog
// for platforms which don’t provide a native page setup dialog, like Unix.
//
// !An example GtkPageSetupUnixDialog (pagesetupdialog.png)
//
// It can be used very much like any other GTK dialog, at the cost of the
// portability offered by the high-level printing API in
// [class@Gtk.PrintOperation].
type PageSetupUnixDialog interface {
	gextras.Objector

	// PageSetup gets the currently selected page setup from the dialog.
	PageSetup() *PageSetupClass
	// PrintSettings gets the current print settings from the dialog.
	PrintSettings() *PrintSettingsClass
	// SetPageSetup sets the `GtkPageSetup` from which the page setup dialog
	// takes its values.
	SetPageSetup(pageSetup PageSetup)
	// SetPrintSettings sets the `GtkPrintSettings` from which the page setup
	// dialog takes its values.
	SetPrintSettings(printSettings PrintSettings)
}

// PageSetupUnixDialogClass implements the PageSetupUnixDialog interface.
type PageSetupUnixDialogClass struct {
	*externglib.Object
	DialogClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
	NativeInterface
	RootInterface
	ShortcutManagerInterface
}

var _ PageSetupUnixDialog = (*PageSetupUnixDialogClass)(nil)

func wrapPageSetupUnixDialog(obj *externglib.Object) PageSetupUnixDialog {
	return &PageSetupUnixDialogClass{
		Object: obj,
		DialogClass: DialogClass{
			Object: obj,
			WindowClass: WindowClass{
				Object: obj,
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					AccessibleInterface: AccessibleInterface{
						Object: obj,
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
					ConstraintTargetInterface: ConstraintTargetInterface{
						Object: obj,
					},
				},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
				NativeInterface: NativeInterface{
					WidgetClass: WidgetClass{
						InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
						AccessibleInterface: AccessibleInterface{
							Object: obj,
						},
						BuildableInterface: BuildableInterface{
							Object: obj,
						},
						ConstraintTargetInterface: ConstraintTargetInterface{
							Object: obj,
						},
					},
				},
				RootInterface: RootInterface{
					Object: obj,
					NativeInterface: NativeInterface{
						WidgetClass: WidgetClass{
							InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
							AccessibleInterface: AccessibleInterface{
								Object: obj,
							},
							BuildableInterface: BuildableInterface{
								Object: obj,
							},
							ConstraintTargetInterface: ConstraintTargetInterface{
								Object: obj,
							},
						},
					},
					WidgetClass: WidgetClass{
						InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
						AccessibleInterface: AccessibleInterface{
							Object: obj,
						},
						BuildableInterface: BuildableInterface{
							Object: obj,
						},
						ConstraintTargetInterface: ConstraintTargetInterface{
							Object: obj,
						},
					},
				},
				ShortcutManagerInterface: ShortcutManagerInterface{
					Object: obj,
				},
			},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
			NativeInterface: NativeInterface{
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					AccessibleInterface: AccessibleInterface{
						Object: obj,
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
					ConstraintTargetInterface: ConstraintTargetInterface{
						Object: obj,
					},
				},
			},
			RootInterface: RootInterface{
				Object: obj,
				NativeInterface: NativeInterface{
					WidgetClass: WidgetClass{
						InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
						AccessibleInterface: AccessibleInterface{
							Object: obj,
						},
						BuildableInterface: BuildableInterface{
							Object: obj,
						},
						ConstraintTargetInterface: ConstraintTargetInterface{
							Object: obj,
						},
					},
				},
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					AccessibleInterface: AccessibleInterface{
						Object: obj,
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
					ConstraintTargetInterface: ConstraintTargetInterface{
						Object: obj,
					},
				},
			},
			ShortcutManagerInterface: ShortcutManagerInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
		NativeInterface: NativeInterface{
			WidgetClass: WidgetClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
			},
		},
		RootInterface: RootInterface{
			Object: obj,
			NativeInterface: NativeInterface{
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					AccessibleInterface: AccessibleInterface{
						Object: obj,
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
					ConstraintTargetInterface: ConstraintTargetInterface{
						Object: obj,
					},
				},
			},
			WidgetClass: WidgetClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
			},
		},
		ShortcutManagerInterface: ShortcutManagerInterface{
			Object: obj,
		},
	}
}

func marshalPageSetupUnixDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPageSetupUnixDialog(obj), nil
}

// NewPageSetupUnixDialog creates a new page setup dialog.
func NewPageSetupUnixDialog(title string, parent Window) *PageSetupUnixDialogClass {
	var _arg1 *C.char      // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer((&Window).Native()))

	_cret = C.gtk_page_setup_unix_dialog_new(_arg1, _arg2)

	var _pageSetupUnixDialog *PageSetupUnixDialogClass // out

	_pageSetupUnixDialog = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*PageSetupUnixDialogClass)

	return _pageSetupUnixDialog
}

// PageSetup gets the currently selected page setup from the dialog.
func (d *PageSetupUnixDialogClass) PageSetup() *PageSetupClass {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _cret *C.GtkPageSetup           // in

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer((&PageSetupUnixDialog).Native()))

	_cret = C.gtk_page_setup_unix_dialog_get_page_setup(_arg0)

	var _pageSetup *PageSetupClass // out

	_pageSetup = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*PageSetupClass)

	return _pageSetup
}

// PrintSettings gets the current print settings from the dialog.
func (d *PageSetupUnixDialogClass) PrintSettings() *PrintSettingsClass {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _cret *C.GtkPrintSettings       // in

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer((&PageSetupUnixDialog).Native()))

	_cret = C.gtk_page_setup_unix_dialog_get_print_settings(_arg0)

	var _printSettings *PrintSettingsClass // out

	_printSettings = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*PrintSettingsClass)

	return _printSettings
}

// SetPageSetup sets the `GtkPageSetup` from which the page setup dialog takes
// its values.
func (d *PageSetupUnixDialogClass) SetPageSetup(pageSetup PageSetup) {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _arg1 *C.GtkPageSetup           // out

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer((&PageSetupUnixDialog).Native()))
	_arg1 = (*C.GtkPageSetup)(unsafe.Pointer((&PageSetup).Native()))

	C.gtk_page_setup_unix_dialog_set_page_setup(_arg0, _arg1)
}

// SetPrintSettings sets the `GtkPrintSettings` from which the page setup dialog
// takes its values.
func (d *PageSetupUnixDialogClass) SetPrintSettings(printSettings PrintSettings) {
	var _arg0 *C.GtkPageSetupUnixDialog // out
	var _arg1 *C.GtkPrintSettings       // out

	_arg0 = (*C.GtkPageSetupUnixDialog)(unsafe.Pointer((&PageSetupUnixDialog).Native()))
	_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer((&PrintSettings).Native()))

	C.gtk_page_setup_unix_dialog_set_print_settings(_arg0, _arg1)
}
