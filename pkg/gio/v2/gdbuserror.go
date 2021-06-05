// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
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
import "C"

// DBusErrorEncodeGerror creates a D-Bus error name to use for @error. If @error
// matches a registered error (cf. g_dbus_error_register_error()), the
// corresponding D-Bus error name will be returned.
//
// Otherwise the a name of the form
// `org.gtk.GDBus.UnmappedGError.Quark._ESCAPED_QUARK_NAME.Code_ERROR_CODE` will
// be used. This allows other GDBus applications to map the error on the wire
// back to a #GError using g_dbus_error_new_for_dbus_error().
//
// This function is typically only used in object mappings to put a #GError on
// the wire. Regular applications should not use it.
func DBusErrorEncodeGerror(error *glib.Error) string {
	var arg1 *C.GError

	arg1 = (*C.GError)(unsafe.Pointer(error.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_dbus_error_encode_gerror(error)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

// DBusErrorGetRemoteError gets the D-Bus error name used for @error, if any.
//
// This function is guaranteed to return a D-Bus error name for all #GErrors
// returned from functions handling remote method calls (e.g.
// g_dbus_connection_call_finish()) unless g_dbus_error_strip_remote_error() has
// been used on @error.
func DBusErrorGetRemoteError(error *glib.Error) string {
	var arg1 *C.GError

	arg1 = (*C.GError)(unsafe.Pointer(error.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_dbus_error_get_remote_error(error)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

// DBusErrorIsRemoteError checks if @error represents an error received via
// D-Bus from a remote peer. If so, use g_dbus_error_get_remote_error() to get
// the name of the error.
func DBusErrorIsRemoteError(error *glib.Error) bool {
	var arg1 *C.GError

	arg1 = (*C.GError)(unsafe.Pointer(error.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_dbus_error_is_remote_error(error)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// DBusErrorNewForDBusError creates a #GError based on the contents of
// @dbus_error_name and @dbus_error_message.
//
// Errors registered with g_dbus_error_register_error() will be looked up using
// @dbus_error_name and if a match is found, the error domain and code is used.
// Applications can use g_dbus_error_get_remote_error() to recover
// @dbus_error_name.
//
// If a match against a registered error is not found and the D-Bus error name
// is in a form as returned by g_dbus_error_encode_gerror() the error domain and
// code encoded in the name is used to create the #GError. Also,
// @dbus_error_name is added to the error message such that it can be recovered
// with g_dbus_error_get_remote_error().
//
// Otherwise, a #GError with the error code G_IO_ERROR_DBUS_ERROR in the
// IO_ERROR error domain is returned. Also, @dbus_error_name is added to the
// error message such that it can be recovered with
// g_dbus_error_get_remote_error().
//
// In all three cases, @dbus_error_name can always be recovered from the
// returned #GError using the g_dbus_error_get_remote_error() function (unless
// g_dbus_error_strip_remote_error() hasn't been used on the returned error).
//
// This function is typically only used in object mappings to prepare #GError
// instances for applications. Regular applications should not use it.
func DBusErrorNewForDBusError(dbusErrorName string, dbusErrorMessage string) *glib.Error {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(dbusErrorName))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(dbusErrorMessage))
	defer C.free(unsafe.Pointer(arg2))

	var cret *C.GError
	var goret1 *glib.Error

	cret = C.g_dbus_error_new_for_dbus_error(dbusErrorName, dbusErrorMessage)

	goret1 = glib.WrapError(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *glib.Error) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// DBusErrorRegisterErrorDomain: helper function for associating a #GError error
// domain with D-Bus error names.
func DBusErrorRegisterErrorDomain(errorDomainQuarkName string, quarkVolatile uint, entries []DBusErrorEntry) {

	C.g_dbus_error_register_error_domain(errorDomainQuarkName, quarkVolatile, entries, numEntries)
}

// DBusErrorStripRemoteError looks for extra information in the error message
// used to recover the D-Bus error name and strips it if found. If stripped, the
// message field in @error will correspond exactly to what was received on the
// wire.
//
// This is typically used when presenting errors to the end user.
func DBusErrorStripRemoteError(error *glib.Error) bool {
	var arg1 *C.GError

	arg1 = (*C.GError)(unsafe.Pointer(error.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_dbus_error_strip_remote_error(error)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// DBusErrorEntry: struct used in g_dbus_error_register_error_domain().
type DBusErrorEntry struct {
	native C.GDBusErrorEntry
}

// WrapDBusErrorEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusErrorEntry(ptr unsafe.Pointer) *DBusErrorEntry {
	if ptr == nil {
		return nil
	}

	return (*DBusErrorEntry)(ptr)
}

func marshalDBusErrorEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusErrorEntry(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusErrorEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// ErrorCode gets the field inside the struct.
func (d *DBusErrorEntry) ErrorCode() int {
	v = C.gint(d.native.error_code)
}

// DBusErrorName gets the field inside the struct.
func (d *DBusErrorEntry) DBusErrorName() string {
	v = C.GoString(d.native.dbus_error_name)
}