// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// GType _gotk4_gtk4_BuilderScope_virtual_get_type_from_function(void* fnptr, GtkBuilderScope* arg0, GtkBuilder* arg1, char* arg2) {
//   return ((GType (*)(GtkBuilderScope*, GtkBuilder*, char*))(fnptr))(arg0, arg1, arg2);
// };
// GType _gotk4_gtk4_BuilderScope_virtual_get_type_from_name(void* fnptr, GtkBuilderScope* arg0, GtkBuilder* arg1, char* arg2) {
//   return ((GType (*)(GtkBuilderScope*, GtkBuilder*, char*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeBuilderClosureFlags = coreglib.Type(C.gtk_builder_closure_flags_get_type())
	GTypeBuilderScope        = coreglib.Type(C.gtk_builder_scope_get_type())
	GTypeBuilderCScope       = coreglib.Type(C.gtk_builder_cscope_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBuilderClosureFlags, F: marshalBuilderClosureFlags},
		coreglib.TypeMarshaler{T: GTypeBuilderScope, F: marshalBuilderScope},
		coreglib.TypeMarshaler{T: GTypeBuilderCScope, F: marshalBuilderCScope},
	})
}

// BuilderClosureFlags: list of flags that can be passed to
// gtk_builder_create_closure().
//
// New values may be added in the future for new features, so external
// implementations of gtk.BuilderScope should test the flags for unknown values
// and raise a GTK_BUILDER_ERROR_INVALID_ATTRIBUTE error when they encounter
// one.
type BuilderClosureFlags C.guint

const (
	// BuilderClosureSwapped: closure should be created swapped. See
	// g_cclosure_new_swap() for details.
	BuilderClosureSwapped BuilderClosureFlags = 0b1
)

