// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
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
		{T: externglib.Type(C.gtk_cclosure_expression_get_type()), F: marshalCClosureExpression},
		{T: externglib.Type(C.gtk_closure_expression_get_type()), F: marshalClosureExpression},
		{T: externglib.Type(C.gtk_constant_expression_get_type()), F: marshalConstantExpression},
		{T: externglib.Type(C.gtk_expression_get_type()), F: marshalExpression},
		{T: externglib.Type(C.gtk_object_expression_get_type()), F: marshalObjectExpression},
		{T: externglib.Type(C.gtk_property_expression_get_type()), F: marshalPropertyExpression},
		{T: externglib.Type(C.gtk_expression_watch_get_type()), F: marshalExpressionWatch},
	})
}

// ExpressionNotify: callback called by gtk_expression_watch() when the
// expression value changes.
type ExpressionNotify func(userData interface{})

//export gotk4_ExpressionNotify
func gotk4_ExpressionNotify(arg0 C.gpointer) {
	v := box.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	var userData interface{} // out

	userData = box.Get(uintptr(arg0))

	fn := v.(ExpressionNotify)
	fn(userData)
}

// ValueDupExpression retrieves the `GtkExpression` stored inside the given
// `value`, and acquires a reference to it.
func ValueDupExpression(value externglib.Value) *ExpressionClass {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(&(&externglib.Value).GValue))

	_cret = C.gtk_value_dup_expression(_arg1)

	var _expression *ExpressionClass // out

	_expression = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*ExpressionClass)

	return _expression
}

// ValueGetExpression retrieves the `GtkExpression` stored inside the given
// `value`.
func ValueGetExpression(value externglib.Value) *ExpressionClass {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(&(&externglib.Value).GValue))

	_cret = C.gtk_value_get_expression(_arg1)

	var _expression *ExpressionClass // out

	_expression = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ExpressionClass)

	return _expression
}

// ValueSetExpression stores the given `GtkExpression` inside `value`.
//
// The `GValue` will acquire a reference to the `expression`.
func ValueSetExpression(value externglib.Value, expression Expression) {
	var _arg1 *C.GValue        // out
	var _arg2 *C.GtkExpression // out

	_arg1 = (*C.GValue)(unsafe.Pointer(&(&externglib.Value).GValue))
	_arg2 = (*C.GtkExpression)(unsafe.Pointer((&Expression).Native()))

	C.gtk_value_set_expression(_arg1, _arg2)
}

// ValueTakeExpression stores the given `GtkExpression` inside `value`.
//
// This function transfers the ownership of the `expression` to the `GValue`.
func ValueTakeExpression(value externglib.Value, expression Expression) {
	var _arg1 *C.GValue        // out
	var _arg2 *C.GtkExpression // out

	_arg1 = (*C.GValue)(unsafe.Pointer(&(&externglib.Value).GValue))
	_arg2 = (*C.GtkExpression)(unsafe.Pointer((&Expression).Native()))

	C.gtk_value_take_expression(_arg1, _arg2)
}

// CClosureExpression: variant of `GtkClosureExpression` using a C closure.
type CClosureExpression interface {
	gextras.Objector

	privateCClosureExpressionClass()
}

// CClosureExpressionClass implements the CClosureExpression interface.
type CClosureExpressionClass struct {
	ExpressionClass
}

var _ CClosureExpression = (*CClosureExpressionClass)(nil)

func wrapCClosureExpression(obj *externglib.Object) CClosureExpression {
	return &CClosureExpressionClass{
		ExpressionClass: ExpressionClass{
			Object: obj,
		},
	}
}

func marshalCClosureExpression(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCClosureExpression(obj), nil
}

func (*CClosureExpressionClass) privateCClosureExpressionClass() {}

// ClosureExpression: expression using a custom `GClosure` to compute the value
// from its parameters.
type ClosureExpression interface {
	gextras.Objector

	privateClosureExpressionClass()
}

