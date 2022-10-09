package types

import "github.com/speauty/go.ding/src/consts"

// NewDrMsgMarkdown 生成新的md消息 需要带上关键字
func NewDrMsgMarkdown() *DRMsgMarkdown {
	return (&DRMsgMarkdown{}).New().(*DRMsgMarkdown)
}

// DRMsgMarkdown 钉钉机器人-markdown消息
// @todo 安特方法待补充
type DRMsgMarkdown struct {
	MsgType  string `json:"msgtype"` // fixed:markdown
	Markdown struct {
		Title   string `json:"title"`
		Content string `json:"text"`
	} `json:"markdown"`
	At *At `json:"at"`
}

func (drMsg *DRMsgMarkdown) GetMsgType() consts.MsgType { return consts.MsgMarkdown }

func (drMsg *DRMsgMarkdown) New() any { return &DRMsgMarkdown{MsgType: consts.StrMsgMD} }

func (drMsg *DRMsgMarkdown) SetTitle(title string) { drMsg.Markdown.Title = title }

func (drMsg *DRMsgMarkdown) SetContent(content string) { drMsg.Markdown.Content = content }

func (drMsg *DRMsgMarkdown) SetAt(at *At) { drMsg.At = at }
