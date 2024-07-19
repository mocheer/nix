package main

import (
	"context"
	"log"

	tiler "github.com/mfbonfigli/gocesiumtiler/v2"
)

func main() {
	t, err := tiler.NewGoCesiumTiler()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.TODO()
	err = t.ProcessFiles([]string{"testdata/cloud.ply"}, "/tmp/myoutput", 32632, tiler.NewTilerOptions(
		tiler.WithEightBitColors(true),
		tiler.WithElevationOffset(34),
		tiler.WithWorkerNumber(2),
		tiler.WithMaxDepth(5),
	), ctx)
	if err != nil {
		log.Fatal(err)
	}
}