func marshalBuilderClosureFlags(p uintptr) (interface{}, error) {
	return BuilderClosureFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for BuilderClosureFlags.
func (b BuilderClosureFlags) String() string {
	if b == 0 {
		return "BuilderClosureFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(21)

	for b != 0 {
		next := b & (b - 1)
		bit := b - next

		switch bit {
		case BuilderClosureSwapped:
			builder.WriteString("Swapped|")
		default:
			builder.WriteString(fmt.Sprintf("BuilderClosureFlags(0b%b)|", bit))
		}

		b = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if b contains other.
func (b BuilderClosureFlags) Has(other BuilderClosureFlags) bool {
	return (b & other) == other
}

// BuilderScope: GtkBuilderScope is an interface to provide language binding
// support to GtkBuilder.
//
// The goal of GtkBuilderScope is to look up programming-language-specific
// values for strings that are given in a GtkBuilder UI file.
//
// The primary intended audience is bindings that want to provide deeper
// integration of GtkBuilder into the language.
//
// A GtkBuilderScope instance may be used with multiple GtkBuilder objects, even
// at once.
//
// By default, GTK will use its own implementation of GtkBuilderScope for the C
// language which can be created via gtk.BuilderCScope.New.
//
// BuilderScope wraps an interface. This means the user can get the
// underlying type by calling Cast().
type BuilderScope struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*BuilderScope)(nil)
)

// BuilderScoper describes BuilderScope's interface methods.
type BuilderScoper interface {
	coreglib.Objector

	baseBuilderScope() *BuilderScope
}

var _ BuilderScoper = (*BuilderScope)(nil)

func wrapBuilderScope(obj *coreglib.Object) *BuilderScope {
	return &BuilderScope{
		Object: obj,
	}
}

func marshalBuilderScope(p uintptr) (interface{}, error) {
	return wrapBuilderScope(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *BuilderScope) baseBuilderScope() *BuilderScope {
	return v
}

// BaseBuilderScope returns the underlying base object.
func BaseBuilderScope(obj BuilderScoper) *BuilderScope {
	return obj.baseBuilderScope()
}

// The function takes the following parameters:
//
//    - builder
//    - functionName
//
// The function returns the following values:
//
func (self *BuilderScope) typeFromFunction(builder *Builder, functionName string) coreglib.Type {
	gclass := (*C.GtkBuilderScopeInterface)(coreglib.PeekParentClass(self))
	fnarg := gclass.get_type_from_function

	var _arg0 *C.GtkBuilderScope // out
	var _arg1 *C.GtkBuilder      // out
	var _arg2 *C.char            // out
	var _cret C.GType            // in

	_arg0 = (*C.GtkBuilderScope)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(functionName)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C._gotk4_gtk4_BuilderScope_virtual_get_type_from_function(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(functionName)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// The function takes the following parameters:
//
//    - builder
//    - typeName
//
// The function returns the following values:
//
func (self *BuilderScope) typeFromName(builder *Builder, typeName string) coreglib.Type {
	gclass := (*C.GtkBuilderScopeInterface)(coreglib.PeekParentClass(self))
	fnarg := gclass.get_type_from_name

	var _arg0 *C.GtkBuilderScope // out
	var _arg1 *C.GtkBuilder      // out
	var _arg2 *C.char            // out
	var _cret C.GType            // in

	_arg0 = (*C.GtkBuilderScope)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(coreglib.InternObject(builder).Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(typeName)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C._gotk4_gtk4_BuilderScope_virtual_get_type_from_name(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(typeName)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// BuilderCScopeOverrides contains methods that are overridable.
type BuilderCScopeOverrides struct {
}

func defaultBuilderCScopeOverrides(v *BuilderCScope) BuilderCScopeOverrides {
	return BuilderCScopeOverrides{}
}

// BuilderCScope: GtkBuilderScope implementation for the C language.
//
// GtkBuilderCScope instances use symbols explicitly added to builder with prior
// calls to gtk.BuilderCScope.AddCallbackSymbol(). If developers want to do
// that, they are encouraged to create their own scopes for that purpose.
//
// In the case that symbols are not explicitly added; GTK will uses GModule’s
// introspective features (by opening the module NULL) to look at the
// application’s symbol table. From here it tries to match the signal function
// names given in the interface description with symbols in the application.
//
// Note that unless gtk.BuilderCScope.AddCallbackSymbol() is called for all
// signal callbacks which are referenced by the loaded XML, this functionality
// will require that GModule be supported on the platform.
type BuilderCScope struct {
	_ [0]func() // equal guard
	*coreglib.Object

	BuilderScope
}

var (
	_ coreglib.Objector = (*BuilderCScope)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*BuilderCScope, *BuilderCScopeClass, BuilderCScopeOverrides](
		GTypeBuilderCScope,
		initBuilderCScopeClass,
		wrapBuilderCScope,
		defaultBuilderCScopeOverrides,
	)
}

func initBuilderCScopeClass(gclass unsafe.Pointer, overrides BuilderCScopeOverrides, classInitFunc func(*BuilderCScopeClass)) {
	if classInitFunc != nil {
		class := (*BuilderCScopeClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapBuilderCScope(obj *coreglib.Object) *BuilderCScope {
	return &BuilderCScope{
		Object: obj,
		BuilderScope: BuilderScope{
			Object: obj,
		},
	}
}

func marshalBuilderCScope(p uintptr) (interface{}, error) {
	return wrapBuilderCScope(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBuilderCScope creates a new GtkBuilderCScope object to use with future
// GtkBuilder instances.
//
// Calling this function is only necessary if you want to add custom callbacks
// via gtk.BuilderCScope.AddCallbackSymbol().
//
// The function returns the following values:
//
//    - builderCScope: new GtkBuilderCScope.
//
func NewBuilderCScope() *BuilderCScope {
	var _cret *C.GtkBuilderScope // in

	_cret = C.gtk_builder_cscope_new()

	var _builderCScope *BuilderCScope // out

	_builderCScope = wrapBuilderCScope(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builderCScope
}

// BuilderCScopeClass: instance of this type is always passed by reference.
type BuilderCScopeClass struct {
	*builderCScopeClass
}

// builderCScopeClass is the struct that's finalized.
type builderCScopeClass struct {
	native *C.GtkBuilderCScopeClass
}

// BuilderScopeInterface: virtual function table to implement for BuilderScope
// implementations. Default implementations for each function do exist, but they
// usually just fail, so it is suggested that implementations implement all of
// them.
//
// An instance of this type is always passed by reference.
type BuilderScopeInterface struct {
	*builderScopeInterface
}

// builderScopeInterface is the struct that's finalized.
type builderScopeInterface struct {
	native *C.GtkBuilderScopeInterface
}
