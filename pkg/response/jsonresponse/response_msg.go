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
	msg[200600] = Message{ Code: 200600, Msg: "成功！" }
	msg[200601] = Message{ Code: 200601, Msg: "登录成功！" }
	//msg[200602] = Message{ Code: 200602, Msg: "成功！" }
	//msg[200603] = Message{ Code: 200603, Msg: "成功！" }
	//msg[200604] = Message{ Code: 200604, Msg: "成功！" }
	//msg[200605] = Message{ Code: 200605, Msg: "成功！" }
	//msg[200606] = Message{ Code: 200606, Msg: "成功！" }
	//msg[200607] = Message{ Code: 200607, Msg: "成功！" }
	//msg[200608] = Message{ Code: 200608, Msg: "成功！" }
	//msg[200609] = Message{ Code: 200609, Msg: "成功！" }
	//msg[200610] = Message{ Code: 200610, Msg: "成功！" }

	msg[200700] = Message{ Code: 200700, Msg: "失败！" }
	msg[200701] = Message{ Code: 200701, Msg: "登录失败！" }
	msg[200702] = Message{ Code: 200702, Msg: "用户不存在！" }
	msg[200703] = Message{ Code: 200703, Msg: "验证码错误！" }
	msg[200704] = Message{ Code: 200704, Msg: "密码错误！" }
	msg[200705] = Message{ Code: 200705, Msg: "缺少参数！" }
	msg[200706] = Message{ Code: 200706, Msg: "获取数据失败！" }
	msg[200707] = Message{ Code: 200707, Msg: "没有权限！" }
	msg[200708] = Message{ Code: 200708, Msg: "没有管理员权限！" }
	msg[200709] = Message{ Code: 200709, Msg: "创建失败！" }


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