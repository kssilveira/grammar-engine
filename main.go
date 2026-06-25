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
	phrases := []phrase.Type{{
		Pronoun: pronoun.I, Verb: verb.ToBeCalled, Object: noun.PersonName,
	}, {
		Pronoun: pronoun.My, Subject: noun.Name, Verb: verb.ToBe, Object: noun.PersonName,
	}}
	languages := []language.Type{{
		Pronouns: map[pronoun.Type]string{
			pronoun.I:  "ich",
			pronoun.My: "mein",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "heiß",
			verb.ToBe:       "ist",
		},
		VerbSuffixes: map[pronoun.Type]string{
			pronoun.I: "e",
		},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "Name",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I:  "I",
			pronoun.My: "my",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "am called",
			verb.ToBe:       "is",
		},
		VerbSuffixes: map[pronoun.Type]string{},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "name",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I:  "je",
			pronoun.My: "mon",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "m'appelle",
			verb.ToBe:       "est",
		},
		VerbSuffixes: map[pronoun.Type]string{},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "nom",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I:  "me",
			pronoun.My: "mi",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "llamo",
			verb.ToBe:       "es",
		},
		VerbSuffixes: map[pronoun.Type]string{},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "nombre",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I:  "eu",
			pronoun.My: "meu",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "me chamo",
			verb.ToBe:       "é",
		},
		VerbSuffixes: map[pronoun.Type]string{},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "nome",
		},
	}}
	for _, phrase := range phrases {
		fmt.Println()
		for _, language := range languages {
			fmt.Print(phrase.On(language))
		}
	}
}