// ClosureExpressionClass implements the ClosureExpression interface.
type ClosureExpressionClass struct {
	ExpressionClass
}

var _ ClosureExpression = (*ClosureExpressionClass)(nil)

func wrapClosureExpression(obj *externglib.Object) ClosureExpression {
	return &ClosureExpressionClass{
		ExpressionClass: ExpressionClass{
			Object: obj,
		},
	}
}

func marshalClosureExpression(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapClosureExpression(obj), nil
}

func (*ClosureExpressionClass) privateClosureExpressionClass() {}

// ConstantExpression: constant value in a `GtkExpression`.
type ConstantExpression interface {
	gextras.Objector

	privateConstantExpressionClass()
}

// ConstantExpressionClass implements the ConstantExpression interface.
type ConstantExpressionClass struct {
	ExpressionClass
}

var _ ConstantExpression = (*ConstantExpressionClass)(nil)

func wrapConstantExpression(obj *externglib.Object) ConstantExpression {
	return &ConstantExpressionClass{
		ExpressionClass: ExpressionClass{
			Object: obj,
		},
	}
}

func marshalConstantExpression(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConstantExpression(obj), nil
}

// NewConstantExpressionForValue creates an expression that always evaluates to
// the given `value`.
func NewConstantExpressionForValue(value externglib.Value) *ConstantExpressionClass {
	var _arg1 *C.GValue        // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GValue)(unsafe.Pointer(&(&externglib.Value).GValue))

	_cret = C.gtk_constant_expression_new_for_value(_arg1)

	var _constantExpression *ConstantExpressionClass // out

	_constantExpression = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*ConstantExpressionClass)

	return _constantExpression
}

func (*ConstantExpressionClass) privateConstantExpressionClass() {}

