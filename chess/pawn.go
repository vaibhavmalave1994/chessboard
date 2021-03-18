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
	var predictedMoved []string
	var predictedMovedPositions []Position
	for _,dir := range rules.Directions{
		switch dir{
		case "f" : 
		   forwardMoves, err := getForwardMoves(rules.Step, pawn.Position)
		   if err == nil{
				for _, position:= range forwardMoves{
					predictedMovedPositions = append(predictedMovedPositions, position)
				}
		   }
		case "b" : 
		backwordMoves, err := getBackwordMoves(rules.Step, pawn.Position)
		if err == nil{
			for _, position:= range backwordMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "l" : 
		leftMoves, err := getLeftMoves(rules.Step, pawn.Position)
		if err == nil{
			for _, position:= range leftMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "r" : 
		rightMoves, err := getRightMoves(rules.Step, pawn.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fr" : 
		rightMoves, err := getForwardRightMoves(rules.Step, pawn.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fl" : 
		rightMoves, err := getForwardLeftMoves(rules.Step, pawn.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "br" : 
		rightMoves, err := getBackwardRightMoves(rules.Step, pawn.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "bl" : 
		rightMoves, err := getBackwardLeftMoves(rules.Step, pawn.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		   
		}
	}
	
	for _,position:= range predictedMovedPositions{
		predictedMoved = append(predictedMoved, convertPositionCoordinatesToLetters(&position))
	}
	return predictedMoved
}