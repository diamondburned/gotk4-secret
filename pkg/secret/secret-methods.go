// Code generated by girgen. DO NOT EDIT.

package secret

import (
	"fmt"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: libsecret-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsecret/secret.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.secret_search_flags_get_type()), F: marshalSearchFlags},
	})
}

// SearchFlags various flags to be used with secret_service_search() and
// secret_service_search_sync().
type SearchFlags C.guint

const (
	// SearchNone: no flags.
	SearchNone SearchFlags = 0b0
	// SearchAll: all the items matching the search will be returned, instead of
	// just the first one.
	SearchAll SearchFlags = 0b10
	// SearchUnlock: unlock locked items while searching.
	SearchUnlock SearchFlags = 0b100
	// SearchLoadSecrets: while searching load secrets for items that are not
	// locked.
	SearchLoadSecrets SearchFlags = 0b1000
)

func marshalSearchFlags(p uintptr) (interface{}, error) {
	return SearchFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for SearchFlags.
func (s SearchFlags) String() string {
	if s == 0 {
		return "SearchFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(51)

	for s != 0 {
		next := s & (s - 1)
		bit := s - next

		switch bit {
		case SearchNone:
			builder.WriteString("None|")
		case SearchAll:
			builder.WriteString("All|")
		case SearchUnlock:
			builder.WriteString("Unlock|")
		case SearchLoadSecrets:
			builder.WriteString("LoadSecrets|")
		default:
			builder.WriteString(fmt.Sprintf("SearchFlags(0b%b)|", bit))
		}

		s = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if s contains other.
func (s SearchFlags) Has(other SearchFlags) bool {
	return (s & other) == other
}
