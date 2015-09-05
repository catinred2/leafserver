package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"reflect"
	"server/msg"
	"time"
)

var (
	c      = make(chan (int))
	myname = ""
)

func main() {
	var client int
	cid := flag.Int("id", 1, "client id")
	flag.Parse()
	client = *cid

	fmt.Printf("client = %d\n", client)
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	go recv(conn)
	var r1 msg.Login
	//	r1.Request = new(struct {
	//		S_name string
	//		S_pwd  string
	//	})
	r1.Request.S_name = "name1"
	r1.Request.S_pwd = "pass1"

	data, err := pack(&r1)
	if err != nil {
		fmt.Printf("pack err=%s", err.Error())
		return
	}
	dataLen := data[0]<<8 + data[1]
	dataMsg := data[2:]
	fmt.Printf("####send data\n%s\n##len=%d\n", string(dataMsg), dataLen)

	// 发送消息
	conn.Write(data)
	fmt.Println("------")
	time.Sleep(time.Duration(500 * time.Millisecond))

	var r2 msg.Join

	r2.Request.I_no = 1
	r2.Request.S_name = "name2"
	data, err = pack(&r2)
	conn.Write(data)
	time.Sleep(time.Duration(500 * time.Millisecond))

	greenx := 0
	greeny := 0
	greyx := 0
	greyy := 0
	var r3 msg.Activity
	for {

		r3.Request.I_green_pos_x = fmt.Sprintf("%d", greenx)
		r3.Request.I_green_pos_y = fmt.Sprintf("%d", greeny)
		r3.Request.I_grey_pos_x = fmt.Sprintf("%d", greyx)
		r3.Request.I_grey_pos_y = fmt.Sprintf("%d", greyy)
		data, err = pack(&r3)
		if err != nil {
			fmt.Printf("pack err %s\n", err.Error())
			break
		}
		fmt.Println("updating game info")
		_, err := conn.Write(data)
		if err != nil {
			fmt.Printf("conn.write err:%s", err.Error())
			break
		}
		time.Sleep(time.Duration(2) * time.Second)
		if client == 1 {
			fmt.Println("move green")
			greenx += 1
			greeny += 1
		} else {
			fmt.Println("move grey")
			greyx += 1
			greyy += 1
		}

	}
	<-c

}
func recv(conn net.Conn) {
	for {

		s, err := unpack(conn)
		if err != nil {
			fmt.Printf("read json error:%s\n", err.Error())
			c <- 1
			return
		}
		fmt.Printf("get json:\n%s\n", s)
	}
}
func pack(msg interface{}) ([]byte, error) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return nil, errors.New("json message pointer required")
	}
	msgID := msgType.Elem().Name()

	m := map[string]interface{}{msgID: msg}
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("marshal result:%s len=%d\n", string(data), len(data))
	buf := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(buf, uint16(len(data)))
	copy(buf[2:], data)
	//fmt.Printf("buf len=%d\n", len(buf))
	return buf, nil
}
func unpack(conn net.Conn) (string, error) {
	buf := make([]byte, 2)

	_, err := io.ReadFull(conn, buf)
	if err != nil {
		return "", err
	}
	packetlen := buf[0]<<8 + buf[1]
	//fmt.Printf("packet len=%d\n", packetlen)

	jsonbuff := make([]byte, packetlen)
	_, err = io.ReadFull(conn, jsonbuff)
	return string(jsonbuff), err
}
