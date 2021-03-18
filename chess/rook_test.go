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

	t.Run("Rook: set position to bottom right corner", func(t *testing.T){
		result := SetRook("rook", "A1")
		expected := []string{"A2","A3","A4","A5","A6","A7", "A8","B1","C1","D1","E1","F1","G1", "H1"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Rook: set position to bottom left corner", func(t *testing.T){
		result := SetRook("rook", "A8")
		expected := []string{"A7","A6","A5","A4","A3","A2", "A1","B8","C8","D8","E8","F8","G8", "H8"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Rook: set position to top left corner", func(t *testing.T){
		result := SetRook("rook", "H8")
		expected := []string{"H7","H6","H5","H4","H3","H2", "H1","B8","C8","D8","E8","F8","G8", "A8"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Rook: set position to top right corner", func(t *testing.T){
		result := SetRook("rook", "H1")
		expected := []string{"H7","H6","H5","H4","H3","H2", "H8","B1","C1","D1","E1","F1","G1", "A1"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})
	
	
}

