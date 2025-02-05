// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_DropTarget_ConnectLeave(gpointer, guintptr);
// extern gboolean _gotk4_gtk4_DropTarget_ConnectDrop(gpointer, GValue, gdouble, gdouble, guintptr);
// extern gboolean _gotk4_gtk4_DropTarget_ConnectAccept(gpointer, GdkDrop*, guintptr);
// extern GdkDragAction _gotk4_gtk4_DropTarget_ConnectMotion(gpointer, gdouble, gdouble, guintptr);
// extern GdkDragAction _gotk4_gtk4_DropTarget_ConnectEnter(gpointer, gdouble, gdouble, guintptr);
import "C"

// GType values.
var (
	GTypeDropTarget = coreglib.Type(C.gtk_drop_target_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDropTarget, F: marshalDropTarget},
	})
}

// DropTarget: GtkDropTarget is an event controller to receive Drag-and-Drop
// operations.
//
// The most basic way to use a GtkDropTarget to receive drops on a widget is to
// create it via gtk.DropTarget.New, passing in the GType of the data you want
// to receive and connect to the gtk.DropTarget::drop signal to receive the
// data:
//
//    static gboolean
//    on_drop (GtkDropTarget *target,
//             const GValue  *value,
//             double         x,
//             double         y,
//             gpointer       data)
//    {
//      MyWidget *self = data;
//
//      // Call the appropriate setter depending on the type of data
//      // that we received
//      if (G_VALUE_HOLDS (value, G_TYPE_FILE))
//        my_widget_set_file (self, g_value_get_object (value));
//      else if (G_VALUE_HOLDS (value, GDK_TYPE_PIXBUF))
//        my_widget_set_pixbuf (self, g_value_get_object (value));
//      else
//        return FALSE;
//
//      return TRUE;
//    }
//
//    static void
//    my_widget_init (MyWidget *self)
//    {
//      GtkDropTarget *target =
//        gtk_drop_target_new (G_TYPE_INVALID, GDK_ACTION_COPY);
//
//      // This widget accepts two types of drop types: GFile objects
//      // and GdkPixbuf objects
//      gtk_drop_target_set_gtypes (target, (GTypes [2]) {
//        G_TYPE_FILE,
//        GDK_TYPE_PIXBUF,
//      }, 2);
//
//      gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (target));
//    }
//
//
// GtkDropTarget supports more options, such as:
//
//    * rejecting potential drops via the gtk.DropTarget::accept signal
//      and the gtk.DropTarget.Reject() function to let other drop
//      targets handle the drop
//    * tracking an ongoing drag operation before the drop via the
//      gtk.DropTarget::enter, gtk.DropTarget::motion and
//      gtk.DropTarget::leave signals
//    * configuring how to receive data by setting the
//      gtk.DropTarget:preload property and listening for its
//      availability via the gtk.DropTarget:value property
//
// However, GtkDropTarget is ultimately modeled in a synchronous way and only
// supports data transferred via GType. If you want full control over an ongoing
// drop, the gtk.DropTargetAsync object gives you this ability.
//
// While a pointer is dragged over the drop target's widget and the drop has not
// been rejected, that widget will receive the GTK_STATE_FLAG_DROP_ACTIVE state,
// which can be used to style the widget.
//
// If you are not interested in receiving the drop, but just want to update UI
// state during a Drag-and-Drop operation (e.g. switching tabs), you can use
// gtk.DropControllerMotion.
type DropTarget struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*DropTarget)(nil)
)

