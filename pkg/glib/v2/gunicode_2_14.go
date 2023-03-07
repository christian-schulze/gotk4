// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
)

// #include <stdlib.h>
// #include <glib.h>
import "C"

// UnicharCombiningClass determines the canonical combining class of a Unicode
// character.
//
// The function takes the following parameters:
//
//    - uc: unicode character.
//
// The function returns the following values:
//
//    - gint: combining class of the character.
//
func UnicharCombiningClass(uc uint32) int {
	var _arg1 C.gunichar // out
	var _cret C.gint     // in

	_arg1 = C.gunichar(uc)

	_cret = C.g_unichar_combining_class(_arg1)
	runtime.KeepAlive(uc)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UnicharGetScript looks up the Script for a particular character (as defined
// by Unicode Standard Annex \#24). No check is made for ch being a valid
// Unicode character; if you pass in invalid character, the result is undefined.
//
// This function is equivalent to pango_script_for_unichar() and the two are
// interchangeable.
//
// The function takes the following parameters:
//
//    - ch: unicode character.
//
// The function returns the following values:
//
//    - unicodeScript for the character.
//
func UnicharGetScript(ch uint32) UnicodeScript {
	var _arg1 C.gunichar       // out
	var _cret C.GUnicodeScript // in

	_arg1 = C.gunichar(ch)

	_cret = C.g_unichar_get_script(_arg1)
	runtime.KeepAlive(ch)

	var _unicodeScript UnicodeScript // out

	_unicodeScript = UnicodeScript(_cret)

	return _unicodeScript
}

// UnicharIsmark determines whether a character is a mark (non-spacing mark,
// combining mark, or enclosing mark in Unicode speak). Given some UTF-8 text,
// obtain a character value with g_utf8_get_char().
//
// Note: in most cases where isalpha characters are allowed, ismark characters
// should be allowed to as they are essential for writing most European
// languages as well as many non-Latin scripts.
//
// The function takes the following parameters:
//
//    - c: unicode character.
//
// The function returns the following values:
//
//    - ok: TRUE if c is a mark character.
//
func UnicharIsmark(c uint32) bool {
	var _arg1 C.gunichar // out
	var _cret C.gboolean // in

	_arg1 = C.gunichar(c)

	_cret = C.g_unichar_ismark(_arg1)
	runtime.KeepAlive(c)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnicharIszerowidth determines if a given character typically takes zero width
// when rendered. The return value is TRUE for all non-spacing and enclosing
// marks (e.g., combining accents), format characters, zero-width space, but not
// U+00AD SOFT HYPHEN.
//
// A typical use of this function is with one of g_unichar_iswide() or
// g_unichar_iswide_cjk() to determine the number of cells a string occupies
// when displayed on a grid display (terminals). However, note that not all
// terminals support zero-width rendering of zero-width marks.
//
// The function takes the following parameters:
//
//    - c: unicode character.
//
// The function returns the following values:
//
//    - ok: TRUE if the character has zero width.
//
func UnicharIszerowidth(c uint32) bool {
	var _arg1 C.gunichar // out
	var _cret C.gboolean // in

	_arg1 = C.gunichar(c)

	_cret = C.g_unichar_iszerowidth(_arg1)
	runtime.KeepAlive(c)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}