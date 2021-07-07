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
		{T: externglib.Type(C.gtk_image_cell_accessible_get_type()), F: marshalImageCellAccessible},
	})
}

type ImageCellAccessible interface {
	RendererCellAccessible

	// AsRendererCellAccessible casts the class to the RendererCellAccessible interface.
	AsRendererCellAccessible() RendererCellAccessible
	// AsAction casts the class to the atk.Action interface.
	AsAction() atk.Action
	// AsImage casts the class to the atk.Image interface.
	AsImage() atk.Image

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
	// DoAction: perform the specified action on the object.
	//
	// This method is inherited from atk.Action
	DoAction(i int) bool
	// GetDescription returns a description of the specified action of the
	// object.
	//
	// This method is inherited from atk.Action
	GetDescription(i int) string
	// GetKeybinding gets the keybinding which can be used to activate this
	// action, if one exists. The string returned should contain localized,
	// human-readable, key sequences as they would appear when displayed on
	// screen. It must be in the format "mnemonic;sequence;shortcut".
	//
	// - The mnemonic key activates the object if it is presently enabled
	// onscreen. This typically corresponds to the underlined letter within the
	// widget. Example: "n" in a traditional "New..." menu item or the "a" in
	// "Apply" for a button. - The sequence is the full list of keys which
	// invoke the action even if the relevant element is not currently shown on
	// screen. For instance, for a menu item the sequence is the keybindings
	// used to open the parent menus before invoking. The sequence string is
	// colon-delimited. Example: "Alt+F:N" in a traditional "New..." menu item.
	// - The shortcut, if it exists, will invoke the same action without showing
	// the component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
	// traditional "New..." menu item.
	//
	// Example: For a traditional "New..." menu item, the expected return value
	// would be: "N;Alt+F:N;Ctrl+N" for the English locale and
	// "N;Alt+D:N;Strg+N" for the German locale. If, hypothetically, this menu
	// item lacked a mnemonic, it would be represented by ";;Ctrl+N" and
	// ";;Strg+N" respectively.
	//
	// This method is inherited from atk.Action
	GetKeybinding(i int) string
	// GetLocalizedName returns the localized name of the specified action of
	// the object.
	//
	// This method is inherited from atk.Action
	GetLocalizedName(i int) string
	// GetNActions gets the number of accessible actions available on the
	// object. If there are more than one, the first one is considered the
	// "default" action of the object.
	//
	// This method is inherited from atk.Action
	GetNActions() int
	// GetName returns a non-localized string naming the specified action of the
	// object. This name is generally not descriptive of the end result of the
	// action, but instead names the 'interaction type' which the object
	// supports. By convention, the above strings should be used to represent
	// the actions which correspond to the common point-and-click interaction
	// techniques of the same name: i.e. "click", "press", "release", "drag",
	// "drop", "popup", etc. The "popup" action should be used to pop up a
	// context menu for the object, if one exists.
	//
	// For technical reasons, some toolkits cannot guarantee that the reported
	// action is actually 'bound' to a nontrivial user event; i.e. the result of
	// some actions via atk_action_do_action() may be NIL.
	//
	// This method is inherited from atk.Action
	GetName(i int) string
	// SetDescription sets a description of the specified action of the object.
	//
	// This method is inherited from atk.Action
	SetDescription(i int, desc string) bool
	// DoAction: perform the specified action on the object.
	//
	// This method is inherited from atk.Action
	DoAction(i int) bool
	// GetDescription returns a description of the specified action of the
	// object.
	//
	// This method is inherited from atk.Action
	GetDescription(i int) string
	// GetKeybinding gets the keybinding which can be used to activate this
	// action, if one exists. The string returned should contain localized,
	// human-readable, key sequences as they would appear when displayed on
	// screen. It must be in the format "mnemonic;sequence;shortcut".
	//
	// - The mnemonic key activates the object if it is presently enabled
	// onscreen. This typically corresponds to the underlined letter within the
	// widget. Example: "n" in a traditional "New..." menu item or the "a" in
	// "Apply" for a button. - The sequence is the full list of keys which
	// invoke the action even if the relevant element is not currently shown on
	// screen. For instance, for a menu item the sequence is the keybindings
	// used to open the parent menus before invoking. The sequence string is
	// colon-delimited. Example: "Alt+F:N" in a traditional "New..." menu item.
	// - The shortcut, if it exists, will invoke the same action without showing
	// the component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
	// traditional "New..." menu item.
	//
	// Example: For a traditional "New..." menu item, the expected return value
	// would be: "N;Alt+F:N;Ctrl+N" for the English locale and
	// "N;Alt+D:N;Strg+N" for the German locale. If, hypothetically, this menu
	// item lacked a mnemonic, it would be represented by ";;Ctrl+N" and
	// ";;Strg+N" respectively.
	//
	// This method is inherited from atk.Action
	GetKeybinding(i int) string
	// GetLocalizedName returns the localized name of the specified action of
	// the object.
	//
	// This method is inherited from atk.Action
	GetLocalizedName(i int) string
	// GetNActions gets the number of accessible actions available on the
	// object. If there are more than one, the first one is considered the
	// "default" action of the object.
	//
	// This method is inherited from atk.Action
	GetNActions() int
	// GetName returns a non-localized string naming the specified action of the
	// object. This name is generally not descriptive of the end result of the
	// action, but instead names the 'interaction type' which the object
	// supports. By convention, the above strings should be used to represent
	// the actions which correspond to the common point-and-click interaction
	// techniques of the same name: i.e. "click", "press", "release", "drag",
	// "drop", "popup", etc. The "popup" action should be used to pop up a
	// context menu for the object, if one exists.
	//
	// For technical reasons, some toolkits cannot guarantee that the reported
	// action is actually 'bound' to a nontrivial user event; i.e. the result of
	// some actions via atk_action_do_action() may be NIL.
	//
	// This method is inherited from atk.Action
	GetName(i int) string
	// SetDescription sets a description of the specified action of the object.
	//
	// This method is inherited from atk.Action
	SetDescription(i int, desc string) bool
	// DoAction: perform the specified action on the object.
	//
	// This method is inherited from atk.Action
	DoAction(i int) bool
	// GetDescription returns a description of the specified action of the
	// object.
	//
	// This method is inherited from atk.Action
	GetDescription(i int) string
	// GetKeybinding gets the keybinding which can be used to activate this
	// action, if one exists. The string returned should contain localized,
	// human-readable, key sequences as they would appear when displayed on
	// screen. It must be in the format "mnemonic;sequence;shortcut".
	//
	// - The mnemonic key activates the object if it is presently enabled
	// onscreen. This typically corresponds to the underlined letter within the
	// widget. Example: "n" in a traditional "New..." menu item or the "a" in
	// "Apply" for a button. - The sequence is the full list of keys which
	// invoke the action even if the relevant element is not currently shown on
	// screen. For instance, for a menu item the sequence is the keybindings
	// used to open the parent menus before invoking. The sequence string is
	// colon-delimited. Example: "Alt+F:N" in a traditional "New..." menu item.
	// - The shortcut, if it exists, will invoke the same action without showing
	// the component or its enclosing menus or dialogs. Example: "Ctrl+N" in a
	// traditional "New..." menu item.
	//
	// Example: For a traditional "New..." menu item, the expected return value
	// would be: "N;Alt+F:N;Ctrl+N" for the English locale and
	// "N;Alt+D:N;Strg+N" for the German locale. If, hypothetically, this menu
	// item lacked a mnemonic, it would be represented by ";;Ctrl+N" and
	// ";;Strg+N" respectively.
	//
	// This method is inherited from atk.Action
	GetKeybinding(i int) string
	// GetLocalizedName returns the localized name of the specified action of
	// the object.
	//
	// This method is inherited from atk.Action
	GetLocalizedName(i int) string
	// GetNActions gets the number of accessible actions available on the
	// object. If there are more than one, the first one is considered the
	// "default" action of the object.
	//
	// This method is inherited from atk.Action
	GetNActions() int
	// GetName returns a non-localized string naming the specified action of the
	// object. This name is generally not descriptive of the end result of the
	// action, but instead names the 'interaction type' which the object
	// supports. By convention, the above strings should be used to represent
	// the actions which correspond to the common point-and-click interaction
	// techniques of the same name: i.e. "click", "press", "release", "drag",
	// "drop", "popup", etc. The "popup" action should be used to pop up a
	// context menu for the object, if one exists.
	//
	// For technical reasons, some toolkits cannot guarantee that the reported
	// action is actually 'bound' to a nontrivial user event; i.e. the result of
	// some actions via atk_action_do_action() may be NIL.
	//
	// This method is inherited from atk.Action
	GetName(i int) string
	// SetDescription sets a description of the specified action of the object.
	//
	// This method is inherited from atk.Action
	SetDescription(i int, desc string) bool
	// GetImageDescription: get a textual description of this image.
	//
	// This method is inherited from atk.Image
	GetImageDescription() string
	// GetImageLocale retrieves the locale identifier associated to the Image.
	//
	// This method is inherited from atk.Image
	GetImageLocale() string
	// GetImagePosition gets the position of the image in the form of a point
	// specifying the images top-left corner.
	//
	// If the position can not be obtained (e.g. missing support), x and y are
	// set to -1.
	//
	// This method is inherited from atk.Image
	GetImagePosition(coordType atk.CoordType) (x int, y int)
	// GetImageSize: get the width and height in pixels for the specified image.
	// The values of @width and @height are returned as -1 if the values cannot
	// be obtained (for instance, if the object is not onscreen).
	//
	// If the size can not be obtained (e.g. missing support), x and y are set
	// to -1.
	//
	// This method is inherited from atk.Image
	GetImageSize() (width int, height int)
	// SetImageDescription sets the textual description for this image.
	//
	// This method is inherited from atk.Image
	SetImageDescription(description string) bool
}

