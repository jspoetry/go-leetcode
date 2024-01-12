/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(value string) {
	*s = append(*s, value)
}

func (s *Stack) pop() string {
	if s.isEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

var openBracketMap map[string]string = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

var closedBracketMap map[string]string = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

func isValid(s string) bool {
	stack := new(Stack)

	for _, val := range s {
		stringifiedVal := string(val)
		_, hasKey := openBracketMap[stringifiedVal]
		if (hasKey) {
			// it's open bracket, just push
			stack.push(stringifiedVal)
		} else {
			// it's closed bracket
			
			// get pair of closed bracket
			sameTypeOpenBracket := closedBracketMap[stringifiedVal]
			bracket := stack.pop()
			// if mismatch it's invalid sequence of brackets
			if (bracket != sameTypeOpenBracket) {
				return false
			}
		}
	}

	// check stack is empty: "[" is not valid case
	return stack.isEmpty()
}
// @lc code=end