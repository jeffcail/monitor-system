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
