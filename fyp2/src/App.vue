<script setup>
  import NavComponent from './components/NavComponent.vue';
  import HomeComponent from './components/HomeComponent.vue';
  import LoginComponent from './components/LoginComponent.vue';
  import FooterComponent from './components/FooterComponent.vue';
  import { Vue } from "vue";
  import { VueRouter } from 'vue-router';
 
  Vue.use(VueRouter);
</script>

<template >
  <section id="app">
    <NavComponent/>

    <section class="main-container">
      <main>
        <!-- <HomeComponent /> -->
        <LoginComponent/>
        <router-link to="/">Go Home</router-link>
        <router-view></router-view>
      </main>
    </section>
    <FooterComponent/>
  </section>
</template>

<style scoped>
    #app {
      background-color: rgb(57, 56, 56);
      height: auto;
    }

    main {
      margin-left: 7%;
      margin-right: 8%;
    }

</style>

<script type="text/javascript">
  // export default {
  //     name: 'App',
  //  }#
  const routes = [
   {
      path: '/',
      name: 'home',
      component: HomeComponent
   }];

  const router = new VueRouter
</script>
