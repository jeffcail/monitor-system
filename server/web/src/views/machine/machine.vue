<template>
    <div class="machine-list">
        <el-table :data="tableData" style="width: 100%">
           
            <el-table-column label="机器码" width="290">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <span style="margin-left: 10px">{{ scope.row.machine_code }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="IP" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <span style="margin-left: 10px">{{ scope.row.ip }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="主机名字" width="280">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <span style="margin-left: 10px">{{ scope.row.host_name }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="备注" width="280">
                <template  #default="scope">
                    <el-input
                    v-model="scope.row.remark"
                    class="w-50 m-2"
                    size="large"
                    placeholder="请输入备注"
                    @blur="addMachineRemark(scope.row)"
                    />
                </template>
                
            </el-table-column>

            <el-table-column label="创建时间" width="280">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <el-icon><timer /></el-icon>
                <span style="margin-left: 10px">{{ scope.row.created_at }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="操作" v-slot="{ row }">
                    <el-button
                    size="small"
                    type="success"
                    @click="sendCommond(row)"
                    >发送指令</el-button>

                    <el-button
                    size="small"
                    type="success"
                    @click="upgradeClientView(row)"
                    >升级客户端</el-button>

                    <el-button
                    size="big"
                    type="success"
                    @click="showConsole(row)"
                    ><el-icon><Monitor /></el-icon></el-button
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
    </div>

        <el-dialog v-model="sendCommondDisable" title="发送指令" width="30%" draggable>
        
            <el-form :model="sendCommondForm">
                <el-form-item>
                    指令 &nbsp;&nbsp; <el-input v-model="sendCommondForm.content" autocomplete="off" placeholder="指令为linux命令 如: ls && mkdir && cd /root/xxx/stat.sh" style="width: 400px" />
                </el-form-item>
            </el-form>

            <template #footer>
            <span class="dialog-footer">
                <el-button @click="sendCommondDisable = false">取消</el-button>
                <el-button type="primary" @click="sendCommondSubmit">发送</el-button>
            </span>
            </template>
        </el-dialog>

        <el-dialog v-model="upgradeClientDisable" title="升级客户端" width="40%" draggable>
            <el-form :model="upgradeClientForm">
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
                        <span>1. 包名不能带后缀</span> <br/>
                        <span>2. 包文件大小不能超过100M</span> <br/>
                        <span>3. 版本号不得重复</span> <br/>
                        <span>4. 当前版本号不得小于上次升级版本号</span>
                    </div>
                    </template>
                </el-upload>

                <el-form-item>
                    输入版本号: &nbsp; &nbsp; <el-input v-model="upgradeClientForm.upgrade_version" autocomplete="off" style="width: 200px" />
                </el-form-item>
            </el-form>

         
                <div style="margin-left: -400px;">
                    <el-button @click="upgradeClientDisable = false">取消</el-button>
                    <el-button type="primary" @click="upgradeClient(2)">上传</el-button>  
                </div>



                <div style="marigin-top: 10px; margin-bottom: 10px;" align="left">
                    记录
                </div>
                <el-table :data="upgradeClientRecordTable" style="width: 100%">
                    <el-table-column prop="created_at" label="升级时间" width="200" />
                    <el-table-column prop="package_name" label="包名" width="200" />
                    <el-table-column prop="upgrade_version" label="版本号" />
                </el-table>
       

                <div class="demo-pagination-block">
                    <el-pagination
                    v-model:currentPage="upgradeClientRecordForm.page"
                    v-model:page-size="upgradeClientRecordForm.page_size"
                    :page-sizes="[10, 20, 30, 40, 50]"
                    :small="small"
                    :disabled="disabled"
                    :background="background"
                    layout="total, sizes, prev, pager, next, jumper"
                    :total="upgradeClientTotal"
                    :current-page="upgradeClientRecordForm.page"
                    @size-change="handleSizeUpgradeClientChange"
                    @current-change="handleUpgradeClientCurrentChange"
                    />
                </div>

        </el-dialog>


    <el-dialog v-model="dialSshVisiableShow" title="终端" width="30%" draggable>
        
        <el-form :model="dialSshForm">
            <el-form-item>
                服务器地址: <el-input v-model="dialSshForm.ip" autocomplete="off" placeholder="127.0.0.1:22" style="width: 200px" />
            </el-form-item>

            <el-form-item>
                服务用户名: <el-input v-model="dialSshForm.name" autocomplete="off" style="width: 200px" />
            </el-form-item>

            <el-form-item>
                服务器密码: <el-input v-model="dialSshForm.password" autocomplete="off" style="width: 200px" />
            </el-form-item>
        </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialSshVisiableShow = false">取消</el-button>
        <el-button type="primary" @click="connectSsh">🔗</el-button>
      </span>
    </template>
  </el-dialog>


</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router';
import { machineList, sendMachineCommond, updateMachineRemark, upgradeClientServe, upgradeClientRecord } from '@/request/api'
import { ElMessage } from 'element-plus';

let terminalBox = ref(null)
let term
let socket

// 升级客户端
const upgradeClientDisable = ref(false);

const upgradeClientView = (row) => {
    upgradeClientDisable.value = true
    upgradeClientForm.value.machine_code = row.machine_code
    upgradeClientForm.value.machine_ip = row.ip
    upgradeClientForm.value.machine_hostname = row.host_name
    upgradeClientForm.value.machine_remark = row.remark

    upgradeClientRecordForm.value.machine_ip = row.ip

    upgradeClientRecords()
}

const upgradeClientForm = ref({
    go_file: "",
    machine_code: "",
    machine_ip: "",
    machine_hostname: "",
    machine_remark: "",
    upgrade_version: ""
});

const httpRequest = (param) => {
    let fileObj = param.file // 相当于input里取得的files
    upgradeClientForm.value.go_file = fileObj
    upgradeClient(1)
}

const upgradeClient = async(t) => {
    if (t === 1) {

    } else {
        let fd = new FormData()// FormData 对象
        fd.append('go_file',upgradeClientForm.value.go_file)
        fd.append('machine_code',upgradeClientForm.value.machine_code)
        fd.append('machine_ip', upgradeClientForm.value.machine_ip)
        fd.append('machine_hostname', upgradeClientForm.value.machine_hostname)
        fd.append('machine_remark', upgradeClientForm.value.machine_remark)
        fd.append('upgrade_version', upgradeClientForm.value.upgrade_version)

        let res = await upgradeClientServe(fd)
        if (res) {
            if (res.code === 2000) {
                ElMessage.success(res.msg)
                upgradeClientDisable.value = false
            }
        }
    }
}


// 查看客户端升级记录
const upgradeClientRecordTable = ref([]);
const upgradeClientTotal = ref(0);


const upgradeClientRecordForm = ref({
    page: 1,
    page_size: 10,
    machine_ip: ""
})

const upgradeClientRecords = async() => {
    let request  = {
        page: upgradeClientRecordForm.value.page,
        page_size: upgradeClientRecordForm.value.page_size,
        machine_ip: upgradeClientRecordForm.value.machine_ip,
    }

    let res = await upgradeClientRecord(request)
    // console.log(res);
    if (res.code !== 2000) {
        res.data = []
    }
    upgradeClientRecordTable.value = res.data.list
    upgradeClientTotal.value = res.data.total
}


const handleSizeUpgradeClientChange = () => {
    upgradeClientRecordForm.value.page_size = row
    upgradeClientRecordForm.value.page = 1
    upgradeClientRecords()
}

const handleUpgradeClientCurrentChange = () => {
    upgradeClientRecordForm.value.page = row
    upgradeClientRecords();
}

// 备注
const addMachineRemark = async (row) => {
    let request = {
        machine_code: row.machine_code,
        ip: row.ip,
        remark: row.remark,
    }
    
    let res = await updateMachineRemark(request)
    console.log(res);
    if (res.code === 2000) {
        ElMessage.success(res.msg)
        machine_list()
    }
}


const small = ref(false);
const disabled = ref(false);
const background = ref(false)

const router = useRouter();


// 发送指令
const sendCommondDisable = ref(false)

const sendCommondForm = ref({
    content: "",
    ip: ""
});

const sendCommond = (row) => {
    sendCommondDisable.value = true
    sendCommondForm.value.ip = row.ip
}

const sendCommondSubmit = async () => {
    let request  = {
        ip: sendCommondForm.value.ip,
        content: sendCommondForm.value.content,
    }
    let res = await sendMachineCommond(request)
    if (res.code === 2000) {
        ElMessage.success("指令发送成功")
        sendCommondDisable.value = false
        machine_list()
    }
}

const dialSshVisiableShow = ref(false);
const dialSshForm = ref({
    ip: "",
    name: "",
    password:""
})

const showConsole = (row) => {
    dialSshVisiableShow.value = true
    dialSshForm.value.ip = row.ip
    dialSshForm.value.name = row.name
    dialSshForm.value.password = row.password
}

const connectSsh = (row) => {
    // console.log(dialSshForm.value);
    sessionStorage.setItem(`url`, `/monitor/machine/dial`)
    router.push({
        path: "/monitor/machine/dial",
        query: {ip: row.ip}
    })
}

const total = ref(0);
const tableData = ref([]);

const formJsonIn = ref({
    page: 1,
    page_size: 10,
})

const serve1 = ref();

const machine_list = async () => {
    let request = {
        page: formJsonIn.value.page,
        page_size: formJsonIn.value.page_size,
    }
    let res = await machineList(request)
    // console.log(res)
    if (res.code !== 2000) {
        res.data = []
    }
    tableData.value = res.data.list
    total.value = res.data.total
}
machine_list()

const handleSizeChange = (row) => {
    formJsonIn.value.page_size = row
    formJsonIn.value.page = 1
    machine_list()
}

const handleCurrentChange = (row) => {
    formJsonIn.value.page = row
    machine_list();
}


const handleDelete = (row) => {
    ElMessageBox.confirm(`确定删除${row.host_name}吗?, 删除之前请先确定此客户端是否还在使用，还在使用中请勿删除`, '提示', {
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

</script>

<style lang="scss" scoped>
.machine-list{
    margin-left: 40px;
}
</style>