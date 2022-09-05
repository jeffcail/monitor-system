<template>
    <div class="flex flex-wrap items-left">
        <el-form-item label="选择客户端服务器" prop="">
               <el-select v-model="value" placeholder="" @change="selectChange($event)">
                <el-option
                v-for="item in tableData"
                :key="item.id"
                :label="item.hostname"
                :value="item.id">
                </el-option>
                 </el-select>
            </el-form-item>
    </div>



    <div>
        <div
            class="echart"
            id="mychart"
            :style="{ float: 'left', width: '100%', height: '400px' }"
        ></div>  

    </div>

</template>

<script setup>
import { ArrowDown } from '@element-plus/icons-vue'
import { ref } from '@vue/reactivity';
import * as echarts from "echarts";
import { allMachine } from '@/request/api';

const tableData = ref([]);

const all_machine = async () => {
    let res = await allMachine()
    // console.log(res)
    if (res.code !== 2000) {
        res.data = []
    }
    tableData.value = res.data
}
all_machine()

</script>


<style lang="scss" scoped>
.flex {
    margin-left: 40px;
}

</style>
