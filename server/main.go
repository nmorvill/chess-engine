package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	chess "server/chess"
	"strconv"
	"time"
)

var board chess.Board

func main() {

	board.ResetBoard()

	fmt.Printf("Reading opening files \n")
	chess.LoadOpeningsFile()

	launchServer()
}

func launchServer() {
	http.HandleFunc("/", getRoot)
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

}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Init board")
	board.SetBoardFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")

	json.NewEncoder(w).Encode(board.Board)
}

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
