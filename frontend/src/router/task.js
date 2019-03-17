const routers = [
    {
        path: '/task/listing',
        name: 'TaskListing',
        meta: {
            title: '任务详情'
        },
        component: (resolve) => require(['../views/task/listing.vue'], resolve)
    },
    {
        path: '/task/dependency',
        name: 'TaskDependency',
        meta: {
            title: '任务依赖分析'
        },
        component: (resolve) => require(['../views/task/dependency.vue'], resolve)
    },
];
export default routers;