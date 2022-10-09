package types

import "github.com/speauty/go.ding/src/consts"

// NewDRMsgEmpty 生成新的空消息
func NewDRMsgEmpty() *DRMsgEmpty {
	return (&DRMsgEmpty{}).New().(*DRMsgEmpty)
}

// DRMsgEmpty 钉钉机器人空消息
type DRMsgEmpty struct {
	MsgType string `json:"msgtype"` // fixed:empty
}

func (drMsg *DRMsgEmpty) GetMsgType() consts.MsgType { return consts.MsgEmpty }

func (drMsg *DRMsgEmpty) New() any { return &DRMsgText{MsgType: consts.StrMsgEmpty} }

func (drMsg *DRMsgEmpty) SetTitle(title string) {}

func (drMsg *DRMsgEmpty) SetContent(content string) {}
