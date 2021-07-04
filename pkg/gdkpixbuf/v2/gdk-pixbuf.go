// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
//
// gboolean gotk4_PixbufSaveFunc(gchar*, gsize, GError**, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_get_type()), F: marshalPixbuf},
	})
}

// Pixbuf: a pixel buffer.
//
// `GdkPixbuf` contains information about an image's pixel data, its color
// space, bits per sample, width and height, and the rowstride (the number of
// bytes between the start of one row and the start of the next).
//
// Creating new `GdkPixbuf`
//
// The most basic way to create a pixbuf is to wrap an existing pixel buffer
// with a [class@GdkPixbuf.Pixbuf] instance. You can use the
// [`ctor@GdkPixbuf.Pixbuf.new_from_data`] function to do this.
//
// Every time you create a new `GdkPixbuf` instance for some data, you will need
// to specify the destroy notification function that will be called when the
// data buffer needs to be freed; this will happen when a `GdkPixbuf` is
// finalized by the reference counting functions. If you have a chunk of static
// data compiled into your application, you can pass in `NULL` as the destroy
// notification function so that the data will not be freed.
//
// The [`ctor@GdkPixbuf.Pixbuf.new`] constructor function can be used as a
// convenience to create a pixbuf with an empty buffer; this is equivalent to
// allocating a data buffer using `malloc()` and then wrapping it with
// `gdk_pixbuf_new_from_data()`. The `gdk_pixbuf_new()` function will compute an
// optimal rowstride so that rendering can be performed with an efficient
// algorithm.
//
// As a special case, you can use the
// [`ctor@GdkPixbuf.Pixbuf.new_from_xpm_data`] function to create a pixbuf from
// inline XPM image data.
//
// You can also copy an existing pixbuf with the [method@Pixbuf.copy] function.
// This is not the same as just acquiring a reference to the old pixbuf
// instance: the copy function will actually duplicate the pixel data in memory
// and create a new [class@Pixbuf] instance for it.
//
//
// Reference counting
//
// `GdkPixbuf` structures are reference counted. This means that an application
// can share a single pixbuf among many parts of the code. When a piece of the
// program needs to use a pixbuf, it should acquire a reference to it by calling
// `g_object_ref()`; when it no longer needs the pixbuf, it should release the
// reference it acquired by calling `g_object_unref()`. The resources associated
// with a `GdkPixbuf` will be freed when its reference count drops to zero.
// Newly-created `GdkPixbuf` instances start with a reference count of one.
//
//
// Image Data
//
// Image data in a pixbuf is stored in memory in an uncompressed, packed format.
// Rows in the image are stored top to bottom, and in each row pixels are stored
// from left to right.
//
// There may be padding at the end of a row.
//
// The "rowstride" value of a pixbuf, as returned by
// [`method@GdkPixbuf.Pixbuf.get_rowstride`], indicates the number of bytes
// between rows.
//
// **NOTE**: If you are copying raw pixbuf data with `memcpy()` note that the
// last row in the pixbuf may not be as wide as the full rowstride, but rather
// just as wide as the pixel data needs to be; that is: it is unsafe to do
// `memcpy (dest, pixels, rowstride * height)` to copy a whole pixbuf. Use
// [method@GdkPixbuf.Pixbuf.copy] instead, or compute the width in bytes of the
// last row as:
//
// “`c last_row = width * ((n_channels * bits_per_sample + 7) / 8); “`
//
// The same rule applies when iterating over each row of a `GdkPixbuf` pixels
// array.
//
// The following code illustrates a simple `put_pixel()` function for RGB
// pixbufs with 8 bits per channel with an alpha channel.
//
// “`c static void put_pixel (GdkPixbuf *pixbuf, int x, int y, guchar red,
// guchar green, guchar blue, guchar alpha) { int n_channels =
// gdk_pixbuf_get_n_channels (pixbuf);
//
//    // Ensure that the pixbuf is valid
//    g_assert (gdk_pixbuf_get_colorspace (pixbuf) == GDK_COLORSPACE_RGB);
//    g_assert (gdk_pixbuf_get_bits_per_sample (pixbuf) == 8);
//    g_assert (gdk_pixbuf_get_has_alpha (pixbuf));
//    g_assert (n_channels == 4);
//
//    int width = gdk_pixbuf_get_width (pixbuf);
//    int height = gdk_pixbuf_get_height (pixbuf);
//
//    // Ensure that the coordinates are in a valid range
//    g_assert (x >= 0 && x < width);
//    g_assert (y >= 0 && y < height);
//
//    int rowstride = gdk_pixbuf_get_rowstride (pixbuf);
//
//    // The pixel buffer in the GdkPixbuf instance
//    guchar *pixels = gdk_pixbuf_get_pixels (pixbuf);
//
//    // The pixel we wish to modify
//    guchar *p = pixels + y * rowstride + x * n_channels;
//    p[0] = red;
//    p[1] = green;
//    p[2] = blue;
//    p[3] = alpha;
//
// } “`
//
//
// Loading images
//
// The `GdkPixBuf` class provides a simple mechanism for loading an image from a
// file in synchronous and asynchronous fashion.
//
// For GUI applications, it is recommended to use the asynchronous stream API to
// avoid blocking the control flow of the application.
//
// Additionally, `GdkPixbuf` provides the [class@GdkPixbuf.PixbufLoader`] API
// for progressive image loading.
//
//
// Saving images
//
// The `GdkPixbuf` class provides methods for saving image data in a number of
// file formats. The formatted data can be written to a file or to a memory
// buffer. `GdkPixbuf` can also call a user-defined callback on the data, which
// allows to e.g. write the image to a socket or store it in a database.
type Pixbuf interface {
	gextras.Objector

	AddAlphaPixbuf(substituteColor bool, r byte, g byte, b byte) Pixbuf

	ApplyEmbeddedOrientationPixbuf() Pixbuf

	CompositePixbuf(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int)

	CompositeColorPixbuf(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int, checkX int, checkY int, checkSize int, color1 uint32, color2 uint32)

	CompositeColorSimplePixbuf(destWidth int, destHeight int, interpType InterpType, overallAlpha int, checkSize int, color1 uint32, color2 uint32) Pixbuf

	CopyPixbuf() Pixbuf

	CopyAreaPixbuf(srcX int, srcY int, width int, height int, destPixbuf Pixbuf, destX int, destY int)

	CopyOptionsPixbuf(destPixbuf Pixbuf) bool

	FillPixbuf(pixel uint32)

	FlipPixbuf(horizontal bool) Pixbuf

	BitsPerSample() int

	ByteLength() uint

	Colorspace() Colorspace

	HasAlpha() bool

	Height() int

	NChannels() int

	Option(key string) string

	Options() *glib.HashTable

	Rowstride() int

	Width() int

	NewSubpixbufPixbuf(srcX int, srcY int, width int, height int) Pixbuf

	ReadPixelsPixbuf() *byte

	RemoveOptionPixbuf(key string) bool

	RotateSimplePixbuf(angle PixbufRotation) Pixbuf

	SaturateAndPixelatePixbuf(dest Pixbuf, saturation float32, pixelate bool)

	SaveToBuffervPixbuf(typ string, optionKeys []string, optionValues []string) ([]byte, error)

	SaveToCallbackvPixbuf(saveFunc PixbufSaveFunc, typ string, optionKeys []string, optionValues []string) error

	SaveToStreamvPixbuf(stream gio.OutputStream, typ string, optionKeys []string, optionValues []string, cancellable gio.Cancellable) error

	SavevPixbuf(filename string, typ string, optionKeys []string, optionValues []string) error

	ScalePixbuf(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType)

	ScaleSimplePixbuf(destWidth int, destHeight int, interpType InterpType) Pixbuf

	SetOptionPixbuf(key string, value string) bool
}

