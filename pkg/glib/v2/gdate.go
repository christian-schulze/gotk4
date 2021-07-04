// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_date_get_type()), F: marshalDate},
	})
}

// Date represents a day between January 1, Year 1 and a few thousand years in
// the future. None of its members should be accessed directly.
//
// If the #GDate-struct is obtained from g_date_new(), it will be safe to mutate
// but invalid and thus not safe for calendrical computations.
//
// If it's declared on the stack, it will contain garbage so must be initialized
// with g_date_clear(). g_date_clear() makes the date invalid but safe. An
// invalid date doesn't represent a day, it's "empty." A date becomes valid
// after you set it to a Julian day or you set a day, month, and year.
type Date C.GDate

// WrapDate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDate(ptr unsafe.Pointer) *Date {
	return (*Date)(ptr)
}

func marshalDate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Date)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *Date) Native() unsafe.Pointer {
	return unsafe.Pointer(d)
}