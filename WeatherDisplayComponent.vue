<script setup>
</script>

<template>
    <section class="weather-container">
        <div class="weather-header"><h1>5 day forecast: </h1></div>
        <div class="warning-descriptions">
            <div v-for="(color, index) in definedColours" v-bind:key="index">
                <span class="colour-signal-unavailable" :style="{ backgroundColor: color}"></span>
                <span v-if="index === 0">DO NOT FLY</span>
                <span v-if="index === definedColours.length-1">SAFE TO FLY</span>
            </div>
        </div>
        <div class="displayWindValues">
            <div class="weather-cols">
                <div class="weather-heading">Date and time:</div>
                <div v-for="(date, index) in windDates" v-bind:key="index">
                    <div v-if="index % 24 !== 0" :style="{backgroundColor: 'grey', border: 'solid 1px white', padding: '5px'}">
                        {{ date }}
                    </div>
                    <div v-else :style="{padding: '8.35px', fontSize: '14pt', borderRight: 'solid 1px white', paddingBottom:'3.4px'}">
                        
                        <div v-if="index < 2">
                            {{ windDatesDay[0] }}
                        </div>
                        <div v-else-if="2 < index && index < 25">
                            {{ windDatesDay[1] }}
                        </div>
                        <div v-else-if="25 < index && index < 49">
                            {{ windDatesDay[2] }}
                        </div>
                        <div v-else-if="49 < index && index < 73">
                            {{ windDatesDay[3] }}
                        </div>
                        <div v-else-if="73 < index && index < 97">
                            {{ windDatesDay[4] }}
                        </div>
                        <div v-else-if="97 < index && index < 123">
                            {{ windDatesDay[5] }}
                        </div>
                        <div v-else-if="121 < index && index < 150">
                            {{ windDatesDay[6] }}
                        </div>
                    </div>
                        
                </div>
            </div>
            <div class="weather-cols">
                <div class="weather-heading"> Wind speed(km/h):</div>
                <div v-for="(wind, index) in windValues" v-bind:key="index">
                    <div v-if="index % 24 !== 0" :style="{backgroundColor: 'grey', border: 'solid 1px white', padding: '5px'}">
                        {{ wind }}
                    </div>
                    <div v-else :style="{padding: '10px', paddingTop:'5px', borderLeft: 'solid 1px white'}">
                        <div style="visibility: hidden;">WARN</div>
                    </div>
                </div>
                
            </div>
            <div class="weather-cols">
                <div class="weather-heading"> Wind direction:</div>
                <div v-for="(wind, index) in windDirections" v-bind:key="index">
                    <div v-if="index % 24 !== 0" :style="{backgroundColor: 'grey', border: 'solid 1px white', padding: '5px'}">
                        {{ wind }}
                    </div>
                    <div v-else :style="{padding: '10px', paddingTop:'5px', borderLeft: 'solid 1px white'}">
                        <div style="visibility: hidden;">WARN</div>
                    </div>
                </div>
                
            </div>
            <div class="weather-cols">
                <div class="weather-heading">Status:</div>
                <div v-for="(c, index) in colorList" v-bind:key="index" >
                    <div v-if="index % 24 !== 0" :style="{ backgroundColor: c, border: 'solid 1px white', padding: '5px' }">
                        <div style="visibility: hidden;">WARN</div>
                    </div>
                    <div v-else :style="{ display: none, padding: '10px', paddingTop:'5px'}">
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
    grid-template-columns: 20% 20% 20% 40%;
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
    border: solid 1px white;
}

.warning-descriptions {
    padding: 17px;
    left: 87.5%;
    position:fixed;
    background-color: rgb(80, 79,79);
    font-size: .9em;;
}

.colour-signal-unavailable {
    width: 15px;
    height: 5px;
    padding-left: 15px;
    margin: 10px;
    background-color: rgb(230, 36, 11);
    border: solid .5px rgb(230, 36, 11);
}


</style>

