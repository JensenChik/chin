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
                <span class="expand-key">负责人：</span>
                <span class="expand-value">{{ row.user_name }}</span>
            </Col>
            <Col span="8">
                <span class="expand-key">是否启用：</span>
                <span class="expand-value">***</span>
            </Col>
            <Col span="8">
                <span class="expand-key">优先级：</span>
                <Rate disabled :value="3"/>
            </Col>
        </Row>


        <Row class="expand-row">
            <Col span="8">
                <span class="expand-key">失败是否重跑：</span>
                <span class="expand-value">yes</span>
            </Col>
            <Col span="8">
                <span class="expand-key">延迟重跑：</span>
                <span class="expand-value">10分钟</span>
            </Col>
            <Col span="8">
                <span class="expand-key">最大重跑次数：</span>
                <span class="expand-value">1次</span>
            </Col>
        </Row>


        <Row class="expand-row">
            <Col span="8">
                <span class="expand-key">等待超时报警</span>
                <span class="expand-value">*****</span>
            </Col>
            <Col span="8">
                <span class="expand-key">执行超时报警：</span>
                <span class="expand-value">yes</span>
            </Col>
            <Col span="8">
                <span class="expand-key">执行失败报警：</span>
                <span class="expand-value">yes</span>
            </Col>
        </Row>

        <Row class="expand-row">
            <Col span="8">
                <span class="expand-key">报警组：</span>
                <span class="expand-value"></span>
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
            <Button size="small">依赖分析</Button>
        </Row>

        <Row class="expand-row">
            <span class="expand-key">执行机器池：</span>
            <span class="expand-value">1</span>
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