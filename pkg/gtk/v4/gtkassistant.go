// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
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
		{T: externglib.Type(C.gtk_assistant_page_type_get_type()), F: marshalAssistantPageType},
		{T: externglib.Type(C.gtk_assistant_get_type()), F: marshalAssistant},
		{T: externglib.Type(C.gtk_assistant_page_get_type()), F: marshalAssistantPage},
	})
}

// AssistantPageType determines the page role inside a `GtkAssistant`.
//
// The role is used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// GTK_ASSISTANT_PAGE_CONFIRM, GTK_ASSISTANT_PAGE_SUMMARY or
// GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”. See
// gtk_assistant_commit() for details.
type AssistantPageType int

const (
	// Content: the page has regular contents. Both the Back and forward buttons
	// will be shown.
	AssistantPageTypeContent AssistantPageType = iota
	// Intro: the page contains an introduction to the assistant task. Only the
	// Forward button will be shown if there is a next page.
	AssistantPageTypeIntro
	// Confirm: the page lets the user confirm or deny the changes. The Back and
	// Apply buttons will be shown.
	AssistantPageTypeConfirm
	// Summary: the page informs the user of the changes done. Only the Close
	// button will be shown.
	AssistantPageTypeSummary
	// Progress: used for tasks that take a long time to complete, blocks the
	// assistant until the page is marked as complete. Only the back button will
	// be shown.
	AssistantPageTypeProgress
	// Custom: used for when other page types are not appropriate. No buttons
	// will be shown, and the application must add its own buttons through
	// gtk_assistant_add_action_widget().
	AssistantPageTypeCustom
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	return AssistantPageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AssistantPageFunc: type of callback used to calculate the next page in a
// `GtkAssistant`.
//
// It’s called both for computing the next page when the user presses the
// “forward” button and for handling the behavior of the “last” button.
//
// See [method@Gtk.Assistant.set_forward_page_func].
type AssistantPageFunc func(currentPage int, data interface{}) (gint int)

//export gotk4_AssistantPageFunc
func gotk4_AssistantPageFunc(arg0 C.int, arg1 C.gpointer) (cret C.int) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var currentPage int  // out
	var data interface{} // out

	currentPage = int(arg0)
	data = box.Get(uintptr(arg1))

	fn := v.(AssistantPageFunc)
	gint := fn(currentPage, data)

	cret = C.int(gint)

	return cret
}

// Assistant: `GtkAssistant` is used to represent a complex as a series of
// steps.
//
// !An example GtkAssistant (assistant.png)
//
// Each step consists of one or more pages. `GtkAssistant` guides the user
// through the pages, and controls the page flow to collect the data needed for
// the operation.
//
// `GtkAssistant` handles which buttons to show and to make sensitive based on
// page sequence knowledge and the [enum@Gtk.AssistantPageType] of each page in
// addition to state information like the *completed* and *committed* page
// statuses.
//
// If you have a case that doesn’t quite fit in `GtkAssistant`s way of handling
// buttons, you can use the GTK_ASSISTANT_PAGE_CUSTOM page type and handle
// buttons yourself.
//
// `GtkAssistant` maintains a `GtkAssistantPage` object for each added child,
// which holds additional per-child properties. You obtain the
// `GtkAssistantPage` for a child with [method@Gtk.Assistant.get_page].
//
//
// GtkAssistant as GtkBuildable
//
// The `GtkAssistant` implementation of the `GtkBuildable` interface exposes the
// @action_area as internal children with the name “action_area”.
//
// To add pages to an assistant in `GtkBuilder`, simply add it as a child to the
// `GtkAssistant` object. If you need to set per-object properties, create a
// `GtkAssistantPage` object explicitly, and set the child widget as a property
// on it.
//
//
// CSS nodes
//
// `GtkAssistant` has a single CSS node with the name window and style class
// .assistant.
type Assistant interface {
	gextras.Objector

	// AddActionWidget adds a widget to the action area of a `GtkAssistant`.
	AddActionWidget(child Widget)
	// AppendPage appends a page to the @assistant.
	AppendPage(page Widget) int
	// Commit erases the visited page history.
	//
	// GTK will then hide the back button on the current page, and removes the
	// cancel button from subsequent pages.
	//
	// Use this when the information provided up to the current page is
	// hereafter deemed permanent and cannot be modified or undone. For example,
	// showing a progress page to track a long-running, unreversible operation
	// after the user has clicked apply on a confirmation page.
	Commit()
	// CurrentPage returns the page number of the current page.
	CurrentPage() int
	// NPages returns the number of pages in the @assistant
	NPages() int
	// NthPage returns the child widget contained in page number @page_num.
	NthPage(pageNum int) *WidgetClass
	// Page returns the `GtkAssistantPage` object for @child.
	Page(child Widget) *AssistantPageClass
	// PageComplete gets whether @page is complete.
	PageComplete(page Widget) bool
	// PageTitle gets the title for @page.
	PageTitle(page Widget) string
	// PageType gets the page type of @page.
	PageType(page Widget) AssistantPageType
	// InsertPage inserts a page in the @assistant at a given position.
	InsertPage(page Widget, position int) int
	// NextPage: navigate to the next page.
	//
	// It is a programming error to call this function when there is no next
	// page.
	//
	// This function is for use when creating pages of the
	// GTK_ASSISTANT_PAGE_CUSTOM type.
	NextPage()
	// PrependPage prepends a page to the @assistant.
	PrependPage(page Widget) int
	// PreviousPage: navigate to the previous visited page.
	//
	// It is a programming error to call this function when no previous page is
	// available.
	//
	// This function is for use when creating pages of the
	// GTK_ASSISTANT_PAGE_CUSTOM type.
	PreviousPage()
	// RemoveActionWidget removes a widget from the action area of a
	// `GtkAssistant`.
	RemoveActionWidget(child Widget)
	// RemovePage removes the @page_num’s page from @assistant.
	RemovePage(pageNum int)
	// SetCurrentPage switches the page to @page_num.
	//
	// Note that this will only be necessary in custom buttons, as the
	// @assistant flow can be set with gtk_assistant_set_forward_page_func().
	SetCurrentPage(pageNum int)
	// SetPageComplete sets whether @page contents are complete.
	//
	// This will make @assistant update the buttons state to be able to continue
	// the task.
	SetPageComplete(page Widget, complete bool)
	// SetPageTitle sets a title for @page.
	//
	// The title is displayed in the header area of the assistant when @page is
	// the current page.
	SetPageTitle(page Widget, title string)
	// UpdateButtonsState forces @assistant to recompute the buttons state.
	//
	// GTK automatically takes care of this in most situations, e.g. when the
	// user goes to a different page, or when the visibility or completeness of
	// a page changes.
	//
	// One situation where it can be necessary to call this function is when
	// changing a value on the current page affects the future page flow of the
	// assistant.
	UpdateButtonsState()
}

