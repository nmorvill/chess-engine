package chess

import "math"

type Material struct {
	piece Piece
	pos   int
	color Color
}

type Move struct {
	StartPos  int
	EndPos    int
	enPassant int
}

type Direction int

const (
	Up        Direction = 8
	Down      Direction = -8
	Left      Direction = -1
	Right     Direction = 1
	UpRight   Direction = 9
	UpLeft    Direction = 7
	DownLeft  Direction = -9
	DownRight Direction = -7
)

var directions = []Direction{Up, Down, Left, Right, UpRight, UpLeft, DownLeft, DownRight}

func getAllCaptureMoves(board Board, color Color) []Move {
	var ret []Move
	for _, i := range getAllMaterial(board.Board, color) {
		for _, m := range getLegalMovesOfPiece(board, i) {
			if m.EndPos != 0 {
				if color == White {
					newBoard := MovePiece(board, m)
					if !isChecking(newBoard.Board, getAttackingSquaresOfColor(newBoard, Black)) {
						ret = append(ret, m)
					}
				} else if color == Black {
					newBoard := MovePiece(board, m)
					if !isChecking(newBoard.Board, getAttackingSquaresOfColor(newBoard, White)) {
						ret = append(ret, m)
					}
				} else {
					ret = append(ret, m)
				}
			}
		}
	}

	return ret
}

func GetAllPossibleSquares(board Board, square int) []int {
	var ret []int
	for _, i := range getLegalMovesOfPiece(board, getMaterialFromSquare(board.Board, square)) {
		ret = append(ret, i.EndPos)
	}
	return ret
}

func MovePiece(board Board, move Move) Board {

	board.pawnPushedLastTurn = -1

	material := getMaterialFromSquare(board.Board, move.StartPos)

	//Capture en passant
	if material.piece == Pawn && math.Abs(float64(move.EndPos)-float64(move.StartPos)) != 8 && board.Board[move.EndPos] == 0 {
		if material.color == White {
			board.Board[move.EndPos-8] = 0
		} else {
			board.Board[move.EndPos+8] = 0
		}
	}

	board.Board[move.EndPos] = board.Board[move.StartPos]
	board.Board[move.StartPos] = 0

	//Roques
	if material.piece == Rook || material.piece == King {
		castling(&board, move.StartPos, move.EndPos, material)
	}

	//Promotion d'un pion
	if material.piece == Pawn {
		promotePawn(&board.Board, move.StartPos, move.EndPos, material)
		enPassant(&board, move)
	}

	board.squaresUnderWhiteAttack = getAttackingSquaresOfColor(board, White)
	board.squaresUnderBlackAttack = getAttackingSquaresOfColor(board, Black)

	board.isWhiteChecked = isChecking(board.Board, board.squaresUnderBlackAttack)
	board.isBlackChecked = isChecking(board.Board, board.squaresUnderWhiteAttack)

	board.movesDone = append(board.movesDone, move)

	return board
}

func castling(board *Board, startSquare int, endSquare int, material Material) {
	if material.piece == Rook {
		if material.color == White {
			if startSquare == 0 {
				board.canWhiteQueenCastle = false
			} else if startSquare == 7 {
				board.canWhiteKingCastle = false
			}
		} else {
			if startSquare == 56 {
				board.canBlackQueenCastle = false
			} else if startSquare == 63 {
				board.canBlackKingCastle = false
			}
		}
	}
	if material.piece == King {
		if material.color == White {
			board.canWhiteKingCastle, board.canWhiteQueenCastle = false, false
		} else {
			board.canBlackKingCastle, board.canBlackQueenCastle = false, false
		}

		//On est en train de roquer donc faut tej la tour
		if math.Abs(float64(endSquare)-float64(startSquare)) == 2 {
			if endSquare < startSquare { // Grand roque
				board.Board[startSquare-1] = board.Board[startSquare-4]
				board.Board[startSquare-4] = 0
			} else { //Petit roque
				board.Board[startSquare+1] = board.Board[startSquare+3]
				board.Board[startSquare+3] = 0
			}
		}
	}
}

func promotePawn(board *[64]int, startSquare int, endSquare int, material Material) {
	if material.color == White && endSquare > 55 {
		board[endSquare] = int(White) | int(Queen)
	} else if material.color == White && endSquare < 8 {
		board[endSquare] = int(Black) | int(Queen)
	}
}

func enPassant(board *Board, move Move) {
	if math.Abs(float64(move.EndPos)-float64(move.StartPos)) == 16 {
		board.pawnPushedLastTurn = move.EndPos
	}
}