func wrapDropTarget(obj *coreglib.Object) *DropTarget {
	return &DropTarget{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalDropTarget(p uintptr) (interface{}, error) {
	return wrapDropTarget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAccept is emitted on the drop site when a drop operation is about to
// begin.
//
// If the drop is not accepted, FALSE will be returned and the drop target will
// ignore the drop. If TRUE is returned, the drop is accepted for now but may be
// rejected later via a call to gtk.DropTarget.Reject() or ultimately by
// returning FALSE from a gtk.DropTarget::drop handler.
//
// The default handler for this signal decides whether to accept the drop based
// on the formats provided by the drop.
//
// If the decision whether the drop will be accepted or rejected depends on the
// data, this function should return TRUE, the gtk.DropTarget:preload property
// should be set and the value should be inspected via the ::notify:value
// signal, calling gtk.DropTarget.Reject() if required.
func (self *DropTarget) ConnectAccept(f func(drop gdk.Dropper) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "accept", false, unsafe.Pointer(C._gotk4_gtk4_DropTarget_ConnectAccept), f)
}

// ConnectDrop is emitted on the drop site when the user drops the data onto the
// widget.
//
// The signal handler must determine whether the pointer position is in a drop
// zone or not. If it is not in a drop zone, it returns FALSE and no further
// processing is necessary.
//
// Otherwise, the handler returns TRUE. In this case, this handler will accept
// the drop. The handler is responsible for rading the given value and
// performing the drop operation.
func (self *DropTarget) ConnectDrop(f func(value coreglib.Value, x, y float64) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "drop", false, unsafe.Pointer(C._gotk4_gtk4_DropTarget_ConnectDrop), f)
}

// ConnectEnter is emitted on the drop site when the pointer enters the widget.
//
// It can be used to set up custom highlighting.
func (self *DropTarget) ConnectEnter(f func(x, y float64) (dragAction gdk.DragAction)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "enter", false, unsafe.Pointer(C._gotk4_gtk4_DropTarget_ConnectEnter), f)
}

// ConnectLeave is emitted on the drop site when the pointer leaves the widget.
//
// Its main purpose it to undo things done in gtk.DropTarget::enter.
func (self *DropTarget) ConnectLeave(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "leave", false, unsafe.Pointer(C._gotk4_gtk4_DropTarget_ConnectLeave), f)
}

// ConnectMotion is emitted while the pointer is moving over the drop target.
func (self *DropTarget) ConnectMotion(f func(x, y float64) (dragAction gdk.DragAction)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "motion", false, unsafe.Pointer(C._gotk4_gtk4_DropTarget_ConnectMotion), f)
}

// NewDropTarget creates a new GtkDropTarget object.
//
// If the drop target should support more than 1 type, pass G_TYPE_INVALID for
// type and then call gtk.DropTarget.SetGTypes().
//
// The function takes the following parameters:
//
//    - typ: supported type or G_TYPE_INVALID.
//    - actions: supported actions.
//
// The function returns the following values:
//
//    - dropTarget: new GtkDropTarget.
//
func NewDropTarget(typ coreglib.Type, actions gdk.DragAction) *DropTarget {
	var _arg1 C.GType          // out
	var _arg2 C.GdkDragAction  // out
	var _cret *C.GtkDropTarget // in

	_arg1 = C.GType(typ)
	_arg2 = C.GdkDragAction(actions)

	_cret = C.gtk_drop_target_new(_arg1, _arg2)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(actions)

	var _dropTarget *DropTarget // out

	_dropTarget = wrapDropTarget(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dropTarget
}

// Actions gets the actions that this drop target supports.
//
// The function returns the following values:
//
//    - dragAction actions that this drop target supports.
//
func (self *DropTarget) Actions() gdk.DragAction {
	var _arg0 *C.GtkDropTarget // out
	var _cret C.GdkDragAction  // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_drop_target_get_actions(_arg0)
	runtime.KeepAlive(self)

	var _dragAction gdk.DragAction // out

	_dragAction = gdk.DragAction(_cret)

	return _dragAction
}

// Drop gets the currently handled drop operation.
//
// If no drop operation is going on, NULL is returned.
//
// The function returns the following values:
//
//    - drop (optional): current drop.
//
func (self *DropTarget) Drop() gdk.Dropper {
	var _arg0 *C.GtkDropTarget // out
	var _cret *C.GdkDrop       // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_drop_target_get_drop(_arg0)
	runtime.KeepAlive(self)

	var _drop gdk.Dropper // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Dropper)
				return ok
			})
			rv, ok := casted.(gdk.Dropper)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Dropper")
			}
			_drop = rv
		}
	}

	return _drop
}

