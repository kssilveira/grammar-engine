package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/kssilveira/grammar-engine/pronoun"
)

type Verb string

const (
	VerbToBeCalled = "be called"
)

type Noun string

const (
	NounName = "Rafael"
)

type Sentence struct {
	Pronoun pronoun.Type
	Verb    Verb
	Object  Noun
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
	Verbs        map[Verb]string
	VerbSuffixes map[pronoun.Type]string
	Nouns        map[Noun]string
}

func (l Language) Pronoun(p pronoun.Type) string {
	return l.Pronouns[p]
}

func (l Language) Verb(v Verb, p pronoun.Type) string {
	return l.Verbs[v] + l.VerbSuffixes[p]
}

func (l Language) Noun(n Noun) string {
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
		Pronoun: pronoun.I, Verb: VerbToBeCalled, Object: NounName,
	}}
	languages := []Language{{
		Pronouns: map[pronoun.Type]string{
			pronoun.I: "ich",
		},
		Verbs: map[Verb]string{
			VerbToBeCalled: "heiß",
		},
		VerbSuffixes: map[pronoun.Type]string{
			pronoun.I: "e",
		},
		Nouns: map[Noun]string{
			NounName: "Rafael",
		},
	}}
	for _, sentence := range sentences {
		for _, language := range languages {
			fmt.Print(sentence.On(language))
		}
	}
}
