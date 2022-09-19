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
                    <!-- <el-button
                    size="small"
                    type="primary"
                    @click="upgradeServeView(row)"
                    >升级</el-button
                    > -->
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


        <el-dialog v-model="upgradeDisable" title="升级服务" width="40%" draggable>
            <el-form :model="upgradeForm">
                <el-upload
                    class="upload-demo"
                    drag
                    multiple
                    :http-request="httpRequest"
                    :limit="1"
                >
                    <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                    <div class="el-upload__text">
                    拖到此区域上传 或者 <em>点击上传</em>
                    </div>
                    <template #tip>
                    <div class="el-upload__tip" style="color: red;" align="left">
                        <span>友情提示:</span> <br/>
                        <span>1. 包名必须与当前服务里start.sh 脚本中构建的名称保持一致</span> <br/>
                        <span>2. 包名不能带后缀</span> <br/>
                        <span>3. 包文件大小不能超过100M</span> <br/>
                        <span>4. 版本号不得重复</span> <br/>
                        <span>5. 当前版本号不得小于上次升级版本号</span>
                    </div>
                    </template>
                </el-upload>

                <el-form-item>
                    输入版本号: &nbsp; &nbsp; <el-input v-model="upgradeForm.upgrade_version" autocomplete="off" style="width: 200px" />
                </el-form-item>
            </el-form>

         
                <div style="margin-left: -400px;">
                    <el-button @click="upgradeDisable = false">取消</el-button>
                    <el-button type="primary" @click="upgradeApp(2)">上传</el-button>  
                </div>



                <div style="marigin-top: 10px; margin-bottom: 10px;" align="left">
                    记录
                </div>
                <el-table :data="upgradeRecordTable" style="width: 100%">
                    <el-table-column prop="created_at" label="升级时间" width="200" />
                    <el-table-column prop="serve_name" label="服务名称" width="200" />
                    <el-table-column prop="pack_name" label="包名" width="200" />
                    <el-table-column prop="upgrade_version" label="版本号" />
                </el-table>
       

                <div class="demo-pagination-block">
                    <el-pagination
                    v-model:currentPage="upgradeRecordForm.page"
                    v-model:page-size="upgradeRecordForm.page_size"
                    :page-sizes="[10, 20, 30, 40, 50]"
                    :small="small"
                    :disabled="disabled"
                    :background="background"
                    layout="total, sizes, prev, pager, next, jumper"
                    :total="upgradeTotal"
                    :current-page="upgradeRecordForm.page"
                    @size-change="handleSizeUpgradeChange"
                    @current-change="handleUpgradeCurrentChange"
                    />
                </div>

        </el-dialog>



    </div>
</template>

<script setup>
import { ref, reactive, onUnmounted } from 'vue'
import { useRouter } from 'vue-router';
import { serveList, deleteServe, createServe, upgradeServe, upgradeServeRecord } from '@/request/api'
import { ElMessage, ElMessageBox } from 'element-plus';

const small = ref(false);
const disabled = ref(false);
const background = ref(false)

const router = useRouter();

// 升级服务 - 上传包文件
const upgradeDisable = ref(false);

const upgradeServeView = (row) => {
    upgradeDisable.value = true
    upgradeForm.value.serve_id = row.id
    upgradeForm.value.serve_name = row.serve_name
    upgradeForm.value.serve_address = row.serve_address
    upgradeForm.value.serve_state = row.serve_state
    upgradeRecordForm.value.serve_address = row.serve_address
    upgradeRecord()
}

const upgradeForm = ref({
    go_file: "",
    upgrade_version: "",
    serve_id: 0,
    serve_name: "",
    serve_address: "",
    serve_state: 0,
});

const httpRequest = (param) => {
    let fileObj = param.file // 相当于input里取得的files
    upgradeForm.value.go_file = fileObj
    upgradeApp(1)
}

// 升级服务
const upgradeApp = async (t) => {
    if (t === 1) {
        
    } else {
        let fd = new FormData()// FormData 对象
        fd.append('go_file',upgradeForm.value.go_file)
        fd.append('upgrade_version',upgradeForm.value.upgrade_version)
        fd.append('serve_id',upgradeForm.value.serve_id)
        fd.append('serve_name',upgradeForm.value.serve_name)
        fd.append('serve_address',upgradeForm.value.serve_address)
        fd.append('serve_state',upgradeForm.value.serve_state)

        let res = await upgradeServe(fd)
        if (res) {
            if (res.code === 2000) {
                ElMessage.success(res.msg)
                upgradeDisable.value = false
            }
        }
        // requestServeList()
    }
  
}

// 升级服务记录
const upgradeRecordTable = ref([]);
const upgradeTotal = ref(0);

const upgradeRecordForm = ref({
    page: 1,
    page_size: 10,
    serve_address: ""
})

const upgradeRecord = async () => {
    let request  = {
        page: upgradeRecordForm.value.page,
        page_size: upgradeRecordForm.value.page_size,
        serve_address: upgradeRecordForm.value.serve_address,
    }

    let res = await upgradeServeRecord(request)
    if (res.code !== 2000) {
        res.data = []
    }
    upgradeRecordTable.value = res.data.list
    upgradeTotal.value = res.data.total
}

const handleSizeUpgradeChange = (row) => {
    upgradeRecordForm.value.page_size = row
    upgradeRecordForm.value.page = 1
    upgradeRecord()
}

const handleUpgradeCurrentChange = (row) => {
    upgradeRecordForm.value.page = row
    upgradeRecord();
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