package helper

func inArray(value any, array []any) bool {
	for _, val := range array {
		if val == value {
			return true
		}
	}
	return false
}
