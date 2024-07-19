package cmds

import (
	"context"
	"log"

	tiler "github.com/mfbonfigli/gocesiumtiler/v2"
	"github.com/urfave/cli/v2"
)

// PointCloud
var PointCloud = &cli.Command{
	Name:  "point-cloud",
	Usage: "从点云数据中生成3dtiles模型",
	Action: func(c *cli.Context) error {
		t, err := tiler.NewGoCesiumTiler()
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.TODO()
		err = t.ProcessFiles([]string{"myinput.las"}, "/tmp/myoutput", 32632, tiler.NewTilerOptions(
			tiler.WithEightBitColors(true),
			tiler.WithElevationOffset(34),
			tiler.WithWorkerNumber(2),
			tiler.WithMaxDepth(5),
		), ctx)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	},
}