func GetAllLegalMoves(board Board, color Color) []Move {
	var ret []Move
	for _, i := range getAllMaterial(board.Board, color) {
		for _, m := range getLegalMovesOfPiece(board, i) {
			if color == White {
				newBoard := MovePiece(board, m)
				if !isChecking(newBoard.Board, getAttackingSquaresOfColor(newBoard, Black)) {
					ret = append(ret, m)
				}
			} else if color == Black {
				newBoard := MovePiece(board, m)
				if !isChecking(newBoard.Board, getAttackingSquaresOfColor(newBoard, White)) {
					ret = append(ret, m)
				}
			} else {
				ret = append(ret, m)
			}
		}
	}

	return ret
}

func getDistanceToEdge(pos int, direction Direction) int {
	switch direction {
	case Up:
		return int((63 - pos) / 8)
	case Down:
		return int(pos / 8)
	case Left:
		return pos % 8
	case Right:
		return 7 - (pos % 8)
	case UpRight:
		return Min(7-(pos%8), int((63-pos)/8))
	case UpLeft:
		return Min(pos%8, int((63-pos)/8))
	case DownRight:
		return Min(7-(pos%8), int(pos/8))
	case DownLeft:
		return Min(pos%8, int(pos/8))
	default:
		return 0
	}
}

func getMaterialFromSquare(board [64]int, square int) Material {
	return Material{Piece(board[square] & 7), square, Color(board[square] & 24)}
}

func getAllMaterial(board [64]int, color Color) []Material {
	var ret []Material
	for k, i := range board {
		if i&24 == int(color) {
			ret = append(ret, Material{Piece(i & 7), k, color})
		}
	}
	return ret
}

func getLegalMovesOfPiece(board Board, material Material) []Move {
	var ret []Move

	if material.piece == Knight {
		return getLegalMovesOfKnight(board.Board, material)
	}

	startIndex, endIndex, maxDistance := 0, 8, 8

	if material.piece == King {
		maxDistance = 1
	} else if material.piece == Rook {
		endIndex = 4
	} else if material.piece == Pawn {
		if (material.color) == White {
			endIndex = 1

			//Capture
			if getDistanceToEdge(material.pos, UpRight) > 0 && (board.Board[material.pos+9]&24 == int(Black) || board.pawnPushedLastTurn == material.pos+1) {
				move := Move{material.pos, material.pos + 9, 0}
				if board.pawnPushedLastTurn == material.pos+1 {
					move.enPassant = material.pos + 1
				}
				ret = append(ret, move)
			}
			if getDistanceToEdge(material.pos, UpLeft) > 0 && (board.Board[material.pos+7]&24 == int(Black) || board.pawnPushedLastTurn == material.pos-1) {
				move := Move{material.pos, material.pos + 7, 0}
				if board.pawnPushedLastTurn == material.pos-1 {
					move.enPassant = material.pos - 1
				}
				ret = append(ret, move)
			}
		} else {
			startIndex, endIndex = 1, 2

			if getDistanceToEdge(material.pos, DownRight) > 0 && (board.Board[material.pos-7]&24 == int(White) || board.pawnPushedLastTurn == material.pos+1) {
				move := Move{material.pos, material.pos - 7, 0}
				if board.pawnPushedLastTurn == material.pos+1 {
					move.enPassant = material.pos + 1
				}
				ret = append(ret, move)
			}
			if getDistanceToEdge(material.pos, DownLeft) > 0 && (board.Board[material.pos-9]&24 == int(White) || board.pawnPushedLastTurn == material.pos-1) {
				move := Move{material.pos, material.pos - 9, 0}
				if board.pawnPushedLastTurn == material.pos-1 {
					move.enPassant = material.pos - 1
				}
				ret = append(ret, move)
			}
		}
		if (material.pos < 16 && material.color == White) || (material.pos > 47 && material.color == Black) {
			maxDistance = 2
		} else {
			maxDistance = 1
		}

	} else if material.piece == Bishop {
		startIndex = 4
	}

	for i := startIndex; i < endIndex; i++ {
		for j := 1; j <= Min(getDistanceToEdge(material.pos, directions[i]), maxDistance); j++ {
			endSquare := material.pos + int(directions[i])*j
			if board.Board[endSquare] == 0 {
				//Si la case est vide alors on ajoute et on continue
				ret = append(ret, Move{StartPos: material.pos, EndPos: endSquare})
			} else if board.Board[endSquare]&24 != int(material.color) {
				//Si la case contient une pièce ennemie alors on la prend (sauf si on est un pion) et on ne peut pas aller plus loin
				if material.piece != Pawn {
					ret = append(ret, Move{StartPos: material.pos, EndPos: endSquare})
				}
				break
			} else {
				//Si la case contient une pièce alliée alors on ne peut pas la prendre ni aller plus loin
				break
			}
		}
	}

	if material.piece == King {
		isKingSideEmpty, isQueenSideEmpty := true, true
		for i := 1; i < 4; i++ {
			if i < 3 && material.pos+i < 64 && board.Board[material.pos+i] != 0 {
				isKingSideEmpty = false
			}
			if material.pos-i >= 0 && board.Board[material.pos-i] != 0 {
				isQueenSideEmpty = false
			}
		}

		for i := 0; i < 2; i++ {
			if material.pos+i < 64 {
				if material.color == White {
					if isSquareUnderAttack(board.Board, board.squaresUnderBlackAttack, material.pos+i) {
						isKingSideEmpty = false
					}
				} else {
					if isSquareUnderAttack(board.Board, board.squaresUnderWhiteAttack, material.pos+i) {
						isKingSideEmpty = false
					}
				}
			}
			if material.pos-i >= 0 {
				if material.color == White {
					if isSquareUnderAttack(board.Board, board.squaresUnderBlackAttack, material.pos-i) {
						isQueenSideEmpty = false
					}
				} else {
					if isSquareUnderAttack(board.Board, board.squaresUnderWhiteAttack, material.pos-i) {
						isQueenSideEmpty = false
					}
				}
			}
		}

		if material.color == White {
			if board.canWhiteKingCastle && isKingSideEmpty {
				ret = append(ret, Move{StartPos: material.pos, EndPos: material.pos + 2})
			}
			if board.canWhiteQueenCastle && isQueenSideEmpty {
				ret = append(ret, Move{StartPos: material.pos, EndPos: material.pos - 2})
			}
		} else {
			if board.canBlackKingCastle && isKingSideEmpty {
				ret = append(ret, Move{StartPos: material.pos, EndPos: material.pos + 2})
			}
			if board.canBlackQueenCastle && isQueenSideEmpty {
				ret = append(ret, Move{StartPos: material.pos, EndPos: material.pos - 2})
			}
		}
	}

	return ret
}

