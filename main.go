package main

import(
	"fmt"
	"chessboard/chess"
)

func main(){
	knight := chess.SetPiece("knight", "H8")
	fmt.Println("predicted position", knight.PredictMoves())
}