// imageCellAccessible implements the ImageCellAccessible interface.
type imageCellAccessible struct {
	*externglib.Object
}

var _ ImageCellAccessible = (*imageCellAccessible)(nil)

// WrapImageCellAccessible wraps a GObject to a type that implements
// interface ImageCellAccessible. It is primarily used internally.
func WrapImageCellAccessible(obj *externglib.Object) ImageCellAccessible {
	return imageCellAccessible{obj}
}

func marshalImageCellAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapImageCellAccessible(obj), nil
}

func (i imageCellAccessible) AsRendererCellAccessible() RendererCellAccessible {
	return WrapRendererCellAccessible(gextras.InternObject(i))
}

func (i imageCellAccessible) AsAction() atk.Action {
	return atk.WrapAction(gextras.InternObject(i))
}

func (i imageCellAccessible) AsImage() atk.Image {
	return atk.WrapImage(gextras.InternObject(i))
}

func (a imageCellAccessible) ConnectWidgetDestroyed() {
	WrapAccessible(gextras.InternObject(a)).ConnectWidgetDestroyed()
}

func (a imageCellAccessible) GetWidget() Widget {
	return WrapAccessible(gextras.InternObject(a)).GetWidget()
}

func (a imageCellAccessible) SetWidget(widget Widget) {
	WrapAccessible(gextras.InternObject(a)).SetWidget(widget)
}

