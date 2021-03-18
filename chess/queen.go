package chess


type Queen struct{
	Name string
	Position *Position
}

func SetQueen(name string, position string) *Queen{
	if name == "" || position == ""{
		return nil
	}
	return &Queen{Name : name, Position: getPositionCoordinates(position)}
}

func(queen *Queen)PredictMoves()[]string{
	rules := &Rule{Step:8, Directions:[]string{"f", "b", "l", "r", "fr", "fl", "bl", "br"}}
	return predictMovesUsingRules(rules, queen.Position)
}