package chess

/*
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var openings []string

func LoadOpeningsFile() {

	ret := []string{}

	f, err := os.Open("./out.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	openings = ret
}

func parseMoveString(str string) Move {

	startSquare := int(str[0]) - int('a') + (int(str[1])-int('1'))*8
	endSquare := int(str[2]) - int('a') + (int(str[3])-int('1'))*8

	return Move{StartPos: startSquare, EndPos: endSquare}
}

func GetUCIFromMove(move Move) string {

	startSquare := string(rune(move.StartPos%8+'a')) + string(rune(int(move.StartPos/8)+'1'))
	endSquare := string(rune(move.EndPos%8+'a')) + string(rune(int(move.EndPos/8)+'1'))

	return startSquare + endSquare

}

func getOpeningMove(movesDone []Move) Move {

	stringToLookFor := ""

	for _, move := range movesDone {
		stringToLookFor += GetUCIFromMove(move) + " "
	}

	possibleOpenings := []Move{}
	for _, opening := range openings {
		if strings.HasPrefix(opening, stringToLookFor) {
			nextMove := opening[len(stringToLookFor):]
			possibleOpenings = append(possibleOpenings, parseMoveString(nextMove))
		}
	}

	if len(possibleOpenings) > 0 {
		return possibleOpenings[rand.Intn(len(possibleOpenings))]
	}

	return Move{}

}
*/
