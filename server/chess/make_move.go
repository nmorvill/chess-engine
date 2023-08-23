package chess

/*
https://www.chessprogramming.org/General_Setwise_Operations#UpdateByMove
*/
func MovePiece(b Board, m Move, c Color) Board {
	ret := b
	from := uint64(1 << m.GetFrom())
	to := uint64(1 << m.GetTo())
	fromTo := from ^ to
	ret.setPieceSet(m.piece, c, ret.getPieceSet(m.piece, c)^fromTo)
	ret.enPassant = m.enPassant

	isCapture := to&b.getColorPieces(getOppositeColor(c)) != 0
	if isCapture {
		b.piecesSquares[b.getPieceOfSquare(m.GetTo())] ^= to
		ret.occupiedSquares ^= from
		ret.emptySquares ^= from
	} else {
		ret.occupiedSquares ^= fromTo
		ret.emptySquares ^= fromTo
	}
	return ret
}
