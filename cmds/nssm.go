package cmds

import (
	_ "embed"
	"fmt"

	"github.com/mocheer/pluto/cmd"
	"github.com/mocheer/pluto/fs"
	"github.com/urfave/cli/v2"
)

//go:embed tools/nssm.exe
var nssmBytes []byte

var nssmExePath string = `.nix/nssm.exe`

// nssm install {appName} {execPath}/charon.exe
// nssm start {appName}
// nssm stop {appName}
//
var NSSM = &cli.Command{
	Name:  "nssm",
	Usage: "执行nssm.exe注册windows服务",
	Action: func(c *cli.Context) error {
		isExist := fs.IsExist(nssmExePath)
		if !isExist {
			fs.SaveFile(nssmExePath, nssmBytes)
		}
		result := cmd.Exec(nssmExePath, c.Args().Slice()...)
		fmt.Println(result)
		return nil
	},
}
