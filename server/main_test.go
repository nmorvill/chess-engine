package main

import (
	"fmt"
	chess "server/chess"
	"testing"
)

// Source of numbers : https://www.chessprogramming.org/Perft_Results

const StartingPos = "rnbqkbnr/1ppppppp/8/p7/Q1P5/8/PP1PPPPP/RNB1KBNR"

func BenchmarkStartingPosDepth1(b *testing.B) { testPos(1, b, 20, StartingPos) }
func BenchmarkStartingPosDepth2(b *testing.B) { testPos(2, b, 400, StartingPos) }
func BenchmarkStartingPosDepth3(b *testing.B) { testPos(3, b, 8902, StartingPos) }
func BenchmarkStartingPosDepth4(b *testing.B) { testPos(4, b, 197281, StartingPos) }
func BenchmarkStartingPosDepth5(b *testing.B) { testPos(5, b, 4865609, StartingPos) }

const SecondPos = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/1R2K2R"

func BenchmarkSecondPosDepth1(b *testing.B) { testPos(1, b, 48, SecondPos) }
func BenchmarkSecondPosDepth2(b *testing.B) { testPos(2, b, 2039, SecondPos) }
func BenchmarkSecondPosDepth3(b *testing.B) { testPos(3, b, 97862, SecondPos) }
func BenchmarkSecondPosDepth4(b *testing.B) { testPos(4, b, 4085603, SecondPos) }

func testPos(depth int, b *testing.B, expectedNumber int, fen string) {
	var board chess.Board
	board.SetBoardFromFEN(fen)

	chess.MagicInit()
	chess.InBetweenInit()

	got := getNumberOfCombinations(board, depth, chess.Black)

	if got != expectedNumber {
		b.Errorf("Error, got %d combinations instead of %d", got, expectedNumber)
	}
}

func getNumberOfCombinations(aBoard chess.Board, depth int, colorToMove chess.Color) int {

	if depth == 0 {
		return 1
	}

	moves := chess.MoveGenerator(aBoard, colorToMove)
	nbOfPos := 0

	nextColor := chess.Black

	if colorToMove == chess.Black {
		nextColor = chess.White
	}

	startingBoard := aBoard
	for _, move := range moves {
		newBoard := chess.MovePiece(startingBoard, move, colorToMove)
		newNbOfPos := getNumberOfCombinations(newBoard, depth-1, nextColor)
		nbOfPos += newNbOfPos
		if depth == 1 {
			fmt.Println(move, newNbOfPos)
		}
	}

	return nbOfPos
}
