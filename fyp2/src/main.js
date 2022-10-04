import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
// import router from './router'
// import VueRouter from 'vue-router'
// import Vue from 'vue'
import HomeComponent from './components/HomeComponent.vue'
import LoginComponent from './components/LoginComponent.vue'
// Vue.use(VueRouter)

import './assets/main.css'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', name: 'home', component: HomeComponent},
        {path: '/login', name: 'login', component: LoginComponent}
    ]
})

createApp(App).use(router).mount('#app')
//createApp(App).use(router).mount('#app')

// const routes = [
//     {path: '/', component: HomeComponent}
// ]

// const router = new VueRouter({
//     routes
// })

// new Vue({
//     el: '#app',
//     router,
//     render: h => h(App)
// })