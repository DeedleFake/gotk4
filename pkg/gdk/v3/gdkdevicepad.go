// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_pad_feature_get_type()), F: marshalDevicePadFeature},
		{T: externglib.Type(C.gdk_device_pad_get_type()), F: marshalDevicePad},
	})
}

// DevicePadFeature: pad feature.
type DevicePadFeature int

const (
	// button: button
	DevicePadFeatureButton DevicePadFeature = iota
	// ring: ring-shaped interactive area
	DevicePadFeatureRing
	// strip: straight interactive area
	DevicePadFeatureStrip
)

func marshalDevicePadFeature(p uintptr) (interface{}, error) {
	return DevicePadFeature(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DevicePad is an interface implemented by devices of type
// GDK_SOURCE_TABLET_PAD, it allows querying the features provided by the pad
// device.
//
// Tablet pads may contain one or more groups, each containing a subset of the
// buttons/rings/strips available. gdk_device_pad_get_n_groups() can be used to
// obtain the number of groups, gdk_device_pad_get_n_features() and
// gdk_device_pad_get_feature_group() can be combined to find out the number of
// buttons/rings/strips the device has, and how are they grouped.
//
// Each of those groups have different modes, which may be used to map each
// individual pad feature to multiple actions. Only one mode is effective
// (current) for each given group, different groups may have different current
// modes. The number of available modes in a group can be found out through
// gdk_device_pad_get_group_n_modes(), and the current mode for a given group
// will be notified through the EventPadGroupMode event.
type DevicePad interface {
	Device

	// FeatureGroup returns the group the given @feature and @idx belong to, or
	// -1 if feature/index do not exist in @pad.
	FeatureGroup(feature DevicePadFeature, featureIdx int) int
	// GroupNModes returns the number of modes that @group may have.
	GroupNModes(groupIdx int) int
	// NFeatures returns the number of features a tablet pad has.
	NFeatures(feature DevicePadFeature) int
	// NGroups returns the number of groups this pad device has. Pads have at
	// least one group. A pad group is a subcollection of buttons/strip/rings
	// that is affected collectively by a same current mode.
	NGroups() int
}

// devicePad implements the DevicePad interface.
type devicePad struct {
	Device
}

var _ DevicePad = (*devicePad)(nil)

// WrapDevicePad wraps a GObject to a type that implements
// interface DevicePad. It is primarily used internally.
func WrapDevicePad(obj *externglib.Object) DevicePad {
	return devicePad{
		Device: WrapDevice(obj),
	}
}

func marshalDevicePad(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDevicePad(obj), nil
}

func (p devicePad) FeatureGroup(feature DevicePadFeature, featureIdx int) int {
	var _arg0 *C.GdkDevicePad       // out
	var _arg1 C.GdkDevicePadFeature // out
	var _arg2 C.gint                // out
	var _cret C.gint                // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	_arg1 = C.GdkDevicePadFeature(feature)
	_arg2 = C.gint(featureIdx)

	_cret = C.gdk_device_pad_get_feature_group(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p devicePad) GroupNModes(groupIdx int) int {
	var _arg0 *C.GdkDevicePad // out
	var _arg1 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	_arg1 = C.gint(groupIdx)

	_cret = C.gdk_device_pad_get_group_n_modes(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p devicePad) NFeatures(feature DevicePadFeature) int {
	var _arg0 *C.GdkDevicePad       // out
	var _arg1 C.GdkDevicePadFeature // out
	var _cret C.gint                // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	_arg1 = C.GdkDevicePadFeature(feature)

	_cret = C.gdk_device_pad_get_n_features(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p devicePad) NGroups() int {
	var _arg0 *C.GdkDevicePad // out
	var _cret C.gint          // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_device_pad_get_n_groups(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
