package main

import (
	"github.com/thiagothalves/solid-go/dip"
	"github.com/thiagothalves/solid-go/isp"
	"github.com/thiagothalves/solid-go/lsp"
	"github.com/thiagothalves/solid-go/ocp"
	"github.com/thiagothalves/solid-go/srp"
)

func main() {
	srp.Run()
	ocp.Run()
	lsp.Run()
	isp.Run()
	dip.Run()
}
