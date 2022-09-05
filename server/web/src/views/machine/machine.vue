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

            <el-table-column></el-table-column>
            <!-- <el-table-column label="操作" v-slot="{ row }">
                    <el-button
                    size="small"
                    type="danger"
                    @click="handleDelete(row)"
                    >删除</el-button
                    >
            </el-table-column> -->
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
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router';
import { machineList} from '@/request/api'


const small = ref(false);
const disabled = ref(false);
const background = ref(false)

const router = useRouter();

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
    console.log(res)
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