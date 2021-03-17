package chess

import(
	"fmt"
)


type King struct{
	Name string
	Rules *Rule
	Position *Position
}

func SetKing(name string, position string) *King{
	if name == "" || position == ""{
		return nil
	}
	return &King{Name : name, Rules:getRuleForPice(name), Position: getPositionCoordinates(position)}
}

func(king *King)PredictMoves()[]string{
	rules := &Rule{Step:1, Directions:[]string{"f", "b", "l", "r", "fr", "fl", "bl", "br"}}
	var predictedMoved []string
	var predictedMovedPositions []Position
	for _,dir := range rules.Directions{
		fmt.Println("direction", dir)
		switch dir{
		case "f" : 
		   forwardMoves, err := getForwardMoves(rules.Step, king.Position)
		   if err == nil{
				for _, position:= range forwardMoves{
					predictedMovedPositions = append(predictedMovedPositions, position)
				}
		   }
		case "b" : 
		backwordMoves, err := getBackwordMoves(rules.Step, king.Position)
		if err == nil{
			for _, position:= range backwordMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "l" : 
		leftMoves, err := getLeftMoves(rules.Step, king.Position)
		if err == nil{
			for _, position:= range leftMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "r" : 
		rightMoves, err := getRightMoves(rules.Step, king.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fr" : 
		rightMoves, err := getForwardRightMoves(rules.Step, king.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fl" : 
		rightMoves, err := getForwardLeftMoves(rules.Step, king.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "br" : 
		fmt.Println("------------")
		rightMoves, err := getBackwardRightMoves(rules.Step, king.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "bl" : 
		rightMoves, err := getBackwardLeftMoves(rules.Step, king.Position)
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