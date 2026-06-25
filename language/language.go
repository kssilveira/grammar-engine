package language

import (
	"github.com/kssilveira/grammar-engine/noun"
	"github.com/kssilveira/grammar-engine/pronoun"
	"github.com/kssilveira/grammar-engine/verb"
)

type Type struct {
	Pronouns     map[pronoun.Type]string
	Verbs        map[verb.Type]string
	VerbSuffixes map[pronoun.Type]string
	Nouns        map[noun.Type]string
}

func (t Type) Pronoun(p pronoun.Type) string {
	return t.Pronouns[p]
}

func (t Type) Verb(v verb.Type, p pronoun.Type) string {
	return t.Verbs[v] + t.VerbSuffixes[p]
}

func (t Type) Noun(n noun.Type) string {
	return t.Nouns[n]
}
