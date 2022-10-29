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

	"github.com/EpicStep/go-simple-geo/geo"
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

func checkDBSignup(ctx context.Context, client *mongo.Client, user string, collection string) (init bool) {
	fmt.Printf("\nCHECKING %v\n", user)
	usersCollection := client.Database("fyp_test").Collection(collection)
	usercheck := bson.D{{"fullName", user}}
	cursor, err := usersCollection.Find(context.TODO(), usercheck)

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if len(results) != 0 {
		return true
	}
	return false
}

func checkDBLogin(ctx context.Context, client *mongo.Client, data primitive.D, collection string) (init bool) {
	fmt.Printf("\nCHECKING %v %v\n", data[0], data[1])
	usersCollection := client.Database("fyp_test").Collection(collection)

	cursor, err := usersCollection.Find(context.TODO(), data)

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if len(results) != 0 {
		fmt.Println("user exists")
		return true
	}

	return false
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
	if userExists := checkDBLogin(context.TODO(), client, user, "users"); userExists {
		fmt.Printf("User exists %v", userExists)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

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

func renderTmpl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rendering")
	// render := render.New()
	// r.HTML(w, http.StatusOK, "example", nil)
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

	alreadySignedUp := checkDBSignup(context.TODO(), client, username[0], "users")
	if alreadySignedUp {
		fmt.Println("User already exists")
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	itemByted := encodeToByte(password)

	hashedVal := sha256.New()
	hashedVal.Write([]byte(itemByted))

	user := bson.D{{"fullName", username}, {"email", email}, {"password", hashedVal.Sum(nil)}}
	fmt.Printf("\nINSERTING THIS HASH-->%v\n ", hashedVal.Sum(nil))
	insertDB(context.TODO(), client, user, "users")
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

	distance := calculateDistance(floatLat, floatLong)
	flightDetailsResp(distance)
}

func calculateDistance(long float64, lat float64) float64 {
	c1 := geo.NewCoordinatesFromDegrees(32, 52.22)
	c2 := geo.NewCoordinatesFromDegrees(32, 52.999)
	distance := c1.Distance(c2).ToKilometers()
	fmt.Printf("%v", distance)
	return distance
}

func flightDetailsResp(distance float64) {
	url := "http://localhost:3333/#/reviewFlight"
	distanceStr := fmt.Sprintf("{\"distance\":\"%v\"}", distance)
	var jsonStr = []byte(distanceStr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
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

	// dist := calculateDistance(3.44, 3.44)
	listenerErr := http.ListenAndServe(":3333", nil)
	fmt.Printf("%v", listenerErr)
}