func (o imageCellAccessible) AddRelationship(relationship atk.RelationType, target atk.Object) bool {
	return atk.WrapObject(gextras.InternObject(o)).AddRelationship(relationship, target)
}

func (a imageCellAccessible) GetAccessibleID() string {
	return atk.WrapObject(gextras.InternObject(a)).GetAccessibleID()
}

func (a imageCellAccessible) GetDescription() string {
	return atk.WrapObject(gextras.InternObject(a)).GetDescription()
}

func (a imageCellAccessible) GetIndexInParent() int {
	return atk.WrapObject(gextras.InternObject(a)).GetIndexInParent()
}

func (a imageCellAccessible) GetLayer() atk.Layer {
	return atk.WrapObject(gextras.InternObject(a)).GetLayer()
}

func (a imageCellAccessible) GetMDIZOrder() int {
	return atk.WrapObject(gextras.InternObject(a)).GetMDIZOrder()
}

func (a imageCellAccessible) GetNAccessibleChildren() int {
	return atk.WrapObject(gextras.InternObject(a)).GetNAccessibleChildren()
}

func (a imageCellAccessible) GetName() string {
	return atk.WrapObject(gextras.InternObject(a)).GetName()
}

func (a imageCellAccessible) GetObjectLocale() string {
	return atk.WrapObject(gextras.InternObject(a)).GetObjectLocale()
}

