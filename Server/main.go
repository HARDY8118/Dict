package main

import (
	"log"

	dictschema "github.com/HARDY8118/first-server/Dict/DictSchema"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	schema, err := graphql.NewSchema(dictschema.BuildSchema())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	r := gin.New()

	// r.Any("/api", CORSMiddleware(), gin.WrapH(h))
	r.GET("/api", CORSMiddleware(), gin.WrapH(h))
	r.POST("/api", CORSMiddleware(), gin.WrapH(h))
	r.OPTIONS("/api", CORSMiddleware(), gin.WrapH(h))

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	r.Run(":8080")
}
