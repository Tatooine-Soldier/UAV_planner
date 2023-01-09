<script setup>
  import FinalGoogleMapComponent from "@/components/FinalGoogleMapComponent.vue";
import MyGoogleMapComponent from "../components/MyGoogleMapComponent.vue"
//import FinalGoogleMapComponentVue from "@/components/FinalGoogleMapComponent.vue";
//   const props = defineProps(['title'])
</script>

<template>
    <section class="flight-planner" id="flight-planner">
        <!-- <form action="/planner" method="post"> -->
        <form @submit.prevent="handleSubmit()">
            <section class="flight-details-container" id="flight-details-container">
                <section class="flight-details" id="flight-details">
                    <section class="flight-details-content">
                        <h1>Confirm Flight Details</h1>
                        <section class="flight-details-content-grid">
                            <div>Date: {{ date.day }}</div>
                            <div>Time: {{ date.hour }}:{{ date.minute }}</div>
                            <div>Speed: {{ speed.description }} </div>
                            <div>Calculated distance: {{ distance }} KM</div>
                        </section>
                        <section class="flight-details-buttons">
                            <input id="cancel-but" name="but" type="button" value="Cancel" v-on:click="disappear()"/>
                            <button id="confirm-but" v-on:click="showFinalMap()">Confirm</button>
                        </section>
                    </section>
                </section>
            </section>
            <section id="details-map-container">
                <MyGoogleMapComponent @someEvent="logme"></MyGoogleMapComponent>
            </section>
            <section id="final-map-container">
                <FinalGoogleMapComponent :propcoords="coords" :propspeed="speed"></FinalGoogleMapComponent>
            </section>
            <section class="waypoint-submit-container" id="waypoint-container">
                Enter desired waypoints:
                <section>
                    <input type="text" name="sourcelatitude" size="16"/>
                </section>
            </section>
            <div><img src="../assets/ex-sign.png" id="ex-sign" v-on:click="disappearEx()"/></div>
        <section class="flight-planner-columns" id="flight-planner-columns">
            <section class="fp-info-container">
                        <h1 id="fp-headers">When</h1>
                        <section class="fp-sub-info">
                            <section class="border-decor">
                                <p class="fp-subtitle">Select your time to fly:</p>
                                <!-- <form @submit.prevent="handleSubmit()"> -->
                                    <section class="time-selection">
                                        <div>
                                            <label for="hour">Select hour: </label>
                                            <select id="hour" name="hour" v-model="date.hour">
                                                <option value="00">00</option>
                                                <option value="01">01</option>
                                                <option value="02">02</option>
                                                <option value="03">03</option>
                                                <option value="04">04</option>
                                                <option value="05">05</option>
                                                <option value="06">06</option>
                                                <option value="07">07</option>
                                                <option value="08">08</option>
                                                <option value="09">09</option>
                                                <option value="10">10</option>
                                                <option value="11">11</option>
                                                <option value="12">12</option>
                                            </select>
                                        </div>
                                        <div>
                                            <label for="minute">Select minute: </label>
                                            <select id="minute" name="minute" v-model="date.minute">
                                                <option value="00">00</option>
                                                <option value="15">15</option>
                                                <option value="30">30</option>
                                                <option value="45">45</option>
                                            </select>
                                        </div>
                                        <div id="last-div">
                                            <label for="day">Select day: </label>
                                            <input type="date" name="date" v-model="date.day"/>
                                        </div>
                                    </section>
                                    <section>
                                        <!-- <button>Add</button> -->
                                    </section>
                                <!-- </form> -->
                            </section>
                        </section>
                    </section>
                    <section class="fp-info-container">
                        <h1 id="fp-headers">Where</h1>
                        <section class="fp-sub-info">
                            <p class="fp-subtitle">Select your source and destination:</p>
                            <section class="coords-container">
                                
                                <!-- <form action="/location" method="post"> -->
                                <section class="coords">
                                    <section class="coords-source">
                                        <div class="coords-labels"><b>Source:</b></div>
                                        <section class="coords-inputs">
                                            <label for="latitude">Latitude: </label>
                                            <input type="text" name="sourcelatitude" size="16" ref="mysourcelat"/>
                                            <label for="longitude">Longitude: </label>
                                            <input type="text" name="sourcelongitude" size="16" ref="mysourcelong"/>
                                        </section>
                                    </section>
                                    <section class="coords-destination">
                                        <div class="coords-labels"><b>Destination:</b></div>
                                        <section class="coords-inputs">
                                            <label for="latitude">Latitude: </label>
                                            <input type="text" name="destlatitude" size="16" ref="mydestlat"/>
                                            <label for="longitude">Longitude: </label>
                                            <input type="text" name="destlongitude" size="16" ref="mydestlong">
                                        </section>
                                    </section>
                                    <section class="waypoints-container">
                                        <div class="waypoints-select">
                                            <label for="waypoints">How many Waypoints to visit?</label>
                                            <div>
                                                <select id="waypoints" name="waypoints" v-model="waypoints">
                                                    <option value="0" v-on:click="setWaypoint()">0</option>
                                                    <option value="1" v-on:click="setWaypoint()">1</option>
                                                    <option value="2" v-on:click="setWaypoint()">2</option>
                                                    <option value="3" v-on:click="setWaypoint()">3</option>
                                                    <option value="4" v-on:click="setWaypoint()">4</option>
                                                    <option value="5" v-on:click="setWaypoint()">5</option>
                                                    <option value="6" v-on:click="setWaypoint()">6</option>
                                                    <option value="7" v-on:click="setWaypoint()">7</option>
                                                    <option value="8" v-on:click="setWaypoint()">8</option>
                                                    <option value="9" v-on:click="setWaypoint()">9</option>
                                                    <option value="10" v-on:click="setWaypoint()">10</option>
                                                </select>
                                            </div>
                                        </div>

                                    </section>
                                </section>
                                <!-- <input id="submit" name="submit" type="submit" value="Add"/> -->
                                <!-- </form> -->
                                <p><b id="or">OR</b></p>
                                <p>Select location on the map below</p>
                                <!-- <section></section> Display the map her-->
                                <input id="map-but" name="but" type="button" value="MAP" v-on:click="showMap()"/>  
                            </section>
                        </section>
                    </section>
                    <section class="fp-info-container">
                        <h1 id="fp-headers">How</h1>
                        <section class="fp-sub-info">
                            <p class="fp-subtitle">Select your desired UAV speed:</p>
                                <!-- <form action="/speed" method="post"> -->
                                <section class="speed-container">
                                    <small><i>* Please check the max speed of your UAV before selecting speed *</i></small>
                                    <p>The speed of the UAV will determine which flight corridor it will be assigned to for the flight</p>
                                    <section class="speeds-option">
                                        <section class="speed-containers">
                                            <div class="speed-selectors">
                                                <label for="latitude">Low speed: </label>
                                                <label for="latitude">Mid speed: </label>
                                                <label for="latitude">High speed: </label>
                                            </div>
                                            <div class="speed-selectors">
                                                <input type="radio" name="description" id="low-radio" value="low-speed corridor"/>
                                                <input type="radio" name="description" id="mid-radio" value="mid-speed corridor"/>
                                                <input type="radio" name="description" id="high-radio" value="high-speed corridor"/>
                                            </div>
                                            <div class="speed-selectors">
                                                <img src="../assets/iicon.png" alt="info" v-on:mouseover="lowInfo()" v-on:mouseout="hideInfo()"/>
                                                <img src="../assets/iicon.png" alt="info" v-on:mouseover="midInfo()" v-on:mouseout="hideInfo()"/>
                                                <img src="../assets/iicon.png" alt="info" v-on:mouseover="highInfo()" v-on:mouseout="hideInfo()"/>
                                            </div>
                                        </section>
                                    </section>
                                    <div class="info-container" id="lowc">
                                        Low speeds are between <b>5km-20kmh</b>
                                    </div>
                                    <div class="info-container" id="midc">
                                        Mid speeds are between <b>20kmh-45kmh</b>
                                    </div>
                                    <div class="info-container" id="highc">
                                        High speeds are between <b>45-80kmh</b>
                                    </div>
                                </section>
                                <!-- <input id="submit" name="submit" type="submit" value="Add"/> -->
                                <!-- </form> -->
                        </section>
                    </section>
                    <section class="fp-info-container">
                        <h1 id="fp-headers">Book</h1>
                        <section class="fp-sub-info">
                                <p class="fp-subtitle">Request flight plan</p>
                                <input id="submit" name="submit" type="submit" value="Add"/>
                                
                        </section>
                    </section>
        </section>
        
    </form>
    </section>
