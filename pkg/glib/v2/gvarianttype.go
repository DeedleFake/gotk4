// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_variant_type_get_gtype()), F: marshalVariantType},
	})
}

// VariantType: this section introduces the GVariant type system. It is based,
// in large part, on the D-Bus type system, with two major changes and some
// minor lifting of restrictions. The D-Bus specification
// (http://dbus.freedesktop.org/doc/dbus-specification.html), therefore,
// provides a significant amount of information that is useful when working with
// GVariant.
//
// The first major change with respect to the D-Bus type system is the
// introduction of maybe (or "nullable") types. Any type in GVariant can be
// converted to a maybe type, in which case, "nothing" (or "null") becomes a
// valid value. Maybe types have been added by introducing the character "m" to
// type strings.
//
// The second major change is that the GVariant type system supports the concept
// of "indefinite types" -- types that are less specific than the normal types
// found in D-Bus. For example, it is possible to speak of "an array of any
// type" in GVariant, where the D-Bus type system would require you to speak of
// "an array of integers" or "an array of strings". Indefinite types have been
// added by introducing the characters "*", "?" and "r" to type strings.
//
// Finally, all arbitrary restrictions relating to the complexity of types are
// lifted along with the restriction that dictionary entries may only appear
// nested inside of arrays.
//
// Just as in D-Bus, GVariant types are described with strings ("type strings").
// Subject to the differences mentioned above, these strings are of the same
// form as those found in D-Bus. Note, however: D-Bus always works in terms of
// messages and therefore individual type strings appear nowhere in its
// interface. Instead, "signatures" are a concatenation of the strings of the
// type of each argument in a message. GVariant deals with single values
// directly so GVariant type strings always describe the type of exactly one
// value. This means that a D-Bus signature string is generally not a valid
// GVariant type string -- except in the case that it is the signature of a
// message containing exactly one argument.
//
// An indefinite type is similar in spirit to what may be called an abstract
// type in other type systems. No value can exist that has an indefinite type as
// its type, but values can exist that have types that are subtypes of
// indefinite types. That is to say, g_variant_get_type() will never return an
// indefinite type, but calling g_variant_is_of_type() with an indefinite type
// may return TRUE. For example, you cannot have a value that represents "an
// array of no particular type", but you can have an "array of integers" which
// certainly matches the type of "an array of no particular type", since "array
// of integers" is a subtype of "array of no particular type".
//
// This is similar to how instances of abstract classes may not directly exist
// in other type systems, but instances of their non-abstract subtypes may. For
// example, in GTK, no object that has the type of Bin can exist (since Bin is
// an abstract class), but a Window can certainly be instantiated, and you would
// say that the Window is a Bin (since Window is a subclass of Bin).
//
//
// GVariant Type Strings
//
// A GVariant type string can be any of the following:
//
// - any basic type string (listed below)
//
// - "v", "r" or "*"
//
// - one of the characters 'a' or 'm', followed by another type string
//
// - the character '(', followed by a concatenation of zero or more other type
// strings, followed by the character ')'
//
// - the character '{', followed by a basic type string (see below), followed by
// another type string, followed by the character '}'
//
// A basic type string describes a basic type (as per g_variant_type_is_basic())
// and is always a single character in length. The valid basic type strings are
// "b", "y", "n", "q", "i", "u", "x", "t", "h", "d", "s", "o", "g" and "?".
//
// The above definition is recursive to arbitrary depth. "aaaaai" and
// "(ui(nq((y)))s)" are both valid type strings, as is "a(aa(ui)(qna{ya(yd)}))".
// In order to not hit memory limits, #GVariant imposes a limit on recursion
// depth of 65 nested containers. This is the limit in the D-Bus specification
// (64) plus one to allow a BusMessage to be nested in a top-level tuple.
//
// The meaning of each of the characters is as follows:
//
// - b: the type string of G_VARIANT_TYPE_BOOLEAN; a boolean value.
//
// - y: the type string of G_VARIANT_TYPE_BYTE; a byte.
//
// - n: the type string of G_VARIANT_TYPE_INT16; a signed 16 bit integer.
//
// - q: the type string of G_VARIANT_TYPE_UINT16; an unsigned 16 bit integer.
//
// - i: the type string of G_VARIANT_TYPE_INT32; a signed 32 bit integer.
//
// - u: the type string of G_VARIANT_TYPE_UINT32; an unsigned 32 bit integer.
//
// - x: the type string of G_VARIANT_TYPE_INT64; a signed 64 bit integer.
//
// - t: the type string of G_VARIANT_TYPE_UINT64; an unsigned 64 bit integer.
//
// - h: the type string of G_VARIANT_TYPE_HANDLE; a signed 32 bit value that, by
// convention, is used as an index into an array of file descriptors that are
// sent alongside a D-Bus message.
//
// - d: the type string of G_VARIANT_TYPE_DOUBLE; a double precision floating
// point value.
//
// - s: the type string of G_VARIANT_TYPE_STRING; a string.
//
// - o: the type string of G_VARIANT_TYPE_OBJECT_PATH; a string in the form of a
// D-Bus object path.
//
// - g: the type string of G_VARIANT_TYPE_SIGNATURE; a string in the form of a
// D-Bus type signature.
//
// - ?: the type string of G_VARIANT_TYPE_BASIC; an indefinite type that is a
// supertype of any of the basic types.
//
// - v: the type string of G_VARIANT_TYPE_VARIANT; a container type that contain
// any other type of value.
//
// - a: used as a prefix on another type string to mean an array of that type;
// the type string "ai", for example, is the type of an array of signed 32-bit
// integers.
//
// - m: used as a prefix on another type string to mean a "maybe", or
// "nullable", version of that type; the type string "ms", for example, is the
// type of a value that maybe contains a string, or maybe contains nothing.
//
// - (): used to enclose zero or more other concatenated type strings to create
// a tuple type; the type string "(is)", for example, is the type of a pair of
// an integer and a string.
//
// - r: the type string of G_VARIANT_TYPE_TUPLE; an indefinite type that is a
// supertype of any tuple type, regardless of the number of items.
//
// - {}: used to enclose a basic type string concatenated with another type
// string to create a dictionary entry type, which usually appears inside of an
// array to form a dictionary; the type string "a{sd}", for example, is the type
// of a dictionary that maps strings to double precision floating point values.
//
//    The first type (the basic type) is the key type and the second type is
//    the value type. The reason that the first type is restricted to being a
//    basic type is so that it can easily be hashed.
//
// - *: the type string of G_VARIANT_TYPE_ANY; the indefinite type that is a
// supertype of all types. Note that, as with all type strings, this character
// represents exactly one type. It cannot be used inside of tuples to mean "any
// number of items".
//
// Any type string of a container that contains an indefinite type is, itself,
// an indefinite type. For example, the type string "a*" (corresponding to
// G_VARIANT_TYPE_ARRAY) is an indefinite type that is a supertype of every
// array type. "(*s)" is a supertype of all tuples that contain exactly two
// items where the second item is a string.
//
// "a{?*}" is an indefinite type that is a supertype of all arrays containing
// dictionary entries where the key is any basic type and the value is any type
// at all. This is, by definition, a dictionary, so this type string corresponds
// to G_VARIANT_TYPE_DICTIONARY. Note that, due to the restriction that the key
// of a dictionary entry must be a basic type, "{**}" is not a valid type
// string.
type VariantType struct {
	nocopy gextras.NoCopy
	native *C.GVariantType
}

