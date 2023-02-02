package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"crypto/sha256"

	"github.com/EpicStep/go-simple-geo/geo"
	"github.com/gorilla/sessions"
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

type Flight struct {
	Date        string `json:"date"`
	Hour        string `json:"hour"`
	Minute      string `json:"minute"`
	Srclat      string `json:"srclat"`
	Srclng      string `json:"srclng"`
	Destlat     string `json:"destlat"`
	Destlng     string `json:"destlng"`
	Speed       string `json:"speed"`
	Altitude    string `json:"altitude"`
	Orientation string `json:"orientation"`
	Corridor    string `json:"corridor"`
	Drone       Drone  `json:"drone"`
	// UserID     int64 `json:"userID"`
	// FlightTime int64 `json:"flightTime"`
	// Altitude   int64 `json:"altitude"`
	// UavID      int64 `json:"uavID"`
}

type Drone struct {
	Name   string `json:"name"`
	Model  string `json:"model"`
	Weight string `json:"weight"`
}

type GridofCoordinates struct {
	Coordinates []interface{} `json:"coordinates"`
}

type MidCoord interface {
	get() map[string]interface{}
}

type Coordinate struct {
	Id        string `json:"id"`
	Latitude  string `json:"lat"`
	Longitude string `json:"lng"`
	//Corridor  string `json:"corridor"`
}

type QueryDate struct {
	Date string `json:"date"`
}

type ResponseData struct {
	Message string `json:"message"`
}

type ResponseDataDate struct {
	Message []string `json:"message"`
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
func insertDB(ctx context.Context, client *mongo.Client, user primitive.D, collection string) (err error) {
	fmt.Printf("\nINSERTING %v\n", user)
	usersCollection := client.Database("fyp_test").Collection(collection)
	result, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		fmt.Printf("%v", err)
		panic(err)
	}
	fmt.Printf("%v", result)
	return err
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
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	if err != nil {
		return
	}
	//http.Redirect(w, r, "/planner", http.StatusSeeOther) //this goes to profile page
}

// encode the string array into byte array
// func encodeToByte(pw []string) []byte {
// 	buf := &bytes.Buffer{}
// 	bytedPw := buf.Bytes()
// 	if len(pw) != 0 {
// 		gob.NewEncoder(buf).Encode(pw)
// 		bytedPw = buf.Bytes()
// 		return bytedPw
// 	}
// 	return bytedPw
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
	cookie := verifyCookie(w, r)
	fmt.Printf("COOKIE-->%v", cookie)
	if !cookie {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

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
	fmt.Printf("%v", distance)
	getData(w, r, distance)
	// http.ServeFile(w, r, fmt.Sprintf("../../../src/views/PlanFlightComponent.vue"))
	// handle(w, r, "planner")
	http.Redirect(w, r, "/#/plannernjrgklnjkl", 307)
}

func getData(w http.ResponseWriter, r *http.Request, d float64) {
	s := fmt.Sprintf("%f", d) // s == "123.456000"
	json.NewEncoder(w).Encode(ResponseData{Message: s})
}

func calculateDistance(long float64, lat float64) float64 {
	c1 := geo.NewCoordinatesFromDegrees(32, 52.22)
	c2 := geo.NewCoordinatesFromDegrees(32, 52.999)
	distance := c1.Distance(c2).ToKilometers()
	fmt.Printf("%v", distance)
	return distance
}

