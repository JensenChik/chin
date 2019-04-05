<template>
    <div>
        <Breadcrumb :style="{margin: '16px 0'}">
            <BreadcrumbItem>Home</BreadcrumbItem>
            <BreadcrumbItem>调度管理</BreadcrumbItem>
            <BreadcrumbItem>作业详情</BreadcrumbItem>
        </Breadcrumb>
        <Card>
            <Table style="margin-bottom: 10px" :columns="columns" :data="jobs"></Table>
            <div align="right">
                <Page :total="100" show-elevator show-total show-sizer/>
            </div>

        </Card>
    </div>
</template>

<script>
    export default {
        name: "job",

        created() {
            this.jobs = [
                {
                    id: 1,
                    task_id: 1,
                    instances: [
                        {
                            id: 1,
                            status: 'failed'
                        },
                        {
                            id: 2,
                            status: 'success'
                        }
                    ],
                    create_time: '2019-01-01 01:02:03',
                    update_time: '2019-01-01 01:02:03'
                },
                {
                    id: 2,
                    task_id: 1,
                    instances: [
                        {
                            id: 2,
                            status: 'success'
                        }
                    ],
                    create_time: '2019-01-01 01:02:03',
                    update_time: '2019-01-01 01:02:03'
                },
                {
                    id: 3,
                    task_id: 1,
                    instances: [
                        {
                            id: 2,
                            status: 'success'
                        }
                    ],
                    create_time: '2019-01-01 01:02:03',
                    update_time: '2019-01-01 01:02:03'
                },
                {
                    id: 4,
                    task_id: 1,
                    instances: [],
                    create_time: '2019-01-01 01:02:03',
                    update_time: '2019-01-01 01:02:03'
                },
                {
                    id: 5,
                    task_id: 1,
                    instances: [
                        {
                            id: 1,
                            status: 'failed'
                        },
                        {
                            id: 3,
                            status: 'failed'
                        },
                        {
                            id: 2,
                            status: 'success'
                        }
                    ],
                    create_time: '2019-01-01 01:02:03',
                    update_time: '2019-01-01 01:02:03'
                },
                {
                    id: 6,
                    task_id: 1,
                    instances: [
                        {
                            id: 2,
                            status: 'success'
                        }
                    ],
                    create_time: '2019-01-01 01:02:03',
                    update_time: '2019-01-01 01:02:03'
                },
            ]

        },

        data() {
            return {
                columns: [
                    {
                        renderHeader: (h, params) => {
                            return h('div', [
                                '作业ID',
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
                        maxWidth: 100,
                        align: 'center',
                        key: 'id'
                    },
                    {
                        title: '任务ID',
                        align: 'center',
                        maxWidth: 100,
                        render: (h, params) => {

                            return h('div', [
                                params.row.task_id,
                                h('Button', {
                                    props: {
                                        type: 'text',
                                        icon: 'md-help',
                                        shape: 'circle',
                                        size: 'small'
                                    }

                                })
                            ])

                        }
                    },
                    {
                        title: '实例ID',
                        align: 'center',
                        render: (h, params) => {
                            let instances = params.row.instances;
                            let STATUS_OF = {
                                'success': 'success',
                                'failed': 'error',
                            };
                            return h('div', instances.length === 0 ? '未实例化' : instances.map(instance => h('Button', {
                                    props: {
                                        size: 'small',
                                        shape: 'circle',
                                        type: STATUS_OF[instance.status]
                                    },
                                    style: {
                                        marginRight: '5px'
                                    }

                                }, instance.id))
                            )
                        }
                    },
                    {
                        title: '创建时间',
                        align: 'center',
                        key: 'create_time',
                        width: 150,
                    },
                    {
                        title: '更新时间',
                        align: 'center',
                        width: 150,
                        key: 'update_time'
                    },
                    {
                        title: '操作',
                        align: 'center',
                        width: 120,
                        render: (h, params) => {
                            return h('Button', {
                                props: {
                                    type: 'error',
                                    size: 'small',
                                    shape: 'circle',
                                    icon: 'md-trash',
                                }
                            })

                        }
                    }

                ],
                jobs: []
            }
        }
    }
</script>

<style scoped>

</style>