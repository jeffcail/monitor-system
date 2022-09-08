<template>
    <div class="add-admin">
        
    </div>

    <div class="admin-list">
        <el-table :data="tableData" style="width: 100%">
           
           <el-table-column label="ID" width="140">
           <template #default="scope">
               <div style="display: flex; align-items: center">
               <span style="margin-left: 10px">{{ scope.row.id }}</span>
               </div>
           </template>
           </el-table-column>

           <el-table-column label="用户名" width="120">
           <template #default="scope">
               <div style="display: flex; align-items: center">
               <span style="margin-left: 10px">{{ scope.row.admin_username }}</span>
               </div>
           </template>
           </el-table-column>

           <!-- <el-table-column label="路径" width="180">
           <template #default="scope">
               <div style="display: flex; align-items: center">
               <span style="margin-left: 10px">{{ scope.row.url }}</span>
               </div>
           </template>
           </el-table-column> -->

           <el-table-column label="日志" width="680">
           <template #default="scope">
               <div style="display: flex; align-items: center">
               <span style="margin-left: 10px">{{ scope.row.content }}</span>
               </div>
           </template>
           </el-table-column>


           <el-table-column label="操作时间" width="180">
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
                   type="primary"
                   @click="update_admin(row)"
                   >修改</el-button
                   >

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
import { ref, reactive } from 'vue';
import { operateLogList } from '@/request/api';

const total = ref(0);
const tableData = ref([]);


const formJsonIn = ref({
    page: 1,
    page_size: 10,
})

const log_list = async () => {
    let request = {
        page: formJsonIn.value.page,
        page_size: formJsonIn.value.page_size,
    }
    let res = await operateLogList(request)
    console.log(res)
    if (res.code !== 2000) {
        res.data = []
    }
    tableData.value = res.data.list
    total.value = res.data.total
}
log_list()

const handleSizeChange = (row) => {
    formJsonIn.value.page_size = row
    formJsonIn.value.page = 1
    log_list()
}

const handleCurrentChange = (row) => {
    formJsonIn.value.page = row
    log_list();
}

</script>

<style lang="scss">
.admin-list{
    margin-left: 40px;
}
</style>
