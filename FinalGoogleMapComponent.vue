<script>
import { Loader } from '@googlemaps/js-api-loader'
import axios from 'axios'

    /* eslint-disable no-undef*/
     /* eslint-disable no-unused-vars*/
    import { computed, ref, onMounted, onUnmounted, watch } from 'vue'
    import { Graph, Node } from '../graph'
    import { Grid } from '../gridCoords'
    //import { useGeolocation } from '../useGeolocation'
    // import haversineDistance from './calculateDistance'
    //const GOOGLE_MAPS_API_KEY = 'AIzaSyDTNOMjJP2zMMEHcGy2wMNae1JnHkGVvn0' //real key
    const GOOGLE_MAPS_API_KEY = 'AIzaSyBkU3LEkHvrO8_kpSWGqobpFob-sESKlA8'
    export default {
      name: 'App',
      props: ['propcoords', 'propspeed', 'propdate', 'propway', 'propEndTime', 'propSubgrid', 'propDuration', 'propID'],
      data() {
        return {
          counter: 0
        }
        
      },
      methods: {
        showLog() {
          var t = document.getElementById("timestamps")
          var f = document.getElementById("flightlogbutton")
          if (this.counter % 2 == 0) {
            t.style.display = "block"
            f.value = "Hide LOG"
          } else {
            t.style.display = "none"
            f.value = "Show LOG"
          }
          this.counter += 1
        }

      },
      setup(props) {
        watch(() => props.propcoords, () => {
          console.log("***PROPS***",props.propcoords)
        })
        const currPos = ref(null)
        const otherLoc = ref(null)
        const waypointLoc =  ref(null)
        let clickListener = null;

        const loader = new Loader({ apiKey: GOOGLE_MAPS_API_KEY})
        const mapDivHere = ref(null);
      
        let sourceMarker = ref(null)
        let destMarker = ref(null)
        let map = ref(null)

        //let waypointMarker =  ref(null)
        var graph; 
        var destNode;
        var sourceNode;
        var timestamps = []
       

        var anchors = {startLat: parseFloat(props.propcoords.sourcelatitude) , startLng: parseFloat(props.propcoords.sourcelongitude), endLat: parseFloat(props.propcoords.destlatitude), endLng: parseFloat(props.propcoords.destlongitude)}

        onMounted(async () => {
          await loader.load() 
          console.log("PROPS props.propcoords--->", props.propcoords)

          graph =  new Graph()
          destNode =  new Node()
          sourceNode = new Node()
          

          currPos.value = {lat: parseFloat(props.propcoords.sourcelatitude) ,lng: parseFloat(props.propcoords.sourcelongitude)} 
          otherLoc.value = {lat: parseFloat(props.propcoords.destlatitude), lng: parseFloat(props.propcoords.destlongitude)}
          
          sourceNode.value = currPos.value
          destNode.value =  otherLoc.value
          graph.add(sourceNode)
          graph.add(destNode)

          console.log("Propway", props.propway)

          waypointLoc.value =  {lat: parseFloat(props.propway.lat), lng: parseFloat(props.propway.lng)}


          map.value = new google.maps.Map(mapDivHere.value, {
                center: currPos.value,
                mapId: 'b6fdfcadbdc5d7a0', 
                zoom: 9 
            })
          sourceMarker.value = new google.maps.Marker({
            position: currPos.value,
            map: map.value
          })
          destMarker.value = new google.maps.Marker({
            position: otherLoc.value,
            map: map.value
          })
          // waypointMarker.value = new google.maps.Marker({
          //   position: waypointLoc.value,
          //   map: map.value
          // })
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
        })
        onUnmounted(async () => {
            if (clickListener) clickListener.remove()
        })
        //let line = null
        // watch([map, currPos, otherLoc], () => {
        //   if (line) line.setMap(null)
        //   if (map.value && otherLoc.value != null)
        //     line = new google.maps.Polyline({
        //       path: [currPos.value, otherLoc.value],
        //       map: map.value
        //     })
        // })
        let sourceToWay =  null; //line
        watch([map, currPos, waypointLoc], () => {
          if (sourceToWay) sourceToWay.setMap(null)
          if (map.value && waypointLoc.value != null)
            sourceToWay = new google.maps.Polyline({
              path: [currPos.value, waypointLoc.value],
              map: map.value
            })
        })

        let wayToDest = null; //line
        watch([map, waypointLoc, otherLoc], () => {
          if (wayToDest) wayToDest.setMap(null)
          if (map.value && waypointLoc.value != null)
            wayToDest = new google.maps.Polyline({
              path: [waypointLoc.value, otherLoc.value],
              map: map.value
            })
        })

        

        const haversineDistance = (pos1, pos2) => {
          console.log("POS-->", pos1, pos2)
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
        const t = (duration) => {
          duration = duration.toFixed(2)
          var hours = duration / 1
          var remainder =  hours % 1
          hours = Math.round(hours)
          var minutes = remainder * 60
          minutes = Math.round(minutes)
          var time = hours.toString() + " hours " + minutes.toString() + " minutes"
          return time
        }
        const calculatedTime = computed(() =>
            t(props.propDuration)
        )
        const distance = 0;
        // const distance = computed(() =>
        // otherLoc.value === null && currPos.value === false
        //   ? 0
        //   : haversineDistance(currPos.value, otherLoc.value)
        // )

        var calendarContainer = document.getElementById("calendar-display-afterwards")
        var footer = document.getElementById("footerApp")
        var loadingScreen = document.getElementById("loadingScreen")

        var segments = [] //used to capture segments(grid points) in the flight path
        var cpos = {lat: anchors.startLat, lng: anchors.startLng}
        var opos = {lat: anchors.endLat, lng: anchors.endLng}
        var grid  = new Grid(4); //pass currPos and otherLoc down to grid, get nearest nodes in graph for both and then use those nodes in Dijkstra
        var psos = grid.generateCoords([[{lat: 51.8964507, lng: -8.4908813}]], true, anchors).then((data) => { 
        console.log("Received In FINAL map coords--->", data); 
        var totalDist = 0
        console.log("data[0].path ", data[0].path)
        setTimeout(function() {
          console.log("W", data[0].path);
        }, 5000);
        var l;
        l = data[0].path
        console.log("l: ", l[0])
        console.log("propcoords===>", props.propcoords) //need to add start and end points to segments
        segments.push({lat: props.propcoords.destlatitude, lng: props.propcoords.destlongitude})
        segments.push({lat: l[0].value.coordinate.lat, lng: l[0].value.coordinate.lng})
        console.log("path--> ", l, typeof l)
        for (var i = 1; i < l.length; i++) {
          var latN = parseFloat(l[i].value.coordinate.lat)
          var lngN = parseFloat(l[i].value.coordinate.lng)
          var latS = parseFloat(l[i-1].value.coordinate.lat)
          var lngS = parseFloat(l[i-1].value.coordinate.lng)
          var north = {lat: latN, lng: lngN}
          var south = {lat: latS, lng: lngS}
          segments.push({lat: l[i].value.coordinate.lat, lng: l[i].value.coordinate.lng})
          console.log("Pairs as floats : ",latN, lngN, latS, lngS )
          console.log("source lats: ", currPos.value.lat) //want to find the nearest node in th graph and connect it to the starting point 
          var line = null;
          if (line) line.setMap(null)
          if (north && south != null)
            line = new google.maps.Polyline({
              path: [north, south],
              map: map.value,
              strokeColor: "#FF0000"
            })
        // const d = computed(() => haversineDistance(north, south)) //want to add up distances between each point in the grid path
        // totalDist += parseFloat(totalDist) + parseFloat(d.value)
        // console.log("totalDist-->", totalDist)
        }
        segments.push({lat: props.propcoords.sourcelatitude, lng: props.propcoords.sourcelongitude})
        totalDist = 6.46*(l.length-1)
        console.log("total distance-->", totalDist)
        // var domDist = document.getElementById("dist")
        // domDist.innerHTML = totalDist.toString()
        const gridEntryCircle = new google.maps.Circle({
                  strokeColor: "#FF1122",
                  strokeOpacity: 0.8,
                  strokeWeight: 1,
                  fillColor: "#FF1122",
                  fillOpacity: 0.7,
                  map: map.value,
                  center: {lat: parseFloat(l[0].value.coordinate.lat), lng: parseFloat(l[0].value.coordinate.lng)},
                  radius: 500
              });
        const gridExitCircle = new google.maps.Circle({
                  strokeColor: "#FF1122",
                  strokeOpacity: 0.8,
                  strokeWeight: 1,
                  fillColor: "#FF1122",
                  fillOpacity: 0.7,
                  map: map.value,
                  center: {lat: parseFloat(l[l.length-1].value.coordinate.lat), lng: parseFloat(l[l.length-1].value.coordinate.lng)},
                  radius: 500
                });
              
          var entry =  {lat: parseFloat(data[1].start.lat), lng: parseFloat(data[1].start.lng)}
          var exit = {lat: parseFloat(data[1].end.lat), lng: parseFloat(data[1].end.lng)}
          console.log("start and entry", entry, exit)

          line = null;
          if (line) line.setMap(null)
          if (data[1].start != null)
            line = new google.maps.Polyline({
              path: [entry, cpos],
              map: map.value,
              strokeColor: "#1133FF"
            })

          line = null;
          if (line) line.setMap(null)
          if (data[1].start != null)
            line = new google.maps.Polyline({
              path: [exit, opos],
              map: map.value,
              strokeColor: "#1133FF"
            })
          
          
          var startTime = props.propdate.hour + props.propdate.minute
          var endTime = props.propEndTime
          var speed =  props.propspeed
          var duration = props.propDuration
        

          var segmentedTime =  duration.toFixed(2) / l.length
          console.log("SEGMENTED TIME & DURATION", segmentedTime, duration.toFixed(2))
          var segmentedTimeList = []
          console.log(" props.propdate.hour.  props.propdate.minute ",  props.propdate.hour,   props.propdate.minute)
          segmentedTimeList.push({hour: props.propdate.hour, minute: props.propdate.minute})
          const dateObj = new Date();
          dateObj.setHours(props.propdate.hour)
          dateObj.setMinutes(props.propdate.minute)
          var j = 1
          while(j < l.length+1) {
            var tl = (segmentedTime*60) 
            //convert t to an actual time(add it to start time )
            //t = t * 60
         
            if (tl < 60) {
              console.log("adding ", tl, " minutes")
              dateObj.setMinutes(dateObj.getMinutes() + tl);
            } else {
              dateObj.setHours(dateObj.getHours() + 1);
              var r = tl % 60
              dateObj.setMinutes(dateObj.getMinutes() + r);
            }
            console.log("dateObj", dateObj, "\nt", tl) 
            if (dateObj.getHours().toString().length < 2) {
              segmentedTimeList.push({hour: "0"+dateObj.getHours().toString(), minute: dateObj.getMinutes().toString()})
            } else {
              segmentedTimeList.push({hour: dateObj.getHours().toString(), minute: dateObj.getMinutes().toString()})
            }
            j++
          }
          var e = props.propEndTime.toString()
          segmentedTimeList.push({hour:  e.slice(12, 14), minute:e.slice(15, 17)})
          console.log("segmentedTimeList", segmentedTimeList)
          segmentedTimeList = segmentedTimeList.reverse()
          


          //code to adjust the timing segments to include No Man's Land
          console.log("** HAVERSINE between first two segments --> **", haversineDistance(segments[segments.length-1], segments[segments.length-2]), segments[segments.length-1], segments[segments.length-2])
          var cdistance =  haversineDistance(segments[segments.length-1], segments[segments.length-2])
          var ctime =  (cdistance*1000) / ((speed*1000)/3600)
          console.log("ctime raw: ", ctime, cdistance*1000)
          function convertToMinutes(time) {
            var t = Math.ceil(time)
            var minutes =  t/60
            return minutes
          } 
          ctime = convertToMinutes(ctime)
          ctime = Math.ceil(ctime)
          console.log("ctime---> ", ctime) //minutes to be added on to each segmented time 
          function addMinutes(time, minutesToAdd) { //
            console.log("Adding ",minutesToAdd, " to ",  time )
            var hours = parseInt(time.hour);
            var minutes = parseInt(time.minute);

            var newMinutes = minutes + minutesToAdd;
            console.log(minutes, minutesToAdd)
            var newHours = hours
            if (newMinutes >= 60) {
              newMinutes = newMinutes % 60;
              newHours = hours + 1
            } 
            newMinutes =  Math.ceil(newMinutes) 
            var formattedTime = {hour: newHours.toString(), minute: newMinutes.toString()}
            return formattedTime
          }
       
          var newTime = addMinutes( segmentedTimeList[segmentedTimeList.length-1],ctime)
          console.log(newTime) //newTime is the time at the first grid coord(start location time + distance to closest grid coord)

          const dateOb = new Date();
          dateOb.setHours(newTime.hour)
          dateOb.setMinutes(newTime.minute)
          
          segmentedTimeList[1] = newTime
          segmentedTimeList[0] = segmentedTimeList[segmentedTimeList.length-1]
          console.log("SEGMENTEDTIMELIST before reverse: ", segmentedTimeList)
          for (j=2;j<segmentedTimeList.length; j++) {
            console.log("List item time-->", segmentedTimeList[j], "time chunk ", segmentedTime*60)
            tl = (segmentedTime*60) 
            if (tl < 60) {
              dateOb.setMinutes(dateOb.getMinutes() + tl);
            } else {
              dateOb.setHours(dateOb.getHours() + 1);
              r = tl % 60
              dateOb.setMinutes(dateOb.getMinutes() + r);
            }
            if (dateOb.getHours().toString().length < 2) {
              segmentedTimeList[j] = {hour: "0"+dateOb.getHours().toString(), minute: dateOb.getMinutes().toString()}
            } else {
              segmentedTimeList[j] = {hour: dateOb.getHours().toString(), minute: dateOb.getMinutes().toString()}
            }
          
          }
          segmentedTimeList = segmentedTimeList.reverse()
          updateFlightTimes(segmentedTimeList[segmentedTimeList.length-1], segmentedTimeList[0], speed)
          displayTimestamps(segmentedTimeList, segments, timestamps)

          console.log("SEGMENTEDTIMELIST after reverse: ", segmentedTimeList)
          //**** RETURN HERE THE SEGMENTS//


          //Timestamp Display the circle objects with info objects for each node timestamp
          console.log("segments", segments, "\nlen(segments): ", l.length, "\nduration: ", props.propDuration)

          for (var s in segments) {
            var lngSeg = parseFloat(segments[s].lng)
            var latSeg = parseFloat(segments[s].lat)
            const segCircle = new google.maps.Circle({
                  strokeColor: "#FF1122",
                  strokeOpacity: 0.8,
                  strokeWeight: 1,
                  fillColor: "#FF1122",
                  fillOpacity: 1,
                  map: map.value,
                  center: {lat: latSeg, lng: lngSeg},
                  radius: 400
              })
            var timeD = segmentedTimeList[s].hour.toString() + ":" + segmentedTimeList[s].minute.toString()
            
            const infowindow = new google.maps.InfoWindow({
              content: timeD,
              ariaLabel: "Time_At_Coordinate",
            });
            segCircle.addListener("click", () => {
              if (segmentedTimeList[s].minute.length < 2 ) {
                segmentedTimeList[s].minute = "0"+segmentedTimeList[s].minute.toString()
              }
              if (segmentedTimeList[s].hour.length < 2 ) {
                segmentedTimeList[s].hour = "0"+segmentedTimeList[s].hour.toString()
              }
              console.log("clicked circle", segmentedTimeList[s].hour.toString(), segmentedTimeList[s].minute.toString(), s)
              infowindow.setPosition(segCircle.center)
              infowindow.open({
                anchor: segCircle,
                map,
              });
              
            })
          }
      
          var day;
          if (props.propdate.day === "") {
            day = "0"
          } else {
            day = props.propdate.day
          }

          for (var z in segmentedTimeList) {
            if (segmentedTimeList[z].hour.length !== 2) {
              segmentedTimeList[z].hour = "0"+segmentedTimeList[z].hour
            }
            if (segmentedTimeList[z].minute.length !== 2) {
              segmentedTimeList[z].minute = "0"+segmentedTimeList[z].minute
            }
            console.log("fixing time", segmentedTimeList[z].hour, segmentedTimeList[z].hour)
          }

          var entryPoint = {
            id: props.propID,
            lat: String(entry.lat),
            lng: String(entry.lng)
          }
          var exitPoint = {
            id: props.propID,
            lat: String(exit.lat),
            lng: String(exit.lng)
          }
          var segementedFlight =  { //need to store the times with segment each too
           
            date: day, 
            segmentList: segments,
            segmentTimes: segmentedTimeList.reverse(),
            subGrid: props.propSubgrid, 
            speed: props.propspeed,
            entryPoint: entryPoint, //entry point to the grid 
            exitPoint: exitPoint, //exit point from the grid
            id: props.propID
          }
          //STORE SEGMENTS LIST AS A SINGLE RECORD IN SEGMENTS COLLECTION WITH THE ID OF THE FLIGHT
          axios
          .post("/storeSegmentedFlight", segementedFlight)
          .then((response) => {
            console.log("* stored Segmented flights!!! *")
            //calendarContainer.style.display = "block"; //UNCOMMENT THIS
          //  footer.style.display = "inline-block";
          loadingScreen.style.display = "none"
            var fmap = document.getElementById("final-map-container")
            fmap.style.display = "block"
            var con = document.getElementById("ex-sign")
            con.style.display = "block"
            var details = document.getElementById('flight-details');
            details.style.display = 'None';
          })
          .catch(error => {
          console.error(error);
        }); 

        

        // NEED TO SEGMENT THE FLIGHT BY GETTING THE COORDINATES AT CERTAIN POINTS AND TIME IN THE FLIGHT(PERHAPS JUST GET THE COORDINATES OF EACH GRUD POINT IN THE PATH)
        //THEN CHECK IF THE INTENDED FLIGHT FALLS WITHIN RADIUS OF THESE POINTS AT A THAT INTENDED TIME
      }).catch((error) => {
        console.log("Error with l", error)
      }) 



        console.log("psos: ", psos)
        return { currPos, otherLoc, distance, mapDivHere, calculatedTime, timestamps}
      }
    }
    function displayTimestamps(segmentedTimeList, segments, timestamps) {
      console.log("segments==>", segments)
      var timestampDisplay =  document.getElementById("timestamps")
      timestampDisplay.innerHTML = "<ul>"
      for (var i=0; i<segmentedTimeList.length;i++) {
        console.log("pushing segmentedTimeList[i] and segments[i]", segmentedTimeList[i], segments[i])
        if (segmentedTimeList[i].minute.length === 1) {
          segmentedTimeList[i].minute = "0"+segmentedTimeList[i].minute
        }
        timestampDisplay.innerHTML += "<li>" + segmentedTimeList[i].hour+":"+segmentedTimeList[i].minute +" ---- <i>" + segments[i].lat + " " + segments[i].lng + "</i></li><hr>"
        timestamps.push({time: segmentedTimeList[i], coord: segments[i]})
      }
      timestampDisplay.innerHTML += "</ul>"
    }

    function updateFlightTimes(start, eta, speed) {
      var r =  document.getElementById("take-off-time")
      r.innerHTML = start.hour + ":" + start.minute
      var f = document.getElementById("eta-final")
      f.innerHTML += " "+"<b>"+eta.hour +":"+eta.minute+"</b>"
      var s = document.getElementById("speed-final")
      s.innerHTML = speed
    }
    


</script>

<template>
  <div id="big-container">
    <div class="final-distance-caption-container">
      <div>
        <div>Flight Number: #<b>{{ propID }}</b></div>
        <div>
          Take-off time: <b id="take-off-time"></b>
        </div>
        <div>
          Altitude: <b id="take-off-altitude"></b>m
        </div>
        <div>
          Speed: <b id="speed-final"></b>km/h
        </div>
        <!-- <br>Take-off Time:: {{ propdate.day }}, {{ propdate.hour }}:{{ propdate.minute }}:00 -->
        <br>ETA: <b><span id="eta-final"> </span></b>
      </div>
      <!-- <div>Corridor: {{ propspeed.description }} <br>Distance of path(km): <div id="dist">{{ distance }}</div><br>Flight Duration: {{ calculatedTime }} </div> -->
      <div class="detailz">
            Starting Point:  <b id="plat">{{ propcoords.sourcelatitude }}</b> <b id="plng">{{ propcoords.sourcelongitude }}<br></b>
            Destination Point:  <b id="plat">{{ propcoords.destlatitude }}</b> <b id="plng">{{ propcoords.destlongitude }}<br></b>
            <!-- displayed as none for now -->
            Flight Log: <input type="button" value="Show LOG" @click="showLog()" id="flightlogbutton"> <div id="timestamps"> </div> 
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
    grid-template-columns: 50% 50%;
  }

  .detailsz {
    font-size: 10px;
    text-align: left;
  }

  .details i {
    font-size: 8px;
  }

  #plng {
    font-size: 8px;
  }

  #plat {
    font-size: 8px;
  }

  #timestamps {
    display: none;
    position:absolute;
    z-index: 1;
    background-color: white;
    border: .5px solid rgb(80, 79, 79);
    border-top: none;
    padding: 5px;
  }
</style>
