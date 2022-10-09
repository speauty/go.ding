package types

import "github.com/speauty/go.ding/src/consts"

// NewDRMsgText 生成新的文本消息 需要带上关键字
func NewDRMsgText() *DRMsgText {
	return (&DRMsgText{}).New().(*DRMsgText)
}

// DRMsgText 钉钉机器人文本消息
type DRMsgText struct {
	MsgType string `json:"msgtype"` // fixed:text
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At *At `json:"at"`
}

func (drMsg *DRMsgText) GetMsgType() consts.MsgType { return consts.MsgText }

func (drMsg *DRMsgText) New() any { return &DRMsgText{MsgType: consts.StrMsgText, At: &At{}} }

func (drMsg *DRMsgText) SetTitle(title string) {}

func (drMsg *DRMsgText) SetContent(content string) { drMsg.Text.Content = content }

func (drMsg *DRMsgText) SetAt(at *At) { drMsg.At = at }
