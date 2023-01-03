<script setup>
  import MyGoogleMapComponent from "../components/MyGoogleMapComponent.vue"
</script>

<template>
    <section class="flight-planner" id="flight-planner">
        <!-- <form action="/planner" method="post"> -->
        <form @submit.prevent="handleSubmit()">
            <section class="flight-details-container" id="flight-details-container">
                <section class="flight-details">
                    <section class="flight-details-content">
                        <h1>Confirm Flight Details</h1>
                        <section class="flight-details-content-grid">
                            <div>Date: {{ date.day }}</div>
                            <div>Time: {{ date.hour }}:{{ date.minute }}</div>
                            <div>Speed: {{ speed }} </div>
                            <div>Calculated distance: {{ distance }} KM</div>
                        </section>
                        <section class="flight-details-buttons">
                            <input id="cancel-but" name="but" type="button" value="Cancel" v-on:click="disappear()"/>
                            <button id="confirm-but">Confirm</button>
                        </section>
                    </section>
                </section>
            </section>
            <section id="details-map-container">
                <MyGoogleMapComponent></MyGoogleMapComponent>
            </section>
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
                                    <section class="coords-inputs">
                                        <label for="latitude">Latitude: </label>
                                        <input type="text" name="latitude" size="16" ref="mylat"/>
                                    </section>
                                    <section class="coords-inputs">
                                        <label for="longitude">Longitude: </label>
                                        <input type="text" name="longitude" size="16" ref="mylong">
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
                                    <section class="speeds-option">
                                        <section class="speeds-inputs">
                                            <label for="lowspeed">Low: </label>
                                            <input type="radio" name="speed" value="lowspeed"  ref="speed" />
                                        </section>
                                        <section class="speeds-inputs">
                                            <label for="midspeed">Mid: </label>
                                            <input type="radio" name="speed" value="midspeed"  ref="speed" />
                                        </section>
                                        <section class="speeds-inputs">
                                            <label for="highspeed">High: </label>
                                            <input type="radio" name="speed" value="highspeed" ref="speed" />
                                        </section>
                                    </section>
                                    
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
        grid-template-rows: 50% 50%;
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

    #details-map-container {
        position: absolute;
        display: none;
        z-index: 1;
        left: 10%;
        right: 10%;
        background-color: grey;
        border: solid 1px grey;

    }
  

</style>

<!-- <script defer src="<https://maps.googleapis.com/maps/api/js?key=AIzaSyDTNOMjJP2zMMEHcGy2wMNae1JnHkGVvn0&callback=initMap>"> -->
<script>
  import * as geolib from 'geolib';

export default {
    data() {
      return {
        date: {
          hour: '',
          minute: '',
          day: ''
        },
        coords: {
            longitude: '',
            latitude: ''
        },
        speed: null,
        info: null,
        distance: null
      }
    },
  
    methods: {
      handleSubmit() {
        // Send data to the server or update your stores and such.
        this.coords.latitude = this.$refs.mylat.value;
        this.coords.longitude = this.$refs.mylong.value;
        console.log(this.coords)
        var long = parseInt(this.coords.longitude)
        var lat = parseInt(this.coords.latitude)
        this.distance = geolib.getDistance({latitude: 0, longitude:0}, {
                    latitude: lat,
                    longitude: long,
                })
        this.distance = geolib.convertDistance(this.distance, 'km');
        console.log(
            'DISTANCE-->',
             this.distance,
            '<--DISTANCE'
        )
        var details = document.getElementById('flight-details-container');
        details.style.display = 'block';

        this.speed = this.$refs.speed.value;
    

        var grey = document.getElementById('flight-planner-columns');
        grey.style.opacity = 0.2;
        grey.style.pointerEvents = "none";

        
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
            console.log(event);
      },
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


