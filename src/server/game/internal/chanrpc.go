package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var (
	AgentMap map[string]gate.Agent
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	AgentMap = make(map[string]gate.Agent)
}

func rpcNewAgent(args []interface{}) {
	log.Debug("NewAgent.....")
	a := args[0].(gate.Agent)
	_ = a

}

func rpcCloseAgent(args []interface{}) {
	log.Debug("CloseAgent.....")
	a := args[0].(gate.Agent)
	_ = a
	userdata := a.UserData()
	if userdata != nil {
		name := userdata.(string)
		delete(AgentMap, name)
		for _, v := range manager.roomList {
			log.Debug("%s left room%d", name, v.id)
			manager.Leave(v.id, name)
		}
	}

}
