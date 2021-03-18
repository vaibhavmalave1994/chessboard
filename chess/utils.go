package chess

import(
	"strconv"
	"fmt"
	"errors"
)


func getXCoordinateByLetter(x string) int  {

    xCoordinateMap := map[string]int{"A":1,"B":2,"C":3, "D":4, "E": 5, "F":6, "G": 7, "H":8}
	if value, ok:=xCoordinateMap[x]; ok{
		return value
	}
	return 0
}

func getPositionCoordinates(position string) *Position{
	xLetter := position[0]
	yLetter := position[1]
	
	x:=getXCoordinateByLetter(string(xLetter))
	y, _ := strconv.Atoi(string(yLetter))
	return &Position{x,y}
}



func getForwardMoves(steps int, position *Position)([]Position, error){
	var result []Position
	if position.I == 8{
		return result, errors.New("NO_FURTHER_MOVES") 
	}
	nextI := position.I
	nextJ := position.J
	for steps>0{
		if nextI+1<=8{
			nextI = nextI+1
			result = append(result, Position{I:nextI, J:nextJ})
			fmt.Println("Position", result)
		}else{
			break
		}
		steps--
	}

	return result, nil
}

func getBackwordMoves(steps int, position *Position)([]Position, error){
	var result []Position
	if position.I == 1{
		return result, errors.New("NO_FURTHER_MOVES") 
	}
	nextI := position.I
	nextJ := position.J
	for steps>0{
		if nextI-1>=1{
			nextI = nextI-1
			result = append(result, Position{I:nextI, J:nextJ})
		}else{
			break
		}
		steps--
	}
	return result, nil
}

func getLeftMoves(steps int, position *Position)([]Position, error){
	var result []Position
	fmt.Println("into get left moves",position.J)
	if position.J == 1{
		return result, errors.New("NO_FURTHER_MOVES") 
	}
	nextI := position.I
	nextJ := position.J
	for steps>0{
		if nextJ-1>=1{
			nextJ = nextJ-1
			result = append(result, Position{I:nextI, J:nextJ})
		} 
		steps--
	}
	return result, nil
}

func getRightMoves(steps int, position *Position)([]Position, error){
	var result []Position
	fmt.Println("into get right moves",position.J)
	if position.J == 8{
		return result, errors.New("NO_FURTHER_MOVES") 
	}
	nextI := position.I
	nextJ := position.J
	for steps>0{
		if nextJ+1<=8{
			nextJ = nextJ+1
			result = append(result, Position{I:nextI, J:nextJ})
		}else{
			break
		}
		steps--
	}
	return result, nil
}

func getForwardRightMoves(steps int, position *Position)([]Position, error){
	var result []Position
	fmt.Println("into get forward right moves",position.I,position.J)
	if position.I == 8 || position.J == 8{
		return result, errors.New("NO_FURTHER_MOVES") 
	}
	nextI := position.I
	nextJ := position.J
	for steps>0{
		if nextI+1<=8 && nextJ+1<=8{
			nextI = nextI +1
			nextJ = nextJ+1
			result = append(result, Position{I:nextI, J:nextJ})
		}else{
			break
		} 
		steps--
	}
	return result, nil
}

func getForwardLeftMoves(steps int, position *Position)([]Position, error){
	var result []Position
	fmt.Println("into get forward left moves",position.I,position.J)
	if position.I == 8 || position.J == 1{
		return result, errors.New("NO_FURTHER_MOVES") 
	}
	nextI := position.I
	nextJ := position.J
	for steps>0{
		if nextI+1<=8 && nextJ-1>=1{
			nextI = nextI +1
			nextJ = nextJ-1
			result = append(result, Position{I:nextI, J:nextJ})
		}else{
			break
		} 
		steps--
	}
	return result, nil
}

func getBackwardRightMoves(steps int, position *Position)([]Position, error){
	var result []Position
	fmt.Println("into get backward right moves",position.I,position.J)
	if position.I == 1 || position.J == 8{
		return result, errors.New("NO_FURTHER_MOVES") 
	}
	nextI := position.I
	nextJ := position.J
	for steps>0{
		if nextI-1>=1 && nextJ+1<=8{
			nextI = nextI -1
			nextJ = nextJ+1
			result = append(result, Position{I:nextI, J:nextJ})
		}else{
			break
		} 
		steps--
	}
	return result, nil
}

func getBackwardLeftMoves(steps int, position *Position)([]Position, error){
	var result []Position
	fmt.Println("into get backward left moves",position.I,position.J)
	if position.I == 1 || position.J == 1{
		return result, errors.New("NO_FURTHER_MOVES") 
	}
	nextI := position.I
	nextJ := position.J
	for steps>0{
		if nextI-1>=1 && nextJ-1>=1{
			nextI = nextI-1
			nextJ = nextJ-1
			result = append(result, Position{I:nextI, J:nextJ})
		}else{
			break
		} 
		steps--
	}
	return result, nil
}

func convertPositionCoordinatesToLetters(position *Position) string{
	xCoordinateMap := map[int]string{1:"A",2:"B",3:"C", 4:"D", 5:"E", 6:"F", 7:"G", 8:"H"}
	if value, ok:=xCoordinateMap[position.I]; ok{
		return fmt.Sprintf("%s%d",value,position.J)
	}
	return ""
}
