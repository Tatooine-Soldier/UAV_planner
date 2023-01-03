



<template>
    <GoogleMap id="mymap" v-on:click="createMarker()"
      api-key="AIzaSyDTNOMjJP2zMMEHcGy2wMNae1JnHkGVvn0"
      style="width: 100%; height: 500px"
      :center="center"
      :zoom="25"
    >
      <Marker :options="{ position: center}" />
    </GoogleMap>

  </template>
  
  <script>
  import { defineComponent } from "vue";
  import { GoogleMap, Marker } from "vue3-google-map";
  
  export default defineComponent({
    components: { GoogleMap, Marker },
    setup() {
      const center = { lat: 53.689247, lng: -7.044502 };
      return { center };
    },
    methods: {
        createMarker() {
            console.log("CREATING")
            const position = { lat: 0.0, lng: 0.0 };
            var google;
            const marker = new google.maps.marker.AdvancedMarkerView({
                position: position,
                draggable: true,
                title: "This marker is draggable.",
            });
            marker.addListener("click", (event) => {
                const position = marker.position;
                console.log(position.lat, position.lng)
                console.log(event)
            });
            marker.addListener('drag', (event) => {
                const position = marker.position;
                console.log(position.lat, position.lng)
                console.log(event)
            });
            return marker;
        }
    }
    // getPosition(position) {
    //     console.log(position.coords.latitude, position.coords.longitude);
    // }
  });
  </script> 
