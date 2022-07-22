// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gdk4_Drag_ConnectCancel(gpointer, GdkDragCancelReason, guintptr);
// extern void _gotk4_gdk4_Drag_ConnectDNDFinished(gpointer, guintptr);
// extern void _gotk4_gdk4_Drag_ConnectDropPerformed(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeDragCancelReason = coreglib.Type(C.gdk_drag_cancel_reason_get_type())
	GTypeDrag             = coreglib.Type(C.gdk_drag_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDragCancelReason, F: marshalDragCancelReason},
		coreglib.TypeMarshaler{T: GTypeDrag, F: marshalDrag},
	})
}

// DragCancelReason: used in GdkDrag to the reason of a cancelled DND operation.
type DragCancelReason C.gint

const (
	// DragCancelNoTarget: there is no suitable drop target.
	DragCancelNoTarget DragCancelReason = iota
	// DragCancelUserCancelled: drag cancelled by the user.
	DragCancelUserCancelled
	// DragCancelError: unspecified error.
	DragCancelError
)

func marshalDragCancelReason(p uintptr) (interface{}, error) {
	return DragCancelReason(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DragCancelReason.
func (d DragCancelReason) String() string {
	switch d {
	case DragCancelNoTarget:
		return "NoTarget"
	case DragCancelUserCancelled:
		return "UserCancelled"
	case DragCancelError:
		return "Error"
	default:
		return fmt.Sprintf("DragCancelReason(%d)", d)
	}
}

// DragActionIsUnique checks if action represents a single action or includes
// multiple actions.
//
// When action is 0 - ie no action was given, TRUE is returned.
//
// The function takes the following parameters:
//
//    - action: GdkDragAction.
//
// The function returns the following values:
//
//    - ok: TRUE if exactly one action was given.
//
func DragActionIsUnique(action DragAction) bool {
	var _arg1 C.GdkDragAction // out
	var _cret C.gboolean      // in

	_arg1 = C.GdkDragAction(action)

	_cret = C.gdk_drag_action_is_unique(_arg1)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Drag: GdkDrag object represents the source of an ongoing DND operation.
//
// A GdkDrag is created when a drag is started, and stays alive for duration of
// the DND operation. After a drag has been started with gdk.Drag().Begin, the
// caller gets informed about the status of the ongoing drag operation with
// signals on the GdkDrag object.
//
// GTK provides a higher level abstraction based on top of these functions, and
// so they are not normally needed in GTK applications. See the "Drag and Drop"
// section of the GTK documentation for more information.
type Drag struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Drag)(nil)
)

// Dragger describes types inherited from class Drag.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Dragger interface {
	coreglib.Objector
	baseDrag() *Drag
}

var _ Dragger = (*Drag)(nil)

func wrapDrag(obj *coreglib.Object) *Drag {
	return &Drag{
		Object: obj,
	}
}

func marshalDrag(p uintptr) (interface{}, error) {
	return wrapDrag(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (drag *Drag) baseDrag() *Drag {
	return drag
}

// BaseDrag returns the underlying base object.
func BaseDrag(obj Dragger) *Drag {
	return obj.baseDrag()
}

//export _gotk4_gdk4_Drag_ConnectCancel
func _gotk4_gdk4_Drag_ConnectCancel(arg0 C.gpointer, arg1 C.GdkDragCancelReason, arg2 C.guintptr) {
	var f func(reason DragCancelReason)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(reason DragCancelReason))
	}

	var _reason DragCancelReason // out

	_reason = DragCancelReason(arg1)

	f(_reason)
}

// ConnectCancel is emitted when the drag operation is cancelled.
func (drag *Drag) ConnectCancel(f func(reason DragCancelReason)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drag, "cancel", false, unsafe.Pointer(C._gotk4_gdk4_Drag_ConnectCancel), f)
}

