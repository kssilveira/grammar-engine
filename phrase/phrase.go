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
	Pronoun pronoun.Type
	Verb    verb.Type
	Object  noun.Type
}

func (t Type) On(l language.Type) string {
	return strings.Join([]string{
		capitalizeFirst(l.Pronoun(t.Pronoun)),
		l.Verb(t.Verb, t.Pronoun),
		l.Noun(t.Object),
	}, " ") + ".\n"
}

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
