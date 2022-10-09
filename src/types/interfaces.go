package types

import "github.com/speauty/go.ding/src/consts"

type DrMsg interface {
	New() any
	SetTitle(title string)
	SetContent(content string)
	GetMsgType() consts.MsgType
}
