package chess

import (
	"fmt"
	"strings"
	"unicode"
)

/*
*
https://www.chessprogramming.org/Bitboard_Board-Definition
*/
type Board struct {
	emptySquares    uint64
	occupiedSquares uint64
	piecesSquares   [12]uint64 // 0 : white king, 1 : white queen... 6 : black king, 7 : black queen....
	attackedSquares [12]uint64
	castling        byte // 0000 | BbWw
	enPassant       int
}

func (b Board) getColorPieces(c Color) uint64 {
	if c == White {
		return b.getWhitePieces()
	} else {
		return b.getBlackPieces()
	}
}

func (b Board) canWhiteQueenCastle() bool { return b.castling&2 == 1 }
func (b Board) canWhiteKingCastle() bool  { return b.castling&1 == 1 }
func (b Board) canBlackQueenCastle() bool { return b.castling&4 == 1 }
func (b Board) canBlackKingCastle() bool  { return b.castling&8 == 1 }

func (b Board) getWhiteKing() uint64    { return b.piecesSquares[0] }
func (b Board) getWhiteQueens() uint64  { return b.piecesSquares[1] }
func (b Board) getWhiteRooks() uint64   { return b.piecesSquares[2] }
func (b Board) getWhiteBishops() uint64 { return b.piecesSquares[3] }
func (b Board) getWhiteKnights() uint64 { return b.piecesSquares[4] }
func (b Board) getWhitePawns() uint64   { return b.piecesSquares[5] }
func (b Board) getBlackKing() uint64    { return b.piecesSquares[6] }
func (b Board) getBlackQueens() uint64  { return b.piecesSquares[7] }
func (b Board) getBlackRooks() uint64   { return b.piecesSquares[8] }
func (b Board) getBlackBishops() uint64 { return b.piecesSquares[9] }
func (b Board) getBlackKnights() uint64 { return b.piecesSquares[10] }
func (b Board) getBlackPawns() uint64   { return b.piecesSquares[11] }

func (b Board) isWhiteChecked() bool { return b.getSquaresAttackedByBlack()&b.getWhiteKing() != 0 }
func (b Board) isBlackChecked() bool { return b.getSquaresAttackedByWhite()&b.getBlackKing() != 0 }

func (b *Board) setPieceSet(p PieceType, c Color, u uint64) { b.piecesSquares[int(p)+int(c)] = u }

func (b Board) getPieceSet(p PieceType, c Color) uint64 { return b.piecesSquares[int(p)+int(c)] }
func (b Board) getPieceAttackSet(p PieceType, c Color) uint64 {
	return b.attackedSquares[int(p)+int(c)]
}

func (b Board) getWhitePieces() uint64 {
	return b.piecesSquares[0] | b.piecesSquares[1] | b.piecesSquares[2] | b.piecesSquares[3] | b.piecesSquares[4] | b.piecesSquares[5]
}

func (b Board) getBlackPieces() uint64 {
	return b.piecesSquares[6] | b.piecesSquares[7] | b.piecesSquares[8] | b.piecesSquares[9] | b.piecesSquares[10] | b.piecesSquares[11]
}

func (b Board) getSquaresAttackedByWhite() uint64 {
	return b.attackedSquares[0] | b.attackedSquares[1] | b.attackedSquares[2] | b.attackedSquares[3] | b.attackedSquares[4] | b.attackedSquares[5]
}

func (b Board) getSquaresAttackedByBlack() uint64 {
	return b.attackedSquares[6] | b.attackedSquares[7] | b.attackedSquares[8] | b.attackedSquares[9] | b.attackedSquares[10] | b.attackedSquares[11]
}

func (b Board) getWhitePawnsAbleToPush() uint64 {
	return soutOne(b.emptySquares) & b.getWhitePawns()
}

func (b Board) getWhitePawnsAbleToDoublePush() uint64 {
	return soutOne(soutOne(b.emptySquares)) & soutOne(b.emptySquares) & b.getWhitePawns() & 0xFF00
}

func (b Board) getBlackPawnsAbleToPush() uint64 {
	return nortOne(b.emptySquares) & b.getBlackPawns()
}

func (b Board) getBlackPawnsAbleToDoublePush() uint64 {
	return nortOne(nortOne(b.emptySquares)) & nortOne(b.emptySquares) & b.getBlackPawns() & 0xFF000000000000
}

func (b Board) getPieceOfSquare(square int8) int8 {
	for i := 0; i < 12; i++ {
		if b.piecesSquares[i]&(1<<square) != 0 {
			return int8(i)
		}
	}
	return -1
}

func (b Board) getSquareAttackers(square int8) uint64 {
	knights := b.getBlackKnights() | b.getWhiteKnights()
	kings := b.getBlackKing() | b.getWhiteKing()
	queens := b.getBlackQueens() | b.getWhiteQueens()
	bishopQueens := b.getBlackBishops() | b.getWhiteBishops() | queens
	rookQueens := b.getBlackRooks() | b.getWhiteRooks() | queens

	return (ArrWhitePawnsAttack[square] & b.getWhitePawns()) |
		(ArrBlackPawnsAttack[square] & b.getBlackPawns()) |
		(ArrKnightsAttack(square) & knights) |
		(ArrKingAttack(square) & kings) |
		(bishopMoves(int(square), b.occupiedSquares) & bishopQueens) |
		(rookMoves(int(square), b.occupiedSquares) & rookQueens)
}

func (b Board) GetPinnedPieces(c Color) uint64 {
	kingSquare := bitScanForward(b.getPieceSet(King, c))
	pinned := uint64(0)
	pinner := xrayRookAttacks(b.occupiedSquares, b.getColorPieces(c), kingSquare)
	for k := pinner != 0; k; k = pinner != 0 {
		sq := bitScanForward(pinner)
		pinned |= inBetween[sq][kingSquare] & b.getColorPieces(c)
		pinner &= pinner - 1
	}
	pinner = xrayBishopAttacks(b.occupiedSquares, b.getColorPieces(c), kingSquare)
	for k := pinner != 0; k; k = pinner != 0 {
		sq := bitScanForward(pinner)
		pinned |= inBetween[sq][kingSquare] & b.getColorPieces(c)
		pinner &= pinner - 1
	}
	fmt.Println(VisualizeBitBoard(pinned))
	return pinned
}

func (b *Board) SetBoardFromFEN(fen string) {
	rows := strings.Split(fen, "/")
	j := 0
	for i := 7; i >= 0; i-- {
		for _, c := range rows[i] {
			if unicode.IsDigit(c) {
				j += int(c) - '0'
			} else {
				p := GetPieceTypeFromFen(c)
				cl := GetColorFromFen(c)
				ps := b.getPieceSet(p, cl) | (1 << j)
				b.setPieceSet(p, cl, ps)
				j += 1
			}
		}
	}
	b.occupiedSquares = b.getWhitePieces() | b.getBlackPieces()
	b.emptySquares = ^b.occupiedSquares
	b.enPassant = 8092
}

func VisualizeBitBoard(b uint64) string {
	s := ""
	for i := 0; i < 8; i++ {
		line := ""
		for j := 0; j < 8; j++ {
			lb := b & 1
			if lb == 0 {
				line += " . "
			} else {
				line += " X "
			}
			b >>= 1
		}
		s += Reverse(line)
		s += "\n"
	}

	return Reverse(s)
}
