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