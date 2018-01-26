package idl

type DemoReq struct {
	Name 	string `json:"name"`
	Age 	int	`json:"age"`
}

func NewDemoReq() *DemoReq {
	return &DemoReq{}
}

type DemoResp struct {
	Errno 	int	`json:"errno"`
	Msg 	string	`json:"msg"`
	Result 	string	`json:"result"`
}

func NewDemoResp() *DemoResp {
	return &DemoResp{}
}