</template>

<style>
    .info-container {
        display: none;
    }

    #lowc {
        position: absolute;
        top: 54%;
        right: 17%;
        width: 10%;
        background-color: grey;
        border: solid 2px rgb(184, 255, 36);
        border-radius: 15px;
        z-index: 1;
        padding: 5px;
        font-size: .8em;
    }

    #or {
        font-size: 18px;
    }

    #midc {
        position: absolute;
        top: 58%;
        right: 17%;
        width: 10%;
        background-color: grey;
        border: solid 2px rgb(184, 255, 36);
        border-radius: 15px;
        z-index: 1;
        padding: 5px;
        font-size: .8em;
    }

    #highc {
        position: absolute;
        top: 62%;
        right: 17%;
        width: 10%;
        background-color: grey;
        border: solid 2px rgb(184, 255, 36);
        border-radius: 15px;
        z-index: 1;
        padding: 5px;
        font-size: .8em;
    }

    .flight-planner-columns {
        display: grid;
        grid-template-columns: 25% 25% 25% 25%;
        align-items: center;
        justify-items: center;
        color: white;
    }

    .fp-info-container {
        width: 100%;
        height: 100%;
        text-align: center;
        border-left: solid 2px white;
        border-right: solid 2px white;
        padding-bottom: 5px;
    }

    .fp-info-container h1 {
        margin: 0px;
    }

    #fp-headers {
        font-size: 32px;
        padding-top: 10px;
    }

    .fp-sub-info {
        padding: 10px;
    }

    .time-selection {
        display:grid;
        grid-template-columns: 50% 50%;
        grid-template-rows: 50% 50%;
    }

    .time-selection div{
        padding-bottom: 10%;
    }

    .fp-subtitle {
        padding: 4%;
    }

    #last-div {
        grid-column-start: 1;
        grid-column-end: 3;
        padding: 5px 0px;
    }

    #cancel-but {
        padding: 5px;
        width: 100%;
    }

    #confirm-but {
        margin-left: 5px;
        padding: 5px;
        width: 100%;
    }

    #map-container {
        display: none;
        z-index: 1;
        position: relative;
    }

    .coords-inputs {
        padding: 12px;
        display: grid;
        grid-template-columns: 38% 62%;
    }

    .coords-inputs label {
        text-align: left;
    }

    .coords {
        display: grid;
        grid-template-rows: 40% 40% 20%;
    }

    .coords-labels {
        padding: 5px;
        text-align: left;
    }

    .waypoints-select {
        display: grid;
        grid-template-rows: 50% 50%;
        row-gap: 4px;
    }

    .speeds-option {
        display: grid;
        grid-template-rows: 33% 33% 33%;
    }

    .speeds-inputs {
        padding: 10px;
    }

    .flight-details-container {
        display: none;
        z-index: 1;
        position: relative;
        width: 100%;
    }

    .flight-details {
        border: solid 2px rgb(184, 255, 36); 
        background-color: rgb(77, 76, 76);
        width: 45%;
        position: absolute;
        padding: 10px;
        margin-top: 2%;
        margin-left: 26.5%;
        
        color: white;
    }

    .flight-details-content {
        padding: 10px;
    }

    .flight-details-content h1 {
        font-size: large;
        margin-top: 0px;
        margin-bottom: 0px;
        text-align: center;
    }

    .flight-details-content-grid {
        display: grid;
        grid-template-rows: 50% 50%;
        grid-template-rows: 25% 25% 25% 25%;
        margin: 0% 5%;
    }

    .flight-details-buttons {
        display: grid;
        grid-template-columns: 50% 50%;
        float: right;
    }

    .flight-details-buttons button:first-of-type{
        margin-right: 5px;
    }

    .coords-source {
        padding: 10px;
    }

    .coords-destination {
        padding: 10px;
    }

    #details-map-container {
        position: absolute;
        display: none;
        z-index: 1;
        left: 10%;
        right: 10%;
        background-color: grey;
        border: solid 1px grey;

    }

    #ex-sign {
        position: absolute;
        right: 10.5%;
        margin-top: .3%;
        z-index: 1;
        display: none;
    }

    #ex-sign:hover {
        cursor: grab;
    }

    #final-map-container {
        position: absolute;
        display: none;
        z-index: 1;
        left: 10%;
        right: 10%;
        background-color: grey;
        border: solid 1px grey;
    }

    .speed-containers {
        display: grid;
        grid-template-columns: 40% 40% 20%;
        padding: 10px;
    }

    .speed-selectors {
        display: grid;
        grid-template-rows: 33% 33% 33%;
    }

    .speed-selectors label {
        padding: 5px;
        font-size: .9em;
    }

    .speed-selectors input {
        margin: 5px;
    }
    
    .waypoint-submit-container {
        display: none;
    }
  

