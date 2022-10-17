package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the homepage")
	checkErr(err)
}
func getHello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func connectDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Printf("%v", err)
}

//InsertDB(context.TODO(), client)
func insertDB(ctx context.Context, client *mongo.Client, user primitive.D) {
	fmt.Printf("\nINSERTING %v\n", user)
	usersCollection := client.Database("fyp_test").Collection("test")
	result, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
	fmt.Printf("%v", result)
}

func loginRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
		fmt.Printf("%v %v", r.Form["username"], r.Form["password"])
	}
	username := r.Form["username"]
	password := r.Form["password"]
	//add checks before calling insert
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	user := bson.D{{"fullName", username}, {"age", password}}
	insertDB(context.TODO(), client, user)
}

func main() {
	// http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	fs := http.FileServer(http.Dir("../../../dist"))
	http.Handle("/", fs)
	http.HandleFunc("/login", loginRequest)

	listenerErr := http.ListenAndServe(":3333", nil)
	fmt.Printf("%v", listenerErr)
}
