package internal

import (
	"fmt"
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.Hello{}, handleHello)
	handler(&msg.Login{}, handleLogin)
	handler(&msg.Join{}, handleJoin)
	handler(&msg.Activity{}, handleActivity)
	//	s := `
	//	{"Activity":{"request":{"grey_pos_x":"218.99989318848","grey_pos_y":"87","green_pos_y":"486.99819946289","name":"aa","green_pos_x":"68.997802734375"}}}
	//	`
	//	data := []byte(s)
	//	var m map[string]json.RawMessage
	//	err := json.Unmarshal(data, &m)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//	if len(m) != 1 {
	//		log.Debug("invalid json data")
	//	}
	//	log.Debug("raw=%+v", m)
	//	for msgID, data := range m {

	//		// msg
	//		mm := msg.Activity{}
	//		err = json.Unmarshal(data, &mm)
	//		log.Fatal("msgId =%s data=%s  %+v", msgID, string(data), mm)
	//	}

}
func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)
	a.SetUserData(m.Request.S_name)
	// 输出收到的消息的内容
	log.Debug("hello %+v", m)

	// 给发送者回应一个 Hello 消息
	var t msg.Hello
	t.Request.S_name = "abc"
	//t.Response.S_name = "def"
	a.WriteMsg(&t)
}
func handleLogin(args []interface{}) {
	log.Debug("handleLogin")
	m := args[0].(*msg.Login)
	a := args[1].(gate.Agent)
	id := GetUniqId()
	username := fmt.Sprintf("%s%d", "user", id)
	a.SetUserData(username)
	log.Debug("Set Agent Map %s", username)
	AgentMap[username] = a

	log.Debug("%s Logged in", m.Request.S_name)
	var r msg.Login
	r.Response.S_name = username
	r.Response.I_code = 0
	a.WriteMsg(&r)
}
func handleJoin(args []interface{}) {
	m := args[0].(*msg.Join)
	a := args[1].(gate.Agent)
	name := a.UserData().(string)
	log.Debug("%s wants to join room %d", name, m.Request.I_no)
	err := manager.Join(m.Request.I_no, name)

	var r msg.Join
	if err == nil {
		r.Response.I_code = 0
	} else {
		r.Response.I_code = 1
		r.Response.S_text = err.Error()
	}

	a.WriteMsg(&r)
}
func handleActivity(args []interface{}) {
	m := args[0].(*msg.Activity)
	a := args[1].(gate.Agent)
	name := a.UserData().(string)
	room, err := manager.whereAmI(name)
	log.Debug("Get Parameter:%+v", m)

	if err == nil {
		var gameinfo msg.GameInfo
		gameinfo.I_green_pos_x = m.Request.I_green_pos_x
		gameinfo.I_green_pos_y = m.Request.I_green_pos_y
		gameinfo.I_grey_pos_x = m.Request.I_grey_pos_x
		gameinfo.I_grey_pos_y = m.Request.I_grey_pos_y
		gameinfo.S_name = name
		room.players_data[name] = gameinfo
		log.Debug("updating gameinfo %s===%+v", name, gameinfo)
	} else {
		//var r msg.Activity
		//		r.Response.S_name = name
		//		r.Response.I_code = 1
		//		r.Response.S_text = err.Error()
		//		log.Debug(err.Error())
	}
}
