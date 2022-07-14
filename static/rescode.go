package static



const (
	SUCCESS 	string 	= "success"	//成功
	FAILED		string 	= "failed"	//失败
	PARAMLOST	string	= "param_lost"	//参数缺失
	PARAMERR	string	= "param_error"	//参数错误
	SERVERERR	string	= "server_error"	//服务器异常
	DBERROR		string	= "db_error"	//数据异常
	
	TOKENERROR	string   = "token_error"	//登录TOKEN错误
	DATAEXIST	string 	= "data_exist"	//数据已存在
	DATANOTEXIST string 	= "data_not_exist"	//数据不存在
	NOTOWNER string = "not_owner"	//不是拥有者
	FORMATERROR string = "format_error"	//格式错误
	
	TIMEOUT string = "time_out" //超时
	NO_PERMISSION string = "no_permission" //没有权限
	STATUSERROR string =  "status_error" //状态异常
	DUPLICATE_DATA string = "duplicate_data" //重复数据
	PWDERROR string = "pwd_error" //密码错误
	NEED_LOGIN string = "need_login" //需要登录

	NEEDSHOW string = "need_show" //需要显示返回

	MOREACTION string 	= "more_action"	//跟多操作
	NEEDAUTH string = "need_auth"	//需要鉴权
	AUTH_FAILED string = "auth_failed" //认证失败

	NOT_ENOUGH string = "not_enough"//数据不足， 积分不足
	


	DEFAULT_SUCCESS_DESC string = "成功"
	DATA_NOT_EXIST_DESC string = "数据不存在"
	DATA_ALREADY_EXIST_DESC string = "数据已存在"
	
	NOTOWNER_DESC string = "不是拥有者"
	FORMATERROR_DESC string = "格式错误"
	
	PARAMERR_DESC string = "参数异常"
	DBERROR_DESC string = "数据库异常"
	PARAMNULL_DESC string = "参数为空"
	DEFAULT_FAILED_DESC string = "失败"
	
	NO_PERMISSION_DESC string = "没有权限"
	NEED_LOGIN_DESC string = "需要登录"
	STATUSERROR_DESC string = "状态异常"
	DUPLICATE_DATA_DESC string = "重复数据"
	PWDERROR_DESC string = "密码错误"
)