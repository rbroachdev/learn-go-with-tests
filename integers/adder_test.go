package integers

import "testing"

func TestAdd(t *testing.T) {

	t.Run("Adding 2 + 2", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		if got != want {
			t.Errorf("Expected %d got %d", want, got)
		}
	})

}
