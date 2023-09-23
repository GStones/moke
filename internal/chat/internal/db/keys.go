package db

import "github.com/gstones/moke-kit/orm/nosql/key"

func makeChatMsgKey(uid string) (key.Key, error) {
	return key.NewKeyFromParts("chat", "message", "queue", uid)
}
