package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/kssilveira/grammar-engine/noun"
	"github.com/kssilveira/grammar-engine/pronoun"
	"github.com/kssilveira/grammar-engine/verb"
)

type Sentence struct {
	Pronoun pronoun.Type
	Verb    verb.Type
	Object  noun.Type
}

func (s Sentence) On(l Language) string {
	return strings.Join([]string{
		CapitalizeFirst(l.Pronoun(s.Pronoun)),
		l.Verb(s.Verb, s.Pronoun),
		l.Noun(s.Object),
	}, " ") + ".\n"
}

type Language struct {
	Pronouns     map[pronoun.Type]string
	Verbs        map[verb.Type]string
	VerbSuffixes map[pronoun.Type]string
	Nouns        map[noun.Type]string
}

func (l Language) Pronoun(p pronoun.Type) string {
	return l.Pronouns[p]
}

func (l Language) Verb(v verb.Type, p pronoun.Type) string {
	return l.Verbs[v] + l.VerbSuffixes[p]
}

func (l Language) Noun(n noun.Type) string {
	return l.Nouns[n]
}

func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

func main() {
	sentences := []Sentence{{
		Pronoun: pronoun.I, Verb: verb.ToBeCalled, Object: noun.Name,
	}}
	languages := []Language{{
		Pronouns: map[pronoun.Type]string{
			pronoun.I: "ich",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "heiß",
		},
		VerbSuffixes: map[pronoun.Type]string{
			pronoun.I: "e",
		},
		Nouns: map[noun.Type]string{
			noun.Name: "Rafael",
		},
	}}
	for _, sentence := range sentences {
		for _, language := range languages {
			fmt.Print(sentence.On(language))
		}
	}
}
