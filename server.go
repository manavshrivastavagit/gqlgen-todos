package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"os"

	// "github.com/99designs/gqlgen/graphql/handler"
	// "github.com/99designs/gqlgen/graphql/playground"
	// "github.com/vektah/gqlgen-todos/graph"
	// "github.com/vektah/gqlgen-todos/graph/generated"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	defaultPort := "8080"

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://user1:nlvduxngA56PvW5B@cluster0-8zwv9.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	// defer cur.Close(ctx)
	// ctx, err = context.WithTimeout(context.Background(), 2*time.Second)
	// err = client.Ping(ctx, readpref.Primary())

	databases, err1 := client.ListDatabaseNames(ctx, bson.M{})
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(databases)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
