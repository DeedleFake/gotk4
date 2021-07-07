// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_cell_renderer_spin_get_type()), F: marshalCellRendererSpin},
	})
}

// CellRendererSpin renders a spin button in a cell
//
// CellRendererSpin renders text in a cell like CellRendererText from which it
// is derived. But while CellRendererText offers a simple entry to edit the
// text, CellRendererSpin offers a SpinButton widget. Of course, that means that
// the text has to be parseable as a floating point number.
//
// The range of the spinbutton is taken from the adjustment property of the cell
// renderer, which can be set explicitly or mapped to a column in the tree
// model, like all properties of cell renders. CellRendererSpin also has
// properties for the CellRendererSpin:climb-rate and the number of
// CellRendererSpin:digits to display. Other SpinButton properties can be set in
// a handler for the CellRenderer::editing-started signal.
//
// The CellRendererSpin cell renderer was added in GTK 2.10.
type CellRendererSpin interface {
	CellRendererText

	// AsCellRendererText casts the class to the CellRendererText interface.
	AsCellRendererText() CellRendererText

	// SetFixedHeightFromFont sets the height of a renderer to explicitly be
	// determined by the “font” and “y_pad” property set on it. Further changes
	// in these properties do not affect the height, so they must be accompanied
	// by a subsequent call to this function. Using this function is inflexible,
	// and should really only be used if calculating the size of a cell is too
	// slow (ie, a massive number of cells displayed). If @number_of_rows is -1,
	// then the fixed height is unset, and the height is determined by the
	// properties again.
	//
	// This method is inherited from CellRendererText
	SetFixedHeightFromFont(numberOfRows int)
	// Activate passes an activate event to the cell renderer for possible
	// processing. Some cell renderers may use events; for example,
	// CellRendererToggle toggles when it gets a mouse click.
	//
	// This method is inherited from CellRenderer
	Activate(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) bool
	// GetAlignedArea gets the aligned area used by @cell inside @cell_area.
	// Used for finding the appropriate edit and focus rectangle.
	//
	// This method is inherited from CellRenderer
	GetAlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle
	// GetAlignment fills in @xalign and @yalign with the appropriate values of
	// @cell.
	//
	// This method is inherited from CellRenderer
	GetAlignment() (xalign float32, yalign float32)
	// GetFixedSize fills in @width and @height with the appropriate size of
	// @cell.
	//
	// This method is inherited from CellRenderer
	GetFixedSize() (width int, height int)
	// GetIsExpanded checks whether the given CellRenderer is expanded.
	//
	// This method is inherited from CellRenderer
	GetIsExpanded() bool
	// GetIsExpander checks whether the given CellRenderer is an expander.
	//
	// This method is inherited from CellRenderer
	GetIsExpander() bool
	// GetPadding fills in @xpad and @ypad with the appropriate values of @cell.
	//
	// This method is inherited from CellRenderer
	GetPadding() (xpad int, ypad int)
	// GetPreferredHeight retrieves a renderer’s natural size when rendered to
	// @widget.
	//
	// This method is inherited from CellRenderer
	GetPreferredHeight(widget Widget) (minimumSize int, naturalSize int)
	// GetPreferredHeightForWidth retrieves a cell renderers’s minimum and
	// natural height if it were rendered to @widget with the specified @width.
	//
	// This method is inherited from CellRenderer
	GetPreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int)
	// GetPreferredSize retrieves the minimum and natural size of a cell taking
	// into account the widget’s preference for height-for-width management.
	//
	// This method is inherited from CellRenderer
	GetPreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition)
	// GetPreferredWidth retrieves a renderer’s natural size when rendered to
	// @widget.
	//
	// This method is inherited from CellRenderer
	GetPreferredWidth(widget Widget) (minimumSize int, naturalSize int)
	// GetPreferredWidthForHeight retrieves a cell renderers’s minimum and
	// natural width if it were rendered to @widget with the specified @height.
	//
	// This method is inherited from CellRenderer
	GetPreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int)
	// GetRequestMode gets whether the cell renderer prefers a height-for-width
	// layout or a width-for-height layout.
	//
	// This method is inherited from CellRenderer
	GetRequestMode() SizeRequestMode
	// GetSensitive returns the cell renderer’s sensitivity.
	//
	// This method is inherited from CellRenderer
	GetSensitive() bool
	// GetState translates the cell renderer state to StateFlags, based on the
	// cell renderer and widget sensitivity, and the given CellRendererState.
	//
	// This method is inherited from CellRenderer
	GetState(widget Widget, cellState CellRendererState) StateFlags
	// GetVisible returns the cell renderer’s visibility.
	//
	// This method is inherited from CellRenderer
	GetVisible() bool
	// IsActivatable checks whether the cell renderer can do something when
	// activated.
	//
	// This method is inherited from CellRenderer
	IsActivatable() bool
	// SetAlignment sets the renderer’s alignment within its available space.
	//
	// This method is inherited from CellRenderer
	SetAlignment(xalign float32, yalign float32)
	// SetFixedSize sets the renderer size to be explicit, independent of the
	// properties set.
	//
	// This method is inherited from CellRenderer
	SetFixedSize(width int, height int)
	// SetIsExpanded sets whether the given CellRenderer is expanded.
	//
	// This method is inherited from CellRenderer
	SetIsExpanded(isExpanded bool)
	// SetIsExpander sets whether the given CellRenderer is an expander.
	//
	// This method is inherited from CellRenderer
	SetIsExpander(isExpander bool)
	// SetPadding sets the renderer’s padding.
	//
	// This method is inherited from CellRenderer
	SetPadding(xpad int, ypad int)
	// SetSensitive sets the cell renderer’s sensitivity.
	//
	// This method is inherited from CellRenderer
	SetSensitive(sensitive bool)
	// SetVisible sets the cell renderer’s visibility.
	//
	// This method is inherited from CellRenderer
	SetVisible(visible bool)
	// Snapshot invokes the virtual render function of the CellRenderer. The
	// three passed-in rectangles are areas in @cr. Most renderers will draw
	// within @cell_area; the xalign, yalign, xpad, and ypad fields of the
	// CellRenderer should be honored with respect to @cell_area.
	// @background_area includes the blank space around the cell, and also the
	// area containing the tree expander; so the @background_area rectangles for
	// all cells tile to cover the entire @window.
	//
	// This method is inherited from CellRenderer
	Snapshot(snapshot Snapshot, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState)
	// StartEditing starts editing the contents of this @cell, through a new
	// CellEditable widget created by the CellRendererClass.start_editing
	// virtual function.
	//
	// This method is inherited from CellRenderer
	StartEditing(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditable
	// StopEditing informs the cell renderer that the editing is stopped. If
	// @canceled is true, the cell renderer will emit the
	// CellRenderer::editing-canceled signal.
	//
	// This function should be called by cell renderer implementations in
	// response to the CellEditable::editing-done signal of CellEditable.
	//
	// This method is inherited from CellRenderer
	StopEditing(canceled bool)
}

