package chess


type Pawn struct{
	Name string
	Position *Position
}

func SetPawn(name string, position string) *Pawn{
	if name == "" || position == ""{
		return nil
	}
	return &Pawn{Name : name, Position: getPositionCoordinates(position)}
}

func(pawn *Pawn)PredictMoves()[]string{
	rules := &Rule{Step:1, Directions:[]string{"f"}}
	return predictMovesUsingRules(rules, pawn.Position)
}