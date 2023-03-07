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
// #include <glib-object.h>
// GIOCondition _gotk4_gio2_DatagramBased_virtual_condition_check(void* fnptr, GDatagramBased* arg0, GIOCondition arg1) {
//   return ((GIOCondition (*)(GDatagramBased*, GIOCondition))(fnptr))(arg0, arg1);
// };
// GSource* _gotk4_gio2_DatagramBased_virtual_create_source(void* fnptr, GDatagramBased* arg0, GIOCondition arg1, GCancellable* arg2) {
//   return ((GSource* (*)(GDatagramBased*, GIOCondition, GCancellable*))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gio2_DatagramBased_virtual_condition_wait(void* fnptr, GDatagramBased* arg0, GIOCondition arg1, gint64 arg2, GCancellable* arg3, GError** arg4) {
//   return ((gboolean (*)(GDatagramBased*, GIOCondition, gint64, GCancellable*, GError**))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// gint _gotk4_gio2_DatagramBased_virtual_receive_messages(void* fnptr, GDatagramBased* arg0, GInputMessage* arg1, guint arg2, gint arg3, gint64 arg4, GCancellable* arg5, GError** arg6) {
//   return ((gint (*)(GDatagramBased*, GInputMessage*, guint, gint, gint64, GCancellable*, GError**))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
// };
// gint _gotk4_gio2_DatagramBased_virtual_send_messages(void* fnptr, GDatagramBased* arg0, GOutputMessage* arg1, guint arg2, gint arg3, gint64 arg4, GCancellable* arg5, GError** arg6) {
//   return ((gint (*)(GDatagramBased*, GOutputMessage*, guint, gint, gint64, GCancellable*, GError**))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5, arg6);
// };
import "C"

// GType values.
var (
	GTypeDatagramBased = coreglib.Type(C.g_datagram_based_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDatagramBased, F: marshalDatagramBased},
	})
}

// DatagramBased is a networking interface for representing datagram-based
// communications. It is a more or less direct mapping of the core parts of the
// BSD socket API in a portable GObject interface. It is implemented by
// #GSocket, which wraps the UNIX socket API on UNIX and winsock2 on Windows.
//
// Based is entirely platform independent, and is intended to be used alongside
// higher-level networking APIs such as OStream.
//
// It uses vectored scatter/gather I/O by default, allowing for many messages to
// be sent or received in a single call. Where possible, implementations of the
// interface should take advantage of vectored I/O to minimise processing or
// system calls. For example, #GSocket uses recvmmsg() and sendmmsg() where
// possible. Callers should take advantage of scatter/gather I/O (the use of
// multiple buffers per message) to avoid unnecessary copying of data to
// assemble or disassemble a message.
//
// Each Based operation has a timeout parameter which may be negative for
// blocking behaviour, zero for non-blocking behaviour, or positive for timeout
// behaviour. A blocking operation blocks until finished or there is an error. A
// non-blocking operation will return immediately with a G_IO_ERROR_WOULD_BLOCK
// error if it cannot make progress. A timeout operation will block until the
// operation is complete or the timeout expires; if the timeout expires it will
// return what progress it made, or G_IO_ERROR_TIMED_OUT if no progress was
// made. To know when a call would successfully run you can call
// g_datagram_based_condition_check() or g_datagram_based_condition_wait(). You
// can also use g_datagram_based_create_source() and attach it to a Context to
// get callbacks when I/O is possible.
//
// When running a non-blocking operation applications should always be able to
// handle getting a G_IO_ERROR_WOULD_BLOCK error even when some other function
// said that I/O was possible. This can easily happen in case of a race
// condition in the application, but it can also happen for other reasons. For
// instance, on Windows a socket is always seen as writable until a write
// returns G_IO_ERROR_WOULD_BLOCK.
//
// As with #GSocket, Baseds can be either connection oriented (for example,
// SCTP) or connectionless (for example, UDP). Baseds must be datagram-based,
// not stream-based. The interface does not cover connection establishment — use
// methods on the underlying type to establish a connection before sending and
// receiving data through the Based API. For connectionless socket types the
// target/source address is specified or received in each I/O operation.
//
// Like most other APIs in GLib, Based is not inherently thread safe. To use a
// Based concurrently from multiple threads, you must implement your own
// locking.
//
// DatagramBased wraps an interface. This means the user can get the
// underlying type by calling Cast().
type DatagramBased struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DatagramBased)(nil)
)

