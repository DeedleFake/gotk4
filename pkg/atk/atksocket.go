// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_socket_get_type()), F: marshalSocket},
	})
}

// SocketOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketOverrider interface {
	// Embed embeds the children of an Plug as the children of the Socket. The
	// plug may be in the same process or in a different process.
	//
	// The class item used by this function should be filled in by the IPC layer
	// (usually at-spi2-atk). The implementor of the AtkSocket should call this
	// function and pass the id for the plug as returned by atk_plug_get_id().
	// It is the responsibility of the application to pass the plug id on to the
	// process implementing the Socket as needed.
	Embed(plugId string)
}

// Socket: together with Plug, Socket provides the ability to embed accessibles
// from one process into another in a fashion that is transparent to assistive
// technologies. Socket works as the container of Plug, embedding it using the
// method atk_socket_embed(). Any accessible contained in the Plug will appear
// to the assistive technologies as being inside the application that created
// the Socket.
//
// The communication between a Socket and a Plug is done by the IPC layer of the
// accessibility framework, normally implemented by the D-Bus based
// implementation of AT-SPI (at-spi2). If that is the case, at-spi-atk2 is the
// responsible to implement the abstract methods atk_plug_get_id() and
// atk_socket_embed(), so an ATK implementor shouldn't reimplement them. The
// process that contains the Plug is responsible to send the ID returned by
// atk_plug_id() to the process that contains the Socket, so it could call the
// method atk_socket_embed() in order to embed it.
//
// For the same reasons, an implementor doesn't need to implement
// atk_object_get_n_accessible_children() and atk_object_ref_accessible_child().
// All the logic related to those functions will be implemented by the IPC
// layer.
type Socket interface {
	Object

	// AsObject casts the class to the Object interface.
	AsObject() Object
	// AsComponent casts the class to the Component interface.
	AsComponent() Component

	// AddRelationship adds a relationship of the specified type with the
	// specified target.
	//
	// This method is inherited from Object
	AddRelationship(relationship RelationType, target Object) bool
	// GetAccessibleID gets the accessible id of the accessible.
	//
	// This method is inherited from Object
	GetAccessibleID() string
	// GetDescription gets the accessible description of the accessible.
	//
	// This method is inherited from Object
	GetDescription() string
	// GetIndexInParent gets the 0-based index of this accessible in its parent;
	// returns -1 if the accessible does not have an accessible parent.
	//
	// This method is inherited from Object
	GetIndexInParent() int
	// GetLayer gets the layer of the accessible.
	//
	// Deprecated.
	//
	// This method is inherited from Object
	GetLayer() Layer
	// GetMDIZOrder gets the zorder of the accessible. The value G_MININT will
	// be returned if the layer of the accessible is not ATK_LAYER_MDI.
	//
	// Deprecated.
	//
	// This method is inherited from Object
	GetMDIZOrder() int
	// GetNAccessibleChildren gets the number of accessible children of the
	// accessible.
	//
	// This method is inherited from Object
	GetNAccessibleChildren() int
	// GetName gets the accessible name of the accessible.
	//
	// This method is inherited from Object
	GetName() string
	// GetObjectLocale gets a UTF-8 string indicating the POSIX-style
	// LC_MESSAGES locale of @accessible.
	//
	// This method is inherited from Object
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
	// This method is inherited from Object
	GetParent() Object
	// GetRole gets the role of the accessible.
	//
	// This method is inherited from Object
	GetRole() Role
	// Initialize: this function is called when implementing subclasses of
	// Object. It does initialization required for the new object. It is
	// intended that this function should called only in the ..._new() functions
	// used to create an instance of a subclass of Object
	//
	// This method is inherited from Object
	Initialize(data interface{})
	// PeekParent gets the accessible parent of the accessible, if it has been
	// manually assigned with atk_object_set_parent. Otherwise, this function
	// returns nil.
	//
	// This method is intended as an utility for ATK implementors, and not to be
	// exposed to accessible tools. See atk_object_get_parent() for further
	// reference.
	//
	// This method is inherited from Object
	PeekParent() Object
	// RefAccessibleChild gets a reference to the specified accessible child of
	// the object. The accessible children are 0-based so the first accessible
	// child is at index 0, the second at index 1 and so on.
	//
	// This method is inherited from Object
	RefAccessibleChild(i int) Object
	// RefRelationSet gets the RelationSet associated with the object.
	//
	// This method is inherited from Object
	RefRelationSet() RelationSet
	// RefStateSet gets a reference to the state set of the accessible; the
	// caller must unreference it when it is no longer needed.
	//
	// This method is inherited from Object
	RefStateSet() StateSet
	// RemovePropertyChangeHandler removes a property change handler.
	//
	// Deprecated: since version 2.12.
	//
	// This method is inherited from Object
	RemovePropertyChangeHandler(handlerId uint)
	// RemoveRelationship removes a relationship of the specified type with the
	// specified target.
	//
	// This method is inherited from Object
	RemoveRelationship(relationship RelationType, target Object) bool
	// SetAccessibleID sets the accessible ID of the accessible. This is not
	// meant to be presented to the user, but to be an ID which is stable over
	// application development. Typically, this is the gtkbuilder ID. Such an ID
	// will be available for instance to identify a given well-known accessible
	// object for tailored screen reading, or for automatic regression testing.
	//
	// This method is inherited from Object
	SetAccessibleID(name string)
	// SetDescription sets the accessible description of the accessible. You
	// can't set the description to NULL. This is reserved for the initial
	// value. In this aspect NULL is similar to ATK_ROLE_UNKNOWN. If you want to
	// set the name to a empty value you can use "".
	//
	// This method is inherited from Object
	SetDescription(description string)
	// SetName sets the accessible name of the accessible. You can't set the
	// name to NULL. This is reserved for the initial value. In this aspect NULL
	// is similar to ATK_ROLE_UNKNOWN. If you want to set the name to a empty
	// value you can use "".
	//
	// This method is inherited from Object
	SetName(name string)
	// SetParent sets the accessible parent of the accessible. @parent can be
	// NULL.
	//
	// This method is inherited from Object
	SetParent(parent Object)
	// SetRole sets the role of the accessible.
	//
	// This method is inherited from Object
	SetRole(role Role)
	// Contains checks whether the specified point is within the extent of the
	// @component.
	//
	// Toolkit implementor note: ATK provides a default implementation for this
	// virtual method. In general there are little reason to re-implement it.
	//
	// This method is inherited from Component
	Contains(x int, y int, coordType CoordType) bool
	// GetAlpha returns the alpha value (i.e. the opacity) for this @component,
	// on a scale from 0 (fully transparent) to 1.0 (fully opaque).
	//
	// This method is inherited from Component
	GetAlpha() float64
	// GetExtents gets the rectangle which gives the extent of the @component.
	//
	// If the extent can not be obtained (e.g. a non-embedded plug or missing
	// support), all of x, y, width, height are set to -1.
	//
	// This method is inherited from Component
	GetExtents(coordType CoordType) (x int, y int, width int, height int)
	// GetLayer gets the layer of the component.
	//
	// This method is inherited from Component
	GetLayer() Layer
	// GetMDIZOrder gets the zorder of the component. The value G_MININT will be
	// returned if the layer of the component is not ATK_LAYER_MDI or
	// ATK_LAYER_WINDOW.
	//
	// This method is inherited from Component
	GetMDIZOrder() int
	// GetPosition gets the position of @component in the form of a point
	// specifying @component's top-left corner.
	//
	// If the position can not be obtained (e.g. a non-embedded plug or missing
	// support), x and y are set to -1.
	//
	// Deprecated.
	//
	// This method is inherited from Component
	GetPosition(coordType CoordType) (x int, y int)
	// GetSize gets the size of the @component in terms of width and height.
	//
	// If the size can not be obtained (e.g. a non-embedded plug or missing
	// support), width and height are set to -1.
	//
	// Deprecated.
	//
	// This method is inherited from Component
	GetSize() (width int, height int)
	// GrabFocus grabs focus for this @component.
	//
	// This method is inherited from Component
	GrabFocus() bool
	// RefAccessibleAtPoint gets a reference to the accessible child, if one
	// exists, at the coordinate point specified by @x and @y.
	//
	// This method is inherited from Component
	RefAccessibleAtPoint(x int, y int, coordType CoordType) Object
	// RemoveFocusHandler: remove the handler specified by @handler_id from the
	// list of functions to be executed when this object receives focus events
	// (in or out).
	//
	// Deprecated: since version 2.9.4.
	//
	// This method is inherited from Component
	RemoveFocusHandler(handlerId uint)
	// ScrollTo makes @component visible on the screen by scrolling all
	// necessary parents.
	//
	// Contrary to atk_component_set_position, this does not actually move
	// @component in its parent, this only makes the parents scroll so that the
	// object shows up on the screen, given its current position within the
	// parents.
	//
	// This method is inherited from Component
	ScrollTo(typ ScrollType) bool
	// ScrollToPoint: move the top-left of @component to a given position of the
	// screen by scrolling all necessary parents.
	//
	// This method is inherited from Component
	ScrollToPoint(coords CoordType, x int, y int) bool
	// SetExtents sets the extents of @component.
	//
	// This method is inherited from Component
	SetExtents(x int, y int, width int, height int, coordType CoordType) bool
	// SetPosition sets the position of @component.
	//
	// Contrary to atk_component_scroll_to, this does not trigger any scrolling,
	// this just moves @component in its parent.
	//
	// This method is inherited from Component
	SetPosition(x int, y int, coordType CoordType) bool
	// SetSize: set the size of the @component in terms of width and height.
	//
	// This method is inherited from Component
	SetSize(width int, height int) bool

	// Embed embeds the children of an Plug as the children of the Socket. The
	// plug may be in the same process or in a different process.
	//
	// The class item used by this function should be filled in by the IPC layer
	// (usually at-spi2-atk). The implementor of the AtkSocket should call this
	// function and pass the id for the plug as returned by atk_plug_get_id().
	// It is the responsibility of the application to pass the plug id on to the
	// process implementing the Socket as needed.
	Embed(plugId string)
	// IsOccupied determines whether or not the socket has an embedded plug.
	IsOccupied() bool
}

