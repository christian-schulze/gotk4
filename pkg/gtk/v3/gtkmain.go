// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// KeySnoopFunc: key snooper functions are called before normal event delivery.
// They can be used to implement custom key event handling.
type KeySnoopFunc func(grabWidget Widget, event *gdk.EventKey) int

//export gotk4_KeySnoopFunc
func gotk4_KeySnoopFunc(arg0 *C.GtkWidget, arg1 *C.GdkEventKey, arg2 C.gpointer) C.gint {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}
	fn := v.(KeySnoopFunc)

	ret := fn(grabWidget, event, funcData)

	cret = C.gint(ret)

	return cret
}

// CheckVersion checks that the GTK+ library in use is compatible with the given
// version. Generally you would pass in the constants K_MAJOR_VERSION,
// K_MINOR_VERSION, K_MICRO_VERSION as the three arguments to this function;
// that produces a check that the library in use is compatible with the version
// of GTK+ the application or module was compiled against.
//
// Compatibility is defined by two things: first the version of the running
// library is newer than the version
// @required_major.required_minor.@required_micro. Second the running library
// must be binary compatible with the version
// @required_major.required_minor.@required_micro (same major version.)
//
// This function is primarily for GTK+ modules; the module can call this
// function to check that it wasn’t loaded into an incompatible version of GTK+.
// However, such a check isn’t completely reliable, since the module may be
// linked against an old version of GTK+ and calling the old version of
// gtk_check_version(), but still get loaded into an application using a newer
// version of GTK+.
func CheckVersion(requiredMajor uint, requiredMinor uint, requiredMicro uint) string {
	var arg1 C.guint
	var arg2 C.guint
	var arg3 C.guint

	arg1 = C.guint(requiredMajor)
	arg2 = C.guint(requiredMinor)
	arg3 = C.guint(requiredMicro)

	var cret *C.gchar
	var goret1 string

	cret = C.gtk_check_version(requiredMajor, requiredMinor, requiredMicro)

	goret1 = C.GoString(cret)

	return goret1
}

// DeviceGrabAdd adds a GTK+ grab on @device, so all the events on @device and
// its associated pointer or keyboard (if any) are delivered to @widget. If the
// @block_others parameter is true, any other devices will be unable to interact
// with @widget during the grab.
func DeviceGrabAdd(widget Widget, device gdk.Device, blockOthers bool) {
	var arg1 *C.GtkWidget
	var arg2 *C.GdkDevice
	var arg3 C.gboolean

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	if blockOthers {
		arg3 = C.gboolean(1)
	}

	C.gtk_device_grab_add(widget, device, blockOthers)
}

// DeviceGrabRemove removes a device grab from the given widget.
//
// You have to pair calls to gtk_device_grab_add() and gtk_device_grab_remove().
func DeviceGrabRemove(widget Widget, device gdk.Device) {
	var arg1 *C.GtkWidget
	var arg2 *C.GdkDevice

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	C.gtk_device_grab_remove(widget, device)
}

// DisableSetlocale prevents gtk_init(), gtk_init_check(), gtk_init_with_args()
// and gtk_parse_args() from automatically calling `setlocale (LC_ALL, "")`. You
// would want to use this function if you wanted to set the locale for your
// program to something other than the user’s locale, or if you wanted to set
// different values for different locale categories.
//
// Most programs should not need to call this function.
func DisableSetlocale() {
	C.gtk_disable_setlocale()
}

// EventsPending checks if any events are pending.
//
// This can be used to update the UI and invoke timeouts etc. while doing some
// time intensive computation.
//
// Updating the UI during a long computation
//
//     // computation going on...
//
//     while (gtk_events_pending ())
//       gtk_main_iteration ();
//
//     // ...computation continued
func EventsPending() bool {
	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_events_pending()

	goret1 = C.bool(cret) != C.false

	return goret1
}

// False: analogical to gtk_true(), this function does nothing but always
// returns false.
func False() bool {
	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_false()

	goret1 = C.bool(cret) != C.false

	return goret1
}

// GetBinaryAge returns the binary age as passed to `libtool` when building the
// GTK+ library the process is running against. If `libtool` means nothing to
// you, don't worry about it.
func GetBinaryAge() uint {
	var cret C.guint
	var goret1 uint

	cret = C.gtk_get_binary_age()

	goret1 = C.guint(cret)

	return goret1
}