// Expression: `GtkExpression` provides a way to describe references to values.
//
// An important aspect of expressions is that the value can be obtained from a
// source that is several steps away. For example, an expression may describe
// ‘the value of property A of `object1`, which is itself the value of a
// property of `object2`’. And `object1` may not even exist yet at the time that
// the expression is created. This is contrast to `GObject` property bindings,
// which can only create direct connections between the properties of two
// objects that must both exist for the duration of the binding.
//
// An expression needs to be "evaluated" to obtain the value that it currently
// refers to. An evaluation always happens in the context of a current object
// called `this` (it mirrors the behavior of object-oriented languages), which
// may or may not influence the result of the evaluation. Use
// [method@Gtk.Expression.evaluate] for evaluating an expression.
//
// Various methods for defining expressions exist, from simple constants via
// [ctor@Gtk.ConstantExpression.new] to looking up properties in a `GObject`
// (even recursively) via [ctor@Gtk.PropertyExpression.new] or providing custom
// functions to transform and combine expressions via
// [ctor@Gtk.ClosureExpression.new].
//
// Here is an example of a complex expression:
//
// “`c color_expr = gtk_property_expression_new (GTK_TYPE_LIST_ITEM, NULL,
// "item"); expression = gtk_property_expression_new (GTK_TYPE_COLOR,
// color_expr, "name"); “`
//
// when evaluated with `this` being a `GtkListItem`, it will obtain the "item"
// property from the `GtkListItem`, and then obtain the "name" property from the
// resulting object (which is assumed to be of type `GTK_TYPE_COLOR`).
//
//
// A more concise way to describe this would be
//
// “` this->item->name “`
//
// The most likely place where you will encounter expressions is in the context
// of list models and list widgets using them. For example, `GtkDropDown` is
// evaluating a `GtkExpression` to obtain strings from the items in its model
// that it can then use to match against the contents of its search entry.
// `GtkStringFilter` is using a `GtkExpression` for similar reasons.
//
// By default, expressions are not paying attention to changes and evaluation is
// just a snapshot of the current state at a given time. To get informed about
// changes, an expression needs to be "watched" via a
// [struct@Gtk.ExpressionWatch], which will cause a callback to be called
// whenever the value of the expression may have changed;
// [method@Gtk.Expression.watch] starts watching an expression, and
// [method@Gtk.ExpressionWatch.unwatch] stops.
//
// Watches can be created for automatically updating the property of an object,
// similar to GObject's `GBinding` mechanism, by using
// [method@Gtk.Expression.bind].
//
//
// GtkExpression in GObject properties
//
// In order to use a `GtkExpression` as a `GObject` property, you must use the
// [id@gtk_param_spec_expression] when creating a `GParamSpec` to install in the
// `GObject` class being defined; for instance:
//
// “`c obj_props[PROP_EXPRESSION] = gtk_param_spec_expression ("expression",
// "Expression", "The expression used by the widget", G_PARAM_READWRITE |
// G_PARAM_STATIC_STRINGS | G_PARAM_EXPLICIT_NOTIFY); “`
//
// When implementing the `GObjectClass.set_property` and
// `GObjectClass.get_property` virtual functions, you must use
// [id@gtk_value_get_expression], to retrieve the stored `GtkExpression` from
// the `GValue` container, and [id@gtk_value_set_expression], to store the
// `GtkExpression` into the `GValue`; for instance:
//
// “`c // in set_property()... case PROP_EXPRESSION: foo_widget_set_expression
// (foo, gtk_value_get_expression (value)); break;
//
//    // in get_property()...
//    case PROP_EXPRESSION:
//      gtk_value_set_expression (value, foo->expression);
//      break;
//
// “`
//
//
// GtkExpression in .ui files
//
// `GtkBuilder` has support for creating expressions. The syntax here can be
// used where a `GtkExpression` object is needed like in a `<property>` tag for
// an expression property, or in a `<binding>` tag to bind a property to an
// expression.
//
// To create an property expression, use the `<lookup>` element. It can have a
// `type` attribute to specify the object type, and a `name` attribute to
// specify the property to look up. The content of `<lookup>` can either be an
// element specfiying the expression to use the object, or a string that
// specifies the name of the object to use.
//
// Example:
//
// “`xml <lookup name='search'>string_filter</lookup> “`
//
// To create a constant expression, use the `<constant>` element. If the type
// attribute is specified, the element content is interpreted as a value of that
// type. Otherwise, it is assumed to be an object. For instance:
//
// “`xml <constant>string_filter</constant> <constant type='gchararray'>Hello,
// world</constant> “`
//
// To create a closure expression, use the `<closure>` element. The `type` and
// `function` attributes specify what function to use for the closure, the
// content of the element contains the expressions for the parameters. For
// instance:
//
// “`xml <closure type='gchararray' function='combine_args_somehow'> <constant
// type='gchararray'>File size:</constant> <lookup type='GFile'
// name='size'>myfile</lookup> </closure> “`
type Expression interface {
	gextras.Objector

	// ValueType gets the `GType` that this expression evaluates to.
	//
	// This type is constant and will not change over the lifetime of this
	// expression.
	ValueType() externglib.Type
	// IsStatic checks if the expression is static.
	//
	// A static expression will never change its result when
	// [method@Gtk.Expression.evaluate] is called on it with the same arguments.
	//
	// That means a call to [method@Gtk.Expression.watch] is not necessary
	// because it will never trigger a notify.
	IsStatic() bool
	// Ref acquires a reference on the given `GtkExpression`.
	ref() *ExpressionClass
	// Unref releases a reference on the given `GtkExpression`.
	//
	// If the reference was the last, the resources associated to the `self` are
	// freed.
	unref()
}

// ExpressionClass implements the Expression interface.
type ExpressionClass struct {
	*externglib.Object
}

var _ Expression = (*ExpressionClass)(nil)

func wrapExpression(obj *externglib.Object) Expression {
	return &ExpressionClass{
		Object: obj,
	}
}

func marshalExpression(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapExpression(obj), nil
}

// ValueType gets the `GType` that this expression evaluates to.
//
// This type is constant and will not change over the lifetime of this
// expression.
func (s *ExpressionClass) ValueType() externglib.Type {
	var _arg0 *C.GtkExpression // out
	var _cret C.GType          // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer((&Expression).Native()))

	_cret = C.gtk_expression_get_value_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(C.GType)

	return _gType
}

