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
		{T: externglib.Type(C.gtk_cell_renderer_pixbuf_get_type()), F: marshalCellRendererPixbuf},
	})
}

// CellRendererPixbuf renders a pixbuf in a cell
//
// A CellRendererPixbuf can be used to render an image in a cell. It allows to
// render either a given Pixbuf (set via the CellRendererPixbuf:pixbuf property)
// or a named icon (set via the CellRendererPixbuf:icon-name property).
//
// To support the tree view, CellRendererPixbuf also supports rendering two
// alternative pixbufs, when the CellRenderer:is-expander property is true. If
// the CellRenderer:is-expanded property is true and the
// CellRendererPixbuf:pixbuf-expander-open property is set to a pixbuf, it
// renders that pixbuf, if the CellRenderer:is-expanded property is false and
// the CellRendererPixbuf:pixbuf-expander-closed property is set to a pixbuf, it
// renders that one.
type CellRendererPixbuf interface {
	CellRenderer

	// AsCellRenderer casts the class to the CellRenderer interface.
	AsCellRenderer() CellRenderer

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

// cellRendererPixbuf implements the CellRendererPixbuf interface.
type cellRendererPixbuf struct {
	*externglib.Object
}

var _ CellRendererPixbuf = (*cellRendererPixbuf)(nil)

// WrapCellRendererPixbuf wraps a GObject to a type that implements
// interface CellRendererPixbuf. It is primarily used internally.
func WrapCellRendererPixbuf(obj *externglib.Object) CellRendererPixbuf {
	return cellRendererPixbuf{obj}
}

func marshalCellRendererPixbuf(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererPixbuf(obj), nil
}

// NewCellRendererPixbuf creates a new CellRendererPixbuf. Adjust rendering
// parameters using object properties. Object properties can be set globally
// (with g_object_set()). Also, with TreeViewColumn, you can bind a property to
// a value in a TreeModel. For example, you can bind the “pixbuf” property on
// the cell renderer to a pixbuf value in the model, thus rendering a different
// image in each row of the TreeView.
func NewCellRendererPixbuf() CellRendererPixbuf {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_pixbuf_new()

	var _cellRendererPixbuf CellRendererPixbuf // out

	_cellRendererPixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRendererPixbuf)

	return _cellRendererPixbuf
}

func (c cellRendererPixbuf) AsCellRenderer() CellRenderer {
	return WrapCellRenderer(gextras.InternObject(c))
}

func (c cellRendererPixbuf) Activate(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) bool {
	return WrapCellRenderer(gextras.InternObject(c)).Activate(event, widget, path, backgroundArea, cellArea, flags)
}

func (c cellRendererPixbuf) GetAlignedArea(widget Widget, flags CellRendererState, cellArea *gdk.Rectangle) gdk.Rectangle {
	return WrapCellRenderer(gextras.InternObject(c)).GetAlignedArea(widget, flags, cellArea)
}

func (c cellRendererPixbuf) GetAlignment() (xalign float32, yalign float32) {
	return WrapCellRenderer(gextras.InternObject(c)).GetAlignment()
}

func (c cellRendererPixbuf) GetFixedSize() (width int, height int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetFixedSize()
}

func (c cellRendererPixbuf) GetIsExpanded() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetIsExpanded()
}

func (c cellRendererPixbuf) GetIsExpander() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetIsExpander()
}

func (c cellRendererPixbuf) GetPadding() (xpad int, ypad int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPadding()
}

func (c cellRendererPixbuf) GetPreferredHeight(widget Widget) (minimumSize int, naturalSize int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredHeight(widget)
}

func (c cellRendererPixbuf) GetPreferredHeightForWidth(widget Widget, width int) (minimumHeight int, naturalHeight int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredHeightForWidth(widget, width)
}

func (c cellRendererPixbuf) GetPreferredSize(widget Widget) (minimumSize Requisition, naturalSize Requisition) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredSize(widget)
}

func (c cellRendererPixbuf) GetPreferredWidth(widget Widget) (minimumSize int, naturalSize int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredWidth(widget)
}

func (c cellRendererPixbuf) GetPreferredWidthForHeight(widget Widget, height int) (minimumWidth int, naturalWidth int) {
	return WrapCellRenderer(gextras.InternObject(c)).GetPreferredWidthForHeight(widget, height)
}

func (c cellRendererPixbuf) GetRequestMode() SizeRequestMode {
	return WrapCellRenderer(gextras.InternObject(c)).GetRequestMode()
}

func (c cellRendererPixbuf) GetSensitive() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetSensitive()
}

func (c cellRendererPixbuf) GetState(widget Widget, cellState CellRendererState) StateFlags {
	return WrapCellRenderer(gextras.InternObject(c)).GetState(widget, cellState)
}

func (c cellRendererPixbuf) GetVisible() bool {
	return WrapCellRenderer(gextras.InternObject(c)).GetVisible()
}

func (c cellRendererPixbuf) IsActivatable() bool {
	return WrapCellRenderer(gextras.InternObject(c)).IsActivatable()
}

func (c cellRendererPixbuf) SetAlignment(xalign float32, yalign float32) {
	WrapCellRenderer(gextras.InternObject(c)).SetAlignment(xalign, yalign)
}

func (c cellRendererPixbuf) SetFixedSize(width int, height int) {
	WrapCellRenderer(gextras.InternObject(c)).SetFixedSize(width, height)
}

func (c cellRendererPixbuf) SetIsExpanded(isExpanded bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetIsExpanded(isExpanded)
}

func (c cellRendererPixbuf) SetIsExpander(isExpander bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetIsExpander(isExpander)
}

func (c cellRendererPixbuf) SetPadding(xpad int, ypad int) {
	WrapCellRenderer(gextras.InternObject(c)).SetPadding(xpad, ypad)
}

func (c cellRendererPixbuf) SetSensitive(sensitive bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetSensitive(sensitive)
}

func (c cellRendererPixbuf) SetVisible(visible bool) {
	WrapCellRenderer(gextras.InternObject(c)).SetVisible(visible)
}

func (c cellRendererPixbuf) Snapshot(snapshot Snapshot, widget Widget, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) {
	WrapCellRenderer(gextras.InternObject(c)).Snapshot(snapshot, widget, backgroundArea, cellArea, flags)
}

func (c cellRendererPixbuf) StartEditing(event gdk.Event, widget Widget, path string, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState) CellEditable {
	return WrapCellRenderer(gextras.InternObject(c)).StartEditing(event, widget, path, backgroundArea, cellArea, flags)
}

func (c cellRendererPixbuf) StopEditing(canceled bool) {
	WrapCellRenderer(gextras.InternObject(c)).StopEditing(canceled)
}
