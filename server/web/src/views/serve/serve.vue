<template>

    <div>
        <div class="addServe-info">
        <div class="addServe-title" align="left">添加服务检测</div>
        <el-form
            ref="ruleFormRef"
            :model="ruleForm"
            :rules="rules"
            status-icon
            label-width="120px"
            class="demo-ruleForm"
        >
            <el-form-item label="地址" prop="channelId">
                <el-input v-model="ruleForm.serve_address" style="width: 500px" type="text" autocomplete="off" ></el-input>
            </el-form-item>
            <el-form-item label="名称" prop="url">
                <el-input v-model="ruleForm.serve_name" style="width: 500px" type="text" autocomplete="off"></el-input>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="submitForm()">添加</el-button>
            </el-form-item>
        </el-form>
    </div>
    </div>



    <div class="serve-list">
        <el-table :data="tableData" style="width: 100%">
           
            <el-table-column label="名称" width="280">
            <template #default="scope">
                <el-popover effect="light" trigger="hover" placement="top" width="auto">
                <template #default>
                    <!-- <div>serve_name: {{ scope.row.serve_name }}</div> -->
                    <div>{{ scope.row.serve_address }}</div>
                </template>
                <template #reference>
                    <el-tag>{{ scope.row.serve_name }}</el-tag>
                </template>
                </el-popover>
            </template>
            </el-table-column>

            <el-table-column label="上次检测时间" width="280">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <el-icon><timer /></el-icon>
                <span style="margin-left: 10px">{{ scope.row.last_check_time }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="状态" width="280">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <span style="margin-left: 10px; color: green;" v-if="scope.row.serve_state == 1"> <el-icon><SuccessFilled /></el-icon> </span>
                <span style="margin-left: 10px; color: red;" v-else> <el-icon><WarningFilled /></el-icon> </span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="操作" v-slot="{ row }">
                    <el-button
                    size="small"
                    type="danger"
                    @click="handleDelete(row)"
                    >删除</el-button
                    >
                    <el-button
                    size="small"
                    type="primary"
                    @click="upgradeDisable= true"
                    >升级</el-button
                    >
            </el-table-column>
        </el-table>

        <div class="demo-pagination-block">
            <el-pagination
            v-model:currentPage="formJsonIn.page"
            v-model:page-size="formJsonIn.page_size"
            :page-sizes="[10, 20, 30, 40, 50]"
            :small="small"
            :disabled="disabled"
            :background="background"
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
            :current-page="formJsonIn.page"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            />
        </div>


        <el-dialog v-model="upgradeDisable" title="升级服务" width="30%" draggable>
    
            <el-form :model="upgradeForm">

                <el-form-item>
                    服务地址: <el-input v-model="upgradeForm.serve_ip" placeholder="127.0.0.1" autocomplete="off" style="width: 200px" />
                </el-form-item>                

                <el-form-item>
                    服务包名: <el-input v-model="upgradeForm.package_name" autocomplete="off" style="width: 200px" />
                </el-form-item>

                <el-form-item>
                    服务路径: <el-input v-model="upgradeForm.package_path" autocomplete="off" style="width: 200px" />
                </el-form-item>
            </el-form>

            <template #footer>
            <span class="dialog-footer">
                <el-button @click="upgradeDisable = false">取消升级</el-button>
                <el-button type="primary" @click="upgradeApp">开始升级</el-button>
            </span>
            </template>
        </el-dialog>



    </div>
</template>

<script setup>
import { ref, reactive, onUnmounted } from 'vue'
import { useRouter } from 'vue-router';
import { serveList, deleteServe, createServe, upgradeServe } from '@/request/api'
import { ElMessage, ElMessageBox } from 'element-plus';

const small = ref(false);
const disabled = ref(false);
const background = ref(false)

const router = useRouter();

const upgradeDisable = ref(false);
const upgradeForm = ref({
    package_name: "",
    package_path : "",
    serve_ip: "",
});


// 升级服务
const upgradeApp = async () => {
    let request = {
        serve_ip: upgradeForm.value.serve_ip,
        package_name: upgradeForm.value.package_name,
        package_path: upgradeForm.value.package_path,
    }
    // console.log(request);

    let res = await upgradeServe(request)
    if (res) {
        if (res.code === 2000) {
            ElMessage.success(res.msg)
            upgradeDisable.value = false
        }
    }
    requestServeList()
}

const ruleForm = ref({
    serve_address: "",
    serve_name: "",
})

const rules = reactive({
    serve_address: { required: true, message: "请输入服务地址", trigger: "blur" },
    serve_name: { required: true, message: "请输入服务名称", trigger: "blur" },
})

const ruleFormRef = ref(null);

const submitForm = () => {
    ruleFormRef.value.validate(async volid => {
        if (volid) {
            let request = {
                serve_name: ruleForm.value.serve_name,
                serve_address: ruleForm.value.serve_address,
            }
            let res = await createServe(request)
            if (res.code === 2000) {
                ElMessage.success("添加成功")
                requestServeList()
            }
        } else {
            ElMessage.warning("请填写完整")
        }
    })
}
const handleDelete = (row) => {
    ElMessageBox.confirm(`确定删除${row.serve_name}吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
    }).then(()=>{
        let res = deleteServe({
            id: row.id
        })
        requestServeList()
    }).catch(()=>{
            ElMessage.warning('取消删除')
    })
}

const total = ref(0);
const tableData = ref([]);

const formJsonIn = ref({
    page: 1,
    page_size: 10,
})

const serve1 = ref();

const requestServeList = async () => {

    clearTimeout(serve1.value)
    serve1.value = setTimeout(()=>requestServeList(),20000)

    let request = {
        page: formJsonIn.value.page,
        page_size: formJsonIn.value.page_size,
    }
    let res = await serveList(request)
    // console.log(res)
    if (res.code !== 2000) {
        res.data = []
    }
    tableData.value = res.data.list
    total.value = res.data.total
}
requestServeList()

const handleSizeChange = (row) => {
    formJsonIn.value.page_size = row
    formJsonIn.value.page = 1
    requestServeList()
}

const handleCurrentChange = (row) => {
    formJsonIn.value.page = row
    requestServeList();
}

onUnmounted (()=>{
    clearTimeout(serve1.value)
    clearTimeout(serve1.value)
});

</script>

<style lang="scss" scoped>
.serve-list{
    margin-left: 40px;
}
.addServe-title {
    margin-left: 40px;
    margin-bottom: 20px;
}
</style>