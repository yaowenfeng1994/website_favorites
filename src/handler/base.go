package handler

type BaseResponse struct {
	ErrCode  int                     `json:"err_code"`
	ErrMsg   string                  `json:"err_msg"`
	ErrMsgEn string                  `json:"err_msg_en"`
	Data     map[string]interface{}  `json:"data"`
}

var (
	ErrMapping map[int][2]string
	//Pool *libs.SQLConnPool
)

func init() {
	ErrMapping = make(map[int][2]string)
	ErrMapping[0x0000] = [2]string{"Request success", "请求成功"}
	ErrMapping[0x0001] = [2]string{"Unknown error", "未知错误"}
	ErrMapping[0x0002] = [2]string{"Mismatch parameter", "参数不匹配"}
	ErrMapping[0x0003] = [2]string{"Create account fail", "创建账号失败"}
	//Pool = libs.InitMySQLPool("127.0.0.1", "website_favorites", "root", "123456", "utf8", 200, 100)
}

func (b *BaseResponse) InitBaseResponse(errCode int, d map[string]interface{}) {
	b.ErrCode = errCode
	_, ok := ErrMapping[errCode]
	if ok == true {
		b.ErrMsgEn = ErrMapping[errCode][0]
		b.ErrMsg = ErrMapping[errCode][1]
	} else {
		b.ErrMsgEn = "Unknown error"
		b.ErrMsg = "未知错误"
	}
	b.Data = d
}
