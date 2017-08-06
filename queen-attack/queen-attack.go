package queenattack

import (
	"errors"
)

const testVersion = 2

// CanQueenAttack determines if queens at given potions can attack themself.
func CanQueenAttack(p1, p2 string) (bool, error) {
	if len(p1) != 2 || len(p2) != 2 {
		return false, errors.New("Invalid piece position")
	}

	x1, y1 := p1[0]-'a', p1[1]-'1'
	x2, y2 := p2[0]-'a', p2[1]-'1'

	if x1 < 0 || y1 < 0 || x2 < 0 || y2 < 0 || x1 > 7 || y1 > 7 || x2 > 7 || y2 > 7 {
		return false, errors.New("Piece position out of board")
	}

	dx := x1 - x2
	dy := y1 - y2

	dx *= dx
	dy *= dy

	if dx == 0 && dy == 0 {
		return false, errors.New("Pieces at the same position")
	}

	return (dx*dy == 0 && dx != dy) || (dx*dy != 0 && dx == dy), nil
}
