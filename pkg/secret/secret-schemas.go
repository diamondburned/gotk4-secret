// Code generated by girgen. DO NOT EDIT.

package secret

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsecret/secret.h>
import "C"

// glib.Type values for secret-schemas.go.
var GTypeSchemaType = externglib.Type(C.secret_schema_type_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSchemaType, F: marshalSchemaType},
	})
}

// SchemaType: different types of schemas for storing secrets, intended for use
// with secret_get_schema().
type SchemaType C.gint

const (
	// SchemaTypeNote: personal passwords; see SECRET_SCHEMA_NOTE.
	SchemaTypeNote SchemaType = iota
	// SchemaTypeCompatNetwork: network passwords from older libgnome-keyring
	// storage; see SECRET_SCHEMA_COMPAT_NETWORK.
	SchemaTypeCompatNetwork
)

func marshalSchemaType(p uintptr) (interface{}, error) {
	return SchemaType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SchemaType.
func (s SchemaType) String() string {
	switch s {
	case SchemaTypeNote:
		return "Note"
	case SchemaTypeCompatNetwork:
		return "CompatNetwork"
	default:
		return fmt.Sprintf("SchemaType(%d)", s)
	}
}

// GetSchema: get a secret storage schema of the given type.
//
// C code may access the schemas (such as SECRET_SCHEMA_NOTE) directly, but
// language bindings cannot, and must use this accessor.
//
// The function takes the following parameters:
//
//    - typ: type of schema to get.
//
// The function returns the following values:
//
//    - schema type.
//
func GetSchema(typ SchemaType) *Schema {
	var _arg1 C.SecretSchemaType // out
	var _cret *C.SecretSchema    // in

	_arg1 = C.SecretSchemaType(typ)

	_cret = C.secret_get_schema(_arg1)
	runtime.KeepAlive(typ)

	var _schema *Schema // out

	_schema = (*Schema)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.secret_schema_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_schema)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.secret_schema_unref((*C.SecretSchema)(intern.C))
		},
	)

	return _schema
}
