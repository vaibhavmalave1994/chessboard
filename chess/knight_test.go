package chess



import(
	"testing"
)

func TestSetKnight(t *testing.T){
	t.Run("Knight: empty piece", func(t *testing.T){
		result := SetKnight("", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Knight: set name, position is empty", func(t *testing.T){
		result := SetKnight("knight", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Knight: set name and position", func(t *testing.T){
		result := SetKnight("knight", "H1")
		if result.Name == "" || result.Position == nil{
			t.Errorf("King: Expected type of piece got %v", result)
		}
	})
}

func TestPredictKnightMoves(t *testing.T){
	t.Run("Knight: set position in the middle", func(t *testing.T){
		result := SetKnight("knight", "D5")
		expected := []string{"F6","F4","E7","C7","B6","B4","C3","E3"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Knight: set position to bottom right corner", func(t *testing.T){
		result := SetKnight("knight", "A1")
		expected := []string{"B3","C2"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Knight: set position to bottom left corner", func(t *testing.T){
		result := SetKnight("knight", "A8")
		expected := []string{"B6","C7"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Knight: set position to top left corner", func(t *testing.T){
		result := SetKnight("knight", "H8")
		expected := []string{"G6","F7"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Knight: set position to top right corner", func(t *testing.T){
		result := SetKnight("knight", "H1")
		expected := []string{"G3","F2"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})
	
}
