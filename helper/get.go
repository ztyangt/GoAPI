package helper

func getType(value any) string {
	attr, _ := typeof(value)
	return attr
}
