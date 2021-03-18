package chess

import(
	"fmt"
)


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
	var predictedMoved []string
	var predictedMovedPositions []Position
	for _,dir := range rules.Directions{
		fmt.Println("direction", dir)
		switch dir{
		case "f" : 
		   forwardMoves, err := getForwardMoves(rules.Step, bishop.Position)
		   if err == nil{
				for _, position:= range forwardMoves{
					predictedMovedPositions = append(predictedMovedPositions, position)
				}
		   }
		case "b" : 
		backwordMoves, err := getBackwordMoves(rules.Step, bishop.Position)
		if err == nil{
			for _, position:= range backwordMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "l" : 
		leftMoves, err := getLeftMoves(rules.Step, bishop.Position)
		if err == nil{
			for _, position:= range leftMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "r" : 
		rightMoves, err := getRightMoves(rules.Step, bishop.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fr" : 
		rightMoves, err := getForwardRightMoves(rules.Step, bishop.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fl" : 
		rightMoves, err := getForwardLeftMoves(rules.Step, bishop.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "br" : 
		rightMoves, err := getBackwardRightMoves(rules.Step, bishop.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "bl" : 
		rightMoves, err := getBackwardLeftMoves(rules.Step, bishop.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		   
		}
	}
	fmt.Println("predictedMovedPositions", predictedMovedPositions)
	for _,position:= range predictedMovedPositions{
		predictedMoved = append(predictedMoved, convertPositionCoordinatesToLetters(&position))
	}
	return predictedMoved
}