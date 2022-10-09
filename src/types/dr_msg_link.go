package types

import "github.com/speauty/go.ding/src/consts"

// NewDRMsgLink 生成新的链接消息 需要带上关键字
func NewDRMsgLink() *DRMsgLink {
	return (&DRMsgLink{}).New().(*DRMsgLink)
}

// DRMsgLink 钉钉机器人-链接消息
type DRMsgLink struct {
	MsgType string `json:"msgtype"` // fixed:link
	Link    struct {
		Title      string `json:"title"`
		Content    string `json:"text"`
		MessageUrl string `json:"messageUrl"`
		PicUrl     string `json:"picUrl"` // optional
	} `json:"link"`
}

func (drMsg *DRMsgLink) GetMsgType() consts.MsgType { return consts.MsgLink }

func (drMsg *DRMsgLink) New() any { return &DRMsgLink{MsgType: consts.StrMsgLink} }

func (drMsg *DRMsgLink) SetTitle(title string) { drMsg.Link.Title = title }

func (drMsg *DRMsgLink) SetContent(content string) { drMsg.Link.Content = content }

func (drMsg *DRMsgLink) SetMessageUrl(messageUrl string) { drMsg.Link.MessageUrl = messageUrl }

func (drMsg *DRMsgLink) SetPicUrl(picUrl string) { drMsg.Link.PicUrl = picUrl }
