package utils

import "testing"

func Test_is_prime_7(t *testing.T) {
	//arrange
	no := 7
	expectedResult := true

	//act
	actualResult := isPrime(no)

	//assert
	if actualResult != expectedResult {
		//t.Fail()
		t.Errorf("isPrime(%d) : expected => %t but got %t\n", no, expectedResult, actualResult)
	}
}

func Test_is_prime(t *testing.T) {
	testDataList := []struct {
		name     string
		no       int
		expected bool
		actual   bool
	}{
		{name: "Testing 7 is prime", no: 7, expected: true},
		{name: "Testing 9 is prime", no: 9, expected: false},
		{name: "Testing 11 is prime", no: 11, expected: false},
		{name: "Testing 13 is prime", no: 13, expected: true},
		{name: "Testing 15 is prime", no: 15, expected: false},
	}
	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testData.actual = isPrime(testData.no)
			if testData.actual != testData.expected {
				//t.Fail()
				t.Errorf("isPrime(%d) : expected => %t but got %t\n", testData.no, testData.expected, testData.actual)
			}
		})
	}
}
