package chess

import (
	"math/bits"
)

/*
https://www.chessprogramming.org/Move_Generation
*/
type MoveGeneratorParameters struct {
}

/*
Generate all possible moves given some parameters
*/
func MoveGenerator(b Board, c Color) []Move {

	var ret []Move

	ret = append(ret, allLegalMovesGenerator(b, c)...)

	return ret
}

func allLegalMovesGenerator(b Board, c Color) []Move {

	var allLegalMoves []Move
	kingThreaths := attacksToKing(b, c)

	if kingThreaths == 0 {
		allLegalMoves = append(allLegalMoves, generateMovesForAllPieces(b, c, all1)...)
	} else {
		nbChecks := bits.OnesCount64(kingThreaths)
		allLegalMoves = append(allLegalMoves, kingEscapeMoveGenerator(b, c)...)
		if nbChecks == 1 {
			allLegalMoves = append(allLegalMoves, pieceInterposingMoveGenerator(b, c, kingThreaths)...)
			allLegalMoves = append(allLegalMoves, generateMovesForAllPieces(b, c, kingThreaths)...)
		}
	}

	return allLegalMoves
}

func pieceInterposingMoveGenerator(b Board, c Color, kingThreaths uint64) []Move {

	kingThreatSquare := bitScanForward(b.getPieceSet(King, c))
	attackingCheckSquare := bitScanForward(b.getPieceSet(King, c))
	possibleTo := inBetween[kingThreatSquare][attackingCheckSquare]

	return generateMovesForAllPieces(b, c, possibleTo)
}

func kingEscapeMoveGenerator(b Board, c Color) []Move {
	var ret []Move
	kingMoves := generateKingMoves(b, c, all1)
	for _, m := range kingMoves {
		newBoard := MovePiece(b, m, c)
		if attacksToKing(newBoard, c) == 0 {
			ret = append(ret, m)
		}
	}
	return ret
}

func generateMovesForAllPieces(b Board, c Color, possibleSquares uint64) []Move {
	var ret []Move

	ret = append(ret, generateKnightsMoves(b, c, possibleSquares)...)
	ret = append(ret, generatePawnsPushes(b, c, possibleSquares)...)
	ret = append(ret, generateKingMoves(b, c, possibleSquares)...)
	ret = append(ret, generateBishopsMoves(b, c, possibleSquares)...)
	ret = append(ret, generateQueensMoves(b, c, possibleSquares)...)
	ret = append(ret, generateRooksMoves(b, c, possibleSquares)...)
	ret = append(ret, generatePawnsAttacks(b, c, possibleSquares)...)

	return ret
}

func generateKingMoves(b Board, c Color, pl uint64) []Move {
	var ret []Move
	kingSet := b.getPieceSet(King, c)
	for k := kingSet != 0; k; k = kingSet != 0 {
		from := bitScanForward(kingSet)
		possibleTo := ArrKingAttack(from) & ^b.getColorPieces(c) & pl
		for l := possibleTo != 0; l; l = possibleTo != 0 {
			to := bitScanForward(possibleTo)
			ret = append(ret, createMove(from, to, Knight))
			possibleTo &= possibleTo - 1
		}
		kingSet &= kingSet - 1
	}
	return ret
}

/*
Generate all the possible knight moves
*/
func generateKnightsMoves(b Board, c Color, pl uint64) []Move {
	var ret []Move
	knightsSet := b.getPieceSet(Knight, c)
	for k := knightsSet != 0; k; k = knightsSet != 0 {
		from := bitScanForward(knightsSet)
		possibleTo := ArrKnightsAttack(from) & ^b.getColorPieces(c) & pl
		for l := possibleTo != 0; l; l = possibleTo != 0 {
			to := bitScanForward(possibleTo)
			ret = append(ret, createMove(from, to, Knight))
			possibleTo &= possibleTo - 1
		}
		knightsSet &= knightsSet - 1
	}
	return ret
}

func generateBishopsMoves(b Board, c Color, pl uint64) []Move {
	var ret []Move
	bishops := b.getPieceSet(Bishop, c)
	for k := bishops != 0; k; k = bishops != 0 {
		from := bitScanForward(bishops)
		possibleTo := bishopMoves(int(from), b.occupiedSquares) & ^b.getColorPieces(c) & pl
		for l := possibleTo != 0; l; l = possibleTo != 0 {
			to := bitScanForward(possibleTo)
			ret = append(ret, createMove(from, to, Bishop))
			possibleTo &= possibleTo - 1
		}
		bishops &= bishops - 1
	}
	return ret
}

