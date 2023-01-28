package main

import (
	_ "embed"

	"github.com/metaplayer-one/metad/command/root"
	"github.com/metaplayer-one/metad/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
