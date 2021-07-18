// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_app_info_get_type()), F: marshalAppInfor},
		{T: externglib.Type(C.g_app_launch_context_get_type()), F: marshalAppLaunchContexter},
	})
}

// AppInfoOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AppInfoOverrider interface {
	// AddSupportsType adds a content type to the application information to
	// indicate the application is capable of opening files with the given
	// content type.
	AddSupportsType(contentType string) error
	// CanDelete obtains the information whether the Info can be deleted. See
	// g_app_info_delete().
	CanDelete() bool
	// CanRemoveSupportsType checks if a supported content type can be removed
	// from an application.
	CanRemoveSupportsType() bool
	// DoDelete tries to delete a Info.
	//
	// On some platforms, there may be a difference between user-defined Infos
	// which can be deleted, and system-wide ones which cannot. See
	// g_app_info_can_delete().
	DoDelete() bool
	// Dup creates a duplicate of a Info.
	Dup() AppInfor
	// Equal checks if two Infos are equal.
	//
	// Note that the check *may not* compare each individual field, and only
	// does an identity check. In case detecting changes in the contents is
	// needed, program code must additionally compare relevant fields.
	Equal(appinfo2 AppInfor) bool
	// Commandline gets the commandline with which the application will be
	// started.
	Commandline() string
	// Description gets a human-readable description of an installed
	// application.
	Description() string
	// DisplayName gets the display name of the application. The display name is
	// often more descriptive to the user than the name itself.
	DisplayName() string
	// Executable gets the executable's name for the installed application.
	Executable() string
	// Icon gets the icon for the application.
	Icon() Iconer
	// ID gets the ID of an application. An id is a string that identifies the
	// application. The exact format of the id is platform dependent. For
	// instance, on Unix this is the desktop file id from the xdg menu
	// specification.
	//
	// Note that the returned ID may be NULL, depending on how the appinfo has
	// been constructed.
	ID() string
	// Name gets the installed name of the application.
	Name() string
	// SupportedTypes retrieves the list of content types that app_info claims
	// to support. If this information is not provided by the environment, this
	// function will return NULL. This function does not take in consideration
	// associations added with g_app_info_add_supports_type(), but only those
	// exported directly by the application.
	SupportedTypes() []string
	// LaunchUrisFinish finishes a g_app_info_launch_uris_async() operation.
	LaunchUrisFinish(result AsyncResulter) error
	// RemoveSupportsType removes a supported type from an application, if
	// possible.
	RemoveSupportsType(contentType string) error
	// SetAsDefaultForExtension sets the application as the default handler for
	// the given file extension.
	SetAsDefaultForExtension(extension string) error
	// SetAsDefaultForType sets the application as the default handler for a
	// given type.
	SetAsDefaultForType(contentType string) error
	// SetAsLastUsedForType sets the application as the last used application
	// for a given type. This will make the application appear as first in the
	// list returned by g_app_info_get_recommended_for_type(), regardless of the
	// default application for that content type.
	SetAsLastUsedForType(contentType string) error
	// ShouldShow checks if the application info should be shown in menus that
	// list available applications.
	ShouldShow() bool
	// SupportsFiles checks if the application accepts files as arguments.
	SupportsFiles() bool
	// SupportsUris checks if the application supports reading files and
	// directories from URIs.
	SupportsUris() bool
}

