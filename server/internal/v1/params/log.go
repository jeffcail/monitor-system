package params

// log query params
type LogQueryParam struct {
	Username  string `json:"username" `  //用户昵称
	StateTime string `json:"state_time"` //开始时间(格式 yyyy-MM-dd hh:mm:ss)
	StopTime  string `json:"stop_time"`  //截至时间
	Page      int    `json:"page"  validate:"required"`
	PageSize  int    `json:"page_size"  validate:"required"`
}
