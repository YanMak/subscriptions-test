package validatorХ

import "strconv"

func ptrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func parseOrDefault(str string, def int) int {
	if str == "" {
		return def
	}
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return def
}
