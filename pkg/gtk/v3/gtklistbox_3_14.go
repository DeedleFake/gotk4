// Code generated by girgen. DO NOT EDIT.

package gtk

import ()

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ListBoxForEachFunc: function used by gtk_list_box_selected_foreach(). It will
// be called on every selected child of the box.
type ListBoxForEachFunc func(box *ListBox, row *ListBoxRow)
