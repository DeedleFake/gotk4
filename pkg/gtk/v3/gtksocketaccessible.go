// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_socket_accessible_get_type()), F: marshalSocketAccessible},
	})
}

type SocketAccessible interface {
	ContainerAccessible

	// AsContainerAccessible casts the class to the ContainerAccessible interface.
	AsContainerAccessible() ContainerAccessible

	// ConnectWidgetDestroyed: this function specifies the callback function to
	// be called when the widget corresponding to a GtkAccessible is destroyed.
	//
	// Deprecated: since version 3.4.
	//
	// This method is inherited from Accessible
	ConnectWidgetDestroyed()
	// GetWidget gets the Widget corresponding to the Accessible. The returned
	// widget does not have a reference added, so you do not need to unref it.
	//
	// This method is inherited from Accessible
	GetWidget() Widget
	// SetWidget sets the Widget corresponding to the Accessible.
	//
	// @accessible will not hold a reference to @widget. It is the caller’s
	// responsibility to ensure that when @widget is destroyed, the widget is
	// unset by calling this function again with @widget set to nil.
	//
	// This method is inherited from Accessible
	SetWidget(widget Widget)
	// AddRelationship adds a relationship of the specified type with the
	// specified target.
	//
	// This method is inherited from atk.Object
	AddRelationship(relationship atk.RelationType, target atk.Object) bool
	// GetAccessibleID gets the accessible id of the accessible.
	//
	// This method is inherited from atk.Object
	GetAccessibleID() string
	// GetDescription gets the accessible description of the accessible.
	//
	// This method is inherited from atk.Object
	GetDescription() string
	// GetIndexInParent gets the 0-based index of this accessible in its parent;
	// returns -1 if the accessible does not have an accessible parent.
	//
	// This method is inherited from atk.Object
	GetIndexInParent() int
	// GetLayer gets the layer of the accessible.
	//
	// Deprecated.
	//
	// This method is inherited from atk.Object
	GetLayer() atk.Layer
	// GetMDIZOrder gets the zorder of the accessible. The value G_MININT will
	// be returned if the layer of the accessible is not ATK_LAYER_MDI.
	//
	// Deprecated.
	//
	// This method is inherited from atk.Object
	GetMDIZOrder() int
	// GetNAccessibleChildren gets the number of accessible children of the
	// accessible.
	//
	// This method is inherited from atk.Object
	GetNAccessibleChildren() int
	// GetName gets the accessible name of the accessible.
	//
	// This method is inherited from atk.Object
	GetName() string
	// GetObjectLocale gets a UTF-8 string indicating the POSIX-style
	// LC_MESSAGES locale of @accessible.
	//
	// This method is inherited from atk.Object
	GetObjectLocale() string
	// GetParent gets the accessible parent of the accessible. By default this
	// is the one assigned with atk_object_set_parent(), but it is assumed that
	// ATK implementors have ways to get the parent of the object without the
	// need of assigning it manually with atk_object_set_parent(), and will
	// return it with this method.
	//
	// If you are only interested on the parent assigned with
	// atk_object_set_parent(), use atk_object_peek_parent().
	//
	// This method is inherited from atk.Object
	GetParent() atk.Object
	// GetRole gets the role of the accessible.
	//
	// This method is inherited from atk.Object
	GetRole() atk.Role
	// Initialize: this function is called when implementing subclasses of
	// Object. It does initialization required for the new object. It is
	// intended that this function should called only in the ..._new() functions
	// used to create an instance of a subclass of Object
	//
	// This method is inherited from atk.Object
	Initialize(data interface{})
	// PeekParent gets the accessible parent of the accessible, if it has been
	// manually assigned with atk_object_set_parent. Otherwise, this function
	// returns nil.
	//
	// This method is intended as an utility for ATK implementors, and not to be
	// exposed to accessible tools. See atk_object_get_parent() for further
	// reference.
	//
	// This method is inherited from atk.Object
	PeekParent() atk.Object
	// RefAccessibleChild gets a reference to the specified accessible child of
	// the object. The accessible children are 0-based so the first accessible
	// child is at index 0, the second at index 1 and so on.
	//
	// This method is inherited from atk.Object
	RefAccessibleChild(i int) atk.Object
	// RefRelationSet gets the RelationSet associated with the object.
	//
	// This method is inherited from atk.Object
	RefRelationSet() atk.RelationSet
	// RefStateSet gets a reference to the state set of the accessible; the
	// caller must unreference it when it is no longer needed.
	//
	// This method is inherited from atk.Object
	RefStateSet() atk.StateSet
	// RemovePropertyChangeHandler removes a property change handler.
	//
	// Deprecated: since version 2.12.
	//
	// This method is inherited from atk.Object
	RemovePropertyChangeHandler(handlerId uint)
	// RemoveRelationship removes a relationship of the specified type with the
	// specified target.
	//
	// This method is inherited from atk.Object
	RemoveRelationship(relationship atk.RelationType, target atk.Object) bool
	// SetAccessibleID sets the accessible ID of the accessible. This is not
	// meant to be presented to the user, but to be an ID which is stable over
	// application development. Typically, this is the gtkbuilder ID. Such an ID
	// will be available for instance to identify a given well-known accessible
	// object for tailored screen reading, or for automatic regression testing.
	//
	// This method is inherited from atk.Object
	SetAccessibleID(name string)
	// SetDescription sets the accessible description of the accessible. You
	// can't set the description to NULL. This is reserved for the initial
	// value. In this aspect NULL is similar to ATK_ROLE_UNKNOWN. If you want to
	// set the name to a empty value you can use "".
	//
	// This method is inherited from atk.Object
	SetDescription(description string)
	// SetName sets the accessible name of the accessible. You can't set the
	// name to NULL. This is reserved for the initial value. In this aspect NULL
	// is similar to ATK_ROLE_UNKNOWN. If you want to set the name to a empty
	// value you can use "".
	//
	// This method is inherited from atk.Object
	SetName(name string)
	// SetParent sets the accessible parent of the accessible. @parent can be
	// NULL.
	//
	// This method is inherited from atk.Object
	SetParent(parent atk.Object)
	// SetRole sets the role of the accessible.
	//
	// This method is inherited from atk.Object
	SetRole(role atk.Role)

	Embed(path string)
}