func generateRooksMoves(b Board, c Color, pl uint64) []Move {
	var ret []Move
	rooks := b.getPieceSet(Rook, c)
	for k := rooks != 0; k; k = rooks != 0 {
		from := bitScanForward(rooks)
		possibleTo := rookMoves(int(from), b.occupiedSquares) & ^b.getColorPieces(c) & pl
		for l := possibleTo != 0; l; l = possibleTo != 0 {
			to := bitScanForward(possibleTo)
			ret = append(ret, createMove(from, to, Rook))
			possibleTo &= possibleTo - 1
		}
		rooks &= rooks - 1
	}
	return ret
}

func generateQueensMoves(b Board, c Color, pl uint64) []Move {
	var ret []Move
	queens := b.getPieceSet(Queen, c)
	for k := queens != 0; k; k = queens != 0 {
		from := bitScanForward(queens)
		possibleTo := queenMoves(int(from), b.occupiedSquares) & ^b.getColorPieces(c) & pl
		for l := possibleTo != 0; l; l = possibleTo != 0 {
			to := bitScanForward(possibleTo)
			ret = append(ret, createMove(from, to, Queen))
			possibleTo &= possibleTo - 1
		}
		queens &= queens - 1
	}
	return ret
}

func generatePawnsAttacks(b Board, c Color, pl uint64) []Move {
	var ret []Move
	pawns := b.getPieceSet(Pawn, c)
	for k := pawns != 0; k; k = pawns != 0 {
		from := bitScanForward(pawns)
		var possibleTo uint64
		if c == White {
			possibleTo = ArrWhitePawnsAttack[from] & (b.getColorPieces(Black) | uint64(1<<b.enPassant)) & pl
		} else {
			possibleTo = ArrBlackPawnsAttack[from] & (b.getColorPieces(White) | uint64(1<<b.enPassant)) & pl
		}
		for l := possibleTo != 0; l; l = possibleTo != 0 {
			to := bitScanForward(possibleTo)
			ret = append(ret, createMove(from, to, Pawn))
			possibleTo &= possibleTo - 1
		}
		pawns &= pawns - 1
	}

	return ret
}

/*
Generate pawn pushes based on color
*/
func generatePawnsPushes(b Board, c Color, pl uint64) []Move {
	if c == White {
		return generateWhitePawnsPushes(b, pl)
	}
	return generateBlackPawnsPushes(b, pl)
}

/*
Generate all the possible white pawns pushes
*/
func generateWhitePawnsPushes(b Board, pl uint64) []Move {
	var ret []Move
	push := b.getWhitePawnsAbleToPush() & pl
	for k := push != 0; k; k = push != 0 {
		from := bitScanForward(push)
		ret = append(ret, createMove(from, from+8, Pawn))
		push &= push - 1
	}
	doublePushs := b.getWhitePawnsAbleToDoublePush() & pl
	for k := doublePushs != 0; k; k = doublePushs != 0 {
		from := bitScanForward(doublePushs)
		move := createMove(from, from+16, Pawn)
		move.enPassant = int(from) + 8
		ret = append(ret, move)
		doublePushs &= doublePushs - 1
	}
	return ret
}

/*
Generate all the possible black pawns pushes
*/
func generateBlackPawnsPushes(b Board, pl uint64) []Move {
	var ret []Move
	push := b.getBlackPawnsAbleToPush() & pl
	for k := push != 0; k; k = push != 0 {
		from := bitScanForward(push)
		ret = append(ret, createMove(from, from-8, Pawn))
		push &= push - 1
	}
	doublePushs := b.getBlackPawnsAbleToDoublePush() & pl
	for k := doublePushs != 0; k; k = doublePushs != 0 {
		from := bitScanForward(doublePushs)
		move := createMove(from, from-16, Pawn)
		move.enPassant = int(from) - 8
		ret = append(ret, move)
		doublePushs &= doublePushs - 1
	}

	return ret
}

func attacksToKing(b Board, c Color) uint64 {
	kingSquare := bitScanForward(b.getPieceSet(King, c))
	return b.getSquareAttackers(kingSquare) & b.getColorPieces(getOppositeColor(c))
}
