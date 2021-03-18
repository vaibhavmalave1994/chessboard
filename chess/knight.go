package chess

type Knight struct{
	Name string
	Position *Position
}

func SetKnight(name string, position string) *Knight{
	if name == "" || position == ""{
		return nil
	}
	return &Knight{Name : name, Position: getPositionCoordinates(position)}
}

func(knight *Knight)PredictMoves()[]string{
	nextI := knight.Position.I
	nextJ := knight.Position.J
	var predictedMoved []string
	var predictedMovedPositions []Position
	//forward right
	if (nextJ+1>1 && nextJ+1<=8) && nextI+2<=8{
		predictedMovedPositions = append(predictedMovedPositions, Position{I:nextI+2,J:nextJ+1 })
	}
	//forwad left
	
	if (nextJ-1>=1 && nextJ-1<=8) && nextI+2<=8{
		predictedMovedPositions = append(predictedMovedPositions, Position{I:nextI+2,J:nextJ-1 })
	}
	//right forward
	if nextJ+2<=8 && nextI+1<=8{
		predictedMovedPositions = append(predictedMovedPositions, Position{I:nextI+1,J:nextJ+2 })
	}
	
	//right backward
	if nextJ+2<=8 && nextI-1>=1{
		predictedMovedPositions = append(predictedMovedPositions, Position{I:nextI-1,J:nextJ+2 })
	}

	//backward right
	if (nextJ+1>1 && nextJ+1<=8) && nextI-2>=1{
		predictedMovedPositions = append(predictedMovedPositions, Position{I:nextI-2,J:nextJ+1 })
	}
	//backward left
	if nextJ-1>=1 && nextI-2>=1{
		predictedMovedPositions = append(predictedMovedPositions, Position{I:nextI-2,J:nextJ-1 })
	}

	//left forward
	if nextJ-2>=1 && nextI+1<=8{
		predictedMovedPositions = append(predictedMovedPositions, Position{I:nextI+1,J:nextJ-2 })
	}
	
	//left backward
	if nextJ-2>=1 && nextI-1>=1{
		predictedMovedPositions = append(predictedMovedPositions, Position{I:nextI-1,J:nextJ-2 })
	}

	for _,position:= range predictedMovedPositions{
		predictedMoved = append(predictedMoved, convertPositionCoordinatesToLetters(&position))
	}

	return predictedMoved
}