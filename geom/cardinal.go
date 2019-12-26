package geom

// Cardinal is a cardinal compass direction.
type Cardinal int

// Enumerated values.
const (
	North Cardinal = iota
	East
	South
	West
)

var Cardinals = []Cardinal{
	North,
	East,
	South,
	West,
}

// RotateCW returns c rotated 90 degrees clockwise.
func (c Cardinal) RotateCW() Cardinal {
	return (c + 1) % 4
}

// RotateCCW returns c rotated 90 degrees counterclockwise.
func (c Cardinal) RotateCCW() Cardinal {
	return (c + 3) % 4
}

// Vec returns the unit vector of a Cardinal.
func (c Cardinal) Vec() Vec {
	switch c {
	case North:
		return Vec{0, -1}
	case East:
		return Vec{1, 0}
	case South:
		return Vec{0, 1}
	case West:
		return Vec{-1, 0}
	default:
		return Vec{}
	}
}

func (c Cardinal) String() string {
	switch c {
	case North:
		return "N"
	case East:
		return "E"
	case South:
		return "S"
	case West:
		return "W"
	default:
		return "?"
	}
}
