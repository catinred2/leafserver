package msg

import (
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
)

var (
	JSONProcessor     = json.NewProcessor()
	ProtobufProcessor = protobuf.NewProcessor()
)

type _Base struct {
	Request *struct {
	} `json:"request,omitempty"`
	Response *struct {
	} `json:"response,omitempty"`
}
type Hello struct {
	Request *struct {
		S_name string `json:"name,omitempty"`
	} `json:"request,omitempty"`
	Response *struct {
		S_name string `json:"name,omitempty"`
	} `json:"response,omitempty"`
}

type Test struct {
	Request *struct {
		S_name string `json:"name,omitempty"`
	} `json:"request,omitempty"`
	Response *struct {
		S_name string `json:"name,omitempty"`
		I_code int    `json:"code"`
	} `json:"response,omitempty"`
}

type Login struct {
	Request struct {
		S_name string `json:"name,omitempty"`
		S_pwd  string `json:"pwd,omitempty"`
	} `json:"request,omitempty"`
	Response struct {
		S_name string `json:"name,omitempty"`
		I_code int    `json:"code"`
	} `json:"response,omitempty"`
}
type Join struct {
	Request struct {
		S_name string `json:"name,omitempty"`
		I_no   int    `json:"no,omitempty"`
	} `json:"request,omitempty"`
	Response struct {
		S_text string `json:"text,omitempty"`
		I_code int    `json:"code"`
	} `json:"response,omitempty"`
}
type Activity struct {
	Request struct {
		S_name        string `json:"name,omitempty"`
		I_green_pos_x string `json:"green_pos_x,omitempty"`
		I_green_pos_y string `json:"green_pos_y,omitempty"`
		I_grey_pos_x  string `json:"grey_pos_x,omitempty"`
		I_grey_pos_y  string `json:"grey_pos_y,omitempty"`
	} `json:"request,omitempty"`
	Response struct {
		S_text        string `json:"text,omitempty"`
		I_code        int    `json:"code"`
		S_name        string `json:"name,omitempty"`
		I_green_pos_x string `json:"green_pos_x,omitempty"`
		I_green_pos_y string `json:"green_pos_y,omitempty"`
		I_grey_pos_x  string `json:"grey_pos_x,omitempty"`
		I_grey_pos_y  string `json:"grey_pos_y,omitempty"`
	} `json:"response,omitempty"`
}
type GameInfo struct {
	S_name        string `json:"name,omitempty"`
	I_green_pos_x string `json:"green_pos_x,omitempty"`
	I_green_pos_y string `json:"green_pos_y,omitempty"`
	I_grey_pos_x  string `json:"grey_pos_x,omitempty"`
	I_grey_pos_y  string `json:"grey_pos_y,omitempty"`
}

func init() {
	//helper.INFO("msg init")
	JSONProcessor.Register(&Hello{})
	JSONProcessor.Register(&Login{})
	JSONProcessor.Register(&Join{})
	JSONProcessor.Register(&Activity{})
	JSONProcessor.Register(&GameInfo{})
}
