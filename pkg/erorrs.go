package game

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	ErrInvalidMove          = constError("Invalid Move")
	ErrInvalidMoveDirection = constError("Invalid Move Direction")
	ErrInvalidTurnOwner     = constError("Invalid Turn Owner")
	ErrAxisOutOfRange       = constError("Invalid Axis range")
	ErrImplementationIssue  = constError("Implementation Issue")
	ErrWIP                  = constError("Work In Progress")
)