// AssistantClass implements the Assistant interface.
type AssistantClass struct {
	*externglib.Object
	WindowClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
	NativeInterface
	RootInterface
	ShortcutManagerInterface
}

var _ Assistant = (*AssistantClass)(nil)

func wrapAssistant(obj *externglib.Object) Assistant {
	return &AssistantClass{
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
	}
}

func marshalAssistant(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAssistant(obj), nil
}

// NewAssistant creates a new `GtkAssistant`.
func NewAssistant() *AssistantClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_assistant_new()

	var _assistant *AssistantClass // out

	_assistant = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*AssistantClass)

	return _assistant
}

// AddActionWidget adds a widget to the action area of a `GtkAssistant`.
func (a *AssistantClass) AddActionWidget(child Widget) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	C.gtk_assistant_add_action_widget(_arg0, _arg1)
}

// AppendPage appends a page to the @assistant.
func (a *AssistantClass) AppendPage(page Widget) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.int           // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	_cret = C.gtk_assistant_append_page(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Commit erases the visited page history.
//
// GTK will then hide the back button on the current page, and removes the
// cancel button from subsequent pages.
//
// Use this when the information provided up to the current page is hereafter
// deemed permanent and cannot be modified or undone. For example, showing a
// progress page to track a long-running, unreversible operation after the user
// has clicked apply on a confirmation page.
func (a *AssistantClass) Commit() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))

	C.gtk_assistant_commit(_arg0)
}

// CurrentPage returns the page number of the current page.
func (a *AssistantClass) CurrentPage() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.int           // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))

	_cret = C.gtk_assistant_get_current_page(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NPages returns the number of pages in the @assistant
func (a *AssistantClass) NPages() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.int           // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))

	_cret = C.gtk_assistant_get_n_pages(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthPage returns the child widget contained in page number @page_num.
func (a *AssistantClass) NthPage(pageNum int) *WidgetClass {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.int           // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = C.int(pageNum)

	_cret = C.gtk_assistant_get_nth_page(_arg0, _arg1)

	var _widget *WidgetClass // out

	_widget = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*WidgetClass)

	return _widget
}

// Page returns the `GtkAssistantPage` object for @child.
func (a *AssistantClass) Page(child Widget) *AssistantPageClass {
	var _arg0 *C.GtkAssistant     // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.GtkAssistantPage // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	_cret = C.gtk_assistant_get_page(_arg0, _arg1)

	var _assistantPage *AssistantPageClass // out

	_assistantPage = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*AssistantPageClass)

	return _assistantPage
}

