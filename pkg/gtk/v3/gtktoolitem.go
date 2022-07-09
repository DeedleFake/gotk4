// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern gboolean _gotk4_gtk3_ToolItemClass_create_menu_proxy(void*);
// extern gboolean _gotk4_gtk3_ToolItem_ConnectCreateMenuProxy(gpointer, guintptr);
// extern void _gotk4_gtk3_ToolItemClass_toolbar_reconfigured(void*);
// extern void _gotk4_gtk3_ToolItem_ConnectToolbarReconfigured(gpointer, guintptr);
import "C"

// GTypeToolItem returns the GType for the type ToolItem.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeToolItem() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ToolItem").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalToolItem)
	return gtype
}

// ToolItemOverrider contains methods that are overridable.
type ToolItemOverrider interface {
	// The function returns the following values:
	//
	CreateMenuProxy() bool
	// ToolbarReconfigured emits the signal ToolItem::toolbar_reconfigured on
	// tool_item. Toolbar and other ToolShell implementations use this function
	// to notify children, when some aspect of their configuration changes.
	ToolbarReconfigured()
}

// ToolItem are widgets that can appear on a toolbar. To create a toolbar item
// that contain something else than a button, use gtk_tool_item_new(). Use
// gtk_container_add() to add a child widget to the tool item.
//
// For toolbar items that contain buttons, see the ToolButton, ToggleToolButton
// and RadioToolButton classes.
//
// See the Toolbar class for a description of the toolbar widget, and ToolShell
// for a description of the tool shell interface.
type ToolItem struct {
	_ [0]func() // equal guard
	Bin

	*coreglib.Object
	Activatable
}

var (
	_ Binner            = (*ToolItem)(nil)
	_ coreglib.Objector = (*ToolItem)(nil)
)

func classInitToolItemmer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ToolItemClass")

	if _, ok := goval.(interface{ CreateMenuProxy() bool }); ok {
		o := pclass.StructFieldOffset("create_menu_proxy")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ToolItemClass_create_menu_proxy)
	}

	if _, ok := goval.(interface{ ToolbarReconfigured() }); ok {
		o := pclass.StructFieldOffset("toolbar_reconfigured")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ToolItemClass_toolbar_reconfigured)
	}
}

//export _gotk4_gtk3_ToolItemClass_create_menu_proxy
func _gotk4_gtk3_ToolItemClass_create_menu_proxy(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ CreateMenuProxy() bool })

	ok := iface.CreateMenuProxy()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_ToolItemClass_toolbar_reconfigured
func _gotk4_gtk3_ToolItemClass_toolbar_reconfigured(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ToolbarReconfigured() })

	iface.ToolbarReconfigured()
}

func wrapToolItem(obj *coreglib.Object) *ToolItem {
	return &ToolItem{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
		},
		Object: obj,
		Activatable: Activatable{
			Object: obj,
		},
	}
}

