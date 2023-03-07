// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSettingsSchema       = coreglib.Type(C.g_settings_schema_get_type())
	GTypeSettingsSchemaSource = coreglib.Type(C.g_settings_schema_source_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSettingsSchema, F: marshalSettingsSchema},
		coreglib.TypeMarshaler{T: GTypeSettingsSchemaSource, F: marshalSettingsSchemaSource},
	})
}

// SettingsSchema and Schema APIs provide a mechanism for advanced control over
// the loading of schemas and a mechanism for introspecting their content.
//
// Plugin loading systems that wish to provide plugins a way to access settings
// face the problem of how to make the schemas for these settings visible to
// GSettings. Typically, a plugin will want to ship the schema along with itself
// and it won't be installed into the standard system directories for schemas.
//
// SchemaSource provides a mechanism for dealing with this by allowing the
// creation of a new 'schema source' from which schemas can be acquired. This
// schema source can then become part of the metadata associated with the plugin
// and queried whenever the plugin requires access to some settings.
//
// Consider the following example:
//
//    {
//      GSettings *settings;
//      gint some_value;
//
//      settings = plugin_get_settings (self, NULL);
//      some_value = g_settings_get_int (settings, "some-value");
//      ...
//    }
//
// It's also possible that the plugin system expects the schema source files
// (ie: .gschema.xml files) instead of a gschemas.compiled file. In that case,
// the plugin loading system must compile the schemas for itself before
// attempting to create the settings source.
//
// An instance of this type is always passed by reference.
type SettingsSchema struct {
	*settingsSchema
}

// settingsSchema is the struct that's finalized.
type settingsSchema struct {
	native *C.GSettingsSchema
}

func marshalSettingsSchema(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &SettingsSchema{&settingsSchema{(*C.GSettingsSchema)(b)}}, nil
}

// ID: get the ID of schema.
//
// The function returns the following values:
//
//    - utf8: ID.
//
func (schema *SettingsSchema) ID() string {
	var _arg0 *C.GSettingsSchema // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GSettingsSchema)(gextras.StructNative(unsafe.Pointer(schema)))

	_cret = C.g_settings_schema_get_id(_arg0)
	runtime.KeepAlive(schema)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Key gets the key named name from schema.
//
// It is a programmer error to request a key that does not exist. See
// g_settings_schema_list_keys().
//
// The function takes the following parameters:
//
//    - name of a key.
//
// The function returns the following values:
//
//    - settingsSchemaKey for name.
//
func (schema *SettingsSchema) Key(name string) *SettingsSchemaKey {
	var _arg0 *C.GSettingsSchema    // out
	var _arg1 *C.gchar              // out
	var _cret *C.GSettingsSchemaKey // in

	_arg0 = (*C.GSettingsSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_schema_get_key(_arg0, _arg1)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(name)

	var _settingsSchemaKey *SettingsSchemaKey // out

	_settingsSchemaKey = (*SettingsSchemaKey)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_settingsSchemaKey)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_settings_schema_key_unref((*C.GSettingsSchemaKey)(intern.C))
		},
	)

	return _settingsSchemaKey
}

