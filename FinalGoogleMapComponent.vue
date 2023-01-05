<script>
import { Loader } from '@googlemaps/js-api-loader'
    /* eslint-disable no-undef*/
    import { computed, ref, onMounted, onUnmounted, watch } from 'vue'
    //import { useGeolocation } from '../useGeolocation'
    // import haversineDistance from './calculateDistance'
    const GOOGLE_MAPS_API_KEY = 'AIzaSyDTNOMjJP2zMMEHcGy2wMNae1JnHkGVvn0'
    export default {
      name: 'App',
      props: ['propcoords'],
      data: function() {
        var myData = {
            myProp: this.propcoords
        }
        return myData
      },
      setup(props) {
        //const { coords } = useGeolocation()
        const currPos = computed(() => ({
          lat: parseFloat(props.propcoords.sourcelatitude),
          lng: parseFloat(props.propcoords.sourcelongitude)
        }))
        const otherLoc = computed(() => ({
          lat: parseFloat(props.propcoords.destlatitude),
          lng: parseFloat(props.propcoords.destlongitude)
        }))
        let clickListener = null;

        const loader = new Loader({ apiKey: GOOGLE_MAPS_API_KEY})
        const mapDivHere = ref(null);
      
        let sourceMarker = ref(null)
        let destMarker = ref(null)
        let map = ref(null)
        onMounted(async () => {
          await loader.load()
          console.log("currPos CENTER--->", currPos.value)
          map.value = new google.maps.Map(mapDivHere.value, {
                center: currPos.value,
                zoom: 8
            })
          sourceMarker.value = new google.maps.Marker({
            position: currPos.value,
            map: map.value
          })
          destMarker.value = new google.maps.Marker({
            position: otherLoc.value,
            map: map.value
          })
          const sourceinfowindow = new google.maps.InfoWindow({
            content: "Starting Point",
            ariaLabel: "Source",
          });
          sourceMarker.value.addListener("click", () => {
            sourceinfowindow.open({
                anchor: sourceMarker.value,
                map,
                });
            });

          const destinfowindow = new google.maps.InfoWindow({
            content: "Destination Point",
            ariaLabel: "Destination",
          });
          destMarker.value.addListener("click", () => {
            destinfowindow.open({
                anchor: destMarker.value,
                map,
                });
            });
        //     window.initMap = () => {
        //     map.value = new google.maps.Map(mapDivHere.value, {
        //         center: currPos.value,
        //         zoom: 8
        //     })
        //   };
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
  <div id="big-container">
    <div class="final-distance-caption-container">
      <div>Distance of path(km): {{ distance }}<br> Estimated arrival: {{ calculatedTime }}</div>
      <div></div>
      <div class="detailz">
            Source Latitude:  <i id="plat">{{ propcoords.sourcelatitude }}</i><br>Source Longitude:  <i id="plng">{{ propcoords.sourcelongitude }}<br></i>
            Destination Latitude:  <i id="plat">{{ propcoords.destlatitude }}</i><br>Destination Longitude:  <i id="plng">{{ propcoords.destlongitude }}<br></i>
      </div>
    </div>
    <div ref="mapDivHere" style="width:100%; height:80vh;"/>
  </div>

</template>

<style>
  .final-distance-caption-container {
    background-color: white;
    padding: 5px;
    display: grid;
    grid-template-columns: 35% 35% 30%;
  }

  .detailsz {
    font-size: 10px;
    text-align: left;
  }
</style>