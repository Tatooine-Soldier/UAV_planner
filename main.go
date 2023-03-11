package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"crypto/sha256"

	"github.com/EpicStep/go-simple-geo/geo"
	"github.com/google/uuid"
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

type LoginSuccessObj struct {
	Name   string `json:"name"`
	Result bool   `json:"result"`
	Id     string `json:"id"`
}

type Flight struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	Hour        string `json:"hour"`
	Minute      string `json:"minute"`
	Srclat      string `json:"srclat"`
	Srclng      string `json:"srclng"`
	Destlat     string `json:"destlat"`
	Destlng     string `json:"destlng"`
	Speed       string `json:"speed"`
	EndTime     string `json:"endtime"`
	Altitude    string `json:"altitude"`
	Orientation string `json:"orientation"`
	Corridor    string `json:"corridor"`
	Drone       Drone  `json:"drone"`
	Subgrid     string `json:"subgrid"`
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

type SegmentedFlightData struct {
	SegmentedTimes []interface{} `json:"segmentTimes"`
	SegmentedList  []interface{} `json:"segmentList"`
	Date           string        `json:"date"`
	SubGrid        string        `json:"subGrid"`
	ID             string        `json:"id"`
}

type TimeComparisonObj struct {
	SegmentedTimes  []string
	SegmentedCoords []Coordinate
}

type TimeUpdate struct {
	Id     string `json:"id"`
	Date   string `json:"date"`
	Hour   string `json:"hour"`
	Minute string `json:minute"`
}

