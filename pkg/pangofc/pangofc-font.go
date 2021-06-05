// Code generated by girgen. DO NOT EDIT.

package pangofc

import (
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangofc pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <pango/pangofc-fontmap.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fc_font_get_type()), F: marshalFont},
	})
}

// Font: `PangoFcFont` is a base class for font implementations using the
// Fontconfig and FreeType libraries.
//
// It is used in onjunction with [class@PangoFc.FontMap]. When deriving from
// this class, you need to implement all of its virtual functions other than
// shutdown() along with the get_glyph_extents() virtual function from
// `PangoFont`.
type Font interface {
	pango.Font

	// Glyph gets the glyph index for a given Unicode character for @font.
	//
	// If you only want to determine whether the font has the glyph, use
	// [method@PangoFc.Font.has_char].
	Glyph(wc uint32) uint
	// Languages returns the languages that are supported by @font.
	//
	// This corresponds to the FC_LANG member of the FcPattern.
	//
	// The returned array is only valid as long as the font and its fontmap are
	// valid.
	Languages() **pango.Language
	// HasChar determines whether @font has a glyph for the codepoint @wc.
	HasChar(wc uint32) bool
	// KernGlyphs: this function used to adjust each adjacent pair of glyphs in
	// @glyphs according to kerning information in @font.
	//
	// Since 1.44, it does nothing.
	KernGlyphs(glyphs *pango.GlyphString)
	// UnlockFace releases a font previously obtained with
	// [method@PangoFc.Font.lock_face].
	UnlockFace()
}

// font implements the Font interface.
type font struct {
	pango.Font
}

var _ Font = (*font)(nil)

// WrapFont wraps a GObject to the right type. It is
// primarily used internally.
func WrapFont(obj *externglib.Object) Font {
	return Font{
		pango.Font: pango.WrapFont(obj),
	}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFont(obj), nil
}

// Glyph gets the glyph index for a given Unicode character for @font.
//
// If you only want to determine whether the font has the glyph, use
// [method@PangoFc.Font.has_char].
func (f font) Glyph(wc uint32) uint {
	var arg0 *C.PangoFcFont
	var arg1 C.gunichar

	arg0 = (*C.PangoFcFont)(unsafe.Pointer(f.Native()))
	arg1 = C.gunichar(wc)

	var cret C.guint
	var goret1 uint

	cret = C.pango_fc_font_get_glyph(arg0, wc)

	goret1 = C.guint(cret)

	return goret1
}

// Languages returns the languages that are supported by @font.
//
// This corresponds to the FC_LANG member of the FcPattern.
//
// The returned array is only valid as long as the font and its fontmap are
// valid.
func (f font) Languages() **pango.Language {
	var arg0 *C.PangoFcFont

	arg0 = (*C.PangoFcFont)(unsafe.Pointer(f.Native()))

	var cret **C.PangoLanguage
	var goret1 **pango.Language

	cret = C.pango_fc_font_get_languages(arg0)

	goret1 = pango.WrapLanguage(unsafe.Pointer(cret))

	return goret1
}

// HasChar determines whether @font has a glyph for the codepoint @wc.
func (f font) HasChar(wc uint32) bool {
	var arg0 *C.PangoFcFont
	var arg1 C.gunichar

	arg0 = (*C.PangoFcFont)(unsafe.Pointer(f.Native()))
	arg1 = C.gunichar(wc)

	var cret C.gboolean
	var goret1 bool

	cret = C.pango_fc_font_has_char(arg0, wc)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// KernGlyphs: this function used to adjust each adjacent pair of glyphs in
// @glyphs according to kerning information in @font.
//
// Since 1.44, it does nothing.
func (f font) KernGlyphs(glyphs *pango.GlyphString) {
	var arg0 *C.PangoFcFont
	var arg1 *C.PangoGlyphString

	arg0 = (*C.PangoFcFont)(unsafe.Pointer(f.Native()))
	arg1 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))

	C.pango_fc_font_kern_glyphs(arg0, glyphs)
}

// UnlockFace releases a font previously obtained with
// [method@PangoFc.Font.lock_face].
func (f font) UnlockFace() {
	var arg0 *C.PangoFcFont

	arg0 = (*C.PangoFcFont)(unsafe.Pointer(f.Native()))

	C.pango_fc_font_unlock_face(arg0)
}