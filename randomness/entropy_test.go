package randomness

import "testing"

func TestCalculateNonUniformEntropy(t *testing.T){
	p1 := .5
	p2 := .5
	expectedResult := 1.0
	result := CalculateNonUniformEntropy(p1,p2)
	if result != expectedResult{
		t.Errorf("Expected Entropy to be %4.f but the result was %4.f",expectedResult,result)
	}
}

func TestCalculateUniformEntropy(t *testing.T) {
	posibilities := .5
	numberOfposibilities := 2
	expectedResult := 1.0
	result := CalculateUniformEntropy(posibilities,numberOfposibilities)
	if result != expectedResult{
		t.Errorf("Expected Entropy to be %4.f but the result was %4.f",expectedResult,result)
	}
}