package main

import (
	"fmt"
	"github.com/j4/gosm"
)

func main() {
	t1 := gosm.NewTileWithLatLong(48.8165, 2.3216, 10)
	t2 := gosm.NewTileWithLatLong(48.7992, 2.2346, 10)
	x1, y1 := t1.Deg2num()
	lat1, long2 := t1.Num2deg()
	fmt.Printf("%d %d %f %f\n", x1, y1, lat1, long2)
	tiles, _ := gosm.BBoxTiles(*t1, *t2)
	for _, t := range tiles {
		fmt.Printf("http://a.tile.openstreetmap.org/%d/%d/%d.png\n", t.Z, t.X, t.Y)
	}
}
