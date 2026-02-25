export default [
  {
    path: '/',
    name: 'home',
    component: () => import('@/pages/HomePage.vue')
  },
  {
    path: '/items',
    name: 'items',
    component: () => import('@/pages/ItemsPage.vue')
  }
]
