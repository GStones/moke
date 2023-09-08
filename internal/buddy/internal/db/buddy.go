package db

import (
	"regexp"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"github.com/gstones/moke-kit/orm/nosql/key"

	"moke/internal/buddy/internal/db/model"
)

const (
	validatePathPattern = `[a-zA-Z0-9_.-]`
)

var (
	validatePathExp *regexp.Regexp
)

func init() {
	validatePathExp = regexp.MustCompile(validatePathPattern)
}

type BuddyQueue struct {
	nosql.DocumentBase `bson:"-"`
	Data               *model.BuddyQueue `bson:"data"`
}

func (b *BuddyQueue) init(appId, id string, ros diface.ICollection) error {
	if ros == nil {
		return nerrors.ErrDocumentStoreIsNil
	}
	k, e := newBuddyQueueKey(appId, id)
	if e != nil {
		return e
	}
	b.Data = model.NewBuddyQueue(id)
	b.Init(&b.Data, b.clear, ros, k)
	return nil
}

func (b *BuddyQueue) clear() {
	b.Data.Clear()
}

func (b *BuddyQueue) InitDefault() error {
	return nil
}

func NewRelativeBuddyQueuePath(appId string, id string) (string, error) {
	if !validatePathExp.MatchString(id) {
		return "", key.ErrInvalidKeyFormat
	} else if !validatePathExp.MatchString(appId) {
		return "", key.ErrInvalidKeyFormat
	} else {
		return "/" + appId + "/buddies/" + id, nil
	}
}

func newBuddyQueueKey(appId string, id string) (key.Key, error) {
	if result, err := NewRelativeBuddyQueuePath(appId, id); err != nil {
		return key.Key{}, err
	} else {
		return key.NewKeyFromString(result)
	}
}
