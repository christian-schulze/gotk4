// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_renderer_get_type()), F: marshalRenderer},
	})
}

// Renderer: `PangoRenderer` is a base class for objects that can render text
// provided as `PangoGlyphString` or `PangoLayout`.
//
// By subclassing `PangoRenderer` and overriding operations such as @draw_glyphs
// and @draw_rectangle, renderers for particular font backends and destinations
// can be created.
type Renderer interface {
	gextras.Objector

	// Activate does initial setup before rendering operations on @renderer.
	//
	// [method@Pango.Renderer.deactivate] should be called when done drawing.
	// Calls such as [method@Pango.Renderer.draw_layout] automatically activate
	// the layout before drawing on it. Calls to `pango_renderer_activate()` and
	// `pango_renderer_deactivate()` can be nested and the renderer will only be
	// initialized and deinitialized once.
	Activate()
	// Deactivate cleans up after rendering operations on @renderer.
	//
	// See docs for [method@Pango.Renderer.activate].
	Deactivate()
	// DrawErrorUnderline: draw a squiggly line that approximately covers the
	// given rectangle in the style of an underline used to indicate a spelling
	// error.
	//
	// The width of the underline is rounded to an integer number of up/down
	// segments and the resulting rectangle is centered in the original
	// rectangle.
	//
	// This should be called while @renderer is already active. Use
	// [method@Pango.Renderer.activate] to activate a renderer.
	DrawErrorUnderline(x int, y int, width int, height int)
	// DrawGlyphItem draws the glyphs in @glyph_item with the specified
	// `PangoRenderer`, embedding the text associated with the glyphs in the
	// output if the output format supports it.
	//
	// This is useful for rendering text in PDF.
	//
	// Note that @text is the start of the text for layout, which is then
	// indexed by `glyph_item->item->offset`.
	//
	// If @text is nil, this simply calls [method@Pango.Renderer.draw_glyphs].
	//
	// The default implementation of this method simply falls back to
	// [method@Pango.Renderer.draw_glyphs].
	DrawGlyphItem(text string, glyphItem *GlyphItem, x int, y int)
	// DrawGlyphs draws the glyphs in @glyphs with the specified
	// `PangoRenderer`.
	DrawGlyphs(font Font, glyphs *GlyphString, x int, y int)
	// DrawLayout draws @layout with the specified `PangoRenderer`.
	DrawLayout(layout Layout, x int, y int)
	// DrawLayoutLine draws @line with the specified `PangoRenderer`.
	DrawLayoutLine(line *LayoutLine, x int, y int)
	// DrawRectangle draws an axis-aligned rectangle in user space coordinates
	// with the specified `PangoRenderer`.
	//
	// This should be called while @renderer is already active. Use
	// [method@Pango.Renderer.activate] to activate a renderer.
	DrawRectangle(part RenderPart, x int, y int, width int, height int)
	// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the
	// X axis using the given `PangoRenderer`; coordinates are in device space.
	DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64)
	// Alpha gets the current alpha for the specified part.
	Alpha(part RenderPart) uint16
	// Color gets the current rendering color for the specified part.
	Color(part RenderPart) *Color
	// Layout gets the layout currently being rendered using @renderer.
	//
	// Calling this function only makes sense from inside a subclass's methods,
	// like in its draw_shape vfunc, for example.
	//
	// The returned layout should not be modified while still being rendered.
	Layout() Layout
	// LayoutLine gets the layout line currently being rendered using @renderer.
	//
	// Calling this function only makes sense from inside a subclass's methods,
	// like in its draw_shape vfunc, for example.
	//
	// The returned layout line should not be modified while still being
	// rendered.
	LayoutLine() *LayoutLine
	// Matrix gets the transformation matrix that will be applied when
	// rendering.
	//
	// See [method@Pango.Renderer.set_matrix].
	Matrix() *Matrix
	// PartChanged informs Pango that the way that the rendering is done for
	// @part has changed.
	//
	// This should be called if the rendering changes in a way that would
	// prevent multiple pieces being joined together into one drawing call. For
	// instance, if a subclass of `PangoRenderer` was to add a stipple option
	// for drawing underlines, it needs to call
	//
	// “` pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE); “`
	//
	// When the stipple changes or underlines with different stipples might be
	// joined together. Pango automatically calls this for changes to colors.
	// (See [method@Pango.Renderer.set_color])
	PartChanged(part RenderPart)
	// SetAlpha sets the alpha for part of the rendering.
	//
	// Note that the alpha may only be used if a color is specified for @part as
	// well.
	SetAlpha(part RenderPart, alpha uint16)
	// SetColor sets the color for part of the rendering.
	//
	// Also see [method@Pango.Renderer.set_alpha].
	SetColor(part RenderPart, color *Color)
	// SetMatrix sets the transformation matrix that will be applied when
	// rendering.
	SetMatrix(matrix *Matrix)
}

