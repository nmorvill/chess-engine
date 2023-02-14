package chess

import "unicode"

type Piece int

const (
	King   Piece = 1
	Queen  Piece = 2
	Rook   Piece = 3
	Bishop Piece = 4
	Knight Piece = 5
	Pawn   Piece = 6
)

type Color int

const (
	White Color = 8
	Black Color = 16
)

func GetPieceCodeFromFEN(c rune) int {
	piecesCode := map[rune]Piece{
		'k': King,
		'q': Queen,
		'r': Rook,
		'b': Bishop,
		'n': Knight,
		'p': Pawn,
	}

	ret := int(piecesCode[unicode.ToLower(c)])

	if unicode.IsLower(c) {
		ret += 6
	}

	return ret
}

func getOppositeColor(color Color) Color {
	if color == White {
		return Black
	} else {
		return White
	}
}
