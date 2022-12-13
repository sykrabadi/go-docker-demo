package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongo(ctx context.Context)(*mongo.Database, error){
	conn, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		return nil, err
	}

	err = conn.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return conn.Database("go-docker-demo"), nil
}

func insert()error{
	ctx := context.Background()
	db, err := connectToMongo(ctx)
	if err != nil {
		log.Fatal()
		return err
	}

	document := bson.D{
		{"name", "rocky balboa"},
	}
	res, err := db.Collection("document").InsertOne(ctx, document)

	if err != nil {
		log.Fatalf("error InsertOne document with error %v \n", err)
		return err
	}

	log.Printf("successfully inserted document with _id : %v \n", res.InsertedID)
	return nil
}

func HelloDocker(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Halo Docker! \n")
	err := insert()
	if err != nil {
		log.Fatalf("error with error : %v \n", err)
	}
}

func serveHTTP(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloDocker)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("error listen to port %v \n", addr)
	}
	log.Printf("Listening to %v", addr)
}

func main() {
	serveHTTP(":8080")
}
