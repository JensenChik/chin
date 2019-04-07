<template>
    <div>
        <Breadcrumb :style="{margin: '16px 0'}">
            <BreadcrumbItem>Home</BreadcrumbItem>
            <BreadcrumbItem>机器管理</BreadcrumbItem>
            <BreadcrumbItem>机器详情</BreadcrumbItem>
        </Breadcrumb>
        <Card>
            <Table style="margin-bottom: 10px" :columns="columns" :data="machines"></Table>
            <div align="right">
                <Page :total="100" show-elevator show-total show-sizer/>
            </div>
        </Card>
    </div>

</template>

<script>
    import {Circle} from 'iview';

    export default {
        name: "listing",
        created() {
            this.machines = [
                {
                    ip: '123.456.789.012',
                    mac: '04:7d:7b:e0:3e:6b',
                    load: {
                        cpu: 10,
                        memory: 10,
                        disk: 20
                    }
                }
            ]
        },
        data() {
            return {
                columns: [
                    {
                        title: '机器ID',
                        width: 80,
                        align: 'center',
                    },
                    {
                        title: '机器名',
                        align: 'center',
                        key: 'name',
                    },
                    {
                        title: 'IP',
                        align: 'center',
                        width: 130,
                        key: 'ip',
                    },
                    {
                        title: 'MAC',
                        align: 'center',
                        width: 150,
                        key: 'mac',
                    },
                    {
                        title: '负载',
                        align: 'center',
                        width: 120,
                        render: (h, params) => {
                            let loads = params.row.load;
                            return h('div',
                                ['cpu', 'memory', 'disk'].map(type => h(Circle, {
                                    props: {
                                        size: 20,
                                        percent: loads[type]
                                    },
                                    style: {
                                        marginRight: '5px'
                                    }
                                }, loads[type]))
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
                machines: []
            }

        },
    }
</script>

<style scoped>

</style>