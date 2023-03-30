<template>
    <section class="profile-page">
        <section class="profile-main">
            <p>Profile Page</p>
            <h2>Welcome {{ user }}</h2>
            <div>
                <p>Past Flights: </p>
                <p>{{ flightData }}</p>
            </div>
        </section>
    </section>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            user: "",
            flightData: []
        }
    },
    async created() {
        axios
        .post("/userProfile")
        .then((response) => {
          const data = response.data;
          console.log("data: ", data)
          var dataArray = data.split("|")
          console.log("dataArray--> ", dataArray)
          this.user = dataArray[0]
          this.flightData = dataArray[1]
        })
        .catch(function (error) {
            console.log("ERROR:", error); 
        })
    }
}
</script>

<style>
.profile-main {
    color: white;
}
</style>