// IsStatic checks if the expression is static.
//
// A static expression will never change its result when
// [method@Gtk.Expression.evaluate] is called on it with the same arguments.
//
// That means a call to [method@Gtk.Expression.watch] is not necessary because
// it will never trigger a notify.
func (s *ExpressionClass) IsStatic() bool {
	var _arg0 *C.GtkExpression // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer((&Expression).Native()))

	_cret = C.gtk_expression_is_static(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref acquires a reference on the given `GtkExpression`.
func (s *ExpressionClass) ref() *ExpressionClass {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer((&Expression).Native()))

	_cret = C.gtk_expression_ref(_arg0)

	var _expression *ExpressionClass // out

	_expression = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*ExpressionClass)

	return _expression
}

// Unref releases a reference on the given `GtkExpression`.
//
// If the reference was the last, the resources associated to the `self` are
// freed.
func (s *ExpressionClass) unref() {
	var _arg0 *C.GtkExpression // out

	_arg0 = (*C.GtkExpression)(unsafe.Pointer((&Expression).Native()))

	C.gtk_expression_unref(_arg0)
}

// ObjectExpression: `GObject` value in a `GtkExpression`.
type ObjectExpression interface {
	gextras.Objector

	// Object gets the object that the expression evaluates to.
	Object() *externglib.Object
}

// ObjectExpressionClass implements the ObjectExpression interface.
type ObjectExpressionClass struct {
	ExpressionClass
}

var _ ObjectExpression = (*ObjectExpressionClass)(nil)

func wrapObjectExpression(obj *externglib.Object) ObjectExpression {
	return &ObjectExpressionClass{
		ExpressionClass: ExpressionClass{
			Object: obj,
		},
	}
}

func marshalObjectExpression(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapObjectExpression(obj), nil
}

// NewObjectExpression creates an expression evaluating to the given `object`
// with a weak reference.
//
// Once the `object` is disposed, it will fail to evaluate.
//
// This expression is meant to break reference cycles.
//
// If you want to keep a reference to `object`, use
// [ctor@Gtk.ConstantExpression.new].
func NewObjectExpression(object gextras.Objector) *ObjectExpressionClass {
	var _arg1 *C.GObject       // out
	var _cret *C.GtkExpression // in

	_arg1 = (*C.GObject)(unsafe.Pointer((&gextras.Objector).Native()))

	_cret = C.gtk_object_expression_new(_arg1)

	var _objectExpression *ObjectExpressionClass // out

	_objectExpression = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*ObjectExpressionClass)

	return _objectExpression
}

// Object gets the object that the expression evaluates to.
func (e *ObjectExpressionClass) Object() *externglib.Object {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GObject       // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer((&ObjectExpression).Native()))

	_cret = C.gtk_object_expression_get_object(_arg0)

	var _object *externglib.Object // out

	_object = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*externglib.Object)

	return _object
}

// PropertyExpression: `GObject` property value in a `GtkExpression`.
type PropertyExpression interface {
	gextras.Objector

	// Expression gets the expression specifying the object of a property
	// expression.
	Expression() *ExpressionClass
}

// PropertyExpressionClass implements the PropertyExpression interface.
type PropertyExpressionClass struct {
	ExpressionClass
}

var _ PropertyExpression = (*PropertyExpressionClass)(nil)

func wrapPropertyExpression(obj *externglib.Object) PropertyExpression {
	return &PropertyExpressionClass{
		ExpressionClass: ExpressionClass{
			Object: obj,
		},
	}
}

func marshalPropertyExpression(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPropertyExpression(obj), nil
}

