<template>
    <section class="profile-page">
        <section class="profile-main">
            <header>
                <p>Profile Page</p>
                <h2>Welcome {{ user }}</h2>
            </header>
            <div>
                <p>Past Flights: </p>
                <!-- <p v-for="(flight, index) in flightData" v-bind:key="index">
                    {{ flight }}
                </p> -->

                <div class="flight-details-headings">
                    <b>Flight ID</b>
                    <b>Date</b>
                    <b>Start Time</b>
                    <b>End Time</b>
                    <b>Altitude</b>
                </div>
                <div class="displayFlightDetails" v-for="(i, index) in flightData" v-bind:key="index">
                        <p class="flight-details-records">
                            {{ flightIDs[index] }}
                        </p>
                        <p class="flight-details-records">
                            {{ flightDates[index] }}
                        </p>
                        <p class="flight-details-records">
                            {{ flightStartTimes[index] }}
                        </p>
                        <p class="flight-details-records">
                            {{ flightEndTimes[index] }}
                        </p>
                        <p class="flight-details-records">
                            {{ flightAltitudes[index] }}
                        </p>
                </div>
            </div>
        </section>
    </section>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            user: "",
            flightData: [],
            flightIDs: [],
            flightDates: [],
            flightStartTimes: [],
            flightEndTimes: [],
            flightAltitudes: [],
            pastCoordinates: []
        }
    },
    async created() {
        axios
        .post("/userProfile")
        .then((response) => {
          const data = response.data;
          console.log("data: ", data)
          var dataArray = data.split("|")
          console.log("dataArray--> ", dataArray)
          this.user = dataArray[0]

          var flights = dataArray[1]
          const jsonArray = JSON.parse(flights);
          console.log("jsonArray-->", jsonArray)
          this.flightData = jsonArray
          this.arrangeDetails()
        })
        .catch(function (error) {
            console.log("ERROR:", error); 
        })
    },
    methods: {
        getIDs() {

        },
        arrangeDetails() {
            for (var flight in this.flightData) {
                console.log(this.flightData[flight].time)
                this.flightIDs.push(this.flightData[flight].id)
                this.flightDates.push(this.flightData[flight].date)
                this.flightStartTimes.push(this.flightData[flight].time)
                this.flightEndTimes.push(this.flightData[flight].endTime)
                this.flightAltitudes.push(this.flightData[flight].altitude)
            }
        }
    }
}
</script>

<style>
.profile-main {
    color: white;
}

header {
    text-align: center;
    border-bottom: solid 2px rgb(120, 255, 36);
    width: 60%;
    left: 20%;
    right: 20%;
}

.flight-details-records {
    text-align: center;
    border: solid grey 1px;
}

.displayFlightDetails {
    display: grid;
    grid-template-columns: 20% 20% 20% 20% 20%;
    background-color: rgb(50, 50, 50);
    border: solid grey 3px;
}

.displayFlightDetails:hover {
    background-color: rgb(70, 75, 70);
    cursor: pointer;
}

.flight-details-headings {
    display: grid;
    grid-template-columns: 20% 20% 20% 20% 20%;
    text-align: center;
}
</style>