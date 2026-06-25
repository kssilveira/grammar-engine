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
	}, {
		Pronoun: pronoun.You, Verb: verb.ToBeCalled, Question: pronoun.How,
	}}
	languages := []language.Type{{
		Pronouns: map[pronoun.Type]string{
			pronoun.I:   "ich",
			pronoun.My:  "mein",
			pronoun.You: "Sie",
			pronoun.How: "wie",
		},
		Verbs: map[verb.Type]string{
			verb.ToBeCalled: "heiß",
			verb.ToBe:       "ist",
		},
		VerbSuffixes: map[pronoun.Type]string{
			pronoun.I:   "e",
			pronoun.You: "en",
		},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "Name",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I:   "I",
			pronoun.My:  "my",
			pronoun.You: "you",
			pronoun.How: "how",
		},
		Verbs:        map[verb.Type]string{},
		VerbSuffixes: map[pronoun.Type]string{},
		VerbsPassive: map[verb.Type]string{
			verb.ToBeCalled: "called",
		},
		VerbsIrregular: map[verb.Type]map[pronoun.Type]string{
			verb.ToBe: {
				pronoun.I:   "am",
				pronoun.My:  "is",
				pronoun.You: "are",
			},
		},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "name",
		},
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I:   "je",
			pronoun.My:  "mon",
			pronoun.You: "vous",
			pronoun.How: "comment",
		},
		Verbs:        map[verb.Type]string{},
		VerbSuffixes: map[pronoun.Type]string{},
		VerbsIrregular: map[verb.Type]map[pronoun.Type]string{
			verb.ToBeCalled: {
				pronoun.I:   "m'appelle",
				pronoun.You: "appelez-vous",
			},
			verb.ToBe: {
				pronoun.My: "est",
			},
		},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "nom",
		},
		IsQuestionPronounBeforeVerb: true,
	}, {
		Pronouns: map[pronoun.Type]string{
			pronoun.I:   "yo",
			pronoun.My:  "mi",
			pronoun.You: "usted",
			pronoun.How: "cómo",
		},
		Verbs:        map[verb.Type]string{},
		VerbSuffixes: map[pronoun.Type]string{},
		VerbsIrregular: map[verb.Type]map[pronoun.Type]string{
			verb.ToBeCalled: {
				pronoun.I:   "me llamo",
				pronoun.You: "se llama",
			},
			verb.ToBe: {
				pronoun.My: "es",
			},
		},
		Nouns: map[noun.Type]string{
			noun.PersonName: "Rafael",
			noun.Name:       "nombre",
		},
		IsUpsideDownQuestionMark: true,
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