func marshalToolItem(p uintptr) (interface{}, error) {
	return wrapToolItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ToolItem_ConnectCreateMenuProxy
func _gotk4_gtk3_ToolItem_ConnectCreateMenuProxy(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectCreateMenuProxy: this signal is emitted when the toolbar needs
// information from tool_item about whether the item should appear in the
// toolbar overflow menu. In response the tool item should either
//
// - call gtk_tool_item_set_proxy_menu_item() with a NULL pointer and return
// TRUE to indicate that the item should not appear in the overflow menu
//
// - call gtk_tool_item_set_proxy_menu_item() with a new menu item and return
// TRUE, or
//
// - return FALSE to indicate that the signal was not handled by the item. This
// means that the item will not appear in the overflow menu unless a later
// handler installs a menu item.
//
// The toolbar may cache the result of this signal. When the tool item changes
// how it will respond to this signal it must call gtk_tool_item_rebuild_menu()
// to invalidate the cache and ensure that the toolbar rebuilds its overflow
// menu.
func (toolItem *ToolItem) ConnectCreateMenuProxy(f func() (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(toolItem, "create-menu-proxy", false, unsafe.Pointer(C._gotk4_gtk3_ToolItem_ConnectCreateMenuProxy), f)
}

//export _gotk4_gtk3_ToolItem_ConnectToolbarReconfigured
func _gotk4_gtk3_ToolItem_ConnectToolbarReconfigured(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectToolbarReconfigured: this signal is emitted when some property of the
// toolbar that the item is a child of changes. For custom subclasses of
// ToolItem, the default handler of this signal use the functions
//
// - gtk_tool_shell_get_orientation()
//
// - gtk_tool_shell_get_style()
//
// - gtk_tool_shell_get_icon_size()
//
// - gtk_tool_shell_get_relief_style() to find out what the toolbar should look
// like and change themselves accordingly.
func (toolItem *ToolItem) ConnectToolbarReconfigured(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(toolItem, "toolbar-reconfigured", false, unsafe.Pointer(C._gotk4_gtk3_ToolItem_ConnectToolbarReconfigured), f)
}

// NewToolItem creates a new ToolItem.
//
// The function returns the following values:
//
//    - toolItem: new ToolItem.
//
func NewToolItem() *ToolItem {
	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("new_ToolItem", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _toolItem *ToolItem // out

	_toolItem = wrapToolItem(coreglib.Take(unsafe.Pointer(_cret)))

	return _toolItem
}

// Expand returns whether tool_item is allocated extra space. See
// gtk_tool_item_set_expand().
//
// The function returns the following values:
//
//    - ok: TRUE if tool_item is allocated extra space.
//
func (toolItem *ToolItem) Expand() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_expand", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous returns whether tool_item is the same size as other homogeneous
// items. See gtk_tool_item_set_homogeneous().
//
// The function returns the following values:
//
//    - ok: TRUE if the item is the same size as other homogeneous items.
//
func (toolItem *ToolItem) Homogeneous() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_homogeneous", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsImportant returns whether tool_item is considered important. See
// gtk_tool_item_set_is_important().
//
// The function returns the following values:
//
//    - ok: TRUE if tool_item is considered important.
//
func (toolItem *ToolItem) IsImportant() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_is_important", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ProxyMenuItem: if menu_item_id matches the string passed to
// gtk_tool_item_set_proxy_menu_item() return the corresponding MenuItem.
//
// Custom subclasses of ToolItem should use this function to update their menu
// item when the ToolItem changes. That the menu_item_ids must match ensures
// that a ToolItem will not inadvertently change a menu item that they did not
// create.
//
// The function takes the following parameters:
//
//    - menuItemId: string used to identify the menu item.
//
// The function returns the following values:
//
//    - widget (optional) passed to gtk_tool_item_set_proxy_menu_item(), if the
//      menu_item_ids match.
//
func (toolItem *ToolItem) ProxyMenuItem(menuItemId string) Widgetter {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(menuItemId)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_proxy_menu_item", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(menuItemId)

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

// TextAlignment returns the text alignment used for tool_item. Custom
// subclasses of ToolItem should call this function to find out how text should
// be aligned.
//
// The function returns the following values:
//
//    - gfloat indicating the horizontal text alignment used for tool_item.
//
func (toolItem *ToolItem) TextAlignment() float32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_text_alignment", _args[:], nil)
	_cret := *(*C.gfloat)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _gfloat float32 // out

	_gfloat = float32(*(*C.gfloat)(unsafe.Pointer(&_cret)))

	return _gfloat
}

// TextSizeGroup returns the size group used for labels in tool_item. Custom
// subclasses of ToolItem should call this function and use the size group for
// labels.
//
// The function returns the following values:
//
//    - sizeGroup: SizeGroup.
//
func (toolItem *ToolItem) TextSizeGroup() *SizeGroup {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_text_size_group", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _sizeGroup *SizeGroup // out

	_sizeGroup = wrapSizeGroup(coreglib.Take(unsafe.Pointer(_cret)))

	return _sizeGroup
}

// UseDragWindow returns whether tool_item has a drag window. See
// gtk_tool_item_set_use_drag_window().
//
// The function returns the following values:
//
//    - ok: TRUE if tool_item uses a drag window.
//
func (toolItem *ToolItem) UseDragWindow() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_use_drag_window", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// VisibleHorizontal returns whether the tool_item is visible on toolbars that
// are docked horizontally.
//
// The function returns the following values:
//
//    - ok: TRUE if tool_item is visible on toolbars that are docked
//      horizontally.
//
func (toolItem *ToolItem) VisibleHorizontal() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_visible_horizontal", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// VisibleVertical returns whether tool_item is visible when the toolbar is
// docked vertically. See gtk_tool_item_set_visible_vertical().
//
// The function returns the following values:
//
//    - ok: whether tool_item is visible when the toolbar is docked vertically.
//
func (toolItem *ToolItem) VisibleVertical() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("get_visible_vertical", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// RebuildMenu: calling this function signals to the toolbar that the overflow
// menu item for tool_item has changed. If the overflow menu is visible when
// this function it called, the menu will be rebuilt.
//
// The function must be called when the tool item changes what it will do in
// response to the ToolItem::create-menu-proxy signal.
func (toolItem *ToolItem) RebuildMenu() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("rebuild_menu", _args[:], nil)

	runtime.KeepAlive(toolItem)
}

// RetrieveProxyMenuItem returns the MenuItem that was last set by
// gtk_tool_item_set_proxy_menu_item(), ie. the MenuItem that is going to appear
// in the overflow menu.
//
// The function returns the following values:
//
//    - widget that is going to appear in the overflow menu for tool_item.
//
func (toolItem *ToolItem) RetrieveProxyMenuItem() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_gret := _info.InvokeClassMethod("retrieve_proxy_menu_item", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(toolItem)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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

	return _widget
}

// SetExpand sets whether tool_item is allocated extra space when there is more
// room on the toolbar then needed for the items. The effect is that the item
// gets bigger when the toolbar gets bigger and smaller when the toolbar gets
// smaller.
//
// The function takes the following parameters:
//
//    - expand: whether tool_item is allocated extra space.
//
func (toolItem *ToolItem) SetExpand(expand bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	if expand {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_expand", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(expand)
}

// SetHomogeneous sets whether tool_item is to be allocated the same size as
// other homogeneous items. The effect is that all homogeneous items will have
// the same width as the widest of the items.
//
// The function takes the following parameters:
//
//    - homogeneous: whether tool_item is the same size as other homogeneous
//      items.
//
func (toolItem *ToolItem) SetHomogeneous(homogeneous bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	if homogeneous {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_homogeneous", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(homogeneous)
}

// SetIsImportant sets whether tool_item should be considered important. The
// ToolButton class uses this property to determine whether to show or hide its
// label when the toolbar style is GTK_TOOLBAR_BOTH_HORIZ. The result is that
// only tool buttons with the “is_important” property set have labels, an effect
// known as “priority text”.
//
// The function takes the following parameters:
//
//    - isImportant: whether the tool item should be considered important.
//
func (toolItem *ToolItem) SetIsImportant(isImportant bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	if isImportant {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_is_important", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(isImportant)
}

// SetProxyMenuItem sets the MenuItem used in the toolbar overflow menu. The
// menu_item_id is used to identify the caller of this function and should also
// be used with gtk_tool_item_get_proxy_menu_item().
//
// See also ToolItem::create-menu-proxy.
//
// The function takes the following parameters:
//
//    - menuItemId: string used to identify menu_item.
//    - menuItem (optional) to use in the overflow menu, or NULL.
//
func (toolItem *ToolItem) SetProxyMenuItem(menuItemId string, menuItem Widgetter) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(menuItemId)))
	defer C.free(unsafe.Pointer(_args[1]))
	if menuItem != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(menuItem).Native()))
	}

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_proxy_menu_item", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(menuItemId)
	runtime.KeepAlive(menuItem)
}

// SetTooltipMarkup sets the markup text to be displayed as tooltip on the item.
// See gtk_widget_set_tooltip_markup().
//
// The function takes the following parameters:
//
//    - markup text to be used as tooltip for tool_item.
//
func (toolItem *ToolItem) SetTooltipMarkup(markup string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_tooltip_markup", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(markup)
}

// SetTooltipText sets the text to be displayed as tooltip on the item. See
// gtk_widget_set_tooltip_text().
//
// The function takes the following parameters:
//
//    - text to be used as tooltip for tool_item.
//
func (toolItem *ToolItem) SetTooltipText(text string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_tooltip_text", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(text)
}

// SetUseDragWindow sets whether tool_item has a drag window. When TRUE the
// toolitem can be used as a drag source through gtk_drag_source_set(). When
// tool_item has a drag window it will intercept all events, even those that
// would otherwise be sent to a child of tool_item.
//
// The function takes the following parameters:
//
//    - useDragWindow: whether tool_item has a drag window.
//
func (toolItem *ToolItem) SetUseDragWindow(useDragWindow bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	if useDragWindow {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_use_drag_window", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(useDragWindow)
}

// SetVisibleHorizontal sets whether tool_item is visible when the toolbar is
// docked horizontally.
//
// The function takes the following parameters:
//
//    - visibleHorizontal: whether tool_item is visible when in horizontal mode.
//
func (toolItem *ToolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	if visibleHorizontal {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_visible_horizontal", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(visibleHorizontal)
}

// SetVisibleVertical sets whether tool_item is visible when the toolbar is
// docked vertically. Some tool items, such as text entries, are too wide to be
// useful on a vertically docked toolbar. If visible_vertical is FALSE tool_item
// will not appear on toolbars that are docked vertically.
//
// The function takes the following parameters:
//
//    - visibleVertical: whether tool_item is visible when the toolbar is in
//      vertical mode.
//
func (toolItem *ToolItem) SetVisibleVertical(visibleVertical bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))
	if visibleVertical {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("set_visible_vertical", _args[:], nil)

	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(visibleVertical)
}

// ToolbarReconfigured emits the signal ToolItem::toolbar_reconfigured on
// tool_item. Toolbar and other ToolShell implementations use this function to
// notify children, when some aspect of their configuration changes.
func (toolItem *ToolItem) ToolbarReconfigured() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(toolItem).Native()))

	_info := girepository.MustFind("Gtk", "ToolItem")
	_info.InvokeClassMethod("toolbar_reconfigured", _args[:], nil)

	runtime.KeepAlive(toolItem)
}
