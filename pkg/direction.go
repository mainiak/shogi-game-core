package game

type Direction uint8

const (
	NOPE      Direction = 0               // for errors only
	North     Direction = 1 << (iota - 1) // start with n=1, for 2^n
	NorthEast                             //Direction
	East                                  //Direction
	SouthEast                             //Direction
	South                                 //Direction
	SouthWest                             //Direction
	West                                  //Direction
	NorthWest                             //Direction
)

func (d Direction) String() string {
	switch d {

	case North:
		return "North"

	case NorthEast:
		return "North East"

	case East:
		return "East"

	case SouthEast:
		return "South East"

	case South:
		return "South"

	case SouthWest:
		return "South West"

	case West:
		return "West"

	case NorthWest:
		return "North West"

	default:
		return "Error: Invalid Direction type"
	}
}
