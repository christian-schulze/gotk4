// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_SocketListener_ConnectEvent(gpointer, GSocketListenerEvent, GSocket*, guintptr);
// extern void _gotk4_gio2_SocketListenerClass_event(GSocketListener*, GSocketListenerEvent, GSocket*);
// extern void _gotk4_gio2_SocketListenerClass_changed(GSocketListener*);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// void _gotk4_gio2_SocketListener_virtual_changed(void* fnptr, GSocketListener* arg0) {
//   ((void (*)(GSocketListener*))(fnptr))(arg0);
// };
// void _gotk4_gio2_SocketListener_virtual_event(void* fnptr, GSocketListener* arg0, GSocketListenerEvent arg1, GSocket* arg2) {
//   ((void (*)(GSocketListener*, GSocketListenerEvent, GSocket*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeSocketListener = coreglib.Type(C.g_socket_listener_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSocketListener, F: marshalSocketListener},
	})
}

// SocketListenerOverrides contains methods that are overridable.
type SocketListenerOverrides struct {
	Changed func()
	// The function takes the following parameters:
	//
	//    - event
	//    - socket
	//
	Event func(event SocketListenerEvent, socket *Socket)
}

func defaultSocketListenerOverrides(v *SocketListener) SocketListenerOverrides {
	return SocketListenerOverrides{
		Changed: v.changed,
		Event:   v.event,
	}
}

// SocketListener is an object that keeps track of a set of server sockets and
// helps you accept sockets from any of the socket, either sync or async.
//
// Add addresses and ports to listen on using g_socket_listener_add_address()
// and g_socket_listener_add_inet_port(). These will be listened on until
// g_socket_listener_close() is called. Dropping your final reference to the
// Listener will not cause g_socket_listener_close() to be called implicitly, as
// some references to the Listener may be held internally.
//
// If you want to implement a network server, also look at Service and
// SocketService which are subclasses of Listener that make this even easier.
type SocketListener struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*SocketListener)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SocketListener, *SocketListenerClass, SocketListenerOverrides](
		GTypeSocketListener,
		initSocketListenerClass,
		wrapSocketListener,
		defaultSocketListenerOverrides,
	)
}

