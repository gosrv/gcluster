package entity

import "time"

/**
在线玩家数据，只存在于内存中
*/
type PlayerOnlineData struct {
	loginTime time.Time
}

func NewPlayerOnlineData() *PlayerOnlineData {
	return &PlayerOnlineData{}
}

var PlayerIdKey = struct{}{}
