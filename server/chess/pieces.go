package chess

import (
	"fmt"
	"unicode"
)

type PieceType int8
type Color int8

const (
	King   PieceType = 0
	Queen  PieceType = 1
	Rook   PieceType = 2
	Bishop PieceType = 3
	Knight PieceType = 4
	Pawn   PieceType = 5
)

const (
	White Color = 0
	Black Color = 6
)

var piecesName = [...]string{"King", "Queen", "Rook", "Bishop", "Knight", "Pawn"}

func GetPieceTypeFromFen(c rune) PieceType {
	piecesCode := map[rune]PieceType{
		'k': King,
		'q': Queen,
		'r': Rook,
		'b': Bishop,
		'n': Knight,
		'p': Pawn,
	}

	return piecesCode[unicode.ToLower(c)]
}

func GetColorFromFen(c rune) Color {
	if unicode.IsLower(c) {
		return Black
	}
	return White
}

func getOppositeColor(color Color) Color {
	if color == White {
		return Black
	} else {
		return White
	}
}

func (p PieceType) String() string {
	return fmt.Sprintf(piecesName[p])
}
