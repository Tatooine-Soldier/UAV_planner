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

          // const airports = { //red zones
          //   dub: {
          //     center: { lat: 53.428923, lng: -6.265767 } 
          //   },
          //   sha: {
          //     center: { lat: 52.6996575, lng: -8.9806044 }
          //   },
          //   crk: {
          //     center: { lat: 51.8490624, lng: -8.4920732 }
          //   },
          // };

          const airports = {
            u2: {
              center: {lat: 53.273556, lng: -6.231778 },
              name: "Currency Centre Dublin",
              rad: 500,
              color: "#FBCEB1"
            },
            u3: {
              center: {lat: 53.343, lng: -6.302167 },
              name: "Royal Hospital Kilmainham",
              rad: 800,
              color: "#FBCEB1"
            },
            u7: {
              center: {lat: 53.302992, lng: -6.455406},
              name: "Casement Aerodrome",
              rad: 5000,
              color: "#FBCEB1"
            },
            u12: {
              center: {lat: 53.072778, lng: -6.036389},
              name: "Newcastle Aerodrome",
              rad: 2778,
              color: "#FBCEB1"
            },
            u17: {
              center: {lat: 53.424444, lng: -7.947778},
              name: "Custume Barracks",
              rad: 3704,
              color: "#FBCEB1"
            },
            u19: {
              center: {lat: 53.327939, lng: -6.270469},
              name: "Cathal Brugha Barracks",
              rad: 371,
              color: "#FBCEB1"
            },
            u21: {
              center: {lat: 53.352222, lng: -6.488333},
              name: "Weston Airport",
              rad: 3000,
              color: "#FBCEB1"
            },
            u24: {
              center: {lat: 53.439478, lng: -6.342445},
              name: "Dublin Control Terrain 1",
              rad: 650,
              color: "#FBCEB1"
            },
            u25: {
              center: {lat: 53.773674, lng: -6.336483},
              name: "Dublin Control Terrain 2",
              rad: 1500,
              color: "#FBCEB1"
            },
            u26: {
              center: {lat: 53.412841, lng: -6.337846},
              name: "Dublin Control Terrain 3",
              rad: 400,
              color: "#FBCEB1"
            },
            u27: {
              center: {lat: 53.428924, lng: -6.265767},
              name: "Dublin Control Red Zone",
              rad: 5000,
              color: "#FF0000"
            },
            u28: {
              center: {lat: 53.428924, lng: -6.265767},
              name: "Dublin Control Amber Zone",
              rad: 12100,
              color: "#FF8833"
            },
            u31: {
              center: {lat: 52.701976, lng: -8.924816},
              name: "Shannon Control Red Zone",
              rad: 5000,
              color: "#FF0000"
            },
            u64: {
              center: {lat: 52.701976, lng: -8.924816},
              name: "Shannon Control Amber Zone",
              rad: 12000,
              color: "#FBCEB1"
            },
            u32: {
              center: {lat: 53.910297, lng: -8.818491},
              name: "Connaught Control Red Zone",
              rad: 5000,
              color: "#FF0000"
            },
            u63: {
              center: {lat: 53.910297, lng: -8.818491},
              name: "Connaught Control Amber Zone",
              rad: 12000,
              color: "#FBCEB1"
            },
            u33: {
              center: {lat: 51.841269, lng: -8.491112},
              name: "Cork Control Red Zone",
              rad: 5000,
              color: "#FF0000"
            },
            u61: {
              center: {lat: 51.841269, lng: -8.491112},
              name: "Cork Control Amber Zone",
              rad: 12000,
              color: "#FBCEB1"
            },
            u67: {
              center: {lat: 51.825383, lng: -8.5958},
              name: "Cork Control Terrain 1",
              rad: 625,
              color: "#FBCEB1"
            },
            u68: {
              center: {lat: 51.789706, lng: -8.456769},
              name: "Cork Control Terrain 2",
              rad: 700,
              color: "#FBCEB1"
            },
            u34: {
              center: {lat: 52.180878, lng: -9.523784},
              name: "Kerry Control Red Zone",
              rad: 5000,
              color: "#FBCEB1"
            },
            u62: {
              center: {lat: 52.180878, lng: -9.523784},
              name: "Kerry Control Amber Zone",
              rad: 18520,
              color: "#FBCEB1"
            },
            u41: {
              center: {lat: 53.072778, lng: -6.036389},
              name: "Newcastle Aerodrome",
              rad: 2778,
              color: "#FBCEB1"
            },
            u42: {
              center: {lat: 53.350158, lng: -6.288103},
              name: "Arbour Hill Prison",
              rad: 800,
              color: "#FBCEB1"
            },
            u43: {
              center: {lat: 53.754108, lng: -8.487152},
              name: "Castlerea Prison",
              rad: 800,
              color: "#FBCEB1"
            },
            u44: {
              center: {lat: 53.341092, lng: -6.383033},
              name: "Cloverhill Prison",
              rad: 800,
              color: "#FBCEB1"
            },
            u45: {
              center: {lat: 51.909278, lng: -8.459992},
              name: "Cork Prison",
              rad: 800,
              color: "#FBCEB1"
            },
            u46: {
              center: {lat: 54.288706, lng: -7.91565},
              name: "Loughan House",
              rad: 800,
              color: "#FBCEB1"
            },
            u47: {
              center: {lat: 52.815725, lng: -6.190419},
              name: "Shelton Abbey",
              rad: 800,
              color: "#FBCEB1"
            },
            u48: {
              center: {lat: 53.4069, lng: -6.236819},
              name: "IPS BSD",
              rad: 500,
              color: "#FBCEB1"
            },
            u49: {
              center: {lat: 53.733286, lng: -7.774978},
              name: "IPS Headquarters",
              rad: 500,
              color: "#FBCEB1"
            },
            u50: {
              center: {lat: 53.356389, lng: 6.340556},
              name: "Leinster Model Flying Club",
              rad: 300,
              color: "#FBCEB1"
            },
            u51: {
              center: {lat: 53.506111, lng: -6.235278},
              name: "Ballyheary Flying Club",
              rad: 800,
              color: "#FBCEB1"
            },
            u52: {
              center: {lat: 53.5375, lng: -6.084167},
              name: "Fingal Model Flying Club",
              rad: 800,
              color: "#FBCEB1"
            },
            u53: {
              center: {lat: 53.225, lng: -6.318333},
              name: "Island Slope Rebel Flying Club",
              rad: 800,
              color: "#FBCEB1"
            },
            u69: {
              center: {lat: 55.044191, lng: -8.340999},
              name: "Donegal Control Red Zone",
              rad: 5000,
              color: "#FF0000"
            },
            u70: {
              center: {lat: 55.044191, lng: -8.340999},
              name: "Donegal Control Amber Zone",
              rad: 12000,
              color: "#FBCEB1"
            },
            u73: {
              center: {lat: 51.739722, lng: -8.694167},
              name: "Bandon Model Flying Club",
              rad: 800,
              color: "#FBCEB1"
            },
            u74: {
              center: {lat: 51.78, lng: -8.72},
              name: "Cork Model Aero Club",
              rad: 800,
              color: "#FBCEB1"
            },
            u75: {
              center: {lat: 51.620556, lng: -8.545},
              name: "Island Slope Rebels Club",
              rad: 800,
              color: "#FBCEB1"
            },
            u76: {
              center: {lat: 53.504917, lng: -6.234667},
              name: "BallyBoughal Airfield",
              rad: 1000,
              color: "#FBCEB1"
            },
            u77: {
              center: {lat: 52.698611, lng: -8.854722},
              name: "Shannon Model Flying Club",
              rad: 300,
              color: "#FBCEB1"
            },
            u78: {
              center: {lat: 52.712061, lng: -8.747769},
              name: "Shannon Control Terrain 1",
              rad: 1200,
              color: "#FBCEB1"
            },
            u79: {
              center: {lat: 52.657506, lng: -8.890856},
              name: "Shannon Control Terrain 2",
              rad: 200,
              color: "#FBCEB1"
            },
            u80: {
              center: {lat: 52.703039, lng: -9.077886},
              name: "Shannon Control Terrain 3",
              rad: 1600,
              color: "#FBCEB1"
            },
            u81: {
              center: {lat: 52.709694, lng: -8.829472},
              name: "Shannon Control Terrain 4",
              rad: 1200,
              color: "#FBCEB1"
            },
            u82: {
              center: {lat: 52.711411, lng: -9.000114},
              name: "Shannon Control Terrain 5",
              rad: 800,
              color: "#FBCEB1"
            },
            u85: {
              center: {lat: 52.762108, lng: -9.048511},
              name: "Shannon Control Terrain 8",
              rad: 600,
              color: "#FBCEB1"
            },


            u87: {
              center: {lat: 54.280213, lng: -8.599208},
              name: "Sligo Control Red Zone",
              rad: 5000,
              color: "#FF0000"
            },
            u88: {
              center: {lat: 54.280213, lng: -8.599208},
              name: "Sligo Control Amber Zone",
              rad: 12000,
              color: "#FBCEB1"
            },
            


            u89: {
              center: {lat: 52.1872, lng: -7.086963},
              name: "Waterford Control Red Zone",
              rad: 5000,
              color: "#FF0000"
            },
            u90: {
              center: {lat: 52.1872, lng: -7.086963},
              name: "Waterford Control Amber Zone",
              rad: 12000,
              color: "#FBCEB1"
            },
            u91: {
              center: {lat: 52.184514, lng: -7.179378},
              name: "Waterford Control Terrain 1",
              rad: 3800,
              color: "#FBCEB1"
            },
            u92: {
              center: {lat: 52.231403, lng: -7.158811},
              name: "Waterford Control Terrain 2",
              rad: 600,
              color: "#FBCEB1"
            },
            u93: {
              center: {lat: 52.265714, lng: -7.006911},
              name: "Waterford Control Terrain 3",
              rad: 600,
              color: "#FBCEB1"
            },
            u94: {
              center: {lat: 52.216361, lng: -7.016533},
              name: "Waterford Control Terrain 4",
              rad: 1700,
              color: "#FBCEB1"
            },
            u95: {
              center: {lat: 52.167847, lng: -7.025103},
              name: "Waterford Control Terrain 5",
              rad: 3500,
              color: "#FBCEB1"
            },

            u96: {
              center: {lat: 53.077778, lng: -6.231389},
              name: "Roundwood Flying Club",
              rad: 800,
              color: "#FBCEB1"
            }
      
          }
          

          // const smallerAirports = {
          //   waterford: {
          //     center: { lat: 53.4264, lng: -6.2499 }
          //   },
          //   donegal: {
          //     center: { lat: 52.6996575, lng: -8.9806044 }
          //   },
          //   crk: {
          //     center: { lat: 51.8490624, lng: -8.4920732 }
          //   },
          // };



          for (const airport in airports) {
              const cityCircle = new google.maps.Circle({
              strokeColor: "#FF0000",
              strokeOpacity: 0.8,
              strokeWeight: 2,
              fillColor: airports[airport].color,
              fillOpacity: 0.35,
              map: map.value,
              center: airports[airport].center,
              radius: airports[airport].rad
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
