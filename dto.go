package go_ding

//text msgtype content atMobiles atUserIds isAtAll
//link msgtype title text messageUrl picUrl
//markdown msgtype title text atMobiles atUserIds isAtAll
//整体跳转ActionCard msgtype title text singleTitle singleURL btnOrientation
//独立跳转ActionCard msgtype title btns actionURL btnOrientation
//FeedCard类型 msgtype title messageURL picURL

type DRMsgText struct {
	MsgType   string   `json:"msgtype"` // fixed:text
	Content   string   `json:"content"`
	AtMobiles []string `json:"atMobiles"` // optional
	AtUserIds []string `json:"atUserIds"` // optional
	IsAtAll   bool     `json:"isAtAll"`   // optional
}

type DRMsgLink struct {
	MsgType    string `json:"msgtype"` // fixed:link
	Title      string `json:"title"`
	Content    string `json:"content"`
	MessageUrl string `json:"messageUrl"`
	PicUrl     string `json:"picUrl"` // optional
}

type DRMsgMarkdown struct {
	MsgType   string   `json:"msgtype"` // fixed:markdown
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	AtMobiles []string `json:"atMobiles"` // optional
	AtUserIds []string `json:"atUserIds"` // optional
	IsAtAll   bool     `json:"isAtAll"`   // optional
}
