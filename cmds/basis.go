package cmds

import (
	_ "embed"

	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//go:embed tools/basisu.exe
var embedBasisu []byte

// Basis
// basisu -ktx2 x.png
var Basis = &cli.Command{
	Name:  "basis",
	Usage: "将图片转成basis纹理图",
	Action: func(c *cli.Context) error {
		// ktx2 使用ktx2，内置纹理格式减少加载量
		// comp_level 压缩等级
		var args = []string{"-ktx2"}
		args = append(args, c.Args().Slice()...)
		sys.MemExec(embedBasisu, args...)
		return nil
	},
}
