// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// PangoContextGetForDisplay creates a Context for display.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for the widget you
// intend to render text onto.
//
// The newly created context will have the default font options (see
// #cairo_font_options_t) for the display; if these options change it will not
// be updated. Using gtk_widget_get_pango_context() is more convenient if you
// want to keep a context around and track changes to the font rendering
// settings.
//
// The function takes the following parameters:
//
//    - display for which the context is to be created.
//
// The function returns the following values:
//
//    - context: new Context for display.
//
func PangoContextGetForDisplay(display *Display) *pango.Context {
	var _arg1 *C.GdkDisplay   // out
	var _cret *C.PangoContext // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gdk_pango_context_get_for_display(_arg1)
	runtime.KeepAlive(display)

	var _context *pango.Context // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_context = &pango.Context{
			Object: obj,
		}
	}

	return _context
}