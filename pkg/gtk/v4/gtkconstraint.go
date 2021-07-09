// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_constraint_target_get_type()), F: marshalConstraintTarget},
		{T: externglib.Type(C.gtk_constraint_get_type()), F: marshalConstraint},
	})
}

// ConstraintTarget: the `GtkConstraintTarget` interface is implemented by
// objects that can be used as source or target in `GtkConstraint`s.
//
// Besides `GtkWidget`, it is also implemented by `GtkConstraintGuide`.
type ConstraintTarget interface {
	gextras.Objector

	privateConstraintTargetInterface()
}

// ConstraintTargetInterface implements the ConstraintTarget interface.
type ConstraintTargetInterface struct {
	*externglib.Object
}

var _ ConstraintTarget = (*ConstraintTargetInterface)(nil)

func wrapConstraintTarget(obj *externglib.Object) ConstraintTarget {
	return &ConstraintTargetInterface{
		Object: obj,
	}
}

func marshalConstraintTarget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConstraintTarget(obj), nil
}

func (*ConstraintTargetInterface) privateConstraintTargetInterface() {}

// Constraint: `GtkConstraint` describes a constraint between attributes of two
// widgets, expressed as a linear equation.
//
// The typical equation for a constraint is:
//
// “` target.target_attr = source.source_attr × multiplier + constant “`
//
// Each `GtkConstraint` is part of a system that will be solved by a
// [class@Gtk.ConstraintLayout] in order to allocate and position each child
// widget or guide.
//
// The source and target, as well as their attributes, of a `GtkConstraint`
// instance are immutable after creation.
type Constraint interface {
	gextras.Objector

	// Constant retrieves the constant factor added to the source attributes'
	// value.
	Constant() float64
	// Multiplier retrieves the multiplication factor applied to the source
	// attribute's value.
	Multiplier() float64
	// Relation: the order relation between the terms of the constraint.
	Relation() ConstraintRelation
	// Source retrieves the [iface@Gtk.ConstraintTarget] used as the source for
	// the constraint.
	//
	// If the source is set to `NULL` at creation, the constraint will use the
	// widget using the [class@Gtk.ConstraintLayout] as the source.
	Source() *ConstraintTargetInterface
	// SourceAttribute retrieves the attribute of the source to be read by the
	// constraint.
	SourceAttribute() ConstraintAttribute
	// Strength retrieves the strength of the constraint.
	Strength() int
	// Target retrieves the [iface@Gtk.ConstraintTarget] used as the target for
	// the constraint.
	//
	// If the targe is set to `NULL` at creation, the constraint will use the
	// widget using the [class@Gtk.ConstraintLayout] as the target.
	Target() *ConstraintTargetInterface
	// TargetAttribute retrieves the attribute of the target to be set by the
	// constraint.
	TargetAttribute() ConstraintAttribute
	// IsAttached checks whether the constraint is attached to a
	// [class@Gtk.ConstraintLayout], and it is contributing to the layout.
	IsAttached() bool
	// IsConstant checks whether the constraint describes a relation between an
	// attribute on the [property@Gtk.Constraint:target] and a constant value.
	IsConstant() bool
	// IsRequired checks whether the constraint is a required relation for
	// solving the constraint layout.
	IsRequired() bool
}

// ConstraintClass implements the Constraint interface.
type ConstraintClass struct {
	*externglib.Object
}

var _ Constraint = (*ConstraintClass)(nil)

func wrapConstraint(obj *externglib.Object) Constraint {
	return &ConstraintClass{
		Object: obj,
	}
}

func marshalConstraint(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConstraint(obj), nil
}

// Constant retrieves the constant factor added to the source attributes' value.
func (c *ConstraintClass) Constant() float64 {
	var _arg0 *C.GtkConstraint // out
	var _cret C.double         // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_get_constant(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Multiplier retrieves the multiplication factor applied to the source
// attribute's value.
func (c *ConstraintClass) Multiplier() float64 {
	var _arg0 *C.GtkConstraint // out
	var _cret C.double         // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_get_multiplier(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Relation: the order relation between the terms of the constraint.
func (c *ConstraintClass) Relation() ConstraintRelation {
	var _arg0 *C.GtkConstraint        // out
	var _cret C.GtkConstraintRelation // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_get_relation(_arg0)

	var _constraintRelation ConstraintRelation // out

	_constraintRelation = (ConstraintRelation)(C.GtkConstraintRelation)

	return _constraintRelation
}

// Source retrieves the [iface@Gtk.ConstraintTarget] used as the source for the
// constraint.
//
// If the source is set to `NULL` at creation, the constraint will use the
// widget using the [class@Gtk.ConstraintLayout] as the source.
func (c *ConstraintClass) Source() *ConstraintTargetInterface {
	var _arg0 *C.GtkConstraint       // out
	var _cret *C.GtkConstraintTarget // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_get_source(_arg0)

	var _constraintTarget *ConstraintTargetInterface // out

	_constraintTarget = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ConstraintTargetInterface)

	return _constraintTarget
}

// SourceAttribute retrieves the attribute of the source to be read by the
// constraint.
func (c *ConstraintClass) SourceAttribute() ConstraintAttribute {
	var _arg0 *C.GtkConstraint         // out
	var _cret C.GtkConstraintAttribute // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_get_source_attribute(_arg0)

	var _constraintAttribute ConstraintAttribute // out

	_constraintAttribute = (ConstraintAttribute)(C.GtkConstraintAttribute)

	return _constraintAttribute
}

// Strength retrieves the strength of the constraint.
func (c *ConstraintClass) Strength() int {
	var _arg0 *C.GtkConstraint // out
	var _cret C.int            // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_get_strength(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Target retrieves the [iface@Gtk.ConstraintTarget] used as the target for the
// constraint.
//
// If the targe is set to `NULL` at creation, the constraint will use the widget
// using the [class@Gtk.ConstraintLayout] as the target.
func (c *ConstraintClass) Target() *ConstraintTargetInterface {
	var _arg0 *C.GtkConstraint       // out
	var _cret *C.GtkConstraintTarget // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_get_target(_arg0)

	var _constraintTarget *ConstraintTargetInterface // out

	_constraintTarget = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ConstraintTargetInterface)

	return _constraintTarget
}

// TargetAttribute retrieves the attribute of the target to be set by the
// constraint.
func (c *ConstraintClass) TargetAttribute() ConstraintAttribute {
	var _arg0 *C.GtkConstraint         // out
	var _cret C.GtkConstraintAttribute // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_get_target_attribute(_arg0)

	var _constraintAttribute ConstraintAttribute // out

	_constraintAttribute = (ConstraintAttribute)(C.GtkConstraintAttribute)

	return _constraintAttribute
}

// IsAttached checks whether the constraint is attached to a
// [class@Gtk.ConstraintLayout], and it is contributing to the layout.
func (c *ConstraintClass) IsAttached() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_is_attached(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsConstant checks whether the constraint describes a relation between an
// attribute on the [property@Gtk.Constraint:target] and a constant value.
func (c *ConstraintClass) IsConstant() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_is_constant(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRequired checks whether the constraint is a required relation for solving
// the constraint layout.
func (c *ConstraintClass) IsRequired() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer((&Constraint).Native()))

	_cret = C.gtk_constraint_is_required(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
