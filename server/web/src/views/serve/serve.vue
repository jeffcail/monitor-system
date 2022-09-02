<template>
    <div class="serve-list">
        <el-table :data="tableData" style="width: 100%">
           
            <el-table-column label="名称" width="180">
            <template #default="scope">
                <el-popover effect="light" trigger="hover" placement="top" width="auto">
                <template #default>
                    <div>serve_name: {{ scope.row.serve_name }}</div>
                    <div>address: {{ scope.row.address }}</div>
                </template>
                <template #reference>
                    <el-tag>{{ scope.row.serve_name }}</el-tag>
                </template>
                </el-popover>
            </template>
            </el-table-column>

            <el-table-column label="上次检测时间" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <el-icon><timer /></el-icon>
                <span style="margin-left: 10px">{{ scope.row.last_check_time }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="状态" width="180">
            <template #default="scope">
                <div style="display: flex; align-items: center">
                <el-icon><timer /></el-icon>
                <span style="margin-left: 10px">{{ scope.row.last_check_time }}</span>
                </div>
            </template>
            </el-table-column>

            <el-table-column label="操作">
            <template #default="scope">
                <el-button
                size="small"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)"
                >删除</el-button
                >
            </template>
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
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { serveList } from '@/request/api'

const small = ref(false);
const disabled = ref(false);
const background = ref(false)


interface User {
  serve_name: string
  last_check_time: string
  serve_state: string
}

const handleDelete = (index: number, row: User) => {
    console.log(index, row)
}

const total = ref(0);
const tableData: User[] = [];

const formJsonIn = ref({
    page: 1,
    page_size: 20,
})

const requestServeList = async () => {
    let request = {
        page: formJsonIn.value.page,
        page_size: formJsonIn.value.page_size,
    }
    let res = await serveList(request)
    console.log(res)
    if (res.code == []) {
        res.data = []
    }
    tableData.values = res.data.list
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

</script>

<style lang="scss" scoped>
.serve-list{
    margin-left: 40px;
}
</style>