<template>
    <div class="add-admin">
        <el-button type="primary" @click="addAdminDisable = true">添加</el-button>
    </div>

    <div class="admin-list">
        <el-table :data="tableData" style="width: 100%">
           
           <el-table-column label="ID" width="180">
           <template #default="scope">
               <div style="display: flex; align-items: center">
               <span style="margin-left: 10px">{{ scope.row.id }}</span>
               </div>
           </template>
           </el-table-column>

           <el-table-column label="用户名" width="180">
           <template #default="scope">
               <div style="display: flex; align-items: center">
               <span style="margin-left: 10px">{{ scope.row.username }}</span>
               </div>
           </template>
           </el-table-column>

           <el-table-column label="手机号" width="180">
           <template #default="scope">
               <div style="display: flex; align-items: center">
               <span style="margin-left: 10px">{{ scope.row.phone }}</span>
               </div>
           </template>
           </el-table-column>

           <el-table-column label="职位" width="180">
           <template #default="scope">
               <div style="display: flex; align-items: center">
               <span style="margin-left: 10px">{{ scope.row.office_post }}</span>
               </div>
           </template>
           </el-table-column>

           <el-table-column prop="state" label="状态" width="120" v-slot="{ row }">
                <el-switch 
                class="switch"
                :active-value="1"
                :inactive-value="2"
                active-text="启用"
                inactive-text="禁用"
                v-model="row.state" 
                @change="changeState(row)"
                />
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

       <el-dialog v-model="updateAdminDisable" title="修改管理员账号" width="30%" draggable>
            
            <el-form 
            ref="ruleFormRef"
            :model="updateForm"
            :rules="rules"
            status-icon
            label-width="120px"
            class="demo-ruleForm"
            >
                <el-form-item>
                    用 户 名: &nbsp; &nbsp;<el-input v-model="updateAdminForm.username" autocomplete="off" style="width: 200px" />
                </el-form-item>

                <el-form-item>
                    真实姓名:  &nbsp;<el-input v-model="updateAdminForm.real_name" autocomplete="off" style="width: 200px" />
                </el-form-item>

                <el-form-item>
                    密   码:  &nbsp; &nbsp; &nbsp; &nbsp;<el-input v-model="updateAdminForm.password" autocomplete="off" style="width: 200px" />
                </el-form-item>

                <el-form-item>
                    手 机 号: &nbsp; &nbsp; <el-input v-model="updateAdminForm.phone" autocomplete="off" style="width: 200px" />
                </el-form-item>
                <el-form-item>
                    部    门： &nbsp; &nbsp; &nbsp;<el-input v-model="updateAdminForm.department" autocomplete="off" style="width: 200px" />
                </el-form-item>
                <el-form-item>
                    职    位:  &nbsp; &nbsp; &nbsp; &nbsp;<el-input v-model="updateAdminForm.office_post" autocomplete="off" style="width: 200px" />
                </el-form-item>
            </el-form>

            <template #footer>
            <span class="dialog-footer">
                <el-button @click="updateAdminDisable = false">取消</el-button>
                <el-button type="primary" @click="submitUpdateForm">确定</el-button>
            </span>
            </template>
        </el-dialog>


       <el-dialog v-model="addAdminDisable" title="添加管理员账号" width="30%" draggable>
            
            <el-form 
            ref="ruleFormRef"
            :model="dialSshForm"
            :rules="rules"
            status-icon
            label-width="120px"
            class="demo-ruleForm"
            >
                <el-form-item>
                    用 户 名: &nbsp; &nbsp;<el-input v-model="addAdminForm.username" autocomplete="off" style="width: 200px" />
                </el-form-item>

                <el-form-item>
                    真实姓名:  &nbsp;<el-input v-model="addAdminForm.real_name" autocomplete="off" style="width: 200px" />
                </el-form-item>

                <el-form-item>
                    密   码:  &nbsp; &nbsp; &nbsp; &nbsp;<el-input v-model="addAdminForm.password" autocomplete="off" style="width: 200px" />
                </el-form-item>

                <el-form-item>
                    手 机 号: &nbsp; &nbsp; <el-input v-model="addAdminForm.phone" autocomplete="off" style="width: 200px" />
                </el-form-item>
                <el-form-item>
                    部    门： &nbsp; &nbsp; &nbsp;<el-input v-model="addAdminForm.department" autocomplete="off" style="width: 200px" />
                </el-form-item>
                <el-form-item>
                    职    位:  &nbsp; &nbsp; &nbsp; &nbsp;<el-input v-model="addAdminForm.office_post" autocomplete="off" style="width: 200px" />
                </el-form-item>
            </el-form>

            <template #footer>
            <span class="dialog-footer">
                <el-button @click="addAdminDisable = false">取消</el-button>
                <el-button type="primary" @click="submitForm">确定</el-button>
            </span>
            </template>
        </el-dialog>

   </div>