// AppInfo and LaunchContext are used for describing and launching applications
// installed on the system.
//
// As of GLib 2.20, URIs will always be converted to POSIX paths (using
// g_file_get_path()) when using g_app_info_launch() even if the application
// requested an URI and not a POSIX path. For example for a desktop-file based
// application with Exec key totem U and a single URI, sftp://foo/file.avi, then
// /home/user/.gvfs/sftp on foo/file.avi will be passed. This will only work if
// a set of suitable GIO extensions (such as gvfs 2.26 compiled with FUSE
// support), is available and operational; if this is not the case, the URI will
// be passed unmodified to the application. Some URIs, such as mailto:, of
// course cannot be mapped to a POSIX path (in gvfs there's no FUSE mount for
// it); such URIs will be passed unmodified to the application.
//
// Specifically for gvfs 2.26 and later, the POSIX URI will be mapped back to
// the GIO URI in the #GFile constructors (since gvfs implements the #GVfs
// extension point). As such, if the application needs to examine the URI, it
// needs to use g_file_get_uri() or similar on #GFile. In other words, an
// application cannot assume that the URI passed to e.g.
// g_file_new_for_commandline_arg() is equal to the result of g_file_get_uri().
// The following snippet illustrates this:
//
//    GFile *f;
//    char *uri;
//
//    file = g_file_new_for_commandline_arg (uri_from_commandline);
//
//    uri = g_file_get_uri (file);
//    strcmp (uri, uri_from_commandline) == 0;
//    g_free (uri);
//
//    if (g_file_has_uri_scheme (file, "cdda"))
//      {
//        // do something special with uri
//      }
//    g_object_unref (file);
//
// This code will work when both cdda://sr0/Track 1.wav and
// /home/user/.gvfs/cdda on sr0/Track 1.wav is passed to the application. It
// should be noted that it's generally not safe for applications to rely on the
// format of a particular URIs. Different launcher applications (e.g. file
// managers) may have different ideas of what a given URI means.
type AppInfo struct {
	*externglib.Object
}

var _ gextras.Nativer = (*AppInfo)(nil)

// AppInfor describes AppInfo's abstract methods.
type AppInfor interface {
	// AddSupportsType adds a content type to the application information to
	// indicate the application is capable of opening files with the given
	// content type.
	AddSupportsType(contentType string) error
	// CanDelete obtains the information whether the Info can be deleted.
	CanDelete() bool
	// CanRemoveSupportsType checks if a supported content type can be removed
	// from an application.
	CanRemoveSupportsType() bool
	// Delete tries to delete a Info.
	Delete() bool
	// Dup creates a duplicate of a Info.
	Dup() AppInfor
	// Equal checks if two Infos are equal.
	Equal(appinfo2 AppInfor) bool
	// Commandline gets the commandline with which the application will be
	// started.
	Commandline() string
	// Description gets a human-readable description of an installed
	// application.
	Description() string
	// DisplayName gets the display name of the application.
	DisplayName() string
	// Executable gets the executable's name for the installed application.
	Executable() string
	// Icon gets the icon for the application.
	Icon() Iconer
	// ID gets the ID of an application.
	ID() string
	// Name gets the installed name of the application.
	Name() string
	// SupportedTypes retrieves the list of content types that app_info claims
	// to support.
	SupportedTypes() []string
	// LaunchUrisFinish finishes a g_app_info_launch_uris_async() operation.
	LaunchUrisFinish(result AsyncResulter) error
	// RemoveSupportsType removes a supported type from an application, if
	// possible.
	RemoveSupportsType(contentType string) error
	// SetAsDefaultForExtension sets the application as the default handler for
	// the given file extension.
	SetAsDefaultForExtension(extension string) error
	// SetAsDefaultForType sets the application as the default handler for a
	// given type.
	SetAsDefaultForType(contentType string) error
	// SetAsLastUsedForType sets the application as the last used application
	// for a given type.
	SetAsLastUsedForType(contentType string) error
	// ShouldShow checks if the application info should be shown in menus that
	// list available applications.
	ShouldShow() bool
	// SupportsFiles checks if the application accepts files as arguments.
	SupportsFiles() bool
	// SupportsUris checks if the application supports reading files and
	// directories from URIs.
	SupportsUris() bool
}

var _ AppInfor = (*AppInfo)(nil)

func wrapAppInfo(obj *externglib.Object) *AppInfo {
	return &AppInfo{
		Object: obj,
	}
}

func marshalAppInfor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppInfo(obj), nil
}

