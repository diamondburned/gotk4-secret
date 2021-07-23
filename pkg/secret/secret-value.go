// Code generated by girgen. DO NOT EDIT.

package secret

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libsecret-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <libsecret/secret.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.secret_value_get_type()), F: marshalValue},
	})
}

// Value: secret value, like a password or other binary secret.
type Value struct {
	nocopy gextras.NoCopy
	native *C.SecretValue
}

func marshalValue(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Value{native: (*C.SecretValue)(unsafe.Pointer(b))}, nil
}

// NewValue constructs a struct Value.
func NewValue(secret string, length int, contentType string) *Value {
	var _arg1 *C.gchar       // out
	var _arg2 C.gssize       // out
	var _arg3 *C.gchar       // out
	var _cret *C.SecretValue // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(secret)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.secret_value_new(_arg1, _arg2, _arg3)

	var _value *Value // out

	_value = (*Value)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.secret_value_ref(_cret)
	runtime.SetFinalizer(_value, func(v *Value) {
		C.secret_value_unref((C.gpointer)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _value
}

// ContentType: get the content type of the secret value, such as
// <literal>text/plain</literal>.
func (value *Value) ContentType() string {
	var _arg0 *C.SecretValue // out
	var _cret *C.gchar       // in

	_arg0 = (*C.SecretValue)(gextras.StructNative(unsafe.Pointer(value)))

	_cret = C.secret_value_get_content_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Text: get the secret data in the Value if it contains a textual value. The
// content type must be <literal>text/plain</literal>.
func (value *Value) Text() string {
	var _arg0 *C.SecretValue // out
	var _cret *C.gchar       // in

	_arg0 = (*C.SecretValue)(gextras.StructNative(unsafe.Pointer(value)))

	_cret = C.secret_value_get_text(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UnrefToPassword: unreference a Value and steal the secret data in Value as
// nonpageable memory.
func (value *Value) UnrefToPassword(length *uint) string {
	var _arg0 *C.SecretValue // out
	var _arg1 *C.gsize       // out
	var _cret *C.gchar       // in

	_arg0 = (*C.SecretValue)(gextras.StructNative(unsafe.Pointer(value)))
	_arg1 = (*C.gsize)(unsafe.Pointer(length))

	_cret = C.secret_value_unref_to_password(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
