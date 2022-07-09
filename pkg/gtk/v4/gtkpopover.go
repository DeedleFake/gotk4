// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_PopoverClass_activate_default(void*);
// extern void _gotk4_gtk4_PopoverClass_closed(void*);
// extern void _gotk4_gtk4_Popover_ConnectActivateDefault(gpointer, guintptr);
// extern void _gotk4_gtk4_Popover_ConnectClosed(gpointer, guintptr);
import "C"

// GTypePopover returns the GType for the type Popover.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePopover() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Popover").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPopover)
	return gtype
}

// PopoverOverrider contains methods that are overridable.
type PopoverOverrider interface {
	ActivateDefault()
	Closed()
}

// Popover: GtkPopover is a bubble-like context popup.
//
// !An example GtkPopover (popover.png)
//
// It is primarily meant to provide context-dependent information or options.
// Popovers are attached to a parent widget. By default, they point to the whole
// widget area, although this behavior can be changed with
// gtk.Popover.SetPointingTo().
//
// The position of a popover relative to the widget it is attached to can also
// be changed with gtk.Popover.SetPosition()
//
// By default, GtkPopover performs a grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Escape key being
// pressed). If no such modal behavior is desired on a popover,
// gtk.Popover.SetAutohide() may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// GtkPopover is often used to replace menus. The best was to do this is to use
// the gtk.PopoverMenu subclass which supports being populated from a GMenuModel
// with gtk.PopoverMenu.NewFromModel.
//
//    <section>
//      <attribute name="display-hint">horizontal-buttons</attribute>
//      <item>
//        <attribute name="label">Cut</attribute>
//        <attribute name="action">app.cut</attribute>
//        <attribute name="verb-icon">edit-cut-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Copy</attribute>
//        <attribute name="action">app.copy</attribute>
//        <attribute name="verb-icon">edit-copy-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Paste</attribute>
//        <attribute name="action">app.paste</attribute>
//        <attribute name="verb-icon">edit-paste-symbolic</attribute>
//      </item>
//    </section>
//
//
// CSS nodes
//
//    popover[.menu]
//    ├── arrow
//    ╰── contents.background
//        ╰── <child>
//
//
// The contents child node always gets the .background style class and the
// popover itself gets the .menu style class if the popover is menu-like (i.e.
// GtkPopoverMenu).
//
// Particular uses of GtkPopover, such as touch selection popups or magnifiers
// in GtkEntry or GtkTextView get style classes like .touch-selection or
// .magnifier to differentiate from plain popovers.
//
// When styling a popover directly, the popover node should usually not have any
// background. The visible part of the popover can have a shadow. To specify it
// in CSS, set the box-shadow of the contents node.
//
// Note that, in order to accomplish appropriate arrow visuals, GtkPopover uses
// custom drawing for the arrow node. This makes it possible for the arrow to
// change its shape dynamically, but it also limits the possibilities of styling
// it using CSS. In particular, the arrow gets drawn over the content node's
// border and shadow, so they look like one shape, which means that the border
// width of the content node and the arrow node should be the same. The arrow
// also does not support any border shape other than solid, no border-radius,
// only one border width (border-bottom-width is used) and no box-shadow.
type Popover struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	NativeSurface
	ShortcutManager
}

var (
	_ Widgetter         = (*Popover)(nil)
	_ coreglib.Objector = (*Popover)(nil)
)

func classInitPopoverer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "PopoverClass")

	if _, ok := goval.(interface{ ActivateDefault() }); ok {
		o := pclass.StructFieldOffset("activate_default")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_PopoverClass_activate_default)
	}

	if _, ok := goval.(interface{ Closed() }); ok {
		o := pclass.StructFieldOffset("closed")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk4_PopoverClass_closed)
	}
}

//export _gotk4_gtk4_PopoverClass_activate_default
func _gotk4_gtk4_PopoverClass_activate_default(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ActivateDefault() })

	iface.ActivateDefault()
}

