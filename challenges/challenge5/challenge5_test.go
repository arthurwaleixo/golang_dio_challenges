package main

import "testing"

// ShouldSumCorrect
func TestShouldSumCorrect(t *testing.T) {
	teste := somar(1, 3, 5)
	resultado := 9

	if teste != resultado {
		t.Errorf("Valor esperado: %d, valor retornado: %d.", resultado, teste)
	}
}

// ShouldSumIncorrect
func TestShouldSumIncorrect(t *testing.T) {
	teste := somar(1, 3, 5)
	resultado := 8

	if teste != resultado {
		t.Errorf("Valor esperado: %d, valor retornado: %d.", resultado, teste)
	}
}

// ShouldMultiplyCorrect
func TestShouldMultiplyCorrect(t *testing.T) {
	teste := multiplicar(4, 5)
	resultado := 20

	if teste != resultado {
		t.Errorf("Valor esperado: %d, valor retornado: %d.", resultado, teste)
	}
}

// ShouldMultiplyIncorrect
func TestShouldMultiplyIncorrect(t *testing.T) {
	teste := multiplicar(4, 5)
	resultado := 19

	if teste != resultado {
		t.Errorf("Valor esperado: %d, valor retornado: %d.", resultado, teste)
	}
}

// ShouldSubtractCorrect
func TestShouldSubtractCorrect(t *testing.T) {
	teste := subtrair(3, 10)
	resultado := 7

	if teste != resultado {
		t.Errorf("Valor esperado: %d, valor retornado: %d.", resultado, teste)
	}
}

// ShouldSubtractIncorrect
func TestShouldSubtractIncorrect(t *testing.T) {
	teste := subtrair(3, 10)
	resultado := 6

	if teste != resultado {
		t.Errorf("Valor esperado: %d, valor retornado: %d.", resultado, teste)
	}
}

// ShouldDivideCorrect
func TestShouldDivideCorrect(t *testing.T) {
	teste := dividir(2, 10)
	resultado := 5

	if teste != resultado {
		t.Errorf("Valor esperado: %d, valor retornado: %d.", resultado, teste)
	}
}

// ShouldDivideIncorrect
func TestShouldDivideIncorrect(t *testing.T) {
	teste := dividir(2, 10)
	resultado := 4

	if teste != resultado {
		t.Errorf("Valor esperado: %d, valor retornado: %d.", resultado, teste)
	}
}
