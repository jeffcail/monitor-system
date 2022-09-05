package params

type AdminParam struct {
	Username   string `json:"User_name" from:"User_name"`     //用户昵称
	RealName   string `json:"Real_name" from:"Real_name"`     //真实姓名
	Password   string `json:"Password" from:"Password"`       //密码
	Phone      string `json:"Phone" from:"Phone"`             //手机号
	Department string `json:"Department" from:"Department"`   //部门
	OfficePost string `json:"Office_post" from:"Office_post"` //职位
}

// 管理员查询入参
type SelAdminParam struct {
	Username   string `json:"User_name" from:"User_name"`     //用户昵称
	RealName   string `json:"Real_name" from:"Real_name"`     //真实姓名
	Phone      string `json:"Phone" from:"Phone"`             //手机号
	Department string `json:"Department" from:"Department"`   //部门
	OfficePost string `json:"Office_post" from:"Office_post"` //职位
	Page       int    `json:"page" from:"page" validate:"required"`
	PageSize   int    `json:"page_size" from:"page_size" validate:"required"`
}
