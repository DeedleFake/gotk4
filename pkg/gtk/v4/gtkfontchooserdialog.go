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
		{T: externglib.Type(C.gtk_font_chooser_dialog_get_type()), F: marshalFontChooserDialog},
	})
}

// FontChooserDialog: the `GtkFontChooserDialog` widget is a dialog for
// selecting a font.
//
// !An example GtkFontChooserDialog (fontchooser.png)
//
// `GtkFontChooserDialog` implements the [iface@Gtk.FontChooser] interface and
// does not provide much API of its own.
//
// To create a `GtkFontChooserDialog`, use [ctor@Gtk.FontChooserDialog.new].
//
//
// GtkFontChooserDialog as GtkBuildable
//
// The `GtkFontChooserDialog` implementation of the `GtkBuildable` interface
// exposes the buttons with the names “select_button” and “cancel_button”.
type FontChooserDialog interface {
	gextras.Objector

	privateFontChooserDialogClass()
}

// FontChooserDialogClass implements the FontChooserDialog interface.
type FontChooserDialogClass struct {
	*externglib.Object
	DialogClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
	FontChooserInterface
	NativeInterface
	RootInterface
	ShortcutManagerInterface
}

var _ FontChooserDialog = (*FontChooserDialogClass)(nil)

func wrapFontChooserDialog(obj *externglib.Object) FontChooserDialog {
	return &FontChooserDialogClass{
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
		FontChooserInterface: FontChooserInterface{
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

func marshalFontChooserDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontChooserDialog(obj), nil
}

// NewFontChooserDialog creates a new `GtkFontChooserDialog`.
func NewFontChooserDialog(title string, parent Window) *FontChooserDialogClass {
	var _arg1 *C.char      // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer((&Window).Native()))

	_cret = C.gtk_font_chooser_dialog_new(_arg1, _arg2)

	var _fontChooserDialog *FontChooserDialogClass // out

	_fontChooserDialog = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*FontChooserDialogClass)

	return _fontChooserDialog
}

func (*FontChooserDialogClass) privateFontChooserDialogClass() {}
