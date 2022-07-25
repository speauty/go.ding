package go_ding

type At struct {
	AtMobiles []string `json:"atMobiles"` // optional
	AtUserIds []string `json:"atUserIds"` // optional
	IsAtAll   bool     `json:"isAtAll"`   // optional
}
type DRMsgText struct {
	MsgType string `json:"msgtype"` // fixed:text
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	At At `json:"at"`
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

type DRMsgActionCard struct {
	MsgType    string `json:"msgtype"` // fixed: actionCard
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation uint8  `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
	} `json:"actionCard"`
}

type DRErr struct {
	Code int32  `json:"errcode"`
	Msg  string `json:"errmsg"`
}