func marshalVariantType(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &VariantType{native: (*C.GVariantType)(unsafe.Pointer(b))}, nil
}

// NewVariantType constructs a struct VariantType.
func NewVariantType(typeString string) *VariantType {
	var _arg1 *C.gchar        // out
	var _cret *C.GVariantType // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typeString)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_variant_type_new(_arg1)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_variantType, func(v *VariantType) {
		C.g_variant_type_free((*C.GVariantType)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variantType
}

// NewVariantTypeArray constructs a struct VariantType.
func NewVariantTypeArray(element *VariantType) *VariantType {
	var _arg1 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg1 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(element)))

	_cret = C.g_variant_type_new_array(_arg1)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_variantType, func(v *VariantType) {
		C.g_variant_type_free((*C.GVariantType)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variantType
}

// NewVariantTypeDictEntry constructs a struct VariantType.
func NewVariantTypeDictEntry(key *VariantType, value *VariantType) *VariantType {
	var _arg1 *C.GVariantType // out
	var _arg2 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg1 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(key)))
	_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(value)))

	_cret = C.g_variant_type_new_dict_entry(_arg1, _arg2)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_variantType, func(v *VariantType) {
		C.g_variant_type_free((*C.GVariantType)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variantType
}

// NewVariantTypeMaybe constructs a struct VariantType.
func NewVariantTypeMaybe(element *VariantType) *VariantType {
	var _arg1 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg1 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(element)))

	_cret = C.g_variant_type_new_maybe(_arg1)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_variantType, func(v *VariantType) {
		C.g_variant_type_free((*C.GVariantType)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variantType
}

// NewVariantTypeTuple constructs a struct VariantType.
func NewVariantTypeTuple(items []*VariantType) *VariantType {
	var _arg1 **C.GVariantType // out
	var _arg2 C.gint
	var _cret *C.GVariantType // in

	_arg2 = (C.gint)(len(items))
	if len(items) > 0 {
		_arg1 = (**C.GVariantType)(unsafe.Pointer(&items[0]))
	}

	_cret = C.g_variant_type_new_tuple(_arg1, _arg2)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_variantType, func(v *VariantType) {
		C.g_variant_type_free((*C.GVariantType)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variantType
}

// Copy makes a copy of a Type. It is appropriate to call g_variant_type_free()
// on the return value. type may not be NULL.
func (typ *VariantType) Copy() *VariantType {
	var _arg0 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_copy(_arg0)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_variantType, func(v *VariantType) {
		C.g_variant_type_free((*C.GVariantType)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variantType
}

// DupString returns a newly-allocated copy of the type string corresponding to
// type. The returned string is nul-terminated. It is appropriate to call
// g_free() on the return value.
func (typ *VariantType) DupString() string {
	var _arg0 *C.GVariantType // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_dup_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Element determines the element type of an array or maybe type.
//
// This function may only be used with array or maybe types.
func (typ *VariantType) Element() *VariantType {
	var _arg0 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_element(_arg0)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _variantType
}

// Equal compares type1 and type2 for equality.
//
// Only returns TRUE if the types are exactly equal. Even if one type is an
// indefinite type and the other is a subtype of it, FALSE will be returned if
// they are not exactly equal. If you want to check for subtypes, use
// g_variant_type_is_subtype_of().
//
// The argument types of type1 and type2 are only #gconstpointer to allow use
// with Table without function pointer casting. For both arguments, a valid Type
// must be provided.
func (type1 *VariantType) Equal(type2 *VariantType) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(type1)))
	_arg1 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(type2)))

	_cret = C.g_variant_type_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// First determines the first item type of a tuple or dictionary entry type.
//
// This function may only be used with tuple or dictionary entry types, but must
// not be used with the generic tuple type G_VARIANT_TYPE_TUPLE.
//
// In the case of a dictionary entry type, this returns the type of the key.
//
// NULL is returned in case of type being G_VARIANT_TYPE_UNIT.
//
// This call, together with g_variant_type_next() provides an iterator interface
// over tuple and dictionary entry types.
func (typ *VariantType) First() *VariantType {
	var _arg0 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_first(_arg0)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _variantType
}

// StringLength returns the length of the type string corresponding to the given
// type. This function must be used to determine the valid extent of the memory
// region returned by g_variant_type_peek_string().
func (typ *VariantType) StringLength() uint {
	var _arg0 *C.GVariantType // out
	var _cret C.gsize         // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_get_string_length(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Hash hashes type.
//
// The argument type of type is only #gconstpointer to allow use with Table
// without function pointer casting. A valid Type must be provided.
func (typ *VariantType) Hash() uint {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_hash(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IsArray determines if the given type is an array type. This is true if the
// type string for type starts with an 'a'.
//
// This function returns TRUE for any indefinite type for which every definite
// subtype is an array type -- G_VARIANT_TYPE_ARRAY, for example.
func (typ *VariantType) IsArray() bool {
	var _arg0 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_is_array(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsBasic determines if the given type is a basic type.
//
// Basic types are booleans, bytes, integers, doubles, strings, object paths and
// signatures.
//
// Only a basic type may be used as the key of a dictionary entry.
//
// This function returns FALSE for all indefinite types except
// G_VARIANT_TYPE_BASIC.
func (typ *VariantType) IsBasic() bool {
	var _arg0 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_is_basic(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsContainer determines if the given type is a container type.
//
// Container types are any array, maybe, tuple, or dictionary entry types plus
// the variant type.
//
// This function returns TRUE for any indefinite type for which every definite
// subtype is a container -- G_VARIANT_TYPE_ARRAY, for example.
func (typ *VariantType) IsContainer() bool {
	var _arg0 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_is_container(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsDefinite determines if the given type is definite (ie: not indefinite).
//
// A type is definite if its type string does not contain any indefinite type
// characters ('*', '?', or 'r').
//
// A #GVariant instance may not have an indefinite type, so calling this
// function on the result of g_variant_get_type() will always result in TRUE
// being returned. Calling this function on an indefinite type like
// G_VARIANT_TYPE_ARRAY, however, will result in FALSE being returned.
func (typ *VariantType) IsDefinite() bool {
	var _arg0 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_is_definite(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsDictEntry determines if the given type is a dictionary entry type. This is
// true if the type string for type starts with a '{'.
//
// This function returns TRUE for any indefinite type for which every definite
// subtype is a dictionary entry type -- G_VARIANT_TYPE_DICT_ENTRY, for example.
func (typ *VariantType) IsDictEntry() bool {
	var _arg0 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_is_dict_entry(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsMaybe determines if the given type is a maybe type. This is true if the
// type string for type starts with an 'm'.
//
// This function returns TRUE for any indefinite type for which every definite
// subtype is a maybe type -- G_VARIANT_TYPE_MAYBE, for example.
func (typ *VariantType) IsMaybe() bool {
	var _arg0 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_is_maybe(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSubtypeOf checks if type is a subtype of supertype.
//
// This function returns TRUE if type is a subtype of supertype. All types are
// considered to be subtypes of themselves. Aside from that, only indefinite
// types can have subtypes.
func (typ *VariantType) IsSubtypeOf(supertype *VariantType) bool {
	var _arg0 *C.GVariantType // out
	var _arg1 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))
	_arg1 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(supertype)))

	_cret = C.g_variant_type_is_subtype_of(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsTuple determines if the given type is a tuple type. This is true if the
// type string for type starts with a '(' or if type is G_VARIANT_TYPE_TUPLE.
//
// This function returns TRUE for any indefinite type for which every definite
// subtype is a tuple type -- G_VARIANT_TYPE_TUPLE, for example.
func (typ *VariantType) IsTuple() bool {
	var _arg0 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_is_tuple(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsVariant determines if the given type is the variant type.
func (typ *VariantType) IsVariant() bool {
	var _arg0 *C.GVariantType // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_is_variant(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Key determines the key type of a dictionary entry type.
//
// This function may only be used with a dictionary entry type. Other than the
// additional restriction, this call is equivalent to g_variant_type_first().
func (typ *VariantType) Key() *VariantType {
	var _arg0 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_key(_arg0)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _variantType
}

// NItems determines the number of items contained in a tuple or dictionary
// entry type.
//
// This function may only be used with tuple or dictionary entry types, but must
// not be used with the generic tuple type G_VARIANT_TYPE_TUPLE.
//
// In the case of a dictionary entry type, this function will always return 2.
func (typ *VariantType) NItems() uint {
	var _arg0 *C.GVariantType // out
	var _cret C.gsize         // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_n_items(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Next determines the next item type of a tuple or dictionary entry type.
//
// type must be the result of a previous call to g_variant_type_first() or
// g_variant_type_next().
//
// If called on the key type of a dictionary entry then this call returns the
// value type. If called on the value type of a dictionary entry then this call
// returns NULL.
//
// For tuples, NULL is returned when type is the last item in a tuple.
func (typ *VariantType) Next() *VariantType {
	var _arg0 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_next(_arg0)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _variantType
}

// Value determines the value type of a dictionary entry type.
//
// This function may only be used with a dictionary entry type.
func (typ *VariantType) Value() *VariantType {
	var _arg0 *C.GVariantType // out
	var _cret *C.GVariantType // in

	_arg0 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_variant_type_value(_arg0)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _variantType
}

func VariantTypeChecked_(arg0 string) *VariantType {
	var _arg1 *C.gchar        // out
	var _cret *C.GVariantType // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(arg0)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_variant_type_checked_(_arg1)

	var _variantType *VariantType // out

	_variantType = (*VariantType)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _variantType
}

func VariantTypeStringGetDepth_(typeString string) uint {
	var _arg1 *C.gchar // out
	var _cret C.gsize  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typeString)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_variant_type_string_get_depth_(_arg1)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// VariantTypeStringIsValid checks if type_string is a valid GVariant type
// string. This call is equivalent to calling g_variant_type_string_scan() and
// confirming that the following character is a nul terminator.
func VariantTypeStringIsValid(typeString string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(typeString)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_variant_type_string_is_valid(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VariantTypeStringScan: scan for a single complete and valid GVariant type
// string in string. The memory pointed to by limit (or bytes beyond it) is
// never accessed.
//
// If a valid type string is found, endptr is updated to point to the first
// character past the end of the string that was found and TRUE is returned.
//
// If there is no valid type string starting at string, or if the type string
// does not end before limit then FALSE is returned.
//
// For the simple case of checking if a string is a valid type string, see
// g_variant_type_string_is_valid().
func VariantTypeStringScan(_string string, limit string) (string, bool) {
	var _arg1 *C.gchar   // out
	var _arg2 *C.gchar   // out
	var _arg3 *C.gchar   // in
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(limit)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_variant_type_string_scan(_arg1, _arg2, &_arg3)

	var _endptr string // out
	var _ok bool       // out

	_endptr = C.GoString((*C.gchar)(unsafe.Pointer(_arg3)))
	defer C.free(unsafe.Pointer(_arg3))
	if _cret != 0 {
		_ok = true
	}

	return _endptr, _ok
}
