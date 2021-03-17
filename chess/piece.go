package chess

import(
	
)

// type Piece struct{
// 	Name string 
// 	Rules *Rule
// 	Position *Position
// }

type Piece interface{

}

type Position struct{
	I int 
	J int
}

type Rule struct{
	Step int
	Directions []string
}

func SetPiece(name string, position string)Piece{
	if name == "" || position == ""{
		return nil
	}
	
	switch name{
	case "king":
		return &King{Name : name, Rules:getRuleForPice(name), Position: getPositionCoordinates(position)}
	}
	
	return nil
	
}




