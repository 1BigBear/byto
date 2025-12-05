package parser

type Parser interface {
	Parse(input string) (map[string]string, error)
}
