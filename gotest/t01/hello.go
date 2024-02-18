package hello

const englishHelloPrefix = "Hello,"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}

func main() {
}
