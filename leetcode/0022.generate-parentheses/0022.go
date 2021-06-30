package leetcode

type answer []string

var (
	answers [9]answer
)

func _generateParenthesis(n int) []string {
	var result []string
	for _, item := range answers[n-1] {
		result = append(result, "("+item+")")
	}
	for i := 1; i < n; i++ {
		for _, leftitem := range answers[i-1] {
			for _, rightitem := range answers[n-i] {
				result = append(result, "("+leftitem+")"+rightitem)
			}
		}
	}
	return result
}

func generateParenthesis(n int) []string {
	answers[0] = []string{""}
	answers[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		answers[i] = _generateParenthesis(i)
	}
	return answers[n]
}
