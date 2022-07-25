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

// {
//    "actionCard": {
//        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
//        "text": "![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)
// ### 乔布斯 20 年前想打造的苹果咖啡厅
// Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
//        "btnOrientation": "0",
//        "singleTitle" : "阅读全文",
//        "singleURL" : "https://www.dingtalk.com/"
//    },
//    "msgtype": "actionCard"
//}

type DRMsgActionCard struct {
	MsgType    string `json:"msgtype"` // fixed: actionCard
	ActionCard struct {
		Title          string `json:"title"`
		Text           string `json:"text"`
		BtnOrientation uint8  `json:"btnOrientation"`
		SingleTitle    string `json:"singleTitle"`
	} `json:"actionCard"`
}