// renderer implements the Renderer interface.
type renderer struct {
	gextras.Objector
}

var _ Renderer = (*renderer)(nil)

// WrapRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapRenderer(obj *externglib.Object) Renderer {
	return Renderer{
		Objector: obj,
	}
}

func marshalRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRenderer(obj), nil
}

// Activate does initial setup before rendering operations on @renderer.
//
// [method@Pango.Renderer.deactivate] should be called when done drawing.
// Calls such as [method@Pango.Renderer.draw_layout] automatically activate
// the layout before drawing on it. Calls to `pango_renderer_activate()` and
// `pango_renderer_deactivate()` can be nested and the renderer will only be
// initialized and deinitialized once.
func (r renderer) Activate() {
	var arg0 *C.PangoRenderer

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	C.pango_renderer_activate(arg0)
}

// Deactivate cleans up after rendering operations on @renderer.
//
// See docs for [method@Pango.Renderer.activate].
func (r renderer) Deactivate() {
	var arg0 *C.PangoRenderer

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	C.pango_renderer_deactivate(arg0)
}

// DrawErrorUnderline: draw a squiggly line that approximately covers the
// given rectangle in the style of an underline used to indicate a spelling
// error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original
// rectangle.
//
// This should be called while @renderer is already active. Use
// [method@Pango.Renderer.activate] to activate a renderer.
func (r renderer) DrawErrorUnderline(x int, y int, width int, height int) {
	var arg0 *C.PangoRenderer
	var arg1 C.int
	var arg2 C.int
	var arg3 C.int
	var arg4 C.int

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = C.int(x)
	arg2 = C.int(y)
	arg3 = C.int(width)
	arg4 = C.int(height)

	C.pango_renderer_draw_error_underline(arg0, x, y, width, height)
}

// DrawGlyphItem draws the glyphs in @glyph_item with the specified
// `PangoRenderer`, embedding the text associated with the glyphs in the
// output if the output format supports it.
//
// This is useful for rendering text in PDF.
//
// Note that @text is the start of the text for layout, which is then
// indexed by `glyph_item->item->offset`.
//
// If @text is nil, this simply calls [method@Pango.Renderer.draw_glyphs].
//
// The default implementation of this method simply falls back to
// [method@Pango.Renderer.draw_glyphs].
func (r renderer) DrawGlyphItem(text string, glyphItem *GlyphItem, x int, y int) {
	var arg0 *C.PangoRenderer
	var arg1 *C.char
	var arg2 *C.PangoGlyphItem
	var arg3 C.int
	var arg4 C.int

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.PangoGlyphItem)(unsafe.Pointer(glyphItem.Native()))
	arg3 = C.int(x)
	arg4 = C.int(y)

	C.pango_renderer_draw_glyph_item(arg0, text, glyphItem, x, y)
}

