package chess

import(
	// "fmt"
)

// type Piece struct{
// 	Name string 
// 	Rules *Rule
// 	Position *Position
// }

type Piece interface{
	PredictMoves() []string
}

type Position struct{
	I int 
	J int
}

type Rule struct{
	Step int
	Directions []string
}

func SetPiece(name string, position string) Piece{
	if name == "" || position == ""{
		return nil
	}

	switch name{
	case "king":
		return SetKing("king", position)
	case "pawn":
		return SetPawn("pawn", position)
	case "queen":
		return SetQueen("queen", position)
	case "rook":
		return SetRook("rook", position)
	case "bishop":
		return SetBishop("bishop", position)
	case "knight":
		return SetKnight("knight", position)
	}
	return nil
	
}




