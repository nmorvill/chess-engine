package chess

/*
import (
	"fmt"
	"sort"
)

var valueArr = []int{0, 0, 900, 500, 300, 300, 100}

var bestMoveThisIteration Move
var bestEval int
var nodesVisited int

const MATE_SCORE = 100000
const POSITIVE_INF = 999999
const NEGATIVE_INF = -POSITIVE_INF

func GetComputerMove(board Board, color Color) Board {

	var bestMove Move

	if len(board.movesDone) < 10 {
		fmt.Println("Getting move from book...")
		bestMove = getOpeningMove(board.movesDone)
	}
	if (len(board.movesDone) >= 10 || bestMove == Move{}) {
		fmt.Println("Calculating move...")
		bestMove = searchBestMove(board, color, 4)
	}

	return MovePiece(board, bestMove)
}

func evaluateBoard(board Board, color Color) int {
	ret := 0

	//Simple comptage des pièces
	ret += countPieces(board, color)

	return ret
}

func countPieces(board Board, color Color) int {
	ret := 0

	for _, k := range board.Board {
		if k&24 == int(color) {
			ret += valueArr[k&7]
		} else {
			ret -= valueArr[k&7]
		}
	}

	return ret
}

func searchBestMove(board Board, color Color, depth int) Move {
	nodesVisited = 0
	bestMoveThisIteration = Move{}

	bestEval = minMaxSearch(board, color, depth, 0, NEGATIVE_INF, POSITIVE_INF)

	fmt.Println(nodesVisited)

	return bestMoveThisIteration
}

func minMaxSearch(board Board, color Color, depth int, distFromRoot int, alpha int, beta int) int {

	if depth == 0 {
		return evaluateBoard(board, color)
		//return searchAllCaptures(board, color, alpha, beta)
	}

	legalMoves := getOrderedMoves(board, GetAllLegalMoves(board, color))
	//legalMoves := GetAllLegalMoves(board, color)

	if len(legalMoves) == 0 {
		if board.isInCheck(color) {
			return -MATE_SCORE
		}
		return 0
	}

	for _, m := range legalMoves {
		newBoard := MovePiece(board, m)
		eval := -minMaxSearch(newBoard, getOppositeColor(color), depth-1, distFromRoot+1, -beta, -alpha)
		nodesVisited += 1

		if eval >= beta {
			return beta
		}

		if eval > alpha {
			alpha = eval
			if distFromRoot == 0 {
				bestMoveThisIteration = m
			}
		}
	}

	return alpha

}

func searchAllCaptures(board Board, color Color, alpha int, beta int) int {

	evaluation := evaluateBoard(board, color)
	if evaluation >= beta {
		return beta
	}

	if evaluation > alpha {
		alpha = evaluation
	}

	moves := getOrderedMoves(board, GetAllLegalMoves(board, color))
	for _, m := range moves {
		newBoard := MovePiece(board, m)
		evaluation = -searchAllCaptures(newBoard, getOppositeColor(color), -beta, -alpha)

		if evaluation >= beta {
			return beta
		}

		if evaluation > alpha {
			alpha = evaluation
		}
	}

	return alpha

}

func getMovePredictedValue(board Board, move Move) int {
	predictedValue := 0
	startSquarePiece := board.Board[move.StartPos]
	endSquarePiece := board.Board[move.EndPos]
	if endSquarePiece != 0 {
		predictedValue += (valueArr[endSquarePiece&7] - valueArr[startSquarePiece&7]) * 10
	}
	return predictedValue
}

func getOrderedMoves(board Board, moves []Move) []Move {

	sort.SliceStable(moves, func(i, j int) bool {
		return getMovePredictedValue(board, moves[i]) < getMovePredictedValue(board, moves[j])
	})

	return moves

}
*/