//export _gotk4_gtk4_PopoverClass_closed
func _gotk4_gtk4_PopoverClass_closed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Closed() })

	iface.Closed()
}

func wrapPopover(obj *coreglib.Object) *Popover {
	return &Popover{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		NativeSurface: NativeSurface{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
		ShortcutManager: ShortcutManager{
			Object: obj,
		},
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	return wrapPopover(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Popover_ConnectActivateDefault
func _gotk4_gtk4_Popover_ConnectActivateDefault(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectActivateDefault is emitted whend the user activates the default
// widget.
//
// This is a keybinding signal (class.SignalAction.html).
func (popover *Popover) ConnectActivateDefault(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(popover, "activate-default", false, unsafe.Pointer(C._gotk4_gtk4_Popover_ConnectActivateDefault), f)
}

//export _gotk4_gtk4_Popover_ConnectClosed
func _gotk4_gtk4_Popover_ConnectClosed(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectClosed is emitted when the popover is closed.
func (popover *Popover) ConnectClosed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(popover, "closed", false, unsafe.Pointer(C._gotk4_gtk4_Popover_ConnectClosed), f)
}

// NewPopover creates a new GtkPopover.
//
// The function returns the following values:
//
//    - popover: new GtkPopover.
//
func NewPopover() *Popover {
	_info := girepository.MustFind("Gtk", "Popover")
	_gret := _info.InvokeClassMethod("new_Popover", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _popover *Popover // out

	_popover = wrapPopover(coreglib.Take(unsafe.Pointer(_cret)))

	return _popover
}

// Autohide returns whether the popover is modal.
//
// See gtk.Popover.SetAutohide() for the implications of this.
//
// The function returns the following values:
//
//    - ok: TRUE if popover is modal.
//
func (popover *Popover) Autohide() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_gret := _info.InvokeClassMethod("get_autohide", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// CascadePopdown returns whether the popover will close after a modal child is
// closed.
//
// The function returns the following values:
//
//    - ok: TRUE if popover will close after a modal child.
//
func (popover *Popover) CascadePopdown() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_gret := _info.InvokeClassMethod("get_cascade_popdown", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child widget of popover.
//
// The function returns the following values:
//
//    - widget (optional): child widget of popover.
//
func (popover *Popover) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_gret := _info.InvokeClassMethod("get_child", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// HasArrow gets whether this popover is showing an arrow pointing at the widget
// that it is relative to.
//
// The function returns the following values:
//
//    - ok: whether the popover has an arrow.
//
func (popover *Popover) HasArrow() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_gret := _info.InvokeClassMethod("get_has_arrow", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// MnemonicsVisible gets whether mnemonics are visible.
//
// The function returns the following values:
//
//    - ok: TRUE if mnemonics are supposed to be visible in this popover.
//
func (popover *Popover) MnemonicsVisible() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_gret := _info.InvokeClassMethod("get_mnemonics_visible", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Offset gets the offset previous set with gtk_popover_set_offset().
//
// The function returns the following values:
//
//    - xOffset (optional): location for the x_offset.
//    - yOffset (optional): location for the y_offset.
//
func (popover *Popover) Offset() (xOffset, yOffset int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("get_offset", _args[:], _outs[:])

	runtime.KeepAlive(popover)

	var _xOffset int32 // out
	var _yOffset int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_xOffset = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_yOffset = *(*int32)(unsafe.Pointer(_outs[1]))
	}

	return _xOffset, _yOffset
}

// PointingTo gets the rectangle that the popover points to.
//
// If a rectangle to point to has been set, this function will return TRUE and
// fill in rect with such rectangle, otherwise it will return FALSE and fill in
// rect with the parent widget coordinates.
//
// The function returns the following values:
//
//    - rect: location to store the rectangle.
//    - ok: TRUE if a rectangle to point to was set.
//
func (popover *Popover) PointingTo() (*gdk.Rectangle, bool) {
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_gret := _info.InvokeClassMethod("get_pointing_to", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(_outs[0])))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _rect, _ok
}

// Popdown pops popover down.
//
// This is different from a gtk.Widget.Hide() call in that it may show the
// popover with a transition. If you want to hide the popover without a
// transition, just use gtk.Widget.Hide().
func (popover *Popover) Popdown() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("popdown", _args[:], nil)

	runtime.KeepAlive(popover)
}

// Popup pops popover up.
//
// This is different from a gtk.Widget.Show() call in that it may show the
// popover with a transition(). If you want to show the popover without a
// transition, just use [methodGtk.Widget.show.
func (popover *Popover) Popup() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("popup", _args[:], nil)

	runtime.KeepAlive(popover)
}

// Present presents the popover to the user.
func (popover *Popover) Present() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("present", _args[:], nil)

	runtime.KeepAlive(popover)
}

// SetAutohide sets whether popover is modal.
//
// A modal popover will grab the keyboard focus on it when being displayed.
// Clicking outside the popover area or pressing Esc will dismiss the popover.
//
// Called this function on an already showing popup with a new autohide value
// different from the current one, will cause the popup to be hidden.
//
// The function takes the following parameters:
//
//    - autohide: TRUE to dismiss the popover on outside clicks.
//
func (popover *Popover) SetAutohide(autohide bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if autohide {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("set_autohide", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(autohide)
}

// SetCascadePopdown: if cascade_popdown is TRUE, the popover will be closed
// when a child modal popover is closed.
//
// If FALSE, popover will stay visible.
//
// The function takes the following parameters:
//
//    - cascadePopdown: TRUE if the popover should follow a child closing.
//
func (popover *Popover) SetCascadePopdown(cascadePopdown bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if cascadePopdown {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("set_cascade_popdown", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(cascadePopdown)
}

// SetChild sets the child widget of popover.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (popover *Popover) SetChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("set_child", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(child)
}

// SetDefaultWidget sets the default widget of a GtkPopover.
//
// The default widget is the widget that’s activated when the user presses Enter
// in a dialog (for example). This function sets or unsets the default widget
// for a GtkPopover.
//
// The function takes the following parameters:
//
//    - widget (optional): child widget of popover to set as the default, or NULL
//      to unset the default widget for the popover.
//
func (popover *Popover) SetDefaultWidget(widget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if widget != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("set_default_widget", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(widget)
}

// SetHasArrow sets whether this popover should draw an arrow pointing at the
// widget it is relative to.
//
// The function takes the following parameters:
//
//    - hasArrow: TRUE to draw an arrow.
//
func (popover *Popover) SetHasArrow(hasArrow bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if hasArrow {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("set_has_arrow", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(hasArrow)
}

// SetMnemonicsVisible sets whether mnemonics should be visible.
//
// The function takes the following parameters:
//
//    - mnemonicsVisible: new value.
//
func (popover *Popover) SetMnemonicsVisible(mnemonicsVisible bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if mnemonicsVisible {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("set_mnemonics_visible", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(mnemonicsVisible)
}

// SetOffset sets the offset to use when calculating the position of the
// popover.
//
// These values are used when preparing the gdk.PopupLayout for positioning the
// popover.
//
// The function takes the following parameters:
//
//    - xOffset: x offset to adjust the position by.
//    - yOffset: y offset to adjust the position by.
//
func (popover *Popover) SetOffset(xOffset, yOffset int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(xOffset)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(yOffset)

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("set_offset", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(xOffset)
	runtime.KeepAlive(yOffset)
}

// SetPointingTo sets the rectangle that popover points to.
//
// This is in the coordinate space of the popover parent.
//
// The function takes the following parameters:
//
//    - rect: rectangle to point to.
//
func (popover *Popover) SetPointingTo(rect *gdk.Rectangle) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(rect)))

	_info := girepository.MustFind("Gtk", "Popover")
	_info.InvokeClassMethod("set_pointing_to", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(rect)
}