// socketAccessible implements the SocketAccessible interface.
type socketAccessible struct {
	*externglib.Object
}

var _ SocketAccessible = (*socketAccessible)(nil)

// WrapSocketAccessible wraps a GObject to a type that implements
// interface SocketAccessible. It is primarily used internally.
func WrapSocketAccessible(obj *externglib.Object) SocketAccessible {
	return socketAccessible{obj}
}

func marshalSocketAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketAccessible(obj), nil
}

func (s socketAccessible) AsContainerAccessible() ContainerAccessible {
	return WrapContainerAccessible(gextras.InternObject(s))
}

func (a socketAccessible) ConnectWidgetDestroyed() {
	WrapAccessible(gextras.InternObject(a)).ConnectWidgetDestroyed()
}

func (a socketAccessible) GetWidget() Widget {
	return WrapAccessible(gextras.InternObject(a)).GetWidget()
}

func (a socketAccessible) SetWidget(widget Widget) {
	WrapAccessible(gextras.InternObject(a)).SetWidget(widget)
}

func (o socketAccessible) AddRelationship(relationship atk.RelationType, target atk.Object) bool {
	return atk.WrapObject(gextras.InternObject(o)).AddRelationship(relationship, target)
}

func (a socketAccessible) GetAccessibleID() string {
	return atk.WrapObject(gextras.InternObject(a)).GetAccessibleID()
}

func (a socketAccessible) GetDescription() string {
	return atk.WrapObject(gextras.InternObject(a)).GetDescription()
}

func (a socketAccessible) GetIndexInParent() int {
	return atk.WrapObject(gextras.InternObject(a)).GetIndexInParent()
}

func (a socketAccessible) GetLayer() atk.Layer {
	return atk.WrapObject(gextras.InternObject(a)).GetLayer()
}

func (a socketAccessible) GetMDIZOrder() int {
	return atk.WrapObject(gextras.InternObject(a)).GetMDIZOrder()
}

func (a socketAccessible) GetNAccessibleChildren() int {
	return atk.WrapObject(gextras.InternObject(a)).GetNAccessibleChildren()
}

func (a socketAccessible) GetName() string {
	return atk.WrapObject(gextras.InternObject(a)).GetName()
}

func (a socketAccessible) GetObjectLocale() string {
	return atk.WrapObject(gextras.InternObject(a)).GetObjectLocale()
}

func (a socketAccessible) GetParent() atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).GetParent()
}

func (a socketAccessible) GetRole() atk.Role {
	return atk.WrapObject(gextras.InternObject(a)).GetRole()
}

func (a socketAccessible) Initialize(data interface{}) {
	atk.WrapObject(gextras.InternObject(a)).Initialize(data)
}

func (a socketAccessible) PeekParent() atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).PeekParent()
}

func (a socketAccessible) RefAccessibleChild(i int) atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).RefAccessibleChild(i)
}

func (a socketAccessible) RefRelationSet() atk.RelationSet {
	return atk.WrapObject(gextras.InternObject(a)).RefRelationSet()
}

func (a socketAccessible) RefStateSet() atk.StateSet {
	return atk.WrapObject(gextras.InternObject(a)).RefStateSet()
}

func (a socketAccessible) RemovePropertyChangeHandler(handlerId uint) {
	atk.WrapObject(gextras.InternObject(a)).RemovePropertyChangeHandler(handlerId)
}

func (o socketAccessible) RemoveRelationship(relationship atk.RelationType, target atk.Object) bool {
	return atk.WrapObject(gextras.InternObject(o)).RemoveRelationship(relationship, target)
}

func (a socketAccessible) SetAccessibleID(name string) {
	atk.WrapObject(gextras.InternObject(a)).SetAccessibleID(name)
}

func (a socketAccessible) SetDescription(description string) {
	atk.WrapObject(gextras.InternObject(a)).SetDescription(description)
}

func (a socketAccessible) SetName(name string) {
	atk.WrapObject(gextras.InternObject(a)).SetName(name)
}

func (a socketAccessible) SetParent(parent atk.Object) {
	atk.WrapObject(gextras.InternObject(a)).SetParent(parent)
}

func (a socketAccessible) SetRole(role atk.Role) {
	atk.WrapObject(gextras.InternObject(a)).SetRole(role)
}

func (s socketAccessible) Embed(path string) {
	var _arg0 *C.GtkSocketAccessible // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkSocketAccessible)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_socket_accessible_embed(_arg0, _arg1)
}
