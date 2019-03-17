import auth from './auth'
import machine from './machine'
import schedule from './schedule'
import task from './task'

let children = [...auth, ...machine, ...schedule, ...task];

const routers = [
    {
        path: '/',
        meta: {
            title: 'Chin'
        },
        component: (resolve) => require(['../views/index.vue'], resolve),
        children: children
    }
];
export default routers;