import axios from "./axios";

// 登陆
export const login = (params: any) => {
    return axios.postJson(params, "/api/login")
}

// 左侧菜单列表
export const menu = () => {
    return axios.getJson("", "/api/menus/list")
}

// 服务检测
// 服务检测列表
export const serveList = (params: any) => {
    return axios.postJson(params, "/api/serve/list")
}

// 创建服务
export const createServe = (params: any) => {
    return axios.postJson(params, "/api/serve/create")
}

// 删除服务
export const deleteServe = (params: any) => {
    return axios.postJson(params, "/api/serve/delete")
}

// 服务升级
export const upgradeServe = (params: any) => {
    return axios.postJson(params, "/api/serve/upgrade")
}

// 机器
// 机器列表
export const machineList = (params: any) => {
    return axios.postJson(params, "/api/machine/list")
}

// 发送指令
export const sendMachineCommond = (params: any) => {
    return axios.postJson(params, "/api/machine/send/com")
}

// 所有机器
export const allMachine = () => {
    return axios.getJson("", "/api/machine/all")
}

// cpu使用率
export const clientSysCpu = (params: any) => {
    return axios.postJson(params, "/api/client/sys/cpu")
}

// 内存使用率
export const clientSysMen = (params: any) => {
    return axios.postJson(params, "/api/client/sys/men")
}

// 磁盘使用率
export const clientSysDisk = (params: any) => {
    return axios.postJson(params, "/api/client/sys/disk")
}

// 管理员
// 管理员列表
export const adminList = (params: any) => {
    return axios.postJson(params, "/api/admin/select")
}

// 添加管理员
export const addAdmin = (params: any) => {
    return axios.postJson(params, "/api/admin/register")
}

// 更新管理员
export const updateAdmin = (params: any) => {
    return axios.postJson(params, "/api/admin/update")
}

// 删除管理员
export const deleteAdmin = (params: any) => {
    return axios.postJson(params, "/api/admin/delete")
}

// 启用禁用管理员
export const adminEnableDisable = (params: any) => {
    return axios.postJson(params, "/api/admin/enable/disable")
}