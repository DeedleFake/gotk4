// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ClipboardImageReceivedFunc: function to be called when the results of
// gtk_clipboard_request_image() are received, or when the request fails.
type ClipboardImageReceivedFunc func(clipboard *Clipboard, pixbuf *gdkpixbuf.Pixbuf)
