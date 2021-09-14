// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "stubs.h"
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_mode_get_type()), F: marshalCellRendererMode},
		{T: externglib.Type(C.gtk_cell_renderer_state_get_type()), F: marshalCellRendererState},
		{T: externglib.Type(C.gtk_cell_renderer_get_type()), F: marshalCellRendererer},
	})
}

// CellRendererMode identifies how the user can interact with a particular cell.
type CellRendererMode int

const (
	// CellRendererModeInert: cell is just for display and cannot be interacted
	// with. Note that this doesn’t mean that eg. the row being drawn can’t be
	// selected -- just that a particular element of it cannot be individually
	// modified.
	CellRendererModeInert CellRendererMode = iota
	// CellRendererModeActivatable: cell can be clicked.
	CellRendererModeActivatable
	// CellRendererModeEditable: cell can be edited or otherwise modified.
	CellRendererModeEditable
)

func marshalCellRendererMode(p uintptr) (interface{}, error) {
	return CellRendererMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for CellRendererMode.
func (c CellRendererMode) String() string {
	switch c {
	case CellRendererModeInert:
		return "Inert"
	case CellRendererModeActivatable:
		return "Activatable"
	case CellRendererModeEditable:
		return "Editable"
	default:
		return fmt.Sprintf("CellRendererMode(%d)", c)
	}
}

// CellRendererState tells how a cell is to be rendered.
type CellRendererState int

const (
	// CellRendererSelected: cell is currently selected, and probably has a
	// selection colored background to render to.
	CellRendererSelected CellRendererState = 0b1
	// CellRendererPrelit: mouse is hovering over the cell.
	CellRendererPrelit CellRendererState = 0b10
	// CellRendererInsensitive: cell is drawn in an insensitive manner
	CellRendererInsensitive CellRendererState = 0b100
	// CellRendererSorted: cell is in a sorted row
	CellRendererSorted CellRendererState = 0b1000
	// CellRendererFocused: cell is in the focus row.
	CellRendererFocused CellRendererState = 0b10000
	// CellRendererExpandable: cell is in a row that can be expanded
	CellRendererExpandable CellRendererState = 0b100000
	// CellRendererExpanded: cell is in a row that is expanded
	CellRendererExpanded CellRendererState = 0b1000000
)

func marshalCellRendererState(p uintptr) (interface{}, error) {
	return CellRendererState(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for CellRendererState.
func (c CellRendererState) String() string {
	if c == 0 {
		return "CellRendererState(0)"
	}

	var builder strings.Builder
	builder.Grow(146)

	for c != 0 {
		next := c & (c - 1)
		bit := c - next

		switch bit {
		case CellRendererSelected:
			builder.WriteString("Selected|")
		case CellRendererPrelit:
			builder.WriteString("Prelit|")
		case CellRendererInsensitive:
			builder.WriteString("Insensitive|")
		case CellRendererSorted:
			builder.WriteString("Sorted|")
		case CellRendererFocused:
			builder.WriteString("Focused|")
		case CellRendererExpandable:
			builder.WriteString("Expandable|")
		case CellRendererExpanded:
			builder.WriteString("Expanded|")
		default:
			builder.WriteString(fmt.Sprintf("CellRendererState(0b%b)|", bit))
		}

		c = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if c contains other.
func (c CellRendererState) Has(other CellRendererState) bool {
	return (c & other) == other
}

// CellRendererOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellRendererOverrider interface {
	// Activate passes an activate event to the cell renderer for possible
	// processing. Some cell renderers may use events; for example,
	// CellRendererToggle toggles when it gets a mouse click.
	Activate(event gdk.Eventer, widget Widgetter, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) bool
	EditingCanceled()
	EditingStarted(editable CellEditabler, path string)
	// AlignedArea gets the aligned area used by cell inside cell_area. Used for
	// finding the appropriate edit and focus rectangle.
	AlignedArea(widget Widgetter, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle
	// PreferredHeight retrieves a renderer’s natural size when rendered to
	// widget.
	PreferredHeight(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
	// height if it were rendered to widget with the specified width.
	PreferredHeightForWidth(widget Widgetter, width int) (minimumHeight int, naturalHeight int)
	// PreferredWidth retrieves a renderer’s natural size when rendered to
	// widget.
	PreferredWidth(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
	// width if it were rendered to widget with the specified height.
	PreferredWidthForHeight(widget Widgetter, height int) (minimumWidth int, naturalWidth int)
	// RequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	RequestMode() SizeRequestMode
	// Snapshot invokes the virtual render function of the CellRenderer. The
	// three passed-in rectangles are areas in cr. Most renderers will draw
	// within cell_area; the xalign, yalign, xpad, and ypad fields of the
	// CellRenderer should be honored with respect to cell_area. background_area
	// includes the blank space around the cell, and also the area containing
	// the tree expander; so the background_area rectangles for all cells tile
	// to cover the entire window.
	Snapshot(snapshot *Snapshot, widget Widgetter, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState)
	// StartEditing starts editing the contents of this cell, through a new
	// CellEditable widget created by the CellRendererClass.start_editing
	// virtual function.
	StartEditing(event gdk.Eventer, widget Widgetter, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditabler
}

// CellRenderer: object for rendering a single cell
//
// The CellRenderer is a base class of a set of objects used for rendering a
// cell to a #cairo_t. These objects are used primarily by the TreeView widget,
// though they aren’t tied to them in any specific way. It is worth noting that
// CellRenderer is not a Widget and cannot be treated as such.
//
// The primary use of a CellRenderer is for drawing a certain graphical elements
// on a #cairo_t. Typically, one cell renderer is used to draw many cells on the
// screen. To this extent, it isn’t expected that a CellRenderer keep any
// permanent state around. Instead, any state is set just prior to use using
// #GObjects property system. Then, the cell is measured using
// gtk_cell_renderer_get_preferred_size(). Finally, the cell is rendered in the
// correct location using gtk_cell_renderer_snapshot().
//
// There are a number of rules that must be followed when writing a new
// CellRenderer. First and foremost, it’s important that a certain set of
// properties will always yield a cell renderer of the same size, barring a
// style change. The CellRenderer also has a number of generic properties that
// are expected to be honored by all children.
//
// Beyond merely rendering a cell, cell renderers can optionally provide active
// user interface elements. A cell renderer can be “activatable” like
// CellRendererToggle, which toggles when it gets activated by a mouse click, or
// it can be “editable” like CellRendererText, which allows the user to edit the
// text using a widget implementing the CellEditable interface, e.g. Entry. To
// make a cell renderer activatable or editable, you have to implement the
// CellRendererClass.activate or CellRendererClass.start_editing virtual
// functions, respectively.
//
// Many properties of CellRenderer and its subclasses have a corresponding “set”
// property, e.g. “cell-background-set” corresponds to “cell-background”. These
// “set” properties reflect whether a property has been set or not. You should
// not set them independently.
type CellRenderer struct {
	externglib.InitiallyUnowned
}

// CellRendererer describes CellRenderer's abstract methods.
type CellRendererer interface {
	externglib.Objector

	// Activate passes an activate event to the cell renderer for possible
	// processing.
	Activate(event gdk.Eventer, widget Widgetter, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) bool
	// AlignedArea gets the aligned area used by cell inside cell_area.
	AlignedArea(widget Widgetter, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle
	// Alignment fills in xalign and yalign with the appropriate values of cell.
	Alignment() (xalign float32, yalign float32)
	// FixedSize fills in width and height with the appropriate size of cell.
	FixedSize() (width int, height int)
	// IsExpanded checks whether the given CellRenderer is expanded.
	IsExpanded() bool
	// IsExpander checks whether the given CellRenderer is an expander.
	IsExpander() bool
	// Padding fills in xpad and ypad with the appropriate values of cell.
	Padding() (xpad int, ypad int)
	// PreferredHeight retrieves a renderer’s natural size when rendered to
	// widget.
	PreferredHeight(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
	// height if it were rendered to widget with the specified width.
	PreferredHeightForWidth(widget Widgetter, width int) (minimumHeight int, naturalHeight int)
	// PreferredSize retrieves the minimum and natural size of a cell taking
	// into account the widget’s preference for height-for-width management.
	PreferredSize(widget Widgetter) (minimumSize Requisition, naturalSize Requisition)
	// PreferredWidth retrieves a renderer’s natural size when rendered to
	// widget.
	PreferredWidth(widget Widgetter) (minimumSize int, naturalSize int)
	// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
	// width if it were rendered to widget with the specified height.
	PreferredWidthForHeight(widget Widgetter, height int) (minimumWidth int, naturalWidth int)
	// RequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	RequestMode() SizeRequestMode
	// Sensitive returns the cell renderer’s sensitivity.
	Sensitive() bool
	// State translates the cell renderer state to StateFlags, based on the cell
	// renderer and widget sensitivity, and the given CellRendererState.
	State(widget Widgetter, cellState CellRendererState) StateFlags
	// Visible returns the cell renderer’s visibility.
	Visible() bool
	// IsActivatable checks whether the cell renderer can do something when
	// activated.
	IsActivatable() bool
	// SetAlignment sets the renderer’s alignment within its available space.
	SetAlignment(xalign float32, yalign float32)
	// SetFixedSize sets the renderer size to be explicit, independent of the
	// properties set.
	SetFixedSize(width int, height int)
	// SetIsExpanded sets whether the given CellRenderer is expanded.
	SetIsExpanded(isExpanded bool)
	// SetIsExpander sets whether the given CellRenderer is an expander.
	SetIsExpander(isExpander bool)
	// SetPadding sets the renderer’s padding.
	SetPadding(xpad int, ypad int)
	// SetSensitive sets the cell renderer’s sensitivity.
	SetSensitive(sensitive bool)
	// SetVisible sets the cell renderer’s visibility.
	SetVisible(visible bool)
	// Snapshot invokes the virtual render function of the CellRenderer.
	Snapshot(snapshot *Snapshot, widget Widgetter, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState)
	// StartEditing starts editing the contents of this cell, through a new
	// CellEditable widget created by the CellRendererClass.start_editing
	// virtual function.
	StartEditing(event gdk.Eventer, widget Widgetter, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditabler
	// StopEditing informs the cell renderer that the editing is stopped.
	StopEditing(canceled bool)
}

var _ CellRendererer = (*CellRenderer)(nil)

func wrapCellRenderer(obj *externglib.Object) *CellRenderer {
	return &CellRenderer{
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
	}
}

func marshalCellRendererer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRenderer(obj), nil
}

// Activate passes an activate event to the cell renderer for possible
// processing. Some cell renderers may use events; for example,
// CellRendererToggle toggles when it gets a mouse click.
func (cell *CellRenderer) Activate(event gdk.Eventer, widget Widgetter, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) bool {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GdkEvent            // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.char                // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 *C.GdkRectangle        // out
	var _arg6 C.GtkCellRendererState // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(backgroundArea)))
	_arg5 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))
	_arg6 = C.GtkCellRendererState(flags)

	_cret = C.gtk_cell_renderer_activate(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(event)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(path)
	runtime.KeepAlive(backgroundArea)
	runtime.KeepAlive(cellArea)
	runtime.KeepAlive(flags)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AlignedArea gets the aligned area used by cell inside cell_area. Used for
// finding the appropriate edit and focus rectangle.
func (cell *CellRenderer) AlignedArea(widget Widgetter, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkCellRendererState // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 C.GdkRectangle         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkCellRendererState(flags)
	_arg3 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))

	C.gtk_cell_renderer_get_aligned_area(_arg0, _arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(cellArea)

	var _alignedArea gdk.Rectangle // out

	_alignedArea = *(*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))

	return _alignedArea
}

// Alignment fills in xalign and yalign with the appropriate values of cell.
func (cell *CellRenderer) Alignment() (xalign float32, yalign float32) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.float            // in
	var _arg2 C.float            // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_alignment(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cell)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
}

// FixedSize fills in width and height with the appropriate size of cell.
func (cell *CellRenderer) FixedSize() (width int, height int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // in
	var _arg2 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_fixed_size(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cell)

	var _width int  // out
	var _height int // out

	_width = int(_arg1)
	_height = int(_arg2)

	return _width, _height
}

// IsExpanded checks whether the given CellRenderer is expanded.
func (cell *CellRenderer) IsExpanded() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_is_expanded(_arg0)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsExpander checks whether the given CellRenderer is an expander.
func (cell *CellRenderer) IsExpander() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_is_expander(_arg0)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Padding fills in xpad and ypad with the appropriate values of cell.
func (cell *CellRenderer) Padding() (xpad int, ypad int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // in
	var _arg2 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_renderer_get_padding(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(cell)

	var _xpad int // out
	var _ypad int // out

	_xpad = int(_arg1)
	_ypad = int(_arg2)

	return _xpad, _ypad
}

// PreferredHeight retrieves a renderer’s natural size when rendered to widget.
func (cell *CellRenderer) PreferredHeight(widget Widgetter) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // in
	var _arg3 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_height(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = int(_arg2)
	_naturalSize = int(_arg3)

	return _minimumSize, _naturalSize
}

// PreferredHeightForWidth retrieves a cell renderers’s minimum and natural
// height if it were rendered to widget with the specified width.
func (cell *CellRenderer) PreferredHeightForWidth(widget Widgetter, width int) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // in
	var _arg4 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.int(width)

	C.gtk_cell_renderer_get_preferred_height_for_width(_arg0, _arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(width)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = int(_arg3)
	_naturalHeight = int(_arg4)

	return _minimumHeight, _naturalHeight
}

// PreferredSize retrieves the minimum and natural size of a cell taking into
// account the widget’s preference for height-for-width management.
func (cell *CellRenderer) PreferredSize(widget Widgetter) (minimumSize Requisition, naturalSize Requisition) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.GtkRequisition   // in
	var _arg3 C.GtkRequisition   // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_size(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)

	var _minimumSize Requisition // out
	var _naturalSize Requisition // out

	_minimumSize = *(*Requisition)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_naturalSize = *(*Requisition)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _minimumSize, _naturalSize
}

// PreferredWidth retrieves a renderer’s natural size when rendered to widget.
func (cell *CellRenderer) PreferredWidth(widget Widgetter) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // in
	var _arg3 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_cell_renderer_get_preferred_width(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = int(_arg2)
	_naturalSize = int(_arg3)

	return _minimumSize, _naturalSize
}

// PreferredWidthForHeight retrieves a cell renderers’s minimum and natural
// width if it were rendered to widget with the specified height.
func (cell *CellRenderer) PreferredWidthForHeight(widget Widgetter, height int) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.int              // out
	var _arg3 C.int              // in
	var _arg4 C.int              // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.int(height)

	C.gtk_cell_renderer_get_preferred_width_for_height(_arg0, _arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(height)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = int(_arg3)
	_naturalWidth = int(_arg4)

	return _minimumWidth, _naturalWidth
}

// RequestMode gets whether the cell renderer prefers a height-for-width layout
// or a width-for-height layout.
func (cell *CellRenderer) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkCellRenderer   // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_request_mode(_arg0)
	runtime.KeepAlive(cell)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = SizeRequestMode(_cret)

	return _sizeRequestMode
}

// Sensitive returns the cell renderer’s sensitivity.
func (cell *CellRenderer) Sensitive() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_sensitive(_arg0)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// State translates the cell renderer state to StateFlags, based on the cell
// renderer and widget sensitivity, and the given CellRendererState.
func (cell *CellRenderer) State(widget Widgetter, cellState CellRendererState) StateFlags {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkCellRendererState // out
	var _cret C.GtkStateFlags        // in

	if cell != nil {
		_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	}
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}
	_arg2 = C.GtkCellRendererState(cellState)

	_cret = C.gtk_cell_renderer_get_state(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(cellState)

	var _stateFlags StateFlags // out

	_stateFlags = StateFlags(_cret)

	return _stateFlags
}

// Visible returns the cell renderer’s visibility.
func (cell *CellRenderer) Visible() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_get_visible(_arg0)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsActivatable checks whether the cell renderer can do something when
// activated.
func (cell *CellRenderer) IsActivatable() bool {
	var _arg0 *C.GtkCellRenderer // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_cell_renderer_is_activatable(_arg0)
	runtime.KeepAlive(cell)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAlignment sets the renderer’s alignment within its available space.
func (cell *CellRenderer) SetAlignment(xalign float32, yalign float32) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.float(xalign)
	_arg2 = C.float(yalign)

	C.gtk_cell_renderer_set_alignment(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
}

// SetFixedSize sets the renderer size to be explicit, independent of the
// properties set.
func (cell *CellRenderer) SetFixedSize(width int, height int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_cell_renderer_set_fixed_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetIsExpanded sets whether the given CellRenderer is expanded.
func (cell *CellRenderer) SetIsExpanded(isExpanded bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if isExpanded {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_is_expanded(_arg0, _arg1)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(isExpanded)
}

// SetIsExpander sets whether the given CellRenderer is an expander.
func (cell *CellRenderer) SetIsExpander(isExpander bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if isExpander {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_is_expander(_arg0, _arg1)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(isExpander)
}

// SetPadding sets the renderer’s padding.
func (cell *CellRenderer) SetPadding(xpad int, ypad int) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = C.int(xpad)
	_arg2 = C.int(ypad)

	C.gtk_cell_renderer_set_padding(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(xpad)
	runtime.KeepAlive(ypad)
}

// SetSensitive sets the cell renderer’s sensitivity.
func (cell *CellRenderer) SetSensitive(sensitive bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_sensitive(_arg0, _arg1)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(sensitive)
}

// SetVisible sets the cell renderer’s visibility.
func (cell *CellRenderer) SetVisible(visible bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_set_visible(_arg0, _arg1)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(visible)
}

// Snapshot invokes the virtual render function of the CellRenderer. The three
// passed-in rectangles are areas in cr. Most renderers will draw within
// cell_area; the xalign, yalign, xpad, and ypad fields of the CellRenderer
// should be honored with respect to cell_area. background_area includes the
// blank space around the cell, and also the area containing the tree expander;
// so the background_area rectangles for all cells tile to cover the entire
// window.
func (cell *CellRenderer) Snapshot(snapshot *Snapshot, widget Widgetter, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GtkSnapshot         // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 C.GtkCellRendererState // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg1 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(backgroundArea)))
	_arg4 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))
	_arg5 = C.GtkCellRendererState(flags)

	C.gtk_cell_renderer_snapshot(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(snapshot)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(backgroundArea)
	runtime.KeepAlive(cellArea)
	runtime.KeepAlive(flags)
}

// StartEditing starts editing the contents of this cell, through a new
// CellEditable widget created by the CellRendererClass.start_editing virtual
// function.
func (cell *CellRenderer) StartEditing(event gdk.Eventer, widget Widgetter, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditabler {
	var _arg0 *C.GtkCellRenderer     // out
	var _arg1 *C.GdkEvent            // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.char                // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 *C.GdkRectangle        // out
	var _arg6 C.GtkCellRendererState // out
	var _cret *C.GtkCellEditable     // in

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if event != nil {
		_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	}
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(backgroundArea)))
	_arg5 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(cellArea)))
	_arg6 = C.GtkCellRendererState(flags)

	_cret = C.gtk_cell_renderer_start_editing(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(event)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(path)
	runtime.KeepAlive(backgroundArea)
	runtime.KeepAlive(cellArea)
	runtime.KeepAlive(flags)

	var _cellEditable CellEditabler // out

	if _cret != nil {
		_cellEditable = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(CellEditabler)
	}

	return _cellEditable
}

// StopEditing informs the cell renderer that the editing is stopped. If
// canceled is TRUE, the cell renderer will emit the
// CellRenderer::editing-canceled signal.
//
// This function should be called by cell renderer implementations in response
// to the CellEditable::editing-done signal of CellEditable.
func (cell *CellRenderer) StopEditing(canceled bool) {
	var _arg0 *C.GtkCellRenderer // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if canceled {
		_arg1 = C.TRUE
	}

	C.gtk_cell_renderer_stop_editing(_arg0, _arg1)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(canceled)
}
