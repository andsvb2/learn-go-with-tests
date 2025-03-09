package helloWorld

const (
	french     = "French"
	portuguese = "Portuguese"
	spanish    = "Spanish"

	englishHelloPrefix    = "Hello, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Oi, "
	spanishHelloPrefix    = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
