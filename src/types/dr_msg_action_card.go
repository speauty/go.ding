package types

import "github.com/speauty/go.ding/src/consts"

// NewDRMsgActionCard 生成新的卡片消息(整体跳转)
func NewDRMsgActionCard() *DRMsgActionCard {
	return (&DRMsgActionCard{}).New().(*DRMsgActionCard)
}

// DRMsgActionCard 钉钉机器人-跳转卡片消息
type DRMsgActionCard struct {
	MsgType    string `json:"msgtype"` // fixed: actionCard
	ActionCard struct {
		Title       string `json:"title"`
		Text        string `json:"text"`
		SingleTitle string `json:"singleTitle"`
		SingleURL   string `json:"singleURL"`
	} `json:"actionCard"`
}

func (drMsg *DRMsgActionCard) GetMsgType() consts.MsgType { return consts.MsgActionCard }

func (drMsg *DRMsgActionCard) New() any { return &DRMsgActionCard{MsgType: consts.StrMsgActionCard} }

func (drMsg *DRMsgActionCard) SetTitle(title string) { drMsg.ActionCard.Title = title }

func (drMsg *DRMsgActionCard) SetContent(content string) { drMsg.ActionCard.Text = content }

func (drMsg *DRMsgActionCard) SetSingleTitle(title string) { drMsg.ActionCard.SingleTitle = title }
func (drMsg *DRMsgActionCard) SetSingleURL(url string)     { drMsg.ActionCard.SingleURL = url }