// DrawGlyphs draws the glyphs in @glyphs with the specified
// `PangoRenderer`.
func (r renderer) DrawGlyphs(font Font, glyphs *GlyphString, x int, y int) {
	var arg0 *C.PangoRenderer
	var arg1 *C.PangoFont
	var arg2 *C.PangoGlyphString
	var arg3 C.int
	var arg4 C.int

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))
	arg2 = (*C.PangoGlyphString)(unsafe.Pointer(glyphs.Native()))
	arg3 = C.int(x)
	arg4 = C.int(y)

	C.pango_renderer_draw_glyphs(arg0, font, glyphs, x, y)
}

// DrawLayout draws @layout with the specified `PangoRenderer`.
func (r renderer) DrawLayout(layout Layout, x int, y int) {
	var arg0 *C.PangoRenderer
	var arg1 *C.PangoLayout
	var arg2 C.int
	var arg3 C.int

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	arg2 = C.int(x)
	arg3 = C.int(y)

	C.pango_renderer_draw_layout(arg0, layout, x, y)
}

// DrawLayoutLine draws @line with the specified `PangoRenderer`.
func (r renderer) DrawLayoutLine(line *LayoutLine, x int, y int) {
	var arg0 *C.PangoRenderer
	var arg1 *C.PangoLayoutLine
	var arg2 C.int
	var arg3 C.int

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (*C.PangoLayoutLine)(unsafe.Pointer(line.Native()))
	arg2 = C.int(x)
	arg3 = C.int(y)

	C.pango_renderer_draw_layout_line(arg0, line, x, y)
}

// DrawRectangle draws an axis-aligned rectangle in user space coordinates
// with the specified `PangoRenderer`.
//
// This should be called while @renderer is already active. Use
// [method@Pango.Renderer.activate] to activate a renderer.
func (r renderer) DrawRectangle(part RenderPart, x int, y int, width int, height int) {
	var arg0 *C.PangoRenderer
	var arg1 C.PangoRenderPart
	var arg2 C.int
	var arg3 C.int
	var arg4 C.int
	var arg5 C.int

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoRenderPart)(part)
	arg2 = C.int(x)
	arg3 = C.int(y)
	arg4 = C.int(width)
	arg5 = C.int(height)

	C.pango_renderer_draw_rectangle(arg0, part, x, y, width, height)
}

// DrawTrapezoid draws a trapezoid with the parallel sides aligned with the
// X axis using the given `PangoRenderer`; coordinates are in device space.
func (r renderer) DrawTrapezoid(part RenderPart, y1 float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	var arg0 *C.PangoRenderer
	var arg1 C.PangoRenderPart
	var arg2 C.double
	var arg3 C.double
	var arg4 C.double
	var arg5 C.double
	var arg6 C.double
	var arg7 C.double

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoRenderPart)(part)
	arg2 = C.double(y1)
	arg3 = C.double(x11)
	arg4 = C.double(x21)
	arg5 = C.double(y2)
	arg6 = C.double(x12)
	arg7 = C.double(x22)

	C.pango_renderer_draw_trapezoid(arg0, part, y1, x11, x21, y2, x12, x22)
}

// Alpha gets the current alpha for the specified part.
func (r renderer) Alpha(part RenderPart) uint16 {
	var arg0 *C.PangoRenderer
	var arg1 C.PangoRenderPart

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoRenderPart)(part)

	var cret C.guint16
	var goret1 uint16

	cret = C.pango_renderer_get_alpha(arg0, part)

	goret1 = C.guint16(cret)

	return goret1
}

// Color gets the current rendering color for the specified part.
func (r renderer) Color(part RenderPart) *Color {
	var arg0 *C.PangoRenderer
	var arg1 C.PangoRenderPart

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoRenderPart)(part)

	var cret *C.PangoColor
	var goret1 *Color

	cret = C.pango_renderer_get_color(arg0, part)

	goret1 = WrapColor(unsafe.Pointer(cret))

	return goret1
}

