package main

import (
	"fmt"

	"github.com/kssilveira/grammar-engine/language"
	"github.com/kssilveira/grammar-engine/noun"
	"github.com/kssilveira/grammar-engine/phrase"
	"github.com/kssilveira/grammar-engine/pronoun"
	"github.com/kssilveira/grammar-engine/verb"
)

func main() {
	sentences := []phrase.Type{{
		Pronoun: pronoun.I, Verb: verb.ToBeCalled, Object: noun.Name,
	}}
	languages := []language.Type{{
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
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I: "I",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "am called",
		},
		VerbSuffixes: map[pronoun.Type]string{},
		Nouns: map[noun.Type]string{
			noun.Name: "Rafael",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I: "je",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "m'appelle",
		},
		VerbSuffixes: map[pronoun.Type]string{},
		Nouns: map[noun.Type]string{
			noun.Name: "Rafael",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I: "me",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "llamo",
		},
		VerbSuffixes: map[pronoun.Type]string{},
		Nouns: map[noun.Type]string{
			noun.Name: "Rafael",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I: "eu",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "me chamo",
		},
		VerbSuffixes: map[pronoun.Type]string{},
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
