package gamelogic

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

type Result int

const (
	Lose Result = iota
	Win
	Draw
)

func Play(PlayerMove Move, oponentMove Move) Result {
	if PlayerMove == oponentMove {
		return Draw
	}
	if PlayerMove == Rock && oponentMove == Scissors {
		return Win
	}
	if PlayerMove == Paper && oponentMove == Rock {
		return Win
	}
	if PlayerMove == Scissors && oponentMove == Paper {
		return Win
	}
	return Lose
}
