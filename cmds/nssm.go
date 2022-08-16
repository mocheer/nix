package cmds

import (
	_ "embed"

	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//go:embed tools/nssm.exe
var embedNSSM []byte

// nssm install {appName} {execPath}/charon.exe
// nssm start {appName}
// nssm stop {appName}
var NSSM = &cli.Command{
	Name:  "nssm",
	Usage: "执行nssm.exe注册windows服务",
	Action: func(c *cli.Context) error {
		sys.MemExec(embedNSSM, c.Args().Slice()...)
		return nil
	},
}
