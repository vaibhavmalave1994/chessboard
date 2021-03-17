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

	t.Run("set name and position", func(t *testing.T){
		result := SetPiece("king", "H1")
		if reflect.TypeOf(result).Implements(reflect.TypeOf((*Piece)(nil)).Elem()) == false{
			t.Errorf("Expected type of piece got %v", result)
		}
	})
}
