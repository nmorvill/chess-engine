package chess

import (
	"fmt"
	"strings"
	"unicode"
)

type Board struct {
	Board                   [64]int
	canWhiteKingCastle      bool
	canWhiteQueenCastle     bool
	canBlackKingCastle      bool
	canBlackQueenCastle     bool
	squaresUnderWhiteAttack []int
	squaresUnderBlackAttack []int
	isWhiteChecked          bool
	isBlackChecked          bool
	pawnPushedLastTurn      int
	movesDone               []Move
}

func (board *Board) SetBoardFromFEN(fen string) {

	board.ResetBoard()
	rows := strings.Split(fen, "/")

	for i := 0; i < 8; i++ {
		j := 0
		for _, c := range rows[i] {
			if unicode.IsDigit(c) {
				j += int(c) - '0'
			} else {
				board.Board[8*(7-i)+j] = GetPieceCodeFromFEN(c)
				j += 1
			}
		}
	}

	board.squaresUnderBlackAttack = getAttackingSquaresOfColor(*board, Black)
	board.squaresUnderWhiteAttack = getAttackingSquaresOfColor(*board, White)

	fmt.Println(board.squaresUnderBlackAttack)
}

func (board *Board) ResetBoard() {
	board.Board = [64]int{}
	board.canWhiteKingCastle, board.canWhiteQueenCastle, board.canBlackKingCastle, board.canBlackQueenCastle = true, true, true, true
	board.isWhiteChecked, board.isBlackChecked = false, false
	board.movesDone = []Move{}
}

func (board *Board) isInCheck(color Color) bool {
	if color == Black {
		return board.isBlackChecked
	} else {
		return board.isWhiteChecked
	}
}
