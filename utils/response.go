package utils

var Res *Result

type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ResponseJson
func (r *Result) ResponseJson(success bool, code int, msg string, data interface{}) *Result {
	res := &Result{
		Success: success,
		Code:    code,
		Msg:     msg,
		Data:    data,
	}
	return res
}

var Resp *Page

type Page struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

// ResponsePagination
func (p *Page) ResponsePagination(count int64, list interface{}) *Page {
	pg := Page{}
	pg.Total = count
	pg.List = list
	return &pg
}
