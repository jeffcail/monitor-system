<template>  
    <div class="go_out">
        <button @click="go_out">关闭终端</button>
    </div>

    <div class="term1">
        <div ref="terminalBox" style="height: 60vh;"></div>
    </div>
</template>

<script setup>
import 'xterm/css/xterm.css'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { ref, onMounted } from 'vue'
import { ElMessage } from "naive-ui";

let terminalBox = ref(null)
let term
let socket

onMounted(() => {
    //创建一个客户端
    term = new Terminal({
        rendererType: 'canvas', //使用这个能解决vim不显示或乱码
        cursorBlink: true,
        cursorStyle: "bar",
    })
    // term.write
    // 将客户端挂载到dom上
    const fitAddon = new FitAddon()
    term.loadAddon(fitAddon)
    term.open(terminalBox.value)
    fitAddon.fit()

    // 创建socket连接
    term.write('正在连接...\r\n');
    socket = new WebSocket('ws://127.0.0.1:9092/ssh')

    socket.binaryType = "arraybuffer";

    // 打开socket监听事件的方法
    socket.onopen = function () {
        fitAddon.fit()
        term.onData(function (data) {
            // socket.send(JSON.stringify({ type: "stdin", data: data }))
            socket.send(data)
            console.log(data)
        });
        // ElMessage.success("会话成功连接！")
    }
    socket.onclose = function () {
        term.writeln('连接关闭');
    }
    socket.onerror = function (err) {
        // console.log(err)
        term.writeln('读取数据异常：', err);
    }
    // 接收数据
    socket.onmessage = function (recv) {
        try {
            term.write(recv.data)
        } catch (e) {
            console.log('unsupport data', recv.data)
        }
    }

    window.addEventListener("resize", () => {
        fitAddon.fit()
    }, false)

})

const go_out = () => {
    window.history.go(-1);
}


</script>

<style lang="scss" scoped>
.upload {
    min-height: 100px;
}
.term1 {
    margin-left: 60px;
}
.go_out {
    margin-left: 40px;
    margin-top: 20px;
    margin-bottom: 20px;
}
</style>