// pixbuf implements the Pixbuf class.
type pixbuf struct {
	gextras.Objector
}

// WrapPixbuf wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbuf(obj *externglib.Object) Pixbuf {
	return pixbuf{
		Objector: obj,
	}
}

func marshalPixbuf(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbuf(obj), nil
}

func NewPixbuf(colorspace Colorspace, hasAlpha bool, bitsPerSample int, width int, height int) Pixbuf {
	var _arg1 C.GdkColorspace // out
	var _arg2 C.gboolean      // out
	var _arg3 C.int           // out
	var _arg4 C.int           // out
	var _arg5 C.int           // out
	var _cret *C.GdkPixbuf    // in

	_arg1 = C.GdkColorspace(colorspace)
	if hasAlpha {
		_arg2 = C.TRUE
	}
	_arg3 = C.int(bitsPerSample)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	_cret = C.gdk_pixbuf_new(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _pixbuf Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _pixbuf
}

func NewPixbufFromFile(filename string) (Pixbuf, error) {
	var _arg1 *C.char      // out
	var _cret *C.GdkPixbuf // in
	var _cerr *C.GError    // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_new_from_file(_arg1, &_cerr)

	var _pixbuf Pixbuf // out
	var _goerr error   // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

func NewPixbufFromFileAtScale(filename string, width int, height int, preserveAspectRatio bool) (Pixbuf, error) {
	var _arg1 *C.char      // out
	var _arg2 C.int        // out
	var _arg3 C.int        // out
	var _arg4 C.gboolean   // out
	var _cret *C.GdkPixbuf // in
	var _cerr *C.GError    // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(width)
	_arg3 = C.int(height)
	if preserveAspectRatio {
		_arg4 = C.TRUE
	}

	_cret = C.gdk_pixbuf_new_from_file_at_scale(_arg1, _arg2, _arg3, _arg4, &_cerr)

	var _pixbuf Pixbuf // out
	var _goerr error   // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

func NewPixbufFromFileAtSize(filename string, width int, height int) (Pixbuf, error) {
	var _arg1 *C.char      // out
	var _arg2 C.int        // out
	var _arg3 C.int        // out
	var _cret *C.GdkPixbuf // in
	var _cerr *C.GError    // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(width)
	_arg3 = C.int(height)

	_cret = C.gdk_pixbuf_new_from_file_at_size(_arg1, _arg2, _arg3, &_cerr)

	var _pixbuf Pixbuf // out
	var _goerr error   // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

func NewPixbufFromInline(data []byte, copyPixels bool) (Pixbuf, error) {
	var _arg2 *C.guint8
	var _arg1 C.gint
	var _arg3 C.gboolean   // out
	var _cret *C.GdkPixbuf // in
	var _cerr *C.GError    // in

	_arg1 = C.gint(len(data))
	_arg2 = (*C.guint8)(unsafe.Pointer(&data[0]))
	if copyPixels {
		_arg3 = C.TRUE
	}

	_cret = C.gdk_pixbuf_new_from_inline(_arg1, _arg2, _arg3, &_cerr)

	var _pixbuf Pixbuf // out
	var _goerr error   // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

func NewPixbufFromResource(resourcePath string) (Pixbuf, error) {
	var _arg1 *C.char      // out
	var _cret *C.GdkPixbuf // in
	var _cerr *C.GError    // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_new_from_resource(_arg1, &_cerr)

	var _pixbuf Pixbuf // out
	var _goerr error   // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

func NewPixbufFromResourceAtScale(resourcePath string, width int, height int, preserveAspectRatio bool) (Pixbuf, error) {
	var _arg1 *C.char      // out
	var _arg2 C.int        // out
	var _arg3 C.int        // out
	var _arg4 C.gboolean   // out
	var _cret *C.GdkPixbuf // in
	var _cerr *C.GError    // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(width)
	_arg3 = C.int(height)
	if preserveAspectRatio {
		_arg4 = C.TRUE
	}

	_cret = C.gdk_pixbuf_new_from_resource_at_scale(_arg1, _arg2, _arg3, _arg4, &_cerr)

	var _pixbuf Pixbuf // out
	var _goerr error   // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

func NewPixbufFromStream(stream gio.InputStream, cancellable gio.Cancellable) (Pixbuf, error) {
	var _arg1 *C.GInputStream // out
	var _arg2 *C.GCancellable // out
	var _cret *C.GdkPixbuf    // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.gdk_pixbuf_new_from_stream(_arg1, _arg2, &_cerr)

	var _pixbuf Pixbuf // out
	var _goerr error   // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

func NewPixbufFromStreamAtScale(stream gio.InputStream, width int, height int, preserveAspectRatio bool, cancellable gio.Cancellable) (Pixbuf, error) {
	var _arg1 *C.GInputStream // out
	var _arg2 C.gint          // out
	var _arg3 C.gint          // out
	var _arg4 C.gboolean      // out
	var _arg5 *C.GCancellable // out
	var _cret *C.GdkPixbuf    // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))
	_arg2 = C.gint(width)
	_arg3 = C.gint(height)
	if preserveAspectRatio {
		_arg4 = C.TRUE
	}
	_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.gdk_pixbuf_new_from_stream_at_scale(_arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)

	var _pixbuf Pixbuf // out
	var _goerr error   // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pixbuf, _goerr
}

func NewPixbufFromXpmData(data []string) Pixbuf {
	var _arg1 **C.char
	var _cret *C.GdkPixbuf // in

	_arg1 = (**C.char)(C.malloc(C.ulong(len(data)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(data))
		for i := range data {
			out[i] = (*C.char)(C.CString(data[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	_cret = C.gdk_pixbuf_new_from_xpm_data(_arg1)

	var _pixbuf Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _pixbuf
}

func (p pixbuf) AddAlphaPixbuf(substituteColor bool, r byte, g byte, b byte) Pixbuf {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 C.gboolean   // out
	var _arg2 C.guchar     // out
	var _arg3 C.guchar     // out
	var _arg4 C.guchar     // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	if substituteColor {
		_arg1 = C.TRUE
	}
	_arg2 = C.guchar(r)
	_arg3 = C.guchar(g)
	_arg4 = C.guchar(b)

	_cret = C.gdk_pixbuf_add_alpha(_arg0, _arg1, _arg2, _arg3, _arg4)

	var _ret Pixbuf // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _ret
}

func (s pixbuf) ApplyEmbeddedOrientationPixbuf() Pixbuf {
	var _arg0 *C.GdkPixbuf // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_pixbuf_apply_embedded_orientation(_arg0)

	var _pixbuf Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _pixbuf
}

func (s pixbuf) CompositePixbuf(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int) {
	var _arg0 *C.GdkPixbuf     // out
	var _arg1 *C.GdkPixbuf     // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out
	var _arg4 C.int            // out
	var _arg5 C.int            // out
	var _arg6 C.double         // out
	var _arg7 C.double         // out
	var _arg8 C.double         // out
	var _arg9 C.double         // out
	var _arg10 C.GdkInterpType // out
	var _arg11 C.int           // out

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(dest.Native()))
	_arg2 = C.int(destX)
	_arg3 = C.int(destY)
	_arg4 = C.int(destWidth)
	_arg5 = C.int(destHeight)
	_arg6 = C.double(offsetX)
	_arg7 = C.double(offsetY)
	_arg8 = C.double(scaleX)
	_arg9 = C.double(scaleY)
	_arg10 = C.GdkInterpType(interpType)
	_arg11 = C.int(overallAlpha)

	C.gdk_pixbuf_composite(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11)
}

func (s pixbuf) CompositeColorPixbuf(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType, overallAlpha int, checkX int, checkY int, checkSize int, color1 uint32, color2 uint32) {
	var _arg0 *C.GdkPixbuf     // out
	var _arg1 *C.GdkPixbuf     // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out
	var _arg4 C.int            // out
	var _arg5 C.int            // out
	var _arg6 C.double         // out
	var _arg7 C.double         // out
	var _arg8 C.double         // out
	var _arg9 C.double         // out
	var _arg10 C.GdkInterpType // out
	var _arg11 C.int           // out
	var _arg12 C.int           // out
	var _arg13 C.int           // out
	var _arg14 C.int           // out
	var _arg15 C.guint32       // out
	var _arg16 C.guint32       // out

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(dest.Native()))
	_arg2 = C.int(destX)
	_arg3 = C.int(destY)
	_arg4 = C.int(destWidth)
	_arg5 = C.int(destHeight)
	_arg6 = C.double(offsetX)
	_arg7 = C.double(offsetY)
	_arg8 = C.double(scaleX)
	_arg9 = C.double(scaleY)
	_arg10 = C.GdkInterpType(interpType)
	_arg11 = C.int(overallAlpha)
	_arg12 = C.int(checkX)
	_arg13 = C.int(checkY)
	_arg14 = C.int(checkSize)
	_arg15 = C.guint32(color1)
	_arg16 = C.guint32(color2)

	C.gdk_pixbuf_composite_color(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10, _arg11, _arg12, _arg13, _arg14, _arg15, _arg16)
}

func (s pixbuf) CompositeColorSimplePixbuf(destWidth int, destHeight int, interpType InterpType, overallAlpha int, checkSize int, color1 uint32, color2 uint32) Pixbuf {
	var _arg0 *C.GdkPixbuf    // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _arg3 C.GdkInterpType // out
	var _arg4 C.int           // out
	var _arg5 C.int           // out
	var _arg6 C.guint32       // out
	var _arg7 C.guint32       // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(destWidth)
	_arg2 = C.int(destHeight)
	_arg3 = C.GdkInterpType(interpType)
	_arg4 = C.int(overallAlpha)
	_arg5 = C.int(checkSize)
	_arg6 = C.guint32(color1)
	_arg7 = C.guint32(color2)

	_cret = C.gdk_pixbuf_composite_color_simple(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _pixbuf Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _pixbuf
}

func (p pixbuf) CopyPixbuf() Pixbuf {
	var _arg0 *C.GdkPixbuf // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_copy(_arg0)

	var _ret Pixbuf // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _ret
}

func (s pixbuf) CopyAreaPixbuf(srcX int, srcY int, width int, height int, destPixbuf Pixbuf, destX int, destY int) {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out
	var _arg3 C.int        // out
	var _arg4 C.int        // out
	var _arg5 *C.GdkPixbuf // out
	var _arg6 C.int        // out
	var _arg7 C.int        // out

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(srcX)
	_arg2 = C.int(srcY)
	_arg3 = C.int(width)
	_arg4 = C.int(height)
	_arg5 = (*C.GdkPixbuf)(unsafe.Pointer(destPixbuf.Native()))
	_arg6 = C.int(destX)
	_arg7 = C.int(destY)

	C.gdk_pixbuf_copy_area(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

func (s pixbuf) CopyOptionsPixbuf(destPixbuf Pixbuf) bool {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 *C.GdkPixbuf // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(destPixbuf.Native()))

	_cret = C.gdk_pixbuf_copy_options(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p pixbuf) FillPixbuf(pixel uint32) {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 C.guint32    // out

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	_arg1 = C.guint32(pixel)

	C.gdk_pixbuf_fill(_arg0, _arg1)
}

func (s pixbuf) FlipPixbuf(horizontal bool) Pixbuf {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 C.gboolean   // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	if horizontal {
		_arg1 = C.TRUE
	}

	_cret = C.gdk_pixbuf_flip(_arg0, _arg1)

	var _pixbuf Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _pixbuf
}

func (p pixbuf) BitsPerSample() int {
	var _arg0 *C.GdkPixbuf // out
	var _cret C.int        // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_bits_per_sample(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p pixbuf) ByteLength() uint {
	var _arg0 *C.GdkPixbuf // out
	var _cret C.gsize      // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_byte_length(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

func (p pixbuf) Colorspace() Colorspace {
	var _arg0 *C.GdkPixbuf    // out
	var _cret C.GdkColorspace // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_colorspace(_arg0)

	var _colorspace Colorspace // out

	_colorspace = Colorspace(_cret)

	return _colorspace
}

func (p pixbuf) HasAlpha() bool {
	var _arg0 *C.GdkPixbuf // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_has_alpha(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p pixbuf) Height() int {
	var _arg0 *C.GdkPixbuf // out
	var _cret C.int        // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p pixbuf) NChannels() int {
	var _arg0 *C.GdkPixbuf // out
	var _cret C.int        // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_n_channels(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p pixbuf) Option(key string) string {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 *C.gchar     // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_get_option(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p pixbuf) Options() *glib.HashTable {
	var _arg0 *C.GdkPixbuf  // out
	var _cret *C.GHashTable // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_options(_arg0)

	var _hashTable *glib.HashTable // out

	_hashTable = (*glib.HashTable)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_hashTable, func(v **glib.HashTable) {
		C.free(unsafe.Pointer(v))
	})

	return _hashTable
}

func (p pixbuf) Rowstride() int {
	var _arg0 *C.GdkPixbuf // out
	var _cret C.int        // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_rowstride(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p pixbuf) Width() int {
	var _arg0 *C.GdkPixbuf // out
	var _cret C.int        // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_get_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s pixbuf) NewSubpixbufPixbuf(srcX int, srcY int, width int, height int) Pixbuf {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out
	var _arg3 C.int        // out
	var _arg4 C.int        // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(srcX)
	_arg2 = C.int(srcY)
	_arg3 = C.int(width)
	_arg4 = C.int(height)

	_cret = C.gdk_pixbuf_new_subpixbuf(_arg0, _arg1, _arg2, _arg3, _arg4)

	var _pixbuf Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _pixbuf
}

func (p pixbuf) ReadPixelsPixbuf() *byte {
	var _arg0 *C.GdkPixbuf // out
	var _cret *C.guint8    // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_pixbuf_read_pixels(_arg0)

	var _guint8 *byte // out

	_guint8 = (*byte)(unsafe.Pointer(_cret))

	return _guint8
}

func (p pixbuf) RemoveOptionPixbuf(key string) bool {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_remove_option(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s pixbuf) RotateSimplePixbuf(angle PixbufRotation) Pixbuf {
	var _arg0 *C.GdkPixbuf        // out
	var _arg1 C.GdkPixbufRotation // out
	var _cret *C.GdkPixbuf        // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkPixbufRotation(angle)

	_cret = C.gdk_pixbuf_rotate_simple(_arg0, _arg1)

	var _pixbuf Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _pixbuf
}

func (s pixbuf) SaturateAndPixelatePixbuf(dest Pixbuf, saturation float32, pixelate bool) {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 *C.GdkPixbuf // out
	var _arg2 C.gfloat     // out
	var _arg3 C.gboolean   // out

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(dest.Native()))
	_arg2 = C.gfloat(saturation)
	if pixelate {
		_arg3 = C.TRUE
	}

	C.gdk_pixbuf_saturate_and_pixelate(_arg0, _arg1, _arg2, _arg3)
}

func (p pixbuf) SaveToBuffervPixbuf(typ string, optionKeys []string, optionValues []string) ([]byte, error) {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 *C.gchar
	var _arg2 C.gsize // in
	var _arg3 *C.char // out
	var _arg4 **C.char
	var _arg5 **C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	_arg3 = (*C.char)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (**C.char)(C.malloc(C.ulong(len(optionKeys)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice(_arg4, len(optionKeys))
		for i := range optionKeys {
			out[i] = (*C.char)(C.CString(optionKeys[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	_arg5 = (**C.char)(C.malloc(C.ulong(len(optionValues)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg5))
	{
		out := unsafe.Slice(_arg5, len(optionValues))
		for i := range optionValues {
			out[i] = (*C.char)(C.CString(optionValues[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gdk_pixbuf_save_to_bufferv(_arg0, &_arg1, &_arg2, _arg3, _arg4, _arg5, &_cerr)

	var _buffer []byte
	var _goerr error // out

	_buffer = unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), _arg2)
	runtime.SetFinalizer(&_buffer, func(v *[]byte) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _buffer, _goerr
}

func (p pixbuf) SaveToCallbackvPixbuf(saveFunc PixbufSaveFunc, typ string, optionKeys []string, optionValues []string) error {
	var _arg0 *C.GdkPixbuf        // out
	var _arg1 C.GdkPixbufSaveFunc // out
	var _arg2 C.gpointer
	var _arg3 *C.char // out
	var _arg4 **C.char
	var _arg5 **C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	_arg1 = (*[0]byte)(C.gotk4_PixbufSaveFunc)
	_arg2 = C.gpointer(box.Assign(saveFunc))
	_arg3 = (*C.char)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (**C.char)(C.malloc(C.ulong(len(optionKeys)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice(_arg4, len(optionKeys))
		for i := range optionKeys {
			out[i] = (*C.char)(C.CString(optionKeys[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	_arg5 = (**C.char)(C.malloc(C.ulong(len(optionValues)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg5))
	{
		out := unsafe.Slice(_arg5, len(optionValues))
		for i := range optionValues {
			out[i] = (*C.char)(C.CString(optionValues[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gdk_pixbuf_save_to_callbackv(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (p pixbuf) SaveToStreamvPixbuf(stream gio.OutputStream, typ string, optionKeys []string, optionValues []string, cancellable gio.Cancellable) error {
	var _arg0 *C.GdkPixbuf     // out
	var _arg1 *C.GOutputStream // out
	var _arg2 *C.char          // out
	var _arg3 **C.char
	var _arg4 **C.char
	var _arg5 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GOutputStream)(unsafe.Pointer(stream.Native()))
	_arg2 = (*C.char)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (**C.char)(C.malloc(C.ulong(len(optionKeys)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice(_arg3, len(optionKeys))
		for i := range optionKeys {
			out[i] = (*C.char)(C.CString(optionKeys[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	_arg4 = (**C.char)(C.malloc(C.ulong(len(optionValues)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice(_arg4, len(optionValues))
		for i := range optionValues {
			out[i] = (*C.char)(C.CString(optionValues[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.gdk_pixbuf_save_to_streamv(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (p pixbuf) SavevPixbuf(filename string, typ string, optionKeys []string, optionValues []string) error {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out
	var _arg3 **C.char
	var _arg4 **C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(typ))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (**C.char)(C.malloc(C.ulong(len(optionKeys)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice(_arg3, len(optionKeys))
		for i := range optionKeys {
			out[i] = (*C.char)(C.CString(optionKeys[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	_arg4 = (**C.char)(C.malloc(C.ulong(len(optionValues)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice(_arg4, len(optionValues))
		for i := range optionValues {
			out[i] = (*C.char)(C.CString(optionValues[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gdk_pixbuf_savev(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (s pixbuf) ScalePixbuf(dest Pixbuf, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType InterpType) {
	var _arg0 *C.GdkPixbuf     // out
	var _arg1 *C.GdkPixbuf     // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out
	var _arg4 C.int            // out
	var _arg5 C.int            // out
	var _arg6 C.double         // out
	var _arg7 C.double         // out
	var _arg8 C.double         // out
	var _arg9 C.double         // out
	var _arg10 C.GdkInterpType // out

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(dest.Native()))
	_arg2 = C.int(destX)
	_arg3 = C.int(destY)
	_arg4 = C.int(destWidth)
	_arg5 = C.int(destHeight)
	_arg6 = C.double(offsetX)
	_arg7 = C.double(offsetY)
	_arg8 = C.double(scaleX)
	_arg9 = C.double(scaleY)
	_arg10 = C.GdkInterpType(interpType)

	C.gdk_pixbuf_scale(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9, _arg10)
}

func (s pixbuf) ScaleSimplePixbuf(destWidth int, destHeight int, interpType InterpType) Pixbuf {
	var _arg0 *C.GdkPixbuf    // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _arg3 C.GdkInterpType // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(destWidth)
	_arg2 = C.int(destHeight)
	_arg3 = C.GdkInterpType(interpType)

	_cret = C.gdk_pixbuf_scale_simple(_arg0, _arg1, _arg2, _arg3)

	var _pixbuf Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Pixbuf)

	return _pixbuf
}

func (p pixbuf) SetOptionPixbuf(key string, value string) bool {
	var _arg0 *C.GdkPixbuf // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkPixbuf)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gdk_pixbuf_set_option(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}