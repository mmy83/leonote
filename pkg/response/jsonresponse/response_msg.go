/*
@Time : 2020/4/25 11:07 上午
@Author : mmy83
@File : response_msg.go
@Software: GoLand
*/

package jsonresponse

type Message struct {
	Code int
	Msg string
}

var msg = make(map[int]Message)

func init(){
	msg[200000] = Message{ Code: 200000, Msg: "成功！" }

	msg[200500] = Message{ Code: 200500, Msg: "失败！" }
	msg[200501] = Message{ Code: 200501, Msg: "登录失败！" }
	msg[200502] = Message{ Code: 200502, Msg: "用户不存在！" }
	msg[200503] = Message{ Code: 200503, Msg: "验证码错误！" }
	msg[200504] = Message{ Code: 200504, Msg: "密码错误！" }
	msg[200505] = Message{ Code: 200505, Msg: "缺少参数！" }
	msg[200506] = Message{ Code: 200506, Msg: "获取数据失败！" }
	msg[200507] = Message{ Code: 200507, Msg: "没有权限！" }
	msg[200508] = Message{ Code: 200508, Msg: "没有管理员权限！" }
	msg[200509] = Message{ Code: 200509, Msg: "创建失败！" }

}

func getMsg(code int) Message {
	if msg, ok := msg[code]; ok {
		return msg
	}
	return Message{
		Code: 200999,
		Msg: "未知返回！",
	}
}