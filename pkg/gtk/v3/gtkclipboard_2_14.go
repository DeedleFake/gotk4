// Code generated by girgen. DO NOT EDIT.

package gtk

import ()

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ClipboardURIReceivedFunc: function to be called when the results of
// gtk_clipboard_request_uris() are received, or when the request fails.
type ClipboardURIReceivedFunc func(clipboard *Clipboard, uris []string)
