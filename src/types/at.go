package types

// At 消息安特对象(可通过手机号, 或用户id, 或者全体)
type At struct {
	AtMobiles []string `json:"atMobiles"` // optional
	AtUserIds []string `json:"atUserIds"` // optional
	IsAtAll   bool     `json:"isAtAll"`   // optional
}
