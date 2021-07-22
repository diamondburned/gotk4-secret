package main

import (
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module   = "github.com/diamondburned/gotk4/pkg"
	adwaitaModule = "github.com/diamondburned/gotk4-adwaita/pkg"
)

var packages = []gendata.Package{
	{PkgName: "libsecret-1", Namespaces: nil},
}

var pkgGenerated = []string{
	"secret",
}

var pkgExceptions = []string{
	"go.mod",
	"go.sum",
	"LICENSE",
}

var filters = []types.FilterMatcher{
	// Not found.
	types.FileFilter("secret-backend."),
	types.AbsoluteFilter("Secret-1.Service.encode_dbus_secret"),
	types.AbsoluteFilter("Secret-1.Service.decode_dbus_secret"),
	types.AbsoluteFilter("Secret-1.Service.get_session_dbus_path"),
	types.AbsoluteFilter("Secret-1.Service.create_item_dbus_path_sync"),
}
