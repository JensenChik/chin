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

    .menu-item span{
        display: inline-block;
        overflow: hidden;
        width: 69px;
        text-overflow: ellipsis;
        white-space: nowrap;
        vertical-align: bottom;
        transition: width .2s ease .2s;
    }
    .menu-item i{
        transform: translateX(0px);
        transition: font-size .2s ease, transform .2s ease;
        vertical-align: middle;
        font-size: 16px;
    }

    .layout-logo {
        width: 150px;
        height: 30px;
        background: #5b6270;
        border-radius: 3px;
        float: left;
        position: relative;
        top: 15px;
        left: 20px;
        color: #BBB;
        font-weight: bolder;
        font-size: 1.8em;
    }


    .collapsed-menu span {
        width: 0px;
        transition: width .2s ease;
    }

    .collapsed-menu i {
        transform: translateX(5px);
        transition: font-size .2s ease .2s, transform .2s ease .2s;
        vertical-align: middle;
        font-size: 22px;
    }

</style>
<template>
    <div class="layout">
        <Layout :style="{minHeight: '100vh'}">
            <Sider collapsible :collapsed-width="78" v-model="isCollapsed">
                <div class="layout-logo" align="center">
                    Chin
                </div>
                <Menu theme="dark" width="auto" :class="menuitemClasses">
                    <Submenu name="1" style="margin-top: 50px">
                        <template slot="title">
                            <Icon type="logo-buffer"></Icon>
                            <span> 任务管理 </span>
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
                            <Icon type="md-alarm"></Icon>
                            <span> 调度管理 </span>
                        </template>
                        <MenuItem name="2-1">
                            <router-link :to="{name: 'ScheduleJob'}" tag="div">
                                作业详情
                            </router-link>
                        </MenuItem>
                        <MenuItem name="2-2">
                            <router-link :to="{name: 'ScheduleInstance'}" tag="div">
                                实例详情
                            </router-link>
                        </MenuItem>
                        <MenuItem name="2-3">
                            <router-link :to="{name: 'ScheduleAnalysis'}" tag="div">
                                调度分析
                            </router-link>
                        </MenuItem>
                    </Submenu>
                    <Submenu name="3">
                        <template slot="title">
                            <Icon type="md-desktop"></Icon>
                            <span> 机器管理 </span>
                        </template>
                        <MenuItem name="3-1">
                            <router-link :to="{name: 'MachineListing'}" tag="div">
                                机器详情
                            </router-link>
                        </MenuItem>
                        <MenuItem name="3-2">
                            <router-link :to="{name: 'MachineStatusAnalysis'}" tag="div">
                                状态分析
                            </router-link>
                        </MenuItem>
                    </Submenu>
                    <Submenu name="4">
                        <template slot="title">
                            <Icon type="md-lock"></Icon>
                            <span> 权限管理 </span>
                        </template>
                        <MenuItem name="4-1">
                            <router-link :to="{name: 'AuthAction'}" tag="div">
                                操作日志
                            </router-link>
                        </MenuItem>
                        <MenuItem name="4-2">
                            <router-link :to="{name: 'AuthGroup'}" tag="div">
                                组权限
                            </router-link>
                        </MenuItem>
                        <MenuItem name="4-3">
                            <router-link :to="{name: 'AuthUser'}" tag="div">
                                用户权限
                            </router-link>
                        </MenuItem>
                    </Submenu>
                </Menu>
            </Sider>
            <Layout>
                <Header :style="{background: '#fff', boxShadow: '0 2px 3px 2px rgba(0,0,0,.1)'}">
                    <div align="right">
                        超级管理员
                        <Button type="text" size="large" icon="md-log-out" shape="circle"
                                style="margin-left: 5px"></Button>
                    </div>
                </Header>
                <Content :style="{padding: '0 16px 16px'}">
                    <router-view></router-view>
                </Content>

            </Layout>
        </Layout>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                isCollapsed: false
            }
        },
        computed: {
            menuitemClasses: function () {
                return [
                    'menu-item',
                    this.isCollapsed ? 'collapsed-menu' : ''
                ]
            }
        }
    }
</script>