</template>

<script setup>
import { ref, reactive } from 'vue'
import {  adminList, addAdmin, updateAdmin, deleteAdmin, adminEnableDisable } from '@/request/api'
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';

const router = useRouter()

// 状态修改
const changeState = (row) => {
  let res =  adminEnableDisable({
    id: row.id,
    state: row.state
  })
  
  console.log(res);
  if (res) {
    if (res.code == "2000") {
      ElMessage.success(res.msg)
      admin_list()
    }
  }
}

// 删除
const handleDelete = (row) => {
    ElMessageBox.confirm(`确定删除 ${row.username} 吗?`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
    }).then(()=>{
        let res = deleteAdmin({
            id: row.id
        })
        admin_list()
    }).catch(()=>{
        ElMessage.warning('取消删除')
    })
}

// 修改
const updateAdminDisable = ref(false)
const updateAdminForm = ref({
    id: 0,
    username: "",
    real_name: "",
    password: "",
    phone: "",
    department: "",
    office_post: "",
});

const update_admin = (row) => {
    updateAdminDisable.value = true
    updateAdminForm.value.id = row.id
    updateAdminForm.value.username = row.username
    updateAdminForm.value.real_name = row.real_name
    updateAdminForm.value.phone = row.phone
    updateAdminForm.value.department = row.department
    updateAdminForm.value.office_post = row.office_post
}

const submitUpdateForm = () => {
    let request = {
        id: updateAdminForm.value.id,
        username: updateAdminForm.value.username,
        real_name: updateAdminForm.value.real_name,
        password: updateAdminForm.value.password,
        phone: updateAdminForm.value.phone,
        department: updateAdminForm.value.department,
        office_post: updateAdminForm.value.office_post,
    }
    let res = updateAdmin(request)
    if (res.code === 2000) {
        ElMessage.success("修改成功")
        router.push("/monitor/serve/list")
        admin_list()
    }
}




const addAdminDisable = ref(false);
// 入参
const addAdminForm = ref({
    username: "",
    real_name: "",
    password: "",
    phone: "",
    department: "",
    office_post: "",
});

// 表单验证规则
// const rules = reactive({
//     username: { required: true, message: '请输入用户名', trigger: 'blur' },
//     real_name: { required: true, message: '请输入真实姓名', trigger: 'blur' },
//     password: { required: true, message: "请输入密码", trigger: 'blur' },
//     phone: { required: true, message: "请输入手机号", trigger: 'blur' },
// })

const ruleFormRef = ref(null)

const submitForm = () => {
    let request = {
        username: addAdminForm.value.username,
        real_name: addAdminForm.value.real_name,
        password: addAdminForm.value.password,
        phone: addAdminForm.value.phone,
        department: addAdminForm.value.department,
        office_post: addAdminForm.value.office_post,
    }
    let res = addAdmin(request)
    if (res.code === 2000) {
        ElMessage.success("添加成功")
        router.push("/monitor/serve/list")
        admin_list()
    }
}


const total = ref(0);
const tableData = ref([]);

const formJsonIn = ref({
    page: 1,
    page_size: 10,
})

const admin_list = async () => {
    let request = {
        page: formJsonIn.value.page,
        page_size: formJsonIn.value.page_size,
    }
    let res = await adminList(request)
    // console.log(res)
    if (res.code !== 2000) {
        res.data = []
    }
    tableData.value = res.data.list
    total.value = res.data.total
}
admin_list()

const handleSizeChange = (row) => {
    formJsonIn.value.page_size = row
    formJsonIn.value.page = 1
    admin_list()
}

const handleCurrentChange = (row) => {
    formJsonIn.value.page = row
    admin_list();
}


</script>
<style lang="scss">
    .admin-list{
        margin-left: 40px;
    }
    .add-admin {
        margin-left: -1500px;
        margin-bottom: 20px;
    }

    /* switch按钮样式 */
.switch .el-switch__label {
    position: absolute;
    display: none;
    color: #fff !important;
}
/*打开时文字位置设置*/
.switch .el-switch__label--right {
    z-index: 1;
}
/* 调整打开时文字的显示位子 */
.switch .el-switch__label--right span{
    margin-right: 9px;
}
/*关闭时文字位置设置*/
.switch .el-switch__label--left {
    z-index: 1;
}
/* 调整关闭时文字的显示位子 */
.switch .el-switch__label--left span{
    margin-left: 9px;
}
/*显示文字*/
.switch .el-switch__label.is-active {
    display: block;
}
/* 调整按钮的宽度 */
.switch.el-switch .el-switch__core,
.el-switch .el-switch__label {
     width: 50px !important;
     margin: 0;
}
</style>