// Layout gets the layout currently being rendered using @renderer.
//
// Calling this function only makes sense from inside a subclass's methods,
// like in its draw_shape vfunc, for example.
//
// The returned layout should not be modified while still being rendered.
func (r renderer) Layout() Layout {
	var arg0 *C.PangoRenderer

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	var cret *C.PangoLayout
	var goret1 Layout

	cret = C.pango_renderer_get_layout(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Layout)

	return goret1
}

// LayoutLine gets the layout line currently being rendered using @renderer.
//
// Calling this function only makes sense from inside a subclass's methods,
// like in its draw_shape vfunc, for example.
//
// The returned layout line should not be modified while still being
// rendered.
func (r renderer) LayoutLine() *LayoutLine {
	var arg0 *C.PangoRenderer

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	var cret *C.PangoLayoutLine
	var goret1 *LayoutLine

	cret = C.pango_renderer_get_layout_line(arg0)

	goret1 = WrapLayoutLine(unsafe.Pointer(cret))

	return goret1
}

// Matrix gets the transformation matrix that will be applied when
// rendering.
//
// See [method@Pango.Renderer.set_matrix].
func (r renderer) Matrix() *Matrix {
	var arg0 *C.PangoRenderer

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))

	var cret *C.PangoMatrix
	var goret1 *Matrix

	cret = C.pango_renderer_get_matrix(arg0)

	goret1 = WrapMatrix(unsafe.Pointer(cret))

	return goret1
}

// PartChanged informs Pango that the way that the rendering is done for
// @part has changed.
//
// This should be called if the rendering changes in a way that would
// prevent multiple pieces being joined together into one drawing call. For
// instance, if a subclass of `PangoRenderer` was to add a stipple option
// for drawing underlines, it needs to call
//
// “` pango_renderer_part_changed (render, PANGO_RENDER_PART_UNDERLINE); “`
//
// When the stipple changes or underlines with different stipples might be
// joined together. Pango automatically calls this for changes to colors.
// (See [method@Pango.Renderer.set_color])
func (r renderer) PartChanged(part RenderPart) {
	var arg0 *C.PangoRenderer
	var arg1 C.PangoRenderPart

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoRenderPart)(part)

	C.pango_renderer_part_changed(arg0, part)
}

// SetAlpha sets the alpha for part of the rendering.
//
// Note that the alpha may only be used if a color is specified for @part as
// well.
func (r renderer) SetAlpha(part RenderPart, alpha uint16) {
	var arg0 *C.PangoRenderer
	var arg1 C.PangoRenderPart
	var arg2 C.guint16

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoRenderPart)(part)
	arg2 = C.guint16(alpha)

	C.pango_renderer_set_alpha(arg0, part, alpha)
}

// SetColor sets the color for part of the rendering.
//
// Also see [method@Pango.Renderer.set_alpha].
func (r renderer) SetColor(part RenderPart, color *Color) {
	var arg0 *C.PangoRenderer
	var arg1 C.PangoRenderPart
	var arg2 *C.PangoColor

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoRenderPart)(part)
	arg2 = (*C.PangoColor)(unsafe.Pointer(color.Native()))

	C.pango_renderer_set_color(arg0, part, color)
}

// SetMatrix sets the transformation matrix that will be applied when
// rendering.
func (r renderer) SetMatrix(matrix *Matrix) {
	var arg0 *C.PangoRenderer
	var arg1 *C.PangoMatrix

	arg0 = (*C.PangoRenderer)(unsafe.Pointer(r.Native()))
	arg1 = (*C.PangoMatrix)(unsafe.Pointer(matrix.Native()))

	C.pango_renderer_set_matrix(arg0, matrix)
}

type RendererPrivate struct {
	native C.PangoRendererPrivate
}

// WrapRendererPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRendererPrivate(ptr unsafe.Pointer) *RendererPrivate {
	if ptr == nil {
		return nil
	}

	return (*RendererPrivate)(ptr)
}

func marshalRendererPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRendererPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RendererPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}