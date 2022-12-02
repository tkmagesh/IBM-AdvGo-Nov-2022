package utils

import "testing"

func Test_is_prime(t *testing.T) {
	//arrange
	no := 12
	expectedResult := true

	//act
	actualResult := isPrime(no)

	//assert
	if actualResult != expectedResult {
		//t.Fail()
		t.Errorf("isPrime(%d) : expected => %t but got %t\n", no, expectedResult, actualResult)
	}
}