//export _gotk4_gdk4_Drag_ConnectDNDFinished
func _gotk4_gdk4_Drag_ConnectDNDFinished(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDNDFinished is emitted when the destination side has finished reading
// all data.
//
// The drag object can now free all miscellaneous data.
func (drag *Drag) ConnectDNDFinished(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drag, "dnd-finished", false, unsafe.Pointer(C._gotk4_gdk4_Drag_ConnectDNDFinished), f)
}

//export _gotk4_gdk4_Drag_ConnectDropPerformed
func _gotk4_gdk4_Drag_ConnectDropPerformed(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectDropPerformed is emitted when the drop operation is performed on an
// accepting client.
func (drag *Drag) ConnectDropPerformed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(drag, "drop-performed", false, unsafe.Pointer(C._gotk4_gdk4_Drag_ConnectDropPerformed), f)
}

// DropDone informs GDK that the drop ended.
//
// Passing FALSE for success may trigger a drag cancellation animation.
//
// This function is called by the drag source, and should be the last call
// before dropping the reference to the drag.
//
// The GdkDrag will only take the first gdk.Drag.DropDone() call as effective,
// if this function is called multiple times, all subsequent calls will be
// ignored.
//
// The function takes the following parameters:
//
//    - success: whether the drag was ultimatively successful.
//
func (drag *Drag) DropDone(success bool) {
	var _arg0 *C.GdkDrag // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	if success {
		_arg1 = C.TRUE
	}

	C.gdk_drag_drop_done(_arg0, _arg1)
	runtime.KeepAlive(drag)
	runtime.KeepAlive(success)
}

// Actions determines the bitmask of possible actions proposed by the source.
//
// The function returns the following values:
//
//    - dragAction: GdkDragAction flags.
//
func (drag *Drag) Actions() DragAction {
	var _arg0 *C.GdkDrag      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gdk_drag_get_actions(_arg0)
	runtime.KeepAlive(drag)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// Content returns the GdkContentProvider associated to the GdkDrag object.
//
// The function returns the following values:
//
//    - contentProvider: GdkContentProvider associated to drag.
//
func (drag *Drag) Content() *ContentProvider {
	var _arg0 *C.GdkDrag            // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gdk_drag_get_content(_arg0)
	runtime.KeepAlive(drag)

	var _contentProvider *ContentProvider // out

	_contentProvider = wrapContentProvider(coreglib.Take(unsafe.Pointer(_cret)))

	return _contentProvider
}

// Device returns the GdkDevice associated to the GdkDrag object.
//
// The function returns the following values:
//
//    - device: GdkDevice associated to drag.
//
func (drag *Drag) Device() Devicer {
	var _arg0 *C.GdkDrag   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gdk_drag_get_device(_arg0)
	runtime.KeepAlive(drag)

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	return _device
}

// Display gets the GdkDisplay that the drag object was created for.
//
// The function returns the following values:
//
//    - display: GdkDisplay.
//
func (drag *Drag) Display() *Display {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gdk_drag_get_display(_arg0)
	runtime.KeepAlive(drag)

	var _display *Display // out

	_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// DragSurface returns the surface on which the drag icon should be rendered
// during the drag operation.
//
// Note that the surface may not be available until the drag operation has
// begun. GDK will move the surface in accordance with the ongoing drag
// operation. The surface is owned by drag and will be destroyed when the drag
// operation is over.
//
// The function returns the following values:
//
//    - surface (optional): drag surface, or NULL.
//
func (drag *Drag) DragSurface() Surfacer {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gdk_drag_get_drag_surface(_arg0)
	runtime.KeepAlive(drag)

	var _surface Surfacer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Surfacer)
				return ok
			})
			rv, ok := casted.(Surfacer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
			}
			_surface = rv
		}
	}

	return _surface
}

// Formats retrieves the formats supported by this GdkDrag object.
//
// The function returns the following values:
//
//    - contentFormats: GdkContentFormats.
//
func (drag *Drag) Formats() *ContentFormats {
	var _arg0 *C.GdkDrag           // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gdk_drag_get_formats(_arg0)
	runtime.KeepAlive(drag)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _contentFormats
}

