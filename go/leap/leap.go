package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear takes a year and returns true if it is a leap year
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
