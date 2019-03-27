package hello

func Hello(msg string) string {
	if msg == "" {
		msg = "world"
	}
	return "[v3] Hello, " + msg
}
