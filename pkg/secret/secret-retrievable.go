// Code generated by girgen. DO NOT EDIT.

package secret

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <libsecret/secret.h>
// extern SecretValue* _gotk4_secret1_RetrievableInterface_retrieve_secret_finish(SecretRetrievable*, GAsyncResult*, GError**);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// glib.Type values for secret-retrievable.go.
var GTypeRetrievable = externglib.Type(C.secret_retrievable_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeRetrievable, F: marshalRetrievable},
	})
}

// Retrievable: object representing a read-only view of a secret item in the
// Secret Service.
//
// Retrievable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Retrievable struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Retrievable)(nil)
)

// Retrievabler describes Retrievable's interface methods.
type Retrievabler interface {
	externglib.Objector

	// Attributes: get the attributes of this object.
	Attributes() map[string]string
	// Created: get the created date and time of the object.
	Created() uint64
	// Label: get the label of this item.
	Label() string
	// Modified: get the modified date and time of the object.
	Modified() uint64
	// RetrieveSecret: retrieve the secret value of this object.
	RetrieveSecret(ctx context.Context, callback gio.AsyncReadyCallback)
	// RetrieveSecretFinish: complete asynchronous operation to retrieve the
	// secret value of this object.
	RetrieveSecretFinish(result gio.AsyncResulter) (*Value, error)
	// RetrieveSecretSync: retrieve the secret value of this object
	// synchronously.
	RetrieveSecretSync(ctx context.Context) (*Value, error)
}

var _ Retrievabler = (*Retrievable)(nil)

func wrapRetrievable(obj *externglib.Object) *Retrievable {
	return &Retrievable{
		Object: obj,
	}
}

func marshalRetrievable(p uintptr) (interface{}, error) {
	return wrapRetrievable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Attributes: get the attributes of this object.
//
// The attributes are a mapping of string keys to string values. Attributes are
// used to search for items. Attributes are not stored or transferred securely
// by the secret service.
//
// Do not modify the attribute returned by this method.
//
// The function returns the following values:
//
//    - hashTable: new reference to the attributes, which should not be modified,
//      and released with g_hash_table_unref().
//
func (self *Retrievable) Attributes() map[string]string {
	var _arg0 *C.SecretRetrievable // out
	var _cret *C.GHashTable        // in

	_arg0 = (*C.SecretRetrievable)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.secret_retrievable_get_attributes(_arg0)
	runtime.KeepAlive(self)

	var _hashTable map[string]string // out

	_hashTable = make(map[string]string, gextras.HashTableSize(unsafe.Pointer(_cret)))
	gextras.MoveHashTable(unsafe.Pointer(_cret), true, func(k, v unsafe.Pointer) {
		ksrc := *(**C.gchar)(k)
		vsrc := *(**C.gchar)(v)
		var kdst string // out
		var vdst string // out
		kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
		defer C.free(unsafe.Pointer(ksrc))
		vdst = C.GoString((*C.gchar)(unsafe.Pointer(vsrc)))
		defer C.free(unsafe.Pointer(vsrc))
		_hashTable[kdst] = vdst
	})

	return _hashTable
}

// Created: get the created date and time of the object. The return value is the
// number of seconds since the unix epoch, January 1st 1970.
//
// The function returns the following values:
//
//    - guint64: created date and time.
//
func (self *Retrievable) Created() uint64 {
	var _arg0 *C.SecretRetrievable // out
	var _cret C.guint64            // in

	_arg0 = (*C.SecretRetrievable)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.secret_retrievable_get_created(_arg0)
	runtime.KeepAlive(self)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Label: get the label of this item.
//
// The function returns the following values:
//
//    - utf8: label, which should be freed with g_free().
//
func (self *Retrievable) Label() string {
	var _arg0 *C.SecretRetrievable // out
	var _cret *C.gchar             // in

	_arg0 = (*C.SecretRetrievable)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.secret_retrievable_get_label(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Modified: get the modified date and time of the object. The return value is
// the number of seconds since the unix epoch, January 1st 1970.
//
// The function returns the following values:
//
//    - guint64: modified date and time.
//
func (self *Retrievable) Modified() uint64 {
	var _arg0 *C.SecretRetrievable // out
	var _cret C.guint64            // in

	_arg0 = (*C.SecretRetrievable)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.secret_retrievable_get_modified(_arg0)
	runtime.KeepAlive(self)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// RetrieveSecret: retrieve the secret value of this object.
//
// Each retrievable object has a single secret which might be a password or some
// other secret binary value.
//
// This function returns immediately and completes asynchronously.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - callback (optional): called when the operation completes.
//
func (self *Retrievable) RetrieveSecret(ctx context.Context, callback gio.AsyncReadyCallback) {
	var _arg0 *C.SecretRetrievable  // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.SecretRetrievable)(unsafe.Pointer(externglib.InternObject(self).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.secret_retrievable_retrieve_secret(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// RetrieveSecretFinish: complete asynchronous operation to retrieve the secret
// value of this object.
//
// The function takes the following parameters:
//
//    - result asynchronous result passed to callback.
//
// The function returns the following values:
//
//    - value (optional): secret value which should be released with
//      secret_value_unref(), or NULL.
//
func (self *Retrievable) RetrieveSecretFinish(result gio.AsyncResulter) (*Value, error) {
	var _arg0 *C.SecretRetrievable // out
	var _arg1 *C.GAsyncResult      // out
	var _cret *C.SecretValue       // in
	var _cerr *C.GError            // in

	_arg0 = (*C.SecretRetrievable)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.secret_retrievable_retrieve_secret_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(result)

	var _value *Value // out
	var _goerr error  // out

	if _cret != nil {
		_value = (*Value)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_value)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.secret_value_unref((C.gpointer)(intern.C))
			},
		)
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _value, _goerr
}

// RetrieveSecretSync: retrieve the secret value of this object synchronously.
//
// Each retrievable object has a single secret which might be a password or some
// other secret binary value.
//
// This method may block indefinitely and should not be used in user interface
// threads.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//
// The function returns the following values:
//
//    - value (optional): secret value which should be released with
//      secret_value_unref(), or NULL.
//
func (self *Retrievable) RetrieveSecretSync(ctx context.Context) (*Value, error) {
	var _arg0 *C.SecretRetrievable // out
	var _arg1 *C.GCancellable      // out
	var _cret *C.SecretValue       // in
	var _cerr *C.GError            // in

	_arg0 = (*C.SecretRetrievable)(unsafe.Pointer(externglib.InternObject(self).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.secret_retrievable_retrieve_secret_sync(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)

	var _value *Value // out
	var _goerr error  // out

	if _cret != nil {
		_value = (*Value)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_value)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.secret_value_unref((C.gpointer)(intern.C))
			},
		)
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _value, _goerr
}