// PageComplete gets whether @page is complete.
func (a *AssistantClass) PageComplete(page Widget) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	_cret = C.gtk_assistant_get_page_complete(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageTitle gets the title for @page.
func (a *AssistantClass) PageTitle(page Widget) string {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	_cret = C.gtk_assistant_get_page_title(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PageType gets the page type of @page.
func (a *AssistantClass) PageType(page Widget) AssistantPageType {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _cret C.GtkAssistantPageType // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	_cret = C.gtk_assistant_get_page_type(_arg0, _arg1)

	var _assistantPageType AssistantPageType // out

	_assistantPageType = (AssistantPageType)(C.GtkAssistantPageType)

	return _assistantPageType
}

// InsertPage inserts a page in the @assistant at a given position.
func (a *AssistantClass) InsertPage(page Widget, position int) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.int           // out
	var _cret C.int           // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))
	_arg2 = C.int(position)

	_cret = C.gtk_assistant_insert_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NextPage: navigate to the next page.
//
// It is a programming error to call this function when there is no next page.
//
// This function is for use when creating pages of the GTK_ASSISTANT_PAGE_CUSTOM
// type.
func (a *AssistantClass) NextPage() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))

	C.gtk_assistant_next_page(_arg0)
}

// PrependPage prepends a page to the @assistant.
func (a *AssistantClass) PrependPage(page Widget) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.int           // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	_cret = C.gtk_assistant_prepend_page(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PreviousPage: navigate to the previous visited page.
//
// It is a programming error to call this function when no previous page is
// available.
//
// This function is for use when creating pages of the GTK_ASSISTANT_PAGE_CUSTOM
// type.
func (a *AssistantClass) PreviousPage() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))

	C.gtk_assistant_previous_page(_arg0)
}

// RemoveActionWidget removes a widget from the action area of a `GtkAssistant`.
func (a *AssistantClass) RemoveActionWidget(child Widget) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))

	C.gtk_assistant_remove_action_widget(_arg0, _arg1)
}

// RemovePage removes the @page_num’s page from @assistant.
func (a *AssistantClass) RemovePage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.int           // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = C.int(pageNum)

	C.gtk_assistant_remove_page(_arg0, _arg1)
}

// SetCurrentPage switches the page to @page_num.
//
// Note that this will only be necessary in custom buttons, as the @assistant
// flow can be set with gtk_assistant_set_forward_page_func().
func (a *AssistantClass) SetCurrentPage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.int           // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = C.int(pageNum)

	C.gtk_assistant_set_current_page(_arg0, _arg1)
}

// SetPageComplete sets whether @page contents are complete.
//
// This will make @assistant update the buttons state to be able to continue the
// task.
func (a *AssistantClass) SetPageComplete(page Widget, complete bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))
	if complete {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_complete(_arg0, _arg1, _arg2)
}

// SetPageTitle sets a title for @page.
//
// The title is displayed in the header area of the assistant when @page is the
// current page.
func (a *AssistantClass) SetPageTitle(page Widget, title string) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.char         // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((&Widget).Native()))
	_arg2 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_assistant_set_page_title(_arg0, _arg1, _arg2)
}

// UpdateButtonsState forces @assistant to recompute the buttons state.
//
// GTK automatically takes care of this in most situations, e.g. when the user
// goes to a different page, or when the visibility or completeness of a page
// changes.
//
// One situation where it can be necessary to call this function is when
// changing a value on the current page affects the future page flow of the
// assistant.
func (a *AssistantClass) UpdateButtonsState() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer((&Assistant).Native()))

	C.gtk_assistant_update_buttons_state(_arg0)
}

// AssistantPage: `GtkAssistantPage` is an auxiliary object used by
// `GtkAssistant.
type AssistantPage interface {
	gextras.Objector

	// Child returns the child to which @page belongs.
	Child() *WidgetClass
}

// AssistantPageClass implements the AssistantPage interface.
type AssistantPageClass struct {
	*externglib.Object
}

var _ AssistantPage = (*AssistantPageClass)(nil)

func wrapAssistantPage(obj *externglib.Object) AssistantPage {
	return &AssistantPageClass{
		Object: obj,
	}
}

func marshalAssistantPage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAssistantPage(obj), nil
}

// Child returns the child to which @page belongs.
func (p *AssistantPageClass) Child() *WidgetClass {
	var _arg0 *C.GtkAssistantPage // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkAssistantPage)(unsafe.Pointer((&AssistantPage).Native()))

	_cret = C.gtk_assistant_page_get_child(_arg0)

	var _widget *WidgetClass // out

	_widget = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*WidgetClass)

	return _widget
}