// GetCurrentEventDevice: if there is a current event and it has a device,
// return that device, otherwise return nil.
func GetCurrentEventDevice() gdk.Device {
	var cret *C.GdkDevice
	var goret1 gdk.Device

	cret = C.gtk_get_current_event_device()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Device)

	return goret1
}

// GetCurrentEventState: if there is a current event and it has a state field,
// place that state field in @state and return true, otherwise return false.
func GetCurrentEventState() (state gdk.ModifierType, ok bool) {
	var arg1 *C.GdkModifierType
	var ret1 *gdk.ModifierType
	var cret C.gboolean
	var goret2 bool

	cret = C.gtk_get_current_event_state(&arg1)

	ret1 = *gdk.ModifierType(arg1)
	goret2 = C.bool(cret) != C.false

	return ret1, goret2
}

// GetCurrentEventTime: if there is a current event and it has a timestamp,
// return that timestamp, otherwise return GDK_CURRENT_TIME.
func GetCurrentEventTime() uint32 {
	var cret C.guint32
	var goret1 uint32

	cret = C.gtk_get_current_event_time()

	goret1 = C.guint32(cret)

	return goret1
}

// GetDefaultLanguage returns the Language for the default language currently in
// effect. (Note that this can change over the life of an application.) The
// default language is derived from the current locale. It determines, for
// example, whether GTK+ uses the right-to-left or left-to-right text direction.
//
// This function is equivalent to pango_language_get_default(). See that
// function for details.
func GetDefaultLanguage() *pango.Language {
	var cret *C.PangoLanguage
	var goret1 *pango.Language

	cret = C.gtk_get_default_language()

	goret1 = pango.WrapLanguage(unsafe.Pointer(cret))

	return goret1
}

// GetInterfaceAge returns the interface age as passed to `libtool` when
// building the GTK+ library the process is running against. If `libtool` means
// nothing to you, don't worry about it.
func GetInterfaceAge() uint {
	var cret C.guint
	var goret1 uint

	cret = C.gtk_get_interface_age()

	goret1 = C.guint(cret)

	return goret1
}

// GetLocaleDirection: get the direction of the current locale. This is the
// expected reading direction for text and UI.
//
// This function depends on the current locale being set with setlocale() and
// will default to setting the GTK_TEXT_DIR_LTR direction otherwise.
// GTK_TEXT_DIR_NONE will never be returned.
//
// GTK+ sets the default text direction according to the locale during
// gtk_init(), and you should normally use gtk_widget_get_direction() or
// gtk_widget_get_default_direction() to obtain the current direcion.
//
// This function is only needed rare cases when the locale is changed after GTK+
// has already been initialized. In this case, you can use it to update the
// default text direction as follows:
//
//    setlocale (LC_ALL, new_locale);
//    direction = gtk_get_locale_direction ();
//    gtk_widget_set_default_direction (direction);
func GetLocaleDirection() TextDirection {
	var cret C.GtkTextDirection
	var goret1 TextDirection

	cret = C.gtk_get_locale_direction()

	goret1 = TextDirection(cret)

	return goret1
}

// GetMajorVersion returns the major version number of the GTK+ library. (e.g.
// in GTK+ version 3.1.5 this is 3.)
//
// This function is in the library, so it represents the GTK+ library your code
// is running against. Contrast with the K_MAJOR_VERSION macro, which represents
// the major version of the GTK+ headers you have included when compiling your
// code.
func GetMajorVersion() uint {
	var cret C.guint
	var goret1 uint

	cret = C.gtk_get_major_version()

	goret1 = C.guint(cret)

	return goret1
}

// GetMicroVersion returns the micro version number of the GTK+ library. (e.g.
// in GTK+ version 3.1.5 this is 5.)
//
// This function is in the library, so it represents the GTK+ library your code
// is are running against. Contrast with the K_MICRO_VERSION macro, which
// represents the micro version of the GTK+ headers you have included when
// compiling your code.
func GetMicroVersion() uint {
	var cret C.guint
	var goret1 uint

	cret = C.gtk_get_micro_version()

	goret1 = C.guint(cret)

	return goret1
}