// Formats gets the data formats that this drop target accepts.
//
// If the result is NULL, all formats are expected to be supported.
//
// The function returns the following values:
//
//    - contentFormats (optional): supported data formats.
//
func (self *DropTarget) Formats() *gdk.ContentFormats {
	var _arg0 *C.GtkDropTarget     // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_drop_target_get_formats(_arg0)
	runtime.KeepAlive(self)

	var _contentFormats *gdk.ContentFormats // out

	if _cret != nil {
		_contentFormats = (*gdk.ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_contentFormats)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
			},
		)
	}

	return _contentFormats
}

// GTypes gets the list of supported GTypes for self.
//
// If no type have been set, NULL will be returned.
//
// The function returns the following values:
//
//    - gTypes (optional): G_TYPE_INVALID-terminated array of types included in
//      formats or NULL if none.
//
func (self *DropTarget) GTypes() []coreglib.Type {
	var _arg0 *C.GtkDropTarget // out
	var _cret *C.GType         // in
	var _arg1 C.gsize          // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_drop_target_get_gtypes(_arg0, &_arg1)
	runtime.KeepAlive(self)

	var _gTypes []coreglib.Type // out

	if _cret != nil {
		{
			src := unsafe.Slice((*C.GType)(_cret), _arg1)
			_gTypes = make([]coreglib.Type, _arg1)
			for i := 0; i < int(_arg1); i++ {
				_gTypes[i] = coreglib.Type(src[i])
			}
		}
	}

	return _gTypes
}

// Preload gets whether data should be preloaded on hover.
//
// The function returns the following values:
//
//    - ok: TRUE if drop data should be preloaded.
//
func (self *DropTarget) Preload() bool {
	var _arg0 *C.GtkDropTarget // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_drop_target_get_preload(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Value gets the current drop data, as a GValue.
//
// The function returns the following values:
//
//    - value (optional): current drop data.
//
func (self *DropTarget) Value() *coreglib.Value {
	var _arg0 *C.GtkDropTarget // out
	var _cret *C.GValue        // in

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_drop_target_get_value(_arg0)
	runtime.KeepAlive(self)

	var _value *coreglib.Value // out

	if _cret != nil {
		_value = coreglib.ValueFromNative(unsafe.Pointer(_cret))
	}

	return _value
}

// Reject rejects the ongoing drop operation.
//
// If no drop operation is ongoing, i.e when gtk.DropTarget:drop is NULL, this
// function does nothing.
//
// This function should be used when delaying the decision on whether to accept
// a drag or not until after reading the data.
func (self *DropTarget) Reject() {
	var _arg0 *C.GtkDropTarget // out

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.gtk_drop_target_reject(_arg0)
	runtime.KeepAlive(self)
}

// SetActions sets the actions that this drop target supports.
//
// The function takes the following parameters:
//
//    - actions: supported actions.
//
func (self *DropTarget) SetActions(actions gdk.DragAction) {
	var _arg0 *C.GtkDropTarget // out
	var _arg1 C.GdkDragAction  // out

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GdkDragAction(actions)

	C.gtk_drop_target_set_actions(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(actions)
}

// SetGTypes sets the supported GTypes for this drop target.
//
// The function takes the following parameters:
//
//    - types (optional): all supported #GTypes that can be dropped.
//
func (self *DropTarget) SetGTypes(types []coreglib.Type) {
	var _arg0 *C.GtkDropTarget // out
	var _arg1 *C.GType         // out
	var _arg2 C.gsize

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg2 = (C.gsize)(len(types))
	_arg1 = (*C.GType)(C.calloc(C.size_t(len(types)), C.size_t(C.sizeof_GType)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GType)(_arg1), len(types))
		for i := range types {
			out[i] = C.GType(types[i])
		}
	}

	C.gtk_drop_target_set_gtypes(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(types)
}

// SetPreload sets whether data should be preloaded on hover.
//
// The function takes the following parameters:
//
//    - preload: TRUE to preload drop data.
//
func (self *DropTarget) SetPreload(preload bool) {
	var _arg0 *C.GtkDropTarget // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkDropTarget)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if preload {
		_arg1 = C.TRUE
	}

	C.gtk_drop_target_set_preload(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(preload)
}
