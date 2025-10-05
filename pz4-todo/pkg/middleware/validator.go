package middleware

func TitleValidation(title string) bool {
	return len(title) >= 3 && len(title) <= 100
}
