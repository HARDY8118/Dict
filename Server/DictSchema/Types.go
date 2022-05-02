package dictschema

import "github.com/graphql-go/graphql"

var PhoneticType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Phonetic",
	Fields: graphql.Fields{
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"audio": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var MeaningType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Meanings",
	Fields: graphql.Fields{
		"partOfSpeech": &graphql.Field{
			Type: graphql.String,
		},
		"definitions": &graphql.Field{
			Type: &graphql.List{
				OfType: graphql.NewObject(graphql.ObjectConfig{
					Name: "Definition",
					Fields: graphql.Fields{
						"definition": &graphql.Field{
							Type: graphql.String,
						},
						"synonyms": &graphql.Field{
							Type: graphql.String,
						},
						"antonyms": &graphql.Field{
							Type: graphql.String,
						},
					},
				}),
			},
		},
		"synonyms": &graphql.Field{
			Type: &graphql.List{
				OfType: graphql.String,
			},
		},
		"antonyms": &graphql.Field{
			Type: &graphql.List{
				OfType: graphql.String,
			},
		},
	},
})

var DefinitionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Word",
	Fields: graphql.Fields{
		"word": &graphql.Field{
			Type: graphql.String,
		},
		// "phonetic": &graphql.Field{
		// 	Type: graphql.String,
		// },
		"phonetics": &graphql.Field{
			Type: &graphql.List{
				OfType: PhoneticType,
			},
		},
		// "origin": &graphql.Field{
		// 	Type: graphql.String,
		// },
		"meanings": &graphql.Field{
			Type: &graphql.List{
				OfType: MeaningType,
			},
		},
	},
})

type WordPhonetic struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}

type WordMeanings struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Definitions  []struct {
		Definition string   `json:"definition"`
		Synonyms   []string `json:"synonyms"`
		Antonyms   []string `json:"antonyms"`
		Example    []string `json:"example"`
	} `json:"definitions"`
	Synonyms []string `json:"synonyms"`
	Antonyms []string `json:"antonyms"`
}

type WordDefinition struct {
	Word string `json:"word"`
	// Phonetic  string         `json:"phonetic"`
	Phonetics []WordPhonetic `json:"phonetics"`
	// Origin    string         `json:"origin"`
	Meanings []WordMeanings `json:"meanings"`
}
