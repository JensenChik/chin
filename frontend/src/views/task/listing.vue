<template>
    <div>
        <Breadcrumb :style="{margin: '16px 0'}">
            <BreadcrumbItem>Home</BreadcrumbItem>
            <BreadcrumbItem>任务管理</BreadcrumbItem>
            <BreadcrumbItem>任务详情</BreadcrumbItem>
        </Breadcrumb>
        <Card>
            <Button style="margin-bottom: 10px">新建</Button>
            <Table style="margin-bottom: 10px" :columns="columns" :data="tasks"></Table>
            <div align="right">
                <Page :total="100" show-elevator show-total show-sizer/>
            </div>

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
                        renderHeader: (h, params) => {
                            return h('div', [
                                '任务ID',
                                h('Icon', {
                                    props: {
                                        type: 'md-search'
                                    },
                                    style: {
                                        marginLeft: '5px'

                                    },
                                    on: {
                                        click: () => {
                                            alert(1)
                                        }
                                    }
                                })
                            ])
                        },
                        align: 'center',
                        maxWidth: 100,
                        key: 'id'
                    },
                    {
                        renderHeader: (h, params) => {
                            return h('div', [
                                '任务名',
                                h('Icon', {
                                    props: {
                                        type: 'md-search'
                                    },
                                    style: {
                                        marginLeft: '5px'

                                    },
                                    on: {
                                        click: () => {
                                            alert(1)
                                        }
                                    }
                                })
                            ])
                        },
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
                        title: '最近三天状态',
                        align: 'center',
                        width: 200,
                        render: (h, param) => {
                            return h('div', [
                                param.row.status.map(s => h('Button', {
                                    props: {
                                        type: 'success',
                                        shape: 'circle',
                                        size: 'small',
                                        icon: 'md-checkmark'
                                    },
                                    style: {
                                        marginLeft: '5px'
                                    }
                                }))
                            ])

                        }
                    },
                    {
                        title: '操作',
                        align: 'center',
                        width: 200,
                        render: (h, param) => {
                            return h('div', [
                                h('Button', {
                                    props: {
                                        type: 'primary',
                                        shape: 'circle',
                                        size: 'small',
                                        icon: 'md-create'
                                    }
                                }),
                                h('Button', {
                                    props: {
                                        type: 'info',
                                        shape: 'circle',
                                        size: 'small',
                                        icon: 'md-git-branch'
                                    },
                                    style: {
                                        marginLeft: '5px'
                                    }
                                }),
                                h('Button', {
                                    props: {
                                        type: 'warning',
                                        shape: 'circle',
                                        size: 'small',
                                        icon: 'md-play'
                                    },
                                    style: {
                                        marginLeft: '5px'
                                    }
                                }),
                                h('Button', {
                                    props: {
                                        type: 'error',
                                        shape: 'circle',
                                        size: 'small',
                                        icon: 'md-trash'
                                    },
                                    style: {
                                        marginLeft: '5px'
                                    }
                                })
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