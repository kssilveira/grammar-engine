package phrase

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/kssilveira/grammar-engine/language"
	"github.com/kssilveira/grammar-engine/noun"
	"github.com/kssilveira/grammar-engine/pronoun"
	"github.com/kssilveira/grammar-engine/verb"
)

type Type struct {
	Pronoun  pronoun.Type
	Subject  noun.Type
	Verb     verb.Type
	Object   noun.Type
	Question pronoun.Type
}

func (t Type) On(l language.Type) string {
	words := [][]string{
		{capitalizeFirst(l.Pronoun(t.Pronoun)), l.Noun(t.Subject)},
		l.Verb(t.Verb, t.Pronoun),
		{l.Noun(t.Object)},
	}
	end := "."
	if t.Question != "" {
		words = [][]string{
			{capitalizeFirst(l.Pronoun(t.Question))},
			l.Verb(t.Verb, t.Pronoun),
			{l.Pronoun(t.Pronoun)},
		}
		end = "?"
	}
	valid := []string{}
	for _, list := range words {
		for _, word := range list {
			if word != "" {
				valid = append(valid, word)
			}
		}
	}
	return strings.Join(valid, " ") + end + "\n"
}

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
