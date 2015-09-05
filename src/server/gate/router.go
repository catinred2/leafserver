package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	msg.JSONProcessor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.Login{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.Join{}, game.ChanRPC)
	msg.JSONProcessor.SetRouter(&msg.Activity{}, game.ChanRPC)
}
