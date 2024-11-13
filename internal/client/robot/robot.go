package robot

import "github.com/go-vgo/robotgo"

type Pos struct {
	x int
	y int
}

var dpi float64

func NewPos(x, y int) Pos {
	return Pos{x: x, y: y}
}

func init() {
	robotgo.Move(1000, 10)
	x, _ := robotgo.Location()
	dpi = float64(x) / 1000
}
