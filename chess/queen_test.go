package chess



import(
	"testing"
)

func TestSetQueen(t *testing.T){
	t.Run("Queen: empty piece", func(t *testing.T){
		result := SetKing("", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Queen: set name, position is empty", func(t *testing.T){
		result := SetKing("queen", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Queen: set name and position", func(t *testing.T){
		result := SetKing("queen", "H1")
		if result.Name == "" || result.Position == nil{
			t.Errorf("King: Expected type of piece got %v", result)
		}
	})
}

func TestPredictQueenMoves(t *testing.T){
	t.Run("Queen: set position in the middle", func(t *testing.T){
		result := SetQueen("queen", "D5")
		expected := []string{"E5","F5","G5","H5","D4","D3","D2","D1","C5", "B5", "A5", "D6", "D7", "D8", "E6", "F7", "G8", "C4", "B3", "A2", "C6", "B7", "A8", "E4", "F3", "G2", "H1"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	// t.Run("Queen: set position to bottom right corner", func(t *testing.T){
	// 	result := SetKing("queen", "A1")
	// 	expected := []string{"B1","A2","B2"}
	// 	positions := result.PredictMoves()
	// 	if comparePostions(expected, positions) == false{
	// 		t.Errorf("expected %v got %v", expected, positions)
	// 	}
		
	// })

	// t.Run("Queen: set position to bottom left corner", func(t *testing.T){
	// 	result := SetKing("queen", "A8")
	// 	expected := []string{"B7","A7","B8"}
	// 	positions := result.PredictMoves()
	// 	if comparePostions(expected, positions) == false{
	// 		t.Errorf("expected %v got %v", expected, positions)
	// 	}
		
	// })

	// t.Run("Queen: set position to top left corner", func(t *testing.T){
	// 	result := SetKing("queen", "H8")
	// 	expected := []string{"G7","H7","G8"}
	// 	positions := result.PredictMoves()
	// 	if comparePostions(expected, positions) == false{
	// 		t.Errorf("expected %v got %v", expected, positions)
	// 	}
		
	// })

	// t.Run("Queen: set position to top right corner", func(t *testing.T){
	// 	result := SetKing("queen", "H1")
	// 	expected := []string{"G1","H2","G2"}
	// 	positions := result.PredictMoves()
	// 	if comparePostions(expected, positions) == false{
	// 		t.Errorf("expected %v got %v", expected, positions)
	// 	}
		
	// })
	
}

