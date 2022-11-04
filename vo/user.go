package vo

type RequestLogin struct {
	UserName string `json:"user_name"  binding:"required"`
	PassWord string `json:"pass_word"  binding:"required"`
	Ua       string
}

type RequestRegister struct {
	UserName string `json:"user_name"  binding:"required"`
	PassWord string `json:"pass_word"  binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required,len=6"`
}

type RequestLoginLogQuery struct {
	Uid uint `form:"uid" binding:"required"`
}

type RequestUpdateUserInfo struct {
	Avatar       string `json:"avatar,omitempty" `
	Introduction string `json:"introduction,omitempty"`
	Grade        uint8  `json:"grade,omitempty" binding:"oneof=0 1"`
	Phone        string `json:"phone,omitempty" binding:"len=11"`
	// todo 这里接收前端数据，目前写的是接收string,不知正确与否。如果不写字符窜，而是time.Time 会导致前端传一个标准时间
	// todo 但是服务端转换失败
	Birthday string `json:"birthday,omitempty" time_format:"2006-01-02 15:04:05"`
}

type RequestUpdatePassWord struct {
	OldPassWord string `json:"old_pass_word"`
	NewPassWord string `json:"new_pass_word"`
}
