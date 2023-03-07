// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <pango/pango.h>
import "C"

// ScriptForUnichar looks up the script for a particular character.
//
// The script of a character is defined by Unicode Standard Annex \#24. No check
// is made for ch being a valid Unicode character; if you pass in invalid
// character, the result is undefined.
//
// Note that while the return type of this function is declared as PangoScript,
// as of Pango 1.18, this function simply returns the return value of
// g_unichar_get_script(). Callers must be prepared to handle unknown values.
//
// Deprecated: Use g_unichar_get_script().
//
// The function takes the following parameters:
//
//    - ch: unicode character.
//
// The function returns the following values:
//
//    - script: PangoScript for the character.
//
func ScriptForUnichar(ch uint32) Script {
	var _arg1 C.gunichar    // out
	var _cret C.PangoScript // in

	_arg1 = C.gunichar(ch)

	_cret = C.pango_script_for_unichar(_arg1)
	runtime.KeepAlive(ch)

	var _script Script // out

	_script = Script(_cret)

	return _script
}

// ScriptGetSampleLanguage finds a language tag that is reasonably
// representative of script.
//
// The language will usually be the most widely spoken or used language written
// in that script: for instance, the sample language for PANGO_SCRIPT_CYRILLIC
// is ru (Russian), the sample language for PANGO_SCRIPT_ARABIC is ar.
//
// For some scripts, no sample language will be returned because there is no
// language that is sufficiently representative. The best example of this is
// PANGO_SCRIPT_HAN, where various different variants of written Chinese,
// Japanese, and Korean all use significantly different sets of Han characters
// and forms of shared characters. No sample language can be provided for many
// historical scripts as well.
//
// As of 1.18, this function checks the environment variables PANGO_LANGUAGE and
// LANGUAGE (checked in that order) first. If one of them is set, it is parsed
// as a list of language tags separated by colons or other separators. This
// function will return the first language in the parsed list that Pango
// believes may use script for writing. This last predicate is tested using
// pango.Language.IncludesScript(). This can be used to control Pango's font
// selection for non-primary languages. For example, a PANGO_LANGUAGE enviroment
// variable set to "en:fa" makes Pango choose fonts suitable for Persian (fa)
// instead of Arabic (ar) when a segment of Arabic text is found in an otherwise
// non-Arabic text. The same trick can be used to choose a default language for
// PANGO_SCRIPT_HAN when setting context language is not feasible.
//
// The function takes the following parameters:
//
//    - script: PangoScript.
//
// The function returns the following values:
//
//    - language (optional): PangoLanguage that is representative of the script,
//      or NULL if no such language exists.
//
func ScriptGetSampleLanguage(script Script) *Language {
	var _arg1 C.PangoScript    // out
	var _cret *C.PangoLanguage // in

	_arg1 = C.PangoScript(script)

	_cret = C.pango_script_get_sample_language(_arg1)
	runtime.KeepAlive(script)

	var _language *Language // out

	if _cret != nil {
		_language = (*Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_language)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _language
}