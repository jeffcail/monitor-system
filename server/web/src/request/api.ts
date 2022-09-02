import axios from "./axios";

// 登陆
export const login = (params: any) => {
    return axios.postJson(params, "/api/login")
}

// 左侧菜单列表
export const menu = () => {
    return axios.getJson("", "/api/menus/list")
}