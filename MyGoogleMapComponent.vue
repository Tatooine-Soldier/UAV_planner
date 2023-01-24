<script>
import { Loader } from '@googlemaps/js-api-loader'
    /* eslint-disable no-undef*/
    /* eslint-disable no-unused-vars*/
    import { computed, ref, onMounted, onUnmounted, watch } from 'vue'
    import { useGeolocation } from '../useGeolocation'
    import { airports, getAirports } from '../airports'
    import { LinkedList, Node } from '../linkedList'
    import { checkD } from '@/withinAirspace';
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
        const waypointLoc = ref(null)
        let clickListener = null;
        let dragListener = null;
        let waypointDragListener = null;
        let waypointListener = null;
        let destDragListener = null;
        let cursor = {lat: 0.0, lng: 0.0}
        let oldMarkers = [];

        const loader = new Loader({ apiKey: GOOGLE_MAPS_API_KEY})
        const mapDivHere = ref(null);
      
        let sourceMarker = ref(null)
        let destMarker = ref(null)
        let waypoint = ref(null)
        let oldMarker = ref(null)
        let waypointMarker = ref(null)
        var waypointDiv;
        var waypointList = new LinkedList(sourceMarker, destMarker);
        let wl = [];
        var flightPlanCoordinates = [];
        const RED_ZONE = "Entering Red Zone airspace is prohibited by law";

        let map = ref(null)
        var circle;
        var mycircle;
        var circleRef =  ref(null);
        var destinfowindow;

        
        
        onMounted(async () => {
          await loader.load()
          currPos.value = {lat: initial.value.lat, lng: initial.value.lng}
          console.log(currPos.value)
          var circleList = [];
          function displayWarning() { 
              
              if (circleRef.value.result) {
                var doc = document.getElementById("locationWarning");
                doc.style.display = "block"; 
                var colorDiv = document.getElementById("colorAirspace");
                colorDiv.style.color = circleRef.value.color;
                var msgDiv = document.getElementById("airspaceMessage");
                var contactDiv = document.getElementById("contactMessage");
                msgDiv.style.color = "#FFFFFF";
                var markerDiv = document.getElementById("markerName");
                if (circleRef.value.color === "#FF8833") {
                  colorDiv.innerHTML = "AMBER";
                  markerDiv.innerHTML =  circleRef.value.marker;
                  msgDiv.innerHTML = "Please contact <span id='airName'>"+ circleRef.value.name + "</span> UAS zone authority before beginning your flight."; //Unmanned Aerial System
                  contactDiv.style.color = "#FFFFFF"
                  contactDiv.innerHTML = "Email: <span><i><a id='emailAirport' href = 'mailto: "+circleRef.value.contact+"'>"+ circleRef.value.contact +"</a></i></span>"; 
                  var email = document.getElementById("emailAirport");
                  email.style.color = "#FFFFFF"
                  email.style.fontSize = "12px";
                  var nameDiv = document.getElementById("airName");
                  nameDiv.style.color =  "#FF8833";
                } else if (circleRef.value.color === "#FF0000") {
                  colorDiv.innerHTML = "RED";
                  msgDiv.innerHTML = RED_ZONE;
                }
              }
            }

          map.value = new google.maps.Map(mapDivHere.value, {
            center: currPos.value,
            zoom: 9
          })
          clickListener = map.value.addListener(
            'click',
            ({latLng: {lat, lng}}) => 
              (otherLoc.value = {lat: lat(), lng: lng()},
              getAirports(), 
              destMarker.value = new google.maps.Marker({
                position: otherLoc.value,
                draggable: true,
                map: map.value
              }),
              destinfowindow = new google.maps.InfoWindow({
                content: "Destination Point",
                ariaLabel: "Destination",
              }),
              destMarker.value.addListener("click", () => {
                destinfowindow.open({
                    anchor: destMarker.value,
                    map,
                });
              }),
              destDragListener = destMarker.value.addListener(
                'drag',
                function(event) {
                  otherLoc.value = {lat: event.latLng.lat(), lng: event.latLng.lng()},
                  
                  circle = function() {
                    var airportsList =  getAirports();
                    for (var airspace=0; airspace<airportsList.length; airspace++) { //for circle on map
                        console.log("marker lat", otherLoc.value)
                        var distnace = haversineDistance(airportsList[airspace].center, otherLoc.value)
                        distnace = distnace*1000;     //convert to metres
                        if (distnace < (airportsList[airspace].rad)) {  // the radius was not accurately represented on the map so I multiplied by .6 to be more accuarte?????
                            console.log("*WITHIN RADIUS*:\n", "distance: ", distnace, "airport name and rad:", airportsList[airspace].name, airportsList[airspace].rad);
                            const ret = {
                                result: true, 
                                name: airportsList[airspace].name, 
                                color: airportsList[airspace].color,
                                marker: "Destination marker",
                                contact: airportsList[airspace].contact
                            }
                            circleRef.value =  ret
                            displayWarning()
                        } 
                    }
                  },
                  circle.call(),
                  console.log("OTHER-->",otherLoc.value.lat)
                  var d = document.getElementById('destCursorLat')
                  d.innerHTML = otherLoc.value.lat;
                  var c = document.getElementById('destCursorLng')
                  c.innerHTML = otherLoc.value.lng;
                }
              ),
              displayWaypointClicker(),
               waypointList = new LinkedList(sourceMarker.value, destMarker.value)
            )
          )


          waypointDiv = document.getElementById("addWaypoint");
          waypointListener = waypointDiv.addEventListener("click", function(){
            waypointLoc.value = {lat: map.value.center.lat(), lng: map.value.center.lng()}
            flightPlanCoordinates.push(waypointLoc.value)
            drawPath()
            console.log("flightPlanCoordinates::::",flightPlanCoordinates)
            waypointMarker.value = new google.maps.Marker({
              map: map.value,
              position: waypointLoc.value,
              draggable: true,
              icon: {
                url: "http://maps.google.com/mapfiles/ms/icons/blue-dot.png"
              }
            }),
            waypointDragListener = waypointMarker.value.addListener(
              'drag',
              function(e) {
                waypointLoc.value = {lat: e.latLng.lat(), lng: e.latLng.lng()}
                flightPlanCoordinates[flightPlanCoordinates.length - 1] = waypointLoc.value
                console.log("PATH", flightPlanCoordinates)
                //flightPlanCoordinates.push(waypointLoc.value)
                updatePath()
              },
              
              console.log("After dragging", flightPlanCoordinates)
            )

          
            function drawPath() {
                const flightPath = new google.maps.Polyline({
                path: flightPlanCoordinates,
                geodesic: true,
                strokeColor: "#FF0000",
                strokeOpacity: 0.5,
                strokeWeight: 3,
                });
                flightPath.setMap(map.value) 
                return flightPath
              }
            
            function updatePath() {
              const flightPath = new google.maps.Polyline({
                path: flightPlanCoordinates,
                geodesic: true,
                strokeColor: "#FF0000",
                strokeOpacity: 0.5,
                strokeWeight: 3,
                });
                flightPath.setMap(map.value) 
                return flightPath
            }

     
            //   if (waypointLinesD) waypointLinesD.setMap(null)
            // waypointLinesD = new google.maps.Polyline({
            //   path: [otherLoc.value, flightPlanCoordinates[flightPlanCoordinates.length-1]],
            //   map: map.value,
            //   strokeColor: "#FF0000",
            //   strokeOpacity: 0.5,
            // })
            console.log(flightPlanCoordinates)
          })

          sourceMarker.value = new google.maps.Marker({
            position: currPos.value,
            draggable: true,
            map: map.value
          })
          //add in info window that appears when clicked  
          const sourceinfowindow = new google.maps.InfoWindow({
            content: "Starting Point",
            ariaLabel: "Start",
          });
          sourceMarker.value.addListener("click", () => {
            sourceinfowindow.open({
                anchor: sourceMarker.value,
                map,
            });
          });
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
              circle = function() {
                  var airportsList =  getAirports();
                  for (var airspace=0; airspace<airportsList.length; airspace++) { //for circle on map
                        console.log("marker lat", currPos.value)
                        var distnace = haversineDistance(airportsList[airspace].center, currPos.value)
                        distnace = distnace*1000;     //convert to metres
                        if (distnace < (airportsList[airspace].rad)) {  // the radius was not accurately represented on the map so I multiplied by .6 to be more accuarte?????
                            console.log("*WITHIN RADIUS*:\n", "distance: ", distnace, "airport name and rad:", airportsList[airspace].name, airportsList[airspace].rad);
                            const ret = {
                                result: true, 
                                name: airportsList[airspace].name, 
                                color: airportsList[airspace].color,
                                marker: "Starting point",
                                contact: airportsList[airspace].contact
                            }
                            circleRef.value =  ret
                            displayWarning()
                        } 
                       
                    }
                  },
                circle.call()
            }
          )


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
            cityCircle.addListener("click", function() {
              console.log(airports[airport].name)
              var a = document.getElementById("airportClicked");
              a.style.display="block";
              a.innerHTML = airports[airport].name
              setTimeout(function() {
                a.style.display="none"; //fade out would be nice here
              },1300)
            })
            circleList.push(cityCircle);
          }
          //do smaller airports next, 10 nautical mile radius, the 3 larger airports hav an extended area outside area already given so add it in
          //add onclick listener to the circle
         
          function check(marker, circle, radius) 
          {  
              var km = radius/1000;
              var kx = Math.cos(Math.PI * circle.lat / 180) * 111;
              var dx = Math.abs(circle.center.lng() - marker.lng) * kx;
              var dy = Math.abs(circle.center.lat() - marker.lat) * 111;  
              return Math.sqrt(dx * dx + dy * dy) <= km;
          }

          function displayWaypointClicker() {
            var t = document.getElementById("addWaypoint")
            t.style.display = "block";
          }

          
        
        })
        onUnmounted(async () => {
            if (clickListener) clickListener.remove()
            if (dragListener) dragListener.remove()
            if (destDragListener) destDragListener.remove()
            if (waypointDragListener) waypointDragListener.remove()
        
        })
        let line = null //source to dest
        watch([map, currPos, otherLoc], () => {
          if (line) line.setMap(null)
          if (map.value && otherLoc.value != null)
            line = new google.maps.Polyline({
              path: [currPos.value, otherLoc.value],
              map: map.value
            })
          
        })

        

        let waypointLines = null //source to waypoint
        watch([map,currPos, waypointLoc], () => {
          if (waypointLines) waypointLines.setMap(null)
            waypointLines = new google.maps.Polyline({
              path: [currPos.value, flightPlanCoordinates[0]],
              map: map.value,
              strokeColor: "#FF0000",
              strokeOpacity: 0.5,
            })
        })
        // flightPlanCoordinates[flightPlanCoordinates.length-1]
        let waypointLinesD = null  //waypoint to dest
        watch([map,otherLoc, waypointLoc], () => {
          if (waypointLinesD) waypointLinesD.setMap(null)
            waypointLinesD = new google.maps.Polyline({
              path: [otherLoc.value, flightPlanCoordinates[flightPlanCoordinates.length-1]],
              map: map.value,
              strokeColor: "#FF0000",
              strokeOpacity: 0.5,
            })
        }) 

     

       

        


        // let waypointLineZ = null   //between final marker in array and destMarker
        // watch([map,otherLoc, wl[wl.length-1]], () => {
        //   if (waypointLineZ) waypointLineZ.setMap(null)
        //   waypointLineZ = new new google.maps.Polyline({
        //       path: [otherLoc.value, wl[wl.length-1]],
        //       map: map.value,
        //       strokeColor: "#FF0000",
        //       strokeOpacity: 0.5,
        //     })
        // })

        // let waypointLineA = null   //between final marker in array and destMarker
        // watch([map, currPos, wl[0]], () => {
        //   if (waypointLineA) waypointLineA.setMap(null)
        //   waypointLineA = new google.maps.Polyline({
        //       path: [currPos.value, wl[0]],
        //       map: map.value,
        //       strokeColor: "#FF0000",
        //       strokeOpacity: 0.5,
        //     })
        // })


        // let waypointPath = null 
        // watch([map, wl[0], wl[wl.length-1]], () => {
        //   if (waypointPath) waypointPath.setMap(null)
        //   waypointPath = new google.maps.Polyline({
        //       path: wayPositionsList,
        //       map: map.value,
        //       strokeColor: "#FF0000",
        //       strokeOpacity: 0.5,
        //     })
        // })

        // var wayPositionsList = function() {
        //   var item
        //   var markerList = []
        //   for (item in wl) {
        //     markerList.push(wl[item].value.position)
        //   } 
        //   return markerList
        // }
     

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
      methods: {
        addWaypoint() {


        }
      }
    
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
    <div id="airportClicked"></div>
    <div id="addWaypoint">Add Waypoint</div>
    <div id="locationWarning">
      <span id="markerName"></span> is within a <span id="colorAirspace"></span> area.
      <br><br>
      <div id="airspaceMessage"></div>
      <br><br>
      <div id="contactMessage"></div>
    </div>
  </div>

</template>

<style>
  #locationWarning {
    display: none;
    color: white;
    background-color: rgb(101, 100, 100);
    border: solid 1px rgb(101, 100, 100);
    border-radius: 5px;
    box-shadow: 0px 1px 3px #576481;
    position: absolute;
    top: 42%;
    margin-left: 1.2%;
    width: 18.5%;
    padding: 7px;
  }
  

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

  #airportClicked {
    position: absolute;
    z-index: 1;
    color: white;
    font-size: 1em;
    width: 25%;
    left: 37%;
    top: 20%;
    text-align: center;
    background-color: rgb(101, 100, 100);
    padding: 10px;
    border: solid 1.5px grey;
    border-radius: 20px;
    opacity: .9;
    display: none;
  }

  #addWaypoint {
    background-color: white;
    color: black;
    width: 19%;
    position: absolute;
    top: 35%;
    margin-left: 1.2%;
    text-align: center;
    padding: 5px;
    border: solid .1px #9a9a9a;
    border-radius: 3px;;
    box-shadow: 0px 1px 3px #576481;
    display: none;
  }

  #addWaypoint:hover {
    background-color: rgb(235, 234, 234);
    cursor: pointer;
  }
</style>
