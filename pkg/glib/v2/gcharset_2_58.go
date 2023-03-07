// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// GetLanguageNamesWithCategory computes a list of applicable locale names with
// a locale category name, which can be used to construct the fallback
// locale-dependent filenames or search paths. The returned list is sorted from
// most desirable to least desirable and always contains the default locale "C".
//
// This function consults the environment variables LANGUAGE, LC_ALL,
// category_name, and LANG to find the list of locales specified by the user.
//
// g_get_language_names() returns
// g_get_language_names_with_category("LC_MESSAGES").
//
// The function takes the following parameters:
//
//    - categoryName: locale category name.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of strings owned by the thread
//      g_get_language_names_with_category was called from. It must not be
//      modified or freed. It must be copied if planned to be used in another
//      thread.
//
func GetLanguageNamesWithCategory(categoryName string) []string {
	var _arg1 *C.gchar  // out
	var _cret **C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(categoryName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_get_language_names_with_category(_arg1)
	runtime.KeepAlive(categoryName)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}