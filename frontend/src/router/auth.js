const routers = [
    {
        path: '/auth/action',
        name: 'AuthAction',
        meta: {
            title: '操作日志'
        },
        component: (resolve) => require(['../views/auth/action.vue'], resolve)
    },
    {
        path: '/auth/group',
        name: 'AuthGroup',
        meta: {
            title: '组管理'
        },
        component: (resolve) => require(['../views/auth/group.vue'], resolve)
    },
    {
        path: '/auth/user',
        name: 'AuthUser',
        meta: {
            title: '用户管理'
        },
        component: (resolve) => require(['../views/auth/user.vue'], resolve)
    }
];
export default routers;