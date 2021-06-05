// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_fixed_layout_get_type()), F: marshalFixedLayout},
		{T: externglib.Type(C.gtk_fixed_layout_child_get_type()), F: marshalFixedLayoutChild},
	})
}

// FixedLayout is a layout manager which can place child widgets at fixed
// positions, and with fixed sizes.
//
// Most applications should never use this layout manager; fixed positioning and
// sizing requires constant recalculations on where children need to be
// positioned and sized. Other layout managers perform this kind of work
// internally so that application developers don't need to do it. Specifically,
// widgets positioned in a fixed layout manager will need to take into account:
//
// - Themes, which may change widget sizes.
//
// - Fonts other than the one you used to write the app will of course change
// the size of widgets containing text; keep in mind that users may use a larger
// font because of difficulty reading the default, or they may be using a
// different OS that provides different fonts.
//
// - Translation of text into other languages changes its size. Also, display of
// non-English text will use a different font in many cases.
//
// In addition, FixedLayout does not pay attention to text direction and thus
// may produce unwanted results if your app is run under right-to-left languages
// such as Hebrew or Arabic. That is: normally GTK will order containers
// appropriately depending on the text direction, e.g. to put labels to the
// right of the thing they label when using an RTL language; FixedLayout won't
// be able to do that for you.
//
// Finally, fixed positioning makes it kind of annoying to add/remove GUI
// elements, since you have to reposition all the other elements. This is a
// long-term maintenance problem for your application.
type FixedLayout interface {
	LayoutManager
}

// fixedLayout implements the FixedLayout interface.
type fixedLayout struct {
	LayoutManager
}

var _ FixedLayout = (*fixedLayout)(nil)

// WrapFixedLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapFixedLayout(obj *externglib.Object) FixedLayout {
	return FixedLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalFixedLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFixedLayout(obj), nil
}

// NewFixedLayout constructs a class FixedLayout.
func NewFixedLayout() FixedLayout {
	var cret C.GtkFixedLayout
	var goret1 FixedLayout

	cret = C.gtk_fixed_layout_new()

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(FixedLayout)

	return goret1
}

type FixedLayoutChild interface {
	LayoutChild

	// Transform retrieves the transformation of the child of a FixedLayout.
	Transform() *gsk.Transform
	// SetTransform sets the transformation of the child of a FixedLayout.
	SetTransform(transform *gsk.Transform)
}

// fixedLayoutChild implements the FixedLayoutChild interface.
type fixedLayoutChild struct {
	LayoutChild
}

var _ FixedLayoutChild = (*fixedLayoutChild)(nil)

// WrapFixedLayoutChild wraps a GObject to the right type. It is
// primarily used internally.
func WrapFixedLayoutChild(obj *externglib.Object) FixedLayoutChild {
	return FixedLayoutChild{
		LayoutChild: WrapLayoutChild(obj),
	}
}

func marshalFixedLayoutChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFixedLayoutChild(obj), nil
}

// Transform retrieves the transformation of the child of a FixedLayout.
func (c fixedLayoutChild) Transform() *gsk.Transform {
	var arg0 *C.GtkFixedLayoutChild

	arg0 = (*C.GtkFixedLayoutChild)(unsafe.Pointer(c.Native()))

	var cret *C.GskTransform
	var goret1 *gsk.Transform

	cret = C.gtk_fixed_layout_child_get_transform(arg0)

	goret1 = gsk.WrapTransform(unsafe.Pointer(cret))

	return goret1
}

// SetTransform sets the transformation of the child of a FixedLayout.
func (c fixedLayoutChild) SetTransform(transform *gsk.Transform) {
	var arg0 *C.GtkFixedLayoutChild
	var arg1 *C.GskTransform

	arg0 = (*C.GtkFixedLayoutChild)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GskTransform)(unsafe.Pointer(transform.Native()))

	C.gtk_fixed_layout_child_set_transform(arg0, transform)
}