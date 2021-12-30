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
// #include <libsecret/secret.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// PasswordClearFinish: finish an asynchronous operation to remove passwords
// from the secret service.
//
// The function takes the following parameters:
//
//    - result asynchronous result passed to the callback.
//
func PasswordClearFinish(result gio.AsyncResulter) error {
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.secret_password_clear_finish(_arg1, &_cerr)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PasswordClear: remove unlocked matching passwords from the secret service.
//
// The attributes should be a set of key and value string pairs.
//
// All unlocked items that match the attributes will be deleted.
//
// This method will return immediately and complete asynchronously.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for the attributes.
//    - attributes: attribute keys and values.
//    - callback (optional): called when the operation completes.
//
func PasswordClear(ctx context.Context, schema *Schema, attributes map[string]string, callback gio.AsyncReadyCallback) {
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.SecretSchema       // out
	var _arg2 *C.GHashTable         // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.secret_password_clearv(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(callback)
}

// PasswordClearSync: remove unlocked matching passwords from the secret
// service.
//
// The attributes should be a set of key and value string pairs.
//
// All unlocked items that match the attributes will be deleted.
//
// This method may block indefinitely and should not be used in user interface
// threads.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for the attributes.
//    - attributes: attribute keys and values.
//
func PasswordClearSync(ctx context.Context, schema *Schema, attributes map[string]string) error {
	var _arg3 *C.GCancellable // out
	var _arg1 *C.SecretSchema // out
	var _arg2 *C.GHashTable   // out
	var _cerr *C.GError       // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		defer C.free(unsafe.Pointer(vdst))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg2)

	C.secret_password_clearv_sync(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PasswordLookupFinish: finish an asynchronous operation to lookup a password
// in the secret service.
//
// The function takes the following parameters:
//
//    - result asynchronous result passed to the callback.
//
// The function returns the following values:
//
//    - utf8: new password string which should be freed with
//      secret_password_free() or may be freed with g_free() when done.
//
func PasswordLookupFinish(result gio.AsyncResulter) (string, error) {
	var _arg1 *C.GAsyncResult // out
	var _cret *C.gchar        // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.secret_password_lookup_finish(_arg1, &_cerr)
	runtime.KeepAlive(result)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8, _goerr
}

// PasswordLookup: lookup a password in the secret service.
//
// The attributes should be a set of key and value string pairs.
//
// If no secret is found then NULL is returned.
//
// This method will return immediately and complete asynchronously.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for attributes.
//    - attributes: attribute keys and values.
//    - callback (optional): called when the operation completes.
//
func PasswordLookup(ctx context.Context, schema *Schema, attributes map[string]string, callback gio.AsyncReadyCallback) {
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.SecretSchema       // out
	var _arg2 *C.GHashTable         // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.secret_password_lookupv(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(callback)
}

// PasswordLookupSync: lookup a password in the secret service.
//
// The attributes should be a set of key and value string pairs.
//
// If no secret is found then NULL is returned.
//
// This method may block indefinitely and should not be used in user interface
// threads.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for attributes.
//    - attributes: attribute keys and values.
//
// The function returns the following values:
//
//    - utf8: new password string which should be freed with
//      secret_password_free() or may be freed with g_free() when done.
//
func PasswordLookupSync(ctx context.Context, schema *Schema, attributes map[string]string) (string, error) {
	var _arg3 *C.GCancellable // out
	var _arg1 *C.SecretSchema // out
	var _arg2 *C.GHashTable   // out
	var _cret *C.gchar        // in
	var _cerr *C.GError       // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		defer C.free(unsafe.Pointer(vdst))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg2)

	_cret = C.secret_password_lookupv_sync(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8, _goerr
}

// PasswordSearchFinish: finish an asynchronous operation to search for items in
// the secret service.
//
// The function takes the following parameters:
//
//    - result asynchronous result passed to the callback.
//
// The function returns the following values:
//
//    - list of Retrievable containing attributes of the matched items.
//
func PasswordSearchFinish(result gio.AsyncResulter) ([]Retrievabler, error) {
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GList        // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.secret_password_search_finish(_arg1, &_cerr)
	runtime.KeepAlive(result)

	var _list []Retrievabler // out
	var _goerr error         // out

	_list = make([]Retrievabler, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.SecretRetrievable)(v)
		var dst Retrievabler // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type secret.Retrievabler is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Retrievabler)
				return ok
			})
			rv, ok := casted.(Retrievabler)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching secret.Retrievabler")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// PasswordSearch: search for items in the secret service.
//
// The attributes should be a set of key and value string pairs.
//
// This method will return immediately and complete asynchronously.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for attributes.
//    - attributes: attribute keys and values.
//    - flags: search option flags.
//    - callback (optional): called when the operation completes.
//
func PasswordSearch(ctx context.Context, schema *Schema, attributes map[string]string, flags SearchFlags, callback gio.AsyncReadyCallback) {
	var _arg4 *C.GCancellable       // out
	var _arg1 *C.SecretSchema       // out
	var _arg2 *C.GHashTable         // out
	var _arg3 C.SecretSearchFlags   // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	_arg3 = C.SecretSearchFlags(flags)
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.secret_password_searchv(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// PasswordSearchSync: search for items in the secret service.
//
// The attributes should be a set of key and value string pairs.
//
// If no secret is found then NULL is returned.
//
// This method may block indefinitely and should not be used in user interface
// threads.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for attributes.
//    - attributes: attribute keys and values.
//    - flags: search option flags.
//
// The function returns the following values:
//
//    - list of Retrievable containing attributes of the matched items.
//
func PasswordSearchSync(ctx context.Context, schema *Schema, attributes map[string]string, flags SearchFlags) ([]Retrievabler, error) {
	var _arg4 *C.GCancellable     // out
	var _arg1 *C.SecretSchema     // out
	var _arg2 *C.GHashTable       // out
	var _arg3 C.SecretSearchFlags // out
	var _cret *C.GList            // in
	var _cerr *C.GError           // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		defer C.free(unsafe.Pointer(vdst))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg2)
	_arg3 = C.SecretSearchFlags(flags)

	_cret = C.secret_password_searchv_sync(_arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(flags)

	var _list []Retrievabler // out
	var _goerr error         // out

	_list = make([]Retrievabler, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.SecretRetrievable)(v)
		var dst Retrievabler // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type secret.Retrievabler is nil")
			}

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Retrievabler)
				return ok
			})
			rv, ok := casted.(Retrievabler)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching secret.Retrievabler")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// PasswordStoreFinish: finish asynchronous operation to store a password in the
// secret service.
//
// The function takes the following parameters:
//
//    - result asynchronous result passed to the callback.
//
func PasswordStoreFinish(result gio.AsyncResulter) error {
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.secret_password_store_finish(_arg1, &_cerr)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PasswordStore: store a password in the secret service.
//
// The attributes should be a set of key and value string pairs.
//
// If the attributes match a secret item already stored in the collection, then
// the item will be updated with these new values.
//
// If collection is NULL, then the default collection will be used. Use
// CRET_COLLECTION_SESSION to store the password in the session collection,
// which doesn't get stored across login sessions.
//
// This method will return immediately and complete asynchronously.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for attributes.
//    - attributes: attribute keys and values.
//    - collection (optional) alias, or D-Bus object path of the collection where
//      to store the secret.
//    - label for the secret.
//    - password: null-terminated password to store.
//    - callback (optional): called when the operation completes.
//
func PasswordStore(ctx context.Context, schema *Schema, attributes map[string]string, collection, label, password string, callback gio.AsyncReadyCallback) {
	var _arg6 *C.GCancellable       // out
	var _arg1 *C.SecretSchema       // out
	var _arg2 *C.GHashTable         // out
	var _arg3 *C.gchar              // out
	var _arg4 *C.gchar              // out
	var _arg5 *C.gchar              // out
	var _arg7 C.GAsyncReadyCallback // out
	var _arg8 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	if collection != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(collection)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg5))
	if callback != nil {
		_arg7 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg8 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.secret_password_storev(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(collection)
	runtime.KeepAlive(label)
	runtime.KeepAlive(password)
	runtime.KeepAlive(callback)
}

// PasswordStoreBinary: store a password in the secret service.
//
// This is similar to secret_password_storev(), but takes a Value as the
// argument instead of a null-terminated password.
//
// This method will return immediately and complete asynchronously.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for attributes.
//    - attributes: attribute keys and values.
//    - collection (optional) alias, or D-Bus object path of the collection where
//      to store the secret.
//    - label for the secret.
//    - value: Value.
//    - callback (optional): called when the operation completes.
//
func PasswordStoreBinary(ctx context.Context, schema *Schema, attributes map[string]string, collection, label string, value *Value, callback gio.AsyncReadyCallback) {
	var _arg6 *C.GCancellable       // out
	var _arg1 *C.SecretSchema       // out
	var _arg2 *C.GHashTable         // out
	var _arg3 *C.gchar              // out
	var _arg4 *C.gchar              // out
	var _arg5 *C.SecretValue        // out
	var _arg7 C.GAsyncReadyCallback // out
	var _arg8 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	if collection != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(collection)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.SecretValue)(gextras.StructNative(unsafe.Pointer(value)))
	if callback != nil {
		_arg7 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg8 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.secret_password_storev_binary(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(collection)
	runtime.KeepAlive(label)
	runtime.KeepAlive(value)
	runtime.KeepAlive(callback)
}

// PasswordStoreBinarySync: store a password in the secret service.
//
// This is similar to secret_password_storev_sync(), but takes a Value as the
// argument instead of a null-terminated passwords.
//
// This method may block indefinitely and should not be used in user interface
// threads.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for attributes.
//    - attributes: attribute keys and values.
//    - collection (optional) alias, or D-Bus object path of the collection where
//      to store the secret.
//    - label for the secret.
//    - value: Value.
//
func PasswordStoreBinarySync(ctx context.Context, schema *Schema, attributes map[string]string, collection, label string, value *Value) error {
	var _arg6 *C.GCancellable // out
	var _arg1 *C.SecretSchema // out
	var _arg2 *C.GHashTable   // out
	var _arg3 *C.gchar        // out
	var _arg4 *C.gchar        // out
	var _arg5 *C.SecretValue  // out
	var _cerr *C.GError       // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		defer C.free(unsafe.Pointer(vdst))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg2)
	if collection != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(collection)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.SecretValue)(gextras.StructNative(unsafe.Pointer(value)))

	C.secret_password_storev_binary_sync(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(collection)
	runtime.KeepAlive(label)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PasswordStoreSync: store a password in the secret service.
//
// The attributes should be a set of key and value string pairs.
//
// If the attributes match a secret item already stored in the collection, then
// the item will be updated with these new values.
//
// If collection is NULL, then the default collection will be used. Use
// CRET_COLLECTION_SESSION to store the password in the session collection,
// which doesn't get stored across login sessions.
//
// This method may block indefinitely and should not be used in user interface
// threads.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional cancellation object.
//    - schema (optional) for attributes.
//    - attributes: attribute keys and values.
//    - collection (optional) alias, or D-Bus object path of the collection where
//      to store the secret.
//    - label for the secret.
//    - password: null-terminated password to store.
//
func PasswordStoreSync(ctx context.Context, schema *Schema, attributes map[string]string, collection, label, password string) error {
	var _arg6 *C.GCancellable // out
	var _arg1 *C.SecretSchema // out
	var _arg2 *C.GHashTable   // out
	var _arg3 *C.gchar        // out
	var _arg4 *C.gchar        // out
	var _arg5 *C.gchar        // out
	var _cerr *C.GError       // in

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if schema != nil {
		_arg1 = (*C.SecretSchema)(gextras.StructNative(unsafe.Pointer(schema)))
	}
	_arg2 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range attributes {
		var kdst *C.gchar // out
		var vdst *C.gchar // out
		kdst = (*C.gchar)(unsafe.Pointer(C.CString(ksrc)))
		defer C.free(unsafe.Pointer(kdst))
		vdst = (*C.gchar)(unsafe.Pointer(C.CString(vsrc)))
		defer C.free(unsafe.Pointer(vdst))
		C.g_hash_table_insert(_arg2, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg2)
	if collection != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(collection)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg5))

	C.secret_password_storev_sync(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, &_cerr)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(schema)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(collection)
	runtime.KeepAlive(label)
	runtime.KeepAlive(password)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PasswordWipe: clear the memory used by a password.
//
// The function takes the following parameters:
//
//    - password (optional) to clear.
//
func PasswordWipe(password string) {
	var _arg1 *C.gchar // out

	if password != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(password)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.secret_password_wipe(_arg1)
	runtime.KeepAlive(password)
}
