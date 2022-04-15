// Code generated by girgen. DO NOT EDIT.

package secret

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsecret/secret.h>
import "C"

// glib.Type values for secret-schema.go.
var (
	GTypeSchemaAttributeType = externglib.Type(C.secret_schema_attribute_type_get_type())
	GTypeSchemaFlags         = externglib.Type(C.secret_schema_flags_get_type())
	GTypeSchema              = externglib.Type(C.secret_schema_get_type())
	GTypeSchemaAttribute     = externglib.Type(C.secret_schema_attribute_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSchemaAttributeType, F: marshalSchemaAttributeType},
		{T: GTypeSchemaFlags, F: marshalSchemaFlags},
		{T: GTypeSchema, F: marshalSchema},
		{T: GTypeSchemaAttribute, F: marshalSchemaAttribute},
	})
}

// SchemaAttributeType: type of an attribute in a Schema. Attributes are stored
// as strings in the Secret Service, and the attribute types simply define
// standard ways to store integer and boolean values as strings.
type SchemaAttributeType C.gint

const (
	// SchemaAttributeString: utf-8 string attribute.
	SchemaAttributeString SchemaAttributeType = iota
	// SchemaAttributeInteger: integer attribute, stored as a decimal.
	SchemaAttributeInteger
	// SchemaAttributeBoolean: boolean attribute, stored as 'true' or 'false'.
	SchemaAttributeBoolean
)

func marshalSchemaAttributeType(p uintptr) (interface{}, error) {
	return SchemaAttributeType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SchemaAttributeType.
func (s SchemaAttributeType) String() string {
	switch s {
	case SchemaAttributeString:
		return "String"
	case SchemaAttributeInteger:
		return "Integer"
	case SchemaAttributeBoolean:
		return "Boolean"
	default:
		return fmt.Sprintf("SchemaAttributeType(%d)", s)
	}
}

// SchemaFlags flags for a Schema definition.
type SchemaFlags C.guint

const (
	// SchemaNone: no flags for the schema.
	SchemaNone SchemaFlags = 0b0
	// SchemaDontMatchName: don't match the schema name when looking up or
	// removing passwords.
	SchemaDontMatchName SchemaFlags = 0b10
)

func marshalSchemaFlags(p uintptr) (interface{}, error) {
	return SchemaFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SchemaFlags.
func (s SchemaFlags) String() string {
	if s == 0 {
		return "SchemaFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(30)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SchemaNone:
			builder.WriteString("None|")
		case SchemaDontMatchName:
			builder.WriteString("DontMatchName|")
		default:
			builder.WriteString(fmt.Sprintf("SchemaFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SchemaFlags) Has(other SchemaFlags) bool {
	return (s & other) == other
}

// Schema represents a set of attributes that are stored with an item. These
// schemas are used for interoperability between various services storing the
// same types of items.
//
// Each schema has a name like "org.gnome.keyring.NetworkPassword", and defines
// a set of attributes, and types (string, integer, boolean) for those
// attributes.
//
// Attributes are stored as strings in the Secret Service, and the attribute
// types simply define standard ways to store integer and boolean values as
// strings. Attributes are represented in libsecret via a Table with string keys
// and values. Even for values that defined as an integer or boolean in the
// schema, the attribute values in the Table are strings. Boolean values are
// stored as the strings 'true' and 'false'. Integer values are stored in
// decimal, with a preceding negative sign for negative integers.
//
// Schemas are handled entirely on the client side by this library. The name of
// the schema is automatically stored as an attribute on the item.
//
// Normally when looking up passwords only those with matching schema names are
// returned. If the schema flags contain the SECRET_SCHEMA_DONT_MATCH_NAME flag,
// then lookups will not check that the schema name matches that on the item,
// only the schema's attributes are matched. This is useful when you are looking
// up items that are not stored by the libsecret library. Other libraries such
// as libgnome-keyring don't store the schema name.
//
// An instance of this type is always passed by reference.
type Schema struct {
	*schema
}

// schema is the struct that's finalized.
type schema struct {
	native *C.SecretSchema
}

func marshalSchema(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Schema{&schema{(*C.SecretSchema)(b)}}, nil
}

// NewSchema constructs a struct Schema.
func NewSchema(name string, flags SchemaFlags, attributeNamesAndTypes map[string]SchemaAttributeType) *Schema {
	var _arg1 *C.gchar            // out
	var _arg2 C.SecretSchemaFlags // out
	var _arg3 *C.GHashTable       // out
	var _cret *C.SecretSchema     // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.SecretSchemaFlags(flags)
	_arg3 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributeNamesAndTypes {
		var kdst *C.gchar                     // out
		var vdst *C.SecretSchemaAttributeType // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.SecretSchemaAttributeType)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg3, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg3)

	_cret = C.secret_schema_newv(_arg1, _arg2, _arg3)
	runtime.KeepAlive(name)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(attributeNamesAndTypes)

	var _schema *Schema // out

	_schema = (*Schema)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_schema)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.secret_schema_unref((*C.SecretSchema)(intern.C))
		},
	)

	return _schema
}

// Name: dotted name of the schema.
func (s *Schema) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(s.native.name)))
	return v
}

// Flags flags for the schema.
func (s *Schema) Flags() SchemaFlags {
	var v SchemaFlags // out
	v = SchemaFlags(s.native.flags)
	return v
}

// Attributes: attribute names and types of those attributes.
func (s *Schema) Attributes() [32]SchemaAttribute {
	var v [32]SchemaAttribute // out
	{
		src := &s.native.attributes
		for i := 0; i < 32; i++ {
			v[i] = *(*SchemaAttribute)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}
	return v
}

// SchemaAttribute: attribute in a Schema.
//
// An instance of this type is always passed by reference.
type SchemaAttribute struct {
	*schemaAttribute
}

// schemaAttribute is the struct that's finalized.
type schemaAttribute struct {
	native *C.SecretSchemaAttribute
}

func marshalSchemaAttribute(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &SchemaAttribute{&schemaAttribute{(*C.SecretSchemaAttribute)(b)}}, nil
}

// Name: name of the attribute.
func (s *SchemaAttribute) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(s.native.name)))
	return v
}

// Type: type of the attribute.
func (s *SchemaAttribute) Type() SchemaAttributeType {
	var v SchemaAttributeType // out
	v = SchemaAttributeType(s.native._type)
	return v
}
