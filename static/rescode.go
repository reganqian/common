package static

type BaseReply struct {
	ResCode string `json:"resCode"` //返回码
	ResDesc string `json:"resDesc"` //返回描述
}

func (s *BaseReply) Success() {
	s.ResCode = SUCCESS
	s.ResDesc = DEFAULT_SUCCESS_DESC
}

func (s *BaseReply) Failed(resCode, resDesc string) {
	s.ResCode = resCode
	s.ResDesc = resDesc
}

type StringReply struct {
	BaseReply
	Data string `json:"data"`
}

type IntReply struct {
	BaseReply
	Data int32 `json:"data"`
}

type Int64Reply struct {
	BaseReply
	Data int64 `json:"data"`
}

type Uint32Reply struct {
	BaseReply
	Data uint32 `json:"data"`
}

type UintReply struct {
	BaseReply
	Data uint `json:"data"`
}

type StringsReply struct {
	BaseReply
	DataList []string `json:"dataList"`
}

const (
	SUCCESS   string = "SUCCESS"      //成功
	FAILED    string = "FAILED"       //失败
	PARAMLOST string = "PARAM_LOST"   //参数缺失
	PARAMERR  string = "PARAM_ERROR"  //参数错误
	SERVERERR string = "SERVER_ERROR" //服务器异常
	DBERROR   string = "DB_ERROR"     //数据异常
	STATUSERR string = "STATUS_ERR"

	TOKENERROR   string = "TOKEN_ERROR"    //登录TOKEN错误
	DATAEXIST    string = "DATA_EXIST"     //数据已存在
	DATANOTEXIST string = "DATA_NOT_EXIST" //数据不存在
	NOTOWNER     string = "NOT_OWNER"      //不是拥有者
	FORMATERROR  string = "FORMAT_ERROR"   //格式错误

	TIMEOUT        string = "TIME_OUT"       //超时
	NO_PERMISSION  string = "NO_PERMISSION"  //没有权限
	STATUSERROR    string = "STATUS_ERROR"   //状态异常
	DUPLICATE_DATA string = "DUPLICATE_DATA" //重复数据
	PWDERROR       string = "PWD_ERROR"      //密码错误
	NEED_LOGIN     string = "NEED_LOGIN"     //需要登录

	NEEDSHOW string = "NEED_SHOW" //需要显示返回

	MOREACTION  string = "MORE_ACTION" //跟多操作
	NEEDAUTH    string = "NEED_AUTH"   //需要鉴权
	AUTH_FAILED string = "AUTH_FAILED" //认证失败

	NOT_ENOUGH string = "NOT_ENOUGH" //数据不足， 积分不足

	DEFAULT_SUCCESS_DESC    string = "success"
	DATA_NOT_EXIST_DESC     string = "data is not exist"
	DATA_ALREADY_EXIST_DESC string = "data is already exist"

	NOTOWNER_DESC    string = "not the owner"
	FORMATERROR_DESC string = "format error"

	PARAMERR_DESC       string = "param error"
	DBERROR_DESC        string = "data base error"
	PARAMNULL_DESC      string = "param is null"
	DEFAULT_FAILED_DESC string = "failed"

	NO_PERMISSION_DESC  string = "permission deny"
	NEED_LOGIN_DESC     string = "need login"
	STATUSERROR_DESC    string = "status error"
	DUPLICATE_DATA_DESC string = "duplicate data"
	PWDERROR_DESC       string = "password is error"
)
