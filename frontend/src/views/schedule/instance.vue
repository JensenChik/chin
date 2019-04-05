<template>
    <div>
        <Breadcrumb :style="{margin: '16px 0'}">
            <BreadcrumbItem>Home</BreadcrumbItem>
            <BreadcrumbItem>调度管理</BreadcrumbItem>
            <BreadcrumbItem>实例详情</BreadcrumbItem>
        </Breadcrumb>
        <Card>
            <Table style="margin-bottom: 10px" :columns="columns" :data="instances"></Table>
            <div align="right">
                <Page :total="100" show-elevator show-total show-sizer/>
            </div>
        </Card>
    </div>
</template>

<script>
    export default {
        name: "instance",
        created() {
            this.instances = []
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
                        key: '实例id'
                    },
                    {
                        title: '作业ID',
                        align: 'center',
                        maxWidth: 100,
                        render: (h, params) => {
                            return h('div', [
                                params.row.job_id,
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
                        title: '执行机器',
                        align: 'center'
                    },
                    {
                        title: '状态',
                        align: 'center'
                    },
                    {
                        title: '执行日志',

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
                instances: []
            }
        }
    }
</script>

<style scoped>

</style>