// SelectedAction determines the action chosen by the drag destination.
//
// The function returns the following values:
//
//    - dragAction: GdkDragAction value.
//
func (drag *Drag) SelectedAction() DragAction {
	var _arg0 *C.GdkDrag      // out
	var _cret C.GdkDragAction // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gdk_drag_get_selected_action(_arg0)
	runtime.KeepAlive(drag)

	var _dragAction DragAction // out

	_dragAction = DragAction(_cret)

	return _dragAction
}

// Surface returns the GdkSurface where the drag originates.
//
// The function returns the following values:
//
//    - surface: GdkSurface where the drag originates.
//
func (drag *Drag) Surface() Surfacer {
	var _arg0 *C.GdkDrag    // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))

	_cret = C.gdk_drag_get_surface(_arg0)
	runtime.KeepAlive(drag)

	var _surface Surfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Surfacer)
			return ok
		})
		rv, ok := casted.(Surfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
		}
		_surface = rv
	}

	return _surface
}

// SetHotspot sets the position of the drag surface that will be kept under the
// cursor hotspot.
//
// Initially, the hotspot is at the top left corner of the drag surface.
//
// The function takes the following parameters:
//
//    - hotX: x coordinate of the drag surface hotspot.
//    - hotY: y coordinate of the drag surface hotspot.
//
func (drag *Drag) SetHotspot(hotX, hotY int) {
	var _arg0 *C.GdkDrag // out
	var _arg1 C.int      // out
	var _arg2 C.int      // out

	_arg0 = (*C.GdkDrag)(unsafe.Pointer(coreglib.InternObject(drag).Native()))
	_arg1 = C.int(hotX)
	_arg2 = C.int(hotY)

	C.gdk_drag_set_hotspot(_arg0, _arg1, _arg2)
	runtime.KeepAlive(drag)
	runtime.KeepAlive(hotX)
	runtime.KeepAlive(hotY)
}

// DragBegin starts a drag and creates a new drag context for it.
//
// This function is called by the drag source. After this call, you probably
// want to set up the drag icon using the surface returned by
// gdk.Drag.GetDragSurface().
//
// This function returns a reference to the gdk.Drag object, but GTK keeps its
// own reference as well, as long as the DND operation is going on.
//
// Note: if actions include GDK_ACTION_MOVE, you need to listen for the
// gdk.Drag::dnd-finished signal and delete the data at the source if
// gdk.Drag.GetSelectedAction() returns GDK_ACTION_MOVE.
//
// The function takes the following parameters:
//
//    - surface: source surface for this drag.
//    - device that controls this drag.
//    - content: offered content.
//    - actions supported by this drag.
//    - dx: x offset to device's position where the drag nominally started.
//    - dy: y offset to device's position where the drag nominally started.
//
// The function returns the following values:
//
//    - drag (optional): newly created gdk.Drag or NULL on error.
//
func DragBegin(surface Surfacer, device Devicer, content *ContentProvider, actions DragAction, dx, dy float64) Dragger {
	var _arg1 *C.GdkSurface         // out
	var _arg2 *C.GdkDevice          // out
	var _arg3 *C.GdkContentProvider // out
	var _arg4 C.GdkDragAction       // out
	var _arg5 C.double              // out
	var _arg6 C.double              // out
	var _cret *C.GdkDrag            // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer(coreglib.InternObject(surface).Native()))
	_arg2 = (*C.GdkDevice)(unsafe.Pointer(coreglib.InternObject(device).Native()))
	_arg3 = (*C.GdkContentProvider)(unsafe.Pointer(coreglib.InternObject(content).Native()))
	_arg4 = C.GdkDragAction(actions)
	_arg5 = C.double(dx)
	_arg6 = C.double(dy)

	_cret = C.gdk_drag_begin(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(device)
	runtime.KeepAlive(content)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(dx)
	runtime.KeepAlive(dy)

	var _drag Dragger // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Dragger)
				return ok
			})
			rv, ok := casted.(Dragger)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dragger")
			}
			_drag = rv
		}
	}

	return _drag
}