// cellRendererSpin implements the CellRendererSpin interface.
type cellRendererSpin struct {
	*externglib.Object
}

var _ CellRendererSpin = (*cellRendererSpin)(nil)

// WrapCellRendererSpin wraps a GObject to a type that implements
// interface CellRendererSpin. It is primarily used internally.
func WrapCellRendererSpin(obj *externglib.Object) CellRendererSpin {
	return cellRendererSpin{obj}
}

func marshalCellRendererSpin(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererSpin(obj), nil
}

// NewCellRendererSpin creates a new CellRendererSpin.
func NewCellRendererSpin() CellRendererSpin {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_spin_new()

	var _cellRendererSpin CellRendererSpin // out

	_cellRendererSpin = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRendererSpin)

	return _cellRendererSpin
}

func (c cellRendererSpin) AsCellRendererText() CellRendererText {
	return WrapCellRendererText(gextras.InternObject(c))
}

func (r cellRendererSpin) SetFixedHeightFromFont(numberOfRows int) {
	WrapCellRendererText(gextras.InternObject(r)).SetFixedHeightFromFont(numberOfRows)
}

func (c cellRendererSpin) Activate(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) bool {
	return WrapCellRenderer(gextras.InternObject(c)).Activate(event, widget, path, backgroundArea, cellArea, flags)
}

