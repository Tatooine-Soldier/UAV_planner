<script>
import { Loader } from '@googlemaps/js-api-loader'
    /* eslint-disable no-undef*/
    import { computed, ref, onMounted, onUnmounted, watch } from 'vue'
    import { useGeolocation } from '../useGeolocation'
    // import haversineDistance from './calculateDistance'
    const GOOGLE_MAPS_API_KEY = 'AIzaSyDTNOMjJP2zMMEHcGy2wMNae1JnHkGVvn0'
    export default {
      name: 'App',
      setup() {
        const { coords } = useGeolocation()
        const currPos = computed(() => ({
          lat: coords.value.latitude,
          lng: coords.value.longitude
        }))
        const otherLoc = ref(null)
        let clickListener = null;

        const loader = new Loader({ apiKey: GOOGLE_MAPS_API_KEY})
        const mapDivHere = ref(null);
      
        let sourceMarker = ref(null)
        let destMarker = ref(null)
        let map = ref(null)
        onMounted(async () => {
          await loader.load()
          map.value = new google.maps.Map(mapDivHere.value, {
            center: currPos.value,
            zoom: 9
          })
          clickListener = map.value.addListener(
            'click',
            ({latLng: {lat, lng}}) => 
              (otherLoc.value = {lat: lat(), lng: lng()},
              destMarker.value = new google.maps.Marker({
                position: otherLoc.value,
                draggable: true,
                map: map.value
              })
            )
          )
          sourceMarker.value = new google.maps.Marker({
            position: currPos.value,
            draggable: true,
            map: map.value
          })
        })
        onUnmounted(async () => {
            if (clickListener) clickListener.remove()
        })
        let line = null
        watch([map, currPos, otherLoc], () => {
          if (line) line.setMap(null)
          if (map.value && otherLoc.value != null)
            line = new google.maps.Polyline({
              path: [currPos.value, otherLoc.value],
              map: map.value
            })
        })
        const haversineDistance = (pos1, pos2) => {
        const R = 3958.8 // Radius of the Earth in miles
        const rlat1 = pos1.lat * (Math.PI / 180) // Convert degrees to radians
        const rlat2 = pos2.lat * (Math.PI / 180) // Convert degrees to radians
        const difflat = rlat2 - rlat1 // Radian difference (latitudes)
        const difflon = (pos2.lng - pos1.lng) * (Math.PI / 180) // Radian difference (longitudes)
        const d =
          2 *
          R *
          Math.asin(
            Math.sqrt(
              Math.sin(difflat / 2) * Math.sin(difflat / 2) +
                Math.cos(rlat1) *
                  Math.cos(rlat2) *
                  Math.sin(difflon / 2) *
                  Math.sin(difflon / 2)
            )
          )*1.609344  //convert to kilometres
          return d.toFixed(2)
        }
        const distance = computed(() =>
        otherLoc.value === null
          ? 0
          : haversineDistance(currPos.value, otherLoc.value)
        )
        const t = (distance, speed) => {
          const tme = distance/speed
          const remainder = tme%1
          console.log("REMAINDER",remainder)
          const minutes = 60*remainder
          console.log("MINUTES",minutes)
          var hours = tme/1
          console.log("HOURS",hours)
          if (hours < 1) {
            hours = 0
            const time = minutes.toFixed(2).toString()+" minutes"
            return time 
          }
          const time = hours.toFixed(0).toString() + " hours " + minutes.toFixed(2).toString() + " minutes"
          return time
        }
        const calculatedTime = computed(() =>
        distance.value === null 
          ? 0
          : t(distance.value, 30)
        )
        return { currPos, otherLoc, distance, mapDivHere, calculatedTime }
      }
    }
  

</script>

<template>

  <div class="distance-caption-container">
    <div>Distance of path(km): {{ distance }}</div>
    <div>Estimated arrival: {{ calculatedTime }}</div>
  </div>
  <div ref="mapDivHere" style="width:100%; height:80vh;"/>

</template>

<style>
  .distance-caption-container {
    background-color: white;
    padding: 5px;
    display: grid;
    grid-template-columns: 50% 50%;
  }
</style>
