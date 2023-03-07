// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

//export _gotk4_gdk4_Display_ConnectClosed
func _gotk4_gdk4_Display_ConnectClosed(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(isError bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(isError bool))
	}

	var _isError bool // out

	if arg1 != 0 {
		_isError = true
	}

	f(_isError)
}

//export _gotk4_gdk4_Display_ConnectOpened
func _gotk4_gdk4_Display_ConnectOpened(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gdk4_Display_ConnectSeatAdded
func _gotk4_gdk4_Display_ConnectSeatAdded(arg0 C.gpointer, arg1 *C.GdkSeat, arg2 C.guintptr) {
	var f func(seat Seater)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(seat Seater))
	}

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}

	f(_seat)
}

//export _gotk4_gdk4_Display_ConnectSeatRemoved
func _gotk4_gdk4_Display_ConnectSeatRemoved(arg0 C.gpointer, arg1 *C.GdkSeat, arg2 C.guintptr) {
	var f func(seat Seater)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(seat Seater))
	}

	var _seat Seater // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gdk.Seater is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Seater)
			return ok
		})
		rv, ok := casted.(Seater)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Seater")
		}
		_seat = rv
	}

	f(_seat)
}

//export _gotk4_gdk4_Display_ConnectSettingChanged
func _gotk4_gdk4_Display_ConnectSettingChanged(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(setting string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(setting string))
	}

	var _setting string // out

	_setting = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_setting)
}