func (a imageCellAccessible) GetParent() atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).GetParent()
}

func (a imageCellAccessible) GetRole() atk.Role {
	return atk.WrapObject(gextras.InternObject(a)).GetRole()
}

func (a imageCellAccessible) Initialize(data interface{}) {
	atk.WrapObject(gextras.InternObject(a)).Initialize(data)
}

func (a imageCellAccessible) PeekParent() atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).PeekParent()
}

func (a imageCellAccessible) RefAccessibleChild(i int) atk.Object {
	return atk.WrapObject(gextras.InternObject(a)).RefAccessibleChild(i)
}

func (a imageCellAccessible) RefRelationSet() atk.RelationSet {
	return atk.WrapObject(gextras.InternObject(a)).RefRelationSet()
}

func (a imageCellAccessible) RefStateSet() atk.StateSet {
	return atk.WrapObject(gextras.InternObject(a)).RefStateSet()
}

func (a imageCellAccessible) RemovePropertyChangeHandler(handlerId uint) {
	atk.WrapObject(gextras.InternObject(a)).RemovePropertyChangeHandler(handlerId)
}

func (o imageCellAccessible) RemoveRelationship(relationship atk.RelationType, target atk.Object) bool {
	return atk.WrapObject(gextras.InternObject(o)).RemoveRelationship(relationship, target)
}

func (a imageCellAccessible) SetAccessibleID(name string) {
	atk.WrapObject(gextras.InternObject(a)).SetAccessibleID(name)
}

func (a imageCellAccessible) SetDescription(description string) {
	atk.WrapObject(gextras.InternObject(a)).SetDescription(description)
}

func (a imageCellAccessible) SetName(name string) {
	atk.WrapObject(gextras.InternObject(a)).SetName(name)
}

func (a imageCellAccessible) SetParent(parent atk.Object) {
	atk.WrapObject(gextras.InternObject(a)).SetParent(parent)
}

func (a imageCellAccessible) SetRole(role atk.Role) {
	atk.WrapObject(gextras.InternObject(a)).SetRole(role)
}

func (a imageCellAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a imageCellAccessible) GetDescription(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetDescription(i)
}

func (a imageCellAccessible) GetKeybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetKeybinding(i)
}

func (a imageCellAccessible) GetLocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetLocalizedName(i)
}

func (a imageCellAccessible) GetNActions() int {
	return atk.WrapAction(gextras.InternObject(a)).GetNActions()
}

func (a imageCellAccessible) GetName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetName(i)
}

func (a imageCellAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}

func (a imageCellAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a imageCellAccessible) GetDescription(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetDescription(i)
}

func (a imageCellAccessible) GetKeybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetKeybinding(i)
}

func (a imageCellAccessible) GetLocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetLocalizedName(i)
}

func (a imageCellAccessible) GetNActions() int {
	return atk.WrapAction(gextras.InternObject(a)).GetNActions()
}

func (a imageCellAccessible) GetName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetName(i)
}

func (a imageCellAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}

func (a imageCellAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a imageCellAccessible) GetDescription(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetDescription(i)
}

func (a imageCellAccessible) GetKeybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetKeybinding(i)
}

func (a imageCellAccessible) GetLocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetLocalizedName(i)
}

func (a imageCellAccessible) GetNActions() int {
	return atk.WrapAction(gextras.InternObject(a)).GetNActions()
}

func (a imageCellAccessible) GetName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).GetName(i)
}

func (a imageCellAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}

func (i imageCellAccessible) GetImageDescription() string {
	return atk.WrapImage(gextras.InternObject(i)).GetImageDescription()
}

func (i imageCellAccessible) GetImageLocale() string {
	return atk.WrapImage(gextras.InternObject(i)).GetImageLocale()
}

func (i imageCellAccessible) GetImagePosition(coordType atk.CoordType) (x int, y int) {
	return atk.WrapImage(gextras.InternObject(i)).GetImagePosition(coordType)
}

func (i imageCellAccessible) GetImageSize() (width int, height int) {
	return atk.WrapImage(gextras.InternObject(i)).GetImageSize()
}

func (i imageCellAccessible) SetImageDescription(description string) bool {
	return atk.WrapImage(gextras.InternObject(i)).SetImageDescription(description)
}
