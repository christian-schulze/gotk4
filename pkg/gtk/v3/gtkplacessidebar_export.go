// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_PlacesSidebar_ConnectDragActionAsk
func _gotk4_gtk3_PlacesSidebar_ConnectDragActionAsk(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) (cret C.gint) {
	var f func(actions int) (gint int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(actions int) (gint int))
	}

	var _actions int // out

	_actions = int(arg1)

	gint := f(_actions)

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_gtk3_PlacesSidebar_ConnectMount
func _gotk4_gtk3_PlacesSidebar_ConnectMount(arg0 C.gpointer, arg1 *C.GMountOperation, arg2 C.guintptr) {
	var f func(mountOperation *gio.MountOperation)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(mountOperation *gio.MountOperation))
	}

	var _mountOperation *gio.MountOperation // out

	{
		obj := coreglib.Take(unsafe.Pointer(arg1))
		_mountOperation = &gio.MountOperation{
			Object: obj,
		}
	}

	f(_mountOperation)
}

//export _gotk4_gtk3_PlacesSidebar_ConnectOpenLocation
func _gotk4_gtk3_PlacesSidebar_ConnectOpenLocation(arg0 C.gpointer, arg1 *C.GFile, arg2 C.GtkPlacesOpenFlags, arg3 C.guintptr) {
	var f func(location gio.Filer, openFlags PlacesOpenFlags)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(location gio.Filer, openFlags PlacesOpenFlags))
	}

	var _location gio.Filer        // out
	var _openFlags PlacesOpenFlags // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.Filer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.Filer)
			return ok
		})
		rv, ok := casted.(gio.Filer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
		}
		_location = rv
	}
	_openFlags = PlacesOpenFlags(arg2)

	f(_location, _openFlags)
}

//export _gotk4_gtk3_PlacesSidebar_ConnectPopulatePopup
func _gotk4_gtk3_PlacesSidebar_ConnectPopulatePopup(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 *C.GFile, arg3 *C.GVolume, arg4 C.guintptr) {
	var f func(container Widgetter, selectedItem gio.Filer, selectedVolume gio.Volumer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(container Widgetter, selectedItem gio.Filer, selectedVolume gio.Volumer))
	}

	var _container Widgetter        // out
	var _selectedItem gio.Filer     // out
	var _selectedVolume gio.Volumer // out

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
		_container = rv
	}
	if arg2 != nil {
		{
			objptr := unsafe.Pointer(arg2)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.Filer)
				return ok
			})
			rv, ok := casted.(gio.Filer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Filer")
			}
			_selectedItem = rv
		}
	}
	if arg3 != nil {
		{
			objptr := unsafe.Pointer(arg3)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.Volumer)
				return ok
			})
			rv, ok := casted.(gio.Volumer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Volumer")
			}
			_selectedVolume = rv
		}
	}

	f(_container, _selectedItem, _selectedVolume)
}

//export _gotk4_gtk3_PlacesSidebar_ConnectShowConnectToServer
func _gotk4_gtk3_PlacesSidebar_ConnectShowConnectToServer(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_PlacesSidebar_ConnectShowEnterLocation
func _gotk4_gtk3_PlacesSidebar_ConnectShowEnterLocation(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_PlacesSidebar_ConnectShowErrorMessage
func _gotk4_gtk3_PlacesSidebar_ConnectShowErrorMessage(arg0 C.gpointer, arg1 *C.gchar, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(primary, secondary string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(primary, secondary string))
	}

	var _primary string   // out
	var _secondary string // out

	_primary = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_secondary = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_primary, _secondary)
}

//export _gotk4_gtk3_PlacesSidebar_ConnectShowOtherLocations
func _gotk4_gtk3_PlacesSidebar_ConnectShowOtherLocations(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gtk3_PlacesSidebar_ConnectShowOtherLocationsWithFlags
func _gotk4_gtk3_PlacesSidebar_ConnectShowOtherLocationsWithFlags(arg0 C.gpointer, arg1 C.GtkPlacesOpenFlags, arg2 C.guintptr) {
	var f func(openFlags PlacesOpenFlags)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(openFlags PlacesOpenFlags))
	}

	var _openFlags PlacesOpenFlags // out

	_openFlags = PlacesOpenFlags(arg1)

	f(_openFlags)
}

//export _gotk4_gtk3_PlacesSidebar_ConnectShowStarredLocation
func _gotk4_gtk3_PlacesSidebar_ConnectShowStarredLocation(arg0 C.gpointer, arg1 C.GtkPlacesOpenFlags, arg2 C.guintptr) {
	var f func(openFlags PlacesOpenFlags)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(openFlags PlacesOpenFlags))
	}

	var _openFlags PlacesOpenFlags // out

	_openFlags = PlacesOpenFlags(arg1)

	f(_openFlags)
}

//export _gotk4_gtk3_PlacesSidebar_ConnectUnmount
func _gotk4_gtk3_PlacesSidebar_ConnectUnmount(arg0 C.gpointer, arg1 *C.GMountOperation, arg2 C.guintptr) {
	var f func(mountOperation *gio.MountOperation)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(mountOperation *gio.MountOperation))
	}

	var _mountOperation *gio.MountOperation // out

	{
		obj := coreglib.Take(unsafe.Pointer(arg1))
		_mountOperation = &gio.MountOperation{
			Object: obj,
		}
	}

	f(_mountOperation)
}