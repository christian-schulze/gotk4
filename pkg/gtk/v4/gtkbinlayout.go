// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_bin_layout_get_type()), F: marshalBinLayout},
	})
}

// BinLayout: gtkBinLayout is a LayoutManager subclass useful for create "bins"
// of widgets. GtkBinLayout will stack each child of a widget on top of each
// other, using the Widget:hexpand, Widget:vexpand, Widget:halign, and
// Widget:valign properties of each child to determine where they should be
// positioned.
type BinLayout interface {
	LayoutManager
}

// binLayout implements the BinLayout interface.
type binLayout struct {
	LayoutManager
}

var _ BinLayout = (*binLayout)(nil)

// WrapBinLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapBinLayout(obj *externglib.Object) BinLayout {
	return BinLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalBinLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBinLayout(obj), nil
}

// NewBinLayout constructs a class BinLayout.
func NewBinLayout() BinLayout {
	var cret C.GtkBinLayout
	var goret1 BinLayout

	cret = C.gtk_bin_layout_new()

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(BinLayout)

	return goret1
}