func (c cellRendererSpin) GetAlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle {
	return WrapCellRenderer(gextras.InternObject(c)).GetAlignedArea(widget, flags, cellArea)
}

func (c cellRendererSpin) GetAlignment() (xalign float32, yalign float32) {
	return WrapCellRenderer(gextras.InternObject(c)).GetAlignment()
}

func (c cellRendererSpin) GetFixedSize() (width int, height int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetFixedSize()
}

func (c cellRendererSpin) GetIsExpanded() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetIsExpanded()
}

func (c cellRendererSpin) GetIsExpander() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetIsExpander()
}

func (c cellRendererSpin) GetPadding() (xpad int, ypad int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPadding()
}

func (c cellRendererSpin) GetPreferredHeight(widget Widget) (minimumSize int, naturalSize int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredHeight(widget)
}

func (c cellRendererSpin) GetPreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredHeightForWidth(widget, width)
}

func (c cellRendererSpin) GetPreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredSize(widget)
}

func (c cellRendererSpin) GetPreferredWidth(widget Widget) (minimumSize int, naturalSize int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredWidth(widget)
}

func (c cellRendererSpin) GetPreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredWidthForHeight(widget, height)
}

func (c cellRendererSpin) GetRequestMode() SizeRequestMode {
	return WrapCellRenderer(gextras.InternObject(c)).GetRequestMode()
}

func (c cellRendererSpin) GetSensitive() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetSensitive()
}

func (c cellRendererSpin) GetState(widget Widget, cellState CellRendererState) StateFlags {
	return WrapCellRenderer(gextras.InternObject(c)).GetState(widget, cellState)
}

func (c cellRendererSpin) GetVisible() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetVisible()
}

func (c cellRendererSpin) IsActivatable() bool {
	return WrapCellRenderer(gextras.InternObject(c)).IsActivatable()
}

func (c cellRendererSpin) SetAlignment(xalign float32, yalign float32) {
	WrapCellRenderer(gextras.InternObject(c)).SetAlignment(xalign, yalign)
}

func (c cellRendererSpin) SetFixedSize(width int, height int) {
	WrapCellRenderer(gextras.InternObject(c)).SetFixedSize(width, height)
}

func (c cellRendererSpin) SetIsExpanded(isExpanded bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetIsExpanded(isExpanded)
}

func (c cellRendererSpin) SetIsExpander(isExpander bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetIsExpander(isExpander)
}

func (c cellRendererSpin) SetPadding(xpad int, ypad int) {
	WrapCellRenderer(gextras.InternObject(c)).SetPadding(xpad, ypad)
}

func (c cellRendererSpin) SetSensitive(sensitive bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetSensitive(sensitive)
}

func (c cellRendererSpin) SetVisible(visible bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetVisible(visible)
}

func (c cellRendererSpin) Snapshot(snapshot Snapshot, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) {
	WrapCellRenderer(gextras.InternObject(c)).Snapshot(snapshot, widget, backgroundArea, cellArea, flags)
}

func (c cellRendererSpin) StartEditing(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditable {
	return WrapCellRenderer(gextras.InternObject(c)).StartEditing(event, widget, path, backgroundArea, cellArea, flags)
}

func (c cellRendererSpin) StopEditing(canceled bool) {
	WrapCellRenderer(gextras.InternObject(c)).StopEditing(canceled)
}
