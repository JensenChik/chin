<template>
    <div>
        <Breadcrumb :style="{margin: '16px 0'}">
            <BreadcrumbItem>Home</BreadcrumbItem>
            <BreadcrumbItem>任务管理</BreadcrumbItem>
            <BreadcrumbItem>依赖分析</BreadcrumbItem>
        </Breadcrumb>
        <Card>
            <div style="height: 600px">Content</div>
        </Card>
    </div>
</template>

<script>
    export default {
        name: "dependence"
    }
</script>

<style scoped>

</style>