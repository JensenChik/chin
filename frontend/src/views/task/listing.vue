<template>
    <div>
        <Breadcrumb :style="{margin: '16px 0'}">
            <BreadcrumbItem>Home</BreadcrumbItem>
            <BreadcrumbItem>任务管理</BreadcrumbItem>
            <BreadcrumbItem>任务详情</BreadcrumbItem>
        </Breadcrumb>
        <Card style="height: 600px">
            <Table :columns="columns" :data="tasks"></Table>
        </Card>
    </div>

</template>

<script>
    import TaskExpand from './task_expand.vue';

    export default {
        components: {TaskExpand},
        data() {
            return {
                columns: [
                    {
                        type: 'expand',
                        width: 50,
                        render: (h, params) => {
                            return h(TaskExpand, {
                                props: {
                                    row: params.row
                                }
                            })
                        }
                    },
                    {
                        title: '任务ID',
                        align: 'center',
                        maxWidth: 80,
                        key: 'id'
                    },
                    {
                        title: '任务名',
                        align: 'center',
                        key: 'name'
                    },
                    {
                        title: '调度周期',
                        align: 'center',
                        width: 90,
                        key: 'scheduleType'
                    },
                    {
                        title: '调度时间',
                        align: 'center',
                        width: 100,
                        key: 'scheduleFormat'
                    },
                    {
                        title: '最近状态',
                        align: 'center',
                        width: 200,
                        render: (h, param) => {
                            return h('div', [
                                param.row.status.map(s => h('Tag', s))
                            ])

                        }
                    },
                    {
                        title: '操作',
                        align: 'center',
                        width: 200,
                        render: (h, param) => {
                            return h('div', [
                                h('Button', '编辑'),
                                h('Button', '删除'),
                            ])

                        }
                    }
                ],
                tasks: [
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                    {
                        id: 123456,
                        name: '任务1',
                        user_name: '管理员',
                        scheduleType: '每天',
                        command: 'echo hello world',
                        scheduleFormat: '00:10:00',
                        status: ['成功', '失败', '成功'],
                        create_time: '2019-01-01 01:02:03',
                        update_time: '2019-01-02 01:02:03',
                        father_tasks: [
                            {
                                id: 1,
                                name: '父任务1'
                            },
                            {
                                id: 2,
                                name: '父任务2'
                            }
                        ]
                    },
                ]
            }
        }
    }
</script>

<style scoped>
    .expand-row {
        margin-bottom: 16px;
    }
</style>