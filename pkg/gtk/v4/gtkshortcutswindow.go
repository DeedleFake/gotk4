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
		{T: externglib.Type(C.gtk_shortcuts_window_get_type()), F: marshalShortcutsWindow},
	})
}

// ShortcutsWindow: `GtkShortcutsWindow` shows information about the keyboard
// shortcuts and gestures of an application.
//
// The shortcuts can be grouped, and you can have multiple sections in this
// window, corresponding to the major modes of your application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a `GtkShortcutsWindow` is with
// [class@Gtk.Builder], by populating a `GtkShortcutsWindow` with one or more
// `GtkShortcutsSection` objects, which contain `GtkShortcutsGroups` that in
// turn contain objects of class `GtkShortcutsShortcut`.
//
// A simple example:
//
// ! (gedit-shortcuts.png)
//
// This example has as single section. As you can see, the shortcut groups are
// arranged in columns, and spread across several pages if there are too many to
// find on a single page.
//
// The .ui file for this example can be found here
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-gedit.ui).
//
// An example with multiple views:
//
// ! (clocks-shortcuts.png)
//
// This example shows a `GtkShortcutsWindow` that has been configured to show
// only the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found here
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-clocks.ui).
//
// An example with multiple sections:
//
// ! (builder-shortcuts.png)
//
// This example shows a `GtkShortcutsWindow` with two sections, "Editor
// Shortcuts" and "Terminal Shortcuts".
//
// The .ui file for this example can be found here
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-builder.ui).
type ShortcutsWindow interface {
	gextras.Objector

	privateShortcutsWindowClass()
}

// ShortcutsWindowClass implements the ShortcutsWindow interface.
type ShortcutsWindowClass struct {
	*externglib.Object
	WindowClass
	AccessibleIface
	BuildableIface
	ConstraintTargetIface
	NativeIface
	RootIface
	ShortcutManagerIface
}

var _ ShortcutsWindow = (*ShortcutsWindowClass)(nil)

func wrapShortcutsWindow(obj *externglib.Object) ShortcutsWindow {
	return &ShortcutsWindowClass{
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
	}
}

func marshalShortcutsWindow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapShortcutsWindow(obj), nil
}

func (*ShortcutsWindowClass) privateShortcutsWindowClass() {}
