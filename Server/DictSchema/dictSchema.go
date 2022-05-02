package dictschema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func BuildSchema() graphql.SchemaConfig {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"query": &graphql.Field{
			Type: graphql.NewList(DefinitionType),
			Args: graphql.FieldConfigArgument{
				"word": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// fmt.Println(p.Args["word"])
				resp, err := http.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", p.Args["word"]))

				if err != nil {
					log.Fatalln(err)
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}

				// fmt.Println(string(body))
				// definition := WordDefinition{
				// 	Word:     "abcde",
				// 	Phonetic: "phone",
				// 	Phonetics: []WordPhonetic{
				// 		{
				// 			Text:  "phoned",
				// 			Audio: "sound",
				// 		},
				// 	},
				// 	Origin: "xmen",
				// }
				var definition []WordDefinition
				json.Unmarshal(body, &definition)
				// fmt.Printf("%+v", definition)

				return definition, nil

				// return resp.Body, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return schemaConfig

}
