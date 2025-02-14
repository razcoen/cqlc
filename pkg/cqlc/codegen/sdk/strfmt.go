package sdk

import (
	"regexp"
	"strings"

	"github.com/jinzhu/inflection"
)

func ToSingularPascalCase(s string) string {
	delimiters := regexp.MustCompile(`[-_]+`)
	normalized := delimiters.ReplaceAllString(s, " ")
	camelOrPascal := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	normalized = camelOrPascal.ReplaceAllString(normalized, `$1 $2`)
	words := strings.Fields(normalized)
	var initialism = map[string]struct{}{
		"API": {}, "HTTP": {}, "ID": {}, "URL": {}, "JSON": {}, "XML": {}, "ACL": {}, "CPU": {}, "UID": {}, "SQL": {}, "IP": {},
		"UUID": {}, "IO": {}, "SSH": {}, "TLS": {}, "SSL": {}, "DNS": {}, "XSS": {}, "CSRF": {}, "JWT": {}, "HTML": {}, "CSS": {},
	}
	capitalizeWord := func(word string) string {
		upper := strings.ToUpper(word)
		if _, isInitialism := initialism[upper]; isInitialism {
			return upper
		}
		return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}
	for i, word := range words {
		if i == len(words)-1 {
			word = inflection.Singular(word)
		}
		word = capitalizeWord(word)
		words[i] = word
	}

	s = strings.Join(words, "")
	return s
}

func ToSingularSnakeCase(s string) string {
	delimiters := regexp.MustCompile(`[-_]+`)
	normalized := delimiters.ReplaceAllString(s, " ")
	camelOrPascal := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	normalized = camelOrPascal.ReplaceAllString(normalized, `$1 $2`)
	words := strings.Fields(normalized)
	for i, word := range words {
		word = strings.ToLower(word)
		if i == len(words)-1 {
			word = inflection.Singular(word)
		}
		words[i] = word
	}
	s = strings.Join(words, "_")
	return s
}

func ToSnakeCase(s string) string {
	delimiters := regexp.MustCompile(`[-_]+`)
	normalized := delimiters.ReplaceAllString(s, " ")
	camelOrPascal := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	normalized = camelOrPascal.ReplaceAllString(normalized, `$1 $2`)
	words := strings.Fields(normalized)
	for i, word := range words {
		word = strings.ToLower(word)
		words[i] = word
	}
	s = strings.Join(words, "_")
	return s
}
