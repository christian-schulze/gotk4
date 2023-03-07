// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_Notebook_ConnectCreateWindow
func _gotk4_gtk4_Notebook_ConnectCreateWindow(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guintptr) (cret *C.GtkNotebook) {
	var f func(page Widgetter) (notebook *Notebook)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(page Widgetter) (notebook *Notebook))
	}

	var _page Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_page = rv
	}

	notebook := f(_page)

	var _ *Notebook

	cret = (*C.GtkNotebook)(unsafe.Pointer(coreglib.InternObject(notebook).Native()))

	return cret
}

//export _gotk4_gtk4_Notebook_ConnectPageAdded
func _gotk4_gtk4_Notebook_ConnectPageAdded(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guint, arg3 C.guintptr) {
	var f func(child Widgetter, pageNum uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(child Widgetter, pageNum uint))
	}

	var _child Widgetter // out
	var _pageNum uint    // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}
	_pageNum = uint(arg2)

	f(_child, _pageNum)
}

//export _gotk4_gtk4_Notebook_ConnectPageRemoved
func _gotk4_gtk4_Notebook_ConnectPageRemoved(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guint, arg3 C.guintptr) {
	var f func(child Widgetter, pageNum uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(child Widgetter, pageNum uint))
	}

	var _child Widgetter // out
	var _pageNum uint    // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}
	_pageNum = uint(arg2)

	f(_child, _pageNum)
}

//export _gotk4_gtk4_Notebook_ConnectPageReordered
func _gotk4_gtk4_Notebook_ConnectPageReordered(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guint, arg3 C.guintptr) {
	var f func(child Widgetter, pageNum uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(child Widgetter, pageNum uint))
	}

	var _child Widgetter // out
	var _pageNum uint    // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_child = rv
	}
	_pageNum = uint(arg2)

	f(_child, _pageNum)
}

//export _gotk4_gtk4_Notebook_ConnectSwitchPage
func _gotk4_gtk4_Notebook_ConnectSwitchPage(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guint, arg3 C.guintptr) {
	var f func(page Widgetter, pageNum uint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(page Widgetter, pageNum uint))
	}

	var _page Widgetter // out
	var _pageNum uint   // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_page = rv
	}
	_pageNum = uint(arg2)

	f(_page, _pageNum)
}