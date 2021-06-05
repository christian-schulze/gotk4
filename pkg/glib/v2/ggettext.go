// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// Dcgettext: this is a variant of g_dgettext() that allows specifying a locale
// category instead of always using `LC_MESSAGES`. See g_dgettext() for more
// information about how this functions differs from calling dcgettext()
// directly.
func Dcgettext(domain string, msgid string, category int) string {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 C.gint

	arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(msgid))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gint(category)

	var cret *C.gchar
	var goret1 string

	cret = C.g_dcgettext(domain, msgid, category)

	goret1 = C.GoString(cret)

	return goret1
}

// Dgettext: this function is a wrapper of dgettext() which does not translate
// the message if the default domain as set with textdomain() has no
// translations for the current locale.
//
// The advantage of using this function over dgettext() proper is that libraries
// using this function (like GTK+) will not use translations if the application
// using the library does not have translations for the current locale. This
// results in a consistent English-only interface instead of one having partial
// translations. For this feature to work, the call to textdomain() and
// setlocale() should precede any g_dgettext() invocations. For GTK+, it means
// calling textdomain() before gtk_init or its variants.
//
// This function disables translations if and only if upon its first call all
// the following conditions hold:
//
// - @domain is not nil
//
// - textdomain() has been called to set a default text domain
//
// - there is no translations available for the default text domain and the
// current locale
//
// - current locale is not "C" or any English locales (those starting with
// "en_")
//
// Note that this behavior may not be desired for example if an application has
// its untranslated messages in a language other than English. In those cases
// the application should call textdomain() after initializing GTK+.
//
// Applications should normally not use this function directly, but use the _()
// macro for translations.
func Dgettext(domain string, msgid string) string {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(msgid))
	defer C.free(unsafe.Pointer(arg2))

	var cret *C.gchar
	var goret1 string

	cret = C.g_dgettext(domain, msgid)

	goret1 = C.GoString(cret)

	return goret1
}

// Dngettext: this function is a wrapper of dngettext() which does not translate
// the message if the default domain as set with textdomain() has no
// translations for the current locale.
//
// See g_dgettext() for details of how this differs from dngettext() proper.
func Dngettext(domain string, msgid string, msgidPlural string, n uint32) string {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gulong

	arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(msgid))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(msgidPlural))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gulong(n)

	var cret *C.gchar
	var goret1 string

	cret = C.g_dngettext(domain, msgid, msgidPlural, n)

	goret1 = C.GoString(cret)

	return goret1
}

// Dpgettext: this function is a variant of g_dgettext() which supports a
// disambiguating message context. GNU gettext uses the '\004' character to
// separate the message context and message id in @msgctxtid. If 0 is passed as
// @msgidoffset, this function will fall back to trying to use the deprecated
// convention of using "|" as a separation character.
//
// This uses g_dgettext() internally. See that functions for differences with
// dgettext() proper.
//
// Applications should normally not use this function directly, but use the C_()
// macro for translations with context.
func Dpgettext(domain string, msgctxtid string, msgidoffset uint) string {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 C.gsize

	arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(msgctxtid))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gsize(msgidoffset)

	var cret *C.gchar
	var goret1 string

	cret = C.g_dpgettext(domain, msgctxtid, msgidoffset)

	goret1 = C.GoString(cret)

	return goret1
}

// Dpgettext2: this function is a variant of g_dgettext() which supports a
// disambiguating message context. GNU gettext uses the '\004' character to
// separate the message context and message id in @msgctxtid.
//
// This uses g_dgettext() internally. See that functions for differences with
// dgettext() proper.
//
// This function differs from C_() in that it is not a macro and thus you may
// use non-string-literals as context and msgid arguments.
func Dpgettext2(domain string, context string, msgid string) string {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(context))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(msgid))
	defer C.free(unsafe.Pointer(arg3))

	var cret *C.gchar
	var goret1 string

	cret = C.g_dpgettext2(domain, context, msgid)

	goret1 = C.GoString(cret)

	return goret1
}

// StripContext: an auxiliary function for gettext() support (see Q_()).
func StripContext(msgid string, msgval string) string {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(msgid))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(msgval))
	defer C.free(unsafe.Pointer(arg2))

	var cret *C.gchar
	var goret1 string

	cret = C.g_strip_context(msgid, msgval)

	goret1 = C.GoString(cret)

	return goret1
}