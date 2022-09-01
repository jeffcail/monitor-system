<template>
    <div class="manageLogin-info" id="manageLogin-info">
        <Particles id="tsparticles"  :options= options />
        <div class="manageLogin-signin">
            <div class="manageLogin-signin-title">服务云监控系统登录</div>
            <el-form ref="formRef" class="login-form" :rules="rules">
                <el-form-item prop="username">
                    <i class="iconfont icon-denglu"></i>
                    <el-input placeholder="请输入您用户名"></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <i class="iconfont icon-mima"></i>
                    <el-input type="password" placeholder="请输入您密码"></el-input>
                </el-form-item>
                <el-form-item >
                    <el-button type="primary" >登录</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script setup>

</script>

<style lang="scss" scoped>
    .manageLogin-info{
        min-height: 100%;
        width: 100%;
        height: 100vh;
        position: relative;
        background-image: url("../../assets/images/bg.png");
        background-position: 50% 50%;
        background-size: cover;
        background-repeat: no-repeat;
        margin-left: auto;
        margin-right:auto;
        display: flex;
        justify-content: center;
        align-items: center;
        
    }
    .manageLogin-signin{
        background-color: #ffffff;
        position: relative;
        width: 350px;
        max-width: 100%;
        height: 520px;
        margin: 0 auto;
        overflow: hidden;
        border-radius: 5px;

        .manageLogin-signin-title{
            font-size: 24px;
            margin-top: 100px;
            padding-left: 0px;
            box-sizing: border-box;
            color: #333;
            margin-bottom: 50px;
        }

        .login-form{
            position: relative;
            max-width: 100%;
            display: flex;
            flex-wrap: wrap;
            justify-content: center;
            align-items: center;

            ::v-deep .el-input{
              --el-border-color:0px solid #ffffff !important;
            }
            ::v-deep .el-form-item {
                width: 70%;
                border-radius: 5px;
                color: #454545;
                margin-bottom: 0px;
            }
            ::v-deep .el-form-item__content{
                    flex-wrap: nowrap
            }
            ::v-deep .el-form-item__error{
                top: 60%;
            }
            ::v-deep .el-input__inner {
                display: inline-block;
                margin-bottom: 30px;
                box-shadow: 0 0 0 0px var(--el-input-border-color,var(--el-border-color-base)) inset;
                border-bottom: 1px solid #ccc;
                border-top: 1px solid #fff;
                border-left: 1px solid #fff;
                border-right: 1px solid #fff;
            }
            ::v-deep button {
                width: 100%;
                box-sizing: border-box;
            }
        }
    }
    .iconfont{
      margin-bottom: 30px;
      margin-right: 12px;
    }
</style>

