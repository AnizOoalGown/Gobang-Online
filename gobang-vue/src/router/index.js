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
            name: 'Game',
            component: () => import('@/views/Game')
        }, {
            path: '/settings',
            name: 'Settings',
            component: () => import('@/views/Settings')
        }, {
            path: '/about',
            name: 'About',
            component: () => import('@/views/About')
        }]
    }],
    mode: 'history'
})