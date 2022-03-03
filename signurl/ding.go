package ding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

type Webhook struct {
	AccessToken string
	Secret      string
}

func (t *Webhook) hmacSha256(strToSing string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(strToSing))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (t *Webhook) GetURL() string {
	wh := "https://oapi.dingtalk.com/robot/send?access_token=" + t.AccessToken
	timestamp := time.Now().UnixNano() / 1e6
	strToSing := fmt.Sprintf("%d\n%s", timestamp, t.Secret)
	sign := t.hmacSha256(strToSing, t.Secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", wh, timestamp, sign)
	return url
}
