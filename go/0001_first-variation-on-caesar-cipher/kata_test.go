package kata

import (
	"testing"
)

func TestMovingShift(t *testing.T) {
	ciphered := MovingShift("I should have known that you would have a perfect answer for me!!!", 1)
	deciphered := DemovingShift(ciphered, 1)
	t.Error(ciphered)
	t.Error(deciphered)
	t.Error(len(ciphered))
}
