package cmds

import (
	_ "embed"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/mocheer/pluto/pkg/ts/clock"
	"github.com/urfave/cli/v2"
)

// active
var Active = &cli.Command{
	Name:  "active",
	Usage: "模拟鼠标和键盘操作，使电脑保持活动状态",
	Action: func(c *cli.Context) error {
		clock.SetInterval(moveTo, time.Second*30, false)
		select {}
	},
}

func moveTo() {
	x := rand.Intn(4) - 2
	y := rand.Intn(4) - 2
	fmt.Println(x, y)
	robotgo.MoveSmoothRelative(x, y)
}
