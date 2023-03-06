<script setup>
</script>

<template>
    <section class="weather-container">
        <div>5 day forecast: </div>
        <div @="getWind()">Click</div>
        <div class="displayWindValues">
            <div class="weather-cols">
                Date and time:
                <div v-for="(date, index) in windDates" v-bind:key="index">
                    {{ date }}
                </div>
            </div>
            <div class="weather-cols">
                Estimated wind speed:
                <div v-for="(wind, index) in windValues" v-bind:key="index">
                    {{ wind }}
                </div>
            </div>
        </div>
    </section>
</template>

<style>
.displayWindValues {
    display: grid;
    grid-template-columns: 50% 50%;
}

.weather-container {
    color: white;
}

.weather-cols {
    text-align: center;
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
        windDates: []
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
            this.windDates = this.windData.hourly.time
        }
    }
}

</script>