</style>

<!-- <script defer src="<https://maps.googleapis.com/maps/api/js?key=AIzaSyDTNOMjJP2zMMEHcGy2wMNae1JnHkGVvn0&callback=initMap>"> -->
<script>
 // import * as geolib from 'geolib';

export default {
    data() {
      return {
        date: {
          hour: '',
          minute: '',
          day: ''
        },
        coords: {
            sourcelongitude: '',
            sourcelatitude: '',
            destlongitude: '',
            destlatitude: ''
        },
        speed: {
            description: '',  //corridor
            velocity: ''
        },
        info: null,
        distance: null,
        waypoints: null,
      }
    },
    components: {
        MyGoogleMapComponent,
        FinalGoogleMapComponent
    },
    methods: {
      handleSubmit() {
        // Send data to the server or update your stores and such.
        // this.coords.latitude = this.$refs.mylat.value;
        // this.coords.longitude = this.$refs.mylong.value;
        this.coords.sourcelatitude = this.$refs.mysourcelat.value;
        this.coords.sourcelongitude = this.$refs.mysourcelong.value;
        this.coords.destlatitude = this.$refs.mydestlat.value;
        this.coords.destlongitude = this.$refs.mydestlong.value;
        if (document.getElementById('low-radio').checked) {
            this.speed.description = document.getElementById('low-radio').value;
            this.speed.velocity = 20;
        } else if (document.getElementById('mid-radio').checked) {
            this.speed.description = document.getElementById('mid-radio').value;
            this.speed.velocity = 45;
        } else if (document.getElementById('high-radio').checked) {
            this.speed.description = document.getElementById('high-radio').value;
            this.speed.velocity = 65;
        } else {
            this.speed.description = "** Please select speed **"
        }
        console.log(document.getElementById('mid-radio').checked)
        console.log("COOORDS--->",this.coords)
        // var long = parseInt(this.coords.longitude)
        // var lat = parseInt(this.coords.latitude)
        // this.distance = geolib.getDistance({latitude: 0, longitude:0}, {
        //             latitude: lat,
        //             longitude: long,
        //         })
        // this.distance = geolib.convertDistance(this.distance, 'km');
        // console.log(
        //     'DISTANCE-->',
        //      this.distance,
        //     '<--DISTANCE'
        // )
        var details = document.getElementById('flight-details-container');
        details.style.display = 'block';

        
    

        // var grey = document.getElementById('flight-planner-columns');
        // grey.style.opacity = 0.2;
        // grey.style.pointerEvents = "none";

        
      },
      disappear: function (event) {
        var details = document.getElementById('flight-details-container');
        details.style.display = 'None';
        console.log(event)

        var grey = document.getElementById('flight-planner-columns');
        grey.style.opacity = 1;
        grey.style.pointerEvents = "initial";
      },
      handleSubmitCoords() {
        console.log(this.coords)
      }, 
      showMap(event) {
            var map = document.getElementById("details-map-container")
            map.style.display = "block"
            var con = document.getElementById("ex-sign")
            con.style.display = "block"
            console.log(event);
      }, 
      disappearEx() {
        var map = document.getElementById("details-map-container")
        map.style.display = "none"
        var fmap = document.getElementById("final-map-container")
        fmap.style.display = "none"
        var ex = document.getElementById("ex-sign")
        ex.style.display = "none"
      },
      showFinalMap() {
        var map = document.getElementById("final-map-container")
        map.style.display = "block"
        var con = document.getElementById("ex-sign")
        con.style.display = "block"
        var details = document.getElementById('flight-details');
        details.style.display = 'None';
        
      },
      lowInfo() {
        var c = document.getElementById('lowc');
        c.style.display = 'block';
      },
      midInfo() {
        var c = document.getElementById('midc');
        c.style.display = 'block';
      },
      highInfo() {
        var c = document.getElementById('highc');
        c.style.display = 'block';
      },
      hideInfo() {
        const infos = document.getElementsByClassName('info-container');
        var i = 0;
        while (i < 3) {
            infos[i].style.display = 'none';
            i ++;
        }
      }, 
      logme({c, d}) {
        console.log("RECEIVED IN PARENT",c.lat, c.lng, d.lat, d.lng);
        this.coords.sourcelatitude = c.lat;
        this.coords.sourcelongitude = c.lng;
        this.coords.destlatitude = d.lat;
        this.coords.destlongitude = d.lng;
      },
      setWaypoint() {
        var w = document.getElementById("waypoint-container")
        w.style.display = "block";
      }
    //   initMap() {
    //     var options = {     
    //         zoom: 10,
    //         center: { lat: 33.933241, lng: -84.340288 }
    //     }
    //     var mapContainer = document.getElementById('map-container');
    //     var myMap = new google.maps.Map(mapContainer, options);
    //     var marker = new google.maps.Marker({
    //         position: { lat: 33.933241, lng: -84.340288 },
    //         map: myMap
    //     });
    //     console.log(marker);
    //   }
    }
    // }, mounted() {

    // fetch('http://localhost:3333', {
    //                 method: 'GET'
    //             })
    // .then(res => res.json())
    // .then(data => this.info = data.message, console.log(this.info))
    // .catch(err => console.log(err.message))
}


</script>