// socket implements the Socket interface.
type socket struct {
	*externglib.Object
}

var _ Socket = (*socket)(nil)

// WrapSocket wraps a GObject to a type that implements
// interface Socket. It is primarily used internally.
func WrapSocket(obj *externglib.Object) Socket {
	return socket{obj}
}

func marshalSocket(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocket(obj), nil
}

// NewSocket creates a new Socket.
func NewSocket() Socket {
	var _cret *C.AtkObject // in

	_cret = C.atk_socket_new()

	var _socket Socket // out

	_socket = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Socket)

	return _socket
}

func (s socket) AsObject() Object {
	return WrapObject(gextras.InternObject(s))
}

func (s socket) AsComponent() Component {
	return WrapComponent(gextras.InternObject(s))
}

func (o socket) AddRelationship(relationship RelationType, target Object) bool {
	return WrapObject(gextras.InternObject(o)).AddRelationship(relationship, target)
}

func (a socket) GetAccessibleID() string {
	return WrapObject(gextras.InternObject(a)).GetAccessibleID()
}

func (a socket) GetDescription() string {
	return WrapObject(gextras.InternObject(a)).GetDescription()
}

func (a socket) GetIndexInParent() int {
	return WrapObject(gextras.InternObject(a)).GetIndexInParent()
}

