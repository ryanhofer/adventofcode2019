package geom

import "fmt"

// Vec is a vector in 2D space.
type Vec struct {
	X, Y int
}

// Add sums two vectors.
func (v Vec) Add(w Vec) Vec {
	v.X += w.X
	v.Y += w.Y
	return v
}

func (v Vec) String() string {
	return fmt.Sprintf("<%d,%d>", v.X, v.Y)
}
