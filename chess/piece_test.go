package chess


import(
	"testing"
	"reflect"
)

func TestPiece(t *testing.T){
	t.Run("empty piece", func(t *testing.T){
		result := SetPiece("", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("set name, position is empty", func(t *testing.T){
		result := SetPiece("king", "")
		if result != nil{
			t.Errorf("Expected nil got %v", result)
		}
	})

	t.Run("set name and position for king", func(t *testing.T){
		result := SetPiece("king", "H1")
		if reflect.TypeOf(result).Implements(reflect.TypeOf((*Piece)(nil)).Elem()) == false{
			t.Errorf("Expected type of piece got %v", result)
		}
	})

	t.Run("set name and position for pawn", func(t *testing.T){
		result := SetPiece("pawn", "H1")
		if reflect.TypeOf(result).Implements(reflect.TypeOf((*Piece)(nil)).Elem()) == false{
			t.Errorf("Expected type of piece got %v", result)
		}
	})


	t.Run("set name and position for queen", func(t *testing.T){
		result := SetPiece("queen", "H1")
		if reflect.TypeOf(result).Implements(reflect.TypeOf((*Piece)(nil)).Elem()) == false{
			t.Errorf("Expected type of piece got %v", result)
		}
	})

	t.Run("set name and position for rook", func(t *testing.T){
		result := SetPiece("rook", "H1")
		if reflect.TypeOf(result).Implements(reflect.TypeOf((*Piece)(nil)).Elem()) == false{
			t.Errorf("Expected type of piece got %v", result)
		}
	})

	t.Run("set name and position for bishop", func(t *testing.T){
		result := SetPiece("bishop", "H1")
		if reflect.TypeOf(result).Implements(reflect.TypeOf((*Piece)(nil)).Elem()) == false{
			t.Errorf("Expected type of piece got %v", result)
		}
	})

}
