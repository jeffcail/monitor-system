package params

// 新管理员创建入参
type AdminParam struct {
	Username   string `json:"username"  validate:"required"`              //用户昵称
	RealName   string `json:"real_name"  validate:"required"`             //真实姓名
	Password   string `json:"password"  validate:"required,gte=6,lte=12"` //密码
	Phone      string `json:"phone"  validate:"required"`                 //手机号
	Department string `json:"department" `                                //部门
	OfficePost string `json:"office_post" `                               //职位
}

// 管理员查询入参
type SelAdminParam struct {
	Username   string `json:"username" `    //用户昵称
	RealName   string `json:"real_name" `   //真实姓名
	Phone      string `json:"phone" `       //手机号
	Department string `json:"department" `  //部门
	OfficePost string `json:"office_post" ` //职位
	Page       int    `json:"page"  validate:"required"`
	PageSize   int    `json:"page_size"  validate:"required"`
}

// 管理员信息变更入参
type UpdAdminParamById struct {
	Id         int64  `json:"id" validate:"required"`
	Username   string `json:"username" `  //用户昵称
	RealName   string `json:"real_name" ` //真实姓名
	Password   string `json:"password" `
	Phone      string `json:"phone" `       //手机号
	RoleId     string `json:"role_id" `     //角色Id
	Department string `json:"department"`   //部门
	OfficePost string `json:"office_post" ` //职位
	State      int    `json:"state"  `      //账号状态 1: 正常  2:禁用
}

// DeleteParam
type DeleteParam struct {
	Id int64 `json:"id" validate:"required"`
}

// EnableDisableAdminParam
type EnableDisableAdminParam struct {
	Id    int64 `json:"id" validate:"required"`
	State int   `json:"state" validate:"required"`
}
