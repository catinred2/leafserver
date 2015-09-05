package internal

import (
	"errors"
	"fmt"
	"server/msg"
	. "server/util"
	"sync"
	"time"

	"github.com/name5566/leaf/log"
)

type Room struct {
	id           int
	name         string
	max          int
	players      *Set
	players_data map[string]msg.GameInfo
}
type RoomManager struct {
	roomList []*Room
	sync.RWMutex
}

var (
	manager   *RoomManager
	timer     *time.Timer
	closeChan chan (int)
)

func init() {
	var room1 Room
	room1.id = 1
	room1.name = "room1"
	room1.max = 2
	room1.players = NewSet()
	room1.players_data = make(map[string]msg.GameInfo)
	manager = new(RoomManager)
	manager.roomList = make([]*Room, 1)
	manager.roomList[0] = &room1
	closeChan = make(chan (int))
	go manager.update(&room1)

}
func (p *RoomManager) GetRooms() []*Room {
	return p.roomList
}
func (p *RoomManager) whereAmI(name string) (*Room, error) {

	p.RLock()
	defer p.RUnlock()
	for _, v := range p.roomList {
		if v.players.Has(name) {
			return v, nil
		}
	}
	return nil, errors.New("you are not in any room")
}
func (p *RoomManager) Join(id int, name string) (err error) {
	room, e := p.whereAmI(name)
	if e == nil {
		err = errors.New(fmt.Sprintf("already in room:%d", room.id))
		return
	}
	p.Lock()
	defer p.Unlock()
	for _, v := range p.roomList {
		log.Debug("checking room %d", v.id)
		if v.id == id {

			if v.players.Len() < v.max {
				//ok. player can join
				v.players.Add(name)
				return nil
			} else {
				err = errors.New("room full")
				return
			}
		}
	}
	err = errors.New("no such room")
	return
}
func (p *RoomManager) Leave(id int, name string) error {
	room, err := p.whereAmI(name)

	if err != nil {
		return errors.New("you are not in any room")
	}
	if room.id != id {
		return errors.New("you are not in this room")
	}
	p.Lock()
	defer p.Unlock()
	room.players.Remove(name)
	log.Error("player %s removed", name)
	return nil
}
func (p *RoomManager) updateroom(r *Room) {
	players := r.players.List()
	for _, v := range players {

		gameinfo := r.players_data[v]
		for _, v2 := range players {
			if v != v2 {
				agent, ok := AgentMap[v2]
				if ok {
					log.Debug("sending user:%s info to %s content=%+v", v, v2, gameinfo)
					var act msg.Activity
					act.Response.I_code = 0
					act.Response.I_green_pos_x = gameinfo.I_green_pos_x
					act.Response.I_green_pos_y = gameinfo.I_green_pos_y
					act.Response.I_grey_pos_x = gameinfo.I_grey_pos_x
					act.Response.I_grey_pos_y = gameinfo.I_grey_pos_y

					agent.WriteMsg(&act)
				}
			}
		}

	}
}
func (p *RoomManager) update(r *Room) {
	for {
		select {
		case <-time.After(1000 * time.Millisecond):
			p.updateroom(r)
		case <-closeChan:
			//return here
			return
		}
	}

}
