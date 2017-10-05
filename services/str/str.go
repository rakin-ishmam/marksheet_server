package str

// ConcatBySpliter returns string tokenized by spliter
func ConcatBySpliter(spliter string, vals ...string) string {
	format := ""
	for i, v := range vals {
		if i > 0 {
			format += spliter
		}

		format += v
	}

	return format
}
