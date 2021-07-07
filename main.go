package main

import (
	"log"
	"os"

	"github.com/mocheer/nix/cmds"
	"github.com/urfave/cli/v2"
)

func main() {
	//
	app := &cli.App{
		Commands: []*cli.Command{
			cmds.Dev,
			cmds.Build,
			cmds.Serve,
			cmds.Fs,
			cmds.Run,
			cmds.Exec,
			cmds.NSSM,
			cmds.Upx,
			cmds.Tif,
			cmds.Scoop,
			cmds.Rsa,
			cmds.Tileset,
			cmds.Struct,
			cmds.About,
		},
	}
	//
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
