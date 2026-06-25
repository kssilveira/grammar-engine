package language

import (
	"github.com/kssilveira/grammar-engine/noun"
	"github.com/kssilveira/grammar-engine/pronoun"
	"github.com/kssilveira/grammar-engine/verb"
)

type Type struct {
	Pronouns                    map[pronoun.Type]string
	Verbs                       map[verb.Type]string
	VerbSuffixes                map[pronoun.Type]string
	VerbsIrregular              map[verb.Type]map[pronoun.Type]string
	VerbsPassive                map[verb.Type]string
	Nouns                       map[noun.Type]string
	IsQuestionPronounBeforeVerb bool
	IsUpsideDownQuestionMark    bool
}

func (t Type) Pronoun(p pronoun.Type) string {
	return t.Pronouns[p]
}

func (t Type) Verb(v verb.Type, p pronoun.Type) []string {
	passive, ok := t.VerbsPassive[v]
	if ok {
		v = verb.ToBe
	}
	res := t.Verbs[v] + t.VerbSuffixes[p]
	if verb, ok := t.VerbsIrregular[v]; ok {
		res = verb[p]
	}
	return []string{res, passive}
}

func (t Type) Noun(n noun.Type) string {
	return t.Nouns[n]
}
