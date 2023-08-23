package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	chess "server/chess"
)

var board chess.Board

func main() {

	fmt.Printf("Reading opening files \n")

	launchServer()
}

func launchServer() {
	http.HandleFunc("/", getRoot)
	chess.MagicInit()
	board.SetBoardFromFEN("rnbqkbnr/ppppppp1/7p/1B6/4P3/8/PPPP1PPP/RNBQK1NR")

	fmt.Println(chess.VisualizeBitBoard(board.GetPinnedPieces(chess.Black)))
	fmt.Println(chess.VisualizeBitBoard(board.GetPinnedPieces(chess.White)))

	for i := 0; i < 64; i++ {
	}

	/*
		chess.MagicInit()

		moves := chess.MoveGenerator(board, chess.Black)
		fmt.Println(len(moves))
	*/

	/*
			http.HandleFunc("/pieceSelected", getPieceSelected)
			http.HandleFunc("/move", getMove)
			http.HandleFunc("/computerMove", getComputerMove)

		fmt.Printf("Launching server \n")

		err := http.ListenAndServe(":44240", nil)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	*/

}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Init board")
	board.SetBoardFromFEN("rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R")

	json.NewEncoder(w).Encode(chess.GetArrayFromBitboard(board))
}

/*
func getMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	startSquare, _ := strconv.Atoi(r.URL.Query().Get("startSquare"))
	endSquare, _ := strconv.Atoi(r.URL.Query().Get("endSquare"))

	board = chess.MovePiece(board, chess.Move{StartPos: startSquare, EndPos: endSquare})

	json.NewEncoder(w).Encode(board.Board)

}

func getPieceSelected(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	square, _ := strconv.Atoi(r.URL.Query().Get("piece"))
	possibleMoves := chess.GetAllPossibleSquares(board, int(square))

	json.NewEncoder(w).Encode(possibleMoves)
}

func getComputerMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	start := time.Now()
	color := r.URL.Query().Get("color")

	if color == "black" {
		board = chess.GetComputerMove(board, chess.Black)
	} else {
		board = chess.GetComputerMove(board, chess.White)
	}

	fmt.Printf("Move calculated in %s \n", time.Since(start))

	json.NewEncoder(w).Encode(board.Board)
}
*/
