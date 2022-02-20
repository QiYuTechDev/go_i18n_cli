package main

import (
	"go_i18n_cli/cli"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		cli.I18nHello(os.Args[1])
	} else {
		i18n := cli.GetPrinter()
		cli.I18nHello(i18n.Sprintf("i18n"))
	}
}