// Path gets the path associated with schema, or NULL.
//
// Schemas may be single-instance or relocatable. Single-instance schemas
// correspond to exactly one set of keys in the backend database: those located
// at the path returned by this function.
//
// Relocatable schemas can be referenced by other schemas and can therefore
// describe multiple sets of keys at different locations. For relocatable
// schemas, this function will return NULL.
//
// The function returns the following values:
//
//    - utf8 (optional): path of the schema, or NULL.
//
func (schema *SettingsSchema) Path() string {
	var _arg0 *C.GSettingsSchema // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GSettingsSchema)(gextras.StructNative(unsafe.Pointer(schema)))

	_cret = C.g_settings_schema_get_path(_arg0)
	runtime.KeepAlive(schema)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// HasKey checks if schema has a key named name.
//
// The function takes the following parameters:
//
//    - name of a key.
//
// The function returns the following values:
//
//    - ok: TRUE if such a key exists.
//
func (schema *SettingsSchema) HasKey(name string) bool {
	var _arg0 *C.GSettingsSchema // out
	var _arg1 *C.gchar           // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GSettingsSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_schema_has_key(_arg0, _arg1)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListChildren gets the list of children in schema.
//
// You should free the return value with g_strfreev() when you are done with it.
//
// The function returns the following values:
//
//    - utf8s: list of the children on settings, in no defined order.
//
func (schema *SettingsSchema) ListChildren() []string {
	var _arg0 *C.GSettingsSchema // out
	var _cret **C.gchar          // in

	_arg0 = (*C.GSettingsSchema)(gextras.StructNative(unsafe.Pointer(schema)))

	_cret = C.g_settings_schema_list_children(_arg0)
	runtime.KeepAlive(schema)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// ListKeys introspects the list of keys on schema.
//
// You should probably not be calling this function from "normal" code (since
// you should already know what keys are in your schema). This function is
// intended for introspection reasons.
//
// The function returns the following values:
//
//    - utf8s: list of the keys on schema, in no defined order.
//
func (schema *SettingsSchema) ListKeys() []string {
	var _arg0 *C.GSettingsSchema // out
	var _cret **C.gchar          // in

	_arg0 = (*C.GSettingsSchema)(gextras.StructNative(unsafe.Pointer(schema)))

	_cret = C.g_settings_schema_list_keys(_arg0)
	runtime.KeepAlive(schema)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// SettingsSchemaSource: this is an opaque structure type. You may not access it
// directly.
//
// An instance of this type is always passed by reference.
type SettingsSchemaSource struct {
	*settingsSchemaSource
}

// settingsSchemaSource is the struct that's finalized.
type settingsSchemaSource struct {
	native *C.GSettingsSchemaSource
}

func marshalSettingsSchemaSource(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &SettingsSchemaSource{&settingsSchemaSource{(*C.GSettingsSchemaSource)(b)}}, nil
}

// NewSettingsSchemaSourceFromDirectory constructs a struct SettingsSchemaSource.
func NewSettingsSchemaSourceFromDirectory(directory string, parent *SettingsSchemaSource, trusted bool) (*SettingsSchemaSource, error) {
	var _arg1 *C.gchar                 // out
	var _arg2 *C.GSettingsSchemaSource // out
	var _arg3 C.gboolean               // out
	var _cret *C.GSettingsSchemaSource // in
	var _cerr *C.GError                // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(directory)))
	defer C.free(unsafe.Pointer(_arg1))
	if parent != nil {
		_arg2 = (*C.GSettingsSchemaSource)(gextras.StructNative(unsafe.Pointer(parent)))
	}
	if trusted {
		_arg3 = C.TRUE
	}

	_cret = C.g_settings_schema_source_new_from_directory(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(directory)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(trusted)

	var _settingsSchemaSource *SettingsSchemaSource // out
	var _goerr error                                // out

	_settingsSchemaSource = (*SettingsSchemaSource)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_settingsSchemaSource)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_settings_schema_source_unref((*C.GSettingsSchemaSource)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _settingsSchemaSource, _goerr
}

// ListSchemas lists the schemas in a given source.
//
// If recursive is TRUE then include parent sources. If FALSE then only include
// the schemas from one source (ie: one directory). You probably want TRUE.
//
// Non-relocatable schemas are those for which you can call g_settings_new().
// Relocatable schemas are those for which you must use
// g_settings_new_with_path().
//
// Do not call this function from normal programs. This is designed for use by
// database editors, commandline tools, etc.
//
// The function takes the following parameters:
//
//    - recursive: if we should recurse.
//
// The function returns the following values:
//
//    - nonRelocatable: the list of non-relocatable schemas, in no defined order.
//    - relocatable: list of relocatable schemas, in no defined order.
//
func (source *SettingsSchemaSource) ListSchemas(recursive bool) (nonRelocatable []string, relocatable []string) {
	var _arg0 *C.GSettingsSchemaSource // out
	var _arg1 C.gboolean               // out
	var _arg2 **C.gchar                // in
	var _arg3 **C.gchar                // in

	_arg0 = (*C.GSettingsSchemaSource)(gextras.StructNative(unsafe.Pointer(source)))
	if recursive {
		_arg1 = C.TRUE
	}

	C.g_settings_schema_source_list_schemas(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(source)
	runtime.KeepAlive(recursive)

	var _nonRelocatable []string // out
	var _relocatable []string    // out

	defer C.free(unsafe.Pointer(_arg2))
	{
		var i int
		var z *C.gchar
		for p := _arg2; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_arg2, i)
		_nonRelocatable = make([]string, i)
		for i := range src {
			_nonRelocatable[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	defer C.free(unsafe.Pointer(_arg3))
	{
		var i int
		var z *C.gchar
		for p := _arg3; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_arg3, i)
		_relocatable = make([]string, i)
		for i := range src {
			_relocatable[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _nonRelocatable, _relocatable
}

// Lookup looks up a schema with the identifier schema_id in source.
//
// This function is not required for normal uses of #GSettings but it may be
// useful to authors of plugin management systems or to those who want to
// introspect the content of schemas.
//
// If the schema isn't found directly in source and recursive is TRUE then the
// parent sources will also be checked.
//
// If the schema isn't found, NULL is returned.
//
// The function takes the following parameters:
//
//    - schemaId: schema ID.
//    - recursive: TRUE if the lookup should be recursive.
//
// The function returns the following values:
//
//    - settingsSchema (optional): new Schema.
//
func (source *SettingsSchemaSource) Lookup(schemaId string, recursive bool) *SettingsSchema {
	var _arg0 *C.GSettingsSchemaSource // out
	var _arg1 *C.gchar                 // out
	var _arg2 C.gboolean               // out
	var _cret *C.GSettingsSchema       // in

	_arg0 = (*C.GSettingsSchemaSource)(gextras.StructNative(unsafe.Pointer(source)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(schemaId)))
	defer C.free(unsafe.Pointer(_arg1))
	if recursive {
		_arg2 = C.TRUE
	}

	_cret = C.g_settings_schema_source_lookup(_arg0, _arg1, _arg2)
	runtime.KeepAlive(source)
	runtime.KeepAlive(schemaId)
	runtime.KeepAlive(recursive)

	var _settingsSchema *SettingsSchema // out

	if _cret != nil {
		_settingsSchema = (*SettingsSchema)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_settingsSchema)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_settings_schema_unref((*C.GSettingsSchema)(intern.C))
			},
		)
	}

	return _settingsSchema
}

// SettingsSchemaSourceGetDefault gets the default system schema source.
//
// This function is not required for normal uses of #GSettings but it may be
// useful to authors of plugin management systems or to those who want to
// introspect the content of schemas.
//
// If no schemas are installed, NULL will be returned.
//
// The returned source may actually consist of multiple schema sources from
// different directories, depending on which directories were given in
// XDG_DATA_DIRS and GSETTINGS_SCHEMA_DIR. For this reason, all lookups
// performed against the default source should probably be done recursively.
//
// The function returns the following values:
//
//    - settingsSchemaSource (optional): default schema source.
//
func SettingsSchemaSourceGetDefault() *SettingsSchemaSource {
	var _cret *C.GSettingsSchemaSource // in

	_cret = C.g_settings_schema_source_get_default()

	var _settingsSchemaSource *SettingsSchemaSource // out

	if _cret != nil {
		_settingsSchemaSource = (*SettingsSchemaSource)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_settings_schema_source_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_settingsSchemaSource)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_settings_schema_source_unref((*C.GSettingsSchemaSource)(intern.C))
			},
		)
	}

	return _settingsSchemaSource
}