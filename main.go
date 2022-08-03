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
			cmds.Version,
			cmds.Exec,
			cmds.NSSM,
			cmds.Upx,
			cmds.Tif,
			cmds.Scoop,
			cmds.Rsa,
			cmds.Tile,
			cmds.TileXYZ,
			cmds.TileArcgis,
			cmds.Tileset,
			cmds.TilesetDem,
			cmds.Struct,
			cmds.About,
			cmds.Package,
			cmds.Esbuild,
			cmds.Publish,
			cmds.Install,
			cmds.Init,
		},
	}
	//
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
