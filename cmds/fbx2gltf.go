package cmds

import (
	_ "embed"

	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//go:embed tools/FBX2glTF.exe
var embedFBX2glTF []byte

// @see https://github.com/facebookincubator/FBX2glTF
// FBX2glTF -b -d -i xx.fbx -o xx.glb
var FBX2glTF = &cli.Command{
	Name:  "FBX2glTF",
	Usage: "将fbx模型格式转换为gltf",
	Action: func(c *cli.Context) error {
		sys.MemExec(embedFBX2glTF, c.Args().Slice()...)
		return nil
	},
}
