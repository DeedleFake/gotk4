// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// gint _gotk4_gtk3_AssistantPageFunc(gint, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_assistant_page_type_get_type()), F: marshalAssistantPageType},
		{T: externglib.Type(C.gtk_assistant_get_type()), F: marshalAssistanter},
	})
}

// AssistantPageType: enum for determining the page role inside the Assistant.
// It's used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// GTK_ASSISTANT_PAGE_CONFIRM, GTK_ASSISTANT_PAGE_SUMMARY or
// GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”. See
// gtk_assistant_commit() for details.
type AssistantPageType int

const (
	// AssistantPageContent: page has regular contents. Both the Back and
	// forward buttons will be shown.
	AssistantPageContent AssistantPageType = iota
	// AssistantPageIntro: page contains an introduction to the assistant task.
	// Only the Forward button will be shown if there is a next page.
	AssistantPageIntro
	// AssistantPageConfirm: page lets the user confirm or deny the changes. The
	// Back and Apply buttons will be shown.
	AssistantPageConfirm
	// AssistantPageSummary: page informs the user of the changes done. Only the
	// Close button will be shown.
	AssistantPageSummary
	// AssistantPageProgress: used for tasks that take a long time to complete,
	// blocks the assistant until the page is marked as complete. Only the back
	// button will be shown.
	AssistantPageProgress
	// AssistantPageCustom: used for when other page types are not appropriate.
	// No buttons will be shown, and the application must add its own buttons
	// through gtk_assistant_add_action_widget().
	AssistantPageCustom
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	return AssistantPageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for AssistantPageType.
func (a AssistantPageType) String() string {
	switch a {
	case AssistantPageContent:
		return "Content"
	case AssistantPageIntro:
		return "Intro"
	case AssistantPageConfirm:
		return "Confirm"
	case AssistantPageSummary:
		return "Summary"
	case AssistantPageProgress:
		return "Progress"
	case AssistantPageCustom:
		return "Custom"
	default:
		return fmt.Sprintf("AssistantPageType(%d)", a)
	}
}

// AssistantPageFunc: function used by gtk_assistant_set_forward_page_func() to
// know which is the next page given a current one. It’s called both for
// computing the next page when the user presses the “forward” button and for
// handling the behavior of the “last” button.
type AssistantPageFunc func(currentPage int) (gint int)

//export _gotk4_gtk3_AssistantPageFunc
func _gotk4_gtk3_AssistantPageFunc(arg0 C.gint, arg1 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var currentPage int // out

	currentPage = int(arg0)

	fn := v.(AssistantPageFunc)
	gint := fn(currentPage)

	cret = C.gint(gint)

	return cret
}

// AssistantOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AssistantOverrider interface {
	Apply()
	Cancel()
	Close()
	Prepare(page Widgetter)
}

// Assistant is a widget used to represent a generally complex operation
// splitted in several steps, guiding the user through its pages and controlling
// the page flow to collect the necessary data.
//
// The design of GtkAssistant is that it controls what buttons to show and to
// make sensitive, based on what it knows about the page sequence and the
// [type][GtkAssistantPageType] of each page, in addition to state information
// like the page [completion][gtk-assistant-set-page-complete] and
// [committed][gtk-assistant-commit] status.
//
// If you have a case that doesn’t quite fit in Assistants way of handling
// buttons, you can use the K_ASSISTANT_PAGE_CUSTOM page type and handle buttons
// yourself.
//
//
// GtkAssistant as GtkBuildable
//
// The GtkAssistant implementation of the Buildable interface exposes the
// action_area as internal children with the name “action_area”.
//
// To add pages to an assistant in Builder, simply add it as a child to the
// GtkAssistant object, and set its child properties as necessary.
//
//
// CSS nodes
//
// GtkAssistant has a single CSS node with the name assistant.
type Assistant struct {
	Window
}

