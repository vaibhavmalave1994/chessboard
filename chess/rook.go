package chess

import(
	"fmt"
)


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
	var predictedMoved []string
	var predictedMovedPositions []Position
	for _,dir := range rules.Directions{
		fmt.Println("direction", dir)
		switch dir{
		case "f" : 
		   forwardMoves, err := getForwardMoves(rules.Step, rook.Position)
		   if err == nil{
				for _, position:= range forwardMoves{
					predictedMovedPositions = append(predictedMovedPositions, position)
				}
		   }
		case "b" : 
		backwordMoves, err := getBackwordMoves(rules.Step, rook.Position)
		if err == nil{
			for _, position:= range backwordMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "l" : 
		leftMoves, err := getLeftMoves(rules.Step, rook.Position)
		if err == nil{
			for _, position:= range leftMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "r" : 
		rightMoves, err := getRightMoves(rules.Step, rook.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fr" : 
		rightMoves, err := getForwardRightMoves(rules.Step, rook.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "fl" : 
		rightMoves, err := getForwardLeftMoves(rules.Step, rook.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "br" : 
		fmt.Println("------------")
		rightMoves, err := getBackwardRightMoves(rules.Step, rook.Position)
		if err == nil{
			for _, position:= range rightMoves{
				predictedMovedPositions = append(predictedMovedPositions, position)
			}
		}
		case "bl" : 
		rightMoves, err := getBackwardLeftMoves(rules.Step, rook.Position)
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