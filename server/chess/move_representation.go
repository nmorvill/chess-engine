package chess

import "fmt"

/*
We represent the move by a int16 with the 10  first bits representing from and
to position and the last 6 bits representing diverse information about the move itself

	Info	  To	  From
	XXXX | 000000 | 000000
*/
type Move struct {
	moveCode  uint16
	piece     PieceType
	enPassant int
}

var squares = [64]string{
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
}

func (m Move) GetFrom() int8 { return int8(m.moveCode & 63) }
func (m Move) GetTo() int8   { return int8((m.moveCode >> 6) & 63) }

func (m *Move) setFrom(from int8) { m.moveCode = m.moveCode | uint16(from) }
func (m *Move) setTo(to int8)     { m.moveCode = m.moveCode | uint16(to)<<6 }

func createMove(from int8, to int8, piece PieceType) Move {
	var m Move
	m.moveCode = uint16(from) | uint16(to)<<6
	m.piece = piece
	m.enPassant = 8092
	return m
}

func (m Move) String() string {
	return fmt.Sprintf("%s%s", squares[m.GetFrom()], squares[m.GetTo()])
}
