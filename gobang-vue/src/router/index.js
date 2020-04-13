import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
    routes: [{
        path: '/',
        redirect: '/login'
    }, {
        path: '/login',
        component: () => import('@/views/Login')
    }, {
        path: '/',
        component: () => import('@/views/Layout'),
        children: [{
            path: '/game',
            name: 'game',
            component: () => import('@/views/Game')
        }, {
            path: '/settings',
            name: 'settings',
            component: () => import('@/views/Settings')
        }, {
            path: '/about',
            name: 'about' ,
            component: () => import('@/views/About')
        }]
    }],
    mode: 'history'
})