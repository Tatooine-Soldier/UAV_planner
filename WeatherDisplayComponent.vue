<script setup>
</script>

<template>
    <section class="weather-container">
        <div class="weather-header"><h1>5 day forecast: </h1></div>
        <div @click="getWind()">Click</div>
        <div class="displayWindValues">
            <div class="weather-cols">
                <div class="weather-heading">Date and time:</div>
                <div v-for="(date, index) in windDates" v-bind:key="index">
                    <div v-if="index % 24 !== 0">
                        {{ date }}
                    </div>
                    <div v-else :style="{padding: '10px'}">
                        New Date
                    </div>
                        
                </div>
            </div>
            <div class="weather-cols">
                <div class="weather-heading">Estimated wind speed:</div>
                <div v-for="(wind, index) in windValues" v-bind:key="index">
                    <div v-if="index % 24 !== 0">
                        {{ wind }}
                    </div>
                    <div v-else :style="{padding: '10px'}">
                        New Date
                    </div>
                </div>
                
            </div>
            <div class="weather-cols">
                <div class="weather-heading">Status:</div>
                <div v-for="(c, index) in colorList" v-bind:key="index" >
                    <div v-if="index % 24 !== 0" :style="{ backgroundColor: c }">
                        <div style="visibility: hidden;">WARN</div>
                    </div>
                    <div v-else :style="{ display: none, padding: '10px'}">
                        <div style="visibility: hidden;">WARN</div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<style>
.displayWindValues {
    display: grid;
    grid-template-columns: 30% 30% 40%;
}

.weather-container {
    color: white;
}

.weather-cols {
    text-align: center;
}

.weather-header {
    text-align: center;
    font-size: 20pt;;
}

h1 {
    margin-top: 10px;
    margin-bottom: 0px;
}

.weather-heading {
    padding: 10px;
    font-size: 16pt;
}
</style>

<script>
import { useGeolocation } from '../useGeolocation'
import { computed } from 'vue'
export default {
    data() {
      return {
        windData: 0,
        windValues: [],
        windDates: [],
        windDatesDay: [],
        colorList: []
      }
    },
    name: "WeatherComponent",
    methods: {
        async getWind() {
            console.log("calling")
            const { coords } = useGeolocation()
            const initial = computed(() => ({
            lat: coords.value.latitude.toString(),
            lng: coords.value.longitude.toString()
            }))
            console.log("LAT--->", initial.value.lat)
            var res = await fetch("https://api.open-meteo.com/v1/forecast?latitude="+initial.value.lat+"&longitude="+initial.value.lng+"&hourly=windspeed_80m")
            var final = await res.json()
            console.log("final", final)
            this.windData = final
            this.windValues = this.windData.hourly.windspeed_80m
            var tempDates = this.windData.hourly.time
            this.getColors()

            var cursor;
            for (var val in tempDates) {
                this.windDates.push(tempDates[val].slice(11,16))
                
                if (cursor != tempDates[val].slice(0, 10)) {
                    cursor =  tempDates[val].slice(0, 10)
                    this.windDatesDay.push(cursor)
                }
                
                
            }
            console.log("DAYS", this.windDatesDay)

        },
        getColors() {
            for (var val in this.windValues) {
                var intWindVal = parseInt(this.windValues[val])
                console.log(intWindVal)
                if (intWindVal >= 40) {      
                    this.colorList.push("#c00000")
                } else if (intWindVal < 40 && intWindVal >= 35) {
                    this.colorList.push("#f30606")
                } else if (intWindVal < 35 && intWindVal >= 29) {
                    this.colorList.push("#ff5700")
                } else if (intWindVal < 29 && intWindVal >= 25) {
                    this.colorList.push("#f79a35")
                } else if (intWindVal < 25 && intWindVal >= 19) {
                    this.colorList.push("#ffd027")
                } else if (intWindVal < 19 && intWindVal >= 14) {
                    this.colorList.push("#ffee37")
                } else {
                    this.colorList.push("#83a7f9")
                }
            }
            console.log("colours", this.colorList)
        }
    }
}

</script>