package lib

func hasSuffix(s, suffix string) bool {
	return s[len(s)-len(suffix):] == suffix
}

func isValidFile(fileName string) bool {
	return hasSuffix(fileName, ".txt")
}