func wrapAssistant(obj *externglib.Object) *Assistant {
	return &Assistant{
		Window: Window{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalAssistanter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAssistant(obj), nil
}

// NewAssistant creates a new Assistant.
func NewAssistant() *Assistant {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_assistant_new()

	var _assistant *Assistant // out

	_assistant = wrapAssistant(externglib.Take(unsafe.Pointer(_cret)))

	return _assistant
}

// AddActionWidget adds a widget to the action area of a Assistant.
func (assistant *Assistant) AddActionWidget(child Widgetter) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_assistant_add_action_widget(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(child)
}

// AppendPage appends a page to the assistant.
func (assistant *Assistant) AppendPage(page Widgetter) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_append_page(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Commit erases the visited page history so the back button is not shown on the
// current page, and removes the cancel button from subsequent pages.
//
// Use this when the information provided up to the current page is hereafter
// deemed permanent and cannot be modified or undone. For example, showing a
// progress page to track a long-running, unreversible operation after the user
// has clicked apply on a confirmation page.
func (assistant *Assistant) Commit() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	C.gtk_assistant_commit(_arg0)
	runtime.KeepAlive(assistant)
}

// CurrentPage returns the page number of the current page.
func (assistant *Assistant) CurrentPage() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	_cret = C.gtk_assistant_get_current_page(_arg0)
	runtime.KeepAlive(assistant)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NPages returns the number of pages in the assistant
func (assistant *Assistant) NPages() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	_cret = C.gtk_assistant_get_n_pages(_arg0)
	runtime.KeepAlive(assistant)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthPage returns the child widget contained in page number page_num.
func (assistant *Assistant) NthPage(pageNum int) Widgetter {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = C.gint(pageNum)

	_cret = C.gtk_assistant_get_nth_page(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// PageComplete gets whether page is complete.
func (assistant *Assistant) PageComplete(page Widgetter) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_complete(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageHasPadding gets whether page has padding.
func (assistant *Assistant) PageHasPadding(page Widgetter) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_has_padding(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageHeaderImage gets the header image for page.
//
// Deprecated: Since GTK+ 3.2, a header is no longer shown; add your header
// decoration to the page content instead.
func (assistant *Assistant) PageHeaderImage(page Widgetter) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_header_image(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// PageSideImage gets the side image for page.
//
// Deprecated: Since GTK+ 3.2, sidebar images are not shown anymore.
func (assistant *Assistant) PageSideImage(page Widgetter) *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_side_image(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// PageTitle gets the title for page.
func (assistant *Assistant) PageTitle(page Widgetter) string {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_title(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PageType gets the page type of page.
func (assistant *Assistant) PageType(page Widgetter) AssistantPageType {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _cret C.GtkAssistantPageType // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_type(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _assistantPageType AssistantPageType // out

	_assistantPageType = AssistantPageType(_cret)

	return _assistantPageType
}

// InsertPage inserts a page in the assistant at a given position.
func (assistant *Assistant) InsertPage(page Widgetter, position int) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = C.gint(position)

	_cret = C.gtk_assistant_insert_page(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(position)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NextPage: navigate to the next page.
//
// It is a programming error to call this function when there is no next page.
//
// This function is for use when creating pages of the K_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) NextPage() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	C.gtk_assistant_next_page(_arg0)
	runtime.KeepAlive(assistant)
}

// PrependPage prepends a page to the assistant.
func (assistant *Assistant) PrependPage(page Widgetter) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_prepend_page(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PreviousPage: navigate to the previous visited page.
//
// It is a programming error to call this function when no previous page is
// available.
//
// This function is for use when creating pages of the K_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) PreviousPage() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	C.gtk_assistant_previous_page(_arg0)
	runtime.KeepAlive(assistant)
}

// RemoveActionWidget removes a widget from the action area of a Assistant.
func (assistant *Assistant) RemoveActionWidget(child Widgetter) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_assistant_remove_action_widget(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(child)
}

// RemovePage removes the page_num’s page from assistant.
func (assistant *Assistant) RemovePage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_remove_page(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)
}

// SetCurrentPage switches the page to page_num.
//
// Note that this will only be necessary in custom buttons, as the assistant
// flow can be set with gtk_assistant_set_forward_page_func().
func (assistant *Assistant) SetCurrentPage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_set_current_page(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)
}

// SetForwardPageFunc sets the page forwarding function to be page_func.
//
// This function will be used to determine what will be the next page when the
// user presses the forward button. Setting page_func to NULL will make the
// assistant to use the default forward function, which just goes to the next
// visible page.
func (assistant *Assistant) SetForwardPageFunc(pageFunc AssistantPageFunc) {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 C.GtkAssistantPageFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	if pageFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_AssistantPageFunc)
		_arg2 = C.gpointer(gbox.Assign(pageFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_assistant_set_forward_page_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageFunc)
}

// SetPageComplete sets whether page contents are complete.
//
// This will make assistant update the buttons state to be able to continue the
// task.
func (assistant *Assistant) SetPageComplete(page Widgetter, complete bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if complete {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_complete(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(complete)
}

// SetPageHasPadding sets whether the assistant is adding padding around the
// page.
func (assistant *Assistant) SetPageHasPadding(page Widgetter, hasPadding bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if hasPadding {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_has_padding(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(hasPadding)
}

// SetPageHeaderImage sets a header image for page.
//
// Deprecated: Since GTK+ 3.2, a header is no longer shown; add your header
// decoration to the page content instead.
func (assistant *Assistant) SetPageHeaderImage(page Widgetter, pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if pixbuf != nil {
		_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	}

	C.gtk_assistant_set_page_header_image(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(pixbuf)
}

// SetPageSideImage sets a side image for page.
//
// This image used to be displayed in the side area of the assistant when page
// is the current page.
//
// Deprecated: Since GTK+ 3.2, sidebar images are not shown anymore.
func (assistant *Assistant) SetPageSideImage(page Widgetter, pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if pixbuf != nil {
		_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	}

	C.gtk_assistant_set_page_side_image(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(pixbuf)
}

// SetPageTitle sets a title for page.
//
// The title is displayed in the header area of the assistant when page is the
// current page.
func (assistant *Assistant) SetPageTitle(page Widgetter, title string) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_assistant_set_page_title(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(title)
}

// SetPageType sets the page type for page.
//
// The page type determines the page behavior in the assistant.
func (assistant *Assistant) SetPageType(page Widgetter, typ AssistantPageType) {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkAssistantPageType // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = C.GtkAssistantPageType(typ)

	C.gtk_assistant_set_page_type(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(typ)
}

// UpdateButtonsState forces assistant to recompute the buttons state.
//
// GTK+ automatically takes care of this in most situations, e.g. when the user
// goes to a different page, or when the visibility or completeness of a page
// changes.
//
// One situation where it can be necessary to call this function is when
// changing a value on the current page affects the future page flow of the
// assistant.
func (assistant *Assistant) UpdateButtonsState() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(assistant.Native()))

	C.gtk_assistant_update_buttons_state(_arg0)
	runtime.KeepAlive(assistant)
}
