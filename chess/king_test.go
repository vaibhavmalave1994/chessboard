package chess



import(
	"testing"
)

func TestSetKing(t *testing.T){
	t.Run("King: empty piece", func(t *testing.T){
		result := SetKing("", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("King: set name, position is empty", func(t *testing.T){
		result := SetKing("king", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("King: set name and position", func(t *testing.T){
		result := SetKing("king", "H1")
		if result.Name == "" || result.Position == nil{
			t.Errorf("King: Expected type of piece got %v", result)
		}
	})
}

func TestPredictKingMoves(t *testing.T){
	t.Run("King: set position in the middle", func(t *testing.T){
		result := SetKing("king", "D5")
		expected := []string{"E5","C5","D4","D6","E6","E4","C4","C6"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("King: set position to bottom right corner", func(t *testing.T){
		result := SetKing("king", "A1")
		expected := []string{"B1","A2","B2"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("King: set position to bottom left corner", func(t *testing.T){
		result := SetKing("king", "A8")
		expected := []string{"B7","A7","B8"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("King: set position to top left corner", func(t *testing.T){
		result := SetKing("king", "H8")
		expected := []string{"G7","H7","G8"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("King: set position to top right corner", func(t *testing.T){
		result := SetKing("king", "H1")
		expected := []string{"G1","H2","G2"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})
	
}

func comparePostions(expected []string, positions []string) bool{
	var result bool 
	if len(expected) == 0 && len(positions)==0{
		return true
	}
	for _, value:= range expected{
		result = false
		for _,position:= range positions{
			if value== position{
				result = true
			}
		}
		if result== false{
			return result
		}
	}
	return result;
}
