<template>
      <el-container class="app-wrapper">
        <!-- 菜单 -->
        <el-aside :width="asideWidth" class="sidebar-container">
            <el-scrollbar style="height:'100%'">
                <el-menu
                    active-text-color="#ffd04b"
                    background-color="#20222A"
                    class="el-menu-vertical-demo"
                    :default-active="activePath"
                    text-color="#fff"
                    unique-opened
                    router
                >
                    <el-menu-item index="board" @click="goUrl('服务云监控系统','','board')" style="height: 50px !important;background-color: var(--el-menu-hover-bg-color)  !important;">
                        <img src="@/assets/images/logo.png">
                        <span class="backstage-manage">服务云监控系统</span>
                    </el-menu-item>
            
                    <el-sub-menu :index="index" v-for="(item,index) in store.state.menuListAll" :key="index">
                        <template #title>
                         
                            <component :is="item.icons" style="width: 16px;height: 16px;"></component>

                            <span class="caidan-auth">{{ item.menu_name }}</span>
                        </template>
                        <el-menu-item :index="item2.front_url" v-for="item2 in item.children" :key="item2.id" @click="goUrl(item.menu_name,item2.menu_name,item2.front_url)">
                            <component :is="item2.icons" style="width: 16px;height: 16px;"></component>
                            <span class="caidan-auth">{{ item2.menu_name }}</span>
                        </el-menu-item>
                    </el-sub-menu>
                </el-menu>
            </el-scrollbar>
        </el-aside>

        <!-- 菜单end -->
        <el-container class="container">
            <!-- 头部 -->
            <el-header  style="font-size: 16px;background-color: #ffffff">
                <div class="toolbar">
                    <div class="toolbar-open-close">
                            <div>
                                <i class="layui-icon"></i>
                            </div>
                    </div>
                    <div class="toolbar-static">

                    <div class="static">
                        <el-popover placement="bottom" :width="400" trigger="hover">
                            <template #reference>
                                <div class="jcbj">风险报警：<p> {{ machineCount }} </p></div>
                            </template>
                            <div style="display: flex;justify-content: space-between;width: 375px;" v-for="(item, index) in MachineCheckRecordList" :key="index">
                                <div>{{ item.machine_ip }}: {{ item. category}}已使用 {{ item.percent }}'%' 状态: {{ item.level }}</div>
                                <el-button type="info" size="small" round @click="ignoreMachineWarning(item.id)">忽略</el-button>
                            </div>
                        </el-popover>
                    </div>


                    <div class="static">
                        <el-popover placement="bottom" :width="400" trigger="hover">
                            <template #reference>
                                <div class="jcbj">检测报警：<p> {{ serveCount }} </p></div>
                            </template>
                            <div style="display: flex;justify-content: space-between;width: 375px;" v-for="(item, index) in ServeCheckRecodList" :key="index">
                                <div>{{ item.serve_name }} : {{ item.serve_address }} (状态: ⚠️)</div>
                                <el-button type="info" size="small" round @click="ignoreWarning(item.serve_id)">忽略</el-button>
                            </div>
                        </el-popover>
                    </div>
                  
                    

                    <div class="toolbar-select">
                        <el-dropdown>
                            <el-icon class="elIcon"><setting/></el-icon>
                            
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <!-- <el-dropdown-item @click="updatapwd()">修改密码</el-dropdown-item> -->
                                    <el-dropdown-item @click="loginOut()">退出</el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
                        <div class="username-sel">
                            <span>{{ username }}</span>
                        </div>
                    </div>

                </div>
                </div>
            </el-header>
            <!-- 头部end -->
            <!-- 面包屑 -->
            <!-- <breadcrumb></breadcrumb> -->
            <!-- 内容 -->
            <el-main style="background-color: #f2f2f2">
                <router-view />
            </el-main>
            <!-- 内容end -->
        </el-container>
    </el-container>
</template>

<script setup>
import { ref,computed,watch } from "vue"                                                                                                                                                                                                                 
import { useRouter } from "vue-router";
import {Location,Document,Menu as IconMenu,Setting,} from '@element-plus/icons-vue'
import variables from "@/assets/styles/variables.scss"
import store from "@/store/index.ts"  
import { serveCheckRecodList, ignoreServeCheckRecord, machineCheckRecordList, ignoreMachineCheckRecord } from '@/request/api'
import { ElMessage } from 'element-plus';



const router = useRouter()
const menuList = ref([])

const serveCount = ref(0);
const ServeCheckRecodList = ref([]);

const serve_check_record_list = async () => {
    let res = await serveCheckRecodList()
    if (res.code !== 2000) {
        res.data = []
    } 
    serveCount.value = res.data.total
    ServeCheckRecodList.value = res.data.list
}
serve_check_record_list()

// 忽略
const ignoreWarning = async (row) => {
    let request = {
        serve_id: row
    }
    let res = await ignoreServeCheckRecord(request)
    if (res.code === 2000) {
        ElMessage.success(res.msg)
        serve_check_record_list()
    }
}

const machineCount = ref(0);
const MachineCheckRecordList = ref([])

const machine_check_record_list = async () => {
    let res = await machineCheckRecordList()
    // console.log(res);
    if (res.code !== 2000) {
        res.data = []
    }
    machineCount.value = res.data.total
    MachineCheckRecordList.value = res.data.list
}
machine_check_record_list()
const ignoreMachineWarning = async (row) => {
    let request = {
        id: row
    }
    // console.log(request);
    let res = await ignoreMachineCheckRecord(request)
    if (res.code === 2000) {
        ElMessage.success(res.msg)
        machine_check_record_list()
    }
}


const username = ref(window.localStorage.getItem("username"))
const initMenusList = async () => {
    if(!window.localStorage.getItem("token")){
        router.push("/monitor/login")                                                                                                                                                      
    }else{
        // console.log(store.state.menuListAll);
        store.state.menuListAll.length || store.dispatch("loadmenuList");
    }
}
initMenusList()

const asideWidth = computed(() => {
    return "210"
})

//点击刷新之后
const activePath = ref(sessionStorage.getItem(`url`) || '/board')
//判断缓存里面是否有url 如果有跳转到对应页面   如果没有就直接跳转到首页看板
if(sessionStorage.getItem(`url`)){
        router.push(sessionStorage.getItem(`url`))
}else{
    if(localStorage.getItem("token")){
        router.push('/monitor/board')
    }else{
        router.push('/monitor/login')
    }
}
// 保存信息和url到缓存里面
const goUrl = (authName, authName2, url) => {
    store.commit("openNames", authName);
    store.commit("activeName", authName2);
    sessionStorage.setItem(`url`,`${url}`)
    activePath.value = sessionStorage.getItem(`url`)   
}

// 退出登陆
const loginOut = () => {
    localStorage.clear()
    sessionStorage.removeItem("url")
    router.push("/monitor/login")
}

</script>

<style lang="scss" scoped>
@import "../../assets/css/index.scss";
</style>