// AddSupportsType adds a content type to the application information to
// indicate the application is capable of opening files with the given content
// type.
func (appinfo *AppInfo) AddSupportsType(contentType string) error {
	var _arg0 *C.GAppInfo // out
	var _arg1 *C.char     // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))

	C.g_app_info_add_supports_type(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// CanDelete obtains the information whether the Info can be deleted. See
// g_app_info_delete().
func (appinfo *AppInfo) CanDelete() bool {
	var _arg0 *C.GAppInfo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_can_delete(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanRemoveSupportsType checks if a supported content type can be removed from
// an application.
func (appinfo *AppInfo) CanRemoveSupportsType() bool {
	var _arg0 *C.GAppInfo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_can_remove_supports_type(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Delete tries to delete a Info.
//
// On some platforms, there may be a difference between user-defined Infos which
// can be deleted, and system-wide ones which cannot. See
// g_app_info_can_delete().
func (appinfo *AppInfo) Delete() bool {
	var _arg0 *C.GAppInfo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_delete(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Dup creates a duplicate of a Info.
func (appinfo *AppInfo) Dup() AppInfor {
	var _arg0 *C.GAppInfo // out
	var _cret *C.GAppInfo // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_dup(_arg0)

	var _appInfo AppInfor // out

	_appInfo = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(AppInfor)

	return _appInfo
}

// Equal checks if two Infos are equal.
//
// Note that the check *may not* compare each individual field, and only does an
// identity check. In case detecting changes in the contents is needed, program
// code must additionally compare relevant fields.
func (appinfo1 *AppInfo) Equal(appinfo2 AppInfor) bool {
	var _arg0 *C.GAppInfo // out
	var _arg1 *C.GAppInfo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo1.Native()))
	_arg1 = (*C.GAppInfo)(unsafe.Pointer((appinfo2).(gextras.Nativer).Native()))

	_cret = C.g_app_info_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Commandline gets the commandline with which the application will be started.
func (appinfo *AppInfo) Commandline() string {
	var _arg0 *C.GAppInfo // out
	var _cret *C.char     // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_get_commandline(_arg0)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// Description gets a human-readable description of an installed application.
func (appinfo *AppInfo) Description() string {
	var _arg0 *C.GAppInfo // out
	var _cret *C.char     // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// DisplayName gets the display name of the application. The display name is
// often more descriptive to the user than the name itself.
func (appinfo *AppInfo) DisplayName() string {
	var _arg0 *C.GAppInfo // out
	var _cret *C.char     // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_get_display_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Executable gets the executable's name for the installed application.
func (appinfo *AppInfo) Executable() string {
	var _arg0 *C.GAppInfo // out
	var _cret *C.char     // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_get_executable(_arg0)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// Icon gets the icon for the application.
func (appinfo *AppInfo) Icon() Iconer {
	var _arg0 *C.GAppInfo // out
	var _cret *C.GIcon    // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_get_icon(_arg0)

	var _icon Iconer // out

	_icon = (*gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Iconer)

	return _icon
}

// ID gets the ID of an application. An id is a string that identifies the
// application. The exact format of the id is platform dependent. For instance,
// on Unix this is the desktop file id from the xdg menu specification.
//
// Note that the returned ID may be NULL, depending on how the appinfo has been
// constructed.
func (appinfo *AppInfo) ID() string {
	var _arg0 *C.GAppInfo // out
	var _cret *C.char     // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_get_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Name gets the installed name of the application.
func (appinfo *AppInfo) Name() string {
	var _arg0 *C.GAppInfo // out
	var _cret *C.char     // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SupportedTypes retrieves the list of content types that app_info claims to
// support. If this information is not provided by the environment, this
// function will return NULL. This function does not take in consideration
// associations added with g_app_info_add_supports_type(), but only those
// exported directly by the application.
func (appinfo *AppInfo) SupportedTypes() []string {
	var _arg0 *C.GAppInfo // out
	var _cret **C.char

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_get_supported_types(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// LaunchUrisFinish finishes a g_app_info_launch_uris_async() operation.
func (appinfo *AppInfo) LaunchUrisFinish(result AsyncResulter) error {
	var _arg0 *C.GAppInfo     // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_app_info_launch_uris_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// RemoveSupportsType removes a supported type from an application, if possible.
func (appinfo *AppInfo) RemoveSupportsType(contentType string) error {
	var _arg0 *C.GAppInfo // out
	var _arg1 *C.char     // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))

	C.g_app_info_remove_supports_type(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetAsDefaultForExtension sets the application as the default handler for the
// given file extension.
func (appinfo *AppInfo) SetAsDefaultForExtension(extension string) error {
	var _arg0 *C.GAppInfo // out
	var _arg1 *C.char     // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(extension)))

	C.g_app_info_set_as_default_for_extension(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetAsDefaultForType sets the application as the default handler for a given
// type.
func (appinfo *AppInfo) SetAsDefaultForType(contentType string) error {
	var _arg0 *C.GAppInfo // out
	var _arg1 *C.char     // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))

	C.g_app_info_set_as_default_for_type(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetAsLastUsedForType sets the application as the last used application for a
// given type. This will make the application appear as first in the list
// returned by g_app_info_get_recommended_for_type(), regardless of the default
// application for that content type.
func (appinfo *AppInfo) SetAsLastUsedForType(contentType string) error {
	var _arg0 *C.GAppInfo // out
	var _arg1 *C.char     // out
	var _cerr *C.GError   // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))

	C.g_app_info_set_as_last_used_for_type(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ShouldShow checks if the application info should be shown in menus that list
// available applications.
func (appinfo *AppInfo) ShouldShow() bool {
	var _arg0 *C.GAppInfo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_should_show(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsFiles checks if the application accepts files as arguments.
func (appinfo *AppInfo) SupportsFiles() bool {
	var _arg0 *C.GAppInfo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_supports_files(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsUris checks if the application supports reading files and directories
// from URIs.
func (appinfo *AppInfo) SupportsUris() bool {
	var _arg0 *C.GAppInfo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GAppInfo)(unsafe.Pointer(appinfo.Native()))

	_cret = C.g_app_info_supports_uris(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AppInfoCreateFromCommandline creates a new Info from the given information.
//
// Note that for commandline, the quoting rules of the Exec key of the
// freedesktop.org Desktop Entry Specification
// (http://freedesktop.org/Standards/desktop-entry-spec) are applied. For
// example, if the commandline contains percent-encoded URIs, the
// percent-character must be doubled in order to prevent it from being swallowed
// by Exec key unquoting. See the specification for exact quoting rules.
func AppInfoCreateFromCommandline(commandline string, applicationName string, flags AppInfoCreateFlags) (AppInfor, error) {
	var _arg1 *C.char               // out
	var _arg2 *C.char               // out
	var _arg3 C.GAppInfoCreateFlags // out
	var _cret *C.GAppInfo           // in
	var _cerr *C.GError             // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(commandline)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(applicationName)))
	_arg3 = C.GAppInfoCreateFlags(flags)

	_cret = C.g_app_info_create_from_commandline(_arg1, _arg2, _arg3, &_cerr)

	var _appInfo AppInfor // out
	var _goerr error      // out

	_appInfo = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(AppInfor)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _appInfo, _goerr
}

// AppInfoAll gets a list of all of the applications currently registered on
// this system.
//
// For desktop files, this includes applications that have NoDisplay=true set or
// are excluded from display by means of OnlyShowIn or NotShowIn. See
// g_app_info_should_show(). The returned list does not include applications
// which have the Hidden key set.
func AppInfoAll() *externglib.List {
	var _cret *C.GList // in

	_cret = C.g_app_info_get_all()

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.GAppInfo)(_p)
		var dst AppInfor // out
		dst = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(AppInfor)
		return dst
	})
	_list.AttachFinalizer(func(v uintptr) {
		C.g_object_unref(C.gpointer(uintptr(unsafe.Pointer(v))))
	})

	return _list
}

// AppInfoAllForType gets a list of all Infos for a given content type,
// including the recommended and fallback Infos. See
// g_app_info_get_recommended_for_type() and g_app_info_get_fallback_for_type().
func AppInfoAllForType(contentType string) *externglib.List {
	var _arg1 *C.char  // out
	var _cret *C.GList // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))

	_cret = C.g_app_info_get_all_for_type(_arg1)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.GAppInfo)(_p)
		var dst AppInfor // out
		dst = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(AppInfor)
		return dst
	})
	_list.AttachFinalizer(func(v uintptr) {
		C.g_object_unref(C.gpointer(uintptr(unsafe.Pointer(v))))
	})

	return _list
}

// AppInfoDefaultForType gets the default Info for a given content type.
func AppInfoDefaultForType(contentType string, mustSupportUris bool) AppInfor {
	var _arg1 *C.char     // out
	var _arg2 C.gboolean  // out
	var _cret *C.GAppInfo // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))
	if mustSupportUris {
		_arg2 = C.TRUE
	}

	_cret = C.g_app_info_get_default_for_type(_arg1, _arg2)

	var _appInfo AppInfor // out

	_appInfo = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(AppInfor)

	return _appInfo
}

// AppInfoDefaultForURIScheme gets the default application for handling URIs
// with the given URI scheme. A URI scheme is the initial part of the URI, up to
// but not including the ':', e.g. "http", "ftp" or "sip".
func AppInfoDefaultForURIScheme(uriScheme string) AppInfor {
	var _arg1 *C.char     // out
	var _cret *C.GAppInfo // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uriScheme)))

	_cret = C.g_app_info_get_default_for_uri_scheme(_arg1)

	var _appInfo AppInfor // out

	_appInfo = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(AppInfor)

	return _appInfo
}

// AppInfoFallbackForType gets a list of fallback Infos for a given content
// type, i.e. those applications which claim to support the given content type
// by MIME type subclassing and not directly.
func AppInfoFallbackForType(contentType string) *externglib.List {
	var _arg1 *C.gchar // out
	var _cret *C.GList // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))

	_cret = C.g_app_info_get_fallback_for_type(_arg1)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.GAppInfo)(_p)
		var dst AppInfor // out
		dst = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(AppInfor)
		return dst
	})
	_list.AttachFinalizer(func(v uintptr) {
		C.g_object_unref(C.gpointer(uintptr(unsafe.Pointer(v))))
	})

	return _list
}

// AppInfoRecommendedForType gets a list of recommended Infos for a given
// content type, i.e. those applications which claim to support the given
// content type exactly, and not by MIME type subclassing. Note that the first
// application of the list is the last used one, i.e. the last one for which
// g_app_info_set_as_last_used_for_type() has been called.
func AppInfoRecommendedForType(contentType string) *externglib.List {
	var _arg1 *C.gchar // out
	var _cret *C.GList // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))

	_cret = C.g_app_info_get_recommended_for_type(_arg1)

	var _list *externglib.List // out

	_list = externglib.WrapList(uintptr(unsafe.Pointer(_cret)))
	_list.DataWrapper(func(_p unsafe.Pointer) interface{} {
		src := (*C.GAppInfo)(_p)
		var dst AppInfor // out
		dst = (*gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(src)))).(AppInfor)
		return dst
	})
	_list.AttachFinalizer(func(v uintptr) {
		C.g_object_unref(C.gpointer(uintptr(unsafe.Pointer(v))))
	})

	return _list
}

// AppInfoLaunchDefaultForURI: utility function that launches the default
// application registered to handle the specified uri. Synchronous I/O is done
// on the uri to detect the type of the file if required.
//
// The D-Bus–activated applications don't have to be started if your application
// terminates too soon after this function. To prevent this, use
// g_app_info_launch_default_for_uri_async() instead.
func AppInfoLaunchDefaultForURI(uri string, context *AppLaunchContext) error {
	var _arg1 *C.char              // out
	var _arg2 *C.GAppLaunchContext // out
	var _cerr *C.GError            // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))

	C.g_app_info_launch_default_for_uri(_arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// AppInfoLaunchDefaultForURIAsync: async version of
// g_app_info_launch_default_for_uri().
//
// This version is useful if you are interested in receiving error information
// in the case where the application is sandboxed and the portal may present an
// application chooser dialog to the user.
//
// This is also useful if you want to be sure that the D-Bus–activated
// applications are really started before termination and if you are interested
// in receiving error information from their activation.
func AppInfoLaunchDefaultForURIAsync(ctx context.Context, uri string, context *AppLaunchContext, callback AsyncReadyCallback) {
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.char               // out
	var _arg2 *C.GAppLaunchContext  // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg5 = C.gpointer(gbox.AssignOnce(callback))

	C.g_app_info_launch_default_for_uri_async(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// AppInfoLaunchDefaultForURIFinish finishes an asynchronous
// launch-default-for-uri operation.
func AppInfoLaunchDefaultForURIFinish(result AsyncResulter) error {
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((result).(gextras.Nativer).Native()))

	C.g_app_info_launch_default_for_uri_finish(_arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// AppInfoResetTypeAssociations removes all changes to the type associations
// done by g_app_info_set_as_default_for_type(),
// g_app_info_set_as_default_for_extension(), g_app_info_add_supports_type() or
// g_app_info_remove_supports_type().
func AppInfoResetTypeAssociations(contentType string) {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))

	C.g_app_info_reset_type_associations(_arg1)
}

// AppInfoMonitor gets the InfoMonitor for the current thread-default main
// context.
//
// The InfoMonitor will emit a "changed" signal in the thread-default main
// context whenever the list of installed applications (as reported by
// g_app_info_get_all()) may have changed.
//
// You must only call g_object_unref() on the return value from under the same
// main context as you created it.
func AppInfoMonitor() *AppInfoMonitor {
	var _cret *C.GAppInfoMonitor // in

	_cret = C.g_app_info_monitor_get()

	var _appInfoMonitor *AppInfoMonitor // out

	_appInfoMonitor = wrapAppInfoMonitor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _appInfoMonitor
}

// AppLaunchContextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AppLaunchContextOverrider interface {
	// LaunchFailed: called when an application has failed to launch, so that it
	// can cancel the application startup notification started in
	// g_app_launch_context_get_startup_notify_id().
	LaunchFailed(startupNotifyId string)
	Launched(info AppInfor, platformData *glib.Variant)
}

// AppLaunchContext: integrating the launch with the launching application. This
// is used to handle for instance startup notification and launching the new
// application on the same screen as the launching window.
type AppLaunchContext struct {
	*externglib.Object
}

var _ gextras.Nativer = (*AppLaunchContext)(nil)

func wrapAppLaunchContext(obj *externglib.Object) *AppLaunchContext {
	return &AppLaunchContext{
		Object: obj,
	}
}

func marshalAppLaunchContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppLaunchContext(obj), nil
}

// NewAppLaunchContext creates a new application launch context. This is not
// normally used, instead you instantiate a subclass of this, such as
// AppLaunchContext.
func NewAppLaunchContext() *AppLaunchContext {
	var _cret *C.GAppLaunchContext // in

	_cret = C.g_app_launch_context_new()

	var _appLaunchContext *AppLaunchContext // out

	_appLaunchContext = wrapAppLaunchContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _appLaunchContext
}

// Environment gets the complete environment variable list to be passed to the
// child process when context is used to launch an application. This is a
// NULL-terminated array of strings, where each string has the form KEY=VALUE.
func (context *AppLaunchContext) Environment() []string {
	var _arg0 *C.GAppLaunchContext // out
	var _cret **C.char

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))

	_cret = C.g_app_launch_context_get_environment(_arg0)

	var _filenames []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _filenames
}

// LaunchFailed: called when an application has failed to launch, so that it can
// cancel the application startup notification started in
// g_app_launch_context_get_startup_notify_id().
func (context *AppLaunchContext) LaunchFailed(startupNotifyId string) {
	var _arg0 *C.GAppLaunchContext // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(startupNotifyId)))

	C.g_app_launch_context_launch_failed(_arg0, _arg1)
}

// Setenv arranges for variable to be set to value in the child's environment
// when context is used to launch an application.
func (context *AppLaunchContext) Setenv(variable string, value string) {
	var _arg0 *C.GAppLaunchContext // out
	var _arg1 *C.char              // out
	var _arg2 *C.char              // out

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(variable)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))

	C.g_app_launch_context_setenv(_arg0, _arg1, _arg2)
}

// Unsetenv arranges for variable to be unset in the child's environment when
// context is used to launch an application.
func (context *AppLaunchContext) Unsetenv(variable string) {
	var _arg0 *C.GAppLaunchContext // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(variable)))

	C.g_app_launch_context_unsetenv(_arg0, _arg1)
}
