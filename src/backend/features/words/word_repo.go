package words

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	WordRepositoryRead interface {
		IsWordExistsInGroup(ctx context.Context, groupId, userId bson.ObjectID, word string) (bool, error)
		GetWordsByGroupID(ctx context.Context, groupID bson.ObjectID, page, limit int) ([]Word, int64, error)
	}

	WordRepostioryWrite interface {
		CreateWord(ctx context.Context, word Word) (bson.ObjectID, error)
	}
)

var (
	wordRepositoryRead  WordRepositoryRead
	wordRepositoryWrite WordRepostioryWrite
)

func InitWordRepositoryRead(read WordRepositoryRead) {
	wordRepositoryRead = read
}

func NewWordRepositoryRead() WordRepositoryRead {
	return wordRepositoryRead
}

func InitWordRepositoryWrite(write WordRepostioryWrite) {
	wordRepositoryWrite = write
}

func NewWordRepositoryWrite() WordRepostioryWrite {
	return wordRepositoryWrite
}
