<script>
import { Loader } from '@googlemaps/js-api-loader'
    /* eslint-disable no-undef*/
    /* eslint-disable no-unused-vars*/
    import { computed, ref, onMounted, onUnmounted, watch } from 'vue'
    import { useGeolocation } from '../useGeolocation'
    // import haversineDistance from './calculateDistance'
    const GOOGLE_MAPS_API_KEY = 'AIzaSyDTNOMjJP2zMMEHcGy2wMNae1JnHkGVvn0'
    export default {
      name: 'App',
      setup() {
        const { coords } = useGeolocation()
        const initial = computed(() => ({
          lat: coords.value.latitude,
          lng: coords.value.longitude
        }))
        
        const currPos = ref(null)
        const otherLoc = ref(null)
        let clickListener = null;
        let dragListener = null;
        let destDragListener = null;
        let cursor = {lat: 0.0, lng: 0.0}
        

        const loader = new Loader({ apiKey: GOOGLE_MAPS_API_KEY})
        const mapDivHere = ref(null);
      
        let sourceMarker = ref(null)
        let destMarker = ref(null)
  
        let map = ref(null)
        
        onMounted(async () => {
          await loader.load()
          currPos.value = {lat: initial.value.lat, lng: initial.value.lng}
          console.log(currPos.value)
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
              }),

              destDragListener = destMarker.value.addListener(
                'drag',
                function(event) {
                  otherLoc.value = {lat: event.latLng.lat(), lng: event.latLng.lng()}
                  console.log("OTHER-->",otherLoc.value.lat)
                  var d = document.getElementById('destCursorLat')
                  d.innerHTML = otherLoc.value.lat;
                  var c = document.getElementById('destCursorLng')
                  c.innerHTML = otherLoc.value.lng;
                }
              )
            )
          )
        

          sourceMarker.value = new google.maps.Marker({
            position: currPos.value,
            draggable: true,
            map: map.value
          })
          var d = document.getElementById('cursorLat')
          d.innerHTML = currPos.value.lat;
          var c = document.getElementById('cursorLng')
          c.innerHTML = currPos.value.lng;
          dragListener = sourceMarker.value.addListener(
            'drag',
            function(event) {
              currPos.value = {lat: event.latLng.lat(), lng: event.latLng.lng()}
              console.log("CURRPOS.LAT",currPos.value.lat)
              cursor.lat = event.latLng.lat(),
              cursor.lng = event.latLng.lng()
              var d = document.getElementById('cursorLat')
              d.innerHTML = cursor.lat;
              var c = document.getElementById('cursorLng')
              c.innerHTML = cursor.lng;
              // this.$emit('changeLocation',cursor);
            }
          )

          const airports = {
            dub: {
              center: { lat: 53.4264, lng: -6.2499 }
            },
            sha: {
              center: { lat: 52.6996575, lng: -8.9806044 }
            },
            crk: {
              center: { lat: 51.8490624, lng: -8.4920732 }
            },
          };

          for (const airport in airports) {
              const cityCircle = new google.maps.Circle({
              strokeColor: "#FF0000",
              strokeOpacity: 0.8,
              strokeWeight: 2,
              fillColor: "#FF0000",
              fillOpacity: 0.35,
              map: map.value,
              center: airports[airport].center,
              radius: 5000
            });
          }
          //do smaller airports next, 10 nautical mile radius, the 3 larger airports hav an extended area outside area already given so add it in
          //add onclick listener to the circle

        
        })
        onUnmounted(async () => {
            if (clickListener) clickListener.remove()
            if (dragListener) dragListener.remove()
            if (destDragListener) destDragListener.remove()
        
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

      //   const flightPathCoordinates = [
      //   { lat: currPos.value.lat+2, lng: currPos.value.lng-1 },
      //   { lat: currPos.value.lat+5, lng: currPos.value.lng-1.7 },
      //   { lat: currPos.value.lat+10, lng: currPos.value.lng-2.3 }
      // ];        

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
        return { currPos, otherLoc, distance, mapDivHere, calculatedTime}
      },
    
    }
  

</script>

<template>
  <div id="big-container">
    <div class="distance-caption-container">
      <div class="distance-and-time">
        <div>
          Distance of path(km): {{ distance }}
        </div>
        <div>
          Estimated arrival: {{ calculatedTime }}
        </div>
      </div>
      <div class="source-coords">
        <div class="source-coords-title">
          STARTING POINT:
        </div>
        <div>
          Lat: <span class="coords-display" id="cursorLat"></span>
        </div>
        <div>
          Lng: <span class="coords-display" id="cursorLng"></span>
        </div>
      </div>
      <div class="dest-coords">
        <div class="source-coords-title">
          DESTINATION POINT
        </div>
        <div>
          Lat: <span class="coords-display" id="destCursorLat"></span>
        </div>
        <div>
          Lng: <span class="coords-display" id="destCursorLng"></span>
        </div>
      </div>
    </div>
    <!-- <button id="coords-confirm-button" @click="$emit('someEvent', {c:currPos, d:otherLoc} )">Confirm</button> -->
    <div class="submit-div" @click="$emit('someEvent', {c:currPos, d:otherLoc, distance:distance} )">CONFIRM</div>
    <div ref="mapDivHere" style="width:100%; height:80vh;"/>
  </div>

</template>

<style>
  .distance-caption-container {
    background-color: white;
    padding: 5px;
    display: grid;
    grid-template-columns: 35% 35% 30%;
  }

  .distance-and-time {
    display: grid;
    grid-template-rows: 50% 50%;
  }

  .source-coords {
    display: grid;
    grid-template-rows: 20% 40% 40%;
  }

  .dest-coords {
    display: grid;
    grid-template-rows: 20% 40% 40%;
  }

  .coords-display {
    font-size: 12px;
    width: 60%;
    padding: 2px;
  }

  .source-coords-title {
    padding-bottom:20px;
  }

  #coords-confirm-button {
    position: static;
    padding: 2px;
    width: 100%;
  }

  .submit-div {
    background-color: white;
    border: solid 1px grey;
    text-align: center;
    padding: 3px;
    transition: 0.4s;
  }

  .submit-div:hover {
    background-color: rgb(101, 100, 100);
    color: white;
  }
</style>
