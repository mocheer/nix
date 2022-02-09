package cmds

import (
	"fmt"
	"os"

	"github.com/mocheer/pluto/pkg/sys"
	"github.com/urfave/cli/v2"
)

//
var About = &cli.Command{
	Name:  "about",
	Usage: "相关信息",
	Action: func(c *cli.Context) error {
		exePath, _ := sys.GetExePath()
		execPath, _ := os.Getwd()
		fmt.Printf(`
%c[1;31;40m---关于---%c[0m
执行程序的所在位置：%s
当前程序执行位置：%s`, 0x1B, 0x1B, exePath, execPath)
		return nil
	},
}
