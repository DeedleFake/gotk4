// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeRecentChooserDialog = coreglib.Type(C.gtk_recent_chooser_dialog_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRecentChooserDialog, F: marshalRecentChooserDialog},
	})
}

// RecentChooserDialogOverrider contains methods that are overridable.
type RecentChooserDialogOverrider interface {
}

// RecentChooserDialog is a dialog box suitable for displaying the recently used
// documents. This widgets works by putting a RecentChooserWidget inside a
// Dialog. It exposes the RecentChooserIface interface, so you can use all the
// RecentChooser functions on the recent chooser dialog as well as those for
// Dialog.
//
// Note that RecentChooserDialog does not have any methods of its own. Instead,
// you should use the functions that work on a RecentChooser.
//
//
// Typical usage
//
// In the simplest of cases, you can use the following code to use a
// RecentChooserDialog to select a recently used file:
//
//    GtkWidget *dialog;
//    gint res;
//
//    dialog = gtk_recent_chooser_dialog_new ("Recent Documents",
//                                            parent_window,
//                                            _("_Cancel"),
//                                            GTK_RESPONSE_CANCEL,
//                                            _("_Open"),
//                                            GTK_RESPONSE_ACCEPT,
//                                            NULL);
//
//    res = gtk_dialog_run (GTK_DIALOG (dialog));
//    if (res == GTK_RESPONSE_ACCEPT)
//      {
//        GtkRecentInfo *info;
//        GtkRecentChooser *chooser = GTK_RECENT_CHOOSER (dialog);
//
//        info = gtk_recent_chooser_get_current_item (chooser);
//        open_file (gtk_recent_info_get_uri (info));
//        gtk_recent_info_unref (info);
//      }
//
//    gtk_widget_destroy (dialog);
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooserDialog struct {
	_ [0]func() // equal guard
	Dialog

	*coreglib.Object
	RecentChooser
}

var (
	_ coreglib.Objector = (*RecentChooserDialog)(nil)
	_ Binner            = (*RecentChooserDialog)(nil)
)

func classInitRecentChooserDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRecentChooserDialog(obj *coreglib.Object) *RecentChooserDialog {
	return &RecentChooserDialog{
		Dialog: Dialog{
			Window: Window{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: coreglib.InitiallyUnowned{
								Object: obj,
							},
							Object: obj,
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
						},
					},
				},
			},
		},
		Object: obj,
		RecentChooser: RecentChooser{
			Object: obj,
		},
	}
}

func marshalRecentChooserDialog(p uintptr) (interface{}, error) {
	return wrapRecentChooserDialog(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
