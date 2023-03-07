// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// PollableSourceNewFull: utility method for InputStream and OutputStream
// implementations. Creates a new #GSource, as with g_pollable_source_new(), but
// also attaching child_source (with a dummy callback), and cancellable, if they
// are non-NULL.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable to attach.
//    - pollableStream: stream associated with the new source.
//    - childSource (optional): optional child source to attach.
//
// The function returns the following values:
//
//    - source: new #GSource.
//
func PollableSourceNewFull(ctx context.Context, pollableStream *coreglib.Object, childSource *glib.Source) *glib.Source {
	var _arg3 *C.GCancellable // out
	var _arg1 C.gpointer      // out
	var _arg2 *C.GSource      // out
	var _cret *C.GSource      // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gpointer(unsafe.Pointer(pollableStream.Native()))
	if childSource != nil {
		_arg2 = (*C.GSource)(gextras.StructNative(unsafe.Pointer(childSource)))
	}

	_cret = C.g_pollable_source_new_full(_arg1, _arg2, _arg3)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(pollableStream)
	runtime.KeepAlive(childSource)

	var _source *glib.Source // out

	_source = (*glib.Source)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_source)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_source_unref((*C.GSource)(intern.C))
		},
	)

	return _source
}

// PollableStreamRead tries to read from stream, as with g_input_stream_read()
// (if blocking is TRUE) or g_pollable_input_stream_read_nonblocking() (if
// blocking is FALSE). This can be used to more easily share code between
// blocking and non-blocking implementations of a method.
//
// If blocking is FALSE, then stream must be a InputStream for which
// g_pollable_input_stream_can_poll() returns TRUE, or else the behavior is
// undefined. If blocking is TRUE, then stream does not need to be a
// InputStream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - stream: Stream.
//    - buffer to read data into.
//    - blocking: whether to do blocking I/O.
//
// The function returns the following values:
//
//    - gssize: number of bytes read, or -1 on error.
//
func PollableStreamRead(ctx context.Context, stream InputStreamer, buffer []byte, blocking bool) (int, error) {
	var _arg5 *C.GCancellable // out
	var _arg1 *C.GInputStream // out
	var _arg2 *C.void         // out
	var _arg3 C.gsize
	var _arg4 C.gboolean // out
	var _cret C.gssize   // in
	var _cerr *C.GError  // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GInputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg3 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg2 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	if blocking {
		_arg4 = C.TRUE
	}

	_cret = C.g_pollable_stream_read(_arg1, unsafe.Pointer(_arg2), _arg3, _arg4, _arg5, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(blocking)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// PollableStreamWrite tries to write to stream, as with g_output_stream_write()
// (if blocking is TRUE) or g_pollable_output_stream_write_nonblocking() (if
// blocking is FALSE). This can be used to more easily share code between
// blocking and non-blocking implementations of a method.
//
// If blocking is FALSE, then stream must be a OutputStream for which
// g_pollable_output_stream_can_poll() returns TRUE or else the behavior is
// undefined. If blocking is TRUE, then stream does not need to be a
// OutputStream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - stream: Stream.
//    - buffer: buffer containing the data to write.
//    - blocking: whether to do blocking I/O.
//
// The function returns the following values:
//
//    - gssize: number of bytes written, or -1 on error.
//
func PollableStreamWrite(ctx context.Context, stream OutputStreamer, buffer []byte, blocking bool) (int, error) {
	var _arg5 *C.GCancellable  // out
	var _arg1 *C.GOutputStream // out
	var _arg2 *C.void          // out
	var _arg3 C.gsize
	var _arg4 C.gboolean // out
	var _cret C.gssize   // in
	var _cerr *C.GError  // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg3 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg2 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	if blocking {
		_arg4 = C.TRUE
	}

	_cret = C.g_pollable_stream_write(_arg1, unsafe.Pointer(_arg2), _arg3, _arg4, _arg5, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(blocking)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gssize, _goerr
}

// PollableStreamWriteAll tries to write count bytes to stream, as with
// g_output_stream_write_all(), but using g_pollable_stream_write() rather than
// g_output_stream_write().
//
// On a successful write of count bytes, TRUE is returned, and bytes_written is
// set to count.
//
// If there is an error during the operation (including G_IO_ERROR_WOULD_BLOCK
// in the non-blocking case), FALSE is returned and error is set to indicate the
// error status, bytes_written is updated to contain the number of bytes written
// into the stream before the error occurred.
//
// As with g_pollable_stream_write(), if blocking is FALSE, then stream must be
// a OutputStream for which g_pollable_output_stream_can_poll() returns TRUE or
// else the behavior is undefined. If blocking is TRUE, then stream does not
// need to be a OutputStream.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - stream: Stream.
//    - buffer: buffer containing the data to write.
//    - blocking: whether to do blocking I/O.
//
// The function returns the following values:
//
//    - bytesWritten: location to store the number of bytes that was written to
//      the stream.
//
func PollableStreamWriteAll(ctx context.Context, stream OutputStreamer, buffer []byte, blocking bool) (uint, error) {
	var _arg6 *C.GCancellable  // out
	var _arg1 *C.GOutputStream // out
	var _arg2 *C.void          // out
	var _arg3 C.gsize
	var _arg4 C.gboolean // out
	var _arg5 C.gsize    // in
	var _cerr *C.GError  // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GOutputStream)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	_arg3 = (C.gsize)(len(buffer))
	if len(buffer) > 0 {
		_arg2 = (*C.void)(unsafe.Pointer(&buffer[0]))
	}
	if blocking {
		_arg4 = C.TRUE
	}

	C.g_pollable_stream_write_all(_arg1, unsafe.Pointer(_arg2), _arg3, _arg4, &_arg5, _arg6, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(blocking)

	var _bytesWritten uint // out
	var _goerr error       // out

	_bytesWritten = uint(_arg5)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytesWritten, _goerr
}