type GridofCoordinates struct {
	Coordinates []interface{} `json:"coordinates"`
	Layers      []string      `json:"layers"` //layers are confined to only 3 due to height restrictions imposed by the IAA at 120metres maximum
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

type TimeRecord struct {
	Id     string `json:"id"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

type FlightSegmented struct {
	Id          string
	Date        string
	Coordinates []Coordinate
	Times       []string
	SubGrid     string
}

type QueryDate struct {
	Date  string `json:"date"`
	ID    string `json:"id"`
	Reset bool   `json:"reset"`
}

type ResponseData struct {
	Message string `json:"message"`
}

type ResponseDataDate struct {
	Message []string `json:"message"`
}

type Userobj struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserSignobj struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var GRID_INCREMENT int
var LAYER_ONE = "60"
var LAYER_TWO = "90"
var LAYER_THREE = "120"

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

func getID(ctx context.Context, client *mongo.Client, user Userobj) string {
	//use filter to get all id records where name = user.name
	//need the id in order to set it to the global variable in vue
	usersCollection := client.Database("fyp_test").Collection("users")
	filter := bson.D{{"fullName", user.Name}}
	result, err := usersCollection.Find(ctx, filter)
	var results []bson.M
	if err = result.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	var returnId = ""
	var temp []interface{}
	for _, doc := range results {
		temp = append(temp, doc)
		returnId = fmt.Sprint(doc["id"])
		return returnId
	}

	return returnId

}

func loginRequest(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "GET" {
	// 	panic("GET method not permitted")
	// } else {
	// 	r.ParseForm()
	// 	fmt.Printf("BEFORE HASH AND STORAGE--> %v %v", r.Form["username"], r.Form["password"])
	// }
	// username := r.Form["username"]
	// password := r.Form["password"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY:", string(body))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	var user Userobj
	err = json.Unmarshal(body, &user)
	fmt.Printf("User in main.go:", user)

	username := user.Name
	var password []string
	password = append(password, user.Password)

	itemByted := encodeToByte(password)

	hashedVal := sha256.New()
	hashedVal.Write([]byte(itemByted))

	userDoc := bson.D{{"fullName", username}, {"password", hashedVal.Sum(nil)}}
	if userExists := checkDBLogin(context.TODO(), client, userDoc, "users"); userExists {
		fmt.Printf("User exists %v", userExists)
		returnLoginSucces(w, r, user, true)
		return
	} else {
		// http.Redirect(w, r, "/login", http.StatusFound)
		returnLoginSucces(w, r, user, false)
		return
	}
	//http.Redirect(w, r, "/planner", http.StatusSeeOther) //this goes to profile page
}

func returnLoginSucces(w http.ResponseWriter, r *http.Request, user Userobj, success bool) {
	if success {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			panic(err)
		}
		uid := getID(context.TODO(), client, user)

		l := &LoginSuccessObj{Name: user.Name, Result: success, Id: uid}
		b, err := json.Marshal(l)
		if err != nil {
			return
		}
		fmt.Fprintf(w, string(b))
		return
	} else {
		fmt.Fprintf(w, "", success)
	}

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
	// if r.Method == "GET" {
	// 	panic("GET method not permitted")
	// } else {
	// 	r.ParseForm()
	// 	fmt.Printf("BEFORE HASH AND STORAGE--> %v %v", r.Form["username"], r.Form["password"])
	// }
	// username := r.Form["username"]
	// password := r.Form["password"]
	// email := r.Form["email"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY:", string(body))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	var user UserSignobj
	err = json.Unmarshal(body, &user)
	fmt.Printf("User in main.go:", user)

	username := user.Name
	email := user.Email
	var password []string
	password = append(password, user.Password)

	itemByted := encodeToByte(password)

	hashedVal := sha256.New()
	hashedVal.Write([]byte(itemByted))

	uid := uuid.New()
	id := fmt.Sprintf("%v", uid)
	// idInt := rand.Intn(1000)
	// id := strconv.Itoa(idInt)

	userDoc := bson.D{{"fullName", username}, {"email", email}, {"password", hashedVal.Sum(nil)}, {"id", id}}

	alreadySignedUp := checkDBSignup(context.TODO(), client, username, "users")
	if alreadySignedUp {
		fmt.Println("User already exists")
		returnSignupSuccess(w, r, user, false)
		return
	} else {
		fmt.Printf("\nINSERTING THIS HASH-->%v\n ", hashedVal.Sum(nil))
		insertDB(context.TODO(), client, userDoc, "users")
		returnSignupSuccess(w, r, user, true)
	}

	// itemByted := encodeToByte(password)

	// hashedVal := sha256.New()
	// hashedVal.Write([]byte(itemByted))

	// user := bson.D{{"fullName", username}, {"email", email}, {"password", hashedVal.Sum(nil)}}
	// fmt.Printf("\nINSERTING THIS HASH-->%v\n ", hashedVal.Sum(nil))
	// insertDB(context.TODO(), client, user, "users")
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

func returnSignupSuccess(w http.ResponseWriter, r *http.Request, user UserSignobj, success bool) {
	if success {
		l := &LoginSuccessObj{Name: user.Name, Result: success}
		b, err := json.Marshal(l)
		if err != nil {
			return
		}
		fmt.Fprintf(w, string(b))
		return
	} else {
		fmt.Fprintf(w, "", success)
	}
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
	http.HandleFunc("/getUsername", getUsername)
	http.HandleFunc("/storeGridCoordinates", storeGridCoordinates)
	http.HandleFunc("/fetchGridCoordinates", fetchGridCoordinates)
	http.HandleFunc("/storeSegmentedFlight", storeSegmentedFlight)
	http.HandleFunc("/getFlightsWithinRadius", getFlightsWithinRadius)
	http.HandleFunc("/updateFlightTime", updateFlightTime)

	// dist := calculateDistance(3.44, 3.44)
	listenerErr := http.ListenAndServe(":3333", nil)
	fmt.Printf("%v", listenerErr)
}

func handle(w http.ResponseWriter, r *http.Request, name string) {
	fs := http.FileServer(http.Dir("../../../dist"))
	url := fmt.Sprintf("//%v", name)
	http.Handle(url, fs)
}

func storeSegmentedFlight(w http.ResponseWriter, r *http.Request) {
	//should be receiving an array of coordinates in body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY:", string(body))

	var d SegmentedFlightData
	err = json.Unmarshal(body, &d)
	fmt.Printf("UNMARSHAL--->%v", d)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	//###########################################################
	//if a flight with same id already exists then delete and replace with this current flight
	// collection := client.Database("fyp_test").Collection("segmentedFlight")
	// filter := bson.M{"id": d.ID}

	// docExists, err := documentExists(collection, filter)
	// if docExists {
	// 	_, err := collection.DeleteOne(context.Background(), filter)
	// 	if err != nil {
	// 		fmt.Printf("err %v", err)
	// 		return
	// 	}
	// 	fmt.Printf("deleted %v", d.ID)
	// }

	//############################################################
	slist := []Coordinate{}
	for _, val := range d.SegmentedList {
		//fmt.Printf("val---> %v", val)x
		t := val.(map[string]interface{})
		var lat = t["lat"].(string)
		var lng = t["lng"].(string)
		fmt.Printf("\ni--> %v\tval---> %v \n", lat, lng)

		var coord Coordinate
		coord.Latitude = lat
		coord.Longitude = lng
		coord.Id = d.ID

		slist = append(slist, coord)
	}
	slist = reverseSegments(slist)

	timesList := []TimeRecord{}
	for _, val := range d.SegmentedTimes {
		t := val.(map[string]interface{})
		var hour = t["hour"].(string)
		var minute = t["minute"].(string)
		fmt.Printf("\ni--> %v\tval---> %v \n", hour, minute)
		//TODO:
		var timeRecord TimeRecord
		timeRecord.Id = d.ID
		timeRecord.Hour = hour
		timeRecord.Minute = minute

		timesList = append(timesList, timeRecord)
	}
	fmt.Printf("TIMESLIST--->%v", timesList)

	gridDoc := bson.D{{"id", d.ID}, {"date", d.Date}, {"subGrid", d.SubGrid}, {"segments", slist}, {"times", timesList}}
	err = insertDB(context.TODO(), client, gridDoc, "segmentedFlight")
	fmt.Printf("\nERROR-->\n", err)

}

func documentExists(collection *mongo.Collection, filter bson.M) (bool, error) {
	var result bson.M
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

//want to update the date of the flight that is stored when the date is selected in the timeslot picker
//need to call the segment function after setting the new time
func updateFlightTime(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY:", string(body))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	var updateTime TimeUpdate
	err = json.Unmarshal(body, &updateTime)
	fmt.Printf("UNMARSHAL UPDATE--->%v", updateTime)

	ctx := context.TODO()
	usersCollection := client.Database("fyp_test").Collection("flights")
	filter := bson.D{{"id", bson.D{{"$eq", updateTime.Id}}}}

	fullTime := updateTime.Hour + ":" + updateTime.Minute
	update := bson.D{{"$set", bson.D{{"date", updateTime.Date}, {"time", fullTime}}}}

	result, err := usersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
	returnTime := fullTime + "," + updateTime.Date
	fmt.Fprintf(w, returnTime)
	// fmt.Printf("\n'%v' matching docs found\n", len(results))
	// if len(results) == 0 {
	// 	fmt.Fprintf(w, "ALL AVAILABLE")
	// 	return
	// }

}

func getFlightsWithinRadius(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY:", string(body))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	var queryDate QueryDate
	err = json.Unmarshal(body, &queryDate)
	fmt.Printf("UNMARSHAL--->%v", queryDate)
	//####################################################
	//reset the segments list after time has been selected
	if queryDate.Reset {
		//var myA primitive.A = primitive.A{}

		collection := client.Database("fyp_test").Collection("segmentedFlight")
		filter := bson.M{"id": queryDate.ID}
		result, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted %v documents.\n", result.DeletedCount)
	}

	//Query data and id are passed into the function, query date is used to find all other flights(segmented) that occur on this date
	//#############################################################################################################################
	//THIS SECTION FOCUSES ON SEGMENTED FLIGHTS ALREADY IN THE DATABASE FOR THIS DATE WHICH I WILL BE COMPARING AGAINST
	ctx := context.TODO()
	usersCollection := client.Database("fyp_test").Collection("segmentedFlight")
	filterForAllFlights := bson.D{{"date", bson.D{{"$eq", queryDate.Date}}}}

	result, err := usersCollection.Find(ctx, filterForAllFlights)
	var results []bson.M
	if err = result.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Printf("\n'%v' matching docs found\n", len(results))
	if len(results) == 0 {
		fmt.Fprintf(w, "ALL AVAILABLE")
		return
	}

	var segs interface{} // for the dates
	var times interface{}
	var reservedFlightsOnThisDate []FlightSegmented
	// var totalList []FlightSegmented
	//var coordStringList []string
	//mapCoordinates := make(map[string][]Coordinate)

	for _, doc := range results {
		var timesStringList []string
		var coordStringList []Coordinate
		fmt.Printf("\ndoc:===:=== %v", doc)
		times = doc["times"]
		for _, x := range times.(primitive.A) {
			f := x.(primitive.M)
			//fmt.Printf("\n%v %v %v\n", f["hour"], f["id"], f["minute"])
			var timeStr = f["hour"].(string) + ":" + f["minute"].(string)
			timesStringList = append(timesStringList, timeStr)
		}
		segs = doc["segments"]
		for _, y := range segs.(primitive.A) {
			f := y.(primitive.M)
			//fmt.Printf("\n%v %v %v\n", f["id"], f["latitude"], f["longitude"])
			var c Coordinate
			c.Id = f["id"].(string)
			c.Latitude = f["latitude"].(string)
			c.Longitude = f["longitude"].(string)
			coordStringList = append(coordStringList, c)
			// var timeStr = f["hour"].(string) + ":" + f["minute"].(string)
			// timesStringList = append(timesStringList, timeStr)
		}
		var c FlightSegmented
		c.Id = doc["id"].(string)
		c.Date = doc["date"].(string)
		c.SubGrid = doc["subGrid"].(string)
		c.Coordinates = coordStringList
		c.Times = timesStringList
		reservedFlightsOnThisDate = append(reservedFlightsOnThisDate, c)
		// fmt.Printf("\ntimes per flight-->%v\n", timesStringList)
		// fmt.Printf("\ncoords per flight-->%v\n", coordStringList)
		//loop for same index length(both lists have same length) and create an object eg. key is time and value is coordinate. Then stick these in a new list
	}

	for _, n := range reservedFlightsOnThisDate { //list of FlightSegmented structs{id, []Coordinate, []string}
		fmt.Printf("\n--> flights on this date: %v", n)
	}
	// i loop through reserved flights, check if the starting time(1st eleemtn in .times is in the intended time-->if not then move on to next flight else check coordinates)

	//################################################################################################################################
	//THIS SECTION BELOW FOCUSES ON THE INTENDED FLIGHT DATA, GETS THE INTENDED FLIGHT
	filterIntendedFlight := bson.D{{"id", bson.D{{"$eq", queryDate.ID}}}}
	resultIntendedFlight, err := usersCollection.Find(ctx, filterIntendedFlight)
	var resultsIntendedFlight []bson.M
	if err = resultIntendedFlight.All(context.TODO(), &resultsIntendedFlight); err != nil {
		panic(err)
	}

	var intendedTimesList []string
	var intendedCoordsList []Coordinate
	var intendedFlight FlightSegmented
	for _, d := range resultsIntendedFlight {
		//fmt.Printf("MATCHING DOC: %v\n", d)
		t := d["times"]
		for _, a := range (t).(primitive.A) {
			f := a.(primitive.M)
			var timeStr = f["hour"].(string) + ":" + f["minute"].(string)
			intendedTimesList = append(intendedTimesList, timeStr)
		}
		s := d["segments"]
		for _, b := range (s).(primitive.A) {
			f := b.(primitive.M)
			var c Coordinate
			c.Id = f["id"].(string)
			c.Latitude = f["latitude"].(string)
			c.Longitude = f["longitude"].(string)
			intendedCoordsList = append(intendedCoordsList, c)
		}

		intendedFlight.Id = d["id"].(string)
		intendedFlight.Date = d["date"].(string)
		intendedFlight.SubGrid = d["subGrid"].(string)
		intendedFlight.Coordinates = intendedCoordsList
		intendedFlight.Times = intendedTimesList
		fmt.Printf("\nIntended Object-->%v", intendedFlight)

	}

	//##########################################################################################################################
	//For each segmented coordinate of intended flight, check all other segmented flights coordinates to see if one is within 120m
	//If a reserved flight has a coordinate within 120m of a intended flight coordinate, check what time both those coordinates are within that distance at
	var flightWatchList []FlightSegmented
	for _, f := range reservedFlightsOnThisDate {
		if intendedFlight.Id != f.Id {
			if checkSubGridLevel(intendedFlight.SubGrid, f.SubGrid) {
				//check these flights
				flightWatchList = append(flightWatchList, f) //flights occuring in the same sub grid on the same date
			}
		}
	}
	var unavailableTimes = schedule(intendedFlight, flightWatchList)
	fmt.Printf("\nOriginal time check %v", unavailableTimes)
	if len(unavailableTimes) > 0 { //if there is a collision at this time
		//add 5 minutes onto each of these times and then rerun the schedule function
		var fiveMinuteWaitSegments []string
		for _, segTime := range intendedFlight.Times {
			if len(segTime) < 5 {
				segTime = segTime[0:3] + "0" + segTime[3:4]
			}
			timePlusFive, err := time.Parse("15:04", segTime)
			if err != nil {
				fmt.Printf("Error parsing %v as a time: %v", timePlusFive, err)
				return
			}
			timePlusFive = timePlusFive.Add(5 * time.Minute)
			timeStr := timePlusFive.String()
			timeStr = timeStr[10:16]
			fiveMinuteWaitSegments = append(fiveMinuteWaitSegments, timeStr)
		}
		fmt.Printf("New 5 min timestamps: %v", fiveMinuteWaitSegments)
		intendedFlight.Times = fiveMinuteWaitSegments
		unavailableTimes = schedule(intendedFlight, flightWatchList)

		//this code below is reached if there is a coord&time collision in the origianl grid and if there is no colllision after waiting 5 minutes in orginial hgrid
		if len(unavailableTimes) > 0 { //if there is still a delay after waiting 5 mins, check if flight can be allocated to another sub grid
			availableGrids, gridIsEmpty := checkOtherSubGridAvailability(intendedFlight.SubGrid)
			if gridIsEmpty {
				fmt.Printf("Empty grids to switch to %v", availableGrids)
			}
		}
	}

	// var unavailableTimes []string
	// for i := 0; i < len(intendedFlight.Coordinates); i++ { //for segmented coord in intended flight path
	// 	var collisionsOnThisDate string
	// 	for _, val := range reservedFlightsOnThisDate { //for reserved flight on this date
	// 		if intendedFlight.Id != val.Id {
	// 			fmt.Println("\n\n#################################################below####\n")
	// 			if checkSubGridLevel(intendedFlight.Id, val.Id) { //check that both flights are on the same grid level
	// 				g := intendedFlight.Coordinates[i]
	// 				ifCollisions := checkCoordinatesRadius(g, val) //index, times at which coordinate collisions should occur
	// 				if ifCollisions != 0.0 {
	// 					//old way using the intended dlights time only
	// 					collisionTimeOfIntendedFlight := intendedFlight.Times[i]
	// 					collisionTimeOfResrvedFlight := val.Times[ifCollisions]
	// 					if checkTimeCollisions(collisionTimeOfIntendedFlight, collisionTimeOfResrvedFlight, intendedFlight.Date) {
	// 						unavailableTimes = append(unavailableTimes, collisionTimeOfIntendedFlight)
	// 						collisionsOnThisDate += collisionTimeOfIntendedFlight
	// 						fmt.Fprintf(w, collisionTimeOfIntendedFlight+",")
	// 						fmt.Printf("POSSIBLE COLLISION AT THIS TIME IN THIS GRID %v", intendedFlight.Times[ifCollisions])
	// 						//check when this coordinate will next be free, predict new starting time to accomodate this, append that time to a list
	// 						//check other subgrid for a start time that occurs before the start times contained in the list above
	// 						//add to queue to enter new subgrid
	// 					} else {
	// 						fmt.Printf("Coordinate collision but no time collison")
	// 					}
	// 				} else {
	// 					fmt.Printf("\nNO Coordinate COLLISIONS PREDICTED")
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fmt.Printf("\nAny Collisions for (%v %v) --> %v\n", intendedFlight.Coordinates[i].Latitude, intendedFlight.Coordinates[i].Longitude, collisionsOnThisDate)
	// }
	//fmt.Printf("Collision times--> %v", unavailableTimes)

}

func checkOtherSubGridAvailability(gridLevel string) ([]string, bool) {
	fmt.Println("\n\n*Checking other grids: *")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database("fyp_test").Collection("segmentedFlight")

	layers := [3]string{LAYER_ONE, LAYER_TWO, LAYER_THREE}
	var availableGrids []string
	isEmpty := false //returns true if any of the other grids is empty
	for _, l := range layers {
		fmt.Printf("Checking layer '%v'", l)
		if l != gridLevel { //only check grids that haven't been checked already
			filter := bson.M{"subGrid": l} //NO: need to get grids that are != gridLevel ie. the grids that are empty
			cursor, err := collection.Find(context.Background(), filter)
			if err != nil {
				fmt.Println(err)
			}
			defer cursor.Close(context.Background())

			if cursor.Next(context.Background()) {
				// cursor is not empty
				fmt.Println("CURSOR CHECK-->cursor not empty")
			} else {
				// cursor is empty
				fmt.Println("CURSOR CHECK-->cursor is empty")
			}

			//seems like this isn't really working, says a record exists for each of these layers
			if err = cursor.Err(); err != nil {
				if err == mongo.ErrNoDocuments { //if grid is empty
					fmt.Printf("\nSubGrid empty %v\n", l)
					availableGrids = append(availableGrids, l)
					isEmpty = true
				}
				fmt.Println(err)
			} else {
				fmt.Println("Sub grid not empty")
			}
		}
	}
	return availableGrids, isEmpty
}

func schedule(intendedFlight FlightSegmented, flightWatchList []FlightSegmented) []string {
	var unavailableTimes []string
	for i := 0; i < len(intendedFlight.Coordinates); i++ { //for segmented coord in intended flight path
		for _, val := range flightWatchList {
			g := intendedFlight.Coordinates[i]
			ifCollisions := checkCoordinatesRadius(g, val) //index, times at which coordinate collisions will occur
			if ifCollisions != 0.0 {
				collisionTimeOfIntendedFlight := intendedFlight.Times[i]
				collisionTimeOfResrvedFlight := val.Times[ifCollisions]
				if checkTimeCollisions(collisionTimeOfIntendedFlight, collisionTimeOfResrvedFlight, intendedFlight.Date) {
					unavailableTimes = append(unavailableTimes, collisionTimeOfIntendedFlight)
					// fmt.Fprintf(w, collisionTimeOfIntendedFlight+",")
					fmt.Printf("POSSIBLE COLLISION AT THIS TIME IN THIS GRID %v", intendedFlight.Times[ifCollisions])
					//check when this coordinate will next be free, predict new starting time to accomodate this, append that time to a list
					//check other subgrid for a start time that occurs before the start times contained in the list above
					//add to queue to enter new subgrid
				} else {
					fmt.Printf("Coordinate collision but no time collison")
				}
			} else {
				fmt.Printf("\nNO Coordinate COLLISIONS PREDICTED")
			}
		}
		fmt.Printf("\nAny Collisions for (%v %v) --> \n", intendedFlight.Coordinates[i].Latitude, intendedFlight.Coordinates[i].Longitude)
		fmt.Printf("\nCollision times--> %v", unavailableTimes)
	}
	return unavailableTimes
}

//returns true if one of the reserved coordinates is within 120m of one of the indended coordinate
func checkCoordinatesRadius(intended Coordinate, reserved FlightSegmented) int {
	intendedLat, err := strconv.ParseFloat(intended.Latitude, 64)
	if err != nil {
		fmt.Println("can't convert")
	}
	intendedLng, err := strconv.ParseFloat(intended.Longitude, 64)
	if err != nil {
		fmt.Println("can't convert")
	}
	for j, flightCoord := range reserved.Coordinates {
		reservedLat, err := strconv.ParseFloat(flightCoord.Latitude, 64)
		if err != nil {
			fmt.Println("can't convert")
		}
		reservedLng, err := strconv.ParseFloat(flightCoord.Longitude, 64)
		if err != nil {
			fmt.Println("can't convert")
		}
		if calculateCoordDistance(intendedLat, intendedLng, reservedLat, reservedLng) < .120 { //if two coordinates are within 120 metres of eachother
			fmt.Printf("\n\nReturned True for these coordinates, flight IDs (%v %v) \tcoords: (%v %v)\t(%v %v) \n", intended.Id, reserved.Id, intendedLat, intendedLng, reservedLat, reservedLng)
			return j //return time that the collision occured at

		}
	}
	return 0.0
}

// *THIS FUNCTION IS FREELY AVAILABLE AND HAS BEEN TAKEN FROM https://www.geodatasource.com/developers/go	*
func calculateCoordDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344 //Convert to km

	//fmt.Printf("\nDist: %v\n", dist)
	return dist
}

//need to convert the string times into actual time objects and see if theres a collision 5 minutes on either side(before and after)
func checkTimeCollisions(intendedTime string, reservedTime string, date string) bool {
	//need to add "0" before time with only one minute digit
	if len(intendedTime) < 5 {
		intendedTime = intendedTime[0:3] + "0" + intendedTime[3:4]
	}
	if len(reservedTime) < 5 {
		reservedTime = reservedTime[0:3] + "0" + reservedTime[3:4]
	}

	fullIntended := date + " " + intendedTime + ":00"
	layout := "2006-01-02 15:04:05" // Example layout string
	t, err := time.Parse(layout, fullIntended)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	intendedEpochTime := t.Unix()
	// fiveMinutesBefore := (intendedEpochTime - 5*60*1000)
	// fiveMinutesAfter := (intendedEpochTime + 5*60*1000)

	fullReserved := date + " " + reservedTime + ":00"
	f, err := time.Parse(layout, fullReserved)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	reservedEpochTime := f.Unix()

	fmt.Printf("Checking %v %v", fullIntended, fullReserved)

	//check if UAV passes through this coordinate five five minutes befor or after the intended flight, this ensures no collision if there is an unexpected delay
	if math.Abs(float64(reservedEpochTime-intendedEpochTime)) <= 300 {
		fmt.Printf("Gap ahead of 5 minutes %v %v %v\n", reservedTime, intendedTime, math.Abs(float64(reservedEpochTime-intendedEpochTime)))
		return true
	}
	if math.Abs(float64(intendedEpochTime-reservedEpochTime)) <= 300 {
		fmt.Printf("Gap behind of 5 minutes %v %v %v\n", reservedTime, intendedTime, math.Abs(float64(intendedEpochTime-reservedEpochTime)))
		return true
	}

	if intendedTime == reservedTime { //need to expand this 5 minutes either side(before and after)
		return true
	}

	fmt.Printf("\nNO TIME COLLISION --> %v\n", math.Abs(float64(reservedEpochTime-intendedEpochTime)))
	return false

	// reservedTimes := reserved.Times
	// for _, time := range reservedTimes {
	// 	if time == intendedTime {
	// 		fmt.Printf("Collision at this time: %v %v", time, intendedTime)
	// 		return true
	// 	}
	// }
	// return false
}

func checkSubGridLevel(intendedSubGrid string, reservedSubGrid string) bool {
	if intendedSubGrid == reservedSubGrid {
		return true
	}
	return false
}

// func checkSubGridLevel(intendedID string, reservedID string) bool {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	collection := client.Database("fyp_test").Collection("flights")
// 	filter := bson.M{"id": intendedID}
// 	var flight bson.M
// 	err = collection.FindOne(context.Background(), filter).Decode(&flight)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	intendedSubGrid := flight["subGridLayer"].(string)

// 	filter = bson.M{"id": reservedID}
// 	err = collection.FindOne(context.Background(), filter).Decode(&flight)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	reservedSubGrid := flight["subGridLayer"].(string)

// 	fmt.Printf("\nComparing grid levels %v %v", intendedSubGrid, reservedSubGrid)
// 	if intendedSubGrid == reservedSubGrid {
// 		fmt.Println("same sub grid level")
// 		return true
// 	}
// 	fmt.Println("disfferent sub grids")
// 	return false
// }

func dateTimeCheck(hour string, minute string) {
	if len(minute) < 2 {
		minute = "0" + minute
	}

}

func reverseSegments(s []Coordinate) []Coordinate {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func getUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		panic("GET method not permitted")
	} else {
		r.ParseForm()
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("BODY in getUsername:", string(body))

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	panic(err)
	// }
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

	//fmt.Printf("%v results", results)

	// var coords []string
	// for _, doc := range results {
	// 	stringDoc := fmt.Sprint(doc) + ","
	// 	fmt.Printf(stringDoc)
	// 	coords = append(coords, fmt.Sprint(stringDoc))
	// 	fmt.Fprintf(w, stringDoc)
	// }

	//need this to be json
	var temp []interface{}
	for _, doc := range results {
		temp = append(temp, doc)
		//fmt.Printf("doc:%v \nc.Id:%v\n\n", doc, c.Id)
	}
	c := &GridofCoordinates{
		Coordinates: temp,
	}

	j, err := json.Marshal(c.Coordinates)
	if err != nil {
		fmt.Println(err)
		return
	}

	jstr := string(j)
	fmt.Printf("%v<---jstr", jstr)

	fmt.Printf("\n'%v' matching docs found\n", len(results))
	fmt.Fprintf(w, jstr)

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

	filter := bson.D{}
	usersCollection := client.Database("fyp_test").Collection("grid")
	usersCollection.DeleteMany(context.TODO(), filter)

	var grid GridofCoordinates
	err = json.Unmarshal(body, &grid)
	fmt.Printf("COORDSLIST: %vFINISHED", grid.Coordinates)

	GRID_INCREMENT++
	var gridID = fmt.Sprintf("%d", GRID_INCREMENT)

	for _, layer := range grid.Layers {
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
			gridDoc := bson.D{{"id", coord.Id}, {"coordinate", gridCoord}, {"layer", layer}, {"gridID", gridID}}
			err = insertDB(context.TODO(), client, gridDoc, "grid")
			fmt.Printf("\nERROR-->\n", err)

		}
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

	parseEndTime(flight)

	ftime := flight.Hour + ":" + flight.Minute
	startCoord := bson.D{{"lat", flight.Srclat}, {"lng", flight.Srclng}}
	destCoord := bson.D{{"lat", flight.Destlat}, {"lng", flight.Destlng}}
	flightDoc := bson.D{{"id", flight.ID}, {"date", flight.Date}, {"time", ftime}, {"startCoord", startCoord}, {"destCoord", destCoord}, {"endTime", flight.EndTime}, {"speed", flight.Speed}, {"corridor", flight.Corridor}, {"subGridLayer", flight.Subgrid}, {"altitude", flight.Altitude}, {"orientation", flight.Orientation}, {"drone", flight.Drone.Name}}
	err = insertDB(context.TODO(), client, flightDoc, "flights")

	droneDoc := bson.D{{"name", flight.Drone.Name}, {"model", flight.Drone.Model}, {"weight", flight.Drone.Weight}}
	err = insertDB(context.TODO(), client, droneDoc, "drones")

	fmt.Printf("\nERROR-->\n", err)

	fmt.Fprint(w, "stored")
}

func storeTimestampedFlight(f Flight) {
	hourToMins, err := strconv.Atoi(f.Hour)
	rawMins, err := strconv.Atoi(f.Minute)
	speed, err := strconv.Atoi(f.Speed)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	mins := hourToMins * 60
	mins = mins + rawMins

	duration := mins
	lengthOfTimestamps := duration / 50 //arbitrary number, may be changed in future
	var segmented []int
	distanceTravelled := speed * lengthOfTimestamps
	segmented = append(segmented, distanceTravelled)
	// add distance traveeled
	// for i := 0; i < 50; i++ {

	// }
}

func parseEndTime(f Flight) {
	fmt.Println("f", f.EndTime)
	// if len(f.EndTime) == 5 {
	// 	hour := f.EndTime[0:2]
	// 	min := f.EndTime[3:5]
	// 	hourInt, err := strconv.Atoi(hour)
	// 	minInt, err := strconv.Atoi(min)

	// 	hourFint, err := strconv.Atoi(f.Hour)
	// 	minFint, err := strconv.Atoi(f.Minute)
	// 	if err != nil {
	// 		fmt.Println("Error during conversion")
	// 		return
	// 	}
	// 	fmt.Printf("End Min + End Hour: %v %v", minInt, hourInt)
	// 	fmt.Printf("\nStart Min + Start Hour: %v %v", minFint, hourFint)

	// 	var totalHours = hourInt + hourFint //if over 24 --> next day
	// 	var totalMinutes = minInt + minFint //if over 60

	// 	if totalMinutes >= 60 {
	// 		totalHours += 1
	// 		m := totalMinutes % 60
	// 		totalMinutes = m
	// 	}

	// 	if totalHours >= 24 {
	// 		fmt.Printf("next day")
	// 		//f.Date += 1

	// 		d := f.Date[8:10]
	// 		date, err := strconv.Atoi(d)
	// 		if err != nil {
	// 			fmt.Println("Error during conversion")
	// 			return
	// 		}
	// 		date += 1
	// 		dateString := strconv.Itoa(date) + "00"
	// 		strings.ReplaceAll(f.Date, d, dateString)

	// 		t := totalHours % 24
	// 		totalHours = t
	// 	}

	//date := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	//fmt.Println(date.Unix())

	//fmt.Printf("%v, %v", minInt, hourInt)
	//}

	// if len(e) == 2 {
	// 	min := e[0:2]
	// 	fmt.Printf("%v", min)
	// }

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
