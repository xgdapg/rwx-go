package main

var keywords map[string]struct{}

func init() {
	keywords = map[string]struct{}{}
	kws := []string{"fn", "if", "else", "var", "const", "return", "struct", "impl", "while", "each", "in", "continue", "break", "interface", "true", "false"}
	for _, kw := range kws {
		keywords[kw] = struct{}{}
	}
}
