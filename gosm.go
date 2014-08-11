package gosm

import (
	"fmt"
	"math"
)

type Tile struct {
	Z    int
	X    int
	Y    int
	Lat  float64
	Long float64
}

type BBox struct {
	TopLeftTile     Tile
	BottomRightTile Tile
}

type Conversion interface {
	Deg2num(t *Tile) (x int, y int)
	Num2deg(t *Tile) (lat float64, long float64)
}

func (*Tile) Deg2num(t *Tile) (x int, y int) {
	x = int(math.Floor((t.Long + 180.0) / 360.0 * (math.Exp2(float64(t.Z)))))
	y = int(math.Floor((1.0 - math.Log(math.Tan(t.Lat*math.Pi/180.0)+1.0/math.Cos(t.Lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(t.Z)))))
	return
}

func (*Tile) Num2deg(t *Tile) (lat float64, long float64) {
	n := math.Pi - 2.0*math.Pi*float64(t.Y)/math.Exp2(float64(t.Z))
	lat = 180.0 / math.Pi * math.Atan(0.5*(math.Exp(n)-math.Exp(-n)))
	long = float64(t.X)/math.Exp2(float64(t.Z))*360.0 - 180.0
	return lat, long
}

func NewTileWithLatLong(lat float64, long float64, z int) (t *Tile) {
	t = new(Tile)
	t.Lat = lat
	t.Long = long
	t.Z = z
	t.X, t.Y = t.Deg2num(t)
	return
}

func NewTileWithXY(x int, y int, z int) (t *Tile) {
	t = new(Tile)
	t.Z = z
	t.X = x
	t.Y = y
	t.Lat, t.Long = t.Num2deg(t)
	return
}

func BBoxTiles(bbox BBox) ([]*Tile, error) {

	if bbox.TopLeftTile.X > bbox.BottomRightTile.X || bbox.TopLeftTile.Y > bbox.BottomRightTile.Y {
		return nil, fmt.Errorf("Your bbox is not correct")
	}

	nbtiles := ((bbox.BottomRightTile.X - bbox.TopLeftTile.X) + 1) * ((bbox.BottomRightTile.Y - bbox.TopLeftTile.Y) + 1)

	tiles := make([]*Tile, nbtiles) // ? + 2

	for x, i := 0, 0; x <= (bbox.BottomRightTile.X - bbox.TopLeftTile.X); x++ {
		for y := 0; y <= (bbox.BottomRightTile.Y - bbox.TopLeftTile.Y); y++ {
			tiles[i] = NewTileWithXY(bbox.BottomRightTile.X+x, bbox.BottomRightTile.Y+y, bbox.BottomRightTile.Z)
			i += 1
		}
	}
	return tiles, nil
}

func PngBBoxTiles(bbox BBox) ([]*Tile, error) {

	whith := (bbox.BottomRightTile.X - bbox.TopLeftTile.X) * 256
	height := (bbox.BottomRightTile.Y - bbox.TopLeftTile.Y) * 256

	if bbox.TopLeftTile.X > bbox.BottomRightTile.X || bbox.TopLeftTile.Y > bbox.BottomRightTile.Y {
		return nil, fmt.Errorf("Your bbox is not correct")
	}

	nbtiles := ((bbox.BottomRightTile.X - bbox.TopLeftTile.X) + 1) * ((bbox.BottomRightTile.Y - bbox.TopLeftTile.Y) + 1)

	tiles := make([]*Tile, nbtiles) // ? + 2

	for x, i := 0, 0; x <= (bbox.BottomRightTile.X - bbox.TopLeftTile.X); x++ {
		for y := 0; y <= (bbox.BottomRightTile.Y - bbox.TopLeftTile.Y); y++ {
			tiles[i] = NewTileWithXY(bbox.BottomRightTile.X+x, bbox.BottomRightTile.Y+y, bbox.BottomRightTile.Z)
			i += 1
		}
	}
	return tiles, nil
}
