const routers = [
    {
        path: '/machine/listing',
        name: 'MachineListing',
        meta: {
            title: '机器详情'
        },
        component: (resolve) => require(['../views/machine/listing.vue'], resolve)
    },
    {
        path: '/machine/status_analysis',
        name: 'MachineStatusAnalysis',
        meta: {
            title: '机器状态分析'
        },
        component: (resolve) => require(['../views/machine/status_analysis.vue'], resolve)
    },
];
export default routers;