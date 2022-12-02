package utils

func isPrime(no int) bool {
	for i := int(2); i < (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
