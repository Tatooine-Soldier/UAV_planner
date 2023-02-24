<script>
import { Loader } from '@googlemaps/js-api-loader'
    /* eslint-disable no-undef*/
     /* eslint-disable no-unused-vars*/
    import { computed, ref, onMounted, onUnmounted, watch } from 'vue'
    import { Graph, Node } from '../graph'
    import { Grid } from '../gridCoords'
    //import { useGeolocation } from '../useGeolocation'
    // import haversineDistance from './calculateDistance'
    const GOOGLE_MAPS_API_KEY = 'AIzaSyDTNOMjJP2zMMEHcGy2wMNae1JnHkGVvn0'
    export default {
      name: 'App',
      props: ['propcoords', 'propspeed', 'propdate', 'propway', 'propEndTime', 'propDuration'],
      data: function() {
        var myData = {
            myProp: this.propcoords
        }
        return myData
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

        var cpos = {lat: anchors.startLat, lng: anchors.startLng}
        var opos = {lat: anchors.endLat, lng: anchors.endLng}
        var grid  = new Grid(3); //pass currPos and otherLoc down to grid, get nearest nodes in graph for both and then use those nodes in Dijkstra
        var psos = grid.generateCoords([[{lat: 53.531386134765576, lng: -7.925040162129355}]], true, anchors).then(data => { 
        console.log("Received In FINAL map coords--->", data); 
        var totalDist = 0
        var l = data[0].path
        console.log("path--> ", l, typeof l)
        for (var i = 1; i < l.length; i++) {
          var latN = parseFloat(l[i].value.coordinate.lat)
          var lngN = parseFloat(l[i].value.coordinate.lng)
          var latS = parseFloat(l[i-1].value.coordinate.lat)
          var lngS = parseFloat(l[i-1].value.coordinate.lng)
          var north = {lat: latN, lng: lngN}
          var south = {lat: latS, lng: lngS}
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
        totalDist = 6.46*(l.length-1)
        console.log("total distance-->", totalDist)
        var domDist = document.getElementById("dist")
        domDist.innerHTML = totalDist.toString()
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

      
        // need to check where data is being sent to Final Map component, do i just pass path as a Prop up to planner like all the other data?
        })
          .catch(error => {
          console.error(error);
        }); //centerPoint for Ireland grid

        console.log("psos: ", psos)
        return { currPos, otherLoc, distance, mapDivHere, calculatedTime }
      }
    }


</script>

<template>
  <div id="big-container">
    <div class="final-distance-caption-container">
      <div>
        <br>Take-off Time:: {{ propdate.day }}, {{ propdate.hour }}:{{ propdate.minute }}:00
        <br>Arrival Time: {{ propEndTime }}
        </div>
      <div>Corridor: {{ propspeed.description }} <br>Distance of path(km): <div id="dist">{{ distance }}</div><br>Flight Duration: {{ calculatedTime }} </div>
      <div class="detailz">
            Starting Point:  <i id="plat">{{ propcoords.sourcelatitude }}</i> <i id="plng">{{ propcoords.sourcelongitude }}<br></i>
            Destination Point:  <i id="plat">{{ propcoords.destlatitude }}</i> <i id="plng">{{ propcoords.destlongitude }}<br></i>
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

  .details i {
    font-size: 8px;
  }

  #plng {
    font-size: 8px;
  }

  #plat {
    font-size: 8px;
  }
</style>