func initSocketListenerClass(gclass unsafe.Pointer, overrides SocketListenerOverrides, classInitFunc func(*SocketListenerClass)) {
	pclass := (*C.GSocketListenerClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeSocketListener))))

	if overrides.Changed != nil {
		pclass.changed = (*[0]byte)(C._gotk4_gio2_SocketListenerClass_changed)
	}

	if overrides.Event != nil {
		pclass.event = (*[0]byte)(C._gotk4_gio2_SocketListenerClass_event)
	}

	if classInitFunc != nil {
		class := (*SocketListenerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSocketListener(obj *coreglib.Object) *SocketListener {
	return &SocketListener{
		Object: obj,
	}
}

func marshalSocketListener(p uintptr) (interface{}, error) {
	return wrapSocketListener(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEvent is emitted when listener's activity on socket changes state.
// Note that when listener is used to listen on both IPv4 and IPv6, a separate
// set of signals will be emitted for each, and the order they happen in is
// undefined.
func (listener *SocketListener) ConnectEvent(f func(event SocketListenerEvent, socket *Socket)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(listener, "event", false, unsafe.Pointer(C._gotk4_gio2_SocketListener_ConnectEvent), f)
}

// NewSocketListener creates a new Listener with no sockets to listen for. New
// listeners can be added with e.g. g_socket_listener_add_address() or
// g_socket_listener_add_inet_port().
//
// The function returns the following values:
//
//    - socketListener: new Listener.
//
func NewSocketListener() *SocketListener {
	var _cret *C.GSocketListener // in

	_cret = C.g_socket_listener_new()

	var _socketListener *SocketListener // out

	_socketListener = wrapSocketListener(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _socketListener
}

// Accept blocks waiting for a client to connect to any of the sockets added to
// the listener. Returns a Connection for the socket that was accepted.
//
// If source_object is not NULL it will be filled out with the source object
// specified when the corresponding socket or address was added to the listener.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
// The function returns the following values:
//
//    - sourceObject (optional): location where #GObject pointer will be stored,
//      or NULL.
//    - socketConnection on success, NULL on error.
//
func (listener *SocketListener) Accept(ctx context.Context) (*coreglib.Object, *SocketConnection, error) {
	var _arg0 *C.GSocketListener   // out
	var _arg2 *C.GCancellable      // out
	var _arg1 *C.GObject           // in
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.g_socket_listener_accept(_arg0, &_arg1, _arg2, &_cerr)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(ctx)

	var _sourceObject *coreglib.Object      // out
	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	if _arg1 != nil {
		_sourceObject = coreglib.Take(unsafe.Pointer(_arg1))
	}
	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _sourceObject, _socketConnection, _goerr
}

// AcceptAsync: this is the asynchronous version of g_socket_listener_accept().
//
// When the operation is finished callback will be called. You can then call
// g_socket_listener_accept_socket() to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - callback (optional): ReadyCallback.
//
func (listener *SocketListener) AcceptAsync(ctx context.Context, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketListener    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_socket_listener_accept_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// AcceptFinish finishes an async accept operation. See
// g_socket_listener_accept_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - sourceObject (optional): optional #GObject identifying this source.
//    - socketConnection on success, NULL on error.
//
func (listener *SocketListener) AcceptFinish(result AsyncResulter) (*coreglib.Object, *SocketConnection, error) {
	var _arg0 *C.GSocketListener   // out
	var _arg1 *C.GAsyncResult      // out
	var _arg2 *C.GObject           // in
	var _cret *C.GSocketConnection // in
	var _cerr *C.GError            // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.g_socket_listener_accept_finish(_arg0, _arg1, &_arg2, &_cerr)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(result)

	var _sourceObject *coreglib.Object      // out
	var _socketConnection *SocketConnection // out
	var _goerr error                        // out

	if _arg2 != nil {
		_sourceObject = coreglib.Take(unsafe.Pointer(_arg2))
	}
	_socketConnection = wrapSocketConnection(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _sourceObject, _socketConnection, _goerr
}

// AcceptSocket blocks waiting for a client to connect to any of the sockets
// added to the listener. Returns the #GSocket that was accepted.
//
// If you want to accept the high-level Connection, not a #GSocket, which is
// often the case, then you should use g_socket_listener_accept() instead.
//
// If source_object is not NULL it will be filled out with the source object
// specified when the corresponding socket or address was added to the listener.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be returned.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//
// The function returns the following values:
//
//    - sourceObject (optional): location where #GObject pointer will be stored,
//      or NULL.
//    - socket on success, NULL on error.
//
func (listener *SocketListener) AcceptSocket(ctx context.Context) (*coreglib.Object, *Socket, error) {
	var _arg0 *C.GSocketListener // out
	var _arg2 *C.GCancellable    // out
	var _arg1 *C.GObject         // in
	var _cret *C.GSocket         // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.g_socket_listener_accept_socket(_arg0, &_arg1, _arg2, &_cerr)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(ctx)

	var _sourceObject *coreglib.Object // out
	var _socket *Socket                // out
	var _goerr error                   // out

	if _arg1 != nil {
		_sourceObject = coreglib.Take(unsafe.Pointer(_arg1))
	}
	_socket = wrapSocket(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _sourceObject, _socket, _goerr
}

// AcceptSocketAsync: this is the asynchronous version of
// g_socket_listener_accept_socket().
//
// When the operation is finished callback will be called. You can then call
// g_socket_listener_accept_socket_finish() to get the result of the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional) or NULL.
//    - callback (optional): ReadyCallback.
//
func (listener *SocketListener) AcceptSocketAsync(ctx context.Context, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketListener    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_socket_listener_accept_socket_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// AcceptSocketFinish finishes an async accept operation. See
// g_socket_listener_accept_socket_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - sourceObject (optional): optional #GObject identifying this source.
//    - socket on success, NULL on error.
//
func (listener *SocketListener) AcceptSocketFinish(result AsyncResulter) (*coreglib.Object, *Socket, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GAsyncResult    // out
	var _arg2 *C.GObject         // in
	var _cret *C.GSocket         // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_cret = C.g_socket_listener_accept_socket_finish(_arg0, _arg1, &_arg2, &_cerr)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(result)

	var _sourceObject *coreglib.Object // out
	var _socket *Socket                // out
	var _goerr error                   // out

	if _arg2 != nil {
		_sourceObject = coreglib.Take(unsafe.Pointer(_arg2))
	}
	_socket = wrapSocket(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _sourceObject, _socket, _goerr
}

// AddAddress creates a socket of type type and protocol protocol, binds it to
// address and adds it to the set of sockets we're accepting sockets from.
//
// Note that adding an IPv6 address, depending on the platform, may or may not
// result in a listener that also accepts IPv4 connections. For more
// deterministic behavior, see g_socket_listener_add_inet_port().
//
// source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// If successful and effective_address is non-NULL then it will be set to the
// address that the binding actually occurred at. This is helpful for
// determining the port number that was used for when requesting a binding to
// port 0 (ie: "any port"). This address, if requested, belongs to the caller
// and must be freed.
//
// Call g_socket_listener_close() to stop listening on address; this will not be
// done automatically when you drop your final reference to listener, as
// references may be held internally.
//
// The function takes the following parameters:
//
//    - address: Address.
//    - typ: Type.
//    - protocol: Protocol.
//    - sourceObject (optional): optional #GObject identifying this source.
//
// The function returns the following values:
//
//    - effectiveAddress (optional): location to store the address that was bound
//      to, or NULL.
//
func (listener *SocketListener) AddAddress(address SocketAddresser, typ SocketType, protocol SocketProtocol, sourceObject *coreglib.Object) (SocketAddresser, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GSocketAddress  // out
	var _arg2 C.GSocketType      // out
	var _arg3 C.GSocketProtocol  // out
	var _arg4 *C.GObject         // out
	var _arg5 *C.GSocketAddress  // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	_arg1 = (*C.GSocketAddress)(unsafe.Pointer(coreglib.InternObject(address).Native()))
	_arg2 = C.GSocketType(typ)
	_arg3 = C.GSocketProtocol(protocol)
	if sourceObject != nil {
		_arg4 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))
	}

	C.g_socket_listener_add_address(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_cerr)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(address)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(sourceObject)

	var _effectiveAddress SocketAddresser // out
	var _goerr error                      // out

	if _arg5 != nil {
		{
			objptr := unsafe.Pointer(_arg5)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(SocketAddresser)
				return ok
			})
			rv, ok := casted.(SocketAddresser)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddresser")
			}
			_effectiveAddress = rv
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _effectiveAddress, _goerr
}

// AddAnyInetPort listens for TCP connections on any available port number for
// both IPv6 and IPv4 (if each is available).
//
// This is useful if you need to have a socket for incoming connections but
// don't care about the specific port number.
//
// source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// The function takes the following parameters:
//
//    - sourceObject (optional): optional #GObject identifying this source.
//
// The function returns the following values:
//
//    - guint16: port number, or 0 in case of failure.
//
func (listener *SocketListener) AddAnyInetPort(sourceObject *coreglib.Object) (uint16, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GObject         // out
	var _cret C.guint16          // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	if sourceObject != nil {
		_arg1 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))
	}

	_cret = C.g_socket_listener_add_any_inet_port(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(sourceObject)

	var _guint16 uint16 // out
	var _goerr error    // out

	_guint16 = uint16(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint16, _goerr
}

// AddInetPort: helper function for g_socket_listener_add_address() that creates
// a TCP/IP socket listening on IPv4 and IPv6 (if supported) on the specified
// port on all interfaces.
//
// source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// Call g_socket_listener_close() to stop listening on port; this will not be
// done automatically when you drop your final reference to listener, as
// references may be held internally.
//
// The function takes the following parameters:
//
//    - port: IP port number (non-zero).
//    - sourceObject (optional): optional #GObject identifying this source.
//
func (listener *SocketListener) AddInetPort(port uint16, sourceObject *coreglib.Object) error {
	var _arg0 *C.GSocketListener // out
	var _arg1 C.guint16          // out
	var _arg2 *C.GObject         // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	_arg1 = C.guint16(port)
	if sourceObject != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))
	}

	C.g_socket_listener_add_inet_port(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(port)
	runtime.KeepAlive(sourceObject)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// AddSocket adds socket to the set of sockets that we try to accept new clients
// from. The socket must be bound to a local address and listened to.
//
// source_object will be passed out in the various calls to accept to identify
// this particular source, which is useful if you're listening on multiple
// addresses and do different things depending on what address is connected to.
//
// The socket will not be automatically closed when the listener is finalized
// unless the listener held the final reference to the socket. Before GLib 2.42,
// the socket was automatically closed on finalization of the listener, even if
// references to it were held elsewhere.
//
// The function takes the following parameters:
//
//    - socket: listening #GSocket.
//    - sourceObject (optional): optional #GObject identifying this source.
//
func (listener *SocketListener) AddSocket(socket *Socket, sourceObject *coreglib.Object) error {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GSocket         // out
	var _arg2 *C.GObject         // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	_arg1 = (*C.GSocket)(unsafe.Pointer(coreglib.InternObject(socket).Native()))
	if sourceObject != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))
	}

	C.g_socket_listener_add_socket(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(socket)
	runtime.KeepAlive(sourceObject)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Close closes all the sockets in the listener.
func (listener *SocketListener) Close() {
	var _arg0 *C.GSocketListener // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))

	C.g_socket_listener_close(_arg0)
	runtime.KeepAlive(listener)
}

// SetBacklog sets the listen backlog on the sockets in the listener. This must
// be called before adding any sockets, addresses or ports to the Listener (for
// example, by calling g_socket_listener_add_inet_port()) to be effective.
//
// See g_socket_set_listen_backlog() for details.
//
// The function takes the following parameters:
//
//    - listenBacklog: integer.
//
func (listener *SocketListener) SetBacklog(listenBacklog int) {
	var _arg0 *C.GSocketListener // out
	var _arg1 C.int              // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	_arg1 = C.int(listenBacklog)

	C.g_socket_listener_set_backlog(_arg0, _arg1)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(listenBacklog)
}

func (listener *SocketListener) changed() {
	gclass := (*C.GSocketListenerClass)(coreglib.PeekParentClass(listener))
	fnarg := gclass.changed

	var _arg0 *C.GSocketListener // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))

	C._gotk4_gio2_SocketListener_virtual_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(listener)
}

// The function takes the following parameters:
//
//    - event
//    - socket
//
func (listener *SocketListener) event(event SocketListenerEvent, socket *Socket) {
	gclass := (*C.GSocketListenerClass)(coreglib.PeekParentClass(listener))
	fnarg := gclass.event

	var _arg0 *C.GSocketListener     // out
	var _arg1 C.GSocketListenerEvent // out
	var _arg2 *C.GSocket             // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(coreglib.InternObject(listener).Native()))
	_arg1 = C.GSocketListenerEvent(event)
	_arg2 = (*C.GSocket)(unsafe.Pointer(coreglib.InternObject(socket).Native()))

	C._gotk4_gio2_SocketListener_virtual_event(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(listener)
	runtime.KeepAlive(event)
	runtime.KeepAlive(socket)
}