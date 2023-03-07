// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// IconSizeLookupForSettings obtains the pixel size of a semantic icon size,
// possibly modified by user preferences for a particular Settings. Normally
// size would be K_ICON_SIZE_MENU, K_ICON_SIZE_BUTTON, etc. This function isn’t
// normally needed, gtk_widget_render_icon_pixbuf() is the usual way to get an
// icon for rendering, then just look at the size of the rendered pixbuf. The
// rendered pixbuf may not even correspond to the width/height returned by
// gtk_icon_size_lookup(), because themes are free to render the pixbuf however
// they like, including changing the usual size.
//
// Deprecated: Use gtk_icon_size_lookup() instead.
//
// The function takes the following parameters:
//
//    - settings object, used to determine which set of user preferences to used.
//    - size: icon size (IconSize).
//
// The function returns the following values:
//
//    - width (optional): location to store icon width.
//    - height (optional): location to store icon height.
//    - ok: TRUE if size was a valid size.
//
func IconSizeLookupForSettings(settings *Settings, size int) (width, height int, ok bool) {
	var _arg1 *C.GtkSettings // out
	var _arg2 C.GtkIconSize  // out
	var _arg3 C.gint         // in
	var _arg4 C.gint         // in
	var _cret C.gboolean     // in

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(coreglib.InternObject(settings).Native()))
	_arg2 = C.GtkIconSize(size)

	_cret = C.gtk_icon_size_lookup_for_settings(_arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(size)

	var _width int  // out
	var _height int // out
	var _ok bool    // out

	_width = int(_arg3)
	_height = int(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _width, _height, _ok
}