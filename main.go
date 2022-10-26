package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"crypto/sha256"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Head struct {
	ContentType string
	Status      int
}

type HTML struct {
	Head      Head
	Name      string
	Templates *template.Template
	// contains filtered or unexported fields
}

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
func insertDB(ctx context.Context, client *mongo.Client, user primitive.D, collection string) {
	fmt.Printf("\nINSERTING %v\n", user)
	usersCollection := client.Database("fyp_test").Collection(collection)
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
		fmt.Printf("BEFORE HASH AND STORAGE--> %v %v", r.Form["username"], r.Form["password"])
	}
	username := r.Form["username"]
	password := r.Form["password"]
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	itemByted := encodeToByte(password)

	hashedVal := sha256.New()
	hashedVal.Write([]byte(itemByted))

	user := bson.D{{"fullName", username}, {"password", hashedVal.Sum(nil)}}
	// checkDB(context.TODO(), client, user)
	insertDB(context.TODO(), client, user, "users")
}

// encode the string array into byte array
func encodeToByte(pw []string) []byte {
	buf := &bytes.Buffer{}
	bytedPw := buf.Bytes()
	if len(pw) != 0 {
		gob.NewEncoder(buf).Encode(pw)
		bytedPw = buf.Bytes()
		return bytedPw
	}
	return bytedPw
}

// func checkDB(ctx context.Context, client *mongo.Client, user primitive.D) {
// }

func locationRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
		fmt.Printf("BEFORE HASH AND STORAGE--> %v %v", r.Form["latitude"], r.Form["longitude"])
	}

	latitude := r.Form["latitude"]
	longitude := r.Form["longitude"]
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	user := bson.D{{"latitude", latitude}, {"longitude", longitude}}
	// checkDB(context.TODO(), client, user)
	insertDB(context.TODO(), client, user, "coordinates")
}

func speedRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
		fmt.Printf("BEFORE HASH AND STORAGE--> %v %v %v", r.Form["lowspeed"], r.Form["midspeed"], r.Form["highspeed"])
	}

	lowspeed := r.Form["lowspeed"]
	midspeed := r.Form["midspeed"]
	highspeed := r.Form["highspeed"]

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	user := bson.D{{"lowspeed", lowspeed}, {"midspeed", midspeed}, {"highspeed", highspeed}}
	// checkDB(context.TODO(), client, user)
	insertDB(context.TODO(), client, user, "speeds")
}

func totalRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
		fmt.Printf("BEFORE HASH AND STORAGE-->%v %v %v %v %v %v %v", r.Form["latitude"], r.Form["longitude"], r.Form["lowspeed"], r.Form["highspeed"], r.Form["date"], r.Form["hour"], r.Form["minute"])
	}

	latitude := r.Form["latitude"]
	longitude := r.Form["longitude"]
	date := r.Form["date"]
	hour := r.Form["hour"]
	minute := r.Form["minute"]

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	user := bson.D{{"latitude", latitude}, {"longitude", longitude}, {"date", date}, {"hour", hour}, {"minute", minute}}
	insertDB(context.TODO(), client, user, "total")
	// http.Redirect(w, r, "/planReview", http.StatusSeeOther)

	lat := latitude[0]
	long := longitude[0]
	floatLat, _ := strconv.ParseFloat(lat, 32)
	floatLong, _ := strconv.ParseFloat(long, 32)

	calculateDistance(floatLat, floatLong)
}

func renderTmpl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rendering")
	// render := render.New()
	// r.HTML(w, http.StatusOK, "example", nil)
}

func calculateDistance(lat float64, long float64) {

}

func signupRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
		fmt.Printf("BEFORE HASH AND STORAGE--> %v %v", r.Form["username"], r.Form["password"])
	}
	username := r.Form["username"]
	password := r.Form["password"]
	email := r.Form["email"]
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	itemByted := encodeToByte(password)

	hashedVal := sha256.New()
	hashedVal.Write([]byte(itemByted))

	user := bson.D{{"fullName", username}, {"email", email}, {"password", hashedVal.Sum(nil)}}
	// checkDB(context.TODO(), client, user)
	insertDB(context.TODO(), client, user, "users")
}

func main() {
	// http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	fs := http.FileServer(http.Dir("../../../dist"))
	http.Handle("/", fs)
	http.HandleFunc("/login", loginRequest)
	http.HandleFunc("/signup", signupRequest)
	http.HandleFunc("/location", locationRequest)
	http.HandleFunc("/speed", speedRequest)
	http.HandleFunc("/detailsSubmit", totalRequest)
	http.HandleFunc("/planReview", renderTmpl)

	listenerErr := http.ListenAndServe(":3333", nil)
	fmt.Printf("%v", listenerErr)
}
