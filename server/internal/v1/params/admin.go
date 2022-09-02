package params

type AdminParam struct {
	Username   string `json:"Username" from:"Username"`     //用户昵称
	RealName   string `json:"RealName" from:"RealName"`     //真实姓名
	Password   string `json:"Password" from:"Password"`     //密码
	Phone      string `json:"Phone" from:"Phone"`           //手机号
	Department string `json:"Department" from:"Department"` //部门
	OfficePost string `json:"OfficePost" from:"OfficePost"` //职位
}
type SelAdminParam struct {
	Username   string `json:"Username" from:"Username"`     //用户昵称
	RealName   string `json:"RealName" from:"RealName"`     //真实姓名
	Phone      string `json:"Phone" from:"Phone"`           //手机号
	Department string `json:"Department" from:"Department"` //部门
	OfficePost string `json:"OfficePost" from:"OfficePost"` //职位
}
