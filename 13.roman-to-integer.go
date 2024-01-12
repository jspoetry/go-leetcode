/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
var romanNumeralsMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
var romanDoubledNumeralsMap = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	acc := 0
	// we use pointer instead of ordinary for loop
	// cause we need increase by 2 on doubled literals
	pointer := 0

	for pointer < len(s) {
		literal := string(s[pointer])
		literalValue := romanNumeralsMap[literal]
		nextLiteralIdx := pointer + 1
		// check next literal
		if nextLiteralIdx < len(s) {
			nextLiteral := string(s[nextLiteralIdx])
			doubledLiteral := literal + nextLiteral
			nextLiteralValue, isExist := romanDoubledNumeralsMap[doubledLiteral]

			// if the doubled literal exists on map so we replace the value
			if isExist {
				literalValue = nextLiteralValue
				pointer += 1
			}
		}

		acc += literalValue
		pointer += 1
	}

	return acc
}

// @lc code=end