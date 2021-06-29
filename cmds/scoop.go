package cmds

import (
	_ "embed"

	"github.com/mocheer/pluto/sys"
	"github.com/urfave/cli/v2"
)

var Scoop = &cli.Command{
	Name:  "scoop",
	Usage: "安装scoop",
	Action: func(c *cli.Context) error {
		// Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')
		// wr -useb get.scoop.sh | iex
		// 这里执行失败，但命令行执行成功，待排查
		sys.Shell("Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')")
		return nil
	},
}
