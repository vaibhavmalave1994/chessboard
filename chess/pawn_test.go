package chess



import(
	"testing"
)

func TestSetPawn(t *testing.T){
	t.Run("Pawn: empty piece", func(t *testing.T){
		result := SetPawn("", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Pawn: set name, position is empty", func(t *testing.T){
		result := SetPawn("pawn", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("Pawn: set name and position", func(t *testing.T){
		result := SetPawn("pawn", "H1")
		if result.Name == "" || result.Position == nil{
			t.Errorf("Pawn: Expected type of piece got %v", result)
		}
	})
}

func TestPredictPawnMoves(t *testing.T){
	t.Run("Pawn: set position in the middle", func(t *testing.T){
		result := SetPawn("pawn", "D5")
		expected := []string{"E5"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Pawn: set position to bottom right corner", func(t *testing.T){
		result := SetPawn("pawn", "A1")
		expected := []string{"B1"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Pawn: set position to bottom left corner", func(t *testing.T){
		result := SetPawn("pawn", "A8")
		expected := []string{"B8"}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Pawn: set position to top left corner", func(t *testing.T){
		result := SetPawn("pawn", "H8")
		expected := []string{}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})

	t.Run("Pawn: set position to top right corner", func(t *testing.T){
		result := SetPawn("pawn", "H1")
		expected := []string{}
		positions := result.PredictMoves()
		if comparePostions(expected, positions) == false{
			t.Errorf("expected %v got %v", expected, positions)
		}
		
	})
	
}
