package main

import(
	"fmt"
	"chessboard/chess"
)

func main(){
	
	var choice int
	var pieceName string
	var position string
cont:
	fmt.Println("1. Next Piece, 2. Bye")
	fmt.Scanf("%d\n", &choice)
	

	switch choice {
	case 1:
		fmt.Println("enter piece name ")
		fmt.Scanf("%s\n", &pieceName)
		fmt.Println("enter piece position ")
		fmt.Scanf("%s\n", &position)
		knight := chess.SetPiece(pieceName,position)
		fmt.Println("predicted position", knight.PredictMoves())
		goto cont
	default:
		fmt.Println("BYE")
	}
}