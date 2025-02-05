// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeConstraintTarget = coreglib.Type(C.gtk_constraint_target_get_type())
	GTypeConstraint       = coreglib.Type(C.gtk_constraint_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeConstraintTarget, F: marshalConstraintTarget},
		coreglib.TypeMarshaler{T: GTypeConstraint, F: marshalConstraint},
	})
}

// ConstraintTarget: GtkConstraintTarget interface is implemented by objects
// that can be used as source or target in GtkConstraints.
//
// Besides GtkWidget, it is also implemented by GtkConstraintGuide.
//
// ConstraintTarget wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ConstraintTarget struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ConstraintTarget)(nil)
)

// ConstraintTargetter describes ConstraintTarget's interface methods.
type ConstraintTargetter interface {
	coreglib.Objector

	baseConstraintTarget() *ConstraintTarget
}

var _ ConstraintTargetter = (*ConstraintTarget)(nil)

func wrapConstraintTarget(obj *coreglib.Object) *ConstraintTarget {
	return &ConstraintTarget{
		Object: obj,
	}
}

func marshalConstraintTarget(p uintptr) (interface{}, error) {
	return wrapConstraintTarget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ConstraintTarget) baseConstraintTarget() *ConstraintTarget {
	return v
}

// BaseConstraintTarget returns the underlying base object.
func BaseConstraintTarget(obj ConstraintTargetter) *ConstraintTarget {
	return obj.baseConstraintTarget()
}

// ConstraintOverrides contains methods that are overridable.
type ConstraintOverrides struct {
}

func defaultConstraintOverrides(v *Constraint) ConstraintOverrides {
	return ConstraintOverrides{}
}

// Constraint: GtkConstraint describes a constraint between attributes of two
// widgets, expressed as a linear equation.
//
// The typical equation for a constraint is:
//
//      target.target_attr = source.source_attr × multiplier + constant
//
//
// Each GtkConstraint is part of a system that will be solved by a
// gtk.ConstraintLayout in order to allocate and position each child widget or
// guide.
//
// The source and target, as well as their attributes, of a GtkConstraint
// instance are immutable after creation.
type Constraint struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Constraint)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Constraint, *ConstraintClass, ConstraintOverrides](
		GTypeConstraint,
		initConstraintClass,
		wrapConstraint,
		defaultConstraintOverrides,
	)
}

