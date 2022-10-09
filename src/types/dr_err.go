package types

// DRErr 钉钉机器人异常消息
type DRErr struct {
	Code int32  `json:"errcode"`
	Msg  string `json:"errmsg"`
}
