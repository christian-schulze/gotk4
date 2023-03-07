// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// Dpgettext: this function is a variant of g_dgettext() which supports a
// disambiguating message context. GNU gettext uses the '\004' character to
// separate the message context and message id in msgctxtid. If 0 is passed as
// msgidoffset, this function will fall back to trying to use the deprecated
// convention of using "|" as a separation character.
//
// This uses g_dgettext() internally. See that functions for differences with
// dgettext() proper.
//
// Applications should normally not use this function directly, but use the C_()
// macro for translations with context.
//
// The function takes the following parameters:
//
//    - domain (optional): translation domain to use, or NULL to use the domain
//      set with textdomain().
//    - msgctxtid: combined message context and message id, separated by a \004
//      character.
//    - msgidoffset: offset of the message id in msgctxid.
//
// The function returns the following values:
//
//    - utf8: translated string.
//
func Dpgettext(domain, msgctxtid string, msgidoffset uint) string {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _arg3 C.gsize  // out
	var _cret *C.gchar // in

	if domain != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(msgctxtid)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gsize(msgidoffset)

	_cret = C.g_dpgettext(_arg1, _arg2, _arg3)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(msgctxtid)
	runtime.KeepAlive(msgidoffset)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}