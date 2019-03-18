<style scoped>
    .expand-row {
        margin-bottom: 16px;
    }

    .expand-key {
        font-weight: bold;

    }
</style>
<template>
    <div>
        <Row class="expand-row">
            <Col span="8">
                <span class="expand-key">创建者：</span>
                <span class="expand-value">{{ row.user_name }}</span>
            </Col>
            <Col span="8">
                <span class="expand-key">创建时间：</span>
                <span class="expand-value">{{ row.create_time }}</span>
            </Col>
            <Col span="8">
                <span class="expand-key">更新时间：</span>
                <span class="expand-value">{{ row.update_time }}</span>
            </Col>
        </Row>
        <Row class="expand-row">
            <span class="expand-key">父任务：</span>
            <span v-for="task in row.father_tasks">
                <Button type="success" size="small" style="margin-right: 5px">{{ task.id }}</Button>
            </span>
        </Row>
        <Row>
            <span class="expand-key">调度命令：</span>
            <span class="expand-value">{{ row.command }}</span>
        </Row>
    </div>
</template>
<script>
    export default {
        props: {
            row: Object
        }
    };
</script>