func loginSubmitted(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

var store = sessions.NewCookieStore([]byte("super-secret"))

func main() {
	//AIzaSyCIuCS2q9yO9Pj_X-xFB7tSI187n5ivS_A
	// http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	fs := http.FileServer(http.Dir("../../../dist"))
	//http.Handle("/", http.StripPrefix("/", fs))
	http.Handle("/", fs)
	http.HandleFunc("/login", loginRequest)
	http.HandleFunc("/signup", signupRequest)
	http.HandleFunc("/location", locationRequest)
	http.HandleFunc("/speed", speedRequest)
	http.HandleFunc("/planner", totalRequest)
	http.HandleFunc("/loginSubmitted", loginSubmitted)
	http.HandleFunc("/getAllTimes", getAllTimes)
	http.HandleFunc("/storeFlight", storeFlight)
	http.HandleFunc("/getDateFlight", getDateFlight)
	http.HandleFunc("/storeGridCoordinates", storeGridCoordinates)
	http.HandleFunc("/fetchGridCoordinates", fetchGridCoordinates)

	// dist := calculateDistance(3.44, 3.44)
	listenerErr := http.ListenAndServe(":3333", nil)
	fmt.Printf("%v", listenerErr)
}

func handle(w http.ResponseWriter, r *http.Request, name string) {
	fs := http.FileServer(http.Dir("../../../dist"))
	url := fmt.Sprintf("//%v", name)
	http.Handle(url, fs)
}

func fetchGridCoordinates(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY:", string(body))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	usersCollection := client.Database("fyp_test").Collection("grid")
	filter := bson.M{}
	result, err := usersCollection.Find(ctx, filter)
	var results []bson.M
	if err = result.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Printf("%v results", results)

	var coords []string
	for _, doc := range results {
		stringDoc := fmt.Sprint(doc) + ","
		fmt.Printf(stringDoc)
		coords = append(coords, fmt.Sprint(stringDoc))
		// c := &Coordinate{
		// 	Id: doc["id"].(string),
		// }
		// b, err := json.Marshal(c)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Printf("\njson.Marshal(c)--> ", string(b))
		fmt.Fprintf(w, stringDoc)
	}

	//convert times (type interface) to type string
	// var cdsStr []string
	// for _, v := range cds {
	// 	valStr := fmt.Sprint(v)
	// 	cdsStr = append(cdsStr, valStr)
	// 	fmt.Fprintf(w, valStr+",")
	// }

	fmt.Printf("\n'%v' matching docs found\n", len(results))
	fmt.Fprintf(w, "Hola")

}

func storeGridCoordinates(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY:", string(body))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	var grid GridofCoordinates
	err = json.Unmarshal(body, &grid)
	fmt.Printf("COORDSLIST: %vFINISHED", grid.Coordinates)

	var coord Coordinate
	for _, val := range grid.Coordinates {
		//convert from interface to map[string]interface{}
		t := val.(map[string]interface{})
		coord.Id = t["id"].(string)
		coord.Latitude = t["lat"].(string)
		coord.Longitude = t["lng"].(string)
		fmt.Printf("coord: %v\t %v\t %v\n", coord.Id, coord.Latitude, coord.Longitude)

		//build mongo record
		gridCoord := bson.D{{"lat", coord.Latitude}, {"lng", coord.Longitude}}
		gridDoc := bson.D{{"id", coord.Id}, {"coordinate", gridCoord}}
		err = insertDB(context.TODO(), client, gridDoc, "grid")
		fmt.Printf("\nERROR-->\n", err)

	}
	fmt.Print("done")

	fmt.Fprint(w, "stored")

}

func getDateFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY:", string(body))

	// connect to client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	var queryDate QueryDate
	err = json.Unmarshal(body, &queryDate)
	fmt.Printf("UNMARSHAL--->%v", queryDate)

	ctx := context.TODO()
	usersCollection := client.Database("fyp_test").Collection("flights")
	filter := bson.D{{"date", bson.D{{"$eq", queryDate.Date}}}}

	result, err := usersCollection.Find(ctx, filter)
	var results []bson.M
	if err = result.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Printf("\n'%v' matching docs found\n", len(results))

	for _, d := range results {
		fmt.Printf("MATCHING DOC: %v\n", d)
	}

	//NEED TO SEND THE HOUR TIMES IN THESE DOCS BACK TO FRONTEND TO DISPLAY

	var times []interface{} // for the dates
	var time []interface{}  //for the actual hours and minutes
	for _, doc := range results {
		times = append(times, doc["date"])
		time = append(time, doc["time"])
	}

	//convert times (type interface) to type string
	var timesStr []string
	for _, v := range times {
		valStr := fmt.Sprint(v)
		timesStr = append(timesStr, valStr)
		// fmt.Fprintf(w, valStr+",")
	}

	var timeStr []string
	for _, x := range time {
		valStr := fmt.Sprint(x)
		timeStr = append(timeStr, valStr)
		// fmt.Fprintf(w, valStr+",")
	}

	var pairs []string
	for i, _ := range timeStr {
		timeval := timeStr[i]
		timesval := timesStr[i]
		timeval = timeval + " " + timesval
		pairs = append(pairs, timeval)
		fmt.Fprintf(w, timeval+",")
	}
}

func storeFlight(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// connect to client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// unmarshal the json into a Flight struct
	var flight Flight
	err = json.Unmarshal(body, &flight)
	fmt.Printf("UNMARSHAL--->%v", flight)

	ftime := flight.Hour + ":" + flight.Minute
	startCoord := bson.D{{"lat", flight.Srclat}, {"lng", flight.Srclng}}
	destCoord := bson.D{{"lat", flight.Destlat}, {"lng", flight.Destlng}}
	flightDoc := bson.D{{"date", flight.Date}, {"time", ftime}, {"startCoord", startCoord}, {"destCoord", destCoord}, {"speed", flight.Speed}, {"corridor", flight.Corridor}, {"altitude", flight.Altitude}, {"orientation", flight.Orientation}, {"drone", flight.Drone.Name}}
	err = insertDB(context.TODO(), client, flightDoc, "flights")

	droneDoc := bson.D{{"name", flight.Drone.Name}, {"model", flight.Drone.Model}, {"weight", flight.Drone.Weight}}
	err = insertDB(context.TODO(), client, droneDoc, "drones")

	fmt.Printf("\nERROR-->\n", err)

	fmt.Fprint(w, "stored")
}

// func storeDrone() {

// 	droneDoc := bson.D{{"name", flight.Drone.Name}, {"model", flight.Drone.Model}, {"weight", flight.Drone.Weight}}
// 	err = insertDB(context.TODO(), client, droneDoc, "drones")
// }

func getAllTimes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	usersCollection := client.Database("fyp_test").Collection("flights")
	filter := bson.M{}

	result, err := usersCollection.Find(ctx, filter)
	var results []bson.M
	if err = result.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var times []interface{} // for the dates
	var time []interface{}  //for the actual hours and minutes
	for _, doc := range results {
		times = append(times, doc["date"])
		time = append(time, doc["time"])
	}

	//convert times (type interface) to type string
	var timesStr []string
	for _, v := range times {
		valStr := fmt.Sprint(v)
		timesStr = append(timesStr, valStr)
		// fmt.Fprintf(w, valStr+",")
	}

	var timeStr []string
	for _, x := range time {
		valStr := fmt.Sprint(x)
		timeStr = append(timeStr, valStr)
		// fmt.Fprintf(w, valStr+",")
	}

	var pairs []string
	for i, _ := range timeStr {
		timeval := timeStr[i]
		timesval := timesStr[i]
		timeval = timeval + " " + timesval
		pairs = append(pairs, timeval)
		fmt.Fprintf(w, timeval+",")
	}

	fmt.Printf("RESULTS:----->%v", results)
	fmt.Print("RECEIVED REQUEST FROM FRONTEND")
	// fmt.Fprintf(w, "Hola")

}
