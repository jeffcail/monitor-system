<template>
    <div class="flex flex-wrap items-left">
        <el-form-item label="选择客户端服务器" prop="">
            <el-select v-model="hostname" placeholder="" @change="selectChange">
                <el-option
                    v-for="item in tableData"
                    :key="item.id"
                    :label="item.content"
                    :value="item.ip"
                />
            </el-select>
        </el-form-item>

        <div class="aaaa">
            <div id="echartsInfoOne" :style="{ width: '50%', height: '500px' }"></div>
            <div id="echartsInfoOne2" :style="{ width: '50%', height: '500px' }"></div>
            <div id="echartsInfoOne3" :style="{ width: '50%', height: '500px' }"></div>
        </div>
    </div>
</template>

<script setup>
import { ArrowDown } from '@element-plus/icons-vue'
import { ref,onUnmounted } from 'vue';
import * as echart from 'echarts';
import { allMachine, clientSysCpu, clientSysMen, clientSysDisk } from '@/request/api';

const tableData = ref([]);
const hostname = ref("")
const all_machine = async () => {
    let res = await allMachine()
    if (res.code !== 2000) {
        res.data = []
    }
    tableData.value = res.data
    hostname.value = res.data[0].ip
    getOrderCount()
    getMemPercent()
    getDiskPercent()
}
all_machine()
    const selectChange = (e) => {
        hostname.value = e
        getOrderCount()
        getMemPercent()
        getDiskPercent()
    }
    const data1 = ref([])
    const data2 = ref([])
    const data3 = ref([])
    const channel1 = ref("")
    const channel2 = ref("")
    const channel3 = ref("")
    const num = ref(0)
    const num2 = ref(0)
    const num3 = ref(0)
    const getOrderCount = async () => {
        clearTimeout(channel1.value)
        channel1.value = setTimeout(()=>{
            getOrderCount()
            let echartsInfoOne = document.getElementById("echartsInfoOne")
            let chart = echart.init(echartsInfoOne);
            if(60 < data1.value.length) {
                data1.value.shift()
            }
            chart.setOption({
                series: [
                    {
                        data: data1.value
                    }
                ]
            });
            num.value++
        },5000)
        const res = await clientSysCpu({
            ip: hostname.value
        })
        if(res){
            data1.value.push({
                num: num.value,
                value: res.data
            })
            if(num.value == 0){
                initEcharts()
            }
        }
    }

    const getMemPercent = async () => {
        clearTimeout(channel2.value)
        channel2.value = setTimeout(()=>{
            getMemPercent()
            let echartsInfoOne2 = document.getElementById("echartsInfoOne2")
            let chart2 = echart.init(echartsInfoOne2);

            if(60 < data2.value.length) {
                data2.value.shift()
            }
            chart2.setOption({
                series: [
                    {
                        data: data2.value
                    }
                ]
            });
            num2.value++
        },5000)


        const res = await clientSysMen({
            ip: hostname.value
        })
        
        if(res){
            data2.value.push({
                num: num2.value,
                value: res.data
            })
            if(num2.value == 0){
                initEcharts2()
            }
        }
    }


    const getDiskPercent = async () => {
        clearTimeout(channel3.value)
        channel3.value = setTimeout(()=>{
            getDiskPercent()
            let echartsInfoOne3 = document.getElementById("echartsInfoOne3")
            let chart3 = echart.init(echartsInfoOne3);

            if(60 < data3.value.length) {
                data3.value.shift()
            }
            chart3.setOption({
                series: [
                    {
                        data: data3.value
                    }
                ]
            });
            num3.value++
        },5000)


        const res = await clientSysDisk({
            ip: hostname.value
        })

        if(res){
            data3.value.push({
                num: num3.value,
                value: res.data
            })
            if(num3.value == 0){
                initEcharts3()
            }
        }
    }


    const initEcharts = () => {
        let echartsInfoOne = document.getElementById("echartsInfoOne")
        echartsInfoOne.removeAttribute("_echarts_instance_");
        let chart = echart.init(echartsInfoOne);
        chart.setOption({
            title: {
                text: 'cpu使用率',
                left: 'center'
            },
            xAxis: {
                type: 'category',
                boundaryGap: false,
                data: data1.value.num
            },
            tooltip: {
                trigger: 'axis'
            },
            yAxis: {
                type: 'value'
            },
            grid: {
                left: '10%',
                right: '4%',
                bottom: '6%',
                containLabel: false
            },
            series: [
                {
                    name: 'cpu使用率',
                    data: data1.value.value,
                    type: 'line',
                    
                }
            ]
        });
        window.onresize = function() {
            //自适应大小
            chart.resize();
        };
    }
    const initEcharts2 = () => {
        let echartsInfoOne = document.getElementById("echartsInfoOne2")
        echartsInfoOne.removeAttribute("_echarts_instance_");
        let chart = echart.init(echartsInfoOne);
        chart.setOption({
            title: {
                text: '内存使用率',
                left: 'center'
            },
            xAxis: {
                type: 'category',
                boundaryGap: false,
                data: data2.value.num
            },
            tooltip: {
                trigger: 'axis'
            },
            yAxis: {
                type: 'value'
            },
            grid: {
                left: '10%',
                right: '4%',
                bottom: '6%',
                containLabel: false
            },
            series: [
                {
                    name: '内存使用率',
                    data: data2.value.value,
                    type: 'line',
                    
                }
            ]
        });
        window.onresize = function() {
            //自适应大小
            chart.resize();
        };
    }
    const initEcharts3 = () => {
        let echartsInfoOne = document.getElementById("echartsInfoOne3")
        echartsInfoOne.removeAttribute("_echarts_instance_");
        let chart = echart.init(echartsInfoOne);
        chart.setOption({
            title: {
                text: '磁盘使用率',
                left: 'center'
            },
            xAxis: {
                type: 'category',
                boundaryGap: false,
                data: data3.value.num
            },
            tooltip: {
                trigger: 'axis'
            },
            yAxis: {
                type: 'value'
            },
            grid: {
                left: '10%',
                right: '4%',
                bottom: '6%',
                containLabel: false
            },
            series: [
                {
                    name: '磁盘使用率',
                    data: data3.value.value,
                    type: 'line',
                    
                }
            ]
        });
        window.onresize = function() {
            //自适应大小
            chart.resize();
        };
    }

    onUnmounted (()=>{
        clearTimeout(channel1.value)
        clearTimeout(channel2.value)
        clearTimeout(channel3.value)
    });

</script>


<style lang="scss" scoped>
.flex {
    margin-left: 40px;
}
.aaaa{
    width: 100%;
    display: flex;
    flex-wrap: wrap;
}

</style>
