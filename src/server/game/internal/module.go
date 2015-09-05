package internal

import (
	"server/base"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	IdChan   = make(chan (int))
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {

	m.Skeleton = skeleton
	//helper.("GameModule init")
	IdGenerator()
}

func (m *Module) OnDestroy() {
	closeChan <- 0
}
func IdGenerator() {

	go func() {
		id := 0
		for {
			IdChan <- id
			id++
		}

	}()
}
func GetUniqId() int {
	return <-IdChan
}