func initConstraintClass(gclass unsafe.Pointer, overrides ConstraintOverrides, classInitFunc func(*ConstraintClass)) {
	if classInitFunc != nil {
		class := (*ConstraintClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapConstraint(obj *coreglib.Object) *Constraint {
	return &Constraint{
		Object: obj,
	}
}

func marshalConstraint(p uintptr) (interface{}, error) {
	return wrapConstraint(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewConstraint creates a new constraint representing a relation between a
// layout attribute on a source and a layout attribute on a target.
//
// The function takes the following parameters:
//
//    - target (optional) of the constraint.
//    - targetAttribute: attribute of target to be set.
//    - relation equivalence between target_attribute and source_attribute.
//    - source (optional) of the constraint.
//    - sourceAttribute: attribute of source to be read.
//    - multiplier: multiplication factor to be applied to source_attribute.
//    - constant factor to be added to source_attribute.
//    - strength of the constraint.
//
// The function returns the following values:
//
//    - constraint: newly created constraint.
//
func NewConstraint(target ConstraintTargetter, targetAttribute ConstraintAttribute, relation ConstraintRelation, source ConstraintTargetter, sourceAttribute ConstraintAttribute, multiplier, constant float64, strength int) *Constraint {
	var _arg1 C.gpointer               // out
	var _arg2 C.GtkConstraintAttribute // out
	var _arg3 C.GtkConstraintRelation  // out
	var _arg4 C.gpointer               // out
	var _arg5 C.GtkConstraintAttribute // out
	var _arg6 C.double                 // out
	var _arg7 C.double                 // out
	var _arg8 C.int                    // out
	var _cret *C.GtkConstraint         // in

	if target != nil {
		_arg1 = *(*C.gpointer)(unsafe.Pointer(coreglib.InternObject(target).Native()))
	}
	_arg2 = C.GtkConstraintAttribute(targetAttribute)
	_arg3 = C.GtkConstraintRelation(relation)
	if source != nil {
		_arg4 = *(*C.gpointer)(unsafe.Pointer(coreglib.InternObject(source).Native()))
	}
	_arg5 = C.GtkConstraintAttribute(sourceAttribute)
	_arg6 = C.double(multiplier)
	_arg7 = C.double(constant)
	_arg8 = C.int(strength)

	_cret = C.gtk_constraint_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
	runtime.KeepAlive(target)
	runtime.KeepAlive(targetAttribute)
	runtime.KeepAlive(relation)
	runtime.KeepAlive(source)
	runtime.KeepAlive(sourceAttribute)
	runtime.KeepAlive(multiplier)
	runtime.KeepAlive(constant)
	runtime.KeepAlive(strength)

	var _constraint *Constraint // out

	_constraint = wrapConstraint(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constraint
}

// NewConstraintConstant creates a new constraint representing a relation
// between a layout attribute on a target and a constant value.
//
// The function takes the following parameters:
//
//    - target (optional): the target of the constraint.
//    - targetAttribute: attribute of target to be set.
//    - relation equivalence between target_attribute and constant.
//    - constant factor to be set on target_attribute.
//    - strength of the constraint.
//
// The function returns the following values:
//
//    - constraint: newly created constraint.
//
func NewConstraintConstant(target ConstraintTargetter, targetAttribute ConstraintAttribute, relation ConstraintRelation, constant float64, strength int) *Constraint {
	var _arg1 C.gpointer               // out
	var _arg2 C.GtkConstraintAttribute // out
	var _arg3 C.GtkConstraintRelation  // out
	var _arg4 C.double                 // out
	var _arg5 C.int                    // out
	var _cret *C.GtkConstraint         // in

	if target != nil {
		_arg1 = *(*C.gpointer)(unsafe.Pointer(coreglib.InternObject(target).Native()))
	}
	_arg2 = C.GtkConstraintAttribute(targetAttribute)
	_arg3 = C.GtkConstraintRelation(relation)
	_arg4 = C.double(constant)
	_arg5 = C.int(strength)

	_cret = C.gtk_constraint_new_constant(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(target)
	runtime.KeepAlive(targetAttribute)
	runtime.KeepAlive(relation)
	runtime.KeepAlive(constant)
	runtime.KeepAlive(strength)

	var _constraint *Constraint // out

	_constraint = wrapConstraint(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _constraint
}

// Constant retrieves the constant factor added to the source attributes' value.
//
// The function returns the following values:
//
//    - gdouble: constant factor.
//
func (constraint *Constraint) Constant() float64 {
	var _arg0 *C.GtkConstraint // out
	var _cret C.double         // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_get_constant(_arg0)
	runtime.KeepAlive(constraint)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Multiplier retrieves the multiplication factor applied to the source
// attribute's value.
//
// The function returns the following values:
//
//    - gdouble: multiplication factor.
//
func (constraint *Constraint) Multiplier() float64 {
	var _arg0 *C.GtkConstraint // out
	var _cret C.double         // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_get_multiplier(_arg0)
	runtime.KeepAlive(constraint)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Relation: order relation between the terms of the constraint.
//
// The function returns the following values:
//
//    - constraintRelation: relation type.
//
func (constraint *Constraint) Relation() ConstraintRelation {
	var _arg0 *C.GtkConstraint        // out
	var _cret C.GtkConstraintRelation // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_get_relation(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintRelation ConstraintRelation // out

	_constraintRelation = ConstraintRelation(_cret)

	return _constraintRelation
}

// Source retrieves the gtk.ConstraintTarget used as the source for the
// constraint.
//
// If the source is set to NULL at creation, the constraint will use the widget
// using the gtk.ConstraintLayout as the source.
//
// The function returns the following values:
//
//    - constraintTarget (optional): source of the constraint.
//
func (constraint *Constraint) Source() *ConstraintTarget {
	var _arg0 *C.GtkConstraint       // out
	var _cret *C.GtkConstraintTarget // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_get_source(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintTarget *ConstraintTarget // out

	if _cret != nil {
		_constraintTarget = wrapConstraintTarget(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _constraintTarget
}

// SourceAttribute retrieves the attribute of the source to be read by the
// constraint.
//
// The function returns the following values:
//
//    - constraintAttribute source's attribute.
//
func (constraint *Constraint) SourceAttribute() ConstraintAttribute {
	var _arg0 *C.GtkConstraint         // out
	var _cret C.GtkConstraintAttribute // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_get_source_attribute(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintAttribute ConstraintAttribute // out

	_constraintAttribute = ConstraintAttribute(_cret)

	return _constraintAttribute
}

// Strength retrieves the strength of the constraint.
//
// The function returns the following values:
//
//    - gint: strength value.
//
func (constraint *Constraint) Strength() int {
	var _arg0 *C.GtkConstraint // out
	var _cret C.int            // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_get_strength(_arg0)
	runtime.KeepAlive(constraint)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Target retrieves the gtk.ConstraintTarget used as the target for the
// constraint.
//
// If the targe is set to NULL at creation, the constraint will use the widget
// using the gtk.ConstraintLayout as the target.
//
// The function returns the following values:
//
//    - constraintTarget (optional): ConstraintTarget.
//
func (constraint *Constraint) Target() *ConstraintTarget {
	var _arg0 *C.GtkConstraint       // out
	var _cret *C.GtkConstraintTarget // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_get_target(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintTarget *ConstraintTarget // out

	if _cret != nil {
		_constraintTarget = wrapConstraintTarget(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _constraintTarget
}

// TargetAttribute retrieves the attribute of the target to be set by the
// constraint.
//
// The function returns the following values:
//
//    - constraintAttribute target's attribute.
//
func (constraint *Constraint) TargetAttribute() ConstraintAttribute {
	var _arg0 *C.GtkConstraint         // out
	var _cret C.GtkConstraintAttribute // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_get_target_attribute(_arg0)
	runtime.KeepAlive(constraint)

	var _constraintAttribute ConstraintAttribute // out

	_constraintAttribute = ConstraintAttribute(_cret)

	return _constraintAttribute
}

// IsAttached checks whether the constraint is attached to a
// gtk.ConstraintLayout, and it is contributing to the layout.
//
// The function returns the following values:
//
//    - ok: TRUE if the constraint is attached.
//
func (constraint *Constraint) IsAttached() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_is_attached(_arg0)
	runtime.KeepAlive(constraint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsConstant checks whether the constraint describes a relation between an
// attribute on the gtk.Constraint:target and a constant value.
//
// The function returns the following values:
//
//    - ok: TRUE if the constraint is a constant relation.
//
func (constraint *Constraint) IsConstant() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_is_constant(_arg0)
	runtime.KeepAlive(constraint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRequired checks whether the constraint is a required relation for solving
// the constraint layout.
//
// The function returns the following values:
//
//    - ok: TRUE if the constraint is required.
//
func (constraint *Constraint) IsRequired() bool {
	var _arg0 *C.GtkConstraint // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkConstraint)(unsafe.Pointer(coreglib.InternObject(constraint).Native()))

	_cret = C.gtk_constraint_is_required(_arg0)
	runtime.KeepAlive(constraint)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ConstraintClass: instance of this type is always passed by reference.
type ConstraintClass struct {
	*constraintClass
}

// constraintClass is the struct that's finalized.
type constraintClass struct {
	native *C.GtkConstraintClass
}