// DatagramBasedder describes DatagramBased's interface methods.
type DatagramBasedder interface {
	coreglib.Objector

	// ConditionCheck checks on the readiness of datagram_based to perform
	// operations.
	ConditionCheck(condition glib.IOCondition) glib.IOCondition
	// ConditionWait waits for up to timeout microseconds for condition to
	// become true on datagram_based.
	ConditionWait(ctx context.Context, condition glib.IOCondition, timeout int64) error
	// CreateSource creates a #GSource that can be attached to a Context to
	// monitor for the availability of the specified condition on the Based.
	CreateSource(ctx context.Context, condition glib.IOCondition) *glib.Source
	// ReceiveMessages: receive one or more data messages from datagram_based in
	// one go.
	ReceiveMessages(ctx context.Context, messages []InputMessage, flags int, timeout int64) (int, error)
	// SendMessages: send one or more data messages from datagram_based in one
	// go.
	SendMessages(ctx context.Context, messages []OutputMessage, flags int, timeout int64) (int, error)
}

var _ DatagramBasedder = (*DatagramBased)(nil)

func wrapDatagramBased(obj *coreglib.Object) *DatagramBased {
	return &DatagramBased{
		Object: obj,
	}
}

func marshalDatagramBased(p uintptr) (interface{}, error) {
	return wrapDatagramBased(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConditionCheck checks on the readiness of datagram_based to perform
// operations. The operations specified in condition are checked for and masked
// against the currently-satisfied conditions on datagram_based. The result is
// returned.
//
// G_IO_IN will be set in the return value if data is available to read with
// g_datagram_based_receive_messages(), or if the connection is closed remotely
// (EOS); and if the datagram_based has not been closed locally using some
// implementation-specific method (such as g_socket_close() or
// g_socket_shutdown() with shutdown_read set, if it’s a #GSocket).
//
// If the connection is shut down or closed (by calling g_socket_close() or
// g_socket_shutdown() with shutdown_read set, if it’s a #GSocket, for example),
// all calls to this function will return G_IO_ERROR_CLOSED.
//
// G_IO_OUT will be set if it is expected that at least one byte can be sent
// using g_datagram_based_send_messages() without blocking. It will not be set
// if the datagram_based has been closed locally.
//
// G_IO_HUP will be set if the connection has been closed locally.
//
// G_IO_ERR will be set if there was an asynchronous error in transmitting data
// previously enqueued using g_datagram_based_send_messages().
//
// Note that on Windows, it is possible for an operation to return
// G_IO_ERROR_WOULD_BLOCK even immediately after
// g_datagram_based_condition_check() has claimed that the Based is ready for
// writing. Rather than calling g_datagram_based_condition_check() and then
// writing to the Based if it succeeds, it is generally better to simply try
// writing right away, and try again later if the initial attempt returns
// G_IO_ERROR_WOULD_BLOCK.
//
// It is meaningless to specify G_IO_ERR or G_IO_HUP in condition; these
// conditions will always be set in the output if they are true. Apart from
// these flags, the output is guaranteed to be masked by condition.
//
// This call never blocks.
//
// The function takes the following parameters:
//
//    - condition mask to check.
//
// The function returns the following values:
//
//    - ioCondition mask of the current state.
//
func (datagramBased *DatagramBased) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	var _arg0 *C.GDatagramBased // out
	var _arg1 C.GIOCondition    // out
	var _cret C.GIOCondition    // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	_arg1 = C.GIOCondition(condition)

	_cret = C.g_datagram_based_condition_check(_arg0, _arg1)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(condition)

	var _ioCondition glib.IOCondition // out

	_ioCondition = glib.IOCondition(_cret)

	return _ioCondition
}

// ConditionWait waits for up to timeout microseconds for condition to become
// true on datagram_based. If the condition is met, TRUE is returned.
//
// If cancellable is cancelled before the condition is met, or if timeout is
// reached before the condition is met, then FALSE is returned and error is set
// appropriately (G_IO_ERROR_CANCELLED or G_IO_ERROR_TIMED_OUT).
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//    - condition mask to wait for.
//    - timeout: maximum time (in microseconds) to wait, 0 to not block, or -1 to
//      block indefinitely.
//
func (datagramBased *DatagramBased) ConditionWait(ctx context.Context, condition glib.IOCondition, timeout int64) error {
	var _arg0 *C.GDatagramBased // out
	var _arg3 *C.GCancellable   // out
	var _arg1 C.GIOCondition    // out
	var _arg2 C.gint64          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GIOCondition(condition)
	_arg2 = C.gint64(timeout)

	C.g_datagram_based_condition_wait(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(condition)
	runtime.KeepAlive(timeout)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CreateSource creates a #GSource that can be attached to a Context to monitor
// for the availability of the specified condition on the Based. The #GSource
// keeps a reference to the datagram_based.
//
// The callback on the source is of the BasedSourceFunc type.
//
// It is meaningless to specify G_IO_ERR or G_IO_HUP in condition; these
// conditions will always be reported in the callback if they are true.
//
// If non-NULL, cancellable can be used to cancel the source, which will cause
// the source to trigger, reporting the current condition (which is likely 0
// unless cancellation happened at the same time as a condition change). You can
// check for this in the callback using g_cancellable_is_cancelled().
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//    - condition mask to monitor.
//
// The function returns the following values:
//
//    - source: newly allocated #GSource.
//
func (datagramBased *DatagramBased) CreateSource(ctx context.Context, condition glib.IOCondition) *glib.Source {
	var _arg0 *C.GDatagramBased // out
	var _arg2 *C.GCancellable   // out
	var _arg1 C.GIOCondition    // out
	var _cret *C.GSource        // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GIOCondition(condition)

	_cret = C.g_datagram_based_create_source(_arg0, _arg1, _arg2)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(condition)

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

// ReceiveMessages: receive one or more data messages from datagram_based in one
// go.
//
// messages must point to an array of Message structs and num_messages must be
// the length of this array. Each Message contains a pointer to an array of
// Vector structs describing the buffers that the data received in each message
// will be written to.
//
// flags modify how all messages are received. The commonly available arguments
// for this are available in the MsgFlags enum, but the values there are the
// same as the system values, and the flags are passed in as-is, so you can pass
// in system-specific flags too. These flags affect the overall receive
// operation. Flags affecting individual messages are returned in Message.flags.
//
// The other members of Message are treated as described in its documentation.
//
// If timeout is negative the call will block until num_messages have been
// received, the connection is closed remotely (EOS), cancellable is cancelled,
// or an error occurs.
//
// If timeout is 0 the call will return up to num_messages without blocking, or
// G_IO_ERROR_WOULD_BLOCK if no messages are queued in the operating system to
// be received.
//
// If timeout is positive the call will block on the same conditions as if
// timeout were negative. If the timeout is reached before any messages are
// received, G_IO_ERROR_TIMED_OUT is returned, otherwise it will return the
// number of messages received before timing out. (Note: This is effectively the
// behaviour of MSG_WAITFORONE with recvmmsg().)
//
// To be notified when messages are available, wait for the G_IO_IN condition.
// Note though that you may still receive G_IO_ERROR_WOULD_BLOCK from
// g_datagram_based_receive_messages() even if you were previously notified of a
// G_IO_IN condition.
//
// If the remote peer closes the connection, any messages queued in the
// underlying receive buffer will be returned, and subsequent calls to
// g_datagram_based_receive_messages() will return 0 (with no error set).
//
// If the connection is shut down or closed (by calling g_socket_close() or
// g_socket_shutdown() with shutdown_read set, if it’s a #GSocket, for example),
// all calls to this function will return G_IO_ERROR_CLOSED.
//
// On error -1 is returned and error is set accordingly. An error will only be
// returned if zero messages could be received; otherwise the number of messages
// successfully received before the error will be returned. If cancellable is
// cancelled, G_IO_ERROR_CANCELLED is returned as with any other error.
//
// The function takes the following parameters:
//
//    - ctx (optional): GCancellable.
//    - messages: array of Message structs.
//    - flags: int containing MsgFlags flags for the overall operation.
//    - timeout: maximum time (in microseconds) to wait, 0 to not block, or -1 to
//      block indefinitely.
//
// The function returns the following values:
//
//    - gint: number of messages received, or -1 on error. Note that the number
//      of messages received may be smaller than num_messages if timeout is zero
//      or positive, if the peer closed the connection, or if num_messages was
//      larger than UIO_MAXIOV (1024), in which case the caller may re-try to
//      receive the remaining messages.
//
func (datagramBased *DatagramBased) ReceiveMessages(ctx context.Context, messages []InputMessage, flags int, timeout int64) (int, error) {
	var _arg0 *C.GDatagramBased // out
	var _arg5 *C.GCancellable   // out
	var _arg1 *C.GInputMessage  // out
	var _arg2 C.guint
	var _arg3 C.gint    // out
	var _arg4 C.gint64  // out
	var _cret C.gint    // in
	var _cerr *C.GError // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.guint)(len(messages))
	_arg1 = (*C.GInputMessage)(C.calloc(C.size_t(len(messages)), C.size_t(C.sizeof_GInputMessage)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GInputMessage)(_arg1), len(messages))
		for i := range messages {
			out[i] = *(*C.GInputMessage)(gextras.StructNative(unsafe.Pointer((&messages[i]))))
		}
	}
	_arg3 = C.gint(flags)
	_arg4 = C.gint64(timeout)

	_cret = C.g_datagram_based_receive_messages(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(messages)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(timeout)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// SendMessages: send one or more data messages from datagram_based in one go.
//
// messages must point to an array of Message structs and num_messages must be
// the length of this array. Each Message contains an address to send the data
// to, and a pointer to an array of Vector structs to describe the buffers that
// the data to be sent for each message will be gathered from.
//
// flags modify how the message is sent. The commonly available arguments for
// this are available in the MsgFlags enum, but the values there are the same as
// the system values, and the flags are passed in as-is, so you can pass in
// system-specific flags too.
//
// The other members of Message are treated as described in its documentation.
//
// If timeout is negative the call will block until num_messages have been sent,
// cancellable is cancelled, or an error occurs.
//
// If timeout is 0 the call will send up to num_messages without blocking, or
// will return G_IO_ERROR_WOULD_BLOCK if there is no space to send messages.
//
// If timeout is positive the call will block on the same conditions as if
// timeout were negative. If the timeout is reached before any messages are
// sent, G_IO_ERROR_TIMED_OUT is returned, otherwise it will return the number
// of messages sent before timing out.
//
// To be notified when messages can be sent, wait for the G_IO_OUT condition.
// Note though that you may still receive G_IO_ERROR_WOULD_BLOCK from
// g_datagram_based_send_messages() even if you were previously notified of a
// G_IO_OUT condition. (On Windows in particular, this is very common due to the
// way the underlying APIs work.)
//
// If the connection is shut down or closed (by calling g_socket_close() or
// g_socket_shutdown() with shutdown_write set, if it’s a #GSocket, for
// example), all calls to this function will return G_IO_ERROR_CLOSED.
//
// On error -1 is returned and error is set accordingly. An error will only be
// returned if zero messages could be sent; otherwise the number of messages
// successfully sent before the error will be returned. If cancellable is
// cancelled, G_IO_ERROR_CANCELLED is returned as with any other error.
//
// The function takes the following parameters:
//
//    - ctx (optional): GCancellable.
//    - messages: array of Message structs.
//    - flags: int containing MsgFlags flags.
//    - timeout: maximum time (in microseconds) to wait, 0 to not block, or -1 to
//      block indefinitely.
//
// The function returns the following values:
//
//    - gint: number of messages sent, or -1 on error. Note that the number of
//      messages sent may be smaller than num_messages if timeout is zero or
//      positive, or if num_messages was larger than UIO_MAXIOV (1024), in which
//      case the caller may re-try to send the remaining messages.
//
func (datagramBased *DatagramBased) SendMessages(ctx context.Context, messages []OutputMessage, flags int, timeout int64) (int, error) {
	var _arg0 *C.GDatagramBased // out
	var _arg5 *C.GCancellable   // out
	var _arg1 *C.GOutputMessage // out
	var _arg2 C.guint
	var _arg3 C.gint    // out
	var _arg4 C.gint64  // out
	var _cret C.gint    // in
	var _cerr *C.GError // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.guint)(len(messages))
	_arg1 = (*C.GOutputMessage)(C.calloc(C.size_t(len(messages)), C.size_t(C.sizeof_GOutputMessage)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GOutputMessage)(_arg1), len(messages))
		for i := range messages {
			out[i] = *(*C.GOutputMessage)(gextras.StructNative(unsafe.Pointer((&messages[i]))))
		}
	}
	_arg3 = C.gint(flags)
	_arg4 = C.gint64(timeout)

	_cret = C.g_datagram_based_send_messages(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(messages)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(timeout)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// conditionCheck checks on the readiness of datagram_based to perform
// operations. The operations specified in condition are checked for and masked
// against the currently-satisfied conditions on datagram_based. The result is
// returned.
//
// G_IO_IN will be set in the return value if data is available to read with
// g_datagram_based_receive_messages(), or if the connection is closed remotely
// (EOS); and if the datagram_based has not been closed locally using some
// implementation-specific method (such as g_socket_close() or
// g_socket_shutdown() with shutdown_read set, if it’s a #GSocket).
//
// If the connection is shut down or closed (by calling g_socket_close() or
// g_socket_shutdown() with shutdown_read set, if it’s a #GSocket, for example),
// all calls to this function will return G_IO_ERROR_CLOSED.
//
// G_IO_OUT will be set if it is expected that at least one byte can be sent
// using g_datagram_based_send_messages() without blocking. It will not be set
// if the datagram_based has been closed locally.
//
// G_IO_HUP will be set if the connection has been closed locally.
//
// G_IO_ERR will be set if there was an asynchronous error in transmitting data
// previously enqueued using g_datagram_based_send_messages().
//
// Note that on Windows, it is possible for an operation to return
// G_IO_ERROR_WOULD_BLOCK even immediately after
// g_datagram_based_condition_check() has claimed that the Based is ready for
// writing. Rather than calling g_datagram_based_condition_check() and then
// writing to the Based if it succeeds, it is generally better to simply try
// writing right away, and try again later if the initial attempt returns
// G_IO_ERROR_WOULD_BLOCK.
//
// It is meaningless to specify G_IO_ERR or G_IO_HUP in condition; these
// conditions will always be set in the output if they are true. Apart from
// these flags, the output is guaranteed to be masked by condition.
//
// This call never blocks.
//
// The function takes the following parameters:
//
//    - condition mask to check.
//
// The function returns the following values:
//
//    - ioCondition mask of the current state.
//
func (datagramBased *DatagramBased) conditionCheck(condition glib.IOCondition) glib.IOCondition {
	gclass := (*C.GDatagramBasedInterface)(coreglib.PeekParentClass(datagramBased))
	fnarg := gclass.condition_check

	var _arg0 *C.GDatagramBased // out
	var _arg1 C.GIOCondition    // out
	var _cret C.GIOCondition    // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	_arg1 = C.GIOCondition(condition)

	_cret = C._gotk4_gio2_DatagramBased_virtual_condition_check(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(condition)

	var _ioCondition glib.IOCondition // out

	_ioCondition = glib.IOCondition(_cret)

	return _ioCondition
}

// conditionWait waits for up to timeout microseconds for condition to become
// true on datagram_based. If the condition is met, TRUE is returned.
//
// If cancellable is cancelled before the condition is met, or if timeout is
// reached before the condition is met, then FALSE is returned and error is set
// appropriately (G_IO_ERROR_CANCELLED or G_IO_ERROR_TIMED_OUT).
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//    - condition mask to wait for.
//    - timeout: maximum time (in microseconds) to wait, 0 to not block, or -1 to
//      block indefinitely.
//
func (datagramBased *DatagramBased) conditionWait(ctx context.Context, condition glib.IOCondition, timeout int64) error {
	gclass := (*C.GDatagramBasedInterface)(coreglib.PeekParentClass(datagramBased))
	fnarg := gclass.condition_wait

	var _arg0 *C.GDatagramBased // out
	var _arg3 *C.GCancellable   // out
	var _arg1 C.GIOCondition    // out
	var _arg2 C.gint64          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GIOCondition(condition)
	_arg2 = C.gint64(timeout)

	C._gotk4_gio2_DatagramBased_virtual_condition_wait(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(condition)
	runtime.KeepAlive(timeout)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// createSource creates a #GSource that can be attached to a Context to monitor
// for the availability of the specified condition on the Based. The #GSource
// keeps a reference to the datagram_based.
//
// The callback on the source is of the BasedSourceFunc type.
//
// It is meaningless to specify G_IO_ERR or G_IO_HUP in condition; these
// conditions will always be reported in the callback if they are true.
//
// If non-NULL, cancellable can be used to cancel the source, which will cause
// the source to trigger, reporting the current condition (which is likely 0
// unless cancellation happened at the same time as a condition change). You can
// check for this in the callback using g_cancellable_is_cancelled().
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//    - condition mask to monitor.
//
// The function returns the following values:
//
//    - source: newly allocated #GSource.
//
func (datagramBased *DatagramBased) createSource(ctx context.Context, condition glib.IOCondition) *glib.Source {
	gclass := (*C.GDatagramBasedInterface)(coreglib.PeekParentClass(datagramBased))
	fnarg := gclass.create_source

	var _arg0 *C.GDatagramBased // out
	var _arg2 *C.GCancellable   // out
	var _arg1 C.GIOCondition    // out
	var _cret *C.GSource        // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GIOCondition(condition)

	_cret = C._gotk4_gio2_DatagramBased_virtual_create_source(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(condition)

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

// receiveMessages: receive one or more data messages from datagram_based in one
// go.
//
// messages must point to an array of Message structs and num_messages must be
// the length of this array. Each Message contains a pointer to an array of
// Vector structs describing the buffers that the data received in each message
// will be written to.
//
// flags modify how all messages are received. The commonly available arguments
// for this are available in the MsgFlags enum, but the values there are the
// same as the system values, and the flags are passed in as-is, so you can pass
// in system-specific flags too. These flags affect the overall receive
// operation. Flags affecting individual messages are returned in Message.flags.
//
// The other members of Message are treated as described in its documentation.
//
// If timeout is negative the call will block until num_messages have been
// received, the connection is closed remotely (EOS), cancellable is cancelled,
// or an error occurs.
//
// If timeout is 0 the call will return up to num_messages without blocking, or
// G_IO_ERROR_WOULD_BLOCK if no messages are queued in the operating system to
// be received.
//
// If timeout is positive the call will block on the same conditions as if
// timeout were negative. If the timeout is reached before any messages are
// received, G_IO_ERROR_TIMED_OUT is returned, otherwise it will return the
// number of messages received before timing out. (Note: This is effectively the
// behaviour of MSG_WAITFORONE with recvmmsg().)
//
// To be notified when messages are available, wait for the G_IO_IN condition.
// Note though that you may still receive G_IO_ERROR_WOULD_BLOCK from
// g_datagram_based_receive_messages() even if you were previously notified of a
// G_IO_IN condition.
//
// If the remote peer closes the connection, any messages queued in the
// underlying receive buffer will be returned, and subsequent calls to
// g_datagram_based_receive_messages() will return 0 (with no error set).
//
// If the connection is shut down or closed (by calling g_socket_close() or
// g_socket_shutdown() with shutdown_read set, if it’s a #GSocket, for example),
// all calls to this function will return G_IO_ERROR_CLOSED.
//
// On error -1 is returned and error is set accordingly. An error will only be
// returned if zero messages could be received; otherwise the number of messages
// successfully received before the error will be returned. If cancellable is
// cancelled, G_IO_ERROR_CANCELLED is returned as with any other error.
//
// The function takes the following parameters:
//
//    - ctx (optional): GCancellable.
//    - messages: array of Message structs.
//    - flags: int containing MsgFlags flags for the overall operation.
//    - timeout: maximum time (in microseconds) to wait, 0 to not block, or -1 to
//      block indefinitely.
//
// The function returns the following values:
//
//    - gint: number of messages received, or -1 on error. Note that the number
//      of messages received may be smaller than num_messages if timeout is zero
//      or positive, if the peer closed the connection, or if num_messages was
//      larger than UIO_MAXIOV (1024), in which case the caller may re-try to
//      receive the remaining messages.
//
func (datagramBased *DatagramBased) receiveMessages(ctx context.Context, messages []InputMessage, flags int, timeout int64) (int, error) {
	gclass := (*C.GDatagramBasedInterface)(coreglib.PeekParentClass(datagramBased))
	fnarg := gclass.receive_messages

	var _arg0 *C.GDatagramBased // out
	var _arg5 *C.GCancellable   // out
	var _arg1 *C.GInputMessage  // out
	var _arg2 C.guint
	var _arg3 C.gint    // out
	var _arg4 C.gint64  // out
	var _cret C.gint    // in
	var _cerr *C.GError // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.guint)(len(messages))
	_arg1 = (*C.GInputMessage)(C.calloc(C.size_t(len(messages)), C.size_t(C.sizeof_GInputMessage)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GInputMessage)(_arg1), len(messages))
		for i := range messages {
			out[i] = *(*C.GInputMessage)(gextras.StructNative(unsafe.Pointer((&messages[i]))))
		}
	}
	_arg3 = C.gint(flags)
	_arg4 = C.gint64(timeout)

	_cret = C._gotk4_gio2_DatagramBased_virtual_receive_messages(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(messages)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(timeout)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// sendMessages: send one or more data messages from datagram_based in one go.
//
// messages must point to an array of Message structs and num_messages must be
// the length of this array. Each Message contains an address to send the data
// to, and a pointer to an array of Vector structs to describe the buffers that
// the data to be sent for each message will be gathered from.
//
// flags modify how the message is sent. The commonly available arguments for
// this are available in the MsgFlags enum, but the values there are the same as
// the system values, and the flags are passed in as-is, so you can pass in
// system-specific flags too.
//
// The other members of Message are treated as described in its documentation.
//
// If timeout is negative the call will block until num_messages have been sent,
// cancellable is cancelled, or an error occurs.
//
// If timeout is 0 the call will send up to num_messages without blocking, or
// will return G_IO_ERROR_WOULD_BLOCK if there is no space to send messages.
//
// If timeout is positive the call will block on the same conditions as if
// timeout were negative. If the timeout is reached before any messages are
// sent, G_IO_ERROR_TIMED_OUT is returned, otherwise it will return the number
// of messages sent before timing out.
//
// To be notified when messages can be sent, wait for the G_IO_OUT condition.
// Note though that you may still receive G_IO_ERROR_WOULD_BLOCK from
// g_datagram_based_send_messages() even if you were previously notified of a
// G_IO_OUT condition. (On Windows in particular, this is very common due to the
// way the underlying APIs work.)
//
// If the connection is shut down or closed (by calling g_socket_close() or
// g_socket_shutdown() with shutdown_write set, if it’s a #GSocket, for
// example), all calls to this function will return G_IO_ERROR_CLOSED.
//
// On error -1 is returned and error is set accordingly. An error will only be
// returned if zero messages could be sent; otherwise the number of messages
// successfully sent before the error will be returned. If cancellable is
// cancelled, G_IO_ERROR_CANCELLED is returned as with any other error.
//
// The function takes the following parameters:
//
//    - ctx (optional): GCancellable.
//    - messages: array of Message structs.
//    - flags: int containing MsgFlags flags.
//    - timeout: maximum time (in microseconds) to wait, 0 to not block, or -1 to
//      block indefinitely.
//
// The function returns the following values:
//
//    - gint: number of messages sent, or -1 on error. Note that the number of
//      messages sent may be smaller than num_messages if timeout is zero or
//      positive, or if num_messages was larger than UIO_MAXIOV (1024), in which
//      case the caller may re-try to send the remaining messages.
//
func (datagramBased *DatagramBased) sendMessages(ctx context.Context, messages []OutputMessage, flags int, timeout int64) (int, error) {
	gclass := (*C.GDatagramBasedInterface)(coreglib.PeekParentClass(datagramBased))
	fnarg := gclass.send_messages

	var _arg0 *C.GDatagramBased // out
	var _arg5 *C.GCancellable   // out
	var _arg1 *C.GOutputMessage // out
	var _arg2 C.guint
	var _arg3 C.gint    // out
	var _arg4 C.gint64  // out
	var _cret C.gint    // in
	var _cerr *C.GError // in

	_arg0 = (*C.GDatagramBased)(unsafe.Pointer(coreglib.InternObject(datagramBased).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg2 = (C.guint)(len(messages))
	_arg1 = (*C.GOutputMessage)(C.calloc(C.size_t(len(messages)), C.size_t(C.sizeof_GOutputMessage)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GOutputMessage)(_arg1), len(messages))
		for i := range messages {
			out[i] = *(*C.GOutputMessage)(gextras.StructNative(unsafe.Pointer((&messages[i]))))
		}
	}
	_arg3 = C.gint(flags)
	_arg4 = C.gint64(timeout)

	_cret = C._gotk4_gio2_DatagramBased_virtual_send_messages(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)
	runtime.KeepAlive(datagramBased)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(messages)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(timeout)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// DatagramBasedInterface provides an interface for socket-like objects which
// have datagram semantics, following the Berkeley sockets API. The interface
// methods are thin wrappers around the corresponding virtual methods, and no
// pre-processing of inputs is implemented — so implementations of this API must
// handle all functionality documented in the interface methods.
//
// An instance of this type is always passed by reference.
type DatagramBasedInterface struct {
	*datagramBasedInterface
}

// datagramBasedInterface is the struct that's finalized.
type datagramBasedInterface struct {
	native *C.GDatagramBasedInterface
}