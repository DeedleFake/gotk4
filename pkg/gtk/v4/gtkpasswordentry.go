// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_password_entry_get_type()), F: marshalPasswordEntry},
	})
}

// PasswordEntry: `GtkPasswordEntry` is an entry that has been tailored for
// entering secrets.
//
// !An example GtkPasswordEntry (password-entry.png)
//
// It does not show its contents in clear text, does not allow to copy it to the
// clipboard, and it shows a warning when Caps Lock is engaged. If the
// underlying platform allows it, `GtkPasswordEntry` will also place the text in
// a non-pageable memory area, to avoid it being written out to disk by the
// operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// `GtkPasswordEntry` provides only minimal API and should be used with the
// [iface@Gtk.Editable] API.
//
//
// CSS Nodes
//
// “` entry.password ╰── text ├── image.caps-lock-indicator ┊ “`
//
// `GtkPasswordEntry` has a single CSS node with name entry that carries a
// .passwordstyle class. The text Css node below it has a child with name image
// and style class .caps-lock-indicator for the Caps Lock icon, and possibly
// other children.
//
//
// Accessibility
//
// `GtkPasswordEntry` uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry interface {
	Editable

	ExtraMenu() gio.MenuModel

	ShowPeekIcon() bool

	SetExtraMenuPasswordEntry(model gio.MenuModel)

	SetShowPeekIconPasswordEntry(showPeekIcon bool)
}

// passwordEntry implements the PasswordEntry class.
type passwordEntry struct {
	Widget
}

// WrapPasswordEntry wraps a GObject to the right type. It is
// primarily used internally.
func WrapPasswordEntry(obj *externglib.Object) PasswordEntry {
	return passwordEntry{
		Widget: WrapWidget(obj),
	}
}

func marshalPasswordEntry(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPasswordEntry(obj), nil
}

func NewPasswordEntry() PasswordEntry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_password_entry_new()

	var _passwordEntry PasswordEntry // out

	_passwordEntry = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(PasswordEntry)

	return _passwordEntry
}

func (e passwordEntry) ExtraMenu() gio.MenuModel {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret *C.GMenuModel       // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_password_entry_get_extra_menu(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.MenuModel)

	return _menuModel
}

func (e passwordEntry) ShowPeekIcon() bool {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_password_entry_get_show_peek_icon(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e passwordEntry) SetExtraMenuPasswordEntry(model gio.MenuModel) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 *C.GMenuModel       // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_password_entry_set_extra_menu(_arg0, _arg1)
}

func (e passwordEntry) SetShowPeekIconPasswordEntry(showPeekIcon bool) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(e.Native()))
	if showPeekIcon {
		_arg1 = C.TRUE
	}

	C.gtk_password_entry_set_show_peek_icon(_arg0, _arg1)
}

func (e passwordEntry) DeleteSelection() {
	WrapEditable(gextras.InternObject(e)).DeleteSelection()
}

func (e passwordEntry) DeleteText(startPos int, endPos int) {
	WrapEditable(gextras.InternObject(e)).DeleteText(startPos, endPos)
}

func (e passwordEntry) FinishDelegate() {
	WrapEditable(gextras.InternObject(e)).FinishDelegate()
}

func (e passwordEntry) Alignment() float32 {
	return WrapEditable(gextras.InternObject(e)).Alignment()
}

func (e passwordEntry) Chars(startPos int, endPos int) string {
	return WrapEditable(gextras.InternObject(e)).Chars(startPos, endPos)
}

func (e passwordEntry) Delegate() Editable {
	return WrapEditable(gextras.InternObject(e)).Delegate()
}

func (e passwordEntry) Editable() bool {
	return WrapEditable(gextras.InternObject(e)).Editable()
}

func (e passwordEntry) EnableUndo() bool {
	return WrapEditable(gextras.InternObject(e)).EnableUndo()
}

func (e passwordEntry) MaxWidthChars() int {
	return WrapEditable(gextras.InternObject(e)).MaxWidthChars()
}

func (e passwordEntry) Position() int {
	return WrapEditable(gextras.InternObject(e)).Position()
}

func (e passwordEntry) SelectionBounds() (startPos int, endPos int, ok bool) {
	return WrapEditable(gextras.InternObject(e)).SelectionBounds()
}

func (e passwordEntry) Text() string {
	return WrapEditable(gextras.InternObject(e)).Text()
}

func (e passwordEntry) WidthChars() int {
	return WrapEditable(gextras.InternObject(e)).WidthChars()
}

func (e passwordEntry) InitDelegate() {
	WrapEditable(gextras.InternObject(e)).InitDelegate()
}

func (e passwordEntry) SelectRegion(startPos int, endPos int) {
	WrapEditable(gextras.InternObject(e)).SelectRegion(startPos, endPos)
}

func (e passwordEntry) SetAlignment(xalign float32) {
	WrapEditable(gextras.InternObject(e)).SetAlignment(xalign)
}

func (e passwordEntry) SetEditable(isEditable bool) {
	WrapEditable(gextras.InternObject(e)).SetEditable(isEditable)
}

func (e passwordEntry) SetEnableUndo(enableUndo bool) {
	WrapEditable(gextras.InternObject(e)).SetEnableUndo(enableUndo)
}

func (e passwordEntry) SetMaxWidthChars(nChars int) {
	WrapEditable(gextras.InternObject(e)).SetMaxWidthChars(nChars)
}

func (e passwordEntry) SetPosition(position int) {
	WrapEditable(gextras.InternObject(e)).SetPosition(position)
}

func (e passwordEntry) SetText(text string) {
	WrapEditable(gextras.InternObject(e)).SetText(text)
}

func (e passwordEntry) SetWidthChars(nChars int) {
	WrapEditable(gextras.InternObject(e)).SetWidthChars(nChars)
}

func (s passwordEntry) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s passwordEntry) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s passwordEntry) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s passwordEntry) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s passwordEntry) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s passwordEntry) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s passwordEntry) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b passwordEntry) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