func (a socket) GetLayer() Layer {
	return WrapObject(gextras.InternObject(a)).GetLayer()
}

func (a socket) GetMDIZOrder() int {
	return WrapObject(gextras.InternObject(a)).GetMDIZOrder()
}

func (a socket) GetNAccessibleChildren() int {
	return WrapObject(gextras.InternObject(a)).GetNAccessibleChildren()
}

func (a socket) GetName() string {
	return WrapObject(gextras.InternObject(a)).GetName()
}

func (a socket) GetObjectLocale() string {
	return WrapObject(gextras.InternObject(a)).GetObjectLocale()
}

func (a socket) GetParent() Object {
	return WrapObject(gextras.InternObject(a)).GetParent()
}

func (a socket) GetRole() Role {
	return WrapObject(gextras.InternObject(a)).GetRole()
}

func (a socket) Initialize(data interface{}) {
	WrapObject(gextras.InternObject(a)).Initialize(data)
}

func (a socket) PeekParent() Object {
	return WrapObject(gextras.InternObject(a)).PeekParent()
}

func (a socket) RefAccessibleChild(i int) Object {
	return WrapObject(gextras.InternObject(a)).RefAccessibleChild(i)
}

func (a socket) RefRelationSet() RelationSet {
	return WrapObject(gextras.InternObject(a)).RefRelationSet()
}

