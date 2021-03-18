package chess



import(
	"testing"
)

func TestSetRook(t *testing.T){
	t.Run("Rook: empty piece", func(t *testing.T){
		result := SetRook("", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Rook: set name, position is empty", func(t *testing.T){
		result := SetRook("rook", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Rook: set name and position", func(t *testing.T){
		result := SetRook("rook", "H1")
		if result.Name == "" || result.Position == nil{
			t.Errorf("King: Expected type of piece got %v", result)
		}
	})
}

func TestPredictRookMoves(t *testing.T){
	t.Run("Rook: set position in the middle", func(t *testing.T){
		result := SetRook("rook", "D5")
		expected := []string{"E5","F5","G5","H5","D4","D3","D2","D1","C5", "B5", "A5", "D6", "D7", "D8"}
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