<script>
import { useGeolocation } from '../useGeolocation'
import { computed } from 'vue'
//import { getWindSpeed } from '@/fetchWindSpeed';
export default {
    data() {
      return {
        windData: 0,
        windValues: [],
        windDates: [],
        windDirections: [],
        windDatesDay: [],
        colorList: [],
        definedColours: ["#c00000", "#f30606", "#ff5700", "#f79a35", "#ffd027", "#ffee37", "#83a7f9"],
        counter: 0,
      }
    },
    name: "WeatherComponent",
    mounted() {
        this.getWind()
    },
    methods: {
        async getWind() {
            if (this.counter === 0) { //prevent multiple calls
                console.log("calling")
                const { coords } = useGeolocation()
                const initial = computed(() => ({
                lat: coords.value.latitude.toString(),
                lng: coords.value.longitude.toString()
                }))
                var lat = "51.892609777851305"
                var lng = "-8.50142240524292"
                console.log("LAT--->", initial.value.lat)
                //var res = await fetch("https://api.open-meteo.com/v1/forecast?latitude="+initial.value.lat+"&longitude="+initial.value.lng+"&hourly=windspeed_80m")
                var res = await fetch("https://api.open-meteo.com/v1/forecast?latitude="+lat+"&longitude="+lng+"&hourly=windspeed_80m,winddirection_80m")
                var final = await res.json()
                console.log("final", final)
                this.windData = final
                this.windValues = this.windData.hourly.windspeed_80m
                var tempDates = this.windData.hourly.time
                this.windDirections = this.windData.hourly.winddirection_80m
                this.windDirections = this.convertDegreesToDirection(this.windDirections)

                this.getColors()

                var cursor;
                for (var val in tempDates) {
                    this.windDates.push(tempDates[val].slice(11,16))
                    
                    if (cursor != tempDates[val].slice(0, 10)) {
                        cursor =  tempDates[val].slice(0, 10)
                        this.windDatesDay.push(cursor)
                    }
                    
                    
                }
                console.log("len(index)", this.windDates.length)
                console.log(this.windDatesDay)
                this.counter++
            }

        },
        getColors() {
            for (var val in this.windValues) {
                var intWindVal = parseInt(this.windValues[val])
                console.log(intWindVal)
                if (intWindVal >= 55) {      
                    this.colorList.push("#c00000")
                } else if (intWindVal < 55 && intWindVal >= 51) {
                    this.colorList.push("#f30606")
                } else if (intWindVal < 51 && intWindVal >= 45) {
                    this.colorList.push("#ff5700")
                } else if (intWindVal < 45 && intWindVal >= 41) {
                    this.colorList.push("#f79a35")
                } else if (intWindVal < 41 && intWindVal >= 35) {
                    this.colorList.push("#ffd027")
                } else if (intWindVal < 35 && intWindVal >= 31) {
                    this.colorList.push("#ffee37")
                } else {
                    this.colorList.push("#83a7f9")
                }
            }
            console.log("colours", this.colorList)
        },
        convertDegreesToDirection(degreeList) {
            var listOfDirections = []
            for (var i in degreeList) {
                var degree = degreeList[i]
                if (degree <= 180) {
                    if (degree <= 90) {
                        if (degree <= 22.5) {
                            listOfDirections.push("North-NorthEast")
                            console.log("degreeList[i]", degreeList[i])
                        }
                        if (22.5 < degree && degree <= 45) {
                            listOfDirections.push("NorthEast")
                            console.log("degreeList[i]", degreeList[i])
                        }
                        if (45 < degree && degree <= 67.5) {
                            listOfDirections.push("East-NorthEast")
                            console.log("degreeList[i]", degreeList[i])
                        }
                        if (67.5 < degree && degree <= 90) {
                            listOfDirections.push("East")
                            console.log("degreeList[i]", degreeList[i])
                        }
                    }
                    else {
                        if (90 < degree && degree <= 112.5) {
                            listOfDirections.push("East-SouthEast")
                            console.log("degreeList[i] ESE", degreeList[i])
                        }
                        if (112.5 < degree && degree <= 135) {
                            listOfDirections.push("SouthEast")
                            console.log("degreeList[i] S", degreeList[i])
                        }
                        if (135 < degree && degree <= 157.5) {
                            listOfDirections.push("South-SouthEast")
                            console.log("degreeList[i] SS", degreeList[i])
                        }
                        if (157.5 < degree && degree <= 180) {
                            listOfDirections.push("South")
                            console.log("degreeList[i] S", degreeList[i])
                        }
                    }
                } 
                else {
                    if (degree <= 270) {
                        if (180 < degree && degree <= 202.5) {
                            listOfDirections.push("South-SouthWest")
                            console.log("degreeList[i] SWS", degreeList[i])
                        }
                        if (202.5 < degree && degree <= 225) {
                            listOfDirections.push("SouthWest")
                            console.log("degreeList[i] SW", degreeList[i])
                        }
                        if (225 < degree && degree <= 247.5) {
                            listOfDirections.push("West-SouthWest")
                            console.log("degreeList[i] WSW", degreeList[i])
                        }
                        if (247.5 < degree && degree <= 270) {
                            listOfDirections.push("West")
                        }
                    }
                    else {
                        if (270 < degree && degree <= 292.5) {
                            listOfDirections.push("West-NorthWest")
                        }
                        if (292.5 < degree && degree <= 315) {
                            listOfDirections.push("NorthWest")
                        }
                        if (315 < degree && degree <= 337.5) {
                            listOfDirections.push("North-NorthWest")
                        }
                        if (337.5 < degree && degree <= 360)  {
                            listOfDirections.push("North")
                        }
                    }
                }
            }
            console.log(listOfDirections.length)
            return listOfDirections
            
        }
    }
}

</script>