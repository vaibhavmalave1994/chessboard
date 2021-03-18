package chess

type Bishop struct{
	Name string
	Position *Position
}

func SetBishop(name string, position string) *Bishop{
	if name == "" || position == ""{
		return nil
	}
	return &Bishop{Name : name, Position: getPositionCoordinates(position)}
}

func(bishop *Bishop)PredictMoves()[]string{
	rules := &Rule{Step:8, Directions:[]string{ "fr", "fl", "bl", "br"}}
	return predictMovesUsingRules(rules, bishop.Position)
}