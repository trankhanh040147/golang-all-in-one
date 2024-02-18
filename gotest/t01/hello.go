package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const franceHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "France":
		prefix = franceHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
}
