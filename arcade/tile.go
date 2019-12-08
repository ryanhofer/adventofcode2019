package arcade

type Tile int

const (
	Empty Tile = iota
	Wall
	Block
	HorizontalPaddle
	Ball
)

func (t Tile) String() string {
	switch t {
	case Empty:
		return " "
	case Wall:
		return "+"
	case Block:
		return "#"
	case HorizontalPaddle:
		return "-"
	case Ball:
		return "o"
	default:
		return "?"
	}
}
