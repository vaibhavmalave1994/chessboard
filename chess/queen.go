package chess

import(
	"fmt"
)


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
	var predictedMoved []string
	var predictedMovedPositions []Position
	for _,dir := range rules.Directions{
		fmt.Println("direction", dir)
		switch dir{
		case "f" : 
		   forwardMoves, err := getForwardMoves(rules.Step, queen.Position)
		   if err == nil{
				for _, position:= range forwardMoves{
					predictedMovedPositions = append(predictedMovedPositions, position)
				}
		   }
		case "b" : 
		backwordMoves, err := getBackwordMoves(rules.Step, queen.Position)
		if err == nil{
			for _, position:= range backwordMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "l" : 
		leftMoves, err := getLeftMoves(rules.Step, queen.Position)
		if err == nil{
			for _, position:= range leftMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "r" : 
		rightMoves, err := getRightMoves(rules.Step, queen.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fr" : 
		rightMoves, err := getForwardRightMoves(rules.Step, queen.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fl" : 
		rightMoves, err := getForwardLeftMoves(rules.Step, queen.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "br" : 
		fmt.Println("------------")
		rightMoves, err := getBackwardRightMoves(rules.Step, queen.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "bl" : 
		rightMoves, err := getBackwardLeftMoves(rules.Step, queen.Position)
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