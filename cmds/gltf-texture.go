package cmds

import (
	_ "embed"

	"github.com/mocheer/xena/pkg/tileset/transform"
	"github.com/urfave/cli/v2"
)

// GltfTexture
var GltfTexture = &cli.Command{
	Name:  "gltf-texture",
	Usage: "获取gltf模型的纹理图片",
	Action: func(c *cli.Context) error {
		args := c.Args().Slice()
		t := transform.FromPath(args[0])
		err := t.RunGetTextures()
		if err != nil {
			return err
		}
		return nil
	},
}
