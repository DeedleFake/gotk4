// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include "gtk.h"
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_editable_properties_get_type()), F: marshalEditableProperties},
		{T: externglib.Type(C.gtk_debug_flags_get_type()), F: marshalDebugFlags},
	})
}

func BuilderErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_builder_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

func ConstraintVflParserErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_constraint_vfl_parser_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type EditableProperties int

const (
	EditablePropText EditableProperties = iota
	EditablePropCursorPosition
	EditablePropSelectionBound
	EditablePropEditable
	EditablePropWidthChars
	EditablePropMaxWidthChars
	EditablePropXAlign
	EditablePropEnableUndo
	EditableNumProperties
)

func marshalEditableProperties(p uintptr) (interface{}, error) {
	return EditableProperties(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for EditableProperties.
func (e EditableProperties) String() string {
	switch e {
	case EditablePropText:
		return "PropText"
	case EditablePropCursorPosition:
		return "PropCursorPosition"
	case EditablePropSelectionBound:
		return "PropSelectionBound"
	case EditablePropEditable:
		return "PropEditable"
	case EditablePropWidthChars:
		return "PropWidthChars"
	case EditablePropMaxWidthChars:
		return "PropMaxWidthChars"
	case EditablePropXAlign:
		return "PropXAlign"
	case EditablePropEnableUndo:
		return "PropEnableUndo"
	case EditableNumProperties:
		return "NumProperties"
	default:
		return fmt.Sprintf("EditableProperties(%d)", e)
	}
}

func IconThemeErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_icon_theme_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

func RecentManagerErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_recent_manager_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

type DebugFlags int

const (
	DebugText           DebugFlags = 0b1
	DebugTree           DebugFlags = 0b10
	DebugKeybindings    DebugFlags = 0b100
	DebugModules        DebugFlags = 0b1000
	DebugGeometry       DebugFlags = 0b10000
	DebugIcontheme      DebugFlags = 0b100000
	DebugPrinting       DebugFlags = 0b1000000
	DebugBuilder        DebugFlags = 0b10000000
	DebugSizeRequest    DebugFlags = 0b100000000
	DebugNoCSSCache     DebugFlags = 0b1000000000
	DebugInteractive    DebugFlags = 0b10000000000
	DebugTouchscreen    DebugFlags = 0b100000000000
	DebugActions        DebugFlags = 0b1000000000000
	DebugLayout         DebugFlags = 0b10000000000000
	DebugSnapshot       DebugFlags = 0b100000000000000
	DebugConstraints    DebugFlags = 0b1000000000000000
	DebugBuilderObjects DebugFlags = 0b10000000000000000
	DebugA11Y           DebugFlags = 0b100000000000000000
	DebugIconfallback   DebugFlags = 0b1000000000000000000
)

func marshalDebugFlags(p uintptr) (interface{}, error) {
	return DebugFlags(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for DebugFlags.
func (d DebugFlags) String() string {
	if d == 0 {
		return "DebugFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(276)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case DebugText:
			builder.WriteString("Text|")
		case DebugTree:
			builder.WriteString("Tree|")
		case DebugKeybindings:
			builder.WriteString("Keybindings|")
		case DebugModules:
			builder.WriteString("Modules|")
		case DebugGeometry:
			builder.WriteString("Geometry|")
		case DebugIcontheme:
			builder.WriteString("Icontheme|")
		case DebugPrinting:
			builder.WriteString("Printing|")
		case DebugBuilder:
			builder.WriteString("Builder|")
		case DebugSizeRequest:
			builder.WriteString("SizeRequest|")
		case DebugNoCSSCache:
			builder.WriteString("NoCSSCache|")
		case DebugInteractive:
			builder.WriteString("Interactive|")
		case DebugTouchscreen:
			builder.WriteString("Touchscreen|")
		case DebugActions:
			builder.WriteString("Actions|")
		case DebugLayout:
			builder.WriteString("Layout|")
		case DebugSnapshot:
			builder.WriteString("Snapshot|")
		case DebugConstraints:
			builder.WriteString("Constraints|")
		case DebugBuilderObjects:
			builder.WriteString("BuilderObjects|")
		case DebugA11Y:
			builder.WriteString("A11Y|")
		case DebugIconfallback:
			builder.WriteString("Iconfallback|")
		default:
			builder.WriteString(fmt.Sprintf("DebugFlags(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DebugFlags) Has(other DebugFlags) bool {
	return (d & other) == other
}

func CSSParserErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_css_parser_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}