func (a socket) RefStateSet() StateSet {
	return WrapObject(gextras.InternObject(a)).RefStateSet()
}

func (a socket) RemovePropertyChangeHandler(handlerId uint) {
	WrapObject(gextras.InternObject(a)).RemovePropertyChangeHandler(handlerId)
}

func (o socket) RemoveRelationship(relationship RelationType, target Object) bool {
	return WrapObject(gextras.InternObject(o)).RemoveRelationship(relationship, target)
}

func (a socket) SetAccessibleID(name string) {
	WrapObject(gextras.InternObject(a)).SetAccessibleID(name)
}

func (a socket) SetDescription(description string) {
	WrapObject(gextras.InternObject(a)).SetDescription(description)
}

func (a socket) SetName(name string) {
	WrapObject(gextras.InternObject(a)).SetName(name)
}

func (a socket) SetParent(parent Object) {
	WrapObject(gextras.InternObject(a)).SetParent(parent)
}

func (a socket) SetRole(role Role) {
	WrapObject(gextras.InternObject(a)).SetRole(role)
}

func (c socket) Contains(x int, y int, coordType CoordType) bool {
	return WrapComponent(gextras.InternObject(c)).Contains(x, y, coordType)
}

func (c socket) GetAlpha() float64 {
	return WrapComponent(gextras.InternObject(c)).GetAlpha()
}

func (c socket) GetExtents(coordType CoordType) (x int, y int, width int, height int) {
	return WrapComponent(gextras.InternObject(c)).GetExtents(coordType)
}

func (c socket) GetLayer() Layer {
	return WrapComponent(gextras.InternObject(c)).GetLayer()
}

func (c socket) GetMDIZOrder() int {
	return WrapComponent(gextras.InternObject(c)).GetMDIZOrder()
}

func (c socket) GetPosition(coordType CoordType) (x int, y int) {
	return WrapComponent(gextras.InternObject(c)).GetPosition(coordType)
}

func (c socket) GetSize() (width int, height int) {
	return WrapComponent(gextras.InternObject(c)).GetSize()
}

func (c socket) GrabFocus() bool {
	return WrapComponent(gextras.InternObject(c)).GrabFocus()
}

func (c socket) RefAccessibleAtPoint(x int, y int, coordType CoordType) Object {
	return WrapComponent(gextras.InternObject(c)).RefAccessibleAtPoint(x, y, coordType)
}

func (c socket) RemoveFocusHandler(handlerId uint) {
	WrapComponent(gextras.InternObject(c)).RemoveFocusHandler(handlerId)
}

func (c socket) ScrollTo(typ ScrollType) bool {
	return WrapComponent(gextras.InternObject(c)).ScrollTo(typ)
}

func (c socket) ScrollToPoint(coords CoordType, x int, y int) bool {
	return WrapComponent(gextras.InternObject(c)).ScrollToPoint(coords, x, y)
}

func (c socket) SetExtents(x int, y int, width int, height int, coordType CoordType) bool {
	return WrapComponent(gextras.InternObject(c)).SetExtents(x, y, width, height, coordType)
}

func (c socket) SetPosition(x int, y int, coordType CoordType) bool {
	return WrapComponent(gextras.InternObject(c)).SetPosition(x, y, coordType)
}

func (c socket) SetSize(width int, height int) bool {
	return WrapComponent(gextras.InternObject(c)).SetSize(width, height)
}

func (o socket) Embed(plugId string) {
	var _arg0 *C.AtkSocket // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.AtkSocket)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.gchar)(C.CString(plugId))
	defer C.free(unsafe.Pointer(_arg1))

	C.atk_socket_embed(_arg0, _arg1)
}

func (o socket) IsOccupied() bool {
	var _arg0 *C.AtkSocket // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AtkSocket)(unsafe.Pointer(o.Native()))

	_cret = C.atk_socket_is_occupied(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