// GetMinorVersion returns the minor version number of the GTK+ library. (e.g.
// in GTK+ version 3.1.5 this is 1.)
//
// This function is in the library, so it represents the GTK+ library your code
// is are running against. Contrast with the K_MINOR_VERSION macro, which
// represents the minor version of the GTK+ headers you have included when
// compiling your code.
func GetMinorVersion() uint {
	var cret C.guint
	var goret1 uint

	cret = C.gtk_get_minor_version()

	goret1 = C.guint(cret)

	return goret1
}

// GetOptionGroup returns a Group for the commandline arguments recognized by
// GTK+ and GDK.
//
// You should add this group to your Context with g_option_context_add_group(),
// if you are using g_option_context_parse() to parse your commandline
// arguments.
func GetOptionGroup(openDefaultDisplay bool) *glib.OptionGroup {
	var arg1 C.gboolean

	if openDefaultDisplay {
		arg1 = C.gboolean(1)
	}

	var cret *C.GOptionGroup
	var goret1 *glib.OptionGroup

	cret = C.gtk_get_option_group(openDefaultDisplay)

	goret1 = glib.WrapOptionGroup(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *glib.OptionGroup) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// GrabGetCurrent queries the current grab of the default window group.
func GrabGetCurrent() Widget {
	var cret *C.GtkWidget
	var goret1 Widget

	cret = C.gtk_grab_get_current()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret1
}

// KeySnooperInstall installs a key snooper function, which will get called on
// all key events before delivering them normally.
func KeySnooperInstall(snooper KeySnoopFunc) uint {

	var cret C.guint
	var goret1 uint

	cret = C.gtk_key_snooper_install(snooper, funcData)

	goret1 = C.guint(cret)

	return goret1
}

// KeySnooperRemove removes the key snooper function with the given id.
func KeySnooperRemove(snooperHandlerID uint) {
	var arg1 C.guint

	arg1 = C.guint(snooperHandlerID)

	C.gtk_key_snooper_remove(snooperHandlerID)
}

// Main runs the main loop until gtk_main_quit() is called.
//
// You can nest calls to gtk_main(). In that case gtk_main_quit() will make the
// innermost invocation of the main loop return.
func Main() {
	C.gtk_main()
}

// MainIteration runs a single iteration of the mainloop.
//
// If no events are waiting to be processed GTK+ will block until the next event
// is noticed. If you don’t want to block look at gtk_main_iteration_do() or
// check if any events are pending with gtk_events_pending() first.
func MainIteration() bool {
	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_main_iteration()

	goret1 = C.bool(cret) != C.false

	return goret1
}

// MainIterationDo runs a single iteration of the mainloop. If no events are
// available either return or block depending on the value of @blocking.
func MainIterationDo(blocking bool) bool {
	var arg1 C.gboolean

	if blocking {
		arg1 = C.gboolean(1)
	}

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_main_iteration_do(blocking)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// MainLevel asks for the current nesting level of the main loop.
func MainLevel() uint {
	var cret C.guint
	var goret1 uint

	cret = C.gtk_main_level()

	goret1 = C.guint(cret)

	return goret1
}

// MainQuit makes the innermost invocation of the main loop return when it
// regains control.
func MainQuit() {
	C.gtk_main_quit()
}

// True: all this function does it to return true.
//
// This can be useful for example if you want to inhibit the deletion of a
// window. Of course you should not do this as the user expects a reaction from
// clicking the close icon of the window...
//
// A persistent window
//
//    #include <gtk/gtk.h>
//
//    int
//    main (int argc, char **argv)
//    {
//      GtkWidget *win, *but;
//      const char *text = "Close yourself. I mean it!";
//
//      gtk_init (&argc, &argv);
//
//      win = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      g_signal_connect (win,
//                        "delete-event",
//                        G_CALLBACK (gtk_true),
//                        NULL);
//      g_signal_connect (win, "destroy",
//                        G_CALLBACK (gtk_main_quit),
//                        NULL);
//
//      but = gtk_button_new_with_label (text);
//      g_signal_connect_swapped (but, "clicked",
//                                G_CALLBACK (gtk_object_destroy),
//                                win);
//      gtk_container_add (GTK_CONTAINER (win), but);
//
//      gtk_widget_show_all (win);
//
//      gtk_main ();
//
//      return 0;
//    }
func True() bool {
	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_true()

	goret1 = C.bool(cret) != C.false

	return goret1
}