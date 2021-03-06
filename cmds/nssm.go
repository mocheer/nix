package cmds

import (
	_ "embed"
	"path"

	"github.com/mocheer/nix/global"
	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//go:embed tools/nssm.exe
var nssmBytes []byte

// nssm install {appName} {execPath}/charon.exe
// nssm start {appName}
// nssm stop {appName}
//
var NSSM = &cli.Command{
	Name:  "nssm",
	Usage: "执行nssm.exe注册windows服务",
	Action: func(c *cli.Context) error {
		nssmExePath := path.Join(global.ExportDir, "nssm.exe")
		isExist := ds.IsExist(nssmExePath)
		if !isExist {
			ds.Save(nssmExePath, nssmBytes)
		}
		sys.Exec(nssmExePath, c.Args().Slice()...)
		return nil
	},
}