// NewPropertyExpression creates an expression that looks up a property via the
// given `expression` or the `this` argument when `expression` is `NULL`.
//
// If the resulting object conforms to `this_type`, its property named
// `property_name` will be queried. Otherwise, this expression's evaluation will
// fail.
//
// The given `this_type` must have a property with `property_name`.
func NewPropertyExpression(thisType externglib.Type, expression Expression, propertyName string) *PropertyExpressionClass {
	var _arg1 C.GType          // out
	var _arg2 *C.GtkExpression // out
	var _arg3 *C.char          // out
	var _cret *C.GtkExpression // in

	_arg1 = (C.GType)(externglib.Type)
	_arg2 = (*C.GtkExpression)(unsafe.Pointer((&Expression).Native()))
	_arg3 = (*C.char)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_property_expression_new(_arg1, _arg2, _arg3)

	var _propertyExpression *PropertyExpressionClass // out

	_propertyExpression = gextras.CastObject(
		externglib.AssumeOwnership(unsafe.Pointer(_cret))).(*PropertyExpressionClass)

	return _propertyExpression
}

// Expression gets the expression specifying the object of a property
// expression.
func (e *PropertyExpressionClass) Expression() *ExpressionClass {
	var _arg0 *C.GtkExpression // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkExpression)(unsafe.Pointer((&PropertyExpression).Native()))

	_cret = C.gtk_property_expression_get_expression(_arg0)

	var _ret *ExpressionClass // out

	_ret = gextras.CastObject(
		externglib.Take(unsafe.Pointer(_cret))).(*ExpressionClass)

	return _ret
}

// ExpressionWatch: opaque structure representing a watched `GtkExpression`.
//
// The contents of `GtkExpressionWatch` should only be accessed through the
// provided API.
type ExpressionWatch struct {
	native C.GtkExpressionWatch
}

// WrapExpressionWatch wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapExpressionWatch(ptr unsafe.Pointer) *ExpressionWatch {
	return (*ExpressionWatch)(ptr)
}

func marshalExpressionWatch(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*ExpressionWatch)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *ExpressionWatch) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// Evaluate evaluates the watched expression and on success stores the result in
// `value`.
//
// This is equivalent to calling [method@Gtk.Expression.evaluate] with the
// expression and this pointer originally used to create `watch`.
func (w *ExpressionWatch) Evaluate(value externglib.Value) bool {
	var _arg0 *C.GtkExpressionWatch // out
	var _arg1 *C.GValue             // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(*ExpressionWatch))
	_arg1 = (*C.GValue)(unsafe.Pointer(&(&externglib.Value).GValue))

	_cret = C.gtk_expression_watch_evaluate(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref acquires a reference on the given `GtkExpressionWatch`.
func (w *ExpressionWatch) ref() *ExpressionWatch {
	var _arg0 *C.GtkExpressionWatch // out
	var _cret *C.GtkExpressionWatch // in

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(*ExpressionWatch))

	_cret = C.gtk_expression_watch_ref(_arg0)

	var _expressionWatch *ExpressionWatch // out

	_expressionWatch = (*ExpressionWatch)(unsafe.Pointer(*C.GtkExpressionWatch))
	C.gtk_expression_watch_ref(_cret)
	runtime.SetFinalizer(_expressionWatch, func(v *ExpressionWatch) {
		C.gtk_expression_watch_unref((*C.GtkExpressionWatch)(unsafe.Pointer(v)))
	})

	return _expressionWatch
}

// Unref releases a reference on the given `GtkExpressionWatch`.
//
// If the reference was the last, the resources associated to `self` are freed.
func (w *ExpressionWatch) unref() {
	var _arg0 *C.GtkExpressionWatch // out

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(*ExpressionWatch))

	C.gtk_expression_watch_unref(_arg0)
}

// Unwatch stops watching an expression.
//
// See [method@Gtk.Expression.watch] for how the watch was established.
func (w *ExpressionWatch) Unwatch() {
	var _arg0 *C.GtkExpressionWatch // out

	_arg0 = (*C.GtkExpressionWatch)(unsafe.Pointer(*ExpressionWatch))

	C.gtk_expression_watch_unwatch(_arg0)
}
