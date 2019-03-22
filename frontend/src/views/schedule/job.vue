<template>
    <div>
        <Breadcrumb :style="{margin: '16px 0'}">
            <BreadcrumbItem>Home</BreadcrumbItem>
            <BreadcrumbItem>调度管理</BreadcrumbItem>
            <BreadcrumbItem>作业详情</BreadcrumbItem>
        </Breadcrumb>
        <Card>
            <Button style="margin-bottom: 10px" type="primary">新建</Button>
            <Table style="margin-bottom: 10px" :columns="columns" :data="tasks"></Table>
            <div align="right">
                <Page :total="100" show-elevator show-total show-sizer/>
            </div>

        </Card>
    </div>
</template>

<script>
    export default {
        name: "job"
    }
</script>

<style scoped>

</style>