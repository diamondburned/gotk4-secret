// Code generated by girgen. DO NOT EDIT.

package secret

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <libsecret/secret.h>
import "C"

// COLLECTION_DEFAULT alias to the default collection. This can be passed to
// secret_password_store() secret_collection_for_alias().
const COLLECTION_DEFAULT = "default"

// COLLECTION_SESSION alias to the session collection, which will be cleared
// when the user ends the session. This can be passed to
// secret_password_store(), secret_collection_for_alias() or similar functions.
const COLLECTION_SESSION = "session"

// The function returns the following values:
//
func ErrorGetQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.secret_error_get_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}
