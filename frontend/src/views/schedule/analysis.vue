<template>
    <div>
        <Breadcrumb :style="{margin: '16px 0'}">
            <BreadcrumbItem>Home</BreadcrumbItem>
            <BreadcrumbItem>调度管理</BreadcrumbItem>
            <BreadcrumbItem>调度分析</BreadcrumbItem>
        </Breadcrumb>
        <Card>
            <div style="height: 600px">Content</div>
        </Card>
    </div>

</template>

<script>
    export default {
        name: "analysis"
    }
</script>

<style scoped>

</style>