<template>
    <section class="profile-page">
        <section class="profile-main">
            <p>Profile Page</p>
            <h2>Welcome {{ user }}</h2>
        </section>
    </section>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            user: ""
        }
    },
    async created() {
        axios
        .post("/userProfile")
        .then((response) => {
          const data = response.data;
          console.log("data: ", data)
          this.user = data
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