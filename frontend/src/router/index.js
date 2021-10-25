import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routerOptions = [
  { path: '/', view: 'Meals' },
  { path: '/suggestions', view: 'Suggestions' },
]

const routes = routerOptions.map(route => {
  return {
    ...route,
    component: () => import(`@/views/${route.view}.vue`)
  }
})

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
