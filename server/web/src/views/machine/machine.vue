<template>
    <div class="machine-list">
        <el-table :data="tableData" style="width: 100%">
           
            <el-table-column label="机器码" width="280">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <span style="margin-left: 10px">{{ scope.row.machine_code }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="IP" width="280">
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
<!-- 
            <el-table-column label="备注" width="280">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <span style="margin-left: 10px">{{ scope.row.remark }}</span>
                </div>
            </template>
            </el-table-column> -->

            <el-table-column label="创建时间" width="280">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <el-icon><timer /></el-icon>
                <span style="margin-left: 10px">{{ scope.row.created_at }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="操作" v-slot="{ row }">
                    <!-- <el-button
                    size="big"
                    type="success"
                    @click="handleDelete(row)"
                    ><el-icon><Monitor /></el-icon></el-button
                    > -->

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

    <!-- <el-dialog v-model="dialSshVisiable" title="Tips" width="30%" draggable>
        
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
        <el-button @click="dialSshVisiable = false">取消</el-button>
        <el-button type="primary" @click="dialSsh">🔗</el-button>
      </span>
    </template>
  </el-dialog> -->


</template>

<script setup>
import 'xterm/css/xterm.css'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { reactive, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router';
import { machineList} from '@/request/api'
import { ElMessage } from 'element-plus';

let terminalBox = ref(null)
let term
let socket


const small = ref(false);
const disabled = ref(false);
const background = ref(false)

const router = useRouter();

const showConsole = (row) => {
    sessionStorage.setItem(`url`, `/monitor/machine/dial`)
    router.push({
        path: "/monitor/machine/dial",
        query: {id: row.ip}
    })
}

// onMounted(() => {
//     // console.log(dialSshForm.ip);

//     term = new Terminal({
//         rendererType: 'canvas', //使用这个能解决vim不显示或乱码
//         cursorBlink: true,
//         cursorStyle: "bar",
//     })

//     const fitAddon = new FitAddon()
//     term.loadAddon(fitAddon)
//     term.open(terminalBox.value)
//     fitAddon.fit()

//     // 创建socket连接
//     term.write('正在连接...\r\n');
//     socket = new WebSocket('ws://127.0.0.1:9092/ssh')
//     socket.binaryType = "arraybuffer";

//     socket.onopen = function () {
//         fitAddon.fit()
//         term.onData(function (data) {
//             socket.send(data)
//             console.log(data)
//         })
//         ElMessage.success("会话连接成功!")
//     }
//     socket.onclose = function () {
//         term.writeln('连接关闭');
//     }
//     socket.onerror = function (err) {
//         term.writeln('读取数据异常：', err);
//     }

//     // 接受数据
//     socket.onmessage = function (recv) {
//         try {
//             term.write(recv.data)
//         } catch (e) {
//             console.log('unsupport data', recv.data)
//         }
//     }

//     window.addEventListener("resize", () => {
//         fitAddon.fit()
//     }, false)


//     console.log(term);
// })




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