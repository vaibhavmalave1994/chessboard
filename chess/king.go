package chess

type King struct{
	Name string
	Position *Position
}

func SetKing(name string, position string) *King{
	if name == "" || position == ""{
		return nil
	}
	return &King{Name : name, Position: getPositionCoordinates(position)}
}

func(king *King)PredictMoves()[]string{
	rules := &Rule{Step:1, Directions:[]string{"f", "b", "l", "r", "fr", "fl", "bl", "br"}}
	return predictMovesUsingRules(rules, king.Position)
}