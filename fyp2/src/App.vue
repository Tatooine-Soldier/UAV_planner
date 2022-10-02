<script setup>
  import NavComponent from './components/NavComponent.vue';
  import HomeComponent from './components/HomeComponent.vue';
  import Vue from 'vue';
</script>

<template >
  <section id="app">
    <NavComponent/>

    <section class="main-container">
      <main>
        <HomeComponent />
        
        <!-- <router-view/>
        <router-link to="/login">Router Link 1</router-link> -->
      </main>
    </section>
  </section>
</template>

<style scoped>
    #app {
      background-color: rgb(57, 56, 56);
    }

    main {
      margin-left: 7%;
      margin-right: 8%;
    }

</style>

<script type="text/javascript">
  // export default {
  //     name: 'App',
  //  }
</script>
