
<template>
    <section class="flight-planner">
        <!-- <form action="/planner" method="post"> -->
        <form @submit.prevent="handleSubmit()">
        <section class="flight-planner-columns">
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
                                            <input type="radio" name="lowspeed" value="lowspeed"  />
                                        </section>
                                        <section class="speeds-inputs">
                                            <label for="midspeed">Mid: </label>
                                            <input type="radio" name="midspeed" value="midspeed"  />
                                        </section>
                                        <section class="speeds-inputs">
                                            <label for="highspeed">High: </label>
                                            <input type="radio" name="highspeed" value="highspeed" />
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
                            {{ info }}
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

    button {
        margin-top: 10%;;
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


  

</style>


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
        speed: {
            lowspeed: '',
            midspeed: '',
            highspeed: ''
        },
        info: null,
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
        console.log(
            'DISTANCE-->',
            geolib.getDistance({latitude: 0, longitude:0}, {
                    latitude: lat,
                    longitude: long,
                }),
            '<--DISTANCE'
        )
      },
      handleSubmitCoords() {
        console.log(this.coords)
      }
    // }, mounted() {

    // fetch('http://localhost:3333', {
    //                 method: 'GET'
    //             })
    // .then(res => res.json())
    // .then(data => this.info = data.message, console.log(this.info))
    // .catch(err => console.log(err.message))
}
}
</script>