func getLegalMovesOfKnight(board [64]int, material Material) []Move {
	var ret []Move
	var endSquares []int

	leftMax := getDistanceToEdge(material.pos, Left)
	upMax := getDistanceToEdge(material.pos, Up)
	rightMax := getDistanceToEdge(material.pos, Right)
	downMax := getDistanceToEdge(material.pos, Down)

	if leftMax >= 2 && upMax >= 1 {
		endSquares = append(endSquares, material.pos+6)
	}
	if leftMax >= 1 && upMax >= 2 {
		endSquares = append(endSquares, material.pos+15)
	}
	if rightMax >= 1 && upMax >= 2 {
		endSquares = append(endSquares, material.pos+17)
	}
	if rightMax >= 2 && upMax >= 1 {
		endSquares = append(endSquares, material.pos+10)
	}
	if rightMax >= 2 && downMax >= 1 {
		endSquares = append(endSquares, material.pos-6)
	}
	if rightMax >= 1 && downMax >= 2 {
		endSquares = append(endSquares, material.pos-15)
	}
	if leftMax >= 1 && downMax >= 2 {
		endSquares = append(endSquares, material.pos-17)
	}
	if leftMax >= 2 && downMax >= 1 {
		endSquares = append(endSquares, material.pos-10)
	}

	for _, i := range endSquares {
		if board[i] == 0 || board[i]&24 != int(material.color) {
			ret = append(ret, Move{StartPos: material.pos, EndPos: i})
		}
	}

	return ret
}

func getAttackingSquaresOfColor(board Board, color Color) []int {
	var ret []int
	for _, m := range getAllMaterial(board.Board, color) {
		ret = append(ret, getAttackingSquaresOfPiece(board, m)...)
	}

	return ret
}

func getAttackingSquaresOfPiece(board Board, material Material) []int {
	var ret []int
	for _, move := range getLegalMovesOfPiece(board, material) {
		ret = append(ret, move.EndPos)
	}

	return ret
}

func isChecking(board [64]int, squaresUnderAttack []int) bool {
	for _, i := range squaresUnderAttack {
		if board[i]&7 == int(King) {
			return true
		}
	}
	return false
}

func isSquareUnderAttack(board [64]int, squaresUnderAttack []int, square int) bool {
	for _, i := range squaresUnderAttack {
		if i == square {
			return true
		}
	}
	return false
}
