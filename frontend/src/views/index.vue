<style scoped>
    .layout {
        border: 1px solid #d7dde4;
        background: #f5f7f9;
        position: relative;
        border-radius: 4px;
        overflow: hidden;
    }

    .layout-header-bar {
        background: #fff;
        box-shadow: 0 1px 1px rgba(0, 0, 0, .1);
    }

</style>
<template>
    <div class="layout">
        <Sider :style="{position: 'fixed', height: '100vh', left: 0, overflow: 'auto'}">
            <Menu theme="dark" width="auto">
                <Submenu name="1">
                    <template slot="title">
                        <Icon type="ios-navigate"></Icon>
                        任务管理
                    </template>
                    <MenuItem name="1-1">
                        <router-link :to="{name: 'TaskListing'}" tag="div">
                            任务详情
                        </router-link>
                    </MenuItem>
                    <MenuItem name="1-2">
                        <router-link :to="{name: 'TaskDependency'}" tag="div">
                        依赖分析
                        </router-link>
                    </MenuItem>
                </Submenu>
                <Submenu name="2">
                    <template slot="title">
                        <Icon type="ios-keypad"></Icon>
                        调度管理
                    </template>
                    <MenuItem name="2-1">作业详情</MenuItem>
                    <MenuItem name="2-2">实例详情</MenuItem>
                    <MenuItem name="2-3">调度分析</MenuItem>
                </Submenu>
                <Submenu name="3">
                    <template slot="title">
                        <Icon type="ios-analytics"></Icon>
                        机器管理
                    </template>
                    <MenuItem name="3-1">机器详情</MenuItem>
                    <MenuItem name="3-2">状态分析</MenuItem>
                </Submenu>
                <Submenu name="4">
                    <template slot="title">
                        <Icon type="ios-analytics"></Icon>
                        权限管理
                    </template>
                    <MenuItem name="4-1">操作日志</MenuItem>
                    <MenuItem name="4-2">组权限</MenuItem>
                    <MenuItem name="4-3">用户权限</MenuItem>
                </Submenu>
            </Menu>
        </Sider>
        <Layout :style="{marginLeft: '200px'}">
            <Header :style="{background: '#fff', boxShadow: '0 2px 3px 2px rgba(0,0,0,.1)'}"></Header>
            <Content :style="{padding: '0 16px 16px'}">
            <router-view></router-view>
            </Content>

        </Layout>
    </div>
</template>
<script>
    export default {}
</script>
