const routers = [
    {
        path: '/schedule/job',
        name: 'ScheduleJob',
        meta: {
            title: '作业详情'
        },
        component: (resolve) => require(['../views/schedule/job.vue'], resolve)
    },
    {
        path: '/schedule/instance',
        name: 'ScheduleInstance',
        meta: {
            title: '实例详情'
        },
        component: (resolve) => require(['../views/schedule/instance.vue'], resolve)
    },
    {
        path: '/schedule/analysis',
        name: 'ScheduleAnalysis',
        meta: {
            title: '调度分析'
        },
        component: (resolve) => require(['../views/schedule/analysis.vue'], resolve)
    },
];
export default routers;