# gotk4-secret

[Package documentation](https://pkg.go.dev/github.com/diamondburned/gotk4-secret/pkg/secret)

Go generated bindings for [libsecret][secret].

[secret]: https://developer.gnome.org/libsecret/0.20/

## Status: Broken

This package is currently broken, as `secret.NewSchema` cannot yet be generated.
This function won't be generated until full `GHashTable` support is added into
[gotk4][gotk4].

[gotk4]: https://github.com/diamondburned/gotk4
