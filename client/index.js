const SQUARE_SIZE = 64
const BOARD_SIZE = SQUARE_SIZE * 8

const SERVER_URL = "http://localhost:44240"

const pieces = {
    1: "King",
    2: "Queen",
    3: "Rook",
    4: "Bishop",
    5: "Knight",
    6: "Pawn"
}

const colors = {
    8 : "White",
    16 : "Black"
}

let board = Array(64).fill(0)

let selection = {
    isSelected : false,
    selectedSquare : 0,
    possibleSquares : []
}

let isUserWhite = true
let isWhiteToMove = true


drawBoard()
fetchBoardState().then(() => {
    drawPieces()
})


async function aiVsAi() {
    while(true) {
        getMoveFromComputer("white")
        await new Promise(resolve => setTimeout(resolve, 3000))
        drawPieces()
        getMoveFromComputer("black")
        await new Promise(resolve => setTimeout(resolve, 3000))
        drawPieces()
    }
}

function drawBoard() {
    for(let i=0;i<8;i++) {
        for(let j=0; j<8;j++) {
            let div = document.createElement("div")
            div.classList.add((i + j) % 2 == 0 ? "white" : "black")
            div.style.top = SQUARE_SIZE * j + "px"
            div.style.left = SQUARE_SIZE * i + "px"
            div.style.width = SQUARE_SIZE + "px"
            div.style.height = SQUARE_SIZE + "px"
            div.id = 8 * (7 - j) + i
            document.getElementById("main").appendChild(div)
            

            div.addEventListener("click", (e) => clickHandler(e))
        }
    }
}

function drawPieces() {
    for(let i=0;i<64;i++) {
        let div = document.getElementById(i)
        div.innerHTML = ""
        if(board[i] != 0) {
            let img = document.createElement("img")
            img.src = "./img/" + (board[i] >> 3 == 1 ? "w" : "b") + (pieces[board[i] & 7]) + ".svg"
            div.appendChild(img)
        }
    }
}

async function fetchBoardState() {
    board = await fetch(SERVER_URL).then(r => r.json())
}

async function fetchWhiteMove() {
    await fetch(SERVER_URL + "/whiteMove")
}

function clickHandler(e) {
    if(isUserWhite == isWhiteToMove) {
        var id = parseInt(e.target.parentNode.id == "main" ? e.target.id : e.target.parentNode.id)
        if(selection.possibleSquares == null || (board[id] != 0 && !selection.possibleSquares.includes(id))) {
            fetch(SERVER_URL + "/pieceSelected?piece=" + id)
                .then(r => r.json())
                .then(r => {
                    selection = {
                        isSelected: true,
                        selectedSquare: id,
                        possibleSquares: r
                    }
                    colorSquares()
                    if(r == []) {
                        alert("MATE")
                    }
                })
        } else if(selection.possibleSquares != null && selection.possibleSquares.includes(id)) {
            fetch(SERVER_URL + "/move?startSquare=" + selection.selectedSquare + "&endSquare=" + id)
                .then(r => r.json())
                .then(r => {
                    board = r
                    drawPieces() 
                    selection = {
                        isSelected : false,
                        selectedSquare : 0,
                        possibleSquares : []
                    }
                    colorSquares()
                    isWhiteToMove = !isWhiteToMove
                })
                .then(new Promise(resolve => setTimeout(resolve, 500)))
                .then(r => getMoveFromComputer("black"))
        }
    }
}

function getMoveFromComputer(color){
    fetch(SERVER_URL + "/computerMove?color=" + color)
        .then(r => r.json())
        .then(r => {
            board = r
            drawPieces()
            isWhiteToMove = !isWhiteToMove
        })
}

function colorSquares() {
    for(const square of document.getElementById("main").children) {
        square.classList.remove("possible")
        square.classList.remove("selected")
    }
        
    if(selection.isSelected) {
        document.getElementById(selection.selectedSquare).classList.add("selected")

        if(selection.possibleSquares != null) {
            for(const square of selection.possibleSquares) {
                document.getElementById(square).classList.add("possible")
            }
        }
    }
}