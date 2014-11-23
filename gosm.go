package gosm

import (
	"fmt"
	"log"
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

func (t *Tile) Deg2num() (x int, y int) {
	x = int(math.Floor((t.Long + 180.0) / 360.0 * (math.Exp2(float64(t.Z)))))
	y = int(math.Floor((1.0 - math.Log(math.Tan(t.Lat*math.Pi/180.0)+1.0/math.Cos(t.Lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(t.Z)))))
	return
}

func (t *Tile) Num2deg() (lat float64, long float64) {
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

func BBoxTiles(topTile Tile, bottomTile Tile) ([]*Tile, error) {

	if topTile.X > bottomTile.X || topTile.Y > bottomTile.Y {
		return nil, fmt.Errorf("Your bbox is not correct")
	}

	nbtiles := ((bottomTile.X - topTile.X) + 1) * ((bottomTile.Y - topTile.Y) + 1)

	log.Printf("%d", nbtiles)

	tiles := make([]*Tile, nbtiles) // ? + 2

	for x, i := 0, 0; x <= (bottomTile.X - topTile.X); x++ {
		for y := 0; y <= (bottomTile.Y - topTile.Y); y++ {
			tiles[i] = NewTileWithXY(bottomTile.X+x, bottomTile.Y+y, bottomTile.Z)
			i += 1
		}
	}
	return tiles, nil
}
