package chess



import(
	"testing"
)

func TestSetBishop(t *testing.T){
	t.Run("Bishop: empty piece", func(t *testing.T){
		result := SetBishop("", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Bishop: set name, position is empty", func(t *testing.T){
		result := SetBishop("bishop", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Bishop: set name and position", func(t *testing.T){
		result := SetBishop("bishop", "H1")
		if result.Name == "" || result.Position == nil{
			t.Errorf("Bishop: Expected type of piece got %v", result)
		}
	})
}

func TestPredictBishopMoves(t *testing.T){
	t.Run("Bishop: set position in the middle", func(t *testing.T){
		result := SetBishop("bishop", "D5")
		expected := []string{"E6", "F7", "G8", "C4", "B3", "A2", "C6", "B7", "A8", "E4", "F3", "G2", "H1"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Bishop: set position to bottom right corner", func(t *testing.T){
		result := SetBishop("queen", "A1")
		expected := []string{"B2","C3","D4","E5", "F6", "G7", "H8"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Bishop: set position to bottom left corner", func(t *testing.T){
		result := SetBishop("queen", "A8")
		expected := []string{"B7","C6","D5","E4", "F3", "G2", "H1"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Bishop: set position to top left corner", func(t *testing.T){
		result := SetBishop("queen", "H1")
		expected := []string{"A8","B7","C6","D5","E4", "F3", "G2"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Bishop: set position to top right corner", func(t *testing.T){
		result := SetBishop("queen", "H8")
		expected := []string{"A1","B2","C3","D4","E5", "F6", "G7"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})
	
}

