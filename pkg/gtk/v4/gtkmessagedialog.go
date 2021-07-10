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
		{T: externglib.Type(C.gtk_buttons_type_get_type()), F: marshalButtonsType},
		{T: externglib.Type(C.gtk_message_dialog_get_type()), F: marshalMessageDialog},
	})
}

// ButtonsType: prebuilt sets of buttons for `GtkDialog`.
//
// If none of these choices are appropriate, simply use GTK_BUTTONS_NONE and
// call [method@Gtk.Dialog.add_buttons].
//
// > Please note that GTK_BUTTONS_OK, GTK_BUTTONS_YES_NO > and
// GTK_BUTTONS_OK_CANCEL are discouraged by the > GNOME Human Interface
// Guidelines (http://library.gnome.org/devel/hig-book/stable/).
type ButtonsType int

const (
	// None: no buttons at all
	ButtonsTypeNone ButtonsType = iota
	// Ok: OK button
	ButtonsTypeOk
	// Close: close button
	ButtonsTypeClose
	// Cancel: cancel button
	ButtonsTypeCancel
	// YesNo yes and No buttons
	ButtonsTypeYesNo
	// OkCancel: OK and Cancel buttons
	ButtonsTypeOkCancel
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	return ButtonsType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// MessageDialog: `GtkMessageDialog` presents a dialog with some message text.
//
// !An example GtkMessageDialog (messagedialog.png)
//
// It’s simply a convenience widget; you could construct the equivalent of
// `GtkMessageDialog` from `GtkDialog` without too much effort, but
// `GtkMessageDialog` saves typing.
//
// The easiest way to do a modal message dialog is to use the GTK_DIALOG_MODAL
// flag, which will call [method@Gtk.Window.set_modal] internally. The dialog
// will prevent interaction with the parent window until it's hidden or
// destroyed. You can use the [signal@Gtk.Dialog::response] signal to know when
// the user dismissed the dialog.
//
// An example for using a modal dialog: “`c GtkDialogFlags flags =
// GTK_DIALOG_DESTROY_WITH_PARENT | GTK_DIALOG_MODAL; dialog =
// gtk_message_dialog_new (parent_window, flags, GTK_MESSAGE_ERROR,
// GTK_BUTTONS_CLOSE, "Error reading “s”: s", filename, g_strerror (errno)); //
// Destroy the dialog when the user responds to it // (e.g. clicks a button)
//
// g_signal_connect (dialog, "response", G_CALLBACK (gtk_window_destroy), NULL);
// “`
//
// You might do a non-modal `GtkMessageDialog` simply by omitting the
// GTK_DIALOG_MODAL flag:
//
// “`c GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT; dialog =
// gtk_message_dialog_new (parent_window, flags, GTK_MESSAGE_ERROR,
// GTK_BUTTONS_CLOSE, "Error reading “s”: s", filename, g_strerror (errno));
//
// // Destroy the dialog when the user responds to it // (e.g. clicks a button)
// // g_signal_connect (dialog, "response", G_CALLBACK (gtk_window_destroy), NULL);
// // “`
//
//
// GtkMessageDialog as GtkBuildable
//
// The `GtkMessageDialog` implementation of the `GtkBuildable` interface exposes
// the message area as an internal child with the name “message_area”.
type MessageDialog interface {
	gextras.Objector

	// MessageArea returns the message area of the dialog.
	//
	// This is the box where the dialog’s primary and secondary labels are
	// packed. You can add your own extra content to that box and it will appear
	// below those labels. See [method@Gtk.Dialog.get_content_area] for the
	// corresponding function in the parent [class@Gtk.Dialog].
	MessageArea() *WidgetClass
	// SetMarkup sets the text of the message dialog.
	SetMarkup(str string)
}

// MessageDialogClass implements the MessageDialog interface.
type MessageDialogClass struct {
	*externglib.Object
	DialogClass
	AccessibleIface
	BuildableIface
	ConstraintTargetIface
	NativeIface
	RootIface
	ShortcutManagerIface
}

var _ MessageDialog = (*MessageDialogClass)(nil)

func wrapMessageDialog(obj *externglib.Object) MessageDialog {
	return &MessageDialogClass{
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

func marshalMessageDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMessageDialog(obj), nil
}

// MessageArea returns the message area of the dialog.
//
// This is the box where the dialog’s primary and secondary labels are packed.
// You can add your own extra content to that box and it will appear below those
// labels. See [method@Gtk.Dialog.get_content_area] for the corresponding
// function in the parent [class@Gtk.Dialog].
func (messageDialog *MessageDialogClass) MessageArea() *WidgetClass {
	var _arg0 *C.GtkMessageDialog // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(messageDialog.Native()))

	_cret = C.gtk_message_dialog_get_message_area(_arg0)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// SetMarkup sets the text of the message dialog.
func (messageDialog *MessageDialogClass) SetMarkup(str string) {
	var _arg0 *C.GtkMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(messageDialog.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_message_dialog_set_markup(_arg0, _arg1)
}
