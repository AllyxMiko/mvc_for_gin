package libs

// 封装状态码（非http状态码）这个状态码用于消息提示，与项目无关，仅作示例

var S = map[int]string{
	0:     "请求成功",
	30001: "短链接已存在，请勿重复创建！",
	40001: "请求参数不合法！",
	40002: "未查询到此短链接！",
	50001: "短链接生成失败！",
}

// 3开头为提示类信息
// 4开头的为客户端错误
// 5开头的表示服务端错误
const (
	OK = 0
	// UrlAlreadyExisted 短链接已存在
	UrlAlreadyExisted = 30001
	// RequestParamIllegal 请求参数不合法
	RequestParamIllegal = 40001
	// ShortUrlNotFound 未查询到此短链接
	ShortUrlNotFound = 40002
	// CreateUrlFaild 数据添加失败
	CreateUrlFaild = 50001
)
