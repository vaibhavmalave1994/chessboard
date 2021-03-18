package chess

type Rook struct{
	Name string
	Position *Position
}

func SetRook(name string, position string) *Rook{
	if name == "" || position == ""{
		return nil
	}
	return &Rook{Name : name, Position: getPositionCoordinates(position)}
}

func(rook *Rook)PredictMoves()[]string{
	rules := &Rule{Step:8, Directions:[]string{"f", "b", "l", "r"}}
	return predictMovesUsingRules(rules, rook.Position)
}